package keeper

import (
	"context"

	"github.com/cosmonaut/vote/x/vote/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/cosmos/cosmos-sdk/types/query" 
)



func (k Keeper) Voteresult(goCtx context.Context, req *types.QueryVoteresultRequest) (*types.QueryVoteresultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	

	store := ctx.KVStore(k.storeKey)

        voteStore := prefix.NewStore(store, []byte(types.VoteKey))
        
        iterator := voteStore.Iterator(nil, nil)
	defer iterator.Close()

	var votes []*types.Vote
	_, err := query.Paginate(voteStore, req.Pagination, func(key []byte, value []byte) error {
	    var vote types.Vote
	    if err := k.cdc.Unmarshal(value, &vote); err != nil {
	      return err
	    }
	    votes = append(votes, &vote)
	    return nil
	  })
      if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
      }
  	
	return &types.QueryVoteresultResponse{Votes: votes}, nil
}
