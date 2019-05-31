// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/click_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
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

// A click view with metrics aggregated at each click level, including both
// valid and invalid clicks. For non-Search campaigns, metrics.clicks
// represents the number of valid and invalid interactions.
type ClickView struct {
	// The resource name of the click view.
	// Click view resource names have the form:
	//
	// `customers/{customer_id}/clickViews/{date (yyyy-MM-dd)}~{gclid}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The Google Click ID.
	Gclid *wrappers.StringValue `protobuf:"bytes,2,opt,name=gclid,proto3" json:"gclid,omitempty"`
	// The location criteria matching the area of interest associated with the
	// impression.
	AreaOfInterest *common.ClickLocation `protobuf:"bytes,3,opt,name=area_of_interest,json=areaOfInterest,proto3" json:"area_of_interest,omitempty"`
	// The location criteria matching the location of presence associated with the
	// impression.
	LocationOfPresence *common.ClickLocation `protobuf:"bytes,4,opt,name=location_of_presence,json=locationOfPresence,proto3" json:"location_of_presence,omitempty"`
	// Page number in search results where the ad was shown.
	PageNumber           *wrappers.Int64Value `protobuf:"bytes,5,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClickView) Reset()         { *m = ClickView{} }
func (m *ClickView) String() string { return proto.CompactTextString(m) }
func (*ClickView) ProtoMessage()    {}
func (*ClickView) Descriptor() ([]byte, []int) {
	return fileDescriptor_c61fa672950e615e, []int{0}
}

func (m *ClickView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickView.Unmarshal(m, b)
}
func (m *ClickView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickView.Marshal(b, m, deterministic)
}
func (m *ClickView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickView.Merge(m, src)
}
func (m *ClickView) XXX_Size() int {
	return xxx_messageInfo_ClickView.Size(m)
}
func (m *ClickView) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickView.DiscardUnknown(m)
}

var xxx_messageInfo_ClickView proto.InternalMessageInfo

func (m *ClickView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ClickView) GetGclid() *wrappers.StringValue {
	if m != nil {
		return m.Gclid
	}
	return nil
}

func (m *ClickView) GetAreaOfInterest() *common.ClickLocation {
	if m != nil {
		return m.AreaOfInterest
	}
	return nil
}

func (m *ClickView) GetLocationOfPresence() *common.ClickLocation {
	if m != nil {
		return m.LocationOfPresence
	}
	return nil
}

func (m *ClickView) GetPageNumber() *wrappers.Int64Value {
	if m != nil {
		return m.PageNumber
	}
	return nil
}

func init() {
	proto.RegisterType((*ClickView)(nil), "google.ads.googleads.v1.resources.ClickView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/click_view.proto", fileDescriptor_c61fa672950e615e)
}

var fileDescriptor_c61fa672950e615e = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0x87, 0x49, 0x6a, 0x85, 0x4e, 0xb5, 0x48, 0xf0, 0x22, 0xd4, 0x22, 0xad, 0x52, 0xe8, 0x8d,
	0x13, 0xb2, 0x15, 0x2f, 0xa2, 0x37, 0xa9, 0x17, 0xa5, 0x22, 0xed, 0xb2, 0x42, 0x04, 0x59, 0x08,
	0xb3, 0x93, 0xb3, 0xc3, 0x60, 0x32, 0x13, 0x66, 0x92, 0xdd, 0xf7, 0xf1, 0xd2, 0x47, 0xf1, 0x19,
	0x7c, 0x02, 0x5f, 0xc1, 0x1b, 0xc9, 0xfc, 0xbb, 0x91, 0x55, 0xe8, 0xdd, 0xd9, 0x3d, 0xdf, 0xef,
	0xcb, 0x39, 0x33, 0x83, 0x66, 0x4c, 0x4a, 0xd6, 0x42, 0x46, 0x1a, 0x9d, 0xd9, 0x72, 0xaa, 0x36,
	0x79, 0xa6, 0x40, 0xcb, 0x51, 0x51, 0xd0, 0x19, 0x6d, 0x39, 0xfd, 0x5a, 0x6f, 0x38, 0x6c, 0x71,
	0xaf, 0xe4, 0x20, 0x93, 0x33, 0x0b, 0x62, 0xd2, 0x68, 0x1c, 0x32, 0x78, 0x93, 0xe3, 0x90, 0x39,
	0xbe, 0xdc, 0xa5, 0xa5, 0xb2, 0xeb, 0xa4, 0x70, 0xce, 0x56, 0x52, 0x32, 0x70, 0x29, 0xac, 0xf7,
	0xf8, 0xb9, 0x0b, 0x99, 0x5f, 0xab, 0x71, 0x9d, 0x6d, 0x15, 0xe9, 0x7b, 0x50, 0xda, 0xf5, 0x4f,
	0xbc, 0xb4, 0xe7, 0x19, 0x11, 0x42, 0x0e, 0x26, 0xec, 0xba, 0x2f, 0x7e, 0xc6, 0xe8, 0xe0, 0xfd,
	0xa4, 0xad, 0x38, 0x6c, 0x93, 0x97, 0xe8, 0xb1, 0x9f, 0xa6, 0x16, 0xa4, 0x83, 0x34, 0x3a, 0x8d,
	0x2e, 0x0e, 0x16, 0x8f, 0xfc, 0x9f, 0xb7, 0xa4, 0x83, 0x64, 0x86, 0xf6, 0x19, 0x6d, 0x79, 0x93,
	0xc6, 0xa7, 0xd1, 0xc5, 0xe1, 0xec, 0xc4, 0x6d, 0x83, 0xfd, 0x00, 0xf8, 0xd3, 0xa0, 0xb8, 0x60,
	0x15, 0x69, 0x47, 0x58, 0x58, 0x34, 0xf9, 0x8c, 0x9e, 0x10, 0x05, 0xa4, 0x96, 0xeb, 0x9a, 0x8b,
	0x01, 0x14, 0xe8, 0x21, 0xdd, 0x33, 0xf1, 0x57, 0x78, 0xd7, 0xb9, 0xd8, 0xa5, 0xb1, 0x99, 0xee,
	0xa3, 0xdb, 0x79, 0x71, 0x34, 0x69, 0xee, 0xd6, 0x37, 0x4e, 0x92, 0xd4, 0xe8, 0xa9, 0x3f, 0x8f,
	0x49, 0xde, 0x2b, 0xd0, 0x20, 0x28, 0xa4, 0x0f, 0xee, 0x23, 0x4f, 0xbc, 0xea, 0x6e, 0x3d, 0x77,
	0xa2, 0xe4, 0x1d, 0x3a, 0xec, 0x09, 0x83, 0x5a, 0x8c, 0xdd, 0x0a, 0x54, 0xba, 0x6f, 0xbc, 0xcf,
	0xfe, 0xda, 0xf9, 0x46, 0x0c, 0x6f, 0x5e, 0xdb, 0x95, 0xd1, 0xc4, 0xdf, 0x1a, 0xfc, 0xea, 0x77,
	0x84, 0xce, 0xa9, 0xec, 0xf0, 0x7f, 0xef, 0xfe, 0xea, 0x28, 0xdc, 0xc2, 0x7c, 0x72, 0xce, 0xa3,
	0x2f, 0x1f, 0x5c, 0x88, 0xc9, 0x96, 0x08, 0x86, 0xa5, 0x62, 0x19, 0x03, 0x61, 0xbe, 0xe8, 0x5f,
	0x47, 0xcf, 0xf5, 0x3f, 0xde, 0xe0, 0xdb, 0x50, 0x7d, 0x8b, 0xf7, 0xae, 0xcb, 0xf2, 0x7b, 0x7c,
	0x76, 0x6d, 0x95, 0x65, 0xa3, 0xb1, 0x2d, 0xa7, 0xaa, 0xca, 0xf1, 0xc2, 0x93, 0x3f, 0x3c, 0xb3,
	0x2c, 0x1b, 0xbd, 0x0c, 0xcc, 0xb2, 0xca, 0x97, 0x81, 0xf9, 0x15, 0x9f, 0xdb, 0x46, 0x51, 0x94,
	0x8d, 0x2e, 0x8a, 0x40, 0x15, 0x45, 0x95, 0x17, 0x45, 0xe0, 0x56, 0x0f, 0xcd, 0xb0, 0x97, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xc6, 0x81, 0x37, 0x2f, 0x03, 0x00, 0x00,
}
