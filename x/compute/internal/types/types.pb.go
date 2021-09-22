// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: secret/compute/v1beta1/types.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_tendermint_tendermint_libs_bytes "github.com/tendermint/tendermint/libs/bytes"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AccessType int32

const (
	AccessTypeUndefined   AccessType = 0
	AccessTypeNobody      AccessType = 1
	AccessTypeOnlyAddress AccessType = 2
	AccessTypeEverybody   AccessType = 3
)

var AccessType_name = map[int32]string{
	0: "UNDEFINED",
	1: "NOBODY",
	2: "ONLY_ADDRESS",
	3: "EVERYBODY",
}

var AccessType_value = map[string]int32{
	"UNDEFINED":    0,
	"NOBODY":       1,
	"ONLY_ADDRESS": 2,
	"EVERYBODY":    3,
}

func (AccessType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ba7f40a6d1951b3, []int{0}
}

type AccessTypeParam struct {
	Value AccessType `protobuf:"varint,1,opt,name=value,proto3,enum=SecretNetwork.x.compute.v1beta1.AccessType" json:"value,omitempty" yaml:"value"`
}

func (m *AccessTypeParam) Reset()         { *m = AccessTypeParam{} }
func (m *AccessTypeParam) String() string { return proto.CompactTextString(m) }
func (*AccessTypeParam) ProtoMessage()    {}
func (*AccessTypeParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba7f40a6d1951b3, []int{0}
}
func (m *AccessTypeParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessTypeParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessTypeParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessTypeParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTypeParam.Merge(m, src)
}
func (m *AccessTypeParam) XXX_Size() int {
	return m.Size()
}
func (m *AccessTypeParam) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTypeParam.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTypeParam proto.InternalMessageInfo

// CodeInfo is data for the uploaded contract WASM code
type CodeInfo struct {
	CodeHash []byte                                        `protobuf:"bytes,1,opt,name=code_hash,json=codeHash,proto3" json:"code_hash,omitempty"`
	Creator  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Source   string                                        `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Builder  string                                        `protobuf:"bytes,4,opt,name=builder,proto3" json:"builder,omitempty"`
}

func (m *CodeInfo) Reset()         { *m = CodeInfo{} }
func (m *CodeInfo) String() string { return proto.CompactTextString(m) }
func (*CodeInfo) ProtoMessage()    {}
func (*CodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba7f40a6d1951b3, []int{1}
}
func (m *CodeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeInfo.Merge(m, src)
}
func (m *CodeInfo) XXX_Size() int {
	return m.Size()
}
func (m *CodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CodeInfo proto.InternalMessageInfo

// ContractInfo stores a WASM contract instance
type ContractInfo struct {
	CodeID  uint64                                        `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	Creator github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	//    bytes admin = 3 [(gogoproto.casttype) = "github.com/cosmos/cosmos-sdk/types.AccAddress"];
	Label string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	// never show this in query results, just use for sorting
	// (Note: when using json tag "-" amino refused to serialize it...)
	Created *AbsoluteTxPosition `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
}

func (m *ContractInfo) Reset()         { *m = ContractInfo{} }
func (m *ContractInfo) String() string { return proto.CompactTextString(m) }
func (*ContractInfo) ProtoMessage()    {}
func (*ContractInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba7f40a6d1951b3, []int{2}
}
func (m *ContractInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractInfo.Merge(m, src)
}
func (m *ContractInfo) XXX_Size() int {
	return m.Size()
}
func (m *ContractInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContractInfo proto.InternalMessageInfo

// AbsoluteTxPosition can be used to sort contracts
type AbsoluteTxPosition struct {
	// BlockHeight is the block the contract was created at
	BlockHeight int64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// TxIndex is a monotonic counter within the block (actual transaction index, or gas consumed)
	TxIndex uint64 `protobuf:"varint,2,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
}

func (m *AbsoluteTxPosition) Reset()         { *m = AbsoluteTxPosition{} }
func (m *AbsoluteTxPosition) String() string { return proto.CompactTextString(m) }
func (*AbsoluteTxPosition) ProtoMessage()    {}
func (*AbsoluteTxPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba7f40a6d1951b3, []int{3}
}
func (m *AbsoluteTxPosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AbsoluteTxPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AbsoluteTxPosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AbsoluteTxPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AbsoluteTxPosition.Merge(m, src)
}
func (m *AbsoluteTxPosition) XXX_Size() int {
	return m.Size()
}
func (m *AbsoluteTxPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_AbsoluteTxPosition.DiscardUnknown(m)
}

var xxx_messageInfo_AbsoluteTxPosition proto.InternalMessageInfo

