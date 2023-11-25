// Code generated by protoc-gen-go. DO NOT EDIT.
// source: survey.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	math "math"
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

type Uint64ToString struct {
	Key                  uint64   `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uint64ToString) Reset()         { *m = Uint64ToString{} }
func (m *Uint64ToString) String() string { return proto.CompactTextString(m) }
func (*Uint64ToString) ProtoMessage()    {}
func (*Uint64ToString) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{0}
}

func (m *Uint64ToString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uint64ToString.Unmarshal(m, b)
}
func (m *Uint64ToString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uint64ToString.Marshal(b, m, deterministic)
}
func (m *Uint64ToString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uint64ToString.Merge(m, src)
}
func (m *Uint64ToString) XXX_Size() int {
	return xxx_messageInfo_Uint64ToString.Size(m)
}
func (m *Uint64ToString) XXX_DiscardUnknown() {
	xxx_messageInfo_Uint64ToString.DiscardUnknown(m)
}

var xxx_messageInfo_Uint64ToString proto.InternalMessageInfo

func (m *Uint64ToString) GetKey() uint64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Uint64ToString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type StringToUint64 struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Key                  uint64   `protobuf:"varint,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringToUint64) Reset()         { *m = StringToUint64{} }
func (m *StringToUint64) String() string { return proto.CompactTextString(m) }
func (*StringToUint64) ProtoMessage()    {}
func (*StringToUint64) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{1}
}

func (m *StringToUint64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringToUint64.Unmarshal(m, b)
}
func (m *StringToUint64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringToUint64.Marshal(b, m, deterministic)
}
func (m *StringToUint64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringToUint64.Merge(m, src)
}
func (m *StringToUint64) XXX_Size() int {
	return xxx_messageInfo_StringToUint64.Size(m)
}
func (m *StringToUint64) XXX_DiscardUnknown() {
	xxx_messageInfo_StringToUint64.DiscardUnknown(m)
}

var xxx_messageInfo_StringToUint64 proto.InternalMessageInfo

func (m *StringToUint64) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *StringToUint64) GetKey() uint64 {
	if m != nil {
		return m.Key
	}
	return 0
}

type Survey struct {
	UserSurvey           *UserSurvey       `protobuf:"bytes,1,opt,name=userSurvey,proto3" json:"userSurvey,omitempty"`
	Answers              []*Uint64ToString `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Survey) Reset()         { *m = Survey{} }
func (m *Survey) String() string { return proto.CompactTextString(m) }
func (*Survey) ProtoMessage()    {}
func (*Survey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{2}
}

func (m *Survey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Survey.Unmarshal(m, b)
}
func (m *Survey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Survey.Marshal(b, m, deterministic)
}
func (m *Survey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Survey.Merge(m, src)
}
func (m *Survey) XXX_Size() int {
	return xxx_messageInfo_Survey.Size(m)
}
func (m *Survey) XXX_DiscardUnknown() {
	xxx_messageInfo_Survey.DiscardUnknown(m)
}

var xxx_messageInfo_Survey proto.InternalMessageInfo

func (m *Survey) GetUserSurvey() *UserSurvey {
	if m != nil {
		return m.UserSurvey
	}
	return nil
}

func (m *Survey) GetAnswers() []*Uint64ToString {
	if m != nil {
		return m.Answers
	}
	return nil
}

type UserSurvey struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	SurveyId             uint64   `protobuf:"varint,2,opt,name=surveyId,proto3" json:"surveyId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSurvey) Reset()         { *m = UserSurvey{} }
func (m *UserSurvey) String() string { return proto.CompactTextString(m) }
func (*UserSurvey) ProtoMessage()    {}
func (*UserSurvey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{3}
}

func (m *UserSurvey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSurvey.Unmarshal(m, b)
}
func (m *UserSurvey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSurvey.Marshal(b, m, deterministic)
}
func (m *UserSurvey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSurvey.Merge(m, src)
}
func (m *UserSurvey) XXX_Size() int {
	return xxx_messageInfo_UserSurvey.Size(m)
}
func (m *UserSurvey) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSurvey.DiscardUnknown(m)
}

