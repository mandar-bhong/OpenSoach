// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/bidding_strategy_type.proto

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

// Enum describing possible bidding strategy types.
type BiddingStrategyTypeEnum_BiddingStrategyType int32

const (
	// Not specified.
	BiddingStrategyTypeEnum_UNSPECIFIED BiddingStrategyTypeEnum_BiddingStrategyType = 0
	// Used for return value only. Represents value unknown in this version.
	BiddingStrategyTypeEnum_UNKNOWN BiddingStrategyTypeEnum_BiddingStrategyType = 1
	// Enhanced CPC is a bidding strategy that raises bids for clicks
	// that seem more likely to lead to a conversion and lowers
	// them for clicks where they seem less likely.
	BiddingStrategyTypeEnum_ENHANCED_CPC BiddingStrategyTypeEnum_BiddingStrategyType = 2
	// Manual click based bidding where user pays per click.
	BiddingStrategyTypeEnum_MANUAL_CPC BiddingStrategyTypeEnum_BiddingStrategyType = 3
	// Manual impression based bidding
	// where user pays per thousand impressions.
	BiddingStrategyTypeEnum_MANUAL_CPM BiddingStrategyTypeEnum_BiddingStrategyType = 4
	// A bidding strategy that pays a configurable amount per video view.
	BiddingStrategyTypeEnum_MANUAL_CPV BiddingStrategyTypeEnum_BiddingStrategyType = 13
	// A bidding strategy that automatically maximizes number of conversions
	// given a daily budget.
	BiddingStrategyTypeEnum_MAXIMIZE_CONVERSIONS BiddingStrategyTypeEnum_BiddingStrategyType = 10
	// An automated bidding strategy that automatically sets bids to maximize
	// revenue while spending your budget.
	BiddingStrategyTypeEnum_MAXIMIZE_CONVERSION_VALUE BiddingStrategyTypeEnum_BiddingStrategyType = 11
	// Page-One Promoted bidding scheme, which sets max cpc bids to
	// target impressions on page one or page one promoted slots on google.com.
	BiddingStrategyTypeEnum_PAGE_ONE_PROMOTED BiddingStrategyTypeEnum_BiddingStrategyType = 5
	// Percent Cpc is bidding strategy where bids are a fraction of the
	// advertised price for some good or service.
	BiddingStrategyTypeEnum_PERCENT_CPC BiddingStrategyTypeEnum_BiddingStrategyType = 12
	// Target CPA is an automated bid strategy that sets bids
	// to help get as many conversions as possible
	// at the target cost-per-acquisition (CPA) you set.
	BiddingStrategyTypeEnum_TARGET_CPA BiddingStrategyTypeEnum_BiddingStrategyType = 6
	// Target CPM is an automated bid strategy that sets bids to help get
	// as many impressions as possible at the target cost per one thousand
	// impressions (CPM) you set.
	BiddingStrategyTypeEnum_TARGET_CPM BiddingStrategyTypeEnum_BiddingStrategyType = 14
	// An automated bidding strategy that sets bids so that a certain percentage
	// of search ads are shown at the top of the first page (or other targeted
	// location).
	BiddingStrategyTypeEnum_TARGET_IMPRESSION_SHARE BiddingStrategyTypeEnum_BiddingStrategyType = 15
	// Target Outrank Share is an automated bidding strategy that sets bids
	// based on the target fraction of auctions where the advertiser
	// should outrank a specific competitor.
	BiddingStrategyTypeEnum_TARGET_OUTRANK_SHARE BiddingStrategyTypeEnum_BiddingStrategyType = 7
	// Target ROAS is an automated bidding strategy
	// that helps you maximize revenue while averaging
	// a specific target Return On Average Spend (ROAS).
	BiddingStrategyTypeEnum_TARGET_ROAS BiddingStrategyTypeEnum_BiddingStrategyType = 8
	// Target Spend is an automated bid strategy that sets your bids
	// to help get as many clicks as possible within your budget.
	BiddingStrategyTypeEnum_TARGET_SPEND BiddingStrategyTypeEnum_BiddingStrategyType = 9
)

