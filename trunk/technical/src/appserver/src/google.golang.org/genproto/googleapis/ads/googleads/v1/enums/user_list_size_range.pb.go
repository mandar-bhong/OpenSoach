// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/user_list_size_range.proto

package enums

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

// Enum containing possible user list size ranges.
type UserListSizeRangeEnum_UserListSizeRange int32

const (
	// Not specified.
	UserListSizeRangeEnum_UNSPECIFIED UserListSizeRangeEnum_UserListSizeRange = 0
	// Used for return value only. Represents value unknown in this version.
	UserListSizeRangeEnum_UNKNOWN UserListSizeRangeEnum_UserListSizeRange = 1
	// User list has less than 500 users.
	UserListSizeRangeEnum_LESS_THAN_FIVE_HUNDRED UserListSizeRangeEnum_UserListSizeRange = 2
	// User list has number of users in range of 500 to 1000.
	UserListSizeRangeEnum_LESS_THAN_ONE_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 3
	// User list has number of users in range of 1000 to 10000.
	UserListSizeRangeEnum_ONE_THOUSAND_TO_TEN_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 4
	// User list has number of users in range of 10000 to 50000.
	UserListSizeRangeEnum_TEN_THOUSAND_TO_FIFTY_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 5
	// User list has number of users in range of 50000 to 100000.
	UserListSizeRangeEnum_FIFTY_THOUSAND_TO_ONE_HUNDRED_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 6
	// User list has number of users in range of 100000 to 300000.
	UserListSizeRangeEnum_ONE_HUNDRED_THOUSAND_TO_THREE_HUNDRED_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 7
	// User list has number of users in range of 300000 to 500000.
	UserListSizeRangeEnum_THREE_HUNDRED_THOUSAND_TO_FIVE_HUNDRED_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 8
	// User list has number of users in range of 500000 to 1 million.
	UserListSizeRangeEnum_FIVE_HUNDRED_THOUSAND_TO_ONE_MILLION UserListSizeRangeEnum_UserListSizeRange = 9
	// User list has number of users in range of 1 to 2 millions.
	UserListSizeRangeEnum_ONE_MILLION_TO_TWO_MILLION UserListSizeRangeEnum_UserListSizeRange = 10
	// User list has number of users in range of 2 to 3 millions.
	UserListSizeRangeEnum_TWO_MILLION_TO_THREE_MILLION UserListSizeRangeEnum_UserListSizeRange = 11
	// User list has number of users in range of 3 to 5 millions.
	UserListSizeRangeEnum_THREE_MILLION_TO_FIVE_MILLION UserListSizeRangeEnum_UserListSizeRange = 12
	// User list has number of users in range of 5 to 10 millions.
	UserListSizeRangeEnum_FIVE_MILLION_TO_TEN_MILLION UserListSizeRangeEnum_UserListSizeRange = 13
	// User list has number of users in range of 10 to 20 millions.
	UserListSizeRangeEnum_TEN_MILLION_TO_TWENTY_MILLION UserListSizeRangeEnum_UserListSizeRange = 14
	// User list has number of users in range of 20 to 30 millions.
	UserListSizeRangeEnum_TWENTY_MILLION_TO_THIRTY_MILLION UserListSizeRangeEnum_UserListSizeRange = 15
	// User list has number of users in range of 30 to 50 millions.
	UserListSizeRangeEnum_THIRTY_MILLION_TO_FIFTY_MILLION UserListSizeRangeEnum_UserListSizeRange = 16
	// User list has over 50 million users.
	UserListSizeRangeEnum_OVER_FIFTY_MILLION UserListSizeRangeEnum_UserListSizeRange = 17
)

var UserListSizeRangeEnum_UserListSizeRange_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "LESS_THAN_FIVE_HUNDRED",
	3:  "LESS_THAN_ONE_THOUSAND",
	4:  "ONE_THOUSAND_TO_TEN_THOUSAND",
	5:  "TEN_THOUSAND_TO_FIFTY_THOUSAND",
	6:  "FIFTY_THOUSAND_TO_ONE_HUNDRED_THOUSAND",
	7:  "ONE_HUNDRED_THOUSAND_TO_THREE_HUNDRED_THOUSAND",
	8:  "THREE_HUNDRED_THOUSAND_TO_FIVE_HUNDRED_THOUSAND",
	9:  "FIVE_HUNDRED_THOUSAND_TO_ONE_MILLION",
	10: "ONE_MILLION_TO_TWO_MILLION",
	11: "TWO_MILLION_TO_THREE_MILLION",
	12: "THREE_MILLION_TO_FIVE_MILLION",
	13: "FIVE_MILLION_TO_TEN_MILLION",
	14: "TEN_MILLION_TO_TWENTY_MILLION",
	15: "TWENTY_MILLION_TO_THIRTY_MILLION",
	16: "THIRTY_MILLION_TO_FIFTY_MILLION",
	17: "OVER_FIFTY_MILLION",
}

