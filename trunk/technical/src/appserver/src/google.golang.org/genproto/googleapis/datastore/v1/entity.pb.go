// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/datastore/v1/entity.proto

package datastore

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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

// A partition ID identifies a grouping of entities. The grouping is always
// by project and namespace, however the namespace ID may be empty.
//
// A partition ID contains several dimensions:
// project ID and namespace ID.
//
// Partition dimensions:
//
// - May be `""`.
// - Must be valid UTF-8 bytes.
// - Must have values that match regex `[A-Za-z\d\.\-_]{1,100}`
// If the value of any dimension matches regex `__.*__`, the partition is
// reserved/read-only.
// A reserved/read-only partition ID is forbidden in certain documented
// contexts.
//
// Foreign partition IDs (in which the project ID does
// not match the context project ID ) are discouraged.
// Reads and writes of foreign partition IDs may fail if the project is not in
// an active state.
type PartitionId struct {
	// The ID of the project to which the entities belong.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// If not empty, the ID of the namespace to which the entities belong.
	NamespaceId          string   `protobuf:"bytes,4,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionId) Reset()         { *m = PartitionId{} }
func (m *PartitionId) String() string { return proto.CompactTextString(m) }
func (*PartitionId) ProtoMessage()    {}
func (*PartitionId) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbfdafa45100300, []int{0}
}

func (m *PartitionId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionId.Unmarshal(m, b)
}
func (m *PartitionId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionId.Marshal(b, m, deterministic)
}
func (m *PartitionId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionId.Merge(m, src)
}
func (m *PartitionId) XXX_Size() int {
	return xxx_messageInfo_PartitionId.Size(m)
}
func (m *PartitionId) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionId.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionId proto.InternalMessageInfo

func (m *PartitionId) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *PartitionId) GetNamespaceId() string {
	if m != nil {
		return m.NamespaceId
	}
	return ""
}

// A unique identifier for an entity.
// If a key's partition ID or any of its path kinds or names are
// reserved/read-only, the key is reserved/read-only.
// A reserved/read-only key is forbidden in certain documented contexts.
type Key struct {
	// Entities are partitioned into subsets, currently identified by a project
	// ID and namespace ID.
	// Queries are scoped to a single partition.
	PartitionId *PartitionId `protobuf:"bytes,1,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	// The entity path.
	// An entity path consists of one or more elements composed of a kind and a
	// string or numerical identifier, which identify entities. The first
	// element identifies a _root entity_, the second element identifies
	// a _child_ of the root entity, the third element identifies a child of the
	// second entity, and so forth. The entities identified by all prefixes of
	// the path are called the element's _ancestors_.
	//
	// An entity path is always fully complete: *all* of the entity's ancestors
	// are required to be in the path along with the entity identifier itself.
	// The only exception is that in some documented cases, the identifier in the
	// last path element (for the entity) itself may be omitted. For example,
	// the last path element of the key of `Mutation.insert` may have no
	// identifier.
	//
	// A path can never be empty, and a path can have at most 100 elements.
	Path                 []*Key_PathElement `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbfdafa45100300, []int{1}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetPartitionId() *PartitionId {
	if m != nil {
		return m.PartitionId
	}
	return nil
}

func (m *Key) GetPath() []*Key_PathElement {
	if m != nil {
		return m.Path
	}
	return nil
}

// A (kind, ID/name) pair used to construct a key path.
//
// If either name or ID is set, the element is complete.
// If neither is set, the element is incomplete.
type Key_PathElement struct {
	// The kind of the entity.
	// A kind matching regex `__.*__` is reserved/read-only.
	// A kind must not contain more than 1500 bytes when UTF-8 encoded.
	// Cannot be `""`.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// The type of ID.
	//
	// Types that are valid to be assigned to IdType:
	//	*Key_PathElement_Id
	//	*Key_PathElement_Name
	IdType               isKey_PathElement_IdType `protobuf_oneof:"id_type"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Key_PathElement) Reset()         { *m = Key_PathElement{} }
