package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"nameservice/x/nameservice/types"
)

func (k msgServer) ByName(goCtx context.Context, msg *types.MsgByName) (*types.MsgByNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgByNameResponse{}, nil
}
