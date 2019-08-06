// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registry.proto

package plugins

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type PluginCapability int32

const (
	PluginCapability_CONNECTION PluginCapability = 0
)

var PluginCapability_name = map[int32]string{
	0: "CONNECTION",
}

var PluginCapability_value = map[string]int32{
	"CONNECTION": 0,
}

func (x PluginCapability) String() string {
	return proto.EnumName(PluginCapability_name, int32(x))
}

func (PluginCapability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{0}
}

type PluginInfo struct {
	Endpoint             string             `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Capabilities         []PluginCapability `protobuf:"varint,2,rep,packed,name=capabilities,proto3,enum=plugins.PluginCapability" json:"capabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PluginInfo) Reset()         { *m = PluginInfo{} }
func (m *PluginInfo) String() string { return proto.CompactTextString(m) }
func (*PluginInfo) ProtoMessage()    {}
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{0}
}

func (m *PluginInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginInfo.Unmarshal(m, b)
}
func (m *PluginInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginInfo.Marshal(b, m, deterministic)
}
func (m *PluginInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginInfo.Merge(m, src)
}
func (m *PluginInfo) XXX_Size() int {
	return xxx_messageInfo_PluginInfo.Size(m)
}
func (m *PluginInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PluginInfo proto.InternalMessageInfo

func (m *PluginInfo) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *PluginInfo) GetCapabilities() []PluginCapability {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

func init() {
	proto.RegisterEnum("plugins.PluginCapability", PluginCapability_name, PluginCapability_value)
	proto.RegisterType((*PluginInfo)(nil), "plugins.PluginInfo")
}

func init() { proto.RegisterFile("registry.proto", fileDescriptor_41af05d40a615591) }

var fileDescriptor_41af05d40a615591 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0xc8, 0x29, 0x4d,
	0xcf, 0xcc, 0x2b, 0x96, 0x92, 0x28, 0x28, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0xcd, 0x2d, 0x28,
	0xa9, 0x84, 0x90, 0x10, 0x25, 0x4a, 0xe9, 0x5c, 0x5c, 0x01, 0x60, 0x45, 0x9e, 0x79, 0x69, 0xf9,
	0x42, 0x52, 0x5c, 0x1c, 0xa9, 0x79, 0x29, 0x05, 0xf9, 0x99, 0x79, 0x25, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x2d, 0x17, 0x4f, 0x72, 0x62, 0x41, 0x62, 0x52, 0x66, 0x4e,
	0x66, 0x49, 0x66, 0x6a, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0x9f, 0x91, 0xa4, 0x1e, 0xd4, 0x0e,
	0x3d, 0x88, 0x31, 0xce, 0x30, 0x25, 0x95, 0x41, 0x28, 0xca, 0xb5, 0x94, 0xb8, 0x04, 0xd0, 0x55,
	0x08, 0xf1, 0x71, 0x71, 0x39, 0xfb, 0xfb, 0xf9, 0xb9, 0x3a, 0x87, 0x78, 0xfa, 0xfb, 0x09, 0x30,
	0x18, 0x79, 0x72, 0xf1, 0x41, 0xd4, 0x04, 0x41, 0xfd, 0x21, 0x64, 0xce, 0xc5, 0x01, 0x61, 0xa7,
	0x16, 0x09, 0x09, 0xa3, 0x59, 0x05, 0x72, 0xb1, 0x94, 0x98, 0x5e, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e,
	0x2a, 0xc4, 0x3b, 0x49, 0xa5, 0x69, 0x7a, 0xae, 0x20, 0xdf, 0x25, 0xb1, 0x81, 0xf9, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xa3, 0x24, 0xd2, 0x13, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PluginRegistryClient is the client API for PluginRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PluginRegistryClient interface {
	Register(ctx context.Context, in *PluginInfo, opts ...grpc.CallOption) (*empty.Empty, error)
}

type pluginRegistryClient struct {
	cc *grpc.ClientConn
}

func NewPluginRegistryClient(cc *grpc.ClientConn) PluginRegistryClient {
	return &pluginRegistryClient{cc}
}

func (c *pluginRegistryClient) Register(ctx context.Context, in *PluginInfo, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/plugins.PluginRegistry/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginRegistryServer is the server API for PluginRegistry service.
type PluginRegistryServer interface {
	Register(context.Context, *PluginInfo) (*empty.Empty, error)
}

// UnimplementedPluginRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedPluginRegistryServer struct {
}

func (*UnimplementedPluginRegistryServer) Register(ctx context.Context, req *PluginInfo) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterPluginRegistryServer(s *grpc.Server, srv PluginRegistryServer) {
	s.RegisterService(&_PluginRegistry_serviceDesc, srv)
}

func _PluginRegistry_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginRegistryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.PluginRegistry/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginRegistryServer).Register(ctx, req.(*PluginInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _PluginRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "plugins.PluginRegistry",
	HandlerType: (*PluginRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _PluginRegistry_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry.proto",
}