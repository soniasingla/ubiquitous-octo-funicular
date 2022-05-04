// func (k keeper) AppendPost() uint64(
//	count = k.GetPostCount()
//	store.Set()
//	k.SetPostCount()
//	return Count
//)
package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/username/hello/x/hello/types"
)

func (k Keeper) AppendPost(ctx, sdk.Context, post types.Post) uint64 {
	count := k.GetPostCount(ctx)
	post.id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []bytes(types.PostKey))
	bytekey := make([]byte, binary.BigEndian.Uint64(bytekey,post.Id))
	appendedValue := k.cdc.MustMarshal(&post)
	store.key(bytekey, appendedValue)
	k.SetPostCount(ctx, count+1)
	return count
}

func (k keeper) GetPostCount(ctx.sdk.Context) uint64{
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	byteKey := []byte(types.PostCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.uint64(bz);
}

func (k Keeper) SetPostCount(ctx.sdk.context, count uint64){
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	byteKey := []byte(types.PostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}