func (m *Key_PathElement) String() string { return proto.CompactTextString(m) }
func (*Key_PathElement) ProtoMessage()    {}
func (*Key_PathElement) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbfdafa45100300, []int{1, 0}
}

func (m *Key_PathElement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key_PathElement.Unmarshal(m, b)
}
func (m *Key_PathElement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key_PathElement.Marshal(b, m, deterministic)
}
func (m *Key_PathElement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key_PathElement.Merge(m, src)
}
func (m *Key_PathElement) XXX_Size() int {
	return xxx_messageInfo_Key_PathElement.Size(m)
}
func (m *Key_PathElement) XXX_DiscardUnknown() {
	xxx_messageInfo_Key_PathElement.DiscardUnknown(m)
}

var xxx_messageInfo_Key_PathElement proto.InternalMessageInfo

func (m *Key_PathElement) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

type isKey_PathElement_IdType interface {
	isKey_PathElement_IdType()
}

type Key_PathElement_Id struct {
	Id int64 `protobuf:"varint,2,opt,name=id,proto3,oneof"`
}

type Key_PathElement_Name struct {
	Name string `protobuf:"bytes,3,opt,name=name,proto3,oneof"`
}

func (*Key_PathElement_Id) isKey_PathElement_IdType() {}

func (*Key_PathElement_Name) isKey_PathElement_IdType() {}

func (m *Key_PathElement) GetIdType() isKey_PathElement_IdType {
	if m != nil {
		return m.IdType
	}
	return nil
}

func (m *Key_PathElement) GetId() int64 {
	if x, ok := m.GetIdType().(*Key_PathElement_Id); ok {
		return x.Id
	}
	return 0
}

func (m *Key_PathElement) GetName() string {
	if x, ok := m.GetIdType().(*Key_PathElement_Name); ok {
		return x.Name
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Key_PathElement) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Key_PathElement_Id)(nil),
		(*Key_PathElement_Name)(nil),
	}
}

// An array value.
type ArrayValue struct {
	// Values in the array.
	// The order of this array may not be preserved if it contains a mix of
	// indexed and unindexed values.
	Values               []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrayValue) Reset()         { *m = ArrayValue{} }
func (m *ArrayValue) String() string { return proto.CompactTextString(m) }
func (*ArrayValue) ProtoMessage()    {}
func (*ArrayValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbfdafa45100300, []int{2}
}

