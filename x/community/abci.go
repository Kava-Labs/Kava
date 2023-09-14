package community

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kava-labs/kava/x/community/keeper"
	"github.com/kava-labs/kava/x/community/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	if k.ShouldStartDisableInflationUpgrade(ctx) {
		k.StartDisableInflationUpgrade(ctx)
	}

	if err := k.PayCommunityRewards(ctx); err != nil {
		panic(err)
	}
}
