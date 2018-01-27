// Code generated by protoc-gen-go. DO NOT EDIT.
// source: block/block.proto

/*
Package block is a generated protocol buffer package.

It is generated from these files:
	block/block.proto

It has these top-level messages:
	Signature
	Transaction
	Petition
	Block
	PublicKey
	PrivateKey
*/
package block

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// metadat about a block signature
type Signature struct {
	Timestamp  int64  `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Hash       []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	SignatureA []byte `protobuf:"bytes,3,opt,name=signatureA,proto3" json:"signatureA,omitempty"`
	SignatureB []byte `protobuf:"bytes,4,opt,name=signatureB,proto3" json:"signatureB,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Signature) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Signature) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Signature) GetSignatureA() []byte {
	if m != nil {
		return m.SignatureA
	}
	return nil
}

func (m *Signature) GetSignatureB() []byte {
	if m != nil {
		return m.SignatureB
	}
	return nil
}

type Transaction struct {
	Signature *Signature `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	Payload   []byte     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Transaction) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Transaction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Petition struct {
	Hash       []byte       `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Signatures []*Signature `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
}

func (m *Petition) Reset()                    { *m = Petition{} }
func (m *Petition) String() string            { return proto.CompactTextString(m) }
func (*Petition) ProtoMessage()               {}
func (*Petition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Petition) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Petition) GetSignatures() []*Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// Metadata about a block
type Block struct {
	Timestamp    int64          `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Parent       []byte         `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Petition     *Petition      `protobuf:"bytes,3,opt,name=petition" json:"petition,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,4,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Block) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Block) GetParent() []byte {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *Block) GetPetition() *Petition {
	if m != nil {
		return m.Petition
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type PublicKey struct {
	X []byte `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
	Y []byte `protobuf:"bytes,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PublicKey) GetX() []byte {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *PublicKey) GetY() []byte {
	if m != nil {
		return m.Y
	}
	return nil
}

type PrivateKey struct {
	Public *PublicKey `protobuf:"bytes,1,opt,name=public" json:"public,omitempty"`
	D      []byte     `protobuf:"bytes,2,opt,name=d,proto3" json:"d,omitempty"`
}

func (m *PrivateKey) Reset()                    { *m = PrivateKey{} }
func (m *PrivateKey) String() string            { return proto.CompactTextString(m) }
func (*PrivateKey) ProtoMessage()               {}
func (*PrivateKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PrivateKey) GetPublic() *PublicKey {
	if m != nil {
		return m.Public
	}
	return nil
}

func (m *PrivateKey) GetD() []byte {
	if m != nil {
		return m.D
	}
	return nil
}

func init() {
	proto.RegisterType((*Signature)(nil), "block.Signature")
	proto.RegisterType((*Transaction)(nil), "block.Transaction")
	proto.RegisterType((*Petition)(nil), "block.Petition")
	proto.RegisterType((*Block)(nil), "block.Block")
	proto.RegisterType((*PublicKey)(nil), "block.PublicKey")
	proto.RegisterType((*PrivateKey)(nil), "block.PrivateKey")
}

func init() { proto.RegisterFile("block/block.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xc9, 0xfe, 0xbd, 0xeb, 0xd9, 0xe0, 0xd5, 0x5c, 0x48, 0x10, 0xd1, 0x51, 0x11, 0x07,
	0x42, 0x2b, 0x13, 0xbc, 0xb7, 0x78, 0xe7, 0x4d, 0xa9, 0x82, 0xe0, 0x5d, 0xda, 0xc5, 0x36, 0xd8,
	0x36, 0xa5, 0x49, 0x65, 0xbd, 0xf0, 0xcb, 0xf8, 0x49, 0xa5, 0x59, 0xda, 0x75, 0x43, 0xf0, 0xa6,
	0xf4, 0x9c, 0xe7, 0x39, 0x27, 0xbf, 0x3c, 0x04, 0x8e, 0xc3, 0x54, 0x44, 0x1f, 0xae, 0xfe, 0x3a,
	0x45, 0x29, 0x94, 0xc0, 0x63, 0x5d, 0x9c, 0x5e, 0xc4, 0x42, 0xc4, 0x29, 0x73, 0x75, 0x33, 0xac,
	0xde, 0x5d, 0xc5, 0x33, 0x26, 0x15, 0xcd, 0x8a, 0xad, 0xcf, 0xfe, 0x02, 0xeb, 0x99, 0xc7, 0x39,
	0x55, 0x55, 0xc9, 0xf0, 0x19, 0x58, 0x9d, 0x4e, 0xd0, 0x02, 0x2d, 0x87, 0xc1, 0xae, 0x81, 0x31,
	0x8c, 0x12, 0x2a, 0x13, 0x32, 0x58, 0xa0, 0xe5, 0x3c, 0xd0, 0xff, 0xf8, 0x1c, 0x40, 0xb6, 0xe3,
	0x0f, 0x64, 0xa8, 0x95, 0x5e, 0x67, 0x4f, 0xf7, 0xc8, 0xe8, 0x40, 0xf7, 0xec, 0x57, 0x98, 0xbd,
	0x94, 0x34, 0x97, 0x34, 0x52, 0x5c, 0xe4, 0xd8, 0x01, 0xab, 0x13, 0x35, 0xc0, 0x6c, 0x75, 0xe4,
	0x6c, 0xaf, 0xd5, 0x51, 0x06, 0x3b, 0x0b, 0x26, 0xf0, 0xaf, 0xa0, 0x75, 0x2a, 0xe8, 0xda, 0x50,
	0xb5, 0xa5, 0xed, 0xc3, 0xd4, 0x67, 0x8a, 0xeb, 0xad, 0x2d, 0x38, 0xea, 0x81, 0xdf, 0xf6, 0xc0,
	0x24, 0x19, 0x2c, 0x86, 0xbf, 0x1e, 0xd5, 0xf3, 0xd8, 0xdf, 0x08, 0xc6, 0x5e, 0xa3, 0xff, 0x11,
	0xd3, 0x09, 0x4c, 0x0a, 0x5a, 0xb2, 0x5c, 0x19, 0x24, 0x53, 0xe1, 0x1b, 0x98, 0x16, 0x86, 0x48,
	0x07, 0x35, 0x5b, 0xfd, 0x37, 0xe7, 0xb5, 0xa0, 0x41, 0x67, 0xc0, 0xf7, 0x30, 0x57, 0xbb, 0x5c,
	0x24, 0x19, 0x69, 0x40, 0x6c, 0x06, 0x7a, 0x91, 0x05, 0x7b, 0x3e, 0xfb, 0x1a, 0x2c, 0xbf, 0x0a,
	0x53, 0x1e, 0x3d, 0xb1, 0x1a, 0xcf, 0x01, 0x6d, 0xcc, 0xa5, 0xd1, 0xa6, 0xa9, 0x6a, 0x83, 0x84,
	0x6a, 0xfb, 0x11, 0xc0, 0x2f, 0xf9, 0x27, 0x55, 0xac, 0x71, 0x2e, 0x61, 0x52, 0xe8, 0xb1, 0x83,
	0xd0, 0xbb, 0x5d, 0x81, 0xd1, 0x9b, 0x2d, 0x6d, 0xd6, 0x68, 0xed, 0x5d, 0xbd, 0x5d, 0xc6, 0x5c,
	0x25, 0x55, 0xe8, 0x44, 0x22, 0x73, 0x59, 0xca, 0x85, 0x4a, 0xd8, 0x9a, 0x65, 0x34, 0x77, 0x13,
	0x46, 0x55, 0xb2, 0x7d, 0x92, 0xe1, 0x44, 0xbf, 0xb5, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x22, 0xa9, 0xb8, 0xfb, 0xa8, 0x02, 0x00, 0x00,
}
