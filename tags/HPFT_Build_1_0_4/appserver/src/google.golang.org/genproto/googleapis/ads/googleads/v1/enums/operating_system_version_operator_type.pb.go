// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/operating_system_version_operator_type.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The type of operating system version.
type OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType int32

const (
	// Not specified.
	OperatingSystemVersionOperatorTypeEnum_UNSPECIFIED OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType = 0
	// Used for return value only. Represents value unknown in this version.
	OperatingSystemVersionOperatorTypeEnum_UNKNOWN OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType = 1
	// Equals to the specified version.
	OperatingSystemVersionOperatorTypeEnum_EQUALS_TO OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType = 2
	// Greater than or equals to the specified version.
	OperatingSystemVersionOperatorTypeEnum_GREATER_THAN_EQUALS_TO OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType = 4
)

var OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "EQUALS_TO",
	4: "GREATER_THAN_EQUALS_TO",
}

var OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"EQUALS_TO":              2,
	"GREATER_THAN_EQUALS_TO": 4,
}

func (x OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType) String() string {
	return proto.EnumName(OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType_name, int32(x))
}

func (OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_87a572f6e7ce491f, []int{0, 0}
}

// Container for enum describing the type of OS operators.
type OperatingSystemVersionOperatorTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatingSystemVersionOperatorTypeEnum) Reset() {
	*m = OperatingSystemVersionOperatorTypeEnum{}
}
func (m *OperatingSystemVersionOperatorTypeEnum) String() string { return proto.CompactTextString(m) }
func (*OperatingSystemVersionOperatorTypeEnum) ProtoMessage()    {}
func (*OperatingSystemVersionOperatorTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_87a572f6e7ce491f, []int{0}
}

func (m *OperatingSystemVersionOperatorTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatingSystemVersionOperatorTypeEnum.Unmarshal(m, b)
}
func (m *OperatingSystemVersionOperatorTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatingSystemVersionOperatorTypeEnum.Marshal(b, m, deterministic)
}
func (m *OperatingSystemVersionOperatorTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatingSystemVersionOperatorTypeEnum.Merge(m, src)
}
func (m *OperatingSystemVersionOperatorTypeEnum) XXX_Size() int {
	return xxx_messageInfo_OperatingSystemVersionOperatorTypeEnum.Size(m)
}
func (m *OperatingSystemVersionOperatorTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatingSystemVersionOperatorTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_OperatingSystemVersionOperatorTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType", OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType_name, OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType_value)
	proto.RegisterType((*OperatingSystemVersionOperatorTypeEnum)(nil), "google.ads.googleads.v1.enums.OperatingSystemVersionOperatorTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/operating_system_version_operator_type.proto", fileDescriptor_87a572f6e7ce491f)
}

var fileDescriptor_87a572f6e7ce491f = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x4f, 0x4b, 0xc3, 0x30,
	0x1c, 0xb5, 0x55, 0x14, 0x33, 0xc4, 0xd1, 0x83, 0x87, 0xe1, 0x0e, 0xdb, 0x41, 0x6f, 0x29, 0xc5,
	0x5b, 0x3c, 0x65, 0x5a, 0xe7, 0x54, 0xda, 0xb9, 0x75, 0x15, 0xa4, 0x10, 0xaa, 0x0d, 0xa1, 0xb0,
	0x26, 0xa5, 0xc9, 0x06, 0xfb, 0x24, 0xde, 0x3d, 0xfa, 0x51, 0xfc, 0x28, 0x7e, 0x06, 0x0f, 0xd2,
	0x64, 0x9d, 0x27, 0xdd, 0x25, 0x3c, 0xf2, 0xde, 0xef, 0xbd, 0xdf, 0x1f, 0x70, 0xc7, 0x84, 0x60,
	0x73, 0xea, 0xa6, 0x99, 0x74, 0x0d, 0xac, 0xd1, 0xd2, 0x73, 0x29, 0x5f, 0x14, 0xd2, 0x15, 0x25,
	0xad, 0x52, 0x95, 0x73, 0x46, 0xe4, 0x4a, 0x2a, 0x5a, 0x90, 0x25, 0xad, 0x64, 0x2e, 0x38, 0x31,
	0x84, 0xa8, 0x88, 0x5a, 0x95, 0x14, 0x96, 0x95, 0x50, 0xc2, 0xe9, 0x1a, 0x03, 0x98, 0x66, 0x12,
	0x6e, 0xbc, 0xe0, 0xd2, 0x83, 0xda, 0xab, 0x73, 0xda, 0x44, 0x95, 0xb9, 0x9b, 0x72, 0x2e, 0x54,
	0xaa, 0x72, 0xc1, 0xa5, 0x29, 0xee, 0xbf, 0x59, 0xe0, 0x2c, 0x6c, 0xd2, 0xa6, 0x3a, 0x2c, 0x36,
	0x59, 0xe1, 0x3a, 0x2a, 0x5a, 0x95, 0xd4, 0xe7, 0x8b, 0xa2, 0x5f, 0x80, 0xfe, 0x76, 0xa5, 0x73,
	0x0c, 0x5a, 0xb3, 0x60, 0x3a, 0xf6, 0xaf, 0x46, 0x37, 0x23, 0xff, 0xba, 0xbd, 0xe3, 0xb4, 0xc0,
	0xc1, 0x2c, 0xb8, 0x0f, 0xc2, 0xa7, 0xa0, 0x6d, 0x39, 0x47, 0xe0, 0xd0, 0x7f, 0x9c, 0xe1, 0x87,
	0x29, 0x89, 0xc2, 0xb6, 0xed, 0x74, 0xc0, 0xc9, 0x70, 0xe2, 0xe3, 0xc8, 0x9f, 0x90, 0xe8, 0x16,
	0x07, 0xe4, 0x97, 0xdb, 0x1b, 0x7c, 0x5b, 0xa0, 0xf7, 0x2a, 0x0a, 0xf8, 0xef, 0x74, 0x83, 0xf3,
	0xed, 0x2d, 0x8d, 0xeb, 0x41, 0xc7, 0xd6, 0xf3, 0x60, 0xed, 0xc4, 0xc4, 0x3c, 0xe5, 0x0c, 0x8a,
	0x8a, 0xb9, 0x8c, 0x72, 0xbd, 0x86, 0xe6, 0x06, 0x65, 0x2e, 0xff, 0x38, 0xc9, 0xa5, 0x7e, 0xdf,
	0xed, 0xdd, 0x21, 0xc6, 0x1f, 0x76, 0x77, 0x68, 0xac, 0x70, 0x26, 0xa1, 0x81, 0x35, 0x8a, 0x3d,
	0x58, 0x2f, 0x4a, 0x7e, 0x36, 0x7c, 0x82, 0x33, 0x99, 0x6c, 0xf8, 0x24, 0xf6, 0x12, 0xcd, 0x7f,
	0xd9, 0x3d, 0xf3, 0x89, 0x10, 0xce, 0x24, 0x42, 0x1b, 0x05, 0x42, 0xb1, 0x87, 0x90, 0xd6, 0xbc,
	0xec, 0xeb, 0xc6, 0x2e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x40, 0xf6, 0xcb, 0x8a, 0x2a, 0x02,
	0x00, 0x00,
}
