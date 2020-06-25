// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/annotation_payload.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
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

// Contains annotation information that is relevant to AutoML.
type AnnotationPayload struct {
	// Output only . Additional information about the annotation
	// specific to the AutoML domain.
	//
	// Types that are valid to be assigned to Detail:
	//	*AnnotationPayload_Translation
	//	*AnnotationPayload_Classification
	//	*AnnotationPayload_ImageObjectDetection
	//	*AnnotationPayload_TextExtraction
	//	*AnnotationPayload_TextSentiment
	Detail isAnnotationPayload_Detail `protobuf_oneof:"detail"`
	// Output only . The resource ID of the annotation spec that
	// this annotation pertains to. The annotation spec comes from either an
	// ancestor dataset, or the dataset that was used to train the model in use.
	AnnotationSpecId string `protobuf:"bytes,1,opt,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
	// Output only. The value of
	// [display_name][google.cloud.automl.v1.AnnotationSpec.display_name]
	// when the model was trained. Because this field returns a value at model
	// training time, for different models trained using the same dataset, the
	// returned value could be different as model owner could update the
	// `display_name` between any two model training.
	DisplayName          string   `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnnotationPayload) Reset()         { *m = AnnotationPayload{} }
func (m *AnnotationPayload) String() string { return proto.CompactTextString(m) }
func (*AnnotationPayload) ProtoMessage()    {}
func (*AnnotationPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_a945962355943c79, []int{0}
}

func (m *AnnotationPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotationPayload.Unmarshal(m, b)
}
func (m *AnnotationPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotationPayload.Marshal(b, m, deterministic)
}
func (m *AnnotationPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotationPayload.Merge(m, src)
}
func (m *AnnotationPayload) XXX_Size() int {
	return xxx_messageInfo_AnnotationPayload.Size(m)
}
func (m *AnnotationPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotationPayload.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotationPayload proto.InternalMessageInfo

type isAnnotationPayload_Detail interface {
	isAnnotationPayload_Detail()
}

type AnnotationPayload_Translation struct {
	Translation *TranslationAnnotation `protobuf:"bytes,2,opt,name=translation,proto3,oneof"`
}

type AnnotationPayload_Classification struct {
	Classification *ClassificationAnnotation `protobuf:"bytes,3,opt,name=classification,proto3,oneof"`
}

type AnnotationPayload_ImageObjectDetection struct {
	ImageObjectDetection *ImageObjectDetectionAnnotation `protobuf:"bytes,4,opt,name=image_object_detection,json=imageObjectDetection,proto3,oneof"`
}

type AnnotationPayload_TextExtraction struct {
	TextExtraction *TextExtractionAnnotation `protobuf:"bytes,6,opt,name=text_extraction,json=textExtraction,proto3,oneof"`
}

type AnnotationPayload_TextSentiment struct {
	TextSentiment *TextSentimentAnnotation `protobuf:"bytes,7,opt,name=text_sentiment,json=textSentiment,proto3,oneof"`
}

func (*AnnotationPayload_Translation) isAnnotationPayload_Detail() {}

func (*AnnotationPayload_Classification) isAnnotationPayload_Detail() {}

func (*AnnotationPayload_ImageObjectDetection) isAnnotationPayload_Detail() {}

func (*AnnotationPayload_TextExtraction) isAnnotationPayload_Detail() {}

func (*AnnotationPayload_TextSentiment) isAnnotationPayload_Detail() {}

func (m *AnnotationPayload) GetDetail() isAnnotationPayload_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (m *AnnotationPayload) GetTranslation() *TranslationAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_Translation); ok {
		return x.Translation
	}
	return nil
}

func (m *AnnotationPayload) GetClassification() *ClassificationAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_Classification); ok {
		return x.Classification
	}
	return nil
}

func (m *AnnotationPayload) GetImageObjectDetection() *ImageObjectDetectionAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_ImageObjectDetection); ok {
		return x.ImageObjectDetection
	}
	return nil
}

func (m *AnnotationPayload) GetTextExtraction() *TextExtractionAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_TextExtraction); ok {
		return x.TextExtraction
	}
	return nil
}

func (m *AnnotationPayload) GetTextSentiment() *TextSentimentAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_TextSentiment); ok {
		return x.TextSentiment
	}
	return nil
}

func (m *AnnotationPayload) GetAnnotationSpecId() string {
	if m != nil {
		return m.AnnotationSpecId
	}
	return ""
}

func (m *AnnotationPayload) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AnnotationPayload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AnnotationPayload_Translation)(nil),
		(*AnnotationPayload_Classification)(nil),
		(*AnnotationPayload_ImageObjectDetection)(nil),
		(*AnnotationPayload_TextExtraction)(nil),
		(*AnnotationPayload_TextSentiment)(nil),
	}
}

func init() {
	proto.RegisterType((*AnnotationPayload)(nil), "google.cloud.automl.v1.AnnotationPayload")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1/annotation_payload.proto", fileDescriptor_a945962355943c79)
}

