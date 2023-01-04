package dataocean_test

import (
	"testing"

	keepertest "dataocean/testutil/keeper"
	"dataocean/testutil/nullify"
	"dataocean/x/dataocean"
	"dataocean/x/dataocean/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		VideoList: []types.Video{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		VideoCount: 2,
		VideoLinkList: []types.VideoLink{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DataoceanKeeper(t)
	dataocean.InitGenesis(ctx, *k, genesisState)
	got := dataocean.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.VideoList, got.VideoList)
	require.Equal(t, genesisState.VideoCount, got.VideoCount)
	require.ElementsMatch(t, genesisState.VideoLinkList, got.VideoLinkList)
	// this line is used by starport scaffolding # genesis/test/assert
}
