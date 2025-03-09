package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "simple-chain/testutil/keeper"
	"simple-chain/x/kudos/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.KudosKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
