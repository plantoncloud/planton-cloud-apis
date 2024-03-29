// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/integration/v1/github/service/command.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/codeproject/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/codeserver/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/github/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GithubCommandController_CreRepository_FullMethodName                 = "/cloud.planton.apis.integration.v1.github.service.GithubCommandController/creRepository"
	GithubCommandController_ApplyTemplate_FullMethodName                 = "/cloud.planton.apis.integration.v1.github.service.GithubCommandController/applyTemplate"
	GithubCommandController_AddSecretsToRepo_FullMethodName              = "/cloud.planton.apis.integration.v1.github.service.GithubCommandController/addSecretsToRepo"
	GithubCommandController_AddSecretsToOrg_FullMethodName               = "/cloud.planton.apis.integration.v1.github.service.GithubCommandController/addSecretsToOrg"
	GithubCommandController_SynchronizeMagicPipelineFiles_FullMethodName = "/cloud.planton.apis.integration.v1.github.service.GithubCommandController/synchronizeMagicPipelineFiles"
)

// GithubCommandControllerClient is the client API for GithubCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GithubCommandControllerClient interface {
	// create a new repository on github
	// https://docs.github.com/en/rest/apps/installations?apiVersion=2022-11-28#add-a-repository-to-an-app-installation
	// https://docs.github.com/en/rest/repos/repos#create-a-repository-for-the-authenticated-user
	// https://docs.github.com/en/rest/repos/repos#create-an-organization-repository
	CreRepository(ctx context.Context, in *model.CreRepositoryCommandInput, opts ...grpc.CallOption) (*model1.CodeProject, error)
	// apply a cookiecutter template on a code project created on github
	ApplyTemplate(ctx context.Context, in *model.GithubApplyTemplateCommandInput, opts ...grpc.CallOption) (*model1.CodeProjectProfile, error)
	// add a list of secrets to a github repository
	AddSecretsToRepo(ctx context.Context, in *model.AddSecretsToRepoCommandInput, opts ...grpc.CallOption) (*model1.CodeProject, error)
	// add a list of secrets to a github organization
	AddSecretsToOrg(ctx context.Context, in *model.AddSecretsToOrgCommandInput, opts ...grpc.CallOption) (*model2.CodeServer, error)
	// synchronization is achieved by first removing all yaml files prefixed with "pc:" inside .github/workflows and
	// then add files in the input to .github/workflows directory.
	SynchronizeMagicPipelineFiles(ctx context.Context, in *model.SynchronizeGithubMagicPipelineFilesCommandInput, opts ...grpc.CallOption) (*model1.CodeProject, error)
}

type githubCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGithubCommandControllerClient(cc grpc.ClientConnInterface) GithubCommandControllerClient {
	return &githubCommandControllerClient{cc}
}

