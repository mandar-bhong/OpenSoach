// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/recommendation_error.proto

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

// Enum describing possible errors from applying a recommendation.
type RecommendationErrorEnum_RecommendationError int32

const (
	// Enum unspecified.
	RecommendationErrorEnum_UNSPECIFIED RecommendationErrorEnum_RecommendationError = 0
	// The received error code is not known in this version.
	RecommendationErrorEnum_UNKNOWN RecommendationErrorEnum_RecommendationError = 1
	// The specified budget amount is too low e.g. lower than minimum currency
	// unit or lower than ad group minimum cost-per-click.
	RecommendationErrorEnum_BUDGET_AMOUNT_TOO_SMALL RecommendationErrorEnum_RecommendationError = 2
	// The specified budget amount is too large.
	RecommendationErrorEnum_BUDGET_AMOUNT_TOO_LARGE RecommendationErrorEnum_RecommendationError = 3
	// The specified budget amount is not a valid amount. e.g. not a multiple
	// of minimum currency unit.
	RecommendationErrorEnum_INVALID_BUDGET_AMOUNT RecommendationErrorEnum_RecommendationError = 4
	// The specified keyword or ad violates ad policy.
	RecommendationErrorEnum_POLICY_ERROR RecommendationErrorEnum_RecommendationError = 5
	// The specified bid amount is not valid. e.g. too many fractional digits,
	// or negative amount.
	RecommendationErrorEnum_INVALID_BID_AMOUNT RecommendationErrorEnum_RecommendationError = 6
	// The number of keywords in ad group have reached the maximum allowed.
	RecommendationErrorEnum_ADGROUP_KEYWORD_LIMIT RecommendationErrorEnum_RecommendationError = 7
	// The recommendation requested to apply has already been applied.
	RecommendationErrorEnum_RECOMMENDATION_ALREADY_APPLIED RecommendationErrorEnum_RecommendationError = 8
	// The recommendation requested to apply has been invalidated.
	RecommendationErrorEnum_RECOMMENDATION_INVALIDATED RecommendationErrorEnum_RecommendationError = 9
	// The number of operations in a single request exceeds the maximum allowed.
	RecommendationErrorEnum_TOO_MANY_OPERATIONS RecommendationErrorEnum_RecommendationError = 10
	// There are no operations in the request.
	RecommendationErrorEnum_NO_OPERATIONS RecommendationErrorEnum_RecommendationError = 11
	// Operations with multiple recommendation types are not supported when
	// partial failure mode is not enabled.
	RecommendationErrorEnum_DIFFERENT_TYPES_NOT_SUPPORTED RecommendationErrorEnum_RecommendationError = 12
	// Request contains multiple operations with the same resource_name.
	RecommendationErrorEnum_DUPLICATE_RESOURCE_NAME RecommendationErrorEnum_RecommendationError = 13
	// The recommendation requested to dismiss has already been dismissed.
	RecommendationErrorEnum_RECOMMENDATION_ALREADY_DISMISSED RecommendationErrorEnum_RecommendationError = 14
	// The recommendation apply request was malformed and invalid.
	RecommendationErrorEnum_INVALID_APPLY_REQUEST RecommendationErrorEnum_RecommendationError = 15
)

var RecommendationErrorEnum_RecommendationError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "BUDGET_AMOUNT_TOO_SMALL",
	3:  "BUDGET_AMOUNT_TOO_LARGE",
	4:  "INVALID_BUDGET_AMOUNT",
	5:  "POLICY_ERROR",
	6:  "INVALID_BID_AMOUNT",
	7:  "ADGROUP_KEYWORD_LIMIT",
	8:  "RECOMMENDATION_ALREADY_APPLIED",
	9:  "RECOMMENDATION_INVALIDATED",
	10: "TOO_MANY_OPERATIONS",
	11: "NO_OPERATIONS",
	12: "DIFFERENT_TYPES_NOT_SUPPORTED",
	13: "DUPLICATE_RESOURCE_NAME",
	14: "RECOMMENDATION_ALREADY_DISMISSED",
	15: "INVALID_APPLY_REQUEST",
}

