// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/feed_item_target_error.proto

package errors

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

// Enum describing possible feed item target errors.
type FeedItemTargetErrorEnum_FeedItemTargetError int32

const (
	// Enum unspecified.
	FeedItemTargetErrorEnum_UNSPECIFIED FeedItemTargetErrorEnum_FeedItemTargetError = 0
	// The received error code is not known in this version.
	FeedItemTargetErrorEnum_UNKNOWN FeedItemTargetErrorEnum_FeedItemTargetError = 1
	// On CREATE, the FeedItemTarget must have a populated field in the oneof
	// target.
	FeedItemTargetErrorEnum_MUST_SET_TARGET_ONEOF_ON_CREATE FeedItemTargetErrorEnum_FeedItemTargetError = 2
	// The specified feed item target already exists, so it cannot be added.
	FeedItemTargetErrorEnum_FEED_ITEM_TARGET_ALREADY_EXISTS FeedItemTargetErrorEnum_FeedItemTargetError = 3
	// The schedules for a given feed item cannot overlap.
	FeedItemTargetErrorEnum_FEED_ITEM_SCHEDULES_CANNOT_OVERLAP FeedItemTargetErrorEnum_FeedItemTargetError = 4
	// Too many targets of a given type were added for a single feed item.
	FeedItemTargetErrorEnum_TARGET_LIMIT_EXCEEDED_FOR_GIVEN_TYPE FeedItemTargetErrorEnum_FeedItemTargetError = 5
	// Too many AdSchedules are enabled for the feed item for the given day.
	FeedItemTargetErrorEnum_TOO_MANY_SCHEDULES_PER_DAY FeedItemTargetErrorEnum_FeedItemTargetError = 6
	// A feed item may either have an enabled campaign target or an enabled ad
	// group target.
	FeedItemTargetErrorEnum_CANNOT_HAVE_ENABLED_CAMPAIGN_AND_ENABLED_AD_GROUP_TARGETS FeedItemTargetErrorEnum_FeedItemTargetError = 7
)

var FeedItemTargetErrorEnum_FeedItemTargetError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MUST_SET_TARGET_ONEOF_ON_CREATE",
	3: "FEED_ITEM_TARGET_ALREADY_EXISTS",
	4: "FEED_ITEM_SCHEDULES_CANNOT_OVERLAP",
	5: "TARGET_LIMIT_EXCEEDED_FOR_GIVEN_TYPE",
	6: "TOO_MANY_SCHEDULES_PER_DAY",
	7: "CANNOT_HAVE_ENABLED_CAMPAIGN_AND_ENABLED_AD_GROUP_TARGETS",
}

var FeedItemTargetErrorEnum_FeedItemTargetError_value = map[string]int32{
	"UNSPECIFIED":                                               0,
	"UNKNOWN":                                                   1,
	"MUST_SET_TARGET_ONEOF_ON_CREATE":                           2,
	"FEED_ITEM_TARGET_ALREADY_EXISTS":                           3,
	"FEED_ITEM_SCHEDULES_CANNOT_OVERLAP":                        4,
	"TARGET_LIMIT_EXCEEDED_FOR_GIVEN_TYPE":                      5,
	"TOO_MANY_SCHEDULES_PER_DAY":                                6,
	"CANNOT_HAVE_ENABLED_CAMPAIGN_AND_ENABLED_AD_GROUP_TARGETS": 7,
}

func (x FeedItemTargetErrorEnum_FeedItemTargetError) String() string {
	return proto.EnumName(FeedItemTargetErrorEnum_FeedItemTargetError_name, int32(x))
}

func (FeedItemTargetErrorEnum_FeedItemTargetError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e158dc4de886fe3b, []int{0, 0}
}

// Container for enum describing possible feed item target errors.
type FeedItemTargetErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItemTargetErrorEnum) Reset()         { *m = FeedItemTargetErrorEnum{} }
func (m *FeedItemTargetErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FeedItemTargetErrorEnum) ProtoMessage()    {}
func (*FeedItemTargetErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e158dc4de886fe3b, []int{0}
}

func (m *FeedItemTargetErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemTargetErrorEnum.Unmarshal(m, b)
}
func (m *FeedItemTargetErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemTargetErrorEnum.Marshal(b, m, deterministic)
}
func (m *FeedItemTargetErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemTargetErrorEnum.Merge(m, src)
}
func (m *FeedItemTargetErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FeedItemTargetErrorEnum.Size(m)
}
func (m *FeedItemTargetErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemTargetErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemTargetErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.FeedItemTargetErrorEnum_FeedItemTargetError", FeedItemTargetErrorEnum_FeedItemTargetError_name, FeedItemTargetErrorEnum_FeedItemTargetError_value)
	proto.RegisterType((*FeedItemTargetErrorEnum)(nil), "google.ads.googleads.v1.errors.FeedItemTargetErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/feed_item_target_error.proto", fileDescriptor_e158dc4de886fe3b)
}

