package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// REST variable names
// nolint
const ()

// RegisterRoutes registers swap-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}

// PostCreateDepositReq defines the properties of a deposit create request's body
type PostCreateDepositReq struct {
	BaseReq  rest.BaseReq   `json:"base_req" yaml:"base_req"`
	From     sdk.AccAddress `json:"from" yaml:"from"`
	TokenA   sdk.Coin       `json:"token_a" yaml:"token_a"`
	TokenB   sdk.Coin       `json:"token_b" yaml:"token_b"`
	Slippage sdk.Dec        `json:"slippage" yaml:"slippage"`
	Deadline int64          `json:"deadline" yaml:"deadline"`
}

// PostCreateWithdrawReq defines the properties of a withdraw create request's body
type PostCreateWithdrawReq struct {
	BaseReq   rest.BaseReq   `json:"base_req" yaml:"base_req"`
	From      sdk.AccAddress `json:"from" yaml:"from"`
	Shares    sdk.Int        `json:"shares" yaml:"shares"`
	MinTokenA sdk.Coin       `json:"token_a" yaml:"token_a"`
	MinTokenB sdk.Coin       `json:"token_b" yaml:"token_b"`
	Deadline  int64          `json:"deadline" yaml:"deadline"`
}
