// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/click_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [ClickViewService.GetClickView][google.ads.googleads.v1.services.ClickViewService.GetClickView].
type GetClickViewRequest struct {
	// The resource name of the click view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClickViewRequest) Reset()         { *m = GetClickViewRequest{} }
func (m *GetClickViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetClickViewRequest) ProtoMessage()    {}
func (*GetClickViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3c35caeb519e64a, []int{0}
}

func (m *GetClickViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClickViewRequest.Unmarshal(m, b)
}
func (m *GetClickViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClickViewRequest.Marshal(b, m, deterministic)
}
func (m *GetClickViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClickViewRequest.Merge(m, src)
}
func (m *GetClickViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetClickViewRequest.Size(m)
}
func (m *GetClickViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClickViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClickViewRequest proto.InternalMessageInfo

func (m *GetClickViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetClickViewRequest)(nil), "google.ads.googleads.v1.services.GetClickViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/click_view_service.proto", fileDescriptor_a3c35caeb519e64a)
}

var fileDescriptor_a3c35caeb519e64a = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0x4a, 0xc3, 0x50,
	0x14, 0x86, 0x49, 0x04, 0xc1, 0x50, 0x41, 0x22, 0x82, 0x14, 0x87, 0x52, 0x3b, 0x48, 0x29, 0xf7,
	0x92, 0xaa, 0x83, 0x57, 0x1c, 0x52, 0x87, 0x3a, 0x49, 0xa9, 0x90, 0x41, 0x02, 0xe5, 0x9a, 0x1c,
	0xc2, 0xc5, 0x26, 0xb7, 0xe6, 0xa4, 0xe9, 0x20, 0x2e, 0xbe, 0x82, 0x6f, 0xe0, 0xa6, 0x0f, 0xe1,
	0x03, 0xb8, 0xfa, 0x0a, 0x4e, 0x3e, 0x83, 0x83, 0xa4, 0xb7, 0x37, 0x58, 0x31, 0x74, 0xfb, 0x39,
	0xf9, 0xbf, 0x3f, 0xff, 0x39, 0x89, 0x75, 0x12, 0x49, 0x19, 0x8d, 0x81, 0xf2, 0x10, 0xa9, 0x92,
	0x85, 0xca, 0x1d, 0x8a, 0x90, 0xe6, 0x22, 0x00, 0xa4, 0xc1, 0x58, 0x04, 0xb7, 0xa3, 0x5c, 0xc0,
	0x6c, 0xb4, 0x98, 0x91, 0x49, 0x2a, 0x33, 0x69, 0x37, 0x94, 0x9f, 0xf0, 0x10, 0x49, 0x89, 0x92,
	0xdc, 0x21, 0x1a, 0xad, 0x77, 0xab, 0xc2, 0x53, 0x40, 0x39, 0x4d, 0x97, 0xd3, 0x55, 0x6a, 0x7d,
	0x4f, 0x33, 0x13, 0x41, 0x79, 0x92, 0xc8, 0x8c, 0x67, 0x42, 0x26, 0xa8, 0x9e, 0x36, 0x99, 0xb5,
	0xdd, 0x87, 0xec, 0xbc, 0x80, 0x3c, 0x01, 0xb3, 0x21, 0xdc, 0x4d, 0x01, 0x33, 0x7b, 0xdf, 0xda,
	0xd4, 0x91, 0xa3, 0x84, 0xc7, 0xb0, 0x6b, 0x34, 0x8c, 0x83, 0x8d, 0x61, 0x4d, 0x0f, 0x2f, 0x79,
	0x0c, 0xdd, 0x37, 0xc3, 0xda, 0x2a, 0xc9, 0x2b, 0xd5, 0xd1, 0x7e, 0x31, 0xac, 0xda, 0xef, 0x44,
	0xfb, 0x98, 0xac, 0x5a, 0x8b, 0xfc, 0xd3, 0xa0, 0xde, 0xa9, 0xc4, 0xca, 0x5d, 0x49, 0x09, 0x35,
	0x8f, 0x1e, 0x3f, 0x3e, 0x9f, 0x4c, 0x62, 0x77, 0x8a, 0x63, 0xdc, 0x2f, 0x55, 0x3f, 0x0b, 0xa6,
	0x98, 0xc9, 0x18, 0x52, 0xa4, 0x6d, 0x75, 0x9d, 0x82, 0x40, 0xda, 0x7e, 0xe8, 0x7d, 0x1b, 0x56,
	0x2b, 0x90, 0xf1, 0xca, 0x82, 0xbd, 0x9d, 0xbf, 0x6b, 0x0e, 0x8a, 0xe3, 0x0d, 0x8c, 0xeb, 0x8b,
	0x05, 0x1a, 0xc9, 0x31, 0x4f, 0x22, 0x22, 0xd3, 0x88, 0x46, 0x90, 0xcc, 0x4f, 0xab, 0x3f, 0xd0,
	0x44, 0x60, 0xf5, 0xcf, 0x70, 0xaa, 0xc5, 0xb3, 0xb9, 0xd6, 0x77, 0xdd, 0x57, 0xb3, 0xd1, 0x57,
	0x81, 0x6e, 0x88, 0x44, 0xc9, 0x42, 0x79, 0x0e, 0x59, 0xbc, 0x18, 0xdf, 0xb5, 0xc5, 0x77, 0x43,
	0xf4, 0x4b, 0x8b, 0xef, 0x39, 0xbe, 0xb6, 0x7c, 0x99, 0x2d, 0x35, 0x67, 0xcc, 0x0d, 0x91, 0xb1,
	0xd2, 0xc4, 0x98, 0xe7, 0x30, 0xa6, 0x6d, 0x37, 0xeb, 0xf3, 0x9e, 0x87, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xe1, 0xc6, 0x06, 0x3b, 0xb3, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClickViewServiceClient is the client API for ClickViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClickViewServiceClient interface {
	// Returns the requested click view in full detail.
	GetClickView(ctx context.Context, in *GetClickViewRequest, opts ...grpc.CallOption) (*resources.ClickView, error)
}

type clickViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewClickViewServiceClient(cc *grpc.ClientConn) ClickViewServiceClient {
	return &clickViewServiceClient{cc}
}

func (c *clickViewServiceClient) GetClickView(ctx context.Context, in *GetClickViewRequest, opts ...grpc.CallOption) (*resources.ClickView, error) {
	out := new(resources.ClickView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ClickViewService/GetClickView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClickViewServiceServer is the server API for ClickViewService service.
type ClickViewServiceServer interface {
	// Returns the requested click view in full detail.
	GetClickView(context.Context, *GetClickViewRequest) (*resources.ClickView, error)
}

func RegisterClickViewServiceServer(s *grpc.Server, srv ClickViewServiceServer) {
	s.RegisterService(&_ClickViewService_serviceDesc, srv)
}

func _ClickViewService_GetClickView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClickViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickViewServiceServer).GetClickView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ClickViewService/GetClickView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickViewServiceServer).GetClickView(ctx, req.(*GetClickViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClickViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ClickViewService",
	HandlerType: (*ClickViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClickView",
			Handler:    _ClickViewService_GetClickView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/click_view_service.proto",
}
