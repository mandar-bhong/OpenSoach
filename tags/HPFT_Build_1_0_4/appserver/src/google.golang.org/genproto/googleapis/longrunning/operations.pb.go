// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/longrunning/operations.proto

package longrunning

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// The server-assigned name, which is only unique within the same service that
	// originally returns it. If you use the default HTTP mapping, the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Service-specific metadata associated with the operation.  It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.  Any method that returns a
	// long-running operation should document the metadata type, if any.
	Metadata *any.Any `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// If the value is `false`, it means the operation is still in progress.
	// If `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	// The operation result, which can be either an `error` or a valid `response`.
	// If `done` == `false`, neither `error` nor `response` is set.
	// If `done` == `true`, exactly one of `error` or `response` is set.
	//
	// Types that are valid to be assigned to Result:
	//	*Operation_Error
	//	*Operation_Response
	Result               isOperation_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_bff5d3ff9032d7eb, []int{0}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Operation) GetMetadata() *any.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Operation) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type isOperation_Result interface {
	isOperation_Result()
}

type Operation_Error struct {
	Error *status.Status `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

type Operation_Response struct {
	Response *any.Any `protobuf:"bytes,5,opt,name=response,proto3,oneof"`
}

func (*Operation_Error) isOperation_Result() {}

func (*Operation_Response) isOperation_Result() {}

func (m *Operation) GetResult() isOperation_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Operation) GetError() *status.Status {
	if x, ok := m.GetResult().(*Operation_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Operation) GetResponse() *any.Any {
	if x, ok := m.GetResult().(*Operation_Response); ok {
		return x.Response
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Operation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Operation_Error)(nil),
		(*Operation_Response)(nil),
	}
}

// The request message for [Operations.GetOperation][google.longrunning.Operations.GetOperation].
type GetOperationRequest struct {
	// The name of the operation resource.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperationRequest) Reset()         { *m = GetOperationRequest{} }
func (m *GetOperationRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperationRequest) ProtoMessage()    {}
func (*GetOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bff5d3ff9032d7eb, []int{1}
}

func (m *GetOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationRequest.Unmarshal(m, b)
}
func (m *GetOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationRequest.Marshal(b, m, deterministic)
}
func (m *GetOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationRequest.Merge(m, src)
}
func (m *GetOperationRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperationRequest.Size(m)
}
func (m *GetOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationRequest proto.InternalMessageInfo

func (m *GetOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Operations.ListOperations][google.longrunning.Operations.ListOperations].
type ListOperationsRequest struct {
	// The name of the operation's parent resource.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The standard list filter.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The standard list page size.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The standard list page token.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOperationsRequest) Reset()         { *m = ListOperationsRequest{} }
func (m *ListOperationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListOperationsRequest) ProtoMessage()    {}
func (*ListOperationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bff5d3ff9032d7eb, []int{2}
}

func (m *ListOperationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOperationsRequest.Unmarshal(m, b)
}
func (m *ListOperationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOperationsRequest.Marshal(b, m, deterministic)
}
func (m *ListOperationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOperationsRequest.Merge(m, src)
}
func (m *ListOperationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListOperationsRequest.Size(m)
}
func (m *ListOperationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOperationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListOperationsRequest proto.InternalMessageInfo

func (m *ListOperationsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListOperationsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListOperationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListOperationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for [Operations.ListOperations][google.longrunning.Operations.ListOperations].
type ListOperationsResponse struct {
	// A list of operations that matches the specified filter in the request.
	Operations []*Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	// The standard List next-page token.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOperationsResponse) Reset()         { *m = ListOperationsResponse{} }
func (m *ListOperationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListOperationsResponse) ProtoMessage()    {}
func (*ListOperationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bff5d3ff9032d7eb, []int{3}
}

func (m *ListOperationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOperationsResponse.Unmarshal(m, b)
}
func (m *ListOperationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOperationsResponse.Marshal(b, m, deterministic)
}
func (m *ListOperationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOperationsResponse.Merge(m, src)
}
func (m *ListOperationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListOperationsResponse.Size(m)
}
func (m *ListOperationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOperationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListOperationsResponse proto.InternalMessageInfo

func (m *ListOperationsResponse) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *ListOperationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for [Operations.CancelOperation][google.longrunning.Operations.CancelOperation].
type CancelOperationRequest struct {
	// The name of the operation resource to be cancelled.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelOperationRequest) Reset()         { *m = CancelOperationRequest{} }
func (m *CancelOperationRequest) String() string { return proto.CompactTextString(m) }
func (*CancelOperationRequest) ProtoMessage()    {}
func (*CancelOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bff5d3ff9032d7eb, []int{4}
}

func (m *CancelOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelOperationRequest.Unmarshal(m, b)
}
func (m *CancelOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelOperationRequest.Marshal(b, m, deterministic)
}
func (m *CancelOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelOperationRequest.Merge(m, src)
}
func (m *CancelOperationRequest) XXX_Size() int {
	return xxx_messageInfo_CancelOperationRequest.Size(m)
}
func (m *CancelOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelOperationRequest proto.InternalMessageInfo

func (m *CancelOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Operations.DeleteOperation][google.longrunning.Operations.DeleteOperation].
type DeleteOperationRequest struct {
	// The name of the operation resource to be deleted.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOperationRequest) Reset()         { *m = DeleteOperationRequest{} }
func (m *DeleteOperationRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteOperationRequest) ProtoMessage()    {}
func (*DeleteOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bff5d3ff9032d7eb, []int{5}
}

func (m *DeleteOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOperationRequest.Unmarshal(m, b)
}
func (m *DeleteOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOperationRequest.Marshal(b, m, deterministic)
}
func (m *DeleteOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOperationRequest.Merge(m, src)
}
func (m *DeleteOperationRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteOperationRequest.Size(m)
}
func (m *DeleteOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOperationRequest proto.InternalMessageInfo

func (m *DeleteOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Operations.WaitOperation][google.longrunning.Operations.WaitOperation].
type WaitOperationRequest struct {
	// The name of the operation resource to wait on.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The maximum duration to wait before timing out. If left blank, the wait
	// will be at most the time permitted by the underlying HTTP/RPC protocol.
	// If RPC context deadline is also specified, the shorter one will be used.
	Timeout              *duration.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *WaitOperationRequest) Reset()         { *m = WaitOperationRequest{} }
func (m *WaitOperationRequest) String() string { return proto.CompactTextString(m) }
func (*WaitOperationRequest) ProtoMessage()    {}
func (*WaitOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bff5d3ff9032d7eb, []int{6}
}

func (m *WaitOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitOperationRequest.Unmarshal(m, b)
}
func (m *WaitOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitOperationRequest.Marshal(b, m, deterministic)
}
func (m *WaitOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitOperationRequest.Merge(m, src)
}
func (m *WaitOperationRequest) XXX_Size() int {
	return xxx_messageInfo_WaitOperationRequest.Size(m)
}
func (m *WaitOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WaitOperationRequest proto.InternalMessageInfo

func (m *WaitOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WaitOperationRequest) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// A message representing the message types used by a long-running operation.
//
// Example:
//
//   rpc LongRunningRecognize(LongRunningRecognizeRequest)
//       returns (google.longrunning.Operation) {
//     option (google.longrunning.operation_info) = {
//       response_type: "LongRunningRecognizeResponse"
//       metadata_type: "LongRunningRecognizeMetadata"
//     };
//   }
type OperationInfo struct {
	// Required. The message name of the primary return type for this
	// long-running operation.
	// This type will be used to deserialize the LRO's response.
	//
	// If the response is in a different package from the rpc, a fully-qualified
	// message name must be used (e.g. `google.protobuf.Struct`).
	//
	// Note: Altering this value constitutes a breaking change.
	ResponseType string `protobuf:"bytes,1,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`
	// Required. The message name of the metadata type for this long-running
	// operation.
	//
	// If the response is in a different package from the rpc, a fully-qualified
	// message name must be used (e.g. `google.protobuf.Struct`).
	//
	// Note: Altering this value constitutes a breaking change.
	MetadataType         string   `protobuf:"bytes,2,opt,name=metadata_type,json=metadataType,proto3" json:"metadata_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationInfo) Reset()         { *m = OperationInfo{} }
func (m *OperationInfo) String() string { return proto.CompactTextString(m) }
func (*OperationInfo) ProtoMessage()    {}
func (*OperationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bff5d3ff9032d7eb, []int{7}
}

func (m *OperationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationInfo.Unmarshal(m, b)
}
func (m *OperationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationInfo.Marshal(b, m, deterministic)
}
func (m *OperationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationInfo.Merge(m, src)
}
func (m *OperationInfo) XXX_Size() int {
	return xxx_messageInfo_OperationInfo.Size(m)
}
func (m *OperationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OperationInfo proto.InternalMessageInfo

func (m *OperationInfo) GetResponseType() string {
	if m != nil {
		return m.ResponseType
	}
	return ""
}

func (m *OperationInfo) GetMetadataType() string {
	if m != nil {
		return m.MetadataType
	}
	return ""
}

var E_OperationInfo = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*OperationInfo)(nil),
	Field:         1049,
	Name:          "google.longrunning.operation_info",
	Tag:           "bytes,1049,opt,name=operation_info",
	Filename:      "google/longrunning/operations.proto",
}

func init() {
	proto.RegisterType((*Operation)(nil), "google.longrunning.Operation")
	proto.RegisterType((*GetOperationRequest)(nil), "google.longrunning.GetOperationRequest")
	proto.RegisterType((*ListOperationsRequest)(nil), "google.longrunning.ListOperationsRequest")
	proto.RegisterType((*ListOperationsResponse)(nil), "google.longrunning.ListOperationsResponse")
	proto.RegisterType((*CancelOperationRequest)(nil), "google.longrunning.CancelOperationRequest")
	proto.RegisterType((*DeleteOperationRequest)(nil), "google.longrunning.DeleteOperationRequest")
	proto.RegisterType((*WaitOperationRequest)(nil), "google.longrunning.WaitOperationRequest")
	proto.RegisterType((*OperationInfo)(nil), "google.longrunning.OperationInfo")
	proto.RegisterExtension(E_OperationInfo)
}

func init() {
	proto.RegisterFile("google/longrunning/operations.proto", fileDescriptor_bff5d3ff9032d7eb)
}

var fileDescriptor_bff5d3ff9032d7eb = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x4e, 0x13, 0x51,
	0x14, 0x66, 0x4a, 0xc1, 0xf6, 0x40, 0x69, 0x72, 0x85, 0x52, 0x8a, 0x60, 0x1d, 0x8c, 0x96, 0x86,
	0xcc, 0x68, 0xd9, 0x61, 0x58, 0x88, 0x18, 0x30, 0xc1, 0x40, 0x06, 0x12, 0x23, 0x31, 0x69, 0x86,
	0xf6, 0x30, 0x4e, 0x6c, 0xef, 0x1d, 0xef, 0xdc, 0x51, 0x8a, 0x21, 0x44, 0x17, 0xbe, 0x80, 0x0b,
	0xe3, 0x2b, 0xf8, 0x28, 0x6e, 0x5c, 0xf8, 0x0a, 0x2e, 0x7c, 0x04, 0x97, 0x66, 0xee, 0xcc, 0xb4,
	0x43, 0x3b, 0xc5, 0xee, 0xe6, 0x9e, 0xf3, 0xdd, 0xef, 0x3b, 0x3f, 0xdf, 0x6d, 0x61, 0xc5, 0x62,
	0xcc, 0x6a, 0xa1, 0xde, 0x62, 0xd4, 0xe2, 0x1e, 0xa5, 0x36, 0xb5, 0x74, 0xe6, 0x20, 0x37, 0x85,
	0xcd, 0xa8, 0xab, 0x39, 0x9c, 0x09, 0x46, 0x48, 0x00, 0xd2, 0x62, 0xa0, 0xd2, 0xad, 0xf0, 0xa2,
	0xe9, 0xd8, 0xba, 0x49, 0x29, 0x13, 0xf1, 0x1b, 0xa5, 0x85, 0x30, 0x2b, 0x4f, 0x27, 0xde, 0xa9,
	0x6e, 0xd2, 0x4e, 0x98, 0x5a, 0xee, 0x4f, 0x35, 0xbd, 0x40, 0x2d, 0xcc, 0x2f, 0xf6, 0xe7, 0xb1,
	0xed, 0x88, 0xe8, 0xf2, 0x7c, 0x98, 0xe4, 0x4e, 0x43, 0x77, 0x85, 0x29, 0xbc, 0x48, 0xb0, 0x3c,
	0xc0, 0x8a, 0x6e, 0x83, 0xdb, 0x8e, 0x60, 0x3c, 0x40, 0xa8, 0x3f, 0x15, 0xc8, 0xee, 0x47, 0x9d,
	0x11, 0x02, 0x69, 0x6a, 0xb6, 0xb1, 0xa8, 0x94, 0x95, 0x4a, 0xd6, 0x90, 0xdf, 0xe4, 0x01, 0x64,
	0xda, 0x28, 0xcc, 0xa6, 0x29, 0xcc, 0x62, 0xaa, 0xac, 0x54, 0xa6, 0x6a, 0xb3, 0x5a, 0xd8, 0x79,
	0x44, 0xab, 0x3d, 0xa6, 0x1d, 0xa3, 0x8b, 0xf2, 0x59, 0x9a, 0x8c, 0x62, 0x71, 0xbc, 0xac, 0x54,
	0x32, 0x86, 0xfc, 0x26, 0x55, 0x98, 0x40, 0xce, 0x19, 0x2f, 0xa6, 0x25, 0x05, 0x89, 0x28, 0xb8,
	0xd3, 0xd0, 0x0e, 0x65, 0xc9, 0xbb, 0x63, 0x46, 0x00, 0x21, 0x35, 0xc8, 0x70, 0x74, 0x1d, 0x46,
	0x5d, 0x2c, 0x4e, 0x0c, 0x57, 0xdc, 0x1d, 0x33, 0xba, 0xb8, 0xad, 0x0c, 0x4c, 0x72, 0x74, 0xbd,
	0x96, 0x50, 0x57, 0xe1, 0xe6, 0x0e, 0x8a, 0x6e, 0x4f, 0x06, 0xbe, 0xf5, 0xd0, 0x15, 0x49, 0xad,
	0xa9, 0x97, 0x30, 0xb7, 0x67, 0xbb, 0x3d, 0xac, 0xdb, 0x0f, 0x4e, 0xc7, 0xe6, 0x50, 0x80, 0xc9,
	0x53, 0xbb, 0x25, 0x90, 0x87, 0x14, 0xe1, 0x89, 0x2c, 0x42, 0xd6, 0x31, 0x2d, 0xac, 0xbb, 0xf6,
	0x39, 0xca, 0x01, 0x4d, 0x18, 0x19, 0x3f, 0x70, 0x68, 0x9f, 0x23, 0x59, 0x02, 0x90, 0x49, 0xc1,
	0xde, 0x20, 0x95, 0x03, 0xc9, 0x1a, 0x12, 0x7e, 0xe4, 0x07, 0xd4, 0x4b, 0x28, 0xf4, 0x17, 0x10,
	0xf4, 0x43, 0x36, 0x01, 0x7a, 0x86, 0x2b, 0x2a, 0xe5, 0xf1, 0xca, 0x54, 0x6d, 0x49, 0x1b, 0x74,
	0x9c, 0xd6, 0x6b, 0x34, 0x76, 0x81, 0xdc, 0x83, 0x3c, 0xc5, 0x33, 0x51, 0x8f, 0x89, 0xa7, 0xa4,
	0x78, 0xce, 0x0f, 0x1f, 0x74, 0x0b, 0x58, 0x83, 0xc2, 0x13, 0x93, 0x36, 0xb0, 0x35, 0xd2, 0xbc,
	0xd6, 0xa0, 0xb0, 0x8d, 0x2d, 0x14, 0x38, 0x12, 0xba, 0x0e, 0xb3, 0x2f, 0x4c, 0x7b, 0xa4, 0x4d,
	0x90, 0x75, 0xb8, 0x21, 0xec, 0x36, 0x32, 0x4f, 0x84, 0x1e, 0x5b, 0x18, 0xd8, 0xf8, 0x76, 0xf8,
	0x20, 0x8c, 0x08, 0xa9, 0xbe, 0x84, 0x5c, 0x97, 0xfc, 0x19, 0x3d, 0x65, 0x64, 0x05, 0x72, 0x91,
	0x21, 0xea, 0xa2, 0xe3, 0x44, 0x12, 0xd3, 0x51, 0xf0, 0xa8, 0xe3, 0xa0, 0x0f, 0x8a, 0x9c, 0x1a,
	0x80, 0x82, 0xc1, 0x4c, 0x47, 0x41, 0x1f, 0x54, 0xfb, 0x93, 0x06, 0xe8, 0x6d, 0x85, 0x7c, 0x56,
	0x60, 0xe6, 0xea, 0xa2, 0xc8, 0x6a, 0xd2, 0x32, 0x12, 0xdd, 0x54, 0xaa, 0x8e, 0x02, 0x0d, 0x2a,
	0x54, 0x97, 0x3e, 0xfd, 0xfa, 0xfd, 0x25, 0x35, 0x4f, 0xe6, 0xf4, 0x77, 0x0f, 0xf5, 0x0f, 0xfe,
	0x6c, 0x36, 0x7b, 0x6b, 0xbd, 0x20, 0x67, 0x30, 0x1d, 0x37, 0x37, 0xb9, 0x9f, 0x44, 0x9d, 0x60,
	0xff, 0xd2, 0xf5, 0xde, 0x51, 0xcb, 0x52, 0xb6, 0x44, 0x8a, 0x49, 0xb2, 0x7a, 0xb5, 0x7a, 0x41,
	0xde, 0x43, 0xbe, 0x6f, 0xf7, 0x24, 0xb1, 0xaf, 0x64, 0x83, 0x94, 0x0a, 0x03, 0xfb, 0x7c, 0xea,
	0xff, 0x80, 0x45, 0xc2, 0xd5, 0xe1, 0xc2, 0x1f, 0x15, 0xc8, 0xf7, 0x79, 0x34, 0x59, 0x39, 0xd9,
	0xc8, 0x43, 0x95, 0xab, 0x52, 0xf9, 0xae, 0x7a, 0x7b, 0x98, 0xf2, 0x46, 0x43, 0x12, 0x6e, 0x28,
	0x55, 0x72, 0x0c, 0xb9, 0x2b, 0x56, 0x26, 0x95, 0xa4, 0x02, 0x92, 0xdc, 0xfe, 0xbf, 0xc1, 0x8f,
	0x6d, 0xd8, 0x30, 0xd3, 0xd5, 0xad, 0xdb, 0xbe, 0x8d, 0x97, 0x07, 0x2a, 0x7e, 0x8e, 0xe2, 0x35,
	0x6b, 0xee, 0x3b, 0xc1, 0xaf, 0xc1, 0xb7, 0x8c, 0x7c, 0x22, 0x77, 0xae, 0x65, 0xf6, 0x1f, 0x84,
	0x91, 0x63, 0xf1, 0xe3, 0xd6, 0x57, 0x05, 0x0a, 0x0d, 0xd6, 0x4e, 0xb8, 0xb7, 0x95, 0xef, 0x79,
	0xf1, 0xc0, 0x57, 0x3d, 0x50, 0x8e, 0x37, 0x43, 0x98, 0xc5, 0x5a, 0x26, 0xb5, 0x34, 0xc6, 0x2d,
	0xdd, 0x42, 0x2a, 0x6b, 0xd2, 0x83, 0x94, 0xe9, 0xd8, 0x6e, 0xfc, 0x3f, 0xf2, 0x51, 0xec, 0xfb,
	0xaf, 0xa2, 0x7c, 0x4f, 0x91, 0x9d, 0x80, 0x62, 0x8f, 0x51, 0xcb, 0x08, 0xe2, 0x3f, 0xa2, 0xe0,
	0xab, 0x58, 0xf0, 0x64, 0x52, 0xd2, 0xae, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x06, 0x86, 0x63,
	0xe5, 0x79, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OperationsClient is the client API for Operations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperationsClient interface {
	// Lists operations that match the specified filter in the request. If the
	// server doesn't support this method, it returns `UNIMPLEMENTED`.
	//
	// NOTE: the `name` binding allows API services to override the binding
	// to use different resource name schemes, such as `users/*/operations`. To
	// override the binding, API services can add a binding such as
	// `"/v1/{name=users/*}/operations"` to their service configuration.
	// For backwards compatibility, the default name includes the operations
	// collection id, however overriding users must ensure the name binding
	// is the parent resource, without the operations collection id.
	ListOperations(ctx context.Context, in *ListOperationsRequest, opts ...grpc.CallOption) (*ListOperationsResponse, error)
	// Gets the latest state of a long-running operation.  Clients can use this
	// method to poll the operation result at intervals as recommended by the API
	// service.
	GetOperation(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*Operation, error)
	// Deletes a long-running operation. This method indicates that the client is
	// no longer interested in the operation result. It does not cancel the
	// operation. If the server doesn't support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.
	DeleteOperation(ctx context.Context, in *DeleteOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Starts asynchronous cancellation on a long-running operation.  The server
	// makes a best effort to cancel the operation, but success is not
	// guaranteed.  If the server doesn't support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.  Clients can use
	// [Operations.GetOperation][google.longrunning.Operations.GetOperation] or
	// other methods to check whether the cancellation succeeded or whether the
	// operation completed despite cancellation. On successful cancellation,
	// the operation is not deleted; instead, it becomes an operation with
	// an [Operation.error][google.longrunning.Operation.error] value with a [google.rpc.Status.code][google.rpc.Status.code] of 1,
	// corresponding to `Code.CANCELLED`.
	CancelOperation(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Waits for the specified long-running operation until it is done or reaches
	// at most a specified timeout, returning the latest state.  If the operation
	// is already done, the latest state is immediately returned.  If the timeout
	// specified is greater than the default HTTP/RPC timeout, the HTTP/RPC
	// timeout is used.  If the server does not support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.
	// Note that this method is on a best-effort basis.  It may return the latest
	// state before the specified timeout (including immediately), meaning even an
	// immediate response is no guarantee that the operation is done.
	WaitOperation(ctx context.Context, in *WaitOperationRequest, opts ...grpc.CallOption) (*Operation, error)
}

type operationsClient struct {
	cc *grpc.ClientConn
}

func NewOperationsClient(cc *grpc.ClientConn) OperationsClient {
	return &operationsClient{cc}
}

func (c *operationsClient) ListOperations(ctx context.Context, in *ListOperationsRequest, opts ...grpc.CallOption) (*ListOperationsResponse, error) {
	out := new(ListOperationsResponse)
	err := c.cc.Invoke(ctx, "/google.longrunning.Operations/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) GetOperation(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/google.longrunning.Operations/GetOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) DeleteOperation(ctx context.Context, in *DeleteOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.longrunning.Operations/DeleteOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) CancelOperation(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.longrunning.Operations/CancelOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) WaitOperation(ctx context.Context, in *WaitOperationRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/google.longrunning.Operations/WaitOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationsServer is the server API for Operations service.
type OperationsServer interface {
	// Lists operations that match the specified filter in the request. If the
	// server doesn't support this method, it returns `UNIMPLEMENTED`.
	//
	// NOTE: the `name` binding allows API services to override the binding
	// to use different resource name schemes, such as `users/*/operations`. To
	// override the binding, API services can add a binding such as
	// `"/v1/{name=users/*}/operations"` to their service configuration.
	// For backwards compatibility, the default name includes the operations
	// collection id, however overriding users must ensure the name binding
	// is the parent resource, without the operations collection id.
	ListOperations(context.Context, *ListOperationsRequest) (*ListOperationsResponse, error)
	// Gets the latest state of a long-running operation.  Clients can use this
	// method to poll the operation result at intervals as recommended by the API
	// service.
	GetOperation(context.Context, *GetOperationRequest) (*Operation, error)
	// Deletes a long-running operation. This method indicates that the client is
	// no longer interested in the operation result. It does not cancel the
	// operation. If the server doesn't support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.
	DeleteOperation(context.Context, *DeleteOperationRequest) (*empty.Empty, error)
	// Starts asynchronous cancellation on a long-running operation.  The server
	// makes a best effort to cancel the operation, but success is not
	// guaranteed.  If the server doesn't support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.  Clients can use
	// [Operations.GetOperation][google.longrunning.Operations.GetOperation] or
	// other methods to check whether the cancellation succeeded or whether the
	// operation completed despite cancellation. On successful cancellation,
	// the operation is not deleted; instead, it becomes an operation with
	// an [Operation.error][google.longrunning.Operation.error] value with a [google.rpc.Status.code][google.rpc.Status.code] of 1,
	// corresponding to `Code.CANCELLED`.
	CancelOperation(context.Context, *CancelOperationRequest) (*empty.Empty, error)
	// Waits for the specified long-running operation until it is done or reaches
	// at most a specified timeout, returning the latest state.  If the operation
	// is already done, the latest state is immediately returned.  If the timeout
	// specified is greater than the default HTTP/RPC timeout, the HTTP/RPC
	// timeout is used.  If the server does not support this method, it returns
	// `google.rpc.Code.UNIMPLEMENTED`.
	// Note that this method is on a best-effort basis.  It may return the latest
	// state before the specified timeout (including immediately), meaning even an
	// immediate response is no guarantee that the operation is done.
	WaitOperation(context.Context, *WaitOperationRequest) (*Operation, error)
}

func RegisterOperationsServer(s *grpc.Server, srv OperationsServer) {
	s.RegisterService(&_Operations_serviceDesc, srv)
}

func _Operations_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.longrunning.Operations/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).ListOperations(ctx, req.(*ListOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_GetOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).GetOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.longrunning.Operations/GetOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).GetOperation(ctx, req.(*GetOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_DeleteOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).DeleteOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.longrunning.Operations/DeleteOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).DeleteOperation(ctx, req.(*DeleteOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_CancelOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).CancelOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.longrunning.Operations/CancelOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).CancelOperation(ctx, req.(*CancelOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_WaitOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).WaitOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.longrunning.Operations/WaitOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).WaitOperation(ctx, req.(*WaitOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Operations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.longrunning.Operations",
	HandlerType: (*OperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOperations",
			Handler:    _Operations_ListOperations_Handler,
		},
		{
			MethodName: "GetOperation",
			Handler:    _Operations_GetOperation_Handler,
		},
		{
			MethodName: "DeleteOperation",
			Handler:    _Operations_DeleteOperation_Handler,
		},
		{
			MethodName: "CancelOperation",
			Handler:    _Operations_CancelOperation_Handler,
		},
		{
			MethodName: "WaitOperation",
			Handler:    _Operations_WaitOperation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/longrunning/operations.proto",
}
