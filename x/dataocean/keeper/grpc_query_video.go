package keeper

import (
	"context"

	"dataocean/x/dataocean/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VideoAll(c context.Context, req *types.QueryAllVideoRequest) (*types.QueryAllVideoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var videos []types.Video
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	videoStore := prefix.NewStore(store, types.KeyPrefix(types.VideoKey))

	pageRes, err := query.Paginate(videoStore, req.Pagination, func(key []byte, value []byte) error {
		var video types.Video
		if err := k.cdc.Unmarshal(value, &video); err != nil {
			return err
		}

		videos = append(videos, video)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVideoResponse{Video: videos, Pagination: pageRes}, nil
}

func (k Keeper) Video(c context.Context, req *types.QueryGetVideoRequest) (*types.QueryGetVideoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	video, found := k.GetVideo(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetVideoResponse{Video: video}, nil
}
