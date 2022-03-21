package keeper

import (
	"encoding/binary"
	"github.com/cosmonaut/vote/x/vote/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendName(ctx sdk.Context, name types.Name) uint64 {

	count := k.GetNameCount(ctx)

	name.Id = count
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NameKey))
	// Convert the name ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, name.Id)
	// Marshal the name into bytes
	appendedValue := k.cdc.MustMarshal(&name)
	// Insert the name bytes using post ID as a key
	store.Set(byteKey, appendedValue)
	// Update the name count
	k.SetNameCount(ctx, count+1)
	return count
}

func (k Keeper) GetNameCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NameCountKey))
	byteKey := []byte(types.NameCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetNameCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NameCountKey))
	// Convert the NameCountKey to bytes
	byteKey := []byte(types.NameCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of Name-count- to count
	store.Set(byteKey, bz)
}