var RecommendationErrorEnum_RecommendationError_value = map[string]int32{
	"UNSPECIFIED":                      0,
	"UNKNOWN":                          1,
	"BUDGET_AMOUNT_TOO_SMALL":          2,
	"BUDGET_AMOUNT_TOO_LARGE":          3,
	"INVALID_BUDGET_AMOUNT":            4,
	"POLICY_ERROR":                     5,
	"INVALID_BID_AMOUNT":               6,
	"ADGROUP_KEYWORD_LIMIT":            7,
	"RECOMMENDATION_ALREADY_APPLIED":   8,
	"RECOMMENDATION_INVALIDATED":       9,
	"TOO_MANY_OPERATIONS":              10,
	"NO_OPERATIONS":                    11,
	"DIFFERENT_TYPES_NOT_SUPPORTED":    12,
	"DUPLICATE_RESOURCE_NAME":          13,
	"RECOMMENDATION_ALREADY_DISMISSED": 14,
	"INVALID_APPLY_REQUEST":            15,
}

func (x RecommendationErrorEnum_RecommendationError) String() string {
	return proto.EnumName(RecommendationErrorEnum_RecommendationError_name, int32(x))
}

func (RecommendationErrorEnum_RecommendationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_81944be64bae24a3, []int{0, 0}
}

// Container for enum describing possible errors from applying a recommendation.
type RecommendationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecommendationErrorEnum) Reset()         { *m = RecommendationErrorEnum{} }
func (m *RecommendationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*RecommendationErrorEnum) ProtoMessage()    {}
func (*RecommendationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_81944be64bae24a3, []int{0}
}

func (m *RecommendationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendationErrorEnum.Unmarshal(m, b)
}
func (m *RecommendationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendationErrorEnum.Marshal(b, m, deterministic)
}
func (m *RecommendationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendationErrorEnum.Merge(m, src)
}
func (m *RecommendationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_RecommendationErrorEnum.Size(m)
}
func (m *RecommendationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.RecommendationErrorEnum_RecommendationError", RecommendationErrorEnum_RecommendationError_name, RecommendationErrorEnum_RecommendationError_value)
	proto.RegisterType((*RecommendationErrorEnum)(nil), "google.ads.googleads.v1.errors.RecommendationErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/recommendation_error.proto", fileDescriptor_81944be64bae24a3)
}

