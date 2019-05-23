// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/customer_manager_link.proto

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

// Represents customer-manager link relationship.
type CustomerManagerLink struct {
	// Name of the resource.
	// CustomerManagerLink resource names have the form:
	//
	// `customers/{customer_id}/customerManagerLinks/{manager_customer_id}~{manager_link_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The manager customer linked to the customer.
	ManagerCustomer *wrappers.StringValue `protobuf:"bytes,3,opt,name=manager_customer,json=managerCustomer,proto3" json:"manager_customer,omitempty"`
	// ID of the customer-manager link. This field is read only.
	ManagerLinkId *wrappers.Int64Value `protobuf:"bytes,4,opt,name=manager_link_id,json=managerLinkId,proto3" json:"manager_link_id,omitempty"`
	// Status of the link between the customer and the manager.
	Status               enums.ManagerLinkStatusEnum_ManagerLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.ManagerLinkStatusEnum_ManagerLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *CustomerManagerLink) Reset()         { *m = CustomerManagerLink{} }
func (m *CustomerManagerLink) String() string { return proto.CompactTextString(m) }
func (*CustomerManagerLink) ProtoMessage()    {}
func (*CustomerManagerLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e3318e064b05720, []int{0}
}

func (m *CustomerManagerLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerManagerLink.Unmarshal(m, b)
}
func (m *CustomerManagerLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerManagerLink.Marshal(b, m, deterministic)
}
func (m *CustomerManagerLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerManagerLink.Merge(m, src)
}
func (m *CustomerManagerLink) XXX_Size() int {
	return xxx_messageInfo_CustomerManagerLink.Size(m)
}
func (m *CustomerManagerLink) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerManagerLink.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerManagerLink proto.InternalMessageInfo

func (m *CustomerManagerLink) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerManagerLink) GetManagerCustomer() *wrappers.StringValue {
	if m != nil {
		return m.ManagerCustomer
	}
	return nil
}

func (m *CustomerManagerLink) GetManagerLinkId() *wrappers.Int64Value {
	if m != nil {
		return m.ManagerLinkId
	}
	return nil
}

func (m *CustomerManagerLink) GetStatus() enums.ManagerLinkStatusEnum_ManagerLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.ManagerLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CustomerManagerLink)(nil), "google.ads.googleads.v1.resources.CustomerManagerLink")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/customer_manager_link.proto", fileDescriptor_7e3318e064b05720)
}

