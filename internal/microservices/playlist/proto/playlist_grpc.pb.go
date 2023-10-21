// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "main/internal/microservices/track/proto"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PlaylistServiceClient is the client API for PlaylistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaylistServiceClient interface {
	Create(ctx context.Context, in *Base, opts ...grpc.CallOption) (*proto.Status, error)
	Get(ctx context.Context, in *PlaylistId, opts ...grpc.CallOption) (*Response, error)
	GetUserPlaylists(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*PlaylistsBase, error)
	AddTrack(ctx context.Context, in *PlaylistToTrackId, opts ...grpc.CallOption) (*proto.Status, error)
	UpdatePreview(ctx context.Context, in *PlaylistIdToImageUrl, opts ...grpc.CallOption) (*proto.Status, error)
	DeleteById(ctx context.Context, in *PlaylistId, opts ...grpc.CallOption) (*proto.Status, error)
}

type playlistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaylistServiceClient(cc grpc.ClientConnInterface) PlaylistServiceClient {
	return &playlistServiceClient{cc}
}

func (c *playlistServiceClient) Create(ctx context.Context, in *Base, opts ...grpc.CallOption) (*proto.Status, error) {
	out := new(proto.Status)
	err := c.cc.Invoke(ctx, "/PlaylistService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) Get(ctx context.Context, in *PlaylistId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/PlaylistService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) GetUserPlaylists(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*PlaylistsBase, error) {
	out := new(PlaylistsBase)
	err := c.cc.Invoke(ctx, "/PlaylistService/GetUserPlaylists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) AddTrack(ctx context.Context, in *PlaylistToTrackId, opts ...grpc.CallOption) (*proto.Status, error) {
	out := new(proto.Status)
	err := c.cc.Invoke(ctx, "/PlaylistService/AddTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) UpdatePreview(ctx context.Context, in *PlaylistIdToImageUrl, opts ...grpc.CallOption) (*proto.Status, error) {
	out := new(proto.Status)
	err := c.cc.Invoke(ctx, "/PlaylistService/UpdatePreview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) DeleteById(ctx context.Context, in *PlaylistId, opts ...grpc.CallOption) (*proto.Status, error) {
	out := new(proto.Status)
	err := c.cc.Invoke(ctx, "/PlaylistService/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaylistServiceServer is the server API for PlaylistService service.
// All implementations must embed UnimplementedPlaylistServiceServer
// for forward compatibility
type PlaylistServiceServer interface {
	Create(context.Context, *Base) (*proto.Status, error)
	Get(context.Context, *PlaylistId) (*Response, error)
	GetUserPlaylists(context.Context, *UserId) (*PlaylistsBase, error)
	AddTrack(context.Context, *PlaylistToTrackId) (*proto.Status, error)
	UpdatePreview(context.Context, *PlaylistIdToImageUrl) (*proto.Status, error)
	DeleteById(context.Context, *PlaylistId) (*proto.Status, error)
	mustEmbedUnimplementedPlaylistServiceServer()
}

// UnimplementedPlaylistServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlaylistServiceServer struct {
}

func (UnimplementedPlaylistServiceServer) Create(context.Context, *Base) (*proto.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPlaylistServiceServer) Get(context.Context, *PlaylistId) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPlaylistServiceServer) GetUserPlaylists(context.Context, *UserId) (*PlaylistsBase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPlaylists not implemented")
}
func (UnimplementedPlaylistServiceServer) AddTrack(context.Context, *PlaylistToTrackId) (*proto.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTrack not implemented")
}
func (UnimplementedPlaylistServiceServer) UpdatePreview(context.Context, *PlaylistIdToImageUrl) (*proto.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePreview not implemented")
}
func (UnimplementedPlaylistServiceServer) DeleteById(context.Context, *PlaylistId) (*proto.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (UnimplementedPlaylistServiceServer) mustEmbedUnimplementedPlaylistServiceServer() {}

// UnsafePlaylistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaylistServiceServer will
// result in compilation errors.
type UnsafePlaylistServiceServer interface {
	mustEmbedUnimplementedPlaylistServiceServer()
}

func RegisterPlaylistServiceServer(s grpc.ServiceRegistrar, srv PlaylistServiceServer) {
	s.RegisterService(&PlaylistService_ServiceDesc, srv)
}

func _PlaylistService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Base)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaylistService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).Create(ctx, req.(*Base))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaylistService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).Get(ctx, req.(*PlaylistId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_GetUserPlaylists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).GetUserPlaylists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaylistService/GetUserPlaylists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).GetUserPlaylists(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_AddTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistToTrackId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).AddTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaylistService/AddTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).AddTrack(ctx, req.(*PlaylistToTrackId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_UpdatePreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistIdToImageUrl)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).UpdatePreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaylistService/UpdatePreview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).UpdatePreview(ctx, req.(*PlaylistIdToImageUrl))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlaylistService/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).DeleteById(ctx, req.(*PlaylistId))
	}
	return interceptor(ctx, in, info, handler)
}

// PlaylistService_ServiceDesc is the grpc.ServiceDesc for PlaylistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlaylistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PlaylistService",
	HandlerType: (*PlaylistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PlaylistService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PlaylistService_Get_Handler,
		},
		{
			MethodName: "GetUserPlaylists",
			Handler:    _PlaylistService_GetUserPlaylists_Handler,
		},
		{
			MethodName: "AddTrack",
			Handler:    _PlaylistService_AddTrack_Handler,
		},
		{
			MethodName: "UpdatePreview",
			Handler:    _PlaylistService_UpdatePreview_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _PlaylistService_DeleteById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playlist.proto",
}