var fileDescriptor_e158dc4de886fe3b = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8e, 0xd3, 0x3e,
	0x10, 0xc6, 0xff, 0xcd, 0xfe, 0xd9, 0x95, 0xbc, 0x07, 0xa2, 0x70, 0x00, 0xad, 0x50, 0x91, 0x0a,
	0x42, 0x9c, 0x12, 0x55, 0x9c, 0xc8, 0x8a, 0x83, 0x1b, 0x4f, 0xb3, 0x11, 0xad, 0x1d, 0x25, 0x4e,
	0xd8, 0xa2, 0x4a, 0xa3, 0x40, 0x4c, 0x54, 0x69, 0x1b, 0x57, 0x49, 0xd8, 0x07, 0xe2, 0xc8, 0x81,
	0x07, 0xd9, 0x47, 0xe1, 0xca, 0x0b, 0xa0, 0xc4, 0x6d, 0xe1, 0xb0, 0x70, 0xf2, 0xa7, 0xf1, 0xef,
	0xfb, 0x3c, 0xf2, 0x0c, 0xb9, 0xac, 0xb4, 0xae, 0x6e, 0x94, 0x57, 0x94, 0xad, 0x67, 0x64, 0xaf,
	0x6e, 0xa7, 0x9e, 0x6a, 0x1a, 0xdd, 0xb4, 0xde, 0x67, 0xa5, 0x4a, 0xdc, 0x74, 0x6a, 0x8b, 0x5d,
	0xd1, 0x54, 0xaa, 0xc3, 0xa1, 0xee, 0xee, 0x1a, 0xdd, 0x69, 0x67, 0x6c, 0x1c, 0x6e, 0x51, 0xb6,
	0xee, 0xd1, 0xec, 0xde, 0x4e, 0x5d, 0x63, 0xbe, 0x78, 0x7a, 0x08, 0xdf, 0x6d, 0xbc, 0xa2, 0xae,
	0x75, 0x57, 0x74, 0x1b, 0x5d, 0xb7, 0xc6, 0x3d, 0xb9, 0xb3, 0xc8, 0xe3, 0xb9, 0x52, 0x65, 0xd4,
	0xa9, 0xad, 0x1c, 0xc2, 0xa1, 0xb7, 0x41, 0xfd, 0x65, 0x3b, 0xf9, 0x6e, 0x91, 0x47, 0xf7, 0xdc,
	0x39, 0x0f, 0xc9, 0x79, 0xc6, 0xd3, 0x18, 0x82, 0x68, 0x1e, 0x01, 0xb3, 0xff, 0x73, 0xce, 0xc9,
	0x59, 0xc6, 0xdf, 0x71, 0xf1, 0x9e, 0xdb, 0x23, 0xe7, 0x39, 0x79, 0xb6, 0xcc, 0x52, 0x89, 0x29,
	0x48, 0x94, 0x34, 0x09, 0x41, 0xa2, 0xe0, 0x20, 0xe6, 0x28, 0x38, 0x06, 0x09, 0x50, 0x09, 0xb6,
	0xd5, 0x43, 0x73, 0x00, 0x86, 0x91, 0x84, 0xe5, 0x81, 0xa2, 0x8b, 0x04, 0x28, 0x5b, 0x21, 0x5c,
	0x47, 0xa9, 0x4c, 0xed, 0x13, 0xe7, 0x25, 0x99, 0xfc, 0x86, 0xd2, 0xe0, 0x0a, 0x58, 0xb6, 0x80,
	0x14, 0x03, 0xca, 0xb9, 0x90, 0x28, 0x72, 0x48, 0x16, 0x34, 0xb6, 0xff, 0x77, 0x5e, 0x91, 0x17,
	0xfb, 0x88, 0x45, 0xb4, 0x8c, 0x24, 0xc2, 0x75, 0x00, 0xc0, 0x80, 0xe1, 0x5c, 0x24, 0x18, 0x46,
	0x39, 0x70, 0x94, 0xab, 0x18, 0xec, 0x07, 0xce, 0x98, 0x5c, 0x48, 0x21, 0x70, 0x49, 0xf9, 0xea,
	0x8f, 0xc0, 0x18, 0x12, 0x64, 0x74, 0x65, 0x9f, 0x3a, 0x6f, 0xc9, 0x9b, 0x7d, 0xfa, 0x15, 0xcd,
	0x01, 0x81, 0xd3, 0xd9, 0x02, 0x18, 0x06, 0x74, 0x19, 0xd3, 0x28, 0xe4, 0x48, 0x39, 0x3b, 0x16,
	0x29, 0xc3, 0x30, 0x11, 0x59, 0xbc, 0x6f, 0x3f, 0xb5, 0xcf, 0x66, 0x3f, 0x47, 0x64, 0xf2, 0x49,
	0x6f, 0xdd, 0x7f, 0x4f, 0x64, 0xf6, 0xe4, 0x9e, 0x4f, 0x8d, 0xfb, 0x69, 0xc4, 0xa3, 0x0f, 0x6c,
	0xef, 0xad, 0xf4, 0x4d, 0x51, 0x57, 0xae, 0x6e, 0x2a, 0xaf, 0x52, 0xf5, 0x30, 0xab, 0xc3, 0x6a,
	0xec, 0x36, 0xed, 0xdf, 0x36, 0xe5, 0xd2, 0x1c, 0x5f, 0xad, 0x93, 0x90, 0xd2, 0x6f, 0xd6, 0x38,
	0x34, 0x61, 0xb4, 0x6c, 0x5d, 0x23, 0x7b, 0x95, 0x4f, 0xdd, 0xe1, 0xc9, 0xf6, 0xee, 0x00, 0xac,
	0x69, 0xd9, 0xae, 0x8f, 0xc0, 0x3a, 0x9f, 0xae, 0x0d, 0xf0, 0xc3, 0x9a, 0x98, 0xaa, 0xef, 0xd3,
	0xb2, 0xf5, 0xfd, 0x23, 0xe2, 0xfb, 0xf9, 0xd4, 0xf7, 0x0d, 0xf4, 0xf1, 0x74, 0xe8, 0xee, 0xf5,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x09, 0xce, 0x26, 0xc6, 0x02, 0x00, 0x00,
}
