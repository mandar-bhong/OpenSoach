// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/streetview/publish/v1/rpcmessages.proto

package publish

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// Specifies which view of the [Photo][google.streetview.publish.v1.Photo]
// to include in the response.
type PhotoView int32

const (
	// Server reponses do not include the download URL for the photo bytes.
	// The default value.
	PhotoView_BASIC PhotoView = 0
	// Server responses include the download URL for the photo bytes.
	PhotoView_INCLUDE_DOWNLOAD_URL PhotoView = 1
)

var PhotoView_name = map[int32]string{
	0: "BASIC",
	1: "INCLUDE_DOWNLOAD_URL",
}

var PhotoView_value = map[string]int32{
	"BASIC":                0,
	"INCLUDE_DOWNLOAD_URL": 1,
}

func (x PhotoView) String() string {
	return proto.EnumName(PhotoView_name, int32(x))
}

func (PhotoView) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{0}
}

// Request to create a [Photo][google.streetview.publish.v1.Photo].
type CreatePhotoRequest struct {
	// Required. Photo to create.
	Photo                *Photo   `protobuf:"bytes,1,opt,name=photo,proto3" json:"photo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePhotoRequest) Reset()         { *m = CreatePhotoRequest{} }
func (m *CreatePhotoRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePhotoRequest) ProtoMessage()    {}
func (*CreatePhotoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{0}
}

func (m *CreatePhotoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePhotoRequest.Unmarshal(m, b)
}
func (m *CreatePhotoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePhotoRequest.Marshal(b, m, deterministic)
}
func (m *CreatePhotoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePhotoRequest.Merge(m, src)
}
func (m *CreatePhotoRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePhotoRequest.Size(m)
}
func (m *CreatePhotoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePhotoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePhotoRequest proto.InternalMessageInfo

func (m *CreatePhotoRequest) GetPhoto() *Photo {
	if m != nil {
		return m.Photo
	}
	return nil
}

// Request to get a [Photo][google.streetview.publish.v1.Photo].
//
// By default
//
// * does not return the download URL for the photo bytes.
//
// Parameters:
//
// * `view` controls if the download URL for the photo bytes is returned.
type GetPhotoRequest struct {
	// Required. ID of the [Photo][google.streetview.publish.v1.Photo].
	PhotoId string `protobuf:"bytes,1,opt,name=photo_id,json=photoId,proto3" json:"photo_id,omitempty"`
	// Specifies if a download URL for the photo bytes should be returned in the
	// [Photo][google.streetview.publish.v1.Photo] response.
	View PhotoView `protobuf:"varint,2,opt,name=view,proto3,enum=google.streetview.publish.v1.PhotoView" json:"view,omitempty"`
	// The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	// information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	// If language_code is unspecified, the user's language preference for Google
	// services is used.
	LanguageCode         string   `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPhotoRequest) Reset()         { *m = GetPhotoRequest{} }
func (m *GetPhotoRequest) String() string { return proto.CompactTextString(m) }
func (*GetPhotoRequest) ProtoMessage()    {}
func (*GetPhotoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{1}
}

