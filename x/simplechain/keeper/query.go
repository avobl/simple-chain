package keeper

import (
	"simple-chain/x/simplechain/types"
)

var _ types.QueryServer = Keeper{}
