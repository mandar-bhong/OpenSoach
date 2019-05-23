// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/oslogin/v1beta/oslogin.proto

package oslogin

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	common "google.golang.org/genproto/googleapis/cloud/oslogin/common"
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

// The user profile information used for logging in to a virtual machine on
// Google Compute Engine.
type LoginProfile struct {
	// The primary email address that uniquely identifies the user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of POSIX accounts associated with the user.
	PosixAccounts []*common.PosixAccount `protobuf:"bytes,2,rep,name=posix_accounts,json=posixAccounts,proto3" json:"posix_accounts,omitempty"`
	// A map from SSH public key fingerprint to the associated key object.
	SshPublicKeys map[string]*common.SshPublicKey `protobuf:"bytes,3,rep,name=ssh_public_keys,json=sshPublicKeys,proto3" json:"ssh_public_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Indicates if the user is suspended. A suspended user cannot log in but
	// their profile information is retained.
	Suspended            bool     `protobuf:"varint,4,opt,name=suspended,proto3" json:"suspended,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginProfile) Reset()         { *m = LoginProfile{} }
func (m *LoginProfile) String() string { return proto.CompactTextString(m) }
func (*LoginProfile) ProtoMessage()    {}
func (*LoginProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51261020b887e3b, []int{0}
}

func (m *LoginProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginProfile.Unmarshal(m, b)
}
func (m *LoginProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginProfile.Marshal(b, m, deterministic)
}
func (m *LoginProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginProfile.Merge(m, src)
}
func (m *LoginProfile) XXX_Size() int {
	return xxx_messageInfo_LoginProfile.Size(m)
}
func (m *LoginProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginProfile.DiscardUnknown(m)
}

var xxx_messageInfo_LoginProfile proto.InternalMessageInfo

func (m *LoginProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginProfile) GetPosixAccounts() []*common.PosixAccount {
	if m != nil {
		return m.PosixAccounts
	}
	return nil
}

func (m *LoginProfile) GetSshPublicKeys() map[string]*common.SshPublicKey {
	if m != nil {
		return m.SshPublicKeys
	}
	return nil
}

func (m *LoginProfile) GetSuspended() bool {
	if m != nil {
		return m.Suspended
	}
	return false
}

// A request message for deleting a POSIX account entry.
type DeletePosixAccountRequest struct {
	// A reference to the POSIX account to update. POSIX accounts are identified
	// by the project ID they are associated with. A reference to the POSIX
	// account is in format `users/{user}/projects/{project}`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePosixAccountRequest) Reset()         { *m = DeletePosixAccountRequest{} }
func (m *DeletePosixAccountRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePosixAccountRequest) ProtoMessage()    {}
func (*DeletePosixAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51261020b887e3b, []int{1}
}

func (m *DeletePosixAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePosixAccountRequest.Unmarshal(m, b)
}
func (m *DeletePosixAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePosixAccountRequest.Marshal(b, m, deterministic)
}
func (m *DeletePosixAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePosixAccountRequest.Merge(m, src)
}
func (m *DeletePosixAccountRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePosixAccountRequest.Size(m)
}
func (m *DeletePosixAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePosixAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePosixAccountRequest proto.InternalMessageInfo

func (m *DeletePosixAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A request message for deleting an SSH public key.
type DeleteSshPublicKeyRequest struct {
	// The fingerprint of the public key to update. Public keys are identified by
	// their SHA-256 fingerprint. The fingerprint of the public key is in format
	// `users/{user}/sshPublicKeys/{fingerprint}`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSshPublicKeyRequest) Reset()         { *m = DeleteSshPublicKeyRequest{} }
func (m *DeleteSshPublicKeyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSshPublicKeyRequest) ProtoMessage()    {}
func (*DeleteSshPublicKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51261020b887e3b, []int{2}
}

func (m *DeleteSshPublicKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSshPublicKeyRequest.Unmarshal(m, b)
}
func (m *DeleteSshPublicKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSshPublicKeyRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSshPublicKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSshPublicKeyRequest.Merge(m, src)
}
func (m *DeleteSshPublicKeyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSshPublicKeyRequest.Size(m)
}
func (m *DeleteSshPublicKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSshPublicKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSshPublicKeyRequest proto.InternalMessageInfo

func (m *DeleteSshPublicKeyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A request message for retrieving the login profile information for a user.
type GetLoginProfileRequest struct {
	// The unique ID for the user in format `users/{user}`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLoginProfileRequest) Reset()         { *m = GetLoginProfileRequest{} }
func (m *GetLoginProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetLoginProfileRequest) ProtoMessage()    {}
func (*GetLoginProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51261020b887e3b, []int{3}
}

func (m *GetLoginProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLoginProfileRequest.Unmarshal(m, b)
}
func (m *GetLoginProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLoginProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetLoginProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLoginProfileRequest.Merge(m, src)
}
func (m *GetLoginProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetLoginProfileRequest.Size(m)
}
func (m *GetLoginProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLoginProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLoginProfileRequest proto.InternalMessageInfo

func (m *GetLoginProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A request message for retrieving an SSH public key.
type GetSshPublicKeyRequest struct {
	// The fingerprint of the public key to retrieve. Public keys are identified
	// by their SHA-256 fingerprint. The fingerprint of the public key is in
	// format `users/{user}/sshPublicKeys/{fingerprint}`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSshPublicKeyRequest) Reset()         { *m = GetSshPublicKeyRequest{} }
func (m *GetSshPublicKeyRequest) String() string { return proto.CompactTextString(m) }
func (*GetSshPublicKeyRequest) ProtoMessage()    {}
func (*GetSshPublicKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51261020b887e3b, []int{4}
}

func (m *GetSshPublicKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSshPublicKeyRequest.Unmarshal(m, b)
}
func (m *GetSshPublicKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSshPublicKeyRequest.Marshal(b, m, deterministic)
}
func (m *GetSshPublicKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSshPublicKeyRequest.Merge(m, src)
}
func (m *GetSshPublicKeyRequest) XXX_Size() int {
	return xxx_messageInfo_GetSshPublicKeyRequest.Size(m)
}
func (m *GetSshPublicKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSshPublicKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSshPublicKeyRequest proto.InternalMessageInfo

func (m *GetSshPublicKeyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A request message for importing an SSH public key.
type ImportSshPublicKeyRequest struct {
	// The unique ID for the user in format `users/{user}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The SSH public key and expiration time.
	SshPublicKey *common.SshPublicKey `protobuf:"bytes,2,opt,name=ssh_public_key,json=sshPublicKey,proto3" json:"ssh_public_key,omitempty"`
	// The project ID of the Google Cloud Platform project.
	ProjectId            string   `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImportSshPublicKeyRequest) Reset()         { *m = ImportSshPublicKeyRequest{} }
func (m *ImportSshPublicKeyRequest) String() string { return proto.CompactTextString(m) }
func (*ImportSshPublicKeyRequest) ProtoMessage()    {}
func (*ImportSshPublicKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51261020b887e3b, []int{5}
}

func (m *ImportSshPublicKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportSshPublicKeyRequest.Unmarshal(m, b)
}
func (m *ImportSshPublicKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportSshPublicKeyRequest.Marshal(b, m, deterministic)
}
func (m *ImportSshPublicKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportSshPublicKeyRequest.Merge(m, src)
}
func (m *ImportSshPublicKeyRequest) XXX_Size() int {
	return xxx_messageInfo_ImportSshPublicKeyRequest.Size(m)
}
func (m *ImportSshPublicKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportSshPublicKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImportSshPublicKeyRequest proto.InternalMessageInfo

func (m *ImportSshPublicKeyRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ImportSshPublicKeyRequest) GetSshPublicKey() *common.SshPublicKey {
	if m != nil {
		return m.SshPublicKey
	}
	return nil
}

func (m *ImportSshPublicKeyRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

// A response message for importing an SSH public key.
type ImportSshPublicKeyResponse struct {
	// The login profile information for the user.
	LoginProfile         *LoginProfile `protobuf:"bytes,1,opt,name=login_profile,json=loginProfile,proto3" json:"login_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ImportSshPublicKeyResponse) Reset()         { *m = ImportSshPublicKeyResponse{} }
func (m *ImportSshPublicKeyResponse) String() string { return proto.CompactTextString(m) }
func (*ImportSshPublicKeyResponse) ProtoMessage()    {}
func (*ImportSshPublicKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51261020b887e3b, []int{6}
}

func (m *ImportSshPublicKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportSshPublicKeyResponse.Unmarshal(m, b)
}
func (m *ImportSshPublicKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportSshPublicKeyResponse.Marshal(b, m, deterministic)
}
func (m *ImportSshPublicKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportSshPublicKeyResponse.Merge(m, src)
}
func (m *ImportSshPublicKeyResponse) XXX_Size() int {
	return xxx_messageInfo_ImportSshPublicKeyResponse.Size(m)
}
func (m *ImportSshPublicKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportSshPublicKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImportSshPublicKeyResponse proto.InternalMessageInfo

func (m *ImportSshPublicKeyResponse) GetLoginProfile() *LoginProfile {
	if m != nil {
		return m.LoginProfile
	}
	return nil
}

// A request message for updating an SSH public key.
type UpdateSshPublicKeyRequest struct {
	// The fingerprint of the public key to update. Public keys are identified by
	// their SHA-256 fingerprint. The fingerprint of the public key is in format
	// `users/{user}/sshPublicKeys/{fingerprint}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The SSH public key and expiration time.
	SshPublicKey *common.SshPublicKey `protobuf:"bytes,2,opt,name=ssh_public_key,json=sshPublicKey,proto3" json:"ssh_public_key,omitempty"`
	// Mask to control which fields get updated. Updates all if not present.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateSshPublicKeyRequest) Reset()         { *m = UpdateSshPublicKeyRequest{} }
func (m *UpdateSshPublicKeyRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSshPublicKeyRequest) ProtoMessage()    {}
func (*UpdateSshPublicKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51261020b887e3b, []int{7}
}

func (m *UpdateSshPublicKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSshPublicKeyRequest.Unmarshal(m, b)
}
func (m *UpdateSshPublicKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSshPublicKeyRequest.Marshal(b, m, deterministic)
}
func (m *UpdateSshPublicKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSshPublicKeyRequest.Merge(m, src)
}
func (m *UpdateSshPublicKeyRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSshPublicKeyRequest.Size(m)
}
func (m *UpdateSshPublicKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSshPublicKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSshPublicKeyRequest proto.InternalMessageInfo

func (m *UpdateSshPublicKeyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateSshPublicKeyRequest) GetSshPublicKey() *common.SshPublicKey {
	if m != nil {
		return m.SshPublicKey
	}
	return nil
}

func (m *UpdateSshPublicKeyRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginProfile)(nil), "google.cloud.oslogin.v1beta.LoginProfile")
	proto.RegisterMapType((map[string]*common.SshPublicKey)(nil), "google.cloud.oslogin.v1beta.LoginProfile.SshPublicKeysEntry")
	proto.RegisterType((*DeletePosixAccountRequest)(nil), "google.cloud.oslogin.v1beta.DeletePosixAccountRequest")
	proto.RegisterType((*DeleteSshPublicKeyRequest)(nil), "google.cloud.oslogin.v1beta.DeleteSshPublicKeyRequest")
	proto.RegisterType((*GetLoginProfileRequest)(nil), "google.cloud.oslogin.v1beta.GetLoginProfileRequest")
	proto.RegisterType((*GetSshPublicKeyRequest)(nil), "google.cloud.oslogin.v1beta.GetSshPublicKeyRequest")
	proto.RegisterType((*ImportSshPublicKeyRequest)(nil), "google.cloud.oslogin.v1beta.ImportSshPublicKeyRequest")
	proto.RegisterType((*ImportSshPublicKeyResponse)(nil), "google.cloud.oslogin.v1beta.ImportSshPublicKeyResponse")
	proto.RegisterType((*UpdateSshPublicKeyRequest)(nil), "google.cloud.oslogin.v1beta.UpdateSshPublicKeyRequest")
}

func init() {
	proto.RegisterFile("google/cloud/oslogin/v1beta/oslogin.proto", fileDescriptor_c51261020b887e3b)
}

var fileDescriptor_c51261020b887e3b = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x4f, 0x4f, 0x13, 0x4f,
	0x18, 0xc7, 0xb3, 0x2d, 0x3f, 0x02, 0x43, 0x81, 0x5f, 0xe6, 0x40, 0xca, 0x82, 0xb1, 0x2e, 0xd1,
	0x94, 0x62, 0x76, 0x43, 0x49, 0x94, 0x80, 0x68, 0x2c, 0x22, 0x21, 0xfe, 0xa1, 0x29, 0x91, 0x83,
	0x21, 0x69, 0xa6, 0xbb, 0xc3, 0xb2, 0x76, 0x77, 0x67, 0xdc, 0xd9, 0x25, 0x36, 0x86, 0x8b, 0x89,
	0x67, 0x0f, 0x7a, 0x36, 0x31, 0xde, 0xbc, 0x79, 0xf2, 0xe4, 0x1b, 0xe0, 0xea, 0x5b, 0xd0, 0xf7,
	0x61, 0x76, 0x76, 0x16, 0xb6, 0xed, 0xb6, 0x5d, 0x12, 0x4f, 0xdd, 0x99, 0xe7, 0xdf, 0x67, 0xbe,
	0x33, 0xcf, 0x93, 0x82, 0x65, 0x93, 0x10, 0xd3, 0xc6, 0x9a, 0x6e, 0x93, 0xc0, 0xd0, 0x08, 0xb3,
	0x89, 0x69, 0xb9, 0xda, 0xe9, 0x6a, 0x0b, 0xfb, 0x28, 0x5e, 0xaa, 0xd4, 0x23, 0x3e, 0x81, 0x0b,
	0x91, 0xab, 0xca, 0x5d, 0xd5, 0xd8, 0x16, 0xb9, 0xca, 0x8b, 0x22, 0x0f, 0xa2, 0x96, 0x86, 0x5c,
	0x97, 0xf8, 0xc8, 0xb7, 0x88, 0xcb, 0xa2, 0x50, 0xb9, 0x9c, 0x5a, 0x45, 0x27, 0x8e, 0x43, 0xe2,
	0x1f, 0xe1, 0x29, 0x8a, 0x68, 0x7c, 0xd5, 0x0a, 0x8e, 0x35, 0xec, 0x50, 0xbf, 0x23, 0x8c, 0xa5,
	0x5e, 0xe3, 0xb1, 0x85, 0x6d, 0xa3, 0xe9, 0x20, 0xd6, 0x8e, 0x3c, 0x94, 0x3f, 0x39, 0x50, 0x78,
	0x1a, 0x26, 0xaf, 0x7b, 0xe4, 0xd8, 0xb2, 0x31, 0x84, 0x60, 0xcc, 0x45, 0x0e, 0x2e, 0x4a, 0x25,
	0xa9, 0x3c, 0xd9, 0xe0, 0xdf, 0xb0, 0x0e, 0x66, 0x28, 0x61, 0xd6, 0x9b, 0x26, 0xd2, 0x75, 0x12,
	0xb8, 0x3e, 0x2b, 0xe6, 0x4a, 0xf9, 0xf2, 0x54, 0x75, 0x59, 0x4d, 0x3d, 0xa1, 0xe0, 0xab, 0x87,
	0x21, 0x0f, 0xa3, 0x88, 0xc6, 0x34, 0x4d, 0xac, 0x18, 0x34, 0xc0, 0x2c, 0x63, 0x27, 0x4d, 0x1a,
	0xb4, 0x6c, 0x4b, 0x6f, 0xb6, 0x71, 0x87, 0x15, 0xf3, 0x3c, 0xe5, 0x3d, 0x75, 0x88, 0x68, 0x6a,
	0x92, 0x54, 0x3d, 0x60, 0x27, 0x75, 0x1e, 0xff, 0x04, 0x77, 0xd8, 0x8e, 0xeb, 0x7b, 0x9d, 0xc6,
	0x34, 0x4b, 0xee, 0xc1, 0x45, 0x30, 0xc9, 0x02, 0x46, 0xb1, 0x6b, 0x60, 0xa3, 0x38, 0x56, 0x92,
	0xca, 0x13, 0x8d, 0xcb, 0x0d, 0xb9, 0x0d, 0x60, 0x7f, 0x0a, 0xf8, 0x3f, 0xc8, 0xb7, 0x71, 0x47,
	0x1c, 0x3f, 0xfc, 0x84, 0x0f, 0xc0, 0x7f, 0xa7, 0xc8, 0x0e, 0x70, 0x31, 0x57, 0x92, 0x46, 0x1e,
	0x3a, 0x99, 0xb1, 0x11, 0xc5, 0x6d, 0xe4, 0xd6, 0x25, 0x45, 0x03, 0xf3, 0x8f, 0xb0, 0x8d, 0x7d,
	0xdc, 0xa5, 0x0a, 0x7e, 0x1d, 0x60, 0xe6, 0xa7, 0x69, 0x7e, 0x19, 0xd0, 0x95, 0x71, 0x48, 0xc0,
	0x6d, 0x30, 0xb7, 0x8b, 0xfd, 0xa4, 0x42, 0xa3, 0xbd, 0xb3, 0xe6, 0xfe, 0x2a, 0x81, 0xf9, 0x3d,
	0x87, 0x12, 0x2f, 0x35, 0x62, 0x0e, 0x8c, 0x53, 0xe4, 0x61, 0xd7, 0x17, 0x31, 0x62, 0x05, 0xf7,
	0xc1, 0x4c, 0xf7, 0x25, 0x5f, 0x5d, 0xc1, 0x42, 0xf2, 0x42, 0xe1, 0x35, 0x00, 0xa8, 0x47, 0x5e,
	0x61, 0xdd, 0x6f, 0x5a, 0x46, 0x31, 0xcf, 0x8b, 0x4d, 0x8a, 0x9d, 0x3d, 0x43, 0xb1, 0x81, 0x9c,
	0x06, 0xc9, 0x28, 0x71, 0x19, 0x86, 0xcf, 0xc1, 0x34, 0xaf, 0xd3, 0xa4, 0x91, 0x3a, 0x1c, 0x76,
	0x20, 0x4c, 0xca, 0x83, 0x6b, 0x14, 0xec, 0xc4, 0x4a, 0xf9, 0x29, 0x81, 0xf9, 0x17, 0xd4, 0x40,
	0x99, 0x6f, 0xe8, 0xdf, 0xeb, 0xb1, 0x09, 0xa6, 0x02, 0x4e, 0xc0, 0x3b, 0x9a, 0x0b, 0x32, 0x55,
	0x95, 0xe3, 0x6c, 0x71, 0xd3, 0xab, 0x8f, 0xc3, 0xa6, 0x7f, 0x86, 0x58, 0xbb, 0x01, 0x22, 0xf7,
	0xf0, 0xbb, 0xfa, 0x7e, 0x02, 0xcc, 0xec, 0x33, 0x7e, 0xc0, 0x03, 0xec, 0x9d, 0x5a, 0x3a, 0x86,
	0x1f, 0x24, 0x00, 0xfb, 0x5f, 0x29, 0xbc, 0x33, 0x54, 0xa2, 0x81, 0xcf, 0x5a, 0x9e, 0xeb, 0x23,
	0xd9, 0x09, 0x67, 0x93, 0xb2, 0xfc, 0xee, 0xd7, 0xef, 0x8f, 0xb9, 0xa5, 0xca, 0x8d, 0x78, 0x6c,
	0xbe, 0x0d, 0xe5, 0xd9, 0x0a, 0x18, 0xf6, 0x98, 0x56, 0xd1, 0xc4, 0x85, 0x32, 0xad, 0x72, 0x06,
	0x3f, 0x5d, 0x10, 0x25, 0x65, 0xc8, 0x44, 0x94, 0x72, 0x2b, 0x03, 0x89, 0x54, 0x4e, 0x54, 0xae,
	0xdc, 0x4a, 0x27, 0xea, 0x9a, 0x2a, 0x21, 0xd6, 0x67, 0x09, 0xcc, 0xf6, 0x34, 0x1b, 0x5c, 0x1b,
	0xca, 0x94, 0xde, 0x9a, 0x72, 0xf6, 0xd7, 0xa7, 0xac, 0x70, 0xc6, 0x9b, 0x70, 0x29, 0x95, 0xf1,
	0x4c, 0x4b, 0x3e, 0x4e, 0xf8, 0x25, 0x02, 0xec, 0x12, 0x6d, 0x24, 0x60, 0x9a, 0x62, 0xd9, 0xdf,
	0x66, 0x2c, 0x22, 0xcc, 0x2a, 0xe2, 0xb9, 0x04, 0x60, 0x7f, 0xbf, 0x8e, 0xb8, 0xdb, 0x81, 0x53,
	0x48, 0xbe, 0x7b, 0xe5, 0xb8, 0x68, 0x30, 0x28, 0xdb, 0x9c, 0x7b, 0x4b, 0x59, 0xb9, 0xe0, 0x8e,
	0xe6, 0xd7, 0x85, 0xb4, 0x1b, 0x56, 0x5f, 0xf0, 0x46, 0x4f, 0x27, 0xc3, 0x1f, 0x12, 0x80, 0xfd,
	0xd3, 0x60, 0xc4, 0x61, 0x06, 0x8e, 0x8f, 0xab, 0xc8, 0x7e, 0x9f, 0xe3, 0xaf, 0x57, 0x33, 0xca,
	0xde, 0x4b, 0x5e, 0xfb, 0x2e, 0x81, 0xeb, 0x3a, 0x71, 0x86, 0x81, 0xd6, 0x0a, 0x62, 0x50, 0xd4,
	0xc3, 0xb6, 0xa9, 0x4b, 0x2f, 0x6b, 0xc2, 0xd9, 0x24, 0x36, 0x72, 0x4d, 0x95, 0x78, 0xa6, 0x66,
	0x62, 0x97, 0x37, 0x95, 0x16, 0x99, 0x10, 0xb5, 0x58, 0xea, 0x7f, 0xa4, 0x4d, 0xb1, 0xfc, 0x96,
	0x5b, 0xd8, 0x8d, 0x92, 0x6c, 0xf3, 0x8a, 0xa2, 0x82, 0x7a, 0xb8, 0x5a, 0xc3, 0x3e, 0x3a, 0x8f,
	0xad, 0x47, 0xdc, 0x7a, 0x24, 0xac, 0x47, 0x87, 0x3c, 0x47, 0x6b, 0x9c, 0x97, 0x5a, 0xfb, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0xc7, 0x96, 0xd9, 0xc2, 0x8d, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OsLoginServiceClient is the client API for OsLoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OsLoginServiceClient interface {
	// Deletes a POSIX account.
	DeletePosixAccount(ctx context.Context, in *DeletePosixAccountRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Deletes an SSH public key.
	DeleteSshPublicKey(ctx context.Context, in *DeleteSshPublicKeyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Retrieves the profile information used for logging in to a virtual machine
	// on Google Compute Engine.
	GetLoginProfile(ctx context.Context, in *GetLoginProfileRequest, opts ...grpc.CallOption) (*LoginProfile, error)
	// Retrieves an SSH public key.
	GetSshPublicKey(ctx context.Context, in *GetSshPublicKeyRequest, opts ...grpc.CallOption) (*common.SshPublicKey, error)
	// Adds an SSH public key and returns the profile information. Default POSIX
	// account information is set when no username and UID exist as part of the
	// login profile.
	ImportSshPublicKey(ctx context.Context, in *ImportSshPublicKeyRequest, opts ...grpc.CallOption) (*ImportSshPublicKeyResponse, error)
	// Updates an SSH public key and returns the profile information. This method
	// supports patch semantics.
	UpdateSshPublicKey(ctx context.Context, in *UpdateSshPublicKeyRequest, opts ...grpc.CallOption) (*common.SshPublicKey, error)
}

type osLoginServiceClient struct {
	cc *grpc.ClientConn
}

func NewOsLoginServiceClient(cc *grpc.ClientConn) OsLoginServiceClient {
	return &osLoginServiceClient{cc}
}

func (c *osLoginServiceClient) DeletePosixAccount(ctx context.Context, in *DeletePosixAccountRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.oslogin.v1beta.OsLoginService/DeletePosixAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osLoginServiceClient) DeleteSshPublicKey(ctx context.Context, in *DeleteSshPublicKeyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.oslogin.v1beta.OsLoginService/DeleteSshPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osLoginServiceClient) GetLoginProfile(ctx context.Context, in *GetLoginProfileRequest, opts ...grpc.CallOption) (*LoginProfile, error) {
	out := new(LoginProfile)
	err := c.cc.Invoke(ctx, "/google.cloud.oslogin.v1beta.OsLoginService/GetLoginProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osLoginServiceClient) GetSshPublicKey(ctx context.Context, in *GetSshPublicKeyRequest, opts ...grpc.CallOption) (*common.SshPublicKey, error) {
	out := new(common.SshPublicKey)
	err := c.cc.Invoke(ctx, "/google.cloud.oslogin.v1beta.OsLoginService/GetSshPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osLoginServiceClient) ImportSshPublicKey(ctx context.Context, in *ImportSshPublicKeyRequest, opts ...grpc.CallOption) (*ImportSshPublicKeyResponse, error) {
	out := new(ImportSshPublicKeyResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.oslogin.v1beta.OsLoginService/ImportSshPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osLoginServiceClient) UpdateSshPublicKey(ctx context.Context, in *UpdateSshPublicKeyRequest, opts ...grpc.CallOption) (*common.SshPublicKey, error) {
	out := new(common.SshPublicKey)
	err := c.cc.Invoke(ctx, "/google.cloud.oslogin.v1beta.OsLoginService/UpdateSshPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OsLoginServiceServer is the server API for OsLoginService service.
type OsLoginServiceServer interface {
	// Deletes a POSIX account.
	DeletePosixAccount(context.Context, *DeletePosixAccountRequest) (*empty.Empty, error)
	// Deletes an SSH public key.
	DeleteSshPublicKey(context.Context, *DeleteSshPublicKeyRequest) (*empty.Empty, error)
	// Retrieves the profile information used for logging in to a virtual machine
	// on Google Compute Engine.
	GetLoginProfile(context.Context, *GetLoginProfileRequest) (*LoginProfile, error)
	// Retrieves an SSH public key.
	GetSshPublicKey(context.Context, *GetSshPublicKeyRequest) (*common.SshPublicKey, error)
	// Adds an SSH public key and returns the profile information. Default POSIX
	// account information is set when no username and UID exist as part of the
	// login profile.
	ImportSshPublicKey(context.Context, *ImportSshPublicKeyRequest) (*ImportSshPublicKeyResponse, error)
	// Updates an SSH public key and returns the profile information. This method
	// supports patch semantics.
	UpdateSshPublicKey(context.Context, *UpdateSshPublicKeyRequest) (*common.SshPublicKey, error)
}

func RegisterOsLoginServiceServer(s *grpc.Server, srv OsLoginServiceServer) {
	s.RegisterService(&_OsLoginService_serviceDesc, srv)
}

func _OsLoginService_DeletePosixAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePosixAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsLoginServiceServer).DeletePosixAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.oslogin.v1beta.OsLoginService/DeletePosixAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsLoginServiceServer).DeletePosixAccount(ctx, req.(*DeletePosixAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsLoginService_DeleteSshPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSshPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsLoginServiceServer).DeleteSshPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.oslogin.v1beta.OsLoginService/DeleteSshPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsLoginServiceServer).DeleteSshPublicKey(ctx, req.(*DeleteSshPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsLoginService_GetLoginProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsLoginServiceServer).GetLoginProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.oslogin.v1beta.OsLoginService/GetLoginProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsLoginServiceServer).GetLoginProfile(ctx, req.(*GetLoginProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsLoginService_GetSshPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSshPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsLoginServiceServer).GetSshPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.oslogin.v1beta.OsLoginService/GetSshPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsLoginServiceServer).GetSshPublicKey(ctx, req.(*GetSshPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsLoginService_ImportSshPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportSshPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsLoginServiceServer).ImportSshPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.oslogin.v1beta.OsLoginService/ImportSshPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsLoginServiceServer).ImportSshPublicKey(ctx, req.(*ImportSshPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsLoginService_UpdateSshPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSshPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsLoginServiceServer).UpdateSshPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.oslogin.v1beta.OsLoginService/UpdateSshPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsLoginServiceServer).UpdateSshPublicKey(ctx, req.(*UpdateSshPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OsLoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.oslogin.v1beta.OsLoginService",
	HandlerType: (*OsLoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeletePosixAccount",
			Handler:    _OsLoginService_DeletePosixAccount_Handler,
		},
		{
			MethodName: "DeleteSshPublicKey",
			Handler:    _OsLoginService_DeleteSshPublicKey_Handler,
		},
		{
			MethodName: "GetLoginProfile",
			Handler:    _OsLoginService_GetLoginProfile_Handler,
		},
		{
			MethodName: "GetSshPublicKey",
			Handler:    _OsLoginService_GetSshPublicKey_Handler,
		},
		{
			MethodName: "ImportSshPublicKey",
			Handler:    _OsLoginService_ImportSshPublicKey_Handler,
		},
		{
			MethodName: "UpdateSshPublicKey",
			Handler:    _OsLoginService_UpdateSshPublicKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/oslogin/v1beta/oslogin.proto",
}
