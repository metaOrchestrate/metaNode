package keeper

import (
	"testing"

	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	"github.com/celestiaorg/celestia-app/x/blob/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func keeper(t *testing.T) (*Keeper, store.CommitMultiStore) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	tempCtx := sdk.NewContext(stateStore, tmproto.Header{}, false, nil)

	aminoCdc := codec.NewLegacyAmino()
	paramsSubspace := typesparams.NewSubspace(cdc,
		aminoCdc,
		storeKey,
		memStoreKey,
		"Blob",
	)
	k := NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
	)
	k.SetParams(tempCtx, types.DefaultParams())

	return k, stateStore
}

func TestPayForBlobGas(t *testing.T) {
	type testCase struct {
		name            string
		msg             types.MsgPayForBlobs
		wantGasConsumed uint64
	}

	paramLookUpCost := uint32(1060)

	testCases := []testCase{
		{
			name:            "1 byte blob", // occupies 1 share
			msg:             types.MsgPayForBlobs{BlobSizes: []uint32{1}},
			wantGasConsumed: uint64(1*appconsts.ShareSize*types.DefaultGasPerBlobByte + paramLookUpCost), // 1 share * 512 bytes per share * 8 gas per byte + 1060 gas for fetching param = 5156 gas
		},
		{
			name:            "100 byte blob", // occupies 1 share
			msg:             types.MsgPayForBlobs{BlobSizes: []uint32{100}},
			wantGasConsumed: uint64(1*appconsts.ShareSize*types.DefaultGasPerBlobByte + paramLookUpCost),
		},
		{
			name:            "1024 byte blob", // occupies 3 shares because share prefix (e.g. namespace, info byte)
			msg:             types.MsgPayForBlobs{BlobSizes: []uint32{1024}},
			wantGasConsumed: uint64(3*appconsts.ShareSize*types.DefaultGasPerBlobByte + paramLookUpCost), // 3 shares * 512 bytes per share * 8 gas per byte + 1060 gas for fetching param = 13348 gas
		},
		{
			name:            "3 blobs, 1 share each",
			msg:             types.MsgPayForBlobs{BlobSizes: []uint32{1, 1, 1}},
			wantGasConsumed: uint64(3*appconsts.ShareSize*types.DefaultGasPerBlobByte + paramLookUpCost), // 3 shares * 512 bytes per share * 8 gas per byte + 1060 gas for fetching param = 13348 gas
		},
		{
			name:            "3 blobs, 6 shares total",
			msg:             types.MsgPayForBlobs{BlobSizes: []uint32{1024, 1000, 100}},
			wantGasConsumed: uint64(6*appconsts.ShareSize*types.DefaultGasPerBlobByte + paramLookUpCost), // 6 shares * 512 bytes per share * 8 gas per byte + 1060 gas for fetching param = 25636 gas
		},
	}

	for _, tc := range testCases {
		k, stateStore := keeper(t)
		ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, nil)
		_, err := k.PayForBlobs(sdk.WrapSDKContext(ctx), &tc.msg)
		require.NoError(t, err)
		if tc.wantGasConsumed != ctx.GasMeter().GasConsumed() {
			t.Errorf("Gas consumed by %s: %d, want: %d", tc.name, ctx.GasMeter().GasConsumed(), tc.wantGasConsumed)
		}
	}
}

func TestChangingGasParam(t *testing.T) {
	msg := types.MsgPayForBlobs{BlobSizes: []uint32{1024}}
	k, stateStore := keeper(t)
	tempCtx := sdk.NewContext(stateStore, tmproto.Header{}, false, nil)

	ctx1 := sdk.NewContext(stateStore, tmproto.Header{}, false, nil)
	_, err := k.PayForBlobs(sdk.WrapSDKContext(ctx1), &msg)
	require.NoError(t, err)

	params := k.GetParams(tempCtx)
	params.GasPerBlobByte++
	k.SetParams(tempCtx, params)

	ctx2 := sdk.NewContext(stateStore, tmproto.Header{}, false, nil)
	_, err = k.PayForBlobs(sdk.WrapSDKContext(ctx2), &msg)
	require.NoError(t, err)

	if ctx1.GasMeter().GasConsumed() >= ctx2.GasMeter().GasConsumedToLimit() {
		t.Errorf("Gas consumed was not increased upon incrementing param, before: %d, after: %d",
			ctx1.GasMeter().GasConsumed(), ctx2.GasMeter().GasConsumedToLimit(),
		)
	}
}
