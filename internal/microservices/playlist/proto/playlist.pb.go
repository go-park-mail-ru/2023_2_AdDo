// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: playlist.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	proto "main/internal/microservices/image/proto"
	proto2 "main/internal/microservices/session/proto"
	proto1 "main/internal/microservices/track/proto"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlaylistIdToNewTitle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId uint64 `protobuf:"varint,1,opt,name=playlistId,proto3" json:"playlistId,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *PlaylistIdToNewTitle) Reset() {
	*x = PlaylistIdToNewTitle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistIdToNewTitle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistIdToNewTitle) ProtoMessage() {}

func (x *PlaylistIdToNewTitle) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistIdToNewTitle.ProtoReflect.Descriptor instead.
func (*PlaylistIdToNewTitle) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{0}
}

func (x *PlaylistIdToNewTitle) GetPlaylistId() uint64 {
	if x != nil {
		return x.PlaylistId
	}
	return 0
}

func (x *PlaylistIdToNewTitle) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type IsLikedPlaylist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLiked bool `protobuf:"varint,1,opt,name=IsLiked,proto3" json:"IsLiked,omitempty"`
}

func (x *IsLikedPlaylist) Reset() {
	*x = IsLikedPlaylist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLikedPlaylist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLikedPlaylist) ProtoMessage() {}

func (x *IsLikedPlaylist) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLikedPlaylist.ProtoReflect.Descriptor instead.
func (*IsLikedPlaylist) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{1}
}

func (x *IsLikedPlaylist) GetIsLiked() bool {
	if x != nil {
		return x.IsLiked
	}
	return false
}

type HasAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAccess bool `protobuf:"varint,1,opt,name=IsAccess,proto3" json:"IsAccess,omitempty"`
}

func (x *HasAccess) Reset() {
	*x = HasAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasAccess) ProtoMessage() {}

func (x *HasAccess) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasAccess.ProtoReflect.Descriptor instead.
func (*HasAccess) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{2}
}

func (x *HasAccess) GetIsAccess() bool {
	if x != nil {
		return x.IsAccess
	}
	return false
}

type IsPlaylistCreator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsCreator bool `protobuf:"varint,1,opt,name=IsCreator,proto3" json:"IsCreator,omitempty"`
}

func (x *IsPlaylistCreator) Reset() {
	*x = IsPlaylistCreator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsPlaylistCreator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsPlaylistCreator) ProtoMessage() {}

func (x *IsPlaylistCreator) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsPlaylistCreator.ProtoReflect.Descriptor instead.
func (*IsPlaylistCreator) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{3}
}

func (x *IsPlaylistCreator) GetIsCreator() bool {
	if x != nil {
		return x.IsCreator
	}
	return false
}

type PlaylistToUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	PlaylistId uint64 `protobuf:"varint,2,opt,name=playlistId,proto3" json:"playlistId,omitempty"`
}

func (x *PlaylistToUserId) Reset() {
	*x = PlaylistToUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistToUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistToUserId) ProtoMessage() {}

func (x *PlaylistToUserId) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistToUserId.ProtoReflect.Descriptor instead.
func (*PlaylistToUserId) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{4}
}

func (x *PlaylistToUserId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PlaylistToUserId) GetPlaylistId() uint64 {
	if x != nil {
		return x.PlaylistId
	}
	return 0
}

type PlaylistId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PlaylistId) Reset() {
	*x = PlaylistId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistId) ProtoMessage() {}

func (x *PlaylistId) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistId.ProtoReflect.Descriptor instead.
func (*PlaylistId) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{5}
}

func (x *PlaylistId) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PlaylistToTrackId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId uint64 `protobuf:"varint,1,opt,name=playlistId,proto3" json:"playlistId,omitempty"`
	TrackId    uint64 `protobuf:"varint,2,opt,name=trackId,proto3" json:"trackId,omitempty"`
}

func (x *PlaylistToTrackId) Reset() {
	*x = PlaylistToTrackId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistToTrackId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistToTrackId) ProtoMessage() {}

func (x *PlaylistToTrackId) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistToTrackId.ProtoReflect.Descriptor instead.
func (*PlaylistToTrackId) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{6}
}

func (x *PlaylistToTrackId) GetPlaylistId() uint64 {
	if x != nil {
		return x.PlaylistId
	}
	return 0
}

func (x *PlaylistToTrackId) GetTrackId() uint64 {
	if x != nil {
		return x.TrackId
	}
	return 0
}

type PlaylistIdToImageUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Url *proto.ImageUrl `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *PlaylistIdToImageUrl) Reset() {
	*x = PlaylistIdToImageUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistIdToImageUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistIdToImageUrl) ProtoMessage() {}

