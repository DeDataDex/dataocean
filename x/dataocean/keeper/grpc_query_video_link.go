package keeper

import (
	"context"

	"dataocean/x/dataocean/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VideoLinkAll(c context.Context, req *types.QueryAllVideoLinkRequest) (*types.QueryAllVideoLinkResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var videoLinks []types.VideoLink
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	videoLinkStore := prefix.NewStore(store, types.KeyPrefix(types.VideoLinkKeyPrefix))

	pageRes, err := query.Paginate(videoLinkStore, req.Pagination, func(key []byte, value []byte) error {
		var videoLink types.VideoLink
		if err := k.cdc.Unmarshal(value, &videoLink); err != nil {
			return err
		}

		videoLinks = append(videoLinks, videoLink)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVideoLinkResponse{VideoLink: videoLinks, Pagination: pageRes}, nil
}

func (k Keeper) VideoLink(c context.Context, req *types.QueryGetVideoLinkRequest) (*types.QueryGetVideoLinkResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVideoLink(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVideoLinkResponse{VideoLink: val}, nil
}
