package keeper_test

import (
	"testing"

	keepertest "dataocean/testutil/keeper"
	"dataocean/testutil/nullify"
	"dataocean/x/dataocean/keeper"
	"dataocean/x/dataocean/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNVideo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Video {
	items := make([]types.Video, n)
	for i := range items {
		items[i].Id = keeper.AppendVideo(ctx, items[i])
	}
	return items
}

func TestVideoGet(t *testing.T) {
	keeper, ctx := keepertest.DataoceanKeeper(t)
	items := createNVideo(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetVideo(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestVideoRemove(t *testing.T) {
	keeper, ctx := keepertest.DataoceanKeeper(t)
	items := createNVideo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVideo(ctx, item.Id)
		_, found := keeper.GetVideo(ctx, item.Id)
		require.False(t, found)
	}
}

func TestVideoGetAll(t *testing.T) {
	keeper, ctx := keepertest.DataoceanKeeper(t)
	items := createNVideo(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVideo(ctx)),
	)
}

func TestVideoCount(t *testing.T) {
	keeper, ctx := keepertest.DataoceanKeeper(t)
	items := createNVideo(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetVideoCount(ctx))
}
