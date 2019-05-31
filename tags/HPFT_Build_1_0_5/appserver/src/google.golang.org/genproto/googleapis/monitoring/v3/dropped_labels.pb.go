// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/dropped_labels.proto

package monitoring

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

// A set of (label, value) pairs which were dropped during aggregation, attached
// to google.api.Distribution.Exemplars in google.api.Distribution values during
// aggregation.
//
// These values are used in combination with the label values that remain on the
// aggregated Distribution timeseries to construct the full label set for the
// exemplar values.  The resulting full label set may be used to identify the
// specific task/job/instance (for example) which may be contributing to a
// long-tail, while allowing the storage savings of only storing aggregated
// distribution values for a large group.
//
// Note that there are no guarantees on ordering of the labels from
// exemplar-to-exemplar and from distribution-to-distribution in the same
// stream, and there may be duplicates.  It is up to clients to resolve any
// ambiguities.
type DroppedLabels struct {
	// Map from label to its value, for all labels dropped in any aggregation.
	Label                map[string]string `protobuf:"bytes,1,rep,name=label,proto3" json:"label,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DroppedLabels) Reset()         { *m = DroppedLabels{} }
func (m *DroppedLabels) String() string { return proto.CompactTextString(m) }
func (*DroppedLabels) ProtoMessage()    {}
func (*DroppedLabels) Descriptor() ([]byte, []int) {
	return fileDescriptor_15749142c06d7f43, []int{0}
}

func (m *DroppedLabels) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DroppedLabels.Unmarshal(m, b)
}
func (m *DroppedLabels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DroppedLabels.Marshal(b, m, deterministic)
}
func (m *DroppedLabels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DroppedLabels.Merge(m, src)
}
func (m *DroppedLabels) XXX_Size() int {
	return xxx_messageInfo_DroppedLabels.Size(m)
}
func (m *DroppedLabels) XXX_DiscardUnknown() {
	xxx_messageInfo_DroppedLabels.DiscardUnknown(m)
}

var xxx_messageInfo_DroppedLabels proto.InternalMessageInfo

func (m *DroppedLabels) GetLabel() map[string]string {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*DroppedLabels)(nil), "google.monitoring.v3.DroppedLabels")
	proto.RegisterMapType((map[string]string)(nil), "google.monitoring.v3.DroppedLabels.LabelEntry")
}

func init() {
	proto.RegisterFile("google/monitoring/v3/dropped_labels.proto", fileDescriptor_15749142c06d7f43)
}

var fileDescriptor_15749142c06d7f43 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4b, 0x43, 0x31,
	0x10, 0xc7, 0x49, 0x4b, 0x05, 0x4f, 0x04, 0x09, 0x1d, 0x42, 0x71, 0x28, 0x4e, 0x75, 0xb9, 0x80,
	0x6f, 0x29, 0x0a, 0x0e, 0xa5, 0x6e, 0x0e, 0xa5, 0xa3, 0x8b, 0xa4, 0x36, 0x84, 0x60, 0x7a, 0x17,
	0xd2, 0xf8, 0xa0, 0x9f, 0xc2, 0xaf, 0x5c, 0x9a, 0x3c, 0x78, 0x3c, 0xe8, 0x94, 0xdc, 0xe5, 0xf7,
	0xff, 0x1d, 0x17, 0x78, 0x76, 0xcc, 0x2e, 0x58, 0x7d, 0x60, 0xf2, 0x99, 0x93, 0x27, 0xa7, 0xdb,
	0x46, 0xef, 0x13, 0xc7, 0x68, 0xf7, 0xdf, 0xc1, 0xec, 0x6c, 0x38, 0x62, 0x4c, 0x9c, 0x59, 0x4e,
	0x2b, 0x8a, 0x3d, 0x8a, 0x6d, 0x33, 0x7b, 0xec, 0x04, 0x26, 0x7a, 0x6d, 0x88, 0x38, 0x9b, 0xec,
	0x99, 0xba, 0xcc, 0xd3, 0xbf, 0x80, 0xfb, 0x75, 0x95, 0x7d, 0x16, 0x97, 0x5c, 0xc3, 0xa4, 0x58,
	0x95, 0x98, 0x8f, 0x17, 0x77, 0x2f, 0x88, 0xd7, 0xac, 0x38, 0xc8, 0x60, 0x39, 0x3e, 0x28, 0xa7,
	0xd3, 0xb6, 0x86, 0x67, 0x4b, 0x80, 0xbe, 0x29, 0x1f, 0x60, 0xfc, 0x6b, 0x4f, 0x4a, 0xcc, 0xc5,
	0xe2, 0x76, 0x7b, 0xb9, 0xca, 0x29, 0x4c, 0x5a, 0x13, 0xfe, 0xac, 0x1a, 0x95, 0x5e, 0x2d, 0x5e,
	0x47, 0x4b, 0xb1, 0x8a, 0xa0, 0x7e, 0xf8, 0x70, 0x75, 0xea, 0x4a, 0x0e, 0xc6, 0x6e, 0x2e, 0x1b,
	0x6c, 0xc4, 0xd7, 0x7b, 0xc7, 0x3a, 0x0e, 0x86, 0x1c, 0x72, 0x72, 0xda, 0x59, 0x2a, 0xfb, 0xe9,
	0xfa, 0x64, 0xa2, 0x3f, 0x0e, 0x7f, 0xf0, 0xad, 0xaf, 0x76, 0x37, 0x05, 0x6d, 0xce, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x41, 0xd0, 0x8d, 0xe9, 0x6b, 0x01, 0x00, 0x00,
}
