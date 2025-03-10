package simplechain_test

import (
	"testing"

	keepertest "github.com/avobl/simple-chain/testutil/keeper"
	"github.com/avobl/simple-chain/testutil/nullify"
	"github.com/avobl/simple-chain/x/simplechain/module"
	"github.com/avobl/simple-chain/x/simplechain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SimplechainKeeper(t)
	simplechain.InitGenesis(ctx, k, genesisState)
	got := simplechain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
