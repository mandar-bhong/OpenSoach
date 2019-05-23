// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/asset/v1beta1/asset_service.proto

package asset

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
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

// Asset content type.
type ContentType int32

const (
	// Unspecified content type.
	ContentType_CONTENT_TYPE_UNSPECIFIED ContentType = 0
	// Resource metadata.
	ContentType_RESOURCE ContentType = 1
	// The actual IAM policy set on a resource.
	ContentType_IAM_POLICY ContentType = 2
)

var ContentType_name = map[int32]string{
	0: "CONTENT_TYPE_UNSPECIFIED",
	1: "RESOURCE",
	2: "IAM_POLICY",
}

var ContentType_value = map[string]int32{
	"CONTENT_TYPE_UNSPECIFIED": 0,
	"RESOURCE":                 1,
	"IAM_POLICY":               2,
}

func (x ContentType) String() string {
	return proto.EnumName(ContentType_name, int32(x))
}

func (ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{0}
}

// Export asset request.
type ExportAssetsRequest struct {
	// Required. The relative name of the root asset. This can only be an
	// organization number (such as "organizations/123"), a project ID (such as
	// "projects/my-project-id"), a project number (such as "projects/12345"), or
	// a folder number (such as "folders/123").
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Timestamp to take an asset snapshot. This can only be set to a timestamp
	// between 2018-10-02 UTC (inclusive) and the current time. If not specified,
	// the current time will be used. Due to delays in resource data collection
	// and indexing, there is a volatile window during which running the same
	// query may get different results.
	ReadTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=read_time,json=readTime,proto3" json:"read_time,omitempty"`
	// A list of asset types of which to take a snapshot for. For example:
	// "google.compute.Disk". If specified, only matching assets will be returned.
	// See [Introduction to Cloud Asset
	// Inventory](https://cloud.google.com/resource-manager/docs/cloud-asset-inventory/overview)
	// for all supported asset types.
	AssetTypes []string `protobuf:"bytes,3,rep,name=asset_types,json=assetTypes,proto3" json:"asset_types,omitempty"`
	// Asset content type. If not specified, no content but the asset name will be
	// returned.
	ContentType ContentType `protobuf:"varint,4,opt,name=content_type,json=contentType,proto3,enum=google.cloud.asset.v1beta1.ContentType" json:"content_type,omitempty"`
	// Required. Output configuration indicating where the results will be output
	// to. All results will be in newline delimited JSON format.
	OutputConfig         *OutputConfig `protobuf:"bytes,5,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExportAssetsRequest) Reset()         { *m = ExportAssetsRequest{} }
func (m *ExportAssetsRequest) String() string { return proto.CompactTextString(m) }
func (*ExportAssetsRequest) ProtoMessage()    {}
func (*ExportAssetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{0}
}

func (m *ExportAssetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportAssetsRequest.Unmarshal(m, b)
}
func (m *ExportAssetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportAssetsRequest.Marshal(b, m, deterministic)
}
func (m *ExportAssetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportAssetsRequest.Merge(m, src)
}
func (m *ExportAssetsRequest) XXX_Size() int {
	return xxx_messageInfo_ExportAssetsRequest.Size(m)
}
func (m *ExportAssetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportAssetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportAssetsRequest proto.InternalMessageInfo

func (m *ExportAssetsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ExportAssetsRequest) GetReadTime() *timestamp.Timestamp {
	if m != nil {
		return m.ReadTime
	}
	return nil
}

func (m *ExportAssetsRequest) GetAssetTypes() []string {
	if m != nil {
		return m.AssetTypes
	}
	return nil
}

func (m *ExportAssetsRequest) GetContentType() ContentType {
	if m != nil {
		return m.ContentType
	}
	return ContentType_CONTENT_TYPE_UNSPECIFIED
}

func (m *ExportAssetsRequest) GetOutputConfig() *OutputConfig {
	if m != nil {
		return m.OutputConfig
	}
	return nil
}

// The export asset response. This message is returned by the
// [google.longrunning.Operations.GetOperation][google.longrunning.Operations.GetOperation]
// method in the returned
// [google.longrunning.Operation.response][google.longrunning.Operation.response]
// field.
type ExportAssetsResponse struct {
	// Time the snapshot was taken.
	ReadTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=read_time,json=readTime,proto3" json:"read_time,omitempty"`
	// Output configuration indicating where the results were output to.
	// All results are in JSON format.
	OutputConfig         *OutputConfig `protobuf:"bytes,2,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExportAssetsResponse) Reset()         { *m = ExportAssetsResponse{} }
func (m *ExportAssetsResponse) String() string { return proto.CompactTextString(m) }
func (*ExportAssetsResponse) ProtoMessage()    {}
func (*ExportAssetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{1}
}

func (m *ExportAssetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportAssetsResponse.Unmarshal(m, b)
}
func (m *ExportAssetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportAssetsResponse.Marshal(b, m, deterministic)
}
func (m *ExportAssetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportAssetsResponse.Merge(m, src)
}
func (m *ExportAssetsResponse) XXX_Size() int {
	return xxx_messageInfo_ExportAssetsResponse.Size(m)
}
func (m *ExportAssetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportAssetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportAssetsResponse proto.InternalMessageInfo

func (m *ExportAssetsResponse) GetReadTime() *timestamp.Timestamp {
	if m != nil {
		return m.ReadTime
	}
	return nil
}

func (m *ExportAssetsResponse) GetOutputConfig() *OutputConfig {
	if m != nil {
		return m.OutputConfig
	}
	return nil
}

// Batch get assets history request.
type BatchGetAssetsHistoryRequest struct {
	// Required. The relative name of the root asset. It can only be an
	// organization number (such as "organizations/123"), a project ID (such as
	// "projects/my-project-id")", or a project number (such as "projects/12345").
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// A list of the full names of the assets. For example:
	// `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`.
	// See [Resource
	// Names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	// for more info.
	//
	// The request becomes a no-op if the asset name list is empty, and the max
	// size of the asset name list is 100 in one request.
	AssetNames []string `protobuf:"bytes,2,rep,name=asset_names,json=assetNames,proto3" json:"asset_names,omitempty"`
	// Required. The content type.
	ContentType ContentType `protobuf:"varint,3,opt,name=content_type,json=contentType,proto3,enum=google.cloud.asset.v1beta1.ContentType" json:"content_type,omitempty"`
	// Optional. The time window for the asset history. Both start_time and
	// end_time are optional and if set, it must be after 2018-10-02 UTC. If
	// end_time is not set, it is default to current timestamp. If start_time is
	// not set, the snapshot of the assets at end_time will be returned. The
	// returned results contain all temporal assets whose time window overlap with
	// read_time_window.
	ReadTimeWindow       *TimeWindow `protobuf:"bytes,4,opt,name=read_time_window,json=readTimeWindow,proto3" json:"read_time_window,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BatchGetAssetsHistoryRequest) Reset()         { *m = BatchGetAssetsHistoryRequest{} }
func (m *BatchGetAssetsHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*BatchGetAssetsHistoryRequest) ProtoMessage()    {}
func (*BatchGetAssetsHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{2}
}

func (m *BatchGetAssetsHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchGetAssetsHistoryRequest.Unmarshal(m, b)
}
func (m *BatchGetAssetsHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchGetAssetsHistoryRequest.Marshal(b, m, deterministic)
}
func (m *BatchGetAssetsHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchGetAssetsHistoryRequest.Merge(m, src)
}
func (m *BatchGetAssetsHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_BatchGetAssetsHistoryRequest.Size(m)
}
func (m *BatchGetAssetsHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchGetAssetsHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchGetAssetsHistoryRequest proto.InternalMessageInfo

func (m *BatchGetAssetsHistoryRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *BatchGetAssetsHistoryRequest) GetAssetNames() []string {
	if m != nil {
		return m.AssetNames
	}
	return nil
}

func (m *BatchGetAssetsHistoryRequest) GetContentType() ContentType {
	if m != nil {
		return m.ContentType
	}
	return ContentType_CONTENT_TYPE_UNSPECIFIED
}

func (m *BatchGetAssetsHistoryRequest) GetReadTimeWindow() *TimeWindow {
	if m != nil {
		return m.ReadTimeWindow
	}
	return nil
}

// Batch get assets history response.
type BatchGetAssetsHistoryResponse struct {
	// A list of assets with valid time windows.
	Assets               []*TemporalAsset `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BatchGetAssetsHistoryResponse) Reset()         { *m = BatchGetAssetsHistoryResponse{} }
func (m *BatchGetAssetsHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*BatchGetAssetsHistoryResponse) ProtoMessage()    {}
func (*BatchGetAssetsHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{3}
}

func (m *BatchGetAssetsHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchGetAssetsHistoryResponse.Unmarshal(m, b)
}
func (m *BatchGetAssetsHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchGetAssetsHistoryResponse.Marshal(b, m, deterministic)
}
func (m *BatchGetAssetsHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchGetAssetsHistoryResponse.Merge(m, src)
}
func (m *BatchGetAssetsHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_BatchGetAssetsHistoryResponse.Size(m)
}
func (m *BatchGetAssetsHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchGetAssetsHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchGetAssetsHistoryResponse proto.InternalMessageInfo

func (m *BatchGetAssetsHistoryResponse) GetAssets() []*TemporalAsset {
	if m != nil {
		return m.Assets
	}
	return nil
}

// Output configuration for export assets destination.
type OutputConfig struct {
	// Asset export destination.
	//
	// Types that are valid to be assigned to Destination:
	//	*OutputConfig_GcsDestination
	Destination          isOutputConfig_Destination `protobuf_oneof:"destination"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *OutputConfig) Reset()         { *m = OutputConfig{} }
func (m *OutputConfig) String() string { return proto.CompactTextString(m) }
func (*OutputConfig) ProtoMessage()    {}
func (*OutputConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{4}
}

func (m *OutputConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputConfig.Unmarshal(m, b)
}
func (m *OutputConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputConfig.Marshal(b, m, deterministic)
}
func (m *OutputConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputConfig.Merge(m, src)
}
func (m *OutputConfig) XXX_Size() int {
	return xxx_messageInfo_OutputConfig.Size(m)
}
func (m *OutputConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OutputConfig proto.InternalMessageInfo

type isOutputConfig_Destination interface {
	isOutputConfig_Destination()
}

type OutputConfig_GcsDestination struct {
	GcsDestination *GcsDestination `protobuf:"bytes,1,opt,name=gcs_destination,json=gcsDestination,proto3,oneof"`
}

func (*OutputConfig_GcsDestination) isOutputConfig_Destination() {}

func (m *OutputConfig) GetDestination() isOutputConfig_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *OutputConfig) GetGcsDestination() *GcsDestination {
	if x, ok := m.GetDestination().(*OutputConfig_GcsDestination); ok {
		return x.GcsDestination
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OutputConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OutputConfig_GcsDestination)(nil),
	}
}

// A Cloud Storage location.
type GcsDestination struct {
	// Required.
	//
	// Types that are valid to be assigned to ObjectUri:
	//	*GcsDestination_Uri
	ObjectUri            isGcsDestination_ObjectUri `protobuf_oneof:"object_uri"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GcsDestination) Reset()         { *m = GcsDestination{} }
func (m *GcsDestination) String() string { return proto.CompactTextString(m) }
func (*GcsDestination) ProtoMessage()    {}
func (*GcsDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_565cfbddaa85b7d6, []int{5}
}

func (m *GcsDestination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GcsDestination.Unmarshal(m, b)
}
func (m *GcsDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GcsDestination.Marshal(b, m, deterministic)
}
func (m *GcsDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GcsDestination.Merge(m, src)
}
func (m *GcsDestination) XXX_Size() int {
	return xxx_messageInfo_GcsDestination.Size(m)
}
func (m *GcsDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_GcsDestination.DiscardUnknown(m)
}

var xxx_messageInfo_GcsDestination proto.InternalMessageInfo

type isGcsDestination_ObjectUri interface {
	isGcsDestination_ObjectUri()
}

type GcsDestination_Uri struct {
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3,oneof"`
}

func (*GcsDestination_Uri) isGcsDestination_ObjectUri() {}

func (m *GcsDestination) GetObjectUri() isGcsDestination_ObjectUri {
	if m != nil {
		return m.ObjectUri
	}
	return nil
}

func (m *GcsDestination) GetUri() string {
	if x, ok := m.GetObjectUri().(*GcsDestination_Uri); ok {
		return x.Uri
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GcsDestination) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GcsDestination_Uri)(nil),
	}
}

