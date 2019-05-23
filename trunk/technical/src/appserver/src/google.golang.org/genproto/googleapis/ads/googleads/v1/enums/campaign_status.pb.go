// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/campaign_status.proto

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

// Possible statuses of a campaign.
type CampaignStatusEnum_CampaignStatus int32

const (
	// Not specified.
	CampaignStatusEnum_UNSPECIFIED CampaignStatusEnum_CampaignStatus = 0
	// Used for return value only. Represents value unknown in this version.
	CampaignStatusEnum_UNKNOWN CampaignStatusEnum_CampaignStatus = 1
	// Campaign is currently serving ads depending on budget information.
	CampaignStatusEnum_ENABLED CampaignStatusEnum_CampaignStatus = 2
	// Campaign has been paused by the user.
	CampaignStatusEnum_PAUSED CampaignStatusEnum_CampaignStatus = 3
	// Campaign has been removed.
	CampaignStatusEnum_REMOVED CampaignStatusEnum_CampaignStatus = 4
)

var CampaignStatusEnum_CampaignStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PAUSED",
	4: "REMOVED",
}

var CampaignStatusEnum_CampaignStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PAUSED":      3,
	"REMOVED":     4,
}

func (x CampaignStatusEnum_CampaignStatus) String() string {
	return proto.EnumName(CampaignStatusEnum_CampaignStatus_name, int32(x))
}

func (CampaignStatusEnum_CampaignStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6d9d0ca16fb2fd74, []int{0, 0}
}

// Container for enum describing possible statuses of a campaign.
type CampaignStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignStatusEnum) Reset()         { *m = CampaignStatusEnum{} }
func (m *CampaignStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignStatusEnum) ProtoMessage()    {}
func (*CampaignStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d9d0ca16fb2fd74, []int{0}
}

func (m *CampaignStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignStatusEnum.Unmarshal(m, b)
}
func (m *CampaignStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignStatusEnum.Marshal(b, m, deterministic)
}
func (m *CampaignStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignStatusEnum.Merge(m, src)
}
func (m *CampaignStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignStatusEnum.Size(m)
}
func (m *CampaignStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.CampaignStatusEnum_CampaignStatus", CampaignStatusEnum_CampaignStatus_name, CampaignStatusEnum_CampaignStatus_value)
	proto.RegisterType((*CampaignStatusEnum)(nil), "google.ads.googleads.v1.enums.CampaignStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/campaign_status.proto", fileDescriptor_6d9d0ca16fb2fd74)
}

var fileDescriptor_6d9d0ca16fb2fd74 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdd, 0x4a, 0xc3, 0x30,
	0x14, 0x76, 0x9d, 0x4c, 0xc8, 0x40, 0x4b, 0xbd, 0x13, 0x77, 0xb1, 0x3d, 0x40, 0x42, 0xd9, 0x5d,
	0xbc, 0x4a, 0xd7, 0x38, 0x86, 0xda, 0x15, 0x67, 0x2b, 0x48, 0x41, 0xe2, 0x5a, 0x42, 0x65, 0x4d,
	0xca, 0xd2, 0xee, 0x81, 0xbc, 0xf4, 0x51, 0x7c, 0x12, 0xf1, 0x29, 0x24, 0xe9, 0x0f, 0xec, 0x42,
	0x6f, 0xca, 0xd7, 0xf3, 0xfd, 0xe4, 0x3b, 0x07, 0xcc, 0xb9, 0x94, 0x7c, 0x97, 0x21, 0x96, 0x2a,
	0xd4, 0x40, 0x8d, 0x0e, 0x2e, 0xca, 0x44, 0x5d, 0x28, 0xb4, 0x65, 0x45, 0xc9, 0x72, 0x2e, 0x5e,
	0x55, 0xc5, 0xaa, 0x5a, 0xc1, 0x72, 0x2f, 0x2b, 0xe9, 0x4c, 0x1a, 0x25, 0x64, 0xa9, 0x82, 0xbd,
	0x09, 0x1e, 0x5c, 0x68, 0x4c, 0x57, 0xd7, 0x5d, 0x66, 0x99, 0x23, 0x26, 0x84, 0xac, 0x58, 0x95,
	0x4b, 0xd1, 0x9a, 0x67, 0xef, 0xc0, 0x59, 0xb4, 0xa9, 0x1b, 0x13, 0x4a, 0x45, 0x5d, 0xcc, 0x9e,
	0xc0, 0xf9, 0xf1, 0xd4, 0xb9, 0x00, 0xe3, 0x28, 0xd8, 0x84, 0x74, 0xb1, 0xba, 0x5d, 0x51, 0xdf,
	0x3e, 0x71, 0xc6, 0xe0, 0x2c, 0x0a, 0xee, 0x82, 0xf5, 0x73, 0x60, 0x0f, 0xf4, 0x0f, 0x0d, 0x88,
	0x77, 0x4f, 0x7d, 0xdb, 0x72, 0x00, 0x18, 0x85, 0x24, 0xda, 0x50, 0xdf, 0x1e, 0x6a, 0xe2, 0x91,
	0x3e, 0xac, 0x63, 0xea, 0xdb, 0xa7, 0xde, 0xf7, 0x00, 0x4c, 0xb7, 0xb2, 0x80, 0xff, 0xf6, 0xf5,
	0x2e, 0x8f, 0x5f, 0x0e, 0x75, 0xcd, 0x70, 0xf0, 0xe2, 0xb5, 0x2e, 0x2e, 0x77, 0x4c, 0x70, 0x28,
	0xf7, 0x1c, 0xf1, 0x4c, 0x98, 0x25, 0xba, 0x53, 0x95, 0xb9, 0xfa, 0xe3, 0x72, 0x37, 0xe6, 0xfb,
	0x61, 0x0d, 0x97, 0x84, 0x7c, 0x5a, 0x93, 0x65, 0x13, 0x45, 0x52, 0x05, 0x1b, 0xa8, 0x51, 0xec,
	0x42, 0xbd, 0xbb, 0xfa, 0xea, 0xf8, 0x84, 0xa4, 0x2a, 0xe9, 0xf9, 0x24, 0x76, 0x13, 0xc3, 0xff,
	0x58, 0xd3, 0x66, 0x88, 0x31, 0x49, 0x15, 0xc6, 0xbd, 0x02, 0xe3, 0xd8, 0xc5, 0xd8, 0x68, 0xde,
	0x46, 0xa6, 0xd8, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x1f, 0xe1, 0x6b, 0xd1, 0x01, 0x00,
	0x00,
}
