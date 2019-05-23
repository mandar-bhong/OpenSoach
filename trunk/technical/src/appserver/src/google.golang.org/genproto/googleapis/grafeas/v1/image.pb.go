// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grafeas/v1/image.proto

package grafeas

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

// Layer holds metadata specific to a layer of a Docker image.
type Layer struct {
	// Required. The recovered Dockerfile directive used to construct this layer.
	// See https://docs.docker.com/engine/reference/builder/ for more information.
	Directive string `protobuf:"bytes,1,opt,name=directive,proto3" json:"directive,omitempty"`
	// The recovered arguments to the Dockerfile directive.
	Arguments            string   `protobuf:"bytes,2,opt,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Layer) Reset()         { *m = Layer{} }
func (m *Layer) String() string { return proto.CompactTextString(m) }
func (*Layer) ProtoMessage()    {}
func (*Layer) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c8e3d6d73ed76c1, []int{0}
}

func (m *Layer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer.Unmarshal(m, b)
}
func (m *Layer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer.Marshal(b, m, deterministic)
}
func (m *Layer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer.Merge(m, src)
}
func (m *Layer) XXX_Size() int {
	return xxx_messageInfo_Layer.Size(m)
}
func (m *Layer) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer.DiscardUnknown(m)
}

var xxx_messageInfo_Layer proto.InternalMessageInfo

func (m *Layer) GetDirective() string {
	if m != nil {
		return m.Directive
	}
	return ""
}

func (m *Layer) GetArguments() string {
	if m != nil {
		return m.Arguments
	}
	return ""
}

// A set of properties that uniquely identify a given Docker image.
type Fingerprint struct {
	// Required. The layer ID of the final layer in the Docker image's v1
	// representation.
	V1Name string `protobuf:"bytes,1,opt,name=v1_name,json=v1Name,proto3" json:"v1_name,omitempty"`
	// Required. The ordered list of v2 blobs that represent a given image.
	V2Blob []string `protobuf:"bytes,2,rep,name=v2_blob,json=v2Blob,proto3" json:"v2_blob,omitempty"`
	// Output only. The name of the image's v2 blobs computed via:
	//   [bottom] := v2_blob[bottom]
	//   [N] := sha256(v2_blob[N] + " " + v2_name[N+1])
	// Only the name of the final blob is kept.
	V2Name               string   `protobuf:"bytes,3,opt,name=v2_name,json=v2Name,proto3" json:"v2_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fingerprint) Reset()         { *m = Fingerprint{} }
func (m *Fingerprint) String() string { return proto.CompactTextString(m) }
func (*Fingerprint) ProtoMessage()    {}
func (*Fingerprint) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c8e3d6d73ed76c1, []int{1}
}

func (m *Fingerprint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fingerprint.Unmarshal(m, b)
}
func (m *Fingerprint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fingerprint.Marshal(b, m, deterministic)
}
func (m *Fingerprint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fingerprint.Merge(m, src)
}
func (m *Fingerprint) XXX_Size() int {
	return xxx_messageInfo_Fingerprint.Size(m)
}
func (m *Fingerprint) XXX_DiscardUnknown() {
	xxx_messageInfo_Fingerprint.DiscardUnknown(m)
}

var xxx_messageInfo_Fingerprint proto.InternalMessageInfo

func (m *Fingerprint) GetV1Name() string {
	if m != nil {
		return m.V1Name
	}
	return ""
}

func (m *Fingerprint) GetV2Blob() []string {
	if m != nil {
		return m.V2Blob
	}
	return nil
}

func (m *Fingerprint) GetV2Name() string {
	if m != nil {
		return m.V2Name
	}
	return ""
}

// Basis describes the base image portion (Note) of the DockerImage
// relationship. Linked occurrences are derived from this or an equivalent image
// via:
//   FROM <Basis.resource_url>
// Or an equivalent reference, e.g., a tag of the resource_url.
type ImageNote struct {
	// Required. Immutable. The resource_url for the resource representing the
	// basis of associated occurrence images.
	ResourceUrl string `protobuf:"bytes,1,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// Required. Immutable. The fingerprint of the base image.
	Fingerprint          *Fingerprint `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ImageNote) Reset()         { *m = ImageNote{} }
