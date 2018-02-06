package consensus

import (
	"testing"

	"fmt"
	"math"

	"github.com/tclchiam/block_n_go/blockchain/engine/mining"
	"github.com/tclchiam/block_n_go/blockchain/entity"
)

func TestVerifyBlock(t *testing.T) {
	err := VerifyBlock(entity.DefaultGenesisBlock())

	//calculateBlockHashTestSuite(entity.DefaultGenesisBlock().Header())
	if err != nil {
		t.Errorf("verifying block: %s", err)
	}
}

func calculateBlockHashTestSuite(header *entity.BlockHeader) {
	input := &mining.BlockHashingInput{
		Index:            header.Index,
		PreviousHash:     header.PreviousHash,
		Timestamp:        header.Timestamp,
		TransactionsHash: header.TransactionsHash,
		Difficulty:       header.Difficulty,
	}

	for n := uint64(0); n < math.MaxUint64; n++ {
		result := mining.CalculateBlockHash(input, n)
		if mining.HasDifficulty(result, input.Difficulty) {
			fmt.Printf("nonce: %d, hash: %s\n", n, result)
			break
		}
	}
}
