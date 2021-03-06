// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/cpf/generator.proto

package cpf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type GenerateCPFRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateCPFRequest) Reset()         { *m = GenerateCPFRequest{} }
func (m *GenerateCPFRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateCPFRequest) ProtoMessage()    {}
func (*GenerateCPFRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acd96ee7ee43acb6, []int{0}
}

func (m *GenerateCPFRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateCPFRequest.Unmarshal(m, b)
}
func (m *GenerateCPFRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateCPFRequest.Marshal(b, m, deterministic)
}
func (m *GenerateCPFRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateCPFRequest.Merge(m, src)
}
func (m *GenerateCPFRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateCPFRequest.Size(m)
}
func (m *GenerateCPFRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateCPFRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateCPFRequest proto.InternalMessageInfo

type GenerateCPFResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateCPFResponse) Reset()         { *m = GenerateCPFResponse{} }
func (m *GenerateCPFResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateCPFResponse) ProtoMessage()    {}
func (*GenerateCPFResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_acd96ee7ee43acb6, []int{1}
}

func (m *GenerateCPFResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateCPFResponse.Unmarshal(m, b)
}
func (m *GenerateCPFResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateCPFResponse.Marshal(b, m, deterministic)
}
func (m *GenerateCPFResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateCPFResponse.Merge(m, src)
}
func (m *GenerateCPFResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateCPFResponse.Size(m)
}
func (m *GenerateCPFResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateCPFResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateCPFResponse proto.InternalMessageInfo

func (m *GenerateCPFResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*GenerateCPFRequest)(nil), "api.cpf.GenerateCPFRequest")
	proto.RegisterType((*GenerateCPFResponse)(nil), "api.cpf.GenerateCPFResponse")
}

func init() { proto.RegisterFile("api/cpf/generator.proto", fileDescriptor_acd96ee7ee43acb6) }

var fileDescriptor_acd96ee7ee43acb6 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2c, 0xc8, 0xd4,
	0x4f, 0x2e, 0x48, 0xd3, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x2c, 0xc8, 0xd4, 0x4b, 0x2e, 0x48, 0x53, 0x12, 0xe1, 0x12,
	0x72, 0x87, 0xc8, 0xa5, 0x3a, 0x07, 0xb8, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0xe9,
	0x72, 0x09, 0xa3, 0x88, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x71, 0xb1, 0x15, 0xa5,
	0x16, 0x97, 0xe6, 0x94, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x46, 0x09, 0x28,
	0x86, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x79, 0x71, 0x71, 0x23, 0x89, 0x0a, 0x49,
	0xeb, 0x41, 0xed, 0xd4, 0xc3, 0xb4, 0x50, 0x4a, 0x06, 0xbb, 0x24, 0xc4, 0x5e, 0x25, 0x06, 0x27,
	0xd6, 0x28, 0xe6, 0xe4, 0x82, 0xb4, 0x24, 0x36, 0xb0, 0xeb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x6a, 0x7f, 0x88, 0xe1, 0xd8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GenerateCPFServiceClient is the client API for GenerateCPFService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GenerateCPFServiceClient interface {
	GenerateCPF(ctx context.Context, in *GenerateCPFRequest, opts ...grpc.CallOption) (*GenerateCPFResponse, error)
}

type generateCPFServiceClient struct {
	cc *grpc.ClientConn
}

func NewGenerateCPFServiceClient(cc *grpc.ClientConn) GenerateCPFServiceClient {
	return &generateCPFServiceClient{cc}
}

func (c *generateCPFServiceClient) GenerateCPF(ctx context.Context, in *GenerateCPFRequest, opts ...grpc.CallOption) (*GenerateCPFResponse, error) {
	out := new(GenerateCPFResponse)
	err := c.cc.Invoke(ctx, "/api.cpf.GenerateCPFService/GenerateCPF", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenerateCPFServiceServer is the server API for GenerateCPFService service.
type GenerateCPFServiceServer interface {
	GenerateCPF(context.Context, *GenerateCPFRequest) (*GenerateCPFResponse, error)
}

func RegisterGenerateCPFServiceServer(s *grpc.Server, srv GenerateCPFServiceServer) {
	s.RegisterService(&_GenerateCPFService_serviceDesc, srv)
}

func _GenerateCPFService_GenerateCPF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateCPFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerateCPFServiceServer).GenerateCPF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.cpf.GenerateCPFService/GenerateCPF",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerateCPFServiceServer).GenerateCPF(ctx, req.(*GenerateCPFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GenerateCPFService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.cpf.GenerateCPFService",
	HandlerType: (*GenerateCPFServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateCPF",
			Handler:    _GenerateCPFService_GenerateCPF_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/cpf/generator.proto",
}
