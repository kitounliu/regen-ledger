// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/v1alpha1/genesis.proto

package ecocredit

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines ecocredit module's genesis state.
type GenesisState struct {
	// Params contains the updateable global parameters for use with the x/params
	// module
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// class_infos is the list of credit class info.
	ClassInfos []*ClassInfo `protobuf:"bytes,2,rep,name=class_infos,json=classInfos,proto3" json:"class_infos,omitempty"`
	// batch_infos is the list of credit batch info.
	BatchInfos []*BatchInfo `protobuf:"bytes,3,rep,name=batch_infos,json=batchInfos,proto3" json:"batch_infos,omitempty"`
	// id_seq is used to get next class/batch id.
	IdSeq uint64 `protobuf:"varint,4,opt,name=id_seq,json=idSeq,proto3" json:"id_seq,omitempty"`
	// tradable_balances is the list of credit batch tradable units.
	TradableBalances []*Balance `protobuf:"bytes,5,rep,name=tradable_balances,json=tradableBalances,proto3" json:"tradable_balances,omitempty"`
	// retired_balances is the list of credit batch retired units.
	RetiredBalances []*Balance `protobuf:"bytes,6,rep,name=retired_balances,json=retiredBalances,proto3" json:"retired_balances,omitempty"`
	// tradable_supplies is the list of credit batch tradable supply.
	TradableSupplies []*Supply `protobuf:"bytes,7,rep,name=tradable_supplies,json=tradableSupplies,proto3" json:"tradable_supplies,omitempty"`
	// retired_supplies is the list of credit batch retired supply.
	RetiredSupplies []*Supply `protobuf:"bytes,8,rep,name=retired_supplies,json=retiredSupplies,proto3" json:"retired_supplies,omitempty"`
	// precisions is the list of decimal precision of a credit batch.
	Precisions []*Precision `protobuf:"bytes,9,rep,name=precisions,proto3" json:"precisions,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9cb84fe1853321, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetClassInfos() []*ClassInfo {
	if m != nil {
		return m.ClassInfos
	}
	return nil
}

func (m *GenesisState) GetBatchInfos() []*BatchInfo {
	if m != nil {
		return m.BatchInfos
	}
	return nil
}

func (m *GenesisState) GetIdSeq() uint64 {
	if m != nil {
		return m.IdSeq
	}
	return 0
}

func (m *GenesisState) GetTradableBalances() []*Balance {
	if m != nil {
		return m.TradableBalances
	}
	return nil
}

func (m *GenesisState) GetRetiredBalances() []*Balance {
	if m != nil {
		return m.RetiredBalances
	}
	return nil
}

func (m *GenesisState) GetTradableSupplies() []*Supply {
	if m != nil {
		return m.TradableSupplies
	}
	return nil
}

func (m *GenesisState) GetRetiredSupplies() []*Supply {
	if m != nil {
		return m.RetiredSupplies
	}
	return nil
}

func (m *GenesisState) GetPrecisions() []*Precision {
	if m != nil {
		return m.Precisions
	}
	return nil
}

// Precision represents a credit batch precision with a batch_denom and max_decimal_places.
type Precision struct {
	// batch_denom is the unique ID of the credit batch.
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// max_decimal_places is the new maximum number of decimal places that can be
	// used to represent some quantity of credit units. It is an experimental
	// feature to concretely explore an idea proposed in
	// https://github.com/cosmos/cosmos-sdk/issues/7113.
	MaxDecimalPlaces uint32 `protobuf:"varint,3,opt,name=max_decimal_places,json=maxDecimalPlaces,proto3" json:"max_decimal_places,omitempty"`
}

func (m *Precision) Reset()         { *m = Precision{} }
func (m *Precision) String() string { return proto.CompactTextString(m) }
func (*Precision) ProtoMessage()    {}
func (*Precision) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9cb84fe1853321, []int{1}
}
func (m *Precision) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Precision) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Precision.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Precision) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Precision.Merge(m, src)
}
func (m *Precision) XXX_Size() int {
	return m.Size()
}
func (m *Precision) XXX_DiscardUnknown() {
	xxx_messageInfo_Precision.DiscardUnknown(m)
}

var xxx_messageInfo_Precision proto.InternalMessageInfo

func (m *Precision) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *Precision) GetMaxDecimalPlaces() uint32 {
	if m != nil {
		return m.MaxDecimalPlaces
	}
	return 0
}

// Balance represents tradable or retired units of a credit batch with an account address,
// batch_denom, and balance.
type Balance struct {
	// address is the account address of the account holding credits.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// batch_denom is the unique ID of the credit batch.
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// balance is the tradable or retired balance of the credit batch.
	Balance string `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9cb84fe1853321, []int{2}
}
func (m *Balance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(m, src)
}
func (m *Balance) XXX_Size() int {
	return m.Size()
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func (m *Balance) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Balance) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *Balance) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

