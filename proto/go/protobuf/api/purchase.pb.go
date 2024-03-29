// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/api/purchase.proto

package api

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

type PurchaseRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurchaseRequest) Reset()         { *m = PurchaseRequest{} }
func (m *PurchaseRequest) String() string { return proto.CompactTextString(m) }
func (*PurchaseRequest) ProtoMessage()    {}
func (*PurchaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9155614e99cf9a02, []int{0}
}

func (m *PurchaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseRequest.Unmarshal(m, b)
}
func (m *PurchaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseRequest.Marshal(b, m, deterministic)
}
func (m *PurchaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseRequest.Merge(m, src)
}
func (m *PurchaseRequest) XXX_Size() int {
	return xxx_messageInfo_PurchaseRequest.Size(m)
}
func (m *PurchaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseRequest proto.InternalMessageInfo

func (m *PurchaseRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PurchaseResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurchaseResponse) Reset()         { *m = PurchaseResponse{} }
func (m *PurchaseResponse) String() string { return proto.CompactTextString(m) }
func (*PurchaseResponse) ProtoMessage()    {}
func (*PurchaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9155614e99cf9a02, []int{1}
}

func (m *PurchaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseResponse.Unmarshal(m, b)
}
func (m *PurchaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseResponse.Marshal(b, m, deterministic)
}
func (m *PurchaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseResponse.Merge(m, src)
}
func (m *PurchaseResponse) XXX_Size() int {
	return xxx_messageInfo_PurchaseResponse.Size(m)
}
func (m *PurchaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseResponse proto.InternalMessageInfo

func (m *PurchaseResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PurchaseResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*PurchaseRequest)(nil), "api.PurchaseRequest")
	proto.RegisterType((*PurchaseResponse)(nil), "api.PurchaseResponse")
}

func init() {
	proto.RegisterFile("protobuf/api/purchase.proto", fileDescriptor_9155614e99cf9a02)
}

var fileDescriptor_9155614e99cf9a02 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x28, 0x2d, 0x4a, 0xce, 0x48, 0x2c,
	0x4e, 0xd5, 0x03, 0x8b, 0x0a, 0x31, 0x27, 0x16, 0x64, 0x2a, 0x29, 0x72, 0xf1, 0x07, 0x40, 0x85,
	0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x94, 0x9c, 0xb8, 0x04, 0x10, 0x4a, 0x8a, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0x8b, 0x4b, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0xc1, 0x0a,
	0x39, 0x82, 0x60, 0x5c, 0x21, 0x11, 0x2e, 0xd6, 0xd4, 0xa2, 0xa2, 0xfc, 0x22, 0x09, 0x26, 0xb0,
	0x01, 0x10, 0x8e, 0x91, 0x0f, 0xc2, 0x9a, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x4b,
	0x2e, 0x0e, 0x98, 0x90, 0x90, 0x88, 0x5e, 0x62, 0x41, 0xa6, 0x1e, 0x9a, 0x43, 0xa4, 0x44, 0xd1,
	0x44, 0x21, 0x76, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0x3d, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xe8, 0x6f, 0xa7, 0x73, 0xdf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PurchaseServiceClient is the client API for PurchaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PurchaseServiceClient interface {
	Purchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error)
}

type purchaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPurchaseServiceClient(cc grpc.ClientConnInterface) PurchaseServiceClient {
	return &purchaseServiceClient{cc}
}

func (c *purchaseServiceClient) Purchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error) {
	out := new(PurchaseResponse)
	err := c.cc.Invoke(ctx, "/api.PurchaseService/Purchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PurchaseServiceServer is the server API for PurchaseService service.
type PurchaseServiceServer interface {
	Purchase(context.Context, *PurchaseRequest) (*PurchaseResponse, error)
}

// UnimplementedPurchaseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPurchaseServiceServer struct {
}

func (*UnimplementedPurchaseServiceServer) Purchase(ctx context.Context, req *PurchaseRequest) (*PurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purchase not implemented")
}

func RegisterPurchaseServiceServer(s *grpc.Server, srv PurchaseServiceServer) {
	s.RegisterService(&_PurchaseService_serviceDesc, srv)
}

func _PurchaseService_Purchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchaseServiceServer).Purchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PurchaseService/Purchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchaseServiceServer).Purchase(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PurchaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.PurchaseService",
	HandlerType: (*PurchaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Purchase",
			Handler:    _PurchaseService_Purchase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/purchase.proto",
}
