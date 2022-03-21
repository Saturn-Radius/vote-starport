package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmonaut/vote/testutil/keeper"
	"github.com/cosmonaut/vote/x/vote/keeper"
	"github.com/cosmonaut/vote/x/vote/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.VoteKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
