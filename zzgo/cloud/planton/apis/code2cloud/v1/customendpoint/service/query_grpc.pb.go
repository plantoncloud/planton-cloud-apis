// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/customendpoint/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/customendpoint/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/resourcemanager/v1/product/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CustomEndpointQueryController_GetById_FullMethodName                              = "/cloud.planton.apis.code2cloud.v1.customendpoint.service.CustomEndpointQueryController/getById"
	CustomEndpointQueryController_FindByProductId_FullMethodName                      = "/cloud.planton.apis.code2cloud.v1.customendpoint.service.CustomEndpointQueryController/findByProductId"
	CustomEndpointQueryController_GetCustomEndpointCertStatus_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.customendpoint.service.CustomEndpointQueryController/getCustomEndpointCertStatus"
	CustomEndpointQueryController_GetCustomEndpointDsnResolutionStatus_FullMethodName = "/cloud.planton.apis.code2cloud.v1.customendpoint.service.CustomEndpointQueryController/getCustomEndpointDsnResolutionStatus"
)

// CustomEndpointQueryControllerClient is the client API for CustomEndpointQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomEndpointQueryControllerClient interface {
	// look up custom-endpoint using custom-endpoint id
	GetById(ctx context.Context, in *model.CustomEndpointId, opts ...grpc.CallOption) (*model.CustomEndpoint, error)
	// find custom-endpoints by product id.
	// response contains only the endpoint domains that the authenticated user account id as viewer access to.
	FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.CustomEndpointList, error)
	// check cert status for custom-endpoint
	GetCustomEndpointCertStatus(ctx context.Context, in *model.CustomEndpointId, opts ...grpc.CallOption) (*model.CustomEndpointCertStatus, error)
	// check status of dns resolution for custom-endpoint.
	// confirms if the dns of the custom-endpoint domain is resolving to the correct address.
	GetCustomEndpointDsnResolutionStatus(ctx context.Context, in *model.CustomEndpointId, opts ...grpc.CallOption) (*model.CustomEndpointDnsResolutionStatus, error)
}

type customEndpointQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomEndpointQueryControllerClient(cc grpc.ClientConnInterface) CustomEndpointQueryControllerClient {
	return &customEndpointQueryControllerClient{cc}
}

func (c *customEndpointQueryControllerClient) GetById(ctx context.Context, in *model.CustomEndpointId, opts ...grpc.CallOption) (*model.CustomEndpoint, error) {
	out := new(model.CustomEndpoint)
	err := c.cc.Invoke(ctx, CustomEndpointQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customEndpointQueryControllerClient) FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.CustomEndpointList, error) {
	out := new(model.CustomEndpointList)
	err := c.cc.Invoke(ctx, CustomEndpointQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customEndpointQueryControllerClient) GetCustomEndpointCertStatus(ctx context.Context, in *model.CustomEndpointId, opts ...grpc.CallOption) (*model.CustomEndpointCertStatus, error) {
	out := new(model.CustomEndpointCertStatus)
	err := c.cc.Invoke(ctx, CustomEndpointQueryController_GetCustomEndpointCertStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customEndpointQueryControllerClient) GetCustomEndpointDsnResolutionStatus(ctx context.Context, in *model.CustomEndpointId, opts ...grpc.CallOption) (*model.CustomEndpointDnsResolutionStatus, error) {
	out := new(model.CustomEndpointDnsResolutionStatus)
	err := c.cc.Invoke(ctx, CustomEndpointQueryController_GetCustomEndpointDsnResolutionStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomEndpointQueryControllerServer is the server API for CustomEndpointQueryController service.
// All implementations should embed UnimplementedCustomEndpointQueryControllerServer
// for forward compatibility
type CustomEndpointQueryControllerServer interface {
	// look up custom-endpoint using custom-endpoint id
	GetById(context.Context, *model.CustomEndpointId) (*model.CustomEndpoint, error)
	// find custom-endpoints by product id.
	// response contains only the endpoint domains that the authenticated user account id as viewer access to.
	FindByProductId(context.Context, *model1.ProductId) (*model.CustomEndpointList, error)
	// check cert status for custom-endpoint
	GetCustomEndpointCertStatus(context.Context, *model.CustomEndpointId) (*model.CustomEndpointCertStatus, error)
	// check status of dns resolution for custom-endpoint.
	// confirms if the dns of the custom-endpoint domain is resolving to the correct address.
	GetCustomEndpointDsnResolutionStatus(context.Context, *model.CustomEndpointId) (*model.CustomEndpointDnsResolutionStatus, error)
}

// UnimplementedCustomEndpointQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCustomEndpointQueryControllerServer struct {
}

func (UnimplementedCustomEndpointQueryControllerServer) GetById(context.Context, *model.CustomEndpointId) (*model.CustomEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedCustomEndpointQueryControllerServer) FindByProductId(context.Context, *model1.ProductId) (*model.CustomEndpointList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedCustomEndpointQueryControllerServer) GetCustomEndpointCertStatus(context.Context, *model.CustomEndpointId) (*model.CustomEndpointCertStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomEndpointCertStatus not implemented")
}
func (UnimplementedCustomEndpointQueryControllerServer) GetCustomEndpointDsnResolutionStatus(context.Context, *model.CustomEndpointId) (*model.CustomEndpointDnsResolutionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomEndpointDsnResolutionStatus not implemented")
}

// UnsafeCustomEndpointQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomEndpointQueryControllerServer will
// result in compilation errors.
type UnsafeCustomEndpointQueryControllerServer interface {
	mustEmbedUnimplementedCustomEndpointQueryControllerServer()
}

func RegisterCustomEndpointQueryControllerServer(s grpc.ServiceRegistrar, srv CustomEndpointQueryControllerServer) {
	s.RegisterService(&CustomEndpointQueryController_ServiceDesc, srv)
}

func _CustomEndpointQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CustomEndpointId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEndpointQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomEndpointQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEndpointQueryControllerServer).GetById(ctx, req.(*model.CustomEndpointId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomEndpointQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEndpointQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomEndpointQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEndpointQueryControllerServer).FindByProductId(ctx, req.(*model1.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomEndpointQueryController_GetCustomEndpointCertStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CustomEndpointId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEndpointQueryControllerServer).GetCustomEndpointCertStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomEndpointQueryController_GetCustomEndpointCertStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEndpointQueryControllerServer).GetCustomEndpointCertStatus(ctx, req.(*model.CustomEndpointId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomEndpointQueryController_GetCustomEndpointDsnResolutionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CustomEndpointId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEndpointQueryControllerServer).GetCustomEndpointDsnResolutionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomEndpointQueryController_GetCustomEndpointDsnResolutionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEndpointQueryControllerServer).GetCustomEndpointDsnResolutionStatus(ctx, req.(*model.CustomEndpointId))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomEndpointQueryController_ServiceDesc is the grpc.ServiceDesc for CustomEndpointQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomEndpointQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.customendpoint.service.CustomEndpointQueryController",
	HandlerType: (*CustomEndpointQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _CustomEndpointQueryController_GetById_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _CustomEndpointQueryController_FindByProductId_Handler,
		},
		{
			MethodName: "getCustomEndpointCertStatus",
			Handler:    _CustomEndpointQueryController_GetCustomEndpointCertStatus_Handler,
		},
		{
			MethodName: "getCustomEndpointDsnResolutionStatus",
			Handler:    _CustomEndpointQueryController_GetCustomEndpointDsnResolutionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/customendpoint/service/query.proto",
}
