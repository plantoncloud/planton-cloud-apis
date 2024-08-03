// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/connect/v1/gitlabcredential/service/command.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/gitlabcredential/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GitlabCredentialCommandController_Create_FullMethodName  = "/cloud.planton.apis.connect.v1.gitlabcredential.service.GitlabCredentialCommandController/create"
	GitlabCredentialCommandController_Update_FullMethodName  = "/cloud.planton.apis.connect.v1.gitlabcredential.service.GitlabCredentialCommandController/update"
	GitlabCredentialCommandController_Delete_FullMethodName  = "/cloud.planton.apis.connect.v1.gitlabcredential.service.GitlabCredentialCommandController/delete"
	GitlabCredentialCommandController_Restore_FullMethodName = "/cloud.planton.apis.connect.v1.gitlabcredential.service.GitlabCredentialCommandController/restore"
)

// GitlabCredentialCommandControllerClient is the client API for GitlabCredentialCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitlabCredentialCommandControllerClient interface {
	// create a gitlab-credential resource
	Create(ctx context.Context, in *model.GitlabCredential, opts ...grpc.CallOption) (*model.GitlabCredential, error)
	// update an existing gitlab-credential
	Update(ctx context.Context, in *model.GitlabCredential, opts ...grpc.CallOption) (*model.GitlabCredential, error)
	// delete a gitlab-credential that was previously created
	Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GitlabCredential, error)
	// restore a deleted gitlab-credential.
	Restore(ctx context.Context, in *model.GitlabCredential, opts ...grpc.CallOption) (*model.GitlabCredential, error)
}

type gitlabCredentialCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGitlabCredentialCommandControllerClient(cc grpc.ClientConnInterface) GitlabCredentialCommandControllerClient {
	return &gitlabCredentialCommandControllerClient{cc}
}

func (c *gitlabCredentialCommandControllerClient) Create(ctx context.Context, in *model.GitlabCredential, opts ...grpc.CallOption) (*model.GitlabCredential, error) {
	out := new(model.GitlabCredential)
	err := c.cc.Invoke(ctx, GitlabCredentialCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabCredentialCommandControllerClient) Update(ctx context.Context, in *model.GitlabCredential, opts ...grpc.CallOption) (*model.GitlabCredential, error) {
	out := new(model.GitlabCredential)
	err := c.cc.Invoke(ctx, GitlabCredentialCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabCredentialCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GitlabCredential, error) {
	out := new(model.GitlabCredential)
	err := c.cc.Invoke(ctx, GitlabCredentialCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabCredentialCommandControllerClient) Restore(ctx context.Context, in *model.GitlabCredential, opts ...grpc.CallOption) (*model.GitlabCredential, error) {
	out := new(model.GitlabCredential)
	err := c.cc.Invoke(ctx, GitlabCredentialCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitlabCredentialCommandControllerServer is the server API for GitlabCredentialCommandController service.
// All implementations should embed UnimplementedGitlabCredentialCommandControllerServer
// for forward compatibility
type GitlabCredentialCommandControllerServer interface {
	// create a gitlab-credential resource
	Create(context.Context, *model.GitlabCredential) (*model.GitlabCredential, error)
	// update an existing gitlab-credential
	Update(context.Context, *model.GitlabCredential) (*model.GitlabCredential, error)
	// delete a gitlab-credential that was previously created
	Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.GitlabCredential, error)
	// restore a deleted gitlab-credential.
	Restore(context.Context, *model.GitlabCredential) (*model.GitlabCredential, error)
}

// UnimplementedGitlabCredentialCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGitlabCredentialCommandControllerServer struct {
}

func (UnimplementedGitlabCredentialCommandControllerServer) Create(context.Context, *model.GitlabCredential) (*model.GitlabCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGitlabCredentialCommandControllerServer) Update(context.Context, *model.GitlabCredential) (*model.GitlabCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGitlabCredentialCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.GitlabCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGitlabCredentialCommandControllerServer) Restore(context.Context, *model.GitlabCredential) (*model.GitlabCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeGitlabCredentialCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitlabCredentialCommandControllerServer will
// result in compilation errors.
type UnsafeGitlabCredentialCommandControllerServer interface {
	mustEmbedUnimplementedGitlabCredentialCommandControllerServer()
}

func RegisterGitlabCredentialCommandControllerServer(s grpc.ServiceRegistrar, srv GitlabCredentialCommandControllerServer) {
	s.RegisterService(&GitlabCredentialCommandController_ServiceDesc, srv)
}

func _GitlabCredentialCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabCredentialCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabCredentialCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabCredentialCommandControllerServer).Create(ctx, req.(*model.GitlabCredential))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabCredentialCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabCredentialCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabCredentialCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabCredentialCommandControllerServer).Update(ctx, req.(*model.GitlabCredential))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabCredentialCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabCredentialCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabCredentialCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabCredentialCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabCredentialCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabCredentialCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabCredentialCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabCredentialCommandControllerServer).Restore(ctx, req.(*model.GitlabCredential))
	}
	return interceptor(ctx, in, info, handler)
}

// GitlabCredentialCommandController_ServiceDesc is the grpc.ServiceDesc for GitlabCredentialCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitlabCredentialCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.connect.v1.gitlabcredential.service.GitlabCredentialCommandController",
	HandlerType: (*GitlabCredentialCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _GitlabCredentialCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _GitlabCredentialCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _GitlabCredentialCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _GitlabCredentialCommandController_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/connect/v1/gitlabcredential/service/command.proto",
}
