// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/resourcemanager/v1/environment/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/resourcemanager/v1/environment/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EnvironmentQueryController_Get_FullMethodName                  = "/cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentQueryController/get"
	EnvironmentQueryController_GetByOrgIdAndEnvName_FullMethodName = "/cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentQueryController/getByOrgIdAndEnvName"
)

// EnvironmentQueryControllerClient is the client API for EnvironmentQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentQueryControllerClient interface {
	// look up environment using environment id
	Get(ctx context.Context, in *model.EnvId, opts ...grpc.CallOption) (*model.Environment, error)
	// look up environment using by organization-id and environment name
	GetByOrgIdAndEnvName(ctx context.Context, in *model.GetByOrgIdAndEnvNameQueryInput, opts ...grpc.CallOption) (*model.Environment, error)
}

type environmentQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentQueryControllerClient(cc grpc.ClientConnInterface) EnvironmentQueryControllerClient {
	return &environmentQueryControllerClient{cc}
}

func (c *environmentQueryControllerClient) Get(ctx context.Context, in *model.EnvId, opts ...grpc.CallOption) (*model.Environment, error) {
	out := new(model.Environment)
	err := c.cc.Invoke(ctx, EnvironmentQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentQueryControllerClient) GetByOrgIdAndEnvName(ctx context.Context, in *model.GetByOrgIdAndEnvNameQueryInput, opts ...grpc.CallOption) (*model.Environment, error) {
	out := new(model.Environment)
	err := c.cc.Invoke(ctx, EnvironmentQueryController_GetByOrgIdAndEnvName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentQueryControllerServer is the server API for EnvironmentQueryController service.
// All implementations should embed UnimplementedEnvironmentQueryControllerServer
// for forward compatibility
type EnvironmentQueryControllerServer interface {
	// look up environment using environment id
	Get(context.Context, *model.EnvId) (*model.Environment, error)
	// look up environment using by organization-id and environment name
	GetByOrgIdAndEnvName(context.Context, *model.GetByOrgIdAndEnvNameQueryInput) (*model.Environment, error)
}

// UnimplementedEnvironmentQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedEnvironmentQueryControllerServer struct {
}

func (UnimplementedEnvironmentQueryControllerServer) Get(context.Context, *model.EnvId) (*model.Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedEnvironmentQueryControllerServer) GetByOrgIdAndEnvName(context.Context, *model.GetByOrgIdAndEnvNameQueryInput) (*model.Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByOrgIdAndEnvName not implemented")
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

func _EnvironmentQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EnvId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentQueryControllerServer).Get(ctx, req.(*model.EnvId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentQueryController_GetByOrgIdAndEnvName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetByOrgIdAndEnvNameQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentQueryControllerServer).GetByOrgIdAndEnvName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentQueryController_GetByOrgIdAndEnvName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentQueryControllerServer).GetByOrgIdAndEnvName(ctx, req.(*model.GetByOrgIdAndEnvNameQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

// EnvironmentQueryController_ServiceDesc is the grpc.ServiceDesc for EnvironmentQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnvironmentQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentQueryController",
	HandlerType: (*EnvironmentQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _EnvironmentQueryController_Get_Handler,
		},
		{
			MethodName: "getByOrgIdAndEnvName",
			Handler:    _EnvironmentQueryController_GetByOrgIdAndEnvName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/resourcemanager/v1/environment/service/query.proto",
}

const (
	EnvironmentSecretQueryController_Get_FullMethodName      = "/cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentSecretQueryController/get"
	EnvironmentSecretQueryController_GetValue_FullMethodName = "/cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentSecretQueryController/getValue"
)

// EnvironmentSecretQueryControllerClient is the client API for EnvironmentSecretQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentSecretQueryControllerClient interface {
	// lookup environment secret using environment secret id
	Get(ctx context.Context, in *model.GetByEnvironmentSecretIdInput, opts ...grpc.CallOption) (*model.EnvironmentSecret, error)
	// get value of a environment secret
	GetValue(ctx context.Context, in *model.GetByEnvironmentSecretIdInput, opts ...grpc.CallOption) (*model.EnvironmentSecretValue, error)
}

type environmentSecretQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentSecretQueryControllerClient(cc grpc.ClientConnInterface) EnvironmentSecretQueryControllerClient {
	return &environmentSecretQueryControllerClient{cc}
}

func (c *environmentSecretQueryControllerClient) Get(ctx context.Context, in *model.GetByEnvironmentSecretIdInput, opts ...grpc.CallOption) (*model.EnvironmentSecret, error) {
	out := new(model.EnvironmentSecret)
	err := c.cc.Invoke(ctx, EnvironmentSecretQueryController_Get_FullMethodName, in, out, opts...)
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
	// lookup environment secret using environment secret id
	Get(context.Context, *model.GetByEnvironmentSecretIdInput) (*model.EnvironmentSecret, error)
	// get value of a environment secret
	GetValue(context.Context, *model.GetByEnvironmentSecretIdInput) (*model.EnvironmentSecretValue, error)
}

// UnimplementedEnvironmentSecretQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedEnvironmentSecretQueryControllerServer struct {
}

func (UnimplementedEnvironmentSecretQueryControllerServer) Get(context.Context, *model.GetByEnvironmentSecretIdInput) (*model.EnvironmentSecret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
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

func _EnvironmentSecretQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetByEnvironmentSecretIdInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentSecretQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentSecretQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentSecretQueryControllerServer).Get(ctx, req.(*model.GetByEnvironmentSecretIdInput))
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
	ServiceName: "cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentSecretQueryController",
	HandlerType: (*EnvironmentSecretQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _EnvironmentSecretQueryController_Get_Handler,
		},
		{
			MethodName: "getValue",
			Handler:    _EnvironmentSecretQueryController_GetValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/resourcemanager/v1/environment/service/query.proto",
}

const (
	EnvironmentVariableQueryController_Get_FullMethodName      = "/cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentVariableQueryController/get"
	EnvironmentVariableQueryController_GetValue_FullMethodName = "/cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentVariableQueryController/getValue"
)

// EnvironmentVariableQueryControllerClient is the client API for EnvironmentVariableQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentVariableQueryControllerClient interface {
	// lookup environment variables using environment variable id
	Get(ctx context.Context, in *model.GetByEnvironmentVariableIdInput, opts ...grpc.CallOption) (*model.EnvironmentVariable, error)
	// get value of a environment variable
	GetValue(ctx context.Context, in *model.GetByEnvironmentVariableIdInput, opts ...grpc.CallOption) (*model.EnvironmentVariableValue, error)
}

type environmentVariableQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentVariableQueryControllerClient(cc grpc.ClientConnInterface) EnvironmentVariableQueryControllerClient {
	return &environmentVariableQueryControllerClient{cc}
}

func (c *environmentVariableQueryControllerClient) Get(ctx context.Context, in *model.GetByEnvironmentVariableIdInput, opts ...grpc.CallOption) (*model.EnvironmentVariable, error) {
	out := new(model.EnvironmentVariable)
	err := c.cc.Invoke(ctx, EnvironmentVariableQueryController_Get_FullMethodName, in, out, opts...)
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
	Get(context.Context, *model.GetByEnvironmentVariableIdInput) (*model.EnvironmentVariable, error)
	// get value of a environment variable
	GetValue(context.Context, *model.GetByEnvironmentVariableIdInput) (*model.EnvironmentVariableValue, error)
}

// UnimplementedEnvironmentVariableQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedEnvironmentVariableQueryControllerServer struct {
}

func (UnimplementedEnvironmentVariableQueryControllerServer) Get(context.Context, *model.GetByEnvironmentVariableIdInput) (*model.EnvironmentVariable, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
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

func _EnvironmentVariableQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetByEnvironmentVariableIdInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentVariableQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentVariableQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentVariableQueryControllerServer).Get(ctx, req.(*model.GetByEnvironmentVariableIdInput))
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
	ServiceName: "cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentVariableQueryController",
	HandlerType: (*EnvironmentVariableQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _EnvironmentVariableQueryController_Get_Handler,
		},
		{
			MethodName: "getValue",
			Handler:    _EnvironmentVariableQueryController_GetValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/resourcemanager/v1/environment/service/query.proto",
}