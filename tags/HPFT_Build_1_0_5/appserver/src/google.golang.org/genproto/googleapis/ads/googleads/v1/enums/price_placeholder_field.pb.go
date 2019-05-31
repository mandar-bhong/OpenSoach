// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/price_placeholder_field.proto

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

// Possible values for Price placeholder fields.
type PricePlaceholderFieldEnum_PricePlaceholderField int32

const (
	// Not specified.
	PricePlaceholderFieldEnum_UNSPECIFIED PricePlaceholderFieldEnum_PricePlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	PricePlaceholderFieldEnum_UNKNOWN PricePlaceholderFieldEnum_PricePlaceholderField = 1
	// Data Type: STRING. The type of your price feed. Must match one of the
	// predefined price feed type exactly.
	PricePlaceholderFieldEnum_TYPE PricePlaceholderFieldEnum_PricePlaceholderField = 2
	// Data Type: STRING. The qualifier of each price. Must match one of the
	// predefined price qualifiers exactly.
	PricePlaceholderFieldEnum_PRICE_QUALIFIER PricePlaceholderFieldEnum_PricePlaceholderField = 3
	// Data Type: URL. Tracking template for the price feed when using Upgraded
	// URLs.
	PricePlaceholderFieldEnum_TRACKING_TEMPLATE PricePlaceholderFieldEnum_PricePlaceholderField = 4
	// Data Type: STRING. Language of the price feed. Must match one of the
	// available available locale codes exactly.
	PricePlaceholderFieldEnum_LANGUAGE PricePlaceholderFieldEnum_PricePlaceholderField = 5
	// Data Type: STRING. Final URL suffix for the price feed when using
	// parallel tracking.
	PricePlaceholderFieldEnum_FINAL_URL_SUFFIX PricePlaceholderFieldEnum_PricePlaceholderField = 6
	// Data Type: STRING. The header of item 1 of the table.
	PricePlaceholderFieldEnum_ITEM_1_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 100
	// Data Type: STRING. The description of item 1 of the table.
	PricePlaceholderFieldEnum_ITEM_1_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 101
	// Data Type: MONEY. The price (money with currency) of item 1 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_1_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 102
	// Data Type: STRING. The price unit of item 1 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_1_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 103
	// Data Type: URL_LIST. The final URLs of item 1 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_1_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 104
	// Data Type: URL_LIST. The final mobile URLs of item 1 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_1_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 105
	// Data Type: STRING. The header of item 2 of the table.
	PricePlaceholderFieldEnum_ITEM_2_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 200
	// Data Type: STRING. The description of item 2 of the table.
	PricePlaceholderFieldEnum_ITEM_2_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 201
	// Data Type: MONEY. The price (money with currency) of item 2 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_2_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 202
	// Data Type: STRING. The price unit of item 2 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_2_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 203
	// Data Type: URL_LIST. The final URLs of item 2 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_2_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 204
	// Data Type: URL_LIST. The final mobile URLs of item 2 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_2_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 205
	// Data Type: STRING. The header of item 3 of the table.
	PricePlaceholderFieldEnum_ITEM_3_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 300
	// Data Type: STRING. The description of item 3 of the table.
	PricePlaceholderFieldEnum_ITEM_3_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 301
	// Data Type: MONEY. The price (money with currency) of item 3 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_3_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 302
	// Data Type: STRING. The price unit of item 3 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_3_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 303
	// Data Type: URL_LIST. The final URLs of item 3 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_3_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 304
	// Data Type: URL_LIST. The final mobile URLs of item 3 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_3_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 305
	// Data Type: STRING. The header of item 4 of the table.
	PricePlaceholderFieldEnum_ITEM_4_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 400
	// Data Type: STRING. The description of item 4 of the table.
	PricePlaceholderFieldEnum_ITEM_4_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 401
	// Data Type: MONEY. The price (money with currency) of item 4 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_4_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 402
	// Data Type: STRING. The price unit of item 4 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_4_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 403
	// Data Type: URL_LIST. The final URLs of item 4 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_4_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 404
	// Data Type: URL_LIST. The final mobile URLs of item 4 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_4_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 405
	// Data Type: STRING. The header of item 5 of the table.
	PricePlaceholderFieldEnum_ITEM_5_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 500
	// Data Type: STRING. The description of item 5 of the table.
	PricePlaceholderFieldEnum_ITEM_5_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 501
	// Data Type: MONEY. The price (money with currency) of item 5 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_5_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 502
	// Data Type: STRING. The price unit of item 5 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_5_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 503
	// Data Type: URL_LIST. The final URLs of item 5 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_5_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 504
	// Data Type: URL_LIST. The final mobile URLs of item 5 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_5_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 505
	// Data Type: STRING. The header of item 6 of the table.
	PricePlaceholderFieldEnum_ITEM_6_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 600
	// Data Type: STRING. The description of item 6 of the table.
	PricePlaceholderFieldEnum_ITEM_6_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 601
	// Data Type: MONEY. The price (money with currency) of item 6 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_6_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 602
	// Data Type: STRING. The price unit of item 6 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_6_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 603
	// Data Type: URL_LIST. The final URLs of item 6 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_6_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 604
	// Data Type: URL_LIST. The final mobile URLs of item 6 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_6_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 605
	// Data Type: STRING. The header of item 7 of the table.
	PricePlaceholderFieldEnum_ITEM_7_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 700
	// Data Type: STRING. The description of item 7 of the table.
	PricePlaceholderFieldEnum_ITEM_7_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 701
	// Data Type: MONEY. The price (money with currency) of item 7 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_7_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 702
	// Data Type: STRING. The price unit of item 7 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_7_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 703
	// Data Type: URL_LIST. The final URLs of item 7 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_7_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 704
	// Data Type: URL_LIST. The final mobile URLs of item 7 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_7_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 705
	// Data Type: STRING. The header of item 8 of the table.
	PricePlaceholderFieldEnum_ITEM_8_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 800
	// Data Type: STRING. The description of item 8 of the table.
	PricePlaceholderFieldEnum_ITEM_8_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 801
	// Data Type: MONEY. The price (money with currency) of item 8 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_8_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 802
	// Data Type: STRING. The price unit of item 8 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_8_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 803
	// Data Type: URL_LIST. The final URLs of item 8 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_8_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 804
	// Data Type: URL_LIST. The final mobile URLs of item 8 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_8_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 805
)

