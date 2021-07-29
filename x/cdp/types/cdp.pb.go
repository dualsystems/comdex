// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/cdp/v1alpha1/cdp.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CDP struct {
	Id              uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner           string                                 `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	Type            string                                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" yaml:"type"`
	Collateral      types.Coin                             `protobuf:"bytes,4,opt,name=collateral,proto3" json:"collateral" yaml:"collateral"`
	Principal       types.Coin                             `protobuf:"bytes,5,opt,name=principal,proto3" json:"principal" yaml:"principal"`
	AccumulatedFees types.Coin                             `protobuf:"bytes,6,opt,name=accumulated_fees,json=accumulatedFees,proto3" json:"accumulated_fees" yaml:"accumulatedFees"`
	FeesUpdated     time.Time                              `protobuf:"bytes,7,opt,name=fees_updated,json=feesUpdated,proto3,stdtime" json:"fees_updated" yaml:"feesUpdated"`
	InterestFactor  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=interest_factor,json=interestFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"interest_factor" yaml:"interestFactor"`
}

func (m *CDP) Reset()         { *m = CDP{} }
func (m *CDP) String() string { return proto.CompactTextString(m) }
func (*CDP) ProtoMessage()    {}
func (*CDP) Descriptor() ([]byte, []int) {
	return fileDescriptor_79abe14dd4273326, []int{0}
}
func (m *CDP) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CDP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CDP.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CDP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDP.Merge(m, src)
}
func (m *CDP) XXX_Size() int {
	return m.Size()
}
func (m *CDP) XXX_DiscardUnknown() {
	xxx_messageInfo_CDP.DiscardUnknown(m)
}

var xxx_messageInfo_CDP proto.InternalMessageInfo

func (m *CDP) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CDP) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CDP) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CDP) GetCollateral() types.Coin {
	if m != nil {
		return m.Collateral
	}
	return types.Coin{}
}

func (m *CDP) GetPrincipal() types.Coin {
	if m != nil {
		return m.Principal
	}
	return types.Coin{}
}

func (m *CDP) GetAccumulatedFees() types.Coin {
	if m != nil {
		return m.AccumulatedFees
	}
	return types.Coin{}
}

func (m *CDP) GetFeesUpdated() time.Time {
	if m != nil {
		return m.FeesUpdated
	}
	return time.Time{}
}

type AugmentedCDP struct {
	Cdp                    CDP        `protobuf:"bytes,1,opt,name=cdp,proto3" json:"cdp"`
	CollateralValue        types.Coin `protobuf:"bytes,2,opt,name=collateral_value,json=collateralValue,proto3" json:"collateral_value" yaml:"collateralValue"`
	CollateralizationRatio float64    `protobuf:"fixed64,3,opt,name=collateralization_ratio,json=collateralizationRatio,proto3" json:"collateralization_ratio,omitempty"`
}

func (m *AugmentedCDP) Reset()         { *m = AugmentedCDP{} }
func (m *AugmentedCDP) String() string { return proto.CompactTextString(m) }
func (*AugmentedCDP) ProtoMessage()    {}
func (*AugmentedCDP) Descriptor() ([]byte, []int) {
	return fileDescriptor_79abe14dd4273326, []int{1}
}
func (m *AugmentedCDP) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AugmentedCDP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AugmentedCDP.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AugmentedCDP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AugmentedCDP.Merge(m, src)
}
func (m *AugmentedCDP) XXX_Size() int {
	return m.Size()
}
func (m *AugmentedCDP) XXX_DiscardUnknown() {
	xxx_messageInfo_AugmentedCDP.DiscardUnknown(m)
}

var xxx_messageInfo_AugmentedCDP proto.InternalMessageInfo

func (m *AugmentedCDP) GetCdp() CDP {
	if m != nil {
		return m.Cdp
	}
	return CDP{}
}

func (m *AugmentedCDP) GetCollateralValue() types.Coin {
	if m != nil {
		return m.CollateralValue
	}
	return types.Coin{}
}

func (m *AugmentedCDP) GetCollateralizationRatio() float64 {
	if m != nil {
		return m.CollateralizationRatio
	}
	return 0
}

type Deposit struct {
	CdpId     uint64     `protobuf:"varint,1,opt,name=cdp_id,json=cdpId,proto3" json:"cdp_id,omitempty"`
	Depositor string     `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty" yaml:"depositor"`
	Amount    types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount" yaml:"amount"`
}

