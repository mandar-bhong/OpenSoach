// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_audience_view.proto

package resources

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

// A campaign audience view.
// Includes performance data from interests and remarketing lists for Display
// Network and YouTube Network ads, and remarketing lists for search ads (RLSA),
// aggregated by campaign and audience criterion. This view only includes
// audiences attached at the campaign level.
type CampaignAudienceView struct {
	// The resource name of the campaign audience view.
	// Campaign audience view resource names have the form:
	//
	//
	// `customers/{customer_id}/campaignAudienceViews/{campaign_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignAudienceView) Reset()         { *m = CampaignAudienceView{} }
func (m *CampaignAudienceView) String() string { return proto.CompactTextString(m) }
func (*CampaignAudienceView) ProtoMessage()    {}
func (*CampaignAudienceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_466ee4be4842c8d4, []int{0}
}

func (m *CampaignAudienceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignAudienceView.Unmarshal(m, b)
}
func (m *CampaignAudienceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignAudienceView.Marshal(b, m, deterministic)
}
func (m *CampaignAudienceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignAudienceView.Merge(m, src)
}
func (m *CampaignAudienceView) XXX_Size() int {
	return xxx_messageInfo_CampaignAudienceView.Size(m)
}
func (m *CampaignAudienceView) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignAudienceView.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignAudienceView proto.InternalMessageInfo

func (m *CampaignAudienceView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*CampaignAudienceView)(nil), "google.ads.googleads.v1.resources.CampaignAudienceView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_audience_view.proto", fileDescriptor_466ee4be4842c8d4)
}

var fileDescriptor_466ee4be4842c8d4 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4f, 0x4a, 0xc4, 0x30,
	0x14, 0xc6, 0x99, 0x0a, 0x82, 0x45, 0x37, 0x83, 0x0b, 0x15, 0x17, 0x8e, 0x32, 0xe0, 0x2a, 0xa1,
	0xb8, 0xcb, 0x80, 0x90, 0x71, 0x31, 0xe0, 0x42, 0x86, 0x59, 0x74, 0x21, 0x85, 0xf2, 0x6c, 0x1e,
	0x21, 0x30, 0xcd, 0x2b, 0x4d, 0xa7, 0x73, 0x03, 0x0f, 0xe2, 0xd2, 0xa3, 0x78, 0x14, 0x4f, 0x21,
	0x9d, 0x34, 0x59, 0x89, 0xee, 0x3e, 0x92, 0xdf, 0xf7, 0x87, 0x97, 0x3e, 0x6a, 0x22, 0xbd, 0x45,
	0x0e, 0xca, 0x71, 0x2f, 0x07, 0xd5, 0x67, 0xbc, 0x45, 0x47, 0xbb, 0xb6, 0x42, 0xc7, 0x2b, 0xa8,
	0x1b, 0x30, 0xda, 0x96, 0xb0, 0x53, 0x06, 0x6d, 0x85, 0x65, 0x6f, 0x70, 0xcf, 0x9a, 0x96, 0x3a,
	0x9a, 0xce, 0xbc, 0x89, 0x81, 0x72, 0x2c, 0xfa, 0x59, 0x9f, 0xb1, 0xe8, 0xbf, 0xba, 0x0e, 0x15,
	0x8d, 0xe1, 0x60, 0x2d, 0x75, 0xd0, 0x19, 0xb2, 0xce, 0x07, 0xdc, 0x2e, 0xd2, 0xf3, 0xa7, 0xb1,
	0x40, 0x8e, 0xf9, 0xb9, 0xc1, 0xfd, 0xf4, 0x2e, 0x3d, 0x0b, 0x11, 0xa5, 0x85, 0x1a, 0x2f, 0x26,
	0x37, 0x93, 0xfb, 0x93, 0xcd, 0x69, 0x78, 0x7c, 0x81, 0x1a, 0x97, 0xef, 0x49, 0x3a, 0xaf, 0xa8,
	0x66, 0xff, 0x8e, 0x58, 0x5e, 0xfe, 0x56, 0xb2, 0x1e, 0x16, 0xac, 0x27, 0xaf, 0xcf, 0xa3, 0x5f,
	0xd3, 0x16, 0xac, 0x66, 0xd4, 0x6a, 0xae, 0xd1, 0x1e, 0xf6, 0x85, 0xa3, 0x34, 0xc6, 0xfd, 0x71,
	0xa3, 0x45, 0x54, 0x1f, 0xc9, 0xd1, 0x4a, 0xca, 0xcf, 0x64, 0xb6, 0xf2, 0x91, 0x52, 0x39, 0xe6,
	0xe5, 0xa0, 0xf2, 0x8c, 0x6d, 0x02, 0xf9, 0x15, 0x98, 0x42, 0x2a, 0x57, 0x44, 0xa6, 0xc8, 0xb3,
	0x22, 0x32, 0xdf, 0xc9, 0xdc, 0x7f, 0x08, 0x21, 0x95, 0x13, 0x22, 0x52, 0x42, 0xe4, 0x99, 0x10,
	0x91, 0x7b, 0x3b, 0x3e, 0x8c, 0x7d, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x40, 0xc3, 0x93,
	0xcf, 0x01, 0x00, 0x00,
}