var fileDescriptor_81944be64bae24a3 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x59, 0x0b, 0x1b, 0x78, 0x1b, 0x33, 0x9e, 0x60, 0x30, 0xa0, 0x82, 0x8a, 0xeb, 0x44,
	0x15, 0x57, 0x84, 0x2b, 0xb7, 0x3e, 0xad, 0xac, 0x25, 0xb6, 0x71, 0x92, 0x4e, 0x41, 0x95, 0xac,
	0xb2, 0x54, 0x51, 0xa5, 0x35, 0xa9, 0x92, 0xb2, 0x07, 0xe2, 0x92, 0x87, 0xe0, 0x01, 0x78, 0x04,
	0x1e, 0x81, 0x5b, 0x5e, 0x00, 0x39, 0x59, 0xab, 0x0d, 0x75, 0x5c, 0xe5, 0xe8, 0xf8, 0xfb, 0x7f,
	0x9f, 0xe3, 0xfc, 0xe8, 0x43, 0x56, 0x14, 0xd9, 0xe5, 0xcc, 0x9d, 0xa6, 0x95, 0xdb, 0x94, 0xb6,
	0xba, 0xea, 0xb9, 0xb3, 0xb2, 0x2c, 0xca, 0xca, 0x2d, 0x67, 0x17, 0xc5, 0x62, 0x31, 0xcb, 0xd3,
	0xe9, 0x6a, 0x5e, 0xe4, 0xa6, 0xee, 0x3a, 0xcb, 0xb2, 0x58, 0x15, 0xa4, 0xd3, 0xf0, 0xce, 0x34,
	0xad, 0x9c, 0x8d, 0xd4, 0xb9, 0xea, 0x39, 0x8d, 0xf4, 0xf4, 0xd5, 0xda, 0x7a, 0x39, 0x77, 0xa7,
	0x79, 0x5e, 0xac, 0x6a, 0x8b, 0xaa, 0x51, 0x77, 0x7f, 0xb5, 0xd1, 0x89, 0xbe, 0x65, 0x0e, 0x56,
	0x06, 0xf9, 0xd7, 0x45, 0xf7, 0x47, 0x1b, 0x1d, 0x6f, 0x39, 0x23, 0x47, 0x68, 0x3f, 0x16, 0xa1,
	0x82, 0x01, 0x1f, 0x72, 0x60, 0xf8, 0x1e, 0xd9, 0x47, 0x7b, 0xb1, 0x38, 0x13, 0xf2, 0x5c, 0xe0,
	0x1d, 0xf2, 0x12, 0x9d, 0xf4, 0x63, 0x36, 0x82, 0xc8, 0xd0, 0x40, 0xc6, 0x22, 0x32, 0x91, 0x94,
	0x26, 0x0c, 0xa8, 0xef, 0xe3, 0xd6, 0xf6, 0x43, 0x9f, 0xea, 0x11, 0xe0, 0x36, 0x79, 0x81, 0x9e,
	0x72, 0x31, 0xa6, 0x3e, 0x67, 0xe6, 0x16, 0x84, 0xef, 0x13, 0x8c, 0x0e, 0x94, 0xf4, 0xf9, 0x20,
	0x31, 0xa0, 0xb5, 0xd4, 0xf8, 0x01, 0x79, 0x86, 0xc8, 0x06, 0xe6, 0x6c, 0x4d, 0xee, 0x5a, 0x13,
	0xca, 0x46, 0x5a, 0xc6, 0xca, 0x9c, 0x41, 0x72, 0x2e, 0x35, 0x33, 0x3e, 0x0f, 0x78, 0x84, 0xf7,
	0x48, 0x17, 0x75, 0x34, 0x0c, 0x64, 0x10, 0x80, 0x60, 0x34, 0xe2, 0x52, 0x18, 0xea, 0x6b, 0xa0,
	0x2c, 0x31, 0x54, 0x29, 0xdf, 0xae, 0xf2, 0x90, 0x74, 0xd0, 0xe9, 0x3f, 0xcc, 0xf5, 0x2d, 0x34,
	0x02, 0x86, 0x1f, 0x91, 0x13, 0x74, 0x6c, 0x47, 0x0e, 0xa8, 0x48, 0x8c, 0x54, 0xa0, 0x6b, 0x26,
	0xc4, 0x88, 0x3c, 0x41, 0x87, 0x42, 0xde, 0x6c, 0xed, 0x93, 0xb7, 0xe8, 0x35, 0xe3, 0xc3, 0x21,
	0x68, 0xb0, 0x8b, 0x26, 0x0a, 0x42, 0x23, 0x64, 0x64, 0xc2, 0x58, 0x29, 0xa9, 0xad, 0xdd, 0x81,
	0x7d, 0x0f, 0x16, 0x2b, 0x9f, 0x0f, 0x68, 0x04, 0x46, 0x43, 0x28, 0x63, 0x3d, 0x00, 0x23, 0x68,
	0x00, 0xf8, 0x90, 0xbc, 0x43, 0x6f, 0xee, 0x98, 0x97, 0xf1, 0x30, 0xe0, 0x61, 0x08, 0x0c, 0x3f,
	0xbe, 0xf9, 0x6a, 0x76, 0x8d, 0xc4, 0x68, 0xf8, 0x14, 0x43, 0x18, 0xe1, 0xa3, 0xfe, 0x9f, 0x1d,
	0xd4, 0xbd, 0x28, 0x16, 0xce, 0xff, 0x13, 0xd2, 0x7f, 0xbe, 0xe5, 0x27, 0x2b, 0x9b, 0x0e, 0xb5,
	0xf3, 0x99, 0x5d, 0x6b, 0xb3, 0xe2, 0x72, 0x9a, 0x67, 0x4e, 0x51, 0x66, 0x6e, 0x36, 0xcb, 0xeb,
	0xec, 0xac, 0x83, 0xba, 0x9c, 0x57, 0x77, 0xe5, 0xf6, 0x63, 0xf3, 0xf9, 0xd6, 0x6a, 0x8f, 0x28,
	0xfd, 0xde, 0xea, 0x8c, 0x1a, 0x33, 0x9a, 0x56, 0x4e, 0x53, 0xda, 0x6a, 0xdc, 0x73, 0xea, 0x2b,
	0xab, 0x9f, 0x6b, 0x60, 0x42, 0xd3, 0x6a, 0xb2, 0x01, 0x26, 0xe3, 0xde, 0xa4, 0x01, 0x7e, 0xb7,
	0xba, 0x4d, 0xd7, 0xf3, 0x68, 0x5a, 0x79, 0xde, 0x06, 0xf1, 0xbc, 0x71, 0xcf, 0xf3, 0x1a, 0xe8,
	0xcb, 0x6e, 0x3d, 0xdd, 0xfb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xb2, 0x52, 0x7c, 0x54,
	0x03, 0x00, 0x00,
}
