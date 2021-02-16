// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file1.proto

package multifiles

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

// The request message containing the user's name.
type Request1 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request1) Reset()         { *m = Request1{} }
func (m *Request1) String() string { return proto.CompactTextString(m) }
func (*Request1) ProtoMessage()    {}
func (*Request1) Descriptor() ([]byte, []int) {
	return fileDescriptor_file1_b87e0295d2d0c495, []int{0}
}
func (m *Request1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request1.Unmarshal(m, b)
}
func (m *Request1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request1.Marshal(b, m, deterministic)
}
func (dst *Request1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request1.Merge(dst, src)
}
func (m *Request1) XXX_Size() int {
	return xxx_messageInfo_Request1.Size(m)
}
func (m *Request1) XXX_DiscardUnknown() {
	xxx_messageInfo_Request1.DiscardUnknown(m)
}

var xxx_messageInfo_Request1 proto.InternalMessageInfo

func (m *Request1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type Reply1 struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ReturnCode           int32    `protobuf:"varint,2,opt,name=return_code,json=returnCode,proto3" json:"return_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply1) Reset()         { *m = Reply1{} }
func (m *Reply1) String() string { return proto.CompactTextString(m) }
func (*Reply1) ProtoMessage()    {}
func (*Reply1) Descriptor() ([]byte, []int) {
	return fileDescriptor_file1_b87e0295d2d0c495, []int{1}
}
func (m *Reply1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply1.Unmarshal(m, b)
}
func (m *Reply1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply1.Marshal(b, m, deterministic)
}
func (dst *Reply1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply1.Merge(dst, src)
}
func (m *Reply1) XXX_Size() int {
	return xxx_messageInfo_Reply1.Size(m)
}
func (m *Reply1) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply1.DiscardUnknown(m)
}

var xxx_messageInfo_Reply1 proto.InternalMessageInfo

func (m *Reply1) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Reply1) GetReturnCode() int32 {
	if m != nil {
		return m.ReturnCode
	}
	return 0
}

func init() {
	proto.RegisterType((*Request1)(nil), "multifiles.Request1")
	proto.RegisterType((*Reply1)(nil), "multifiles.Reply1")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Gripmock1Client is the client API for Gripmock1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Gripmock1Client interface {
	// simple unary method
	SayHello(ctx context.Context, in *Request1, opts ...grpc.CallOption) (*Reply1, error)
}

type gripmock1Client struct {
	cc *grpc.ClientConn
}

func NewGripmock1Client(cc *grpc.ClientConn) Gripmock1Client {
	return &gripmock1Client{cc}
}

func (c *gripmock1Client) SayHello(ctx context.Context, in *Request1, opts ...grpc.CallOption) (*Reply1, error) {
	out := new(Reply1)
	err := c.cc.Invoke(ctx, "/multifiles.Gripmock1/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Gripmock1Server is the server API for Gripmock1 service.
type Gripmock1Server interface {
	// simple unary method
	SayHello(context.Context, *Request1) (*Reply1, error)
}

func RegisterGripmock1Server(s *grpc.Server, srv Gripmock1Server) {
	s.RegisterService(&_Gripmock1_serviceDesc, srv)
}

func _Gripmock1_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Gripmock1Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multifiles.Gripmock1/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Gripmock1Server).SayHello(ctx, req.(*Request1))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gripmock1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "multifiles.Gripmock1",
	HandlerType: (*Gripmock1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Gripmock1_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file1.proto",
}

func init() { proto.RegisterFile("file1.proto", fileDescriptor_file1_b87e0295d2d0c495) }

var fileDescriptor_file1_b87e0295d2d0c495 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcb, 0xcc, 0x49,
	0x35, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x2d, 0xcd, 0x29, 0xc9, 0x04, 0x89,
	0x14, 0x2b, 0xc9, 0x71, 0x71, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x18, 0x0a, 0x09, 0x71,
	0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0xce,
	0x5c, 0x6c, 0x41, 0xa9, 0x05, 0x39, 0x95, 0x86, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5,
	0x89, 0xe9, 0x30, 0x05, 0x30, 0xae, 0x90, 0x3c, 0x17, 0x77, 0x51, 0x6a, 0x49, 0x69, 0x51, 0x5e,
	0x7c, 0x72, 0x7e, 0x4a, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x17, 0x44, 0xc8, 0x39,
	0x3f, 0x25, 0xd5, 0xc8, 0x91, 0x8b, 0xd3, 0xbd, 0x28, 0xb3, 0x20, 0x37, 0x3f, 0x39, 0xdb, 0x50,
	0xc8, 0x84, 0x8b, 0x23, 0x38, 0xb1, 0xd2, 0x23, 0x35, 0x27, 0x27, 0x5f, 0x48, 0x44, 0x0f, 0xe1,
	0x14, 0x3d, 0x98, 0x3b, 0xa4, 0x84, 0x50, 0x45, 0x41, 0xb6, 0x27, 0xb1, 0x81, 0x9d, 0x6e, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xb7, 0x45, 0x56, 0xc9, 0x00, 0x00, 0x00,
}