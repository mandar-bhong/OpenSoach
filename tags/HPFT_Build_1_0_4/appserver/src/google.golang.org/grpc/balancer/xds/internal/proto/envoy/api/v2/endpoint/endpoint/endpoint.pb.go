// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/endpoint/endpoint.proto

package endpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import address "google.golang.org/grpc/balancer/xds/internal/proto/envoy/api/v2/core/address"
import base "google.golang.org/grpc/balancer/xds/internal/proto/envoy/api/v2/core/base"
import health_check "google.golang.org/grpc/balancer/xds/internal/proto/envoy/api/v2/core/health_check"
import _ "google.golang.org/grpc/balancer/xds/internal/proto/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Endpoint struct {
	Address              *address.Address            `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	HealthCheckConfig    *Endpoint_HealthCheckConfig `protobuf:"bytes,2,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_endpoint_2d1a533d75f3064c, []int{0}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (dst *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(dst, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetAddress() *address.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetHealthCheckConfig() *Endpoint_HealthCheckConfig {
	if m != nil {
		return m.HealthCheckConfig
	}
	return nil
}

type Endpoint_HealthCheckConfig struct {
	PortValue            uint32   `protobuf:"varint,1,opt,name=port_value,json=portValue,proto3" json:"port_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint_HealthCheckConfig) Reset()         { *m = Endpoint_HealthCheckConfig{} }
func (m *Endpoint_HealthCheckConfig) String() string { return proto.CompactTextString(m) }
func (*Endpoint_HealthCheckConfig) ProtoMessage()    {}
func (*Endpoint_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_endpoint_2d1a533d75f3064c, []int{0, 0}
}
func (m *Endpoint_HealthCheckConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Unmarshal(m, b)
}
func (m *Endpoint_HealthCheckConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Marshal(b, m, deterministic)
}
func (dst *Endpoint_HealthCheckConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint_HealthCheckConfig.Merge(dst, src)
}
func (m *Endpoint_HealthCheckConfig) XXX_Size() int {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Size(m)
}
func (m *Endpoint_HealthCheckConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint_HealthCheckConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint_HealthCheckConfig proto.InternalMessageInfo

func (m *Endpoint_HealthCheckConfig) GetPortValue() uint32 {
	if m != nil {
		return m.PortValue
	}
	return 0
}

type LbEndpoint struct {
	// Types that are valid to be assigned to HostIdentifier:
	//	*LbEndpoint_Endpoint
	//	*LbEndpoint_EndpointName
	HostIdentifier       isLbEndpoint_HostIdentifier `protobuf_oneof:"host_identifier"`
	HealthStatus         health_check.HealthStatus   `protobuf:"varint,2,opt,name=health_status,json=healthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"health_status,omitempty"`
	Metadata             *base.Metadata              `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	LoadBalancingWeight  *wrappers.UInt32Value       `protobuf:"bytes,4,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *LbEndpoint) Reset()         { *m = LbEndpoint{} }
func (m *LbEndpoint) String() string { return proto.CompactTextString(m) }
func (*LbEndpoint) ProtoMessage()    {}
func (*LbEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_endpoint_2d1a533d75f3064c, []int{1}
}
func (m *LbEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LbEndpoint.Unmarshal(m, b)
}
func (m *LbEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LbEndpoint.Marshal(b, m, deterministic)
}
func (dst *LbEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LbEndpoint.Merge(dst, src)
}
func (m *LbEndpoint) XXX_Size() int {
	return xxx_messageInfo_LbEndpoint.Size(m)
}
func (m *LbEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_LbEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_LbEndpoint proto.InternalMessageInfo

type isLbEndpoint_HostIdentifier interface {
	isLbEndpoint_HostIdentifier()
}

type LbEndpoint_Endpoint struct {
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3,oneof"`
}

type LbEndpoint_EndpointName struct {
	EndpointName string `protobuf:"bytes,5,opt,name=endpoint_name,json=endpointName,proto3,oneof"`
}

func (*LbEndpoint_Endpoint) isLbEndpoint_HostIdentifier() {}

func (*LbEndpoint_EndpointName) isLbEndpoint_HostIdentifier() {}

func (m *LbEndpoint) GetHostIdentifier() isLbEndpoint_HostIdentifier {
	if m != nil {
		return m.HostIdentifier
	}
	return nil
}

func (m *LbEndpoint) GetEndpoint() *Endpoint {
	if x, ok := m.GetHostIdentifier().(*LbEndpoint_Endpoint); ok {
		return x.Endpoint
	}
	return nil
}

func (m *LbEndpoint) GetEndpointName() string {
	if x, ok := m.GetHostIdentifier().(*LbEndpoint_EndpointName); ok {
		return x.EndpointName
	}
	return ""
}

func (m *LbEndpoint) GetHealthStatus() health_check.HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return health_check.HealthStatus_UNKNOWN
}

func (m *LbEndpoint) GetMetadata() *base.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LbEndpoint) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LbEndpoint) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LbEndpoint_OneofMarshaler, _LbEndpoint_OneofUnmarshaler, _LbEndpoint_OneofSizer, []interface{}{
		(*LbEndpoint_Endpoint)(nil),
		(*LbEndpoint_EndpointName)(nil),
	}
}

