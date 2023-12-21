// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/kavadist/v1beta1/proposal.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// CommunityPoolMultiSpendProposal spends from the community pool by sending to one or more
// addresses
type CommunityPoolMultiSpendProposal struct {
	Title         string                `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	RecipientList []MultiSpendRecipient `protobuf:"bytes,3,rep,name=recipient_list,json=recipientList,proto3" json:"recipient_list"`
}

func (m *CommunityPoolMultiSpendProposal) Reset()      { *m = CommunityPoolMultiSpendProposal{} }
func (*CommunityPoolMultiSpendProposal) ProtoMessage() {}
func (*CommunityPoolMultiSpendProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ee2c0b398254fd, []int{0}
}
func (m *CommunityPoolMultiSpendProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommunityPoolMultiSpendProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommunityPoolMultiSpendProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommunityPoolMultiSpendProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityPoolMultiSpendProposal.Merge(m, src)
}
func (m *CommunityPoolMultiSpendProposal) XXX_Size() int {
	return m.Size()
}
func (m *CommunityPoolMultiSpendProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityPoolMultiSpendProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityPoolMultiSpendProposal proto.InternalMessageInfo

// CommunityPoolMultiSpendProposalJSON defines a CommunityPoolMultiSpendProposal with a deposit
type CommunityPoolMultiSpendProposalJSON struct {
	Title         string                                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	RecipientList []MultiSpendRecipient                    `protobuf:"bytes,3,rep,name=recipient_list,json=recipientList,proto3" json:"recipient_list"`
	Deposit       github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=deposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"deposit"`
}

func (m *CommunityPoolMultiSpendProposalJSON) Reset()         { *m = CommunityPoolMultiSpendProposalJSON{} }
func (m *CommunityPoolMultiSpendProposalJSON) String() string { return proto.CompactTextString(m) }
func (*CommunityPoolMultiSpendProposalJSON) ProtoMessage()    {}
func (*CommunityPoolMultiSpendProposalJSON) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ee2c0b398254fd, []int{1}
}
func (m *CommunityPoolMultiSpendProposalJSON) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommunityPoolMultiSpendProposalJSON) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommunityPoolMultiSpendProposalJSON.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommunityPoolMultiSpendProposalJSON) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityPoolMultiSpendProposalJSON.Merge(m, src)
}
func (m *CommunityPoolMultiSpendProposalJSON) XXX_Size() int {
	return m.Size()
}
func (m *CommunityPoolMultiSpendProposalJSON) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityPoolMultiSpendProposalJSON.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityPoolMultiSpendProposalJSON proto.InternalMessageInfo

// MultiSpendRecipient defines a recipient and the amount of coins they are receiving
type MultiSpendRecipient struct {
	Address string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *MultiSpendRecipient) Reset()      { *m = MultiSpendRecipient{} }
func (*MultiSpendRecipient) ProtoMessage() {}
func (*MultiSpendRecipient) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ee2c0b398254fd, []int{2}
}
func (m *MultiSpendRecipient) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiSpendRecipient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiSpendRecipient.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiSpendRecipient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiSpendRecipient.Merge(m, src)
}
func (m *MultiSpendRecipient) XXX_Size() int {
	return m.Size()
}
func (m *MultiSpendRecipient) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiSpendRecipient.DiscardUnknown(m)
}

var xxx_messageInfo_MultiSpendRecipient proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CommunityPoolMultiSpendProposal)(nil), "kava.kavadist.v1beta1.CommunityPoolMultiSpendProposal")
	proto.RegisterType((*CommunityPoolMultiSpendProposalJSON)(nil), "kava.kavadist.v1beta1.CommunityPoolMultiSpendProposalJSON")
	proto.RegisterType((*MultiSpendRecipient)(nil), "kava.kavadist.v1beta1.MultiSpendRecipient")
}

func init() {
	proto.RegisterFile("kava/kavadist/v1beta1/proposal.proto", fileDescriptor_22ee2c0b398254fd)
}

