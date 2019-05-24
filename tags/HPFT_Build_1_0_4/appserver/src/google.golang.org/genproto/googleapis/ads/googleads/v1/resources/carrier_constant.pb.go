// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/carrier_constant.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A carrier criterion that can be used in campaign targeting.
type CarrierConstant struct {
	// The resource name of the carrier criterion.
	// Carrier criterion resource names have the form:
	//
	// `carrierConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the carrier criterion.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The full name of the carrier in English.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The country code of the country where the carrier is located, e.g., "AR",
	// "FR", etc.
	CountryCode          *wrappers.StringValue `protobuf:"bytes,4,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CarrierConstant) Reset()         { *m = CarrierConstant{} }
func (m *CarrierConstant) String() string { return proto.CompactTextString(m) }
func (*CarrierConstant) ProtoMessage()    {}
func (*CarrierConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b8b1da6512447a4, []int{0}
}

func (m *CarrierConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarrierConstant.Unmarshal(m, b)
}
func (m *CarrierConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarrierConstant.Marshal(b, m, deterministic)
}
func (m *CarrierConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarrierConstant.Merge(m, src)
}
func (m *CarrierConstant) XXX_Size() int {
	return xxx_messageInfo_CarrierConstant.Size(m)
}
func (m *CarrierConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_CarrierConstant.DiscardUnknown(m)
}

var xxx_messageInfo_CarrierConstant proto.InternalMessageInfo

func (m *CarrierConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CarrierConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CarrierConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CarrierConstant) GetCountryCode() *wrappers.StringValue {
	if m != nil {
		return m.CountryCode
	}
	return nil
}

func init() {
	proto.RegisterType((*CarrierConstant)(nil), "google.ads.googleads.v1.resources.CarrierConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/carrier_constant.proto", fileDescriptor_6b8b1da6512447a4)
}

var fileDescriptor_6b8b1da6512447a4 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x49, 0x5a, 0x2e, 0xdc, 0xb4, 0x97, 0x0b, 0xc1, 0x45, 0xa9, 0x45, 0x5a, 0xa5, 0x50,
	0x10, 0x26, 0x46, 0x45, 0x64, 0x5c, 0x48, 0xda, 0x45, 0xd1, 0x85, 0x94, 0x0a, 0x59, 0x48, 0xa0,
	0x4c, 0x33, 0x63, 0x08, 0xb4, 0x73, 0xc2, 0xcc, 0xa4, 0xe2, 0xd2, 0x57, 0x71, 0xe9, 0xa3, 0xf8,
	0x00, 0x3e, 0x84, 0x4f, 0x21, 0xcd, 0x64, 0x66, 0xa1, 0xa0, 0xee, 0x7e, 0x72, 0xbe, 0xff, 0xff,
	0x4f, 0xe6, 0x78, 0xe7, 0x19, 0x40, 0xb6, 0x62, 0x01, 0xa1, 0x32, 0xd0, 0x72, 0xab, 0x36, 0x61,
	0x20, 0x98, 0x84, 0x52, 0xa4, 0x4c, 0x06, 0x29, 0x11, 0x22, 0x67, 0x62, 0x91, 0x02, 0x97, 0x8a,
	0x70, 0x85, 0x0a, 0x01, 0x0a, 0xfc, 0x81, 0xc6, 0x11, 0xa1, 0x12, 0x59, 0x27, 0xda, 0x84, 0xc8,
	0x3a, 0xbb, 0x7b, 0x75, 0x78, 0x65, 0x58, 0x96, 0xf7, 0xc1, 0x83, 0x20, 0x45, 0xc1, 0x84, 0xd4,
	0x11, 0xdd, 0x9e, 0x29, 0x2f, 0xf2, 0x80, 0x70, 0x0e, 0x8a, 0xa8, 0x1c, 0x78, 0x3d, 0xdd, 0x7f,
	0x73, 0xbc, 0xff, 0x13, 0xdd, 0x3d, 0xa9, 0xab, 0xfd, 0x03, 0xef, 0x9f, 0x89, 0x5f, 0x70, 0xb2,
	0x66, 0x1d, 0xa7, 0xef, 0x8c, 0xfe, 0xce, 0xdb, 0xe6, 0xe3, 0x0d, 0x59, 0x33, 0xff, 0xd0, 0x73,
	0x73, 0xda, 0x71, 0xfb, 0xce, 0xa8, 0x75, 0xbc, 0x5b, 0xef, 0x86, 0xcc, 0x0e, 0xe8, 0x8a, 0xab,
	0xb3, 0xd3, 0x98, 0xac, 0x4a, 0x36, 0x77, 0x73, 0xea, 0x1f, 0x79, 0xcd, 0x2a, 0xa8, 0x51, 0xe1,
	0xbd, 0x2f, 0xf8, 0xad, 0x12, 0x39, 0xcf, 0x34, 0x5f, 0x91, 0xfe, 0xa5, 0xd7, 0x4e, 0xa1, 0xe4,
	0x4a, 0x3c, 0x2e, 0x52, 0xa0, 0xac, 0xd3, 0xfc, 0x85, 0xb3, 0x55, 0x3b, 0x26, 0x40, 0xd9, 0xf8,
	0xc9, 0xf5, 0x86, 0x29, 0xac, 0xd1, 0x8f, 0x0f, 0x38, 0xde, 0xf9, 0xf4, 0xff, 0xb3, 0x6d, 0xf6,
	0xcc, 0xb9, 0xbb, 0xae, 0xad, 0x19, 0xac, 0x08, 0xcf, 0x10, 0x88, 0x2c, 0xc8, 0x18, 0xaf, 0x9a,
	0xcd, 0x15, 0x8b, 0x5c, 0x7e, 0x73, 0xd4, 0x0b, 0xab, 0x9e, 0xdd, 0xc6, 0x34, 0x8a, 0x5e, 0xdc,
	0xc1, 0x54, 0x47, 0x46, 0x54, 0x22, 0x2d, 0xb7, 0x2a, 0x0e, 0xd1, 0xdc, 0x90, 0xaf, 0x86, 0x49,
	0x22, 0x2a, 0x13, 0xcb, 0x24, 0x71, 0x98, 0x58, 0xe6, 0xdd, 0x1d, 0xea, 0x01, 0xc6, 0x11, 0x95,
	0x18, 0x5b, 0x0a, 0xe3, 0x38, 0xc4, 0xd8, 0x72, 0xcb, 0x3f, 0xd5, 0xb2, 0x27, 0x1f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd2, 0xda, 0x43, 0x48, 0x80, 0x02, 0x00, 0x00,
}
