package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		VideoList: []Video{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in video
	videoIdMap := make(map[uint64]bool)
	videoCount := gs.GetVideoCount()
	for _, elem := range gs.VideoList {
		if _, ok := videoIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for video")
		}
		if elem.Id >= videoCount {
			return fmt.Errorf("video id should be lower or equal than the last id")
		}
		videoIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
