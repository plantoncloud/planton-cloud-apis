// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/microservice/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/microservice/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MicroserviceInstanceCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/previewCreate"
	MicroserviceInstanceCommandController_Create_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/create"
	MicroserviceInstanceCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/previewUpdate"
	MicroserviceInstanceCommandController_Update_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/update"
	MicroserviceInstanceCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/previewDelete"
	MicroserviceInstanceCommandController_Delete_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/delete"
	MicroserviceInstanceCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/previewRestore"
	MicroserviceInstanceCommandController_Restore_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/restore"
	MicroserviceInstanceCommandController_CreateStackJob_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/createStackJob"
	MicroserviceInstanceCommandController_Restart_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/restart"
	MicroserviceInstanceCommandController_Pause_FullMethodName          = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/pause"
	MicroserviceInstanceCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/unpause"
	MicroserviceInstanceCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/previewRefresh"
	MicroserviceInstanceCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController/refresh"
)

// MicroserviceInstanceCommandControllerClient is the client API for MicroserviceInstanceCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicroserviceInstanceCommandControllerClient interface {
	// preview create microservice-instance
	PreviewCreate(ctx context.Context, in *model.MicroserviceInstance, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// create microservice-instance
	Create(ctx context.Context, in *model.MicroserviceInstance, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// preview update microservice-instance
	PreviewUpdate(ctx context.Context, in *model.MicroserviceInstance, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// update microservice-instance
	Update(ctx context.Context, in *model.MicroserviceInstance, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// preview delete microservice-instance
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// delete microservice-instance
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// preview restoring a deleted microservice-instance
	PreviewRestore(ctx context.Context, in *model.MicroserviceInstance, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// restore a deleted microservice-instance of a environment.
	Restore(ctx context.Context, in *model.MicroserviceInstance, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// create-stack-job for microservice-instance
	CreateStackJob(ctx context.Context, in *model2.CreateStackJobCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// restart a microservice-instance running in a environment.
	// microservice-instance is restarted by deleting running pods which will be automatically recreated by kubernetes.
	Restart(ctx context.Context, in *model.MicroserviceInstanceId, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// pause a microservice-instance running in a environment.
	// microservice-instance is paused by scaling down number of replicas of
	// the kubernetes deployment/stateful sets to zero in the environment.
	Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// unpause a previously paused microservice-instance running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the microservice-instance.
	Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// preview refresh a microservice-instance that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// refresh a microservice-instance that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
}

type microserviceInstanceCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroserviceInstanceCommandControllerClient(cc grpc.ClientConnInterface) MicroserviceInstanceCommandControllerClient {
	return &microserviceInstanceCommandControllerClient{cc}
}

func (c *microserviceInstanceCommandControllerClient) PreviewCreate(ctx context.Context, in *model.MicroserviceInstance, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) Create(ctx context.Context, in *model.MicroserviceInstance, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.MicroserviceInstance, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) Update(ctx context.Context, in *model.MicroserviceInstance, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) PreviewRestore(ctx context.Context, in *model.MicroserviceInstance, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) Restore(ctx context.Context, in *model.MicroserviceInstance, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) CreateStackJob(ctx context.Context, in *model2.CreateStackJobCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_CreateStackJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) Restart(ctx context.Context, in *model.MicroserviceInstanceId, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroserviceInstanceCommandControllerServer is the server API for MicroserviceInstanceCommandController service.
// All implementations should embed UnimplementedMicroserviceInstanceCommandControllerServer
// for forward compatibility
type MicroserviceInstanceCommandControllerServer interface {
	// preview create microservice-instance
	PreviewCreate(context.Context, *model.MicroserviceInstance) (*model.MicroserviceInstance, error)
	// create microservice-instance
	Create(context.Context, *model.MicroserviceInstance) (*model.MicroserviceInstance, error)
	// preview update microservice-instance
	PreviewUpdate(context.Context, *model.MicroserviceInstance) (*model.MicroserviceInstance, error)
	// update microservice-instance
	Update(context.Context, *model.MicroserviceInstance) (*model.MicroserviceInstance, error)
	// preview delete microservice-instance
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.MicroserviceInstance, error)
	// delete microservice-instance
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.MicroserviceInstance, error)
	// preview restoring a deleted microservice-instance
	PreviewRestore(context.Context, *model.MicroserviceInstance) (*model.MicroserviceInstance, error)
	// restore a deleted microservice-instance of a environment.
	Restore(context.Context, *model.MicroserviceInstance) (*model.MicroserviceInstance, error)
	// create-stack-job for microservice-instance
	CreateStackJob(context.Context, *model2.CreateStackJobCommandInput) (*model.MicroserviceInstance, error)
	// restart a microservice-instance running in a environment.
	// microservice-instance is restarted by deleting running pods which will be automatically recreated by kubernetes.
	Restart(context.Context, *model.MicroserviceInstanceId) (*model.MicroserviceInstance, error)
	// pause a microservice-instance running in a environment.
	// microservice-instance is paused by scaling down number of replicas of
	// the kubernetes deployment/stateful sets to zero in the environment.
	Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.MicroserviceInstance, error)
	// unpause a previously paused microservice-instance running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the microservice-instance.
	Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.MicroserviceInstance, error)
	// preview refresh a microservice-instance that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.MicroserviceInstance, error)
	// refresh a microservice-instance that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.MicroserviceInstance, error)
}

// UnimplementedMicroserviceInstanceCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedMicroserviceInstanceCommandControllerServer struct {
}

func (UnimplementedMicroserviceInstanceCommandControllerServer) PreviewCreate(context.Context, *model.MicroserviceInstance) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) Create(context.Context, *model.MicroserviceInstance) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) PreviewUpdate(context.Context, *model.MicroserviceInstance) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) Update(context.Context, *model.MicroserviceInstance) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) PreviewRestore(context.Context, *model.MicroserviceInstance) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) Restore(context.Context, *model.MicroserviceInstance) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) CreateStackJob(context.Context, *model2.CreateStackJobCommandInput) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStackJob not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) Restart(context.Context, *model.MicroserviceInstanceId) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedMicroserviceInstanceCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeMicroserviceInstanceCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroserviceInstanceCommandControllerServer will
// result in compilation errors.
type UnsafeMicroserviceInstanceCommandControllerServer interface {
	mustEmbedUnimplementedMicroserviceInstanceCommandControllerServer()
}

