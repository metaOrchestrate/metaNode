package inclusion

import (
	"errors"

	"github.com/celestiaorg/celestia-app/pkg/da"
	"github.com/tendermint/tendermint/crypto/merkle"
)

// GetCommitment gets the share commitment for a blob in the original data
// square.
func GetCommitment(cacher *EDSSubTreeRootCacher, dah da.DataAvailabilityHeader, start, blobShareLen int) ([]byte, error) {
	squareSize := len(dah.RowsRoots) / 2
	if start+blobShareLen > squareSize*squareSize {
		return nil, errors.New("cannot get commitment for blob that doesn't fit in square")
	}
	paths := calculateCommitmentPaths(squareSize, start, blobShareLen)
	subTreeRoots := make([][]byte, len(paths))
	for i, path := range paths {
		// here we prepend false (walk left down the tree) because we only need
		// the subtree roots from the original data square.
		orignalSquarePath := append(append(make([]WalkInstruction, 0, len(path.instructions)+1), WalkLeft), path.instructions...)
		subTreeRoot, err := cacher.getSubTreeRoot(dah, path.row, orignalSquarePath)
		if err != nil {
			return nil, err
		}
		subTreeRoots[i] = subTreeRoot
	}
	return merkle.HashFromByteSlices(subTreeRoots), nil
}
