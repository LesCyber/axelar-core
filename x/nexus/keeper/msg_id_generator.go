package keeper

import (
	"crypto/sha256"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/axelarnetwork/axelar-core/utils"
	"github.com/axelarnetwork/axelar-core/x/nexus/types"
)

var _ types.MsgIDGenerator = &Keeper{}

func getTxHash(ctx sdk.Context) [32]byte {
	return sha256.Sum256(ctx.TxBytes())
}

// IncrID increments the nonce
func (k Keeper) IncrID(ctx sdk.Context) {
	utils.NewCounter[uint64](messageNonceKey, k.getStore(ctx)).Incr(ctx)
}

// currIDAndIncr returns the current transaction hash and nonce, and increments the nonce
func (k Keeper) currIDAndIncr(ctx sdk.Context) ([32]byte, uint64) {
	return getTxHash(ctx), utils.NewCounter[uint64](messageNonceKey, k.getStore(ctx)).Incr(ctx)
}

// CurrID returns the current transaction hash and nonce
func (k Keeper) CurrID(ctx sdk.Context) ([32]byte, uint64) {
	return getTxHash(ctx), utils.NewCounter[uint64](messageNonceKey, k.getStore(ctx)).Curr(ctx)
}
