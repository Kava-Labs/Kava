// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/kavadist/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Params governance parameters for kavadist module
type Params struct {
	Active               bool                 `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	Periods              []Period             `protobuf:"bytes,3,rep,name=periods,proto3" json:"periods"`
	InfrastructureParams InfrastructureParams `protobuf:"bytes,4,opt,name=infrastructure_params,json=infrastructureParams,proto3" json:"infrastructure_params"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7a7a4b0c884a4e, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// InfrastructureParams define the parameters for infrastructure rewards.
type InfrastructureParams struct {
	InfrastructurePeriods Periods        `protobuf:"bytes,1,rep,name=infrastructure_periods,json=infrastructurePeriods,proto3,castrepeated=Periods" json:"infrastructure_periods"`
	CoreRewards           CoreRewards    `protobuf:"bytes,2,rep,name=core_rewards,json=coreRewards,proto3,castrepeated=CoreRewards" json:"core_rewards"`
	PartnerRewards        PartnerRewards `protobuf:"bytes,3,rep,name=partner_rewards,json=partnerRewards,proto3,castrepeated=PartnerRewards" json:"partner_rewards"`
}

func (m *InfrastructureParams) Reset()         { *m = InfrastructureParams{} }
func (m *InfrastructureParams) String() string { return proto.CompactTextString(m) }
func (*InfrastructureParams) ProtoMessage()    {}
func (*InfrastructureParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7a7a4b0c884a4e, []int{1}
}
func (m *InfrastructureParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InfrastructureParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InfrastructureParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InfrastructureParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfrastructureParams.Merge(m, src)
}
func (m *InfrastructureParams) XXX_Size() int {
	return m.Size()
}
func (m *InfrastructureParams) XXX_DiscardUnknown() {
	xxx_messageInfo_InfrastructureParams.DiscardUnknown(m)
}

var xxx_messageInfo_InfrastructureParams proto.InternalMessageInfo

// CoreReward defines the reward weights for core infrastructure providers.
type CoreReward struct {
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	Weight  github_com_cosmos_cosmos_sdk_types.Dec        `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight"`
}

func (m *CoreReward) Reset()         { *m = CoreReward{} }
func (m *CoreReward) String() string { return proto.CompactTextString(m) }
func (*CoreReward) ProtoMessage()    {}
func (*CoreReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7a7a4b0c884a4e, []int{2}
}
func (m *CoreReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoreReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoreReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoreReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreReward.Merge(m, src)
}
func (m *CoreReward) XXX_Size() int {
	return m.Size()
}
func (m *CoreReward) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreReward.DiscardUnknown(m)
}

var xxx_messageInfo_CoreReward proto.InternalMessageInfo

// PartnerRewards defines the reward schedule for partner infrastructure providers.
type PartnerReward struct {
	Address          github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	RewardsPerSecond types.Coin                                    `protobuf:"bytes,2,opt,name=rewards_per_second,json=rewardsPerSecond,proto3" json:"rewards_per_second"`
}

func (m *PartnerReward) Reset()         { *m = PartnerReward{} }
func (m *PartnerReward) String() string { return proto.CompactTextString(m) }
func (*PartnerReward) ProtoMessage()    {}
func (*PartnerReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7a7a4b0c884a4e, []int{3}
}
func (m *PartnerReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PartnerReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PartnerReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PartnerReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartnerReward.Merge(m, src)
}
func (m *PartnerReward) XXX_Size() int {
	return m.Size()
}
func (m *PartnerReward) XXX_DiscardUnknown() {
	xxx_messageInfo_PartnerReward.DiscardUnknown(m)
}

var xxx_messageInfo_PartnerReward proto.InternalMessageInfo

// Period stores the specified start and end dates, and the inflation, expressed as a decimal
// representing the yearly APR of KAVA tokens that will be minted during that period
type Period struct {
	// example "2020-03-01T15:20:00Z"
	Start time.Time `protobuf:"bytes,1,opt,name=start,proto3,stdtime" json:"start"`
	// example "2020-06-01T15:20:00Z"
	End time.Time `protobuf:"bytes,2,opt,name=end,proto3,stdtime" json:"end"`
	// example "1.000000003022265980"  - 10% inflation
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation"`
}

