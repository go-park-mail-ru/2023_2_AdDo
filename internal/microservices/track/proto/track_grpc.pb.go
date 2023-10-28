// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TrackServiceClient is the client API for TrackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrackServiceClient interface {
	Listen(ctx context.Context, in *TrackId, opts ...grpc.CallOption) (*empty.Empty, error)
	Like(ctx context.Context, in *TrackToUserId, opts ...grpc.CallOption) (*empty.Empty, error)
}

type trackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackServiceClient(cc grpc.ClientConnInterface) TrackServiceClient {
	return &trackServiceClient{cc}
}

func (c *trackServiceClient) Listen(ctx context.Context, in *TrackId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/TrackService/Listen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackServiceClient) Like(ctx context.Context, in *TrackToUserId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/TrackService/Like", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackServiceServer is the server API for TrackService service.
// All implementations must embed UnimplementedTrackServiceServer
// for forward compatibility
type TrackServiceServer interface {
	Listen(context.Context, *TrackId) (*empty.Empty, error)
	Like(context.Context, *TrackToUserId) (*empty.Empty, error)
	mustEmbedUnimplementedTrackServiceServer()
}

// UnimplementedTrackServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrackServiceServer struct {
}

func (UnimplementedTrackServiceServer) Listen(context.Context, *TrackId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedTrackServiceServer) Like(context.Context, *TrackToUserId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedTrackServiceServer) mustEmbedUnimplementedTrackServiceServer() {}

// UnsafeTrackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrackServiceServer will
// result in compilation errors.
type UnsafeTrackServiceServer interface {
	mustEmbedUnimplementedTrackServiceServer()
}

func RegisterTrackServiceServer(s grpc.ServiceRegistrar, srv TrackServiceServer) {
	s.RegisterService(&TrackService_ServiceDesc, srv)
}

func _TrackService_Listen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackServiceServer).Listen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TrackService/Listen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackServiceServer).Listen(ctx, req.(*TrackId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackService_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackToUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackServiceServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TrackService/Like",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackServiceServer).Like(ctx, req.(*TrackToUserId))
	}
	return interceptor(ctx, in, info, handler)
}

// TrackService_ServiceDesc is the grpc.ServiceDesc for TrackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TrackService",
	HandlerType: (*TrackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Listen",
			Handler:    _TrackService_Listen_Handler,
		},
		{
			MethodName: "Like",
			Handler:    _TrackService_Like_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "track.proto",
}