func _LbEndpoint_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LbEndpoint)
	// host_identifier
	switch x := m.HostIdentifier.(type) {
	case *LbEndpoint_Endpoint:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Endpoint); err != nil {
			return err
		}
	case *LbEndpoint_EndpointName:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.EndpointName)
	case nil:
	default:
		return fmt.Errorf("LbEndpoint.HostIdentifier has unexpected type %T", x)
	}
	return nil
}

func _LbEndpoint_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LbEndpoint)
	switch tag {
	case 1: // host_identifier.endpoint
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Endpoint)
		err := b.DecodeMessage(msg)
		m.HostIdentifier = &LbEndpoint_Endpoint{msg}
		return true, err
	case 5: // host_identifier.endpoint_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.HostIdentifier = &LbEndpoint_EndpointName{x}
		return true, err
	default:
		return false, nil
	}
}

func _LbEndpoint_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LbEndpoint)
	// host_identifier
	switch x := m.HostIdentifier.(type) {
	case *LbEndpoint_Endpoint:
		s := proto.Size(x.Endpoint)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LbEndpoint_EndpointName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.EndpointName)))
		n += len(x.EndpointName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LocalityLbEndpoints struct {
	Locality             *base.Locality        `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	LbEndpoints          []*LbEndpoint         `protobuf:"bytes,2,rep,name=lb_endpoints,json=lbEndpoints,proto3" json:"lb_endpoints,omitempty"`
	LoadBalancingWeight  *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	Priority             uint32                `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LocalityLbEndpoints) Reset()         { *m = LocalityLbEndpoints{} }
func (m *LocalityLbEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityLbEndpoints) ProtoMessage()    {}
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_endpoint_2d1a533d75f3064c, []int{2}
}
func (m *LocalityLbEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalityLbEndpoints.Unmarshal(m, b)
}
func (m *LocalityLbEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalityLbEndpoints.Marshal(b, m, deterministic)
}
func (dst *LocalityLbEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalityLbEndpoints.Merge(dst, src)
}
func (m *LocalityLbEndpoints) XXX_Size() int {
	return xxx_messageInfo_LocalityLbEndpoints.Size(m)
}
func (m *LocalityLbEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalityLbEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_LocalityLbEndpoints proto.InternalMessageInfo

func (m *LocalityLbEndpoints) GetLocality() *base.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLbEndpoints() []*LbEndpoint {
	if m != nil {
		return m.LbEndpoints
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

func (m *LocalityLbEndpoints) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "envoy.api.v2.endpoint.Endpoint")
	proto.RegisterType((*Endpoint_HealthCheckConfig)(nil), "envoy.api.v2.endpoint.Endpoint.HealthCheckConfig")
	proto.RegisterType((*LbEndpoint)(nil), "envoy.api.v2.endpoint.LbEndpoint")
	proto.RegisterType((*LocalityLbEndpoints)(nil), "envoy.api.v2.endpoint.LocalityLbEndpoints")
}

func init() {
	proto.RegisterFile("envoy/api/v2/endpoint/endpoint.proto", fileDescriptor_endpoint_2d1a533d75f3064c)
}

var fileDescriptor_endpoint_2d1a533d75f3064c = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xc1, 0x6e, 0xd3, 0x4c,
	0x18, 0xac, 0xe3, 0xa6, 0x4d, 0x37, 0xc9, 0xff, 0x2b, 0x8e, 0x2a, 0xac, 0x50, 0xd1, 0x12, 0x0a,
	0x8a, 0x72, 0x58, 0x8b, 0x14, 0x89, 0x13, 0x02, 0xdc, 0x22, 0x05, 0xa9, 0xa0, 0x6a, 0x11, 0x20,
	0x71, 0xc0, 0xfa, 0x6c, 0x6f, 0xe2, 0x15, 0x8e, 0xd7, 0xb2, 0x37, 0x2e, 0xb9, 0xf5, 0x41, 0x78,
	0x12, 0x4e, 0x3c, 0x4c, 0x2f, 0x3c, 0x45, 0x91, 0xd7, 0x5e, 0x37, 0x6d, 0x52, 0x71, 0xe1, 0xb6,
	0xde, 0x99, 0xf9, 0xbe, 0x99, 0xf1, 0xa2, 0x43, 0x1a, 0x65, 0x7c, 0x61, 0x41, 0xcc, 0xac, 0x6c,
	0x64, 0xd1, 0xc8, 0x8f, 0x39, 0x8b, 0x44, 0x75, 0xc0, 0x71, 0xc2, 0x05, 0x37, 0x76, 0x25, 0x0b,
	0x43, 0xcc, 0x70, 0x36, 0xc2, 0x0a, 0xec, 0xed, 0xdf, 0x10, 0x7b, 0x3c, 0xa1, 0x16, 0xf8, 0x7e,
	0x42, 0xd3, 0xb4, 0xd0, 0xf5, 0xf6, 0x56, 0x09, 0x2e, 0xa4, 0xb4, 0x44, 0x0f, 0x57, 0xd1, 0x80,
	0x42, 0x28, 0x02, 0xc7, 0x0b, 0xa8, 0xf7, 0xad, 0x64, 0x3d, 0x98, 0x72, 0x3e, 0x0d, 0xa9, 0x25,
	0xbf, 0xdc, 0xf9, 0xc4, 0x3a, 0x4f, 0x20, 0x8e, 0x69, 0xa2, 0x76, 0xdc, 0xcb, 0x20, 0x64, 0x3e,
	0x08, 0x6a, 0xa9, 0x43, 0x01, 0xf4, 0x2f, 0x35, 0xd4, 0x78, 0x53, 0x5a, 0x35, 0x9e, 0xa1, 0xed,
	0xd2, 0x9a, 0xa9, 0x1d, 0x68, 0x83, 0xe6, 0xa8, 0x87, 0x6f, 0x64, 0xca, 0xb7, 0xe3, 0xd7, 0x05,
	0x83, 0x28, 0xaa, 0x01, 0xa8, 0xbb, 0xec, 0xc8, 0xf1, 0x78, 0x34, 0x61, 0x53, 0xb3, 0x26, 0x27,
	0x3c, 0xc5, 0x6b, 0x5b, 0xc1, 0x6a, 0x27, 0x1e, 0x4b, 0xe9, 0x71, 0xae, 0x3c, 0x96, 0x42, 0xd2,
	0x09, 0x6e, 0x5f, 0xf5, 0x5e, 0xa2, 0xce, 0x0a, 0xcf, 0x18, 0x22, 0x14, 0xf3, 0x44, 0x38, 0x19,
	0x84, 0x73, 0x2a, 0x0d, 0xb7, 0xed, 0xe6, 0xcf, 0xdf, 0xbf, 0xf4, 0xad, 0xe1, 0xa6, 0x79, 0x75,
	0xa5, 0x93, 0x9d, 0x1c, 0xfe, 0x94, 0xa3, 0xfd, 0xcb, 0x1a, 0x42, 0xa7, 0x6e, 0x15, 0xf4, 0x05,
	0x6a, 0x28, 0x27, 0x65, 0xd2, 0xfd, 0xbf, 0xf8, 0x1c, 0x6f, 0x90, 0x4a, 0x62, 0x3c, 0x46, 0x6d,
	0x75, 0x76, 0x22, 0x98, 0x51, 0xb3, 0x7e, 0xa0, 0x0d, 0x76, 0xc6, 0x1b, 0xa4, 0xa5, 0xae, 0xdf,
	0xc3, 0x8c, 0x1a, 0x27, 0xa8, 0x5d, 0x16, 0x93, 0x0a, 0x10, 0xf3, 0x54, 0x56, 0xf2, 0xdf, 0xed,
	0x55, 0xb2, 0xd4, 0x22, 0xdd, 0x07, 0x49, 0x23, 0xad, 0x60, 0xe9, 0xcb, 0x78, 0x8e, 0x1a, 0x33,
	0x2a, 0xc0, 0x07, 0x01, 0xa6, 0x2e, 0xbd, 0xde, 0x5f, 0x33, 0xe0, 0x5d, 0x49, 0x21, 0x15, 0xd9,
	0xf8, 0x8a, 0x76, 0x43, 0x0e, 0xbe, 0xe3, 0x42, 0x08, 0x91, 0xc7, 0xa2, 0xa9, 0x73, 0x4e, 0xd9,
	0x34, 0x10, 0xe6, 0xa6, 0x9c, 0xb2, 0x87, 0x8b, 0x37, 0x83, 0xd5, 0x9b, 0xc1, 0x1f, 0xdf, 0x46,
	0xe2, 0x68, 0x24, 0x0b, 0xb3, 0x5b, 0x79, 0x91, 0xdb, 0xc3, 0xba, 0x79, 0xa1, 0x0d, 0x34, 0xd2,
	0xcd, 0x07, 0xd9, 0x6a, 0xce, 0x67, 0x39, 0xc6, 0xee, 0xa0, 0xff, 0x03, 0x9e, 0x0a, 0x87, 0xf9,
	0x34, 0x12, 0x6c, 0xc2, 0x68, 0xd2, 0xff, 0x51, 0x43, 0xdd, 0x53, 0xee, 0x41, 0xc8, 0xc4, 0xe2,
	0xba, 0x6e, 0x99, 0x21, 0x2c, 0xaf, 0xcb, 0xbe, 0xd7, 0x65, 0x50, 0x4a, 0x52, 0x91, 0x8d, 0x13,
	0xd4, 0x0a, 0x5d, 0x47, 0xb5, 0x9a, 0x37, 0xa8, 0x0f, 0x9a, 0xa3, 0x87, 0x77, 0xfc, 0xac, 0xeb,
	0x95, 0xa4, 0x19, 0x2e, 0xad, 0xbf, 0xb3, 0x09, 0xfd, 0x9f, 0x34, 0x61, 0x3c, 0x41, 0x8d, 0x38,
	0x61, 0x3c, 0xc9, 0xe3, 0xd5, 0xe5, 0x3b, 0x44, 0xb9, 0xa8, 0x3e, 0xd4, 0xcd, 0x0b, 0x8d, 0x54,
	0x98, 0xfd, 0x0a, 0x3d, 0x62, 0xbc, 0xf0, 0x1e, 0x27, 0xfc, 0xfb, 0x62, 0x7d, 0x0c, 0xbb, 0xad,
	0x9c, 0x9f, 0xe5, 0x7e, 0xce, 0xb4, 0x2f, 0xd5, 0xcb, 0x73, 0xb7, 0xa4, 0xc5, 0xa3, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x6b, 0xdf, 0xf4, 0xe1, 0x92, 0x04, 0x00, 0x00,
}
