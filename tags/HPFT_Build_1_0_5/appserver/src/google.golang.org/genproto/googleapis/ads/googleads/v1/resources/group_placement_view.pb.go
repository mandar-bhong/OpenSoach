// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/group_placement_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A group placement view.
type GroupPlacementView struct {
	// The resource name of the group placement view.
	// Group placement view resource names have the form:
	//
	//
	// `customers/{customer_id}/groupPlacementViews/{ad_group_id}~{base64_placement}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The automatic placement string at group level, e. g. web domain, mobile
	// app ID, or a YouTube channel ID.
	Placement *wrappers.StringValue `protobuf:"bytes,2,opt,name=placement,proto3" json:"placement,omitempty"`
	// Domain name for websites and YouTube channel name for YouTube channels.
	DisplayName *wrappers.StringValue `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// URL of the group placement, e.g. domain, link to the mobile application in
	// app store, or a YouTube channel URL.
	TargetUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=target_url,json=targetUrl,proto3" json:"target_url,omitempty"`
	// Type of the placement, e.g. Website, YouTube Channel, Mobile Application.
	PlacementType        enums.PlacementTypeEnum_PlacementType `protobuf:"varint,5,opt,name=placement_type,json=placementType,proto3,enum=google.ads.googleads.v1.enums.PlacementTypeEnum_PlacementType" json:"placement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *GroupPlacementView) Reset()         { *m = GroupPlacementView{} }
func (m *GroupPlacementView) String() string { return proto.CompactTextString(m) }
func (*GroupPlacementView) ProtoMessage()    {}
func (*GroupPlacementView) Descriptor() ([]byte, []int) {
	return fileDescriptor_40a38bfc66391954, []int{0}
}

func (m *GroupPlacementView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupPlacementView.Unmarshal(m, b)
}
func (m *GroupPlacementView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupPlacementView.Marshal(b, m, deterministic)
}
func (m *GroupPlacementView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupPlacementView.Merge(m, src)
}
func (m *GroupPlacementView) XXX_Size() int {
	return xxx_messageInfo_GroupPlacementView.Size(m)
}
func (m *GroupPlacementView) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupPlacementView.DiscardUnknown(m)
}

var xxx_messageInfo_GroupPlacementView proto.InternalMessageInfo

func (m *GroupPlacementView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *GroupPlacementView) GetPlacement() *wrappers.StringValue {
	if m != nil {
		return m.Placement
	}
	return nil
}

func (m *GroupPlacementView) GetDisplayName() *wrappers.StringValue {
	if m != nil {
		return m.DisplayName
	}
	return nil
}

func (m *GroupPlacementView) GetTargetUrl() *wrappers.StringValue {
	if m != nil {
		return m.TargetUrl
	}
	return nil
}

func (m *GroupPlacementView) GetPlacementType() enums.PlacementTypeEnum_PlacementType {
	if m != nil {
		return m.PlacementType
	}
	return enums.PlacementTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*GroupPlacementView)(nil), "google.ads.googleads.v1.resources.GroupPlacementView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/group_placement_view.proto", fileDescriptor_40a38bfc66391954)
}

var fileDescriptor_40a38bfc66391954 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x8a, 0xd4, 0x30,
	0x14, 0xa6, 0x5d, 0x15, 0x36, 0xfb, 0x73, 0xd1, 0x1b, 0x87, 0x65, 0x91, 0x59, 0x65, 0x61, 0xae,
	0x12, 0x3a, 0xde, 0x65, 0x45, 0xe9, 0x82, 0x0c, 0x78, 0x21, 0x43, 0xd5, 0x5e, 0x48, 0xa1, 0x64,
	0xa7, 0xc7, 0x50, 0x68, 0x93, 0x90, 0xa4, 0x33, 0xcc, 0xbd, 0x4f, 0xe2, 0xa5, 0x8f, 0xe2, 0x23,
	0xf8, 0x08, 0x3e, 0x85, 0xb4, 0x69, 0xb2, 0x2e, 0xb2, 0xba, 0x77, 0x5f, 0x72, 0xbe, 0xef, 0x3b,
	0xdf, 0xc9, 0x09, 0x7a, 0xc5, 0xa5, 0xe4, 0x2d, 0x10, 0x56, 0x1b, 0xe2, 0xe0, 0x80, 0xb6, 0x29,
	0xd1, 0x60, 0x64, 0xaf, 0x37, 0x60, 0x08, 0xd7, 0xb2, 0x57, 0x95, 0x6a, 0xd9, 0x06, 0x3a, 0x10,
	0xb6, 0xda, 0x36, 0xb0, 0xc3, 0x4a, 0x4b, 0x2b, 0x93, 0x0b, 0x27, 0xc1, 0xac, 0x36, 0x38, 0xa8,
	0xf1, 0x36, 0xc5, 0x41, 0x7d, 0xb6, 0xbc, 0xaf, 0x01, 0x88, 0xbe, 0x33, 0xe4, 0xd6, 0xd6, 0xee,
	0x15, 0x38, 0xdb, 0xb3, 0x67, 0x93, 0x66, 0x3c, 0xdd, 0xf4, 0x5f, 0xc8, 0x4e, 0x33, 0xa5, 0x40,
	0x9b, 0xa9, 0x7e, 0xee, 0x3d, 0x55, 0x43, 0x98, 0x10, 0xd2, 0x32, 0xdb, 0x48, 0x31, 0x55, 0x9f,
	0xff, 0x8c, 0x51, 0xb2, 0x1a, 0x32, 0xaf, 0xbd, 0x77, 0xd1, 0xc0, 0x2e, 0x79, 0x81, 0x4e, 0x7c,
	0xaa, 0x4a, 0xb0, 0x0e, 0x66, 0xd1, 0x3c, 0x5a, 0x1c, 0xe6, 0xc7, 0xfe, 0xf2, 0x3d, 0xeb, 0x20,
	0xa1, 0xe8, 0x30, 0x24, 0x9a, 0xc5, 0xf3, 0x68, 0x71, 0xb4, 0x3c, 0x9f, 0x26, 0xc3, 0x3e, 0x0d,
	0xfe, 0x60, 0x75, 0x23, 0x78, 0xc1, 0xda, 0x1e, 0xf2, 0x5b, 0x7a, 0xf2, 0x06, 0x1d, 0xd7, 0x8d,
	0x51, 0x2d, 0xdb, 0x3b, 0xff, 0x83, 0x07, 0xc8, 0x8f, 0x26, 0xc5, 0xd8, 0xfc, 0x0a, 0x21, 0xcb,
	0x34, 0x07, 0x5b, 0xf5, 0xba, 0x9d, 0x3d, 0x7a, 0x48, 0x77, 0xc7, 0xff, 0xa4, 0xdb, 0x04, 0xd0,
	0xe9, 0xdd, 0xb7, 0x9c, 0x3d, 0x9e, 0x47, 0x8b, 0xd3, 0xe5, 0x6b, 0x7c, 0xdf, 0x8e, 0xc6, 0x05,
	0xe0, 0xf0, 0x48, 0x1f, 0xf7, 0x0a, 0xde, 0x8a, 0xbe, 0xbb, 0x7b, 0x93, 0x9f, 0xa8, 0x3f, 0x8f,
	0xd7, 0x5f, 0x63, 0x74, 0xb9, 0x91, 0x1d, 0xfe, 0xef, 0xe2, 0xaf, 0x9f, 0xfe, 0xbd, 0x83, 0xf5,
	0x30, 0xc3, 0x3a, 0xfa, 0xfc, 0x6e, 0x52, 0x73, 0xd9, 0x32, 0xc1, 0xb1, 0xd4, 0x9c, 0x70, 0x10,
	0xe3, 0x84, 0xfe, 0x8f, 0xa8, 0xc6, 0xfc, 0xe3, 0x4f, 0x5e, 0x05, 0xf4, 0x2d, 0x3e, 0x58, 0x65,
	0xd9, 0xf7, 0xf8, 0x62, 0xe5, 0x2c, 0xb3, 0xda, 0x60, 0x07, 0x07, 0x54, 0xa4, 0x38, 0xf7, 0xcc,
	0x1f, 0x9e, 0x53, 0x66, 0xb5, 0x29, 0x03, 0xa7, 0x2c, 0xd2, 0x32, 0x70, 0x7e, 0xc5, 0x97, 0xae,
	0x40, 0x69, 0x56, 0x1b, 0x4a, 0x03, 0x8b, 0xd2, 0x22, 0xa5, 0x34, 0xf0, 0x6e, 0x9e, 0x8c, 0x61,
	0x5f, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x6f, 0x57, 0xe2, 0x3f, 0x03, 0x00, 0x00,
}
