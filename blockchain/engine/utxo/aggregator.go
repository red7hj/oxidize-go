package utxo

import (
	"github.com/tclchiam/block_n_go/blockchain/entity"
	"github.com/tclchiam/block_n_go/identity"
)

type utxoAggregatorEngine struct {
}

func NewAggregatorEngine() Engine {
	return &utxoAggregatorEngine{}
}

func (utxoAggregatorEngine) FindUnspentOutputs(spender *identity.Identity) (*TransactionOutputSet, error) {
	panic("implement me")
}

func (utxoAggregatorEngine) IndexTransactions(transactions entity.Transactions) error {
	panic("implement me")
}
