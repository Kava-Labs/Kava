package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/kava-labs/kava/x/harvest/types"
)

// LiqData holds liquidation-related data
type LiqData struct {
	price            sdk.Dec
	ltv              sdk.Dec
	conversionFactor sdk.Int
}

// AttemptIndexLiquidations attempts to liquidate the lowest LTV borrows
func (k Keeper) AttemptIndexLiquidations(ctx sdk.Context) error {
	params := k.GetParams(ctx)
	borrowers := k.GetLtvIndexSlice(ctx, params.CheckLtvIndexCount)
	for _, borrower := range borrowers {
		prevLtv, err := k.GetCurrentLTV(ctx, borrower)
		if err != nil {
			panic(err)
		}

		k.SyncOutstandingInterest(ctx, borrower)

		liquidated, _ := k.AttemptKeeperLiquidation(ctx, sdk.AccAddress(types.LiquidatorAccount), borrower)
		if !liquidated {
			k.UpdateItemInLtvIndex(ctx, prevLtv, borrower)
		}
	}
	return nil
}

// AttemptKeeperLiquidation enables a keeper to liquidate an individual borrower's position
func (k Keeper) AttemptKeeperLiquidation(ctx sdk.Context, keeper sdk.AccAddress, borrower sdk.AccAddress) (bool, error) {
	prevLtv, err := k.GetCurrentLTV(ctx, borrower)
	if err != nil {
		return false, err
	}

	// Fetch deposits and parse coin denoms
	deposits := k.GetDepositsByUser(ctx, borrower)
	depositDenoms := []string{}
	for _, deposit := range deposits {
		depositDenoms = append(depositDenoms, deposit.Amount.Denom)
	}

	// Fetch borrow balances and parse coin denoms
	borrowBalances := k.GetBorrowBalance(ctx, borrower)
	borrowDenoms := getDenoms(borrowBalances)

	liqMap := make(map[string]LiqData)

	// Load required liquidation data for every deposit/borrow denom
	denoms := removeDuplicates(borrowDenoms, depositDenoms)
	for _, denom := range denoms {
		mm, found := k.GetMoneyMarket(ctx, denom)
		if !found {
			return false, sdkerrors.Wrapf(types.ErrMarketNotFound, "no market found for denom %s", denom)
		}

		priceData, err := k.pricefeedKeeper.GetCurrentPrice(ctx, mm.SpotMarketID)
		if err != nil {
			return false, err
		}

		liqMap[denom] = LiqData{priceData.Price, mm.BorrowLimit.LoanToValue, mm.ConversionFactor}
	}

	totalBorrowableUSDAmount := sdk.ZeroDec()
	totalDepositedUSDAmount := sdk.ZeroDec()
	for _, deposit := range deposits {
		lData := liqMap[deposit.Amount.Denom]
		usdValue := sdk.NewDecFromInt(deposit.Amount.Amount).Quo(sdk.NewDecFromInt(lData.conversionFactor)).Mul(lData.price)
		totalDepositedUSDAmount = totalDepositedUSDAmount.Add(usdValue)
		borrowableUSDAmountForDeposit := usdValue.Mul(lData.ltv)
		totalBorrowableUSDAmount = totalBorrowableUSDAmount.Add(borrowableUSDAmountForDeposit)
	}

	totalBorrowedUSDAmount := sdk.ZeroDec()
	for _, coin := range borrowBalances {
		lData := liqMap[coin.Denom]
		usdValue := sdk.NewDecFromInt(coin.Amount).Quo(sdk.NewDecFromInt(lData.conversionFactor)).Mul(lData.price)
		totalBorrowedUSDAmount = totalBorrowedUSDAmount.Add(usdValue)
	}

	// Validate that the proposed borrow's USD value is within user's borrowable limit
	if totalBorrowedUSDAmount.LTE(totalBorrowableUSDAmount) {
		return false, sdkerrors.Wrapf(types.ErrBorrowNotLiquidatable, "borrowed %s <= borrowable %s", totalBorrowedUSDAmount, totalBorrowableUSDAmount)
	}

	// Seize deposits and auciton them off
	err = k.SeizeDeposits(ctx, keeper, liqMap, deposits, borrowBalances, depositDenoms, borrowDenoms)
	if err != nil {
		return false, err
	}

	borrow, _ := k.GetBorrow(ctx, borrower)
	k.DeleteBorrow(ctx, borrow)
	k.RemoveFromLtvIndex(ctx, prevLtv, borrower)

	for _, oldDeposit := range deposits {
		k.DeleteDeposit(ctx, oldDeposit)
	}

	return true, err
}

