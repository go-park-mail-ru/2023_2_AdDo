// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: album.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	proto "main/internal/microservices/track/proto"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IsLikedAlbum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLiked bool `protobuf:"varint,1,opt,name=IsLiked,proto3" json:"IsLiked,omitempty"`
}

func (x *IsLikedAlbum) Reset() {
	*x = IsLikedAlbum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLikedAlbum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLikedAlbum) ProtoMessage() {}

func (x *IsLikedAlbum) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLikedAlbum.ProtoReflect.Descriptor instead.
func (*IsLikedAlbum) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{0}
}

func (x *IsLikedAlbum) GetIsLiked() bool {
	if x != nil {
		return x.IsLiked
	}
	return false
}

type AlbumToUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	AlbumId uint64 `protobuf:"varint,2,opt,name=albumId,proto3" json:"albumId,omitempty"`
}

func (x *AlbumToUserId) Reset() {
	*x = AlbumToUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumToUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumToUserId) ProtoMessage() {}

func (x *AlbumToUserId) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumToUserId.ProtoReflect.Descriptor instead.
func (*AlbumToUserId) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{1}
}

func (x *AlbumToUserId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AlbumToUserId) GetAlbumId() uint64 {
	if x != nil {
		return x.AlbumId
	}
	return 0
}

type AlbumId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlbumId uint64 `protobuf:"varint,1,opt,name=albumId,proto3" json:"albumId,omitempty"`
}

func (x *AlbumId) Reset() {
	*x = AlbumId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumId) ProtoMessage() {}

func (x *AlbumId) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumId.ProtoReflect.Descriptor instead.
func (*AlbumId) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{2}
}

func (x *AlbumId) GetAlbumId() uint64 {
	if x != nil {
		return x.AlbumId
	}
	return 0
}

type AlbumBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Preview string `protobuf:"bytes,3,opt,name=Preview,proto3" json:"Preview,omitempty"`
}

func (x *AlbumBase) Reset() {
	*x = AlbumBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumBase) ProtoMessage() {}

func (x *AlbumBase) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumBase.ProtoReflect.Descriptor instead.
func (*AlbumBase) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{3}
}

func (x *AlbumBase) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AlbumBase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlbumBase) GetPreview() string {
	if x != nil {
		return x.Preview
	}
	return ""
}

type AlbumsBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Albums []*AlbumBase `protobuf:"bytes,1,rep,name=albums,proto3" json:"albums,omitempty"`
}

func (x *AlbumsBase) Reset() {
	*x = AlbumsBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumsBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumsBase) ProtoMessage() {}

func (x *AlbumsBase) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumsBase.ProtoReflect.Descriptor instead.
func (*AlbumsBase) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{4}
}

func (x *AlbumsBase) GetAlbums() []*AlbumBase {
	if x != nil {
		return x.Albums
	}
	return nil
}

type AlbumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64                `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name       string                `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Preview    string                `protobuf:"bytes,3,opt,name=Preview,proto3" json:"Preview,omitempty"`
	ArtistId   uint64                `protobuf:"varint,4,opt,name=ArtistId,proto3" json:"ArtistId,omitempty"`
	ArtistName string                `protobuf:"bytes,5,opt,name=ArtistName,proto3" json:"ArtistName,omitempty"`
	Tracks     *proto.TracksResponse `protobuf:"bytes,6,opt,name=Tracks,proto3" json:"Tracks,omitempty"`
}

func (x *AlbumResponse) Reset() {
	*x = AlbumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumResponse) ProtoMessage() {}

func (x *AlbumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumResponse.ProtoReflect.Descriptor instead.
func (*AlbumResponse) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{5}
}

func (x *AlbumResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AlbumResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlbumResponse) GetPreview() string {
	if x != nil {
		return x.Preview
	}
	return ""
}

func (x *AlbumResponse) GetArtistId() uint64 {
	if x != nil {
		return x.ArtistId
	}
	return 0
}

func (x *AlbumResponse) GetArtistName() string {
	if x != nil {
		return x.ArtistName
	}
	return ""
}

func (x *AlbumResponse) GetTracks() *proto.TracksResponse {
	if x != nil {
		return x.Tracks
	}
	return nil
}

type AlbumsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Albums []*AlbumResponse `protobuf:"bytes,1,rep,name=albums,proto3" json:"albums,omitempty"`
}

func (x *AlbumsResponse) Reset() {
	*x = AlbumsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumsResponse) ProtoMessage() {}

func (x *AlbumsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumsResponse.ProtoReflect.Descriptor instead.
func (*AlbumsResponse) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{6}
}