func (m *Deposit) Reset()         { *m = Deposit{} }
func (m *Deposit) String() string { return proto.CompactTextString(m) }
func (*Deposit) ProtoMessage()    {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_79abe14dd4273326, []int{2}
}
func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return m.Size()
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo

func (m *Deposit) GetCdpId() uint64 {
	if m != nil {
		return m.CdpId
	}
	return 0
}

func (m *Deposit) GetDepositor() string {
	if m != nil {
		return m.Depositor
	}
	return ""
}

func (m *Deposit) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*CDP)(nil), "comdex.cdp.v1alpha1.CDP")
	proto.RegisterType((*AugmentedCDP)(nil), "comdex.cdp.v1alpha1.AugmentedCDP")
	proto.RegisterType((*Deposit)(nil), "comdex.cdp.v1alpha1.Deposit")
}

func init() { proto.RegisterFile("comdex/cdp/v1alpha1/cdp.proto", fileDescriptor_79abe14dd4273326) }

var fileDescriptor_79abe14dd4273326 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x4e, 0xd4, 0x40,
	0x18, 0xde, 0xc2, 0xb2, 0xc8, 0x2c, 0xb2, 0x38, 0x02, 0x96, 0x4d, 0x6c, 0x49, 0x4d, 0x08, 0x17,
	0xa6, 0x82, 0x07, 0x13, 0x6f, 0x2e, 0x1b, 0xd0, 0x1b, 0x4e, 0xd0, 0x83, 0x89, 0xd9, 0xcc, 0xce,
	0xcc, 0x2e, 0x13, 0xdb, 0x4e, 0xd3, 0x4e, 0x51, 0x7c, 0x0a, 0x9e, 0xc0, 0x87, 0xf1, 0xc4, 0x91,
	0xa3, 0xf1, 0xb0, 0x18, 0x78, 0x83, 0x7d, 0x02, 0x33, 0x33, 0xdd, 0x2d, 0xa2, 0x91, 0x78, 0xd9,
	0xee, 0x7c, 0xdf, 0xff, 0x7f, 0xff, 0xdf, 0x7e, 0x5f, 0x0b, 0x1e, 0x53, 0x19, 0x33, 0xfe, 0x39,
	0xa4, 0x2c, 0x0d, 0x4f, 0x76, 0x48, 0x94, 0x1e, 0x93, 0x1d, 0x7d, 0x40, 0x69, 0x26, 0x95, 0x84,
	0x0f, 0x2d, 0x8d, 0x34, 0x32, 0xa1, 0xdb, 0x2b, 0x43, 0x39, 0x94, 0x86, 0x0f, 0xf5, 0x3f, 0x5b,
	0xda, 0xf6, 0x87, 0x52, 0x0e, 0x23, 0x1e, 0x9a, 0x53, 0xbf, 0x18, 0x84, 0x4a, 0xc4, 0x3c, 0x57,
	0x24, 0x2e, 0xb5, 0xda, 0x1e, 0x95, 0x79, 0x2c, 0xf3, 0xb0, 0x4f, 0x72, 0x1e, 0x9e, 0xec, 0xf4,
	0xb9, 0xd2, 0xa3, 0xa4, 0x48, 0x2c, 0x1f, 0x7c, 0xab, 0x83, 0xd9, 0xbd, 0xee, 0x21, 0x5c, 0x02,
	0x33, 0x82, 0xb9, 0xce, 0x86, 0xb3, 0x55, 0xc7, 0x33, 0x82, 0xc1, 0x4d, 0x30, 0x27, 0x3f, 0x25,
	0x3c, 0x73, 0x67, 0x36, 0x9c, 0xad, 0x85, 0xce, 0xf2, 0x78, 0xe4, 0x2f, 0x9e, 0x92, 0x38, 0x7a,
	0x11, 0x18, 0x38, 0xc0, 0x96, 0x86, 0x4f, 0x40, 0x5d, 0x9d, 0xa6, 0xdc, 0x9d, 0x35, 0x65, 0xad,
	0xf1, 0xc8, 0x6f, 0xda, 0x32, 0x8d, 0x06, 0xd8, 0x90, 0xf0, 0x08, 0x00, 0x2a, 0xa3, 0x88, 0x28,
	0x9e, 0x91, 0xc8, 0xad, 0x6f, 0x38, 0x5b, 0xcd, 0xdd, 0x75, 0x64, 0x37, 0x43, 0x7a, 0x33, 0x54,
	0x6e, 0x86, 0xf6, 0xa4, 0x48, 0x3a, 0xeb, 0xe7, 0x23, 0xbf, 0x36, 0x1e, 0xf9, 0x0f, 0xac, 0x52,
	0xd5, 0x1a, 0xe0, 0x1b, 0x3a, 0xf0, 0x0d, 0x58, 0x48, 0x33, 0x91, 0x50, 0x91, 0x92, 0xc8, 0x9d,
	0xbb, 0x4b, 0xd4, 0x2d, 0x45, 0x97, 0xad, 0xe8, 0xb4, 0x33, 0xc0, 0x95, 0x0a, 0x64, 0x60, 0x99,
	0x50, 0x5a, 0xc4, 0x85, 0x9e, 0xc1, 0x7a, 0x03, 0xce, 0x73, 0xb7, 0x71, 0x97, 0xb2, 0x57, 0x2a,
	0xaf, 0x59, 0xe5, 0x1b, 0x02, 0xfb, 0x9c, 0xe7, 0x01, 0x6e, 0xdd, 0x42, 0xe0, 0x07, 0xb0, 0xa8,
	0x95, 0x7b, 0x45, 0xca, 0x34, 0xe6, 0xce, 0x9b, 0x09, 0x6d, 0x64, 0xbd, 0x44, 0x13, 0x2f, 0xd1,
	0xd1, 0xc4, 0xcb, 0xe9, 0x08, 0x68, 0x47, 0xe8, 0xee, 0xb7, 0xb6, 0x39, 0x38, 0xbb, 0xf4, 0x1d,
	0xdc, 0xbc, 0x81, 0xc0, 0x14, 0xb4, 0x44, 0xa2, 0x78, 0xc6, 0x73, 0xd5, 0x1b, 0x10, 0xaa, 0x64,
	0xe6, 0xde, 0x33, 0xee, 0x1c, 0x68, 0x95, 0x1f, 0x23, 0x7f, 0x73, 0x28, 0xd4, 0x71, 0xd1, 0x47,
	0x54, 0xc6, 0x61, 0x19, 0x0f, 0x7b, 0xd9, 0xce, 0xd9, 0xc7, 0x50, 0x3b, 0x96, 0xa3, 0x2e, 0xa7,
	0xe3, 0x91, 0xbf, 0x6a, 0xe7, 0x4d, 0xe4, 0xf6, 0x8d, 0x5a, 0x80, 0x97, 0x6e, 0x01, 0x97, 0x0e,
	0x58, 0x7c, 0x59, 0x0c, 0x63, 0x9e, 0x28, 0xce, 0x74, 0x9a, 0x9e, 0x82, 0x59, 0xca, 0x52, 0x13,
	0xa7, 0xe6, 0xae, 0x8b, 0xfe, 0x92, 0x67, 0xb4, 0xd7, 0x3d, 0xec, 0xd4, 0xf5, 0x42, 0x58, 0x97,
	0xea, 0x27, 0x5f, 0x59, 0xdb, 0x3b, 0x21, 0x51, 0xc1, 0x4d, 0xf4, 0xfe, 0xe7, 0xc9, 0x57, 0x02,
	0xef, 0x74, 0x7f, 0x80, 0x5b, 0xb7, 0x10, 0xf8, 0x1c, 0x3c, 0xaa, 0x20, 0xf1, 0x85, 0x28, 0x21,
	0x93, 0x5e, 0xa6, 0x2f, 0x26, 0xc0, 0x0e, 0x5e, 0xfb, 0x83, 0xc6, 0xfa, 0x37, 0xf8, 0xea, 0x80,
	0xf9, 0x2e, 0x4f, 0x65, 0x2e, 0x14, 0x5c, 0x05, 0x0d, 0xca, 0xd2, 0xde, 0xf4, 0x75, 0x99, 0xa3,
	0x2c, 0x7d, 0xcd, 0xe0, 0x2e, 0x58, 0x60, 0xb6, 0x42, 0x4e, 0xde, 0x9a, 0x95, 0x2a, 0x6f, 0x53,
	0x2a, 0xc0, 0x55, 0x19, 0x7c, 0x05, 0x1a, 0x24, 0x96, 0x45, 0xa2, 0xcc, 0xf8, 0x7f, 0xde, 0xeb,
	0x6a, 0x79, 0xaf, 0xf7, 0xcb, 0x94, 0x99, 0xb6, 0x00, 0x97, 0xfd, 0x9d, 0x83, 0xf3, 0x2b, 0xcf,
	0xb9, 0xb8, 0xf2, 0x9c, 0x9f, 0x57, 0x9e, 0x73, 0x76, 0xed, 0xd5, 0x2e, 0xae, 0xbd, 0xda, 0xf7,
	0x6b, 0xaf, 0xf6, 0x7e, 0xfb, 0x37, 0xb7, 0xb5, 0x11, 0xdb, 0x72, 0x30, 0x10, 0x54, 0x90, 0xa8,
	0x3c, 0x87, 0xf6, 0x4b, 0x64, 0x8c, 0xef, 0x37, 0x4c, 0xfc, 0x9e, 0xfd, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x2e, 0x88, 0x9b, 0xd7, 0xa4, 0x04, 0x00, 0x00,
}