// Supply represents a tradable or retired supply of a credit batch.
type Supply struct {
	// batch_denom is the unique ID of the credit batch.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// supply is the tradable or retired supply of the credit batch.
	Supply string `protobuf:"bytes,2,opt,name=supply,proto3" json:"supply,omitempty"`
}

func (m *Supply) Reset()         { *m = Supply{} }
func (m *Supply) String() string { return proto.CompactTextString(m) }
func (*Supply) ProtoMessage()    {}
func (*Supply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9cb84fe1853321, []int{3}
}
func (m *Supply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Supply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Supply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Supply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supply.Merge(m, src)
}
func (m *Supply) XXX_Size() int {
	return m.Size()
}
func (m *Supply) XXX_DiscardUnknown() {
	xxx_messageInfo_Supply.DiscardUnknown(m)
}

var xxx_messageInfo_Supply proto.InternalMessageInfo

func (m *Supply) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *Supply) GetSupply() string {
	if m != nil {
		return m.Supply
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "regen.ecocredit.v1alpha1.GenesisState")
	proto.RegisterType((*Precision)(nil), "regen.ecocredit.v1alpha1.Precision")
	proto.RegisterType((*Balance)(nil), "regen.ecocredit.v1alpha1.Balance")
	proto.RegisterType((*Supply)(nil), "regen.ecocredit.v1alpha1.Supply")
}

func init() {
	proto.RegisterFile("regen/ecocredit/v1alpha1/genesis.proto", fileDescriptor_2f9cb84fe1853321)
}

