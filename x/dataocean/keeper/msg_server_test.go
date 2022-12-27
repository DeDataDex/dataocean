package keeper_test

import (
	"context"
	"testing"

	keepertest "dataocean/testutil/keeper"
	"dataocean/x/dataocean/keeper"
	"dataocean/x/dataocean/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DataoceanKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
