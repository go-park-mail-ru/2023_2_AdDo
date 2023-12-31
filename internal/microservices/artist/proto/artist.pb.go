// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artist.proto

package artist

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	proto2 "main/internal/microservices/album/proto"
	proto3 "main/internal/microservices/playlist/proto"
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

type ArtistBase struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtistBase) Reset()         { *m = ArtistBase{} }
func (m *ArtistBase) String() string { return proto.CompactTextString(m) }
func (*ArtistBase) ProtoMessage()    {}
func (*ArtistBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_43defc8f563de921, []int{0}
}

func (m *ArtistBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtistBase.Unmarshal(m, b)
}
func (m *ArtistBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtistBase.Marshal(b, m, deterministic)
}
func (m *ArtistBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtistBase.Merge(m, src)
}
func (m *ArtistBase) XXX_Size() int {
	return xxx_messageInfo_ArtistBase.Size(m)
}
func (m *ArtistBase) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtistBase.DiscardUnknown(m)
}

var xxx_messageInfo_ArtistBase proto.InternalMessageInfo

func (m *ArtistBase) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArtistBase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtistBase) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type ArtistsBase struct {
	Artists              []*ArtistBase `protobuf:"bytes,1,rep,name=Artists,proto3" json:"Artists,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ArtistsBase) Reset()         { *m = ArtistsBase{} }
func (m *ArtistsBase) String() string { return proto.CompactTextString(m) }
func (*ArtistsBase) ProtoMessage()    {}
func (*ArtistsBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_43defc8f563de921, []int{1}
}

func (m *ArtistsBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtistsBase.Unmarshal(m, b)
}
func (m *ArtistsBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtistsBase.Marshal(b, m, deterministic)
}
func (m *ArtistsBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtistsBase.Merge(m, src)
}
func (m *ArtistsBase) XXX_Size() int {
	return xxx_messageInfo_ArtistsBase.Size(m)
}
func (m *ArtistsBase) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtistsBase.DiscardUnknown(m)
}

var xxx_messageInfo_ArtistsBase proto.InternalMessageInfo

func (m *ArtistsBase) GetArtists() []*ArtistBase {
	if m != nil {
		return m.Artists
	}
	return nil
}

type SearchResponse struct {
	Tracks               *proto1.TracksResponse `protobuf:"bytes,1,opt,name=Tracks,proto3" json:"Tracks,omitempty"`
	Albums               *proto2.AlbumsBase     `protobuf:"bytes,2,opt,name=Albums,proto3" json:"Albums,omitempty"`
	Playlists            *proto3.PlaylistsBase  `protobuf:"bytes,3,opt,name=Playlists,proto3" json:"Playlists,omitempty"`
	Artists              *ArtistsBase           `protobuf:"bytes,4,opt,name=Artists,proto3" json:"Artists,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43defc8f563de921, []int{2}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetTracks() *proto1.TracksResponse {
	if m != nil {
		return m.Tracks
	}
	return nil
}

func (m *SearchResponse) GetAlbums() *proto2.AlbumsBase {
	if m != nil {
		return m.Albums
	}
	return nil
}

func (m *SearchResponse) GetPlaylists() *proto3.PlaylistsBase {
	if m != nil {
		return m.Playlists
	}
	return nil
}

func (m *SearchResponse) GetArtists() *ArtistsBase {
	if m != nil {
		return m.Artists
	}
	return nil
}