func (m *ImageNote) String() string { return proto.CompactTextString(m) }
func (*ImageNote) ProtoMessage()    {}
func (*ImageNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c8e3d6d73ed76c1, []int{2}
}

func (m *ImageNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageNote.Unmarshal(m, b)
}
func (m *ImageNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageNote.Marshal(b, m, deterministic)
}
func (m *ImageNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageNote.Merge(m, src)
}
func (m *ImageNote) XXX_Size() int {
	return xxx_messageInfo_ImageNote.Size(m)
}
func (m *ImageNote) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageNote.DiscardUnknown(m)
}

var xxx_messageInfo_ImageNote proto.InternalMessageInfo

func (m *ImageNote) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *ImageNote) GetFingerprint() *Fingerprint {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

// Details of the derived image portion of the DockerImage relationship. This
// image would be produced from a Dockerfile with FROM <DockerImage.Basis in
// attached Note>.
type ImageOccurrence struct {
	// Required. The fingerprint of the derived image.
	Fingerprint *Fingerprint `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Output only. The number of layers by which this image differs from the
	// associated image basis.
	Distance int32 `protobuf:"varint,2,opt,name=distance,proto3" json:"distance,omitempty"`
	// This contains layer-specific metadata, if populated it has length
	// "distance" and is ordered with [distance] being the layer immediately
	// following the base image and [1] being the final layer.
	LayerInfo []*Layer `protobuf:"bytes,3,rep,name=layer_info,json=layerInfo,proto3" json:"layer_info,omitempty"`
	// Output only. This contains the base image URL for the derived image
	// occurrence.
	BaseResourceUrl      string   `protobuf:"bytes,4,opt,name=base_resource_url,json=baseResourceUrl,proto3" json:"base_resource_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageOccurrence) Reset()         { *m = ImageOccurrence{} }
func (m *ImageOccurrence) String() string { return proto.CompactTextString(m) }
func (*ImageOccurrence) ProtoMessage()    {}
func (*ImageOccurrence) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c8e3d6d73ed76c1, []int{3}
}

func (m *ImageOccurrence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageOccurrence.Unmarshal(m, b)
}
func (m *ImageOccurrence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageOccurrence.Marshal(b, m, deterministic)
}
func (m *ImageOccurrence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageOccurrence.Merge(m, src)
}
func (m *ImageOccurrence) XXX_Size() int {
	return xxx_messageInfo_ImageOccurrence.Size(m)
}
func (m *ImageOccurrence) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageOccurrence.DiscardUnknown(m)
}

var xxx_messageInfo_ImageOccurrence proto.InternalMessageInfo

func (m *ImageOccurrence) GetFingerprint() *Fingerprint {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

func (m *ImageOccurrence) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *ImageOccurrence) GetLayerInfo() []*Layer {
	if m != nil {
		return m.LayerInfo
	}
	return nil
}

func (m *ImageOccurrence) GetBaseResourceUrl() string {
	if m != nil {
		return m.BaseResourceUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Layer)(nil), "grafeas.v1.Layer")
	proto.RegisterType((*Fingerprint)(nil), "grafeas.v1.Fingerprint")
	proto.RegisterType((*ImageNote)(nil), "grafeas.v1.ImageNote")
	proto.RegisterType((*ImageOccurrence)(nil), "grafeas.v1.ImageOccurrence")
}

func init() { proto.RegisterFile("grafeas/v1/image.proto", fileDescriptor_4c8e3d6d73ed76c1) }

