package keeper

import (
	"context"
	"strconv"

	"dataocean/x/dataocean/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateVideo(goCtx context.Context, msg *types.MsgCreateVideo) (*types.MsgCreateVideoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	video := types.Video{
		Creator:     msg.Creator,
		Title:       msg.Title,
		Description: msg.Description,
		CoverLink:   msg.CoverLink,
		VideoLink:   msg.VideoLink,
		PriceMB:     msg.PriceMB,
		CreatedAt:   uint64(ctx.BlockHeight()),
	}

	id := k.AppendVideo(ctx, video)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.TypeMsgCreateVideo, sdk.NewAttribute("video_id", strconv.Itoa(int(id)))))

	return &types.MsgCreateVideoResponse{
		Id: id,
	}, nil
}
