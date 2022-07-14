// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/earn/v1beta1/vault.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// AllowedVault is a vault that is allowed to be created. These can be
// modified via parameter governance.
type AllowedVault struct {
	// Denom is the only supported denomination of the vault for deposits and withdrawals.
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// VaultStrategy is the strategies to use for this vault.
	VaultStrategy StrategyType `protobuf:"varint,2,opt,name=vault_strategy,json=vaultStrategy,proto3,enum=kava.earn.v1beta1.StrategyType" json:"vault_strategy,omitempty"`
}

func (m *AllowedVault) Reset()         { *m = AllowedVault{} }
func (m *AllowedVault) String() string { return proto.CompactTextString(m) }
func (*AllowedVault) ProtoMessage()    {}
func (*AllowedVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_884eb89509fbdc04, []int{0}
}
func (m *AllowedVault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedVault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedVault.Merge(m, src)
}
func (m *AllowedVault) XXX_Size() int {
	return m.Size()
}
func (m *AllowedVault) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedVault.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedVault proto.InternalMessageInfo

func (m *AllowedVault) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *AllowedVault) GetVaultStrategy() StrategyType {
	if m != nil {
		return m.VaultStrategy
	}
	return STRATEGY_TYPE_UNKNOWN
}

// VaultRecord is the state of a vault and is used to store the state of a
// vault.
type VaultRecord struct {
	// Denom is the only supported denomination of the vault for deposits and
	// withdrawals.
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// TotalSupply is the total supply of the vault, denominated **only** in the
	// user deposit/withdrawal denom, must be the same as the Denom field.
	TotalSupply types.Coin `protobuf:"bytes,2,opt,name=total_supply,json=totalSupply,proto3" json:"total_supply"`
}

func (m *VaultRecord) Reset()         { *m = VaultRecord{} }
func (m *VaultRecord) String() string { return proto.CompactTextString(m) }
func (*VaultRecord) ProtoMessage()    {}
func (*VaultRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_884eb89509fbdc04, []int{1}
}
func (m *VaultRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultRecord.Merge(m, src)
}
func (m *VaultRecord) XXX_Size() int {
	return m.Size()
}
func (m *VaultRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultRecord.DiscardUnknown(m)
}

var xxx_messageInfo_VaultRecord proto.InternalMessageInfo

func (m *VaultRecord) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *VaultRecord) GetTotalSupply() types.Coin {
	if m != nil {
		return m.TotalSupply
	}
	return types.Coin{}
}

