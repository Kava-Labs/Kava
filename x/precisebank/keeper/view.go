package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kava-labs/kava/x/precisebank/types"
)

// GetBalance returns the balance of a specific denom for an address. This will
// return the extended balance for the ExtendedCoinDenom, and the regular
// balance for all other denoms.
func (k Keeper) GetBalance(
	ctx sdk.Context,
	addr sdk.AccAddress,
	denom string,
) sdk.Coin {
	// Pass through to x/bank for denoms except ExtendedCoinDenom
	if denom != types.ExtendedCoinDenom {
		return k.bk.GetBalance(ctx, addr, denom)
	}

	// x/bank for integer balance - full balance, including locked
	integerCoins := k.bk.GetBalance(ctx, addr, types.IntegerCoinDenom)

	// x/precisebank for fractional balance
	fractionalAmount := k.GetFractionalBalance(ctx, addr)

	// (Integer * ConversionFactor) + Fractional
	fullAmount := integerCoins.
		Amount.
		Mul(types.ConversionFactor()).
		Add(fractionalAmount)

	return sdk.NewCoin(types.ExtendedCoinDenom, fullAmount)
}

// SpendableCoins returns the total balances of spendable coins for an account
// by address. If the account has no spendable coins, an empty Coins slice is
// returned.
func (k Keeper) SpendableCoin(
	ctx sdk.Context,
	addr sdk.AccAddress,
	denom string,
) sdk.Coin {
	// Pass through to x/bank for denoms except ExtendedCoinDenom
	if denom != types.ExtendedCoinDenom {
		return k.bk.SpendableCoin(ctx, addr, denom)
	}

	// x/bank for integer balance - excluding locked
	integerCoin := k.bk.SpendableCoin(ctx, addr, types.IntegerCoinDenom)

	// x/precisebank for fractional balance
	fractionalAmount := k.GetFractionalBalance(ctx, addr)

	// Spendable = (Integer * ConversionFactor) + Fractional
	fullAmount := integerCoin.Amount.
		Mul(types.ConversionFactor()).
		Add(fractionalAmount)

	return sdk.NewCoin(types.ExtendedCoinDenom, fullAmount)
}
