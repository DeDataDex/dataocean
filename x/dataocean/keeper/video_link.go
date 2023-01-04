package keeper

import (
	"dataocean/x/dataocean/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetVideoLink set a specific videoLink in the store from its index
func (k Keeper) SetVideoLink(ctx sdk.Context, videoLink types.VideoLink) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VideoLinkKeyPrefix))
	b := k.cdc.MustMarshal(&videoLink)
	store.Set(types.VideoLinkKey(
		videoLink.Index,
	), b)
}

// GetVideoLink returns a videoLink from its index
func (k Keeper) GetVideoLink(
	ctx sdk.Context,
	index string,

) (val types.VideoLink, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VideoLinkKeyPrefix))

	b := store.Get(types.VideoLinkKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVideoLink removes a videoLink from the store
func (k Keeper) RemoveVideoLink(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VideoLinkKeyPrefix))
	store.Delete(types.VideoLinkKey(
		index,
	))
}

// GetAllVideoLink returns all videoLink
func (k Keeper) GetAllVideoLink(ctx sdk.Context) (list []types.VideoLink) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VideoLinkKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VideoLink
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
