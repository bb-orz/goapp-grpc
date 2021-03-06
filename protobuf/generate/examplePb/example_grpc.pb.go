// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example_grpc.proto

package examplePb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type EmptyReqMsg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyReqMsg) Reset()         { *m = EmptyReqMsg{} }
func (m *EmptyReqMsg) String() string { return proto.CompactTextString(m) }
func (*EmptyReqMsg) ProtoMessage()    {}
func (*EmptyReqMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_804d5c0a5eff92d4, []int{0}
}

func (m *EmptyReqMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyReqMsg.Unmarshal(m, b)
}
func (m *EmptyReqMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyReqMsg.Marshal(b, m, deterministic)
}
func (m *EmptyReqMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyReqMsg.Merge(m, src)
}
func (m *EmptyReqMsg) XXX_Size() int {
	return xxx_messageInfo_EmptyReqMsg.Size(m)
}
func (m *EmptyReqMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyReqMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyReqMsg proto.InternalMessageInfo

type CommonRespMsg struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonRespMsg) Reset()         { *m = CommonRespMsg{} }
func (m *CommonRespMsg) String() string { return proto.CompactTextString(m) }
func (*CommonRespMsg) ProtoMessage()    {}
func (*CommonRespMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_804d5c0a5eff92d4, []int{1}
}

func (m *CommonRespMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonRespMsg.Unmarshal(m, b)
}
func (m *CommonRespMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonRespMsg.Marshal(b, m, deterministic)
}
func (m *CommonRespMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonRespMsg.Merge(m, src)
}
func (m *CommonRespMsg) XXX_Size() int {
	return xxx_messageInfo_CommonRespMsg.Size(m)
}
func (m *CommonRespMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonRespMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CommonRespMsg proto.InternalMessageInfo

func (m *CommonRespMsg) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type FooReqMsg struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FooReqMsg) Reset()         { *m = FooReqMsg{} }
func (m *FooReqMsg) String() string { return proto.CompactTextString(m) }
func (*FooReqMsg) ProtoMessage()    {}
func (*FooReqMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_804d5c0a5eff92d4, []int{2}
}

func (m *FooReqMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FooReqMsg.Unmarshal(m, b)
}
func (m *FooReqMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FooReqMsg.Marshal(b, m, deterministic)
}
func (m *FooReqMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FooReqMsg.Merge(m, src)
}
func (m *FooReqMsg) XXX_Size() int {
	return xxx_messageInfo_FooReqMsg.Size(m)
}
func (m *FooReqMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_FooReqMsg.DiscardUnknown(m)
}

var xxx_messageInfo_FooReqMsg proto.InternalMessageInfo

func (m *FooReqMsg) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FooRespMsg struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FooRespMsg) Reset()         { *m = FooRespMsg{} }
func (m *FooRespMsg) String() string { return proto.CompactTextString(m) }
func (*FooRespMsg) ProtoMessage()    {}
func (*FooRespMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_804d5c0a5eff92d4, []int{3}
}

func (m *FooRespMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FooRespMsg.Unmarshal(m, b)
}
func (m *FooRespMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FooRespMsg.Marshal(b, m, deterministic)
}
func (m *FooRespMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FooRespMsg.Merge(m, src)
}
func (m *FooRespMsg) XXX_Size() int {
	return xxx_messageInfo_FooRespMsg.Size(m)
}
func (m *FooRespMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_FooRespMsg.DiscardUnknown(m)
}

var xxx_messageInfo_FooRespMsg proto.InternalMessageInfo

func (m *FooRespMsg) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type BarReqMsg struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarReqMsg) Reset()         { *m = BarReqMsg{} }
func (m *BarReqMsg) String() string { return proto.CompactTextString(m) }
func (*BarReqMsg) ProtoMessage()    {}
func (*BarReqMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_804d5c0a5eff92d4, []int{4}
}

func (m *BarReqMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarReqMsg.Unmarshal(m, b)
}
func (m *BarReqMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarReqMsg.Marshal(b, m, deterministic)
}
func (m *BarReqMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarReqMsg.Merge(m, src)
}
func (m *BarReqMsg) XXX_Size() int {
	return xxx_messageInfo_BarReqMsg.Size(m)
}
func (m *BarReqMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_BarReqMsg.DiscardUnknown(m)
}

