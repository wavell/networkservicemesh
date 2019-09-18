// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networkservice.proto

package networkservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	connection "github.com/networkservicemesh/networkservicemesh/controlplane/api/remote/connection"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NetworkServiceRequest struct {
	Connection           *connection.Connection  `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	MechanismPreferences []*connection.Mechanism `protobuf:"bytes,2,rep,name=mechanism_preferences,json=mechanismPreferences,proto3" json:"mechanism_preferences,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *NetworkServiceRequest) Reset()         { *m = NetworkServiceRequest{} }
func (m *NetworkServiceRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceRequest) ProtoMessage()    {}
func (*NetworkServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_361e8247d5a9945c, []int{0}
}

func (m *NetworkServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceRequest.Unmarshal(m, b)
}
func (m *NetworkServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceRequest.Marshal(b, m, deterministic)
}
func (m *NetworkServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceRequest.Merge(m, src)
}
func (m *NetworkServiceRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceRequest.Size(m)
}
func (m *NetworkServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceRequest proto.InternalMessageInfo

func (m *NetworkServiceRequest) GetConnection() *connection.Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *NetworkServiceRequest) GetMechanismPreferences() []*connection.Mechanism {
	if m != nil {
		return m.MechanismPreferences
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkServiceRequest)(nil), "remote.networkservice.NetworkServiceRequest")
}

func init() { proto.RegisterFile("networkservice.proto", fileDescriptor_361e8247d5a9945c) }

var fileDescriptor_361e8247d5a9945c = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xa6, 0x8a, 0x0a, 0x59, 0x58, 0x24, 0x6c, 0xa5, 0x14, 0x85, 0xc5, 0xd3, 0x1e, 0x24, 0x85,
	0x7a, 0xd6, 0x83, 0x8b, 0x47, 0x45, 0x2b, 0x78, 0xf0, 0x22, 0x6d, 0x9c, 0x6d, 0x83, 0x4d, 0x26,
	0x26, 0xa9, 0xb2, 0x6f, 0xa4, 0x6f, 0x29, 0xbb, 0xd9, 0x9f, 0x16, 0x4a, 0x2f, 0x61, 0x98, 0x99,
	0xef, 0x6f, 0x42, 0x26, 0x0a, 0xdc, 0x0f, 0x9a, 0x4f, 0x0b, 0xe6, 0x5b, 0x70, 0x60, 0xda, 0xa0,
	0x43, 0x1a, 0x1a, 0x90, 0xe8, 0x80, 0x75, 0x87, 0xf1, 0x47, 0x29, 0x5c, 0xd5, 0x14, 0x8c, 0xa3,
	0x4c, 0xba, 0x23, 0x09, 0xb6, 0xea, 0x6b, 0x71, 0x54, 0xce, 0x60, 0xad, 0xeb, 0x5c, 0x41, 0x92,
	0x6b, 0x91, 0x78, 0xe2, 0x55, 0x5f, 0x01, 0x77, 0x02, 0x55, 0xab, 0xf4, 0xe2, 0x71, 0xa4, 0xdd,
	0x52, 0x83, 0x4d, 0x40, 0x6a, 0xb7, 0xf4, 0xaf, 0x9f, 0x5c, 0xfe, 0x05, 0x24, 0x7c, 0xf4, 0x22,
	0x2f, 0x5e, 0x24, 0x83, 0xaf, 0x06, 0xac, 0xa3, 0x37, 0x84, 0xec, 0x79, 0xa2, 0x60, 0x1a, 0xcc,
	0x46, 0xe9, 0x05, 0xdb, 0xa4, 0x68, 0x29, 0xcc, 0x77, 0x65, 0xd6, 0x02, 0xd0, 0x67, 0x12, 0x4a,
	0xe0, 0x55, 0xae, 0x84, 0x95, 0xef, 0xda, 0xc0, 0x02, 0x0c, 0x28, 0x0e, 0x36, 0x3a, 0x98, 0x1e,
	0xce, 0x46, 0xe9, 0x79, 0x0f, 0xd3, 0xc3, 0x76, 0x3f, 0x9b, 0xec, 0xa0, 0x4f, 0x7b, 0x64, 0xfa,
	0x1b, 0x90, 0x71, 0xd7, 0x2b, 0x7d, 0x25, 0x27, 0x5b, 0xbf, 0x57, 0xac, 0xf7, 0xc2, 0xac, 0x37,
	0x5d, 0x3c, 0x9c, 0x84, 0xde, 0x92, 0xa3, 0x79, 0x8d, 0x16, 0xe8, 0xf0, 0x5e, 0x7c, 0xc6, 0x4a,
	0xc4, 0xb2, 0xde, 0x7c, 0x72, 0xd1, 0x2c, 0xd8, 0xfd, 0xea, 0xb8, 0x77, 0xa7, 0x6f, 0xe3, 0xae,
	0x8d, 0xe2, 0x78, 0xbd, 0x71, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x51, 0x11, 0xa2, 0x1e,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceClient interface {
	Request(ctx context.Context, in *NetworkServiceRequest, opts ...grpc.CallOption) (*connection.Connection, error)
	Close(ctx context.Context, in *connection.Connection, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) Request(ctx context.Context, in *NetworkServiceRequest, opts ...grpc.CallOption) (*connection.Connection, error) {
	out := new(connection.Connection)
	err := c.cc.Invoke(ctx, "/remote.networkservice.NetworkService/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Close(ctx context.Context, in *connection.Connection, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/remote.networkservice.NetworkService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
type NetworkServiceServer interface {
	Request(context.Context, *NetworkServiceRequest) (*connection.Connection, error)
	Close(context.Context, *connection.Connection) (*empty.Empty, error)
}

// UnimplementedNetworkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkServiceServer struct {
}

func (*UnimplementedNetworkServiceServer) Request(ctx context.Context, req *NetworkServiceRequest) (*connection.Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (*UnimplementedNetworkServiceServer) Close(ctx context.Context, req *connection.Connection) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.networkservice.NetworkService/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Request(ctx, req.(*NetworkServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(connection.Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.networkservice.NetworkService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Close(ctx, req.(*connection.Connection))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remote.networkservice.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _NetworkService_Request_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _NetworkService_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networkservice.proto",
}