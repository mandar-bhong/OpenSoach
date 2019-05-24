// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/model.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Deployment state of the model.
type Model_DeploymentState int32

const (
	// Should not be used, an un-set enum has this value by default.
	Model_DEPLOYMENT_STATE_UNSPECIFIED Model_DeploymentState = 0
	// Model is deployed.
	Model_DEPLOYED Model_DeploymentState = 1
	// Model is not deployed.
	Model_UNDEPLOYED Model_DeploymentState = 2
)

var Model_DeploymentState_name = map[int32]string{
	0: "DEPLOYMENT_STATE_UNSPECIFIED",
	1: "DEPLOYED",
	2: "UNDEPLOYED",
}

var Model_DeploymentState_value = map[string]int32{
	"DEPLOYMENT_STATE_UNSPECIFIED": 0,
	"DEPLOYED":                     1,
	"UNDEPLOYED":                   2,
}

func (x Model_DeploymentState) String() string {
	return proto.EnumName(Model_DeploymentState_name, int32(x))
}

func (Model_DeploymentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_27c9ed3a1b40360b, []int{0, 0}
}

// API proto representing a trained machine learning model.
type Model struct {
	// Required.
	// The model metadata that is specific to the problem type.
	// Must match the metadata type of the dataset used to train the model.
	//
	// Types that are valid to be assigned to ModelMetadata:
	//	*Model_TranslationModelMetadata
	//	*Model_ImageClassificationModelMetadata
	//	*Model_TextClassificationModelMetadata
	//	*Model_ImageObjectDetectionModelMetadata
	//	*Model_VideoClassificationModelMetadata
	//	*Model_TextExtractionModelMetadata
	//	*Model_TablesModelMetadata
	//	*Model_TextSentimentModelMetadata
	ModelMetadata isModel_ModelMetadata `protobuf_oneof:"model_metadata"`
	// Output only.
	// Resource name of the model.
	// Format: `projects/{project_id}/locations/{location_id}/models/{model_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the model to show in the interface. The name can be
	// up to 32 characters long and can consist only of ASCII Latin letters A-Z
	// and a-z, underscores
	// (_), and ASCII digits 0-9. It must start with a letter.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required.
	// The resource ID of the dataset used to create the model. The dataset must
	// come from the same ancestor project and location.
	DatasetId string `protobuf:"bytes,3,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// Output only.
	// Timestamp when this model was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only.
	// Timestamp when this model was last updated.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,11,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Deployment state of the model. A model can only serve
	// prediction requests after it gets deployed.
	DeploymentState      Model_DeploymentState `protobuf:"varint,8,opt,name=deployment_state,json=deploymentState,proto3,enum=google.cloud.automl.v1beta1.Model_DeploymentState" json:"deployment_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_27c9ed3a1b40360b, []int{0}
}

