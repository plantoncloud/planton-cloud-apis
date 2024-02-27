// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/mongodbcluster/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/mongodbcluster/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MongodbClusterCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/previewCreate"
	MongodbClusterCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/create"
	MongodbClusterCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/previewUpdate"
	MongodbClusterCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/update"
	MongodbClusterCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/previewDelete"
	MongodbClusterCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/delete"
	MongodbClusterCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/previewRestore"
	MongodbClusterCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/restore"
	MongodbClusterCommandController_CreateStackJob_FullMethodName = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/createStackJob"
	MongodbClusterCommandController_Restart_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/restart"
	MongodbClusterCommandController_Pause_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/pause"
	MongodbClusterCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/unpause"
	MongodbClusterCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/previewRefresh"
	MongodbClusterCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController/refresh"
)

// MongodbClusterCommandControllerClient is the client API for MongodbClusterCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MongodbClusterCommandControllerClient interface {
	// preview creating mongodb-cluster
	PreviewCreate(ctx context.Context, in *model.MongodbCluster, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// create mongodb-cluster
	Create(ctx context.Context, in *model.MongodbCluster, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// preview updating an existing mongodb-cluster
	PreviewUpdate(ctx context.Context, in *model.MongodbCluster, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// update an existing mongodb-cluster
	Update(ctx context.Context, in *model.MongodbCluster, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// preview deleting an existing mongodb-cluster
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// delete an existing mongodb-cluster
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// preview restoring a previously deleted mongodb-cluster
	PreviewRestore(ctx context.Context, in *model.MongodbCluster, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// restore a previously deleted mongodb-cluster
	Restore(ctx context.Context, in *model.MongodbCluster, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// create-stack-job for mongodb-cluster
	CreateStackJob(ctx context.Context, in *model2.CreateStackJobCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// restart a mongodb-cluster running in a environment.
	// mongodb-cluster is restarted by deleting running "mongodb" pods which will be automatically recreated by kubernetes
	Restart(ctx context.Context, in *model.MongodbClusterId, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// pause a mongodb-cluster running in a environment.
	// mongodb-cluster is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// unpause a previously paused mongodb-cluster running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the mongodb-cluster.
	Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// preview refresh a mongodb-cluster that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// refresh a mongodb-cluster that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error)
}

type mongodbClusterCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMongodbClusterCommandControllerClient(cc grpc.ClientConnInterface) MongodbClusterCommandControllerClient {
	return &mongodbClusterCommandControllerClient{cc}
}

func (c *mongodbClusterCommandControllerClient) PreviewCreate(ctx context.Context, in *model.MongodbCluster, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) Create(ctx context.Context, in *model.MongodbCluster, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.MongodbCluster, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) Update(ctx context.Context, in *model.MongodbCluster, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) PreviewRestore(ctx context.Context, in *model.MongodbCluster, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) Restore(ctx context.Context, in *model.MongodbCluster, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) CreateStackJob(ctx context.Context, in *model2.CreateStackJobCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_CreateStackJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) Restart(ctx context.Context, in *model.MongodbClusterId, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MongodbClusterCommandControllerServer is the server API for MongodbClusterCommandController service.
// All implementations should embed UnimplementedMongodbClusterCommandControllerServer
// for forward compatibility
type MongodbClusterCommandControllerServer interface {
	// preview creating mongodb-cluster
	PreviewCreate(context.Context, *model.MongodbCluster) (*model.MongodbCluster, error)
	// create mongodb-cluster
	Create(context.Context, *model.MongodbCluster) (*model.MongodbCluster, error)
	// preview updating an existing mongodb-cluster
	PreviewUpdate(context.Context, *model.MongodbCluster) (*model.MongodbCluster, error)
	// update an existing mongodb-cluster
	Update(context.Context, *model.MongodbCluster) (*model.MongodbCluster, error)
	// preview deleting an existing mongodb-cluster
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.MongodbCluster, error)
	// delete an existing mongodb-cluster
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.MongodbCluster, error)
	// preview restoring a previously deleted mongodb-cluster
	PreviewRestore(context.Context, *model.MongodbCluster) (*model.MongodbCluster, error)
	// restore a previously deleted mongodb-cluster
	Restore(context.Context, *model.MongodbCluster) (*model.MongodbCluster, error)
	// create-stack-job for mongodb-cluster
	CreateStackJob(context.Context, *model2.CreateStackJobCommandInput) (*model.MongodbCluster, error)
	// restart a mongodb-cluster running in a environment.
	// mongodb-cluster is restarted by deleting running "mongodb" pods which will be automatically recreated by kubernetes
	Restart(context.Context, *model.MongodbClusterId) (*model.MongodbCluster, error)
	// pause a mongodb-cluster running in a environment.
	// mongodb-cluster is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.MongodbCluster, error)
	// unpause a previously paused mongodb-cluster running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the mongodb-cluster.
	Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.MongodbCluster, error)
	// preview refresh a mongodb-cluster that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.MongodbCluster, error)
	// refresh a mongodb-cluster that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.MongodbCluster, error)
}

// UnimplementedMongodbClusterCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedMongodbClusterCommandControllerServer struct {
}

func (UnimplementedMongodbClusterCommandControllerServer) PreviewCreate(context.Context, *model.MongodbCluster) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) Create(context.Context, *model.MongodbCluster) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) PreviewUpdate(context.Context, *model.MongodbCluster) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) Update(context.Context, *model.MongodbCluster) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) PreviewRestore(context.Context, *model.MongodbCluster) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) Restore(context.Context, *model.MongodbCluster) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) CreateStackJob(context.Context, *model2.CreateStackJobCommandInput) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStackJob not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) Restart(context.Context, *model.MongodbClusterId) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedMongodbClusterCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeMongodbClusterCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MongodbClusterCommandControllerServer will
// result in compilation errors.
type UnsafeMongodbClusterCommandControllerServer interface {
	mustEmbedUnimplementedMongodbClusterCommandControllerServer()
}

