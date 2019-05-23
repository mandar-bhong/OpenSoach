// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/bidding_strategy_status.proto

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

// The possible statuses of a BiddingStrategy.
type BiddingStrategyStatusEnum_BiddingStrategyStatus int32

const (
	// No value has been specified.
	BiddingStrategyStatusEnum_UNSPECIFIED BiddingStrategyStatusEnum_BiddingStrategyStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	BiddingStrategyStatusEnum_UNKNOWN BiddingStrategyStatusEnum_BiddingStrategyStatus = 1
	// The bidding strategy is enabled.
	BiddingStrategyStatusEnum_ENABLED BiddingStrategyStatusEnum_BiddingStrategyStatus = 2
	// The bidding strategy is removed.
	BiddingStrategyStatusEnum_REMOVED BiddingStrategyStatusEnum_BiddingStrategyStatus = 4
)

var BiddingStrategyStatusEnum_BiddingStrategyStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	4: "REMOVED",
}

var BiddingStrategyStatusEnum_BiddingStrategyStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     4,
}

func (x BiddingStrategyStatusEnum_BiddingStrategyStatus) String() string {
	return proto.EnumName(BiddingStrategyStatusEnum_BiddingStrategyStatus_name, int32(x))
}

func (BiddingStrategyStatusEnum_BiddingStrategyStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e6f31065c294d03, []int{0, 0}
}

// Message describing BiddingStrategy statuses.
type BiddingStrategyStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiddingStrategyStatusEnum) Reset()         { *m = BiddingStrategyStatusEnum{} }
func (m *BiddingStrategyStatusEnum) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategyStatusEnum) ProtoMessage()    {}
func (*BiddingStrategyStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e6f31065c294d03, []int{0}
}

func (m *BiddingStrategyStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategyStatusEnum.Unmarshal(m, b)
}
func (m *BiddingStrategyStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategyStatusEnum.Marshal(b, m, deterministic)
}
func (m *BiddingStrategyStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategyStatusEnum.Merge(m, src)
}
func (m *BiddingStrategyStatusEnum) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategyStatusEnum.Size(m)
}
func (m *BiddingStrategyStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategyStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategyStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.BiddingStrategyStatusEnum_BiddingStrategyStatus", BiddingStrategyStatusEnum_BiddingStrategyStatus_name, BiddingStrategyStatusEnum_BiddingStrategyStatus_value)
	proto.RegisterType((*BiddingStrategyStatusEnum)(nil), "google.ads.googleads.v1.enums.BiddingStrategyStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/bidding_strategy_status.proto", fileDescriptor_1e6f31065c294d03)
}

var fileDescriptor_1e6f31065c294d03 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x75, 0x55, 0x14, 0xb2, 0x07, 0xcb, 0xc0, 0x07, 0x87, 0x7b, 0xd8, 0x3e, 0x20, 0xa1, 0xf8,
	0x96, 0x3d, 0xa5, 0x2e, 0x8e, 0xa1, 0x76, 0xc3, 0xb1, 0x0a, 0x52, 0x18, 0x99, 0x29, 0xa1, 0xb0,
	0x25, 0xa3, 0x49, 0x07, 0xfe, 0x8e, 0x8f, 0x7e, 0x8a, 0x9f, 0xe2, 0x8b, 0xbf, 0x20, 0x49, 0xda,
	0x3e, 0x4d, 0x5f, 0xc2, 0xb9, 0x39, 0xf7, 0x9c, 0x7b, 0xee, 0x05, 0x63, 0xa1, 0x94, 0xd8, 0xe6,
	0x88, 0x71, 0x8d, 0x3c, 0xb4, 0xe8, 0x10, 0xa1, 0x5c, 0x56, 0x3b, 0x8d, 0x36, 0x05, 0xe7, 0x85,
	0x14, 0x6b, 0x6d, 0x4a, 0x66, 0x72, 0xf1, 0xbe, 0xd6, 0x86, 0x99, 0x4a, 0xc3, 0x7d, 0xa9, 0x8c,
	0xea, 0x0d, 0xbc, 0x02, 0x32, 0xae, 0x61, 0x2b, 0x86, 0x87, 0x08, 0x3a, 0x71, 0xff, 0xa6, 0xf1,
	0xde, 0x17, 0x88, 0x49, 0xa9, 0x0c, 0x33, 0x85, 0x92, 0xb5, 0x78, 0xb4, 0x05, 0xd7, 0xb1, 0x77,
	0x5f, 0xd6, 0xe6, 0x4b, 0xe7, 0x4d, 0x65, 0xb5, 0x1b, 0xcd, 0xc1, 0xd5, 0x51, 0xb2, 0x77, 0x09,
	0xba, 0xab, 0x64, 0xb9, 0xa0, 0x77, 0xb3, 0xfb, 0x19, 0x9d, 0x84, 0x27, 0xbd, 0x2e, 0xb8, 0x58,
	0x25, 0x0f, 0xc9, 0xfc, 0x25, 0x09, 0x3b, 0xb6, 0xa0, 0x09, 0x89, 0x1f, 0xe9, 0x24, 0x0c, 0x6c,
	0xf1, 0x4c, 0x9f, 0xe6, 0x29, 0x9d, 0x84, 0x67, 0xf1, 0x4f, 0x07, 0x0c, 0xdf, 0xd4, 0x0e, 0xfe,
	0x9b, 0x38, 0xee, 0x1f, 0x1d, 0xba, 0xb0, 0x79, 0x17, 0x9d, 0xd7, 0xb8, 0x16, 0x0b, 0xb5, 0x65,
	0x52, 0x40, 0x55, 0x0a, 0x24, 0x72, 0xe9, 0xb6, 0x69, 0x6e, 0xb7, 0x2f, 0xf4, 0x1f, 0xa7, 0x1c,
	0xbb, 0xf7, 0x23, 0x38, 0x9d, 0x12, 0xf2, 0x19, 0x0c, 0xa6, 0xde, 0x8a, 0x70, 0x0d, 0x3d, 0xb4,
	0x28, 0x8d, 0xa0, 0xdd, 0x5e, 0x7f, 0x35, 0x7c, 0x46, 0xb8, 0xce, 0x5a, 0x3e, 0x4b, 0xa3, 0xcc,
	0xf1, 0xdf, 0xc1, 0xd0, 0x7f, 0x62, 0x4c, 0xb8, 0xc6, 0xb8, 0xed, 0xc0, 0x38, 0x8d, 0x30, 0x76,
	0x3d, 0x9b, 0x73, 0x17, 0xec, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xb7, 0x88, 0x78, 0xe2,
	0x01, 0x00, 0x00,
}