func (m *ArrayValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayValue.Unmarshal(m, b)
}
func (m *ArrayValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayValue.Marshal(b, m, deterministic)
}
func (m *ArrayValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayValue.Merge(m, src)
}
func (m *ArrayValue) XXX_Size() int {
	return xxx_messageInfo_ArrayValue.Size(m)
}
func (m *ArrayValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayValue.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayValue proto.InternalMessageInfo

func (m *ArrayValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// A message that can hold any of the supported value types and associated
// metadata.
type Value struct {
	// Must have a value set.
	//
	// Types that are valid to be assigned to ValueType:
	//	*Value_NullValue
	//	*Value_BooleanValue
	//	*Value_IntegerValue
	//	*Value_DoubleValue
	//	*Value_TimestampValue
	//	*Value_KeyValue
	//	*Value_StringValue
	//	*Value_BlobValue
	//	*Value_GeoPointValue
	//	*Value_EntityValue
	//	*Value_ArrayValue
	ValueType isValue_ValueType `protobuf_oneof:"value_type"`
	// The `meaning` field should only be populated for backwards compatibility.
	Meaning int32 `protobuf:"varint,14,opt,name=meaning,proto3" json:"meaning,omitempty"`
	// If the value should be excluded from all indexes including those defined
	// explicitly.
	ExcludeFromIndexes   bool     `protobuf:"varint,19,opt,name=exclude_from_indexes,json=excludeFromIndexes,proto3" json:"exclude_from_indexes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbfdafa45100300, []int{3}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_ValueType interface {
	isValue_ValueType()
}

type Value_NullValue struct {
	NullValue _struct.NullValue `protobuf:"varint,11,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}

type Value_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,1,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type Value_IntegerValue struct {
	IntegerValue int64 `protobuf:"varint,2,opt,name=integer_value,json=integerValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_TimestampValue struct {
	TimestampValue *timestamp.Timestamp `protobuf:"bytes,10,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type Value_KeyValue struct {
	KeyValue *Key `protobuf:"bytes,5,opt,name=key_value,json=keyValue,proto3,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,17,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BlobValue struct {
	BlobValue []byte `protobuf:"bytes,18,opt,name=blob_value,json=blobValue,proto3,oneof"`
}

type Value_GeoPointValue struct {
	GeoPointValue *latlng.LatLng `protobuf:"bytes,8,opt,name=geo_point_value,json=geoPointValue,proto3,oneof"`
}

type Value_EntityValue struct {
	EntityValue *Entity `protobuf:"bytes,6,opt,name=entity_value,json=entityValue,proto3,oneof"`
}

type Value_ArrayValue struct {
	ArrayValue *ArrayValue `protobuf:"bytes,9,opt,name=array_value,json=arrayValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_ValueType() {}

func (*Value_BooleanValue) isValue_ValueType() {}

func (*Value_IntegerValue) isValue_ValueType() {}

func (*Value_DoubleValue) isValue_ValueType() {}

func (*Value_TimestampValue) isValue_ValueType() {}

func (*Value_KeyValue) isValue_ValueType() {}

func (*Value_StringValue) isValue_ValueType() {}

func (*Value_BlobValue) isValue_ValueType() {}

func (*Value_GeoPointValue) isValue_ValueType() {}

func (*Value_EntityValue) isValue_ValueType() {}

func (*Value_ArrayValue) isValue_ValueType() {}

func (m *Value) GetValueType() isValue_ValueType {
	if m != nil {
		return m.ValueType
	}
	return nil
}

func (m *Value) GetNullValue() _struct.NullValue {
	if x, ok := m.GetValueType().(*Value_NullValue); ok {
		return x.NullValue
	}
	return _struct.NullValue_NULL_VALUE
}

func (m *Value) GetBooleanValue() bool {
	if x, ok := m.GetValueType().(*Value_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (m *Value) GetIntegerValue() int64 {
	if x, ok := m.GetValueType().(*Value_IntegerValue); ok {
		return x.IntegerValue
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if x, ok := m.GetValueType().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Value) GetTimestampValue() *timestamp.Timestamp {
	if x, ok := m.GetValueType().(*Value_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (m *Value) GetKeyValue() *Key {
	if x, ok := m.GetValueType().(*Value_KeyValue); ok {
		return x.KeyValue
	}
	return nil
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetValueType().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBlobValue() []byte {
	if x, ok := m.GetValueType().(*Value_BlobValue); ok {
		return x.BlobValue
	}
	return nil
}

func (m *Value) GetGeoPointValue() *latlng.LatLng {
	if x, ok := m.GetValueType().(*Value_GeoPointValue); ok {
		return x.GeoPointValue
	}
	return nil
}

func (m *Value) GetEntityValue() *Entity {
	if x, ok := m.GetValueType().(*Value_EntityValue); ok {
		return x.EntityValue
	}
	return nil
}

func (m *Value) GetArrayValue() *ArrayValue {
	if x, ok := m.GetValueType().(*Value_ArrayValue); ok {
		return x.ArrayValue
	}
	return nil
}

func (m *Value) GetMeaning() int32 {
	if m != nil {
		return m.Meaning
	}
	return 0
}

func (m *Value) GetExcludeFromIndexes() bool {
	if m != nil {
		return m.ExcludeFromIndexes
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Value_NullValue)(nil),
		(*Value_BooleanValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_TimestampValue)(nil),
		(*Value_KeyValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BlobValue)(nil),
		(*Value_GeoPointValue)(nil),
		(*Value_EntityValue)(nil),
		(*Value_ArrayValue)(nil),
	}
}

// A Datastore data object.
//
// An entity is limited to 1 megabyte when stored. That _roughly_
// corresponds to a limit of 1 megabyte for the serialized form of this
// message.
type Entity struct {
	// The entity's key.
	//
	// An entity must have a key, unless otherwise documented (for example,
	// an entity in `Value.entity_value` may have no key).
	// An entity's kind is its key path's last element's kind,
	// or null if it has no key.
	Key *Key `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The entity's properties.
	// The map's keys are property names.
	// A property name matching regex `__.*__` is reserved.
	// A reserved property name is forbidden in certain documented contexts.
	// The name must not contain more than 500 characters.
	// The name cannot be `""`.
	Properties           map[string]*Value `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbfdafa45100300, []int{4}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Entity) GetProperties() map[string]*Value {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*PartitionId)(nil), "google.datastore.v1.PartitionId")
	proto.RegisterType((*Key)(nil), "google.datastore.v1.Key")
	proto.RegisterType((*Key_PathElement)(nil), "google.datastore.v1.Key.PathElement")
	proto.RegisterType((*ArrayValue)(nil), "google.datastore.v1.ArrayValue")
	proto.RegisterType((*Value)(nil), "google.datastore.v1.Value")
	proto.RegisterType((*Entity)(nil), "google.datastore.v1.Entity")
	proto.RegisterMapType((map[string]*Value)(nil), "google.datastore.v1.Entity.PropertiesEntry")
}

func init() { proto.RegisterFile("google/datastore/v1/entity.proto", fileDescriptor_ecbfdafa45100300) }

var fileDescriptor_ecbfdafa45100300 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xff, 0x6e, 0xdc, 0x44,
	0x10, 0xc7, 0xed, 0xbb, 0x5c, 0x1a, 0x8f, 0xdd, 0xa4, 0x6c, 0x2a, 0x61, 0x02, 0x28, 0x26, 0x80,
	0x74, 0x02, 0xc9, 0x6e, 0xc2, 0x1f, 0x54, 0x14, 0xa4, 0x72, 0x25, 0xe0, 0x28, 0x15, 0x9c, 0x56,
	0x55, 0x24, 0x50, 0xa4, 0xd3, 0xde, 0x79, 0xeb, 0x2e, 0x67, 0xef, 0x5a, 0xf6, 0x3a, 0xaa, 0xdf,
	0x05, 0xf1, 0x00, 0x3c, 0x0a, 0x8f, 0x80, 0x78, 0x18, 0xb4, 0x3f, 0xec, 0x0b, 0xed, 0x35, 0xff,
	0x79, 0x67, 0x3e, 0xdf, 0xd9, 0xef, 0xec, 0xce, 0x1a, 0xa2, 0x5c, 0x88, 0xbc, 0xa0, 0x49, 0x46,
	0x24, 0x69, 0xa4, 0xa8, 0x69, 0x72, 0x73, 0x9a, 0x50, 0x2e, 0x99, 0xec, 0xe2, 0xaa, 0x16, 0x52,
	0xa0, 0x43, 0x43, 0xc4, 0x03, 0x11, 0xdf, 0x9c, 0x1e, 0x7d, 0x64, 0x65, 0xa4, 0x62, 0x09, 0xe1,
	0x5c, 0x48, 0x22, 0x99, 0xe0, 0x8d, 0x91, 0x0c, 0x59, 0xbd, 0x5a, 0xb6, 0x2f, 0x93, 0x46, 0xd6,
	0xed, 0x4a, 0xda, 0xec, 0xf1, 0x9b, 0x59, 0xc9, 0x4a, 0xda, 0x48, 0x52, 0x56, 0x16, 0x08, 0x2d,
	0x20, 0xbb, 0x8a, 0x26, 0x05, 0x91, 0x05, 0xcf, 0x4d, 0xe6, 0xe4, 0x17, 0xf0, 0xe7, 0xa4, 0x96,
	0x4c, 0x6d, 0x76, 0x91, 0xa1, 0x8f, 0x01, 0xaa, 0x5a, 0xfc, 0x4e, 0x57, 0x72, 0xc1, 0xb2, 0x70,
	0x14, 0xb9, 0x53, 0x0f, 0x7b, 0x36, 0x72, 0x91, 0xa1, 0x4f, 0x20, 0xe0, 0xa4, 0xa4, 0x4d, 0x45,
	0x56, 0x54, 0x01, 0x3b, 0x1a, 0xf0, 0x87, 0xd8, 0x45, 0x76, 0xf2, 0x8f, 0x0b, 0xe3, 0x4b, 0xda,
	0xa1, 0x67, 0x10, 0x54, 0x7d, 0x61, 0x85, 0xba, 0x91, 0x3b, 0xf5, 0xcf, 0xa2, 0x78, 0x4b, 0xef,
	0xf1, 0x2d, 0x07, 0xd8, 0xaf, 0x6e, 0xd9, 0x79, 0x0c, 0x3b, 0x15, 0x91, 0xaf, 0xc2, 0x51, 0x34,
	0x9e, 0xfa, 0x67, 0x9f, 0x6d, 0x15, 0x5f, 0xd2, 0x2e, 0x9e, 0x13, 0xf9, 0xea, 0xbc, 0xa0, 0x25,
	0xe5, 0x12, 0x6b, 0xc5, 0xd1, 0x0b, 0xd5, 0xd7, 0x10, 0x44, 0x08, 0x76, 0xd6, 0x8c, 0x1b, 0x17,
	0x1e, 0xd6, 0xdf, 0xe8, 0x01, 0x8c, 0x6c, 0x8f, 0xe3, 0xd4, 0xc1, 0x23, 0x96, 0xa1, 0x87, 0xb0,
	0xa3, 0x5a, 0x09, 0xc7, 0x8a, 0x4a, 0x1d, 0xac, 0x57, 0x33, 0x0f, 0xee, 0xb1, 0x6c, 0xa1, 0x8e,
	0xee, 0xe4, 0x29, 0xc0, 0xf7, 0x75, 0x4d, 0xba, 0x2b, 0x52, 0xb4, 0x14, 0x9d, 0xc1, 0xee, 0x8d,
	0xfa, 0x68, 0x42, 0x57, 0xfb, 0x3b, 0xda, 0xea, 0x4f, 0xb3, 0xd8, 0x92, 0x27, 0x7f, 0x4c, 0x60,
	0x62, 0xd4, 0x4f, 0x00, 0x78, 0x5b, 0x14, 0x0b, 0x9d, 0x08, 0xfd, 0xc8, 0x9d, 0xee, 0x6f, 0x2a,
	0xf4, 0x37, 0x19, 0xff, 0xdc, 0x16, 0x85, 0xe6, 0x53, 0x07, 0x7b, 0xbc, 0x5f, 0xa0, 0xcf, 0xe1,
	0xfe, 0x52, 0x88, 0x82, 0x12, 0x6e, 0xf5, 0xaa, 0xb1, 0xbd, 0xd4, 0xc1, 0x81, 0x0d, 0x0f, 0x18,
	0xe3, 0x92, 0xe6, 0xb4, 0xb6, 0x58, 0xdf, 0x6d, 0x60, 0xc3, 0x06, 0xfb, 0x14, 0x82, 0x4c, 0xb4,
	0xcb, 0x82, 0x5a, 0x4a, 0xf5, 0xef, 0xa6, 0x0e, 0xf6, 0x4d, 0xd4, 0x40, 0xe7, 0x70, 0x30, 0x8c,
	0x95, 0xe5, 0x40, 0xdf, 0xe9, 0xdb, 0xa6, 0x5f, 0xf4, 0x5c, 0xea, 0xe0, 0xfd, 0x41, 0x64, 0xca,
	0x7c, 0x0d, 0xde, 0x9a, 0x76, 0xb6, 0xc0, 0x44, 0x17, 0x08, 0xdf, 0x75, 0xaf, 0xa9, 0x83, 0xf7,
	0xd6, 0xb4, 0x1b, 0x4c, 0x36, 0xb2, 0x66, 0x3c, 0xb7, 0xda, 0xf7, 0xec, 0x25, 0xf9, 0x26, 0x6a,
	0xa0, 0x63, 0x80, 0x65, 0x21, 0x96, 0x16, 0x41, 0x91, 0x3b, 0x0d, 0xd4, 0xc1, 0xa9, 0x98, 0x01,
	0xbe, 0x83, 0x83, 0x9c, 0x8a, 0x45, 0x25, 0x18, 0x97, 0x96, 0xda, 0xd3, 0x26, 0x0e, 0x7b, 0x13,
	0xea, 0xa2, 0xe3, 0xe7, 0x44, 0x3e, 0xe7, 0x79, 0xea, 0xe0, 0xfb, 0x39, 0x15, 0x73, 0x05, 0x1b,
	0xf9, 0x53, 0x08, 0xcc, 0x53, 0xb6, 0xda, 0x5d, 0xad, 0xfd, 0x70, 0x6b, 0x03, 0xe7, 0x1a, 0x54,
	0x0e, 0x8d, 0xc4, 0x54, 0x98, 0x81, 0x4f, 0xd4, 0x08, 0xd9, 0x02, 0x9e, 0x2e, 0x70, 0xbc, 0xb5,
	0xc0, 0x66, 0xd4, 0x52, 0x07, 0x03, 0xd9, 0x0c, 0x5e, 0x08, 0xf7, 0x4a, 0x4a, 0x38, 0xe3, 0x79,
	0xb8, 0x1f, 0xb9, 0xd3, 0x09, 0xee, 0x97, 0xe8, 0x11, 0x3c, 0xa4, 0xaf, 0x57, 0x45, 0x9b, 0xd1,
	0xc5, 0xcb, 0x5a, 0x94, 0x0b, 0xc6, 0x33, 0xfa, 0x9a, 0x36, 0xe1, 0xa1, 0x1a, 0x0f, 0x8c, 0x6c,
	0xee, 0xc7, 0x5a, 0x94, 0x17, 0x26, 0x33, 0x0b, 0x00, 0xb4, 0x13, 0x33, 0xe0, 0xff, 0xba, 0xb0,
	0x6b, 0x7c, 0xa3, 0x2f, 0x60, 0xbc, 0xa6, 0x9d, 0x7d, 0xb7, 0xef, 0xbc, 0x22, 0xac, 0x20, 0x74,
	0xa9, 0x7f, 0x1b, 0x15, 0xad, 0x25, 0xa3, 0x4d, 0x38, 0xd6, 0xaf, 0xe1, 0xcb, 0x3b, 0x0e, 0x25,
	0x9e, 0x0f, 0xf4, 0x39, 0x97, 0x75, 0x87, 0x6f, 0xc9, 0x8f, 0x7e, 0x85, 0x83, 0x37, 0xd2, 0xe8,
	0xc1, 0xc6, 0x8b, 0x67, 0x76, 0x7c, 0x04, 0x93, 0xcd, 0x44, 0xdf, 0xfd, 0xf4, 0x0c, 0xf8, 0xcd,
	0xe8, 0xb1, 0x3b, 0xfb, 0xd3, 0x85, 0xf7, 0x57, 0xa2, 0xdc, 0x06, 0xcf, 0x7c, 0x63, 0x6d, 0xae,
	0x86, 0x78, 0xee, 0xfe, 0xf6, 0xad, 0x65, 0x72, 0x51, 0x10, 0x9e, 0xc7, 0xa2, 0xce, 0x93, 0x9c,
	0x72, 0x3d, 0xe2, 0x89, 0x49, 0x91, 0x8a, 0x35, 0xff, 0xfb, 0xcb, 0x3f, 0x19, 0x16, 0x7f, 0x8d,
	0x3e, 0xf8, 0xc9, 0xc8, 0x9f, 0x15, 0xa2, 0xcd, 0xe2, 0x1f, 0x86, 0x8d, 0xae, 0x4e, 0xff, 0xee,
	0x73, 0xd7, 0x3a, 0x77, 0x3d, 0xe4, 0xae, 0xaf, 0x4e, 0x97, 0xbb, 0x7a, 0x83, 0xaf, 0xfe, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0xf3, 0xdd, 0x11, 0x96, 0x45, 0x06, 0x00, 0x00,
}
