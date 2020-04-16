// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package proto

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

type SubscribeReq struct {
	Topics               []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeReq) Reset()         { *m = SubscribeReq{} }
func (m *SubscribeReq) String() string { return proto.CompactTextString(m) }
func (*SubscribeReq) ProtoMessage()    {}
func (*SubscribeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *SubscribeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeReq.Unmarshal(m, b)
}
func (m *SubscribeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeReq.Marshal(b, m, deterministic)
}
func (m *SubscribeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeReq.Merge(m, src)
}
func (m *SubscribeReq) XXX_Size() int {
	return xxx_messageInfo_SubscribeReq.Size(m)
}
func (m *SubscribeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeReq.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeReq proto.InternalMessageInfo

func (m *SubscribeReq) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type SubscribeReply struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeReply) Reset()         { *m = SubscribeReply{} }
func (m *SubscribeReply) String() string { return proto.CompactTextString(m) }
func (*SubscribeReply) ProtoMessage()    {}
func (*SubscribeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *SubscribeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeReply.Unmarshal(m, b)
}
func (m *SubscribeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeReply.Marshal(b, m, deterministic)
}
func (m *SubscribeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeReply.Merge(m, src)
}
func (m *SubscribeReply) XXX_Size() int {
	return xxx_messageInfo_SubscribeReply.Size(m)
}
func (m *SubscribeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeReply proto.InternalMessageInfo

func (m *SubscribeReply) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SubscribeReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SubscribeReq)(nil), "grpc.push.SubscribeReq")
	proto.RegisterType((*SubscribeReply)(nil), "grpc.push.SubscribeReply")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0xb0, 0x4b, 0x8b, 0x33, 0x94, 0xd4, 0xb8,
	0x78, 0x82, 0x4b, 0x93, 0x8a, 0x93, 0x8b, 0x32, 0x93, 0x52, 0x83, 0x52, 0x0b, 0x85, 0xc4, 0xb8,
	0xd8, 0x4a, 0xf2, 0x0b, 0x32, 0x93, 0x8b, 0x25, 0x18, 0x15, 0x98, 0x35, 0x38, 0x83, 0xa0, 0x3c,
	0x25, 0x07, 0x2e, 0x3e, 0x24, 0x75, 0x05, 0x39, 0x95, 0x42, 0x22, 0x5c, 0xac, 0x60, 0x39, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38,
	0x31, 0x3d, 0x55, 0x82, 0x09, 0x2c, 0x0e, 0xe3, 0x1a, 0x05, 0x71, 0x71, 0x07, 0x94, 0x16, 0x67,
	0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x39, 0x73, 0x71, 0xc2, 0x0d, 0x14, 0x12, 0xd7,
	0x83, 0xbb, 0x48, 0x0f, 0xd9, 0x39, 0x52, 0x92, 0xd8, 0x25, 0x0a, 0x72, 0x2a, 0x95, 0x18, 0x0c,
	0x18, 0x9d, 0x38, 0xa2, 0xd8, 0xf4, 0xc1, 0x5e, 0x4a, 0x62, 0x03, 0x53, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xfc, 0x68, 0x82, 0x5e, 0xe7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PushServiceClient is the client API for PushService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PushServiceClient interface {
	Subscribe(ctx context.Context, in *SubscribeReq, opts ...grpc.CallOption) (PushService_SubscribeClient, error)
}

type pushServiceClient struct {
	cc *grpc.ClientConn
}

func NewPushServiceClient(cc *grpc.ClientConn) PushServiceClient {
	return &pushServiceClient{cc}
}

func (c *pushServiceClient) Subscribe(ctx context.Context, in *SubscribeReq, opts ...grpc.CallOption) (PushService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PushService_serviceDesc.Streams[0], "/grpc.push.PushService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PushService_SubscribeClient interface {
	Recv() (*SubscribeReply, error)
	grpc.ClientStream
}

type pushServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pushServiceSubscribeClient) Recv() (*SubscribeReply, error) {
	m := new(SubscribeReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PushServiceServer is the server API for PushService service.
type PushServiceServer interface {
	Subscribe(*SubscribeReq, PushService_SubscribeServer) error
}

// UnimplementedPushServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPushServiceServer struct {
}

func (*UnimplementedPushServiceServer) Subscribe(req *SubscribeReq, srv PushService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterPushServiceServer(s *grpc.Server, srv PushServiceServer) {
	s.RegisterService(&_PushService_serviceDesc, srv)
}

func _PushService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushServiceServer).Subscribe(m, &pushServiceSubscribeServer{stream})
}

type PushService_SubscribeServer interface {
	Send(*SubscribeReply) error
	grpc.ServerStream
}

type pushServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pushServiceSubscribeServer) Send(m *SubscribeReply) error {
	return x.ServerStream.SendMsg(m)
}

var _PushService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.push.PushService",
	HandlerType: (*PushServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PushService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc.proto",
}
