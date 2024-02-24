// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/commons/apiresource/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ApiResourceListViewQueryController_ListByResourceKind_FullMethodName          = "/cloud.planton.apis.commons.apiresource.service.ApiResourceListViewQueryController/listByResourceKind"
	ApiResourceListViewQueryController_GetResourceCountByCompanyId_FullMethodName = "/cloud.planton.apis.commons.apiresource.service.ApiResourceListViewQueryController/getResourceCountByCompanyId"
)

// ApiResourceListViewQueryControllerClient is the client API for ApiResourceListViewQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiResourceListViewQueryControllerClient interface {
	// This method returns a `ResourceList` message, which encapsulates a list of resources that match
	// the input search parameters. Each resource in the list should match the specified resource type,
	// and be associated with the specified company and product.
	ListByResourceKind(ctx context.Context, in *model.GetByApiResourceKindInput, opts ...grpc.CallOption) (*model.ApiResourceRecordList, error)
	// getResourceCountByCompanyId retrieves detailed information about the count of different resources
	// associated with a given company. The request requires a GetResourceCountByCompanyIdInput message
	// containing the company_id of interest. It returns a ResourceCount message, which includes the type
	// and name of the resource along with its total count within the specified company.
	GetResourceCountByCompanyId(ctx context.Context, in *model.GetResourceCountByCompanyIdInput, opts ...grpc.CallOption) (*model.ApiResourcesCount, error)
}

type apiResourceListViewQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewApiResourceListViewQueryControllerClient(cc grpc.ClientConnInterface) ApiResourceListViewQueryControllerClient {
	return &apiResourceListViewQueryControllerClient{cc}
}

func (c *apiResourceListViewQueryControllerClient) ListByResourceKind(ctx context.Context, in *model.GetByApiResourceKindInput, opts ...grpc.CallOption) (*model.ApiResourceRecordList, error) {
	out := new(model.ApiResourceRecordList)
	err := c.cc.Invoke(ctx, ApiResourceListViewQueryController_ListByResourceKind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiResourceListViewQueryControllerClient) GetResourceCountByCompanyId(ctx context.Context, in *model.GetResourceCountByCompanyIdInput, opts ...grpc.CallOption) (*model.ApiResourcesCount, error) {
	out := new(model.ApiResourcesCount)
	err := c.cc.Invoke(ctx, ApiResourceListViewQueryController_GetResourceCountByCompanyId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiResourceListViewQueryControllerServer is the server API for ApiResourceListViewQueryController service.
// All implementations should embed UnimplementedApiResourceListViewQueryControllerServer
// for forward compatibility
type ApiResourceListViewQueryControllerServer interface {
	// This method returns a `ResourceList` message, which encapsulates a list of resources that match
	// the input search parameters. Each resource in the list should match the specified resource type,
	// and be associated with the specified company and product.
	ListByResourceKind(context.Context, *model.GetByApiResourceKindInput) (*model.ApiResourceRecordList, error)
	// getResourceCountByCompanyId retrieves detailed information about the count of different resources
	// associated with a given company. The request requires a GetResourceCountByCompanyIdInput message
	// containing the company_id of interest. It returns a ResourceCount message, which includes the type
	// and name of the resource along with its total count within the specified company.
	GetResourceCountByCompanyId(context.Context, *model.GetResourceCountByCompanyIdInput) (*model.ApiResourcesCount, error)
}

// UnimplementedApiResourceListViewQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedApiResourceListViewQueryControllerServer struct {
}

func (UnimplementedApiResourceListViewQueryControllerServer) ListByResourceKind(context.Context, *model.GetByApiResourceKindInput) (*model.ApiResourceRecordList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByResourceKind not implemented")
}
func (UnimplementedApiResourceListViewQueryControllerServer) GetResourceCountByCompanyId(context.Context, *model.GetResourceCountByCompanyIdInput) (*model.ApiResourcesCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceCountByCompanyId not implemented")
}

// UnsafeApiResourceListViewQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiResourceListViewQueryControllerServer will
// result in compilation errors.
type UnsafeApiResourceListViewQueryControllerServer interface {
	mustEmbedUnimplementedApiResourceListViewQueryControllerServer()
}

func RegisterApiResourceListViewQueryControllerServer(s grpc.ServiceRegistrar, srv ApiResourceListViewQueryControllerServer) {
	s.RegisterService(&ApiResourceListViewQueryController_ServiceDesc, srv)
}

func _ApiResourceListViewQueryController_ListByResourceKind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetByApiResourceKindInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceListViewQueryControllerServer).ListByResourceKind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceListViewQueryController_ListByResourceKind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceListViewQueryControllerServer).ListByResourceKind(ctx, req.(*model.GetByApiResourceKindInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiResourceListViewQueryController_GetResourceCountByCompanyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetResourceCountByCompanyIdInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceListViewQueryControllerServer).GetResourceCountByCompanyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceListViewQueryController_GetResourceCountByCompanyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceListViewQueryControllerServer).GetResourceCountByCompanyId(ctx, req.(*model.GetResourceCountByCompanyIdInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiResourceListViewQueryController_ServiceDesc is the grpc.ServiceDesc for ApiResourceListViewQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiResourceListViewQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.commons.apiresource.service.ApiResourceListViewQueryController",
	HandlerType: (*ApiResourceListViewQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listByResourceKind",
			Handler:    _ApiResourceListViewQueryController_ListByResourceKind_Handler,
		},
		{
			MethodName: "getResourceCountByCompanyId",
			Handler:    _ApiResourceListViewQueryController_GetResourceCountByCompanyId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/commons/apiresource/service/query.proto",
}