// SeizeDeposits seizes a list of deposits and sends them to auction
func (k Keeper) SeizeDeposits(ctx sdk.Context, keeper sdk.AccAddress, liqMap map[string]LiqData,
	deposits []types.Deposit, borrowBalances sdk.Coins, dDenoms, bDenoms []string) error {

	// Seize % of every deposit and send to the keeper
	aucDeposits := sdk.Coins{}
	for _, deposit := range deposits {
		denom := deposit.Amount.Denom
		amount := deposit.Amount.Amount
		mm, _ := k.GetMoneyMarket(ctx, denom)

		// No rewards for anyone if liquidated by LTV index
		if !keeper.Equals(sdk.AccAddress(types.LiquidatorAccount)) {
			keeperReward := mm.KeeperRewardPercentage.MulInt(amount).TruncateInt()
			if keeperReward.GT(sdk.ZeroInt()) {
				// Send keeper their reward
				keeperCoin := sdk.NewCoin(denom, keeperReward)
				err := k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccountName, keeper, sdk.NewCoins(keeperCoin))
				if err != nil {
					return err
				}
				amount = amount.Sub(keeperReward)
			}
		}

		// Add remaining deposit coin to aucDeposits
		aucDeposits = aucDeposits.Add(sdk.NewCoin(denom, amount))
	}

	// Build valuation map to hold deposit coin USD valuations
	depositCoinValues := types.NewValuationMap()
	for _, deposit := range aucDeposits {
		dData := liqMap[deposit.Denom]
		dCoinUsdValue := sdk.NewDecFromInt(deposit.Amount).Quo(sdk.NewDecFromInt(dData.conversionFactor)).Mul(dData.price)
		depositCoinValues.Increment(deposit.Denom, dCoinUsdValue)
	}

	// Build valuation map to hold borrow coin USD valuations
	borrowCoinValues := types.NewValuationMap()
	for _, bCoin := range borrowBalances {
		bData := liqMap[bCoin.Denom]
		bCoinUsdValue := sdk.NewDecFromInt(bCoin.Amount).Quo(sdk.NewDecFromInt(bData.conversionFactor)).Mul(bData.price)
		borrowCoinValues.Increment(bCoin.Denom, bCoinUsdValue)
	}

	// Loan-to-Value ratio after sending keeper their reward
	ltv := borrowCoinValues.Sum().Quo(depositCoinValues.Sum())

	err := k.StartAuctions(ctx, deposits[0].Depositor, borrowBalances, aucDeposits, depositCoinValues, borrowCoinValues, ltv, liqMap)
	if err != nil {
		return err
	}

	return nil
}

