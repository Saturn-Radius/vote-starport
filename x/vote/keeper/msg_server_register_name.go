package keeper

import (
	"context"

	"github.com/cosmonaut/vote/x/vote/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterName(goCtx context.Context, msg *types.MsgRegisterName) (*types.MsgRegisterNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var name = types.Name{
		Creator: msg.Creator,
		Name:    msg.Name,
	}

	k.AppendName(ctx, name)

	return &types.MsgRegisterNameResponse{}, nil
}
