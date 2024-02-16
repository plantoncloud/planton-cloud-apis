// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/resourcemanager/resource/version/query.proto

package version

import (
	context "context"
	version "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/version"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ResourceVersionQueryController_ListByFilters_FullMethodName          = "/cloud.planton.apis.v1.resourcemanager.resource.version.ResourceVersionQueryController/listByFilters"
	ResourceVersionQueryController_GetByIdWithContextSize_FullMethodName = "/cloud.planton.apis.v1.resourcemanager.resource.version.ResourceVersionQueryController/getByIdWithContextSize"
	ResourceVersionQueryController_GetCount_FullMethodName               = "/cloud.planton.apis.v1.resourcemanager.resource.version.ResourceVersionQueryController/getCount"
)

// ResourceVersionQueryControllerClient is the client API for ResourceVersionQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceVersionQueryControllerClient interface {
	// list of resource-versions
	ListByFilters(ctx context.Context, in *version.ListResourceVersionsInput, opts ...grpc.CallOption) (*version.ResourceVersionList, error)
	// look up resource-version by version-id
	GetByIdWithContextSize(ctx context.Context, in *version.ResourceVersionWithContextSizeInput, opts ...grpc.CallOption) (*version.ResourceVersion, error)
	// count of resource-versions
	GetCount(ctx context.Context, in *version.GetResourceVersionCountInput, opts ...grpc.CallOption) (*version.ResourceVersionCount, error)
}

type resourceVersionQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceVersionQueryControllerClient(cc grpc.ClientConnInterface) ResourceVersionQueryControllerClient {
	return &resourceVersionQueryControllerClient{cc}
}

func (c *resourceVersionQueryControllerClient) ListByFilters(ctx context.Context, in *version.ListResourceVersionsInput, opts ...grpc.CallOption) (*version.ResourceVersionList, error) {
	out := new(version.ResourceVersionList)
	err := c.cc.Invoke(ctx, ResourceVersionQueryController_ListByFilters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceVersionQueryControllerClient) GetByIdWithContextSize(ctx context.Context, in *version.ResourceVersionWithContextSizeInput, opts ...grpc.CallOption) (*version.ResourceVersion, error) {
	out := new(version.ResourceVersion)
	err := c.cc.Invoke(ctx, ResourceVersionQueryController_GetByIdWithContextSize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceVersionQueryControllerClient) GetCount(ctx context.Context, in *version.GetResourceVersionCountInput, opts ...grpc.CallOption) (*version.ResourceVersionCount, error) {
	out := new(version.ResourceVersionCount)
	err := c.cc.Invoke(ctx, ResourceVersionQueryController_GetCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceVersionQueryControllerServer is the server API for ResourceVersionQueryController service.
// All implementations should embed UnimplementedResourceVersionQueryControllerServer
// for forward compatibility
type ResourceVersionQueryControllerServer interface {
	// list of resource-versions
	ListByFilters(context.Context, *version.ListResourceVersionsInput) (*version.ResourceVersionList, error)
	// look up resource-version by version-id
	GetByIdWithContextSize(context.Context, *version.ResourceVersionWithContextSizeInput) (*version.ResourceVersion, error)
	// count of resource-versions
	GetCount(context.Context, *version.GetResourceVersionCountInput) (*version.ResourceVersionCount, error)
}

// UnimplementedResourceVersionQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedResourceVersionQueryControllerServer struct {
}

func (UnimplementedResourceVersionQueryControllerServer) ListByFilters(context.Context, *version.ListResourceVersionsInput) (*version.ResourceVersionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByFilters not implemented")
}
func (UnimplementedResourceVersionQueryControllerServer) GetByIdWithContextSize(context.Context, *version.ResourceVersionWithContextSizeInput) (*version.ResourceVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdWithContextSize not implemented")
}
func (UnimplementedResourceVersionQueryControllerServer) GetCount(context.Context, *version.GetResourceVersionCountInput) (*version.ResourceVersionCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCount not implemented")
}

// UnsafeResourceVersionQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceVersionQueryControllerServer will
// result in compilation errors.
type UnsafeResourceVersionQueryControllerServer interface {
	mustEmbedUnimplementedResourceVersionQueryControllerServer()
}

func RegisterResourceVersionQueryControllerServer(s grpc.ServiceRegistrar, srv ResourceVersionQueryControllerServer) {
	s.RegisterService(&ResourceVersionQueryController_ServiceDesc, srv)
}

func _ResourceVersionQueryController_ListByFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(version.ListResourceVersionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceVersionQueryControllerServer).ListByFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceVersionQueryController_ListByFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceVersionQueryControllerServer).ListByFilters(ctx, req.(*version.ListResourceVersionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceVersionQueryController_GetByIdWithContextSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(version.ResourceVersionWithContextSizeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceVersionQueryControllerServer).GetByIdWithContextSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceVersionQueryController_GetByIdWithContextSize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceVersionQueryControllerServer).GetByIdWithContextSize(ctx, req.(*version.ResourceVersionWithContextSizeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceVersionQueryController_GetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(version.GetResourceVersionCountInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceVersionQueryControllerServer).GetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceVersionQueryController_GetCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceVersionQueryControllerServer).GetCount(ctx, req.(*version.GetResourceVersionCountInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceVersionQueryController_ServiceDesc is the grpc.ServiceDesc for ResourceVersionQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceVersionQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.resourcemanager.resource.version.ResourceVersionQueryController",
	HandlerType: (*ResourceVersionQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listByFilters",
			Handler:    _ResourceVersionQueryController_ListByFilters_Handler,
		},
		{
			MethodName: "getByIdWithContextSize",
			Handler:    _ResourceVersionQueryController_GetByIdWithContextSize_Handler,
		},
		{
			MethodName: "getCount",
			Handler:    _ResourceVersionQueryController_GetCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/resourcemanager/resource/version/query.proto",
}
