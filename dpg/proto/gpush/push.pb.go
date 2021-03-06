// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push.proto

package gpush

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

type RetData struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Summary              string   `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetData) Reset()         { *m = RetData{} }
func (m *RetData) String() string { return proto.CompactTextString(m) }
func (*RetData) ProtoMessage()    {}
func (*RetData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{0}
}

func (m *RetData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetData.Unmarshal(m, b)
}
func (m *RetData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetData.Marshal(b, m, deterministic)
}
func (m *RetData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetData.Merge(m, src)
}
func (m *RetData) XXX_Size() int {
	return xxx_messageInfo_RetData.Size(m)
}
func (m *RetData) XXX_DiscardUnknown() {
	xxx_messageInfo_RetData.DiscardUnknown(m)
}

var xxx_messageInfo_RetData proto.InternalMessageInfo

func (m *RetData) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RetData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *RetData) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

type PushKeysReq struct {
	Op                   int64    `protobuf:"varint,1,opt,name=op,proto3" json:"op,omitempty"`
	Keys                 []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushKeysReq) Reset()         { *m = PushKeysReq{} }
func (m *PushKeysReq) String() string { return proto.CompactTextString(m) }
func (*PushKeysReq) ProtoMessage()    {}
func (*PushKeysReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{1}
}

func (m *PushKeysReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushKeysReq.Unmarshal(m, b)
}
func (m *PushKeysReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushKeysReq.Marshal(b, m, deterministic)
}
func (m *PushKeysReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushKeysReq.Merge(m, src)
}
func (m *PushKeysReq) XXX_Size() int {
	return xxx_messageInfo_PushKeysReq.Size(m)
}
func (m *PushKeysReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PushKeysReq.DiscardUnknown(m)
}

var xxx_messageInfo_PushKeysReq proto.InternalMessageInfo

func (m *PushKeysReq) GetOp() int64 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *PushKeysReq) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *PushKeysReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type PushMidsReq struct {
	Op                   int64    `protobuf:"varint,1,opt,name=op,proto3" json:"op,omitempty"`
	Mids                 []int64  `protobuf:"varint,2,rep,packed,name=mids,proto3" json:"mids,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushMidsReq) Reset()         { *m = PushMidsReq{} }
func (m *PushMidsReq) String() string { return proto.CompactTextString(m) }
func (*PushMidsReq) ProtoMessage()    {}
func (*PushMidsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{2}
}

func (m *PushMidsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushMidsReq.Unmarshal(m, b)
}
func (m *PushMidsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushMidsReq.Marshal(b, m, deterministic)
}
func (m *PushMidsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMidsReq.Merge(m, src)
}
func (m *PushMidsReq) XXX_Size() int {
	return xxx_messageInfo_PushMidsReq.Size(m)
}
func (m *PushMidsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMidsReq.DiscardUnknown(m)
}

var xxx_messageInfo_PushMidsReq proto.InternalMessageInfo

func (m *PushMidsReq) GetOp() int64 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *PushMidsReq) GetMids() []int64 {
	if m != nil {
		return m.Mids
	}
	return nil
}

