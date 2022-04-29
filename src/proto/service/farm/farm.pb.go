// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/farm.proto

package farm

import (
	context "context"
	device "farm/src/proto/messages/device"
	user "farm/src/proto/messages/user"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() { proto.RegisterFile("services/farm.proto", fileDescriptor_025aa1765b083615) }

var fileDescriptor_025aa1765b083615 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x4f, 0x4b, 0x2c, 0xca, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x81, 0x0a, 0xea, 0x81, 0xc4, 0xa4, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13,
	0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0x6a,
	0xa5, 0x84, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x8b, 0xf5, 0x4b, 0x8b, 0x53, 0x8b, 0xa0,
	0x82, 0xa2, 0x70, 0xc1, 0x94, 0x54, 0xb0, 0x41, 0x60, 0x61, 0xa3, 0xab, 0x8c, 0x5c, 0x2c, 0x6e,
	0x89, 0x45, 0xb9, 0x42, 0xa1, 0x5c, 0xac, 0x3e, 0xf9, 0xe9, 0x99, 0x79, 0x42, 0xd2, 0x7a, 0x30,
	0x95, 0x7a, 0x60, 0xed, 0x60, 0xd1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x29, 0x19, 0xec,
	0x92, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x4a, 0x82, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0xe2, 0x56,
	0x62, 0xd3, 0xcf, 0x01, 0x89, 0x5b, 0x31, 0x6a, 0x09, 0x95, 0x71, 0xf1, 0x38, 0x17, 0xa5, 0x26,
	0x96, 0xa4, 0xba, 0x80, 0x6d, 0x15, 0x52, 0x41, 0x18, 0x00, 0x75, 0x07, 0xb2, 0x34, 0xcc, 0x1a,
	0x55, 0x02, 0xaa, 0xa0, 0xf6, 0x49, 0x82, 0xed, 0x13, 0x56, 0xe2, 0x83, 0xfa, 0x49, 0x3f, 0x19,
	0xac, 0xca, 0x8a, 0x51, 0xcb, 0x49, 0x36, 0x4a, 0x1a, 0x14, 0x52, 0xfa, 0xc5, 0x45, 0xc9, 0xfa,
	0x60, 0x9f, 0xea, 0x43, 0x03, 0x10, 0x1c, 0xa8, 0x49, 0x6c, 0x60, 0x31, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x81, 0xfc, 0x42, 0xbb, 0x6c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FarmClient is the client API for Farm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FarmClient interface {
	Login(ctx context.Context, in *user.LoginRequest, opts ...grpc.CallOption) (*user.LoginResponse, error)
	CreateDevice(ctx context.Context, in *device.CreateDeviceRequest, opts ...grpc.CallOption) (*device.CreateDeviceResponse, error)
}

type farmClient struct {
	cc *grpc.ClientConn
}

func NewFarmClient(cc *grpc.ClientConn) FarmClient {
	return &farmClient{cc}
}

func (c *farmClient) Login(ctx context.Context, in *user.LoginRequest, opts ...grpc.CallOption) (*user.LoginResponse, error) {
	out := new(user.LoginResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) CreateDevice(ctx context.Context, in *device.CreateDeviceRequest, opts ...grpc.CallOption) (*device.CreateDeviceResponse, error) {
	out := new(device.CreateDeviceResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/CreateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FarmServer is the server API for Farm service.
type FarmServer interface {
	Login(context.Context, *user.LoginRequest) (*user.LoginResponse, error)
	CreateDevice(context.Context, *device.CreateDeviceRequest) (*device.CreateDeviceResponse, error)
}

// UnimplementedFarmServer can be embedded to have forward compatible implementations.
type UnimplementedFarmServer struct {
}

func (*UnimplementedFarmServer) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedFarmServer) CreateDevice(ctx context.Context, req *device.CreateDeviceRequest) (*device.CreateDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}

func RegisterFarmServer(s *grpc.Server, srv FarmServer) {
	s.RegisterService(&_Farm_serviceDesc, srv)
}

func _Farm_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).Login(ctx, req.(*user.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(device.CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).CreateDevice(ctx, req.(*device.CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Farm_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.farm.Farm",
	HandlerType: (*FarmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Farm_Login_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _Farm_CreateDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/farm.proto",
}
