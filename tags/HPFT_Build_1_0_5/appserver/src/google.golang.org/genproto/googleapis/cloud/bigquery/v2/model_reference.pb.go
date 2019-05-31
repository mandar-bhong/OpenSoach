// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/bigquery/v2/model_reference.proto

package bigquery

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

// Id path of a model.
type ModelReference struct {
	// [Required] The ID of the project containing this model.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// [Required] The ID of the dataset containing this model.
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// [Required] The ID of the model. The ID must contain only
	// letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	// length is 1,024 characters.
	ModelId              string   `protobuf:"bytes,3,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelReference) Reset()         { *m = ModelReference{} }
func (m *ModelReference) String() string { return proto.CompactTextString(m) }
func (*ModelReference) ProtoMessage()    {}
func (*ModelReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c8f9aced577e06a, []int{0}
}

func (m *ModelReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelReference.Unmarshal(m, b)
}
func (m *ModelReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelReference.Marshal(b, m, deterministic)
}
func (m *ModelReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelReference.Merge(m, src)
}
func (m *ModelReference) XXX_Size() int {
	return xxx_messageInfo_ModelReference.Size(m)
}
func (m *ModelReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelReference.DiscardUnknown(m)
}

var xxx_messageInfo_ModelReference proto.InternalMessageInfo

func (m *ModelReference) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ModelReference) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *ModelReference) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func init() {
	proto.RegisterType((*ModelReference)(nil), "google.cloud.bigquery.v2.ModelReference")
}

func init() {
	proto.RegisterFile("google/cloud/bigquery/v2/model_reference.proto", fileDescriptor_4c8f9aced577e06a)
}

var fileDescriptor_4c8f9aced577e06a = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0x80, 0x39, 0x05, 0xf5, 0x32, 0x38, 0xd4, 0xa5, 0xca, 0x09, 0xe2, 0xe4, 0x94, 0xc0, 0x39,
	0xba, 0x48, 0xb7, 0x0e, 0x82, 0x74, 0x74, 0x29, 0x69, 0xde, 0x33, 0x44, 0xd2, 0xbc, 0x98, 0xa6,
	0x05, 0xff, 0xbd, 0x34, 0x49, 0x07, 0x91, 0xdb, 0x92, 0xf7, 0x7d, 0x79, 0x7c, 0x84, 0x71, 0x4d,
	0xa4, 0x2d, 0x0a, 0x65, 0x69, 0x06, 0x31, 0x18, 0xfd, 0x3d, 0x63, 0xf8, 0x11, 0xcb, 0x51, 0x8c,
	0x04, 0x68, 0xfb, 0x80, 0x9f, 0x18, 0xd0, 0x29, 0xe4, 0x3e, 0x50, 0xa4, 0xaa, 0xce, 0x3e, 0x4f,
	0x3e, 0xdf, 0x7c, 0xbe, 0x1c, 0xef, 0x0e, 0x65, 0x93, 0xf4, 0x46, 0x48, 0xe7, 0x28, 0xca, 0x68,
	0xc8, 0x4d, 0xf9, 0xdd, 0xa3, 0x61, 0xd7, 0x6f, 0xeb, 0xc2, 0x6e, 0xdb, 0x57, 0xdd, 0x33, 0xe6,
	0x03, 0x7d, 0xa1, 0x8a, 0xbd, 0x81, 0x7a, 0xf7, 0xb0, 0x7b, 0xda, 0x77, 0xfb, 0x32, 0x69, 0x61,
	0xc5, 0x20, 0xa3, 0x9c, 0x30, 0xe1, 0xb3, 0x8c, 0xcb, 0xa4, 0x85, 0xea, 0x96, 0x5d, 0xe5, 0x40,
	0x03, 0xf5, 0x79, 0x82, 0x97, 0xe9, 0xde, 0x42, 0x33, 0xb3, 0x83, 0xa2, 0x91, 0x9f, 0x0a, 0x6d,
	0x6e, 0xfe, 0x86, 0xbc, 0xaf, 0x7d, 0x1f, 0xaf, 0x45, 0xd7, 0x64, 0xa5, 0xd3, 0x9c, 0x82, 0x16,
	0x1a, 0x5d, 0x6a, 0x17, 0x19, 0x49, 0x6f, 0xa6, 0xff, 0xdf, 0xf4, 0xb2, 0x9d, 0x87, 0x8b, 0x24,
	0x3f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x56, 0x77, 0xce, 0xa8, 0x52, 0x01, 0x00, 0x00,
}
