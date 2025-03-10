package keeper

import (
	"github.com/avobl/simple-chain/x/simplechain/types"
)

var _ types.QueryServer = Keeper{}
