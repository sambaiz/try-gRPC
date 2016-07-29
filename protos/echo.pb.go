// Code generated by protoc-gen-go.
// source: protos/echo.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	protos/echo.proto

It has these top-level messages:
	EchoRequest
	EchoReply
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type EchoRequest struct {
	Say string `protobuf:"bytes,1,opt,name=say" json:"say,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EchoReply struct {
	Ret string `protobuf:"bytes,1,opt,name=ret" json:"ret,omitempty"`
}

func (m *EchoReply) Reset()                    { *m = EchoReply{} }
func (m *EchoReply) String() string            { return proto.CompactTextString(m) }
func (*EchoReply) ProtoMessage()               {}
func (*EchoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*EchoRequest)(nil), "protos.EchoRequest")
	proto.RegisterType((*EchoReply)(nil), "protos.EchoReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Echo service

type EchoClient interface {
	RetEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) RetEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := grpc.Invoke(ctx, "/protos.Echo/RetEcho", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Echo service

type EchoServer interface {
	RetEcho(context.Context, *EchoRequest) (*EchoReply, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_RetEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).RetEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Echo/RetEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).RetEcho(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetEcho",
			Handler:    _Echo_RetEcho_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("protos/echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x4d, 0xce, 0xc8, 0xd7, 0x03, 0xb3, 0x85, 0xd8, 0x20, 0x42, 0x4a, 0xf2,
	0x5c, 0xdc, 0xae, 0x40, 0xd1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x01, 0x2e, 0xe6,
	0xe2, 0xc4, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x49, 0x96, 0x8b, 0x13,
	0xa2, 0xa0, 0x20, 0xa7, 0x12, 0x24, 0x5d, 0x94, 0x5a, 0x02, 0x93, 0x06, 0x32, 0x8d, 0xac, 0xb9,
	0x58, 0x40, 0xd2, 0x42, 0xc6, 0x5c, 0xec, 0x41, 0xa9, 0x25, 0x60, 0xa6, 0x30, 0xc4, 0x8a, 0x62,
	0x3d, 0x24, 0x83, 0xa5, 0x04, 0x51, 0x05, 0x81, 0x86, 0x29, 0x31, 0x38, 0xc9, 0x70, 0x49, 0xe5,
	0xa5, 0x96, 0xe8, 0x15, 0x27, 0xe6, 0x26, 0x25, 0x66, 0x56, 0xe9, 0x95, 0x14, 0x55, 0xa6, 0x17,
	0x15, 0x24, 0x43, 0x55, 0x26, 0x41, 0x9c, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xf7,
	0x53, 0x37, 0xbe, 0x00, 0x00, 0x00,
}
