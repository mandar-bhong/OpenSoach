// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/uptime.proto

package monitoring

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"
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

// The regions from which an uptime check can be run.
type UptimeCheckRegion int32

const (
	// Default value if no region is specified. Will result in uptime checks
	// running from all regions.
	UptimeCheckRegion_REGION_UNSPECIFIED UptimeCheckRegion = 0
	// Allows checks to run from locations within the United States of America.
	UptimeCheckRegion_USA UptimeCheckRegion = 1
	// Allows checks to run from locations within the continent of Europe.
	UptimeCheckRegion_EUROPE UptimeCheckRegion = 2
	// Allows checks to run from locations within the continent of South
	// America.
	UptimeCheckRegion_SOUTH_AMERICA UptimeCheckRegion = 3
	// Allows checks to run from locations within the Asia Pacific area (ex:
	// Singapore).
	UptimeCheckRegion_ASIA_PACIFIC UptimeCheckRegion = 4
)

var UptimeCheckRegion_name = map[int32]string{
	0: "REGION_UNSPECIFIED",
	1: "USA",
	2: "EUROPE",
	3: "SOUTH_AMERICA",
	4: "ASIA_PACIFIC",
}

var UptimeCheckRegion_value = map[string]int32{
	"REGION_UNSPECIFIED": 0,
	"USA":                1,
	"EUROPE":             2,
	"SOUTH_AMERICA":      3,
	"ASIA_PACIFIC":       4,
}

func (x UptimeCheckRegion) String() string {
	return proto.EnumName(UptimeCheckRegion_name, int32(x))
}

func (UptimeCheckRegion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7ca0e36dfc8221d8, []int{0}
}

// The supported resource types that can be used as values of
// `group_resource.resource_type`.
// `INSTANCE` includes `gce_instance` and `aws_ec2_instance` resource types.
// The resource types `gae_app` and `uptime_url` are not valid here because
// group checks on App Engine modules and URLs are not allowed.
type GroupResourceType int32

const (
	// Default value (not valid).
	GroupResourceType_RESOURCE_TYPE_UNSPECIFIED GroupResourceType = 0
	// A group of instances from Google Cloud Platform (GCP) or
	// Amazon Web Services (AWS).
	GroupResourceType_INSTANCE GroupResourceType = 1
	// A group of Amazon ELB load balancers.
	GroupResourceType_AWS_ELB_LOAD_BALANCER GroupResourceType = 2
)

var GroupResourceType_name = map[int32]string{
	0: "RESOURCE_TYPE_UNSPECIFIED",
	1: "INSTANCE",
	2: "AWS_ELB_LOAD_BALANCER",
}

var GroupResourceType_value = map[string]int32{
	"RESOURCE_TYPE_UNSPECIFIED": 0,
	"INSTANCE":                  1,
	"AWS_ELB_LOAD_BALANCER":     2,
}

func (x GroupResourceType) String() string {
	return proto.EnumName(GroupResourceType_name, int32(x))
}

func (GroupResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7ca0e36dfc8221d8, []int{1}
}

