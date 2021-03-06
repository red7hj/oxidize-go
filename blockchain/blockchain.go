package blockchain

import (
	"github.com/tclchiam/oxidize-go/blockchain/engine"
	"github.com/tclchiam/oxidize-go/blockchain/engine/iter"
	"github.com/tclchiam/oxidize-go/blockchain/engine/mining"
	"github.com/tclchiam/oxidize-go/blockchain/entity"
)

type Blockchain interface {
	entity.ChainRepository

	ForEachBlock(consume func(*entity.Block)) error

	Headers(hash *entity.Hash, index uint64) (entity.BlockHeaders, error)
	SaveHeaders(headers entity.BlockHeaders) error
	MineBlock(transactions entity.Transactions) (*entity.Block, error)
}

type blockchain struct {
	entity.ChainRepository
	miner mining.Miner
}

func Open(repository entity.ChainRepository, miner mining.Miner) (Blockchain, error) {
	err := engine.ResetGenesis(repository)
	if err != nil {
		return nil, err
	}

	return &blockchain{
		ChainRepository: repository,
		miner:           miner,
	}, nil
}

func (bc *blockchain) ForEachBlock(consume func(*entity.Block)) error {
	return iter.ForEachBlock(bc.ChainRepository, consume)
}

func (bc *blockchain) Headers(hash *entity.Hash, index uint64) (entity.BlockHeaders, error) {
	startingHeader, err := bc.HeaderByHash(hash)
	if err != nil {
		return nil, err
	}
	if startingHeader == nil {
		return entity.NewBlockHeaders(), nil
	}

	headers := entity.BlockHeaders{startingHeader}
	nextHeader := startingHeader
	for {
		nextHeader, err = bc.HeaderByIndex(nextHeader.Index + 1)
		if err != nil {
			return nil, err
		}
		if nextHeader == nil {
			return headers, nil
		}

		headers = headers.Add(nextHeader)
	}

	panic("unexpected")
}

func (bc *blockchain) SaveHeaders(headers entity.BlockHeaders) error {
	return engine.SaveHeaders(headers, bc)
}

func (bc *blockchain) SaveHeader(header *entity.BlockHeader) error {
	// TODO verify header
	return bc.ChainRepository.SaveHeader(header)
}

func (bc *blockchain) SaveBlock(block *entity.Block) error {
	// TODO verify block
	return bc.ChainRepository.SaveBlock(block)
}

func (bc *blockchain) MineBlock(transactions entity.Transactions) (*entity.Block, error) {
	return engine.MineBlock(transactions, bc.miner, bc.ChainRepository)
}
