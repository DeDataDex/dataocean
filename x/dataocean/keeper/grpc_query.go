package keeper

import (
	"dataocean/x/dataocean/types"
)

var _ types.QueryServer = Keeper{}
