// Code generated by protoc-gen-go. DO NOT EDIT.
// source: class6.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	class6.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package helloworld

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "class6.proto",
}

func init() { proto.RegisterFile("class6.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0x49, 0x2c,
	0x2e, 0x36, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x48, 0xcd, 0xc9, 0xc9, 0x2f,
	0xcf, 0x2f, 0xca, 0x49, 0x91, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8,
	0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8, 0x54, 0x52,
	0xe2, 0xe2, 0xf1, 0x00, 0xa9, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62,
	0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xd4, 0xb8,
	0xb8, 0xa0, 0x6a, 0x0a, 0x72, 0x2a, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3,
	0x61, 0x8a, 0x60, 0x5c, 0xa3, 0x14, 0x2e, 0x76, 0xf7, 0xa2, 0xd4, 0xd4, 0x92, 0xd4, 0x22, 0xa1,
	0x48, 0x2e, 0x8e, 0xe0, 0xc4, 0x4a, 0xb0, 0x2e, 0x21, 0x09, 0x3d, 0x84, 0x6b, 0xf4, 0x90, 0x2d,
	0x93, 0x12, 0xc3, 0x22, 0x53, 0x90, 0x53, 0xa9, 0x24, 0xdd, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x51,
	0x25, 0x01, 0xfd, 0x32, 0x43, 0xfd, 0xd4, 0x8a, 0xc4, 0xdc, 0x82, 0x9c, 0x54, 0xfd, 0xd4, 0xe4,
	0x8c, 0x7c, 0x2b, 0x46, 0xad, 0x24, 0x36, 0xb0, 0xc3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xcf, 0x1c, 0x80, 0x9e, 0xf2, 0x00, 0x00, 0x00,
}
