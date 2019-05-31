// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/bigquery/v2/standard_sql.proto

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

type StandardSqlDataType_TypeKind int32

const (
	// Invalid type.
	StandardSqlDataType_TYPE_KIND_UNSPECIFIED StandardSqlDataType_TypeKind = 0
	// Encoded as a string in decimal format.
	StandardSqlDataType_INT64 StandardSqlDataType_TypeKind = 2
	// Encoded as a boolean "false" or "true".
	StandardSqlDataType_BOOL StandardSqlDataType_TypeKind = 5
	// Encoded as a number, or string "NaN", "Infinity" or "-Infinity".
	StandardSqlDataType_FLOAT64 StandardSqlDataType_TypeKind = 7
	// Encoded as a string value.
	StandardSqlDataType_STRING StandardSqlDataType_TypeKind = 8
	// Encoded as a base64 string per RFC 4648, section 4.
	StandardSqlDataType_BYTES StandardSqlDataType_TypeKind = 9
	// Encoded as an RFC 3339 timestamp with mandatory "Z" time zone string:
	// 1985-04-12T23:20:50.52Z
	StandardSqlDataType_TIMESTAMP StandardSqlDataType_TypeKind = 19
	// Encoded as RFC 3339 full-date format string: 1985-04-12
	StandardSqlDataType_DATE StandardSqlDataType_TypeKind = 10
	// Encoded as RFC 3339 partial-time format string: 23:20:50.52
	StandardSqlDataType_TIME StandardSqlDataType_TypeKind = 20
	// Encoded as RFC 3339 full-date "T" partial-time: 1985-04-12T23:20:50.52
	StandardSqlDataType_DATETIME StandardSqlDataType_TypeKind = 21
	// Encoded as WKT
	StandardSqlDataType_GEOGRAPHY StandardSqlDataType_TypeKind = 22
	// Encoded as a decimal string.
	StandardSqlDataType_NUMERIC StandardSqlDataType_TypeKind = 23
	// Encoded as a list with types matching Type.array_type.
	StandardSqlDataType_ARRAY StandardSqlDataType_TypeKind = 16
	// Encoded as a list with fields of type Type.struct_type[i]. List is used
	// because a JSON object cannot have duplicate field names.
	StandardSqlDataType_STRUCT StandardSqlDataType_TypeKind = 17
)

var StandardSqlDataType_TypeKind_name = map[int32]string{
	0:  "TYPE_KIND_UNSPECIFIED",
	2:  "INT64",
	5:  "BOOL",
	7:  "FLOAT64",
	8:  "STRING",
	9:  "BYTES",
	19: "TIMESTAMP",
	10: "DATE",
	20: "TIME",
	21: "DATETIME",
	22: "GEOGRAPHY",
	23: "NUMERIC",
	16: "ARRAY",
	17: "STRUCT",
}

var StandardSqlDataType_TypeKind_value = map[string]int32{
	"TYPE_KIND_UNSPECIFIED": 0,
	"INT64":                 2,
	"BOOL":                  5,
	"FLOAT64":               7,
	"STRING":                8,
	"BYTES":                 9,
	"TIMESTAMP":             19,
	"DATE":                  10,
	"TIME":                  20,
	"DATETIME":              21,
	"GEOGRAPHY":             22,
	"NUMERIC":               23,
	"ARRAY":                 16,
	"STRUCT":                17,
}

func (x StandardSqlDataType_TypeKind) String() string {
	return proto.EnumName(StandardSqlDataType_TypeKind_name, int32(x))
}

func (StandardSqlDataType_TypeKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17cd17721212b1f9, []int{0, 0}
}

