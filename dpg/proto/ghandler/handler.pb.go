// Code generated by protoc-gen-go. DO NOT EDIT.
// source: handler.proto

package ghandler

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

type HandlerReply struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Summary              string   `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandlerReply) Reset()         { *m = HandlerReply{} }
func (m *HandlerReply) String() string { return proto.CompactTextString(m) }
func (*HandlerReply) ProtoMessage()    {}
func (*HandlerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_515968b8e1a22554, []int{0}
}

func (m *HandlerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandlerReply.Unmarshal(m, b)
}
func (m *HandlerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandlerReply.Marshal(b, m, deterministic)
}
func (m *HandlerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandlerReply.Merge(m, src)
}
func (m *HandlerReply) XXX_Size() int {
	return xxx_messageInfo_HandlerReply.Size(m)
}
func (m *HandlerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HandlerReply.DiscardUnknown(m)
}

var xxx_messageInfo_HandlerReply proto.InternalMessageInfo

func (m *HandlerReply) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HandlerReply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *HandlerReply) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

type HandlerReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Group                string   `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Platform             string   `protobuf:"bytes,5,opt,name=platform,proto3" json:"platform,omitempty"`
	Version              string   `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	Timestamp            int64    `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Operation            int64    `protobuf:"varint,8,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandlerReq) Reset()         { *m = HandlerReq{} }
func (m *HandlerReq) String() string { return proto.CompactTextString(m) }
func (*HandlerReq) ProtoMessage()    {}
func (*HandlerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_515968b8e1a22554, []int{1}
}

func (m *HandlerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandlerReq.Unmarshal(m, b)
}
func (m *HandlerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandlerReq.Marshal(b, m, deterministic)
}
func (m *HandlerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandlerReq.Merge(m, src)
}
func (m *HandlerReq) XXX_Size() int {
	return xxx_messageInfo_HandlerReq.Size(m)
}
func (m *HandlerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HandlerReq.DiscardUnknown(m)
}

var xxx_messageInfo_HandlerReq proto.InternalMessageInfo

func (m *HandlerReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *HandlerReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *HandlerReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *HandlerReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *HandlerReq) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *HandlerReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *HandlerReq) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *HandlerReq) GetOperation() int64 {
	if m != nil {
		return m.Operation
	}
	return 0
}

func init() {
	proto.RegisterType((*HandlerReply)(nil), "dpg.handler.HandlerReply")
	proto.RegisterType((*HandlerReq)(nil), "dpg.handler.HandlerReq")
}

func init() { proto.RegisterFile("handler.proto", fileDescriptor_515968b8e1a22554) }

var fileDescriptor_515968b8e1a22554 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3d, 0x4f, 0xf3, 0x40,
	0x10, 0x84, 0xe5, 0xd7, 0x49, 0x6c, 0xef, 0x4b, 0x24, 0xb4, 0x42, 0xe2, 0x88, 0x28, 0xa2, 0x54,
	0xa9, 0x5c, 0x40, 0x4d, 0x43, 0x45, 0x85, 0x90, 0xe9, 0xe8, 0x8e, 0xdc, 0x62, 0x2c, 0x7c, 0x1f,
	0x9c, 0xcf, 0x91, 0xfc, 0x47, 0xf9, 0x3d, 0xe8, 0x3e, 0x70, 0x28, 0xe8, 0x66, 0x9e, 0xf1, 0xce,
	0x6a, 0x7d, 0xb0, 0x7e, 0xe7, 0x4a, 0xf4, 0x64, 0x6b, 0x63, 0xb5, 0xd3, 0xf8, 0x5f, 0x98, 0xb6,
	0x4e, 0x68, 0xf7, 0x04, 0x67, 0x0f, 0x51, 0x36, 0x64, 0xfa, 0x09, 0x11, 0x16, 0x07, 0x2d, 0x88,
	0x65, 0xdb, 0x6c, 0x9f, 0x37, 0x41, 0x7b, 0x26, 0xb8, 0xe3, 0xec, 0xdf, 0x36, 0xdb, 0x57, 0x4d,
	0xd0, 0xc8, 0xa0, 0x18, 0x46, 0x29, 0xb9, 0x9d, 0x58, 0x1e, 0xf0, 0x8f, 0xdd, 0x7d, 0x65, 0x00,
	0x73, 0xe5, 0x27, 0x5e, 0xc0, 0xd2, 0xe9, 0x0f, 0x52, 0xa1, 0xb1, 0x6a, 0xa2, 0xc1, 0x73, 0xc8,
	0xc7, 0x4e, 0xa4, 0x46, 0x2f, 0xfd, 0x77, 0xad, 0xd5, 0xa3, 0x49, 0x75, 0xd1, 0xf8, 0x35, 0x07,
	0xad, 0x1c, 0x29, 0xc7, 0x16, 0x71, 0x4d, 0xb2, 0xb8, 0x81, 0xd2, 0xf4, 0xdc, 0xbd, 0x69, 0x2b,
	0xd9, 0x32, 0x44, 0xb3, 0xf7, 0x53, 0x47, 0xb2, 0x43, 0xa7, 0x15, 0x5b, 0xc5, 0xa9, 0x64, 0xf1,
	0x1a, 0x2a, 0xd7, 0x49, 0x1a, 0x1c, 0x97, 0x86, 0x15, 0xe1, 0xc6, 0x13, 0xf0, 0xa9, 0x36, 0x64,
	0xb9, 0xf3, 0x93, 0x65, 0x4c, 0x67, 0x70, 0xf3, 0x08, 0xeb, 0x74, 0xd7, 0x33, 0xd9, 0x23, 0x59,
	0xbc, 0x83, 0x22, 0x01, 0xbc, 0xac, 0x7f, 0xfd, 0xd4, 0xfa, 0x74, 0xfe, 0xe6, 0xea, 0xef, 0xc0,
	0xf4, 0xd3, 0x3d, 0xbc, 0x94, 0x6d, 0x0a, 0x5e, 0x57, 0xe1, 0x69, 0x6e, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xba, 0x75, 0xed, 0x44, 0xab, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HandlerServerClient is the client API for HandlerServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HandlerServerClient interface {
	Handler(ctx context.Context, in *HandlerReq, opts ...grpc.CallOption) (*HandlerReply, error)
}

type handlerServerClient struct {
	cc *grpc.ClientConn
}

func NewHandlerServerClient(cc *grpc.ClientConn) HandlerServerClient {
	return &handlerServerClient{cc}
}

func (c *handlerServerClient) Handler(ctx context.Context, in *HandlerReq, opts ...grpc.CallOption) (*HandlerReply, error) {
	out := new(HandlerReply)
	err := c.cc.Invoke(ctx, "/dpg.handler.HandlerServer/Handler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandlerServerServer is the server API for HandlerServer service.
type HandlerServerServer interface {
	Handler(context.Context, *HandlerReq) (*HandlerReply, error)
}

// UnimplementedHandlerServerServer can be embedded to have forward compatible implementations.
type UnimplementedHandlerServerServer struct {
}

func (*UnimplementedHandlerServerServer) Handler(ctx context.Context, req *HandlerReq) (*HandlerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handler not implemented")
}

func RegisterHandlerServerServer(s *grpc.Server, srv HandlerServerServer) {
	s.RegisterService(&_HandlerServer_serviceDesc, srv)
}

func _HandlerServer_Handler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandlerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServerServer).Handler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dpg.handler.HandlerServer/Handler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServerServer).Handler(ctx, req.(*HandlerReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HandlerServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dpg.handler.HandlerServer",
	HandlerType: (*HandlerServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handler",
			Handler:    _HandlerServer_Handler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "handler.proto",
}