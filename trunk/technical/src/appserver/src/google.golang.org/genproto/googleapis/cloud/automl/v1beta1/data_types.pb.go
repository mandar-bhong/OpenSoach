// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/data_types.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
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

// `TypeCode` is used as a part of
// [DataType][google.cloud.automl.v1beta1.DataType].
//
// Each legal value of a DataType can be encoded to or decoded from a JSON
// value, using the encodings listed below, and definitions of which can be
// found at
//
// https:
// //developers.google.com/protocol-buffers
// // /docs/reference/google.protobuf#value.
type TypeCode int32

const (
	// Not specified. Should not be used.
	TypeCode_TYPE_CODE_UNSPECIFIED TypeCode = 0
	// Encoded as `number`, or the strings `"NaN"`, `"Infinity"`, or
	// `"-Infinity"`.
	TypeCode_FLOAT64 TypeCode = 3
	// Must be between 0AD and 9999AD. Encoded as `string` according to
	// [time_format][google.cloud.automl.v1beta1.DataType.time_format], or, if
	// that format is not set, then in RFC 3339 `date-time` format, where
	// `time-offset` = `"Z"` (e.g. 1985-04-12T23:20:50.52Z).
	TypeCode_TIMESTAMP TypeCode = 4
	// Encoded as `string`.
	TypeCode_STRING TypeCode = 6
	// Encoded as `list`, where the list elements are represented according to
	//
	// [list_element_type][google.cloud.automl.v1beta1.DataType.list_element_type].
	TypeCode_ARRAY TypeCode = 8
	// Encoded as `struct`, where field values are represented according to
	// [struct_type][google.cloud.automl.v1beta1.DataType.struct_type].
	TypeCode_STRUCT TypeCode = 9
	// Values of this type are not further understood by AutoML,
	// e.g. AutoML is unable to tell the order of values (as it could with
	// FLOAT64), or is unable to say if one value contains another (as it
	// could with STRING).
	// Encoded as `string` (bytes should be base64-encoded, as described in RFC
	// 4648, section 4).
	TypeCode_CATEGORY TypeCode = 10
)

var TypeCode_name = map[int32]string{
	0:  "TYPE_CODE_UNSPECIFIED",
	3:  "FLOAT64",
	4:  "TIMESTAMP",
	6:  "STRING",
	8:  "ARRAY",
	9:  "STRUCT",
	10: "CATEGORY",
}

var TypeCode_value = map[string]int32{
	"TYPE_CODE_UNSPECIFIED": 0,
	"FLOAT64":               3,
	"TIMESTAMP":             4,
	"STRING":                6,
	"ARRAY":                 8,
	"STRUCT":                9,
	"CATEGORY":              10,
}

func (x TypeCode) String() string {
	return proto.EnumName(TypeCode_name, int32(x))
}

func (TypeCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_43071a57be770d81, []int{0}
}

// Indicated the type of data that can be stored in a structured data entity
// (e.g. a table).
type DataType struct {
	// Details of DataType-s that need additional specification.
	//
	// Types that are valid to be assigned to Details:
	//	*DataType_ListElementType
	//	*DataType_StructType
	//	*DataType_TimeFormat
	Details isDataType_Details `protobuf_oneof:"details"`
	// Required. The [TypeCode][google.cloud.automl.v1beta1.TypeCode] for this type.
	TypeCode TypeCode `protobuf:"varint,1,opt,name=type_code,json=typeCode,proto3,enum=google.cloud.automl.v1beta1.TypeCode" json:"type_code,omitempty"`
	// If true, this DataType can also be `null`.
	Nullable             bool     `protobuf:"varint,4,opt,name=nullable,proto3" json:"nullable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataType) Reset()         { *m = DataType{} }
func (m *DataType) String() string { return proto.CompactTextString(m) }
func (*DataType) ProtoMessage()    {}
func (*DataType) Descriptor() ([]byte, []int) {
	return fileDescriptor_43071a57be770d81, []int{0}
}

func (m *DataType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataType.Unmarshal(m, b)
}
func (m *DataType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataType.Marshal(b, m, deterministic)
}
func (m *DataType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataType.Merge(m, src)
}
func (m *DataType) XXX_Size() int {
	return xxx_messageInfo_DataType.Size(m)
}
func (m *DataType) XXX_DiscardUnknown() {
	xxx_messageInfo_DataType.DiscardUnknown(m)
}

var xxx_messageInfo_DataType proto.InternalMessageInfo

type isDataType_Details interface {
	isDataType_Details()
}

type DataType_ListElementType struct {
	ListElementType *DataType `protobuf:"bytes,2,opt,name=list_element_type,json=listElementType,proto3,oneof"`
}

type DataType_StructType struct {
	StructType *StructType `protobuf:"bytes,3,opt,name=struct_type,json=structType,proto3,oneof"`
}

type DataType_TimeFormat struct {
	TimeFormat string `protobuf:"bytes,5,opt,name=time_format,json=timeFormat,proto3,oneof"`
}

func (*DataType_ListElementType) isDataType_Details() {}

func (*DataType_StructType) isDataType_Details() {}

func (*DataType_TimeFormat) isDataType_Details() {}

func (m *DataType) GetDetails() isDataType_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *DataType) GetListElementType() *DataType {
	if x, ok := m.GetDetails().(*DataType_ListElementType); ok {
		return x.ListElementType
	}
	return nil
}

func (m *DataType) GetStructType() *StructType {
	if x, ok := m.GetDetails().(*DataType_StructType); ok {
		return x.StructType
	}
	return nil
}

func (m *DataType) GetTimeFormat() string {
	if x, ok := m.GetDetails().(*DataType_TimeFormat); ok {
		return x.TimeFormat
	}
	return ""
}

func (m *DataType) GetTypeCode() TypeCode {
	if m != nil {
		return m.TypeCode
	}
	return TypeCode_TYPE_CODE_UNSPECIFIED
}

func (m *DataType) GetNullable() bool {
	if m != nil {
		return m.Nullable
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DataType) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DataType_ListElementType)(nil),
		(*DataType_StructType)(nil),
		(*DataType_TimeFormat)(nil),
	}
}

// `StructType` defines the DataType-s of a [STRUCT][google.cloud.automl.v1beta1.TypeCode.STRUCT] type.
type StructType struct {
	// Unordered map of struct field names to their data types.
	// Fields cannot be added or removed via Update. Their names and
	// data types are still mutable.
	Fields               map[string]*DataType `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StructType) Reset()         { *m = StructType{} }
