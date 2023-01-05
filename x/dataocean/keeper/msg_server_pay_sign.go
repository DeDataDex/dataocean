package keeper

import (
	"context"

	"dataocean/x/dataocean/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PaySign(goCtx context.Context, msg *types.MsgPaySign) (*types.MsgPaySignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgPaySignResponse{}, nil
}