// An internal checker allows uptime checks to run on private/internal GCP
// resources.
type InternalChecker struct {
	// A unique resource name for this InternalChecker. The format is:
	//
	//
	//   `projects/[PROJECT_ID]/internalCheckers/[INTERNAL_CHECKER_ID]`.
	//
	// PROJECT_ID is the stackdriver workspace project for the
	// uptime check config associated with the internal checker.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The checker's human-readable name. The display name
	// should be unique within a Stackdriver Workspace in order to make it easier
	// to identify; however, uniqueness is not enforced.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The [GCP VPC network](https://cloud.google.com/vpc/docs/vpc) where the
	// internal resource lives (ex: "default").
	Network string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	// The GCP zone the uptime check should egress from. Only respected for
	// internal uptime checks, where internal_network is specified.
	GcpZone string `protobuf:"bytes,4,opt,name=gcp_zone,json=gcpZone,proto3" json:"gcp_zone,omitempty"`
	// The GCP project_id where the internal checker lives. Not necessary
	// the same as the workspace project.
	PeerProjectId        string   `protobuf:"bytes,6,opt,name=peer_project_id,json=peerProjectId,proto3" json:"peer_project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalChecker) Reset()         { *m = InternalChecker{} }
func (m *InternalChecker) String() string { return proto.CompactTextString(m) }
func (*InternalChecker) ProtoMessage()    {}
func (*InternalChecker) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca0e36dfc8221d8, []int{0}
}

func (m *InternalChecker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalChecker.Unmarshal(m, b)
}
func (m *InternalChecker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalChecker.Marshal(b, m, deterministic)
}
func (m *InternalChecker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalChecker.Merge(m, src)
}
func (m *InternalChecker) XXX_Size() int {
	return xxx_messageInfo_InternalChecker.Size(m)
}
func (m *InternalChecker) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalChecker.DiscardUnknown(m)
}

var xxx_messageInfo_InternalChecker proto.InternalMessageInfo

func (m *InternalChecker) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InternalChecker) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *InternalChecker) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *InternalChecker) GetGcpZone() string {
	if m != nil {
		return m.GcpZone
	}
	return ""
}

func (m *InternalChecker) GetPeerProjectId() string {
	if m != nil {
		return m.PeerProjectId
	}
	return ""
}

// This message configures which resources and services to monitor for
// availability.
type UptimeCheckConfig struct {
	// A unique resource name for this UptimeCheckConfig. The format is:
	//
	//
	//   `projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID]`.
	//
	// This field should be omitted when creating the uptime check configuration;
	// on create, the resource name is assigned by the server and included in the
	// response.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A human-friendly name for the uptime check configuration. The display name
	// should be unique within a Stackdriver Workspace in order to make it easier
	// to identify; however, uniqueness is not enforced. Required.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The resource the check is checking. Required.
	//
	// Types that are valid to be assigned to Resource:
	//	*UptimeCheckConfig_MonitoredResource
	//	*UptimeCheckConfig_ResourceGroup_
	Resource isUptimeCheckConfig_Resource `protobuf_oneof:"resource"`
	// The type of uptime check request.
	//
	// Types that are valid to be assigned to CheckRequestType:
	//	*UptimeCheckConfig_HttpCheck_
	//	*UptimeCheckConfig_TcpCheck_
	CheckRequestType isUptimeCheckConfig_CheckRequestType `protobuf_oneof:"check_request_type"`
	// How often, in seconds, the uptime check is performed.
	// Currently, the only supported values are `60s` (1 minute), `300s`
	// (5 minutes), `600s` (10 minutes), and `900s` (15 minutes). Optional,
	// defaults to `300s`.
	Period *duration.Duration `protobuf:"bytes,7,opt,name=period,proto3" json:"period,omitempty"`
	// The maximum amount of time to wait for the request to complete (must be
	// between 1 and 60 seconds). Required.
	Timeout *duration.Duration `protobuf:"bytes,8,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The expected content on the page the check is run against.
	// Currently, only the first entry in the list is supported, and other entries
	// will be ignored. The server will look for an exact match of the string in
	// the page response's content. This field is optional and should only be
	// specified if a content match is required.
	ContentMatchers []*UptimeCheckConfig_ContentMatcher `protobuf:"bytes,9,rep,name=content_matchers,json=contentMatchers,proto3" json:"content_matchers,omitempty"`
	// The list of regions from which the check will be run.
	// Some regions contain one location, and others contain more than one.
	// If this field is specified, enough regions to include a minimum of
	// 3 locations must be provided, or an error message is returned.
	// Not specifying this field will result in uptime checks running from all
	// regions.
	SelectedRegions []UptimeCheckRegion `protobuf:"varint,10,rep,packed,name=selected_regions,json=selectedRegions,proto3,enum=google.monitoring.v3.UptimeCheckRegion" json:"selected_regions,omitempty"`
	// If this is true, then checks are made only from the 'internal_checkers'.
	// If it is false, then checks are made only from the 'selected_regions'.
	// It is an error to provide 'selected_regions' when is_internal is true,
	// or to provide 'internal_checkers' when is_internal is false.
	IsInternal bool `protobuf:"varint,15,opt,name=is_internal,json=isInternal,proto3" json:"is_internal,omitempty"`
	// The internal checkers that this check will egress from. If `is_internal` is
	// true and this list is empty, the check will egress from all the
	// InternalCheckers configured for the project that owns this CheckConfig.
	InternalCheckers     []*InternalChecker `protobuf:"bytes,14,rep,name=internal_checkers,json=internalCheckers,proto3" json:"internal_checkers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UptimeCheckConfig) Reset()         { *m = UptimeCheckConfig{} }
func (m *UptimeCheckConfig) String() string { return proto.CompactTextString(m) }
func (*UptimeCheckConfig) ProtoMessage()    {}
func (*UptimeCheckConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca0e36dfc8221d8, []int{1}
}

func (m *UptimeCheckConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UptimeCheckConfig.Unmarshal(m, b)
}
func (m *UptimeCheckConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UptimeCheckConfig.Marshal(b, m, deterministic)
}
func (m *UptimeCheckConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeCheckConfig.Merge(m, src)
}
func (m *UptimeCheckConfig) XXX_Size() int {
	return xxx_messageInfo_UptimeCheckConfig.Size(m)
}
func (m *UptimeCheckConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeCheckConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeCheckConfig proto.InternalMessageInfo

func (m *UptimeCheckConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UptimeCheckConfig) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type isUptimeCheckConfig_Resource interface {
	isUptimeCheckConfig_Resource()
}

type UptimeCheckConfig_MonitoredResource struct {
	MonitoredResource *monitoredres.MonitoredResource `protobuf:"bytes,3,opt,name=monitored_resource,json=monitoredResource,proto3,oneof"`
}

type UptimeCheckConfig_ResourceGroup_ struct {
	ResourceGroup *UptimeCheckConfig_ResourceGroup `protobuf:"bytes,4,opt,name=resource_group,json=resourceGroup,proto3,oneof"`
}

func (*UptimeCheckConfig_MonitoredResource) isUptimeCheckConfig_Resource() {}

func (*UptimeCheckConfig_ResourceGroup_) isUptimeCheckConfig_Resource() {}

func (m *UptimeCheckConfig) GetResource() isUptimeCheckConfig_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *UptimeCheckConfig) GetMonitoredResource() *monitoredres.MonitoredResource {
	if x, ok := m.GetResource().(*UptimeCheckConfig_MonitoredResource); ok {
		return x.MonitoredResource
	}
	return nil
}

func (m *UptimeCheckConfig) GetResourceGroup() *UptimeCheckConfig_ResourceGroup {
	if x, ok := m.GetResource().(*UptimeCheckConfig_ResourceGroup_); ok {
		return x.ResourceGroup
	}
	return nil
}

type isUptimeCheckConfig_CheckRequestType interface {
	isUptimeCheckConfig_CheckRequestType()
}

type UptimeCheckConfig_HttpCheck_ struct {
	HttpCheck *UptimeCheckConfig_HttpCheck `protobuf:"bytes,5,opt,name=http_check,json=httpCheck,proto3,oneof"`
}

type UptimeCheckConfig_TcpCheck_ struct {
	TcpCheck *UptimeCheckConfig_TcpCheck `protobuf:"bytes,6,opt,name=tcp_check,json=tcpCheck,proto3,oneof"`
}

func (*UptimeCheckConfig_HttpCheck_) isUptimeCheckConfig_CheckRequestType() {}

func (*UptimeCheckConfig_TcpCheck_) isUptimeCheckConfig_CheckRequestType() {}

func (m *UptimeCheckConfig) GetCheckRequestType() isUptimeCheckConfig_CheckRequestType {
	if m != nil {
		return m.CheckRequestType
	}
	return nil
}

func (m *UptimeCheckConfig) GetHttpCheck() *UptimeCheckConfig_HttpCheck {
	if x, ok := m.GetCheckRequestType().(*UptimeCheckConfig_HttpCheck_); ok {
		return x.HttpCheck
	}
	return nil
}

func (m *UptimeCheckConfig) GetTcpCheck() *UptimeCheckConfig_TcpCheck {
	if x, ok := m.GetCheckRequestType().(*UptimeCheckConfig_TcpCheck_); ok {
		return x.TcpCheck
	}
	return nil
}

func (m *UptimeCheckConfig) GetPeriod() *duration.Duration {
	if m != nil {
		return m.Period
	}
	return nil
}

func (m *UptimeCheckConfig) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *UptimeCheckConfig) GetContentMatchers() []*UptimeCheckConfig_ContentMatcher {
	if m != nil {
		return m.ContentMatchers
	}
	return nil
}

func (m *UptimeCheckConfig) GetSelectedRegions() []UptimeCheckRegion {
	if m != nil {
		return m.SelectedRegions
	}
	return nil
}

func (m *UptimeCheckConfig) GetIsInternal() bool {
	if m != nil {
		return m.IsInternal
	}
	return false
}

func (m *UptimeCheckConfig) GetInternalCheckers() []*InternalChecker {
	if m != nil {
		return m.InternalCheckers
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UptimeCheckConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UptimeCheckConfig_MonitoredResource)(nil),
		(*UptimeCheckConfig_ResourceGroup_)(nil),
		(*UptimeCheckConfig_HttpCheck_)(nil),
		(*UptimeCheckConfig_TcpCheck_)(nil),
	}
}

// The resource submessage for group checks. It can be used instead of a
// monitored resource, when multiple resources are being monitored.
type UptimeCheckConfig_ResourceGroup struct {
	// The group of resources being monitored. Should be only the
	// group_id, not projects/<project_id>/groups/<group_id>.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// The resource type of the group members.
	ResourceType         GroupResourceType `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3,enum=google.monitoring.v3.GroupResourceType" json:"resource_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UptimeCheckConfig_ResourceGroup) Reset()         { *m = UptimeCheckConfig_ResourceGroup{} }