func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (m *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(m, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

type isModel_ModelMetadata interface {
	isModel_ModelMetadata()
}

type Model_TranslationModelMetadata struct {
	TranslationModelMetadata *TranslationModelMetadata `protobuf:"bytes,15,opt,name=translation_model_metadata,json=translationModelMetadata,proto3,oneof"`
}

type Model_ImageClassificationModelMetadata struct {
	ImageClassificationModelMetadata *ImageClassificationModelMetadata `protobuf:"bytes,13,opt,name=image_classification_model_metadata,json=imageClassificationModelMetadata,proto3,oneof"`
}

type Model_TextClassificationModelMetadata struct {
	TextClassificationModelMetadata *TextClassificationModelMetadata `protobuf:"bytes,14,opt,name=text_classification_model_metadata,json=textClassificationModelMetadata,proto3,oneof"`
}

type Model_ImageObjectDetectionModelMetadata struct {
	ImageObjectDetectionModelMetadata *ImageObjectDetectionModelMetadata `protobuf:"bytes,20,opt,name=image_object_detection_model_metadata,json=imageObjectDetectionModelMetadata,proto3,oneof"`
}

type Model_VideoClassificationModelMetadata struct {
	VideoClassificationModelMetadata *VideoClassificationModelMetadata `protobuf:"bytes,23,opt,name=video_classification_model_metadata,json=videoClassificationModelMetadata,proto3,oneof"`
}

type Model_TextExtractionModelMetadata struct {
	TextExtractionModelMetadata *TextExtractionModelMetadata `protobuf:"bytes,19,opt,name=text_extraction_model_metadata,json=textExtractionModelMetadata,proto3,oneof"`
}

type Model_TablesModelMetadata struct {
	TablesModelMetadata *TablesModelMetadata `protobuf:"bytes,24,opt,name=tables_model_metadata,json=tablesModelMetadata,proto3,oneof"`
}

type Model_TextSentimentModelMetadata struct {
	TextSentimentModelMetadata *TextSentimentModelMetadata `protobuf:"bytes,22,opt,name=text_sentiment_model_metadata,json=textSentimentModelMetadata,proto3,oneof"`
}

func (*Model_TranslationModelMetadata) isModel_ModelMetadata() {}

func (*Model_ImageClassificationModelMetadata) isModel_ModelMetadata() {}

func (*Model_TextClassificationModelMetadata) isModel_ModelMetadata() {}

func (*Model_ImageObjectDetectionModelMetadata) isModel_ModelMetadata() {}

func (*Model_VideoClassificationModelMetadata) isModel_ModelMetadata() {}

func (*Model_TextExtractionModelMetadata) isModel_ModelMetadata() {}

func (*Model_TablesModelMetadata) isModel_ModelMetadata() {}

func (*Model_TextSentimentModelMetadata) isModel_ModelMetadata() {}

func (m *Model) GetModelMetadata() isModel_ModelMetadata {
	if m != nil {
		return m.ModelMetadata
	}
	return nil
}

func (m *Model) GetTranslationModelMetadata() *TranslationModelMetadata {
	if x, ok := m.GetModelMetadata().(*Model_TranslationModelMetadata); ok {
		return x.TranslationModelMetadata
	}
	return nil
}

func (m *Model) GetImageClassificationModelMetadata() *ImageClassificationModelMetadata {
	if x, ok := m.GetModelMetadata().(*Model_ImageClassificationModelMetadata); ok {
		return x.ImageClassificationModelMetadata
	}
	return nil
}

func (m *Model) GetTextClassificationModelMetadata() *TextClassificationModelMetadata {
	if x, ok := m.GetModelMetadata().(*Model_TextClassificationModelMetadata); ok {
		return x.TextClassificationModelMetadata
	}
	return nil
}

func (m *Model) GetImageObjectDetectionModelMetadata() *ImageObjectDetectionModelMetadata {
	if x, ok := m.GetModelMetadata().(*Model_ImageObjectDetectionModelMetadata); ok {
		return x.ImageObjectDetectionModelMetadata
	}
	return nil
}

func (m *Model) GetVideoClassificationModelMetadata() *VideoClassificationModelMetadata {
	if x, ok := m.GetModelMetadata().(*Model_VideoClassificationModelMetadata); ok {
		return x.VideoClassificationModelMetadata
	}
	return nil
}

func (m *Model) GetTextExtractionModelMetadata() *TextExtractionModelMetadata {
	if x, ok := m.GetModelMetadata().(*Model_TextExtractionModelMetadata); ok {
		return x.TextExtractionModelMetadata
	}
	return nil
}

func (m *Model) GetTablesModelMetadata() *TablesModelMetadata {
	if x, ok := m.GetModelMetadata().(*Model_TablesModelMetadata); ok {
		return x.TablesModelMetadata
	}
	return nil
}

func (m *Model) GetTextSentimentModelMetadata() *TextSentimentModelMetadata {
	if x, ok := m.GetModelMetadata().(*Model_TextSentimentModelMetadata); ok {
		return x.TextSentimentModelMetadata
	}
	return nil
}

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Model) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *Model) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Model) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Model) GetDeploymentState() Model_DeploymentState {
	if m != nil {
		return m.DeploymentState
	}
	return Model_DEPLOYMENT_STATE_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Model) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Model_TranslationModelMetadata)(nil),
		(*Model_ImageClassificationModelMetadata)(nil),
		(*Model_TextClassificationModelMetadata)(nil),
		(*Model_ImageObjectDetectionModelMetadata)(nil),
		(*Model_VideoClassificationModelMetadata)(nil),
		(*Model_TextExtractionModelMetadata)(nil),
		(*Model_TablesModelMetadata)(nil),
		(*Model_TextSentimentModelMetadata)(nil),
	}
}

func init() {
	proto.RegisterEnum("google.cloud.automl.v1beta1.Model_DeploymentState", Model_DeploymentState_name, Model_DeploymentState_value)
	proto.RegisterType((*Model)(nil), "google.cloud.automl.v1beta1.Model")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/model.proto", fileDescriptor_27c9ed3a1b40360b)
}

