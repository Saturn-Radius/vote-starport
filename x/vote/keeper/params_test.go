package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmonaut/vote/testutil/keeper"
	"github.com/cosmonaut/vote/x/vote/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VoteKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