func (m *UptimeCheckConfig_ResourceGroup) String() string { return proto.CompactTextString(m) }
func (*UptimeCheckConfig_ResourceGroup) ProtoMessage()    {}
func (*UptimeCheckConfig_ResourceGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca0e36dfc8221d8, []int{1, 0}
}

func (m *UptimeCheckConfig_ResourceGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UptimeCheckConfig_ResourceGroup.Unmarshal(m, b)
}
func (m *UptimeCheckConfig_ResourceGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UptimeCheckConfig_ResourceGroup.Marshal(b, m, deterministic)
}
func (m *UptimeCheckConfig_ResourceGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeCheckConfig_ResourceGroup.Merge(m, src)
}
func (m *UptimeCheckConfig_ResourceGroup) XXX_Size() int {
	return xxx_messageInfo_UptimeCheckConfig_ResourceGroup.Size(m)
}
func (m *UptimeCheckConfig_ResourceGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeCheckConfig_ResourceGroup.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeCheckConfig_ResourceGroup proto.InternalMessageInfo

func (m *UptimeCheckConfig_ResourceGroup) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *UptimeCheckConfig_ResourceGroup) GetResourceType() GroupResourceType {
	if m != nil {
		return m.ResourceType
	}
	return GroupResourceType_RESOURCE_TYPE_UNSPECIFIED
}

