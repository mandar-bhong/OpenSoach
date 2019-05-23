// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/ad_group_criterion_label_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for
// [AdGroupCriterionLabelService.GetAdGroupCriterionLabel][google.ads.googleads.v1.services.AdGroupCriterionLabelService.GetAdGroupCriterionLabel].
type GetAdGroupCriterionLabelRequest struct {
	// The resource name of the ad group criterion label to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupCriterionLabelRequest) Reset()         { *m = GetAdGroupCriterionLabelRequest{} }
func (m *GetAdGroupCriterionLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupCriterionLabelRequest) ProtoMessage()    {}
func (*GetAdGroupCriterionLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_812d424f598fe7ab, []int{0}
}

func (m *GetAdGroupCriterionLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupCriterionLabelRequest.Unmarshal(m, b)
}
func (m *GetAdGroupCriterionLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupCriterionLabelRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupCriterionLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupCriterionLabelRequest.Merge(m, src)
}
func (m *GetAdGroupCriterionLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupCriterionLabelRequest.Size(m)
}
func (m *GetAdGroupCriterionLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupCriterionLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupCriterionLabelRequest proto.InternalMessageInfo

func (m *GetAdGroupCriterionLabelRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [AdGroupCriterionLabelService.MutateAdGroupCriterionLabels][google.ads.googleads.v1.services.AdGroupCriterionLabelService.MutateAdGroupCriterionLabels].
type MutateAdGroupCriterionLabelsRequest struct {
	// ID of the customer whose ad group criterion labels are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on ad group criterion labels.
	Operations []*AdGroupCriterionLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupCriterionLabelsRequest) Reset()         { *m = MutateAdGroupCriterionLabelsRequest{} }
func (m *MutateAdGroupCriterionLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriterionLabelsRequest) ProtoMessage()    {}
func (*MutateAdGroupCriterionLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_812d424f598fe7ab, []int{1}
}

func (m *MutateAdGroupCriterionLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupCriterionLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriterionLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.Merge(m, src)
}
func (m *MutateAdGroupCriterionLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.Size(m)
}
func (m *MutateAdGroupCriterionLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriterionLabelsRequest proto.InternalMessageInfo

func (m *MutateAdGroupCriterionLabelsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupCriterionLabelsRequest) GetOperations() []*AdGroupCriterionLabelOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupCriterionLabelsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupCriterionLabelsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an ad group criterion label.
type AdGroupCriterionLabelOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupCriterionLabelOperation_Create
	//	*AdGroupCriterionLabelOperation_Remove
	Operation            isAdGroupCriterionLabelOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *AdGroupCriterionLabelOperation) Reset()         { *m = AdGroupCriterionLabelOperation{} }
func (m *AdGroupCriterionLabelOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionLabelOperation) ProtoMessage()    {}
func (*AdGroupCriterionLabelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_812d424f598fe7ab, []int{2}
}

func (m *AdGroupCriterionLabelOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionLabelOperation.Unmarshal(m, b)
}
func (m *AdGroupCriterionLabelOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionLabelOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterionLabelOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionLabelOperation.Merge(m, src)
}
func (m *AdGroupCriterionLabelOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionLabelOperation.Size(m)
}
func (m *AdGroupCriterionLabelOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionLabelOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionLabelOperation proto.InternalMessageInfo

type isAdGroupCriterionLabelOperation_Operation interface {
	isAdGroupCriterionLabelOperation_Operation()
}

type AdGroupCriterionLabelOperation_Create struct {
	Create *resources.AdGroupCriterionLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupCriterionLabelOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*AdGroupCriterionLabelOperation_Create) isAdGroupCriterionLabelOperation_Operation() {}

func (*AdGroupCriterionLabelOperation_Remove) isAdGroupCriterionLabelOperation_Operation() {}

func (m *AdGroupCriterionLabelOperation) GetOperation() isAdGroupCriterionLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupCriterionLabelOperation) GetCreate() *resources.AdGroupCriterionLabel {
	if x, ok := m.GetOperation().(*AdGroupCriterionLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupCriterionLabelOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupCriterionLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupCriterionLabelOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupCriterionLabelOperation_Create)(nil),
		(*AdGroupCriterionLabelOperation_Remove)(nil),
	}
}

// Response message for an ad group criterion labels mutate.
type MutateAdGroupCriterionLabelsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupCriterionLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *MutateAdGroupCriterionLabelsResponse) Reset()         { *m = MutateAdGroupCriterionLabelsResponse{} }
func (m *MutateAdGroupCriterionLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriterionLabelsResponse) ProtoMessage()    {}
func (*MutateAdGroupCriterionLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_812d424f598fe7ab, []int{3}
}

func (m *MutateAdGroupCriterionLabelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupCriterionLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriterionLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.Merge(m, src)
}
func (m *MutateAdGroupCriterionLabelsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.Size(m)
}
func (m *MutateAdGroupCriterionLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriterionLabelsResponse proto.InternalMessageInfo

func (m *MutateAdGroupCriterionLabelsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupCriterionLabelsResponse) GetResults() []*MutateAdGroupCriterionLabelResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for an ad group criterion label mutate.
type MutateAdGroupCriterionLabelResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupCriterionLabelResult) Reset()         { *m = MutateAdGroupCriterionLabelResult{} }
func (m *MutateAdGroupCriterionLabelResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriterionLabelResult) ProtoMessage()    {}
func (*MutateAdGroupCriterionLabelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_812d424f598fe7ab, []int{4}
}

func (m *MutateAdGroupCriterionLabelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriterionLabelResult.Unmarshal(m, b)
}
func (m *MutateAdGroupCriterionLabelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriterionLabelResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriterionLabelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriterionLabelResult.Merge(m, src)
}
func (m *MutateAdGroupCriterionLabelResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriterionLabelResult.Size(m)
}
func (m *MutateAdGroupCriterionLabelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriterionLabelResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriterionLabelResult proto.InternalMessageInfo

func (m *MutateAdGroupCriterionLabelResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupCriterionLabelRequest)(nil), "google.ads.googleads.v1.services.GetAdGroupCriterionLabelRequest")
	proto.RegisterType((*MutateAdGroupCriterionLabelsRequest)(nil), "google.ads.googleads.v1.services.MutateAdGroupCriterionLabelsRequest")
	proto.RegisterType((*AdGroupCriterionLabelOperation)(nil), "google.ads.googleads.v1.services.AdGroupCriterionLabelOperation")
	proto.RegisterType((*MutateAdGroupCriterionLabelsResponse)(nil), "google.ads.googleads.v1.services.MutateAdGroupCriterionLabelsResponse")
	proto.RegisterType((*MutateAdGroupCriterionLabelResult)(nil), "google.ads.googleads.v1.services.MutateAdGroupCriterionLabelResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/ad_group_criterion_label_service.proto", fileDescriptor_812d424f598fe7ab)
}