// The type of a variable, e.g., a function argument.
// Examples:
// INT64: {type_kind="INT64"}
// ARRAY<STRING>: {type_kind="ARRAY", array_element_type="STRING"}
// STRUCT<x STRING, y ARRAY<DATE>>:
//   {type_kind="STRUCT",
//    struct_type={fields=[
//      {name="x", type={type_kind="STRING"}},
//      {name="y", type={type_kind="ARRAY", array_element_type="DATE"}}
//    ]}}
type StandardSqlDataType struct {
	// Required. The top level type of this field.
	// Can be any standard SQL data type (e.g., "INT64", "DATE", "ARRAY").
	TypeKind StandardSqlDataType_TypeKind `protobuf:"varint,1,opt,name=type_kind,json=typeKind,proto3,enum=google.cloud.bigquery.v2.StandardSqlDataType_TypeKind" json:"type_kind,omitempty"`
	// Types that are valid to be assigned to SubType:
	//	*StandardSqlDataType_ArrayElementType
	//	*StandardSqlDataType_StructType
	SubType              isStandardSqlDataType_SubType `protobuf_oneof:"sub_type"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *StandardSqlDataType) Reset()         { *m = StandardSqlDataType{} }
func (m *StandardSqlDataType) String() string { return proto.CompactTextString(m) }
func (*StandardSqlDataType) ProtoMessage()    {}
func (*StandardSqlDataType) Descriptor() ([]byte, []int) {
	return fileDescriptor_17cd17721212b1f9, []int{0}
}

func (m *StandardSqlDataType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StandardSqlDataType.Unmarshal(m, b)
}
func (m *StandardSqlDataType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StandardSqlDataType.Marshal(b, m, deterministic)
}
func (m *StandardSqlDataType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StandardSqlDataType.Merge(m, src)
}
func (m *StandardSqlDataType) XXX_Size() int {
	return xxx_messageInfo_StandardSqlDataType.Size(m)
}
func (m *StandardSqlDataType) XXX_DiscardUnknown() {
	xxx_messageInfo_StandardSqlDataType.DiscardUnknown(m)
}

var xxx_messageInfo_StandardSqlDataType proto.InternalMessageInfo

func (m *StandardSqlDataType) GetTypeKind() StandardSqlDataType_TypeKind {
	if m != nil {
		return m.TypeKind
	}
	return StandardSqlDataType_TYPE_KIND_UNSPECIFIED
}

type isStandardSqlDataType_SubType interface {
	isStandardSqlDataType_SubType()
}

type StandardSqlDataType_ArrayElementType struct {
	ArrayElementType *StandardSqlDataType `protobuf:"bytes,2,opt,name=array_element_type,json=arrayElementType,proto3,oneof"`
}

type StandardSqlDataType_StructType struct {
	StructType *StandardSqlStructType `protobuf:"bytes,3,opt,name=struct_type,json=structType,proto3,oneof"`
}

func (*StandardSqlDataType_ArrayElementType) isStandardSqlDataType_SubType() {}

func (*StandardSqlDataType_StructType) isStandardSqlDataType_SubType() {}

func (m *StandardSqlDataType) GetSubType() isStandardSqlDataType_SubType {
	if m != nil {
		return m.SubType
	}
	return nil
}

func (m *StandardSqlDataType) GetArrayElementType() *StandardSqlDataType {
	if x, ok := m.GetSubType().(*StandardSqlDataType_ArrayElementType); ok {
		return x.ArrayElementType
	}
	return nil
}

func (m *StandardSqlDataType) GetStructType() *StandardSqlStructType {
	if x, ok := m.GetSubType().(*StandardSqlDataType_StructType); ok {
		return x.StructType
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StandardSqlDataType) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StandardSqlDataType_ArrayElementType)(nil),
		(*StandardSqlDataType_StructType)(nil),
	}
}

// A field or a column.
type StandardSqlField struct {
	// Optional. The name of this field. Can be absent for struct fields.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The type of this parameter. Absent if not explicitly
	// specified (e.g., CREATE FUNCTION statement can omit the return type;
	// in this case the output parameter does not have this "type" field).
	Type                 *StandardSqlDataType `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StandardSqlField) Reset()         { *m = StandardSqlField{} }
func (m *StandardSqlField) String() string { return proto.CompactTextString(m) }
func (*StandardSqlField) ProtoMessage()    {}
func (*StandardSqlField) Descriptor() ([]byte, []int) {
	return fileDescriptor_17cd17721212b1f9, []int{1}
}

func (m *StandardSqlField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StandardSqlField.Unmarshal(m, b)
}
func (m *StandardSqlField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StandardSqlField.Marshal(b, m, deterministic)
}
func (m *StandardSqlField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StandardSqlField.Merge(m, src)
}
func (m *StandardSqlField) XXX_Size() int {
	return xxx_messageInfo_StandardSqlField.Size(m)
}
func (m *StandardSqlField) XXX_DiscardUnknown() {
	xxx_messageInfo_StandardSqlField.DiscardUnknown(m)
}

var xxx_messageInfo_StandardSqlField proto.InternalMessageInfo

