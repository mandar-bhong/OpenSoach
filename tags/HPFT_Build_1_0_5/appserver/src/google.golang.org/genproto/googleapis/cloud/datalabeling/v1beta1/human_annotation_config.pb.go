// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datalabeling/v1beta1/human_annotation_config.proto

package datalabeling

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type StringAggregationType int32

const (
	StringAggregationType_STRING_AGGREGATION_TYPE_UNSPECIFIED StringAggregationType = 0
	// Majority vote to aggregate answers.
	StringAggregationType_MAJORITY_VOTE StringAggregationType = 1
	// Unanimous answers will be adopted.
	StringAggregationType_UNANIMOUS_VOTE StringAggregationType = 2
	// Preserve all answers by crowd compute.
	StringAggregationType_NO_AGGREGATION StringAggregationType = 3
)

var StringAggregationType_name = map[int32]string{
	0: "STRING_AGGREGATION_TYPE_UNSPECIFIED",
	1: "MAJORITY_VOTE",
	2: "UNANIMOUS_VOTE",
	3: "NO_AGGREGATION",
}

var StringAggregationType_value = map[string]int32{
	"STRING_AGGREGATION_TYPE_UNSPECIFIED": 0,
	"MAJORITY_VOTE":                       1,
	"UNANIMOUS_VOTE":                      2,
	"NO_AGGREGATION":                      3,
}

func (x StringAggregationType) String() string {
	return proto.EnumName(StringAggregationType_name, int32(x))
}

func (StringAggregationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{0}
}

// Configuration for how human labeling task should be done.
type HumanAnnotationConfig struct {
	// Required except for LabelAudio case. Instruction resource name.
	Instruction string `protobuf:"bytes,1,opt,name=instruction,proto3" json:"instruction,omitempty"`
	// Required. A human-readable name for AnnotatedDataset defined by
	// users. Maximum of 64 characters
	// .
	AnnotatedDatasetDisplayName string `protobuf:"bytes,2,opt,name=annotated_dataset_display_name,json=annotatedDatasetDisplayName,proto3" json:"annotated_dataset_display_name,omitempty"`
	// Optional. A human-readable description for AnnotatedDataset.
	// The description can be up to 10000 characters long.
	AnnotatedDatasetDescription string `protobuf:"bytes,3,opt,name=annotated_dataset_description,json=annotatedDatasetDescription,proto3" json:"annotated_dataset_description,omitempty"`
	// Optional. A human-readable label used to logically group labeling tasks.
	// This string must match the regular expression `[a-zA-Z\\d_-]{0,128}`.
	LabelGroup string `protobuf:"bytes,4,opt,name=label_group,json=labelGroup,proto3" json:"label_group,omitempty"`
	// Optional. The Language of this question, as a
	// [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt).
	// Default value is en-US.
	// Only need to set this when task is language related. For example, French
	// text classification or Chinese audio transcription.
	LanguageCode string `protobuf:"bytes,5,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. Replication of questions. Each question will be sent to up to
	// this number of contributors to label. Aggregated answers will be returned.
	// Default is set to 1.
	// For image related labeling, valid values are 1, 3, 5.
	ReplicaCount int32 `protobuf:"varint,6,opt,name=replica_count,json=replicaCount,proto3" json:"replica_count,omitempty"`
	// Optional. Maximum duration for contributors to answer a question. Default
	// is 1800 seconds.
	QuestionDuration *duration.Duration `protobuf:"bytes,7,opt,name=question_duration,json=questionDuration,proto3" json:"question_duration,omitempty"`
	// Optional. If you want your own labeling contributors to manage and work on
	// this labeling request, you can set these contributors here. We will give
	// them access to the question types in crowdcompute. Note that these
	// emails must be registered in crowdcompute worker UI:
	// https://crowd-compute.appspot.com/
	ContributorEmails    []string `protobuf:"bytes,9,rep,name=contributor_emails,json=contributorEmails,proto3" json:"contributor_emails,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HumanAnnotationConfig) Reset()         { *m = HumanAnnotationConfig{} }
func (m *HumanAnnotationConfig) String() string { return proto.CompactTextString(m) }
func (*HumanAnnotationConfig) ProtoMessage()    {}
func (*HumanAnnotationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{0}
}

func (m *HumanAnnotationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HumanAnnotationConfig.Unmarshal(m, b)
}
func (m *HumanAnnotationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HumanAnnotationConfig.Marshal(b, m, deterministic)
}
func (m *HumanAnnotationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HumanAnnotationConfig.Merge(m, src)
}
func (m *HumanAnnotationConfig) XXX_Size() int {
	return xxx_messageInfo_HumanAnnotationConfig.Size(m)
}
func (m *HumanAnnotationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HumanAnnotationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HumanAnnotationConfig proto.InternalMessageInfo

func (m *HumanAnnotationConfig) GetInstruction() string {
	if m != nil {
		return m.Instruction
	}
	return ""
}

func (m *HumanAnnotationConfig) GetAnnotatedDatasetDisplayName() string {
	if m != nil {
		return m.AnnotatedDatasetDisplayName
	}
	return ""
}

func (m *HumanAnnotationConfig) GetAnnotatedDatasetDescription() string {
	if m != nil {
		return m.AnnotatedDatasetDescription
	}
	return ""
}

func (m *HumanAnnotationConfig) GetLabelGroup() string {
	if m != nil {
		return m.LabelGroup
	}
	return ""
}

func (m *HumanAnnotationConfig) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *HumanAnnotationConfig) GetReplicaCount() int32 {
	if m != nil {
		return m.ReplicaCount
	}
	return 0
}

func (m *HumanAnnotationConfig) GetQuestionDuration() *duration.Duration {
	if m != nil {
		return m.QuestionDuration
	}
	return nil
}

func (m *HumanAnnotationConfig) GetContributorEmails() []string {
	if m != nil {
		return m.ContributorEmails
	}
	return nil
}

// Config for image classification human labeling task.
type ImageClassificationConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet string `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Optional. If allow_multi_label is true, contributors are able to choose
	// multiple labels for one image.
	AllowMultiLabel bool `protobuf:"varint,2,opt,name=allow_multi_label,json=allowMultiLabel,proto3" json:"allow_multi_label,omitempty"`
	// Optional. The type of how to aggregate answers.
	AnswerAggregationType StringAggregationType `protobuf:"varint,3,opt,name=answer_aggregation_type,json=answerAggregationType,proto3,enum=google.cloud.datalabeling.v1beta1.StringAggregationType" json:"answer_aggregation_type,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}              `json:"-"`
	XXX_unrecognized      []byte                `json:"-"`
	XXX_sizecache         int32                 `json:"-"`
}

