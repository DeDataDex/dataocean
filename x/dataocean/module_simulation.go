package dataocean

import (
	"math/rand"

	"dataocean/testutil/sample"
	dataoceansimulation "dataocean/x/dataocean/simulation"
	"dataocean/x/dataocean/types"
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
	_ = dataoceansimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateVideo = "op_weight_msg_create_video"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVideo int = 100

	opWeightMsgPlayVideo = "op_weight_msg_play_video"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPlayVideo int = 100

	opWeightMsgPaySign = "op_weight_msg_pay_sign"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPaySign int = 100

	opWeightMsgSubmitPaySign = "op_weight_msg_submit_pay_sign"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitPaySign int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dataoceanGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dataoceanGenesis)
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

	var weightMsgCreateVideo int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateVideo, &weightMsgCreateVideo, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVideo = defaultWeightMsgCreateVideo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVideo,
		dataoceansimulation.SimulateMsgCreateVideo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPlayVideo int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPlayVideo, &weightMsgPlayVideo, nil,
		func(_ *rand.Rand) {
			weightMsgPlayVideo = defaultWeightMsgPlayVideo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPlayVideo,
		dataoceansimulation.SimulateMsgPlayVideo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPaySign int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPaySign, &weightMsgPaySign, nil,
		func(_ *rand.Rand) {
			weightMsgPaySign = defaultWeightMsgPaySign
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPaySign,
		dataoceansimulation.SimulateMsgPaySign(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitPaySign int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitPaySign, &weightMsgSubmitPaySign, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitPaySign = defaultWeightMsgSubmitPaySign
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitPaySign,
		dataoceansimulation.SimulateMsgSubmitPaySign(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
