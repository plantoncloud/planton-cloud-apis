// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/resourcemanager/v1/configmap/service/command.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/resourcemanager/v1/configmap/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ConfigMapCommandController_Create_FullMethodName  = "/cloud.planton.apis.resourcemanager.v1.configmap.service.ConfigMapCommandController/create"
	ConfigMapCommandController_Update_FullMethodName  = "/cloud.planton.apis.resourcemanager.v1.configmap.service.ConfigMapCommandController/update"
	ConfigMapCommandController_Delete_FullMethodName  = "/cloud.planton.apis.resourcemanager.v1.configmap.service.ConfigMapCommandController/delete"
	ConfigMapCommandController_Restore_FullMethodName = "/cloud.planton.apis.resourcemanager.v1.configmap.service.ConfigMapCommandController/restore"
)

// ConfigMapCommandControllerClient is the client API for ConfigMapCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigMapCommandControllerClient interface {
	// create a new config-map on planton-cloud
	Create(ctx context.Context, in *model.ConfigMap, opts ...grpc.CallOption) (*model.ConfigMap, error)
	// update an existing config-map on planton-cloud
	Update(ctx context.Context, in *model.ConfigMap, opts ...grpc.CallOption) (*model.ConfigMap, error)
	// delete an existing config-map on planton-cloud using config-map id
	Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.ConfigMap, error)
	// restore a previously deleted config-map.
	Restore(ctx context.Context, in *model.ConfigMap, opts ...grpc.CallOption) (*model.ConfigMap, error)
}

type configMapCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigMapCommandControllerClient(cc grpc.ClientConnInterface) ConfigMapCommandControllerClient {
	return &configMapCommandControllerClient{cc}
}

func (c *configMapCommandControllerClient) Create(ctx context.Context, in *model.ConfigMap, opts ...grpc.CallOption) (*model.ConfigMap, error) {
	out := new(model.ConfigMap)
	err := c.cc.Invoke(ctx, ConfigMapCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMapCommandControllerClient) Update(ctx context.Context, in *model.ConfigMap, opts ...grpc.CallOption) (*model.ConfigMap, error) {
	out := new(model.ConfigMap)
	err := c.cc.Invoke(ctx, ConfigMapCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMapCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.ConfigMap, error) {
	out := new(model.ConfigMap)
	err := c.cc.Invoke(ctx, ConfigMapCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configMapCommandControllerClient) Restore(ctx context.Context, in *model.ConfigMap, opts ...grpc.CallOption) (*model.ConfigMap, error) {
	out := new(model.ConfigMap)
	err := c.cc.Invoke(ctx, ConfigMapCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigMapCommandControllerServer is the server API for ConfigMapCommandController service.
// All implementations should embed UnimplementedConfigMapCommandControllerServer
// for forward compatibility
type ConfigMapCommandControllerServer interface {
	// create a new config-map on planton-cloud
	Create(context.Context, *model.ConfigMap) (*model.ConfigMap, error)
	// update an existing config-map on planton-cloud
	Update(context.Context, *model.ConfigMap) (*model.ConfigMap, error)
	// delete an existing config-map on planton-cloud using config-map id
	Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.ConfigMap, error)
	// restore a previously deleted config-map.
	Restore(context.Context, *model.ConfigMap) (*model.ConfigMap, error)
}

// UnimplementedConfigMapCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedConfigMapCommandControllerServer struct {
}

func (UnimplementedConfigMapCommandControllerServer) Create(context.Context, *model.ConfigMap) (*model.ConfigMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedConfigMapCommandControllerServer) Update(context.Context, *model.ConfigMap) (*model.ConfigMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedConfigMapCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.ConfigMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedConfigMapCommandControllerServer) Restore(context.Context, *model.ConfigMap) (*model.ConfigMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeConfigMapCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigMapCommandControllerServer will
// result in compilation errors.
type UnsafeConfigMapCommandControllerServer interface {
	mustEmbedUnimplementedConfigMapCommandControllerServer()
}

func RegisterConfigMapCommandControllerServer(s grpc.ServiceRegistrar, srv ConfigMapCommandControllerServer) {
	s.RegisterService(&ConfigMapCommandController_ServiceDesc, srv)
}

func _ConfigMapCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ConfigMap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMapCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigMapCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMapCommandControllerServer).Create(ctx, req.(*model.ConfigMap))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMapCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ConfigMap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMapCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigMapCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMapCommandControllerServer).Update(ctx, req.(*model.ConfigMap))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMapCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMapCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigMapCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMapCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigMapCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ConfigMap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigMapCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigMapCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigMapCommandControllerServer).Restore(ctx, req.(*model.ConfigMap))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigMapCommandController_ServiceDesc is the grpc.ServiceDesc for ConfigMapCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigMapCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.resourcemanager.v1.configmap.service.ConfigMapCommandController",
	HandlerType: (*ConfigMapCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _ConfigMapCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _ConfigMapCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _ConfigMapCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _ConfigMapCommandController_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/resourcemanager/v1/configmap/service/command.proto",
}