var fileDescriptor_812d424f598fe7ab = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcf, 0x6b, 0xd4, 0x5e,
	0x10, 0xff, 0x26, 0xfb, 0xa5, 0xda, 0xb7, 0x55, 0xe1, 0x89, 0x18, 0x96, 0xd2, 0x6e, 0xd3, 0x82,
	0x65, 0x0f, 0x09, 0xbb, 0x5e, 0x4a, 0x4a, 0x61, 0xd3, 0xd2, 0xdd, 0x0a, 0x6a, 0x4b, 0x0a, 0x3d,
	0xc8, 0x4a, 0x7c, 0x9b, 0x4c, 0x43, 0x20, 0x9b, 0x17, 0xdf, 0x7b, 0x59, 0x29, 0xa5, 0x17, 0x4f,
	0xde, 0x3d, 0x7a, 0xf3, 0xe8, 0x9f, 0x22, 0x78, 0xf2, 0xe4, 0xdd, 0x93, 0x47, 0x41, 0xcf, 0x92,
	0x1f, 0x6f, 0x6d, 0x65, 0x77, 0x23, 0xf6, 0x36, 0x99, 0x99, 0xfd, 0x7c, 0xe6, 0x33, 0x33, 0x6f,
	0x16, 0xf5, 0x03, 0x4a, 0x83, 0x08, 0x4c, 0xe2, 0x73, 0xb3, 0x30, 0x33, 0x6b, 0xdc, 0x36, 0x39,
	0xb0, 0x71, 0xe8, 0x01, 0x37, 0x89, 0xef, 0x06, 0x8c, 0xa6, 0x89, 0xeb, 0xb1, 0x50, 0x00, 0x0b,
	0x69, 0xec, 0x46, 0x64, 0x08, 0x91, 0x5b, 0x66, 0x18, 0x09, 0xa3, 0x82, 0xe2, 0x66, 0xf1, 0x6b,
	0x83, 0xf8, 0xdc, 0x98, 0x00, 0x19, 0xe3, 0xb6, 0x21, 0x81, 0x1a, 0xdd, 0x59, 0x54, 0x0c, 0x38,
	0x4d, 0xd9, 0x3c, 0xae, 0x82, 0xa3, 0xb1, 0x2c, 0x11, 0x92, 0xd0, 0x24, 0x71, 0x4c, 0x05, 0x11,
	0x21, 0x8d, 0x79, 0x19, 0x5d, 0x29, 0xa3, 0xf9, 0xd7, 0x30, 0x3d, 0x35, 0x5f, 0x31, 0x92, 0x24,
	0xc0, 0x64, 0xfc, 0x7e, 0x19, 0x67, 0x89, 0x67, 0x72, 0x41, 0x44, 0x5a, 0x06, 0xf4, 0x1e, 0x5a,
	0xed, 0x83, 0xb0, 0xfd, 0x7e, 0x46, 0xbd, 0x27, 0x99, 0x1f, 0x67, 0xc4, 0x0e, 0xbc, 0x4c, 0x81,
	0x0b, 0xbc, 0x8e, 0x6e, 0xc9, 0x2a, 0xdd, 0x98, 0x8c, 0x40, 0x53, 0x9a, 0xca, 0xe6, 0xa2, 0xb3,
	0x24, 0x9d, 0x4f, 0xc9, 0x08, 0xf4, 0x9f, 0x0a, 0x5a, 0x7f, 0x92, 0x0a, 0x22, 0x60, 0x2a, 0x16,
	0x97, 0x60, 0xab, 0xa8, 0xee, 0xa5, 0x5c, 0xd0, 0x11, 0x30, 0x37, 0xf4, 0x4b, 0x28, 0x24, 0x5d,
	0x8f, 0x7c, 0xfc, 0x02, 0x21, 0x9a, 0x00, 0x2b, 0xd4, 0x69, 0x6a, 0xb3, 0xb6, 0x59, 0xef, 0x74,
	0x8d, 0xaa, 0x06, 0x1b, 0x53, 0x59, 0x0f, 0x25, 0x90, 0x73, 0x09, 0x13, 0x3f, 0x40, 0x77, 0x12,
	0xc2, 0x44, 0x48, 0x22, 0xf7, 0x94, 0x84, 0x51, 0xca, 0x40, 0xab, 0x35, 0x95, 0xcd, 0x9b, 0xce,
	0xed, 0xd2, 0xdd, 0x2b, 0xbc, 0x99, 0xf0, 0x31, 0x89, 0x42, 0x9f, 0x08, 0x70, 0x69, 0x1c, 0x9d,
	0x69, 0xff, 0xe7, 0x69, 0x4b, 0xd2, 0x79, 0x18, 0x47, 0x67, 0xfa, 0x3b, 0x05, 0xad, 0xcc, 0x27,
	0xc7, 0x0e, 0x5a, 0xf0, 0x18, 0x10, 0x51, 0x74, 0xae, 0xde, 0xd9, 0x9a, 0x29, 0x67, 0xb2, 0x0d,
	0xd3, 0xf5, 0x1c, 0xfc, 0xe7, 0x94, 0x48, 0x58, 0x43, 0x0b, 0x0c, 0x46, 0x74, 0x0c, 0x9a, 0x9a,
	0xb5, 0x30, 0x8b, 0x14, 0xdf, 0xbb, 0x75, 0xb4, 0x38, 0x11, 0xab, 0x7f, 0x52, 0xd0, 0xc6, 0xfc,
	0xb1, 0xf0, 0x84, 0xc6, 0x1c, 0x70, 0x0f, 0xdd, 0xfb, 0xa3, 0x29, 0x2e, 0x30, 0x46, 0x59, 0xde,
	0x9a, 0x7a, 0x07, 0xcb, 0x92, 0x59, 0xe2, 0x19, 0xc7, 0xf9, 0x02, 0x39, 0x77, 0xaf, 0xb6, 0x6b,
	0x3f, 0x4b, 0xc7, 0xcf, 0xd1, 0x0d, 0x06, 0x3c, 0x8d, 0x84, 0x9c, 0xdd, 0x5e, 0xf5, 0xec, 0xe6,
	0x14, 0xe8, 0xe4, 0x58, 0x8e, 0xc4, 0xd4, 0x0f, 0xd0, 0x5a, 0x65, 0xf6, 0x5f, 0x2d, 0x6c, 0xe7,
	0x7b, 0x0d, 0x2d, 0x4f, 0x05, 0x39, 0x2e, 0xca, 0xc2, 0x5f, 0x14, 0xa4, 0xcd, 0x7a, 0x1a, 0xd8,
	0xae, 0x56, 0x55, 0xf1, 0xac, 0x1a, 0xff, 0xbc, 0x05, 0x7a, 0xf7, 0xf5, 0xe7, 0xaf, 0x6f, 0x55,
	0x0b, 0x6f, 0x65, 0x07, 0xe4, 0xfc, 0x8a, 0xd4, 0x1d, 0xf9, 0x92, 0xb8, 0xd9, 0x32, 0xc9, 0xd4,
	0x91, 0x9b, 0xad, 0x0b, 0xfc, 0x43, 0x41, 0xcb, 0xf3, 0xd6, 0x02, 0xef, 0x5f, 0x6b, 0x6a, 0xf2,
	0xb5, 0x37, 0x7a, 0xd7, 0x85, 0x29, 0xb6, 0x53, 0xef, 0xe5, 0x8a, 0xbb, 0xfa, 0x76, 0xa6, 0xf8,
	0xb7, 0xc4, 0xf3, 0x4b, 0xa7, 0x64, 0xa7, 0x75, 0x31, 0x43, 0xb0, 0x35, 0xca, 0x29, 0x2c, 0xa5,
	0xb5, 0xfb, 0x46, 0x45, 0x1b, 0x1e, 0x1d, 0x55, 0x56, 0xb5, 0xbb, 0x36, 0x6f, 0x35, 0x8e, 0xb2,
	0xcb, 0x79, 0xa4, 0x3c, 0x3b, 0x28, 0x61, 0x02, 0x1a, 0x91, 0x38, 0x30, 0x28, 0x0b, 0xcc, 0x00,
	0xe2, 0xfc, 0xae, 0xca, 0x23, 0x9f, 0x84, 0x7c, 0xf6, 0xdf, 0xcb, 0xb6, 0x34, 0xde, 0xab, 0xb5,
	0xbe, 0x6d, 0x7f, 0x50, 0x9b, 0xfd, 0x02, 0xd0, 0xf6, 0xb9, 0x51, 0x98, 0x99, 0x75, 0xd2, 0x36,
	0x4a, 0x62, 0xfe, 0x51, 0xa6, 0x0c, 0x6c, 0x9f, 0x0f, 0x26, 0x29, 0x83, 0x93, 0xf6, 0x40, 0xa6,
	0x7c, 0x53, 0x37, 0x0a, 0xbf, 0x65, 0xd9, 0x3e, 0xb7, 0xac, 0x49, 0x92, 0x65, 0x9d, 0xb4, 0x2d,
	0x4b, 0xa6, 0x0d, 0x17, 0xf2, 0x3a, 0x1f, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x21, 0x33,
	0x9c, 0x05, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdGroupCriterionLabelServiceClient is the client API for AdGroupCriterionLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupCriterionLabelServiceClient interface {
	// Returns the requested ad group criterion label in full detail.
	GetAdGroupCriterionLabel(ctx context.Context, in *GetAdGroupCriterionLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterionLabel, error)
	// Creates and removes ad group criterion labels.
	// Operation statuses are returned.
	MutateAdGroupCriterionLabels(ctx context.Context, in *MutateAdGroupCriterionLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupCriterionLabelsResponse, error)
}

type adGroupCriterionLabelServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdGroupCriterionLabelServiceClient(cc *grpc.ClientConn) AdGroupCriterionLabelServiceClient {
	return &adGroupCriterionLabelServiceClient{cc}
}

func (c *adGroupCriterionLabelServiceClient) GetAdGroupCriterionLabel(ctx context.Context, in *GetAdGroupCriterionLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterionLabel, error) {
	out := new(resources.AdGroupCriterionLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupCriterionLabelService/GetAdGroupCriterionLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupCriterionLabelServiceClient) MutateAdGroupCriterionLabels(ctx context.Context, in *MutateAdGroupCriterionLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupCriterionLabelsResponse, error) {
	out := new(MutateAdGroupCriterionLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupCriterionLabelService/MutateAdGroupCriterionLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupCriterionLabelServiceServer is the server API for AdGroupCriterionLabelService service.
type AdGroupCriterionLabelServiceServer interface {
	// Returns the requested ad group criterion label in full detail.
	GetAdGroupCriterionLabel(context.Context, *GetAdGroupCriterionLabelRequest) (*resources.AdGroupCriterionLabel, error)
	// Creates and removes ad group criterion labels.
	// Operation statuses are returned.
	MutateAdGroupCriterionLabels(context.Context, *MutateAdGroupCriterionLabelsRequest) (*MutateAdGroupCriterionLabelsResponse, error)
}

func RegisterAdGroupCriterionLabelServiceServer(s *grpc.Server, srv AdGroupCriterionLabelServiceServer) {
	s.RegisterService(&_AdGroupCriterionLabelService_serviceDesc, srv)
}

func _AdGroupCriterionLabelService_GetAdGroupCriterionLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupCriterionLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionLabelServiceServer).GetAdGroupCriterionLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupCriterionLabelService/GetAdGroupCriterionLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionLabelServiceServer).GetAdGroupCriterionLabel(ctx, req.(*GetAdGroupCriterionLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupCriterionLabelService_MutateAdGroupCriterionLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupCriterionLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionLabelServiceServer).MutateAdGroupCriterionLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupCriterionLabelService/MutateAdGroupCriterionLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionLabelServiceServer).MutateAdGroupCriterionLabels(ctx, req.(*MutateAdGroupCriterionLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupCriterionLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.AdGroupCriterionLabelService",
	HandlerType: (*AdGroupCriterionLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupCriterionLabel",
			Handler:    _AdGroupCriterionLabelService_GetAdGroupCriterionLabel_Handler,
		},
		{
			MethodName: "MutateAdGroupCriterionLabels",
			Handler:    _AdGroupCriterionLabelService_MutateAdGroupCriterionLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/ad_group_criterion_label_service.proto",
}
