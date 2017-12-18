package blockchain

import (
	"github.com/tclchiam/block_n_go/tx"
	"github.com/imdario/mergo"
)

func (bc *Blockchain) ReadBalance(address string) (int, error) {
	unspentOutputs, err := bc.findUnspentOutputs(address)
	if err != nil {
		return -1, err
	}

	balance := calculateBalance(unspentOutputs)
	return balance, nil
}

func (bc *Blockchain) findUnspentOutputs(address string) (*tx.TransactionSet, error) {
	spentOutputs := make(map[string][]int)
	outputsForAddress := tx.NewTransactionSet()

	isUnspent := func(transactionId string, output *tx.Output) bool {
		if outputs, ok := spentOutputs[transactionId]; ok {
			for _, outputId := range outputs {
				if outputId == output.Id {
					return false
				}
			}
		}
		return true
	}

	err := bc.ForEachBlock(func(block *Block) {
		block.ForEachTransaction(func(transaction *tx.Transaction) {
			mergo.Map(&spentOutputs, transaction.FindSpentOutputs(address))
			outputsForAddress = outputsForAddress.Plus(transaction.FindOutputsForAddress(address))
		})
	})

	unspentOutputs := outputsForAddress.Filter(isUnspent)
	return unspentOutputs, err
}

func calculateBalance(unspentOutputs *tx.TransactionSet) int {
	sumBalance := func(res interface{}, _ string, output *tx.Output) interface{} {
		return res.(int) + output.Value
	}

	return unspentOutputs.Reduce(0, sumBalance).(int)
}