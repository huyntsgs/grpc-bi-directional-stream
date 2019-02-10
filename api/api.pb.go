// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Request
	Response
*/
package protobuf

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

type Request struct {
	Number    int32  `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Sig       []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	PublicKey []byte `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	R         []byte `protobuf:"bytes,4,opt,name=r,proto3" json:"r,omitempty"`
	S         []byte `protobuf:"bytes,5,opt,name=s,proto3" json:"s,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Request) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *Request) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Request) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *Request) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

type Response struct {
	Res int32 `protobuf:"varint,1,opt,name=res" json:"res,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetRes() int32 {
	if m != nil {
		return m.Res
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "protobuf.Request")
	proto.RegisterType((*Response)(nil), "protobuf.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Math service

type MathClient interface {
	FindMaxNumber(ctx context.Context, opts ...grpc.CallOption) (Math_FindMaxNumberClient, error)
}

type mathClient struct {
	cc *grpc.ClientConn
}

func NewMathClient(cc *grpc.ClientConn) MathClient {
	return &mathClient{cc}
}

func (c *mathClient) FindMaxNumber(ctx context.Context, opts ...grpc.CallOption) (Math_FindMaxNumberClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Math_serviceDesc.Streams[0], c.cc, "/protobuf.Math/FindMaxNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &mathFindMaxNumberClient{stream}
	return x, nil
}

type Math_FindMaxNumberClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type mathFindMaxNumberClient struct {
	grpc.ClientStream
}

func (x *mathFindMaxNumberClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mathFindMaxNumberClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Math service

type MathServer interface {
	FindMaxNumber(Math_FindMaxNumberServer) error
}

func RegisterMathServer(s *grpc.Server, srv MathServer) {
	s.RegisterService(&_Math_serviceDesc, srv)
}

func _Math_FindMaxNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MathServer).FindMaxNumber(&mathFindMaxNumberServer{stream})
}

type Math_FindMaxNumberServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type mathFindMaxNumberServer struct {
	grpc.ServerStream
}

func (x *mathFindMaxNumberServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mathFindMaxNumberServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Math_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Math",
	HandlerType: (*MathServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindMaxNumber",
			Handler:       _Math_FindMaxNumber_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x4a, 0xb9, 0x5c, 0xec,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0x79, 0xa5, 0xb9, 0x49, 0xa9,
	0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x50, 0x9e, 0x90, 0x00, 0x17, 0x73, 0x71, 0x66,
	0xba, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x88, 0x29, 0x24, 0xc3, 0xc5, 0x59, 0x50, 0x9a,
	0x94, 0x93, 0x99, 0xec, 0x9d, 0x5a, 0x29, 0xc1, 0x0c, 0x16, 0x47, 0x08, 0x08, 0xf1, 0x70, 0x31,
	0x16, 0x49, 0xb0, 0x80, 0x45, 0x19, 0x8b, 0x40, 0xbc, 0x62, 0x09, 0x56, 0x08, 0xaf, 0x58, 0x49,
	0x86, 0x8b, 0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0x64, 0x6e, 0x51, 0x6a, 0x31,
	0xd4, 0x32, 0x10, 0xd3, 0xc8, 0x89, 0x8b, 0xc5, 0x37, 0xb1, 0x24, 0x43, 0xc8, 0x8a, 0x8b, 0xd7,
	0x2d, 0x33, 0x2f, 0xc5, 0x37, 0xb1, 0xc2, 0x0f, 0xe2, 0x04, 0x41, 0x3d, 0x98, 0x83, 0xf5, 0xa0,
	0xae, 0x95, 0x12, 0x42, 0x16, 0x82, 0x98, 0xa8, 0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06, 0x16, 0x36,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xe7, 0x93, 0xb5, 0xee, 0x00, 0x00, 0x00,
}