var fileDescriptor_7e3318e064b05720 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xd1, 0xca, 0xd3, 0x30,
	0x18, 0xa5, 0xfd, 0xf5, 0x07, 0xab, 0x73, 0x52, 0x6f, 0xca, 0x1c, 0xb2, 0x29, 0x83, 0x5d, 0x25,
	0x74, 0x8a, 0x42, 0xc4, 0x8b, 0x6e, 0xe8, 0x98, 0xa8, 0x8c, 0x0e, 0x7a, 0x21, 0xc5, 0x92, 0xad,
	0x31, 0x94, 0x35, 0x49, 0x49, 0xd2, 0xf9, 0x04, 0xbe, 0x88, 0x57, 0xe2, 0xa3, 0xf8, 0x28, 0x3e,
	0x85, 0xac, 0x69, 0xea, 0x64, 0x4e, 0xef, 0x4e, 0x93, 0x73, 0xce, 0x77, 0xce, 0x97, 0x7a, 0x2f,
	0xa9, 0x10, 0xb4, 0x24, 0x10, 0xe7, 0x0a, 0x1a, 0x78, 0x44, 0x87, 0x10, 0x4a, 0xa2, 0x44, 0x2d,
	0x77, 0x44, 0xc1, 0x5d, 0xad, 0xb4, 0x60, 0x44, 0x66, 0x0c, 0x73, 0x4c, 0x89, 0xcc, 0xca, 0x82,
	0xef, 0x41, 0x25, 0x85, 0x16, 0xfe, 0xd8, 0x68, 0x00, 0xce, 0x15, 0xe8, 0xe4, 0xe0, 0x10, 0x82,
	0x4e, 0x3e, 0x78, 0x7e, 0x69, 0x02, 0xe1, 0x35, 0x53, 0xf0, 0xd4, 0x34, 0x53, 0x1a, 0xeb, 0x5a,
	0x19, 0xef, 0xc1, 0xc3, 0x56, 0xd8, 0x7c, 0x6d, 0xeb, 0x4f, 0xf0, 0xb3, 0xc4, 0x55, 0x45, 0xa4,
	0xbd, 0x1f, 0x5a, 0xe3, 0xaa, 0x80, 0x98, 0x73, 0xa1, 0xb1, 0x2e, 0x04, 0x6f, 0x6f, 0x1f, 0x7d,
	0x73, 0xbd, 0xfb, 0x8b, 0x36, 0xf9, 0x3b, 0x33, 0xe3, 0x6d, 0xc1, 0xf7, 0xfe, 0x63, 0xaf, 0x67,
	0xb3, 0x65, 0x1c, 0x33, 0x12, 0x38, 0x23, 0x67, 0x7a, 0x2b, 0xbe, 0x63, 0x0f, 0xdf, 0x63, 0x46,
	0xfc, 0xa5, 0x77, 0xcf, 0xe6, 0xb2, 0xed, 0x83, 0xab, 0x91, 0x33, 0xbd, 0x3d, 0x1b, 0xb6, 0x35,
	0x81, 0x4d, 0x05, 0x36, 0x5a, 0x16, 0x9c, 0x26, 0xb8, 0xac, 0x49, 0xdc, 0x6f, 0x55, 0x76, 0xb0,
	0xbf, 0xf0, 0xfa, 0x7f, 0x14, 0x2c, 0xf2, 0xe0, 0x46, 0xe3, 0xf3, 0xe0, 0xcc, 0x67, 0xc5, 0xf5,
	0xb3, 0xa7, 0xc6, 0xa6, 0xc7, 0x7e, 0x07, 0x5e, 0xe5, 0xfe, 0x47, 0xef, 0xda, 0x2c, 0x26, 0xb8,
	0x39, 0x72, 0xa6, 0x77, 0x67, 0xaf, 0xc1, 0xa5, 0xad, 0x37, 0x2b, 0x05, 0x27, 0x75, 0x37, 0x8d,
	0xee, 0x15, 0xaf, 0xd9, 0xf9, 0x69, 0xdc, 0xba, 0xce, 0xbf, 0xb8, 0xde, 0x64, 0x27, 0x18, 0xf8,
	0xef, 0x5b, 0xce, 0x83, 0xbf, 0x6c, 0x74, 0x7d, 0x2c, 0xb0, 0x76, 0x3e, 0xbc, 0x69, 0xe5, 0x54,
	0x94, 0x98, 0x53, 0x20, 0x24, 0x85, 0x94, 0xf0, 0xa6, 0x9e, 0x7d, 0xf7, 0xaa, 0x50, 0xff, 0xf8,
	0xd1, 0x5e, 0x74, 0xe8, 0xab, 0x7b, 0xb5, 0x8c, 0xa2, 0xef, 0xee, 0x78, 0x69, 0x2c, 0xa3, 0x5c,
	0x01, 0x03, 0x8f, 0x28, 0x09, 0x41, 0x6c, 0x99, 0x3f, 0x2c, 0x27, 0x8d, 0x72, 0x95, 0x76, 0x9c,
	0x34, 0x09, 0xd3, 0x8e, 0xf3, 0xd3, 0x9d, 0x98, 0x0b, 0x84, 0xa2, 0x5c, 0x21, 0xd4, 0xb1, 0x10,
	0x4a, 0x42, 0x84, 0x3a, 0xde, 0xf6, 0xba, 0x09, 0xfb, 0xe4, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x60, 0xab, 0x53, 0xed, 0x14, 0x03, 0x00, 0x00,
}
