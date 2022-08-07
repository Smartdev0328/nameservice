package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "nameservice/testutil/keeper"
	"nameservice/x/nameservice/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NameserviceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
