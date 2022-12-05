package kavamint

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kava-labs/kava/x/kavamint/keeper"
	"github.com/kava-labs/kava/x/kavamint/types"
)

// InitGenesis new mint genesis
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, ak types.AccountKeeper, data *types.GenesisState) {
	keeper.SetParams(ctx, data.Params)

	if data.PreviousBlockTime.IsZero() {
		panic(fmt.Sprintf("No previous_block_time set in genesis state of %s module", types.ModuleName))
	}
	keeper.SetPreviousBlockTime(ctx, data.PreviousBlockTime)

	if macc := ak.GetModuleAccount(ctx, types.ModuleName); macc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	params := keeper.GetParams(ctx)
	previousBlockTime, found := keeper.GetPreviousBlockTime(ctx)
	if !found || previousBlockTime.IsZero() {
		previousBlockTime = ctx.BlockTime()
	}
	return types.NewGenesisState(params, previousBlockTime)
}