var BiddingStrategyTypeEnum_BiddingStrategyType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "ENHANCED_CPC",
	3:  "MANUAL_CPC",
	4:  "MANUAL_CPM",
	13: "MANUAL_CPV",
	10: "MAXIMIZE_CONVERSIONS",
	11: "MAXIMIZE_CONVERSION_VALUE",
	5:  "PAGE_ONE_PROMOTED",
	12: "PERCENT_CPC",
	6:  "TARGET_CPA",
	14: "TARGET_CPM",
	15: "TARGET_IMPRESSION_SHARE",
	7:  "TARGET_OUTRANK_SHARE",
	8:  "TARGET_ROAS",
	9:  "TARGET_SPEND",
}

var BiddingStrategyTypeEnum_BiddingStrategyType_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"ENHANCED_CPC":              2,
	"MANUAL_CPC":                3,
	"MANUAL_CPM":                4,
	"MANUAL_CPV":                13,
	"MAXIMIZE_CONVERSIONS":      10,
	"MAXIMIZE_CONVERSION_VALUE": 11,
	"PAGE_ONE_PROMOTED":         5,
	"PERCENT_CPC":               12,
	"TARGET_CPA":                6,
	"TARGET_CPM":                14,
	"TARGET_IMPRESSION_SHARE":   15,
	"TARGET_OUTRANK_SHARE":      7,
	"TARGET_ROAS":               8,
	"TARGET_SPEND":              9,
}

func (x BiddingStrategyTypeEnum_BiddingStrategyType) String() string {
	return proto.EnumName(BiddingStrategyTypeEnum_BiddingStrategyType_name, int32(x))
}

func (BiddingStrategyTypeEnum_BiddingStrategyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f0033534dbb333f5, []int{0, 0}
}

// Container for enum describing possible bidding strategy types.
type BiddingStrategyTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiddingStrategyTypeEnum) Reset()         { *m = BiddingStrategyTypeEnum{} }
func (m *BiddingStrategyTypeEnum) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategyTypeEnum) ProtoMessage()    {}
func (*BiddingStrategyTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0033534dbb333f5, []int{0}
}

func (m *BiddingStrategyTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategyTypeEnum.Unmarshal(m, b)
}
func (m *BiddingStrategyTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategyTypeEnum.Marshal(b, m, deterministic)
}
func (m *BiddingStrategyTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategyTypeEnum.Merge(m, src)
}
func (m *BiddingStrategyTypeEnum) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategyTypeEnum.Size(m)
}
func (m *BiddingStrategyTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategyTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategyTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.BiddingStrategyTypeEnum_BiddingStrategyType", BiddingStrategyTypeEnum_BiddingStrategyType_name, BiddingStrategyTypeEnum_BiddingStrategyType_value)
	proto.RegisterType((*BiddingStrategyTypeEnum)(nil), "google.ads.googleads.v1.enums.BiddingStrategyTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/bidding_strategy_type.proto", fileDescriptor_f0033534dbb333f5)
}

