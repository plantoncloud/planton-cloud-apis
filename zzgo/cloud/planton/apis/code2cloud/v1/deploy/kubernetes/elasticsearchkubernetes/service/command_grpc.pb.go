// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/kubernetes/elasticsearchkubernetes/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/kubernetes/elasticsearchkubernetes/model"
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
	ElasticsearchKubernetesCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/previewCreate"
	ElasticsearchKubernetesCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/create"
	ElasticsearchKubernetesCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/previewUpdate"
	ElasticsearchKubernetesCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/update"
	ElasticsearchKubernetesCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/previewDelete"
	ElasticsearchKubernetesCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/delete"
	ElasticsearchKubernetesCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/previewRestore"
	ElasticsearchKubernetesCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/restore"
	ElasticsearchKubernetesCommandController_Restart_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/restart"
	ElasticsearchKubernetesCommandController_Pause_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/pause"
	ElasticsearchKubernetesCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/unpause"
	ElasticsearchKubernetesCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/previewRefresh"
	ElasticsearchKubernetesCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController/refresh"
)

// ElasticsearchKubernetesCommandControllerClient is the client API for ElasticsearchKubernetesCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElasticsearchKubernetesCommandControllerClient interface {
	// preview creating elasticsearch-kubernetes
	PreviewCreate(ctx context.Context, in *model.ElasticsearchKubernetes, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// create elasticsearch-kubernetes
	Create(ctx context.Context, in *model.ElasticsearchKubernetes, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// preview updating an existing elasticsearch-kubernetes
	PreviewUpdate(ctx context.Context, in *model.ElasticsearchKubernetes, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// update an existing elasticsearch-kubernetes
	Update(ctx context.Context, in *model.ElasticsearchKubernetes, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// preview deleting an existing elasticsearch-kubernetes
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// delete an existing elasticsearch-kubernetes
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// preview restoring a previously deleted elasticsearch-kubernetes
	PreviewRestore(ctx context.Context, in *model.ElasticsearchKubernetes, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// restore a previously deleted elasticsearch-kubernetes
	Restore(ctx context.Context, in *model.ElasticsearchKubernetes, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// restart a elasticsearch-kubernetes running in a environment.
	// elasticsearch-kubernetes is restarted by deleting running "elasticsearch" pods which will be automatically recreated by kubernetes
	Restart(ctx context.Context, in *model.ElasticsearchKubernetesId, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// pause a elasticsearch-kubernetes running in a environment.
	// elasticsearch-kubernetes is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// unpause a previously paused elasticsearch-kubernetes running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the elasticsearch-kubernetes.
	Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// preview refresh a elasticsearch-kubernetes that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// refresh a elasticsearch-kubernetes that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
}

type elasticsearchKubernetesCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewElasticsearchKubernetesCommandControllerClient(cc grpc.ClientConnInterface) ElasticsearchKubernetesCommandControllerClient {
	return &elasticsearchKubernetesCommandControllerClient{cc}
}

func (c *elasticsearchKubernetesCommandControllerClient) PreviewCreate(ctx context.Context, in *model.ElasticsearchKubernetes, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesCommandControllerClient) Create(ctx context.Context, in *model.ElasticsearchKubernetes, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.ElasticsearchKubernetes, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesCommandControllerClient) Update(ctx context.Context, in *model.ElasticsearchKubernetes, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesCommandControllerClient) PreviewRestore(ctx context.Context, in *model.ElasticsearchKubernetes, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesCommandControllerClient) Restore(ctx context.Context, in *model.ElasticsearchKubernetes, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesCommandControllerClient) Restart(ctx context.Context, in *model.ElasticsearchKubernetesId, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesCommandControllerClient) Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesCommandControllerClient) Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElasticsearchKubernetesCommandControllerServer is the server API for ElasticsearchKubernetesCommandController service.
// All implementations should embed UnimplementedElasticsearchKubernetesCommandControllerServer
// for forward compatibility
type ElasticsearchKubernetesCommandControllerServer interface {
	// preview creating elasticsearch-kubernetes
	PreviewCreate(context.Context, *model.ElasticsearchKubernetes) (*model.ElasticsearchKubernetes, error)
	// create elasticsearch-kubernetes
	Create(context.Context, *model.ElasticsearchKubernetes) (*model.ElasticsearchKubernetes, error)
	// preview updating an existing elasticsearch-kubernetes
	PreviewUpdate(context.Context, *model.ElasticsearchKubernetes) (*model.ElasticsearchKubernetes, error)
	// update an existing elasticsearch-kubernetes
	Update(context.Context, *model.ElasticsearchKubernetes) (*model.ElasticsearchKubernetes, error)
	// preview deleting an existing elasticsearch-kubernetes
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.ElasticsearchKubernetes, error)
	// delete an existing elasticsearch-kubernetes
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.ElasticsearchKubernetes, error)
	// preview restoring a previously deleted elasticsearch-kubernetes
	PreviewRestore(context.Context, *model.ElasticsearchKubernetes) (*model.ElasticsearchKubernetes, error)
	// restore a previously deleted elasticsearch-kubernetes
	Restore(context.Context, *model.ElasticsearchKubernetes) (*model.ElasticsearchKubernetes, error)
	// restart a elasticsearch-kubernetes running in a environment.
	// elasticsearch-kubernetes is restarted by deleting running "elasticsearch" pods which will be automatically recreated by kubernetes
	Restart(context.Context, *model.ElasticsearchKubernetesId) (*model.ElasticsearchKubernetes, error)
	// pause a elasticsearch-kubernetes running in a environment.
	// elasticsearch-kubernetes is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.ElasticsearchKubernetes, error)
	// unpause a previously paused elasticsearch-kubernetes running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the elasticsearch-kubernetes.
	Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.ElasticsearchKubernetes, error)
	// preview refresh a elasticsearch-kubernetes that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.ElasticsearchKubernetes, error)
	// refresh a elasticsearch-kubernetes that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.ElasticsearchKubernetes, error)
}

// UnimplementedElasticsearchKubernetesCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedElasticsearchKubernetesCommandControllerServer struct {
}

func (UnimplementedElasticsearchKubernetesCommandControllerServer) PreviewCreate(context.Context, *model.ElasticsearchKubernetes) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedElasticsearchKubernetesCommandControllerServer) Create(context.Context, *model.ElasticsearchKubernetes) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedElasticsearchKubernetesCommandControllerServer) PreviewUpdate(context.Context, *model.ElasticsearchKubernetes) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedElasticsearchKubernetesCommandControllerServer) Update(context.Context, *model.ElasticsearchKubernetes) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedElasticsearchKubernetesCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedElasticsearchKubernetesCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedElasticsearchKubernetesCommandControllerServer) PreviewRestore(context.Context, *model.ElasticsearchKubernetes) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedElasticsearchKubernetesCommandControllerServer) Restore(context.Context, *model.ElasticsearchKubernetes) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedElasticsearchKubernetesCommandControllerServer) Restart(context.Context, *model.ElasticsearchKubernetesId) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedElasticsearchKubernetesCommandControllerServer) Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedElasticsearchKubernetesCommandControllerServer) Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}
func (UnimplementedElasticsearchKubernetesCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedElasticsearchKubernetesCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeElasticsearchKubernetesCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElasticsearchKubernetesCommandControllerServer will
// result in compilation errors.
type UnsafeElasticsearchKubernetesCommandControllerServer interface {
	mustEmbedUnimplementedElasticsearchKubernetesCommandControllerServer()
}

func RegisterElasticsearchKubernetesCommandControllerServer(s grpc.ServiceRegistrar, srv ElasticsearchKubernetesCommandControllerServer) {
	s.RegisterService(&ElasticsearchKubernetesCommandController_ServiceDesc, srv)
}

func _ElasticsearchKubernetesCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ElasticsearchKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).PreviewCreate(ctx, req.(*model.ElasticsearchKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ElasticsearchKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Create(ctx, req.(*model.ElasticsearchKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ElasticsearchKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).PreviewUpdate(ctx, req.(*model.ElasticsearchKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ElasticsearchKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Update(ctx, req.(*model.ElasticsearchKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ElasticsearchKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).PreviewRestore(ctx, req.(*model.ElasticsearchKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ElasticsearchKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Restore(ctx, req.(*model.ElasticsearchKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ElasticsearchKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Restart(ctx, req.(*model.ElasticsearchKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourcePauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Pause(ctx, req.(*model1.ApiResourcePauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceUnPauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Unpause(ctx, req.(*model1.ApiResourceUnPauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ElasticsearchKubernetesCommandController_ServiceDesc is the grpc.ServiceDesc for ElasticsearchKubernetesCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ElasticsearchKubernetesCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesCommandController",
	HandlerType: (*ElasticsearchKubernetesCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _ElasticsearchKubernetesCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _ElasticsearchKubernetesCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _ElasticsearchKubernetesCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _ElasticsearchKubernetesCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _ElasticsearchKubernetesCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _ElasticsearchKubernetesCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _ElasticsearchKubernetesCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _ElasticsearchKubernetesCommandController_Restore_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _ElasticsearchKubernetesCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _ElasticsearchKubernetesCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _ElasticsearchKubernetesCommandController_Unpause_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _ElasticsearchKubernetesCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _ElasticsearchKubernetesCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/kubernetes/elasticsearchkubernetes/service/command.proto",
}