package entity

import (
	"testing"

	"github.com/tclchiam/block_n_go/blockchain/chainhash"
)

func TestNewGenesisBlockHeader(t *testing.T) {
	const address = "1DtTgiTBGmtLgawnm3ar5FE526FbKMKzCn"
	genesisBlock := NewGenesisBlockHeader(address)

	if !genesisBlock.PreviousHash.IsEqual(&chainhash.EmptyHash) {
		t.Fatalf("Genesis block has bad PreviousHash, expected [%v], but was [%v]", chainhash.EmptyHash, genesisBlock.PreviousHash)
	}
	if genesisBlock.Index != 0 {
		t.Fatalf("Genesis block has bad Index, expected %d, but was %d", 0, genesisBlock.Index)
	}
	if len(genesisBlock.Transactions) != 1 {
		t.Fatalf("Genesis block has bad Transactions, %s", genesisBlock.Transactions)
	}
}