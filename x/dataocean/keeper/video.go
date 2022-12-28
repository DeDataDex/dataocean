package keeper

import (
	"encoding/binary"

	"dataocean/x/dataocean/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetVideoCount get the total number of video
func (k Keeper) GetVideoCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VideoCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVideoCount set the total number of video
func (k Keeper) SetVideoCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VideoCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVideo appends a video in the store with a new id and update the count
func (k Keeper) AppendVideo(
	ctx sdk.Context,
	video types.Video,
) uint64 {
	// Create the video
	count := k.GetVideoCount(ctx)

	// Set the ID of the appended value
	video.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VideoKey))
	appendedValue := k.cdc.MustMarshal(&video)
	store.Set(GetVideoIDBytes(video.Id), appendedValue)

	// Update video count
	k.SetVideoCount(ctx, count+1)

	return count
}

// SetVideo set a specific video in the store
func (k Keeper) SetVideo(ctx sdk.Context, video types.Video) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VideoKey))
	b := k.cdc.MustMarshal(&video)
	store.Set(GetVideoIDBytes(video.Id), b)
}

// GetVideo returns a video from its id
func (k Keeper) GetVideo(ctx sdk.Context, id uint64) (val types.Video, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VideoKey))
	b := store.Get(GetVideoIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVideo removes a video from the store
func (k Keeper) RemoveVideo(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VideoKey))
	store.Delete(GetVideoIDBytes(id))
}

// GetAllVideo returns all video
func (k Keeper) GetAllVideo(ctx sdk.Context) (list []types.Video) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VideoKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Video
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetVideoIDBytes returns the byte representation of the ID
func GetVideoIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetVideoIDFromBytes returns ID in uint64 format from a byte array
func GetVideoIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
