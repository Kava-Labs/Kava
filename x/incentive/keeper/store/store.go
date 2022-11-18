package store

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type IncentiveStore struct {
	cdc codec.Codec
	key sdk.StoreKey
}

// NewIncentiveStore creates a new IncentiveStore
func NewIncentiveStore(cdc codec.Codec, key sdk.StoreKey) IncentiveStore {
	return IncentiveStore{
		cdc: cdc,
		key: key,
	}
}
