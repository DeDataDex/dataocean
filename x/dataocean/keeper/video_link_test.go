package keeper_test

import (
	"strconv"
	"testing"

	keepertest "dataocean/testutil/keeper"
	"dataocean/testutil/nullify"
	"dataocean/x/dataocean/keeper"
	"dataocean/x/dataocean/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNVideoLink(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VideoLink {
	items := make([]types.VideoLink, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetVideoLink(ctx, items[i])
	}
	return items
}

func TestVideoLinkGet(t *testing.T) {
	keeper, ctx := keepertest.DataoceanKeeper(t)
	items := createNVideoLink(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVideoLink(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestVideoLinkRemove(t *testing.T) {
	keeper, ctx := keepertest.DataoceanKeeper(t)
	items := createNVideoLink(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVideoLink(ctx,
			item.Index,
		)
		_, found := keeper.GetVideoLink(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestVideoLinkGetAll(t *testing.T) {
	keeper, ctx := keepertest.DataoceanKeeper(t)
	items := createNVideoLink(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVideoLink(ctx)),
	)
}
