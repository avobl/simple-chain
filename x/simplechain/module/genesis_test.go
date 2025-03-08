package simplechain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "simple-chain/testutil/keeper"
	"simple-chain/testutil/nullify"
	"simple-chain/x/simplechain/module"
	"simple-chain/x/simplechain/types"
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