func (m *StandardSqlField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StandardSqlField) GetType() *StandardSqlDataType {
	if m != nil {
		return m.Type
	}
	return nil
}

type StandardSqlStructType struct {
	Fields               []*StandardSqlField `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StandardSqlStructType) Reset()         { *m = StandardSqlStructType{} }
func (m *StandardSqlStructType) String() string { return proto.CompactTextString(m) }
func (*StandardSqlStructType) ProtoMessage()    {}
func (*StandardSqlStructType) Descriptor() ([]byte, []int) {
	return fileDescriptor_17cd17721212b1f9, []int{2}
}

func (m *StandardSqlStructType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StandardSqlStructType.Unmarshal(m, b)
}
func (m *StandardSqlStructType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StandardSqlStructType.Marshal(b, m, deterministic)
}
func (m *StandardSqlStructType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StandardSqlStructType.Merge(m, src)
}
func (m *StandardSqlStructType) XXX_Size() int {
	return xxx_messageInfo_StandardSqlStructType.Size(m)
}
func (m *StandardSqlStructType) XXX_DiscardUnknown() {
	xxx_messageInfo_StandardSqlStructType.DiscardUnknown(m)
}

var xxx_messageInfo_StandardSqlStructType proto.InternalMessageInfo

func (m *StandardSqlStructType) GetFields() []*StandardSqlField {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.bigquery.v2.StandardSqlDataType_TypeKind", StandardSqlDataType_TypeKind_name, StandardSqlDataType_TypeKind_value)
	proto.RegisterType((*StandardSqlDataType)(nil), "google.cloud.bigquery.v2.StandardSqlDataType")
	proto.RegisterType((*StandardSqlField)(nil), "google.cloud.bigquery.v2.StandardSqlField")
	proto.RegisterType((*StandardSqlStructType)(nil), "google.cloud.bigquery.v2.StandardSqlStructType")
}

func init() {
	proto.RegisterFile("google/cloud/bigquery/v2/standard_sql.proto", fileDescriptor_17cd17721212b1f9)
}

var fileDescriptor_17cd17721212b1f9 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xd7, 0xad, 0xeb, 0xd2, 0x57, 0x40, 0xc6, 0xa3, 0x50, 0xd0, 0x0e, 0x55, 0x4f, 0x15,
	0x88, 0x44, 0x2a, 0x68, 0x17, 0x2e, 0x24, 0x6d, 0xda, 0x45, 0x5b, 0xd3, 0xc8, 0x49, 0x0f, 0x05,
	0xa1, 0xc8, 0x6d, 0x4c, 0x14, 0x91, 0xda, 0x69, 0x92, 0x4e, 0xea, 0xd7, 0xe3, 0x03, 0xf1, 0x19,
	0x90, 0xdd, 0x14, 0x4d, 0x62, 0x93, 0x06, 0x97, 0xe8, 0xd9, 0xef, 0xff, 0xff, 0x3d, 0x3f, 0x3b,
	0x0f, 0xde, 0xc5, 0x42, 0xc4, 0x29, 0x33, 0x56, 0xa9, 0xd8, 0x46, 0xc6, 0x32, 0x89, 0x37, 0x5b,
	0x96, 0xef, 0x8c, 0xdb, 0x81, 0x51, 0x94, 0x94, 0x47, 0x34, 0x8f, 0xc2, 0x62, 0x93, 0xea, 0x59,
	0x2e, 0x4a, 0x81, 0x3b, 0x7b, 0xb1, 0xae, 0xc4, 0xfa, 0x41, 0xac, 0xdf, 0x0e, 0xde, 0x5c, 0x54,
	0x18, 0x9a, 0x25, 0x06, 0xe5, 0x5c, 0x94, 0xb4, 0x4c, 0x04, 0x2f, 0xf6, 0xbe, 0xde, 0xaf, 0x13,
	0x38, 0xf7, 0x2b, 0x9c, 0xbf, 0x49, 0x47, 0xb4, 0xa4, 0xc1, 0x2e, 0x63, 0xd8, 0x87, 0x66, 0xb9,
	0xcb, 0x58, 0xf8, 0x23, 0xe1, 0x51, 0xa7, 0xd6, 0xad, 0xf5, 0x9f, 0x0d, 0x2e, 0xf5, 0x87, 0x6a,
	0xe8, 0xf7, 0x10, 0x74, 0xf9, 0xb9, 0x4e, 0x78, 0x44, 0xb4, 0xb2, 0x8a, 0xf0, 0x37, 0xc0, 0x34,
	0xcf, 0xe9, 0x2e, 0x64, 0x29, 0x5b, 0x33, 0x5e, 0x86, 0x32, 0xd3, 0x39, 0xee, 0xd6, 0xfa, 0xad,
	0xc1, 0xfb, 0x7f, 0xa2, 0x5f, 0x1d, 0x11, 0xa4, 0x50, 0xf6, 0x9e, 0xa4, 0xce, 0x4c, 0xa0, 0x55,
	0x94, 0xf9, 0x76, 0x55, 0x71, 0x4f, 0x14, 0xd7, 0x78, 0x14, 0xd7, 0x57, 0xbe, 0x8a, 0x0c, 0xc5,
	0x9f, 0x55, 0xef, 0x67, 0x0d, 0xb4, 0x43, 0x27, 0xf8, 0x35, 0xb4, 0x83, 0x85, 0x67, 0x87, 0xd7,
	0x8e, 0x3b, 0x0a, 0xe7, 0xae, 0xef, 0xd9, 0x43, 0x67, 0xec, 0xd8, 0x23, 0x74, 0x84, 0x9b, 0x70,
	0xea, 0xb8, 0xc1, 0xe5, 0x47, 0x74, 0x8c, 0x35, 0xa8, 0x5b, 0xb3, 0xd9, 0x0d, 0x3a, 0xc5, 0x2d,
	0x38, 0x1b, 0xdf, 0xcc, 0x4c, 0xb9, 0x7d, 0x86, 0x01, 0x1a, 0x7e, 0x40, 0x1c, 0x77, 0x82, 0x34,
	0xa9, 0xb6, 0x16, 0x81, 0xed, 0xa3, 0x26, 0x7e, 0x0a, 0xcd, 0xc0, 0x99, 0xda, 0x7e, 0x60, 0x4e,
	0x3d, 0x74, 0x2e, 0xcd, 0x23, 0x33, 0xb0, 0x11, 0xc8, 0x48, 0x26, 0xd0, 0x0b, 0xfc, 0x04, 0x34,
	0xb9, 0xa7, 0x56, 0x6d, 0x69, 0x98, 0xd8, 0xb3, 0x09, 0x31, 0xbd, 0xab, 0x05, 0x7a, 0x29, 0x6b,
	0xb8, 0xf3, 0xa9, 0x4d, 0x9c, 0x21, 0x7a, 0x25, 0xb9, 0x26, 0x21, 0xe6, 0x02, 0xa1, 0xaa, 0xdc,
	0x7c, 0x18, 0xa0, 0xe7, 0x16, 0x80, 0x56, 0x6c, 0x97, 0xea, 0x56, 0x7a, 0x09, 0xa0, 0x3b, 0x7d,
	0x8f, 0x13, 0x96, 0x46, 0x18, 0x43, 0x9d, 0xd3, 0x35, 0x53, 0xef, 0xdc, 0x24, 0x2a, 0xc6, 0x26,
	0xd4, 0xff, 0xfb, 0x75, 0x88, 0xb2, 0xf6, 0xbe, 0x42, 0xfb, 0xde, 0x2b, 0xc6, 0x16, 0x34, 0xbe,
	0xcb, 0xc2, 0x45, 0xa7, 0xd6, 0x3d, 0xe9, 0xb7, 0x06, 0x6f, 0x1f, 0x45, 0x57, 0x67, 0x25, 0x95,
	0xd3, 0xca, 0xe1, 0x62, 0x25, 0xd6, 0x0f, 0x1a, 0xad, 0xbb, 0x5d, 0x7a, 0xf2, 0x57, 0xff, 0xf2,
	0xb9, 0xd2, 0xc6, 0x22, 0xa5, 0x3c, 0xd6, 0x45, 0x1e, 0x1b, 0x31, 0xe3, 0x6a, 0x0c, 0x8c, 0x7d,
	0x8a, 0x66, 0x49, 0xf1, 0xf7, 0xb8, 0x7d, 0x3a, 0xc4, 0xcb, 0x86, 0x12, 0x7f, 0xf8, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0xdf, 0x4e, 0x6d, 0x31, 0x9a, 0x03, 0x00, 0x00,
}