func (x *PlaylistIdToImageUrl) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistIdToImageUrl.ProtoReflect.Descriptor instead.
func (*PlaylistIdToImageUrl) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{7}
}

func (x *PlaylistIdToImageUrl) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PlaylistIdToImageUrl) GetUrl() *proto.ImageUrl {
	if x != nil {
		return x.Url
	}
	return nil
}

type PlaylistBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CreatorId string `protobuf:"bytes,3,opt,name=CreatorId,proto3" json:"CreatorId,omitempty"`
	Preview   string `protobuf:"bytes,4,opt,name=Preview,proto3" json:"Preview,omitempty"`
}

func (x *PlaylistBase) Reset() {
	*x = PlaylistBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistBase) ProtoMessage() {}

func (x *PlaylistBase) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistBase.ProtoReflect.Descriptor instead.
func (*PlaylistBase) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{8}
}

func (x *PlaylistBase) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PlaylistBase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlaylistBase) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *PlaylistBase) GetPreview() string {
	if x != nil {
		return x.Preview
	}
	return ""
}

type PlaylistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64                 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Preview     string                 `protobuf:"bytes,3,opt,name=Preview,proto3" json:"Preview,omitempty"`
	CreatorName string                 `protobuf:"bytes,4,opt,name=CreatorName,proto3" json:"CreatorName,omitempty"`
	CreatorId   string                 `protobuf:"bytes,6,opt,name=CreatorId,proto3" json:"CreatorId,omitempty"`
	Tracks      *proto1.TracksResponse `protobuf:"bytes,7,opt,name=Tracks,proto3" json:"Tracks,omitempty"`
}

func (x *PlaylistResponse) Reset() {
	*x = PlaylistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistResponse) ProtoMessage() {}

func (x *PlaylistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistResponse.ProtoReflect.Descriptor instead.
func (*PlaylistResponse) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{9}
}

func (x *PlaylistResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PlaylistResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlaylistResponse) GetPreview() string {
	if x != nil {
		return x.Preview
	}
	return ""
}

func (x *PlaylistResponse) GetCreatorName() string {
	if x != nil {
		return x.CreatorName
	}
	return ""
}

func (x *PlaylistResponse) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *PlaylistResponse) GetTracks() *proto1.TracksResponse {
	if x != nil {
		return x.Tracks
	}
	return nil
}

type PlaylistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Playlists []*PlaylistResponse `protobuf:"bytes,1,rep,name=playlists,proto3" json:"playlists,omitempty"`
}

func (x *PlaylistsResponse) Reset() {
	*x = PlaylistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistsResponse) ProtoMessage() {}

func (x *PlaylistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistsResponse.ProtoReflect.Descriptor instead.
func (*PlaylistsResponse) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{10}
}

func (x *PlaylistsResponse) GetPlaylists() []*PlaylistResponse {
	if x != nil {
		return x.Playlists
	}
	return nil
}

type PlaylistsBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Playlists []*PlaylistBase `protobuf:"bytes,1,rep,name=playlists,proto3" json:"playlists,omitempty"`
}

func (x *PlaylistsBase) Reset() {
	*x = PlaylistsBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistsBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistsBase) ProtoMessage() {}

func (x *PlaylistsBase) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistsBase.ProtoReflect.Descriptor instead.
func (*PlaylistsBase) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{11}
}

func (x *PlaylistsBase) GetPlaylists() []*PlaylistBase {
	if x != nil {
		return x.Playlists
	}
	return nil
}

var File_playlist_proto protoreflect.FileDescriptor

