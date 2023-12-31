// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/k8s/v1/version_service.proto

package k8s

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	VersionService_List_FullMethodName = "/yandex.cloud.k8s.v1.VersionService/List"
)

// VersionServiceClient is the client API for VersionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VersionServiceClient interface {
	// Retrieves the list of versions in the specified release channel.
	List(ctx context.Context, in *ListVersionsRequest, opts ...grpc.CallOption) (*ListVersionsResponse, error)
}

type versionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVersionServiceClient(cc grpc.ClientConnInterface) VersionServiceClient {
	return &versionServiceClient{cc}
}

func (c *versionServiceClient) List(ctx context.Context, in *ListVersionsRequest, opts ...grpc.CallOption) (*ListVersionsResponse, error) {
	out := new(ListVersionsResponse)
	err := c.cc.Invoke(ctx, VersionService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionServiceServer is the server API for VersionService service.
// All implementations should embed UnimplementedVersionServiceServer
// for forward compatibility
type VersionServiceServer interface {
	// Retrieves the list of versions in the specified release channel.
	List(context.Context, *ListVersionsRequest) (*ListVersionsResponse, error)
}

// UnimplementedVersionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVersionServiceServer struct {
}

func (UnimplementedVersionServiceServer) List(context.Context, *ListVersionsRequest) (*ListVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeVersionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VersionServiceServer will
// result in compilation errors.
type UnsafeVersionServiceServer interface {
	mustEmbedUnimplementedVersionServiceServer()
}

func RegisterVersionServiceServer(s grpc.ServiceRegistrar, srv VersionServiceServer) {
	s.RegisterService(&VersionService_ServiceDesc, srv)
}

func _VersionService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VersionService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).List(ctx, req.(*ListVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VersionService_ServiceDesc is the grpc.ServiceDesc for VersionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VersionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.k8s.v1.VersionService",
	HandlerType: (*VersionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _VersionService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/k8s/v1/version_service.proto",
}