func RegisterMongodbClusterCommandControllerServer(s grpc.ServiceRegistrar, srv MongodbClusterCommandControllerServer) {
	s.RegisterService(&MongodbClusterCommandController_ServiceDesc, srv)
}

func _MongodbClusterCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbCluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).PreviewCreate(ctx, req.(*model.MongodbCluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbCluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).Create(ctx, req.(*model.MongodbCluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbCluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).PreviewUpdate(ctx, req.(*model.MongodbCluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbCluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).Update(ctx, req.(*model.MongodbCluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbCluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).PreviewRestore(ctx, req.(*model.MongodbCluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbCluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).Restore(ctx, req.(*model.MongodbCluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_CreateStackJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.CreateStackJobCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).CreateStackJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_CreateStackJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).CreateStackJob(ctx, req.(*model2.CreateStackJobCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).Restart(ctx, req.(*model.MongodbClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourcePauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).Pause(ctx, req.(*model1.ApiResourcePauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceUnPauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).Unpause(ctx, req.(*model1.ApiResourceUnPauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MongodbClusterCommandController_ServiceDesc is the grpc.ServiceDesc for MongodbClusterCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MongodbClusterCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterCommandController",
	HandlerType: (*MongodbClusterCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _MongodbClusterCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _MongodbClusterCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _MongodbClusterCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _MongodbClusterCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _MongodbClusterCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _MongodbClusterCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _MongodbClusterCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _MongodbClusterCommandController_Restore_Handler,
		},
		{
			MethodName: "createStackJob",
			Handler:    _MongodbClusterCommandController_CreateStackJob_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _MongodbClusterCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _MongodbClusterCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _MongodbClusterCommandController_Unpause_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _MongodbClusterCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _MongodbClusterCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/mongodbcluster/service/command.proto",
}