func (m *ImageClassificationConfig) Reset()         { *m = ImageClassificationConfig{} }
func (m *ImageClassificationConfig) String() string { return proto.CompactTextString(m) }
func (*ImageClassificationConfig) ProtoMessage()    {}
func (*ImageClassificationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{1}
}

func (m *ImageClassificationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageClassificationConfig.Unmarshal(m, b)
}
func (m *ImageClassificationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageClassificationConfig.Marshal(b, m, deterministic)
}
func (m *ImageClassificationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageClassificationConfig.Merge(m, src)
}
func (m *ImageClassificationConfig) XXX_Size() int {
	return xxx_messageInfo_ImageClassificationConfig.Size(m)
}
func (m *ImageClassificationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageClassificationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ImageClassificationConfig proto.InternalMessageInfo

func (m *ImageClassificationConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *ImageClassificationConfig) GetAllowMultiLabel() bool {
	if m != nil {
		return m.AllowMultiLabel
	}
	return false
}

func (m *ImageClassificationConfig) GetAnswerAggregationType() StringAggregationType {
	if m != nil {
		return m.AnswerAggregationType
	}
	return StringAggregationType_STRING_AGGREGATION_TYPE_UNSPECIFIED
}

// Config for image bounding poly (and bounding box) human labeling task.
type BoundingPolyConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet string `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Optional. Instruction message showed on contributors UI.
	InstructionMessage   string   `protobuf:"bytes,2,opt,name=instruction_message,json=instructionMessage,proto3" json:"instruction_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoundingPolyConfig) Reset()         { *m = BoundingPolyConfig{} }
func (m *BoundingPolyConfig) String() string { return proto.CompactTextString(m) }
func (*BoundingPolyConfig) ProtoMessage()    {}
func (*BoundingPolyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{2}
}

func (m *BoundingPolyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingPolyConfig.Unmarshal(m, b)
}
func (m *BoundingPolyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingPolyConfig.Marshal(b, m, deterministic)
}
func (m *BoundingPolyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingPolyConfig.Merge(m, src)
}
func (m *BoundingPolyConfig) XXX_Size() int {
	return xxx_messageInfo_BoundingPolyConfig.Size(m)
}
func (m *BoundingPolyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingPolyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingPolyConfig proto.InternalMessageInfo

func (m *BoundingPolyConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *BoundingPolyConfig) GetInstructionMessage() string {
	if m != nil {
		return m.InstructionMessage
	}
	return ""
}

// Config for image polyline human labeling task.
type PolylineConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet string `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Optional. Instruction message showed on contributors UI.
	InstructionMessage   string   `protobuf:"bytes,2,opt,name=instruction_message,json=instructionMessage,proto3" json:"instruction_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolylineConfig) Reset()         { *m = PolylineConfig{} }
func (m *PolylineConfig) String() string { return proto.CompactTextString(m) }
func (*PolylineConfig) ProtoMessage()    {}
func (*PolylineConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{3}
}

func (m *PolylineConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolylineConfig.Unmarshal(m, b)
}
func (m *PolylineConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolylineConfig.Marshal(b, m, deterministic)
}
func (m *PolylineConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolylineConfig.Merge(m, src)
}
func (m *PolylineConfig) XXX_Size() int {
	return xxx_messageInfo_PolylineConfig.Size(m)
}
func (m *PolylineConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PolylineConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PolylineConfig proto.InternalMessageInfo

func (m *PolylineConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *PolylineConfig) GetInstructionMessage() string {
	if m != nil {
		return m.InstructionMessage
	}
	return ""
}

// Config for image segmentation
type SegmentationConfig struct {
	// Required. Annotation spec set resource name. format:
	// projects/{project_id}/annotationSpecSets/{annotation_spec_set_id}
	AnnotationSpecSet string `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Instruction message showed on labelers UI.
	InstructionMessage   string   `protobuf:"bytes,2,opt,name=instruction_message,json=instructionMessage,proto3" json:"instruction_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentationConfig) Reset()         { *m = SegmentationConfig{} }
func (m *SegmentationConfig) String() string { return proto.CompactTextString(m) }
func (*SegmentationConfig) ProtoMessage()    {}
func (*SegmentationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{4}
}

func (m *SegmentationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentationConfig.Unmarshal(m, b)
}
func (m *SegmentationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentationConfig.Marshal(b, m, deterministic)
}
func (m *SegmentationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentationConfig.Merge(m, src)
}
func (m *SegmentationConfig) XXX_Size() int {
	return xxx_messageInfo_SegmentationConfig.Size(m)
}
func (m *SegmentationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentationConfig proto.InternalMessageInfo

func (m *SegmentationConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *SegmentationConfig) GetInstructionMessage() string {
	if m != nil {
		return m.InstructionMessage
	}
	return ""
}

// Config for video classification human labeling task.
// Currently two types of video classification are supported:
// 1. Assign labels on the entire video.
// 2. Split the video into multiple video clips based on camera shot, and
// assign labels on each video clip.
type VideoClassificationConfig struct {
	// Required. The list of annotation spec set configs.
	// Since watching a video clip takes much longer time than an image, we
	// support label with multiple AnnotationSpecSet at the same time. Labels
	// in each AnnotationSpecSet will be shown in a group to contributors.
	// Contributors can select one or more (depending on whether to allow multi
	// label) from each group.
	AnnotationSpecSetConfigs []*VideoClassificationConfig_AnnotationSpecSetConfig `protobuf:"bytes,1,rep,name=annotation_spec_set_configs,json=annotationSpecSetConfigs,proto3" json:"annotation_spec_set_configs,omitempty"`
	// Optional. Option to apply shot detection on the video.
	ApplyShotDetection   bool     `protobuf:"varint,2,opt,name=apply_shot_detection,json=applyShotDetection,proto3" json:"apply_shot_detection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoClassificationConfig) Reset()         { *m = VideoClassificationConfig{} }
func (m *VideoClassificationConfig) String() string { return proto.CompactTextString(m) }
func (*VideoClassificationConfig) ProtoMessage()    {}
func (*VideoClassificationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{5}
}

func (m *VideoClassificationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoClassificationConfig.Unmarshal(m, b)
}
func (m *VideoClassificationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoClassificationConfig.Marshal(b, m, deterministic)
}
func (m *VideoClassificationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoClassificationConfig.Merge(m, src)
}
func (m *VideoClassificationConfig) XXX_Size() int {
	return xxx_messageInfo_VideoClassificationConfig.Size(m)
}
func (m *VideoClassificationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoClassificationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VideoClassificationConfig proto.InternalMessageInfo

func (m *VideoClassificationConfig) GetAnnotationSpecSetConfigs() []*VideoClassificationConfig_AnnotationSpecSetConfig {
	if m != nil {
		return m.AnnotationSpecSetConfigs
	}
	return nil
}

func (m *VideoClassificationConfig) GetApplyShotDetection() bool {
	if m != nil {
		return m.ApplyShotDetection
	}
	return false
}

// Annotation spec set with the setting of allowing multi labels or not.
type VideoClassificationConfig_AnnotationSpecSetConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet string `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Optional. If allow_multi_label is true, contributors are able to
	// choose multiple labels from one annotation spec set.
	AllowMultiLabel      bool     `protobuf:"varint,2,opt,name=allow_multi_label,json=allowMultiLabel,proto3" json:"allow_multi_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoClassificationConfig_AnnotationSpecSetConfig) Reset() {
	*m = VideoClassificationConfig_AnnotationSpecSetConfig{}
}
func (m *VideoClassificationConfig_AnnotationSpecSetConfig) String() string {
	return proto.CompactTextString(m)
}
func (*VideoClassificationConfig_AnnotationSpecSetConfig) ProtoMessage() {}
func (*VideoClassificationConfig_AnnotationSpecSetConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{5, 0}
}

func (m *VideoClassificationConfig_AnnotationSpecSetConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoClassificationConfig_AnnotationSpecSetConfig.Unmarshal(m, b)
}
func (m *VideoClassificationConfig_AnnotationSpecSetConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoClassificationConfig_AnnotationSpecSetConfig.Marshal(b, m, deterministic)
}
func (m *VideoClassificationConfig_AnnotationSpecSetConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoClassificationConfig_AnnotationSpecSetConfig.Merge(m, src)
}
func (m *VideoClassificationConfig_AnnotationSpecSetConfig) XXX_Size() int {
	return xxx_messageInfo_VideoClassificationConfig_AnnotationSpecSetConfig.Size(m)
}
func (m *VideoClassificationConfig_AnnotationSpecSetConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoClassificationConfig_AnnotationSpecSetConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VideoClassificationConfig_AnnotationSpecSetConfig proto.InternalMessageInfo

func (m *VideoClassificationConfig_AnnotationSpecSetConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *VideoClassificationConfig_AnnotationSpecSetConfig) GetAllowMultiLabel() bool {
	if m != nil {
		return m.AllowMultiLabel
	}
	return false
}

// Config for video object detection human labeling task.
// Object detection will be conducted on the images extracted from the video,
// and those objects will be labeled with bounding boxes.
// User need to specify the number of images to be extracted per second as the
// extraction frame rate.
type ObjectDetectionConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet string `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Optional. Instruction message showed on labelers UI.
	InstructionMessage string `protobuf:"bytes,2,opt,name=instruction_message,json=instructionMessage,proto3" json:"instruction_message,omitempty"`
	// Required. Number of frames per second to be extracted from the video.
	ExtractionFrameRate  float64  `protobuf:"fixed64,3,opt,name=extraction_frame_rate,json=extractionFrameRate,proto3" json:"extraction_frame_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectDetectionConfig) Reset()         { *m = ObjectDetectionConfig{} }
func (m *ObjectDetectionConfig) String() string { return proto.CompactTextString(m) }
func (*ObjectDetectionConfig) ProtoMessage()    {}
func (*ObjectDetectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{6}
}

func (m *ObjectDetectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectDetectionConfig.Unmarshal(m, b)
}
func (m *ObjectDetectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectDetectionConfig.Marshal(b, m, deterministic)
}
func (m *ObjectDetectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectDetectionConfig.Merge(m, src)
}
func (m *ObjectDetectionConfig) XXX_Size() int {
	return xxx_messageInfo_ObjectDetectionConfig.Size(m)
}
func (m *ObjectDetectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectDetectionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectDetectionConfig proto.InternalMessageInfo

func (m *ObjectDetectionConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *ObjectDetectionConfig) GetInstructionMessage() string {
	if m != nil {
		return m.InstructionMessage
	}
	return ""
}

func (m *ObjectDetectionConfig) GetExtractionFrameRate() float64 {
	if m != nil {
		return m.ExtractionFrameRate
	}
	return 0
}

// Config for video object tracking human labeling task.
type ObjectTrackingConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet    string   `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectTrackingConfig) Reset()         { *m = ObjectTrackingConfig{} }
func (m *ObjectTrackingConfig) String() string { return proto.CompactTextString(m) }
func (*ObjectTrackingConfig) ProtoMessage()    {}
func (*ObjectTrackingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{7}
}

func (m *ObjectTrackingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectTrackingConfig.Unmarshal(m, b)
}
func (m *ObjectTrackingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectTrackingConfig.Marshal(b, m, deterministic)
}
func (m *ObjectTrackingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectTrackingConfig.Merge(m, src)
}
func (m *ObjectTrackingConfig) XXX_Size() int {
	return xxx_messageInfo_ObjectTrackingConfig.Size(m)
}
func (m *ObjectTrackingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectTrackingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectTrackingConfig proto.InternalMessageInfo

func (m *ObjectTrackingConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

// Config for video event human labeling task.
type EventConfig struct {
	// Required. The list of annotation spec set resource name. Similar to video
	// classification, we support selecting event from multiple AnnotationSpecSet
	// at the same time.
	AnnotationSpecSets   []string `protobuf:"bytes,1,rep,name=annotation_spec_sets,json=annotationSpecSets,proto3" json:"annotation_spec_sets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventConfig) Reset()         { *m = EventConfig{} }
func (m *EventConfig) String() string { return proto.CompactTextString(m) }
func (*EventConfig) ProtoMessage()    {}
func (*EventConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{8}
}

func (m *EventConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventConfig.Unmarshal(m, b)
}
func (m *EventConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventConfig.Marshal(b, m, deterministic)
}
func (m *EventConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventConfig.Merge(m, src)
}
func (m *EventConfig) XXX_Size() int {
	return xxx_messageInfo_EventConfig.Size(m)
}
func (m *EventConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EventConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EventConfig proto.InternalMessageInfo

func (m *EventConfig) GetAnnotationSpecSets() []string {
	if m != nil {
		return m.AnnotationSpecSets
	}
	return nil
}

// Config for text classification human labeling task.
type TextClassificationConfig struct {
	// Optional. If allow_multi_label is true, contributors are able to choose
	// multiple labels for one text segment.
	AllowMultiLabel bool `protobuf:"varint,1,opt,name=allow_multi_label,json=allowMultiLabel,proto3" json:"allow_multi_label,omitempty"`
	// Required. Annotation spec set resource name.
	AnnotationSpecSet string `protobuf:"bytes,2,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// Optional. Configs for sentiment selection.
	SentimentConfig      *SentimentConfig `protobuf:"bytes,3,opt,name=sentiment_config,json=sentimentConfig,proto3" json:"sentiment_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TextClassificationConfig) Reset()         { *m = TextClassificationConfig{} }
func (m *TextClassificationConfig) String() string { return proto.CompactTextString(m) }
func (*TextClassificationConfig) ProtoMessage()    {}
func (*TextClassificationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{9}
}

func (m *TextClassificationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextClassificationConfig.Unmarshal(m, b)
}
func (m *TextClassificationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextClassificationConfig.Marshal(b, m, deterministic)
}
func (m *TextClassificationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextClassificationConfig.Merge(m, src)
}
func (m *TextClassificationConfig) XXX_Size() int {
	return xxx_messageInfo_TextClassificationConfig.Size(m)
}
func (m *TextClassificationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TextClassificationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TextClassificationConfig proto.InternalMessageInfo

func (m *TextClassificationConfig) GetAllowMultiLabel() bool {
	if m != nil {
		return m.AllowMultiLabel
	}
	return false
}

func (m *TextClassificationConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *TextClassificationConfig) GetSentimentConfig() *SentimentConfig {
	if m != nil {
		return m.SentimentConfig
	}
	return nil
}

// Config for setting up sentiments.
type SentimentConfig struct {
	// If set to true, contributors will have the option to select sentiment of
	// the label they selected, to mark it as negative or positive label. Default
	// is false.
	EnableLabelSentimentSelection bool     `protobuf:"varint,1,opt,name=enable_label_sentiment_selection,json=enableLabelSentimentSelection,proto3" json:"enable_label_sentiment_selection,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *SentimentConfig) Reset()         { *m = SentimentConfig{} }
func (m *SentimentConfig) String() string { return proto.CompactTextString(m) }
func (*SentimentConfig) ProtoMessage()    {}
func (*SentimentConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{10}
}

func (m *SentimentConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SentimentConfig.Unmarshal(m, b)
}
func (m *SentimentConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SentimentConfig.Marshal(b, m, deterministic)
}
func (m *SentimentConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SentimentConfig.Merge(m, src)
}
func (m *SentimentConfig) XXX_Size() int {
	return xxx_messageInfo_SentimentConfig.Size(m)
}
func (m *SentimentConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SentimentConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SentimentConfig proto.InternalMessageInfo

func (m *SentimentConfig) GetEnableLabelSentimentSelection() bool {
	if m != nil {
		return m.EnableLabelSentimentSelection
	}
	return false
}

// Config for text entity extraction human labeling task.
type TextEntityExtractionConfig struct {
	// Required. Annotation spec set resource name.
	AnnotationSpecSet    string   `protobuf:"bytes,1,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextEntityExtractionConfig) Reset()         { *m = TextEntityExtractionConfig{} }
func (m *TextEntityExtractionConfig) String() string { return proto.CompactTextString(m) }
func (*TextEntityExtractionConfig) ProtoMessage()    {}
func (*TextEntityExtractionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_331725facbee63fc, []int{11}
}

func (m *TextEntityExtractionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextEntityExtractionConfig.Unmarshal(m, b)
}
func (m *TextEntityExtractionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextEntityExtractionConfig.Marshal(b, m, deterministic)
}
func (m *TextEntityExtractionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextEntityExtractionConfig.Merge(m, src)
}
func (m *TextEntityExtractionConfig) XXX_Size() int {
	return xxx_messageInfo_TextEntityExtractionConfig.Size(m)
}
func (m *TextEntityExtractionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TextEntityExtractionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TextEntityExtractionConfig proto.InternalMessageInfo

func (m *TextEntityExtractionConfig) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.datalabeling.v1beta1.StringAggregationType", StringAggregationType_name, StringAggregationType_value)
	proto.RegisterType((*HumanAnnotationConfig)(nil), "google.cloud.datalabeling.v1beta1.HumanAnnotationConfig")
	proto.RegisterType((*ImageClassificationConfig)(nil), "google.cloud.datalabeling.v1beta1.ImageClassificationConfig")
	proto.RegisterType((*BoundingPolyConfig)(nil), "google.cloud.datalabeling.v1beta1.BoundingPolyConfig")
	proto.RegisterType((*PolylineConfig)(nil), "google.cloud.datalabeling.v1beta1.PolylineConfig")
	proto.RegisterType((*SegmentationConfig)(nil), "google.cloud.datalabeling.v1beta1.SegmentationConfig")
	proto.RegisterType((*VideoClassificationConfig)(nil), "google.cloud.datalabeling.v1beta1.VideoClassificationConfig")
	proto.RegisterType((*VideoClassificationConfig_AnnotationSpecSetConfig)(nil), "google.cloud.datalabeling.v1beta1.VideoClassificationConfig.AnnotationSpecSetConfig")
	proto.RegisterType((*ObjectDetectionConfig)(nil), "google.cloud.datalabeling.v1beta1.ObjectDetectionConfig")
	proto.RegisterType((*ObjectTrackingConfig)(nil), "google.cloud.datalabeling.v1beta1.ObjectTrackingConfig")
	proto.RegisterType((*EventConfig)(nil), "google.cloud.datalabeling.v1beta1.EventConfig")
	proto.RegisterType((*TextClassificationConfig)(nil), "google.cloud.datalabeling.v1beta1.TextClassificationConfig")
	proto.RegisterType((*SentimentConfig)(nil), "google.cloud.datalabeling.v1beta1.SentimentConfig")
	proto.RegisterType((*TextEntityExtractionConfig)(nil), "google.cloud.datalabeling.v1beta1.TextEntityExtractionConfig")
}

func init() {
	proto.RegisterFile("google/cloud/datalabeling/v1beta1/human_annotation_config.proto", fileDescriptor_331725facbee63fc)
}

var fileDescriptor_331725facbee63fc = []byte{
	// 912 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xd1, 0x6e, 0xdb, 0x36,
	0x14, 0x9d, 0xe2, 0xb5, 0x6b, 0xe8, 0x36, 0xb1, 0x99, 0x1a, 0x55, 0xd2, 0xb5, 0xf3, 0x54, 0x0c,
	0x33, 0x0a, 0x4c, 0x5e, 0xbd, 0x97, 0x01, 0x7b, 0x28, 0x1c, 0xdb, 0xf1, 0x3c, 0xd4, 0x76, 0x20,
	0x3b, 0x05, 0x5a, 0x60, 0x20, 0x68, 0xf9, 0x46, 0xe1, 0x26, 0x91, 0xaa, 0x48, 0xb5, 0x31, 0xf6,
	0x19, 0xfb, 0x85, 0x3d, 0xee, 0x77, 0xf6, 0xb2, 0x3f, 0xd8, 0x5f, 0x0c, 0x22, 0xa5, 0xd8, 0x4b,
	0xed, 0xb5, 0xcb, 0xd0, 0xbc, 0x9e, 0x73, 0xef, 0xb9, 0x97, 0xf7, 0x5c, 0x91, 0x42, 0x4f, 0x03,
	0x21, 0x82, 0x10, 0x9a, 0x7e, 0x28, 0xd2, 0x79, 0x73, 0x4e, 0x15, 0x0d, 0xe9, 0x0c, 0x42, 0xc6,
	0x83, 0xe6, 0xeb, 0x27, 0x33, 0x50, 0xf4, 0x49, 0xf3, 0x2c, 0x8d, 0x28, 0x27, 0x94, 0x73, 0xa1,
	0xa8, 0x62, 0x82, 0x13, 0x5f, 0xf0, 0x53, 0x16, 0xb8, 0x71, 0x22, 0x94, 0xc0, 0x9f, 0x1b, 0x01,
	0x57, 0x0b, 0xb8, 0xab, 0x02, 0x6e, 0x2e, 0x70, 0xf0, 0x30, 0xaf, 0xa1, 0x13, 0x66, 0xe9, 0x69,
	0x73, 0x9e, 0x26, 0x5a, 0xc9, 0x48, 0x1c, 0x7c, 0x9a, 0xf3, 0x34, 0x66, 0xcd, 0x65, 0x19, 0x69,
	0x58, 0xe7, 0xb7, 0x12, 0xaa, 0x7d, 0x9f, 0xb5, 0xd0, 0xbe, 0xa0, 0x3a, 0xba, 0x01, 0x5c, 0x47,
	0x65, 0xc6, 0xa5, 0x4a, 0x52, 0x3f, 0x03, 0x6d, 0xab, 0x6e, 0x35, 0xb6, 0xbd, 0x55, 0x08, 0x77,
	0xd0, 0xc3, 0x5c, 0x10, 0xe6, 0x24, 0xeb, 0x4d, 0x82, 0x22, 0x73, 0x26, 0xe3, 0x90, 0x2e, 0x08,
	0xa7, 0x11, 0xd8, 0x5b, 0x3a, 0xe9, 0xfe, 0x45, 0x54, 0xd7, 0x04, 0x75, 0x4d, 0xcc, 0x88, 0x46,
	0x80, 0x0f, 0xd1, 0x83, 0x35, 0x22, 0x20, 0xfd, 0x84, 0xc5, 0xba, 0x70, 0x69, 0x83, 0xc6, 0x32,
	0x04, 0x7f, 0x86, 0xca, 0x7a, 0x2c, 0x24, 0x48, 0x44, 0x1a, 0xdb, 0x1f, 0xeb, 0x0c, 0xa4, 0xa1,
	0x7e, 0x86, 0xe0, 0x47, 0xe8, 0x4e, 0x48, 0x79, 0x90, 0xd2, 0x00, 0x88, 0x2f, 0xe6, 0x60, 0xdf,
	0xd0, 0x21, 0xb7, 0x0b, 0xb0, 0x23, 0xe6, 0x90, 0x05, 0x25, 0x10, 0x87, 0xcc, 0xa7, 0xc4, 0x17,
	0x29, 0x57, 0xf6, 0xcd, 0xba, 0xd5, 0xb8, 0xe1, 0xdd, 0xce, 0xc1, 0x4e, 0x86, 0xe1, 0x23, 0x54,
	0x7d, 0x95, 0x82, 0xd4, 0x4e, 0x15, 0x83, 0xb6, 0x3f, 0xa9, 0x5b, 0x8d, 0x72, 0x6b, 0xdf, 0xcd,
	0xcd, 0x2a, 0x9c, 0x70, 0xbb, 0x79, 0x80, 0x57, 0x29, 0x72, 0x0a, 0x04, 0x7f, 0x85, 0xb0, 0x2f,
	0xb8, 0x4a, 0xd8, 0x2c, 0x55, 0x22, 0x21, 0x10, 0x51, 0x16, 0x4a, 0x7b, 0xbb, 0x5e, 0x6a, 0x6c,
	0x7b, 0xd5, 0x15, 0xa6, 0xa7, 0x09, 0xe7, 0x2f, 0x0b, 0xed, 0x0f, 0xa2, 0xac, 0xd3, 0x90, 0x4a,
	0xc9, 0x4e, 0x99, 0xbf, 0x6a, 0x95, 0x8b, 0xf6, 0x56, 0x16, 0x48, 0xc6, 0xe0, 0x13, 0x09, 0x2a,
	0xb7, 0xac, 0xba, 0xa4, 0x26, 0x31, 0xf8, 0x13, 0x50, 0xf8, 0x31, 0xaa, 0xd2, 0x30, 0x14, 0x6f,
	0x48, 0x94, 0x86, 0x8a, 0x11, 0x3d, 0x28, 0xed, 0xd5, 0x2d, 0x6f, 0x57, 0x13, 0xc3, 0x0c, 0x7f,
	0x96, 0xc1, 0x38, 0x46, 0xf7, 0x28, 0x97, 0x6f, 0x20, 0x21, 0x34, 0x08, 0x12, 0x08, 0x4c, 0x0d,
	0xb5, 0x88, 0x41, 0x3b, 0xb3, 0xd3, 0xfa, 0xd6, 0x7d, 0xe7, 0x8e, 0xba, 0x13, 0x95, 0x30, 0x1e,
	0xb4, 0x97, 0x02, 0xd3, 0x45, 0x0c, 0x5e, 0xcd, 0x08, 0x5f, 0x82, 0x9d, 0x14, 0xe1, 0x43, 0x91,
	0xf2, 0x39, 0xe3, 0xc1, 0xb1, 0x08, 0x17, 0x57, 0x3c, 0x63, 0x13, 0xed, 0xad, 0xec, 0x2a, 0x89,
	0x40, 0x4a, 0x1a, 0x14, 0x1b, 0x89, 0x57, 0xa8, 0xa1, 0x61, 0x9c, 0x57, 0x68, 0x27, 0x2b, 0x17,
	0x32, 0x0e, 0xd7, 0x55, 0x32, 0x45, 0x78, 0x02, 0x41, 0x04, 0x5c, 0xfd, 0x1f, 0x37, 0xff, 0x73,
	0xd9, 0x3f, 0xb7, 0xd0, 0xfe, 0x73, 0x36, 0x07, 0xb1, 0x76, 0x99, 0x7e, 0xb5, 0xd0, 0xfd, 0x35,
	0xf5, 0xf3, 0x7b, 0x49, 0xda, 0x56, 0xbd, 0xd4, 0x28, 0xb7, 0xa6, 0xef, 0xe1, 0xfa, 0xc6, 0x1a,
	0x6e, 0xfb, 0xf2, 0x21, 0x0c, 0xee, 0xd9, 0x74, 0x3d, 0x21, 0xf1, 0xd7, 0xe8, 0x2e, 0x8d, 0xe3,
	0x70, 0x41, 0xe4, 0x99, 0xc8, 0xee, 0x07, 0x05, 0xe6, 0x5a, 0x32, 0x5b, 0x8b, 0x35, 0x37, 0x39,
	0x13, 0xaa, 0x5b, 0x30, 0x07, 0x29, 0xba, 0xb7, 0xa1, 0xcc, 0x87, 0xfc, 0x5e, 0x9c, 0xdf, 0x2d,
	0x54, 0x1b, 0xcf, 0x7e, 0x02, 0x7f, 0xd9, 0xca, 0x35, 0xf9, 0x8a, 0x5b, 0xa8, 0x06, 0xe7, 0x2a,
	0xa1, 0x26, 0xfe, 0x34, 0xa1, 0x11, 0x90, 0x84, 0x2a, 0xf3, 0xa1, 0x5a, 0xde, 0xde, 0x92, 0x3c,
	0xca, 0x38, 0x8f, 0x2a, 0x70, 0x8e, 0xd0, 0x5d, 0xd3, 0xed, 0x34, 0xa1, 0xfe, 0xcf, 0x8c, 0x07,
	0x57, 0x6b, 0xd6, 0x79, 0x8a, 0xca, 0xbd, 0xd7, 0xc0, 0x8b, 0x09, 0x67, 0x76, 0xbd, 0x9d, 0x6e,
	0x96, 0x67, 0xdb, 0xc3, 0x6f, 0xe5, 0x4b, 0xe7, 0x0f, 0x0b, 0xd9, 0x53, 0x38, 0x57, 0x6b, 0x77,
	0x72, 0xad, 0x01, 0xd6, 0xfa, 0x0b, 0x6b, 0x43, 0xe7, 0x5b, 0x9b, 0xc6, 0xfc, 0x23, 0xaa, 0x48,
	0xe0, 0x8a, 0x65, 0x9f, 0x61, 0xbe, 0xe4, 0x7a, 0x60, 0xe5, 0x56, 0xeb, 0x7d, 0x6e, 0xb6, 0x22,
	0x35, 0xdf, 0xe0, 0x5d, 0xf9, 0x4f, 0xc0, 0x79, 0x89, 0x76, 0x2f, 0xc5, 0xe0, 0x3e, 0xaa, 0x03,
	0xa7, 0xb3, 0x10, 0xcc, 0x41, 0xc8, 0xb2, 0xbc, 0x84, 0x10, 0x96, 0xcf, 0xed, 0x2d, 0xef, 0x81,
	0x89, 0xd3, 0x07, 0xbb, 0x50, 0x99, 0x14, 0x41, 0xce, 0x33, 0x74, 0x90, 0x8d, 0xac, 0xc7, 0x15,
	0x53, 0x8b, 0xde, 0x85, 0xbb, 0x57, 0xb3, 0xf0, 0xf1, 0x2f, 0xa8, 0xb6, 0xf6, 0x9e, 0xc6, 0x5f,
	0xa2, 0x47, 0x93, 0xa9, 0x37, 0x18, 0xf5, 0x49, 0xbb, 0xdf, 0xf7, 0x7a, 0xfd, 0xf6, 0x74, 0x30,
	0x1e, 0x91, 0xe9, 0x8b, 0xe3, 0x1e, 0x39, 0x19, 0x4d, 0x8e, 0x7b, 0x9d, 0xc1, 0xd1, 0xa0, 0xd7,
	0xad, 0x7c, 0x84, 0xab, 0xe8, 0xce, 0xb0, 0xfd, 0xc3, 0xd8, 0x1b, 0x4c, 0x5f, 0x90, 0xe7, 0xe3,
	0x69, 0xaf, 0x62, 0x61, 0x8c, 0x76, 0x4e, 0x46, 0xed, 0xd1, 0x60, 0x38, 0x3e, 0x99, 0x18, 0x6c,
	0x2b, 0xc3, 0x46, 0xe3, 0x55, 0xad, 0x4a, 0xe9, 0xf0, 0x1c, 0x7d, 0xe1, 0x8b, 0xe8, 0xdd, 0x03,
	0x3f, 0xb6, 0x5e, 0x0e, 0xf3, 0xa0, 0x40, 0x64, 0xcf, 0xb7, 0x2b, 0x92, 0xa0, 0x19, 0x00, 0xd7,
	0x8f, 0x6e, 0xd3, 0x50, 0x34, 0x66, 0xf2, 0x5f, 0xfe, 0xb9, 0xbe, 0x5b, 0x05, 0x67, 0x37, 0x75,
	0xe6, 0x37, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x56, 0x6a, 0x3b, 0xa6, 0xac, 0x09, 0x00, 0x00,
}