var fileDescriptor_4c8e3d6d73ed76c1 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6a, 0xa3, 0x50,
	0x14, 0xc6, 0x31, 0x4e, 0x32, 0xe3, 0x71, 0x86, 0x90, 0xbb, 0x98, 0xc8, 0x30, 0x8b, 0x8c, 0xab,
	0x30, 0x0b, 0xad, 0x76, 0xd3, 0xd2, 0x55, 0x53, 0x68, 0x09, 0x94, 0xb4, 0x15, 0x0a, 0xa5, 0x1b,
	0xb9, 0x9a, 0xe3, 0xe5, 0x82, 0xde, 0x1b, 0xae, 0x7f, 0xa0, 0xaf, 0xd3, 0x87, 0xe9, 0x73, 0x15,
	0x6f, 0x4c, 0xb4, 0x5d, 0x75, 0xa7, 0xdf, 0xef, 0x9c, 0xcf, 0xef, 0x1c, 0x0f, 0xfc, 0x66, 0x8a,
	0x66, 0x48, 0x4b, 0xbf, 0x09, 0x7c, 0x5e, 0x50, 0x86, 0xde, 0x4e, 0xc9, 0x4a, 0x12, 0xe8, 0x74,
	0xaf, 0x09, 0xdc, 0x2b, 0x18, 0xdf, 0xd2, 0x17, 0x54, 0xe4, 0x2f, 0x58, 0x5b, 0xae, 0x30, 0xad,
	0x78, 0x83, 0x8e, 0xb1, 0x30, 0x96, 0x56, 0xd4, 0x0b, 0x2d, 0xa5, 0x8a, 0xd5, 0x05, 0x8a, 0xaa,
	0x74, 0x46, 0x7b, 0x7a, 0x14, 0xdc, 0x27, 0xb0, 0xaf, 0xb9, 0x60, 0xa8, 0x76, 0x8a, 0x8b, 0x8a,
	0xcc, 0xe1, 0x7b, 0x13, 0xc4, 0x82, 0x16, 0x07, 0xa3, 0x49, 0x13, 0x6c, 0x68, 0x81, 0x1a, 0x84,
	0x71, 0x92, 0xcb, 0xc4, 0x19, 0x2d, 0x4c, 0x0d, 0xc2, 0x55, 0x2e, 0x93, 0x0e, 0xe8, 0x0e, 0xb3,
	0xeb, 0x08, 0xdb, 0x0e, 0x97, 0x83, 0xb5, 0x6e, 0x93, 0x6f, 0x64, 0x85, 0xe4, 0x1f, 0xfc, 0x54,
	0x58, 0xca, 0x5a, 0xa5, 0x18, 0xd7, 0x2a, 0xef, 0xcc, 0xed, 0x83, 0xf6, 0xa8, 0x72, 0x72, 0x0e,
	0x76, 0xd6, 0x27, 0xd1, 0x49, 0xed, 0x70, 0xee, 0xf5, 0x03, 0x7b, 0x83, 0xa0, 0xd1, 0xb0, 0xd6,
	0x7d, 0x33, 0x60, 0xaa, 0xbf, 0x75, 0x97, 0xa6, 0xb5, 0x52, 0x28, 0x52, 0xfc, 0x6c, 0x67, 0x7c,
	0xdd, 0x8e, 0xfc, 0x81, 0x1f, 0x5b, 0x5e, 0x56, 0x54, 0xa4, 0xa8, 0x63, 0x8c, 0xa3, 0xe3, 0x3b,
	0x39, 0x01, 0xc8, 0xdb, 0xa5, 0xc7, 0x5c, 0x64, 0xd2, 0x31, 0x17, 0xe6, 0xd2, 0x0e, 0x67, 0x43,
	0x57, 0xfd, 0x4b, 0x22, 0x4b, 0x17, 0xad, 0x45, 0x26, 0xc9, 0x7f, 0x98, 0x25, 0xb4, 0xc4, 0xf8,
	0xc3, 0xfc, 0xdf, 0xf4, 0xfc, 0xd3, 0x16, 0x44, 0xfd, 0x0e, 0x56, 0x0f, 0xf0, 0x8b, 0xcb, 0x81,
	0xdb, 0xbd, 0xf1, 0x7c, 0xc6, 0xa4, 0x64, 0x39, 0x7a, 0x4c, 0xe6, 0x54, 0x30, 0x4f, 0x2a, 0xe6,
	0x33, 0x14, 0xfa, 0x1a, 0xfc, 0x3d, 0xa2, 0x3b, 0x5e, 0xfa, 0xfd, 0xbd, 0x5c, 0x74, 0x8f, 0xaf,
	0x23, 0xf3, 0x26, 0xba, 0x4c, 0x26, 0xba, 0xf4, 0xf4, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x16, 0xc9,
	0xde, 0x0b, 0x52, 0x02, 0x00, 0x00,
}
