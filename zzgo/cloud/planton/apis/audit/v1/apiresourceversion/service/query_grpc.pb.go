// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/audit/v1/apiresourceversion/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/audit/v1/apiresourceversion/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ApiResourceVersionQueryController_GetById_FullMethodName                     = "/cloud.planton.apis.audit.v1.apiresourceversion.service.ApiResourceVersionQueryController/getById"
	ApiResourceVersionQueryController_ListByFilters_FullMethodName               = "/cloud.planton.apis.audit.v1.apiresourceversion.service.ApiResourceVersionQueryController/listByFilters"
	ApiResourceVersionQueryController_GetByIdWithContextSize_FullMethodName      = "/cloud.planton.apis.audit.v1.apiresourceversion.service.ApiResourceVersionQueryController/getByIdWithContextSize"
	ApiResourceVersionQueryController_GetCount_FullMethodName                    = "/cloud.planton.apis.audit.v1.apiresourceversion.service.ApiResourceVersionQueryController/getCount"
	ApiResourceVersionQueryController_GetResourceCountByCompanyId_FullMethodName = "/cloud.planton.apis.audit.v1.apiresourceversion.service.ApiResourceVersionQueryController/getResourceCountByCompanyId"
)

// ApiResourceVersionQueryControllerClient is the client API for ApiResourceVersionQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiResourceVersionQueryControllerClient interface {
	// look up a cloud-account by id
	// todo: add authorization
	GetById(ctx context.Context, in *model.ApiResourceVersionId, opts ...grpc.CallOption) (*model.ApiResourceVersion, error)
	// list of api-resource-versions
	ListByFilters(ctx context.Context, in *model.ListApiResourceVersionsInput, opts ...grpc.CallOption) (*model.ApiResourceVersionList, error)
	// look up api-resource-version by version-id
	GetByIdWithContextSize(ctx context.Context, in *model.ApiResourceVersionWithContextSizeInput, opts ...grpc.CallOption) (*model.ApiResourceVersion, error)
	// count of api-resource-versions
	GetCount(ctx context.Context, in *model.GetApiResourceVersionCountInput, opts ...grpc.CallOption) (*model.ApiResourceVersionCount, error)
	// getResourceCountByCompanyId retrieves detailed information about the count of different resources
	// associated with a given company. The request requires a GetResourceCountByCompanyIdInput message
	// containing the company_id of interest. It returns a ResourceCount message, which includes the type
	// and name of the resource along with its total count within the specified company.
	GetResourceCountByCompanyId(ctx context.Context, in *model.GetResourceCountByCompanyIdInput, opts ...grpc.CallOption) (*model.ApiResourcesCount, error)
}

type apiResourceVersionQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewApiResourceVersionQueryControllerClient(cc grpc.ClientConnInterface) ApiResourceVersionQueryControllerClient {
	return &apiResourceVersionQueryControllerClient{cc}
}