// StartAuctions attempts to start auctions for seized assets
func (k Keeper) StartAuctions(ctx sdk.Context, borrower sdk.AccAddress, borrows, deposits sdk.Coins,
	depositCoinValues, borrowCoinValues types.ValuationMap, ltv sdk.Dec, liqMap map[string]LiqData) error {
	// Sort keys to ensure deterministic behavior
	bKeys := borrowCoinValues.GetSortedKeys()
	dKeys := depositCoinValues.GetSortedKeys()

	// Set up auction constants
	returnAddrs := []sdk.AccAddress{borrower}
	weights := []sdk.Int{sdk.NewInt(100)}
	debt := sdk.NewCoin("debt", sdk.ZeroInt())

	for _, bKey := range bKeys {
		bValue := borrowCoinValues.Get(bKey)
		maxLotSize := bValue.Quo(ltv)

		for _, dKey := range dKeys {
			dValue := depositCoinValues.Get(dKey)
			if maxLotSize.Equal(sdk.ZeroDec()) {
				break // exit out of the loop if we have cleared the full amount
			}

			if dValue.GTE(maxLotSize) { // We can start an auction for the whole borrow amount
				bid := sdk.NewCoin(bKey, borrows.AmountOf(bKey))
				lotSize := maxLotSize.MulInt(liqMap[dKey].conversionFactor).Quo(liqMap[dKey].price)
				lot := sdk.NewCoin(dKey, lotSize.TruncateInt())
				// Sanity check that we can deliver coins to the liquidator account
				if deposits.AmountOf(dKey).LT(lot.Amount) {
					return types.ErrInsufficientCoins
				}

				// Start auction: bid = full borrow amount, lot = maxLotSize
				err := k.supplyKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleAccountName, types.LiquidatorAccount, sdk.NewCoins(lot))
				if err != nil {
					return err
				}
				_, err = k.auctionKeeper.StartCollateralAuction(ctx, types.LiquidatorAccount, lot, bid, returnAddrs, weights, debt)
				if err != nil {
					return err
				}

				// Update USD valuation maps
				borrowCoinValues.SetZero(bKey)
				depositCoinValues.Decrement(dKey, maxLotSize)
				// Update deposits, borrows
				borrows = borrows.Sub(sdk.NewCoins(bid))
				deposits = deposits.Sub(sdk.NewCoins(lot))
				// Update max lot size
				maxLotSize = sdk.ZeroDec()
			} else { // We can only start an auction for the partial borrow amount
				maxBid := dValue.Mul(ltv)
				bidSize := maxBid.MulInt(liqMap[bKey].conversionFactor).Quo(liqMap[bKey].price)
				bid := sdk.NewCoin(bKey, bidSize.TruncateInt())
				lot := sdk.NewCoin(dKey, deposits.AmountOf(dKey))

				if bid.Amount.Equal(sdk.ZeroInt()) || lot.Amount.Equal(sdk.ZeroInt()) {
					continue
				}

				// Sanity check that we can deliver coins to the liquidator account
				if deposits.AmountOf(dKey).LT(lot.Amount) {
					return types.ErrInsufficientCoins
				}

				// Start auction: bid = maxBid, lot = whole deposit amount
				err := k.supplyKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleAccountName, types.LiquidatorAccount, sdk.NewCoins(lot))
				if err != nil {
					return err
				}
				_, err = k.auctionKeeper.StartCollateralAuction(ctx, types.LiquidatorAccount, lot, bid, returnAddrs, weights, debt)
				if err != nil {
					return err
				}

				// Update variables to account for partial auction
				borrowCoinValues.Decrement(bKey, maxBid)
				depositCoinValues.SetZero(dKey)
				// Update deposits, borrows
				borrows = borrows.Sub(sdk.NewCoins(bid))
				deposits = deposits.Sub(sdk.NewCoins(lot))
				// Update max lot size
				maxLotSize = borrowCoinValues.Get(bKey).Quo(ltv)
			}
		}
	}

	// Send any remaining deposit back to the original borrower
	for _, dKey := range dKeys {
		remaining := deposits.AmountOf(dKey)
		if remaining.GT(sdk.ZeroInt()) {
			returnCoin := sdk.NewCoins(sdk.NewCoin(dKey, remaining))
			err := k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccountName, borrower, returnCoin)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// GetCurrentLTV calculates the user's current LTV based on their deposits/borrows in the store
func (k Keeper) GetCurrentLTV(ctx sdk.Context, addr sdk.AccAddress) (sdk.Dec, error) {
	// Fetch deposits and parse coin denoms
	deposits := k.GetDepositsByUser(ctx, addr)
	depositDenoms := []string{}
	for _, deposit := range deposits {
		depositDenoms = append(depositDenoms, deposit.Amount.Denom)
	}

	// Fetch borrow balances and parse coin denoms
	borrowBalances := k.GetBorrowBalance(ctx, addr)
	if borrowBalances.IsZero() {
		return sdk.ZeroDec(), nil
	}
	borrowDenoms := getDenoms(borrowBalances)

	liqMap := make(map[string]LiqData)

	// Load required liquidation data for every deposit/borrow denom
	denoms := removeDuplicates(borrowDenoms, depositDenoms)
	for _, denom := range denoms {
		mm, found := k.GetMoneyMarket(ctx, denom)
		if !found {
			return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrMarketNotFound, "no market found for denom %s", denom)
		}

		priceData, err := k.pricefeedKeeper.GetCurrentPrice(ctx, mm.SpotMarketID)
		if err != nil {
			return sdk.ZeroDec(), err
		}

		liqMap[denom] = LiqData{priceData.Price, mm.BorrowLimit.LoanToValue, mm.ConversionFactor}
	}

	// Build valuation map to hold deposit coin USD valuations
	depositCoinValues := types.NewValuationMap()
	for _, deposit := range deposits {
		dData := liqMap[deposit.Amount.Denom]
		dCoinUsdValue := sdk.NewDecFromInt(deposit.Amount.Amount).Quo(sdk.NewDecFromInt(dData.conversionFactor)).Mul(dData.price)
		depositCoinValues.Increment(deposit.Amount.Denom, dCoinUsdValue)
	}

	// Build valuation map to hold borrow coin USD valuations
	borrowCoinValues := types.NewValuationMap()
	for _, bCoin := range borrowBalances {
		bData := liqMap[bCoin.Denom]
		bCoinUsdValue := sdk.NewDecFromInt(bCoin.Amount).Quo(sdk.NewDecFromInt(bData.conversionFactor)).Mul(bData.price)
		borrowCoinValues.Increment(bCoin.Denom, bCoinUsdValue)
	}

	// User doesn't have any deposits, catch divide by 0 error
	sumDeposits := depositCoinValues.Sum()
	if sumDeposits.Equal(sdk.ZeroDec()) {
		return sdk.ZeroDec(), nil
	}

	// Loan-to-Value ratio
	return borrowCoinValues.Sum().Quo(sumDeposits), nil
}

// UpdateItemInLtvIndex updates the key a borrower's address is stored under in the LTV index
func (k Keeper) UpdateItemInLtvIndex(ctx sdk.Context, prevLtv sdk.Dec, borrower sdk.AccAddress) error {
	currLtv, err := k.GetCurrentLTV(ctx, borrower)
	if err != nil {
		return err
	}

	k.RemoveFromLtvIndex(ctx, prevLtv, borrower)
	// If the user doesn't have any borrows their LTV is 0
	if currLtv.GT(sdk.ZeroDec()) {
		k.InsertIntoLtvIndex(ctx, currLtv, borrower)
	}
	return nil
}

func getDenoms(coins sdk.Coins) []string {
	denoms := []string{}
	for _, coin := range coins {
		denoms = append(denoms, coin.Denom)
	}
	return denoms
}

func removeDuplicates(one []string, two []string) []string {
	check := make(map[string]int)
	fullList := append(one, two...)

	res := []string{}
	for _, val := range fullList {
		check[val] = 1
	}

	for key := range check {
		res = append(res, key)
	}
	return res
}
