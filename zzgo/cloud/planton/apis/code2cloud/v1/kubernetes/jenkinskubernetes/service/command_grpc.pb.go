// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/kubernetes/jenkinskubernetes/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/jenkinskubernetes/model"
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
	JenkinsKubernetesCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/previewCreate"
	JenkinsKubernetesCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/create"
	JenkinsKubernetesCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/previewUpdate"
	JenkinsKubernetesCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/update"
	JenkinsKubernetesCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/previewDelete"
	JenkinsKubernetesCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/delete"
	JenkinsKubernetesCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/previewRestore"
	JenkinsKubernetesCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/restore"
	JenkinsKubernetesCommandController_Restart_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/restart"
	JenkinsKubernetesCommandController_Pause_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/pause"
	JenkinsKubernetesCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/unpause"
	JenkinsKubernetesCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/previewRefresh"
	JenkinsKubernetesCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController/refresh"
)

// JenkinsKubernetesCommandControllerClient is the client API for JenkinsKubernetesCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JenkinsKubernetesCommandControllerClient interface {
	// preview creating jenkins-kubernetes
	PreviewCreate(ctx context.Context, in *model.JenkinsKubernetes, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
	// create jenkins-kubernetes
	Create(ctx context.Context, in *model.JenkinsKubernetes, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
	// preview updating an existing jenkins-kubernetes
	PreviewUpdate(ctx context.Context, in *model.JenkinsKubernetes, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
	// update an existing jenkins-kubernetes
	Update(ctx context.Context, in *model.JenkinsKubernetes, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
	// preview deleting an existing jenkins-kubernetes
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
	// delete an existing jenkins-kubernetes
	Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
	// preview restoring a previously deleted jenkins-kubernetes
	PreviewRestore(ctx context.Context, in *model.JenkinsKubernetes, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
	// restore a previously deleted jenkins-kubernetes
	Restore(ctx context.Context, in *model.JenkinsKubernetes, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
	// restart a jenkins-kubernetes running in a environment.
	// jenkins-kubernetes is restarted by deleting running "jenkins" pods which will be automatically recreated by kubernetes
	Restart(ctx context.Context, in *model.JenkinsKubernetesId, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
	// pause a jenkins-kubernetes running in a environment.
	// jenkins-kubernetes is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(ctx context.Context, in *model1.ApiResourcePauseInput, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
	// unpause a previously paused jenkins-kubernetes running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the jenkins-kubernetes.
	Unpause(ctx context.Context, in *model1.ApiResourceUnPauseInput, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
	// preview refresh a jenkins-kubernetes that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
	// refresh a jenkins-kubernetes that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error)
}

type jenkinsKubernetesCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewJenkinsKubernetesCommandControllerClient(cc grpc.ClientConnInterface) JenkinsKubernetesCommandControllerClient {
	return &jenkinsKubernetesCommandControllerClient{cc}
}

func (c *jenkinsKubernetesCommandControllerClient) PreviewCreate(ctx context.Context, in *model.JenkinsKubernetes, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsKubernetesCommandControllerClient) Create(ctx context.Context, in *model.JenkinsKubernetes, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsKubernetesCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.JenkinsKubernetes, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsKubernetesCommandControllerClient) Update(ctx context.Context, in *model.JenkinsKubernetes, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsKubernetesCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsKubernetesCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsKubernetesCommandControllerClient) PreviewRestore(ctx context.Context, in *model.JenkinsKubernetes, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsKubernetesCommandControllerClient) Restore(ctx context.Context, in *model.JenkinsKubernetes, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsKubernetesCommandControllerClient) Restart(ctx context.Context, in *model.JenkinsKubernetesId, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsKubernetesCommandControllerClient) Pause(ctx context.Context, in *model1.ApiResourcePauseInput, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsKubernetesCommandControllerClient) Unpause(ctx context.Context, in *model1.ApiResourceUnPauseInput, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsKubernetesCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsKubernetesCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.JenkinsKubernetes, error) {
	out := new(model.JenkinsKubernetes)
	err := c.cc.Invoke(ctx, JenkinsKubernetesCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JenkinsKubernetesCommandControllerServer is the server API for JenkinsKubernetesCommandController service.
// All implementations should embed UnimplementedJenkinsKubernetesCommandControllerServer
// for forward compatibility
type JenkinsKubernetesCommandControllerServer interface {
	// preview creating jenkins-kubernetes
	PreviewCreate(context.Context, *model.JenkinsKubernetes) (*model.JenkinsKubernetes, error)
	// create jenkins-kubernetes
	Create(context.Context, *model.JenkinsKubernetes) (*model.JenkinsKubernetes, error)
	// preview updating an existing jenkins-kubernetes
	PreviewUpdate(context.Context, *model.JenkinsKubernetes) (*model.JenkinsKubernetes, error)
	// update an existing jenkins-kubernetes
	Update(context.Context, *model.JenkinsKubernetes) (*model.JenkinsKubernetes, error)
	// preview deleting an existing jenkins-kubernetes
	PreviewDelete(context.Context, *model1.ApiResourceDeleteInput) (*model.JenkinsKubernetes, error)
	// delete an existing jenkins-kubernetes
	Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.JenkinsKubernetes, error)
	// preview restoring a previously deleted jenkins-kubernetes
	PreviewRestore(context.Context, *model.JenkinsKubernetes) (*model.JenkinsKubernetes, error)
	// restore a previously deleted jenkins-kubernetes
	Restore(context.Context, *model.JenkinsKubernetes) (*model.JenkinsKubernetes, error)
	// restart a jenkins-kubernetes running in a environment.
	// jenkins-kubernetes is restarted by deleting running "jenkins" pods which will be automatically recreated by kubernetes
	Restart(context.Context, *model.JenkinsKubernetesId) (*model.JenkinsKubernetes, error)
	// pause a jenkins-kubernetes running in a environment.
	// jenkins-kubernetes is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(context.Context, *model1.ApiResourcePauseInput) (*model.JenkinsKubernetes, error)
	// unpause a previously paused jenkins-kubernetes running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the jenkins-kubernetes.
	Unpause(context.Context, *model1.ApiResourceUnPauseInput) (*model.JenkinsKubernetes, error)
	// preview refresh a jenkins-kubernetes that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshInput) (*model.JenkinsKubernetes, error)
	// refresh a jenkins-kubernetes that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshInput) (*model.JenkinsKubernetes, error)
}

// UnimplementedJenkinsKubernetesCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedJenkinsKubernetesCommandControllerServer struct {
}

func (UnimplementedJenkinsKubernetesCommandControllerServer) PreviewCreate(context.Context, *model.JenkinsKubernetes) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedJenkinsKubernetesCommandControllerServer) Create(context.Context, *model.JenkinsKubernetes) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedJenkinsKubernetesCommandControllerServer) PreviewUpdate(context.Context, *model.JenkinsKubernetes) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedJenkinsKubernetesCommandControllerServer) Update(context.Context, *model.JenkinsKubernetes) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedJenkinsKubernetesCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteInput) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedJenkinsKubernetesCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedJenkinsKubernetesCommandControllerServer) PreviewRestore(context.Context, *model.JenkinsKubernetes) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedJenkinsKubernetesCommandControllerServer) Restore(context.Context, *model.JenkinsKubernetes) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedJenkinsKubernetesCommandControllerServer) Restart(context.Context, *model.JenkinsKubernetesId) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedJenkinsKubernetesCommandControllerServer) Pause(context.Context, *model1.ApiResourcePauseInput) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedJenkinsKubernetesCommandControllerServer) Unpause(context.Context, *model1.ApiResourceUnPauseInput) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}
func (UnimplementedJenkinsKubernetesCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshInput) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedJenkinsKubernetesCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshInput) (*model.JenkinsKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeJenkinsKubernetesCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JenkinsKubernetesCommandControllerServer will
// result in compilation errors.
type UnsafeJenkinsKubernetesCommandControllerServer interface {
	mustEmbedUnimplementedJenkinsKubernetesCommandControllerServer()
}

func RegisterJenkinsKubernetesCommandControllerServer(s grpc.ServiceRegistrar, srv JenkinsKubernetesCommandControllerServer) {
	s.RegisterService(&JenkinsKubernetesCommandController_ServiceDesc, srv)
}

func _JenkinsKubernetesCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.JenkinsKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).PreviewCreate(ctx, req.(*model.JenkinsKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsKubernetesCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.JenkinsKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).Create(ctx, req.(*model.JenkinsKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsKubernetesCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.JenkinsKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).PreviewUpdate(ctx, req.(*model.JenkinsKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsKubernetesCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.JenkinsKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).Update(ctx, req.(*model.JenkinsKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsKubernetesCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsKubernetesCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsKubernetesCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.JenkinsKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).PreviewRestore(ctx, req.(*model.JenkinsKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsKubernetesCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.JenkinsKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).Restore(ctx, req.(*model.JenkinsKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsKubernetesCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.JenkinsKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).Restart(ctx, req.(*model.JenkinsKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsKubernetesCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourcePauseInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).Pause(ctx, req.(*model1.ApiResourcePauseInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsKubernetesCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceUnPauseInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).Unpause(ctx, req.(*model1.ApiResourceUnPauseInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsKubernetesCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsKubernetesCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsKubernetesCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsKubernetesCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsKubernetesCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshInput))
	}
	return interceptor(ctx, in, info, handler)
}

// JenkinsKubernetesCommandController_ServiceDesc is the grpc.ServiceDesc for JenkinsKubernetesCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JenkinsKubernetesCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.kubernetes.jenkinskubernetes.service.JenkinsKubernetesCommandController",
	HandlerType: (*JenkinsKubernetesCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _JenkinsKubernetesCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _JenkinsKubernetesCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _JenkinsKubernetesCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _JenkinsKubernetesCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _JenkinsKubernetesCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _JenkinsKubernetesCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _JenkinsKubernetesCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _JenkinsKubernetesCommandController_Restore_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _JenkinsKubernetesCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _JenkinsKubernetesCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _JenkinsKubernetesCommandController_Unpause_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _JenkinsKubernetesCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _JenkinsKubernetesCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/kubernetes/jenkinskubernetes/service/command.proto",
}