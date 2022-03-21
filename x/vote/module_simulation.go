package vote

import (
	"math/rand"

	"github.com/cosmonaut/vote/testutil/sample"
	votesimulation "github.com/cosmonaut/vote/x/vote/simulation"
	"github.com/cosmonaut/vote/x/vote/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = votesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRegisterName = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterName int = 100

	opWeightMsgVote = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVote int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	voteGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&voteGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterName, &weightMsgRegisterName, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterName = defaultWeightMsgRegisterName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterName,
		votesimulation.SimulateMsgRegisterName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVote int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVote, &weightMsgVote, nil,
		func(_ *rand.Rand) {
			weightMsgVote = defaultWeightMsgVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVote,
		votesimulation.SimulateMsgVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
