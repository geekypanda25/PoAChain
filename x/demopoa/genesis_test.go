package demopoa_test

import (
	"testing"

	keepertest "github.com/PaddyMc/demopoa/testutil/keeper"
	"github.com/PaddyMc/demopoa/testutil/nullify"
	"github.com/PaddyMc/demopoa/x/demopoa"
	"github.com/PaddyMc/demopoa/x/demopoa/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DemopoaKeeper(t)
	demopoa.InitGenesis(ctx, *k, genesisState)
	got := demopoa.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