var xxx_messageInfo_BarReqMsg proto.InternalMessageInfo

func (m *BarReqMsg) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type BarRespMsg struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarRespMsg) Reset()         { *m = BarRespMsg{} }
func (m *BarRespMsg) String() string { return proto.CompactTextString(m) }
func (*BarRespMsg) ProtoMessage()    {}
func (*BarRespMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_804d5c0a5eff92d4, []int{5}
}

func (m *BarRespMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarRespMsg.Unmarshal(m, b)
}
func (m *BarRespMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarRespMsg.Marshal(b, m, deterministic)
}
func (m *BarRespMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarRespMsg.Merge(m, src)
}
func (m *BarRespMsg) XXX_Size() int {
	return xxx_messageInfo_BarRespMsg.Size(m)
}
func (m *BarRespMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_BarRespMsg.DiscardUnknown(m)
}

var xxx_messageInfo_BarRespMsg proto.InternalMessageInfo

func (m *BarRespMsg) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*EmptyReqMsg)(nil), "EmptyReqMsg")
	proto.RegisterType((*CommonRespMsg)(nil), "CommonRespMsg")
	proto.RegisterType((*FooReqMsg)(nil), "FooReqMsg")
	proto.RegisterType((*FooRespMsg)(nil), "FooRespMsg")
	proto.RegisterType((*BarReqMsg)(nil), "BarReqMsg")
	proto.RegisterType((*BarRespMsg)(nil), "BarRespMsg")
}

func init() { proto.RegisterFile("example_grpc.proto", fileDescriptor_804d5c0a5eff92d4) }

var fileDescriptor_804d5c0a5eff92d4 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x69, 0x57, 0x8a, 0x9d, 0xb5, 0x3d, 0x0c, 0x22, 0xb2, 0x42, 0x91, 0xa5, 0xa0, 0xa7,
	0x08, 0xfa, 0x0f, 0x56, 0x5a, 0x4f, 0x85, 0xb2, 0xde, 0xbc, 0x48, 0xba, 0x1d, 0x42, 0xa0, 0xd9,
	0x89, 0x93, 0x28, 0xea, 0xaf, 0x17, 0xd2, 0xb2, 0xae, 0xa0, 0x1e, 0xc3, 0xcb, 0x3c, 0xbe, 0xef,
	0x01, 0xd2, 0xbb, 0x76, 0x7e, 0x47, 0xcf, 0x46, 0x7c, 0xa3, 0xbc, 0x70, 0xe4, 0x72, 0x02, 0xf9,
	0xc2, 0xf9, 0xf8, 0x51, 0xd3, 0xcb, 0x2a, 0x98, 0xf2, 0x0a, 0x26, 0xf7, 0xec, 0x1c, 0xb7, 0x35,
	0x05, 0xbf, 0x0a, 0x06, 0xcf, 0x60, 0x24, 0x14, 0x5e, 0x77, 0xf1, 0x7c, 0x70, 0x39, 0xb8, 0x3e,
	0xae, 0x0f, 0xaf, 0xf2, 0x02, 0xc6, 0x4b, 0xe6, 0xfd, 0x15, 0x4e, 0x61, 0x68, 0xb7, 0xe9, 0x43,
	0x56, 0x0f, 0xed, 0xb6, 0x9c, 0x03, 0xa4, 0xf0, 0xb7, 0x8a, 0x71, 0xbf, 0xa2, 0xd2, 0xf2, 0x77,
	0x45, 0x0a, 0xff, 0xa5, 0xb8, 0xfd, 0x04, 0x5c, 0xec, 0x9d, 0x1e, 0xc4, 0x37, 0x8f, 0x24, 0x6f,
	0xb6, 0x21, 0x9c, 0xc3, 0xd1, 0xda, 0xb6, 0x06, 0x4f, 0x54, 0x4f, 0xad, 0x98, 0xaa, 0x9f, 0x66,
	0x33, 0xc8, 0x96, 0xcc, 0x08, 0xaa, 0xf3, 0x28, 0x72, 0xd5, 0xc3, 0x9e, 0x41, 0x56, 0x69, 0x41,
	0x50, 0x1d, 0x64, 0x91, 0xab, 0x6f, 0xa6, 0xea, 0xf4, 0x09, 0x0d, 0xb5, 0x24, 0x3a, 0xd2, 0xcd,
	0x61, 0xd8, 0xf5, 0x66, 0x33, 0x4a, 0xb3, 0xde, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xed,
	0x03, 0x03, 0x6c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleGrpcServiceClient is the client API for ExampleGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleGrpcServiceClient interface {
	Ping(ctx context.Context, in *EmptyReqMsg, opts ...grpc.CallOption) (*CommonRespMsg, error)
	Foo(ctx context.Context, in *FooReqMsg, opts ...grpc.CallOption) (*FooRespMsg, error)
	Bar(ctx context.Context, in *BarReqMsg, opts ...grpc.CallOption) (*BarRespMsg, error)
}

