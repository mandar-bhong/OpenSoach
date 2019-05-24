// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grafeas/v1/discovery.proto

package grafeas

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Whether the resource is continuously analyzed.
type DiscoveryOccurrence_ContinuousAnalysis int32

const (
	// Unknown.
	DiscoveryOccurrence_CONTINUOUS_ANALYSIS_UNSPECIFIED DiscoveryOccurrence_ContinuousAnalysis = 0
	// The resource is continuously analyzed.
	DiscoveryOccurrence_ACTIVE DiscoveryOccurrence_ContinuousAnalysis = 1
	// The resource is ignored for continuous analysis.
	DiscoveryOccurrence_INACTIVE DiscoveryOccurrence_ContinuousAnalysis = 2
)

var DiscoveryOccurrence_ContinuousAnalysis_name = map[int32]string{
	0: "CONTINUOUS_ANALYSIS_UNSPECIFIED",
	1: "ACTIVE",
	2: "INACTIVE",
}

var DiscoveryOccurrence_ContinuousAnalysis_value = map[string]int32{
	"CONTINUOUS_ANALYSIS_UNSPECIFIED": 0,
	"ACTIVE":                          1,
	"INACTIVE":                        2,
}

func (x DiscoveryOccurrence_ContinuousAnalysis) String() string {
	return proto.EnumName(DiscoveryOccurrence_ContinuousAnalysis_name, int32(x))
}

func (DiscoveryOccurrence_ContinuousAnalysis) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ca027ecd4454b7df, []int{1, 0}
}

// Analysis status for a resource. Currently for initial analysis only (not
// updated in continuous analysis).
type DiscoveryOccurrence_AnalysisStatus int32

const (
	// Unknown.
	DiscoveryOccurrence_ANALYSIS_STATUS_UNSPECIFIED DiscoveryOccurrence_AnalysisStatus = 0
	// Resource is known but no action has been taken yet.
	DiscoveryOccurrence_PENDING DiscoveryOccurrence_AnalysisStatus = 1
	// Resource is being analyzed.
	DiscoveryOccurrence_SCANNING DiscoveryOccurrence_AnalysisStatus = 2
	// Analysis has finished successfully.
	DiscoveryOccurrence_FINISHED_SUCCESS DiscoveryOccurrence_AnalysisStatus = 3
	// Analysis has finished unsuccessfully, the analysis itself is in a bad
	// state.
	DiscoveryOccurrence_FINISHED_FAILED DiscoveryOccurrence_AnalysisStatus = 4
	// The resource is known not to be supported
	DiscoveryOccurrence_FINISHED_UNSUPPORTED DiscoveryOccurrence_AnalysisStatus = 5
)

var DiscoveryOccurrence_AnalysisStatus_name = map[int32]string{
	0: "ANALYSIS_STATUS_UNSPECIFIED",
	1: "PENDING",
	2: "SCANNING",
	3: "FINISHED_SUCCESS",
	4: "FINISHED_FAILED",
	5: "FINISHED_UNSUPPORTED",
}

var DiscoveryOccurrence_AnalysisStatus_value = map[string]int32{
	"ANALYSIS_STATUS_UNSPECIFIED": 0,
	"PENDING":                     1,
	"SCANNING":                    2,
	"FINISHED_SUCCESS":            3,
	"FINISHED_FAILED":             4,
	"FINISHED_UNSUPPORTED":        5,
}

func (x DiscoveryOccurrence_AnalysisStatus) String() string {
	return proto.EnumName(DiscoveryOccurrence_AnalysisStatus_name, int32(x))
}

func (DiscoveryOccurrence_AnalysisStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ca027ecd4454b7df, []int{1, 1}
}