func (m *PushMidsReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type PushRoomReq struct {
	Op                   int64    `protobuf:"varint,1,opt,name=op,proto3" json:"op,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Room                 string   `protobuf:"bytes,3,opt,name=room,proto3" json:"room,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushRoomReq) Reset()         { *m = PushRoomReq{} }
func (m *PushRoomReq) String() string { return proto.CompactTextString(m) }
func (*PushRoomReq) ProtoMessage()    {}
func (*PushRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{3}
}

func (m *PushRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushRoomReq.Unmarshal(m, b)
}
func (m *PushRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushRoomReq.Marshal(b, m, deterministic)
}
func (m *PushRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRoomReq.Merge(m, src)
}
func (m *PushRoomReq) XXX_Size() int {
	return xxx_messageInfo_PushRoomReq.Size(m)
}
func (m *PushRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_PushRoomReq proto.InternalMessageInfo

func (m *PushRoomReq) GetOp() int64 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *PushRoomReq) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PushRoomReq) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

func (m *PushRoomReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type PushAllReq struct {
	Op                   int64    `protobuf:"varint,1,opt,name=op,proto3" json:"op,omitempty"`
	Speed                int64    `protobuf:"varint,2,opt,name=speed,proto3" json:"speed,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushAllReq) Reset()         { *m = PushAllReq{} }
func (m *PushAllReq) String() string { return proto.CompactTextString(m) }
func (*PushAllReq) ProtoMessage()    {}
func (*PushAllReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{4}
}

func (m *PushAllReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushAllReq.Unmarshal(m, b)
}
func (m *PushAllReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushAllReq.Marshal(b, m, deterministic)
}
func (m *PushAllReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushAllReq.Merge(m, src)
}
func (m *PushAllReq) XXX_Size() int {
	return xxx_messageInfo_PushAllReq.Size(m)
}
func (m *PushAllReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PushAllReq.DiscardUnknown(m)
}

var xxx_messageInfo_PushAllReq proto.InternalMessageInfo

func (m *PushAllReq) GetOp() int64 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *PushAllReq) GetSpeed() int64 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *PushAllReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*RetData)(nil), "dpg.push.RetData")
	proto.RegisterType((*PushKeysReq)(nil), "dpg.push.PushKeysReq")
	proto.RegisterType((*PushMidsReq)(nil), "dpg.push.PushMidsReq")
	proto.RegisterType((*PushRoomReq)(nil), "dpg.push.PushRoomReq")
	proto.RegisterType((*PushAllReq)(nil), "dpg.push.PushAllReq")
}

func init() { proto.RegisterFile("push.proto", fileDescriptor_d1e4bfd2e9d102bb) }

var fileDescriptor_d1e4bfd2e9d102bb = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0xb6, 0x35, 0x75, 0x04, 0xc1, 0xa5, 0x42, 0xf0, 0x14, 0x7a, 0xea, 0x29, 0x87,
	0xea, 0x17, 0x50, 0xaa, 0x97, 0x22, 0xc8, 0xde, 0xf4, 0xb6, 0xba, 0x4b, 0x5b, 0x4c, 0xd8, 0x35,
	0xbb, 0x39, 0xe4, 0xdb, 0xfa, 0x51, 0x64, 0xf6, 0x4f, 0x5b, 0xcb, 0xb6, 0xb7, 0x37, 0x0f, 0xde,
	0x6f, 0x33, 0x6f, 0x02, 0xa0, 0x7b, 0xb3, 0xa9, 0x75, 0xa7, 0xac, 0xa2, 0x13, 0xa1, 0xd7, 0x35,
	0xce, 0xb3, 0x15, 0x14, 0x4c, 0xda, 0x25, 0xb7, 0x9c, 0x52, 0x18, 0x7d, 0x29, 0x21, 0xcb, 0xac,
	0xca, 0xe6, 0x84, 0x39, 0x8d, 0x9e, 0xe0, 0x96, 0x97, 0x79, 0x95, 0xcd, 0x2f, 0x99, 0xd3, 0xb4,
	0x84, 0xc2, 0xf4, 0x6d, 0xcb, 0xbb, 0xa1, 0x24, 0xce, 0x8e, 0xe3, 0xec, 0x19, 0xae, 0xde, 0x7a,
	0xb3, 0x59, 0xc9, 0xc1, 0x30, 0xf9, 0x43, 0xaf, 0x21, 0x57, 0x3a, 0xe0, 0x72, 0xa5, 0x11, 0xf6,
	0x2d, 0x07, 0x53, 0xe6, 0x15, 0x41, 0x18, 0xea, 0xdd, 0x03, 0x64, 0xff, 0x40, 0xc4, 0xbc, 0x6e,
	0xc5, 0x29, 0x4c, 0xbb, 0x15, 0x1e, 0x43, 0x98, 0xd3, 0x49, 0xcc, 0xbb, 0xc7, 0x30, 0xa5, 0xda,
	0x13, 0x18, 0x3b, 0x68, 0x19, 0x57, 0x43, 0x8d, 0x5e, 0xa7, 0x54, 0x1b, 0x31, 0xa8, 0x77, 0xe8,
	0xd1, 0x01, 0xfa, 0x05, 0x00, 0xd1, 0x8f, 0x4d, 0x93, 0x22, 0x4f, 0x61, 0x6c, 0xb4, 0x94, 0xc2,
	0xa1, 0x09, 0xf3, 0x43, 0xea, 0x13, 0x17, 0xbf, 0x19, 0x4c, 0x10, 0xb4, 0x94, 0xbc, 0xa1, 0x0f,
	0x5e, 0x63, 0x7b, 0xf4, 0xb6, 0x8e, 0x17, 0xaa, 0x0f, 0x1a, 0xbd, 0xbb, 0xd9, 0xdb, 0xf1, 0x6a,
	0x21, 0x85, 0x65, 0x1d, 0xa7, 0x42, 0x81, 0x67, 0x52, 0xd8, 0xcd, 0x71, 0x2a, 0xf4, 0x95, 0x4a,
	0x2d, 0xa0, 0x08, 0x6b, 0xd3, 0xe9, 0xff, 0x90, 0x6f, 0x22, 0x91, 0x79, 0x2a, 0x3e, 0xc6, 0x6b,
	0x34, 0x3e, 0x2f, 0xdc, 0xaf, 0x77, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x61, 0xaa, 0x9c,
	0x88, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PushDealClient is the client API for PushDeal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PushDealClient interface {
	PushKeys(ctx context.Context, in *PushKeysReq, opts ...grpc.CallOption) (*RetData, error)
	PushMids(ctx context.Context, in *PushMidsReq, opts ...grpc.CallOption) (*RetData, error)
	PushRoom(ctx context.Context, in *PushRoomReq, opts ...grpc.CallOption) (*RetData, error)
	PushAll(ctx context.Context, in *PushAllReq, opts ...grpc.CallOption) (*RetData, error)
}

type pushDealClient struct {
	cc *grpc.ClientConn
}

func NewPushDealClient(cc *grpc.ClientConn) PushDealClient {
	return &pushDealClient{cc}
}

func (c *pushDealClient) PushKeys(ctx context.Context, in *PushKeysReq, opts ...grpc.CallOption) (*RetData, error) {
	out := new(RetData)
	err := c.cc.Invoke(ctx, "/dpg.push.PushDeal/PushKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushDealClient) PushMids(ctx context.Context, in *PushMidsReq, opts ...grpc.CallOption) (*RetData, error) {
	out := new(RetData)
	err := c.cc.Invoke(ctx, "/dpg.push.PushDeal/PushMids", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushDealClient) PushRoom(ctx context.Context, in *PushRoomReq, opts ...grpc.CallOption) (*RetData, error) {
	out := new(RetData)
	err := c.cc.Invoke(ctx, "/dpg.push.PushDeal/PushRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushDealClient) PushAll(ctx context.Context, in *PushAllReq, opts ...grpc.CallOption) (*RetData, error) {
	out := new(RetData)
	err := c.cc.Invoke(ctx, "/dpg.push.PushDeal/PushAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushDealServer is the server API for PushDeal service.
type PushDealServer interface {
	PushKeys(context.Context, *PushKeysReq) (*RetData, error)
	PushMids(context.Context, *PushMidsReq) (*RetData, error)
	PushRoom(context.Context, *PushRoomReq) (*RetData, error)
	PushAll(context.Context, *PushAllReq) (*RetData, error)
}

// UnimplementedPushDealServer can be embedded to have forward compatible implementations.
type UnimplementedPushDealServer struct {
}

func (*UnimplementedPushDealServer) PushKeys(ctx context.Context, req *PushKeysReq) (*RetData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushKeys not implemented")
}
func (*UnimplementedPushDealServer) PushMids(ctx context.Context, req *PushMidsReq) (*RetData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMids not implemented")
}
func (*UnimplementedPushDealServer) PushRoom(ctx context.Context, req *PushRoomReq) (*RetData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushRoom not implemented")
}
func (*UnimplementedPushDealServer) PushAll(ctx context.Context, req *PushAllReq) (*RetData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushAll not implemented")
}

func RegisterPushDealServer(s *grpc.Server, srv PushDealServer) {
	s.RegisterService(&_PushDeal_serviceDesc, srv)
}

func _PushDeal_PushKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushKeysReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushDealServer).PushKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dpg.push.PushDeal/PushKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushDealServer).PushKeys(ctx, req.(*PushKeysReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushDeal_PushMids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMidsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushDealServer).PushMids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dpg.push.PushDeal/PushMids",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushDealServer).PushMids(ctx, req.(*PushMidsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushDeal_PushRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushDealServer).PushRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dpg.push.PushDeal/PushRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushDealServer).PushRoom(ctx, req.(*PushRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushDeal_PushAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushDealServer).PushAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dpg.push.PushDeal/PushAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushDealServer).PushAll(ctx, req.(*PushAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PushDeal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dpg.push.PushDeal",
	HandlerType: (*PushDealServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushKeys",
			Handler:    _PushDeal_PushKeys_Handler,
		},
		{
			MethodName: "PushMids",
			Handler:    _PushDeal_PushMids_Handler,
		},
		{
			MethodName: "PushRoom",
			Handler:    _PushDeal_PushRoom_Handler,
		},
		{
			MethodName: "PushAll",
			Handler:    _PushDeal_PushAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "push.proto",
}