func (c *apiResourceVersionQueryControllerClient) GetById(ctx context.Context, in *model.ApiResourceVersionId, opts ...grpc.CallOption) (*model.ApiResourceVersion, error) {
	out := new(model.ApiResourceVersion)
	err := c.cc.Invoke(ctx, ApiResourceVersionQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiResourceVersionQueryControllerClient) ListByFilters(ctx context.Context, in *model.ListApiResourceVersionsInput, opts ...grpc.CallOption) (*model.ApiResourceVersionList, error) {
	out := new(model.ApiResourceVersionList)
	err := c.cc.Invoke(ctx, ApiResourceVersionQueryController_ListByFilters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiResourceVersionQueryControllerClient) GetByIdWithContextSize(ctx context.Context, in *model.ApiResourceVersionWithContextSizeInput, opts ...grpc.CallOption) (*model.ApiResourceVersion, error) {
	out := new(model.ApiResourceVersion)
	err := c.cc.Invoke(ctx, ApiResourceVersionQueryController_GetByIdWithContextSize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiResourceVersionQueryControllerClient) GetCount(ctx context.Context, in *model.GetApiResourceVersionCountInput, opts ...grpc.CallOption) (*model.ApiResourceVersionCount, error) {
	out := new(model.ApiResourceVersionCount)
	err := c.cc.Invoke(ctx, ApiResourceVersionQueryController_GetCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiResourceVersionQueryControllerClient) GetResourceCountByCompanyId(ctx context.Context, in *model.GetResourceCountByCompanyIdInput, opts ...grpc.CallOption) (*model.ApiResourcesCount, error) {
	out := new(model.ApiResourcesCount)
	err := c.cc.Invoke(ctx, ApiResourceVersionQueryController_GetResourceCountByCompanyId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiResourceVersionQueryControllerServer is the server API for ApiResourceVersionQueryController service.
// All implementations should embed UnimplementedApiResourceVersionQueryControllerServer
// for forward compatibility
type ApiResourceVersionQueryControllerServer interface {
	// look up a cloud-account by id
	// todo: add authorization
	GetById(context.Context, *model.ApiResourceVersionId) (*model.ApiResourceVersion, error)
	// list of api-resource-versions
	ListByFilters(context.Context, *model.ListApiResourceVersionsInput) (*model.ApiResourceVersionList, error)
	// look up api-resource-version by version-id
	GetByIdWithContextSize(context.Context, *model.ApiResourceVersionWithContextSizeInput) (*model.ApiResourceVersion, error)
	// count of api-resource-versions
	GetCount(context.Context, *model.GetApiResourceVersionCountInput) (*model.ApiResourceVersionCount, error)
	// getResourceCountByCompanyId retrieves detailed information about the count of different resources
	// associated with a given company. The request requires a GetResourceCountByCompanyIdInput message
	// containing the company_id of interest. It returns a ResourceCount message, which includes the type
	// and name of the resource along with its total count within the specified company.
	GetResourceCountByCompanyId(context.Context, *model.GetResourceCountByCompanyIdInput) (*model.ApiResourcesCount, error)
}

// UnimplementedApiResourceVersionQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedApiResourceVersionQueryControllerServer struct {
}

func (UnimplementedApiResourceVersionQueryControllerServer) GetById(context.Context, *model.ApiResourceVersionId) (*model.ApiResourceVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedApiResourceVersionQueryControllerServer) ListByFilters(context.Context, *model.ListApiResourceVersionsInput) (*model.ApiResourceVersionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByFilters not implemented")
}
func (UnimplementedApiResourceVersionQueryControllerServer) GetByIdWithContextSize(context.Context, *model.ApiResourceVersionWithContextSizeInput) (*model.ApiResourceVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdWithContextSize not implemented")
}
func (UnimplementedApiResourceVersionQueryControllerServer) GetCount(context.Context, *model.GetApiResourceVersionCountInput) (*model.ApiResourceVersionCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCount not implemented")
}
func (UnimplementedApiResourceVersionQueryControllerServer) GetResourceCountByCompanyId(context.Context, *model.GetResourceCountByCompanyIdInput) (*model.ApiResourcesCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceCountByCompanyId not implemented")
}

// UnsafeApiResourceVersionQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiResourceVersionQueryControllerServer will
// result in compilation errors.
type UnsafeApiResourceVersionQueryControllerServer interface {
	mustEmbedUnimplementedApiResourceVersionQueryControllerServer()
}

func RegisterApiResourceVersionQueryControllerServer(s grpc.ServiceRegistrar, srv ApiResourceVersionQueryControllerServer) {
	s.RegisterService(&ApiResourceVersionQueryController_ServiceDesc, srv)
}

func _ApiResourceVersionQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ApiResourceVersionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceVersionQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceVersionQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceVersionQueryControllerServer).GetById(ctx, req.(*model.ApiResourceVersionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiResourceVersionQueryController_ListByFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ListApiResourceVersionsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceVersionQueryControllerServer).ListByFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceVersionQueryController_ListByFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceVersionQueryControllerServer).ListByFilters(ctx, req.(*model.ListApiResourceVersionsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiResourceVersionQueryController_GetByIdWithContextSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ApiResourceVersionWithContextSizeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceVersionQueryControllerServer).GetByIdWithContextSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceVersionQueryController_GetByIdWithContextSize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceVersionQueryControllerServer).GetByIdWithContextSize(ctx, req.(*model.ApiResourceVersionWithContextSizeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiResourceVersionQueryController_GetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetApiResourceVersionCountInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceVersionQueryControllerServer).GetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceVersionQueryController_GetCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceVersionQueryControllerServer).GetCount(ctx, req.(*model.GetApiResourceVersionCountInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiResourceVersionQueryController_GetResourceCountByCompanyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetResourceCountByCompanyIdInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceVersionQueryControllerServer).GetResourceCountByCompanyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceVersionQueryController_GetResourceCountByCompanyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceVersionQueryControllerServer).GetResourceCountByCompanyId(ctx, req.(*model.GetResourceCountByCompanyIdInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiResourceVersionQueryController_ServiceDesc is the grpc.ServiceDesc for ApiResourceVersionQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiResourceVersionQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.audit.v1.apiresourceversion.service.ApiResourceVersionQueryController",
	HandlerType: (*ApiResourceVersionQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _ApiResourceVersionQueryController_GetById_Handler,
		},
		{
			MethodName: "listByFilters",
			Handler:    _ApiResourceVersionQueryController_ListByFilters_Handler,
		},
		{
			MethodName: "getByIdWithContextSize",
			Handler:    _ApiResourceVersionQueryController_GetByIdWithContextSize_Handler,
		},
		{
			MethodName: "getCount",
			Handler:    _ApiResourceVersionQueryController_GetCount_Handler,
		},
		{
			MethodName: "getResourceCountByCompanyId",
			Handler:    _ApiResourceVersionQueryController_GetResourceCountByCompanyId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/audit/v1/apiresourceversion/service/query.proto",
}
