// Code generated by test_proto-gen-go. DO NOT EDIT.
// source: hello.test_proto

package pb

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
// is compatible with the test_proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// test_proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the test_proto package

type HelloRequest_MSGID int32

const (
	HelloRequest_ID HelloRequest_MSGID = 664
)

var HelloRequest_MSGID_name = map[int32]string{
	664: "ID",
}

var HelloRequest_MSGID_value = map[string]int32{
	"ID": 664,
}

func (x HelloRequest_MSGID) Enum() *HelloRequest_MSGID {
	p := new(HelloRequest_MSGID)
	*p = x
	return p
}

func (x HelloRequest_MSGID) String() string {
	return proto.EnumName(HelloRequest_MSGID_name, int32(x))
}

func (x *HelloRequest_MSGID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HelloRequest_MSGID_value, data, "HelloRequest_MSGID")
	if err != nil {
		return err
	}
	*x = HelloRequest_MSGID(value)
	return nil
}

func (HelloRequest_MSGID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0, 0}
}

type HelloResponse_MSGID int32

const (
	HelloResponse_ID HelloResponse_MSGID = 665
)

var HelloResponse_MSGID_name = map[int32]string{
	665: "ID",
}

var HelloResponse_MSGID_value = map[string]int32{
	"ID": 665,
}

func (x HelloResponse_MSGID) Enum() *HelloResponse_MSGID {
	p := new(HelloResponse_MSGID)
	*p = x
	return p
}

func (x HelloResponse_MSGID) String() string {
	return proto.EnumName(HelloResponse_MSGID_name, int32(x))
}

func (x *HelloResponse_MSGID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HelloResponse_MSGID_value, data, "HelloResponse_MSGID")
	if err != nil {
		return err
	}
	*x = HelloResponse_MSGID(value)
	return nil
}

func (HelloResponse_MSGID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1, 0}
}

type HelloRequest struct {
	Username             *string  `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

type HelloResponse struct {
	Message              *string  `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type ClientStream struct {
	Stream               []byte   `protobuf:"bytes,1,opt,name=stream" json:"stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStream) Reset()         { *m = ClientStream{} }
func (m *ClientStream) String() string { return proto.CompactTextString(m) }
func (*ClientStream) ProtoMessage()    {}
func (*ClientStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{2}
}

func (m *ClientStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStream.Unmarshal(m, b)
}
func (m *ClientStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStream.Marshal(b, m, deterministic)
}
func (m *ClientStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStream.Merge(m, src)
}
func (m *ClientStream) XXX_Size() int {
	return xxx_messageInfo_ClientStream.Size(m)
}
func (m *ClientStream) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStream.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStream proto.InternalMessageInfo

func (m *ClientStream) GetStream() []byte {
	if m != nil {
		return m.Stream
	}
	return nil
}

type ServerStream struct {
	Stream               []byte   `protobuf:"bytes,1,opt,name=stream" json:"stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStream) Reset()         { *m = ServerStream{} }
func (m *ServerStream) String() string { return proto.CompactTextString(m) }
func (*ServerStream) ProtoMessage()    {}
func (*ServerStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{3}
}

func (m *ServerStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStream.Unmarshal(m, b)
}
func (m *ServerStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStream.Marshal(b, m, deterministic)
}
func (m *ServerStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStream.Merge(m, src)
}
func (m *ServerStream) XXX_Size() int {
	return xxx_messageInfo_ServerStream.Size(m)
}
func (m *ServerStream) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStream.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStream proto.InternalMessageInfo

func (m *ServerStream) GetStream() []byte {
	if m != nil {
		return m.Stream
	}
	return nil
}

func init() {
	proto.RegisterEnum("test_proto.HelloRequest_MSGID", HelloRequest_MSGID_name, HelloRequest_MSGID_value)
	proto.RegisterEnum("test_proto.HelloResponse_MSGID", HelloResponse_MSGID_name, HelloResponse_MSGID_value)
	proto.RegisterType((*HelloRequest)(nil), "test_proto.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "test_proto.HelloResponse")
	proto.RegisterType((*ClientStream)(nil), "test_proto.ClientStream")
	proto.RegisterType((*ServerStream)(nil), "test_proto.ServerStream")
}

func init() { proto.RegisterFile("hello.test_proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xb2, 0xe1, 0xe2, 0xf1,
	0x00, 0x09, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x71, 0x71, 0x94, 0x16, 0xa7,
	0x16, 0xe5, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x4a, 0x02,
	0x5c, 0xac, 0xbe, 0xc1, 0xee, 0x9e, 0x2e, 0x42, 0xec, 0x5c, 0x4c, 0x9e, 0x2e, 0x02, 0x33, 0x58,
	0x95, 0xac, 0xb9, 0x78, 0xa1, 0xba, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8,
	0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x61, 0xba, 0x61, 0x5c, 0x74, 0xcd, 0x33, 0x59, 0x95, 0xd4,
	0xb8, 0x78, 0x9c, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0x82, 0x4b, 0x8a, 0x52, 0x13, 0x73, 0x85, 0xc4,
	0xb8, 0xd8, 0x8a, 0xc1, 0x2c, 0xb0, 0x56, 0x9e, 0x20, 0x28, 0x0f, 0xa4, 0x2e, 0x38, 0xb5, 0xa8,
	0x2c, 0xb5, 0x08, 0xbf, 0x3a, 0xa3, 0x62, 0xa8, 0x57, 0x40, 0x8a, 0x33, 0x93, 0x53, 0x85, 0x0c,
	0xb9, 0x38, 0x82, 0x13, 0x2b, 0xc1, 0x42, 0x42, 0x02, 0x7a, 0x05, 0x49, 0x7a, 0xc8, 0x1e, 0x95,
	0x12, 0x44, 0x12, 0x81, 0x38, 0x5e, 0x89, 0x41, 0xc8, 0x80, 0x8b, 0xc5, 0x39, 0x23, 0xb1, 0x04,
	0xa2, 0x1c, 0xd9, 0x71, 0x52, 0x60, 0x11, 0x64, 0x67, 0x28, 0x31, 0x68, 0x30, 0x1a, 0x30, 0x02,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x7a, 0x15, 0xbd, 0x51, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	Chat(ctx context.Context, opts ...grpc.CallOption) (HelloService_ChatClient, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/test_proto.HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) Chat(ctx context.Context, opts ...grpc.CallOption) (HelloService_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloService_serviceDesc.Streams[0], "/test_proto.HelloService/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceChatClient{stream}
	return x, nil
}

type HelloService_ChatClient interface {
	Send(*ClientStream) error
	Recv() (*ServerStream, error)
	grpc.ClientStream
}

type helloServiceChatClient struct {
	grpc.ClientStream
}

func (x *helloServiceChatClient) Send(m *ClientStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceChatClient) Recv() (*ServerStream, error) {
	m := new(ServerStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	Chat(HelloService_ChatServer) error
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedHelloServiceServer) Chat(srv HelloService_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test_proto.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).Chat(&helloServiceChatServer{stream})
}

type HelloService_ChatServer interface {
	Send(*ServerStream) error
	Recv() (*ClientStream, error)
	grpc.ServerStream
}

type helloServiceChatServer struct {
	grpc.ServerStream
}

func (x *helloServiceChatServer) Send(m *ServerStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceChatServer) Recv() (*ClientStream, error) {
	m := new(ClientStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test_proto.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _HelloService_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.test_proto",
}