func (m *StructType) String() string { return proto.CompactTextString(m) }
func (*StructType) ProtoMessage()    {}
func (*StructType) Descriptor() ([]byte, []int) {
	return fileDescriptor_43071a57be770d81, []int{1}
}

func (m *StructType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructType.Unmarshal(m, b)
}
func (m *StructType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructType.Marshal(b, m, deterministic)
}
func (m *StructType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructType.Merge(m, src)
}
func (m *StructType) XXX_Size() int {
	return xxx_messageInfo_StructType.Size(m)
}
func (m *StructType) XXX_DiscardUnknown() {
	xxx_messageInfo_StructType.DiscardUnknown(m)
}

var xxx_messageInfo_StructType proto.InternalMessageInfo

func (m *StructType) GetFields() map[string]*DataType {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.automl.v1beta1.TypeCode", TypeCode_name, TypeCode_value)
	proto.RegisterType((*DataType)(nil), "google.cloud.automl.v1beta1.DataType")
	proto.RegisterType((*StructType)(nil), "google.cloud.automl.v1beta1.StructType")
	proto.RegisterMapType((map[string]*DataType)(nil), "google.cloud.automl.v1beta1.StructType.FieldsEntry")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/data_types.proto", fileDescriptor_43071a57be770d81)
}