var fileDescriptor_a945962355943c79 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0x19, 0x05, 0x5c, 0x18, 0x60, 0xa1, 0x2a, 0xab, 0xb8, 0x18, 0x5c, 0xa0, 0x4a,
	0x1b, 0x09, 0x01, 0x89, 0x8b, 0xc0, 0xcd, 0x36, 0x10, 0x4c, 0xe2, 0xcf, 0xd8, 0xa6, 0x0a, 0x8d,
	0x4a, 0xd1, 0x69, 0x72, 0x16, 0x19, 0x39, 0x76, 0xd4, 0x9c, 0x4e, 0xed, 0x2b, 0xed, 0x39, 0xb8,
	0xe2, 0x51, 0x78, 0x0a, 0x84, 0x9d, 0xb6, 0x49, 0xd5, 0xf4, 0x2e, 0xf1, 0xf7, 0xfb, 0xbe, 0x63,
	0x1f, 0x1f, 0x33, 0x3f, 0xd5, 0x3a, 0x95, 0xe8, 0xc7, 0x52, 0x4f, 0x12, 0x1f, 0x26, 0xa4, 0x33,
	0xe9, 0x5f, 0x05, 0x3e, 0x28, 0xa5, 0x09, 0x48, 0x68, 0x15, 0xe5, 0x30, 0x93, 0x1a, 0x12, 0x2f,
	0x1f, 0x6b, 0xd2, 0xbc, 0x6b, 0x0d, 0x9e, 0x31, 0x78, 0xd6, 0xe0, 0x5d, 0x05, 0xbd, 0xbd, 0x86,
	0xa0, 0x58, 0x42, 0x51, 0x88, 0x4b, 0x11, 0x9b, 0x30, 0x1b, 0xd2, 0x7b, 0xde, 0x00, 0x27, 0x48,
	0x18, 0x57, 0xb8, 0xfd, 0x06, 0x8e, 0x70, 0x4a, 0x11, 0x4e, 0x69, 0x0c, 0x55, 0x7a, 0x6f, 0x13,
	0x5d, 0xa0, 0x22, 0x91, 0xa1, 0xa2, 0x12, 0xee, 0x37, 0xc1, 0x63, 0x50, 0x85, 0xac, 0x6e, 0x76,
	0xa7, 0x24, 0xcd, 0xdf, 0x68, 0x72, 0xe9, 0x83, 0x9a, 0x95, 0xd2, 0x93, 0x52, 0x82, 0x5c, 0x54,
	0x3a, 0x56, 0x58, 0xf5, 0xd9, 0xef, 0x2d, 0xf6, 0xe8, 0x60, 0xb1, 0x7a, 0x62, 0xdb, 0xc8, 0xbf,
	0xb3, 0x4e, 0xa5, 0x86, 0xdb, 0xda, 0x75, 0xfa, 0x9d, 0x57, 0x2f, 0xbc, 0xf5, 0x6d, 0xf5, 0xce,
	0x97, 0xe8, 0x32, 0xea, 0xd3, 0x8d, 0xd3, 0x6a, 0x06, 0xbf, 0x60, 0xdb, 0xf5, 0x36, 0xbb, 0x37,
	0x4d, 0xea, 0xcb, 0xa6, 0xd4, 0xa3, 0x1a, 0x5d, 0x0b, 0x5e, 0x49, 0xe2, 0x8a, 0x75, 0x45, 0x06,
	0x29, 0x46, 0x7a, 0xf4, 0x0b, 0x63, 0x8a, 0x16, 0x57, 0xe4, 0x6e, 0x99, 0x1a, 0x6f, 0x9a, 0x6a,
	0x1c, 0xff, 0x77, 0x7d, 0x33, 0xa6, 0xf7, 0x73, 0x4f, 0xad, 0xd2, 0x63, 0xb1, 0x86, 0xe0, 0x3f,
	0xd9, 0x83, 0x95, 0xdb, 0x75, 0xdb, 0x9b, 0x0f, 0x73, 0x8e, 0x53, 0xfa, 0xb0, 0xa0, 0xeb, 0x87,
	0xa1, 0x9a, 0xc6, 0x7f, 0xb0, 0xed, 0xfa, 0x30, 0xb8, 0xb7, 0x4d, 0xb6, 0xbf, 0x29, 0xfb, 0x6c,
	0x0e, 0xd7, 0xa2, 0xef, 0x53, 0x55, 0xe2, 0xfb, 0x8c, 0x57, 0x9e, 0x4c, 0x91, 0x63, 0x1c, 0x89,
	0xc4, 0x75, 0x76, 0x9d, 0xfe, 0xdd, 0xd3, 0x87, 0x4b, 0xe5, 0x2c, 0xc7, 0xf8, 0x38, 0xe1, 0x4f,
	0xd9, 0xbd, 0x44, 0x14, 0xb9, 0x84, 0x59, 0xa4, 0x20, 0x43, 0xf7, 0x96, 0xe1, 0x3a, 0xe5, 0xda,
	0x57, 0xc8, 0xf0, 0xf0, 0x0e, 0x6b, 0x27, 0x48, 0x20, 0xe4, 0xe1, 0xb5, 0xc3, 0x7a, 0xb1, 0xce,
	0x1a, 0xb6, 0x78, 0xe2, 0x5c, 0xbc, 0x2b, 0x95, 0x54, 0x4b, 0x50, 0xa9, 0xa7, 0xc7, 0xa9, 0x9f,
	0xa2, 0x32, 0x33, 0x58, 0x3e, 0x6f, 0xc8, 0x45, 0xb1, 0x3a, 0xe9, 0x6f, 0xed, 0xd7, 0x75, 0xab,
	0xfb, 0xd1, 0xda, 0x8f, 0x4c, 0xf0, 0xc1, 0x84, 0xf4, 0x97, 0xcf, 0xde, 0x20, 0xf8, 0x33, 0x17,
	0x86, 0x46, 0x18, 0x5a, 0x61, 0x38, 0x08, 0xfe, 0xb6, 0x76, 0xac, 0x10, 0x86, 0x46, 0x09, 0x43,
	0x2b, 0x85, 0xe1, 0x20, 0x18, 0xb5, 0x4d, 0xd9, 0xd7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcd,
	0x2e, 0x32, 0x1c, 0x58, 0x04, 0x00, 0x00,
}