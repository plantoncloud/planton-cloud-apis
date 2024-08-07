// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/kubernetes/rediskubernetes/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/rediskubernetes/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RedisKubernetesCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesCommandController/previewCreate"
	RedisKubernetesCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesCommandController/create"
	RedisKubernetesCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesCommandController/previewUpdate"
	RedisKubernetesCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesCommandController/update"
	RedisKubernetesCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesCommandController/previewDelete"
	RedisKubernetesCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesCommandController/delete"
	RedisKubernetesCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesCommandController/previewRestore"
	RedisKubernetesCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesCommandController/restore"
	RedisKubernetesCommandController_Restart_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesCommandController/restart"
	RedisKubernetesCommandController_Pause_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesCommandController/pause"
	RedisKubernetesCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesCommandController/unpause"
)

// RedisKubernetesCommandControllerClient is the client API for RedisKubernetesCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RedisKubernetesCommandControllerClient interface {
	// preview creating redis-kubernetes
	PreviewCreate(ctx context.Context, in *model.RedisKubernetes, opts ...grpc.CallOption) (*model.RedisKubernetes, error)
	// create redis-kubernetes
	Create(ctx context.Context, in *model.RedisKubernetes, opts ...grpc.CallOption) (*model.RedisKubernetes, error)
	// preview updating an existing redis-kubernetes
	PreviewUpdate(ctx context.Context, in *model.RedisKubernetes, opts ...grpc.CallOption) (*model.RedisKubernetes, error)
	// update an existing redis-kubernetes
	Update(ctx context.Context, in *model.RedisKubernetes, opts ...grpc.CallOption) (*model.RedisKubernetes, error)
	// preview deleting an existing redis-kubernetes
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.RedisKubernetes, error)
	// delete an existing redis-kubernetes
	Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.RedisKubernetes, error)
	// preview restoring a previously deleted redis-kubernetes
	PreviewRestore(ctx context.Context, in *model.RedisKubernetes, opts ...grpc.CallOption) (*model.RedisKubernetes, error)
	// restore a previously deleted redis-kubernetes
	Restore(ctx context.Context, in *model.RedisKubernetes, opts ...grpc.CallOption) (*model.RedisKubernetes, error)
	// restart a redis-kubernetes running in a environment.
	// redis-kubernetes is restarted by deleting running "redis" pods which will be automatically recreated by kubernetes
	Restart(ctx context.Context, in *model.RedisKubernetesId, opts ...grpc.CallOption) (*model.RedisKubernetes, error)
	// pause a redis-kubernetes running in a environment.
	// redis-kubernetes is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(ctx context.Context, in *model1.ApiResourcePauseInput, opts ...grpc.CallOption) (*model.RedisKubernetes, error)
	// unpause a previously paused redis-kubernetes running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the redis-kubernetes.
	Unpause(ctx context.Context, in *model1.ApiResourceUnPauseInput, opts ...grpc.CallOption) (*model.RedisKubernetes, error)
}

type redisKubernetesCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewRedisKubernetesCommandControllerClient(cc grpc.ClientConnInterface) RedisKubernetesCommandControllerClient {
	return &redisKubernetesCommandControllerClient{cc}
}

