package keeper

import (
	"nameservice/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
