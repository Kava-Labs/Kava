// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/committee/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgSubmitProposal is used by committee members to create a new proposal that they can vote on.
type MsgSubmitProposal struct {
	PubProposal *types.Any `protobuf:"bytes,1,opt,name=pub_proposal,json=pubProposal,proto3" json:"pub_proposal,omitempty"`
	Proposer    string     `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	CommitteeID uint64     `protobuf:"varint,3,opt,name=committee_id,json=committeeId,proto3" json:"committee_id,omitempty"`
}

func (m *MsgSubmitProposal) Reset()         { *m = MsgSubmitProposal{} }
func (m *MsgSubmitProposal) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitProposal) ProtoMessage()    {}
func (*MsgSubmitProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f3857845b071606, []int{0}
}
func (m *MsgSubmitProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitProposal.Merge(m, src)
}
func (m *MsgSubmitProposal) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitProposal.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitProposal proto.InternalMessageInfo

// MsgSubmitProposalResponse defines the SubmitProposal response type
type MsgSubmitProposalResponse struct {
	ProposalID uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (m *MsgSubmitProposalResponse) Reset()         { *m = MsgSubmitProposalResponse{} }
func (m *MsgSubmitProposalResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitProposalResponse) ProtoMessage()    {}
func (*MsgSubmitProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f3857845b071606, []int{1}
}
func (m *MsgSubmitProposalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitProposalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitProposalResponse.Merge(m, src)
}
func (m *MsgSubmitProposalResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitProposalResponse proto.InternalMessageInfo

// MsgVote is submitted by committee members to vote on proposals.
type MsgVote struct {
	ProposalID uint64   `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Voter      string   `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
	VoteType   VoteType `protobuf:"varint,3,opt,name=vote_type,json=voteType,proto3,enum=kava.committee.v1beta1.VoteType" json:"vote_type,omitempty"`
}

func (m *MsgVote) Reset()         { *m = MsgVote{} }
func (m *MsgVote) String() string { return proto.CompactTextString(m) }
func (*MsgVote) ProtoMessage()    {}
func (*MsgVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f3857845b071606, []int{2}
}
func (m *MsgVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVote.Merge(m, src)
}
func (m *MsgVote) XXX_Size() int {
	return m.Size()
}
func (m *MsgVote) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVote.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVote proto.InternalMessageInfo

// MsgVoteResponse defines the Vote response type
type MsgVoteResponse struct {
}

func (m *MsgVoteResponse) Reset()         { *m = MsgVoteResponse{} }
func (m *MsgVoteResponse) String() string { return proto.CompactTextString(m) }
func (*MsgVoteResponse) ProtoMessage()    {}
func (*MsgVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f3857845b071606, []int{3}
}
func (m *MsgVoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVoteResponse.Merge(m, src)
}
func (m *MsgVoteResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVoteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitProposal)(nil), "kava.committee.v1beta1.MsgSubmitProposal")
	proto.RegisterType((*MsgSubmitProposalResponse)(nil), "kava.committee.v1beta1.MsgSubmitProposalResponse")
	proto.RegisterType((*MsgVote)(nil), "kava.committee.v1beta1.MsgVote")
	proto.RegisterType((*MsgVoteResponse)(nil), "kava.committee.v1beta1.MsgVoteResponse")
}

func init() { proto.RegisterFile("kava/committee/v1beta1/tx.proto", fileDescriptor_3f3857845b071606) }

var fileDescriptor_3f3857845b071606 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0xb4, 0x40, 0x3b, 0xae, 0x52, 0xd5, 0x8a, 0x50, 0xe2, 0x83, 0x13, 0x45, 0x48,
	0x04, 0xa1, 0xee, 0x2a, 0xe1, 0xcc, 0x81, 0xb4, 0x17, 0x4b, 0x44, 0xaa, 0x0c, 0x02, 0x89, 0x4b,
	0x64, 0x37, 0xcb, 0x62, 0x91, 0x78, 0xac, 0xec, 0xda, 0xaa, 0x9f, 0x02, 0x1e, 0x86, 0x23, 0x77,
	0x2a, 0x4e, 0x3d, 0x72, 0xaa, 0xc0, 0x79, 0x11, 0xb4, 0xb6, 0xd7, 0x42, 0x94, 0xf0, 0xe7, 0x36,
	0x33, 0xfe, 0xcd, 0x37, 0xdf, 0x8c, 0x17, 0xfa, 0xef, 0x82, 0x2c, 0x60, 0xe7, 0xb8, 0x5a, 0x45,
	0x4a, 0x71, 0xce, 0xb2, 0x71, 0xc8, 0x55, 0x30, 0x66, 0xea, 0x82, 0x26, 0x6b, 0x54, 0x68, 0xdf,
	0xd3, 0x00, 0x6d, 0x00, 0x5a, 0x03, 0x4e, 0xef, 0x1c, 0xe5, 0x0a, 0xe5, 0xbc, 0xa4, 0x58, 0x95,
	0x54, 0x2d, 0x4e, 0x47, 0xa0, 0xc0, 0xaa, 0xae, 0xa3, 0xba, 0xda, 0x13, 0x88, 0x62, 0xc9, 0x59,
	0x99, 0x85, 0xe9, 0x1b, 0x16, 0xc4, 0x79, 0xfd, 0xe9, 0xfe, 0x16, 0x13, 0x82, 0xc7, 0x5c, 0x46,
	0xb5, 0xec, 0xf0, 0x13, 0x81, 0xa3, 0x99, 0x14, 0xcf, 0xd3, 0x70, 0x15, 0xa9, 0xb3, 0x35, 0x26,
	0x28, 0x83, 0xa5, 0xfd, 0x0a, 0x0e, 0x92, 0x34, 0xd4, 0x36, 0xca, 0xbc, 0x4b, 0x06, 0x64, 0x64,
	0x4d, 0x3a, 0xb4, 0x9a, 0x46, 0xcd, 0x34, 0xfa, 0x34, 0xce, 0xa7, 0xee, 0x97, 0x8f, 0xc7, 0x4e,
	0x6d, 0x55, 0x60, 0x66, 0x76, 0xa1, 0x27, 0x18, 0x2b, 0x1e, 0x2b, 0xdf, 0x4a, 0xd2, 0xb0, 0x11,
	0x76, 0x60, 0xaf, 0x12, 0xe5, 0xeb, 0xee, 0xad, 0x01, 0x19, 0xed, 0xfb, 0x4d, 0x6e, 0x4f, 0xe0,
	0xa0, 0x71, 0x3b, 0x8f, 0x16, 0xdd, 0x9d, 0x01, 0x19, 0xed, 0x4e, 0x0f, 0x8b, 0xeb, 0xbe, 0x75,
	0x62, 0xea, 0xde, 0xa9, 0x6f, 0x35, 0x90, 0xb7, 0x18, 0x3e, 0x83, 0xde, 0x0d, 0xf7, 0x3e, 0x97,
	0x09, 0xc6, 0x92, 0xdb, 0x0c, 0x2c, 0xb3, 0x81, 0xd6, 0x23, 0xa5, 0x5e, 0xbb, 0xb8, 0xee, 0x83,
	0x41, 0xbd, 0x53, 0x1f, 0x0c, 0xe2, 0x2d, 0x86, 0xef, 0x09, 0xdc, 0x9d, 0x49, 0xf1, 0x12, 0xd5,
	0xff, 0x37, 0xdb, 0x1d, 0xb8, 0x9d, 0xa1, 0x6a, 0xf6, 0xaa, 0x12, 0xfb, 0x09, 0xec, 0xeb, 0x60,
	0xae, 0xf2, 0x84, 0x97, 0x1b, 0xb5, 0x27, 0x03, 0xfa, 0xfb, 0xbf, 0x4f, 0xf5, 0xdc, 0x17, 0x79,
	0xc2, 0xfd, 0xbd, 0xac, 0x8e, 0x86, 0x47, 0x70, 0x58, 0x1b, 0x32, 0x5b, 0x4d, 0x3e, 0x13, 0xd8,
	0x99, 0x49, 0x61, 0xc7, 0xd0, 0xfe, 0xe5, 0xaf, 0x3d, 0xdc, 0x26, 0x7c, 0xe3, 0x44, 0xce, 0xf8,
	0x9f, 0xd1, 0xe6, 0x9a, 0x67, 0xb0, 0x5b, 0x1e, 0xa6, 0xff, 0x87, 0x56, 0x0d, 0x38, 0x0f, 0xfe,
	0x02, 0x18, 0xc5, 0xa9, 0x77, 0xf9, 0xdd, 0x6d, 0x5d, 0x16, 0x2e, 0xb9, 0x2a, 0x5c, 0xf2, 0xad,
	0x70, 0xc9, 0x87, 0x8d, 0xdb, 0xba, 0xda, 0xb8, 0xad, 0xaf, 0x1b, 0xb7, 0xf5, 0xfa, 0x91, 0x88,
	0xd4, 0xdb, 0x34, 0xd4, 0x3a, 0x4c, 0x0b, 0x1e, 0x2f, 0x83, 0x50, 0x96, 0x11, 0xbb, 0xf8, 0xe9,
	0x59, 0xeb, 0xc3, 0xca, 0xf0, 0x4e, 0xf9, 0x24, 0x1f, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe5,
	0x91, 0xbf, 0x3d, 0x7a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SubmitProposal defines a method for submitting a committee proposal
	SubmitProposal(ctx context.Context, in *MsgSubmitProposal, opts ...grpc.CallOption) (*MsgSubmitProposalResponse, error)
	// Vote defines a method for voting on a proposal
	Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitProposal(ctx context.Context, in *MsgSubmitProposal, opts ...grpc.CallOption) (*MsgSubmitProposalResponse, error) {
	out := new(MsgSubmitProposalResponse)
	err := c.cc.Invoke(ctx, "/kava.committee.v1beta1.Msg/SubmitProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error) {
	out := new(MsgVoteResponse)
	err := c.cc.Invoke(ctx, "/kava.committee.v1beta1.Msg/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SubmitProposal defines a method for submitting a committee proposal
	SubmitProposal(context.Context, *MsgSubmitProposal) (*MsgSubmitProposalResponse, error)
	// Vote defines a method for voting on a proposal
	Vote(context.Context, *MsgVote) (*MsgVoteResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitProposal(ctx context.Context, req *MsgSubmitProposal) (*MsgSubmitProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitProposal not implemented")
}
func (*UnimplementedMsgServer) Vote(ctx context.Context, req *MsgVote) (*MsgVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kava.committee.v1beta1.Msg/SubmitProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitProposal(ctx, req.(*MsgSubmitProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kava.committee.v1beta1.Msg/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Vote(ctx, req.(*MsgVote))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kava.committee.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitProposal",
			Handler:    _Msg_SubmitProposal_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Msg_Vote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kava/committee/v1beta1/tx.proto",
}

func (m *MsgSubmitProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CommitteeID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CommitteeID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x12
	}
	if m.PubProposal != nil {
		{
			size, err := m.PubProposal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitProposalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitProposalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitProposalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ProposalID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VoteType != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.VoteType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ProposalID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgVoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PubProposal != nil {
		l = m.PubProposal.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.CommitteeID != 0 {
		n += 1 + sovTx(uint64(m.CommitteeID))
	}
	return n
}

func (m *MsgSubmitProposalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalID != 0 {
		n += 1 + sovTx(uint64(m.ProposalID))
	}
	return n
}

func (m *MsgVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalID != 0 {
		n += 1 + sovTx(uint64(m.ProposalID))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.VoteType != 0 {
		n += 1 + sovTx(uint64(m.VoteType))
	}
	return n
}

func (m *MsgVoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubProposal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubProposal == nil {
				m.PubProposal = &types.Any{}
			}
			if err := m.PubProposal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitteeID", wireType)
			}
			m.CommitteeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommitteeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitProposalResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitProposalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitProposalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalID", wireType)
			}
			m.ProposalID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalID", wireType)
			}
			m.ProposalID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteType", wireType)
			}
			m.VoteType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteType |= VoteType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgVoteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgVoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
