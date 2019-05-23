// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/configuration.proto

package resultstore

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Represents a configuration within an Invocation associated with one or more
// ConfiguredTargets. It captures the environment and other settings that
// were used.
type Configuration struct {
	// The format of this Configuration resource name must be:
	// invocations/${INVOCATION_ID}/configs/${CONFIG_ID}
	// The configuration ID of "default" should be preferred for the default
	// configuration in a single-config invocation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource ID components that identify the Configuration. They must match
	// the resource name after proper encoding.
	Id *Configuration_Id `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The aggregate status for this configuration.
	StatusAttributes *StatusAttributes `protobuf:"bytes,3,opt,name=status_attributes,json=statusAttributes,proto3" json:"status_attributes,omitempty"`
	// Attributes that apply only to this configuration.
	ConfigurationAttributes *ConfigurationAttributes `protobuf:"bytes,5,opt,name=configuration_attributes,json=configurationAttributes,proto3" json:"configuration_attributes,omitempty"`
	// Arbitrary name-value pairs.
	// This is implemented as a multi-map. Multiple properties are allowed with
	// the same key. Properties will be returned in lexicographical order by key.
	Properties           []*Property `protobuf:"bytes,6,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cb6c7dfd72e78e, []int{0}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Configuration) GetId() *Configuration_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Configuration) GetStatusAttributes() *StatusAttributes {
	if m != nil {
		return m.StatusAttributes
	}
	return nil
}

func (m *Configuration) GetConfigurationAttributes() *ConfigurationAttributes {
	if m != nil {
		return m.ConfigurationAttributes
	}
	return nil
}

func (m *Configuration) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

