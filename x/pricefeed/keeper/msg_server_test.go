package keeper_test

import (
	sdkmath "cosmossdk.io/math"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kava-labs/kava/app"
	"github.com/kava-labs/kava/x/pricefeed/keeper"
	"github.com/kava-labs/kava/x/pricefeed/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_PostPrice(t *testing.T) {
	_, addrs := app.GeneratePrivKeyAddressPairs(4)
	tApp := app.NewTestApp()
	ctx := tApp.NewContextLegacy(true, tmproto.Header{Time: tmtime.Now().UTC()})
	k := tApp.GetPriceFeedKeeper()
	msgSrv := keeper.NewMsgServerImpl(k)

	authorizedOracles := addrs[:2]
	unauthorizedAddrs := addrs[2:]

	mp := types.Params{
		Markets: []types.Market{
			{MarketID: "tstusd", BaseAsset: "tst", QuoteAsset: "usd", Oracles: authorizedOracles, Active: true},
		},
	}
	k.SetParams(ctx, mp)

	now := time.Now().UTC()

	tests := []struct {
		giveMsg      string
		giveOracle   sdk.AccAddress
		giveMarketId string
		giveExpiry   time.Time
		wantAccepted bool
		errorKind    error
	}{
		{"authorized", authorizedOracles[0], "tstusd", now.Add(time.Hour * 1), true, nil},
		{"expired", authorizedOracles[0], "tstusd", now.Add(-time.Hour * 1), false, types.ErrExpired},
		{"invalid", authorizedOracles[0], "invalid", now.Add(time.Hour * 1), false, types.ErrInvalidMarket},
		{"unauthorized", unauthorizedAddrs[0], "tstusd", now.Add(time.Hour * 1), false, types.ErrInvalidOracle},
	}

	for _, tt := range tests {
		t.Run(tt.giveMsg, func(t *testing.T) {
			// Use MsgServer over keeper methods directly to tests against valid oracles
			msg := types.NewMsgPostPrice(tt.giveOracle.String(), tt.giveMarketId, sdkmath.LegacyMustNewDecFromStr("0.5"), tt.giveExpiry)
			_, err := msgSrv.PostPrice(sdk.WrapSDKContext(ctx), msg)

			if tt.wantAccepted {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.ErrorIs(t, tt.errorKind, err)
			}
		})
	}
}
