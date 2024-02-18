// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/environment/service/query.proto

package service

import (
	context "context"
	gcp "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/cloudaccount/model/provider/gcp"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/environment/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/model"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/rpc"
	model3 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/apiresources/model"
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
	EnvironmentQueryController_List_FullMethodName                                  = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentQueryController/list"
	EnvironmentQueryController_GetById_FullMethodName                               = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentQueryController/getById"
	EnvironmentQueryController_FindByProductId_FullMethodName                       = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentQueryController/findByProductId"
	EnvironmentQueryController_FindByKubeClusterId_FullMethodName                   = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentQueryController/findByKubeClusterId"
	EnvironmentQueryController_GetByProductIdAndEnvironmentName_FullMethodName      = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentQueryController/getByProductIdAndEnvironmentName"
	EnvironmentQueryController_GetBuildEngineEnvironmentByProductId_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentQueryController/getBuildEngineEnvironmentByProductId"
	EnvironmentQueryController_GetSecretsGcpProjectByEnvironmentId_FullMethodName   = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentQueryController/getSecretsGcpProjectByEnvironmentId"
	EnvironmentQueryController_FindWorkloadPodsByEnvironmentId_FullMethodName       = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentQueryController/findWorkloadPodsByEnvironmentId"
	EnvironmentQueryController_FindWorkloadNamespacesByEnvironmentId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentQueryController/findWorkloadNamespacesByEnvironmentId"
)

// EnvironmentQueryControllerClient is the client API for EnvironmentQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentQueryControllerClient interface {
	// list all environments on planton cloud for the requested page. This is intended for use on portal.
	List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.EnvironmentList, error)
	// look up environment using environment id
	GetById(ctx context.Context, in *model.EnvironmentId, opts ...grpc.CallOption) (*model.Environment, error)
	// find environments by product id
	FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.Environments, error)
	// find environments by kube-cluster id
	FindByKubeClusterId(ctx context.Context, in *model2.KubeClusterId, opts ...grpc.CallOption) (*model.Environments, error)
	// look up environment using environment id
	GetByProductIdAndEnvironmentName(ctx context.Context, in *model.GetByProductIdAndEnvironmentNameQueryInput, opts ...grpc.CallOption) (*model.Environment, error)
	// look up the build engine environment for the product required for launching the microservice in build engine
	GetBuildEngineEnvironmentByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.Environment, error)
	// look up the gcp project details by environment id required for fetching secrets for launching project in build engine.
	GetSecretsGcpProjectByEnvironmentId(ctx context.Context, in *model.EnvironmentId, opts ...grpc.CallOption) (*gcp.GcpProject, error)
	// find workload pods part of environment
	FindWorkloadPodsByEnvironmentId(ctx context.Context, in *model.EnvironmentId, opts ...grpc.CallOption) (*model3.WorkloadPods, error)
	// find workload namespaces in a environment.
	FindWorkloadNamespacesByEnvironmentId(ctx context.Context, in *model.EnvironmentId, opts ...grpc.CallOption) (*model3.WorkloadNamespaces, error)
}

type environmentQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentQueryControllerClient(cc grpc.ClientConnInterface) EnvironmentQueryControllerClient {
	return &environmentQueryControllerClient{cc}
}

