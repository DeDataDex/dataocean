package dataocean

import (
	"dataocean/x/dataocean/keeper"
	"dataocean/x/dataocean/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the video
	for _, elem := range genState.VideoList {
		k.SetVideo(ctx, elem)
	}

	// Set video count
	k.SetVideoCount(ctx, genState.VideoCount)
	// Set all the videoLink
	for _, elem := range genState.VideoLinkList {
		k.SetVideoLink(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.VideoList = k.GetAllVideo(ctx)
	genesis.VideoCount = k.GetVideoCount(ctx)
	genesis.VideoLinkList = k.GetAllVideoLink(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
