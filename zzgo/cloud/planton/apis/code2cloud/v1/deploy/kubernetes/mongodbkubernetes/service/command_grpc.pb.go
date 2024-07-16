// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/kubernetes/mongodbkubernetes/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/kubernetes/mongodbkubernetes/model"
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
	MongodbKubernetesCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/previewCreate"
	MongodbKubernetesCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/create"
	MongodbKubernetesCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/previewUpdate"
	MongodbKubernetesCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/update"
	MongodbKubernetesCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/previewDelete"
	MongodbKubernetesCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/delete"
	MongodbKubernetesCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/previewRestore"
	MongodbKubernetesCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/restore"
	MongodbKubernetesCommandController_Restart_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/restart"
	MongodbKubernetesCommandController_Pause_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/pause"
	MongodbKubernetesCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/unpause"
	MongodbKubernetesCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/previewRefresh"
	MongodbKubernetesCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController/refresh"
)

// MongodbKubernetesCommandControllerClient is the client API for MongodbKubernetesCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MongodbKubernetesCommandControllerClient interface {
	// preview creating mongodb-kubernetes
	PreviewCreate(ctx context.Context, in *model.MongodbKubernetes, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// create mongodb-kubernetes
	Create(ctx context.Context, in *model.MongodbKubernetes, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// preview updating an existing mongodb-kubernetes
	PreviewUpdate(ctx context.Context, in *model.MongodbKubernetes, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// update an existing mongodb-kubernetes
	Update(ctx context.Context, in *model.MongodbKubernetes, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// preview deleting an existing mongodb-kubernetes
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// delete an existing mongodb-kubernetes
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// preview restoring a previously deleted mongodb-kubernetes
	PreviewRestore(ctx context.Context, in *model.MongodbKubernetes, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// restore a previously deleted mongodb-kubernetes
	Restore(ctx context.Context, in *model.MongodbKubernetes, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// restart a mongodb-kubernetes running in a environment.
	// mongodb-kubernetes is restarted by deleting running "mongodb" pods which will be automatically recreated by kubernetes
	Restart(ctx context.Context, in *model.MongodbKubernetesId, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// pause a mongodb-kubernetes running in a environment.
	// mongodb-kubernetes is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// unpause a previously paused mongodb-kubernetes running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the mongodb-kubernetes.
	Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// preview refresh a mongodb-kubernetes that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// refresh a mongodb-kubernetes that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
}

type mongodbKubernetesCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMongodbKubernetesCommandControllerClient(cc grpc.ClientConnInterface) MongodbKubernetesCommandControllerClient {
	return &mongodbKubernetesCommandControllerClient{cc}
}

func (c *mongodbKubernetesCommandControllerClient) PreviewCreate(ctx context.Context, in *model.MongodbKubernetes, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesCommandControllerClient) Create(ctx context.Context, in *model.MongodbKubernetes, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.MongodbKubernetes, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesCommandControllerClient) Update(ctx context.Context, in *model.MongodbKubernetes, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesCommandControllerClient) PreviewRestore(ctx context.Context, in *model.MongodbKubernetes, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesCommandControllerClient) Restore(ctx context.Context, in *model.MongodbKubernetes, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesCommandControllerClient) Restart(ctx context.Context, in *model.MongodbKubernetesId, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesCommandControllerClient) Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesCommandControllerClient) Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MongodbKubernetesCommandControllerServer is the server API for MongodbKubernetesCommandController service.
// All implementations should embed UnimplementedMongodbKubernetesCommandControllerServer
// for forward compatibility
type MongodbKubernetesCommandControllerServer interface {
	// preview creating mongodb-kubernetes
	PreviewCreate(context.Context, *model.MongodbKubernetes) (*model.MongodbKubernetes, error)
	// create mongodb-kubernetes
	Create(context.Context, *model.MongodbKubernetes) (*model.MongodbKubernetes, error)
	// preview updating an existing mongodb-kubernetes
	PreviewUpdate(context.Context, *model.MongodbKubernetes) (*model.MongodbKubernetes, error)
	// update an existing mongodb-kubernetes
	Update(context.Context, *model.MongodbKubernetes) (*model.MongodbKubernetes, error)
	// preview deleting an existing mongodb-kubernetes
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.MongodbKubernetes, error)
	// delete an existing mongodb-kubernetes
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.MongodbKubernetes, error)
	// preview restoring a previously deleted mongodb-kubernetes
	PreviewRestore(context.Context, *model.MongodbKubernetes) (*model.MongodbKubernetes, error)
	// restore a previously deleted mongodb-kubernetes
	Restore(context.Context, *model.MongodbKubernetes) (*model.MongodbKubernetes, error)
	// restart a mongodb-kubernetes running in a environment.
	// mongodb-kubernetes is restarted by deleting running "mongodb" pods which will be automatically recreated by kubernetes
	Restart(context.Context, *model.MongodbKubernetesId) (*model.MongodbKubernetes, error)
	// pause a mongodb-kubernetes running in a environment.
	// mongodb-kubernetes is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.MongodbKubernetes, error)
	// unpause a previously paused mongodb-kubernetes running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the mongodb-kubernetes.
	Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.MongodbKubernetes, error)
	// preview refresh a mongodb-kubernetes that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.MongodbKubernetes, error)
	// refresh a mongodb-kubernetes that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.MongodbKubernetes, error)
}

// UnimplementedMongodbKubernetesCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedMongodbKubernetesCommandControllerServer struct {
}

func (UnimplementedMongodbKubernetesCommandControllerServer) PreviewCreate(context.Context, *model.MongodbKubernetes) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedMongodbKubernetesCommandControllerServer) Create(context.Context, *model.MongodbKubernetes) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMongodbKubernetesCommandControllerServer) PreviewUpdate(context.Context, *model.MongodbKubernetes) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedMongodbKubernetesCommandControllerServer) Update(context.Context, *model.MongodbKubernetes) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMongodbKubernetesCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedMongodbKubernetesCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMongodbKubernetesCommandControllerServer) PreviewRestore(context.Context, *model.MongodbKubernetes) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedMongodbKubernetesCommandControllerServer) Restore(context.Context, *model.MongodbKubernetes) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedMongodbKubernetesCommandControllerServer) Restart(context.Context, *model.MongodbKubernetesId) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedMongodbKubernetesCommandControllerServer) Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedMongodbKubernetesCommandControllerServer) Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}
func (UnimplementedMongodbKubernetesCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedMongodbKubernetesCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeMongodbKubernetesCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MongodbKubernetesCommandControllerServer will
// result in compilation errors.
type UnsafeMongodbKubernetesCommandControllerServer interface {
	mustEmbedUnimplementedMongodbKubernetesCommandControllerServer()
}

func RegisterMongodbKubernetesCommandControllerServer(s grpc.ServiceRegistrar, srv MongodbKubernetesCommandControllerServer) {
	s.RegisterService(&MongodbKubernetesCommandController_ServiceDesc, srv)
}

func _MongodbKubernetesCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).PreviewCreate(ctx, req.(*model.MongodbKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).Create(ctx, req.(*model.MongodbKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).PreviewUpdate(ctx, req.(*model.MongodbKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).Update(ctx, req.(*model.MongodbKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).PreviewRestore(ctx, req.(*model.MongodbKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).Restore(ctx, req.(*model.MongodbKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).Restart(ctx, req.(*model.MongodbKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourcePauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).Pause(ctx, req.(*model1.ApiResourcePauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceUnPauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).Unpause(ctx, req.(*model1.ApiResourceUnPauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MongodbKubernetesCommandController_ServiceDesc is the grpc.ServiceDesc for MongodbKubernetesCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MongodbKubernetesCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.service.MongodbKubernetesCommandController",
	HandlerType: (*MongodbKubernetesCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _MongodbKubernetesCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _MongodbKubernetesCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _MongodbKubernetesCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _MongodbKubernetesCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _MongodbKubernetesCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _MongodbKubernetesCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _MongodbKubernetesCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _MongodbKubernetesCommandController_Restore_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _MongodbKubernetesCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _MongodbKubernetesCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _MongodbKubernetesCommandController_Unpause_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _MongodbKubernetesCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _MongodbKubernetesCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/kubernetes/mongodbkubernetes/service/command.proto",
}