var UserListSizeRangeEnum_UserListSizeRange_value = map[string]int32{
	"UNSPECIFIED":                                     0,
	"UNKNOWN":                                         1,
	"LESS_THAN_FIVE_HUNDRED":                          2,
	"LESS_THAN_ONE_THOUSAND":                          3,
	"ONE_THOUSAND_TO_TEN_THOUSAND":                    4,
	"TEN_THOUSAND_TO_FIFTY_THOUSAND":                  5,
	"FIFTY_THOUSAND_TO_ONE_HUNDRED_THOUSAND":          6,
	"ONE_HUNDRED_THOUSAND_TO_THREE_HUNDRED_THOUSAND":  7,
	"THREE_HUNDRED_THOUSAND_TO_FIVE_HUNDRED_THOUSAND": 8,
	"FIVE_HUNDRED_THOUSAND_TO_ONE_MILLION":            9,
	"ONE_MILLION_TO_TWO_MILLION":                      10,
	"TWO_MILLION_TO_THREE_MILLION":                    11,
	"THREE_MILLION_TO_FIVE_MILLION":                   12,
	"FIVE_MILLION_TO_TEN_MILLION":                     13,
	"TEN_MILLION_TO_TWENTY_MILLION":                   14,
	"TWENTY_MILLION_TO_THIRTY_MILLION":                15,
	"THIRTY_MILLION_TO_FIFTY_MILLION":                 16,
	"OVER_FIFTY_MILLION":                              17,
}

func (x UserListSizeRangeEnum_UserListSizeRange) String() string {
	return proto.EnumName(UserListSizeRangeEnum_UserListSizeRange_name, int32(x))
}

func (UserListSizeRangeEnum_UserListSizeRange) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_08c233c74e444694, []int{0, 0}
}

// Size range in terms of number of users of a UserList.
type UserListSizeRangeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListSizeRangeEnum) Reset()         { *m = UserListSizeRangeEnum{} }
func (m *UserListSizeRangeEnum) String() string { return proto.CompactTextString(m) }
func (*UserListSizeRangeEnum) ProtoMessage()    {}
func (*UserListSizeRangeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_08c233c74e444694, []int{0}
}

func (m *UserListSizeRangeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListSizeRangeEnum.Unmarshal(m, b)
}
func (m *UserListSizeRangeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListSizeRangeEnum.Marshal(b, m, deterministic)
}
func (m *UserListSizeRangeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListSizeRangeEnum.Merge(m, src)
}
func (m *UserListSizeRangeEnum) XXX_Size() int {
	return xxx_messageInfo_UserListSizeRangeEnum.Size(m)
}
func (m *UserListSizeRangeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListSizeRangeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListSizeRangeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.UserListSizeRangeEnum_UserListSizeRange", UserListSizeRangeEnum_UserListSizeRange_name, UserListSizeRangeEnum_UserListSizeRange_value)
	proto.RegisterType((*UserListSizeRangeEnum)(nil), "google.ads.googleads.v1.enums.UserListSizeRangeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/user_list_size_range.proto", fileDescriptor_08c233c74e444694)
}