var xxx_messageInfo_UserSurvey proto.InternalMessageInfo

func (m *UserSurvey) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserSurvey) GetSurveyId() uint64 {
	if m != nil {
		return m.SurveyId
	}
	return 0
}

type SurveyId struct {
	SurveyId             uint64   `protobuf:"varint,1,opt,name=surveyId,proto3" json:"surveyId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SurveyId) Reset()         { *m = SurveyId{} }
func (m *SurveyId) String() string { return proto.CompactTextString(m) }
func (*SurveyId) ProtoMessage()    {}
func (*SurveyId) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{4}
}

func (m *SurveyId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyId.Unmarshal(m, b)
}
func (m *SurveyId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyId.Marshal(b, m, deterministic)
}
func (m *SurveyId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyId.Merge(m, src)
}
func (m *SurveyId) XXX_Size() int {
	return xxx_messageInfo_SurveyId.Size(m)
}
func (m *SurveyId) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyId.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyId proto.InternalMessageInfo

func (m *SurveyId) GetSurveyId() uint64 {
	if m != nil {
		return m.SurveyId
	}
	return 0
}

type IsOk struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsOk) Reset()         { *m = IsOk{} }
func (m *IsOk) String() string { return proto.CompactTextString(m) }
func (*IsOk) ProtoMessage()    {}
func (*IsOk) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{5}
}

func (m *IsOk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsOk.Unmarshal(m, b)
}
func (m *IsOk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsOk.Marshal(b, m, deterministic)
}
func (m *IsOk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsOk.Merge(m, src)
}
func (m *IsOk) XXX_Size() int {
	return xxx_messageInfo_IsOk.Size(m)
}
func (m *IsOk) XXX_DiscardUnknown() {
	xxx_messageInfo_IsOk.DiscardUnknown(m)
}

var xxx_messageInfo_IsOk proto.InternalMessageInfo

func (m *IsOk) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type StatResponse struct {
	SurveyId             uint64            `protobuf:"varint,1,opt,name=surveyId,proto3" json:"surveyId,omitempty"`
	QuestionToAverage    []*StringToUint64 `protobuf:"bytes,2,rep,name=questionToAverage,proto3" json:"questionToAverage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StatResponse) Reset()         { *m = StatResponse{} }
func (m *StatResponse) String() string { return proto.CompactTextString(m) }
func (*StatResponse) ProtoMessage()    {}
func (*StatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{6}
}

func (m *StatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatResponse.Unmarshal(m, b)
}
func (m *StatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatResponse.Marshal(b, m, deterministic)
}
func (m *StatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatResponse.Merge(m, src)
}
func (m *StatResponse) XXX_Size() int {
	return xxx_messageInfo_StatResponse.Size(m)
}
func (m *StatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatResponse proto.InternalMessageInfo

func (m *StatResponse) GetSurveyId() uint64 {
	if m != nil {
		return m.SurveyId
	}
	return 0
}

func (m *StatResponse) GetQuestionToAverage() []*StringToUint64 {
	if m != nil {
		return m.QuestionToAverage
	}
	return nil
}

type StatResponses struct {
	StatResponses        []*StatResponse `protobuf:"bytes,1,rep,name=statResponses,proto3" json:"statResponses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StatResponses) Reset()         { *m = StatResponses{} }
func (m *StatResponses) String() string { return proto.CompactTextString(m) }
func (*StatResponses) ProtoMessage()    {}
func (*StatResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{7}
}

func (m *StatResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatResponses.Unmarshal(m, b)
}
func (m *StatResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatResponses.Marshal(b, m, deterministic)
}
func (m *StatResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatResponses.Merge(m, src)
}
func (m *StatResponses) XXX_Size() int {
	return xxx_messageInfo_StatResponses.Size(m)
}
func (m *StatResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_StatResponses.DiscardUnknown(m)
}

var xxx_messageInfo_StatResponses proto.InternalMessageInfo

func (m *StatResponses) GetStatResponses() []*StatResponse {
	if m != nil {
		return m.StatResponses
	}
	return nil
}

type Response struct {
	SurveyId             uint64            `protobuf:"varint,1,opt,name=surveyId,proto3" json:"surveyId,omitempty"`
	QuestionIdToText     []*Uint64ToString `protobuf:"bytes,2,rep,name=questionIdToText,proto3" json:"questionIdToText,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{8}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSurveyId() uint64 {
	if m != nil {
		return m.SurveyId
	}
	return 0
}

func (m *Response) GetQuestionIdToText() []*Uint64ToString {
	if m != nil {
		return m.QuestionIdToText
	}
	return nil
}

func init() {
	proto.RegisterType((*Uint64ToString)(nil), "Uint64ToString")
	proto.RegisterType((*StringToUint64)(nil), "StringToUint64")
	proto.RegisterType((*Survey)(nil), "Survey")
	proto.RegisterType((*UserSurvey)(nil), "UserSurvey")
	proto.RegisterType((*SurveyId)(nil), "SurveyId")
	proto.RegisterType((*IsOk)(nil), "IsOk")
	proto.RegisterType((*StatResponse)(nil), "StatResponse")
	proto.RegisterType((*StatResponses)(nil), "StatResponses")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() {
	proto.RegisterFile("survey.proto", fileDescriptor_a40f94eaa8e6ca46)
}

var fileDescriptor_a40f94eaa8e6ca46 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0x6e, 0xd2, 0xbb, 0x5e, 0x3a, 0xfd, 0xe1, 0xb9, 0x48, 0x29, 0x11, 0xa1, 0xec, 0x83, 0xd4,
	0x1f, 0x6c, 0xb0, 0x27, 0x2a, 0x88, 0xe0, 0x89, 0x52, 0xf2, 0x24, 0x24, 0xb9, 0x17, 0x9f, 0xcc,
	0xb5, 0x63, 0x59, 0x9a, 0x66, 0xeb, 0xee, 0xa6, 0x7a, 0xff, 0xb2, 0x7f, 0x85, 0x24, 0x9b, 0xe4,
	0x12, 0x8f, 0xca, 0x3d, 0x75, 0x66, 0xe7, 0x9b, 0xaf, 0xdf, 0x7c, 0x33, 0x81, 0xa1, 0xca, 0xe4,
	0x01, 0x6f, 0xd8, 0x5e, 0x0a, 0x2d, 0xdc, 0xc7, 0x1b, 0x21, 0x36, 0x09, 0x7a, 0x45, 0x76, 0x9d,
	0xfd, 0xf0, 0x70, 0xb7, 0xd7, 0x65, 0x91, 0xbe, 0x83, 0xf1, 0x15, 0x4f, 0xf5, 0x9b, 0xd7, 0x91,
	0x08, 0xb5, 0xe4, 0xe9, 0x86, 0x9c, 0x43, 0x77, 0x8b, 0x37, 0x53, 0x6b, 0x66, 0xcd, 0x4f, 0x82,
	0x3c, 0x24, 0x8f, 0xe0, 0xf4, 0x10, 0x27, 0x19, 0x4e, 0xed, 0x99, 0x35, 0xef, 0x07, 0x26, 0xc9,
	0x3b, 0x4d, 0x47, 0x24, 0x0c, 0xc3, 0x2d, 0xce, 0x6a, 0xe0, 0x2a, 0x3e, 0xbb, 0xe6, 0xa3, 0xdf,
	0xa1, 0x17, 0x16, 0x02, 0xc9, 0x0b, 0x80, 0x4c, 0xa1, 0x34, 0x59, 0xd1, 0x36, 0x58, 0x0c, 0xd8,
	0x55, 0xfd, 0x14, 0x34, 0xca, 0xe4, 0x19, 0x9c, 0xc5, 0xa9, 0xfa, 0x85, 0x52, 0x4d, 0xed, 0x59,
	0x77, 0x3e, 0x58, 0x3c, 0x60, 0x6d, 0xe9, 0x41, 0x55, 0xa7, 0x1f, 0x01, 0x6e, 0x49, 0xc8, 0x04,
	0x7a, 0x39, 0x8d, 0xbf, 0x2e, 0x85, 0x95, 0x19, 0x71, 0xc1, 0x31, 0x46, 0xf9, 0xeb, 0x52, 0x5e,
	0x9d, 0xd3, 0xa7, 0xe0, 0x84, 0x65, 0xdc, 0xc2, 0x59, 0xff, 0xe0, 0x26, 0x70, 0xe2, 0xab, 0xaf,
	0x5b, 0x32, 0x06, 0x5b, 0x6c, 0x8b, 0xaa, 0x13, 0xd8, 0x62, 0x4b, 0x39, 0x0c, 0x43, 0x1d, 0xeb,
	0x00, 0xd5, 0x5e, 0xa4, 0x0a, 0xff, 0xc7, 0x41, 0x3e, 0xc0, 0xc3, 0x9f, 0x19, 0x2a, 0xcd, 0x45,
	0x1a, 0x89, 0xcb, 0x03, 0xca, 0x78, 0x83, 0xf5, 0x88, 0x6d, 0x8f, 0x83, 0xbb, 0x48, 0xfa, 0x19,
	0x46, 0xcd, 0xbf, 0x52, 0xe4, 0x02, 0x46, 0xaa, 0xf9, 0x30, 0xb5, 0x0a, 0xae, 0x11, 0x6b, 0xc2,
	0x82, 0x36, 0x86, 0xae, 0xc0, 0xb9, 0x97, 0xd8, 0xf7, 0x70, 0x5e, 0x49, 0xf0, 0xd7, 0x91, 0x88,
	0xf0, 0xb7, 0x3e, 0xb6, 0x8e, 0x3b, 0xc0, 0xc5, 0x1f, 0x0b, 0x46, 0xc6, 0xd6, 0x10, 0xe5, 0x81,
	0xaf, 0x90, 0xbc, 0x82, 0x61, 0x98, 0x5d, 0xef, 0xb8, 0x2e, 0x77, 0x75, 0xc6, 0x4c, 0xe0, 0x4e,
	0x98, 0x39, 0x5b, 0x56, 0x9d, 0x2d, 0xfb, 0x92, 0x9f, 0x2d, 0xed, 0x90, 0x19, 0x38, 0xbe, 0x32,
	0x4d, 0xa4, 0x79, 0x2c, 0xee, 0x29, 0xcb, 0x57, 0x41, 0x3b, 0xe4, 0x25, 0x8c, 0x97, 0x58, 0x32,
	0xe6, 0x33, 0x2b, 0xd2, 0x67, 0xd5, 0x36, 0xdd, 0xb6, 0x0d, 0xb4, 0x43, 0x9e, 0x40, 0x77, 0x89,
	0xba, 0x09, 0xe9, 0xb3, 0x46, 0xf9, 0x2d, 0x0c, 0x96, 0xa8, 0x2f, 0x93, 0xc4, 0x30, 0x1d, 0xd1,
	0xe5, 0x8e, 0x5b, 0xb4, 0x8a, 0x76, 0x3e, 0x3d, 0xff, 0x36, 0xdf, 0xc5, 0x3c, 0xf5, 0x78, 0xaa,
	0x51, 0xa6, 0x71, 0xe2, 0xed, 0xf8, 0x4a, 0x0a, 0x65, 0x06, 0x57, 0x9e, 0x71, 0xb4, 0xfc, 0x2a,
	0x7b, 0xc5, 0xcf, 0xc5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0x42, 0x73, 0xf4, 0xba, 0x03,
	0x00, 0x00,
}