var PricePlaceholderFieldEnum_PricePlaceholderField_name = map[int32]string{
	0:   "UNSPECIFIED",
	1:   "UNKNOWN",
	2:   "TYPE",
	3:   "PRICE_QUALIFIER",
	4:   "TRACKING_TEMPLATE",
	5:   "LANGUAGE",
	6:   "FINAL_URL_SUFFIX",
	100: "ITEM_1_HEADER",
	101: "ITEM_1_DESCRIPTION",
	102: "ITEM_1_PRICE",
	103: "ITEM_1_UNIT",
	104: "ITEM_1_FINAL_URLS",
	105: "ITEM_1_FINAL_MOBILE_URLS",
	200: "ITEM_2_HEADER",
	201: "ITEM_2_DESCRIPTION",
	202: "ITEM_2_PRICE",
	203: "ITEM_2_UNIT",
	204: "ITEM_2_FINAL_URLS",
	205: "ITEM_2_FINAL_MOBILE_URLS",
	300: "ITEM_3_HEADER",
	301: "ITEM_3_DESCRIPTION",
	302: "ITEM_3_PRICE",
	303: "ITEM_3_UNIT",
	304: "ITEM_3_FINAL_URLS",
	305: "ITEM_3_FINAL_MOBILE_URLS",
	400: "ITEM_4_HEADER",
	401: "ITEM_4_DESCRIPTION",
	402: "ITEM_4_PRICE",
	403: "ITEM_4_UNIT",
	404: "ITEM_4_FINAL_URLS",
	405: "ITEM_4_FINAL_MOBILE_URLS",
	500: "ITEM_5_HEADER",
	501: "ITEM_5_DESCRIPTION",
	502: "ITEM_5_PRICE",
	503: "ITEM_5_UNIT",
	504: "ITEM_5_FINAL_URLS",
	505: "ITEM_5_FINAL_MOBILE_URLS",
	600: "ITEM_6_HEADER",
	601: "ITEM_6_DESCRIPTION",
	602: "ITEM_6_PRICE",
	603: "ITEM_6_UNIT",
	604: "ITEM_6_FINAL_URLS",
	605: "ITEM_6_FINAL_MOBILE_URLS",
	700: "ITEM_7_HEADER",
	701: "ITEM_7_DESCRIPTION",
	702: "ITEM_7_PRICE",
	703: "ITEM_7_UNIT",
	704: "ITEM_7_FINAL_URLS",
	705: "ITEM_7_FINAL_MOBILE_URLS",
	800: "ITEM_8_HEADER",
	801: "ITEM_8_DESCRIPTION",
	802: "ITEM_8_PRICE",
	803: "ITEM_8_UNIT",
	804: "ITEM_8_FINAL_URLS",
	805: "ITEM_8_FINAL_MOBILE_URLS",
}

