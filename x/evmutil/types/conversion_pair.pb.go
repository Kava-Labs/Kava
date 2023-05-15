// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/evmutil/v1beta1/conversion_pair.proto

package types

import (
	bytes "bytes"
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

// ConversionPair defines a Kava ERC20 address and corresponding denom that is
// allowed to be converted between ERC20 and sdk.Coin
type ConversionPair struct {
	// ERC20 address of the token on the Kava EVM
	KavaERC20Address HexBytes `protobuf:"bytes,1,opt,name=kava_erc20_address,json=kavaErc20Address,proto3,casttype=HexBytes" json:"kava_erc20_address,omitempty"`
	// Denom of the corresponding sdk.Coin
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *ConversionPair) Reset()         { *m = ConversionPair{} }
func (m *ConversionPair) String() string { return proto.CompactTextString(m) }
func (*ConversionPair) ProtoMessage()    {}
func (*ConversionPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1396d08199817d0, []int{0}
}
func (m *ConversionPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConversionPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConversionPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConversionPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionPair.Merge(m, src)
}
func (m *ConversionPair) XXX_Size() int {
	return m.Size()
}
func (m *ConversionPair) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionPair.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionPair proto.InternalMessageInfo

// AllowedNativeCoinERC20Token defines allowed sdk denom & metadata
// for evm token representations of sdk assets.
// NOTE: once evm token contracts are deployed, changes to metadata for a given
// sdk_denom will not change metadata of deployed contract.
type AllowedNativeCoinERC20Token struct {
	// Denom of the sdk.Coin
	SdkDenom string `protobuf:"bytes,1,opt,name=sdk_denom,json=sdkDenom,proto3" json:"sdk_denom,omitempty"`
	// Name of ERC20 contract
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Symbol of ERC20 contract
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Number of decimals ERC20 contract is deployed with.
	Decimal uint32 `protobuf:"varint,4,opt,name=decimal,proto3" json:"decimal,omitempty"`
}

func (m *AllowedNativeCoinERC20Token) Reset()         { *m = AllowedNativeCoinERC20Token{} }
func (m *AllowedNativeCoinERC20Token) String() string { return proto.CompactTextString(m) }
func (*AllowedNativeCoinERC20Token) ProtoMessage()    {}
func (*AllowedNativeCoinERC20Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1396d08199817d0, []int{1}
}
func (m *AllowedNativeCoinERC20Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedNativeCoinERC20Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedNativeCoinERC20Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedNativeCoinERC20Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedNativeCoinERC20Token.Merge(m, src)
}
func (m *AllowedNativeCoinERC20Token) XXX_Size() int {
	return m.Size()
}
func (m *AllowedNativeCoinERC20Token) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedNativeCoinERC20Token.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedNativeCoinERC20Token proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConversionPair)(nil), "kava.evmutil.v1beta1.ConversionPair")
	proto.RegisterType((*AllowedNativeCoinERC20Token)(nil), "kava.evmutil.v1beta1.AllowedNativeCoinERC20Token")
}

func init() {
	proto.RegisterFile("kava/evmutil/v1beta1/conversion_pair.proto", fileDescriptor_e1396d08199817d0)
}

