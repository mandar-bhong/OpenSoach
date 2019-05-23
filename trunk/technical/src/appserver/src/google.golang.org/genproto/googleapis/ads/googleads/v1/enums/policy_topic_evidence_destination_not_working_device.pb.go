// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/policy_topic_evidence_destination_not_working_device.proto

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

// The possible policy topic evidence destination not working devices.
type PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice int32

const (
	// No value has been specified.
	PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_UNSPECIFIED PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_UNKNOWN PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice = 1
	// Landing page doesn't work on desktop device.
	PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_DESKTOP PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice = 2
	// Landing page doesn't work on Android device.
	PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_ANDROID PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice = 3
	// Landing page doesn't work on iOS device.
	PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_IOS PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice = 4
)

var PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DESKTOP",
	3: "ANDROID",
	4: "IOS",
}

var PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"DESKTOP":     2,
	"ANDROID":     3,
	"IOS":         4,
}

func (x PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice) String() string {
	return proto.EnumName(PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice_name, int32(x))
}

func (PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b10d0781dd50ac63, []int{0, 0}
}

// Container for enum describing possible policy topic evidence destination not
// working devices.
type PolicyTopicEvidenceDestinationNotWorkingDeviceEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyTopicEvidenceDestinationNotWorkingDeviceEnum) Reset() {
	*m = PolicyTopicEvidenceDestinationNotWorkingDeviceEnum{}
}
func (m *PolicyTopicEvidenceDestinationNotWorkingDeviceEnum) String() string {
	return proto.CompactTextString(m)
}
func (*PolicyTopicEvidenceDestinationNotWorkingDeviceEnum) ProtoMessage() {}
func (*PolicyTopicEvidenceDestinationNotWorkingDeviceEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b10d0781dd50ac63, []int{0}
}

func (m *PolicyTopicEvidenceDestinationNotWorkingDeviceEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyTopicEvidenceDestinationNotWorkingDeviceEnum.Unmarshal(m, b)
}
func (m *PolicyTopicEvidenceDestinationNotWorkingDeviceEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyTopicEvidenceDestinationNotWorkingDeviceEnum.Marshal(b, m, deterministic)
}
func (m *PolicyTopicEvidenceDestinationNotWorkingDeviceEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyTopicEvidenceDestinationNotWorkingDeviceEnum.Merge(m, src)
}
func (m *PolicyTopicEvidenceDestinationNotWorkingDeviceEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyTopicEvidenceDestinationNotWorkingDeviceEnum.Size(m)
}
func (m *PolicyTopicEvidenceDestinationNotWorkingDeviceEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyTopicEvidenceDestinationNotWorkingDeviceEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyTopicEvidenceDestinationNotWorkingDeviceEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice", PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice_name, PolicyTopicEvidenceDestinationNotWorkingDeviceEnum_PolicyTopicEvidenceDestinationNotWorkingDevice_value)
	proto.RegisterType((*PolicyTopicEvidenceDestinationNotWorkingDeviceEnum)(nil), "google.ads.googleads.v1.enums.PolicyTopicEvidenceDestinationNotWorkingDeviceEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/policy_topic_evidence_destination_not_working_device.proto", fileDescriptor_b10d0781dd50ac63)
}

