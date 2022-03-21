package simulation

import (
	"math/rand"

	"github.com/cosmonaut/vote/x/vote/keeper"
	"github.com/cosmonaut/vote/x/vote/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgVote(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgVote{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Vote simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Vote simulation not implemented"), nil, nil
	}
}