var file_playlist_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x49, 0x64, 0x54, 0x6f, 0x4e, 0x65, 0x77, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x50,
	0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x4c, 0x69, 0x6b,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65,
	0x64, 0x22, 0x27, 0x0a, 0x09, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x49, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x49, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x31, 0x0a, 0x11, 0x49, 0x73,
	0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x49, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x4a, 0x0a,
	0x10, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70,
	0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x50, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x49, 0x64, 0x54, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x6a, 0x0a, 0x0c, 0x50,
	0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0xb9, 0x01, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x06, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x73, 0x22, 0x44, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09,
	0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x22, 0x3c, 0x0a, 0x0d, 0x50, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x42, 0x61, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x73, 0x65, 0x52, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x32, 0xbf, 0x07, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x42, 0x61, 0x73, 0x65, 0x1a, 0x11, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x0b, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x11, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x49, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x11, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x1a, 0x12, 0x2e, 0x49, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x07, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x73, 0x42, 0x61, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x13, 0x50, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x07,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x73, 0x42, 0x61, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x08, 0x41, 0x64, 0x64,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x12, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x54, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x12, 0x12, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x15, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x54, 0x6f,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x29, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x0b, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64,
	0x1a, 0x09, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x00, 0x12, 0x33, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0b, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x11, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x49, 0x73, 0x4c, 0x69, 0x6b,
	0x65, 0x12, 0x11, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x50, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x55, 0x6e, 0x6c, 0x69,
	0x6b, 0x65, 0x12, 0x11, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x32, 0x0a, 0x0f, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x11, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0d, 0x48, 0x61, 0x73, 0x52, 0x65, 0x61, 0x64, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x0b, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49,
	0x64, 0x1a, 0x0a, 0x2e, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x0b, 0x4d, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x0b,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0a, 0x4d, 0x61, 0x6b, 0x65, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x12, 0x0b, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x49, 0x64, 0x54, 0x6f, 0x4e, 0x65, 0x77, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_playlist_proto_rawDescOnce sync.Once
	file_playlist_proto_rawDescData = file_playlist_proto_rawDesc
)

func file_playlist_proto_rawDescGZIP() []byte {
	file_playlist_proto_rawDescOnce.Do(func() {
		file_playlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_playlist_proto_rawDescData)
	})
	return file_playlist_proto_rawDescData
}

