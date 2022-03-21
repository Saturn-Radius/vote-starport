package keeper

import (
	"context"

	"github.com/cosmonaut/vote/x/vote/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var vote = types.Vote{
		Creator: msg.Creator,
		Name:    msg.Name,
	}

	k.AppendVote(ctx, vote)

	return &types.MsgVoteResponse{}, nil
}
