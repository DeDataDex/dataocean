package keeper_test

import (
	"testing"

	testkeeper "dataocean/testutil/keeper"
	"dataocean/x/dataocean/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DataoceanKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
