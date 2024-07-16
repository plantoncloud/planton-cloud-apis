// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/kubernetes/kafkakubernetes/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/kubernetes/kafkakubernetes/model"
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
	KafkaKubernetesCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/previewCreate"
	KafkaKubernetesCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/create"
	KafkaKubernetesCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/previewUpdate"
	KafkaKubernetesCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/update"
	KafkaKubernetesCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/previewDelete"
	KafkaKubernetesCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/delete"
	KafkaKubernetesCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/previewRestore"
	KafkaKubernetesCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/restore"
	KafkaKubernetesCommandController_Restart_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/restart"
	KafkaKubernetesCommandController_Pause_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/pause"
	KafkaKubernetesCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/unpause"
	KafkaKubernetesCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/previewRefresh"
	KafkaKubernetesCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController/refresh"
)

// KafkaKubernetesCommandControllerClient is the client API for KafkaKubernetesCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KafkaKubernetesCommandControllerClient interface {
	// preview create kafka-kubernetes
	PreviewCreate(ctx context.Context, in *model.KafkaKubernetes, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// create kafka-kubernetes
	Create(ctx context.Context, in *model.KafkaKubernetes, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// preview update an existing kafka-kubernetes
	PreviewUpdate(ctx context.Context, in *model.KafkaKubernetes, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// update an existing kafka-kubernetes
	Update(ctx context.Context, in *model.KafkaKubernetes, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// preview deleting an existing kafka-kubernetes
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// delete an existing kafka-kubernetes
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// preview restoring a deleted kafka-kubernetes
	PreviewRestore(ctx context.Context, in *model.KafkaKubernetes, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// restore a deleted kafka-kubernetes
	Restore(ctx context.Context, in *model.KafkaKubernetes, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// restart a kafka-kubernetes running in a environment.
	// kafka-kubernetes is restarted by deleting running "broker" pods which will be automatically recreated by kubernetes
	// note: zookeeper pods are not deleted.
	Restart(ctx context.Context, in *model.KafkaKubernetesId, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// pause a kafka-kubernetes running in a environment.
	// kafka-kubernetes is paused by scaling down number of replicas of
	// the kubernetes deployment/stateful sets to zero in the environment.
	Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// unpause a previously paused kafka-kubernetes running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the kafka-kubernetes.
	Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// preview refresh a kafka-kubernetes that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// refresh a kafka-kubernetes that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
}

type kafkaKubernetesCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKafkaKubernetesCommandControllerClient(cc grpc.ClientConnInterface) KafkaKubernetesCommandControllerClient {
	return &kafkaKubernetesCommandControllerClient{cc}
}

func (c *kafkaKubernetesCommandControllerClient) PreviewCreate(ctx context.Context, in *model.KafkaKubernetes, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesCommandControllerClient) Create(ctx context.Context, in *model.KafkaKubernetes, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.KafkaKubernetes, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesCommandControllerClient) Update(ctx context.Context, in *model.KafkaKubernetes, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesCommandControllerClient) PreviewRestore(ctx context.Context, in *model.KafkaKubernetes, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesCommandControllerClient) Restore(ctx context.Context, in *model.KafkaKubernetes, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesCommandControllerClient) Restart(ctx context.Context, in *model.KafkaKubernetesId, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesCommandControllerClient) Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesCommandControllerClient) Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafkaKubernetesCommandControllerServer is the server API for KafkaKubernetesCommandController service.
// All implementations should embed UnimplementedKafkaKubernetesCommandControllerServer
// for forward compatibility
type KafkaKubernetesCommandControllerServer interface {
	// preview create kafka-kubernetes
	PreviewCreate(context.Context, *model.KafkaKubernetes) (*model.KafkaKubernetes, error)
	// create kafka-kubernetes
	Create(context.Context, *model.KafkaKubernetes) (*model.KafkaKubernetes, error)
	// preview update an existing kafka-kubernetes
	PreviewUpdate(context.Context, *model.KafkaKubernetes) (*model.KafkaKubernetes, error)
	// update an existing kafka-kubernetes
	Update(context.Context, *model.KafkaKubernetes) (*model.KafkaKubernetes, error)
	// preview deleting an existing kafka-kubernetes
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.KafkaKubernetes, error)
	// delete an existing kafka-kubernetes
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.KafkaKubernetes, error)
	// preview restoring a deleted kafka-kubernetes
	PreviewRestore(context.Context, *model.KafkaKubernetes) (*model.KafkaKubernetes, error)
	// restore a deleted kafka-kubernetes
	Restore(context.Context, *model.KafkaKubernetes) (*model.KafkaKubernetes, error)
	// restart a kafka-kubernetes running in a environment.
	// kafka-kubernetes is restarted by deleting running "broker" pods which will be automatically recreated by kubernetes
	// note: zookeeper pods are not deleted.
	Restart(context.Context, *model.KafkaKubernetesId) (*model.KafkaKubernetes, error)
	// pause a kafka-kubernetes running in a environment.
	// kafka-kubernetes is paused by scaling down number of replicas of
	// the kubernetes deployment/stateful sets to zero in the environment.
	Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.KafkaKubernetes, error)
	// unpause a previously paused kafka-kubernetes running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the kafka-kubernetes.
	Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.KafkaKubernetes, error)
	// preview refresh a kafka-kubernetes that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.KafkaKubernetes, error)
	// refresh a kafka-kubernetes that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.KafkaKubernetes, error)
}

// UnimplementedKafkaKubernetesCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKafkaKubernetesCommandControllerServer struct {
}

