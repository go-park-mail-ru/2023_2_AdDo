// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: survey.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Uint64ToString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   uint64 `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Uint64ToString) Reset() {
	*x = Uint64ToString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uint64ToString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint64ToString) ProtoMessage() {}

func (x *Uint64ToString) ProtoReflect() protoreflect.Message {
	mi := &file_survey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uint64ToString.ProtoReflect.Descriptor instead.
func (*Uint64ToString) Descriptor() ([]byte, []int) {
	return file_survey_proto_rawDescGZIP(), []int{0}
}

func (x *Uint64ToString) GetKey() uint64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *Uint64ToString) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type StringToUint64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Key   uint64 `protobuf:"varint,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *StringToUint64) Reset() {
	*x = StringToUint64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringToUint64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringToUint64) ProtoMessage() {}

func (x *StringToUint64) ProtoReflect() protoreflect.Message {
	mi := &file_survey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringToUint64.ProtoReflect.Descriptor instead.
func (*StringToUint64) Descriptor() ([]byte, []int) {
	return file_survey_proto_rawDescGZIP(), []int{1}
}

func (x *StringToUint64) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *StringToUint64) GetKey() uint64 {
	if x != nil {
		return x.Key
	}
	return 0
}

type Survey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Survey  *UserSurvey       `protobuf:"bytes,1,opt,name=survey,proto3" json:"survey,omitempty"`
	Answers []*Uint64ToString `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *Survey) Reset() {
	*x = Survey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Survey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Survey) ProtoMessage() {}

func (x *Survey) ProtoReflect() protoreflect.Message {
	mi := &file_survey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Survey.ProtoReflect.Descriptor instead.
func (*Survey) Descriptor() ([]byte, []int) {
	return file_survey_proto_rawDescGZIP(), []int{2}
}

func (x *Survey) GetSurvey() *UserSurvey {
	if x != nil {
		return x.Survey
	}
	return nil
}

func (x *Survey) GetAnswers() []*Uint64ToString {
	if x != nil {
		return x.Answers
	}
	return nil
}

type UserSurvey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	SurveyId uint64 `protobuf:"varint,2,opt,name=surveyId,proto3" json:"surveyId,omitempty"`
}

func (x *UserSurvey) Reset() {
	*x = UserSurvey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSurvey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSurvey) ProtoMessage() {}

func (x *UserSurvey) ProtoReflect() protoreflect.Message {
	mi := &file_survey_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSurvey.ProtoReflect.Descriptor instead.
func (*UserSurvey) Descriptor() ([]byte, []int) {
	return file_survey_proto_rawDescGZIP(), []int{3}
}

func (x *UserSurvey) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserSurvey) GetSurveyId() uint64 {
	if x != nil {
		return x.SurveyId
	}
	return 0
}

type SurveyId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyId string `protobuf:"bytes,1,opt,name=surveyId,proto3" json:"surveyId,omitempty"`
}

func (x *SurveyId) Reset() {
	*x = SurveyId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyId) ProtoMessage() {}

func (x *SurveyId) ProtoReflect() protoreflect.Message {
	mi := &file_survey_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyId.ProtoReflect.Descriptor instead.
func (*SurveyId) Descriptor() ([]byte, []int) {
	return file_survey_proto_rawDescGZIP(), []int{4}
}

func (x *SurveyId) GetSurveyId() string {
	if x != nil {
		return x.SurveyId
	}
	return ""
}

type IsOk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *IsOk) Reset() {
	*x = IsOk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsOk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsOk) ProtoMessage() {}

func (x *IsOk) ProtoReflect() protoreflect.Message {
	mi := &file_survey_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsOk.ProtoReflect.Descriptor instead.
func (*IsOk) Descriptor() ([]byte, []int) {
	return file_survey_proto_rawDescGZIP(), []int{5}
}

func (x *IsOk) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type StatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyId          uint64            `protobuf:"varint,1,opt,name=surveyId,proto3" json:"surveyId,omitempty"`
	QuestionToAverage []*StringToUint64 `protobuf:"bytes,2,rep,name=questionToAverage,proto3" json:"questionToAverage,omitempty"`
}

func (x *StatResponse) Reset() {
	*x = StatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatResponse) ProtoMessage() {}

func (x *StatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_survey_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatResponse.ProtoReflect.Descriptor instead.
func (*StatResponse) Descriptor() ([]byte, []int) {
	return file_survey_proto_rawDescGZIP(), []int{6}
}

func (x *StatResponse) GetSurveyId() uint64 {
	if x != nil {
		return x.SurveyId
	}
	return 0
}

func (x *StatResponse) GetQuestionToAverage() []*StringToUint64 {
	if x != nil {
		return x.QuestionToAverage
	}
	return nil
}

type StatResponses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatResponses []*StatResponse `protobuf:"bytes,1,rep,name=statResponses,proto3" json:"statResponses,omitempty"`
}