var fileDescriptor_22ee2c0b398254fd = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0xbf, 0x0f, 0xd2, 0x40,
	0x14, 0xc7, 0x7b, 0x80, 0xa0, 0x47, 0x74, 0xa8, 0x98, 0x54, 0x86, 0x96, 0xa0, 0x03, 0x21, 0xe1,
	0x4e, 0x74, 0x73, 0x04, 0x27, 0xe3, 0x0f, 0x52, 0x06, 0x13, 0x17, 0x73, 0x6d, 0x2f, 0x78, 0xa1,
	0xed, 0xbb, 0xf4, 0xae, 0x44, 0xfe, 0x03, 0x47, 0x47, 0x27, 0xc3, 0x66, 0xe2, 0xdf, 0xe0, 0x1f,
	0xc0, 0xc8, 0xe8, 0xa4, 0x06, 0xfe, 0x11, 0xd3, 0x9f, 0x30, 0x90, 0xb8, 0x38, 0xb8, 0xb4, 0xef,
	0xae, 0xef, 0xfb, 0xb9, 0xf7, 0xde, 0xb7, 0x87, 0x1f, 0xae, 0xd9, 0x86, 0xd1, 0xec, 0x11, 0x08,
	0xa5, 0xe9, 0x66, 0xea, 0x71, 0xcd, 0xa6, 0x54, 0x26, 0x20, 0x41, 0xb1, 0x90, 0xc8, 0x04, 0x34,
	0x98, 0xf7, 0xb2, 0x04, 0x52, 0x65, 0x91, 0x32, 0xab, 0x6f, 0xfb, 0xa0, 0x22, 0x50, 0xd4, 0x63,
	0x8a, 0xd7, 0x52, 0x1f, 0x44, 0x5c, 0xc8, 0xfa, 0xbd, 0x15, 0xac, 0x20, 0x0f, 0x69, 0x16, 0x15,
	0xbb, 0xc3, 0xef, 0x08, 0x3b, 0x73, 0x88, 0xa2, 0x34, 0x16, 0x7a, 0xbb, 0x00, 0x08, 0x5f, 0xa6,
	0xa1, 0x16, 0x4b, 0xc9, 0xe3, 0x60, 0x51, 0x1e, 0x6b, 0xf6, 0xf0, 0x0d, 0x2d, 0x74, 0xc8, 0x2d,
	0x34, 0x40, 0xa3, 0x5b, 0x6e, 0xb1, 0x30, 0x07, 0xb8, 0x1b, 0x70, 0xe5, 0x27, 0x42, 0x6a, 0x01,
	0xb1, 0xd5, 0xc8, 0xbf, 0x5d, 0x6e, 0x99, 0x6f, 0xf0, 0x9d, 0x84, 0xfb, 0x42, 0x0a, 0x1e, 0xeb,
	0x77, 0xa1, 0x50, 0xda, 0x6a, 0x0e, 0x9a, 0xa3, 0xee, 0xe3, 0x31, 0xb9, 0xda, 0x01, 0x39, 0x1f,
	0xed, 0x56, 0xb2, 0x59, 0x6b, 0xff, 0xd3, 0x31, 0xdc, 0xdb, 0x35, 0xe7, 0x85, 0x50, 0xfa, 0xe9,
	0xcd, 0x8f, 0x3b, 0xc7, 0xf8, 0xbc, 0x73, 0x8c, 0xe1, 0xd7, 0x06, 0x7e, 0xf0, 0x97, 0xf2, 0x9f,
	0x2f, 0x5f, 0xbf, 0xfa, 0xef, 0x5a, 0x30, 0x39, 0xee, 0x04, 0x5c, 0x82, 0x12, 0xda, 0x6a, 0xe5,
	0xc4, 0xfb, 0xa4, 0xf0, 0x8f, 0x64, 0xfe, 0xd5, 0xbc, 0x39, 0x88, 0x78, 0xf6, 0x28, 0x03, 0x7c,
	0xfb, 0xe5, 0x8c, 0x56, 0x42, 0xbf, 0x4f, 0x3d, 0xe2, 0x43, 0x44, 0x4b, 0xb3, 0x8b, 0xd7, 0x44,
	0x05, 0x6b, 0xaa, 0xb7, 0x92, 0xab, 0x5c, 0xa0, 0xdc, 0x8a, 0x5d, 0x4f, 0x0a, 0x0d, 0xbf, 0x20,
	0x7c, 0xf7, 0x4a, 0x75, 0xa6, 0x85, 0x3b, 0x2c, 0x08, 0x12, 0xae, 0x54, 0x39, 0x9b, 0x6a, 0x69,
	0xfa, 0xb8, 0xcd, 0x22, 0x48, 0x63, 0x6d, 0x35, 0xfe, 0x7d, 0x85, 0x25, 0xfa, 0x6c, 0xe5, 0xec,
	0xd9, 0xfe, 0x68, 0xa3, 0xc3, 0xd1, 0x46, 0xbf, 0x8f, 0x36, 0xfa, 0x74, 0xb2, 0x8d, 0xc3, 0xc9,
	0x36, 0x7e, 0x9c, 0x6c, 0xe3, 0xed, 0xf8, 0x82, 0x9a, 0x4d, 0x7c, 0x12, 0x32, 0x4f, 0xe5, 0x11,
	0xfd, 0x70, 0xbe, 0x2d, 0x39, 0xdd, 0x6b, 0xe7, 0xbf, 0xf5, 0x93, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x02, 0xb8, 0xb8, 0x96, 0x4b, 0x03, 0x00, 0x00,
}

func (m *CommunityPoolMultiSpendProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommunityPoolMultiSpendProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommunityPoolMultiSpendProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RecipientList) > 0 {
		for iNdEx := len(m.RecipientList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RecipientList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommunityPoolMultiSpendProposalJSON) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommunityPoolMultiSpendProposalJSON) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommunityPoolMultiSpendProposalJSON) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deposit) > 0 {
		for iNdEx := len(m.Deposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.RecipientList) > 0 {
		for iNdEx := len(m.RecipientList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RecipientList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MultiSpendRecipient) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiSpendRecipient) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiSpendRecipient) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommunityPoolMultiSpendProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.RecipientList) > 0 {
		for _, e := range m.RecipientList {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *CommunityPoolMultiSpendProposalJSON) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.RecipientList) > 0 {
		for _, e := range m.RecipientList {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if len(m.Deposit) > 0 {
		for _, e := range m.Deposit {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *MultiSpendRecipient) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommunityPoolMultiSpendProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: CommunityPoolMultiSpendProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommunityPoolMultiSpendProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientList = append(m.RecipientList, MultiSpendRecipient{})
			if err := m.RecipientList[len(m.RecipientList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *CommunityPoolMultiSpendProposalJSON) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: CommunityPoolMultiSpendProposalJSON: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommunityPoolMultiSpendProposalJSON: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientList = append(m.RecipientList, MultiSpendRecipient{})
			if err := m.RecipientList[len(m.RecipientList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposit = append(m.Deposit, types.Coin{})
			if err := m.Deposit[len(m.Deposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *MultiSpendRecipient) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: MultiSpendRecipient: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiSpendRecipient: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