func (m *Period) Reset()      { *m = Period{} }
func (*Period) ProtoMessage() {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7a7a4b0c884a4e, []int{4}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Period.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(m, src)
}
func (m *Period) XXX_Size() int {
	return m.Size()
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "kava.kavadist.v1beta1.Params")
	proto.RegisterType((*InfrastructureParams)(nil), "kava.kavadist.v1beta1.InfrastructureParams")
	proto.RegisterType((*CoreReward)(nil), "kava.kavadist.v1beta1.CoreReward")
	proto.RegisterType((*PartnerReward)(nil), "kava.kavadist.v1beta1.PartnerReward")
	proto.RegisterType((*Period)(nil), "kava.kavadist.v1beta1.Period")
}

func init() {
	proto.RegisterFile("kava/kavadist/v1beta1/params.proto", fileDescriptor_2c7a7a4b0c884a4e)
}

var fileDescriptor_2c7a7a4b0c884a4e = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x3f, 0x4f, 0x1b, 0x3f,
	0x18, 0x8e, 0x49, 0x7e, 0x01, 0x1c, 0x7e, 0x50, 0x99, 0x3f, 0x0a, 0x48, 0xbd, 0x4b, 0xa3, 0xaa,
	0x8a, 0x5a, 0xe5, 0x4e, 0xa4, 0x52, 0x07, 0xd4, 0x0e, 0x5c, 0x19, 0x5a, 0x89, 0x4a, 0xe8, 0xca,
	0xd2, 0x2e, 0x91, 0xcf, 0xe7, 0x04, 0x0b, 0x72, 0x3e, 0xd9, 0x0e, 0x94, 0x6f, 0xc1, 0xd8, 0x91,
	0xb9, 0x33, 0x5f, 0xa1, 0x6a, 0x86, 0x0e, 0x88, 0x09, 0x75, 0x80, 0x02, 0x4b, 0x3f, 0x43, 0xa7,
	0xea, 0x6c, 0x87, 0x10, 0x04, 0x52, 0xba, 0x74, 0x49, 0xee, 0x7d, 0xfd, 0x3e, 0xcf, 0xfb, 0x3e,
	0xe7, 0xe7, 0x3d, 0x58, 0xdd, 0xc6, 0xbb, 0xd8, 0xcf, 0x7e, 0x62, 0x26, 0x95, 0xbf, 0xbb, 0x1c,
	0x51, 0x85, 0x97, 0xfd, 0x14, 0x0b, 0xdc, 0x91, 0x5e, 0x2a, 0xb8, 0xe2, 0x68, 0x3e, 0x3b, 0xf6,
	0xfa, 0x35, 0x9e, 0xad, 0x59, 0x72, 0x08, 0x97, 0x1d, 0x2e, 0xfd, 0x08, 0x4b, 0x7a, 0x0d, 0x24,
	0x9c, 0x25, 0x06, 0xb6, 0xb4, 0x68, 0xce, 0x9b, 0x3a, 0xf2, 0x4d, 0x60, 0x8f, 0xe6, 0xda, 0xbc,
	0xcd, 0x4d, 0x3e, 0x7b, 0xb2, 0x59, 0xb7, 0xcd, 0x79, 0x7b, 0x87, 0xfa, 0x3a, 0x8a, 0xba, 0x2d,
	0x5f, 0xb1, 0x0e, 0x95, 0x0a, 0x77, 0x52, 0x53, 0x50, 0xfd, 0x06, 0x60, 0x71, 0x43, 0x4f, 0x86,
	0x16, 0x60, 0x11, 0x13, 0xc5, 0x76, 0x69, 0x19, 0x54, 0x40, 0x6d, 0x22, 0xb4, 0x11, 0x7a, 0x05,
	0xc7, 0x53, 0x2a, 0x18, 0x8f, 0x65, 0x39, 0x5f, 0xc9, 0xd7, 0x4a, 0x8d, 0x87, 0xde, 0x9d, 0xd3,
	0x7b, 0x1b, 0xba, 0x2a, 0x28, 0xf4, 0xce, 0xdc, 0x5c, 0xd8, 0xc7, 0xa0, 0x16, 0x9c, 0x67, 0x49,
	0x4b, 0x60, 0xa9, 0x44, 0x97, 0xa8, 0xae, 0xa0, 0x4d, 0xf3, 0x26, 0xca, 0x85, 0x0a, 0xa8, 0x95,
	0x1a, 0xcf, 0xee, 0x21, 0x7b, 0x3b, 0x84, 0x31, 0x23, 0x5a, 0xea, 0x39, 0x76, 0xc7, 0x59, 0xf5,
	0xeb, 0x18, 0x9c, 0xbb, 0x0b, 0x84, 0x28, 0x5c, 0xb8, 0x3d, 0x80, 0x95, 0x03, 0x46, 0x91, 0x33,
	0x93, 0xf5, 0xfc, 0x72, 0xee, 0x8e, 0x9b, 0x58, 0x86, 0xb7, 0xe4, 0xd8, 0x34, 0xfa, 0x00, 0xa7,
	0x08, 0x17, 0xb4, 0x29, 0xe8, 0x1e, 0x16, 0xb1, 0x2c, 0x8f, 0x69, 0xf2, 0x47, 0xf7, 0x90, 0xbf,
	0xe6, 0x82, 0x86, 0xba, 0x32, 0x98, 0xb5, 0x0d, 0x4a, 0x83, 0x9c, 0x0c, 0x4b, 0x64, 0x10, 0x20,
	0x0a, 0x67, 0x52, 0x2c, 0x54, 0x42, 0xc5, 0x35, 0xbb, 0xb9, 0x89, 0xc7, 0xf7, 0x8d, 0x6e, 0xaa,
	0x6d, 0x83, 0x05, 0xdb, 0x60, 0x7a, 0x28, 0x2d, 0xc3, 0xe9, 0x74, 0x28, 0x5e, 0x29, 0x7c, 0x3e,
	0x74, 0x41, 0xf5, 0x3b, 0x80, 0x70, 0x30, 0x09, 0x8a, 0xe0, 0x38, 0x8e, 0x63, 0x41, 0xa5, 0xd4,
	0xb6, 0x98, 0x0a, 0xde, 0xfc, 0x3e, 0x73, 0xeb, 0x6d, 0xa6, 0xb6, 0xba, 0x91, 0x47, 0x78, 0xc7,
	0xba, 0xd0, 0xfe, 0xd5, 0x65, 0xbc, 0xed, 0xab, 0xfd, 0x94, 0x4a, 0x6f, 0x95, 0x90, 0x55, 0x03,
	0x3c, 0x39, 0xaa, 0xcf, 0x5a, 0xaf, 0xda, 0x4c, 0xb0, 0xaf, 0xa8, 0x0c, 0xfb, 0xc4, 0x68, 0x13,
	0x16, 0xf7, 0x28, 0x6b, 0x6f, 0xa9, 0xf2, 0x58, 0x05, 0xd4, 0x26, 0x83, 0x97, 0xd9, 0xc0, 0x3f,
	0xce, 0xdc, 0x27, 0x23, 0xb4, 0x59, 0xa3, 0xe4, 0xe4, 0xa8, 0x0e, 0x2d, 0xff, 0x1a, 0x25, 0xa1,
	0xe5, 0xb2, 0x72, 0x7a, 0x00, 0xfe, 0x3f, 0xa4, 0xfb, 0x9f, 0x28, 0x7a, 0x07, 0x91, 0xbd, 0xa9,
	0xcc, 0x6c, 0x4d, 0x49, 0x09, 0x4f, 0x62, 0xad, 0xae, 0xd4, 0x58, 0xf4, 0x2c, 0x34, 0xdb, 0xf2,
	0x1b, 0x86, 0x60, 0x89, 0xf5, 0xf7, 0x03, 0x0b, 0xdd, 0xa0, 0xe2, 0xbd, 0x06, 0x5a, 0x29, 0xc7,
	0xd9, 0xae, 0x6a, 0xb7, 0xa1, 0x15, 0xf8, 0x9f, 0x54, 0x58, 0x28, 0xad, 0xa0, 0xd4, 0x58, 0xf2,
	0xcc, 0x9e, 0x7b, 0xfd, 0x3d, 0xf7, 0x36, 0xfb, 0x7b, 0x1e, 0x4c, 0x64, 0x9c, 0x07, 0xe7, 0x2e,
	0x08, 0x0d, 0x04, 0xbd, 0x80, 0x79, 0x7a, 0x3d, 0xcc, 0x68, 0xc8, 0x0c, 0x80, 0xd6, 0xe1, 0x24,
	0x4b, 0x5a, 0x3b, 0x58, 0x31, 0x9e, 0x94, 0xf3, 0xfa, 0xcd, 0x79, 0x7f, 0x77, 0x51, 0xe1, 0x80,
	0x60, 0xa5, 0xf0, 0xeb, 0xd0, 0x05, 0xc1, 0x7a, 0xef, 0xc2, 0xc9, 0x9d, 0x5e, 0x38, 0xb9, 0xde,
	0xa5, 0x03, 0x8e, 0x2f, 0x1d, 0xf0, 0xf3, 0xd2, 0x01, 0x07, 0x57, 0x4e, 0xee, 0xf8, 0xca, 0xc9,
	0x9d, 0x5e, 0x39, 0xb9, 0x8f, 0x4f, 0x6f, 0x50, 0x67, 0x3e, 0xaf, 0xef, 0xe0, 0x48, 0xea, 0x27,
	0xff, 0xd3, 0xe0, 0x23, 0xab, 0x5b, 0x44, 0x45, 0x2d, 0xe2, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xca, 0x25, 0xda, 0x69, 0x82, 0x05, 0x00, 0x00,
}

