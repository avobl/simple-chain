package keeper

import (
	"github.com/avobl/simple-chain/x/kudos/types"
)

var _ types.QueryServer = Keeper{}