var PricePlaceholderFieldEnum_PricePlaceholderField_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"TYPE":                     2,
	"PRICE_QUALIFIER":          3,
	"TRACKING_TEMPLATE":        4,
	"LANGUAGE":                 5,
	"FINAL_URL_SUFFIX":         6,
	"ITEM_1_HEADER":            100,
	"ITEM_1_DESCRIPTION":       101,
	"ITEM_1_PRICE":             102,
	"ITEM_1_UNIT":              103,
	"ITEM_1_FINAL_URLS":        104,
	"ITEM_1_FINAL_MOBILE_URLS": 105,
	"ITEM_2_HEADER":            200,
	"ITEM_2_DESCRIPTION":       201,
	"ITEM_2_PRICE":             202,
	"ITEM_2_UNIT":              203,
	"ITEM_2_FINAL_URLS":        204,
	"ITEM_2_FINAL_MOBILE_URLS": 205,
	"ITEM_3_HEADER":            300,
	"ITEM_3_DESCRIPTION":       301,
	"ITEM_3_PRICE":             302,
	"ITEM_3_UNIT":              303,
	"ITEM_3_FINAL_URLS":        304,
	"ITEM_3_FINAL_MOBILE_URLS": 305,
	"ITEM_4_HEADER":            400,
	"ITEM_4_DESCRIPTION":       401,
	"ITEM_4_PRICE":             402,
	"ITEM_4_UNIT":              403,
	"ITEM_4_FINAL_URLS":        404,
	"ITEM_4_FINAL_MOBILE_URLS": 405,
	"ITEM_5_HEADER":            500,
	"ITEM_5_DESCRIPTION":       501,
	"ITEM_5_PRICE":             502,
	"ITEM_5_UNIT":              503,
	"ITEM_5_FINAL_URLS":        504,
	"ITEM_5_FINAL_MOBILE_URLS": 505,
	"ITEM_6_HEADER":            600,
	"ITEM_6_DESCRIPTION":       601,
	"ITEM_6_PRICE":             602,
	"ITEM_6_UNIT":              603,
	"ITEM_6_FINAL_URLS":        604,
	"ITEM_6_FINAL_MOBILE_URLS": 605,
	"ITEM_7_HEADER":            700,
	"ITEM_7_DESCRIPTION":       701,
	"ITEM_7_PRICE":             702,
	"ITEM_7_UNIT":              703,
	"ITEM_7_FINAL_URLS":        704,
	"ITEM_7_FINAL_MOBILE_URLS": 705,
	"ITEM_8_HEADER":            800,
	"ITEM_8_DESCRIPTION":       801,
	"ITEM_8_PRICE":             802,
	"ITEM_8_UNIT":              803,
	"ITEM_8_FINAL_URLS":        804,
	"ITEM_8_FINAL_MOBILE_URLS": 805,
}