func (this *Period) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Period)
	if !ok {
		that2, ok := that.(Period)
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
	if !this.Start.Equal(that1.Start) {
		return false
	}
	if !this.End.Equal(that1.End) {
		return false
	}
	if !this.Inflation.Equal(that1.Inflation) {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.InfrastructureParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Periods) > 0 {
		for iNdEx := len(m.Periods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Periods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *InfrastructureParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfrastructureParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InfrastructureParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PartnerRewards) > 0 {
		for iNdEx := len(m.PartnerRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PartnerRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.CoreRewards) > 0 {
		for iNdEx := len(m.CoreRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoreRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.InfrastructurePeriods) > 0 {
		for iNdEx := len(m.InfrastructurePeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InfrastructurePeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CoreReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoreReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoreReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PartnerReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartnerReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PartnerReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RewardsPerSecond.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Period) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Period) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Period) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.End, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.End):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintParams(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Start, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Start):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintParams(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Active {
		n += 2
	}
	if len(m.Periods) > 0 {
		for _, e := range m.Periods {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.InfrastructureParams.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *InfrastructureParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.InfrastructurePeriods) > 0 {
		for _, e := range m.InfrastructurePeriods {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.CoreRewards) > 0 {
		for _, e := range m.CoreRewards {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.PartnerRewards) > 0 {
		for _, e := range m.PartnerRewards {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *CoreReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *PartnerReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.RewardsPerSecond.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Start)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.End)
	n += 1 + l + sovParams(uint64(l))
	l = m.Inflation.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Periods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Periods = append(m.Periods, Period{})
			if err := m.Periods[len(m.Periods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InfrastructureParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InfrastructureParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *InfrastructureParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: InfrastructureParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfrastructureParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InfrastructurePeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InfrastructurePeriods = append(m.InfrastructurePeriods, Period{})
			if err := m.InfrastructurePeriods[len(m.InfrastructurePeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoreRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoreRewards = append(m.CoreRewards, CoreReward{})
			if err := m.CoreRewards[len(m.CoreRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartnerRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartnerRewards = append(m.PartnerRewards, PartnerReward{})
			if err := m.PartnerRewards[len(m.PartnerRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *CoreReward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: CoreReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoreReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *PartnerReward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: PartnerReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartnerReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsPerSecond", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardsPerSecond.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Period) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Period: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Start, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.End, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
