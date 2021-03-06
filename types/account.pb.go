// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/account.proto

package types // import "github.com/zhigui-projects/zwasm/types"

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

type State struct {
	Nonce                uint64   `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Balance              uint64   `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	CodeHash             []byte   `protobuf:"bytes,3,opt,name=codeHash,proto3" json:"codeHash,omitempty"`
	StorageRoot          []byte   `protobuf:"bytes,4,opt,name=storageRoot,proto3" json:"storageRoot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_c2d8e4e7ea984bb1, []int{0}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *State) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *State) GetCodeHash() []byte {
	if m != nil {
		return m.CodeHash
	}
	return nil
}

func (m *State) GetStorageRoot() []byte {
	if m != nil {
		return m.StorageRoot
	}
	return nil
}

func init() {
	proto.RegisterType((*State)(nil), "types.State")
}

func init() { proto.RegisterFile("types/account.proto", fileDescriptor_account_c2d8e4e7ea984bb1) }

var fileDescriptor_account_c2d8e4e7ea984bb1 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xce, 0xb1, 0xae, 0x82, 0x40,
	0x10, 0x85, 0xe1, 0x70, 0x2f, 0xa8, 0x59, 0xad, 0x56, 0x8b, 0x8d, 0x15, 0xb1, 0x30, 0x34, 0xb2,
	0x85, 0x6f, 0x60, 0x65, 0x8d, 0x9d, 0xdd, 0x30, 0x4e, 0x00, 0x23, 0x0c, 0x61, 0x67, 0x63, 0xe4,
	0xe9, 0x8d, 0x6b, 0x34, 0x96, 0xff, 0x77, 0x9a, 0xa3, 0x96, 0xf2, 0xe8, 0xc9, 0x59, 0x40, 0x64,
	0xdf, 0x49, 0xde, 0x0f, 0x2c, 0xac, 0x93, 0x80, 0x1b, 0xaf, 0x92, 0x93, 0x80, 0x90, 0x5e, 0xa9,
	0xa4, 0xe3, 0x0e, 0xc9, 0x44, 0x69, 0x94, 0xc5, 0xc5, 0x3b, 0xb4, 0x51, 0xd3, 0x12, 0x6e, 0xf0,
	0xf2, 0xbf, 0xe0, 0x9f, 0xd4, 0x6b, 0x35, 0x43, 0xbe, 0xd0, 0x11, 0x5c, 0x6d, 0xfe, 0xd3, 0x28,
	0x5b, 0x14, 0xdf, 0xd6, 0xa9, 0x9a, 0x3b, 0xe1, 0x01, 0x2a, 0x2a, 0x98, 0xc5, 0xc4, 0x61, 0xfe,
	0xa5, 0x43, 0x76, 0xde, 0x56, 0x8d, 0xd4, 0xbe, 0xcc, 0x91, 0x5b, 0x3b, 0xd6, 0x4d, 0xe5, 0x9b,
	0x5d, 0x3f, 0xf0, 0x95, 0x50, 0x9c, 0x1d, 0xef, 0xe0, 0x5a, 0x1b, 0x0e, 0x96, 0x93, 0x70, 0x77,
	0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x94, 0x56, 0xc1, 0x72, 0xc5, 0x00, 0x00, 0x00,
}