var fileDescriptor_27c9ed3a1b40360b = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0x6f, 0x4f, 0xd3, 0x5e,
	0x14, 0xc7, 0x29, 0xbf, 0x9f, 0x0a, 0x77, 0x38, 0x96, 0x8b, 0x7f, 0x9a, 0x01, 0x32, 0x30, 0xea,
	0x9e, 0xd8, 0x0a, 0xc6, 0x68, 0x82, 0x9a, 0x00, 0xab, 0xba, 0xc4, 0x0d, 0xdc, 0x06, 0x89, 0x06,
	0xd3, 0xdc, 0xb5, 0x97, 0xa6, 0xa6, 0xed, 0x6d, 0xd6, 0x33, 0x02, 0x89, 0x89, 0x4f, 0x7c, 0x66,
	0xe2, 0x03, 0x5f, 0x90, 0x2f, 0xc0, 0x57, 0x65, 0x7a, 0x6e, 0x37, 0xa5, 0x6e, 0xb7, 0x3c, 0xa3,
	0xf7, 0x7c, 0xce, 0xf7, 0x7c, 0xef, 0xf7, 0xd0, 0x8e, 0x3c, 0xf0, 0x84, 0xf0, 0x02, 0x6e, 0x3a,
	0x81, 0x18, 0xba, 0x26, 0x1b, 0x82, 0x08, 0x03, 0xf3, 0x74, 0xb3, 0xcf, 0x81, 0x6d, 0x9a, 0xa1,
	0x70, 0x79, 0x60, 0xc4, 0x03, 0x01, 0x82, 0x2e, 0x4b, 0xd0, 0x40, 0xd0, 0x90, 0xa0, 0x91, 0x81,
	0xd5, 0x95, 0x4c, 0x85, 0xc5, 0xbe, 0xc9, 0xa2, 0x48, 0x00, 0x03, 0x5f, 0x44, 0x89, 0x6c, 0xad,
	0x2a, 0x67, 0xf8, 0x21, 0xf3, 0x78, 0x06, 0xd6, 0x55, 0x20, 0xb0, 0x7e, 0xc0, 0x47, 0x92, 0xf7,
	0x95, 0x24, 0x3f, 0x83, 0x8c, 0x7b, 0xa8, 0xe4, 0x06, 0x2c, 0x4a, 0x02, 0xb4, 0x7a, 0x19, 0xa7,
	0xa7, 0xbe, 0xcb, 0x45, 0x06, 0xae, 0x65, 0x20, 0x3e, 0xf5, 0x87, 0x27, 0x26, 0xf8, 0x21, 0x4f,
	0x80, 0x85, 0xb1, 0x04, 0x36, 0x7e, 0x96, 0xc8, 0x95, 0x56, 0x1a, 0x1f, 0x1d, 0x92, 0xea, 0x5f,
	0x83, 0x6c, 0xcc, 0xd4, 0x0e, 0x39, 0x30, 0x97, 0x01, 0xd3, 0x17, 0x6b, 0x5a, 0xbd, 0xb4, 0xf5,
	0xc4, 0x50, 0xa4, 0x6b, 0xf4, 0xfe, 0xb4, 0xa3, 0x64, 0x2b, 0x6b, 0x7e, 0x33, 0xd3, 0xd1, 0x61,
	0x4a, 0x8d, 0x7e, 0xd7, 0xc8, 0x5d, 0xcc, 0xd6, 0x76, 0x02, 0x96, 0x24, 0xfe, 0x89, 0xef, 0x4c,
	0x34, 0x70, 0x1d, 0x0d, 0xbc, 0x50, 0x1a, 0x68, 0xa6, 0x3a, 0x7b, 0x17, 0x64, 0xf2, 0x46, 0x6a,
	0x7e, 0x01, 0x43, 0xbf, 0x69, 0x64, 0x23, 0xdd, 0x4c, 0x81, 0x9f, 0x32, 0xfa, 0x79, 0xae, 0x0e,
	0x84, 0x9f, 0x81, 0xda, 0xce, 0x1a, 0xa8, 0x11, 0xfa, 0x43, 0x23, 0xf7, 0x64, 0x3c, 0xa2, 0xff,
	0x89, 0x3b, 0x60, 0xbb, 0x1c, 0xb8, 0x33, 0xc9, 0xd0, 0x0d, 0x34, 0xf4, 0xb2, 0x38, 0xa0, 0x7d,
	0x14, 0x6a, 0x8c, 0x74, 0xf2, 0x96, 0xd6, 0xfd, 0x22, 0x08, 0x77, 0x86, 0xff, 0x65, 0x05, 0x19,
	0xdd, 0xbe, 0xc4, 0xce, 0x8e, 0x52, 0x9d, 0x82, 0x9d, 0x9d, 0x16, 0x30, 0xf4, 0x0b, 0xb9, 0x83,
	0x2b, 0xe3, 0x67, 0x30, 0x60, 0x13, 0xd3, 0x59, 0x42, 0x2b, 0xcf, 0x0a, 0xd7, 0x65, 0x8d, 0x15,
	0xf2, 0x2e, 0x96, 0x61, 0x7a, 0x99, 0x9e, 0x90, 0x9b, 0xf2, 0xbd, 0xcf, 0xcf, 0xd5, 0x71, 0xee,
	0x23, 0xf5, 0x5c, 0xec, 0xcc, 0xcf, 0x5b, 0x82, 0x7f, 0x8f, 0xe9, 0x67, 0xb2, 0x8a, 0x17, 0x4d,
	0x78, 0x94, 0xbe, 0xc9, 0x11, 0xe4, 0xe7, 0xdd, 0xc2, 0x79, 0x4f, 0x0b, 0xef, 0xd9, 0x1d, 0x09,
	0xe4, 0xc7, 0x56, 0x61, 0x6a, 0x95, 0x52, 0xf2, 0x7f, 0xc4, 0x42, 0xae, 0x6b, 0x35, 0xad, 0x3e,
	0xdf, 0xc1, 0xbf, 0xe9, 0x3a, 0x59, 0x70, 0xfd, 0x24, 0x0e, 0xd8, 0xb9, 0x8d, 0xb5, 0x59, 0xac,
	0x95, 0xb2, 0xb3, 0x76, 0x8a, 0xac, 0x12, 0x92, 0xb6, 0x27, 0x1c, 0x6c, 0xdf, 0xd5, 0xff, 0x43,
	0x60, 0x3e, 0x3b, 0x69, 0xba, 0x74, 0x9b, 0x94, 0x9c, 0x01, 0x67, 0xc0, 0xed, 0x74, 0xa6, 0x7e,
	0x0d, 0x6f, 0x50, 0x1d, 0xdd, 0x60, 0xf4, 0xe5, 0x32, 0x7a, 0xa3, 0x2f, 0x57, 0x87, 0x48, 0x3c,
	0x3d, 0x48, 0x9b, 0x87, 0xb1, 0x3b, 0x6e, 0x2e, 0x15, 0x37, 0x4b, 0x1c, 0x9b, 0x3f, 0x92, 0x8a,
	0xcb, 0xe3, 0x40, 0x9c, 0x63, 0x92, 0x09, 0x30, 0xe0, 0xfa, 0x5c, 0x4d, 0xab, 0x97, 0xb7, 0xb6,
	0x94, 0x01, 0x62, 0x2a, 0x46, 0x63, 0xdc, 0xda, 0x4d, 0x3b, 0x3b, 0x8b, 0xee, 0xc5, 0x83, 0x8d,
	0x77, 0x64, 0x31, 0xc7, 0xd0, 0x1a, 0x59, 0x69, 0x58, 0x07, 0x6f, 0xf7, 0xdf, 0xb7, 0xac, 0x76,
	0xcf, 0xee, 0xf6, 0x76, 0x7a, 0x96, 0x7d, 0xd8, 0xee, 0x1e, 0x58, 0x7b, 0xcd, 0x57, 0x4d, 0xab,
	0x51, 0x99, 0xa1, 0x0b, 0x64, 0x4e, 0x12, 0x56, 0xa3, 0xa2, 0xd1, 0x32, 0x21, 0x87, 0xed, 0xf1,
	0xf3, 0xec, 0x6e, 0x85, 0x94, 0x2f, 0x2e, 0x7c, 0xf7, 0xab, 0x46, 0xd6, 0x1c, 0x11, 0xaa, 0xfc,
	0x1e, 0x68, 0x1f, 0x76, 0xb2, 0xb2, 0x27, 0x02, 0x16, 0x79, 0x86, 0x18, 0x78, 0xa6, 0xc7, 0x23,
	0x8c, 0xc7, 0x94, 0x25, 0x16, 0xfb, 0xc9, 0xc4, 0xdf, 0x93, 0x6d, 0xf9, 0xf8, 0x6b, 0x76, 0xf9,
	0x35, 0x82, 0xc7, 0x7b, 0x29, 0x74, 0xbc, 0x33, 0x04, 0xd1, 0x0a, 0x8e, 0x8f, 0x24, 0xd4, 0xbf,
	0x8a, 0x5a, 0x8f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xed, 0x1f, 0xbf, 0xcc, 0xa8, 0x07, 0x00,
	0x00,
}
