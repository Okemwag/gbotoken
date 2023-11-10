package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "gbotoken/testutil/keeper"
	"gbotoken/x/gbotoken/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GbotokenKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