type Query struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_43defc8f563de921, []int{3}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type IsLikedArtist struct {
	IsLiked              bool     `protobuf:"varint,1,opt,name=IsLiked,proto3" json:"IsLiked,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsLikedArtist) Reset()         { *m = IsLikedArtist{} }
func (m *IsLikedArtist) String() string { return proto.CompactTextString(m) }
func (*IsLikedArtist) ProtoMessage()    {}
func (*IsLikedArtist) Descriptor() ([]byte, []int) {
	return fileDescriptor_43defc8f563de921, []int{4}
}

func (m *IsLikedArtist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsLikedArtist.Unmarshal(m, b)
}
func (m *IsLikedArtist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsLikedArtist.Marshal(b, m, deterministic)
}
func (m *IsLikedArtist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsLikedArtist.Merge(m, src)
}
func (m *IsLikedArtist) XXX_Size() int {
	return xxx_messageInfo_IsLikedArtist.Size(m)
}
func (m *IsLikedArtist) XXX_DiscardUnknown() {
	xxx_messageInfo_IsLikedArtist.DiscardUnknown(m)
}

var xxx_messageInfo_IsLikedArtist proto.InternalMessageInfo

func (m *IsLikedArtist) GetIsLiked() bool {
	if m != nil {
		return m.IsLiked
	}
	return false
}

type ArtistToUserId struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ArtistId             uint64   `protobuf:"varint,2,opt,name=artistId,proto3" json:"artistId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtistToUserId) Reset()         { *m = ArtistToUserId{} }
func (m *ArtistToUserId) String() string { return proto.CompactTextString(m) }
func (*ArtistToUserId) ProtoMessage()    {}
func (*ArtistToUserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_43defc8f563de921, []int{5}
}

func (m *ArtistToUserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtistToUserId.Unmarshal(m, b)
}
func (m *ArtistToUserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtistToUserId.Marshal(b, m, deterministic)
}
func (m *ArtistToUserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtistToUserId.Merge(m, src)
}
func (m *ArtistToUserId) XXX_Size() int {
	return xxx_messageInfo_ArtistToUserId.Size(m)
}
func (m *ArtistToUserId) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtistToUserId.DiscardUnknown(m)
}

var xxx_messageInfo_ArtistToUserId proto.InternalMessageInfo

func (m *ArtistToUserId) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ArtistToUserId) GetArtistId() uint64 {
	if m != nil {
		return m.ArtistId
	}
	return 0
}

