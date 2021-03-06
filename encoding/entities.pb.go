// Code generated by protoc-gen-go. DO NOT EDIT.
// source: encoding/entities.proto

/*
Package encoding is a generated protocol buffer package.

It is generated from these files:
	encoding/entities.proto

It has these top-level messages:
	Block
	BlockHeader
	Transaction
	Output
	OutputReference
	SignedInput
	UnsignedInput
*/
package encoding

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Block struct {
	Header           *BlockHeader   `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Transactions     []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type BlockHeader struct {
	Index            *uint64 `protobuf:"varint,1,req,name=index" json:"index,omitempty"`
	PreviousHash     []byte  `protobuf:"bytes,2,req,name=previousHash" json:"previousHash,omitempty"`
	Timestamp        *uint64 `protobuf:"varint,3,req,name=timestamp" json:"timestamp,omitempty"`
	TransactionsHash []byte  `protobuf:"bytes,4,req,name=transactionsHash" json:"transactionsHash,omitempty"`
	Nonce            *uint64 `protobuf:"varint,5,req,name=nonce" json:"nonce,omitempty"`
	Hash             []byte  `protobuf:"bytes,6,req,name=hash" json:"hash,omitempty"`
	Difficulty       *uint64 `protobuf:"varint,7,req,name=difficulty" json:"difficulty,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BlockHeader) GetIndex() uint64 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *BlockHeader) GetPreviousHash() []byte {
	if m != nil {
		return m.PreviousHash
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetTransactionsHash() []byte {
	if m != nil {
		return m.TransactionsHash
	}
	return nil
}

func (m *BlockHeader) GetNonce() uint64 {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return 0
}

func (m *BlockHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockHeader) GetDifficulty() uint64 {
	if m != nil && m.Difficulty != nil {
		return *m.Difficulty
	}
	return 0
}

type Transaction struct {
	Id               []byte         `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Inputs           []*SignedInput `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty"`
	Outputs          []*Output      `protobuf:"bytes,3,rep,name=outputs" json:"outputs,omitempty"`
	Secret           []byte         `protobuf:"bytes,4,req,name=secret" json:"secret,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Transaction) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Transaction) GetInputs() []*SignedInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Transaction) GetOutputs() []*Output {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Transaction) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

type Output struct {
	Index            *uint32 `protobuf:"varint,1,req,name=index" json:"index,omitempty"`
	Value            *uint64 `protobuf:"varint,2,req,name=value" json:"value,omitempty"`
	PublicKeyHash    []byte  `protobuf:"bytes,3,req,name=publicKeyHash" json:"publicKeyHash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Output) Reset()                    { *m = Output{} }
func (m *Output) String() string            { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()               {}
func (*Output) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Output) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *Output) GetValue() uint64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *Output) GetPublicKeyHash() []byte {
	if m != nil {
		return m.PublicKeyHash
	}
	return nil
}

type OutputReference struct {
	Id               []byte  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Output           *Output `protobuf:"bytes,2,req,name=output" json:"output,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OutputReference) Reset()                    { *m = OutputReference{} }
func (m *OutputReference) String() string            { return proto.CompactTextString(m) }
func (*OutputReference) ProtoMessage()               {}
func (*OutputReference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OutputReference) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OutputReference) GetOutput() *Output {
	if m != nil {
		return m.Output
	}
	return nil
}

type SignedInput struct {
	Reference        *OutputReference `protobuf:"bytes,1,req,name=reference" json:"reference,omitempty"`
	PublicKey        []byte           `protobuf:"bytes,2,req,name=publicKey" json:"publicKey,omitempty"`
	Signature        []byte           `protobuf:"bytes,3,req,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *SignedInput) Reset()                    { *m = SignedInput{} }
func (m *SignedInput) String() string            { return proto.CompactTextString(m) }
func (*SignedInput) ProtoMessage()               {}
func (*SignedInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SignedInput) GetReference() *OutputReference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *SignedInput) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *SignedInput) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type UnsignedInput struct {
	Reference        *OutputReference `protobuf:"bytes,1,req,name=reference" json:"reference,omitempty"`
	PublicKey        []byte           `protobuf:"bytes,2,req,name=publicKey" json:"publicKey,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *UnsignedInput) Reset()                    { *m = UnsignedInput{} }
func (m *UnsignedInput) String() string            { return proto.CompactTextString(m) }
func (*UnsignedInput) ProtoMessage()               {}
func (*UnsignedInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UnsignedInput) GetReference() *OutputReference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *UnsignedInput) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "entities.Block")
	proto.RegisterType((*BlockHeader)(nil), "entities.BlockHeader")
	proto.RegisterType((*Transaction)(nil), "entities.Transaction")
	proto.RegisterType((*Output)(nil), "entities.Output")
	proto.RegisterType((*OutputReference)(nil), "entities.OutputReference")
	proto.RegisterType((*SignedInput)(nil), "entities.SignedInput")
	proto.RegisterType((*UnsignedInput)(nil), "entities.UnsignedInput")
}

func init() { proto.RegisterFile("encoding/entities.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0xb5, 0xf9, 0xd7, 0x76, 0x92, 0x85, 0xca, 0x2a, 0x10, 0x24, 0x84, 0x56, 0x16, 0x87,
	0x55, 0x25, 0x8a, 0xc4, 0x05, 0x71, 0xed, 0xa9, 0xa8, 0x07, 0x24, 0x03, 0x97, 0xde, 0x42, 0x32,
	0xd9, 0xb5, 0x48, 0xed, 0x60, 0x3b, 0x15, 0xbd, 0xf2, 0x04, 0xbc, 0x1f, 0x2f, 0x83, 0x32, 0x71,
	0x9a, 0x84, 0xf6, 0xda, 0x5b, 0xe6, 0x9b, 0xdf, 0x8c, 0x67, 0x3e, 0x4d, 0xe0, 0x05, 0xaa, 0x52,
	0x57, 0x52, 0xed, 0xde, 0xa1, 0x72, 0xd2, 0x49, 0xb4, 0x67, 0xad, 0xd1, 0x4e, 0xb3, 0xc3, 0x31,
	0xe6, 0x3f, 0x21, 0x3e, 0x6f, 0x74, 0xf9, 0x83, 0xbd, 0x85, 0x64, 0x8f, 0x45, 0x85, 0x26, 0x5f,
	0x6d, 0x82, 0x6d, 0xfa, 0xfe, 0xd9, 0xd9, 0x5d, 0x0d, 0x01, 0x17, 0x94, 0x14, 0x1e, 0x62, 0x1f,
	0x21, 0x73, 0xa6, 0x50, 0xb6, 0x28, 0x9d, 0xd4, 0xca, 0xe6, 0xc1, 0x26, 0x5c, 0x16, 0x7d, 0x9d,
	0xb2, 0x62, 0x81, 0xf2, 0xbf, 0x2b, 0x48, 0x67, 0x2d, 0xd9, 0x09, 0xc4, 0x52, 0x55, 0xf8, 0x8b,
	0x1e, 0x8e, 0xc4, 0x10, 0x30, 0x0e, 0x59, 0x6b, 0xf0, 0x46, 0xea, 0xce, 0x5e, 0x14, 0x76, 0x9f,
	0x07, 0x9b, 0x60, 0x9b, 0x89, 0x85, 0xc6, 0x5e, 0xc1, 0x91, 0x93, 0xd7, 0x68, 0x5d, 0x71, 0xdd,
	0xe6, 0x21, 0x55, 0x4f, 0x02, 0x3b, 0x85, 0xe3, 0xf9, 0xbb, 0xd4, 0x25, 0xa2, 0x2e, 0xf7, 0xf4,
	0x7e, 0x06, 0xa5, 0x55, 0x89, 0x79, 0x3c, 0xcc, 0x40, 0x01, 0x63, 0x10, 0xed, 0xfb, 0xaa, 0x84,
	0xaa, 0xe8, 0x9b, 0xbd, 0x06, 0xa8, 0x64, 0x5d, 0xcb, 0xb2, 0x6b, 0xdc, 0x6d, 0x7e, 0x40, 0xf8,
	0x4c, 0xe1, 0x7f, 0x56, 0x90, 0xce, 0x76, 0x67, 0x4f, 0x20, 0x90, 0x15, 0xad, 0x96, 0x89, 0x40,
	0x56, 0xbd, 0xcf, 0x52, 0xb5, 0x9d, 0x7b, 0xc0, 0xb2, 0x2f, 0x72, 0xa7, 0xb0, 0xfa, 0xd4, 0x67,
	0x85, 0x87, 0xd8, 0x29, 0x1c, 0xe8, 0xce, 0x11, 0x1f, 0x12, 0x7f, 0x3c, 0xf1, 0x9f, 0x29, 0x21,
	0x46, 0x80, 0x3d, 0x87, 0xc4, 0x62, 0x69, 0xd0, 0xf9, 0x35, 0x7d, 0xc4, 0xaf, 0x20, 0x19, 0xd0,
	0xa5, 0xd5, 0xeb, 0xd1, 0xea, 0x13, 0x88, 0x6f, 0x8a, 0xa6, 0x43, 0xf2, 0x38, 0x12, 0x43, 0xc0,
	0xde, 0xc0, 0xba, 0xed, 0xbe, 0x37, 0xb2, 0xbc, 0xc4, 0x5b, 0xf2, 0x2e, 0xa4, 0xa6, 0x4b, 0x91,
	0x5f, 0xc2, 0x53, 0x3f, 0x06, 0xd6, 0x68, 0xb0, 0x77, 0xed, 0xff, 0x8d, 0xb7, 0x90, 0x0c, 0x13,
	0x52, 0xff, 0x87, 0x36, 0xf0, 0x79, 0xfe, 0x7b, 0x05, 0xe9, 0xcc, 0x04, 0xf6, 0x01, 0x8e, 0xcc,
	0xd8, 0xd6, 0x9f, 0xe5, 0xcb, 0x7b, 0xc5, 0x23, 0x20, 0x26, 0xb6, 0x3f, 0x8c, 0xbb, 0x31, 0xfd,
	0xe5, 0x4c, 0x42, 0x9f, 0xb5, 0x72, 0xa7, 0x0a, 0xd7, 0x19, 0xf4, 0x5b, 0x4d, 0x02, 0xaf, 0x61,
	0xfd, 0x4d, 0xd9, 0x47, 0x9f, 0xe2, 0x1c, 0xae, 0x0e, 0xc7, 0xdf, 0xf3, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x71, 0x3a, 0xe6, 0x58, 0xa9, 0x03, 0x00, 0x00,
}