func (x *AlbumsResponse) GetAlbums() []*AlbumResponse {
	if x != nil {
		return x.Albums
	}
	return nil
}

var File_album_proto protoreflect.FileDescriptor

var file_album_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0c, 0x49, 0x73, 0x4c, 0x69, 0x6b,
	0x65, 0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x4c, 0x69, 0x6b,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65,
	0x64, 0x22, 0x41, 0x0a, 0x0d, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x49, 0x64, 0x22, 0x23, 0x0a, 0x07, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x09, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x22, 0x30, 0x0a, 0x0a, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x42, 0x61,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x61, 0x73, 0x65, 0x52, 0x06,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x0d, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x06, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x38, 0x0a, 0x0e, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x73, 0x32, 0xa8, 0x03, 0x0a, 0x0c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0f, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x12, 0x08, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x1a, 0x0e, 0x2e,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x30, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x0e, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x29, 0x0a, 0x06, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x0e, 0x2e, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x49, 0x73,
	0x4c, 0x69, 0x6b, 0x65, 0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06,
	0x55, 0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x0e, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x54, 0x6f,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x29, 0x5a, 0x27, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_album_proto_rawDescOnce sync.Once
	file_album_proto_rawDescData = file_album_proto_rawDesc
)

func file_album_proto_rawDescGZIP() []byte {
	file_album_proto_rawDescOnce.Do(func() {
		file_album_proto_rawDescData = protoimpl.X.CompressGZIP(file_album_proto_rawDescData)
	})
	return file_album_proto_rawDescData
}

var file_album_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_album_proto_goTypes = []interface{}{
	(*IsLikedAlbum)(nil),         // 0: IsLikedAlbum
	(*AlbumToUserId)(nil),        // 1: AlbumToUserId
	(*AlbumId)(nil),              // 2: AlbumId
	(*AlbumBase)(nil),            // 3: AlbumBase
	(*AlbumsBase)(nil),           // 4: AlbumsBase
	(*AlbumResponse)(nil),        // 5: AlbumResponse
	(*AlbumsResponse)(nil),       // 6: AlbumsResponse
	(*proto.TracksResponse)(nil), // 7: TracksResponse
	(*empty.Empty)(nil),          // 8: google.protobuf.Empty
}
var file_album_proto_depIdxs = []int32{
	3,  // 0: AlbumsBase.albums:type_name -> AlbumBase
	7,  // 1: AlbumResponse.Tracks:type_name -> TracksResponse
	5,  // 2: AlbumsResponse.albums:type_name -> AlbumResponse
	8,  // 3: AlbumService.GetRandom:input_type -> google.protobuf.Empty
	8,  // 4: AlbumService.GetMostLiked:input_type -> google.protobuf.Empty
	8,  // 5: AlbumService.GetPopular:input_type -> google.protobuf.Empty
	8,  // 6: AlbumService.GetNew:input_type -> google.protobuf.Empty
	2,  // 7: AlbumService.GetAlbum:input_type -> AlbumId
	1,  // 8: AlbumService.Like:input_type -> AlbumToUserId
	1,  // 9: AlbumService.IsLike:input_type -> AlbumToUserId
	1,  // 10: AlbumService.Unlike:input_type -> AlbumToUserId
	6,  // 11: AlbumService.GetRandom:output_type -> AlbumsResponse
	6,  // 12: AlbumService.GetMostLiked:output_type -> AlbumsResponse
	6,  // 13: AlbumService.GetPopular:output_type -> AlbumsResponse
	6,  // 14: AlbumService.GetNew:output_type -> AlbumsResponse
	5,  // 15: AlbumService.GetAlbum:output_type -> AlbumResponse
	8,  // 16: AlbumService.Like:output_type -> google.protobuf.Empty
	0,  // 17: AlbumService.IsLike:output_type -> IsLikedAlbum
	8,  // 18: AlbumService.Unlike:output_type -> google.protobuf.Empty
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_album_proto_init() }
func file_album_proto_init() {
	if File_album_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_album_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLikedAlbum); i {
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
		file_album_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumToUserId); i {
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
		file_album_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumId); i {
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
		file_album_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumBase); i {
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
		file_album_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumsBase); i {
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
		file_album_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumResponse); i {
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
		file_album_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumsResponse); i {
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
			RawDescriptor: file_album_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_album_proto_goTypes,
		DependencyIndexes: file_album_proto_depIdxs,
		MessageInfos:      file_album_proto_msgTypes,
	}.Build()
	File_album_proto = out.File
	file_album_proto_rawDesc = nil
	file_album_proto_goTypes = nil
	file_album_proto_depIdxs = nil
}