func (UnimplementedKafkaKubernetesCommandControllerServer) PreviewCreate(context.Context, *model.KafkaKubernetes) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedKafkaKubernetesCommandControllerServer) Create(context.Context, *model.KafkaKubernetes) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedKafkaKubernetesCommandControllerServer) PreviewUpdate(context.Context, *model.KafkaKubernetes) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedKafkaKubernetesCommandControllerServer) Update(context.Context, *model.KafkaKubernetes) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedKafkaKubernetesCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedKafkaKubernetesCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKafkaKubernetesCommandControllerServer) PreviewRestore(context.Context, *model.KafkaKubernetes) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedKafkaKubernetesCommandControllerServer) Restore(context.Context, *model.KafkaKubernetes) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedKafkaKubernetesCommandControllerServer) Restart(context.Context, *model.KafkaKubernetesId) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedKafkaKubernetesCommandControllerServer) Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedKafkaKubernetesCommandControllerServer) Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}
func (UnimplementedKafkaKubernetesCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedKafkaKubernetesCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeKafkaKubernetesCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KafkaKubernetesCommandControllerServer will
// result in compilation errors.
type UnsafeKafkaKubernetesCommandControllerServer interface {
	mustEmbedUnimplementedKafkaKubernetesCommandControllerServer()
}

func RegisterKafkaKubernetesCommandControllerServer(s grpc.ServiceRegistrar, srv KafkaKubernetesCommandControllerServer) {
	s.RegisterService(&KafkaKubernetesCommandController_ServiceDesc, srv)
}

func _KafkaKubernetesCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).PreviewCreate(ctx, req.(*model.KafkaKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).Create(ctx, req.(*model.KafkaKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).PreviewUpdate(ctx, req.(*model.KafkaKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).Update(ctx, req.(*model.KafkaKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).PreviewRestore(ctx, req.(*model.KafkaKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).Restore(ctx, req.(*model.KafkaKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).Restart(ctx, req.(*model.KafkaKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourcePauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).Pause(ctx, req.(*model1.ApiResourcePauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceUnPauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).Unpause(ctx, req.(*model1.ApiResourceUnPauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// KafkaKubernetesCommandController_ServiceDesc is the grpc.ServiceDesc for KafkaKubernetesCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KafkaKubernetesCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesCommandController",
	HandlerType: (*KafkaKubernetesCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _KafkaKubernetesCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _KafkaKubernetesCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _KafkaKubernetesCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _KafkaKubernetesCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _KafkaKubernetesCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _KafkaKubernetesCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _KafkaKubernetesCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _KafkaKubernetesCommandController_Restore_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _KafkaKubernetesCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _KafkaKubernetesCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _KafkaKubernetesCommandController_Unpause_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _KafkaKubernetesCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _KafkaKubernetesCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/kubernetes/kafkakubernetes/service/command.proto",
}

const (
	KafkaTopicCommandController_Add_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaTopicCommandController/add"
	KafkaTopicCommandController_AddMultiple_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaTopicCommandController/addMultiple"
	KafkaTopicCommandController_Update_FullMethodName      = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaTopicCommandController/update"
	KafkaTopicCommandController_Delete_FullMethodName      = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaTopicCommandController/delete"
)

// KafkaTopicCommandControllerClient is the client API for KafkaTopicCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KafkaTopicCommandControllerClient interface {
	// add a single kafka topic to existing list of kafka topics of a kafka-kubernetes
	Add(ctx context.Context, in *model.AddOrUpdateKafkaTopicCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// add multiple kafka topics to existing list of kafka topics of a kafka-kubernetes
	AddMultiple(ctx context.Context, in *model.AddKafkaTopicsCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// update a kafka topic.
	Update(ctx context.Context, in *model.AddOrUpdateKafkaTopicCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// delete a kafka topic.
	Delete(ctx context.Context, in *model.DeleteOrRestoreKafkaTopicCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
}

type kafkaTopicCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKafkaTopicCommandControllerClient(cc grpc.ClientConnInterface) KafkaTopicCommandControllerClient {
	return &kafkaTopicCommandControllerClient{cc}
}

func (c *kafkaTopicCommandControllerClient) Add(ctx context.Context, in *model.AddOrUpdateKafkaTopicCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaTopicCommandController_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaTopicCommandControllerClient) AddMultiple(ctx context.Context, in *model.AddKafkaTopicsCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaTopicCommandController_AddMultiple_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaTopicCommandControllerClient) Update(ctx context.Context, in *model.AddOrUpdateKafkaTopicCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaTopicCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaTopicCommandControllerClient) Delete(ctx context.Context, in *model.DeleteOrRestoreKafkaTopicCommandInput, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaTopicCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafkaTopicCommandControllerServer is the server API for KafkaTopicCommandController service.
// All implementations should embed UnimplementedKafkaTopicCommandControllerServer
// for forward compatibility
type KafkaTopicCommandControllerServer interface {
	// add a single kafka topic to existing list of kafka topics of a kafka-kubernetes
	Add(context.Context, *model.AddOrUpdateKafkaTopicCommandInput) (*model.KafkaKubernetes, error)
	// add multiple kafka topics to existing list of kafka topics of a kafka-kubernetes
	AddMultiple(context.Context, *model.AddKafkaTopicsCommandInput) (*model.KafkaKubernetes, error)
	// update a kafka topic.
	Update(context.Context, *model.AddOrUpdateKafkaTopicCommandInput) (*model.KafkaKubernetes, error)
	// delete a kafka topic.
	Delete(context.Context, *model.DeleteOrRestoreKafkaTopicCommandInput) (*model.KafkaKubernetes, error)
}

// UnimplementedKafkaTopicCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKafkaTopicCommandControllerServer struct {
}

func (UnimplementedKafkaTopicCommandControllerServer) Add(context.Context, *model.AddOrUpdateKafkaTopicCommandInput) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedKafkaTopicCommandControllerServer) AddMultiple(context.Context, *model.AddKafkaTopicsCommandInput) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMultiple not implemented")
}
func (UnimplementedKafkaTopicCommandControllerServer) Update(context.Context, *model.AddOrUpdateKafkaTopicCommandInput) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedKafkaTopicCommandControllerServer) Delete(context.Context, *model.DeleteOrRestoreKafkaTopicCommandInput) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeKafkaTopicCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KafkaTopicCommandControllerServer will
// result in compilation errors.
type UnsafeKafkaTopicCommandControllerServer interface {
	mustEmbedUnimplementedKafkaTopicCommandControllerServer()
}

func RegisterKafkaTopicCommandControllerServer(s grpc.ServiceRegistrar, srv KafkaTopicCommandControllerServer) {
	s.RegisterService(&KafkaTopicCommandController_ServiceDesc, srv)
}

func _KafkaTopicCommandController_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AddOrUpdateKafkaTopicCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaTopicCommandControllerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaTopicCommandController_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaTopicCommandControllerServer).Add(ctx, req.(*model.AddOrUpdateKafkaTopicCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaTopicCommandController_AddMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AddKafkaTopicsCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaTopicCommandControllerServer).AddMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaTopicCommandController_AddMultiple_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaTopicCommandControllerServer).AddMultiple(ctx, req.(*model.AddKafkaTopicsCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaTopicCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AddOrUpdateKafkaTopicCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaTopicCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaTopicCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaTopicCommandControllerServer).Update(ctx, req.(*model.AddOrUpdateKafkaTopicCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaTopicCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.DeleteOrRestoreKafkaTopicCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaTopicCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaTopicCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaTopicCommandControllerServer).Delete(ctx, req.(*model.DeleteOrRestoreKafkaTopicCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// KafkaTopicCommandController_ServiceDesc is the grpc.ServiceDesc for KafkaTopicCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KafkaTopicCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaTopicCommandController",
	HandlerType: (*KafkaTopicCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _KafkaTopicCommandController_Add_Handler,
		},
		{
			MethodName: "addMultiple",
			Handler:    _KafkaTopicCommandController_AddMultiple_Handler,
		},
		{
			MethodName: "update",
			Handler:    _KafkaTopicCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _KafkaTopicCommandController_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/kubernetes/kafkakubernetes/service/command.proto",
}