func (m *CDP) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CDP) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CDP) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.InterestFactor.Size()
		i -= size
		if _, err := m.InterestFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.FeesUpdated, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.FeesUpdated):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCdp(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.AccumulatedFees.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Principal.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Collateral.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintCdp(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCdp(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintCdp(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AugmentedCDP) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AugmentedCDP) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AugmentedCDP) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CollateralizationRatio != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.CollateralizationRatio))))
		i--
		dAtA[i] = 0x19
	}
	{
		size, err := m.CollateralValue.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Cdp.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Deposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Deposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCdp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintCdp(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if m.CdpId != 0 {
		i = encodeVarintCdp(dAtA, i, uint64(m.CdpId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCdp(dAtA []byte, offset int, v uint64) int {
	offset -= sovCdp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CDP) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCdp(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCdp(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovCdp(uint64(l))
	}
	l = m.Collateral.Size()
	n += 1 + l + sovCdp(uint64(l))
	l = m.Principal.Size()
	n += 1 + l + sovCdp(uint64(l))
	l = m.AccumulatedFees.Size()
	n += 1 + l + sovCdp(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.FeesUpdated)
	n += 1 + l + sovCdp(uint64(l))
	l = m.InterestFactor.Size()
	n += 1 + l + sovCdp(uint64(l))
	return n
}

func (m *AugmentedCDP) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Cdp.Size()
	n += 1 + l + sovCdp(uint64(l))
	l = m.CollateralValue.Size()
	n += 1 + l + sovCdp(uint64(l))
	if m.CollateralizationRatio != 0 {
		n += 9
	}
	return n
}

func (m *Deposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CdpId != 0 {
		n += 1 + sovCdp(uint64(m.CdpId))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovCdp(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovCdp(uint64(l))
	return n
}

func sovCdp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCdp(x uint64) (n int) {
	return sovCdp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CDP) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCdp
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
			return fmt.Errorf("proto: CDP: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CDP: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
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
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
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
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
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
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Collateral.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Principal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
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
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Principal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccumulatedFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
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
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AccumulatedFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeesUpdated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
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
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.FeesUpdated, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterestFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
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
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InterestFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCdp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCdp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCdp
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
func (m *AugmentedCDP) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCdp
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
			return fmt.Errorf("proto: AugmentedCDP: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AugmentedCDP: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cdp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
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
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cdp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
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
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralizationRatio", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.CollateralizationRatio = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipCdp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCdp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCdp
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
func (m *Deposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCdp
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
			return fmt.Errorf("proto: Deposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CdpId", wireType)
			}
			m.CdpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CdpId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
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
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCdp
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
				return ErrInvalidLengthCdp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCdp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCdp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCdp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCdp
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
func skipCdp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCdp
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
					return 0, ErrIntOverflowCdp
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
					return 0, ErrIntOverflowCdp
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
				return 0, ErrInvalidLengthCdp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCdp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCdp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCdp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCdp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCdp = fmt.Errorf("proto: unexpected end of group")
)