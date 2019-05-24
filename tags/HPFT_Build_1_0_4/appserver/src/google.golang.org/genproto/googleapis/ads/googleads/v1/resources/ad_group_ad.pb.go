// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_group_ad.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// An ad group ad.
type AdGroupAd struct {
	// The resource name of the ad.
	// Ad group ad resource names have the form:
	//
	// `customers/{customer_id}/adGroupAds/{ad_group_id}~{ad_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The status of the ad.
	Status enums.AdGroupAdStatusEnum_AdGroupAdStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.AdGroupAdStatusEnum_AdGroupAdStatus" json:"status,omitempty"`
	// The ad group to which the ad belongs.
	AdGroup *wrappers.StringValue `protobuf:"bytes,4,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The ad.
	Ad *Ad `protobuf:"bytes,5,opt,name=ad,proto3" json:"ad,omitempty"`
	// Policy information for the ad.
	PolicySummary *AdGroupAdPolicySummary `protobuf:"bytes,6,opt,name=policy_summary,json=policySummary,proto3" json:"policy_summary,omitempty"`
	// Overall ad strength for this ad group ad.
	AdStrength           enums.AdStrengthEnum_AdStrength `protobuf:"varint,7,opt,name=ad_strength,json=adStrength,proto3,enum=google.ads.googleads.v1.enums.AdStrengthEnum_AdStrength" json:"ad_strength,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *AdGroupAd) Reset()         { *m = AdGroupAd{} }
func (m *AdGroupAd) String() string { return proto.CompactTextString(m) }
func (*AdGroupAd) ProtoMessage()    {}
func (*AdGroupAd) Descriptor() ([]byte, []int) {
	return fileDescriptor_f498dcb629eaef64, []int{0}
}

func (m *AdGroupAd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAd.Unmarshal(m, b)
}
func (m *AdGroupAd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAd.Marshal(b, m, deterministic)
}
func (m *AdGroupAd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAd.Merge(m, src)
}
func (m *AdGroupAd) XXX_Size() int {
	return xxx_messageInfo_AdGroupAd.Size(m)
}
func (m *AdGroupAd) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAd.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAd proto.InternalMessageInfo

func (m *AdGroupAd) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupAd) GetStatus() enums.AdGroupAdStatusEnum_AdGroupAdStatus {
	if m != nil {
		return m.Status
	}
	return enums.AdGroupAdStatusEnum_UNSPECIFIED
}

func (m *AdGroupAd) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupAd) GetAd() *Ad {
	if m != nil {
		return m.Ad
	}
	return nil
}

func (m *AdGroupAd) GetPolicySummary() *AdGroupAdPolicySummary {
	if m != nil {
		return m.PolicySummary
	}
	return nil
}

func (m *AdGroupAd) GetAdStrength() enums.AdStrengthEnum_AdStrength {
	if m != nil {
		return m.AdStrength
	}
	return enums.AdStrengthEnum_UNSPECIFIED
}