var fileDescriptor_e1396d08199817d0 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x91, 0xcf, 0x6a, 0xea, 0x40,
	0x18, 0xc5, 0x33, 0xf7, 0x7a, 0xbd, 0x3a, 0xdc, 0x5b, 0x64, 0x90, 0x12, 0x2a, 0x8c, 0xc1, 0x95,
	0x2d, 0x34, 0x51, 0xbb, 0xeb, 0x4e, 0xad, 0x50, 0x10, 0xa4, 0x84, 0xae, 0xba, 0x09, 0x93, 0xe4,
	0xc3, 0x0e, 0xf9, 0x33, 0x92, 0x89, 0xa9, 0x42, 0xd7, 0xa5, 0xcb, 0x3e, 0x42, 0x97, 0x7d, 0x94,
	0x2e, 0x5d, 0x76, 0x25, 0x36, 0xbe, 0x45, 0x57, 0x25, 0x7f, 0xea, 0xee, 0x9c, 0xf9, 0xce, 0xef,
	0x30, 0x33, 0x1f, 0x3e, 0xf3, 0x58, 0xc2, 0x0c, 0x48, 0x82, 0x65, 0xcc, 0x7d, 0x23, 0xe9, 0xdb,
	0x10, 0xb3, 0xbe, 0xe1, 0x88, 0x30, 0x81, 0x48, 0x72, 0x11, 0x5a, 0x0b, 0xc6, 0x23, 0x7d, 0x11,
	0x89, 0x58, 0x90, 0x66, 0x96, 0xd5, 0xcb, 0xac, 0x5e, 0x66, 0x4f, 0x9a, 0x73, 0x31, 0x17, 0x79,
	0xc0, 0xc8, 0x54, 0x91, 0xed, 0x3c, 0xe2, 0xa3, 0xf1, 0xa1, 0xe4, 0x86, 0xf1, 0x88, 0xcc, 0x30,
	0xc9, 0x78, 0x0b, 0x22, 0x67, 0xd0, 0xb3, 0x98, 0xeb, 0x46, 0x20, 0xa5, 0x8a, 0x34, 0xd4, 0xfd,
	0x37, 0xd2, 0xd2, 0x6d, 0xbb, 0x31, 0x65, 0x09, 0x9b, 0x98, 0xe3, 0x41, 0x6f, 0x58, 0xcc, 0xbe,
	0xb6, 0xed, 0xda, 0x35, 0xac, 0x46, 0xeb, 0x18, 0xa4, 0xd9, 0xc8, 0xd8, 0x49, 0x86, 0x96, 0x53,
	0xd2, 0xc4, 0x7f, 0x5c, 0x08, 0x45, 0xa0, 0xfe, 0xd2, 0x50, 0xb7, 0x6e, 0x16, 0xe6, 0xb2, 0xf2,
	0xfc, 0xda, 0x56, 0x3a, 0x4f, 0x08, 0xb7, 0x86, 0xbe, 0x2f, 0x1e, 0xc0, 0x9d, 0xb1, 0x98, 0x27,
	0x30, 0x16, 0x3c, 0xcc, 0xbb, 0x6f, 0x85, 0x07, 0x21, 0x69, 0xe1, 0xba, 0x74, 0x3d, 0xab, 0xe0,
	0x51, 0xce, 0xd7, 0xa4, 0xeb, 0x5d, 0x65, 0x9e, 0x10, 0x5c, 0x09, 0x59, 0x00, 0x65, 0x6f, 0xae,
	0xc9, 0x31, 0xae, 0xca, 0x75, 0x60, 0x0b, 0x5f, 0xfd, 0x9d, 0x9f, 0x96, 0x8e, 0xa8, 0xf8, 0xaf,
	0x0b, 0x0e, 0x0f, 0x98, 0xaf, 0x56, 0x34, 0xd4, 0xfd, 0x6f, 0xfe, 0xd8, 0xe2, 0x22, 0xa3, 0xe9,
	0xee, 0x93, 0xa2, 0xb7, 0x94, 0xa2, 0xf7, 0x94, 0xa2, 0x4d, 0x4a, 0xd1, 0x2e, 0xa5, 0xe8, 0x65,
	0x4f, 0x95, 0xcd, 0x9e, 0x2a, 0x1f, 0x7b, 0xaa, 0xdc, 0x9d, 0xce, 0x79, 0x7c, 0xbf, 0xb4, 0x75,
	0x47, 0x04, 0x46, 0xf6, 0xc6, 0x73, 0x9f, 0xd9, 0x32, 0x57, 0xc6, 0xea, 0xb0, 0x97, 0x78, 0xbd,
	0x00, 0x69, 0x57, 0xf3, 0xaf, 0xbd, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x27, 0x58, 0xfe, 0x57,
	0xb4, 0x01, 0x00, 0x00,
}