type exampleGrpcServiceClient struct {
	cc *grpc.ClientConn
}

func NewExampleGrpcServiceClient(cc *grpc.ClientConn) ExampleGrpcServiceClient {
	return &exampleGrpcServiceClient{cc}
}

func (c *exampleGrpcServiceClient) Ping(ctx context.Context, in *EmptyReqMsg, opts ...grpc.CallOption) (*CommonRespMsg, error) {
	out := new(CommonRespMsg)
	err := c.cc.Invoke(ctx, "/ExampleGrpcService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleGrpcServiceClient) Foo(ctx context.Context, in *FooReqMsg, opts ...grpc.CallOption) (*FooRespMsg, error) {
	out := new(FooRespMsg)
	err := c.cc.Invoke(ctx, "/ExampleGrpcService/Foo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleGrpcServiceClient) Bar(ctx context.Context, in *BarReqMsg, opts ...grpc.CallOption) (*BarRespMsg, error) {
	out := new(BarRespMsg)
	err := c.cc.Invoke(ctx, "/ExampleGrpcService/Bar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleGrpcServiceServer is the server API for ExampleGrpcService service.
type ExampleGrpcServiceServer interface {
	Ping(context.Context, *EmptyReqMsg) (*CommonRespMsg, error)
	Foo(context.Context, *FooReqMsg) (*FooRespMsg, error)
	Bar(context.Context, *BarReqMsg) (*BarRespMsg, error)
}

// UnimplementedExampleGrpcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExampleGrpcServiceServer struct {
}

func (*UnimplementedExampleGrpcServiceServer) Ping(ctx context.Context, req *EmptyReqMsg) (*CommonRespMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedExampleGrpcServiceServer) Foo(ctx context.Context, req *FooReqMsg) (*FooRespMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Foo not implemented")
}
func (*UnimplementedExampleGrpcServiceServer) Bar(ctx context.Context, req *BarReqMsg) (*BarRespMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bar not implemented")
}

func RegisterExampleGrpcServiceServer(s *grpc.Server, srv ExampleGrpcServiceServer) {
	s.RegisterService(&_ExampleGrpcService_serviceDesc, srv)
}

func _ExampleGrpcService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleGrpcServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExampleGrpcService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleGrpcServiceServer).Ping(ctx, req.(*EmptyReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleGrpcService_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleGrpcServiceServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExampleGrpcService/Foo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleGrpcServiceServer).Foo(ctx, req.(*FooReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleGrpcService_Bar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BarReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleGrpcServiceServer).Bar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExampleGrpcService/Bar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleGrpcServiceServer).Bar(ctx, req.(*BarReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleGrpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ExampleGrpcService",
	HandlerType: (*ExampleGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ExampleGrpcService_Ping_Handler,
		},
		{
			MethodName: "Foo",
			Handler:    _ExampleGrpcService_Foo_Handler,
		},
		{
			MethodName: "Bar",
			Handler:    _ExampleGrpcService_Bar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example_grpc.proto",
}
