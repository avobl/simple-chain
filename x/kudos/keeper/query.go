package keeper

import (
	"simple-chain/x/kudos/types"
)

var _ types.QueryServer = Keeper{}
