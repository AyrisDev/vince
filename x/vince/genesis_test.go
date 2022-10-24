package vince_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vince/testutil/keeper"
	"vince/testutil/nullify"
	"vince/x/vince"
	"vince/x/vince/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VinceKeeper(t)
	vince.InitGenesis(ctx, *k, genesisState)
	got := vince.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
