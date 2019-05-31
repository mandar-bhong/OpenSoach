// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/user_list.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A user list. This is a list of users a customer may target.
type UserList struct {
	// The resource name of the user list.
	// User list resource names have the form:
	//
	// `customers/{customer_id}/userLists/{user_list_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Id of the user list.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// A flag that indicates if a user may edit a list. Depends on the list
	// ownership and list type. For example, external remarketing user lists are
	// not editable.
	//
	// This field is read-only.
	ReadOnly *wrappers.BoolValue `protobuf:"bytes,3,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	// Name of this user list. Depending on its access_reason, the user list name
	// may not be unique (e.g. if access_reason=SHARED)
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of this user list.
	Description *wrappers.StringValue `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Membership status of this user list. Indicates whether a user list is open
	// or active. Only open user lists can accumulate more users and can be
	// targeted to.
	MembershipStatus enums.UserListMembershipStatusEnum_UserListMembershipStatus `protobuf:"varint,6,opt,name=membership_status,json=membershipStatus,proto3,enum=google.ads.googleads.v1.enums.UserListMembershipStatusEnum_UserListMembershipStatus" json:"membership_status,omitempty"`
	// An ID from external system. It is used by user list sellers to correlate
	// IDs on their systems.
	IntegrationCode *wrappers.StringValue `protobuf:"bytes,7,opt,name=integration_code,json=integrationCode,proto3" json:"integration_code,omitempty"`
	// Number of days a user's cookie stays on your list since its most recent
	// addition to the list. This field must be between 0 and 540 inclusive.
	// However, for CRM based userlists, this field can be set to 10000 which
	// means no expiration.
	//
	// It'll be ignored for logical_user_list.
	MembershipLifeSpan *wrappers.Int64Value `protobuf:"bytes,8,opt,name=membership_life_span,json=membershipLifeSpan,proto3" json:"membership_life_span,omitempty"`
	// Estimated number of users in this user list, on the Google Display Network.
	// This value is null if the number of users has not yet been determined.
	//
	// This field is read-only.
	SizeForDisplay *wrappers.Int64Value `protobuf:"bytes,9,opt,name=size_for_display,json=sizeForDisplay,proto3" json:"size_for_display,omitempty"`
	// Size range in terms of number of users of the UserList, on the Google
	// Display Network.
	//
	// This field is read-only.
	SizeRangeForDisplay enums.UserListSizeRangeEnum_UserListSizeRange `protobuf:"varint,10,opt,name=size_range_for_display,json=sizeRangeForDisplay,proto3,enum=google.ads.googleads.v1.enums.UserListSizeRangeEnum_UserListSizeRange" json:"size_range_for_display,omitempty"`
	// Estimated number of users in this user list in the google.com domain.
	// These are the users available for targeting in Search campaigns.
	// This value is null if the number of users has not yet been determined.
	//
	// This field is read-only.
	SizeForSearch *wrappers.Int64Value `protobuf:"bytes,11,opt,name=size_for_search,json=sizeForSearch,proto3" json:"size_for_search,omitempty"`
	// Size range in terms of number of users of the UserList, for Search ads.
	//
	// This field is read-only.
	SizeRangeForSearch enums.UserListSizeRangeEnum_UserListSizeRange `protobuf:"varint,12,opt,name=size_range_for_search,json=sizeRangeForSearch,proto3,enum=google.ads.googleads.v1.enums.UserListSizeRangeEnum_UserListSizeRange" json:"size_range_for_search,omitempty"`
	// Type of this list.
	//
	// This field is read-only.
	Type enums.UserListTypeEnum_UserListType `protobuf:"varint,13,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.UserListTypeEnum_UserListType" json:"type,omitempty"`
	// Indicating the reason why this user list membership status is closed. It is
	// only populated on lists that were automatically closed due to inactivity,
	// and will be cleared once the list membership status becomes open.
	ClosingReason enums.UserListClosingReasonEnum_UserListClosingReason `protobuf:"varint,14,opt,name=closing_reason,json=closingReason,proto3,enum=google.ads.googleads.v1.enums.UserListClosingReasonEnum_UserListClosingReason" json:"closing_reason,omitempty"`
	// Indicates the reason this account has been granted access to the list.
	// The reason can be SHARED, OWNED, LICENSED or SUBSCRIBED.
	//
	// This field is read-only.
	AccessReason enums.AccessReasonEnum_AccessReason `protobuf:"varint,15,opt,name=access_reason,json=accessReason,proto3,enum=google.ads.googleads.v1.enums.AccessReasonEnum_AccessReason" json:"access_reason,omitempty"`
	// Indicates if this share is still enabled. When a UserList is shared with
	// the user this field is set to ENABLED. Later the userList owner can decide
	// to revoke the share and make it DISABLED.
	// The default value of this field is set to ENABLED.
	AccountUserListStatus enums.UserListAccessStatusEnum_UserListAccessStatus `protobuf:"varint,16,opt,name=account_user_list_status,json=accountUserListStatus,proto3,enum=google.ads.googleads.v1.enums.UserListAccessStatusEnum_UserListAccessStatus" json:"account_user_list_status,omitempty"`
	// Indicates if this user list is eligible for Google Search Network.
	EligibleForSearch *wrappers.BoolValue `protobuf:"bytes,17,opt,name=eligible_for_search,json=eligibleForSearch,proto3" json:"eligible_for_search,omitempty"`
	// Indicates this user list is eligible for Google Display Network.
	//
	// This field is read-only.
	EligibleForDisplay *wrappers.BoolValue `protobuf:"bytes,18,opt,name=eligible_for_display,json=eligibleForDisplay,proto3" json:"eligible_for_display,omitempty"`
	// The user list.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to UserList:
	//	*UserList_CrmBasedUserList
	//	*UserList_SimilarUserList
	//	*UserList_RuleBasedUserList
	//	*UserList_LogicalUserList
	//	*UserList_BasicUserList
	UserList             isUserList_UserList `protobuf_oneof:"user_list"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcfde90717249098, []int{0}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *UserList) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UserList) GetReadOnly() *wrappers.BoolValue {
	if m != nil {
		return m.ReadOnly
	}
	return nil
}

func (m *UserList) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserList) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *UserList) GetMembershipStatus() enums.UserListMembershipStatusEnum_UserListMembershipStatus {
	if m != nil {
		return m.MembershipStatus
	}
	return enums.UserListMembershipStatusEnum_UNSPECIFIED
}

func (m *UserList) GetIntegrationCode() *wrappers.StringValue {
	if m != nil {
		return m.IntegrationCode
	}
	return nil
}

func (m *UserList) GetMembershipLifeSpan() *wrappers.Int64Value {
	if m != nil {
		return m.MembershipLifeSpan
	}
	return nil
}

func (m *UserList) GetSizeForDisplay() *wrappers.Int64Value {
	if m != nil {
		return m.SizeForDisplay
	}
	return nil
}

func (m *UserList) GetSizeRangeForDisplay() enums.UserListSizeRangeEnum_UserListSizeRange {
	if m != nil {
		return m.SizeRangeForDisplay
	}
	return enums.UserListSizeRangeEnum_UNSPECIFIED
}

func (m *UserList) GetSizeForSearch() *wrappers.Int64Value {
	if m != nil {
		return m.SizeForSearch
	}
	return nil
}

func (m *UserList) GetSizeRangeForSearch() enums.UserListSizeRangeEnum_UserListSizeRange {
	if m != nil {
		return m.SizeRangeForSearch
	}
	return enums.UserListSizeRangeEnum_UNSPECIFIED
}

func (m *UserList) GetType() enums.UserListTypeEnum_UserListType {
	if m != nil {
		return m.Type
	}
	return enums.UserListTypeEnum_UNSPECIFIED
}

func (m *UserList) GetClosingReason() enums.UserListClosingReasonEnum_UserListClosingReason {
	if m != nil {
		return m.ClosingReason
	}
	return enums.UserListClosingReasonEnum_UNSPECIFIED
}

func (m *UserList) GetAccessReason() enums.AccessReasonEnum_AccessReason {
	if m != nil {
		return m.AccessReason
	}
	return enums.AccessReasonEnum_UNSPECIFIED
}

func (m *UserList) GetAccountUserListStatus() enums.UserListAccessStatusEnum_UserListAccessStatus {
	if m != nil {
		return m.AccountUserListStatus
	}
	return enums.UserListAccessStatusEnum_UNSPECIFIED
}

func (m *UserList) GetEligibleForSearch() *wrappers.BoolValue {
	if m != nil {
		return m.EligibleForSearch
	}
	return nil
}

func (m *UserList) GetEligibleForDisplay() *wrappers.BoolValue {
	if m != nil {
		return m.EligibleForDisplay
	}
	return nil
}

type isUserList_UserList interface {
	isUserList_UserList()
}

type UserList_CrmBasedUserList struct {
	CrmBasedUserList *common.CrmBasedUserListInfo `protobuf:"bytes,19,opt,name=crm_based_user_list,json=crmBasedUserList,proto3,oneof"`
}

type UserList_SimilarUserList struct {
	SimilarUserList *common.SimilarUserListInfo `protobuf:"bytes,20,opt,name=similar_user_list,json=similarUserList,proto3,oneof"`
}

type UserList_RuleBasedUserList struct {
	RuleBasedUserList *common.RuleBasedUserListInfo `protobuf:"bytes,21,opt,name=rule_based_user_list,json=ruleBasedUserList,proto3,oneof"`
}

type UserList_LogicalUserList struct {
	LogicalUserList *common.LogicalUserListInfo `protobuf:"bytes,22,opt,name=logical_user_list,json=logicalUserList,proto3,oneof"`
}

type UserList_BasicUserList struct {
	BasicUserList *common.BasicUserListInfo `protobuf:"bytes,23,opt,name=basic_user_list,json=basicUserList,proto3,oneof"`
}

func (*UserList_CrmBasedUserList) isUserList_UserList() {}

func (*UserList_SimilarUserList) isUserList_UserList() {}

func (*UserList_RuleBasedUserList) isUserList_UserList() {}

func (*UserList_LogicalUserList) isUserList_UserList() {}

func (*UserList_BasicUserList) isUserList_UserList() {}

func (m *UserList) GetUserList() isUserList_UserList {
	if m != nil {
		return m.UserList
	}
	return nil
}

func (m *UserList) GetCrmBasedUserList() *common.CrmBasedUserListInfo {
	if x, ok := m.GetUserList().(*UserList_CrmBasedUserList); ok {
		return x.CrmBasedUserList
	}
	return nil
}

func (m *UserList) GetSimilarUserList() *common.SimilarUserListInfo {
	if x, ok := m.GetUserList().(*UserList_SimilarUserList); ok {
		return x.SimilarUserList
	}
	return nil
}

func (m *UserList) GetRuleBasedUserList() *common.RuleBasedUserListInfo {
	if x, ok := m.GetUserList().(*UserList_RuleBasedUserList); ok {
		return x.RuleBasedUserList
	}
	return nil
}

func (m *UserList) GetLogicalUserList() *common.LogicalUserListInfo {
	if x, ok := m.GetUserList().(*UserList_LogicalUserList); ok {
		return x.LogicalUserList
	}
	return nil
}

func (m *UserList) GetBasicUserList() *common.BasicUserListInfo {
	if x, ok := m.GetUserList().(*UserList_BasicUserList); ok {
		return x.BasicUserList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UserList) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UserList_CrmBasedUserList)(nil),
		(*UserList_SimilarUserList)(nil),
		(*UserList_RuleBasedUserList)(nil),
		(*UserList_LogicalUserList)(nil),
		(*UserList_BasicUserList)(nil),
	}
}

func init() {
	proto.RegisterType((*UserList)(nil), "google.ads.googleads.v1.resources.UserList")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/user_list.proto", fileDescriptor_fcfde90717249098)
}

var fileDescriptor_fcfde90717249098 = []byte{
	// 927 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x4d, 0x6f, 0xdb, 0x36,
	0x1c, 0xc6, 0x67, 0x37, 0xeb, 0x12, 0x26, 0x8e, 0x6d, 0x26, 0xe9, 0x84, 0xac, 0x18, 0xd2, 0x0d,
	0x05, 0x02, 0x0c, 0x90, 0xe6, 0xb4, 0x7b, 0x81, 0x5b, 0x6c, 0xb0, 0xb3, 0x36, 0x6b, 0xe1, 0x76,
	0x81, 0xdc, 0xe5, 0xb0, 0x05, 0x10, 0x68, 0x89, 0x56, 0x08, 0x50, 0xa4, 0x40, 0x4a, 0x1d, 0xdc,
	0x9d, 0x76, 0xd8, 0x17, 0xd9, 0x71, 0x1f, 0x65, 0x1f, 0x65, 0x1f, 0x60, 0x87, 0x9d, 0x06, 0x51,
	0xa2, 0x4c, 0xd9, 0x75, 0xa5, 0xc3, 0x6e, 0xf4, 0x9f, 0xff, 0xdf, 0xf3, 0xe8, 0xe1, 0x8b, 0x2c,
	0x30, 0x08, 0x39, 0x0f, 0x29, 0x76, 0x50, 0x20, 0x9d, 0x7c, 0x98, 0x8d, 0x5e, 0x0f, 0x1c, 0x81,
	0x25, 0x4f, 0x85, 0x8f, 0xa5, 0x93, 0x4a, 0x2c, 0x3c, 0x4a, 0x64, 0x62, 0xc7, 0x82, 0x27, 0x1c,
	0xde, 0xcb, 0xfb, 0x6c, 0x14, 0x48, 0xbb, 0x44, 0xec, 0xd7, 0x03, 0xbb, 0x44, 0x8e, 0x9d, 0x4d,
	0xaa, 0x3e, 0x8f, 0x22, 0xce, 0x96, 0x92, 0x32, 0xd7, 0x3c, 0xde, 0xf8, 0x18, 0x98, 0xa5, 0x91,
	0x74, 0x90, 0xef, 0x63, 0x29, 0x3d, 0x81, 0x91, 0xe4, 0xac, 0x40, 0x1e, 0xbd, 0x1b, 0x29, 0x2d,
	0xbc, 0x02, 0x96, 0x09, 0x4a, 0x52, 0xed, 0xf7, 0xb8, 0x29, 0xec, 0x53, 0x2e, 0x09, 0x0b, 0xab,
	0xd6, 0xdf, 0x36, 0xa5, 0x23, 0x1c, 0xcd, 0xb0, 0x90, 0x37, 0x24, 0xae, 0xda, 0x7f, 0xdd, 0x54,
	0x40, 0x92, 0x37, 0xd8, 0x13, 0x88, 0x85, 0xb8, 0x20, 0xcf, 0x9a, 0x92, 0xc9, 0x22, 0xd6, 0xcc,
	0xc7, 0x05, 0xa3, 0x7e, 0xcd, 0xd2, 0xb9, 0xf3, 0x8b, 0x40, 0x71, 0x8c, 0x85, 0x7e, 0x9a, 0xbb,
	0x5a, 0x33, 0x26, 0x0e, 0x62, 0x8c, 0x27, 0x28, 0x21, 0x9c, 0x15, 0xb3, 0x9f, 0xfc, 0xd3, 0x05,
	0xdb, 0x3f, 0x4a, 0x2c, 0x26, 0x44, 0x26, 0xf0, 0x53, 0xd0, 0xd1, 0xbb, 0xec, 0x31, 0x14, 0x61,
	0xab, 0x75, 0xd2, 0x3a, 0xdd, 0x71, 0xf7, 0x74, 0xf1, 0x25, 0x8a, 0x30, 0xfc, 0x0c, 0xb4, 0x49,
	0x60, 0xb5, 0x4f, 0x5a, 0xa7, 0xbb, 0x67, 0x1f, 0x15, 0x47, 0xc4, 0xd6, 0xe6, 0xf6, 0x33, 0x96,
	0x7c, 0xf9, 0xf0, 0x0a, 0xd1, 0x14, 0xbb, 0x6d, 0x12, 0xc0, 0xaf, 0xc0, 0x8e, 0xc0, 0x28, 0xf0,
	0x38, 0xa3, 0x0b, 0xeb, 0x96, 0x62, 0x8e, 0xd7, 0x98, 0x31, 0xe7, 0x34, 0x47, 0xb6, 0xb3, 0xe6,
	0x1f, 0x18, 0x5d, 0xc0, 0xcf, 0xc1, 0x96, 0x7a, 0x82, 0x2d, 0xc5, 0xdc, 0x5d, 0x63, 0xa6, 0x89,
	0x20, 0x2c, 0xcc, 0x29, 0xd5, 0x09, 0xbf, 0x01, 0xbb, 0x01, 0x96, 0xbe, 0x20, 0x71, 0x96, 0xcf,
	0x7a, 0xbf, 0x01, 0x68, 0x02, 0xf0, 0xb7, 0x16, 0xe8, 0xaf, 0xed, 0xa8, 0x75, 0xfb, 0xa4, 0x75,
	0xba, 0x7f, 0xf6, 0xca, 0xde, 0x74, 0x2b, 0xd4, 0xc6, 0xd8, 0x7a, 0x05, 0x5f, 0x94, 0xfc, 0x54,
	0xe1, 0x4f, 0x58, 0x1a, 0x6d, 0x9c, 0x74, 0x7b, 0xd1, 0x4a, 0x05, 0x5e, 0x80, 0x1e, 0x61, 0x09,
	0x0e, 0x85, 0xda, 0x23, 0xcf, 0xe7, 0x01, 0xb6, 0x3e, 0x68, 0x10, 0xa4, 0x6b, 0x50, 0xe7, 0x3c,
	0xc0, 0xf0, 0x05, 0x38, 0x34, 0xb2, 0x50, 0x32, 0xc7, 0x9e, 0x8c, 0x11, 0xb3, 0xb6, 0xeb, 0xb7,
	0x0d, 0x2e, 0xc1, 0x09, 0x99, 0xe3, 0x69, 0x8c, 0x18, 0x7c, 0x02, 0x7a, 0xea, 0xac, 0xce, 0xb9,
	0xf0, 0x02, 0x22, 0x63, 0x8a, 0x16, 0xd6, 0x4e, 0xbd, 0xd4, 0x7e, 0x06, 0x3d, 0xe5, 0xe2, 0xbb,
	0x1c, 0x81, 0xbf, 0x82, 0x3b, 0xcb, 0x23, 0x5f, 0x11, 0x03, 0x6a, 0x99, 0x9f, 0x36, 0x5c, 0xe6,
	0x29, 0x79, 0x83, 0xdd, 0x4c, 0xa3, 0xb2, 0xbe, 0x65, 0xd5, 0x3d, 0x90, 0x7a, 0x68, 0x98, 0x9f,
	0x83, 0x6e, 0x99, 0x41, 0x62, 0x24, 0xfc, 0x1b, 0x6b, 0xb7, 0x3e, 0x42, 0xa7, 0x88, 0x30, 0x55,
	0x04, 0x5c, 0x80, 0xa3, 0x95, 0x04, 0x85, 0xd4, 0xde, 0xff, 0x1a, 0x00, 0x9a, 0x01, 0x0a, 0xeb,
	0x4b, 0xb0, 0x95, 0xdd, 0x7a, 0xab, 0xa3, 0x9c, 0x1e, 0x37, 0x74, 0x7a, 0xb5, 0x88, 0xab, 0x26,
	0x59, 0xc1, 0x55, 0x4a, 0x30, 0x05, 0xfb, 0xd5, 0x17, 0xa0, 0xb5, 0xaf, 0xb4, 0x5f, 0x36, 0xd4,
	0x3e, 0xcf, 0x61, 0x57, 0xb1, 0x15, 0x93, 0xca, 0x8c, 0xdb, 0xf1, 0xcd, 0x9f, 0x10, 0x81, 0x4e,
	0xe5, 0x8d, 0x6f, 0x75, 0x1b, 0x25, 0x1a, 0x29, 0xc6, 0x30, 0x33, 0x0b, 0xee, 0x1e, 0x32, 0x7e,
	0xc1, 0xdf, 0x5b, 0xc0, 0x42, 0xbe, 0xcf, 0x53, 0x96, 0x78, 0xc6, 0xeb, 0x36, 0xbf, 0xd2, 0x3d,
	0x65, 0x37, 0x69, 0x18, 0x32, 0x77, 0x79, 0xcb, 0x75, 0x36, 0x27, 0xdc, 0xa3, 0xc2, 0xad, 0xdc,
	0xca, 0xfc, 0x3e, 0x3f, 0x07, 0x07, 0x98, 0x92, 0x90, 0xcc, 0x68, 0xe5, 0xb0, 0xf4, 0x6b, 0x5f,
	0x84, 0x7d, 0x8d, 0x2d, 0xf7, 0x7f, 0x02, 0x0e, 0x2b, 0x5a, 0xfa, 0xea, 0xc0, 0x5a, 0x31, 0x68,
	0x88, 0xe9, 0xdb, 0x80, 0xc1, 0x81, 0x2f, 0x22, 0x6f, 0x86, 0x24, 0x0e, 0x96, 0x4b, 0x64, 0x1d,
	0x28, 0xb1, 0x87, 0x1b, 0xd7, 0x26, 0xff, 0x87, 0xb7, 0xcf, 0x45, 0x34, 0xce, 0x48, 0x1d, 0xf7,
	0x19, 0x9b, 0xf3, 0xef, 0xdf, 0x73, 0x7b, 0xfe, 0x4a, 0x1d, 0x22, 0xd0, 0x97, 0x24, 0x22, 0x14,
	0x09, 0xc3, 0xe4, 0x50, 0x99, 0x3c, 0xa8, 0x33, 0x99, 0xe6, 0xe0, 0x8a, 0x47, 0x57, 0x56, 0xcb,
	0xf0, 0x06, 0x1c, 0x8a, 0x94, 0xe2, 0xb5, 0x28, 0x47, 0xca, 0xe5, 0x8b, 0x3a, 0x17, 0x37, 0xa5,
	0xf8, 0x6d, 0x59, 0xfa, 0x62, 0x75, 0x22, 0x0b, 0x43, 0x79, 0x48, 0x7c, 0x44, 0x0d, 0x9b, 0x3b,
	0xcd, 0xc2, 0x4c, 0x72, 0x70, 0x35, 0x0c, 0xad, 0x96, 0xe1, 0xcf, 0xa0, 0x3b, 0x43, 0x92, 0xf8,
	0x86, 0xc1, 0x87, 0xca, 0x60, 0x50, 0x67, 0x30, 0xce, 0xb0, 0x15, 0xf9, 0xce, 0xcc, 0x2c, 0x8e,
	0x77, 0xc1, 0x4e, 0x29, 0x3b, 0xfe, 0xb7, 0x05, 0xee, 0xfb, 0x3c, 0xb2, 0x6b, 0x3f, 0xf7, 0xc6,
	0x1d, 0x2d, 0x70, 0x99, 0x1d, 0xad, 0xcb, 0xd6, 0x4f, 0xcf, 0x0b, 0x26, 0xe4, 0x14, 0xb1, 0xd0,
	0xe6, 0x22, 0x74, 0x42, 0xcc, 0xd4, 0xc1, 0xd3, 0x5f, 0x2d, 0x31, 0x91, 0xef, 0xf8, 0xe8, 0x7c,
	0x54, 0x8e, 0xfe, 0x68, 0xdf, 0xba, 0x18, 0x8d, 0xfe, 0x6c, 0xdf, 0xbb, 0xc8, 0x25, 0x47, 0x81,
	0xb4, 0xf3, 0x61, 0x36, 0xba, 0x1a, 0xd8, 0xae, 0xee, 0xfc, 0x4b, 0xf7, 0x5c, 0x8f, 0x02, 0x79,
	0x5d, 0xf6, 0x5c, 0x5f, 0x0d, 0xae, 0xcb, 0x9e, 0xbf, 0xdb, 0xf7, 0xf3, 0x89, 0xe1, 0x70, 0x14,
	0xc8, 0xe1, 0xb0, 0xec, 0x1a, 0x0e, 0xaf, 0x06, 0xc3, 0x61, 0xd9, 0x37, 0xbb, 0xad, 0x1e, 0xf6,
	0xc1, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x06, 0xb6, 0xd5, 0x8a, 0x20, 0x0b, 0x00, 0x00,
}
