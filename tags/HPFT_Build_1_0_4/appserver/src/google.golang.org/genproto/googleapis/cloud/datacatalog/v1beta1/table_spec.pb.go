// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1beta1/table_spec.proto

package datacatalog

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Table source type.
type TableSourceType int32

const (
	// Default unknown type.
	TableSourceType_TABLE_SOURCE_TYPE_UNSPECIFIED TableSourceType = 0
	// Table view.
	TableSourceType_BIGQUERY_VIEW TableSourceType = 2
	// BigQuery native table.
	TableSourceType_BIGQUERY_TABLE TableSourceType = 5
)

var TableSourceType_name = map[int32]string{
	0: "TABLE_SOURCE_TYPE_UNSPECIFIED",
	2: "BIGQUERY_VIEW",
	5: "BIGQUERY_TABLE",
}

var TableSourceType_value = map[string]int32{
	"TABLE_SOURCE_TYPE_UNSPECIFIED": 0,
	"BIGQUERY_VIEW":                 2,
	"BIGQUERY_TABLE":                5,
}

func (x TableSourceType) String() string {
	return proto.EnumName(TableSourceType_name, int32(x))
}

func (TableSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2effb41fca72136b, []int{0}
}

// Describes a BigQuery table.
type BigQueryTableSpec struct {
	// The table source type.
	TableSourceType TableSourceType `protobuf:"varint,1,opt,name=table_source_type,json=tableSourceType,proto3,enum=google.cloud.datacatalog.v1beta1.TableSourceType" json:"table_source_type,omitempty"`
	// Table view specification. This field should only be populated if
	// table_source_type is BIGQUERY_VIEW.
	ViewSpec             *ViewSpec `protobuf:"bytes,2,opt,name=view_spec,json=viewSpec,proto3" json:"view_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BigQueryTableSpec) Reset()         { *m = BigQueryTableSpec{} }
func (m *BigQueryTableSpec) String() string { return proto.CompactTextString(m) }
func (*BigQueryTableSpec) ProtoMessage()    {}
func (*BigQueryTableSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_2effb41fca72136b, []int{0}
}

func (m *BigQueryTableSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigQueryTableSpec.Unmarshal(m, b)
}
func (m *BigQueryTableSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigQueryTableSpec.Marshal(b, m, deterministic)
}
func (m *BigQueryTableSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigQueryTableSpec.Merge(m, src)
}
func (m *BigQueryTableSpec) XXX_Size() int {
	return xxx_messageInfo_BigQueryTableSpec.Size(m)
}
func (m *BigQueryTableSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_BigQueryTableSpec.DiscardUnknown(m)
}

var xxx_messageInfo_BigQueryTableSpec proto.InternalMessageInfo

func (m *BigQueryTableSpec) GetTableSourceType() TableSourceType {
	if m != nil {
		return m.TableSourceType
	}
	return TableSourceType_TABLE_SOURCE_TYPE_UNSPECIFIED
}

func (m *BigQueryTableSpec) GetViewSpec() *ViewSpec {
	if m != nil {
		return m.ViewSpec
	}
	return nil
}

// Table view specification.
type ViewSpec struct {
	// The query that defines the table view.
	ViewQuery            string   `protobuf:"bytes,1,opt,name=view_query,json=viewQuery,proto3" json:"view_query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewSpec) Reset()         { *m = ViewSpec{} }
func (m *ViewSpec) String() string { return proto.CompactTextString(m) }
func (*ViewSpec) ProtoMessage()    {}
func (*ViewSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_2effb41fca72136b, []int{1}
}

func (m *ViewSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewSpec.Unmarshal(m, b)
}
func (m *ViewSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewSpec.Marshal(b, m, deterministic)
}
func (m *ViewSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewSpec.Merge(m, src)
}
func (m *ViewSpec) XXX_Size() int {
	return xxx_messageInfo_ViewSpec.Size(m)
}
func (m *ViewSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ViewSpec proto.InternalMessageInfo

func (m *ViewSpec) GetViewQuery() string {
	if m != nil {
		return m.ViewQuery
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.datacatalog.v1beta1.TableSourceType", TableSourceType_name, TableSourceType_value)
	proto.RegisterType((*BigQueryTableSpec)(nil), "google.cloud.datacatalog.v1beta1.BigQueryTableSpec")
	proto.RegisterType((*ViewSpec)(nil), "google.cloud.datacatalog.v1beta1.ViewSpec")
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1beta1/table_spec.proto", fileDescriptor_2effb41fca72136b)
}

var fileDescriptor_2effb41fca72136b = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x1b, 0xa1, 0xd0, 0x89, 0xfc, 0x31, 0x27, 0x0f, 0x09, 0xe6, 0xc9, 0x3c, 0xcc, 0xb2,
	0x76, 0xec, 0xd4, 0xda, 0x24, 0x4b, 0x51, 0xba, 0xae, 0x86, 0x45, 0x2c, 0xe3, 0xf8, 0x18, 0x16,
	0x36, 0x67, 0x5a, 0x47, 0xc5, 0xbf, 0xad, 0x7f, 0xac, 0x63, 0x38, 0x6e, 0x61, 0x81, 0x78, 0x1b,
	0xbe, 0x8f, 0xcf, 0x67, 0xbe, 0x8f, 0x87, 0x5d, 0xa9, 0x94, 0x4c, 0xc0, 0x11, 0x89, 0x5a, 0x4c,
	0x9d, 0x29, 0x37, 0x5c, 0x70, 0xc3, 0x13, 0x25, 0x9d, 0xa5, 0x3b, 0x01, 0xc3, 0x5d, 0xc7, 0xf0,
	0x49, 0x02, 0xd1, 0x5c, 0x83, 0xa0, 0x3a, 0x55, 0x46, 0x91, 0xfa, 0x16, 0xa1, 0x16, 0xa1, 0x3b,
	0x08, 0xcd, 0x90, 0xc6, 0x27, 0xc2, 0x15, 0x2f, 0x96, 0xfd, 0x05, 0xa4, 0xeb, 0x70, 0x83, 0x0f,
	0x34, 0x08, 0xf2, 0x86, 0x2b, 0x99, 0x4b, 0x2d, 0x52, 0x01, 0x91, 0x59, 0x6b, 0xa8, 0xa2, 0x3a,
	0x6a, 0x16, 0xdb, 0x2e, 0x3d, 0xe4, 0xa4, 0x5b, 0x8f, 0x25, 0xc3, 0xb5, 0x86, 0xa0, 0x64, 0xfe,
	0x06, 0xa4, 0x8b, 0x0b, 0xcb, 0x18, 0x56, 0xb6, 0x69, 0x35, 0x57, 0x47, 0xcd, 0xd3, 0x76, 0xeb,
	0xb0, 0x76, 0x14, 0xc3, 0x6a, 0xd3, 0x2e, 0xc8, 0x2f, 0xb3, 0x57, 0xe3, 0x12, 0xe7, 0x7f, 0x52,
	0x52, 0xc3, 0xd8, 0x4a, 0x3f, 0x36, 0xab, 0xd8, 0xb2, 0x85, 0xc0, 0x7e, 0x63, 0x77, 0x6b, 0xbd,
	0xe2, 0xd2, 0xbf, 0x5e, 0xe4, 0x02, 0xd7, 0xc2, 0x1b, 0xef, 0x81, 0x45, 0x83, 0xa7, 0x61, 0xd0,
	0x61, 0x51, 0x38, 0xee, 0xb1, 0x68, 0xf8, 0x38, 0xe8, 0xb1, 0x8e, 0x7f, 0xe7, 0xb3, 0xdb, 0xf2,
	0x11, 0xa9, 0xe0, 0x33, 0xcf, 0xef, 0xf6, 0x87, 0x2c, 0x18, 0x47, 0x23, 0x9f, 0x3d, 0x97, 0x73,
	0x84, 0xe0, 0xe2, 0x6f, 0x64, 0xf1, 0xf2, 0xb1, 0xa7, 0xf1, 0xb9, 0x50, 0xef, 0x7b, 0x57, 0xe8,
	0xa1, 0x97, 0xfb, 0x6c, 0x26, 0x55, 0xc2, 0x67, 0x92, 0xaa, 0x54, 0x3a, 0x12, 0x66, 0xf6, 0x4a,
	0xce, 0x76, 0xc4, 0x75, 0x3c, 0xdf, 0x7f, 0xdb, 0xeb, 0x9d, 0xec, 0x0b, 0xa1, 0xc9, 0x89, 0x45,
	0xaf, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x41, 0xe6, 0x4f, 0xb9, 0x15, 0x02, 0x00, 0x00,
}