func (c *redisKubernetesCommandControllerClient) PreviewCreate(ctx context.Context, in *model.RedisKubernetes, opts ...grpc.CallOption) (*model.RedisKubernetes, error) {
	out := new(model.RedisKubernetes)
	err := c.cc.Invoke(ctx, RedisKubernetesCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisKubernetesCommandControllerClient) Create(ctx context.Context, in *model.RedisKubernetes, opts ...grpc.CallOption) (*model.RedisKubernetes, error) {
	out := new(model.RedisKubernetes)
	err := c.cc.Invoke(ctx, RedisKubernetesCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisKubernetesCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.RedisKubernetes, opts ...grpc.CallOption) (*model.RedisKubernetes, error) {
	out := new(model.RedisKubernetes)
	err := c.cc.Invoke(ctx, RedisKubernetesCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisKubernetesCommandControllerClient) Update(ctx context.Context, in *model.RedisKubernetes, opts ...grpc.CallOption) (*model.RedisKubernetes, error) {
	out := new(model.RedisKubernetes)
	err := c.cc.Invoke(ctx, RedisKubernetesCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisKubernetesCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.RedisKubernetes, error) {
	out := new(model.RedisKubernetes)
	err := c.cc.Invoke(ctx, RedisKubernetesCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisKubernetesCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.RedisKubernetes, error) {
	out := new(model.RedisKubernetes)
	err := c.cc.Invoke(ctx, RedisKubernetesCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisKubernetesCommandControllerClient) PreviewRestore(ctx context.Context, in *model.RedisKubernetes, opts ...grpc.CallOption) (*model.RedisKubernetes, error) {
	out := new(model.RedisKubernetes)
	err := c.cc.Invoke(ctx, RedisKubernetesCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisKubernetesCommandControllerClient) Restore(ctx context.Context, in *model.RedisKubernetes, opts ...grpc.CallOption) (*model.RedisKubernetes, error) {
	out := new(model.RedisKubernetes)
	err := c.cc.Invoke(ctx, RedisKubernetesCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisKubernetesCommandControllerClient) Restart(ctx context.Context, in *model.RedisKubernetesId, opts ...grpc.CallOption) (*model.RedisKubernetes, error) {
	out := new(model.RedisKubernetes)
	err := c.cc.Invoke(ctx, RedisKubernetesCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisKubernetesCommandControllerClient) Pause(ctx context.Context, in *model1.ApiResourcePauseInput, opts ...grpc.CallOption) (*model.RedisKubernetes, error) {
	out := new(model.RedisKubernetes)
	err := c.cc.Invoke(ctx, RedisKubernetesCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisKubernetesCommandControllerClient) Unpause(ctx context.Context, in *model1.ApiResourceUnPauseInput, opts ...grpc.CallOption) (*model.RedisKubernetes, error) {
	out := new(model.RedisKubernetes)
	err := c.cc.Invoke(ctx, RedisKubernetesCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedisKubernetesCommandControllerServer is the server API for RedisKubernetesCommandController service.
// All implementations should embed UnimplementedRedisKubernetesCommandControllerServer
// for forward compatibility
type RedisKubernetesCommandControllerServer interface {
	// preview creating redis-kubernetes
	PreviewCreate(context.Context, *model.RedisKubernetes) (*model.RedisKubernetes, error)
	// create redis-kubernetes
	Create(context.Context, *model.RedisKubernetes) (*model.RedisKubernetes, error)
	// preview updating an existing redis-kubernetes
	PreviewUpdate(context.Context, *model.RedisKubernetes) (*model.RedisKubernetes, error)
	// update an existing redis-kubernetes
	Update(context.Context, *model.RedisKubernetes) (*model.RedisKubernetes, error)
	// preview deleting an existing redis-kubernetes
	PreviewDelete(context.Context, *model1.ApiResourceDeleteInput) (*model.RedisKubernetes, error)
	// delete an existing redis-kubernetes
	Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.RedisKubernetes, error)
	// preview restoring a previously deleted redis-kubernetes
	PreviewRestore(context.Context, *model.RedisKubernetes) (*model.RedisKubernetes, error)
	// restore a previously deleted redis-kubernetes
	Restore(context.Context, *model.RedisKubernetes) (*model.RedisKubernetes, error)
	// restart a redis-kubernetes running in a environment.
	// redis-kubernetes is restarted by deleting running "redis" pods which will be automatically recreated by kubernetes
	Restart(context.Context, *model.RedisKubernetesId) (*model.RedisKubernetes, error)
	// pause a redis-kubernetes running in a environment.
	// redis-kubernetes is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(context.Context, *model1.ApiResourcePauseInput) (*model.RedisKubernetes, error)
	// unpause a previously paused redis-kubernetes running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the redis-kubernetes.
	Unpause(context.Context, *model1.ApiResourceUnPauseInput) (*model.RedisKubernetes, error)
}

// UnimplementedRedisKubernetesCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedRedisKubernetesCommandControllerServer struct {
}

func (UnimplementedRedisKubernetesCommandControllerServer) PreviewCreate(context.Context, *model.RedisKubernetes) (*model.RedisKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedRedisKubernetesCommandControllerServer) Create(context.Context, *model.RedisKubernetes) (*model.RedisKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRedisKubernetesCommandControllerServer) PreviewUpdate(context.Context, *model.RedisKubernetes) (*model.RedisKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedRedisKubernetesCommandControllerServer) Update(context.Context, *model.RedisKubernetes) (*model.RedisKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRedisKubernetesCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteInput) (*model.RedisKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedRedisKubernetesCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.RedisKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRedisKubernetesCommandControllerServer) PreviewRestore(context.Context, *model.RedisKubernetes) (*model.RedisKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedRedisKubernetesCommandControllerServer) Restore(context.Context, *model.RedisKubernetes) (*model.RedisKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedRedisKubernetesCommandControllerServer) Restart(context.Context, *model.RedisKubernetesId) (*model.RedisKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedRedisKubernetesCommandControllerServer) Pause(context.Context, *model1.ApiResourcePauseInput) (*model.RedisKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedRedisKubernetesCommandControllerServer) Unpause(context.Context, *model1.ApiResourceUnPauseInput) (*model.RedisKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}

// UnsafeRedisKubernetesCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedisKubernetesCommandControllerServer will
// result in compilation errors.
type UnsafeRedisKubernetesCommandControllerServer interface {
	mustEmbedUnimplementedRedisKubernetesCommandControllerServer()
}

func RegisterRedisKubernetesCommandControllerServer(s grpc.ServiceRegistrar, srv RedisKubernetesCommandControllerServer) {
	s.RegisterService(&RedisKubernetesCommandController_ServiceDesc, srv)
}

func _RedisKubernetesCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.RedisKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisKubernetesCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisKubernetesCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisKubernetesCommandControllerServer).PreviewCreate(ctx, req.(*model.RedisKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisKubernetesCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.RedisKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisKubernetesCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisKubernetesCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisKubernetesCommandControllerServer).Create(ctx, req.(*model.RedisKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisKubernetesCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.RedisKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisKubernetesCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisKubernetesCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisKubernetesCommandControllerServer).PreviewUpdate(ctx, req.(*model.RedisKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisKubernetesCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.RedisKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisKubernetesCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisKubernetesCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisKubernetesCommandControllerServer).Update(ctx, req.(*model.RedisKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisKubernetesCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisKubernetesCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisKubernetesCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisKubernetesCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisKubernetesCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisKubernetesCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisKubernetesCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisKubernetesCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisKubernetesCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.RedisKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisKubernetesCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisKubernetesCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisKubernetesCommandControllerServer).PreviewRestore(ctx, req.(*model.RedisKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisKubernetesCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.RedisKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisKubernetesCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisKubernetesCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisKubernetesCommandControllerServer).Restore(ctx, req.(*model.RedisKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisKubernetesCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.RedisKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisKubernetesCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisKubernetesCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisKubernetesCommandControllerServer).Restart(ctx, req.(*model.RedisKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisKubernetesCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourcePauseInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisKubernetesCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisKubernetesCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisKubernetesCommandControllerServer).Pause(ctx, req.(*model1.ApiResourcePauseInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisKubernetesCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceUnPauseInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisKubernetesCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisKubernetesCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisKubernetesCommandControllerServer).Unpause(ctx, req.(*model1.ApiResourceUnPauseInput))
	}
	return interceptor(ctx, in, info, handler)
}

// RedisKubernetesCommandController_ServiceDesc is the grpc.ServiceDesc for RedisKubernetesCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedisKubernetesCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesCommandController",
	HandlerType: (*RedisKubernetesCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _RedisKubernetesCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _RedisKubernetesCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _RedisKubernetesCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _RedisKubernetesCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _RedisKubernetesCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _RedisKubernetesCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _RedisKubernetesCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _RedisKubernetesCommandController_Restore_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _RedisKubernetesCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _RedisKubernetesCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _RedisKubernetesCommandController_Unpause_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/kubernetes/rediskubernetes/service/command.proto",
}
