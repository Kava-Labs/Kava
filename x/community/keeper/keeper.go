package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/kava-labs/kava/x/community/types"
)

// Keeper of the community store
type Keeper struct {
	key storetypes.StoreKey
	cdc codec.Codec

	bankKeeper     types.BankKeeper
	cdpKeeper      types.CdpKeeper
	distrKeeper    types.DistributionKeeper
	hardKeeper     types.HardKeeper
	moduleAddress  sdk.AccAddress
	mintKeeper     types.MintKeeper
	kavadistKeeper types.KavadistKeeper

	legacyCommunityPoolAddress sdk.AccAddress
}

// NewKeeper creates a new community Keeper instance
func NewKeeper(
	cdc codec.Codec,
	key storetypes.StoreKey,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	ck types.CdpKeeper,
	dk types.DistributionKeeper,
	hk types.HardKeeper,
	mk types.MintKeeper,
	kk types.KavadistKeeper,
) Keeper {
	// ensure community module account is set
	addr := ak.GetModuleAddress(types.ModuleAccountName)
	if addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleAccountName))
	}
	legacyAddr := ak.GetModuleAddress(types.LegacyCommunityPoolModuleName)
	if addr == nil {
		panic("legacy community pool address not found")
	}

	return Keeper{
		key: key,
		cdc: cdc,

		bankKeeper:     bk,
		cdpKeeper:      ck,
		distrKeeper:    dk,
		hardKeeper:     hk,
		mintKeeper:     mk,
		kavadistKeeper: kk,
		moduleAddress:  addr,

		legacyCommunityPoolAddress: legacyAddr,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// GetModuleAccountBalance returns all the coins held by the community module account
func (k Keeper) GetModuleAccountBalance(ctx sdk.Context) sdk.Coins {
	return k.bankKeeper.GetAllBalances(ctx, k.moduleAddress)
}

// FundCommunityPool transfers coins from the sender to the community module account.
func (k Keeper) FundCommunityPool(ctx sdk.Context, sender sdk.AccAddress, amount sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleAccountName, amount)
}

// DistributeFromCommunityPool transfers coins from the community pool to recipient.
func (k Keeper) DistributeFromCommunityPool(ctx sdk.Context, recipient sdk.AccAddress, amount sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccountName, recipient, amount)
}

// GetPreviousBlockTime returns the blockTime for the previous block after the inflation upgrade.
// If inflation upgrade has not been executed, this will not be set.
func (k Keeper) GetPreviousBlockTime(ctx sdk.Context) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousBlockTimeKey)
	b := store.Get(types.PreviousBlockTimeKey)
	if b == nil {
		return time.Time{}, false
	}
	if err := blockTime.UnmarshalBinary(b); err != nil {
		panic(err)
	}
	return blockTime, true
}

// SetPreviousBlockTime set the time of the previous block
func (k Keeper) SetPreviousBlockTime(ctx sdk.Context, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousBlockTimeKey)
	b, err := blockTime.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set(types.PreviousBlockTimeKey, b)
}
