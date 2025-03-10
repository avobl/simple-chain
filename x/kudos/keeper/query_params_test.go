package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/avobl/simple-chain/testutil/keeper"
	"github.com/avobl/simple-chain/x/kudos/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.KudosKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
