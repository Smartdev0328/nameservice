package nameservice_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "nameservice/testutil/keeper"
	"nameservice/testutil/nullify"
	"nameservice/x/nameservice"
	"nameservice/x/nameservice/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		WhoisList: []types.Whois{
		{
			Index: "0",
},
		{
			Index: "1",
},
	},
	// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NameserviceKeeper(t)
	nameservice.InitGenesis(ctx, *k, genesisState)
	got := nameservice.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.WhoisList, got.WhoisList)
// this line is used by starport scaffolding # genesis/test/assert
}