// Information involved in an HTTP/HTTPS uptime check request.
type UptimeCheckConfig_HttpCheck struct {
	// If true, use HTTPS instead of HTTP to run the check.
	UseSsl bool `protobuf:"varint,1,opt,name=use_ssl,json=useSsl,proto3" json:"use_ssl,omitempty"`
	// The path to the page to run the check against. Will be combined with the
	// host (specified within the MonitoredResource) and port to construct the
	// full URL. Optional (defaults to "/").
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// The port to the page to run the check against. Will be combined with host
	// (specified within the MonitoredResource) and path to construct the full
	// URL. Optional (defaults to 80 without SSL, or 443 with SSL).
	Port int32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// The authentication information. Optional when creating an HTTP check;
	// defaults to empty.
	AuthInfo *UptimeCheckConfig_HttpCheck_BasicAuthentication `protobuf:"bytes,4,opt,name=auth_info,json=authInfo,proto3" json:"auth_info,omitempty"`
	// Boolean specifiying whether to encrypt the header information.
	// Encryption should be specified for any headers related to authentication
	// that you do not wish to be seen when retrieving the configuration. The
	// server will be responsible for encrypting the headers.
	// On Get/List calls, if mask_headers is set to True then the headers
	// will be obscured with ******.
	MaskHeaders bool `protobuf:"varint,5,opt,name=mask_headers,json=maskHeaders,proto3" json:"mask_headers,omitempty"`
	// The list of headers to send as part of the uptime check request.
	// If two headers have the same key and different values, they should
	// be entered as a single header, with the value being a comma-separated
	// list of all the desired values as described at
	// https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31).
	// Entering two separate headers with the same key in a Create call will
	// cause the first to be overwritten by the second.
	// The maximum number of headers allowed is 100.
	Headers              map[string]string `protobuf:"bytes,6,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UptimeCheckConfig_HttpCheck) Reset()         { *m = UptimeCheckConfig_HttpCheck{} }
func (m *UptimeCheckConfig_HttpCheck) String() string { return proto.CompactTextString(m) }
func (*UptimeCheckConfig_HttpCheck) ProtoMessage()    {}
func (*UptimeCheckConfig_HttpCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca0e36dfc8221d8, []int{1, 1}
}

func (m *UptimeCheckConfig_HttpCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UptimeCheckConfig_HttpCheck.Unmarshal(m, b)
}
func (m *UptimeCheckConfig_HttpCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UptimeCheckConfig_HttpCheck.Marshal(b, m, deterministic)
}
func (m *UptimeCheckConfig_HttpCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeCheckConfig_HttpCheck.Merge(m, src)
}
func (m *UptimeCheckConfig_HttpCheck) XXX_Size() int {
	return xxx_messageInfo_UptimeCheckConfig_HttpCheck.Size(m)
}
func (m *UptimeCheckConfig_HttpCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeCheckConfig_HttpCheck.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeCheckConfig_HttpCheck proto.InternalMessageInfo

func (m *UptimeCheckConfig_HttpCheck) GetUseSsl() bool {
	if m != nil {
		return m.UseSsl
	}
	return false
}

func (m *UptimeCheckConfig_HttpCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *UptimeCheckConfig_HttpCheck) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *UptimeCheckConfig_HttpCheck) GetAuthInfo() *UptimeCheckConfig_HttpCheck_BasicAuthentication {
	if m != nil {
		return m.AuthInfo
	}
	return nil
}

func (m *UptimeCheckConfig_HttpCheck) GetMaskHeaders() bool {
	if m != nil {
		return m.MaskHeaders
	}
	return false
}

func (m *UptimeCheckConfig_HttpCheck) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

// A type of authentication to perform against the specified resource or URL
// that uses username and password.
// Currently, only Basic authentication is supported in Uptime Monitoring.
type UptimeCheckConfig_HttpCheck_BasicAuthentication struct {
	// The username to authenticate.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// The password to authenticate.
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UptimeCheckConfig_HttpCheck_BasicAuthentication) Reset() {
	*m = UptimeCheckConfig_HttpCheck_BasicAuthentication{}
}
func (m *UptimeCheckConfig_HttpCheck_BasicAuthentication) String() string {
	return proto.CompactTextString(m)
}
func (*UptimeCheckConfig_HttpCheck_BasicAuthentication) ProtoMessage() {}
func (*UptimeCheckConfig_HttpCheck_BasicAuthentication) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca0e36dfc8221d8, []int{1, 1, 0}
}

func (m *UptimeCheckConfig_HttpCheck_BasicAuthentication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UptimeCheckConfig_HttpCheck_BasicAuthentication.Unmarshal(m, b)
}
func (m *UptimeCheckConfig_HttpCheck_BasicAuthentication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UptimeCheckConfig_HttpCheck_BasicAuthentication.Marshal(b, m, deterministic)
}
func (m *UptimeCheckConfig_HttpCheck_BasicAuthentication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeCheckConfig_HttpCheck_BasicAuthentication.Merge(m, src)
}
func (m *UptimeCheckConfig_HttpCheck_BasicAuthentication) XXX_Size() int {
	return xxx_messageInfo_UptimeCheckConfig_HttpCheck_BasicAuthentication.Size(m)
}
func (m *UptimeCheckConfig_HttpCheck_BasicAuthentication) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeCheckConfig_HttpCheck_BasicAuthentication.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeCheckConfig_HttpCheck_BasicAuthentication proto.InternalMessageInfo

func (m *UptimeCheckConfig_HttpCheck_BasicAuthentication) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UptimeCheckConfig_HttpCheck_BasicAuthentication) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// Information required for a TCP uptime check request.
type UptimeCheckConfig_TcpCheck struct {
	// The port to the page to run the check against. Will be combined with host
	// (specified within the MonitoredResource) to construct the full URL.
	// Required.
	Port                 int32    `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UptimeCheckConfig_TcpCheck) Reset()         { *m = UptimeCheckConfig_TcpCheck{} }
