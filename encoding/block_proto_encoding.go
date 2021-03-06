package encoding

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/tclchiam/oxidize-go/blockchain/entity"
)

type blockProtoEncoder struct{}

var blockProtoEncoderInstance blockProtoEncoder

func BlockProtoEncoder() entity.BlockEncoder {
	return &blockProtoEncoderInstance
}

func (*blockProtoEncoder) EncodeBlock(block *entity.Block) ([]byte, error) {
	message := ToWireBlock(block)

	data, err := proto.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("serializing block to protobuf: %s", err)
	}

	return data, nil
}

func (*blockProtoEncoder) DecodeBlock(input []byte) (*entity.Block, error) {
	message := &Block{}

	err := proto.Unmarshal(input, message)
	if err != nil {
		return nil, fmt.Errorf("deserializing block from protobuf '%s': %s", input, err)
	}

	return FromWireBlock(message)
}

func (*blockProtoEncoder) EncodeHeader(header *entity.BlockHeader) ([]byte, error) {
	message := ToWireBlockHeader(header)

	data, err := proto.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("serializing header to protobuf: %s", err)
	}

	return data, nil
}

func (*blockProtoEncoder) DecodeHeader(input []byte) (*entity.BlockHeader, error) {
	message := &BlockHeader{}

	err := proto.Unmarshal(input, message)
	if err != nil {
		return nil, fmt.Errorf("deserializing header from protobuf '%s': %s", input, err)
	}

	return FromWireBlockHeader(message)
}