var file_playlist_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_playlist_proto_goTypes = []interface{}{
	(*PlaylistIdToNewTitle)(nil),  // 0: PlaylistIdToNewTitle
	(*IsLikedPlaylist)(nil),       // 1: IsLikedPlaylist
	(*HasAccess)(nil),             // 2: HasAccess
	(*IsPlaylistCreator)(nil),     // 3: IsPlaylistCreator
	(*PlaylistToUserId)(nil),      // 4: PlaylistToUserId
	(*PlaylistId)(nil),            // 5: PlaylistId
	(*PlaylistToTrackId)(nil),     // 6: PlaylistToTrackId
	(*PlaylistIdToImageUrl)(nil),  // 7: PlaylistIdToImageUrl
	(*PlaylistBase)(nil),          // 8: PlaylistBase
	(*PlaylistResponse)(nil),      // 9: PlaylistResponse
	(*PlaylistsResponse)(nil),     // 10: PlaylistsResponse
	(*PlaylistsBase)(nil),         // 11: PlaylistsBase
	(*proto.ImageUrl)(nil),        // 12: ImageUrl
	(*proto1.TracksResponse)(nil), // 13: TracksResponse
	(*proto2.UserId)(nil),         // 14: UserId
	(*empty.Empty)(nil),           // 15: google.protobuf.Empty
}
var file_playlist_proto_depIdxs = []int32{
	12, // 0: PlaylistIdToImageUrl.url:type_name -> ImageUrl
	13, // 1: PlaylistResponse.Tracks:type_name -> TracksResponse
	9,  // 2: PlaylistsResponse.playlists:type_name -> PlaylistResponse
	8,  // 3: PlaylistsBase.playlists:type_name -> PlaylistBase
	8,  // 4: PlaylistService.Create:input_type -> PlaylistBase
	5,  // 5: PlaylistService.Get:input_type -> PlaylistId
	4,  // 6: PlaylistService.IsCreator:input_type -> PlaylistToUserId
	14, // 7: PlaylistService.GetUserPlaylists:input_type -> UserId
	14, // 8: PlaylistService.PlaylistCollections:input_type -> UserId
	6,  // 9: PlaylistService.AddTrack:input_type -> PlaylistToTrackId
	6,  // 10: PlaylistService.RemoveTrack:input_type -> PlaylistToTrackId
	7,  // 11: PlaylistService.UpdatePreview:input_type -> PlaylistIdToImageUrl
	5,  // 12: PlaylistService.RemovePreview:input_type -> PlaylistId
	5,  // 13: PlaylistService.DeleteById:input_type -> PlaylistId
	4,  // 14: PlaylistService.Like:input_type -> PlaylistToUserId
	4,  // 15: PlaylistService.IsLike:input_type -> PlaylistToUserId
	4,  // 16: PlaylistService.Unlike:input_type -> PlaylistToUserId
	4,  // 17: PlaylistService.HasModifyAccess:input_type -> PlaylistToUserId
	5,  // 18: PlaylistService.HasReadAccess:input_type -> PlaylistId
	5,  // 19: PlaylistService.MakePrivate:input_type -> PlaylistId
	5,  // 20: PlaylistService.MakePublic:input_type -> PlaylistId
	0,  // 21: PlaylistService.UpdateName:input_type -> PlaylistIdToNewTitle
	9,  // 22: PlaylistService.Create:output_type -> PlaylistResponse
	9,  // 23: PlaylistService.Get:output_type -> PlaylistResponse
	3,  // 24: PlaylistService.IsCreator:output_type -> IsPlaylistCreator
	11, // 25: PlaylistService.GetUserPlaylists:output_type -> PlaylistsBase
	11, // 26: PlaylistService.PlaylistCollections:output_type -> PlaylistsBase
	15, // 27: PlaylistService.AddTrack:output_type -> google.protobuf.Empty
	15, // 28: PlaylistService.RemoveTrack:output_type -> google.protobuf.Empty
	15, // 29: PlaylistService.UpdatePreview:output_type -> google.protobuf.Empty
	12, // 30: PlaylistService.RemovePreview:output_type -> ImageUrl
	15, // 31: PlaylistService.DeleteById:output_type -> google.protobuf.Empty
	15, // 32: PlaylistService.Like:output_type -> google.protobuf.Empty
	1,  // 33: PlaylistService.IsLike:output_type -> IsLikedPlaylist
	15, // 34: PlaylistService.Unlike:output_type -> google.protobuf.Empty
	2,  // 35: PlaylistService.HasModifyAccess:output_type -> HasAccess
	2,  // 36: PlaylistService.HasReadAccess:output_type -> HasAccess
	15, // 37: PlaylistService.MakePrivate:output_type -> google.protobuf.Empty
	15, // 38: PlaylistService.MakePublic:output_type -> google.protobuf.Empty
	15, // 39: PlaylistService.UpdateName:output_type -> google.protobuf.Empty
	22, // [22:40] is the sub-list for method output_type
	4,  // [4:22] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_playlist_proto_init() }
func file_playlist_proto_init() {
	if File_playlist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_playlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistIdToNewTitle); i {
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
		file_playlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLikedPlaylist); i {
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
		file_playlist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasAccess); i {
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
		file_playlist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsPlaylistCreator); i {
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
		file_playlist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistToUserId); i {
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
		file_playlist_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistId); i {
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
		file_playlist_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistToTrackId); i {
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
		file_playlist_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistIdToImageUrl); i {
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
		file_playlist_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistBase); i {
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
		file_playlist_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistResponse); i {
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
		file_playlist_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistsResponse); i {
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
		file_playlist_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistsBase); i {
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
			RawDescriptor: file_playlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_playlist_proto_goTypes,
		DependencyIndexes: file_playlist_proto_depIdxs,
		MessageInfos:      file_playlist_proto_msgTypes,
	}.Build()
	File_playlist_proto = out.File
	file_playlist_proto_rawDesc = nil
	file_playlist_proto_goTypes = nil
	file_playlist_proto_depIdxs = nil
}