// Model is a struct that holds a KV pair
type Model struct {
	// hex-encode key to read it better (this is often ascii)
	Key github_com_tendermint_tendermint_libs_bytes.HexBytes `protobuf:"bytes,1,opt,name=Key,proto3,casttype=github.com/tendermint/tendermint/libs/bytes.HexBytes" json:"Key,omitempty"`
	// base64-encode raw value
	Value []byte `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba7f40a6d1951b3, []int{4}
}
func (m *Model) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Model.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(m, src)
}
func (m *Model) XXX_Size() int {
	return m.Size()
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("SecretNetwork.x.compute.v1beta1.AccessType", AccessType_name, AccessType_value)
	proto.RegisterType((*AccessTypeParam)(nil), "SecretNetwork.x.compute.v1beta1.AccessTypeParam")
	proto.RegisterType((*CodeInfo)(nil), "SecretNetwork.x.compute.v1beta1.CodeInfo")
	proto.RegisterType((*ContractInfo)(nil), "SecretNetwork.x.compute.v1beta1.ContractInfo")
	proto.RegisterType((*AbsoluteTxPosition)(nil), "SecretNetwork.x.compute.v1beta1.AbsoluteTxPosition")
	proto.RegisterType((*Model)(nil), "SecretNetwork.x.compute.v1beta1.Model")
}

func init() {
	proto.RegisterFile("secret/compute/v1beta1/types.proto", fileDescriptor_8ba7f40a6d1951b3)
}

var fileDescriptor_8ba7f40a6d1951b3 = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0x8e, 0xc9, 0x2f, 0x43, 0x74, 0x6f, 0x34, 0x97, 0x7b, 0x6f, 0x48, 0x25, 0x27, 0x75, 0xa5,
	0x0a, 0x15, 0x11, 0x0b, 0xe8, 0xa2, 0xa2, 0xab, 0x98, 0xa4, 0x22, 0xa5, 0x24, 0xc8, 0x01, 0x24,
	0x2a, 0x55, 0x91, 0xed, 0x39, 0x24, 0x16, 0xb6, 0x27, 0xf2, 0x4c, 0xa8, 0xfd, 0x06, 0x15, 0xab,
	0x2e, 0xbb, 0x41, 0xaa, 0xd4, 0xaa, 0xe2, 0x05, 0xfa, 0x0e, 0x2c, 0x59, 0x76, 0x15, 0xb5, 0x41,
	0x7d, 0x81, 0x2e, 0x59, 0x55, 0x9e, 0x98, 0x86, 0xb6, 0x0b, 0x36, 0x5d, 0xf9, 0xfc, 0x7c, 0xfe,
	0xce, 0x77, 0x7e, 0x6c, 0xa4, 0x30, 0xb0, 0x7c, 0xe0, 0xaa, 0x45, 0xdd, 0xc1, 0x90, 0x83, 0x7a,
	0xbc, 0x62, 0x02, 0x37, 0x56, 0x54, 0x1e, 0x0e, 0x80, 0x55, 0x07, 0x3e, 0xe5, 0x14, 0x97, 0x3b,
	0x02, 0xd3, 0x02, 0xfe, 0x92, 0xfa, 0x47, 0xd5, 0xa0, 0x1a, 0x83, 0xab, 0x31, 0xb8, 0x34, 0xdf,
	0xa3, 0x3d, 0x2a, 0xb0, 0x6a, 0x64, 0x4d, 0x5e, 0x53, 0x1c, 0xf4, 0x77, 0xcd, 0xb2, 0x80, 0xb1,
	0xdd, 0x70, 0x00, 0x3b, 0x86, 0x6f, 0xb8, 0xb8, 0x83, 0xd2, 0xc7, 0x86, 0x33, 0x84, 0xa2, 0x54,
	0x91, 0x16, 0xff, 0x5a, 0x5d, 0xaa, 0xde, 0xc2, 0x5c, 0x9d, 0x12, 0x68, 0x85, 0x6f, 0xa3, 0x72,
	0x3e, 0x34, 0x5c, 0x67, 0x5d, 0x11, 0x1c, 0x8a, 0x3e, 0xe1, 0x5a, 0x4f, 0xbd, 0x79, 0x5b, 0x96,
	0x94, 0x0f, 0x12, 0xca, 0x6d, 0x50, 0x02, 0x4d, 0xef, 0x90, 0xe2, 0x3b, 0x68, 0xd6, 0xa2, 0x04,
	0xba, 0x7d, 0x83, 0xf5, 0x45, 0xad, 0xbc, 0x9e, 0x8b, 0x02, 0x9b, 0x06, 0xeb, 0xe3, 0x2d, 0x94,
	0xb5, 0x7c, 0x30, 0x38, 0xf5, 0x8b, 0x33, 0x51, 0x4a, 0x5b, 0xb9, 0x1a, 0x95, 0x97, 0x7b, 0x36,
	0xef, 0x0f, 0xcd, 0x48, 0x80, 0x6a, 0x51, 0xe6, 0x52, 0x16, 0x3f, 0x96, 0x19, 0x39, 0x8a, 0xa7,
	0x51, 0xb3, 0xac, 0x1a, 0x21, 0x3e, 0x30, 0xa6, 0x5f, 0x33, 0xe0, 0xff, 0x50, 0x86, 0xd1, 0xa1,
	0x6f, 0x41, 0x31, 0x59, 0x91, 0x16, 0x67, 0xf5, 0xd8, 0xc3, 0x45, 0x94, 0x35, 0x87, 0xb6, 0x43,
	0xc0, 0x2f, 0xa6, 0x44, 0xe2, 0xda, 0x55, 0xbe, 0x4a, 0x28, 0xbf, 0x41, 0x3d, 0xee, 0x1b, 0x16,
	0x17, 0x62, 0xef, 0xa1, 0xac, 0x10, 0x6b, 0x13, 0x21, 0x35, 0xa5, 0xa1, 0xf1, 0xa8, 0x9c, 0x11,
	0xbd, 0xd4, 0xf5, 0x4c, 0x94, 0x6a, 0x92, 0x3f, 0x2b, 0x7a, 0x1e, 0xa5, 0x1d, 0xc3, 0x04, 0x27,
	0x96, 0x36, 0x71, 0xf0, 0x76, 0x5c, 0x02, 0x48, 0x31, 0x5d, 0x91, 0x16, 0xe7, 0x56, 0xd7, 0x6e,
	0x5f, 0x8f, 0xc9, 0xa8, 0x33, 0xe4, 0xb0, 0x1b, 0xec, 0x50, 0x66, 0x73, 0x9b, 0x7a, 0xfa, 0x35,
	0x87, 0xa2, 0x23, 0xfc, 0x7b, 0x1a, 0xdf, 0x45, 0x79, 0xd3, 0xa1, 0xd6, 0x51, 0xb7, 0x0f, 0x76,
	0xaf, 0xcf, 0x45, 0xc7, 0x49, 0x7d, 0x4e, 0xc4, 0x36, 0x45, 0x08, 0x2f, 0xa0, 0x1c, 0x0f, 0xba,
	0xb6, 0x47, 0x20, 0x10, 0xbd, 0xa6, 0xf4, 0x2c, 0x0f, 0x9a, 0x91, 0xab, 0xd8, 0x28, 0xbd, 0x4d,
	0x09, 0x38, 0xf8, 0x29, 0x4a, 0x6e, 0x41, 0x38, 0x59, 0xad, 0xf6, 0xe8, 0x6a, 0x54, 0x7e, 0x78,
	0x63, 0x14, 0x1c, 0x3c, 0x02, 0xbe, 0x6b, 0x7b, 0xfc, 0xa6, 0xe9, 0xd8, 0x26, 0x53, 0xcd, 0x90,
	0x03, 0xab, 0x6e, 0x42, 0xa0, 0x45, 0x86, 0x1e, 0x91, 0x44, 0xd3, 0xd8, 0x17, 0x47, 0x29, 0x06,
	0xab, 0x4f, 0x9c, 0x07, 0x1f, 0x25, 0x84, 0xa6, 0xd7, 0x87, 0xef, 0xa3, 0xd9, 0xbd, 0x56, 0xbd,
	0xf1, 0xa4, 0xd9, 0x6a, 0xd4, 0x0b, 0x89, 0xd2, 0xff, 0x27, 0xa7, 0x95, 0x7f, 0xa6, 0xe9, 0x3d,
	0x8f, 0xc0, 0xa1, 0xed, 0x01, 0xc1, 0x15, 0x94, 0x69, 0xb5, 0xb5, 0x76, 0xfd, 0xa0, 0x20, 0x95,
	0xe6, 0x4f, 0x4e, 0x2b, 0x85, 0x29, 0xa8, 0x45, 0x4d, 0x4a, 0x42, 0xbc, 0x84, 0xf2, 0xed, 0xd6,
	0xb3, 0x83, 0x6e, 0xad, 0x5e, 0xd7, 0x1b, 0x9d, 0x4e, 0x61, 0xa6, 0xb4, 0x70, 0x72, 0x5a, 0xf9,
	0x77, 0x8a, 0x6b, 0x7b, 0x4e, 0x18, 0xaf, 0x2c, 0x2a, 0xdb, 0xd8, 0x6f, 0xe8, 0x07, 0x82, 0x31,
	0xf9, 0x6b, 0xd9, 0xc6, 0x31, 0xf8, 0x61, 0x44, 0x5a, 0xca, 0xbd, 0x7a, 0x27, 0x27, 0xce, 0xde,
	0xcb, 0x09, 0xed, 0xc5, 0xf9, 0x17, 0x39, 0x71, 0x36, 0x96, 0xa5, 0xf3, 0xb1, 0x2c, 0x5d, 0x8c,
	0x65, 0xe9, 0xf3, 0x58, 0x96, 0x5e, 0x5f, 0xca, 0x89, 0x8b, 0x4b, 0x39, 0xf1, 0xe9, 0x52, 0x4e,
	0x3c, 0x7f, 0x7c, 0x63, 0x54, 0xe0, 0xd9, 0x3d, 0xd7, 0x70, 0x07, 0x96, 0xfa, 0xd3, 0xaa, 0xd5,
	0xe0, 0xc7, 0x0f, 0xc1, 0xf6, 0x38, 0xf8, 0x9e, 0xe1, 0x4c, 0xce, 0xc9, 0xcc, 0x88, 0x6f, 0x7b,
	0xed, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x19, 0x4f, 0x33, 0x38, 0x04, 0x00, 0x00,
}

func (this *AccessTypeParam) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessTypeParam)
	if !ok {
		that2, ok := that.(AccessTypeParam)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *CodeInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CodeInfo)
	if !ok {
		that2, ok := that.(CodeInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.CodeHash, that1.CodeHash) {
		return false
	}
	if !bytes.Equal(this.Creator, that1.Creator) {
		return false
	}
	if this.Source != that1.Source {
		return false
	}
	if this.Builder != that1.Builder {
		return false
	}
	return true
}
func (this *ContractInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ContractInfo)
	if !ok {
		that2, ok := that.(ContractInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CodeID != that1.CodeID {
		return false
	}
	if !bytes.Equal(this.Creator, that1.Creator) {
		return false
	}
	if this.Label != that1.Label {
		return false
	}
	if !this.Created.Equal(that1.Created) {
		return false
	}
	return true
}
func (this *AbsoluteTxPosition) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AbsoluteTxPosition)
	if !ok {
		that2, ok := that.(AbsoluteTxPosition)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.BlockHeight != that1.BlockHeight {
		return false
	}
	if this.TxIndex != that1.TxIndex {
		return false
	}
	return true
}
func (this *Model) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Model)
	if !ok {
		that2, ok := that.(Model)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Key, that1.Key) {
		return false
	}
	if !bytes.Equal(this.Value, that1.Value) {
		return false
	}
	return true
}
func (m *AccessTypeParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessTypeParam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessTypeParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CodeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Builder) > 0 {
		i -= len(m.Builder)
		copy(dAtA[i:], m.Builder)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Builder)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Source) > 0 {
		i -= len(m.Source)
		copy(dAtA[i:], m.Source)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Source)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CodeHash) > 0 {
		i -= len(m.CodeHash)
		copy(dAtA[i:], m.CodeHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CodeHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContractInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Created != nil {
		{
			size, err := m.Created.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.CodeID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AbsoluteTxPosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AbsoluteTxPosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AbsoluteTxPosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TxIndex != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.TxIndex))
		i--
		dAtA[i] = 0x10
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Model) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Model) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Model) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccessTypeParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovTypes(uint64(m.Value))
	}
	return n
}

func (m *CodeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CodeHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Builder)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ContractInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeID != 0 {
		n += 1 + sovTypes(uint64(m.CodeID))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Created != nil {
		l = m.Created.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *AbsoluteTxPosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovTypes(uint64(m.BlockHeight))
	}
	if m.TxIndex != 0 {
		n += 1 + sovTypes(uint64(m.TxIndex))
	}
	return n
}

func (m *Model) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessTypeParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AccessTypeParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessTypeParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= AccessType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CodeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CodeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeHash = append(m.CodeHash[:0], dAtA[iNdEx:postIndex]...)
			if m.CodeHash == nil {
				m.CodeHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Builder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Builder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContractInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContractInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Created == nil {
				m.Created = &AbsoluteTxPosition{}
			}
			if err := m.Created.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AbsoluteTxPosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AbsoluteTxPosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AbsoluteTxPosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIndex", wireType)
			}
			m.TxIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Model) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Model: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Model: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)