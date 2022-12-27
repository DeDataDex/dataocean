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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DataoceanKeeper(t)
	dataocean.InitGenesis(ctx, *k, genesisState)
	got := dataocean.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