func (m *UptimeCheckConfig_TcpCheck) String() string { return proto.CompactTextString(m) }
func (*UptimeCheckConfig_TcpCheck) ProtoMessage()    {}
func (*UptimeCheckConfig_TcpCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca0e36dfc8221d8, []int{1, 2}
}

func (m *UptimeCheckConfig_TcpCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UptimeCheckConfig_TcpCheck.Unmarshal(m, b)
}
func (m *UptimeCheckConfig_TcpCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UptimeCheckConfig_TcpCheck.Marshal(b, m, deterministic)
}
func (m *UptimeCheckConfig_TcpCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeCheckConfig_TcpCheck.Merge(m, src)
}
func (m *UptimeCheckConfig_TcpCheck) XXX_Size() int {
	return xxx_messageInfo_UptimeCheckConfig_TcpCheck.Size(m)
}
func (m *UptimeCheckConfig_TcpCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeCheckConfig_TcpCheck.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeCheckConfig_TcpCheck proto.InternalMessageInfo

func (m *UptimeCheckConfig_TcpCheck) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// Used to perform string matching. It allows substring and regular
// expressions, together with their negations.
type UptimeCheckConfig_ContentMatcher struct {
	// String or regex content to match (max 1024 bytes)
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UptimeCheckConfig_ContentMatcher) Reset()         { *m = UptimeCheckConfig_ContentMatcher{} }
func (m *UptimeCheckConfig_ContentMatcher) String() string { return proto.CompactTextString(m) }
func (*UptimeCheckConfig_ContentMatcher) ProtoMessage()    {}
func (*UptimeCheckConfig_ContentMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca0e36dfc8221d8, []int{1, 3}
}

func (m *UptimeCheckConfig_ContentMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UptimeCheckConfig_ContentMatcher.Unmarshal(m, b)
}
func (m *UptimeCheckConfig_ContentMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UptimeCheckConfig_ContentMatcher.Marshal(b, m, deterministic)
}
func (m *UptimeCheckConfig_ContentMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeCheckConfig_ContentMatcher.Merge(m, src)
}
func (m *UptimeCheckConfig_ContentMatcher) XXX_Size() int {
	return xxx_messageInfo_UptimeCheckConfig_ContentMatcher.Size(m)
}
func (m *UptimeCheckConfig_ContentMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeCheckConfig_ContentMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeCheckConfig_ContentMatcher proto.InternalMessageInfo

func (m *UptimeCheckConfig_ContentMatcher) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// Contains the region, location, and list of IP
// addresses where checkers in the location run from.
type UptimeCheckIp struct {
	// A broad region category in which the IP address is located.
	Region UptimeCheckRegion `protobuf:"varint,1,opt,name=region,proto3,enum=google.monitoring.v3.UptimeCheckRegion" json:"region,omitempty"`
	// A more specific location within the region that typically encodes
	// a particular city/town/metro (and its containing state/province or country)
	// within the broader umbrella region category.
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// The IP address from which the uptime check originates. This is a full
	// IP address (not an IP address range). Most IP addresses, as of this
	// publication, are in IPv4 format; however, one should not rely on the
	// IP addresses being in IPv4 format indefinitely and should support
	// interpreting this field in either IPv4 or IPv6 format.
	IpAddress            string   `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UptimeCheckIp) Reset()         { *m = UptimeCheckIp{} }
func (m *UptimeCheckIp) String() string { return proto.CompactTextString(m) }
func (*UptimeCheckIp) ProtoMessage()    {}
func (*UptimeCheckIp) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca0e36dfc8221d8, []int{2}
}

func (m *UptimeCheckIp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UptimeCheckIp.Unmarshal(m, b)
}
func (m *UptimeCheckIp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UptimeCheckIp.Marshal(b, m, deterministic)
}
func (m *UptimeCheckIp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeCheckIp.Merge(m, src)
}
func (m *UptimeCheckIp) XXX_Size() int {
	return xxx_messageInfo_UptimeCheckIp.Size(m)
}
func (m *UptimeCheckIp) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeCheckIp.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeCheckIp proto.InternalMessageInfo

func (m *UptimeCheckIp) GetRegion() UptimeCheckRegion {
	if m != nil {
		return m.Region
	}
	return UptimeCheckRegion_REGION_UNSPECIFIED
}

func (m *UptimeCheckIp) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *UptimeCheckIp) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.monitoring.v3.UptimeCheckRegion", UptimeCheckRegion_name, UptimeCheckRegion_value)
	proto.RegisterEnum("google.monitoring.v3.GroupResourceType", GroupResourceType_name, GroupResourceType_value)
	proto.RegisterType((*InternalChecker)(nil), "google.monitoring.v3.InternalChecker")
	proto.RegisterType((*UptimeCheckConfig)(nil), "google.monitoring.v3.UptimeCheckConfig")
	proto.RegisterType((*UptimeCheckConfig_ResourceGroup)(nil), "google.monitoring.v3.UptimeCheckConfig.ResourceGroup")
	proto.RegisterType((*UptimeCheckConfig_HttpCheck)(nil), "google.monitoring.v3.UptimeCheckConfig.HttpCheck")
	proto.RegisterMapType((map[string]string)(nil), "google.monitoring.v3.UptimeCheckConfig.HttpCheck.HeadersEntry")
	proto.RegisterType((*UptimeCheckConfig_HttpCheck_BasicAuthentication)(nil), "google.monitoring.v3.UptimeCheckConfig.HttpCheck.BasicAuthentication")
	proto.RegisterType((*UptimeCheckConfig_TcpCheck)(nil), "google.monitoring.v3.UptimeCheckConfig.TcpCheck")
	proto.RegisterType((*UptimeCheckConfig_ContentMatcher)(nil), "google.monitoring.v3.UptimeCheckConfig.ContentMatcher")
	proto.RegisterType((*UptimeCheckIp)(nil), "google.monitoring.v3.UptimeCheckIp")
}

func init() { proto.RegisterFile("google/monitoring/v3/uptime.proto", fileDescriptor_7ca0e36dfc8221d8) }

var fileDescriptor_7ca0e36dfc8221d8 = []byte{
	// 1036 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xed, 0x6e, 0xe3, 0x44,
	0x17, 0xae, 0x9b, 0x36, 0x1f, 0x27, 0xfd, 0x70, 0xe7, 0xed, 0x0b, 0x6e, 0xa4, 0x2e, 0xdd, 0x22,
	0xa0, 0xea, 0x8f, 0x84, 0x6d, 0x04, 0x42, 0x8b, 0xb4, 0xc8, 0x49, 0x4d, 0x13, 0xa9, 0x4d, 0xa2,
	0x49, 0xb3, 0xc0, 0x52, 0x31, 0x72, 0xed, 0xa9, 0x63, 0x9a, 0x78, 0x8c, 0x67, 0xdc, 0xa5, 0xdc,
	0x02, 0x97, 0x81, 0xf8, 0xc3, 0x15, 0x70, 0x0d, 0x5c, 0x00, 0xd7, 0x83, 0x66, 0xec, 0x49, 0x9b,
	0xb6, 0x68, 0x5b, 0xfe, 0xcd, 0x73, 0x3e, 0x9e, 0x39, 0x73, 0xe6, 0x3c, 0x63, 0xc3, 0xf3, 0x80,
	0xb1, 0x60, 0x42, 0x1b, 0x53, 0x16, 0x85, 0x82, 0x25, 0x61, 0x14, 0x34, 0xae, 0x9a, 0x8d, 0x34,
	0x16, 0xe1, 0x94, 0xd6, 0xe3, 0x84, 0x09, 0x86, 0x36, 0xb3, 0x90, 0xfa, 0x4d, 0x48, 0xfd, 0xaa,
	0x59, 0xfb, 0x30, 0x4f, 0x74, 0xe3, 0x50, 0x27, 0x53, 0x9f, 0x24, 0x94, 0xb3, 0x34, 0xf1, 0xf2,
	0xd4, 0xda, 0xb3, 0x3c, 0x48, 0xa1, 0xf3, 0xf4, 0xa2, 0xe1, 0xa7, 0x89, 0x2b, 0x42, 0x16, 0x65,
	0xfe, 0xdd, 0xdf, 0x0d, 0x58, 0xef, 0x46, 0x82, 0x26, 0x91, 0x3b, 0x69, 0x8f, 0xa9, 0x77, 0x49,
	0x13, 0x84, 0x60, 0x29, 0x72, 0xa7, 0xd4, 0x32, 0x76, 0x8c, 0xbd, 0x0a, 0x56, 0x6b, 0xf4, 0x1c,
	0x56, 0xfc, 0x90, 0xc7, 0x13, 0xf7, 0x9a, 0x28, 0xdf, 0xa2, 0xf2, 0x55, 0x73, 0x5b, 0x4f, 0x86,
	0x58, 0x50, 0x8a, 0xa8, 0x78, 0xcb, 0x92, 0x4b, 0xab, 0xa0, 0xbc, 0x1a, 0xa2, 0x2d, 0x28, 0x07,
	0x5e, 0x4c, 0x7e, 0x61, 0x11, 0xb5, 0x96, 0x32, 0x57, 0xe0, 0xc5, 0x6f, 0x58, 0x44, 0xd1, 0xc7,
	0xb0, 0x1e, 0x53, 0x9a, 0x90, 0x38, 0x61, 0x3f, 0x52, 0x4f, 0x90, 0xd0, 0xb7, 0x8a, 0x2a, 0x62,
	0x55, 0x9a, 0x07, 0x99, 0xb5, 0xeb, 0xef, 0xfe, 0x5d, 0x85, 0x8d, 0x91, 0xea, 0x89, 0xaa, 0xb2,
	0xcd, 0xa2, 0x8b, 0x30, 0xf8, 0xaf, 0x95, 0xf6, 0x00, 0xdd, 0x6f, 0x98, 0x2a, 0xba, 0x7a, 0xb0,
	0x5d, 0xcf, 0x9b, 0xed, 0xc6, 0x61, 0xfd, 0x44, 0x47, 0xe1, 0x3c, 0xa8, 0xb3, 0x80, 0x37, 0xa6,
	0x77, 0x8d, 0xe8, 0x07, 0x58, 0xd3, 0x2c, 0x24, 0x48, 0x58, 0x1a, 0xab, 0x53, 0x56, 0x0f, 0x3e,
	0xab, 0x3f, 0x74, 0x71, 0xf5, 0x7b, 0xe7, 0xa8, 0x6b, 0xa6, 0x23, 0x99, 0xdc, 0x59, 0xc0, 0xab,
	0xc9, 0x6d, 0x03, 0xc2, 0x00, 0x63, 0x21, 0x62, 0xe2, 0xc9, 0x14, 0x6b, 0x59, 0x71, 0xbf, 0x78,
	0x2c, 0x77, 0x47, 0x88, 0x58, 0xe1, 0x8e, 0x81, 0x2b, 0x63, 0x0d, 0x50, 0x1f, 0x2a, 0xc2, 0xd3,
	0x94, 0x45, 0x45, 0xf9, 0xe9, 0x63, 0x29, 0x4f, 0xbd, 0x19, 0x63, 0x59, 0xe4, 0x6b, 0xf4, 0x02,
	0x8a, 0x31, 0x4d, 0x42, 0xe6, 0x5b, 0x25, 0xc5, 0xb6, 0xa5, 0xd9, 0xf4, 0xe8, 0xd5, 0x0f, 0xf3,
	0xd1, 0xc3, 0x79, 0x20, 0x6a, 0x42, 0x49, 0x52, 0xb3, 0x54, 0x58, 0xe5, 0x77, 0xe5, 0xe8, 0x48,
	0xe4, 0x82, 0xe9, 0xb1, 0x48, 0xd0, 0x48, 0x90, 0xa9, 0x2b, 0xbc, 0x31, 0x4d, 0xb8, 0x55, 0xd9,
	0x29, 0xec, 0x55, 0x0f, 0x3e, 0x7f, 0x6c, 0xfd, 0xed, 0x2c, 0xff, 0x24, 0x4b, 0xc7, 0xeb, 0xde,
	0x1c, 0xe6, 0x08, 0x83, 0xc9, 0xe9, 0x84, 0x7a, 0x42, 0x8d, 0x47, 0x10, 0xb2, 0x88, 0x5b, 0xb0,
	0x53, 0xd8, 0x5b, 0x3b, 0xf8, 0xe4, 0x9d, 0x5b, 0x60, 0x15, 0x8f, 0xd7, 0x35, 0x41, 0x86, 0x39,
	0xfa, 0x00, 0xaa, 0x21, 0x27, 0x61, 0x2e, 0x35, 0x6b, 0x7d, 0xc7, 0xd8, 0x2b, 0x63, 0x08, 0xb9,
	0x16, 0x1f, 0xc2, 0xb0, 0xa1, 0xbd, 0xd9, 0xad, 0xc8, 0x83, 0xad, 0xa9, 0x83, 0x7d, 0xf4, 0xf0,
	0xae, 0x77, 0x74, 0x8b, 0xcd, 0x70, 0xde, 0xc0, 0x6b, 0x3f, 0xc3, 0xea, 0xdc, 0x68, 0x29, 0x25,
	0xca, 0x85, 0xd4, 0x99, 0x91, 0x2b, 0x51, 0xe2, 0xae, 0x8f, 0x8e, 0x61, 0x36, 0x75, 0x44, 0x5c,
	0xc7, 0x99, 0x70, 0xfe, 0xf5, 0xc4, 0x8a, 0x4e, 0x73, 0x9f, 0x5e, 0xc7, 0x14, 0xaf, 0x24, 0xb7,
	0x50, 0xed, 0xcf, 0x02, 0x54, 0x66, 0x93, 0x87, 0xde, 0x87, 0x52, 0xca, 0x29, 0xe1, 0x7c, 0xa2,
	0x76, 0x2d, 0xe3, 0x62, 0xca, 0xe9, 0x90, 0x4f, 0xa4, 0x80, 0x63, 0x57, 0x8c, 0x73, 0x91, 0xaa,
	0xb5, 0xb2, 0xb1, 0x44, 0x28, 0x3d, 0x2e, 0x63, 0xb5, 0x46, 0xe7, 0x50, 0x71, 0x53, 0x31, 0x26,
	0x61, 0x74, 0xc1, 0x72, 0x71, 0x39, 0x4f, 0x16, 0x40, 0xbd, 0xe5, 0xf2, 0xd0, 0xb3, 0x53, 0x31,
	0xa6, 0x91, 0x08, 0xbd, 0x6c, 0xae, 0xca, 0x92, 0xb7, 0x1b, 0x5d, 0x30, 0xf9, 0x70, 0x4c, 0x5d,
	0x7e, 0x49, 0xc6, 0xd4, 0xf5, 0x65, 0xef, 0x97, 0x55, 0xa5, 0x55, 0x69, 0xeb, 0x64, 0x26, 0xf4,
	0x2d, 0x94, 0xb4, 0xb7, 0xa8, 0x6e, 0xe6, 0xd5, 0xd3, 0x8b, 0xc8, 0xb9, 0x9c, 0x48, 0x24, 0xd7,
	0x58, 0xd3, 0xd5, 0x4e, 0xe0, 0x7f, 0x0f, 0x54, 0x87, 0x6a, 0x50, 0x4e, 0xb9, 0xbc, 0xd3, 0xd9,
	0x23, 0x37, 0xc3, 0xd2, 0x17, 0xbb, 0x9c, 0xbf, 0x65, 0x89, 0x9f, 0xf7, 0x6f, 0x86, 0x6b, 0x2f,
	0x61, 0xe5, 0xf6, 0x3e, 0xc8, 0x84, 0xc2, 0x25, 0xbd, 0xce, 0x29, 0xe4, 0x12, 0x6d, 0xc2, 0xf2,
	0x95, 0x3b, 0x49, 0xf5, 0xfb, 0x98, 0x81, 0x97, 0x8b, 0x5f, 0x18, 0xb5, 0x67, 0x50, 0xd6, 0x02,
	0x9f, 0xdd, 0x85, 0x71, 0x73, 0x17, 0xb5, 0x7d, 0x58, 0x9b, 0x17, 0x90, 0x7c, 0xf9, 0x73, 0x09,
	0xe9, 0xa1, 0xca, 0x61, 0x0b, 0xa0, 0xac, 0xc7, 0xa2, 0xb5, 0x09, 0x48, 0xcd, 0x35, 0x49, 0xe8,
	0x4f, 0x29, 0xe5, 0x42, 0x4d, 0xd9, 0xee, 0xaf, 0x06, 0xac, 0xde, 0x6a, 0x57, 0x37, 0x46, 0x5f,
	0x41, 0x31, 0x13, 0x9d, 0x22, 0x7b, 0x82, 0xe6, 0xf2, 0x34, 0xd9, 0x98, 0x09, 0xcb, 0x1a, 0xa8,
	0x1b, 0xa3, 0x31, 0xda, 0x06, 0x08, 0x63, 0xe2, 0xfa, 0x7e, 0x42, 0x39, 0xcf, 0xbf, 0x53, 0x95,
	0x30, 0xb6, 0x33, 0xc3, 0x3e, 0x9d, 0xfb, 0xca, 0x64, 0xbc, 0xe8, 0x3d, 0x40, 0xd8, 0x39, 0xea,
	0xf6, 0x7b, 0x64, 0xd4, 0x1b, 0x0e, 0x9c, 0x76, 0xf7, 0xeb, 0xae, 0x73, 0x68, 0x2e, 0xa0, 0x12,
	0x14, 0x46, 0x43, 0xdb, 0x34, 0x10, 0x40, 0xd1, 0x19, 0xe1, 0xfe, 0xc0, 0x31, 0x17, 0xd1, 0x06,
	0xac, 0x0e, 0xfb, 0xa3, 0xd3, 0x0e, 0xb1, 0x4f, 0x1c, 0xdc, 0x6d, 0xdb, 0x66, 0x01, 0x99, 0xb0,
	0x62, 0x0f, 0xbb, 0x36, 0x19, 0xd8, 0x32, 0xb5, 0x6d, 0x2e, 0xed, 0x7f, 0x0f, 0x1b, 0xf7, 0x04,
	0x84, 0xb6, 0x61, 0x0b, 0x3b, 0xc3, 0xfe, 0x08, 0xb7, 0x1d, 0x72, 0xfa, 0xdd, 0xc0, 0xb9, 0xb3,
	0xdb, 0x0a, 0x94, 0xbb, 0xbd, 0xe1, 0xa9, 0xdd, 0x6b, 0x3b, 0xa6, 0x81, 0xb6, 0xe0, 0xff, 0xf6,
	0x37, 0x43, 0xe2, 0x1c, 0xb7, 0xc8, 0x71, 0xdf, 0x3e, 0x24, 0x2d, 0xfb, 0x58, 0x7a, 0xb0, 0xb9,
	0xd8, 0xfa, 0xcd, 0x00, 0xcb, 0x63, 0xd3, 0x07, 0xbb, 0xd6, 0xaa, 0x66, 0xc7, 0x1b, 0xc8, 0xf7,
	0x75, 0x60, 0xbc, 0x79, 0x95, 0x07, 0x05, 0x6c, 0xe2, 0x46, 0x41, 0x9d, 0x25, 0x41, 0x23, 0xa0,
	0x91, 0x7a, 0x7d, 0x1b, 0x99, 0xcb, 0x8d, 0x43, 0x3e, 0xff, 0x6f, 0xf2, 0xe5, 0x0d, 0xfa, 0x63,
	0xb1, 0x76, 0x94, 0x11, 0xb4, 0x27, 0x2c, 0xf5, 0xf5, 0xf7, 0x52, 0xee, 0xf5, 0xba, 0xf9, 0x97,
	0x76, 0x9e, 0x29, 0xe7, 0xd9, 0x8d, 0xf3, 0xec, 0x75, 0xf3, 0xbc, 0xa8, 0x36, 0x69, 0xfe, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x0a, 0x0a, 0xbb, 0x6b, 0xff, 0x08, 0x00, 0x00,
}
