package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/avobl/simple-chain/x/kudos/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AppendKudos appends a kudos in the store with a new id and update the count
func (k Keeper) AppendKudos(ctx sdk.Context, kudos types.Kudos) {
	count := k.GetKudosCount(ctx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.KudosKey))
	appendedValue := k.cdc.MustMarshal(&kudos)
	store.Set(GetKudosIDBytes(count), appendedValue)
	k.SetKudosCount(ctx, count+1)
}

// GetKudosCount get the total number of kudos
func (k Keeper) GetKudosCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.KudosCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

// SetKudosCount into blockchain
func (k Keeper) SetKudosCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.KudosCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func GetKudosIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)

	return bz
}
