// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/search_engine_results_page_type.proto

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

// The type of the search engine results page.
type SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType int32

const (
	// Not specified.
	SearchEngineResultsPageTypeEnum_UNSPECIFIED SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType = 0
	// Used for return value only. Represents value unknown in this version.
	SearchEngineResultsPageTypeEnum_UNKNOWN SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType = 1
	// Only ads were contained in the search engine results page.
	SearchEngineResultsPageTypeEnum_ADS_ONLY SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType = 2
	// Only organic results were contained in the search engine results page.
	SearchEngineResultsPageTypeEnum_ORGANIC_ONLY SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType = 3
	// Both ads and organic results were contained in the search engine results
	// page.
	SearchEngineResultsPageTypeEnum_ADS_AND_ORGANIC SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType = 4
)

var SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ADS_ONLY",
	3: "ORGANIC_ONLY",
	4: "ADS_AND_ORGANIC",
}

var SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"ADS_ONLY":        2,
	"ORGANIC_ONLY":    3,
	"ADS_AND_ORGANIC": 4,
}

func (x SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType) String() string {
	return proto.EnumName(SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType_name, int32(x))
}

func (SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db318dd200586d33, []int{0, 0}
}

// The type of the search engine results page.
type SearchEngineResultsPageTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchEngineResultsPageTypeEnum) Reset()         { *m = SearchEngineResultsPageTypeEnum{} }
func (m *SearchEngineResultsPageTypeEnum) String() string { return proto.CompactTextString(m) }
func (*SearchEngineResultsPageTypeEnum) ProtoMessage()    {}
func (*SearchEngineResultsPageTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_db318dd200586d33, []int{0}
}

func (m *SearchEngineResultsPageTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchEngineResultsPageTypeEnum.Unmarshal(m, b)
}
func (m *SearchEngineResultsPageTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchEngineResultsPageTypeEnum.Marshal(b, m, deterministic)
}
func (m *SearchEngineResultsPageTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchEngineResultsPageTypeEnum.Merge(m, src)
}
func (m *SearchEngineResultsPageTypeEnum) XXX_Size() int {
	return xxx_messageInfo_SearchEngineResultsPageTypeEnum.Size(m)
}
func (m *SearchEngineResultsPageTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchEngineResultsPageTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SearchEngineResultsPageTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType", SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType_name, SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType_value)
	proto.RegisterType((*SearchEngineResultsPageTypeEnum)(nil), "google.ads.googleads.v1.enums.SearchEngineResultsPageTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/search_engine_results_page_type.proto", fileDescriptor_db318dd200586d33)
}

var fileDescriptor_db318dd200586d33 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xd1, 0x4e, 0xc2, 0x30,
	0x14, 0x75, 0xc3, 0xa8, 0x29, 0x24, 0x2c, 0xf3, 0x4d, 0x25, 0x0a, 0x1f, 0xd0, 0x65, 0xf1, 0xad,
	0x3e, 0x15, 0x98, 0x84, 0x68, 0xca, 0x02, 0x82, 0xd1, 0x2c, 0x59, 0x2a, 0x6b, 0xea, 0x12, 0x68,
	0x9b, 0x75, 0x90, 0xf0, 0x1d, 0xfe, 0x81, 0x8f, 0x7e, 0x8a, 0x9f, 0xe2, 0xbb, 0xef, 0x66, 0x2d,
	0xf0, 0x26, 0x2f, 0xcd, 0xc9, 0x3d, 0xf7, 0x9e, 0x73, 0xcf, 0x2d, 0xe8, 0x71, 0x29, 0xf9, 0x82,
	0x05, 0x34, 0xd3, 0x81, 0x85, 0x15, 0x5a, 0x87, 0x01, 0x13, 0xab, 0xa5, 0x0e, 0x34, 0xa3, 0xc5,
	0xfc, 0x3d, 0x65, 0x82, 0xe7, 0x82, 0xa5, 0x05, 0xd3, 0xab, 0x45, 0xa9, 0x53, 0x45, 0x39, 0x4b,
	0xcb, 0x8d, 0x62, 0x50, 0x15, 0xb2, 0x94, 0x7e, 0xcb, 0x4e, 0x42, 0x9a, 0x69, 0xb8, 0x17, 0x81,
	0xeb, 0x10, 0x1a, 0x91, 0x8b, 0xab, 0x9d, 0x87, 0xca, 0x03, 0x2a, 0x84, 0x2c, 0x69, 0x99, 0x4b,
	0xa1, 0xed, 0x70, 0xe7, 0xc3, 0x01, 0xd7, 0x13, 0x63, 0x13, 0x19, 0x97, 0xb1, 0x35, 0x89, 0x29,
	0x67, 0x4f, 0x1b, 0xc5, 0x22, 0xb1, 0x5a, 0x76, 0x14, 0xb8, 0x3c, 0xd0, 0xe2, 0x37, 0x41, 0x7d,
	0x4a, 0x26, 0x71, 0xd4, 0x1b, 0xde, 0x0f, 0xa3, 0xbe, 0x77, 0xe4, 0xd7, 0xc1, 0xe9, 0x94, 0x3c,
	0x90, 0xd1, 0x33, 0xf1, 0x1c, 0xbf, 0x01, 0xce, 0x70, 0x7f, 0x92, 0x8e, 0xc8, 0xe3, 0x8b, 0xe7,
	0xfa, 0x1e, 0x68, 0x8c, 0xc6, 0x03, 0x4c, 0x86, 0x3d, 0x5b, 0xa9, 0xf9, 0xe7, 0xa0, 0x59, 0xf1,
	0x98, 0xf4, 0xd3, 0x2d, 0xe3, 0x1d, 0x77, 0x7f, 0x1d, 0xd0, 0x9e, 0xcb, 0x25, 0x3c, 0x98, 0xac,
	0x7b, 0x73, 0x60, 0xab, 0xb8, 0x4a, 0x17, 0x3b, 0xaf, 0xdd, 0xad, 0x04, 0x97, 0x0b, 0x2a, 0x38,
	0x94, 0x05, 0x0f, 0x38, 0x13, 0x26, 0xfb, 0xee, 0xe2, 0x2a, 0xd7, 0xff, 0x7c, 0xc0, 0x9d, 0x79,
	0x3f, 0xdd, 0xda, 0x00, 0xe3, 0x2f, 0xb7, 0x35, 0xb0, 0x52, 0x38, 0xd3, 0xd0, 0xc2, 0x0a, 0xcd,
	0x42, 0x58, 0x1d, 0x49, 0x7f, 0xef, 0xf8, 0x04, 0x67, 0x3a, 0xd9, 0xf3, 0xc9, 0x2c, 0x4c, 0x0c,
	0xff, 0xe3, 0xb6, 0x6d, 0x11, 0x21, 0x9c, 0x69, 0x84, 0xf6, 0x1d, 0x08, 0xcd, 0x42, 0x84, 0x4c,
	0xcf, 0xdb, 0x89, 0x59, 0xec, 0xf6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x36, 0xa1, 0xbf, 0x51, 0x18,
	0x02, 0x00, 0x00,
}
