package kudos_test

import (
	"testing"

	keepertest "github.com/avobl/simple-chain/testutil/keeper"
	"github.com/avobl/simple-chain/testutil/nullify"
	"github.com/avobl/simple-chain/x/kudos/module"
	"github.com/avobl/simple-chain/x/kudos/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KudosKeeper(t)
	kudos.InitGenesis(ctx, k, genesisState)
	got := kudos.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
