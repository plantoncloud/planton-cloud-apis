// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/resourcemanager/v1/environment/service/command.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
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
	EnvironmentCommandController_Create_FullMethodName  = "/cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentCommandController/create"
	EnvironmentCommandController_Update_FullMethodName  = "/cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentCommandController/update"
	EnvironmentCommandController_Delete_FullMethodName  = "/cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentCommandController/delete"
	EnvironmentCommandController_Restore_FullMethodName = "/cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentCommandController/restore"
)

// EnvironmentCommandControllerClient is the client API for EnvironmentCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentCommandControllerClient interface {
	// create environment
	Create(ctx context.Context, in *model.Environment, opts ...grpc.CallOption) (*model.Environment, error)
	// update an existing environment
	Update(ctx context.Context, in *model.Environment, opts ...grpc.CallOption) (*model.Environment, error)
	// delete an existing environment
	// deleting a environment involves cleaning of stack-modules deployed to that environment.
	// microservices, secrets, postgres-clusters, kafka-cluster should be cleaned up in the corresponding environment
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.Environment, error)
	// restore a deleted environment
	// restoring a environment tries to restore all the individual resources that were destroyed as part of the delete operation.
	Restore(ctx context.Context, in *model.Environment, opts ...grpc.CallOption) (*model.Environment, error)
}

type environmentCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentCommandControllerClient(cc grpc.ClientConnInterface) EnvironmentCommandControllerClient {
	return &environmentCommandControllerClient{cc}
}

func (c *environmentCommandControllerClient) Create(ctx context.Context, in *model.Environment, opts ...grpc.CallOption) (*model.Environment, error) {
	out := new(model.Environment)
	err := c.cc.Invoke(ctx, EnvironmentCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentCommandControllerClient) Update(ctx context.Context, in *model.Environment, opts ...grpc.CallOption) (*model.Environment, error) {
	out := new(model.Environment)
	err := c.cc.Invoke(ctx, EnvironmentCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.Environment, error) {
	out := new(model.Environment)
	err := c.cc.Invoke(ctx, EnvironmentCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentCommandControllerClient) Restore(ctx context.Context, in *model.Environment, opts ...grpc.CallOption) (*model.Environment, error) {
	out := new(model.Environment)
	err := c.cc.Invoke(ctx, EnvironmentCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentCommandControllerServer is the server API for EnvironmentCommandController service.
// All implementations should embed UnimplementedEnvironmentCommandControllerServer
// for forward compatibility
type EnvironmentCommandControllerServer interface {
	// create environment
	Create(context.Context, *model.Environment) (*model.Environment, error)
	// update an existing environment
	Update(context.Context, *model.Environment) (*model.Environment, error)
	// delete an existing environment
	// deleting a environment involves cleaning of stack-modules deployed to that environment.
	// microservices, secrets, postgres-clusters, kafka-cluster should be cleaned up in the corresponding environment
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.Environment, error)
	// restore a deleted environment
	// restoring a environment tries to restore all the individual resources that were destroyed as part of the delete operation.
	Restore(context.Context, *model.Environment) (*model.Environment, error)
}

// UnimplementedEnvironmentCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedEnvironmentCommandControllerServer struct {
}

func (UnimplementedEnvironmentCommandControllerServer) Create(context.Context, *model.Environment) (*model.Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEnvironmentCommandControllerServer) Update(context.Context, *model.Environment) (*model.Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEnvironmentCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEnvironmentCommandControllerServer) Restore(context.Context, *model.Environment) (*model.Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeEnvironmentCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvironmentCommandControllerServer will
// result in compilation errors.
type UnsafeEnvironmentCommandControllerServer interface {
	mustEmbedUnimplementedEnvironmentCommandControllerServer()
}

func RegisterEnvironmentCommandControllerServer(s grpc.ServiceRegistrar, srv EnvironmentCommandControllerServer) {
	s.RegisterService(&EnvironmentCommandController_ServiceDesc, srv)
}

func _EnvironmentCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Environment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentCommandControllerServer).Create(ctx, req.(*model.Environment))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Environment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentCommandControllerServer).Update(ctx, req.(*model.Environment))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Environment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentCommandControllerServer).Restore(ctx, req.(*model.Environment))
	}
	return interceptor(ctx, in, info, handler)
}

// EnvironmentCommandController_ServiceDesc is the grpc.ServiceDesc for EnvironmentCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnvironmentCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.resourcemanager.v1.environment.service.EnvironmentCommandController",
	HandlerType: (*EnvironmentCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _EnvironmentCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _EnvironmentCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _EnvironmentCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _EnvironmentCommandController_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/resourcemanager/v1/environment/service/command.proto",
}