func init() {
	proto.RegisterEnum("google.cloud.asset.v1beta1.ContentType", ContentType_name, ContentType_value)
	proto.RegisterType((*ExportAssetsRequest)(nil), "google.cloud.asset.v1beta1.ExportAssetsRequest")
	proto.RegisterType((*ExportAssetsResponse)(nil), "google.cloud.asset.v1beta1.ExportAssetsResponse")
	proto.RegisterType((*BatchGetAssetsHistoryRequest)(nil), "google.cloud.asset.v1beta1.BatchGetAssetsHistoryRequest")
	proto.RegisterType((*BatchGetAssetsHistoryResponse)(nil), "google.cloud.asset.v1beta1.BatchGetAssetsHistoryResponse")
	proto.RegisterType((*OutputConfig)(nil), "google.cloud.asset.v1beta1.OutputConfig")
	proto.RegisterType((*GcsDestination)(nil), "google.cloud.asset.v1beta1.GcsDestination")
}

func init() {
	proto.RegisterFile("google/cloud/asset/v1beta1/asset_service.proto", fileDescriptor_565cfbddaa85b7d6)
}

var fileDescriptor_565cfbddaa85b7d6 = []byte{
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xae, 0x9d, 0xde, 0xaa, 0x9d, 0xa4, 0xb9, 0xb9, 0x73, 0x7f, 0x64, 0x45, 0xed, 0x6d, 0x64,
	0x24, 0x9a, 0x46, 0xc2, 0x56, 0xd3, 0x45, 0x69, 0x11, 0x42, 0x4d, 0x1a, 0xda, 0x20, 0x9a, 0x44,
	0x6e, 0x5a, 0xd4, 0xaa, 0x92, 0xe5, 0x38, 0x53, 0x63, 0x94, 0xcc, 0x18, 0xcf, 0xb8, 0xa5, 0x20,
	0x36, 0xf0, 0x04, 0xc0, 0x8a, 0x0d, 0x0b, 0x96, 0xac, 0x78, 0x0e, 0xb6, 0xbc, 0x02, 0xaf, 0xc0,
	0x86, 0x15, 0xf2, 0x8c, 0xd3, 0x3a, 0x4a, 0xea, 0x0a, 0xba, 0x3c, 0x33, 0xdf, 0xf7, 0x9d, 0x33,
	0xdf, 0x99, 0x39, 0x03, 0x34, 0x87, 0x10, 0xa7, 0x87, 0x74, 0xbb, 0x47, 0x82, 0xae, 0x6e, 0x51,
	0x8a, 0x98, 0x7e, 0xb2, 0xdc, 0x41, 0xcc, 0x5a, 0x16, 0x91, 0x49, 0x91, 0x7f, 0xe2, 0xda, 0x48,
	0xf3, 0x7c, 0xc2, 0x08, 0xcc, 0x0b, 0xbc, 0xc6, 0xf1, 0x1a, 0x47, 0x68, 0x11, 0x3e, 0x3f, 0x17,
	0x69, 0x59, 0x9e, 0xab, 0x5b, 0x18, 0x13, 0x66, 0x31, 0x97, 0x60, 0x2a, 0x98, 0xf9, 0xc5, 0xab,
	0x32, 0x0d, 0x80, 0x37, 0x22, 0x60, 0x8f, 0x60, 0xc7, 0x0f, 0x30, 0x76, 0xb1, 0xa3, 0x13, 0x0f,
	0xf9, 0x43, 0x6a, 0x0b, 0x11, 0x88, 0x47, 0x9d, 0xe0, 0x58, 0x67, 0x6e, 0x1f, 0x51, 0x66, 0xf5,
	0x3d, 0x01, 0x50, 0x3f, 0xca, 0xe0, 0xef, 0xda, 0x33, 0x8f, 0xf8, 0x6c, 0x83, 0x8b, 0x1b, 0xe8,
	0x69, 0x80, 0x28, 0x83, 0xff, 0x81, 0x29, 0xcf, 0xf2, 0x11, 0x66, 0x8a, 0x54, 0x90, 0x8a, 0x33,
	0x46, 0x14, 0xc1, 0x55, 0x30, 0xe3, 0x23, 0xab, 0x6b, 0x86, 0x3a, 0x8a, 0x5c, 0x90, 0x8a, 0xe9,
	0x72, 0x3e, 0x32, 0x47, 0x1b, 0x24, 0xd1, 0xda, 0x83, 0x24, 0xc6, 0x74, 0x08, 0x0e, 0x43, 0xb8,
	0x00, 0xd2, 0xc2, 0x28, 0x76, 0xe6, 0x21, 0xaa, 0xa4, 0x0a, 0xa9, 0xe2, 0x8c, 0x01, 0xf8, 0x52,
	0x3b, 0x5c, 0x81, 0x0f, 0x40, 0xc6, 0x26, 0x98, 0x21, 0x2c, 0x20, 0xca, 0x64, 0x41, 0x2a, 0x66,
	0xcb, 0x8b, 0xda, 0xe5, 0x4e, 0x6a, 0x55, 0x81, 0x0f, 0xf9, 0x46, 0xda, 0xbe, 0x08, 0xe0, 0x0e,
	0x98, 0x25, 0x01, 0xf3, 0x02, 0x66, 0xda, 0x04, 0x1f, 0xbb, 0x8e, 0xf2, 0x07, 0xaf, 0xb4, 0x98,
	0x24, 0xd6, 0xe4, 0x84, 0x2a, 0xc7, 0x1b, 0x19, 0x12, 0x8b, 0xd4, 0x0f, 0x12, 0xf8, 0x67, 0xd8,
	0x24, 0xea, 0x11, 0x4c, 0xd1, 0xb0, 0x1b, 0xd2, 0x2f, 0xb8, 0x31, 0x52, 0xa0, 0x7c, 0xad, 0x02,
	0xbf, 0x4b, 0x60, 0xae, 0x62, 0x31, 0xfb, 0xf1, 0x16, 0x8a, 0x4a, 0xdc, 0x76, 0x29, 0x23, 0xfe,
	0xd9, 0x55, 0xed, 0x3c, 0xef, 0x0a, 0xb6, 0xfa, 0x88, 0x2a, 0x72, 0xac, 0x2b, 0x8d, 0x70, 0x65,
	0xa4, 0x2b, 0xa9, 0x6b, 0x74, 0xa5, 0x05, 0x72, 0xe7, 0x6e, 0x99, 0xa7, 0x2e, 0xee, 0x92, 0x53,
	0xde, 0xe5, 0x74, 0xf9, 0x66, 0x92, 0x5e, 0x68, 0xd8, 0x23, 0x8e, 0x36, 0xb2, 0x03, 0x03, 0x45,
	0xac, 0x76, 0xc0, 0xfc, 0x25, 0xc7, 0x8e, 0x1a, 0xb4, 0x01, 0xa6, 0xc4, 0xa3, 0x51, 0xa4, 0x42,
	0xaa, 0x98, 0x2e, 0x2f, 0x25, 0x26, 0x42, 0x7d, 0x8f, 0xf8, 0x56, 0x8f, 0x4b, 0x19, 0x11, 0x51,
	0x65, 0x20, 0x13, 0x77, 0x1e, 0xee, 0x81, 0x3f, 0x1d, 0x9b, 0x9a, 0x5d, 0x44, 0x99, 0x8b, 0xf9,
	0x63, 0x8b, 0x3a, 0x5f, 0x4a, 0xd2, 0xde, 0xb2, 0xe9, 0xe6, 0x05, 0x63, 0x7b, 0xc2, 0xc8, 0x3a,
	0x43, 0x2b, 0x95, 0x59, 0x90, 0x8e, 0x49, 0xaa, 0x65, 0x90, 0x1d, 0xa6, 0x40, 0x08, 0x52, 0x81,
	0xef, 0x8a, 0xfe, 0x6d, 0x4f, 0x18, 0x61, 0x50, 0xc9, 0x00, 0x40, 0x3a, 0x4f, 0x90, 0xcd, 0xcc,
	0xc0, 0x77, 0x4b, 0x75, 0x90, 0x8e, 0x79, 0x0f, 0xe7, 0x80, 0x52, 0x6d, 0x36, 0xda, 0xb5, 0x46,
	0xdb, 0x6c, 0x1f, 0xb4, 0x6a, 0xe6, 0x5e, 0x63, 0xb7, 0x55, 0xab, 0xd6, 0xef, 0xd7, 0x6b, 0x9b,
	0xb9, 0x09, 0x98, 0x01, 0xd3, 0x46, 0x6d, 0xb7, 0xb9, 0x67, 0x54, 0x6b, 0x39, 0x09, 0x66, 0x01,
	0xa8, 0x6f, 0xec, 0x98, 0xad, 0xe6, 0xc3, 0x7a, 0xf5, 0x20, 0x27, 0x97, 0xdf, 0x4c, 0x82, 0x0c,
	0xb7, 0x61, 0x57, 0x8c, 0x35, 0xf8, 0x43, 0x02, 0x99, 0xf8, 0x13, 0x80, 0x7a, 0xd2, 0x69, 0xc7,
	0x4c, 0x94, 0xfc, 0xfc, 0x80, 0x10, 0x1b, 0x58, 0x5a, 0x73, 0x30, 0xb0, 0xd4, 0xf7, 0xd2, 0xab,
	0xaf, 0xdf, 0xde, 0xc9, 0x6f, 0x25, 0x75, 0xe9, 0x7c, 0xde, 0xbd, 0x10, 0xb7, 0xf4, 0xae, 0xe7,
	0x93, 0xf0, 0x90, 0x54, 0x2f, 0xbd, 0x5c, 0x47, 0x31, 0xe9, 0x75, 0xa9, 0x74, 0x78, 0x4b, 0x2d,
	0x8e, 0xe0, 0x8f, 0x49, 0xaf, 0x8b, 0xfc, 0xb1, 0xf0, 0x15, 0x55, 0x1b, 0x81, 0x13, 0xdf, 0xb1,
	0xb0, 0xfb, 0x5c, 0x0c, 0xce, 0x31, 0x24, 0xf8, 0x5a, 0x06, 0xff, 0x8e, 0xbd, 0x67, 0xf0, 0x76,
	0x92, 0x0b, 0x49, 0x2f, 0x32, 0xbf, 0xf6, 0x1b, 0x4c, 0x71, 0xa9, 0xd5, 0x80, 0x3b, 0x45, 0x60,
	0x39, 0xd1, 0xa8, 0xce, 0x38, 0x8d, 0xc3, 0x35, 0xb8, 0x7a, 0xf5, 0xf9, 0xc7, 0x52, 0x2b, 0x9f,
	0x25, 0xf0, 0xbf, 0x4d, 0xfa, 0x09, 0x75, 0x57, 0xfe, 0x8a, 0xdf, 0x99, 0x56, 0x38, 0x00, 0x5b,
	0xd2, 0xe1, 0xbd, 0x88, 0xe0, 0x90, 0x9e, 0x85, 0x1d, 0x8d, 0xf8, 0x8e, 0xee, 0x20, 0xcc, 0xc7,
	0xa3, 0x2e, 0xb6, 0x2c, 0xcf, 0xa5, 0xe3, 0x3e, 0xbc, 0x3b, 0x3c, 0xfa, 0x24, 0xe7, 0xb7, 0x84,
	0x42, 0x95, 0xa7, 0xe4, 0x39, 0xb4, 0xfd, 0xe5, 0x4a, 0x08, 0xf9, 0x32, 0xd8, 0x3c, 0xe2, 0x9b,
	0x47, 0x7c, 0xf3, 0x68, 0x5f, 0xf0, 0x3b, 0x53, 0x3c, 0xcb, 0xca, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc4, 0xf2, 0x88, 0x01, 0xbf, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssetServiceClient interface {
	// Exports assets with time and resource types to a given Cloud Storage
	// location. The output format is newline-delimited JSON.
	// This API implements the
	// [google.longrunning.Operation][google.longrunning.Operation] API allowing
	// you to keep track of the export.
	ExportAssets(ctx context.Context, in *ExportAssetsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Batch gets the update history of assets that overlap a time window.
	// For RESOURCE content, this API outputs history with asset in both
	// non-delete or deleted status.
	// For IAM_POLICY content, this API outputs history when the asset and its
	// attached IAM POLICY both exist. This can create gaps in the output history.
	BatchGetAssetsHistory(ctx context.Context, in *BatchGetAssetsHistoryRequest, opts ...grpc.CallOption) (*BatchGetAssetsHistoryResponse, error)
}

type assetServiceClient struct {
	cc *grpc.ClientConn
}

func NewAssetServiceClient(cc *grpc.ClientConn) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) ExportAssets(ctx context.Context, in *ExportAssetsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.asset.v1beta1.AssetService/ExportAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) BatchGetAssetsHistory(ctx context.Context, in *BatchGetAssetsHistoryRequest, opts ...grpc.CallOption) (*BatchGetAssetsHistoryResponse, error) {
	out := new(BatchGetAssetsHistoryResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.asset.v1beta1.AssetService/BatchGetAssetsHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
type AssetServiceServer interface {
	// Exports assets with time and resource types to a given Cloud Storage
	// location. The output format is newline-delimited JSON.
	// This API implements the
	// [google.longrunning.Operation][google.longrunning.Operation] API allowing
	// you to keep track of the export.
	ExportAssets(context.Context, *ExportAssetsRequest) (*longrunning.Operation, error)
	// Batch gets the update history of assets that overlap a time window.
	// For RESOURCE content, this API outputs history with asset in both
	// non-delete or deleted status.
	// For IAM_POLICY content, this API outputs history when the asset and its
	// attached IAM POLICY both exist. This can create gaps in the output history.
	BatchGetAssetsHistory(context.Context, *BatchGetAssetsHistoryRequest) (*BatchGetAssetsHistoryResponse, error)
}

func RegisterAssetServiceServer(s *grpc.Server, srv AssetServiceServer) {
	s.RegisterService(&_AssetService_serviceDesc, srv)
}

func _AssetService_ExportAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).ExportAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.asset.v1beta1.AssetService/ExportAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).ExportAssets(ctx, req.(*ExportAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_BatchGetAssetsHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetAssetsHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).BatchGetAssetsHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.asset.v1beta1.AssetService/BatchGetAssetsHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).BatchGetAssetsHistory(ctx, req.(*BatchGetAssetsHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AssetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.asset.v1beta1.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExportAssets",
			Handler:    _AssetService_ExportAssets_Handler,
		},
		{
			MethodName: "BatchGetAssetsHistory",
			Handler:    _AssetService_BatchGetAssetsHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/asset/v1beta1/asset_service.proto",
}
