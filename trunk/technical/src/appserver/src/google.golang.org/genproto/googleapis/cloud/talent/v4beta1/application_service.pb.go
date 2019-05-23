// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/application_service.proto

package talent

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
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

// The Request of the CreateApplication method.
type CreateApplicationRequest struct {
	// Required.
	//
	// Resource name of the profile under which the application is created.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/profiles/{profile_id}", for
	// example, "projects/test-project/tenants/test-tenant/profiles/test-profile".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required.
	//
	// The application to be created.
	Application          *Application `protobuf:"bytes,2,opt,name=application,proto3" json:"application,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateApplicationRequest) Reset()         { *m = CreateApplicationRequest{} }
func (m *CreateApplicationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateApplicationRequest) ProtoMessage()    {}
func (*CreateApplicationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d34b1b7148314cc, []int{0}
}

func (m *CreateApplicationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateApplicationRequest.Unmarshal(m, b)
}
func (m *CreateApplicationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateApplicationRequest.Marshal(b, m, deterministic)
}
func (m *CreateApplicationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateApplicationRequest.Merge(m, src)
}
func (m *CreateApplicationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateApplicationRequest.Size(m)
}
func (m *CreateApplicationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateApplicationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateApplicationRequest proto.InternalMessageInfo

func (m *CreateApplicationRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateApplicationRequest) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

// Request for getting a application by name.
type GetApplicationRequest struct {
	// Required.
	//
	// The resource name of the application to be retrieved.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/profiles/{profile_id}/applications/{application_id}",
	// for example,
	// "projects/test-project/tenants/test-tenant/profiles/test-profile/applications/test-application".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetApplicationRequest) Reset()         { *m = GetApplicationRequest{} }
func (m *GetApplicationRequest) String() string { return proto.CompactTextString(m) }
func (*GetApplicationRequest) ProtoMessage()    {}
func (*GetApplicationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d34b1b7148314cc, []int{1}
}

func (m *GetApplicationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetApplicationRequest.Unmarshal(m, b)
}
func (m *GetApplicationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetApplicationRequest.Marshal(b, m, deterministic)
}
func (m *GetApplicationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetApplicationRequest.Merge(m, src)
}
func (m *GetApplicationRequest) XXX_Size() int {
	return xxx_messageInfo_GetApplicationRequest.Size(m)
}
func (m *GetApplicationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetApplicationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetApplicationRequest proto.InternalMessageInfo

func (m *GetApplicationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request for updating a specified application.
type UpdateApplicationRequest struct {
	// Required.
	//
	// The application resource to replace the current resource in the system.
	Application *Application `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	// Optional but strongly recommended for the best service
	// experience.
	//
	// If [update_mask][google.cloud.talent.v4beta1.UpdateApplicationRequest.update_mask] is provided, only the specified fields in
	// [application][google.cloud.talent.v4beta1.UpdateApplicationRequest.application] are updated. Otherwise all the fields are updated.
	//
	// A field mask to specify the application fields to be updated. Only
	// top level fields of [Application][google.cloud.talent.v4beta1.Application] are supported.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateApplicationRequest) Reset()         { *m = UpdateApplicationRequest{} }
func (m *UpdateApplicationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateApplicationRequest) ProtoMessage()    {}
func (*UpdateApplicationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d34b1b7148314cc, []int{2}
}

func (m *UpdateApplicationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateApplicationRequest.Unmarshal(m, b)
}
func (m *UpdateApplicationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateApplicationRequest.Marshal(b, m, deterministic)
}
func (m *UpdateApplicationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateApplicationRequest.Merge(m, src)
}
func (m *UpdateApplicationRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateApplicationRequest.Size(m)
}
func (m *UpdateApplicationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateApplicationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateApplicationRequest proto.InternalMessageInfo

func (m *UpdateApplicationRequest) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *UpdateApplicationRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Request to delete a application.
type DeleteApplicationRequest struct {
	// Required.
	//
	// The resource name of the application to be deleted.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/profiles/{profile_id}/applications/{application_id}",
	// for example,
	// "projects/test-project/tenants/test-tenant/profiles/test-profile/applications/test-application".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteApplicationRequest) Reset()         { *m = DeleteApplicationRequest{} }
func (m *DeleteApplicationRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteApplicationRequest) ProtoMessage()    {}
func (*DeleteApplicationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d34b1b7148314cc, []int{3}
}

func (m *DeleteApplicationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteApplicationRequest.Unmarshal(m, b)
}
func (m *DeleteApplicationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteApplicationRequest.Marshal(b, m, deterministic)
}
func (m *DeleteApplicationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteApplicationRequest.Merge(m, src)
}
func (m *DeleteApplicationRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteApplicationRequest.Size(m)
}
func (m *DeleteApplicationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteApplicationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteApplicationRequest proto.InternalMessageInfo

func (m *DeleteApplicationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// List applications for which the client has ACL visibility.
type ListApplicationsRequest struct {
	// Required.
	//
	// Resource name of the profile under which the application is created.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/profiles/{profile_id}", for
	// example, "projects/test-project/tenants/test-tenant/profiles/test-profile".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional.
	//
	// The starting indicator from which to return results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional.
	//
	// The maximum number of applications to be returned, at most 100.
	// Default is 100 if a non-positive number is provided.
	PageSize             int32    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListApplicationsRequest) Reset()         { *m = ListApplicationsRequest{} }
func (m *ListApplicationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListApplicationsRequest) ProtoMessage()    {}
func (*ListApplicationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d34b1b7148314cc, []int{4}
}

func (m *ListApplicationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListApplicationsRequest.Unmarshal(m, b)
}
func (m *ListApplicationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListApplicationsRequest.Marshal(b, m, deterministic)
}
func (m *ListApplicationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListApplicationsRequest.Merge(m, src)
}
func (m *ListApplicationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListApplicationsRequest.Size(m)
}
func (m *ListApplicationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListApplicationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListApplicationsRequest proto.InternalMessageInfo

func (m *ListApplicationsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListApplicationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListApplicationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// Output only.
//
// The List applications response object.
type ListApplicationsResponse struct {
	// Applications for the current client.
	Applications []*Application `protobuf:"bytes,1,rep,name=applications,proto3" json:"applications,omitempty"`
	// A token to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Additional information for the API invocation, such as the request
	// tracking id.
	Metadata             *ResponseMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListApplicationsResponse) Reset()         { *m = ListApplicationsResponse{} }
func (m *ListApplicationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListApplicationsResponse) ProtoMessage()    {}
func (*ListApplicationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d34b1b7148314cc, []int{5}
}

func (m *ListApplicationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListApplicationsResponse.Unmarshal(m, b)
}
func (m *ListApplicationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListApplicationsResponse.Marshal(b, m, deterministic)
}
func (m *ListApplicationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListApplicationsResponse.Merge(m, src)
}
func (m *ListApplicationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListApplicationsResponse.Size(m)
}
func (m *ListApplicationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListApplicationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListApplicationsResponse proto.InternalMessageInfo

func (m *ListApplicationsResponse) GetApplications() []*Application {
	if m != nil {
		return m.Applications
	}
	return nil
}

func (m *ListApplicationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListApplicationsResponse) GetMetadata() *ResponseMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateApplicationRequest)(nil), "google.cloud.talent.v4beta1.CreateApplicationRequest")
	proto.RegisterType((*GetApplicationRequest)(nil), "google.cloud.talent.v4beta1.GetApplicationRequest")
	proto.RegisterType((*UpdateApplicationRequest)(nil), "google.cloud.talent.v4beta1.UpdateApplicationRequest")
	proto.RegisterType((*DeleteApplicationRequest)(nil), "google.cloud.talent.v4beta1.DeleteApplicationRequest")
	proto.RegisterType((*ListApplicationsRequest)(nil), "google.cloud.talent.v4beta1.ListApplicationsRequest")
	proto.RegisterType((*ListApplicationsResponse)(nil), "google.cloud.talent.v4beta1.ListApplicationsResponse")
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/application_service.proto", fileDescriptor_5d34b1b7148314cc)
}

var fileDescriptor_5d34b1b7148314cc = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4f, 0x4f, 0xd4, 0x40,
	0x18, 0xc6, 0x33, 0xa0, 0x08, 0xef, 0xfa, 0x8f, 0x49, 0x84, 0x66, 0xd1, 0xb8, 0xe9, 0xc1, 0x6c,
	0xd6, 0xd0, 0x89, 0x2b, 0x5c, 0x24, 0x1a, 0x01, 0xc5, 0x40, 0x20, 0x92, 0x05, 0x2e, 0x5e, 0x36,
	0xc3, 0xee, 0x4b, 0x53, 0x69, 0x3b, 0xb5, 0x33, 0x4b, 0x14, 0x83, 0x07, 0x13, 0x3f, 0x81, 0xdf,
	0x40, 0x6f, 0x7e, 0x00, 0xaf, 0x1e, 0x3d, 0x78, 0xd4, 0x9b, 0x57, 0x3f, 0x88, 0xe9, 0xb4, 0xbb,
	0x14, 0xda, 0xed, 0xb2, 0x70, 0xeb, 0x4c, 0xe7, 0x99, 0xf7, 0x37, 0xcf, 0xf4, 0x79, 0x0b, 0xf3,
	0xb6, 0x10, 0xb6, 0x8b, 0xac, 0xe5, 0x8a, 0x4e, 0x9b, 0x29, 0xee, 0xa2, 0xaf, 0xd8, 0xc1, 0xdc,
	0x2e, 0x2a, 0xfe, 0x80, 0xf1, 0x20, 0x70, 0x9d, 0x16, 0x57, 0x8e, 0xf0, 0x9b, 0x12, 0xc3, 0x03,
	0xa7, 0x85, 0x56, 0x10, 0x0a, 0x25, 0xe8, 0x4c, 0x2c, 0xb3, 0xb4, 0xcc, 0x8a, 0x65, 0x56, 0x22,
	0x2b, 0xdf, 0x4e, 0xf6, 0xe4, 0x81, 0xc3, 0xb8, 0xef, 0x0b, 0xa5, 0x77, 0x90, 0xb1, 0xb4, 0x3c,
	0x7b, 0xc6, 0x8a, 0xc9, 0xf2, 0x6a, 0xd1, 0xf2, 0x96, 0xf0, 0xbc, 0xde, 0xca, 0x84, 0x89, 0xe9,
	0xd1, 0x6e, 0x67, 0x8f, 0xa1, 0x17, 0xa8, 0x77, 0xc9, 0xcb, 0xca, 0xe9, 0x97, 0x7b, 0x0e, 0xba,
	0xed, 0xa6, 0xc7, 0xe5, 0x7e, 0xbc, 0xc2, 0xfc, 0x00, 0xc6, 0x72, 0x88, 0x5c, 0xe1, 0xe2, 0x31,
	0x43, 0x03, 0xdf, 0x74, 0x50, 0x2a, 0x3a, 0x05, 0x63, 0x01, 0x0f, 0xd1, 0x57, 0x06, 0xa9, 0x90,
	0xea, 0x44, 0x23, 0x19, 0xd1, 0x35, 0x28, 0xa5, 0x88, 0x8d, 0x91, 0x0a, 0xa9, 0x96, 0xea, 0x55,
	0xab, 0xc0, 0x1c, 0x2b, 0xbd, 0x7b, 0x5a, 0x6c, 0xde, 0x87, 0x5b, 0x2f, 0x50, 0xe5, 0x14, 0xa7,
	0x70, 0xc9, 0xe7, 0x1e, 0x26, 0xa5, 0xf5, 0xb3, 0xf9, 0x95, 0x80, 0xb1, 0x13, 0xb4, 0xf3, 0x69,
	0x4f, 0x51, 0x91, 0x0b, 0x50, 0xd1, 0x05, 0x28, 0x75, 0x74, 0x1d, 0x6d, 0x55, 0x72, 0xc2, 0x72,
	0x77, 0xaf, 0xae, 0x9b, 0xd6, 0x4a, 0xe4, 0xe6, 0x06, 0x97, 0xfb, 0x0d, 0x88, 0x97, 0x47, 0xcf,
	0xa6, 0x05, 0xc6, 0x33, 0x74, 0x31, 0x17, 0x32, 0xef, 0x54, 0x1e, 0x4c, 0xaf, 0x3b, 0x32, 0xed,
	0x81, 0x1c, 0x74, 0x03, 0x77, 0x00, 0x02, 0x6e, 0x63, 0x53, 0x89, 0x7d, 0x8c, 0x2f, 0x60, 0xa2,
	0x31, 0x11, 0xcd, 0x6c, 0x47, 0x13, 0x74, 0x06, 0xf4, 0xa0, 0x29, 0x9d, 0x43, 0x34, 0x46, 0x2b,
	0xa4, 0x7a, 0xb9, 0x31, 0x1e, 0x4d, 0x6c, 0x39, 0x87, 0x68, 0xfe, 0x25, 0x60, 0x64, 0xeb, 0xc9,
	0x40, 0xf8, 0x12, 0xe9, 0x3a, 0x5c, 0x4d, 0xf9, 0x20, 0x0d, 0x52, 0x19, 0x1d, 0xca, 0xc5, 0x13,
	0x6a, 0x7a, 0x0f, 0x6e, 0xf8, 0xf8, 0x56, 0x35, 0x33, 0xac, 0xd7, 0xa2, 0xe9, 0xcd, 0x1e, 0xef,
	0x2a, 0x8c, 0x7b, 0xa8, 0x78, 0x9b, 0x2b, 0xae, 0x71, 0x4b, 0xf5, 0xd9, 0xc2, 0x8a, 0x5d, 0xdc,
	0x8d, 0x44, 0xd4, 0xe8, 0xc9, 0xeb, 0x7f, 0xae, 0x00, 0x4d, 0x01, 0x6d, 0xc5, 0xf9, 0xa5, 0x3f,
	0x08, 0x4c, 0x66, 0xbe, 0x73, 0x3a, 0x5f, 0x58, 0xa5, 0x5f, 0x2e, 0xca, 0x67, 0xb6, 0xc3, 0x5c,
	0xfd, 0xf8, 0xfb, 0xdf, 0xe7, 0x91, 0x65, 0xf3, 0x49, 0x2f, 0xbb, 0xef, 0xe3, 0x1b, 0x7c, 0x1c,
	0x84, 0xe2, 0x35, 0xb6, 0x94, 0x64, 0x35, 0xa6, 0xd0, 0xe7, 0xbe, 0x7e, 0x0a, 0x42, 0xb1, 0xe7,
	0xb8, 0x28, 0x59, 0xed, 0x28, 0xdd, 0x13, 0xe4, 0x23, 0x52, 0xa3, 0xdf, 0x09, 0x5c, 0x3f, 0x99,
	0x14, 0x5a, 0x2f, 0xe4, 0xc8, 0x8d, 0xd5, 0x10, 0xec, 0x2b, 0x9a, 0xfd, 0x29, 0x4d, 0xb1, 0x47,
	0x9f, 0xeb, 0x20, 0xf2, 0x13, 0xe0, 0xac, 0x76, 0x44, 0x7f, 0x11, 0x98, 0xcc, 0x84, 0x76, 0x80,
	0xf5, 0xfd, 0x42, 0x3e, 0x04, 0xfe, 0x8e, 0xc6, 0x7f, 0x59, 0x5f, 0x3b, 0xc6, 0x4f, 0xb7, 0xd9,
	0xf3, 0x1c, 0x25, 0xba, 0x86, 0x6f, 0x04, 0x26, 0x33, 0xe9, 0x1e, 0x70, 0x9a, 0x7e, 0xdd, 0xa0,
	0x3c, 0x95, 0xe9, 0x28, 0xcf, 0xa3, 0xe6, 0xdd, 0xb5, 0xbe, 0x76, 0x51, 0xeb, 0x7f, 0x12, 0xb8,
	0x79, 0x3a, 0xea, 0x74, 0xae, 0x90, 0xb5, 0x4f, 0x27, 0x2a, 0xcf, 0x0f, 0xa9, 0x8a, 0x03, 0x9a,
	0xf7, 0x11, 0x9d, 0x27, 0x00, 0x4b, 0x9f, 0x08, 0xdc, 0x6d, 0x09, 0xaf, 0x08, 0x62, 0x69, 0x3a,
	0x9b, 0xfb, 0xcd, 0xc8, 0xd7, 0x4d, 0xf2, 0x6a, 0x31, 0xd1, 0xd9, 0xc2, 0xe5, 0xbe, 0x6d, 0x89,
	0xd0, 0x66, 0x36, 0xfa, 0xda, 0x75, 0x16, 0xbf, 0xe2, 0x81, 0x23, 0x73, 0xff, 0xb6, 0x0b, 0xf1,
	0xf0, 0xcb, 0xc8, 0xe8, 0xf2, 0xf6, 0xd6, 0xee, 0x98, 0xd6, 0x3c, 0xfc, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0x39, 0x45, 0xae, 0x62, 0x41, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApplicationServiceClient is the client API for ApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationServiceClient interface {
	// Creates a new application entity.
	CreateApplication(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	// Retrieves specified application.
	GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	// Updates specified application.
	UpdateApplication(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	// Deletes specified application.
	DeleteApplication(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists all applications associated with the profile.
	ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error)
}

type applicationServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServiceClient(cc *grpc.ClientConn) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) CreateApplication(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.ApplicationService/CreateApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.ApplicationService/GetApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) UpdateApplication(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.ApplicationService/UpdateApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) DeleteApplication(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.ApplicationService/DeleteApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error) {
	out := new(ListApplicationsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.ApplicationService/ListApplications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationServiceServer is the server API for ApplicationService service.
type ApplicationServiceServer interface {
	// Creates a new application entity.
	CreateApplication(context.Context, *CreateApplicationRequest) (*Application, error)
	// Retrieves specified application.
	GetApplication(context.Context, *GetApplicationRequest) (*Application, error)
	// Updates specified application.
	UpdateApplication(context.Context, *UpdateApplicationRequest) (*Application, error)
	// Deletes specified application.
	DeleteApplication(context.Context, *DeleteApplicationRequest) (*empty.Empty, error)
	// Lists all applications associated with the profile.
	ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error)
}

func RegisterApplicationServiceServer(s *grpc.Server, srv ApplicationServiceServer) {
	s.RegisterService(&_ApplicationService_serviceDesc, srv)
}

func _ApplicationService_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.ApplicationService/CreateApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).CreateApplication(ctx, req.(*CreateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_GetApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).GetApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.ApplicationService/GetApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).GetApplication(ctx, req.(*GetApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_UpdateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).UpdateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.ApplicationService/UpdateApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).UpdateApplication(ctx, req.(*UpdateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_DeleteApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).DeleteApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.ApplicationService/DeleteApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).DeleteApplication(ctx, req.(*DeleteApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.ApplicationService/ListApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).ListApplications(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.talent.v4beta1.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApplication",
			Handler:    _ApplicationService_CreateApplication_Handler,
		},
		{
			MethodName: "GetApplication",
			Handler:    _ApplicationService_GetApplication_Handler,
		},
		{
			MethodName: "UpdateApplication",
			Handler:    _ApplicationService_UpdateApplication_Handler,
		},
		{
			MethodName: "DeleteApplication",
			Handler:    _ApplicationService_DeleteApplication_Handler,
		},
		{
			MethodName: "ListApplications",
			Handler:    _ApplicationService_ListApplications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/talent/v4beta1/application_service.proto",
}
