package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/kava-labs/kava/x/community/types"
)

// Keeper of the community store
type Keeper struct {
	paramSpace    paramtypes.Subspace
	bankKeeper    types.BankKeeper
	moduleAddress sdk.AccAddress
}

// NewKeeper creates a new community Keeper instance
func NewKeeper(ak types.AccountKeeper, bk types.BankKeeper) Keeper {
	// ensure community module account is set
	addr := ak.GetModuleAddress(types.ModuleName)
	if addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	return Keeper{
		bankKeeper:    bk,
		moduleAddress: addr,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// GetParams returns the total set of x/community parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of x/community parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetModuleAccountBalance returns all the coins held by the community module account
func (k Keeper) GetModuleAccountBalance(ctx sdk.Context) sdk.Coins {
	return k.bankKeeper.GetAllBalances(ctx, k.moduleAddress)
}

// FundCommunityPool transfers coins from the sender to the community module account.
func (k Keeper) FundCommunityPool(ctx sdk.Context, sender sdk.AccAddress, amount sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, amount)
}
