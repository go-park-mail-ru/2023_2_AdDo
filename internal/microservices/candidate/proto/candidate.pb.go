// Code generated by protoc-gen-go. DO NOT EDIT.
// source: candidate.proto

package candidate

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "main/internal/microservices/session/proto"
	proto1 "main/internal/microservices/track/proto"
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

type Candidates struct {
	Tracks               *proto1.TracksResponse `protobuf:"bytes,1,opt,name=Tracks,proto3" json:"Tracks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Candidates) Reset()         { *m = Candidates{} }
func (m *Candidates) String() string { return proto.CompactTextString(m) }
func (*Candidates) ProtoMessage()    {}
func (*Candidates) Descriptor() ([]byte, []int) {
	return fileDescriptor_515fb327dd7e8b3d, []int{0}
}

func (m *Candidates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candidates.Unmarshal(m, b)
}
func (m *Candidates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candidates.Marshal(b, m, deterministic)
}
func (m *Candidates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candidates.Merge(m, src)
}
func (m *Candidates) XXX_Size() int {
	return xxx_messageInfo_Candidates.Size(m)
}
func (m *Candidates) XXX_DiscardUnknown() {
	xxx_messageInfo_Candidates.DiscardUnknown(m)
}

var xxx_messageInfo_Candidates proto.InternalMessageInfo

func (m *Candidates) GetTracks() *proto1.TracksResponse {
	if m != nil {
		return m.Tracks
	}
	return nil
}

func init() {
	proto.RegisterType((*Candidates)(nil), "Candidates")
}

func init() {
	proto.RegisterFile("candidate.proto", fileDescriptor_515fb327dd7e8b3d)
}

var fileDescriptor_515fb327dd7e8b3d = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4e, 0xcc, 0x4b,
	0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x2e, 0x29, 0x4a,
	0x4c, 0xce, 0x86, 0x72, 0x78, 0x8b, 0x53, 0x8b, 0x8b, 0x33, 0xf3, 0xf3, 0x20, 0x5c, 0x25, 0x53,
	0x2e, 0x2e, 0x67, 0x98, 0xf2, 0x62, 0x21, 0x75, 0x2e, 0xb6, 0x10, 0x90, 0xda, 0x62, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x6e, 0x23, 0x7e, 0x3d, 0x08, 0x37, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0x35, 0x08, 0x2a, 0x6d, 0xe4, 0xc4, 0x25, 0x00, 0xd7, 0x16, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c,
	0x2a, 0xa4, 0xc7, 0x25, 0xe2, 0x9e, 0x5a, 0x82, 0x30, 0xcd, 0x2d, 0xbf, 0x28, 0xb4, 0x38, 0xb5,
	0x48, 0x88, 0x5d, 0x0f, 0x44, 0x79, 0xa6, 0x48, 0x71, 0xeb, 0x21, 0x24, 0x95, 0x18, 0x9c, 0x78,
	0xa2, 0xb8, 0xf4, 0xe1, 0x4e, 0x4d, 0x62, 0x03, 0xbb, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xde, 0x73, 0xf7, 0x79, 0xbe, 0x00, 0x00, 0x00,
}