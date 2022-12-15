// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/community/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryBalanceRequest defines the request type for querying x/community balance.
type QueryBalanceRequest struct {
}

func (m *QueryBalanceRequest) Reset()         { *m = QueryBalanceRequest{} }
func (m *QueryBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBalanceRequest) ProtoMessage()    {}
func (*QueryBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f236f06c43149273, []int{0}
}
func (m *QueryBalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalanceRequest.Merge(m, src)
}
func (m *QueryBalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalanceRequest proto.InternalMessageInfo

// QueryBalanceResponse defines the response type for querying x/community balance.
type QueryBalanceResponse struct {
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *QueryBalanceResponse) Reset()         { *m = QueryBalanceResponse{} }
func (m *QueryBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBalanceResponse) ProtoMessage()    {}
func (*QueryBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f236f06c43149273, []int{1}
}
func (m *QueryBalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalanceResponse.Merge(m, src)
}
func (m *QueryBalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalanceResponse proto.InternalMessageInfo

func (m *QueryBalanceResponse) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

// QueryLegacyCommunityPoolRequest defines the request type for querying the legacy community pool balance.
type QueryLegacyCommunityPoolRequest struct {
}

func (m *QueryLegacyCommunityPoolRequest) Reset()         { *m = QueryLegacyCommunityPoolRequest{} }
func (m *QueryLegacyCommunityPoolRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLegacyCommunityPoolRequest) ProtoMessage()    {}
func (*QueryLegacyCommunityPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f236f06c43149273, []int{2}
}
func (m *QueryLegacyCommunityPoolRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLegacyCommunityPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLegacyCommunityPoolRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLegacyCommunityPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLegacyCommunityPoolRequest.Merge(m, src)
}
func (m *QueryLegacyCommunityPoolRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLegacyCommunityPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLegacyCommunityPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLegacyCommunityPoolRequest proto.InternalMessageInfo

// QueryLegacyCommunityPoolResponse defines the response type for querying legacy community pool.
type QueryLegacyCommunityPoolResponse struct {
	Address string                                      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=balance,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"balance"`
}

func (m *QueryLegacyCommunityPoolResponse) Reset()         { *m = QueryLegacyCommunityPoolResponse{} }
func (m *QueryLegacyCommunityPoolResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLegacyCommunityPoolResponse) ProtoMessage()    {}
func (*QueryLegacyCommunityPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f236f06c43149273, []int{3}
}
func (m *QueryLegacyCommunityPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLegacyCommunityPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLegacyCommunityPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLegacyCommunityPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLegacyCommunityPoolResponse.Merge(m, src)
}
func (m *QueryLegacyCommunityPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLegacyCommunityPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLegacyCommunityPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLegacyCommunityPoolResponse proto.InternalMessageInfo

func (m *QueryLegacyCommunityPoolResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QueryLegacyCommunityPoolResponse) GetBalance() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Balance
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryBalanceRequest)(nil), "kava.community.v1beta1.QueryBalanceRequest")
	proto.RegisterType((*QueryBalanceResponse)(nil), "kava.community.v1beta1.QueryBalanceResponse")
	proto.RegisterType((*QueryLegacyCommunityPoolRequest)(nil), "kava.community.v1beta1.QueryLegacyCommunityPoolRequest")
	proto.RegisterType((*QueryLegacyCommunityPoolResponse)(nil), "kava.community.v1beta1.QueryLegacyCommunityPoolResponse")
}

func init() {
	proto.RegisterFile("kava/community/v1beta1/query.proto", fileDescriptor_f236f06c43149273)
}

var fileDescriptor_f236f06c43149273 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x41, 0x25, 0xe2, 0xd8, 0xae, 0x01, 0xa5, 0x51, 0xe5, 0xb4, 0x5e, 0x88, 0x54,
	0xe2, 0xa3, 0xa9, 0x10, 0xac, 0xa4, 0xb0, 0x31, 0x40, 0xd8, 0x58, 0xa2, 0xb3, 0x73, 0x32, 0x56,
	0x9c, 0x7b, 0xae, 0xef, 0x5c, 0xe1, 0x95, 0x1d, 0x09, 0x89, 0xff, 0x82, 0x99, 0x9d, 0x81, 0xa5,
	0x63, 0x05, 0x0b, 0x2c, 0x80, 0x12, 0xfe, 0x10, 0x74, 0x3f, 0x1c, 0x15, 0xc9, 0x46, 0x61, 0x4a,
	0xee, 0xde, 0xfb, 0xbe, 0xf7, 0xf9, 0xbe, 0x77, 0xc6, 0xfe, 0x82, 0x9d, 0x33, 0x1a, 0xc1, 0x72,
	0x59, 0x88, 0x44, 0x95, 0xf4, 0xfc, 0x38, 0xe4, 0x8a, 0x1d, 0xd3, 0xb3, 0x82, 0xe7, 0x65, 0x90,
	0xe5, 0xa0, 0x80, 0xdc, 0xd6, 0x39, 0xc1, 0x26, 0x27, 0x70, 0x39, 0x7d, 0x2f, 0x02, 0xb9, 0x04,
	0x49, 0x43, 0x26, 0xf9, 0x46, 0x18, 0x41, 0x22, 0xac, 0xae, 0xbf, 0x67, 0xe3, 0x33, 0x73, 0xa2,
	0xf6, 0xe0, 0x42, 0xdd, 0x18, 0x62, 0xb0, 0xf7, 0xfa, 0x9f, 0xbb, 0xdd, 0x8f, 0x01, 0xe2, 0x94,
	0x53, 0x96, 0x25, 0x94, 0x09, 0x01, 0x8a, 0xa9, 0x04, 0x84, 0xd3, 0xf8, 0xb7, 0xf0, 0xee, 0x73,
	0x4d, 0x35, 0x61, 0x29, 0x13, 0x11, 0x9f, 0xf2, 0xb3, 0x82, 0x4b, 0xe5, 0x97, 0xb8, 0xfb, 0xf7,
	0xb5, 0xcc, 0x40, 0x48, 0x4e, 0x18, 0xde, 0xd1, 0x2c, 0xb2, 0x87, 0x0e, 0xae, 0x0d, 0x6f, 0x8e,
	0xf7, 0x02, 0x07, 0xa0, 0x69, 0x2b, 0x0b, 0xc1, 0x29, 0x24, 0x62, 0x72, 0xef, 0xe2, 0xc7, 0xa0,
	0xf5, 0xe1, 0xe7, 0x60, 0x18, 0x27, 0xea, 0x55, 0x11, 0x6a, 0xa7, 0x8e, 0xd6, 0xfd, 0x8c, 0xe4,
	0x7c, 0x41, 0x55, 0x99, 0x71, 0x69, 0x04, 0x72, 0x6a, 0x2b, 0xfb, 0x87, 0x78, 0x60, 0x5a, 0x3f,
	0xe5, 0x31, 0x8b, 0xca, 0xd3, 0x6a, 0x40, 0xcf, 0x00, 0xd2, 0x8a, 0xee, 0x33, 0xc2, 0x07, 0xcd,
	0x39, 0x0e, 0x75, 0x8c, 0x3b, 0x6c, 0x3e, 0xcf, 0xb9, 0xd4, 0xb0, 0x68, 0x78, 0x63, 0xd2, 0xfb,
	0xf2, 0x71, 0xd4, 0x75, 0xbc, 0x8f, 0x6c, 0xe4, 0x85, 0xca, 0x13, 0x11, 0x4f, 0xab, 0x44, 0xb2,
	0xc0, 0x9d, 0xd0, 0x3a, 0xee, 0xb5, 0x8d, 0xc1, 0xfd, 0x5a, 0x83, 0x8f, 0x79, 0x64, 0x3c, 0x9e,
	0x38, 0x8f, 0x47, 0x5b, 0x78, 0x74, 0x1a, 0x39, 0xad, 0x3a, 0x8c, 0xbf, 0xb7, 0xf1, 0x8e, 0x71,
	0x41, 0xde, 0x22, 0xdc, 0x71, 0x93, 0x26, 0x47, 0x41, 0xfd, 0xc3, 0x08, 0x6a, 0xd6, 0xd4, 0xbf,
	0xbb, 0x5d, 0xb2, 0x9d, 0x88, 0x7f, 0xe7, 0xcd, 0xd7, 0xdf, 0xef, 0xdb, 0x87, 0x64, 0x40, 0x1b,
	0xde, 0xa7, 0x23, 0x23, 0x9f, 0x10, 0xde, 0xad, 0x19, 0x2d, 0x79, 0xf0, 0xcf, 0x76, 0xcd, 0x0b,
	0xeb, 0x3f, 0xfc, 0x7f, 0xa1, 0x63, 0xbe, 0x6f, 0x98, 0x29, 0x19, 0x35, 0x31, 0xa7, 0x46, 0x3c,
	0xdb, 0x04, 0x66, 0x19, 0x40, 0x3a, 0x79, 0x72, 0xb1, 0xf2, 0xd0, 0xe5, 0xca, 0x43, 0xbf, 0x56,
	0x1e, 0x7a, 0xb7, 0xf6, 0x5a, 0x97, 0x6b, 0xaf, 0xf5, 0x6d, 0xed, 0xb5, 0x5e, 0x5e, 0xdd, 0x95,
	0x2e, 0x39, 0x4a, 0x59, 0x28, 0x6d, 0xf1, 0xd7, 0x57, 0xca, 0x9b, 0xa5, 0x85, 0xd7, 0xcd, 0x47,
	0x72, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x64, 0x99, 0x49, 0xd1, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Balance queries the balance of all coins of x/community module.
	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
	// LegacyCommunityPool queries the balance of all coins of the legacy community pool.
	// The legacy community pool is a subaccount of the fee pool and has been replaced by x/community.
	LegacyCommunityPool(ctx context.Context, in *QueryLegacyCommunityPoolRequest, opts ...grpc.CallOption) (*QueryLegacyCommunityPoolResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, "/kava.community.v1beta1.Query/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LegacyCommunityPool(ctx context.Context, in *QueryLegacyCommunityPoolRequest, opts ...grpc.CallOption) (*QueryLegacyCommunityPoolResponse, error) {
	out := new(QueryLegacyCommunityPoolResponse)
	err := c.cc.Invoke(ctx, "/kava.community.v1beta1.Query/LegacyCommunityPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Balance queries the balance of all coins of x/community module.
	Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
	// LegacyCommunityPool queries the balance of all coins of the legacy community pool.
	// The legacy community pool is a subaccount of the fee pool and has been replaced by x/community.
	LegacyCommunityPool(context.Context, *QueryLegacyCommunityPoolRequest) (*QueryLegacyCommunityPoolResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Balance(ctx context.Context, req *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (*UnimplementedQueryServer) LegacyCommunityPool(ctx context.Context, req *QueryLegacyCommunityPoolRequest) (*QueryLegacyCommunityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LegacyCommunityPool not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kava.community.v1beta1.Query/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LegacyCommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLegacyCommunityPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LegacyCommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kava.community.v1beta1.Query/LegacyCommunityPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LegacyCommunityPool(ctx, req.(*QueryLegacyCommunityPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kava.community.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
		{
			MethodName: "LegacyCommunityPool",
			Handler:    _Query_LegacyCommunityPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kava/community/v1beta1/query.proto",
}

func (m *QueryBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryLegacyCommunityPoolRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLegacyCommunityPoolRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLegacyCommunityPoolRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryLegacyCommunityPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLegacyCommunityPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLegacyCommunityPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Balance) > 0 {
		for iNdEx := len(m.Balance) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Balance[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryBalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryBalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryLegacyCommunityPoolRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryLegacyCommunityPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.Balance) > 0 {
		for _, e := range m.Balance {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryBalanceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryBalanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryLegacyCommunityPoolRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryLegacyCommunityPoolRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLegacyCommunityPoolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryLegacyCommunityPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryLegacyCommunityPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLegacyCommunityPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balance = append(m.Balance, types.DecCoin{})
			if err := m.Balance[len(m.Balance)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