var fileDescriptor_f0033534dbb333f5 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdb, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x69, 0x06, 0x1b, 0xb8, 0x63, 0x33, 0x06, 0xb4, 0x71, 0xe8, 0xc5, 0xf6, 0x00, 0x8e,
	0x22, 0xae, 0x08, 0x57, 0x6e, 0x6a, 0xba, 0x68, 0x8b, 0x63, 0xe5, 0x04, 0x9a, 0x22, 0x45, 0x19,
	0x89, 0xa2, 0x48, 0x6b, 0x1c, 0xd5, 0xd9, 0xa4, 0xbe, 0x0e, 0x97, 0x3c, 0x0a, 0x57, 0x3c, 0x07,
	0x82, 0x77, 0x40, 0x76, 0xd2, 0x4a, 0x95, 0xca, 0x6e, 0xac, 0xef, 0xef, 0xdf, 0x77, 0x92, 0xff,
	0x06, 0x1f, 0x2b, 0x21, 0xaa, 0xdb, 0xd2, 0xcc, 0x0b, 0x69, 0xf6, 0xa1, 0x8a, 0xee, 0x2d, 0xb3,
	0x6c, 0xee, 0x16, 0xd2, 0xbc, 0xa9, 0x8b, 0xa2, 0x6e, 0xaa, 0x4c, 0x76, 0xcb, 0xbc, 0x2b, 0xab,
	0x55, 0xd6, 0xad, 0xda, 0x12, 0xb7, 0x4b, 0xd1, 0x09, 0x34, 0xe9, 0xf3, 0x71, 0x5e, 0x48, 0xbc,
	0x29, 0xc5, 0xf7, 0x16, 0xd6, 0xa5, 0x6f, 0xdf, 0xaf, 0x3b, 0xb7, 0xb5, 0x99, 0x37, 0x8d, 0xe8,
	0xf2, 0xae, 0x16, 0x8d, 0xec, 0x8b, 0xcf, 0xff, 0x18, 0xe0, 0x64, 0xda, 0x37, 0x0f, 0x87, 0xde,
	0xd1, 0xaa, 0x2d, 0x69, 0x73, 0xb7, 0x38, 0xff, 0x65, 0x80, 0x97, 0x3b, 0x18, 0x3a, 0x06, 0xe3,
	0x98, 0x85, 0x9c, 0x3a, 0xee, 0x67, 0x97, 0xce, 0xe0, 0x23, 0x34, 0x06, 0x07, 0x31, 0xbb, 0x64,
	0xfe, 0x17, 0x06, 0x47, 0x08, 0x82, 0x43, 0xca, 0x2e, 0x08, 0x73, 0xe8, 0x2c, 0x73, 0xb8, 0x03,
	0x0d, 0x74, 0x04, 0x80, 0x47, 0x58, 0x4c, 0xae, 0xb4, 0xde, 0xdb, 0xd2, 0x1e, 0x7c, 0xbc, 0xa5,
	0x13, 0xf8, 0x1c, 0x9d, 0x82, 0x57, 0x1e, 0xf9, 0xea, 0x7a, 0xee, 0x35, 0xcd, 0x1c, 0x9f, 0x25,
	0x34, 0x08, 0x5d, 0x9f, 0x85, 0x10, 0xa0, 0x09, 0x78, 0xb3, 0x83, 0x64, 0x09, 0xb9, 0x8a, 0x29,
	0x1c, 0xa3, 0xd7, 0xe0, 0x05, 0x27, 0x73, 0x9a, 0xf9, 0x8c, 0x66, 0x3c, 0xf0, 0x3d, 0x3f, 0xa2,
	0x33, 0xf8, 0x44, 0xed, 0xcb, 0x69, 0xe0, 0x50, 0x16, 0xe9, 0x05, 0x0e, 0xd5, 0xc0, 0x88, 0x04,
	0x73, 0xaa, 0x34, 0x81, 0xfb, 0x5b, 0xda, 0x83, 0x47, 0xe8, 0x1d, 0x38, 0x19, 0xb4, 0xeb, 0xf1,
	0x80, 0x86, 0x7a, 0x48, 0x78, 0x41, 0x02, 0x0a, 0x8f, 0xd5, 0x76, 0x03, 0xf4, 0xe3, 0x28, 0x20,
	0xec, 0x72, 0x20, 0x07, 0x6a, 0xce, 0x40, 0x02, 0x9f, 0x84, 0xf0, 0xa9, 0x7a, 0x8a, 0xe1, 0x22,
	0xe4, 0x94, 0xcd, 0xe0, 0xb3, 0xe9, 0xdf, 0x11, 0x38, 0xfb, 0x26, 0x16, 0xf8, 0x41, 0xcb, 0xa6,
	0xa7, 0x3b, 0x5e, 0x9d, 0x2b, 0xbb, 0xf8, 0xe8, 0x7a, 0x3a, 0x94, 0x56, 0xe2, 0x36, 0x6f, 0x2a,
	0x2c, 0x96, 0x95, 0x59, 0x95, 0x8d, 0x36, 0x73, 0xfd, 0x71, 0xda, 0x5a, 0xfe, 0xe7, 0x1f, 0x7d,
	0xd2, 0xe7, 0x77, 0x63, 0x6f, 0x4e, 0xc8, 0x0f, 0x63, 0x32, 0xef, 0x5b, 0x91, 0x42, 0xe2, 0x3e,
	0x54, 0x51, 0x62, 0x61, 0xe5, 0xbe, 0xfc, 0xb9, 0xe6, 0x29, 0x29, 0x64, 0xba, 0xe1, 0x69, 0x62,
	0xa5, 0x9a, 0xff, 0x36, 0xce, 0xfa, 0x4b, 0xdb, 0x26, 0x85, 0xb4, 0xed, 0x4d, 0x86, 0x6d, 0x27,
	0x96, 0x6d, 0xeb, 0x9c, 0x9b, 0x7d, 0xbd, 0xd8, 0x87, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdd,
	0x75, 0xc4, 0x3f, 0xdf, 0x02, 0x00, 0x00,
}