type ArtistId struct {
	ArtistId             uint64   `protobuf:"varint,1,opt,name=artistId,proto3" json:"artistId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtistId) Reset()         { *m = ArtistId{} }
func (m *ArtistId) String() string { return proto.CompactTextString(m) }
func (*ArtistId) ProtoMessage()    {}
func (*ArtistId) Descriptor() ([]byte, []int) {
	return fileDescriptor_43defc8f563de921, []int{6}
}

func (m *ArtistId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtistId.Unmarshal(m, b)
}
func (m *ArtistId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtistId.Marshal(b, m, deterministic)
}
func (m *ArtistId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtistId.Merge(m, src)
}
func (m *ArtistId) XXX_Size() int {
	return xxx_messageInfo_ArtistId.Size(m)
}
func (m *ArtistId) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtistId.DiscardUnknown(m)
}

var xxx_messageInfo_ArtistId proto.InternalMessageInfo

func (m *ArtistId) GetArtistId() uint64 {
	if m != nil {
		return m.ArtistId
	}
	return 0
}

type Artist struct {
	Id                   uint64                 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Avatar               string                 `protobuf:"bytes,3,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Albums               *proto2.AlbumsBase     `protobuf:"bytes,4,opt,name=Albums,proto3" json:"Albums,omitempty"`
	Tracks               *proto1.TracksResponse `protobuf:"bytes,5,opt,name=Tracks,proto3" json:"Tracks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Artist) Reset()         { *m = Artist{} }
func (m *Artist) String() string { return proto.CompactTextString(m) }
func (*Artist) ProtoMessage()    {}
func (*Artist) Descriptor() ([]byte, []int) {
	return fileDescriptor_43defc8f563de921, []int{7}
}

func (m *Artist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artist.Unmarshal(m, b)
}
func (m *Artist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artist.Marshal(b, m, deterministic)
}
func (m *Artist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artist.Merge(m, src)
}
func (m *Artist) XXX_Size() int {
	return xxx_messageInfo_Artist.Size(m)
}
func (m *Artist) XXX_DiscardUnknown() {
	xxx_messageInfo_Artist.DiscardUnknown(m)
}

var xxx_messageInfo_Artist proto.InternalMessageInfo

func (m *Artist) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Artist) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artist) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Artist) GetAlbums() *proto2.AlbumsBase {
	if m != nil {
		return m.Albums
	}
	return nil
}

func (m *Artist) GetTracks() *proto1.TracksResponse {
	if m != nil {
		return m.Tracks
	}
	return nil
}

func init() {
	proto.RegisterType((*ArtistBase)(nil), "ArtistBase")
	proto.RegisterType((*ArtistsBase)(nil), "ArtistsBase")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
	proto.RegisterType((*Query)(nil), "Query")
	proto.RegisterType((*IsLikedArtist)(nil), "IsLikedArtist")
	proto.RegisterType((*ArtistToUserId)(nil), "ArtistToUserId")
	proto.RegisterType((*ArtistId)(nil), "ArtistId")
	proto.RegisterType((*Artist)(nil), "Artist")
}

func init() {
	proto.RegisterFile("artist.proto", fileDescriptor_43defc8f563de921)
}

var fileDescriptor_43defc8f563de921 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xed, 0x6e, 0xd3, 0x30,
	0x14, 0x6d, 0xba, 0x2c, 0x69, 0x6e, 0xd6, 0x14, 0x59, 0x68, 0x8a, 0x82, 0x90, 0x2a, 0x23, 0xb6,
	0x22, 0x90, 0x2b, 0x3a, 0x5e, 0xa0, 0xe3, 0x33, 0x12, 0x42, 0x90, 0x6d, 0x7f, 0xf8, 0x97, 0x36,
	0xde, 0x88, 0xe6, 0xc6, 0x25, 0x4e, 0x26, 0xf5, 0x35, 0x78, 0x03, 0xde, 0x81, 0x07, 0x44, 0xb9,
	0x76, 0x96, 0x05, 0x04, 0x12, 0xe2, 0xdf, 0x3d, 0xf6, 0xb1, 0xef, 0xb9, 0xe7, 0x1e, 0x38, 0x48,
	0xcb, 0x2a, 0x57, 0x15, 0xdb, 0x96, 0xb2, 0x92, 0x91, 0x9f, 0x8a, 0x55, 0xbd, 0x69, 0x41, 0x55,
	0xa6, 0xeb, 0x6b, 0x03, 0xc6, 0x8a, 0x2b, 0x95, 0xcb, 0xc2, 0xc0, 0x60, 0x2b, 0xd2, 0x9d, 0xe8,
	0x1e, 0x3e, 0xb8, 0x92, 0xf2, 0x4a, 0xf0, 0x39, 0xa2, 0x55, 0x7d, 0x39, 0xe7, 0x9b, 0x6d, 0xb5,
	0xd3, 0x97, 0xf4, 0x1d, 0xc0, 0x12, 0xbb, 0x9c, 0xa6, 0x8a, 0x93, 0x00, 0x86, 0x71, 0x16, 0x5a,
	0x53, 0x6b, 0x66, 0x27, 0xc3, 0x38, 0x23, 0x04, 0xec, 0x0f, 0xe9, 0x86, 0x87, 0xc3, 0xa9, 0x35,
	0xf3, 0x12, 0xac, 0xc9, 0x21, 0x38, 0xcb, 0x9b, 0xb4, 0x4a, 0xcb, 0x70, 0x0f, 0x4f, 0x0d, 0xa2,
	0x2f, 0xc0, 0xd7, 0x3f, 0x29, 0xfc, 0xea, 0x31, 0xb8, 0x06, 0x86, 0xd6, 0x74, 0x6f, 0xe6, 0x2f,
	0x7c, 0xd6, 0x35, 0x4a, 0xda, 0x3b, 0xfa, 0xc3, 0x82, 0xe0, 0x8c, 0xa7, 0xe5, 0xfa, 0x4b, 0xc2,
	0xd5, 0x56, 0x16, 0x8a, 0x93, 0x63, 0x70, 0xce, 0x9b, 0xe9, 0x14, 0x0a, 0xf1, 0x17, 0x13, 0xa6,
	0x61, 0x4b, 0x48, 0xcc, 0x35, 0x79, 0x04, 0xce, 0xb2, 0xf1, 0x44, 0xa1, 0x3e, 0xec, 0x80, 0x10,
	0x3b, 0x98, 0x2b, 0xf2, 0x0c, 0xbc, 0x8f, 0xc6, 0x0f, 0x85, 0x8a, 0xfd, 0x45, 0xc0, 0x6e, 0x4f,
	0x90, 0xda, 0x11, 0xc8, 0x51, 0xa7, 0xda, 0x46, 0xee, 0x01, 0xbb, 0x33, 0x54, 0x27, 0xfb, 0x21,
	0xec, 0x7f, 0xaa, 0x79, 0xb9, 0x23, 0xf7, 0x61, 0xff, 0x6b, 0x53, 0xa0, 0x56, 0x2f, 0xd1, 0x80,
	0x3e, 0x81, 0x71, 0xac, 0xde, 0xe7, 0xd7, 0x3c, 0xd3, 0x0f, 0x48, 0x08, 0xae, 0x39, 0x40, 0xe2,
	0x28, 0x69, 0x21, 0x7d, 0x05, 0x81, 0xe6, 0x9c, 0xcb, 0x0b, 0xc5, 0xcb, 0x38, 0x6b, 0x0c, 0xae,
	0xb1, 0x32, 0x7f, 0x1a, 0x44, 0x22, 0x18, 0xe9, 0x40, 0xc4, 0x19, 0x0e, 0x6c, 0x27, 0xb7, 0x98,
	0x1e, 0xc1, 0x68, 0x69, 0xea, 0x1e, 0xcf, 0xfa, 0x85, 0xf7, 0xcd, 0x02, 0xc7, 0x48, 0xfa, 0x8f,
	0x5d, 0xdf, 0x71, 0xde, 0xfe, 0xb3, 0xf3, 0xdd, 0x1e, 0xf7, 0xff, 0xba, 0xc7, 0xc5, 0xf7, 0x21,
	0x8c, 0xb5, 0xa8, 0x33, 0x5e, 0xde, 0xe4, 0xeb, 0x26, 0x3c, 0xe3, 0xb7, 0xbc, 0x32, 0x13, 0x15,
	0x97, 0x92, 0x78, 0xac, 0x1d, 0x2f, 0x72, 0x4d, 0x49, 0x07, 0xe4, 0x39, 0xd8, 0x8d, 0x89, 0x64,
	0xc2, 0xfa, 0x16, 0x46, 0x87, 0x4c, 0x67, 0x9e, 0xb5, 0x99, 0x67, 0xaf, 0x9b, 0xcc, 0xd3, 0x01,
	0x79, 0x0a, 0x8e, 0x76, 0xfe, 0xf7, 0x47, 0x01, 0xeb, 0xed, 0x8c, 0x0e, 0xc8, 0x09, 0x38, 0x17,
	0x85, 0xf8, 0xc7, 0x0e, 0xc7, 0x00, 0x6f, 0x6a, 0x21, 0x74, 0xa8, 0x89, 0xc3, 0x30, 0x27, 0xd1,
	0x84, 0xf5, 0x53, 0x8e, 0x52, 0xee, 0xbd, 0x94, 0x42, 0xf0, 0x75, 0x95, 0xcb, 0xc2, 0x2c, 0xc5,
	0x65, 0xe6, 0xff, 0x5e, 0xee, 0xe8, 0xe0, 0xd4, 0xfb, 0xec, 0xce, 0xf5, 0x16, 0x57, 0x0e, 0xb6,
	0x3c, 0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xae, 0x94, 0x7b, 0x1f, 0x04, 0x00, 0x00,
}