func RegisterMicroserviceInstanceCommandControllerServer(s grpc.ServiceRegistrar, srv MicroserviceInstanceCommandControllerServer) {
	s.RegisterService(&MicroserviceInstanceCommandController_ServiceDesc, srv)
}

func _MicroserviceInstanceCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MicroserviceInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).PreviewCreate(ctx, req.(*model.MicroserviceInstance))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MicroserviceInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).Create(ctx, req.(*model.MicroserviceInstance))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MicroserviceInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).PreviewUpdate(ctx, req.(*model.MicroserviceInstance))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MicroserviceInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).Update(ctx, req.(*model.MicroserviceInstance))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MicroserviceInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).PreviewRestore(ctx, req.(*model.MicroserviceInstance))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MicroserviceInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).Restore(ctx, req.(*model.MicroserviceInstance))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_CreateStackJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.CreateStackJobCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).CreateStackJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_CreateStackJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).CreateStackJob(ctx, req.(*model2.CreateStackJobCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MicroserviceInstanceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).Restart(ctx, req.(*model.MicroserviceInstanceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourcePauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).Pause(ctx, req.(*model1.ApiResourcePauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceUnPauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).Unpause(ctx, req.(*model1.ApiResourceUnPauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MicroserviceInstanceCommandController_ServiceDesc is the grpc.ServiceDesc for MicroserviceInstanceCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MicroserviceInstanceCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.microservice.service.MicroserviceInstanceCommandController",
	HandlerType: (*MicroserviceInstanceCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _MicroserviceInstanceCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _MicroserviceInstanceCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _MicroserviceInstanceCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _MicroserviceInstanceCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _MicroserviceInstanceCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _MicroserviceInstanceCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _MicroserviceInstanceCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _MicroserviceInstanceCommandController_Restore_Handler,
		},
		{
			MethodName: "createStackJob",
			Handler:    _MicroserviceInstanceCommandController_CreateStackJob_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _MicroserviceInstanceCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _MicroserviceInstanceCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _MicroserviceInstanceCommandController_Unpause_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _MicroserviceInstanceCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _MicroserviceInstanceCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/microservice/service/command.proto",
}