var fileDescriptor_43071a57be770d81 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x5f, 0x8b, 0xd3, 0x40,
	0x10, 0xc0, 0x2f, 0xed, 0xb5, 0x97, 0x4c, 0xfc, 0x13, 0x17, 0x84, 0x5e, 0x4f, 0xb0, 0x1e, 0x8a,
	0x45, 0x24, 0xa1, 0x77, 0x22, 0xe2, 0x3d, 0xa5, 0x69, 0x5a, 0xab, 0xf6, 0x5a, 0xd2, 0x9c, 0x50,
	0x29, 0xc4, 0x6d, 0xb3, 0x0d, 0xc1, 0x6d, 0xb6, 0x24, 0x9b, 0xe3, 0xfa, 0xee, 0x97, 0x12, 0xbf,
	0x81, 0x9f, 0x4a, 0xb2, 0xd9, 0x3b, 0x7d, 0x90, 0xea, 0xbd, 0xed, 0xcc, 0xfc, 0xe6, 0xb7, 0x33,
	0x24, 0x0b, 0x2f, 0x23, 0xc6, 0x22, 0x4a, 0xac, 0x25, 0x65, 0x79, 0x68, 0xe1, 0x9c, 0xb3, 0x35,
	0xb5, 0x2e, 0x3b, 0x0b, 0xc2, 0x71, 0xc7, 0x0a, 0x31, 0xc7, 0x01, 0xdf, 0x6e, 0x48, 0x66, 0x6e,
	0x52, 0xc6, 0x19, 0x3a, 0x2a, 0x69, 0x53, 0xd0, 0x66, 0x49, 0x9b, 0x92, 0x6e, 0x3e, 0x92, 0x2a,
	0xbc, 0x89, 0x2d, 0x9c, 0x24, 0x8c, 0x63, 0x1e, 0xb3, 0x44, 0xb6, 0x36, 0x9f, 0xee, 0xba, 0x28,
	0x66, 0x92, 0xea, 0xec, 0xa2, 0x38, 0xb9, 0xe2, 0x01, 0xb9, 0xe2, 0x29, 0x5e, 0x16, 0x66, 0xd9,
	0x72, 0x28, 0x5b, 0x44, 0xb4, 0xc8, 0x57, 0x16, 0x4e, 0xb6, 0x65, 0xe9, 0xf8, 0x7b, 0x05, 0xd4,
	0x1e, 0xe6, 0xd8, 0xdf, 0x6e, 0x08, 0x9a, 0xc2, 0x03, 0x1a, 0x67, 0x3c, 0x20, 0x94, 0xac, 0x49,
	0xc2, 0xc5, 0x5e, 0x8d, 0x4a, 0x4b, 0x69, 0xeb, 0x27, 0xcf, 0xcc, 0x1d, 0x7b, 0x99, 0xd7, 0x86,
	0x77, 0x7b, 0xde, 0xfd, 0xc2, 0xe0, 0x96, 0x02, 0x21, 0x7d, 0x0f, 0x7a, 0xc6, 0xd3, 0x7c, 0x29,
	0x75, 0x55, 0xa1, 0x7b, 0xbe, 0x53, 0x37, 0x15, 0xbc, 0x14, 0x42, 0x76, 0x13, 0xa1, 0x27, 0xa0,
	0xf3, 0x78, 0x4d, 0x82, 0x15, 0x4b, 0xd7, 0x98, 0x37, 0x6a, 0x2d, 0xa5, 0xad, 0x15, 0x48, 0x91,
	0xec, 0x8b, 0x1c, 0xea, 0x82, 0x56, 0xdc, 0x13, 0x2c, 0x59, 0x48, 0x1a, 0x4a, 0x4b, 0x69, 0xdf,
	0xfb, 0xc7, 0xec, 0x85, 0xd8, 0x61, 0x21, 0xf1, 0x54, 0x2e, 0x4f, 0xa8, 0x09, 0x6a, 0x92, 0x53,
	0x8a, 0x17, 0x94, 0x34, 0xf6, 0x5b, 0x4a, 0x5b, 0xf5, 0x6e, 0xe2, 0xae, 0x06, 0x07, 0x21, 0xe1,
	0x38, 0xa6, 0xd9, 0xf1, 0x0f, 0x05, 0xe0, 0xf7, 0xa8, 0xe8, 0x03, 0xd4, 0x57, 0x31, 0xa1, 0x61,
	0xd6, 0x50, 0x5a, 0xd5, 0xb6, 0x7e, 0x72, 0xfa, 0x9f, 0x3b, 0x9a, 0x7d, 0xd1, 0xe5, 0x26, 0x3c,
	0xdd, 0x7a, 0x52, 0xd1, 0xfc, 0x02, 0xfa, 0x1f, 0x69, 0x64, 0x40, 0xf5, 0x2b, 0xd9, 0x8a, 0x7d,
	0x34, 0xaf, 0x38, 0xa2, 0x33, 0xa8, 0x5d, 0x62, 0x9a, 0xdf, 0xee, 0xfb, 0x78, 0x65, 0xcf, 0xdb,
	0xca, 0x1b, 0xe5, 0x45, 0x0a, 0xea, 0xf5, 0xea, 0xe8, 0x10, 0x1e, 0xfa, 0xb3, 0x89, 0x1b, 0x38,
	0xe3, 0x9e, 0x1b, 0x5c, 0x9c, 0x4f, 0x27, 0xae, 0x33, 0xec, 0x0f, 0xdd, 0x9e, 0xb1, 0x87, 0x74,
	0x38, 0xe8, 0x7f, 0x1c, 0xdb, 0xfe, 0xeb, 0x57, 0x46, 0x15, 0xdd, 0x05, 0xcd, 0x1f, 0x8e, 0xdc,
	0xa9, 0x6f, 0x8f, 0x26, 0xc6, 0x3e, 0x02, 0xa8, 0x4f, 0x7d, 0x6f, 0x78, 0x3e, 0x30, 0xea, 0x48,
	0x83, 0x9a, 0xed, 0x79, 0xf6, 0xcc, 0x50, 0x65, 0xfa, 0xc2, 0xf1, 0x0d, 0x0d, 0xdd, 0x01, 0xd5,
	0xb1, 0x7d, 0x77, 0x30, 0xf6, 0x66, 0x06, 0x74, 0xbf, 0x29, 0xf0, 0x78, 0xc9, 0xd6, 0xbb, 0x66,
	0x9d, 0x28, 0x9f, 0x6d, 0x59, 0x8e, 0x18, 0xc5, 0x49, 0x64, 0xb2, 0x34, 0xb2, 0x22, 0x92, 0x88,
	0xff, 0xd5, 0x2a, 0x4b, 0x78, 0x13, 0x67, 0x7f, 0x7d, 0x00, 0x67, 0x65, 0xf8, 0xb3, 0x72, 0x34,
	0x10, 0xe0, 0xdc, 0x29, 0xa0, 0xb9, 0x9d, 0x73, 0x36, 0xa2, 0xf3, 0x4f, 0x25, 0xb4, 0xa8, 0x0b,
	0xd7, 0xe9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x86, 0x13, 0xbf, 0xda, 0x03, 0x00, 0x00,
}