var fileDescriptor_08c233c74e444694 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x49, 0x7a, 0x83, 0x13, 0xa0, 0xee, 0x48, 0x64, 0x11, 0x9a, 0xb6, 0x09, 0x15, 0x42,
	0x2c, 0xc6, 0x0a, 0xdd, 0x20, 0xb3, 0x72, 0xc8, 0xa4, 0xb1, 0x08, 0xe3, 0xc8, 0xb7, 0xaa, 0x28,
	0x92, 0x65, 0xb0, 0x65, 0x59, 0x4a, 0xec, 0xc8, 0xe3, 0x74, 0xd1, 0xe7, 0xe0, 0x09, 0x58, 0xf2,
	0x28, 0x3c, 0x0a, 0xbc, 0x01, 0x2b, 0x64, 0xbb, 0xe3, 0x0b, 0x35, 0x6c, 0xa2, 0xa3, 0xff, 0xff,
	0xce, 0xc9, 0xaf, 0xf1, 0x39, 0xf0, 0xd6, 0x8f, 0x22, 0x7f, 0xe5, 0x89, 0x8e, 0xcb, 0xc4, 0xbc,
	0x4c, 0xab, 0x9b, 0x91, 0xe8, 0x85, 0xdb, 0x35, 0x13, 0xb7, 0xcc, 0x8b, 0xed, 0x55, 0xc0, 0x12,
	0x9b, 0x05, 0xb7, 0x9e, 0x1d, 0x3b, 0xa1, 0xef, 0xe1, 0x4d, 0x1c, 0x25, 0x11, 0xea, 0xe7, 0x38,
	0x76, 0x5c, 0x86, 0x8b, 0x4e, 0x7c, 0x33, 0xc2, 0x59, 0x67, 0xef, 0x98, 0x0f, 0xde, 0x04, 0xa2,
	0x13, 0x86, 0x51, 0xe2, 0x24, 0x41, 0x14, 0xb2, 0xbc, 0x79, 0xf8, 0x75, 0x0f, 0x9e, 0x99, 0xcc,
	0x8b, 0xe7, 0x01, 0x4b, 0xf4, 0xe0, 0xd6, 0xd3, 0xd2, 0xc1, 0x24, 0xdc, 0xae, 0x87, 0xbf, 0x77,
	0xe1, 0xe8, 0x9e, 0x83, 0x0e, 0xa1, 0x63, 0x52, 0x7d, 0x41, 0xde, 0x2b, 0x53, 0x85, 0x4c, 0x84,
	0x07, 0xa8, 0x03, 0x07, 0x26, 0xfd, 0x40, 0xd5, 0x2b, 0x2a, 0xb4, 0x50, 0x0f, 0xba, 0x73, 0xa2,
	0xeb, 0xb6, 0x31, 0x93, 0xa9, 0x3d, 0x55, 0x2c, 0x62, 0xcf, 0x4c, 0x3a, 0xd1, 0xc8, 0x44, 0x68,
	0xd7, 0x3d, 0x95, 0x12, 0xdb, 0x98, 0xa9, 0xa6, 0x2e, 0xd3, 0x89, 0xb0, 0x83, 0xce, 0xe0, 0xb8,
	0xaa, 0xd8, 0x86, 0x6a, 0x1b, 0x84, 0x96, 0xc4, 0x2e, 0x1a, 0xc2, 0x49, 0x55, 0x49, 0x89, 0xa9,
	0x32, 0x35, 0xae, 0x4b, 0x66, 0x0f, 0xbd, 0x86, 0x97, 0x75, 0x2d, 0xa5, 0xd2, 0xb9, 0x77, 0x21,
	0x4a, 0x76, 0x1f, 0xbd, 0x01, 0xdc, 0xe4, 0x64, 0xff, 0x3c, 0xd3, 0x48, 0x43, 0xcf, 0x01, 0xba,
	0x00, 0xb1, 0xd9, 0xcb, 0xd3, 0x58, 0x0d, 0x4d, 0x0f, 0xd1, 0x2b, 0x38, 0x6f, 0xb4, 0x78, 0xb6,
	0x8f, 0xca, 0x7c, 0xae, 0xa8, 0x54, 0x78, 0x84, 0x4e, 0xa0, 0x57, 0x11, 0xb2, 0x24, 0x57, 0x6a,
	0xe1, 0x43, 0xfa, 0x48, 0x15, 0xa1, 0x4c, 0xca, 0x89, 0x0e, 0x1a, 0x40, 0xbf, 0x26, 0x15, 0xb9,
	0x38, 0xf2, 0x18, 0x9d, 0xc2, 0xf3, 0xaa, 0xc2, 0x5f, 0x9a, 0x03, 0x4f, 0xb2, 0x19, 0xa5, 0x90,
	0xa7, 0x20, 0xd4, 0xb8, 0x2e, 0x90, 0xa7, 0xe8, 0x1c, 0xce, 0xea, 0x5a, 0x9e, 0x45, 0xd1, 0x2a,
	0xd4, 0x21, 0x7a, 0x01, 0xa7, 0x75, 0xad, 0xfc, 0x66, 0x1c, 0x12, 0x50, 0x17, 0x90, 0x6a, 0x11,
	0xed, 0x2f, 0xfd, 0x68, 0xfc, 0xab, 0x05, 0x83, 0x2f, 0xd1, 0x1a, 0xff, 0x77, 0xb5, 0xc7, 0xdd,
	0x7b, 0xfb, 0xb9, 0x48, 0x97, 0x7a, 0xd1, 0xfa, 0x34, 0xbe, 0x6b, 0xf4, 0xa3, 0x95, 0x13, 0xfa,
	0x38, 0x8a, 0x7d, 0xd1, 0xf7, 0xc2, 0x6c, 0xe5, 0xf9, 0x75, 0x6d, 0x02, 0xf6, 0x8f, 0x63, 0x7b,
	0x97, 0xfd, 0x7e, 0x6b, 0xef, 0x5c, 0xca, 0xf2, 0xf7, 0x76, 0xff, 0x32, 0x1f, 0x25, 0xbb, 0x0c,
	0xe7, 0x65, 0x5a, 0x59, 0x23, 0x9c, 0x5e, 0x09, 0xfb, 0xc1, 0xfd, 0xa5, 0xec, 0xb2, 0x65, 0xe1,
	0x2f, 0xad, 0xd1, 0x32, 0xf3, 0x7f, 0xb6, 0x07, 0xb9, 0x28, 0x49, 0xb2, 0xcb, 0x24, 0xa9, 0x20,
	0x24, 0xc9, 0x1a, 0x49, 0x52, 0xc6, 0x7c, 0xde, 0xcf, 0x82, 0x5d, 0xfc, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x64, 0x67, 0xa0, 0x98, 0x04, 0x04, 0x00, 0x00,
}
