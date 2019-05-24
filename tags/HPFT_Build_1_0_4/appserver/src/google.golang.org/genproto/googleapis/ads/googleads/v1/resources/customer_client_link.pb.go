// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/customer_client_link.proto

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

// Represents customer client link relationship.
type CustomerClientLink struct {
	// Name of the resource.
	// CustomerClientLink resource names have the form:
	//
	// `customers/{customer_id}/customerClientLinks/{client_customer_id}~{manager_link_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The client customer linked to this customer.
	ClientCustomer *wrappers.StringValue `protobuf:"bytes,3,opt,name=client_customer,json=clientCustomer,proto3" json:"client_customer,omitempty"`
	// This is uniquely identifies a customer client link. Read only.
	ManagerLinkId *wrappers.Int64Value `protobuf:"bytes,4,opt,name=manager_link_id,json=managerLinkId,proto3" json:"manager_link_id,omitempty"`
	// This is the status of the link between client and manager.
	Status enums.ManagerLinkStatusEnum_ManagerLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.ManagerLinkStatusEnum_ManagerLinkStatus" json:"status,omitempty"`
	// The visibility of the link. Users can choose whether or not to see hidden
	// links in the AdWords UI.
	// Default value is false
	Hidden               *wrappers.BoolValue `protobuf:"bytes,6,opt,name=hidden,proto3" json:"hidden,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CustomerClientLink) Reset()         { *m = CustomerClientLink{} }
func (m *CustomerClientLink) String() string { return proto.CompactTextString(m) }
func (*CustomerClientLink) ProtoMessage()    {}
func (*CustomerClientLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dba9317d3293f16, []int{0}
}

func (m *CustomerClientLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerClientLink.Unmarshal(m, b)
}
func (m *CustomerClientLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerClientLink.Marshal(b, m, deterministic)
}
func (m *CustomerClientLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerClientLink.Merge(m, src)
}
func (m *CustomerClientLink) XXX_Size() int {
	return xxx_messageInfo_CustomerClientLink.Size(m)
}
func (m *CustomerClientLink) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerClientLink.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerClientLink proto.InternalMessageInfo

func (m *CustomerClientLink) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerClientLink) GetClientCustomer() *wrappers.StringValue {
	if m != nil {
		return m.ClientCustomer
	}
	return nil
}

func (m *CustomerClientLink) GetManagerLinkId() *wrappers.Int64Value {
	if m != nil {
		return m.ManagerLinkId
	}
	return nil
}

func (m *CustomerClientLink) GetStatus() enums.ManagerLinkStatusEnum_ManagerLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.ManagerLinkStatusEnum_UNSPECIFIED
}

func (m *CustomerClientLink) GetHidden() *wrappers.BoolValue {
	if m != nil {
		return m.Hidden
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomerClientLink)(nil), "google.ads.googleads.v1.resources.CustomerClientLink")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/customer_client_link.proto", fileDescriptor_1dba9317d3293f16)
}

var fileDescriptor_1dba9317d3293f16 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xdd, 0x6a, 0x14, 0x31,
	0x18, 0x65, 0xa6, 0xba, 0x60, 0xb4, 0x2d, 0xcc, 0x8d, 0xc3, 0x5a, 0x64, 0xab, 0x14, 0xf6, 0x2a,
	0x61, 0x56, 0x51, 0x88, 0xde, 0xcc, 0x2e, 0xb5, 0x54, 0x54, 0xca, 0x16, 0xf6, 0x42, 0x16, 0x97,
	0x74, 0x13, 0x63, 0xe8, 0x4c, 0x32, 0x24, 0x99, 0xfa, 0x02, 0x3e, 0x89, 0x97, 0x3e, 0x8a, 0x6f,
	0xe0, 0x2b, 0xf8, 0x14, 0x32, 0xf9, 0x19, 0x91, 0x61, 0xed, 0xdd, 0x49, 0xbe, 0x73, 0xce, 0x77,
	0xbe, 0x2f, 0x01, 0xaf, 0xb9, 0x52, 0xbc, 0x62, 0x88, 0x50, 0x83, 0x3c, 0xec, 0xd0, 0x4d, 0x81,
	0x34, 0x33, 0xaa, 0xd5, 0x5b, 0x66, 0xd0, 0xb6, 0x35, 0x56, 0xd5, 0x4c, 0x6f, 0xb6, 0x95, 0x60,
	0xd2, 0x6e, 0x2a, 0x21, 0xaf, 0x61, 0xa3, 0x95, 0x55, 0xd9, 0xb1, 0x97, 0x40, 0x42, 0x0d, 0xec,
	0xd5, 0xf0, 0xa6, 0x80, 0xbd, 0x7a, 0xfc, 0x72, 0x57, 0x03, 0x26, 0xdb, 0xda, 0xa0, 0x9a, 0x48,
	0xc2, 0x99, 0x76, 0xa6, 0x1b, 0x63, 0x89, 0x6d, 0x8d, 0xf7, 0x1e, 0x3f, 0x0e, 0x42, 0x77, 0xba,
	0x6a, 0x3f, 0xa3, 0xaf, 0x9a, 0x34, 0x0d, 0xd3, 0xb1, 0x7e, 0x14, 0x8d, 0x1b, 0x81, 0x88, 0x94,
	0xca, 0x12, 0x2b, 0x94, 0x0c, 0xd5, 0x27, 0xbf, 0x52, 0x90, 0x2d, 0x42, 0xf0, 0x85, 0xcb, 0xfd,
	0x4e, 0xc8, 0xeb, 0xec, 0x29, 0xd8, 0x8f, 0xd1, 0x36, 0x92, 0xd4, 0x2c, 0x4f, 0x26, 0xc9, 0xf4,
	0xde, 0xf2, 0x41, 0xbc, 0xfc, 0x40, 0x6a, 0x96, 0x9d, 0x82, 0xc3, 0x30, 0x6a, 0x1c, 0x3d, 0xdf,
	0x9b, 0x24, 0xd3, 0xfb, 0xb3, 0xa3, 0x30, 0x24, 0x8c, 0x99, 0xe0, 0xa5, 0xd5, 0x42, 0xf2, 0x15,
	0xa9, 0x5a, 0xb6, 0x3c, 0xf0, 0xa2, 0xd8, 0x35, 0x5b, 0x80, 0xc3, 0x7f, 0xa6, 0x13, 0x34, 0xbf,
	0xe3, 0x6c, 0x1e, 0x0d, 0x6c, 0xce, 0xa5, 0x7d, 0xf1, 0xdc, 0xbb, 0xec, 0x07, 0x4d, 0x17, 0xf7,
	0x9c, 0x66, 0x9f, 0xc0, 0xc8, 0x6f, 0x25, 0xbf, 0x3b, 0x49, 0xa6, 0x07, 0xb3, 0x37, 0x70, 0xd7,
	0xca, 0xdd, 0x3e, 0xe1, 0xfb, 0xbf, 0xea, 0x4b, 0xa7, 0x3b, 0x95, 0x6d, 0x3d, 0xbc, 0x5d, 0x06,
	0xd7, 0x6c, 0x06, 0x46, 0x5f, 0x04, 0xa5, 0x4c, 0xe6, 0x23, 0x97, 0x6d, 0x3c, 0xc8, 0x36, 0x57,
	0xaa, 0xf2, 0xd1, 0x02, 0x73, 0xfe, 0x2d, 0x05, 0x27, 0x5b, 0x55, 0xc3, 0x5b, 0x1f, 0x7f, 0xfe,
	0x70, 0xf8, 0x04, 0x17, 0x9d, 0xef, 0x45, 0xf2, 0xf1, 0x6d, 0x50, 0x73, 0x55, 0x11, 0xc9, 0xa1,
	0xd2, 0x1c, 0x71, 0x26, 0x5d, 0xd7, 0xf8, 0x4f, 0x1a, 0x61, 0xfe, 0xf3, 0x2f, 0x5f, 0xf5, 0xe8,
	0x7b, 0xba, 0x77, 0x56, 0x96, 0x3f, 0xd2, 0xe3, 0x33, 0x6f, 0x59, 0x52, 0x03, 0x3d, 0xec, 0xd0,
	0xaa, 0x80, 0xcb, 0xc8, 0xfc, 0x19, 0x39, 0xeb, 0x92, 0x9a, 0x75, 0xcf, 0x59, 0xaf, 0x8a, 0x75,
	0xcf, 0xf9, 0x9d, 0x9e, 0xf8, 0x02, 0xc6, 0x25, 0x35, 0x18, 0xf7, 0x2c, 0x8c, 0x57, 0x05, 0xc6,
	0x3d, 0xef, 0x6a, 0xe4, 0xc2, 0x3e, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x73, 0x1a, 0x09, 0xc6,
	0x43, 0x03, 0x00, 0x00,
}
