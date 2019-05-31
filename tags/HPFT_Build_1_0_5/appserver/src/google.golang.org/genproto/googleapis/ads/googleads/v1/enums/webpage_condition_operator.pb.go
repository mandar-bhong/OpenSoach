// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/webpage_condition_operator.proto

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

// The webpage condition operator in webpage criterion.
type WebpageConditionOperatorEnum_WebpageConditionOperator int32

const (
	// Not specified.
	WebpageConditionOperatorEnum_UNSPECIFIED WebpageConditionOperatorEnum_WebpageConditionOperator = 0
	// Used for return value only. Represents value unknown in this version.
	WebpageConditionOperatorEnum_UNKNOWN WebpageConditionOperatorEnum_WebpageConditionOperator = 1
	// The argument web condition is equal to the compared web condition.
	WebpageConditionOperatorEnum_EQUALS WebpageConditionOperatorEnum_WebpageConditionOperator = 2
	// The argument web condition is part of the compared web condition.
	WebpageConditionOperatorEnum_CONTAINS WebpageConditionOperatorEnum_WebpageConditionOperator = 3
)

var WebpageConditionOperatorEnum_WebpageConditionOperator_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "EQUALS",
	3: "CONTAINS",
}

var WebpageConditionOperatorEnum_WebpageConditionOperator_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"EQUALS":      2,
	"CONTAINS":    3,
}

func (x WebpageConditionOperatorEnum_WebpageConditionOperator) String() string {
	return proto.EnumName(WebpageConditionOperatorEnum_WebpageConditionOperator_name, int32(x))
}

func (WebpageConditionOperatorEnum_WebpageConditionOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e856da14d9f47c05, []int{0, 0}
}

// Container for enum describing webpage condition operator in webpage
// criterion.
type WebpageConditionOperatorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebpageConditionOperatorEnum) Reset()         { *m = WebpageConditionOperatorEnum{} }
func (m *WebpageConditionOperatorEnum) String() string { return proto.CompactTextString(m) }
func (*WebpageConditionOperatorEnum) ProtoMessage()    {}
func (*WebpageConditionOperatorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e856da14d9f47c05, []int{0}
}

func (m *WebpageConditionOperatorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebpageConditionOperatorEnum.Unmarshal(m, b)
}
func (m *WebpageConditionOperatorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebpageConditionOperatorEnum.Marshal(b, m, deterministic)
}
func (m *WebpageConditionOperatorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebpageConditionOperatorEnum.Merge(m, src)
}
func (m *WebpageConditionOperatorEnum) XXX_Size() int {
	return xxx_messageInfo_WebpageConditionOperatorEnum.Size(m)
}
func (m *WebpageConditionOperatorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_WebpageConditionOperatorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_WebpageConditionOperatorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.WebpageConditionOperatorEnum_WebpageConditionOperator", WebpageConditionOperatorEnum_WebpageConditionOperator_name, WebpageConditionOperatorEnum_WebpageConditionOperator_value)
	proto.RegisterType((*WebpageConditionOperatorEnum)(nil), "google.ads.googleads.v1.enums.WebpageConditionOperatorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/webpage_condition_operator.proto", fileDescriptor_e856da14d9f47c05)
}

var fileDescriptor_e856da14d9f47c05 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xd1, 0x6a, 0xc2, 0x30,
	0x14, 0x9d, 0x15, 0xdc, 0x88, 0x83, 0x95, 0x3e, 0x8d, 0xa1, 0x0f, 0xfa, 0x01, 0x29, 0x65, 0x6f,
	0x19, 0x0c, 0xa2, 0xeb, 0x44, 0x36, 0xa2, 0x9b, 0x53, 0x61, 0x14, 0x24, 0x9a, 0x10, 0x0a, 0x9a,
	0x5b, 0x9a, 0xea, 0xfe, 0x67, 0x8f, 0xfb, 0x94, 0x7d, 0xca, 0x1e, 0xf7, 0x05, 0xa3, 0x89, 0xf5,
	0xad, 0x7b, 0x09, 0x87, 0x9c, 0x7b, 0xce, 0x3d, 0xf7, 0xa0, 0x7b, 0x05, 0xa0, 0xb6, 0x32, 0xe4,
	0xc2, 0x84, 0x0e, 0x96, 0xe8, 0x10, 0x85, 0x52, 0xef, 0x77, 0x26, 0xfc, 0x90, 0xeb, 0x8c, 0x2b,
	0xb9, 0xda, 0x80, 0x16, 0x69, 0x91, 0x82, 0x5e, 0x41, 0x26, 0x73, 0x5e, 0x40, 0x8e, 0xb3, 0x1c,
	0x0a, 0x08, 0xba, 0x4e, 0x84, 0xb9, 0x30, 0xf8, 0xa4, 0xc7, 0x87, 0x08, 0x5b, 0xfd, 0x4d, 0xa7,
	0xb2, 0xcf, 0xd2, 0x90, 0x6b, 0x0d, 0x05, 0x2f, 0x4d, 0x8c, 0x13, 0xf7, 0x73, 0xd4, 0x59, 0xba,
	0x05, 0xc3, 0xca, 0x7f, 0x72, 0xb4, 0x8f, 0xf5, 0x7e, 0xd7, 0x7f, 0x45, 0xd7, 0x75, 0x7c, 0x70,
	0x85, 0xda, 0x73, 0x36, 0x9b, 0xc6, 0xc3, 0xf1, 0xe3, 0x38, 0x7e, 0xf0, 0xcf, 0x82, 0x36, 0x3a,
	0x9f, 0xb3, 0x27, 0x36, 0x59, 0x32, 0xbf, 0x11, 0x20, 0xd4, 0x8a, 0x5f, 0xe6, 0xf4, 0x79, 0xe6,
	0x7b, 0xc1, 0x25, 0xba, 0x18, 0x4e, 0xd8, 0x1b, 0x1d, 0xb3, 0x99, 0xdf, 0x1c, 0xfc, 0x36, 0x50,
	0x6f, 0x03, 0x3b, 0xfc, 0x6f, 0xee, 0x41, 0xb7, 0x6e, 0xef, 0xb4, 0x0c, 0x3e, 0x6d, 0xbc, 0x0f,
	0x8e, 0x7a, 0x05, 0x5b, 0xae, 0x15, 0x86, 0x5c, 0x85, 0x4a, 0x6a, 0x7b, 0x56, 0xd5, 0x63, 0x96,
	0x9a, 0x9a, 0x5a, 0xef, 0xec, 0xfb, 0xe9, 0x35, 0x47, 0x94, 0x7e, 0x79, 0xdd, 0x91, 0xb3, 0xa2,
	0xc2, 0x60, 0x07, 0x4b, 0xb4, 0x88, 0x70, 0xd9, 0x81, 0xf9, 0xae, 0xf8, 0x84, 0x0a, 0x93, 0x9c,
	0xf8, 0x64, 0x11, 0x25, 0x96, 0xff, 0xf1, 0x7a, 0xee, 0x93, 0x10, 0x2a, 0x0c, 0x21, 0xa7, 0x09,
	0x42, 0x16, 0x11, 0x21, 0x76, 0x66, 0xdd, 0xb2, 0xc1, 0x6e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xc7, 0x74, 0xb9, 0x52, 0xee, 0x01, 0x00, 0x00,
}
