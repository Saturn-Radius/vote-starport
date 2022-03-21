package keeper

import (
	"encoding/binary"
	"github.com/cosmonaut/vote/x/vote/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendVote(ctx sdk.Context, vote types.Vote) uint64 {

	count := k.GetVoteCount(ctx)

	vote.Id = count
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.VoteKey))
	// Convert the name ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, vote.Id)
	// Marshal the name into bytes
	appendedValue := k.cdc.MustMarshal(&vote)
	// Insert the name bytes using post ID as a key
	store.Set(byteKey, appendedValue)
	// Update the name count
	k.SetVoteCount(ctx, count+1)
	return count
}

func (k Keeper) GetVoteCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.VoteCountKey))
	byteKey := []byte(types.VoteCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetVoteCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.VoteCountKey))
	// Convert the VoteCountKey to bytes
	byteKey := []byte(types.VoteCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of Name-count- to count
	store.Set(byteKey, bz)
}
