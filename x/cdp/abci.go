package cdp

import (
	sdkmath "cosmossdk.io/math"
	"errors"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kava-labs/kava/x/cdp/keeper"
	"github.com/kava-labs/kava/x/cdp/types"
	pricefeedtypes "github.com/kava-labs/kava/x/pricefeed/types"
)

// BeginBlocker compounds the debt in outstanding cdps and liquidates cdps that are below the required collateralization ratio
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	params := k.GetParams(ctx)

	// only run CDP liquidations every `LiquidationBlockInterval` blocks
	skipSyncronizeAndLiquidations := ctx.BlockHeight()%params.LiquidationBlockInterval != 0

	fmt.Println("skipSyncronizeAndLiquidations ", skipSyncronizeAndLiquidations)

	for _, cp := range params.CollateralParams {
		fmt.Println("cp ", cp)
		ok := k.UpdatePricefeedStatus(ctx, cp.SpotMarketID)
		if !ok {
			continue
		}

		ok = k.UpdatePricefeedStatus(ctx, cp.LiquidationMarketID)
		if !ok {
			continue
		}

		err := k.AccumulateInterest(ctx, cp.Type)
		if err != nil {
			panic(err)
		}

		if skipSyncronizeAndLiquidations {
			ctx.Logger().Debug(fmt.Sprintf("skipping x/cdp SynchronizeInterestForRiskyCDPs and LiquidateCdps for %s", cp.Type))
			continue
		}

		ctx.Logger().Debug(fmt.Sprintf("running x/cdp SynchronizeInterestForRiskyCDPs and LiquidateCdps for %s", cp.Type))

		err = k.SynchronizeInterestForRiskyCDPs(ctx, sdkmath.LegacyMaxSortableDec, cp)
		if err != nil {
			panic(err)
		}

		err = k.LiquidateCdps(ctx, cp.LiquidationMarketID, cp.Type, cp.LiquidationRatio, cp.CheckCollateralizationIndexCount)
		if err != nil && !errors.Is(err, pricefeedtypes.ErrNoValidPrice) {
			panic(err)
		}
	}

	err := k.RunSurplusAndDebtAuctions(ctx)
	if err != nil {
		panic(err)
	}
}