func (c *environmentQueryControllerClient) List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.EnvironmentList, error) {
	out := new(model.EnvironmentList)
	err := c.cc.Invoke(ctx, EnvironmentQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentQueryControllerClient) GetById(ctx context.Context, in *model.EnvironmentId, opts ...grpc.CallOption) (*model.Environment, error) {
	out := new(model.Environment)
	err := c.cc.Invoke(ctx, EnvironmentQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentQueryControllerClient) FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.Environments, error) {
	out := new(model.Environments)
	err := c.cc.Invoke(ctx, EnvironmentQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentQueryControllerClient) FindByKubeClusterId(ctx context.Context, in *model2.KubeClusterId, opts ...grpc.CallOption) (*model.Environments, error) {
	out := new(model.Environments)
	err := c.cc.Invoke(ctx, EnvironmentQueryController_FindByKubeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentQueryControllerClient) GetByProductIdAndEnvironmentName(ctx context.Context, in *model.GetByProductIdAndEnvironmentNameQueryInput, opts ...grpc.CallOption) (*model.Environment, error) {
	out := new(model.Environment)
	err := c.cc.Invoke(ctx, EnvironmentQueryController_GetByProductIdAndEnvironmentName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentQueryControllerClient) GetBuildEngineEnvironmentByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.Environment, error) {
	out := new(model.Environment)
	err := c.cc.Invoke(ctx, EnvironmentQueryController_GetBuildEngineEnvironmentByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentQueryControllerClient) GetSecretsGcpProjectByEnvironmentId(ctx context.Context, in *model.EnvironmentId, opts ...grpc.CallOption) (*gcp.GcpProject, error) {
	out := new(gcp.GcpProject)
	err := c.cc.Invoke(ctx, EnvironmentQueryController_GetSecretsGcpProjectByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentQueryControllerClient) FindWorkloadPodsByEnvironmentId(ctx context.Context, in *model.EnvironmentId, opts ...grpc.CallOption) (*model3.WorkloadPods, error) {
	out := new(model3.WorkloadPods)
	err := c.cc.Invoke(ctx, EnvironmentQueryController_FindWorkloadPodsByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentQueryControllerClient) FindWorkloadNamespacesByEnvironmentId(ctx context.Context, in *model.EnvironmentId, opts ...grpc.CallOption) (*model3.WorkloadNamespaces, error) {
	out := new(model3.WorkloadNamespaces)
	err := c.cc.Invoke(ctx, EnvironmentQueryController_FindWorkloadNamespacesByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentQueryControllerServer is the server API for EnvironmentQueryController service.
// All implementations should embed UnimplementedEnvironmentQueryControllerServer
// for forward compatibility
type EnvironmentQueryControllerServer interface {
	// list all environments on planton cloud for the requested page. This is intended for use on portal.
	List(context.Context, *rpc.PageInfo) (*model.EnvironmentList, error)
	// look up environment using environment id
	GetById(context.Context, *model.EnvironmentId) (*model.Environment, error)
	// find environments by product id
	FindByProductId(context.Context, *model1.ProductId) (*model.Environments, error)
	// find environments by kube-cluster id
	FindByKubeClusterId(context.Context, *model2.KubeClusterId) (*model.Environments, error)
	// look up environment using environment id
	GetByProductIdAndEnvironmentName(context.Context, *model.GetByProductIdAndEnvironmentNameQueryInput) (*model.Environment, error)
	// look up the build engine environment for the product required for launching the microservice in build engine
	GetBuildEngineEnvironmentByProductId(context.Context, *model1.ProductId) (*model.Environment, error)
	// look up the gcp project details by environment id required for fetching secrets for launching project in build engine.
	GetSecretsGcpProjectByEnvironmentId(context.Context, *model.EnvironmentId) (*gcp.GcpProject, error)
	// find workload pods part of environment
	FindWorkloadPodsByEnvironmentId(context.Context, *model.EnvironmentId) (*model3.WorkloadPods, error)
	// find workload namespaces in a environment.
	FindWorkloadNamespacesByEnvironmentId(context.Context, *model.EnvironmentId) (*model3.WorkloadNamespaces, error)
}

// UnimplementedEnvironmentQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedEnvironmentQueryControllerServer struct {
}

func (UnimplementedEnvironmentQueryControllerServer) List(context.Context, *rpc.PageInfo) (*model.EnvironmentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedEnvironmentQueryControllerServer) GetById(context.Context, *model.EnvironmentId) (*model.Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedEnvironmentQueryControllerServer) FindByProductId(context.Context, *model1.ProductId) (*model.Environments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedEnvironmentQueryControllerServer) FindByKubeClusterId(context.Context, *model2.KubeClusterId) (*model.Environments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKubeClusterId not implemented")
}
func (UnimplementedEnvironmentQueryControllerServer) GetByProductIdAndEnvironmentName(context.Context, *model.GetByProductIdAndEnvironmentNameQueryInput) (*model.Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByProductIdAndEnvironmentName not implemented")
}
func (UnimplementedEnvironmentQueryControllerServer) GetBuildEngineEnvironmentByProductId(context.Context, *model1.ProductId) (*model.Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuildEngineEnvironmentByProductId not implemented")
}
func (UnimplementedEnvironmentQueryControllerServer) GetSecretsGcpProjectByEnvironmentId(context.Context, *model.EnvironmentId) (*gcp.GcpProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecretsGcpProjectByEnvironmentId not implemented")
}
func (UnimplementedEnvironmentQueryControllerServer) FindWorkloadPodsByEnvironmentId(context.Context, *model.EnvironmentId) (*model3.WorkloadPods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindWorkloadPodsByEnvironmentId not implemented")
}
func (UnimplementedEnvironmentQueryControllerServer) FindWorkloadNamespacesByEnvironmentId(context.Context, *model.EnvironmentId) (*model3.WorkloadNamespaces, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindWorkloadNamespacesByEnvironmentId not implemented")
}

// UnsafeEnvironmentQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvironmentQueryControllerServer will
// result in compilation errors.
type UnsafeEnvironmentQueryControllerServer interface {
	mustEmbedUnimplementedEnvironmentQueryControllerServer()
}

func RegisterEnvironmentQueryControllerServer(s grpc.ServiceRegistrar, srv EnvironmentQueryControllerServer) {
	s.RegisterService(&EnvironmentQueryController_ServiceDesc, srv)
}

func _EnvironmentQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentQueryControllerServer).List(ctx, req.(*rpc.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EnvironmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentQueryControllerServer).GetById(ctx, req.(*model.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentQueryControllerServer).FindByProductId(ctx, req.(*model1.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentQueryController_FindByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentQueryControllerServer).FindByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentQueryController_FindByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentQueryControllerServer).FindByKubeClusterId(ctx, req.(*model2.KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentQueryController_GetByProductIdAndEnvironmentName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetByProductIdAndEnvironmentNameQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentQueryControllerServer).GetByProductIdAndEnvironmentName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentQueryController_GetByProductIdAndEnvironmentName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentQueryControllerServer).GetByProductIdAndEnvironmentName(ctx, req.(*model.GetByProductIdAndEnvironmentNameQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentQueryController_GetBuildEngineEnvironmentByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentQueryControllerServer).GetBuildEngineEnvironmentByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentQueryController_GetBuildEngineEnvironmentByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentQueryControllerServer).GetBuildEngineEnvironmentByProductId(ctx, req.(*model1.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentQueryController_GetSecretsGcpProjectByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EnvironmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentQueryControllerServer).GetSecretsGcpProjectByEnvironmentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentQueryController_GetSecretsGcpProjectByEnvironmentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentQueryControllerServer).GetSecretsGcpProjectByEnvironmentId(ctx, req.(*model.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentQueryController_FindWorkloadPodsByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EnvironmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentQueryControllerServer).FindWorkloadPodsByEnvironmentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentQueryController_FindWorkloadPodsByEnvironmentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentQueryControllerServer).FindWorkloadPodsByEnvironmentId(ctx, req.(*model.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentQueryController_FindWorkloadNamespacesByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EnvironmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentQueryControllerServer).FindWorkloadNamespacesByEnvironmentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentQueryController_FindWorkloadNamespacesByEnvironmentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentQueryControllerServer).FindWorkloadNamespacesByEnvironmentId(ctx, req.(*model.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

// EnvironmentQueryController_ServiceDesc is the grpc.ServiceDesc for EnvironmentQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnvironmentQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentQueryController",
	HandlerType: (*EnvironmentQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _EnvironmentQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _EnvironmentQueryController_GetById_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _EnvironmentQueryController_FindByProductId_Handler,
		},
		{
			MethodName: "findByKubeClusterId",
			Handler:    _EnvironmentQueryController_FindByKubeClusterId_Handler,
		},
		{
			MethodName: "getByProductIdAndEnvironmentName",
			Handler:    _EnvironmentQueryController_GetByProductIdAndEnvironmentName_Handler,
		},
		{
			MethodName: "getBuildEngineEnvironmentByProductId",
			Handler:    _EnvironmentQueryController_GetBuildEngineEnvironmentByProductId_Handler,
		},
		{
			MethodName: "getSecretsGcpProjectByEnvironmentId",
			Handler:    _EnvironmentQueryController_GetSecretsGcpProjectByEnvironmentId_Handler,
		},
		{
			MethodName: "findWorkloadPodsByEnvironmentId",
			Handler:    _EnvironmentQueryController_FindWorkloadPodsByEnvironmentId_Handler,
		},
		{
			MethodName: "findWorkloadNamespacesByEnvironmentId",
			Handler:    _EnvironmentQueryController_FindWorkloadNamespacesByEnvironmentId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/environment/service/query.proto",
}

const (
	EnvironmentSecretQueryController_GetById_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentSecretQueryController/getById"
	EnvironmentSecretQueryController_GetValue_FullMethodName = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentSecretQueryController/getValue"
)

// EnvironmentSecretQueryControllerClient is the client API for EnvironmentSecretQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentSecretQueryControllerClient interface {
	// lookup product secrets using product secret id
	GetById(ctx context.Context, in *model.GetByEnvironmentSecretIdInput, opts ...grpc.CallOption) (*model.EnvironmentSecret, error)
	// get value of a product secret
	GetValue(ctx context.Context, in *model.GetByEnvironmentSecretIdInput, opts ...grpc.CallOption) (*model.EnvironmentSecretValue, error)
}

type environmentSecretQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentSecretQueryControllerClient(cc grpc.ClientConnInterface) EnvironmentSecretQueryControllerClient {
	return &environmentSecretQueryControllerClient{cc}
}

func (c *environmentSecretQueryControllerClient) GetById(ctx context.Context, in *model.GetByEnvironmentSecretIdInput, opts ...grpc.CallOption) (*model.EnvironmentSecret, error) {
	out := new(model.EnvironmentSecret)
	err := c.cc.Invoke(ctx, EnvironmentSecretQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentSecretQueryControllerClient) GetValue(ctx context.Context, in *model.GetByEnvironmentSecretIdInput, opts ...grpc.CallOption) (*model.EnvironmentSecretValue, error) {
	out := new(model.EnvironmentSecretValue)
	err := c.cc.Invoke(ctx, EnvironmentSecretQueryController_GetValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentSecretQueryControllerServer is the server API for EnvironmentSecretQueryController service.
// All implementations should embed UnimplementedEnvironmentSecretQueryControllerServer
// for forward compatibility
type EnvironmentSecretQueryControllerServer interface {
	// lookup product secrets using product secret id
	GetById(context.Context, *model.GetByEnvironmentSecretIdInput) (*model.EnvironmentSecret, error)
	// get value of a product secret
	GetValue(context.Context, *model.GetByEnvironmentSecretIdInput) (*model.EnvironmentSecretValue, error)
}

// UnimplementedEnvironmentSecretQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedEnvironmentSecretQueryControllerServer struct {
}

func (UnimplementedEnvironmentSecretQueryControllerServer) GetById(context.Context, *model.GetByEnvironmentSecretIdInput) (*model.EnvironmentSecret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedEnvironmentSecretQueryControllerServer) GetValue(context.Context, *model.GetByEnvironmentSecretIdInput) (*model.EnvironmentSecretValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}

// UnsafeEnvironmentSecretQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvironmentSecretQueryControllerServer will
// result in compilation errors.
type UnsafeEnvironmentSecretQueryControllerServer interface {
	mustEmbedUnimplementedEnvironmentSecretQueryControllerServer()
}

func RegisterEnvironmentSecretQueryControllerServer(s grpc.ServiceRegistrar, srv EnvironmentSecretQueryControllerServer) {
	s.RegisterService(&EnvironmentSecretQueryController_ServiceDesc, srv)
}

func _EnvironmentSecretQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetByEnvironmentSecretIdInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentSecretQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentSecretQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentSecretQueryControllerServer).GetById(ctx, req.(*model.GetByEnvironmentSecretIdInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentSecretQueryController_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetByEnvironmentSecretIdInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentSecretQueryControllerServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentSecretQueryController_GetValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentSecretQueryControllerServer).GetValue(ctx, req.(*model.GetByEnvironmentSecretIdInput))
	}
	return interceptor(ctx, in, info, handler)
}

// EnvironmentSecretQueryController_ServiceDesc is the grpc.ServiceDesc for EnvironmentSecretQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnvironmentSecretQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentSecretQueryController",
	HandlerType: (*EnvironmentSecretQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _EnvironmentSecretQueryController_GetById_Handler,
		},
		{
			MethodName: "getValue",
			Handler:    _EnvironmentSecretQueryController_GetValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/environment/service/query.proto",
}

const (
	EnvironmentVariableQueryController_GetById_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentVariableQueryController/getById"
	EnvironmentVariableQueryController_GetValue_FullMethodName = "/cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentVariableQueryController/getValue"
)

// EnvironmentVariableQueryControllerClient is the client API for EnvironmentVariableQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentVariableQueryControllerClient interface {
	// lookup environment variables using environment variable id
	GetById(ctx context.Context, in *model.GetByEnvironmentVariableIdInput, opts ...grpc.CallOption) (*model.EnvironmentVariable, error)
	// get value of a environment variable
	GetValue(ctx context.Context, in *model.GetByEnvironmentVariableIdInput, opts ...grpc.CallOption) (*model.EnvironmentVariableValue, error)
}

type environmentVariableQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentVariableQueryControllerClient(cc grpc.ClientConnInterface) EnvironmentVariableQueryControllerClient {
	return &environmentVariableQueryControllerClient{cc}
}

func (c *environmentVariableQueryControllerClient) GetById(ctx context.Context, in *model.GetByEnvironmentVariableIdInput, opts ...grpc.CallOption) (*model.EnvironmentVariable, error) {
	out := new(model.EnvironmentVariable)
	err := c.cc.Invoke(ctx, EnvironmentVariableQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentVariableQueryControllerClient) GetValue(ctx context.Context, in *model.GetByEnvironmentVariableIdInput, opts ...grpc.CallOption) (*model.EnvironmentVariableValue, error) {
	out := new(model.EnvironmentVariableValue)
	err := c.cc.Invoke(ctx, EnvironmentVariableQueryController_GetValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentVariableQueryControllerServer is the server API for EnvironmentVariableQueryController service.
// All implementations should embed UnimplementedEnvironmentVariableQueryControllerServer
// for forward compatibility
type EnvironmentVariableQueryControllerServer interface {
	// lookup environment variables using environment variable id
	GetById(context.Context, *model.GetByEnvironmentVariableIdInput) (*model.EnvironmentVariable, error)
	// get value of a environment variable
	GetValue(context.Context, *model.GetByEnvironmentVariableIdInput) (*model.EnvironmentVariableValue, error)
}

// UnimplementedEnvironmentVariableQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedEnvironmentVariableQueryControllerServer struct {
}

func (UnimplementedEnvironmentVariableQueryControllerServer) GetById(context.Context, *model.GetByEnvironmentVariableIdInput) (*model.EnvironmentVariable, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedEnvironmentVariableQueryControllerServer) GetValue(context.Context, *model.GetByEnvironmentVariableIdInput) (*model.EnvironmentVariableValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}

// UnsafeEnvironmentVariableQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvironmentVariableQueryControllerServer will
// result in compilation errors.
type UnsafeEnvironmentVariableQueryControllerServer interface {
	mustEmbedUnimplementedEnvironmentVariableQueryControllerServer()
}

func RegisterEnvironmentVariableQueryControllerServer(s grpc.ServiceRegistrar, srv EnvironmentVariableQueryControllerServer) {
	s.RegisterService(&EnvironmentVariableQueryController_ServiceDesc, srv)
}

func _EnvironmentVariableQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetByEnvironmentVariableIdInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentVariableQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentVariableQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentVariableQueryControllerServer).GetById(ctx, req.(*model.GetByEnvironmentVariableIdInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentVariableQueryController_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetByEnvironmentVariableIdInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentVariableQueryControllerServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentVariableQueryController_GetValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentVariableQueryControllerServer).GetValue(ctx, req.(*model.GetByEnvironmentVariableIdInput))
	}
	return interceptor(ctx, in, info, handler)
}

// EnvironmentVariableQueryController_ServiceDesc is the grpc.ServiceDesc for EnvironmentVariableQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnvironmentVariableQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.environment.service.EnvironmentVariableQueryController",
	HandlerType: (*EnvironmentVariableQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _EnvironmentVariableQueryController_GetById_Handler,
		},
		{
			MethodName: "getValue",
			Handler:    _EnvironmentVariableQueryController_GetValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/environment/service/query.proto",
}