// A note that indicates a type of analysis a provider would perform. This note
// exists in a provider's project. A `Discovery` occurrence is created in a
// consumer's project at the start of analysis.
type DiscoveryNote struct {
	// Required. Immutable. The kind of analysis that is handled by this
	// discovery.
	AnalysisKind         NoteKind `protobuf:"varint,1,opt,name=analysis_kind,json=analysisKind,proto3,enum=grafeas.v1.NoteKind" json:"analysis_kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryNote) Reset()         { *m = DiscoveryNote{} }
func (m *DiscoveryNote) String() string { return proto.CompactTextString(m) }
func (*DiscoveryNote) ProtoMessage()    {}
func (*DiscoveryNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca027ecd4454b7df, []int{0}
}

func (m *DiscoveryNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryNote.Unmarshal(m, b)
}
func (m *DiscoveryNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryNote.Marshal(b, m, deterministic)
}
func (m *DiscoveryNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryNote.Merge(m, src)
}
func (m *DiscoveryNote) XXX_Size() int {
	return xxx_messageInfo_DiscoveryNote.Size(m)
}
func (m *DiscoveryNote) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryNote.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryNote proto.InternalMessageInfo

func (m *DiscoveryNote) GetAnalysisKind() NoteKind {
	if m != nil {
		return m.AnalysisKind
	}
	return NoteKind_NOTE_KIND_UNSPECIFIED
}

// Provides information about the analysis status of a discovered resource.
type DiscoveryOccurrence struct {
	// Whether the resource is continuously analyzed.
	ContinuousAnalysis DiscoveryOccurrence_ContinuousAnalysis `protobuf:"varint,1,opt,name=continuous_analysis,json=continuousAnalysis,proto3,enum=grafeas.v1.DiscoveryOccurrence_ContinuousAnalysis" json:"continuous_analysis,omitempty"`
	// The status of discovery for the resource.
	AnalysisStatus DiscoveryOccurrence_AnalysisStatus `protobuf:"varint,2,opt,name=analysis_status,json=analysisStatus,proto3,enum=grafeas.v1.DiscoveryOccurrence_AnalysisStatus" json:"analysis_status,omitempty"`
	// When an error is encountered this will contain a LocalizedMessage under
	// details to show to the user. The LocalizedMessage is output only and
	// populated by the API.
	AnalysisStatusError  *status.Status `protobuf:"bytes,3,opt,name=analysis_status_error,json=analysisStatusError,proto3" json:"analysis_status_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiscoveryOccurrence) Reset()         { *m = DiscoveryOccurrence{} }
func (m *DiscoveryOccurrence) String() string { return proto.CompactTextString(m) }
func (*DiscoveryOccurrence) ProtoMessage()    {}
func (*DiscoveryOccurrence) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca027ecd4454b7df, []int{1}
}

func (m *DiscoveryOccurrence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryOccurrence.Unmarshal(m, b)
}
func (m *DiscoveryOccurrence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryOccurrence.Marshal(b, m, deterministic)
}
func (m *DiscoveryOccurrence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryOccurrence.Merge(m, src)
}
func (m *DiscoveryOccurrence) XXX_Size() int {
	return xxx_messageInfo_DiscoveryOccurrence.Size(m)
}
func (m *DiscoveryOccurrence) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryOccurrence.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryOccurrence proto.InternalMessageInfo

func (m *DiscoveryOccurrence) GetContinuousAnalysis() DiscoveryOccurrence_ContinuousAnalysis {
	if m != nil {
		return m.ContinuousAnalysis
	}
	return DiscoveryOccurrence_CONTINUOUS_ANALYSIS_UNSPECIFIED
}

func (m *DiscoveryOccurrence) GetAnalysisStatus() DiscoveryOccurrence_AnalysisStatus {
	if m != nil {
		return m.AnalysisStatus
	}
	return DiscoveryOccurrence_ANALYSIS_STATUS_UNSPECIFIED
}

func (m *DiscoveryOccurrence) GetAnalysisStatusError() *status.Status {
	if m != nil {
		return m.AnalysisStatusError
	}
	return nil
}

func init() {
	proto.RegisterEnum("grafeas.v1.DiscoveryOccurrence_ContinuousAnalysis", DiscoveryOccurrence_ContinuousAnalysis_name, DiscoveryOccurrence_ContinuousAnalysis_value)
	proto.RegisterEnum("grafeas.v1.DiscoveryOccurrence_AnalysisStatus", DiscoveryOccurrence_AnalysisStatus_name, DiscoveryOccurrence_AnalysisStatus_value)
	proto.RegisterType((*DiscoveryNote)(nil), "grafeas.v1.DiscoveryNote")
	proto.RegisterType((*DiscoveryOccurrence)(nil), "grafeas.v1.DiscoveryOccurrence")
}

func init() { proto.RegisterFile("grafeas/v1/discovery.proto", fileDescriptor_ca027ecd4454b7df) }