var fileDescriptor_b10d0781dd50ac63 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x6d, 0x2a, 0x16, 0xd2, 0x83, 0x21, 0x47, 0xb1, 0x87, 0xf6, 0x01, 0x36, 0xc4, 0xde,
	0xd6, 0x53, 0x6a, 0x62, 0x09, 0x85, 0x24, 0xd8, 0x7f, 0x22, 0x81, 0x10, 0xb3, 0x4b, 0x58, 0x6c,
	0x77, 0x62, 0x37, 0x8d, 0xf8, 0x06, 0x3e, 0x87, 0x17, 0xc1, 0x47, 0xf1, 0x51, 0x7c, 0x0a, 0xd9,
	0xdd, 0xb6, 0x9e, 0x14, 0x7a, 0x09, 0xdf, 0x64, 0x66, 0xbe, 0xdf, 0xec, 0x8c, 0x79, 0x5f, 0x02,
	0x94, 0x2b, 0xea, 0xe4, 0x44, 0x38, 0x5a, 0x4a, 0xd5, 0xb8, 0x0e, 0xe5, 0xdb, 0xb5, 0x70, 0x2a,
	0x58, 0xb1, 0xe2, 0x35, 0xab, 0xa1, 0x62, 0x45, 0x46, 0x1b, 0x46, 0x28, 0x2f, 0x68, 0x46, 0xa8,
	0xa8, 0x19, 0xcf, 0x6b, 0x06, 0x3c, 0xe3, 0x50, 0x67, 0x2f, 0xb0, 0x79, 0x62, 0xbc, 0xcc, 0x08,
	0x6d, 0x58, 0x41, 0x51, 0xb5, 0x81, 0x1a, 0xec, 0x9e, 0xb6, 0x43, 0x39, 0x11, 0xe8, 0xe0, 0x8c,
	0x1a, 0x17, 0x29, 0xe7, 0x8b, 0xcb, 0x3d, 0xb8, 0x62, 0x4e, 0xce, 0x39, 0xd4, 0xca, 0x4e, 0xe8,
	0xe6, 0xc1, 0x47, 0xcb, 0xbc, 0x4a, 0x14, 0x7b, 0x26, 0xd1, 0xc1, 0x8e, 0xec, 0xff, 0x82, 0x23,
	0xa8, 0x97, 0x1a, 0xeb, 0x2b, 0x6a, 0xc0, 0xb7, 0xeb, 0xc1, 0xb3, 0x89, 0x8e, 0xeb, 0xb2, 0xcf,
	0xcd, 0xee, 0x3c, 0x9a, 0x26, 0xc1, 0x4d, 0x78, 0x1b, 0x06, 0xbe, 0x75, 0x62, 0x77, 0xcd, 0xce,
	0x3c, 0x9a, 0x44, 0xf1, 0x32, 0xb2, 0x5a, 0x32, 0xf0, 0x83, 0xe9, 0x64, 0x16, 0x27, 0x96, 0x21,
	0x03, 0x2f, 0xf2, 0xef, 0xe2, 0xd0, 0xb7, 0xda, 0x76, 0xc7, 0x6c, 0x87, 0xf1, 0xd4, 0x3a, 0x1d,
	0xbd, 0x19, 0x66, 0xbf, 0x80, 0x35, 0xfa, 0xf7, 0xb5, 0xa3, 0xe1, 0x71, 0x63, 0x25, 0x72, 0x09,
	0x49, 0xeb, 0x61, 0xb4, 0x73, 0x2d, 0x61, 0x95, 0xf3, 0x12, 0xc1, 0xa6, 0x74, 0x4a, 0xca, 0xd5,
	0x8a, 0xf6, 0xd7, 0xaa, 0x98, 0xf8, 0xe3, 0x78, 0xd7, 0xea, 0xfb, 0x6e, 0xb4, 0xc7, 0x9e, 0xf7,
	0x69, 0xf4, 0xc6, 0xda, 0xca, 0x23, 0x02, 0x69, 0x29, 0xd5, 0xc2, 0x45, 0x72, 0x71, 0xe2, 0x6b,
	0x9f, 0x4f, 0x3d, 0x22, 0xd2, 0x43, 0x3e, 0x5d, 0xb8, 0xa9, 0xca, 0x7f, 0x1b, 0x7d, 0xfd, 0x13,
	0x63, 0x8f, 0x08, 0x8c, 0x0f, 0x15, 0x18, 0x2f, 0x5c, 0x8c, 0x55, 0xcd, 0xe3, 0x99, 0x1a, 0x6c,
	0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xfa, 0x15, 0x08, 0x54, 0x02, 0x00, 0x00,
}
