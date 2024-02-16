// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/resourcemanager/resource/list/service/query.proto

package service

import (
	context "context"
	list "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/list"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ResourceListViewQueryController_ListByResourceType_FullMethodName = "/cloud.planton.apis.v1.resourcemanager.resource.list.service.ResourceListViewQueryController/listByResourceType"
)

// ResourceListViewQueryControllerClient is the client API for ResourceListViewQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceListViewQueryControllerClient interface {
	// This method returns a `ResourceList` message, which encapsulates a list of resources that match
	// the input search parameters. Each resource in the list should match the specified resource type,
	// and be associated with the specified company and product.
	ListByResourceType(ctx context.Context, in *list.GetByResourceTypeInput, opts ...grpc.CallOption) (*list.RecordList, error)
}

type resourceListViewQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceListViewQueryControllerClient(cc grpc.ClientConnInterface) ResourceListViewQueryControllerClient {
	return &resourceListViewQueryControllerClient{cc}
}

func (c *resourceListViewQueryControllerClient) ListByResourceType(ctx context.Context, in *list.GetByResourceTypeInput, opts ...grpc.CallOption) (*list.RecordList, error) {
	out := new(list.RecordList)
	err := c.cc.Invoke(ctx, ResourceListViewQueryController_ListByResourceType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceListViewQueryControllerServer is the server API for ResourceListViewQueryController service.
// All implementations should embed UnimplementedResourceListViewQueryControllerServer
// for forward compatibility
type ResourceListViewQueryControllerServer interface {
	// This method returns a `ResourceList` message, which encapsulates a list of resources that match
	// the input search parameters. Each resource in the list should match the specified resource type,
	// and be associated with the specified company and product.
	ListByResourceType(context.Context, *list.GetByResourceTypeInput) (*list.RecordList, error)
}

// UnimplementedResourceListViewQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedResourceListViewQueryControllerServer struct {
}

func (UnimplementedResourceListViewQueryControllerServer) ListByResourceType(context.Context, *list.GetByResourceTypeInput) (*list.RecordList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByResourceType not implemented")
}

// UnsafeResourceListViewQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceListViewQueryControllerServer will
// result in compilation errors.
type UnsafeResourceListViewQueryControllerServer interface {
	mustEmbedUnimplementedResourceListViewQueryControllerServer()
}

func RegisterResourceListViewQueryControllerServer(s grpc.ServiceRegistrar, srv ResourceListViewQueryControllerServer) {
	s.RegisterService(&ResourceListViewQueryController_ServiceDesc, srv)
}

func _ResourceListViewQueryController_ListByResourceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(list.GetByResourceTypeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceListViewQueryControllerServer).ListByResourceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceListViewQueryController_ListByResourceType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceListViewQueryControllerServer).ListByResourceType(ctx, req.(*list.GetByResourceTypeInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceListViewQueryController_ServiceDesc is the grpc.ServiceDesc for ResourceListViewQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceListViewQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.resourcemanager.resource.list.service.ResourceListViewQueryController",
	HandlerType: (*ResourceListViewQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listByResourceType",
			Handler:    _ResourceListViewQueryController_ListByResourceType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/resourcemanager/resource/list/service/query.proto",
}
