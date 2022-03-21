package vote_test

import (
	"testing"

	keepertest "github.com/cosmonaut/vote/testutil/keeper"
	"github.com/cosmonaut/vote/testutil/nullify"
	"github.com/cosmonaut/vote/x/vote"
	"github.com/cosmonaut/vote/x/vote/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VoteKeeper(t)
	vote.InitGenesis(ctx, *k, genesisState)
	got := vote.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