func (x *StatResponses) Reset() {
	*x = StatResponses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatResponses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatResponses) ProtoMessage() {}

func (x *StatResponses) ProtoReflect() protoreflect.Message {
	mi := &file_survey_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatResponses.ProtoReflect.Descriptor instead.
func (*StatResponses) Descriptor() ([]byte, []int) {
	return file_survey_proto_rawDescGZIP(), []int{7}
}

func (x *StatResponses) GetStatResponses() []*StatResponse {
	if x != nil {
		return x.StatResponses
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyId         uint64            `protobuf:"varint,1,opt,name=surveyId,proto3" json:"surveyId,omitempty"`
	QuestionIdToText []*Uint64ToString `protobuf:"bytes,2,rep,name=questionIdToText,proto3" json:"questionIdToText,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_survey_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_survey_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_survey_proto_rawDescGZIP(), []int{8}
}

func (x *Response) GetSurveyId() uint64 {
	if x != nil {
		return x.SurveyId
	}
	return 0
}

func (x *Response) GetQuestionIdToText() []*Uint64ToString {
	if x != nil {
		return x.QuestionIdToText
	}
	return nil
}

var File_survey_proto protoreflect.FileDescriptor

var file_survey_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0e, 0x55,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54,
	0x6f, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x58, 0x0a, 0x06, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x06, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x12, 0x29,
	0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22, 0x40, 0x0a, 0x0a, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x08, 0x53,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x04, 0x49, 0x73, 0x4f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x69, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x11, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x55, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x52, 0x11, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x44, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0d, 0x73,
	0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x10, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x54, 0x6f, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x10, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x54, 0x6f, 0x54, 0x65, 0x78,
	0x74, 0x32, 0xea, 0x01, 0x0a, 0x0d, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x12, 0x07, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x08, 0x49, 0x73, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x12, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x1a,
	0x05, 0x2e, 0x49, 0x73, 0x4f, 0x6b, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x09, 0x2e, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x1d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x09, 0x2e,
	0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x00, 0x42, 0x2a,
	0x5a, 0x28, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_survey_proto_rawDescOnce sync.Once
	file_survey_proto_rawDescData = file_survey_proto_rawDesc
)

func file_survey_proto_rawDescGZIP() []byte {
	file_survey_proto_rawDescOnce.Do(func() {
		file_survey_proto_rawDescData = protoimpl.X.CompressGZIP(file_survey_proto_rawDescData)
	})
	return file_survey_proto_rawDescData
}

var file_survey_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_survey_proto_goTypes = []interface{}{
	(*Uint64ToString)(nil), // 0: Uint64ToString
	(*StringToUint64)(nil), // 1: StringToUint64
	(*Survey)(nil),         // 2: Survey
	(*UserSurvey)(nil),     // 3: UserSurvey
	(*SurveyId)(nil),       // 4: SurveyId
	(*IsOk)(nil),           // 5: IsOk
	(*StatResponse)(nil),   // 6: StatResponse
	(*StatResponses)(nil),  // 7: StatResponses
	(*Response)(nil),       // 8: Response
	(*empty.Empty)(nil),    // 9: google.protobuf.Empty
}
var file_survey_proto_depIdxs = []int32{
	3,  // 0: Survey.survey:type_name -> UserSurvey
	0,  // 1: Survey.answers:type_name -> Uint64ToString
	1,  // 2: StatResponse.questionToAverage:type_name -> StringToUint64
	6,  // 3: StatResponses.statResponses:type_name -> StatResponse
	0,  // 4: Response.questionIdToText:type_name -> Uint64ToString
	2,  // 5: SurveyService.SubmitSurvey:input_type -> Survey
	3,  // 6: SurveyService.IsSubmit:input_type -> UserSurvey
	4,  // 7: SurveyService.GetSurveyStats:input_type -> SurveyId
	4,  // 8: SurveyService.Get:input_type -> SurveyId
	9,  // 9: SurveyService.GetAllStats:input_type -> google.protobuf.Empty
	9,  // 10: SurveyService.SubmitSurvey:output_type -> google.protobuf.Empty
	5,  // 11: SurveyService.IsSubmit:output_type -> IsOk
	6,  // 12: SurveyService.GetSurveyStats:output_type -> StatResponse
	8,  // 13: SurveyService.Get:output_type -> Response
	7,  // 14: SurveyService.GetAllStats:output_type -> StatResponses
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_survey_proto_init() }
func file_survey_proto_init() {
	if File_survey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_survey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uint64ToString); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_survey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringToUint64); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_survey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Survey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_survey_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSurvey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_survey_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_survey_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsOk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_survey_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_survey_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatResponses); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_survey_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_survey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_survey_proto_goTypes,
		DependencyIndexes: file_survey_proto_depIdxs,
		MessageInfos:      file_survey_proto_msgTypes,
	}.Build()
	File_survey_proto = out.File
	file_survey_proto_rawDesc = nil
	file_survey_proto_goTypes = nil
	file_survey_proto_depIdxs = nil
}
