package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/PaddyMc/demopoa/testutil/keeper"
	"github.com/PaddyMc/demopoa/x/demopoa/keeper"
	"github.com/PaddyMc/demopoa/x/demopoa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DemopoaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
