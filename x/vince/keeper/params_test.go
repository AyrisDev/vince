package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "vince/testutil/keeper"
	"vince/x/vince/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VinceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