// The resource ID components that identify the Configuration.
type Configuration_Id struct {
	// The Invocation ID.
	InvocationId string `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	// The Configuration ID.
	ConfigurationId      string   `protobuf:"bytes,2,opt,name=configuration_id,json=configurationId,proto3" json:"configuration_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Configuration_Id) Reset()         { *m = Configuration_Id{} }
func (m *Configuration_Id) String() string { return proto.CompactTextString(m) }
func (*Configuration_Id) ProtoMessage()    {}
func (*Configuration_Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cb6c7dfd72e78e, []int{0, 0}
}

func (m *Configuration_Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration_Id.Unmarshal(m, b)
}
func (m *Configuration_Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration_Id.Marshal(b, m, deterministic)
}
func (m *Configuration_Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration_Id.Merge(m, src)
}
func (m *Configuration_Id) XXX_Size() int {
	return xxx_messageInfo_Configuration_Id.Size(m)
}
func (m *Configuration_Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration_Id proto.InternalMessageInfo

func (m *Configuration_Id) GetInvocationId() string {
	if m != nil {
		return m.InvocationId
	}
	return ""
}

func (m *Configuration_Id) GetConfigurationId() string {
	if m != nil {
		return m.ConfigurationId
	}
	return ""
}

// Attributes that apply only to the configuration.
type ConfigurationAttributes struct {
	// The type of cpu. (e.g. "x86", "powerpc")
	Cpu                  string   `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigurationAttributes) Reset()         { *m = ConfigurationAttributes{} }
func (m *ConfigurationAttributes) String() string { return proto.CompactTextString(m) }
func (*ConfigurationAttributes) ProtoMessage()    {}
func (*ConfigurationAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cb6c7dfd72e78e, []int{1}
}

func (m *ConfigurationAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigurationAttributes.Unmarshal(m, b)
}
func (m *ConfigurationAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigurationAttributes.Marshal(b, m, deterministic)
}
func (m *ConfigurationAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigurationAttributes.Merge(m, src)
}
func (m *ConfigurationAttributes) XXX_Size() int {
	return xxx_messageInfo_ConfigurationAttributes.Size(m)
}
func (m *ConfigurationAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigurationAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigurationAttributes proto.InternalMessageInfo

func (m *ConfigurationAttributes) GetCpu() string {
	if m != nil {
		return m.Cpu
	}
	return ""
}

func init() {
	proto.RegisterType((*Configuration)(nil), "google.devtools.resultstore.v2.Configuration")
	proto.RegisterType((*Configuration_Id)(nil), "google.devtools.resultstore.v2.Configuration.Id")
	proto.RegisterType((*ConfigurationAttributes)(nil), "google.devtools.resultstore.v2.ConfigurationAttributes")
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/configuration.proto", fileDescriptor_c7cb6c7dfd72e78e)
}

var fileDescriptor_c7cb6c7dfd72e78e = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4f, 0xb3, 0x40,
	0x10, 0xc6, 0x03, 0xbc, 0x6f, 0x93, 0x4e, 0x6d, 0xac, 0x7b, 0x29, 0xe9, 0xc1, 0x34, 0xf5, 0x82,
	0x69, 0xb2, 0x18, 0x3c, 0x78, 0xf0, 0xe2, 0x9f, 0x8b, 0xdc, 0x1a, 0xf4, 0x64, 0x62, 0x0c, 0x65,
	0xd7, 0xcd, 0x26, 0xc0, 0xe0, 0xee, 0x42, 0xe2, 0x37, 0xf5, 0xe3, 0x18, 0x81, 0x2a, 0x18, 0x6b,
	0xf5, 0x36, 0xcc, 0xf0, 0x7b, 0x9e, 0x99, 0x9d, 0x81, 0x40, 0x20, 0x8a, 0x94, 0xfb, 0x8c, 0x57,
	0x06, 0x31, 0xd5, 0xbe, 0xe2, 0xba, 0x4c, 0x8d, 0x36, 0xa8, 0xb8, 0x5f, 0x05, 0x7e, 0x82, 0xf9,
	0x93, 0x14, 0xa5, 0x8a, 0x8d, 0xc4, 0x9c, 0x16, 0x0a, 0x0d, 0x92, 0xc3, 0x86, 0xa1, 0x1b, 0x86,
	0x76, 0x18, 0x5a, 0x05, 0xb3, 0xe5, 0x4e, 0xcd, 0x2c, 0xdb, 0x88, 0x2d, 0x5e, 0x1d, 0x18, 0x5f,
	0x77, 0x4d, 0x08, 0x81, 0x7f, 0x79, 0x9c, 0x71, 0xd7, 0x9a, 0x5b, 0xde, 0x30, 0xaa, 0x63, 0x72,
	0x01, 0xb6, 0x64, 0xae, 0x3d, 0xb7, 0xbc, 0x51, 0x70, 0x42, 0x7f, 0xf6, 0xa7, 0x3d, 0x39, 0x1a,
	0xb2, 0xc8, 0x96, 0x8c, 0x3c, 0xc0, 0x81, 0x36, 0xb1, 0x29, 0xf5, 0x63, 0x6c, 0x8c, 0x92, 0xeb,
	0xd2, 0x70, 0xed, 0x3a, 0xbf, 0x13, 0xbc, 0xad, 0xc1, 0xcb, 0x0f, 0x2e, 0x9a, 0xe8, 0x2f, 0x19,
	0xa2, 0xc0, 0xed, 0x3d, 0x55, 0xd7, 0xe5, 0x7f, 0xed, 0x72, 0xf6, 0xa7, 0xb6, 0x3b, 0x66, 0xd3,
	0xe4, 0xfb, 0x02, 0xb9, 0x01, 0x28, 0x14, 0x16, 0x5c, 0x19, 0xc9, 0xb5, 0x3b, 0x98, 0x3b, 0xde,
	0x28, 0xf0, 0x76, 0xb9, 0xac, 0x1a, 0xe2, 0x25, 0xea, 0xb0, 0xb3, 0x3b, 0xb0, 0x43, 0x46, 0x8e,
	0x60, 0x2c, 0xf3, 0x0a, 0x93, 0x66, 0x00, 0xc9, 0xda, 0x0d, 0xec, 0x7d, 0x26, 0x43, 0x46, 0x8e,
	0x61, 0xd2, 0x1f, 0xb4, 0xdd, 0xcb, 0x30, 0xda, 0xef, 0xe5, 0x43, 0xb6, 0x58, 0xc2, 0x74, 0xcb,
	0x4c, 0x64, 0x02, 0x4e, 0x52, 0x94, 0xad, 0xc1, 0x7b, 0x78, 0xf5, 0x0c, 0x8b, 0x04, 0xb3, 0x1d,
	0xdd, 0xaf, 0xac, 0xfb, 0xb0, 0xfd, 0x43, 0x60, 0x1a, 0xe7, 0x82, 0xa2, 0x12, 0xbe, 0xe0, 0x79,
	0x7d, 0x4b, 0x7e, 0x53, 0x8a, 0x0b, 0xa9, 0xb7, 0xdd, 0xde, 0x79, 0xe7, 0x73, 0x3d, 0xa8, 0xa9,
	0xd3, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0xea, 0x27, 0xe2, 0x04, 0x03, 0x00, 0x00,
}
