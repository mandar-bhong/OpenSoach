// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/customer_label.proto

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

// Represents a relationship between a customer and a label. This customer may
// not have access to all the labels attached to it. Additional CustomerLabels
// may be returned by increasing permissions with login-customer-id.
type CustomerLabel struct {
	// Name of the resource.
	// Customer label resource names have the form:
	// `customers/{customer_id}/customerLabels/{label_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The resource name of the customer to which the label is attached.
	// Read only.
	Customer *wrappers.StringValue `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
	// The resource name of the label assigned to the customer.
	//
	// Note: the Customer ID portion of the label resource name is not
	// validated when creating a new CustomerLabel.
	Label                *wrappers.StringValue `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CustomerLabel) Reset()         { *m = CustomerLabel{} }
func (m *CustomerLabel) String() string { return proto.CompactTextString(m) }
func (*CustomerLabel) ProtoMessage()    {}
func (*CustomerLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c963ece951af59, []int{0}
}

func (m *CustomerLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerLabel.Unmarshal(m, b)
}
func (m *CustomerLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerLabel.Marshal(b, m, deterministic)
}
func (m *CustomerLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerLabel.Merge(m, src)
}
func (m *CustomerLabel) XXX_Size() int {
	return xxx_messageInfo_CustomerLabel.Size(m)
}
func (m *CustomerLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerLabel.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerLabel proto.InternalMessageInfo

func (m *CustomerLabel) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerLabel) GetCustomer() *wrappers.StringValue {
	if m != nil {
		return m.Customer
	}
	return nil
}

func (m *CustomerLabel) GetLabel() *wrappers.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomerLabel)(nil), "google.ads.googleads.v1.resources.CustomerLabel")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/customer_label.proto", fileDescriptor_85c963ece951af59)
}

var fileDescriptor_85c963ece951af59 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0xc3, 0x30,
	0x00, 0xc6, 0x49, 0x87, 0xa2, 0xd1, 0x5d, 0x7a, 0x1a, 0x63, 0xc8, 0xa6, 0x0c, 0x76, 0x4a, 0xe8,
	0x04, 0x91, 0x78, 0xea, 0x3c, 0x0c, 0x44, 0x64, 0x4c, 0xe8, 0x41, 0x0a, 0x23, 0x6d, 0x63, 0x28,
	0xb4, 0x49, 0x49, 0xda, 0xf9, 0x3e, 0x3b, 0xfa, 0x28, 0x3e, 0x8a, 0x2f, 0xa1, 0xb4, 0x69, 0x02,
	0x5e, 0xd4, 0xdb, 0x47, 0xf3, 0xfb, 0xfe, 0xa4, 0x81, 0x37, 0x5c, 0x4a, 0x5e, 0x30, 0x4c, 0x33,
	0x8d, 0x8d, 0x6c, 0xd5, 0x3e, 0xc0, 0x8a, 0x69, 0xd9, 0xa8, 0x94, 0x69, 0x9c, 0x36, 0xba, 0x96,
	0x25, 0x53, 0xbb, 0x82, 0x26, 0xac, 0x40, 0x95, 0x92, 0xb5, 0xf4, 0x67, 0x06, 0x46, 0x34, 0xd3,
	0xc8, 0xf9, 0xd0, 0x3e, 0x40, 0xce, 0x37, 0xbe, 0xe8, 0xa3, 0x3b, 0x43, 0xd2, 0xbc, 0xe2, 0x37,
	0x45, 0xab, 0x8a, 0x29, 0x6d, 0x22, 0xc6, 0x13, 0x5b, 0x5d, 0xe5, 0x98, 0x0a, 0x21, 0x6b, 0x5a,
	0xe7, 0x52, 0xf4, 0xa7, 0x97, 0x07, 0x00, 0x87, 0xf7, 0x7d, 0xf3, 0x63, 0x5b, 0xec, 0x5f, 0xc1,
	0xa1, 0x0d, 0xdf, 0x09, 0x5a, 0xb2, 0x11, 0x98, 0x82, 0xc5, 0xe9, 0xf6, 0xdc, 0x7e, 0x7c, 0xa2,
	0x25, 0xf3, 0x6f, 0xe1, 0x89, 0xdd, 0x3b, 0xf2, 0xa6, 0x60, 0x71, 0xb6, 0x9c, 0xf4, 0xfb, 0x90,
	0xdd, 0x81, 0x9e, 0x6b, 0x95, 0x0b, 0x1e, 0xd1, 0xa2, 0x61, 0x5b, 0x47, 0xfb, 0x4b, 0x78, 0xd4,
	0x5d, 0x70, 0x34, 0xf8, 0x87, 0xcd, 0xa0, 0xab, 0x2f, 0x00, 0xe7, 0xa9, 0x2c, 0xd1, 0x9f, 0x3f,
	0x63, 0xe5, 0xff, 0xb8, 0xcb, 0xa6, 0xcd, 0xdc, 0x80, 0x97, 0x87, 0xde, 0xc8, 0x65, 0x41, 0x05,
	0x47, 0x52, 0x71, 0xcc, 0x99, 0xe8, 0x1a, 0xed, 0x6b, 0x54, 0xb9, 0xfe, 0xe5, 0x71, 0xee, 0x9c,
	0x3a, 0x78, 0x83, 0x75, 0x18, 0xbe, 0x7b, 0xb3, 0xb5, 0x89, 0x0c, 0x33, 0x8d, 0x8c, 0x6c, 0x55,
	0x14, 0xa0, 0xad, 0x25, 0x3f, 0x2c, 0x13, 0x87, 0x99, 0x8e, 0x1d, 0x13, 0x47, 0x41, 0xec, 0x98,
	0x4f, 0x6f, 0x6e, 0x0e, 0x08, 0x09, 0x33, 0x4d, 0x88, 0xa3, 0x08, 0x89, 0x02, 0x42, 0x1c, 0x97,
	0x1c, 0x77, 0x63, 0xaf, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x55, 0x24, 0xa1, 0x48, 0x02,
	0x00, 0x00,
}