func (this *ConversionPair) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ConversionPair)
	if !ok {
		that2, ok := that.(ConversionPair)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ConversionPair")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ConversionPair but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ConversionPair but is not nil && this == nil")
	}
	if !bytes.Equal(this.KavaERC20Address, that1.KavaERC20Address) {
		return fmt.Errorf("KavaERC20Address this(%v) Not Equal that(%v)", this.KavaERC20Address, that1.KavaERC20Address)
	}
	if this.Denom != that1.Denom {
		return fmt.Errorf("Denom this(%v) Not Equal that(%v)", this.Denom, that1.Denom)
	}
	return nil
}
func (this *ConversionPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConversionPair)
	if !ok {
		that2, ok := that.(ConversionPair)
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
	if !bytes.Equal(this.KavaERC20Address, that1.KavaERC20Address) {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	return true
}
func (this *AllowedNativeCoinERC20Token) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*AllowedNativeCoinERC20Token)
	if !ok {
		that2, ok := that.(AllowedNativeCoinERC20Token)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *AllowedNativeCoinERC20Token")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *AllowedNativeCoinERC20Token but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *AllowedNativeCoinERC20Token but is not nil && this == nil")
	}
	if this.SdkDenom != that1.SdkDenom {
		return fmt.Errorf("SdkDenom this(%v) Not Equal that(%v)", this.SdkDenom, that1.SdkDenom)
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if this.Symbol != that1.Symbol {
		return fmt.Errorf("Symbol this(%v) Not Equal that(%v)", this.Symbol, that1.Symbol)
	}
	if this.Decimal != that1.Decimal {
		return fmt.Errorf("Decimal this(%v) Not Equal that(%v)", this.Decimal, that1.Decimal)
	}
	return nil
}
func (this *AllowedNativeCoinERC20Token) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AllowedNativeCoinERC20Token)
	if !ok {
		that2, ok := that.(AllowedNativeCoinERC20Token)
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
	if this.SdkDenom != that1.SdkDenom {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Symbol != that1.Symbol {
		return false
	}
	if this.Decimal != that1.Decimal {
		return false
	}
	return true
}
func (m *ConversionPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConversionPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConversionPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.KavaERC20Address) > 0 {
		i -= len(m.KavaERC20Address)
		copy(dAtA[i:], m.KavaERC20Address)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.KavaERC20Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AllowedNativeCoinERC20Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedNativeCoinERC20Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedNativeCoinERC20Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimal != 0 {
		i = encodeVarintConversionPair(dAtA, i, uint64(m.Decimal))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SdkDenom) > 0 {
		i -= len(m.SdkDenom)
		copy(dAtA[i:], m.SdkDenom)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.SdkDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConversionPair(dAtA []byte, offset int, v uint64) int {
	offset -= sovConversionPair(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConversionPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KavaERC20Address)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	return n
}

func (m *AllowedNativeCoinERC20Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SdkDenom)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	if m.Decimal != 0 {
		n += 1 + sovConversionPair(uint64(m.Decimal))
	}
	return n
}

func sovConversionPair(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConversionPair(x uint64) (n int) {
	return sovConversionPair(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConversionPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConversionPair
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
			return fmt.Errorf("proto: ConversionPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConversionPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KavaERC20Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KavaERC20Address = append(m.KavaERC20Address[:0], dAtA[iNdEx:postIndex]...)
			if m.KavaERC20Address == nil {
				m.KavaERC20Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConversionPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConversionPair
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
func (m *AllowedNativeCoinERC20Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConversionPair
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
			return fmt.Errorf("proto: AllowedNativeCoinERC20Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedNativeCoinERC20Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SdkDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SdkDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			m.Decimal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimal |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConversionPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConversionPair
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
func skipConversionPair(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConversionPair
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
					return 0, ErrIntOverflowConversionPair
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
					return 0, ErrIntOverflowConversionPair
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
				return 0, ErrInvalidLengthConversionPair
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConversionPair
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConversionPair
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConversionPair        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConversionPair          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConversionPair = fmt.Errorf("proto: unexpected end of group")
)
