package encoding

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	coretypes "github.com/tendermint/tendermint/types"
)

func IndexWrapperDecoder(dec sdk.TxDecoder) sdk.TxDecoder {
	return func(txBytes []byte) (sdk.Tx, error) {
		if indexWrapper, isIndexWrapper := coretypes.UnmarshalIndexWrapper(txBytes); isIndexWrapper {
			return dec(indexWrapper.Tx)
		}
		return dec(txBytes)
	}
}