func (c *githubCommandControllerClient) CreRepository(ctx context.Context, in *model.CreRepositoryCommandInput, opts ...grpc.CallOption) (*model1.CodeProject, error) {
	out := new(model1.CodeProject)
	err := c.cc.Invoke(ctx, GithubCommandController_CreRepository_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubCommandControllerClient) ApplyTemplate(ctx context.Context, in *model.GithubApplyTemplateCommandInput, opts ...grpc.CallOption) (*model1.CodeProjectProfile, error) {
	out := new(model1.CodeProjectProfile)
	err := c.cc.Invoke(ctx, GithubCommandController_ApplyTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubCommandControllerClient) AddSecretsToRepo(ctx context.Context, in *model.AddSecretsToRepoCommandInput, opts ...grpc.CallOption) (*model1.CodeProject, error) {
	out := new(model1.CodeProject)
	err := c.cc.Invoke(ctx, GithubCommandController_AddSecretsToRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubCommandControllerClient) AddSecretsToOrg(ctx context.Context, in *model.AddSecretsToOrgCommandInput, opts ...grpc.CallOption) (*model2.CodeServer, error) {
	out := new(model2.CodeServer)
	err := c.cc.Invoke(ctx, GithubCommandController_AddSecretsToOrg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubCommandControllerClient) SynchronizeMagicPipelineFiles(ctx context.Context, in *model.SynchronizeGithubMagicPipelineFilesCommandInput, opts ...grpc.CallOption) (*model1.CodeProject, error) {
	out := new(model1.CodeProject)
	err := c.cc.Invoke(ctx, GithubCommandController_SynchronizeMagicPipelineFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GithubCommandControllerServer is the server API for GithubCommandController service.
// All implementations should embed UnimplementedGithubCommandControllerServer
// for forward compatibility
type GithubCommandControllerServer interface {
	// create a new repository on github
	// https://docs.github.com/en/rest/apps/installations?apiVersion=2022-11-28#add-a-repository-to-an-app-installation
	// https://docs.github.com/en/rest/repos/repos#create-a-repository-for-the-authenticated-user
	// https://docs.github.com/en/rest/repos/repos#create-an-organization-repository
	CreRepository(context.Context, *model.CreRepositoryCommandInput) (*model1.CodeProject, error)
	// apply a cookiecutter template on a code project created on github
	ApplyTemplate(context.Context, *model.GithubApplyTemplateCommandInput) (*model1.CodeProjectProfile, error)
	// add a list of secrets to a github repository
	AddSecretsToRepo(context.Context, *model.AddSecretsToRepoCommandInput) (*model1.CodeProject, error)
	// add a list of secrets to a github organization
	AddSecretsToOrg(context.Context, *model.AddSecretsToOrgCommandInput) (*model2.CodeServer, error)
	// synchronization is achieved by first removing all yaml files prefixed with "pc:" inside .github/workflows and
	// then add files in the input to .github/workflows directory.
	SynchronizeMagicPipelineFiles(context.Context, *model.SynchronizeGithubMagicPipelineFilesCommandInput) (*model1.CodeProject, error)
}

// UnimplementedGithubCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGithubCommandControllerServer struct {
}

func (UnimplementedGithubCommandControllerServer) CreRepository(context.Context, *model.CreRepositoryCommandInput) (*model1.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreRepository not implemented")
}
func (UnimplementedGithubCommandControllerServer) ApplyTemplate(context.Context, *model.GithubApplyTemplateCommandInput) (*model1.CodeProjectProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyTemplate not implemented")
}
func (UnimplementedGithubCommandControllerServer) AddSecretsToRepo(context.Context, *model.AddSecretsToRepoCommandInput) (*model1.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSecretsToRepo not implemented")
}
func (UnimplementedGithubCommandControllerServer) AddSecretsToOrg(context.Context, *model.AddSecretsToOrgCommandInput) (*model2.CodeServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSecretsToOrg not implemented")
}
func (UnimplementedGithubCommandControllerServer) SynchronizeMagicPipelineFiles(context.Context, *model.SynchronizeGithubMagicPipelineFilesCommandInput) (*model1.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SynchronizeMagicPipelineFiles not implemented")
}

// UnsafeGithubCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GithubCommandControllerServer will
// result in compilation errors.
type UnsafeGithubCommandControllerServer interface {
	mustEmbedUnimplementedGithubCommandControllerServer()
}

func RegisterGithubCommandControllerServer(s grpc.ServiceRegistrar, srv GithubCommandControllerServer) {
	s.RegisterService(&GithubCommandController_ServiceDesc, srv)
}

func _GithubCommandController_CreRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CreRepositoryCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubCommandControllerServer).CreRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubCommandController_CreRepository_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubCommandControllerServer).CreRepository(ctx, req.(*model.CreRepositoryCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubCommandController_ApplyTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GithubApplyTemplateCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubCommandControllerServer).ApplyTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubCommandController_ApplyTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubCommandControllerServer).ApplyTemplate(ctx, req.(*model.GithubApplyTemplateCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubCommandController_AddSecretsToRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AddSecretsToRepoCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubCommandControllerServer).AddSecretsToRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubCommandController_AddSecretsToRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubCommandControllerServer).AddSecretsToRepo(ctx, req.(*model.AddSecretsToRepoCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubCommandController_AddSecretsToOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AddSecretsToOrgCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubCommandControllerServer).AddSecretsToOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubCommandController_AddSecretsToOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubCommandControllerServer).AddSecretsToOrg(ctx, req.(*model.AddSecretsToOrgCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubCommandController_SynchronizeMagicPipelineFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.SynchronizeGithubMagicPipelineFilesCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubCommandControllerServer).SynchronizeMagicPipelineFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubCommandController_SynchronizeMagicPipelineFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubCommandControllerServer).SynchronizeMagicPipelineFiles(ctx, req.(*model.SynchronizeGithubMagicPipelineFilesCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GithubCommandController_ServiceDesc is the grpc.ServiceDesc for GithubCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GithubCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.integration.v1.github.service.GithubCommandController",
	HandlerType: (*GithubCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "creRepository",
			Handler:    _GithubCommandController_CreRepository_Handler,
		},
		{
			MethodName: "applyTemplate",
			Handler:    _GithubCommandController_ApplyTemplate_Handler,
		},
		{
			MethodName: "addSecretsToRepo",
			Handler:    _GithubCommandController_AddSecretsToRepo_Handler,
		},
		{
			MethodName: "addSecretsToOrg",
			Handler:    _GithubCommandController_AddSecretsToOrg_Handler,
		},
		{
			MethodName: "synchronizeMagicPipelineFiles",
			Handler:    _GithubCommandController_SynchronizeMagicPipelineFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/integration/v1/github/service/command.proto",
}
