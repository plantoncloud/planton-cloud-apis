// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/gitlabkubernetes/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gitlabkubernetes/model"
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
	GitlabKubernetesCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/previewCreate"
	GitlabKubernetesCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/create"
	GitlabKubernetesCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/previewUpdate"
	GitlabKubernetesCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/update"
	GitlabKubernetesCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/previewDelete"
	GitlabKubernetesCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/delete"
	GitlabKubernetesCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/previewRestore"
	GitlabKubernetesCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/restore"
	GitlabKubernetesCommandController_Restart_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/restart"
	GitlabKubernetesCommandController_Pause_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/pause"
	GitlabKubernetesCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/unpause"
	GitlabKubernetesCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/previewRefresh"
	GitlabKubernetesCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController/refresh"
)

// GitlabKubernetesCommandControllerClient is the client API for GitlabKubernetesCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitlabKubernetesCommandControllerClient interface {
	// preview creating gitlab-server
	PreviewCreate(ctx context.Context, in *model.GitlabKubernetes, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
	// create gitlab-server
	Create(ctx context.Context, in *model.GitlabKubernetes, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
	// preview updating an existing gitlab-server
	PreviewUpdate(ctx context.Context, in *model.GitlabKubernetes, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
	// update an existing gitlab-server
	Update(ctx context.Context, in *model.GitlabKubernetes, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
	// preview deleting an existing gitlab-server
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
	// delete an existing gitlab-server
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
	// preview restoring a previously deleted gitlab-server
	PreviewRestore(ctx context.Context, in *model.GitlabKubernetes, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
	// restore a previously deleted gitlab-server
	Restore(ctx context.Context, in *model.GitlabKubernetes, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
	// restart a gitlab-server running in a environment.
	// gitlab-server is restarted by deleting running "gitlab" pods which will be automatically recreated by kubernetes
	Restart(ctx context.Context, in *model.GitlabKubernetesId, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
	// pause a gitlab-server running in a environment.
	// gitlab-server is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
	// unpause a previously paused gitlab-server running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the gitlab-server.
	Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
	// preview refresh a gitlab-server that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
	// refresh a gitlab-server that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.GitlabKubernetes, error)
}

type gitlabKubernetesCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGitlabKubernetesCommandControllerClient(cc grpc.ClientConnInterface) GitlabKubernetesCommandControllerClient {
	return &gitlabKubernetesCommandControllerClient{cc}
}

func (c *gitlabKubernetesCommandControllerClient) PreviewCreate(ctx context.Context, in *model.GitlabKubernetes, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabKubernetesCommandControllerClient) Create(ctx context.Context, in *model.GitlabKubernetes, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabKubernetesCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.GitlabKubernetes, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabKubernetesCommandControllerClient) Update(ctx context.Context, in *model.GitlabKubernetes, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabKubernetesCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabKubernetesCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabKubernetesCommandControllerClient) PreviewRestore(ctx context.Context, in *model.GitlabKubernetes, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabKubernetesCommandControllerClient) Restore(ctx context.Context, in *model.GitlabKubernetes, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabKubernetesCommandControllerClient) Restart(ctx context.Context, in *model.GitlabKubernetesId, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabKubernetesCommandControllerClient) Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabKubernetesCommandControllerClient) Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabKubernetesCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabKubernetesCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.GitlabKubernetes, error) {
	out := new(model.GitlabKubernetes)
	err := c.cc.Invoke(ctx, GitlabKubernetesCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitlabKubernetesCommandControllerServer is the server API for GitlabKubernetesCommandController service.
// All implementations should embed UnimplementedGitlabKubernetesCommandControllerServer
// for forward compatibility
type GitlabKubernetesCommandControllerServer interface {
	// preview creating gitlab-server
	PreviewCreate(context.Context, *model.GitlabKubernetes) (*model.GitlabKubernetes, error)
	// create gitlab-server
	Create(context.Context, *model.GitlabKubernetes) (*model.GitlabKubernetes, error)
	// preview updating an existing gitlab-server
	PreviewUpdate(context.Context, *model.GitlabKubernetes) (*model.GitlabKubernetes, error)
	// update an existing gitlab-server
	Update(context.Context, *model.GitlabKubernetes) (*model.GitlabKubernetes, error)
	// preview deleting an existing gitlab-server
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.GitlabKubernetes, error)
	// delete an existing gitlab-server
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.GitlabKubernetes, error)
	// preview restoring a previously deleted gitlab-server
	PreviewRestore(context.Context, *model.GitlabKubernetes) (*model.GitlabKubernetes, error)
	// restore a previously deleted gitlab-server
	Restore(context.Context, *model.GitlabKubernetes) (*model.GitlabKubernetes, error)
	// restart a gitlab-server running in a environment.
	// gitlab-server is restarted by deleting running "gitlab" pods which will be automatically recreated by kubernetes
	Restart(context.Context, *model.GitlabKubernetesId) (*model.GitlabKubernetes, error)
	// pause a gitlab-server running in a environment.
	// gitlab-server is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.GitlabKubernetes, error)
	// unpause a previously paused gitlab-server running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the gitlab-server.
	Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.GitlabKubernetes, error)
	// preview refresh a gitlab-server that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.GitlabKubernetes, error)
	// refresh a gitlab-server that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.GitlabKubernetes, error)
}

// UnimplementedGitlabKubernetesCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGitlabKubernetesCommandControllerServer struct {
}

func (UnimplementedGitlabKubernetesCommandControllerServer) PreviewCreate(context.Context, *model.GitlabKubernetes) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedGitlabKubernetesCommandControllerServer) Create(context.Context, *model.GitlabKubernetes) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGitlabKubernetesCommandControllerServer) PreviewUpdate(context.Context, *model.GitlabKubernetes) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedGitlabKubernetesCommandControllerServer) Update(context.Context, *model.GitlabKubernetes) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGitlabKubernetesCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedGitlabKubernetesCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGitlabKubernetesCommandControllerServer) PreviewRestore(context.Context, *model.GitlabKubernetes) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedGitlabKubernetesCommandControllerServer) Restore(context.Context, *model.GitlabKubernetes) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedGitlabKubernetesCommandControllerServer) Restart(context.Context, *model.GitlabKubernetesId) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedGitlabKubernetesCommandControllerServer) Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedGitlabKubernetesCommandControllerServer) Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}
func (UnimplementedGitlabKubernetesCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedGitlabKubernetesCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.GitlabKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeGitlabKubernetesCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitlabKubernetesCommandControllerServer will
// result in compilation errors.
type UnsafeGitlabKubernetesCommandControllerServer interface {
	mustEmbedUnimplementedGitlabKubernetesCommandControllerServer()
}

func RegisterGitlabKubernetesCommandControllerServer(s grpc.ServiceRegistrar, srv GitlabKubernetesCommandControllerServer) {
	s.RegisterService(&GitlabKubernetesCommandController_ServiceDesc, srv)
}

func _GitlabKubernetesCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).PreviewCreate(ctx, req.(*model.GitlabKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabKubernetesCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).Create(ctx, req.(*model.GitlabKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabKubernetesCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).PreviewUpdate(ctx, req.(*model.GitlabKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabKubernetesCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).Update(ctx, req.(*model.GitlabKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabKubernetesCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabKubernetesCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabKubernetesCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).PreviewRestore(ctx, req.(*model.GitlabKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabKubernetesCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabKubernetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).Restore(ctx, req.(*model.GitlabKubernetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabKubernetesCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).Restart(ctx, req.(*model.GitlabKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabKubernetesCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourcePauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).Pause(ctx, req.(*model1.ApiResourcePauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabKubernetesCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceUnPauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).Unpause(ctx, req.(*model1.ApiResourceUnPauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabKubernetesCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabKubernetesCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabKubernetesCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabKubernetesCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabKubernetesCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GitlabKubernetesCommandController_ServiceDesc is the grpc.ServiceDesc for GitlabKubernetesCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitlabKubernetesCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.gitlabkubernetes.service.GitlabKubernetesCommandController",
	HandlerType: (*GitlabKubernetesCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _GitlabKubernetesCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _GitlabKubernetesCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _GitlabKubernetesCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _GitlabKubernetesCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _GitlabKubernetesCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _GitlabKubernetesCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _GitlabKubernetesCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _GitlabKubernetesCommandController_Restore_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _GitlabKubernetesCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _GitlabKubernetesCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _GitlabKubernetesCommandController_Unpause_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _GitlabKubernetesCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _GitlabKubernetesCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/gitlabkubernetes/service/command.proto",
}