// Contains policy information for an ad.
type AdGroupAdPolicySummary struct {
	// The list of policy findings for this ad.
	PolicyTopicEntries []*common.PolicyTopicEntry `protobuf:"bytes,1,rep,name=policy_topic_entries,json=policyTopicEntries,proto3" json:"policy_topic_entries,omitempty"`
	// Where in the review process this ad is.
	ReviewStatus enums.PolicyReviewStatusEnum_PolicyReviewStatus `protobuf:"varint,2,opt,name=review_status,json=reviewStatus,proto3,enum=google.ads.googleads.v1.enums.PolicyReviewStatusEnum_PolicyReviewStatus" json:"review_status,omitempty"`
	// The overall approval status of this ad, calculated based on the status of
	// its individual policy topic entries.
	ApprovalStatus       enums.PolicyApprovalStatusEnum_PolicyApprovalStatus `protobuf:"varint,3,opt,name=approval_status,json=approvalStatus,proto3,enum=google.ads.googleads.v1.enums.PolicyApprovalStatusEnum_PolicyApprovalStatus" json:"approval_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                            `json:"-"`
	XXX_unrecognized     []byte                                              `json:"-"`
	XXX_sizecache        int32                                               `json:"-"`
}

func (m *AdGroupAdPolicySummary) Reset()         { *m = AdGroupAdPolicySummary{} }
func (m *AdGroupAdPolicySummary) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdPolicySummary) ProtoMessage()    {}
func (*AdGroupAdPolicySummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_f498dcb629eaef64, []int{1}
}

func (m *AdGroupAdPolicySummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdPolicySummary.Unmarshal(m, b)
}
func (m *AdGroupAdPolicySummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdPolicySummary.Marshal(b, m, deterministic)
}
func (m *AdGroupAdPolicySummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdPolicySummary.Merge(m, src)
}
func (m *AdGroupAdPolicySummary) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdPolicySummary.Size(m)
}
func (m *AdGroupAdPolicySummary) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdPolicySummary.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdPolicySummary proto.InternalMessageInfo

func (m *AdGroupAdPolicySummary) GetPolicyTopicEntries() []*common.PolicyTopicEntry {
	if m != nil {
		return m.PolicyTopicEntries
	}
	return nil
}

func (m *AdGroupAdPolicySummary) GetReviewStatus() enums.PolicyReviewStatusEnum_PolicyReviewStatus {
	if m != nil {
		return m.ReviewStatus
	}
	return enums.PolicyReviewStatusEnum_UNSPECIFIED
}

func (m *AdGroupAdPolicySummary) GetApprovalStatus() enums.PolicyApprovalStatusEnum_PolicyApprovalStatus {
	if m != nil {
		return m.ApprovalStatus
	}
	return enums.PolicyApprovalStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AdGroupAd)(nil), "google.ads.googleads.v1.resources.AdGroupAd")
	proto.RegisterType((*AdGroupAdPolicySummary)(nil), "google.ads.googleads.v1.resources.AdGroupAdPolicySummary")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_group_ad.proto", fileDescriptor_f498dcb629eaef64)
}

var fileDescriptor_f498dcb629eaef64 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdf, 0x6a, 0xdb, 0x3e,
	0x18, 0xc5, 0xce, 0xef, 0x97, 0xae, 0xca, 0x9a, 0x81, 0x19, 0xc3, 0x94, 0x32, 0xd2, 0x8e, 0x40,
	0xd8, 0x40, 0x5e, 0x52, 0xb6, 0x75, 0xde, 0x95, 0x03, 0x25, 0x63, 0x8c, 0x51, 0x9c, 0x11, 0x58,
	0x09, 0x78, 0x4a, 0xa4, 0x79, 0x86, 0x58, 0x12, 0x92, 0x9d, 0x92, 0xd7, 0xd9, 0xe5, 0xde, 0x60,
	0xaf, 0xb0, 0xc7, 0xd8, 0xe5, 0x5e, 0x61, 0x37, 0xc3, 0x96, 0xe4, 0xc6, 0xed, 0xd2, 0xe4, 0xee,
	0xfb, 0xa4, 0x73, 0xce, 0xf7, 0x47, 0x47, 0xe0, 0x34, 0x66, 0x2c, 0x5e, 0x10, 0x0f, 0x61, 0xe9,
	0xa9, 0xb0, 0x88, 0x96, 0x7d, 0x4f, 0x10, 0xc9, 0x72, 0x31, 0x27, 0xd2, 0x43, 0x38, 0x8a, 0x05,
	0xcb, 0x79, 0x84, 0x30, 0xe4, 0x82, 0x65, 0xcc, 0x39, 0x56, 0x48, 0x88, 0xb0, 0x84, 0x15, 0x09,
	0x2e, 0xfb, 0xb0, 0x22, 0x1d, 0x3e, 0xdb, 0xa4, 0x3b, 0x67, 0x69, 0xca, 0xa8, 0xc7, 0xd9, 0x22,
	0x99, 0xaf, 0x94, 0xde, 0xe1, 0xcb, 0x4d, 0x60, 0x42, 0xf3, 0xb4, 0xd6, 0x40, 0x24, 0x33, 0x94,
	0xe5, 0x52, 0xf3, 0xbc, 0xad, 0x3c, 0x99, 0x09, 0x42, 0xe3, 0xec, 0xab, 0x26, 0xf8, 0x77, 0x13,
	0x54, 0x53, 0x11, 0xe2, 0x5c, 0xb0, 0x25, 0x5a, 0xd4, 0x8b, 0x9d, 0xed, 0xc4, 0x15, 0x64, 0x99,
	0x90, 0xab, 0x3a, 0xf3, 0xe9, 0x2e, 0x3b, 0xd6, 0xd8, 0xc7, 0x1a, 0x5b, 0x66, 0xb3, 0xfc, 0x8b,
	0x77, 0x25, 0x10, 0xe7, 0x44, 0x18, 0xad, 0x23, 0xa3, 0xc5, 0x13, 0x0f, 0x51, 0xca, 0x32, 0x94,
	0x25, 0x8c, 0xea, 0xdb, 0x93, 0x1f, 0x0d, 0xb0, 0x1f, 0xe0, 0x51, 0xb1, 0xac, 0x00, 0x3b, 0x4f,
	0xc0, 0x81, 0xa9, 0x10, 0x51, 0x94, 0x12, 0xd7, 0xea, 0x58, 0xbd, 0xfd, 0xf0, 0xbe, 0x39, 0xfc,
	0x80, 0x52, 0xe2, 0x5c, 0x82, 0xa6, 0x6a, 0xd6, 0x6d, 0x74, 0xac, 0x5e, 0x7b, 0x30, 0x84, 0x9b,
	0x1e, 0xb7, 0x9c, 0x13, 0x56, 0xf2, 0xe3, 0x92, 0x75, 0x4e, 0xf3, 0xf4, 0xe6, 0x59, 0xa8, 0x15,
	0x9d, 0x57, 0xe0, 0x9e, 0x79, 0x3b, 0xf7, 0xbf, 0x8e, 0xd5, 0x6b, 0x0d, 0x8e, 0x8c, 0xba, 0x99,
	0x0f, 0x8e, 0x33, 0x91, 0xd0, 0x78, 0x82, 0x16, 0x39, 0x09, 0xf7, 0x90, 0x12, 0x72, 0x5e, 0x00,
	0x1b, 0x61, 0xf7, 0xff, 0x92, 0xd2, 0x85, 0x5b, 0xdd, 0x06, 0x03, 0x1c, 0xda, 0x08, 0x3b, 0x9f,
	0x41, 0x5b, 0x3f, 0x83, 0xcc, 0xd3, 0x14, 0x89, 0x95, 0xdb, 0x2c, 0x25, 0x5e, 0xef, 0x24, 0xa1,
	0x67, 0xb8, 0x28, 0x15, 0xc6, 0x4a, 0x20, 0x3c, 0xe0, 0xeb, 0xa9, 0xf3, 0x09, 0xb4, 0xd6, 0x5c,
	0xe5, 0xee, 0x95, 0x2b, 0x3b, 0xdb, 0xba, 0xb2, 0xb1, 0x26, 0xe8, 0x6d, 0x99, 0x34, 0x04, 0xa8,
	0x8a, 0x4f, 0x7e, 0xd9, 0xe0, 0xd1, 0xbf, 0x9b, 0x70, 0x66, 0xe0, 0xa1, 0x9e, 0x2b, 0x63, 0x3c,
	0x99, 0x47, 0x84, 0x66, 0x22, 0x21, 0xd2, 0xb5, 0x3a, 0x8d, 0x5e, 0x6b, 0xf0, 0x7c, 0x63, 0x79,
	0xf5, 0xd7, 0xa0, 0x12, 0xfb, 0x58, 0x50, 0xcf, 0x69, 0x26, 0x56, 0xa1, 0xc3, 0xeb, 0x27, 0x09,
	0x91, 0x4e, 0x5a, 0x98, 0x65, 0xcd, 0xbb, 0xae, 0x5d, 0xce, 0xf6, 0x76, 0xcb, 0x6c, 0x4a, 0x3b,
	0x2c, 0x99, 0x6b, 0x8e, 0xb8, 0x7d, 0x5c, 0xd8, 0xee, 0x3a, 0x73, 0x72, 0xf0, 0xe0, 0xc6, 0x37,
	0xd3, 0xfe, 0x7b, 0xbf, 0x53, 0xc1, 0x40, 0x73, 0x6f, 0x95, 0xac, 0x5f, 0x84, 0x6d, 0x54, 0xcb,
	0x87, 0x7f, 0x2c, 0xd0, 0x9d, 0xb3, 0x74, 0xbb, 0x1f, 0x86, 0xed, 0xeb, 0xb7, 0x28, 0xac, 0x7a,
	0x61, 0x5d, 0xbe, 0xd3, 0xa4, 0x98, 0x2d, 0x10, 0x8d, 0x21, 0x13, 0xb1, 0x17, 0x13, 0x5a, 0x1a,
	0xd9, 0x7c, 0x6b, 0x9e, 0xc8, 0x3b, 0x7e, 0xf9, 0x9b, 0x2a, 0xfa, 0x66, 0x37, 0x46, 0x41, 0xf0,
	0xdd, 0x3e, 0x1e, 0x29, 0xc9, 0x00, 0x4b, 0xa8, 0xc2, 0x22, 0x9a, 0xf4, 0x61, 0x68, 0x90, 0x3f,
	0x0d, 0x66, 0x1a, 0x60, 0x39, 0xad, 0x30, 0xd3, 0x49, 0x7f, 0x5a, 0x61, 0x7e, 0xdb, 0x5d, 0x75,
	0xe1, 0xfb, 0x01, 0x96, 0xbe, 0x5f, 0xa1, 0x7c, 0x7f, 0xd2, 0xf7, 0xfd, 0x0a, 0x37, 0x6b, 0x96,
	0xcd, 0x9e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xda, 0xae, 0x85, 0x3e, 0xf5, 0x05, 0x00, 0x00,
}
