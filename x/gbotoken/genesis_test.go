package gbotoken_test

import (
	"testing"

	keepertest "gbotoken/testutil/keeper"
	"gbotoken/testutil/nullify"
	"gbotoken/x/gbotoken"
	"gbotoken/x/gbotoken/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GbotokenKeeper(t)
	gbotoken.InitGenesis(ctx, *k, genesisState)
	got := gbotoken.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