// VaultShareRecord defines the shares owned by a depositor and vault.
type VaultShareRecord struct {
	// depositor represents the owner of the shares
	Depositor github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=depositor,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"depositor,omitempty"`
	// amount_supplied represents the total amount a depositor has supplied to the
	// vault. The vault is determined by the coin denom.
	AmountSupplied types.Coin `protobuf:"bytes,2,opt,name=amount_supplied,json=amountSupplied,proto3" json:"amount_supplied"`
}

func (m *VaultShareRecord) Reset()         { *m = VaultShareRecord{} }
func (m *VaultShareRecord) String() string { return proto.CompactTextString(m) }
func (*VaultShareRecord) ProtoMessage()    {}
func (*VaultShareRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_884eb89509fbdc04, []int{2}
}
func (m *VaultShareRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultShareRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultShareRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultShareRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultShareRecord.Merge(m, src)
}
func (m *VaultShareRecord) XXX_Size() int {
	return m.Size()
}
func (m *VaultShareRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultShareRecord.DiscardUnknown(m)
}

var xxx_messageInfo_VaultShareRecord proto.InternalMessageInfo

func (m *VaultShareRecord) GetDepositor() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Depositor
	}
	return nil
}

func (m *VaultShareRecord) GetAmountSupplied() types.Coin {
	if m != nil {
		return m.AmountSupplied
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*AllowedVault)(nil), "kava.earn.v1beta1.AllowedVault")
	proto.RegisterType((*VaultRecord)(nil), "kava.earn.v1beta1.VaultRecord")
	proto.RegisterType((*VaultShareRecord)(nil), "kava.earn.v1beta1.VaultShareRecord")
}

func init() { proto.RegisterFile("kava/earn/v1beta1/vault.proto", fileDescriptor_884eb89509fbdc04) }

var fileDescriptor_884eb89509fbdc04 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0xae, 0xd2, 0x40,
	0x14, 0x6e, 0x8d, 0x9a, 0x30, 0x20, 0x6a, 0x65, 0x01, 0x24, 0x16, 0xc2, 0xc2, 0xb0, 0xe9, 0x34,
	0xe0, 0x0b, 0x48, 0x4d, 0x0c, 0xeb, 0xd6, 0xb8, 0x70, 0x43, 0xa6, 0x9d, 0xb1, 0x34, 0xb4, 0x3d,
	0x4d, 0x67, 0x8a, 0xf6, 0x2d, 0x7c, 0x18, 0x1f, 0xc1, 0x05, 0x4b, 0xe2, 0xca, 0x15, 0xb9, 0x81,
	0xb7, 0xb8, 0xab, 0x9b, 0xce, 0x0c, 0xdc, 0x9b, 0x90, 0x9b, 0xdc, 0xd5, 0xfc, 0x7c, 0xe7, 0x3b,
	0xdf, 0xf7, 0x9d, 0x83, 0xde, 0x6f, 0xc8, 0x96, 0xb8, 0x8c, 0x94, 0xb9, 0xbb, 0x9d, 0x85, 0x4c,
	0x90, 0x99, 0xbb, 0x25, 0x55, 0x2a, 0x70, 0x51, 0x82, 0x00, 0xeb, 0x6d, 0x03, 0xe3, 0x06, 0xc6,
	0x1a, 0x1e, 0xf6, 0x62, 0x88, 0x41, 0xa2, 0x6e, 0x73, 0x53, 0x85, 0x43, 0x3b, 0x02, 0x9e, 0x01,
	0x77, 0x43, 0xc2, 0xd9, 0xa5, 0x53, 0x04, 0x49, 0xae, 0xf1, 0x81, 0xc2, 0x57, 0x8a, 0xa8, 0x1e,
	0x1a, 0x1a, 0x5f, 0x5b, 0xe0, 0xa2, 0x24, 0x82, 0xc5, 0xb5, 0xaa, 0x98, 0xa4, 0xa8, 0xb3, 0x48,
	0x53, 0xf8, 0xc9, 0xe8, 0xb7, 0xc6, 0x9b, 0xd5, 0x43, 0x2f, 0x28, 0xcb, 0x21, 0xeb, 0x9b, 0x63,
	0x73, 0xda, 0xf2, 0xd5, 0xc3, 0xfa, 0x82, 0xba, 0xd2, 0xfa, 0xea, 0xcc, 0xee, 0x3f, 0x1b, 0x9b,
	0xd3, 0xee, 0x7c, 0x84, 0xaf, 0x42, 0xe0, 0x40, 0x97, 0x7c, 0xad, 0x0b, 0xe6, 0xbf, 0x92, 0xb4,
	0xf3, 0xd7, 0x24, 0x46, 0x6d, 0x29, 0xe3, 0xb3, 0x08, 0x4a, 0xfa, 0x88, 0x98, 0x87, 0x3a, 0x02,
	0x04, 0x49, 0x57, 0xbc, 0x2a, 0x8a, 0x54, 0x49, 0xb5, 0xe7, 0x03, 0xac, 0x93, 0x35, 0x63, 0xb8,
	0x88, 0x7d, 0x86, 0x24, 0xf7, 0x9e, 0xef, 0x0e, 0x23, 0xc3, 0x6f, 0x4b, 0x52, 0x20, 0x39, 0x93,
	0xbf, 0x26, 0x7a, 0x23, 0x95, 0x82, 0x35, 0x29, 0x99, 0x96, 0xfb, 0x81, 0x5a, 0x94, 0x15, 0xc0,
	0x13, 0x01, 0xa5, 0x94, 0xec, 0x78, 0xcb, 0xdb, 0xc3, 0xc8, 0x89, 0x13, 0xb1, 0xae, 0x42, 0x1c,
	0x41, 0xa6, 0xa7, 0xa7, 0x0f, 0x87, 0xd3, 0x8d, 0x2b, 0xea, 0x82, 0x71, 0xbc, 0x88, 0xa2, 0x05,
	0xa5, 0x25, 0xe3, 0xfc, 0xdf, 0x1f, 0xe7, 0x9d, 0x76, 0xa2, 0x7f, 0xbc, 0x5a, 0x30, 0xee, 0xdf,
	0xb7, 0xb6, 0x96, 0xe8, 0x35, 0xc9, 0xa0, 0xca, 0x85, 0x4a, 0x90, 0x30, 0xfa, 0xd4, 0x0c, 0x5d,
	0xc5, 0x0b, 0x34, 0xcd, 0xfb, 0xb4, 0x3b, 0xda, 0xe6, 0xfe, 0x68, 0x9b, 0x37, 0x47, 0xdb, 0xfc,
	0x7d, 0xb2, 0x8d, 0xfd, 0xc9, 0x36, 0xfe, 0x9f, 0x6c, 0xe3, 0xfb, 0x87, 0x07, 0xa6, 0x9b, 0x1d,
	0x38, 0x29, 0x09, 0xb9, 0xbc, 0xb9, 0xbf, 0xd4, 0xc2, 0xa5, 0xf1, 0xf0, 0xa5, 0x5c, 0xf3, 0xc7,
	0xbb, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x8b, 0xa9, 0xd7, 0x8d, 0x02, 0x00, 0x00,
}

func (m *AllowedVault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedVault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedVault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VaultStrategy != 0 {
		i = encodeVarintVault(dAtA, i, uint64(m.VaultStrategy))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintVault(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VaultRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TotalSupply.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintVault(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VaultShareRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultShareRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultShareRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AmountSupplied.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintVault(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVault(dAtA []byte, offset int, v uint64) int {
	offset -= sovVault(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AllowedVault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovVault(uint64(l))
	}
	if m.VaultStrategy != 0 {
		n += 1 + sovVault(uint64(m.VaultStrategy))
	}
	return n
}

func (m *VaultRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovVault(uint64(l))
	}
	l = m.TotalSupply.Size()
	n += 1 + l + sovVault(uint64(l))
	return n
}

func (m *VaultShareRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovVault(uint64(l))
	}
	l = m.AmountSupplied.Size()
	n += 1 + l + sovVault(uint64(l))
	return n
}

func sovVault(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVault(x uint64) (n int) {
	return sovVault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AllowedVault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVault
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
			return fmt.Errorf("proto: AllowedVault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedVault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultStrategy", wireType)
			}
			m.VaultStrategy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VaultStrategy |= StrategyType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
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
func (m *VaultRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVault
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
			return fmt.Errorf("proto: VaultRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
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
func (m *VaultShareRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVault
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
			return fmt.Errorf("proto: VaultShareRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultShareRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = append(m.Depositor[:0], dAtA[iNdEx:postIndex]...)
			if m.Depositor == nil {
				m.Depositor = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountSupplied", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountSupplied.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
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
func skipVault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVault
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
					return 0, ErrIntOverflowVault
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
					return 0, ErrIntOverflowVault
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
				return 0, ErrInvalidLengthVault
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVault
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVault
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVault        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVault          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVault = fmt.Errorf("proto: unexpected end of group")
)
