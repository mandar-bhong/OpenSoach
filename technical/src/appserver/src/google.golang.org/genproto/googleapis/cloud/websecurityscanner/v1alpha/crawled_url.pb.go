// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1alpha/crawled_url.proto

package websecurityscanner

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

// A CrawledUrl resource represents a URL that was crawled during a ScanRun. Web
// Security Scanner Service crawls the web applications, following all links
// within the scope of sites, to find the URLs to test against.
type CrawledUrl struct {
	// Output only.
	// The http method of the request that was used to visit the URL, in
	// uppercase.
	HttpMethod string `protobuf:"bytes,1,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	// Output only.
	// The URL that was crawled.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Output only.
	// The body of the request that was used to visit the URL.
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CrawledUrl) Reset()         { *m = CrawledUrl{} }
func (m *CrawledUrl) String() string { return proto.CompactTextString(m) }
func (*CrawledUrl) ProtoMessage()    {}
func (*CrawledUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_a45f0d639c4b8a44, []int{0}
}

func (m *CrawledUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrawledUrl.Unmarshal(m, b)
}
func (m *CrawledUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrawledUrl.Marshal(b, m, deterministic)
}
func (m *CrawledUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrawledUrl.Merge(m, src)
}
func (m *CrawledUrl) XXX_Size() int {
	return xxx_messageInfo_CrawledUrl.Size(m)
}
func (m *CrawledUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_CrawledUrl.DiscardUnknown(m)
}

var xxx_messageInfo_CrawledUrl proto.InternalMessageInfo

func (m *CrawledUrl) GetHttpMethod() string {
	if m != nil {
		return m.HttpMethod
	}
	return ""
}

func (m *CrawledUrl) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CrawledUrl) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*CrawledUrl)(nil), "google.cloud.websecurityscanner.v1alpha.CrawledUrl")
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1alpha/crawled_url.proto", fileDescriptor_a45f0d639c4b8a44)
}

var fileDescriptor_a45f0d639c4b8a44 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x59, 0x4f, 0x04, 0xc7, 0x42, 0x49, 0xb5, 0x88, 0xa0, 0xd8, 0x28, 0x08, 0x09, 0x62,
	0x25, 0x76, 0x67, 0x2d, 0x1c, 0x8a, 0x85, 0x36, 0xc7, 0x6c, 0x36, 0x64, 0x17, 0xe6, 0x32, 0x61,
	0x36, 0xeb, 0x71, 0xbf, 0xc5, 0x3f, 0x2b, 0x9b, 0x2c, 0x58, 0x5c, 0xa1, 0xdd, 0xf0, 0x1e, 0xdf,
	0xcb, 0x47, 0xe0, 0xd1, 0x33, 0x7b, 0x72, 0xc6, 0x12, 0x8f, 0xad, 0xd9, 0xba, 0x66, 0x70, 0x76,
	0x94, 0x3e, 0xed, 0x06, 0x8b, 0x21, 0x38, 0x31, 0x5f, 0xf7, 0x48, 0xb1, 0x43, 0x63, 0x05, 0xb7,
	0xe4, 0xda, 0xf5, 0x28, 0xa4, 0xa3, 0x70, 0x62, 0x75, 0x53, 0x50, 0x9d, 0x51, 0xbd, 0x8f, 0xea,
	0x19, 0x3d, 0xbf, 0x98, 0xdf, 0xc0, 0xd8, 0x1b, 0x0c, 0x81, 0x13, 0xa6, 0x9e, 0xc3, 0x50, 0x66,
	0xae, 0xdf, 0x00, 0x9e, 0xcb, 0xf6, 0xbb, 0x90, 0xba, 0x84, 0x93, 0x2e, 0xa5, 0xb8, 0xde, 0xb8,
	0xd4, 0x71, 0x5b, 0x57, 0x57, 0xd5, 0xed, 0xf1, 0x2b, 0x4c, 0xd1, 0x4b, 0x4e, 0xd4, 0x19, 0x2c,
	0x46, 0xa1, 0xfa, 0x20, 0x17, 0xd3, 0xa9, 0x14, 0x1c, 0x36, 0xdc, 0xee, 0xea, 0x45, 0x8e, 0xf2,
	0xbd, 0xfc, 0xae, 0xe0, 0xce, 0xf2, 0x46, 0xff, 0x53, 0x71, 0x79, 0xfa, 0xab, 0xb0, 0x9a, 0xac,
	0x56, 0xd5, 0xe7, 0xc7, 0xcc, 0x7a, 0x26, 0x0c, 0x5e, 0xb3, 0x78, 0xe3, 0x5d, 0xc8, 0xce, 0xa6,
	0x54, 0x18, 0xfb, 0xe1, 0xcf, 0x8f, 0x7b, 0xda, 0xaf, 0x9a, 0xa3, 0xbc, 0xf2, 0xf0, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x7b, 0x55, 0xd9, 0x73, 0x7d, 0x01, 0x00, 0x00,
}