var fileDescriptor_ca027ecd4454b7df = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0x3b, 0x06, 0x7a, 0x5b, 0xbb, 0xc8, 0x2d, 0x5a, 0x55, 0x0e, 0x9b, 0xca, 0x65,
	0x27, 0x47, 0x2b, 0x17, 0x10, 0xa7, 0x90, 0xb8, 0xc3, 0x30, 0xb9, 0x21, 0x4e, 0x40, 0x70, 0x89,
	0x82, 0x1b, 0xa2, 0x88, 0xce, 0xae, 0xec, 0xb4, 0xd2, 0xee, 0x7c, 0x01, 0xbe, 0x02, 0x9f, 0x14,
	0xb5, 0x49, 0xb3, 0x76, 0x43, 0xda, 0xcd, 0xff, 0xf7, 0xfe, 0xff, 0xdf, 0xb3, 0x9f, 0x0c, 0xc3,
	0x5c, 0xa7, 0x3f, 0xb3, 0xd4, 0x38, 0xab, 0x4b, 0x67, 0x56, 0x18, 0xa1, 0x56, 0x99, 0xbe, 0xc5,
	0x0b, 0xad, 0x4a, 0x85, 0xa0, 0xee, 0xe1, 0xd5, 0xe5, 0xf0, 0x34, 0x57, 0x2a, 0x9f, 0x67, 0x8e,
	0x5e, 0x08, 0xc7, 0x94, 0x69, 0xb9, 0x34, 0x95, 0x69, 0x78, 0xba, 0x03, 0x10, 0xea, 0xe6, 0x46,
	0xc9, 0xaa, 0x31, 0xfa, 0x08, 0x1d, 0x7f, 0x0b, 0x64, 0xaa, 0xcc, 0xd0, 0x5b, 0xe8, 0xa4, 0x32,
	0x9d, 0xdf, 0x9a, 0xc2, 0x24, 0xbf, 0x0a, 0x39, 0x1b, 0x58, 0xe7, 0xd6, 0x45, 0x77, 0xdc, 0xc7,
	0x77, 0x63, 0xf0, 0xda, 0xf8, 0xa9, 0x90, 0xb3, 0xf0, 0x78, 0x6b, 0x5d, 0xab, 0xd1, 0xef, 0x03,
	0xe8, 0x35, 0xb0, 0xa9, 0x10, 0x4b, 0xad, 0x33, 0x29, 0x32, 0x24, 0xa0, 0x27, 0x94, 0x2c, 0x0b,
	0xb9, 0x54, 0x4b, 0x93, 0x6c, 0x23, 0x35, 0x78, 0xbc, 0x0b, 0xfe, 0x4f, 0x1a, 0x7b, 0x4d, 0xd4,
	0xad, 0x93, 0x21, 0x12, 0x0f, 0x6a, 0xe8, 0x2b, 0x9c, 0x34, 0xf7, 0xae, 0x9e, 0x3e, 0x68, 0x6d,
	0x06, 0xe0, 0xc7, 0x06, 0x6c, 0x11, 0x7c, 0x93, 0x0a, 0xbb, 0xe9, 0x9e, 0x46, 0x13, 0x78, 0x71,
	0x0f, 0x9c, 0x64, 0x5a, 0x2b, 0x3d, 0x68, 0x9f, 0x5b, 0x17, 0x47, 0x63, 0x84, 0xab, 0x9d, 0x63,
	0xbd, 0x10, 0xb8, 0x46, 0xf4, 0xf6, 0x11, 0x64, 0x6d, 0x1f, 0x71, 0x40, 0x0f, 0x9f, 0x82, 0x5e,
	0xc1, 0x99, 0x37, 0x65, 0x11, 0x65, 0xf1, 0x34, 0xe6, 0x89, 0xcb, 0xdc, 0xeb, 0x6f, 0x9c, 0xf2,
	0x24, 0x66, 0x3c, 0x20, 0x1e, 0x9d, 0x50, 0xe2, 0xdb, 0x4f, 0x10, 0xc0, 0xa1, 0xeb, 0x45, 0xf4,
	0x0b, 0xb1, 0x2d, 0x74, 0x0c, 0xcf, 0x29, 0xab, 0x55, 0x6b, 0xf4, 0xc7, 0x82, 0xee, 0xfe, 0xfd,
	0xd1, 0x19, 0xbc, 0x6c, 0x30, 0x3c, 0x72, 0xa3, 0xf8, 0x3e, 0xed, 0x08, 0x9e, 0x05, 0x84, 0xf9,
	0x94, 0x5d, 0x55, 0x38, 0xee, 0xb9, 0x8c, 0xad, 0x55, 0x0b, 0xf5, 0xc1, 0x9e, 0x50, 0x46, 0xf9,
	0x07, 0xe2, 0x27, 0x3c, 0xf6, 0x3c, 0xc2, 0xb9, 0xdd, 0x46, 0x3d, 0x38, 0x69, 0xaa, 0x13, 0x97,
	0x5e, 0x13, 0xdf, 0x3e, 0x40, 0x03, 0xe8, 0x37, 0xc5, 0x98, 0xf1, 0x38, 0x08, 0xa6, 0x61, 0x44,
	0x7c, 0xfb, 0xe9, 0xfb, 0xcf, 0xd0, 0x29, 0xd4, 0xce, 0xd2, 0x03, 0xeb, 0xfb, 0x9b, 0x7a, 0x47,
	0xb9, 0x9a, 0xa7, 0x32, 0xc7, 0x4a, 0xe7, 0x4e, 0x9e, 0xc9, 0xcd, 0x0f, 0x74, 0xaa, 0x56, 0xba,
	0x28, 0x8c, 0x73, 0xf7, 0x49, 0xdf, 0xd5, 0xc7, 0xbf, 0xad, 0xf6, 0x55, 0xe8, 0xfe, 0x38, 0xdc,
	0x58, 0x5f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x34, 0x10, 0xd3, 0x64, 0x08, 0x03, 0x00, 0x00,
}
