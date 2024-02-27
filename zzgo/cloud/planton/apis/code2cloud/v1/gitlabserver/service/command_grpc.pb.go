// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/gitlabserver/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gitlabserver/model"
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
	GitlabServerCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/previewCreate"
	GitlabServerCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/create"
	GitlabServerCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/previewUpdate"
	GitlabServerCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/update"
	GitlabServerCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/previewDelete"
	GitlabServerCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/delete"
	GitlabServerCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/previewRestore"
	GitlabServerCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/restore"
	GitlabServerCommandController_CreateStackJob_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/createStackJob"
	GitlabServerCommandController_Restart_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/restart"
	GitlabServerCommandController_Pause_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/pause"
	GitlabServerCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/unpause"
	GitlabServerCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/previewRefresh"
	GitlabServerCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController/refresh"
)

// GitlabServerCommandControllerClient is the client API for GitlabServerCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitlabServerCommandControllerClient interface {
	// preview creating gitlab-server
	PreviewCreate(ctx context.Context, in *model.GitlabServer, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// create gitlab-server
	Create(ctx context.Context, in *model.GitlabServer, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// preview updating an existing gitlab-server
	PreviewUpdate(ctx context.Context, in *model.GitlabServer, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// update an existing gitlab-server
	Update(ctx context.Context, in *model.GitlabServer, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// preview deleting an existing gitlab-server
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// delete an existing gitlab-server
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// preview restoring a previously deleted gitlab-server
	PreviewRestore(ctx context.Context, in *model.GitlabServer, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// restore a previously deleted gitlab-server
	Restore(ctx context.Context, in *model.GitlabServer, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// create-stack-job for gitlab-server
	CreateStackJob(ctx context.Context, in *model2.CreateStackJobCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// restart a gitlab-server running in a environment.
	// gitlab-server is restarted by deleting running "gitlab" pods which will be automatically recreated by kubernetes
	Restart(ctx context.Context, in *model.GitlabServerId, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// pause a gitlab-server running in a environment.
	// gitlab-server is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// unpause a previously paused gitlab-server running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the gitlab-server.
	Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// preview refresh a gitlab-server that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// refresh a gitlab-server that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error)
}

type gitlabServerCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGitlabServerCommandControllerClient(cc grpc.ClientConnInterface) GitlabServerCommandControllerClient {
	return &gitlabServerCommandControllerClient{cc}
}

func (c *gitlabServerCommandControllerClient) PreviewCreate(ctx context.Context, in *model.GitlabServer, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) Create(ctx context.Context, in *model.GitlabServer, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.GitlabServer, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) Update(ctx context.Context, in *model.GitlabServer, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) PreviewRestore(ctx context.Context, in *model.GitlabServer, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) Restore(ctx context.Context, in *model.GitlabServer, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) CreateStackJob(ctx context.Context, in *model2.CreateStackJobCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_CreateStackJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) Restart(ctx context.Context, in *model.GitlabServerId, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitlabServerCommandControllerServer is the server API for GitlabServerCommandController service.
// All implementations should embed UnimplementedGitlabServerCommandControllerServer
// for forward compatibility
type GitlabServerCommandControllerServer interface {
	// preview creating gitlab-server
	PreviewCreate(context.Context, *model.GitlabServer) (*model.GitlabServer, error)
	// create gitlab-server
	Create(context.Context, *model.GitlabServer) (*model.GitlabServer, error)
	// preview updating an existing gitlab-server
	PreviewUpdate(context.Context, *model.GitlabServer) (*model.GitlabServer, error)
	// update an existing gitlab-server
	Update(context.Context, *model.GitlabServer) (*model.GitlabServer, error)
	// preview deleting an existing gitlab-server
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.GitlabServer, error)
	// delete an existing gitlab-server
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.GitlabServer, error)
	// preview restoring a previously deleted gitlab-server
	PreviewRestore(context.Context, *model.GitlabServer) (*model.GitlabServer, error)
	// restore a previously deleted gitlab-server
	Restore(context.Context, *model.GitlabServer) (*model.GitlabServer, error)
	// create-stack-job for gitlab-server
	CreateStackJob(context.Context, *model2.CreateStackJobCommandInput) (*model.GitlabServer, error)
	// restart a gitlab-server running in a environment.
	// gitlab-server is restarted by deleting running "gitlab" pods which will be automatically recreated by kubernetes
	Restart(context.Context, *model.GitlabServerId) (*model.GitlabServer, error)
	// pause a gitlab-server running in a environment.
	// gitlab-server is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.GitlabServer, error)
	// unpause a previously paused gitlab-server running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the gitlab-server.
	Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.GitlabServer, error)
	// preview refresh a gitlab-server that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.GitlabServer, error)
	// refresh a gitlab-server that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.GitlabServer, error)
}

// UnimplementedGitlabServerCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGitlabServerCommandControllerServer struct {
}

func (UnimplementedGitlabServerCommandControllerServer) PreviewCreate(context.Context, *model.GitlabServer) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) Create(context.Context, *model.GitlabServer) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) PreviewUpdate(context.Context, *model.GitlabServer) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) Update(context.Context, *model.GitlabServer) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) PreviewRestore(context.Context, *model.GitlabServer) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) Restore(context.Context, *model.GitlabServer) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) CreateStackJob(context.Context, *model2.CreateStackJobCommandInput) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStackJob not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) Restart(context.Context, *model.GitlabServerId) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedGitlabServerCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeGitlabServerCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitlabServerCommandControllerServer will
// result in compilation errors.
type UnsafeGitlabServerCommandControllerServer interface {
	mustEmbedUnimplementedGitlabServerCommandControllerServer()
}

func RegisterGitlabServerCommandControllerServer(s grpc.ServiceRegistrar, srv GitlabServerCommandControllerServer) {
	s.RegisterService(&GitlabServerCommandController_ServiceDesc, srv)
}

func _GitlabServerCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).PreviewCreate(ctx, req.(*model.GitlabServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).Create(ctx, req.(*model.GitlabServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).PreviewUpdate(ctx, req.(*model.GitlabServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).Update(ctx, req.(*model.GitlabServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).PreviewRestore(ctx, req.(*model.GitlabServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).Restore(ctx, req.(*model.GitlabServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_CreateStackJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.CreateStackJobCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).CreateStackJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_CreateStackJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).CreateStackJob(ctx, req.(*model2.CreateStackJobCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).Restart(ctx, req.(*model.GitlabServerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourcePauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).Pause(ctx, req.(*model1.ApiResourcePauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceUnPauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).Unpause(ctx, req.(*model1.ApiResourceUnPauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GitlabServerCommandController_ServiceDesc is the grpc.ServiceDesc for GitlabServerCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitlabServerCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerCommandController",
	HandlerType: (*GitlabServerCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _GitlabServerCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _GitlabServerCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _GitlabServerCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _GitlabServerCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _GitlabServerCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _GitlabServerCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _GitlabServerCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _GitlabServerCommandController_Restore_Handler,
		},
		{
			MethodName: "createStackJob",
			Handler:    _GitlabServerCommandController_CreateStackJob_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _GitlabServerCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _GitlabServerCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _GitlabServerCommandController_Unpause_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _GitlabServerCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _GitlabServerCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/gitlabserver/service/command.proto",
}