var fileDescriptor_2f9cb84fe1853321 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe3, 0x36, 0x75, 0xc8, 0x06, 0x44, 0x58, 0x01, 0xb2, 0x7a, 0x70, 0x4d, 0x40, 0x28,
	0x07, 0xb0, 0x95, 0x72, 0x47, 0x22, 0x8d, 0x84, 0x50, 0x01, 0x45, 0xce, 0xad, 0x07, 0xa2, 0xb5,
	0x77, 0xea, 0xac, 0xb0, 0xbd, 0xee, 0xee, 0x16, 0x92, 0xb7, 0xe0, 0xb1, 0x7a, 0xec, 0x91, 0x13,
	0x42, 0xc9, 0x0b, 0xf0, 0x08, 0xc8, 0xeb, 0x75, 0x52, 0x81, 0xd2, 0xf4, 0x36, 0x33, 0xfe, 0xe7,
	0xfb, 0x47, 0x33, 0x5e, 0xf4, 0x52, 0x40, 0x02, 0x79, 0x00, 0x31, 0x8f, 0x05, 0x50, 0xa6, 0x82,
	0x6f, 0x03, 0x92, 0x16, 0x33, 0x32, 0x08, 0x12, 0xc8, 0x41, 0x32, 0xe9, 0x17, 0x82, 0x2b, 0x8e,
	0x1d, 0xad, 0xf3, 0xd7, 0x3a, 0xbf, 0xd6, 0x1d, 0xbe, 0xd8, 0x4a, 0x50, 0x8b, 0x02, 0x4c, 0xff,
	0xe1, 0xe3, 0x84, 0x27, 0x5c, 0x87, 0x41, 0x19, 0x55, 0xd5, 0xde, 0x9f, 0x26, 0xba, 0xff, 0xbe,
	0xf2, 0x99, 0x28, 0xa2, 0x00, 0xbf, 0x45, 0x76, 0x41, 0x04, 0xc9, 0xa4, 0x63, 0x79, 0x56, 0xbf,
	0x73, 0xec, 0xf9, 0xdb, 0x7c, 0xfd, 0xb1, 0xd6, 0x0d, 0x9b, 0x57, 0xbf, 0x8e, 0x1a, 0xa1, 0xe9,
	0xc2, 0x23, 0xd4, 0x89, 0x53, 0x22, 0xe5, 0x94, 0xe5, 0xe7, 0x5c, 0x3a, 0x7b, 0xde, 0x7e, 0xbf,
	0x73, 0xfc, 0x7c, 0x3b, 0xe4, 0xa4, 0x14, 0x7f, 0xc8, 0xcf, 0x79, 0x88, 0xe2, 0x3a, 0xd4, 0x94,
	0x88, 0xa8, 0x78, 0x66, 0x28, 0xfb, 0xbb, 0x28, 0xc3, 0x52, 0x5c, 0x51, 0xa2, 0x3a, 0x94, 0xf8,
	0x09, 0xb2, 0x19, 0x9d, 0x4a, 0xb8, 0x70, 0x9a, 0x9e, 0xd5, 0x6f, 0x86, 0x07, 0x8c, 0x4e, 0xe0,
	0x02, 0x7f, 0x46, 0x8f, 0x94, 0x20, 0x94, 0x44, 0x29, 0x4c, 0x23, 0x92, 0x92, 0x3c, 0x06, 0xe9,
	0x1c, 0x68, 0x8b, 0x67, 0xb7, 0x59, 0x68, 0x65, 0xd8, 0xad, 0x7b, 0x4d, 0x41, 0xe2, 0x8f, 0xa8,
	0x2b, 0x40, 0x31, 0x01, 0x74, 0x83, 0xb3, 0xef, 0x8a, 0x7b, 0x68, 0x5a, 0xd7, 0xb4, 0x4f, 0x37,
	0xa6, 0x93, 0x97, 0x45, 0x91, 0x32, 0x90, 0x4e, 0x4b, 0xe3, 0x6e, 0xb9, 0xc5, 0xa4, 0x54, 0x2e,
	0x36, 0xc3, 0x4d, 0x4c, 0x27, 0x3e, 0xdd, 0x0c, 0xb7, 0xa6, 0xdd, 0xbb, 0x23, 0xad, 0x9e, 0x6d,
	0x0d, 0x3b, 0x41, 0xa8, 0x10, 0x10, 0x33, 0xc9, 0x78, 0x2e, 0x9d, 0xf6, 0xae, 0xab, 0x8c, 0x6b,
	0x6d, 0x78, 0xa3, 0xad, 0x77, 0x86, 0xda, 0xeb, 0x0f, 0xf8, 0xa8, 0x3e, 0x34, 0x85, 0x9c, 0x67,
	0xce, 0x9e, 0x67, 0xf5, 0xdb, 0xe6, 0x86, 0xa3, 0xb2, 0x82, 0x5f, 0x21, 0x9c, 0x91, 0xf9, 0x94,
	0x42, 0xcc, 0x32, 0x92, 0x4e, 0x8b, 0x94, 0x94, 0xeb, 0xdd, 0xf7, 0xac, 0xfe, 0x83, 0xb0, 0x9b,
	0x91, 0xf9, 0xa8, 0xfa, 0x30, 0xd6, 0xf5, 0xde, 0x17, 0xd4, 0x32, 0x8b, 0xc4, 0x0e, 0x6a, 0x11,
	0x4a, 0x05, 0xc8, 0xea, 0x4f, 0x6e, 0x87, 0x75, 0xba, 0xdb, 0xd3, 0x41, 0x2d, 0x73, 0x48, 0x6d,
	0xd4, 0x0e, 0xeb, 0xb4, 0xf7, 0x0e, 0xd9, 0xd5, 0x6e, 0xfe, 0x85, 0x58, 0xff, 0x41, 0x9e, 0x22,
	0x5b, 0x2f, 0x7c, 0x61, 0x0c, 0x4c, 0x36, 0x3c, 0xbd, 0x5a, 0xba, 0xd6, 0xf5, 0xd2, 0xb5, 0x7e,
	0x2f, 0x5d, 0xeb, 0xc7, 0xca, 0x6d, 0x5c, 0xaf, 0xdc, 0xc6, 0xcf, 0x95, 0xdb, 0x38, 0x1b, 0x24,
	0x4c, 0xcd, 0x2e, 0x23, 0x3f, 0xe6, 0x59, 0xa0, 0x77, 0xfa, 0x3a, 0x07, 0xf5, 0x9d, 0x8b, 0xaf,
	0x26, 0x4b, 0x81, 0x26, 0x20, 0x82, 0xf9, 0xe6, 0xa5, 0x47, 0xb6, 0x7e, 0xc5, 0x6f, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xbe, 0xd8, 0x41, 0xbb, 0x45, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Precisions) > 0 {
		for iNdEx := len(m.Precisions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Precisions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.RetiredSupplies) > 0 {
		for iNdEx := len(m.RetiredSupplies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RetiredSupplies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.TradableSupplies) > 0 {
		for iNdEx := len(m.TradableSupplies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TradableSupplies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.RetiredBalances) > 0 {
		for iNdEx := len(m.RetiredBalances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RetiredBalances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.TradableBalances) > 0 {
		for iNdEx := len(m.TradableBalances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TradableBalances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.IdSeq != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.IdSeq))
		i--
		dAtA[i] = 0x20
	}
	if len(m.BatchInfos) > 0 {
		for iNdEx := len(m.BatchInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BatchInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ClassInfos) > 0 {
		for iNdEx := len(m.ClassInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClassInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Precision) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Precision) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Precision) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxDecimalPlaces != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxDecimalPlaces))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *Balance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Balance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Balance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Balance) > 0 {
		i -= len(m.Balance)
		copy(dAtA[i:], m.Balance)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Balance)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Supply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Supply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Supply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Supply) > 0 {
		i -= len(m.Supply)
		copy(dAtA[i:], m.Supply)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Supply)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ClassInfos) > 0 {
		for _, e := range m.ClassInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BatchInfos) > 0 {
		for _, e := range m.BatchInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.IdSeq != 0 {
		n += 1 + sovGenesis(uint64(m.IdSeq))
	}
	if len(m.TradableBalances) > 0 {
		for _, e := range m.TradableBalances {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RetiredBalances) > 0 {
		for _, e := range m.RetiredBalances {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TradableSupplies) > 0 {
		for _, e := range m.TradableSupplies {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RetiredSupplies) > 0 {
		for _, e := range m.RetiredSupplies {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Precisions) > 0 {
		for _, e := range m.Precisions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Precision) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MaxDecimalPlaces != 0 {
		n += 1 + sovGenesis(uint64(m.MaxDecimalPlaces))
	}
	return n
}

func (m *Balance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Balance)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *Supply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Supply)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassInfos = append(m.ClassInfos, &ClassInfo{})
			if err := m.ClassInfos[len(m.ClassInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchInfos = append(m.BatchInfos, &BatchInfo{})
			if err := m.BatchInfos[len(m.BatchInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdSeq", wireType)
			}
			m.IdSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IdSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradableBalances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradableBalances = append(m.TradableBalances, &Balance{})
			if err := m.TradableBalances[len(m.TradableBalances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetiredBalances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RetiredBalances = append(m.RetiredBalances, &Balance{})
			if err := m.RetiredBalances[len(m.RetiredBalances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradableSupplies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradableSupplies = append(m.TradableSupplies, &Supply{})
			if err := m.TradableSupplies[len(m.TradableSupplies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetiredSupplies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RetiredSupplies = append(m.RetiredSupplies, &Supply{})
			if err := m.RetiredSupplies[len(m.RetiredSupplies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precisions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Precisions = append(m.Precisions, &Precision{})
			if err := m.Precisions[len(m.Precisions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Precision) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Precision: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Precision: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDecimalPlaces", wireType)
			}
			m.MaxDecimalPlaces = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxDecimalPlaces |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Balance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Balance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Balance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Supply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Supply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Supply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Supply = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
