// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/pricefeed/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgPostPrice represents a method for creating a new post price
type MsgPostPrice struct {
	// address of client
	From     string                                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	MarketID string                                 `protobuf:"bytes,2,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	Price    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
	Expiry   time.Time                              `protobuf:"bytes,4,opt,name=expiry,proto3,stdtime" json:"expiry"`
}

func (m *MsgPostPrice) Reset()         { *m = MsgPostPrice{} }
func (m *MsgPostPrice) String() string { return proto.CompactTextString(m) }
func (*MsgPostPrice) ProtoMessage()    {}
func (*MsgPostPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd93c8e4685da16, []int{0}
}
func (m *MsgPostPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPostPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPostPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPostPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPostPrice.Merge(m, src)
}
func (m *MsgPostPrice) XXX_Size() int {
	return m.Size()
}
func (m *MsgPostPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPostPrice.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPostPrice proto.InternalMessageInfo

// MsgPostPriceResponse defines the Msg/PostPrice response type.
type MsgPostPriceResponse struct {
}

func (m *MsgPostPriceResponse) Reset()         { *m = MsgPostPriceResponse{} }
func (m *MsgPostPriceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPostPriceResponse) ProtoMessage()    {}
func (*MsgPostPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd93c8e4685da16, []int{1}
}
func (m *MsgPostPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPostPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPostPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPostPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPostPriceResponse.Merge(m, src)
}
func (m *MsgPostPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPostPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPostPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPostPriceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgPostPrice)(nil), "kava.pricefeed.v1beta1.MsgPostPrice")
	proto.RegisterType((*MsgPostPriceResponse)(nil), "kava.pricefeed.v1beta1.MsgPostPriceResponse")
}

func init() { proto.RegisterFile("kava/pricefeed/v1beta1/tx.proto", fileDescriptor_afd93c8e4685da16) }

var fileDescriptor_afd93c8e4685da16 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbd, 0x6e, 0xea, 0x30,
	0x14, 0x80, 0xe3, 0x0b, 0x17, 0x05, 0x5f, 0xa6, 0x08, 0xa1, 0x28, 0x43, 0x8c, 0xd0, 0x55, 0x45,
	0xd5, 0x62, 0x0b, 0xba, 0x55, 0x9d, 0x10, 0x1d, 0x18, 0x90, 0x50, 0xd4, 0xa9, 0x0b, 0xca, 0x8f,
	0x93, 0x46, 0x90, 0x3a, 0x8a, 0x0d, 0x82, 0x37, 0xe8, 0xc8, 0x23, 0xf0, 0x38, 0x8c, 0x8c, 0x55,
	0x2b, 0xd1, 0x2a, 0x2c, 0x7d, 0x8c, 0x2a, 0x0e, 0xb4, 0x19, 0x3a, 0x74, 0xca, 0x89, 0xcf, 0x77,
	0xce, 0xf1, 0x77, 0x0c, 0xd1, 0xd4, 0x5e, 0xd8, 0x24, 0x4e, 0x42, 0x97, 0xfa, 0x94, 0x7a, 0x64,
	0xd1, 0x75, 0xa8, 0xb0, 0xbb, 0x44, 0x2c, 0x71, 0x9c, 0x30, 0xc1, 0xb4, 0x46, 0x06, 0xe0, 0x2f,
	0x00, 0x1f, 0x01, 0xa3, 0x1e, 0xb0, 0x80, 0x49, 0x84, 0x64, 0x51, 0x4e, 0x1b, 0x28, 0x60, 0x2c,
	0x98, 0x51, 0x22, 0xff, 0x9c, 0xb9, 0x4f, 0x44, 0x18, 0x51, 0x2e, 0xec, 0x28, 0xce, 0x81, 0xd6,
	0x2b, 0x80, 0xb5, 0x11, 0x0f, 0xc6, 0x8c, 0x8b, 0x71, 0xd6, 0x53, 0xd3, 0x60, 0xd9, 0x4f, 0x58,
	0xa4, 0x83, 0x26, 0x68, 0x57, 0x2d, 0x19, 0x6b, 0xe7, 0xb0, 0x1a, 0xd9, 0xc9, 0x94, 0x8a, 0x49,
	0xe8, 0xe9, 0x7f, 0xb2, 0x44, 0xbf, 0x96, 0xee, 0x91, 0x3a, 0x92, 0x87, 0xc3, 0x81, 0xa5, 0xe6,
	0xe9, 0xa1, 0xa7, 0x0d, 0xe0, 0x5f, 0x79, 0x37, 0xbd, 0x24, 0x31, 0xbc, 0xdd, 0x23, 0xe5, 0x65,
	0x8f, 0xce, 0x82, 0x50, 0x3c, 0xcc, 0x1d, 0xec, 0xb2, 0x88, 0xb8, 0x8c, 0x47, 0x8c, 0x1f, 0x3f,
	0x1d, 0xee, 0x4d, 0x89, 0x58, 0xc5, 0x94, 0xe3, 0x01, 0x75, 0xad, 0xbc, 0x58, 0xbb, 0x81, 0x15,
	0xba, 0x8c, 0xc3, 0x64, 0xa5, 0x97, 0x9b, 0xa0, 0xfd, 0xaf, 0x67, 0xe0, 0xdc, 0x03, 0x9f, 0x3c,
	0xf0, 0xdd, 0xc9, 0xa3, 0xaf, 0x66, 0x23, 0xd6, 0x6f, 0x08, 0x58, 0xc7, 0x9a, 0x6b, 0xf5, 0x69,
	0x83, 0x94, 0x8f, 0x0d, 0x52, 0x5a, 0x0d, 0x58, 0x2f, 0xca, 0x59, 0x94, 0xc7, 0xec, 0x91, 0xd3,
	0x9e, 0x0f, 0x4b, 0x23, 0x1e, 0x68, 0x13, 0x58, 0xfd, 0x16, 0xff, 0x8f, 0x7f, 0xde, 0x2c, 0x2e,
	0x76, 0x30, 0x2e, 0x7f, 0x43, 0x9d, 0xe6, 0xf4, 0x6f, 0xb7, 0xa9, 0x09, 0x76, 0xa9, 0x09, 0xde,
	0x53, 0x13, 0xac, 0x0f, 0xa6, 0xb2, 0x3b, 0x98, 0xca, 0xf3, 0xc1, 0x54, 0xee, 0x2f, 0x0a, 0x0b,
	0xc9, 0x3a, 0x76, 0x66, 0xb6, 0xc3, 0x65, 0x44, 0x96, 0x85, 0xe7, 0x97, 0x9b, 0x71, 0x2a, 0x52,
	0xfb, 0xea, 0x33, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xaa, 0xa5, 0x0a, 0x1d, 0x02, 0x00, 0x00,
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
	// PostPrice defines a method for creating a new post price
	PostPrice(ctx context.Context, in *MsgPostPrice, opts ...grpc.CallOption) (*MsgPostPriceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PostPrice(ctx context.Context, in *MsgPostPrice, opts ...grpc.CallOption) (*MsgPostPriceResponse, error) {
	out := new(MsgPostPriceResponse)
	err := c.cc.Invoke(ctx, "/kava.pricefeed.v1beta1.Msg/PostPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// PostPrice defines a method for creating a new post price
	PostPrice(context.Context, *MsgPostPrice) (*MsgPostPriceResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PostPrice(ctx context.Context, req *MsgPostPrice) (*MsgPostPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostPrice not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PostPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPostPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PostPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kava.pricefeed.v1beta1.Msg/PostPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PostPrice(ctx, req.(*MsgPostPrice))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kava.pricefeed.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostPrice",
			Handler:    _Msg_PostPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kava/pricefeed/v1beta1/tx.proto",
}

func (m *MsgPostPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPostPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPostPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiry, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiry):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTx(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.MarketID) > 0 {
		i -= len(m.MarketID)
		copy(dAtA[i:], m.MarketID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MarketID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPostPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPostPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPostPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgPostPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MarketID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovTx(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiry)
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgPostPriceResponse) Size() (n int) {
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
func (m *MsgPostPrice) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgPostPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPostPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketID", wireType)
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
			m.MarketID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
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
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiry, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *MsgPostPriceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgPostPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPostPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