func (m *GetPhotoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPhotoRequest.Unmarshal(m, b)
}
func (m *GetPhotoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPhotoRequest.Marshal(b, m, deterministic)
}
func (m *GetPhotoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPhotoRequest.Merge(m, src)
}
func (m *GetPhotoRequest) XXX_Size() int {
	return xxx_messageInfo_GetPhotoRequest.Size(m)
}
func (m *GetPhotoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPhotoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPhotoRequest proto.InternalMessageInfo

func (m *GetPhotoRequest) GetPhotoId() string {
	if m != nil {
		return m.PhotoId
	}
	return ""
}

func (m *GetPhotoRequest) GetView() PhotoView {
	if m != nil {
		return m.View
	}
	return PhotoView_BASIC
}

func (m *GetPhotoRequest) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

// Request to get one or more [Photos][google.streetview.publish.v1.Photo].
// By default
//
// * does not return the download URL for the photo bytes.
//
// Parameters:
//
// * `view` controls if the download URL for the photo bytes is returned.
type BatchGetPhotosRequest struct {
	// Required. IDs of the [Photos][google.streetview.publish.v1.Photo]. HTTP GET
	// requests require the following syntax for the URL query parameter:
	// `photoIds=<id1>&photoIds=<id2>&...`.
	PhotoIds []string `protobuf:"bytes,1,rep,name=photo_ids,json=photoIds,proto3" json:"photo_ids,omitempty"`
	// Specifies if a download URL for the photo bytes should be returned in the
	// Photo response.
	View PhotoView `protobuf:"varint,2,opt,name=view,proto3,enum=google.streetview.publish.v1.PhotoView" json:"view,omitempty"`
	// The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	// information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	// If language_code is unspecified, the user's language preference for Google
	// services is used.
	LanguageCode         string   `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchGetPhotosRequest) Reset()         { *m = BatchGetPhotosRequest{} }
func (m *BatchGetPhotosRequest) String() string { return proto.CompactTextString(m) }
func (*BatchGetPhotosRequest) ProtoMessage()    {}
func (*BatchGetPhotosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{2}
}

func (m *BatchGetPhotosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchGetPhotosRequest.Unmarshal(m, b)
}
func (m *BatchGetPhotosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchGetPhotosRequest.Marshal(b, m, deterministic)
}
func (m *BatchGetPhotosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchGetPhotosRequest.Merge(m, src)
}
func (m *BatchGetPhotosRequest) XXX_Size() int {
	return xxx_messageInfo_BatchGetPhotosRequest.Size(m)
}
func (m *BatchGetPhotosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchGetPhotosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchGetPhotosRequest proto.InternalMessageInfo

func (m *BatchGetPhotosRequest) GetPhotoIds() []string {
	if m != nil {
		return m.PhotoIds
	}
	return nil
}

func (m *BatchGetPhotosRequest) GetView() PhotoView {
	if m != nil {
		return m.View
	}
	return PhotoView_BASIC
}

func (m *BatchGetPhotosRequest) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

// Response to batch get of [Photos][google.streetview.publish.v1.Photo].
type BatchGetPhotosResponse struct {
	// List of results for each individual
	// [Photo][google.streetview.publish.v1.Photo] requested, in the same order as
	// the requests in
	// [BatchGetPhotos][google.streetview.publish.v1.StreetViewPublishService.BatchGetPhotos].
	Results              []*PhotoResponse `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BatchGetPhotosResponse) Reset()         { *m = BatchGetPhotosResponse{} }
func (m *BatchGetPhotosResponse) String() string { return proto.CompactTextString(m) }
func (*BatchGetPhotosResponse) ProtoMessage()    {}
func (*BatchGetPhotosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{3}
}

func (m *BatchGetPhotosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchGetPhotosResponse.Unmarshal(m, b)
}
func (m *BatchGetPhotosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchGetPhotosResponse.Marshal(b, m, deterministic)
}
func (m *BatchGetPhotosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchGetPhotosResponse.Merge(m, src)
}
func (m *BatchGetPhotosResponse) XXX_Size() int {
	return xxx_messageInfo_BatchGetPhotosResponse.Size(m)
}
func (m *BatchGetPhotosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchGetPhotosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchGetPhotosResponse proto.InternalMessageInfo

func (m *BatchGetPhotosResponse) GetResults() []*PhotoResponse {
	if m != nil {
		return m.Results
	}
	return nil
}

// Response payload for a single
// [Photo][google.streetview.publish.v1.Photo]
// in batch operations including
// [BatchGetPhotos][google.streetview.publish.v1.StreetViewPublishService.BatchGetPhotos]
// and
// [BatchUpdatePhotos][google.streetview.publish.v1.StreetViewPublishService.BatchUpdatePhotos].
type PhotoResponse struct {
	// The status for the operation to get or update a single photo in the batch
	// request.
	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// The [Photo][google.streetview.publish.v1.Photo] resource, if the request
	// was successful.
	Photo                *Photo   `protobuf:"bytes,2,opt,name=photo,proto3" json:"photo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhotoResponse) Reset()         { *m = PhotoResponse{} }
func (m *PhotoResponse) String() string { return proto.CompactTextString(m) }
func (*PhotoResponse) ProtoMessage()    {}
func (*PhotoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{4}
}

func (m *PhotoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhotoResponse.Unmarshal(m, b)
}
func (m *PhotoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhotoResponse.Marshal(b, m, deterministic)
}
func (m *PhotoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhotoResponse.Merge(m, src)
}
func (m *PhotoResponse) XXX_Size() int {
	return xxx_messageInfo_PhotoResponse.Size(m)
}
func (m *PhotoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PhotoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PhotoResponse proto.InternalMessageInfo

func (m *PhotoResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PhotoResponse) GetPhoto() *Photo {
	if m != nil {
		return m.Photo
	}
	return nil
}

// Request to list all photos that belong to the user sending the request.
//
// By default
//
// * does not return the download URL for the photo bytes.
//
// Parameters:
//
// * `view` controls if the download URL for the photo bytes is returned.
// * `pageSize` determines the maximum number of photos to return.
// * `pageToken` is the next page token value returned from a previous
// [ListPhotos][google.streetview.publish.v1.StreetViewPublishService.ListPhotos]
//     request, if any.
// * `filter` allows filtering by a given parameter. 'placeId' is the only
// parameter supported at the moment.
type ListPhotosRequest struct {
	// Specifies if a download URL for the photos bytes should be returned in the
	// Photos response.
	View PhotoView `protobuf:"varint,1,opt,name=view,proto3,enum=google.streetview.publish.v1.PhotoView" json:"view,omitempty"`
	// The maximum number of photos to return.
	// `pageSize` must be non-negative. If `pageSize` is zero or is not provided,
	// the default page size of 100 is used.
	// The number of photos returned in the response may be less than `pageSize`
	// if the number of photos that belong to the user is less than `pageSize`.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The
	// [nextPageToken][google.streetview.publish.v1.ListPhotosResponse.next_page_token]
	// value returned from a previous
	// [ListPhotos][google.streetview.publish.v1.StreetViewPublishService.ListPhotos]
	// request, if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The filter expression. For example: `placeId=ChIJj61dQgK6j4AR4GeTYWZsKWw`.
	//
	// The only filter supported at the moment is `placeId`.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	// The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	// information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	// If language_code is unspecified, the user's language preference for Google
	// services is used.
	LanguageCode         string   `protobuf:"bytes,5,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPhotosRequest) Reset()         { *m = ListPhotosRequest{} }
func (m *ListPhotosRequest) String() string { return proto.CompactTextString(m) }
func (*ListPhotosRequest) ProtoMessage()    {}
func (*ListPhotosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{5}
}

func (m *ListPhotosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPhotosRequest.Unmarshal(m, b)
}
func (m *ListPhotosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPhotosRequest.Marshal(b, m, deterministic)
}
func (m *ListPhotosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPhotosRequest.Merge(m, src)
}
func (m *ListPhotosRequest) XXX_Size() int {
	return xxx_messageInfo_ListPhotosRequest.Size(m)
}
func (m *ListPhotosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPhotosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPhotosRequest proto.InternalMessageInfo

func (m *ListPhotosRequest) GetView() PhotoView {
	if m != nil {
		return m.View
	}
	return PhotoView_BASIC
}

func (m *ListPhotosRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListPhotosRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListPhotosRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListPhotosRequest) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

// Response to list all photos that belong to a user.
type ListPhotosResponse struct {
	// List of photos. The
	// [pageSize][google.streetview.publish.v1.ListPhotosRequest.page_size] field
	// in the request determines the number of items returned.
	Photos []*Photo `protobuf:"bytes,1,rep,name=photos,proto3" json:"photos,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more
	// results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPhotosResponse) Reset()         { *m = ListPhotosResponse{} }
func (m *ListPhotosResponse) String() string { return proto.CompactTextString(m) }
func (*ListPhotosResponse) ProtoMessage()    {}
func (*ListPhotosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{6}
}

func (m *ListPhotosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPhotosResponse.Unmarshal(m, b)
}
func (m *ListPhotosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPhotosResponse.Marshal(b, m, deterministic)
}
func (m *ListPhotosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPhotosResponse.Merge(m, src)
}
func (m *ListPhotosResponse) XXX_Size() int {
	return xxx_messageInfo_ListPhotosResponse.Size(m)
}
func (m *ListPhotosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPhotosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPhotosResponse proto.InternalMessageInfo

func (m *ListPhotosResponse) GetPhotos() []*Photo {
	if m != nil {
		return m.Photos
	}
	return nil
}

func (m *ListPhotosResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request to update the metadata of a
// [Photo][google.streetview.publish.v1.Photo]. Updating the pixels of a photo
// is not supported.
type UpdatePhotoRequest struct {
	// Required. [Photo][google.streetview.publish.v1.Photo] object containing the
	// new metadata.
	Photo *Photo `protobuf:"bytes,1,opt,name=photo,proto3" json:"photo,omitempty"`
	// Mask that identifies fields on the photo metadata to update.
	// If not present, the old [Photo][google.streetview.publish.v1.Photo]
	// metadata is entirely replaced with the
	// new [Photo][google.streetview.publish.v1.Photo] metadata in this request.
	// The update fails if invalid fields are specified. Multiple fields can be
	// specified in a comma-delimited list.
	//
	// The following fields are valid:
	//
	// * `pose.heading`
	// * `pose.latLngPair`
	// * `pose.pitch`
	// * `pose.roll`
	// * `pose.level`
	// * `pose.altitude`
	// * `connections`
	// * `places`
	//
	//
	// <aside class="note"><b>Note:</b>  When
	// [updateMask][google.streetview.publish.v1.UpdatePhotoRequest.update_mask]
	// contains repeated fields, the entire set of repeated values get replaced
	// with the new contents. For example, if
	// [updateMask][google.streetview.publish.v1.UpdatePhotoRequest.update_mask]
	// contains `connections` and `UpdatePhotoRequest.photo.connections` is empty,
	// all connections are removed.</aside>
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdatePhotoRequest) Reset()         { *m = UpdatePhotoRequest{} }
func (m *UpdatePhotoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePhotoRequest) ProtoMessage()    {}
func (*UpdatePhotoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{7}
}

func (m *UpdatePhotoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePhotoRequest.Unmarshal(m, b)
}
func (m *UpdatePhotoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePhotoRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePhotoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePhotoRequest.Merge(m, src)
}
func (m *UpdatePhotoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePhotoRequest.Size(m)
}
func (m *UpdatePhotoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePhotoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePhotoRequest proto.InternalMessageInfo

func (m *UpdatePhotoRequest) GetPhoto() *Photo {
	if m != nil {
		return m.Photo
	}
	return nil
}

func (m *UpdatePhotoRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Request to update the metadata of photos.
// Updating the pixels of photos is not supported.
type BatchUpdatePhotosRequest struct {
	// Required. List of
	// [UpdatePhotoRequests][google.streetview.publish.v1.UpdatePhotoRequest].
	UpdatePhotoRequests  []*UpdatePhotoRequest `protobuf:"bytes,1,rep,name=update_photo_requests,json=updatePhotoRequests,proto3" json:"update_photo_requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BatchUpdatePhotosRequest) Reset()         { *m = BatchUpdatePhotosRequest{} }
func (m *BatchUpdatePhotosRequest) String() string { return proto.CompactTextString(m) }
func (*BatchUpdatePhotosRequest) ProtoMessage()    {}
func (*BatchUpdatePhotosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{8}
}

func (m *BatchUpdatePhotosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchUpdatePhotosRequest.Unmarshal(m, b)
}
func (m *BatchUpdatePhotosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchUpdatePhotosRequest.Marshal(b, m, deterministic)
}
func (m *BatchUpdatePhotosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchUpdatePhotosRequest.Merge(m, src)
}
func (m *BatchUpdatePhotosRequest) XXX_Size() int {
	return xxx_messageInfo_BatchUpdatePhotosRequest.Size(m)
}
func (m *BatchUpdatePhotosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchUpdatePhotosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchUpdatePhotosRequest proto.InternalMessageInfo

func (m *BatchUpdatePhotosRequest) GetUpdatePhotoRequests() []*UpdatePhotoRequest {
	if m != nil {
		return m.UpdatePhotoRequests
	}
	return nil
}

// Response to batch update of metadata of one or more
// [Photos][google.streetview.publish.v1.Photo].
type BatchUpdatePhotosResponse struct {
	// List of results for each individual
	// [Photo][google.streetview.publish.v1.Photo] updated, in the same order as
	// the request.
	Results              []*PhotoResponse `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BatchUpdatePhotosResponse) Reset()         { *m = BatchUpdatePhotosResponse{} }
func (m *BatchUpdatePhotosResponse) String() string { return proto.CompactTextString(m) }
func (*BatchUpdatePhotosResponse) ProtoMessage()    {}
func (*BatchUpdatePhotosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{9}
}

func (m *BatchUpdatePhotosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchUpdatePhotosResponse.Unmarshal(m, b)
}
func (m *BatchUpdatePhotosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchUpdatePhotosResponse.Marshal(b, m, deterministic)
}
func (m *BatchUpdatePhotosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchUpdatePhotosResponse.Merge(m, src)
}
func (m *BatchUpdatePhotosResponse) XXX_Size() int {
	return xxx_messageInfo_BatchUpdatePhotosResponse.Size(m)
}
func (m *BatchUpdatePhotosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchUpdatePhotosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchUpdatePhotosResponse proto.InternalMessageInfo

func (m *BatchUpdatePhotosResponse) GetResults() []*PhotoResponse {
	if m != nil {
		return m.Results
	}
	return nil
}

// Request to delete a [Photo][google.streetview.publish.v1.Photo].
type DeletePhotoRequest struct {
	// Required. ID of the [Photo][google.streetview.publish.v1.Photo].
	PhotoId              string   `protobuf:"bytes,1,opt,name=photo_id,json=photoId,proto3" json:"photo_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePhotoRequest) Reset()         { *m = DeletePhotoRequest{} }
func (m *DeletePhotoRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePhotoRequest) ProtoMessage()    {}
func (*DeletePhotoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{10}
}

func (m *DeletePhotoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePhotoRequest.Unmarshal(m, b)
}
func (m *DeletePhotoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePhotoRequest.Marshal(b, m, deterministic)
}
func (m *DeletePhotoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePhotoRequest.Merge(m, src)
}
func (m *DeletePhotoRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePhotoRequest.Size(m)
}
func (m *DeletePhotoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePhotoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePhotoRequest proto.InternalMessageInfo

func (m *DeletePhotoRequest) GetPhotoId() string {
	if m != nil {
		return m.PhotoId
	}
	return ""
}

// Request to delete multiple [Photos][google.streetview.publish.v1.Photo].
type BatchDeletePhotosRequest struct {
	// Required. IDs of the [Photos][google.streetview.publish.v1.Photo]. HTTP
	// GET requests require the following syntax for the URL query parameter:
	// `photoIds=<id1>&photoIds=<id2>&...`.
	PhotoIds             []string `protobuf:"bytes,1,rep,name=photo_ids,json=photoIds,proto3" json:"photo_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchDeletePhotosRequest) Reset()         { *m = BatchDeletePhotosRequest{} }
func (m *BatchDeletePhotosRequest) String() string { return proto.CompactTextString(m) }
func (*BatchDeletePhotosRequest) ProtoMessage()    {}
func (*BatchDeletePhotosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{11}
}

func (m *BatchDeletePhotosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchDeletePhotosRequest.Unmarshal(m, b)
}
func (m *BatchDeletePhotosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchDeletePhotosRequest.Marshal(b, m, deterministic)
}
func (m *BatchDeletePhotosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchDeletePhotosRequest.Merge(m, src)
}
func (m *BatchDeletePhotosRequest) XXX_Size() int {
	return xxx_messageInfo_BatchDeletePhotosRequest.Size(m)
}
func (m *BatchDeletePhotosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchDeletePhotosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchDeletePhotosRequest proto.InternalMessageInfo

func (m *BatchDeletePhotosRequest) GetPhotoIds() []string {
	if m != nil {
		return m.PhotoIds
	}
	return nil
}

// Response to batch delete of one or more
// [Photos][google.streetview.publish.v1.Photo].
type BatchDeletePhotosResponse struct {
	// The status for the operation to delete a single
	// [Photo][google.streetview.publish.v1.Photo] in the batch request.
	Status               []*status.Status `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BatchDeletePhotosResponse) Reset()         { *m = BatchDeletePhotosResponse{} }
func (m *BatchDeletePhotosResponse) String() string { return proto.CompactTextString(m) }
func (*BatchDeletePhotosResponse) ProtoMessage()    {}
func (*BatchDeletePhotosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e56ff94407a6aca7, []int{12}
}

func (m *BatchDeletePhotosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchDeletePhotosResponse.Unmarshal(m, b)
}
func (m *BatchDeletePhotosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchDeletePhotosResponse.Marshal(b, m, deterministic)
}
func (m *BatchDeletePhotosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchDeletePhotosResponse.Merge(m, src)
}
func (m *BatchDeletePhotosResponse) XXX_Size() int {
	return xxx_messageInfo_BatchDeletePhotosResponse.Size(m)
}
func (m *BatchDeletePhotosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchDeletePhotosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchDeletePhotosResponse proto.InternalMessageInfo

func (m *BatchDeletePhotosResponse) GetStatus() []*status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.streetview.publish.v1.PhotoView", PhotoView_name, PhotoView_value)
	proto.RegisterType((*CreatePhotoRequest)(nil), "google.streetview.publish.v1.CreatePhotoRequest")
	proto.RegisterType((*GetPhotoRequest)(nil), "google.streetview.publish.v1.GetPhotoRequest")
	proto.RegisterType((*BatchGetPhotosRequest)(nil), "google.streetview.publish.v1.BatchGetPhotosRequest")
	proto.RegisterType((*BatchGetPhotosResponse)(nil), "google.streetview.publish.v1.BatchGetPhotosResponse")
	proto.RegisterType((*PhotoResponse)(nil), "google.streetview.publish.v1.PhotoResponse")
	proto.RegisterType((*ListPhotosRequest)(nil), "google.streetview.publish.v1.ListPhotosRequest")
	proto.RegisterType((*ListPhotosResponse)(nil), "google.streetview.publish.v1.ListPhotosResponse")
	proto.RegisterType((*UpdatePhotoRequest)(nil), "google.streetview.publish.v1.UpdatePhotoRequest")
	proto.RegisterType((*BatchUpdatePhotosRequest)(nil), "google.streetview.publish.v1.BatchUpdatePhotosRequest")
	proto.RegisterType((*BatchUpdatePhotosResponse)(nil), "google.streetview.publish.v1.BatchUpdatePhotosResponse")
	proto.RegisterType((*DeletePhotoRequest)(nil), "google.streetview.publish.v1.DeletePhotoRequest")
	proto.RegisterType((*BatchDeletePhotosRequest)(nil), "google.streetview.publish.v1.BatchDeletePhotosRequest")
	proto.RegisterType((*BatchDeletePhotosResponse)(nil), "google.streetview.publish.v1.BatchDeletePhotosResponse")
}

func init() {
	proto.RegisterFile("google/streetview/publish/v1/rpcmessages.proto", fileDescriptor_e56ff94407a6aca7)
}

var fileDescriptor_e56ff94407a6aca7 = []byte{
	// 669 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x5d, 0x4f, 0xd3, 0x5e,
	0x18, 0xff, 0x17, 0xd8, 0x60, 0x0f, 0x7f, 0x04, 0x8f, 0x82, 0x65, 0x62, 0xb2, 0x94, 0x44, 0x17,
	0x34, 0x2d, 0xe0, 0x85, 0x31, 0xbb, 0x62, 0x1b, 0x12, 0x92, 0xf1, 0x92, 0x4e, 0x34, 0xf1, 0xa6,
	0xe9, 0xda, 0x87, 0xd2, 0x50, 0x76, 0x6a, 0xcf, 0xe9, 0x50, 0xae, 0xfc, 0x00, 0xea, 0xa5, 0xdf,
	0xc9, 0x6f, 0x65, 0x7a, 0x7a, 0x0e, 0x8c, 0x6d, 0x2e, 0xd3, 0x10, 0xef, 0x7a, 0x9e, 0x97, 0x5f,
	0x7f, 0xe7, 0xf7, 0xbc, 0x1c, 0x30, 0x03, 0x4a, 0x83, 0x08, 0x2d, 0xc6, 0x13, 0x44, 0xde, 0x0b,
	0xf1, 0xd2, 0x8a, 0xd3, 0x4e, 0x14, 0xb2, 0x33, 0xab, 0xb7, 0x65, 0x25, 0xb1, 0x77, 0x81, 0x8c,
	0xb9, 0x01, 0x32, 0x33, 0x4e, 0x28, 0xa7, 0x64, 0x2d, 0x8f, 0x37, 0x6f, 0xe2, 0x4d, 0x19, 0x6f,
	0xf6, 0xb6, 0xca, 0x15, 0x89, 0x26, 0x62, 0x3b, 0xe9, 0xa9, 0x75, 0x1a, 0x62, 0xe4, 0x3b, 0x17,
	0x2e, 0x3b, 0xcf, 0xf3, 0xcb, 0x8f, 0x64, 0x44, 0x12, 0x7b, 0x16, 0xe3, 0x2e, 0x4f, 0x25, 0x70,
	0xf9, 0xc5, 0x78, 0x22, 0xc8, 0x68, 0x9a, 0x78, 0x8a, 0x86, 0x71, 0x04, 0xa4, 0x91, 0xa0, 0xcb,
	0xf1, 0xf8, 0x8c, 0x72, 0x6a, 0xe3, 0xc7, 0x14, 0x19, 0x27, 0xaf, 0xa1, 0x10, 0x67, 0x67, 0x5d,
	0xab, 0x68, 0xd5, 0xf9, 0xed, 0x75, 0x73, 0x1c, 0x59, 0x33, 0x4f, 0xcd, 0x33, 0x8c, 0xef, 0x1a,
	0x2c, 0xee, 0x21, 0xbf, 0x05, 0xb7, 0x0a, 0x73, 0xc2, 0xe9, 0x84, 0xbe, 0x40, 0x2c, 0xd9, 0xb3,
	0xe2, 0xbc, 0xef, 0x93, 0x1a, 0xcc, 0x64, 0x70, 0xfa, 0x54, 0x45, 0xab, 0xde, 0xdb, 0x7e, 0x36,
	0xc1, 0x8f, 0xde, 0x85, 0x78, 0x69, 0x8b, 0x24, 0xb2, 0x0e, 0x0b, 0x91, 0xdb, 0x0d, 0x52, 0x37,
	0x40, 0xc7, 0xa3, 0x3e, 0xea, 0xd3, 0x02, 0xfc, 0x7f, 0x65, 0x6c, 0x50, 0x1f, 0x8d, 0x1f, 0x1a,
	0x2c, 0xd7, 0x5d, 0xee, 0x9d, 0x29, 0x56, 0x4c, 0xd1, 0x7a, 0x0c, 0x25, 0x45, 0x8b, 0xe9, 0x5a,
	0x65, 0xba, 0x5a, 0xb2, 0xe7, 0x24, 0x2f, 0xf6, 0x0f, 0x88, 0x39, 0xb0, 0x32, 0xc8, 0x8b, 0xc5,
	0xb4, 0xcb, 0x90, 0xec, 0xc2, 0x6c, 0x82, 0x2c, 0x8d, 0x78, 0x4e, 0x6b, 0x7e, 0xfb, 0xf9, 0x24,
	0x05, 0x90, 0xd9, 0xb6, 0xca, 0x35, 0x7a, 0xb0, 0x70, 0xcb, 0x43, 0x36, 0xa0, 0x98, 0xb7, 0x8a,
	0xac, 0x2b, 0x51, 0xb0, 0x49, 0xec, 0x99, 0x6d, 0xe1, 0xb1, 0x65, 0xc4, 0x4d, 0x0b, 0x4c, 0xfd,
	0x71, 0x0b, 0xfc, 0xd4, 0xe0, 0x7e, 0x2b, 0x64, 0x03, 0x6a, 0x2b, 0x41, 0xb5, 0xbf, 0x11, 0x34,
	0x2b, 0x55, 0x26, 0x26, 0x0b, 0xaf, 0x50, 0x30, 0x2a, 0xd8, 0x73, 0x99, 0xa1, 0x1d, 0x5e, 0x21,
	0x79, 0x02, 0x20, 0x9c, 0x9c, 0x9e, 0x63, 0x57, 0x4a, 0x2d, 0xc2, 0xdf, 0x66, 0x06, 0xb2, 0x02,
	0xc5, 0xd3, 0x30, 0xe2, 0x98, 0xe8, 0x33, 0xc2, 0x25, 0x4f, 0xc3, 0x45, 0x2a, 0x8c, 0x28, 0xd2,
	0x67, 0x20, 0xfd, 0x57, 0x91, 0x42, 0xd6, 0xa0, 0x28, 0xae, 0xaa, 0xea, 0x33, 0x91, 0x3a, 0x32,
	0x85, 0x3c, 0x85, 0xc5, 0x2e, 0x7e, 0xe2, 0x4e, 0x1f, 0xe7, 0x29, 0xf1, 0xe7, 0x85, 0xcc, 0x7c,
	0xac, 0x78, 0x1b, 0x5f, 0x35, 0x20, 0x27, 0xb1, 0x7f, 0x77, 0xb3, 0x49, 0x6a, 0x30, 0x9f, 0x0a,
	0x40, 0xb1, 0x48, 0x64, 0x65, 0xcb, 0x0a, 0x40, 0xed, 0x1a, 0xf3, 0x4d, 0xb6, 0x6b, 0x0e, 0x5c,
	0x76, 0x6e, 0x43, 0x1e, 0x9e, 0x7d, 0x1b, 0x5f, 0x34, 0xd0, 0x45, 0xbf, 0xf6, 0x71, 0xba, 0x2e,
	0xae, 0x0f, 0xcb, 0x12, 0x39, 0x9f, 0xa8, 0x24, 0xb7, 0x2b, 0x7d, 0x36, 0xc7, 0x93, 0x1c, 0xbe,
	0xa5, 0xfd, 0x20, 0x1d, 0xb2, 0x31, 0xa3, 0x03, 0xab, 0x23, 0x18, 0xdc, 0xed, 0xd0, 0x58, 0x40,
	0x9a, 0x18, 0xe1, 0x80, 0xe8, 0xbf, 0xdf, 0x60, 0xc6, 0x2b, 0x29, 0x4b, 0x5f, 0xd6, 0x44, 0x1b,
	0xc6, 0xd8, 0x93, 0xb7, 0xb9, 0x9d, 0x38, 0x62, 0x54, 0xa7, 0xc7, 0x8f, 0xea, 0xc6, 0x26, 0x94,
	0xae, 0xe7, 0x85, 0x94, 0xa0, 0x50, 0xdf, 0x69, 0xef, 0x37, 0x96, 0xfe, 0x23, 0x3a, 0x3c, 0xdc,
	0x3f, 0x6c, 0xb4, 0x4e, 0x9a, 0xbb, 0x4e, 0xf3, 0xe8, 0xfd, 0x61, 0xeb, 0x68, 0xa7, 0xe9, 0x9c,
	0xd8, 0xad, 0x25, 0xad, 0xfe, 0x4d, 0x83, 0xaa, 0x47, 0x2f, 0x14, 0x66, 0x80, 0xd4, 0x4c, 0x03,
	0x6f, 0xb4, 0x50, 0xf5, 0xb5, 0xb6, 0x30, 0x67, 0xe8, 0xc7, 0xb9, 0xd5, 0x8e, 0xbd, 0x03, 0xf9,
	0x9a, 0x7d, 0x68, 0x28, 0x0c, 0x9a, 0xcd, 0x8d, 0x49, 0x93, 0xc0, 0x0a, 0xb0, 0x2b, 0x7a, 0xc9,
	0xca, 0x5d, 0x6e, 0x1c, 0xb2, 0xd1, 0xaf, 0x51, 0x4d, 0x7e, 0x76, 0x8a, 0x22, 0xfe, 0xe5, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x7d, 0x68, 0xfd, 0x45, 0x07, 0x00, 0x00,
}