func (x PricePlaceholderFieldEnum_PricePlaceholderField) String() string {
	return proto.EnumName(PricePlaceholderFieldEnum_PricePlaceholderField_name, int32(x))
}

func (PricePlaceholderFieldEnum_PricePlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3d990bd7464db1dd, []int{0, 0}
}

// Values for Price placeholder fields.
type PricePlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PricePlaceholderFieldEnum) Reset()         { *m = PricePlaceholderFieldEnum{} }
func (m *PricePlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*PricePlaceholderFieldEnum) ProtoMessage()    {}
func (*PricePlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d990bd7464db1dd, []int{0}
}

func (m *PricePlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PricePlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *PricePlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PricePlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *PricePlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PricePlaceholderFieldEnum.Merge(m, src)
}
func (m *PricePlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_PricePlaceholderFieldEnum.Size(m)
}
func (m *PricePlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PricePlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PricePlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.PricePlaceholderFieldEnum_PricePlaceholderField", PricePlaceholderFieldEnum_PricePlaceholderField_name, PricePlaceholderFieldEnum_PricePlaceholderField_value)
	proto.RegisterType((*PricePlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.PricePlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/price_placeholder_field.proto", fileDescriptor_3d990bd7464db1dd)
}

var fileDescriptor_3d990bd7464db1dd = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd4, 0x49, 0x4f, 0xdb, 0x5a,
	0x14, 0x07, 0xf0, 0xe7, 0x21, 0x7e, 0x70, 0xe1, 0x89, 0xcb, 0x7d, 0x0f, 0x5e, 0x8b, 0x60, 0x01,
	0x1f, 0xc0, 0x51, 0xe6, 0xc8, 0xac, 0x9c, 0xe0, 0xa4, 0x16, 0xc1, 0xb8, 0x19, 0xe8, 0xa0, 0x48,
	0x56, 0x8a, 0x4d, 0x88, 0x14, 0xe2, 0x28, 0x06, 0x3e, 0x47, 0xc7, 0x7d, 0x4b, 0xdb, 0x4d, 0xd5,
	0x71, 0xdf, 0x71, 0xd7, 0xd2, 0x76, 0xd1, 0x5d, 0xc7, 0x2f, 0xd0, 0x0d, 0x9b, 0x8e, 0xbb, 0xca,
	0xf1, 0xf5, 0x89, 0x2d, 0xb9, 0xdd, 0x58, 0x47, 0xe7, 0xdc, 0x73, 0xfc, 0x5b, 0xfd, 0xd1, 0x72,
	0xdb, 0xb6, 0xdb, 0x5d, 0x2b, 0xde, 0x32, 0x9d, 0xb8, 0x57, 0xba, 0xd5, 0x7e, 0x22, 0x6e, 0xf5,
	0xf6, 0x76, 0x9c, 0x78, 0x7f, 0xd0, 0xd9, 0xb4, 0x8c, 0x7e, 0xb7, 0xb5, 0x69, 0x6d, 0xdb, 0x5d,
	0xd3, 0x1a, 0x18, 0x5b, 0x1d, 0xab, 0x6b, 0x8a, 0xfd, 0x81, 0xbd, 0x6b, 0x93, 0x05, 0x6f, 0x43,
	0x6c, 0x99, 0x8e, 0x08, 0xcb, 0xe2, 0x7e, 0x42, 0x1c, 0x2e, 0xcf, 0xcd, 0xfb, 0xb7, 0xfb, 0x9d,
	0x78, 0xab, 0xd7, 0xb3, 0x77, 0x5b, 0xbb, 0x1d, 0xbb, 0xe7, 0x78, 0xcb, 0x4b, 0x47, 0xe3, 0xe8,
	0xb8, 0xee, 0x9e, 0xd7, 0x47, 0xd7, 0x4b, 0xee, 0x71, 0xa5, 0xb7, 0xb7, 0xb3, 0xf4, 0x66, 0x1c,
	0xcd, 0x44, 0x4e, 0xc9, 0x14, 0x9a, 0x68, 0x68, 0x35, 0x5d, 0x29, 0xaa, 0x25, 0x55, 0x59, 0xc1,
	0x7f, 0x91, 0x09, 0xf4, 0x77, 0x43, 0x5b, 0xd5, 0xd6, 0x4f, 0x69, 0x98, 0x21, 0x63, 0x88, 0xaf,
	0x9f, 0xd1, 0x15, 0xcc, 0x92, 0x7f, 0xd1, 0x94, 0x5e, 0x55, 0x8b, 0x8a, 0x71, 0xb2, 0x21, 0x57,
	0xdc, 0xb7, 0x55, 0xcc, 0x91, 0x19, 0x34, 0x5d, 0xaf, 0xca, 0xc5, 0x55, 0x55, 0x2b, 0x1b, 0x75,
	0x65, 0x4d, 0xaf, 0xc8, 0x75, 0x05, 0xf3, 0x64, 0x12, 0x8d, 0x55, 0x64, 0xad, 0xdc, 0x90, 0xcb,
	0x0a, 0x8e, 0x91, 0xff, 0x10, 0x2e, 0xa9, 0x9a, 0x5c, 0x31, 0x1a, 0xd5, 0x8a, 0x51, 0x6b, 0x94,
	0x4a, 0xea, 0x69, 0x2c, 0x90, 0x69, 0xf4, 0x8f, 0x5a, 0x57, 0xd6, 0x8c, 0x84, 0x71, 0x42, 0x91,
	0x57, 0x94, 0x2a, 0x36, 0xc9, 0x2c, 0x22, 0xb4, 0xb5, 0xa2, 0xd4, 0x8a, 0x55, 0x55, 0xaf, 0xab,
	0xeb, 0x1a, 0xb6, 0x08, 0x46, 0x93, 0xb4, 0x3f, 0x14, 0xe0, 0x2d, 0x17, 0x4d, 0x3b, 0x0d, 0x4d,
	0xad, 0xe3, 0xb6, 0x0b, 0xa1, 0x0d, 0xf8, 0x55, 0x0d, 0x6f, 0x93, 0x79, 0x74, 0x2c, 0xd4, 0x5e,
	0x5b, 0x2f, 0xa8, 0x15, 0xc5, 0x9b, 0x76, 0x08, 0xa1, 0x84, 0xa4, 0x4f, 0x78, 0xce, 0x90, 0xff,
	0xa9, 0x21, 0x19, 0x32, 0xbc, 0x60, 0xc8, 0x34, 0x45, 0x24, 0x29, 0xe2, 0x90, 0x21, 0x98, 0x2a,
	0x92, 0x9e, 0xe2, 0x25, 0x43, 0x66, 0x29, 0x23, 0x19, 0x64, 0xbc, 0x62, 0xc8, 0x02, 0x75, 0x24,
	0x23, 0x1c, 0xaf, 0x19, 0x80, 0xa4, 0x7c, 0xc8, 0x6d, 0x16, 0x20, 0xa9, 0x10, 0xe4, 0x0e, 0x0b,
	0x90, 0x14, 0x85, 0xdc, 0x65, 0x01, 0x92, 0xf2, 0x20, 0xf7, 0x58, 0x80, 0xa4, 0x82, 0x90, 0xfb,
	0x2c, 0x40, 0x52, 0x11, 0x90, 0x07, 0x2c, 0x40, 0xd2, 0x3e, 0xe4, 0x3c, 0x07, 0x90, 0x74, 0x08,
	0x72, 0x81, 0x03, 0x48, 0x9a, 0x42, 0x2e, 0x72, 0x00, 0x49, 0x7b, 0x90, 0x4b, 0x1c, 0x40, 0xd2,
	0x41, 0xc8, 0x65, 0x0e, 0x20, 0xe9, 0x08, 0xc8, 0x15, 0x0e, 0x20, 0x19, 0x1f, 0xf2, 0x65, 0x04,
	0xc9, 0x84, 0x20, 0x5f, 0x47, 0x90, 0x0c, 0x85, 0x7c, 0x1b, 0x41, 0x32, 0x1e, 0xe4, 0xfb, 0x08,
	0x92, 0x09, 0x42, 0x7e, 0x8c, 0x20, 0x99, 0x08, 0xc8, 0xcf, 0x11, 0x24, 0xeb, 0x43, 0xde, 0xf2,
	0x00, 0xc9, 0x86, 0x20, 0xef, 0x78, 0x80, 0x64, 0x29, 0xe4, 0x3d, 0x0f, 0x90, 0xac, 0x07, 0xf9,
	0xc0, 0x03, 0x24, 0x1b, 0x84, 0x7c, 0xe4, 0x01, 0x92, 0x8d, 0x80, 0x7c, 0xe2, 0x01, 0x92, 0xf3,
	0x21, 0x0f, 0x63, 0x00, 0xc9, 0x85, 0x20, 0x8f, 0x62, 0x00, 0xc9, 0x51, 0xc8, 0xe3, 0x18, 0x40,
	0x72, 0x1e, 0xe4, 0x49, 0x0c, 0x20, 0xb9, 0x20, 0xe4, 0x69, 0x0c, 0x20, 0xb9, 0x08, 0xc8, 0xb3,
	0x18, 0x40, 0xf2, 0x3e, 0xe4, 0xaa, 0x00, 0x90, 0x7c, 0x08, 0x72, 0x4d, 0x00, 0x48, 0x9e, 0x42,
	0x0e, 0x04, 0x80, 0xe4, 0x3d, 0xc8, 0x75, 0x01, 0x20, 0xf9, 0x20, 0xe4, 0x86, 0x00, 0x90, 0x7c,
	0x04, 0xe4, 0xa6, 0x50, 0x38, 0x62, 0xd0, 0xe2, 0xa6, 0xbd, 0x23, 0xfe, 0x31, 0x35, 0x0b, 0x73,
	0x91, 0xb1, 0xa7, 0xbb, 0x99, 0xa9, 0x33, 0x67, 0x0b, 0x74, 0xb9, 0x6d, 0x77, 0x5b, 0xbd, 0xb6,
	0x68, 0x0f, 0xda, 0xf1, 0xb6, 0xd5, 0x1b, 0x26, 0xaa, 0x9f, 0xdf, 0xfd, 0x8e, 0xf3, 0x9b, 0x38,
	0x5f, 0x1e, 0x7e, 0x0f, 0x58, 0xae, 0x2c, 0xcb, 0xb7, 0xd8, 0x85, 0xb2, 0x77, 0x4a, 0x36, 0x1d,
	0xd1, 0x2b, 0xdd, 0x6a, 0x23, 0x21, 0xba, 0x01, 0xec, 0x1c, 0xfa, 0xf3, 0xa6, 0x6c, 0x3a, 0x4d,
	0x98, 0x37, 0x37, 0x12, 0xcd, 0xe1, 0xfc, 0x33, 0xbb, 0xe8, 0x35, 0x25, 0x49, 0x36, 0x1d, 0x49,
	0x82, 0x17, 0x92, 0xb4, 0x91, 0x90, 0xa4, 0xe1, 0x9b, 0x73, 0xc2, 0x10, 0x96, 0xfa, 0x15, 0x00,
	0x00, 0xff, 0xff, 0xce, 0x0b, 0xb8, 0x5a, 0x66, 0x06, 0x00, 0x00,
}
