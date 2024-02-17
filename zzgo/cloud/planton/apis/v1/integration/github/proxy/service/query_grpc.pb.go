// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/integration/github/proxy/service/query.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/sourcecode/project/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/integration/github/proxy/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GithubProxyQueryController_ListRepositories_FullMethodName            = "/cloud.planton.apis.v1.integration.github.proxy.service.GithubProxyQueryController/listRepositories"
	GithubProxyQueryController_GetRepository_FullMethodName               = "/cloud.planton.apis.v1.integration.github.proxy.service.GithubProxyQueryController/getRepository"
	GithubProxyQueryController_GetGithubAppInstallation_FullMethodName    = "/cloud.planton.apis.v1.integration.github.proxy.service.GithubProxyQueryController/getGithubAppInstallation"
	GithubProxyQueryController_GetGithubCodeProjectProfile_FullMethodName = "/cloud.planton.apis.v1.integration.github.proxy.service.GithubProxyQueryController/getGithubCodeProjectProfile"
)

// GithubProxyQueryControllerClient is the client API for GithubProxyQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GithubProxyQueryControllerClient interface {
	// list repositories for the requested organization or user on github
	// https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#list-organization-repositories
	// todo: we have to add pagination support for response.
	ListRepositories(ctx context.Context, in *model.ListRepositoriesQueryInput, opts ...grpc.CallOption) (*model1.CodeProjects, error)
	// get the details of a repository on github
	GetRepository(ctx context.Context, in *model.GetRepositoryQueryInput, opts ...grpc.CallOption) (*model1.CodeProject, error)
	// get details of a github app installation
	GetGithubAppInstallation(ctx context.Context, in *model.GithubAppInstallation, opts ...grpc.CallOption) (*model.GithubAppInstallation, error)
	// get code project profile of a code-project hosted on github
	GetGithubCodeProjectProfile(ctx context.Context, in *model.GetGithubCodeProjectProfileQueryInput, opts ...grpc.CallOption) (*model1.CodeProjectProfile, error)
}

type githubProxyQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGithubProxyQueryControllerClient(cc grpc.ClientConnInterface) GithubProxyQueryControllerClient {
	return &githubProxyQueryControllerClient{cc}
}

func (c *githubProxyQueryControllerClient) ListRepositories(ctx context.Context, in *model.ListRepositoriesQueryInput, opts ...grpc.CallOption) (*model1.CodeProjects, error) {
	out := new(model1.CodeProjects)
	err := c.cc.Invoke(ctx, GithubProxyQueryController_ListRepositories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubProxyQueryControllerClient) GetRepository(ctx context.Context, in *model.GetRepositoryQueryInput, opts ...grpc.CallOption) (*model1.CodeProject, error) {
	out := new(model1.CodeProject)
	err := c.cc.Invoke(ctx, GithubProxyQueryController_GetRepository_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubProxyQueryControllerClient) GetGithubAppInstallation(ctx context.Context, in *model.GithubAppInstallation, opts ...grpc.CallOption) (*model.GithubAppInstallation, error) {
	out := new(model.GithubAppInstallation)
	err := c.cc.Invoke(ctx, GithubProxyQueryController_GetGithubAppInstallation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubProxyQueryControllerClient) GetGithubCodeProjectProfile(ctx context.Context, in *model.GetGithubCodeProjectProfileQueryInput, opts ...grpc.CallOption) (*model1.CodeProjectProfile, error) {
	out := new(model1.CodeProjectProfile)
	err := c.cc.Invoke(ctx, GithubProxyQueryController_GetGithubCodeProjectProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GithubProxyQueryControllerServer is the server API for GithubProxyQueryController service.
// All implementations should embed UnimplementedGithubProxyQueryControllerServer
// for forward compatibility
type GithubProxyQueryControllerServer interface {
	// list repositories for the requested organization or user on github
	// https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#list-organization-repositories
	// todo: we have to add pagination support for response.
	ListRepositories(context.Context, *model.ListRepositoriesQueryInput) (*model1.CodeProjects, error)
	// get the details of a repository on github
	GetRepository(context.Context, *model.GetRepositoryQueryInput) (*model1.CodeProject, error)
	// get details of a github app installation
	GetGithubAppInstallation(context.Context, *model.GithubAppInstallation) (*model.GithubAppInstallation, error)
	// get code project profile of a code-project hosted on github
	GetGithubCodeProjectProfile(context.Context, *model.GetGithubCodeProjectProfileQueryInput) (*model1.CodeProjectProfile, error)
}

// UnimplementedGithubProxyQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGithubProxyQueryControllerServer struct {
}

func (UnimplementedGithubProxyQueryControllerServer) ListRepositories(context.Context, *model.ListRepositoriesQueryInput) (*model1.CodeProjects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepositories not implemented")
}
func (UnimplementedGithubProxyQueryControllerServer) GetRepository(context.Context, *model.GetRepositoryQueryInput) (*model1.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepository not implemented")
}
func (UnimplementedGithubProxyQueryControllerServer) GetGithubAppInstallation(context.Context, *model.GithubAppInstallation) (*model.GithubAppInstallation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGithubAppInstallation not implemented")
}
func (UnimplementedGithubProxyQueryControllerServer) GetGithubCodeProjectProfile(context.Context, *model.GetGithubCodeProjectProfileQueryInput) (*model1.CodeProjectProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGithubCodeProjectProfile not implemented")
}

// UnsafeGithubProxyQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GithubProxyQueryControllerServer will
// result in compilation errors.
type UnsafeGithubProxyQueryControllerServer interface {
	mustEmbedUnimplementedGithubProxyQueryControllerServer()
}

func RegisterGithubProxyQueryControllerServer(s grpc.ServiceRegistrar, srv GithubProxyQueryControllerServer) {
	s.RegisterService(&GithubProxyQueryController_ServiceDesc, srv)
}

func _GithubProxyQueryController_ListRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ListRepositoriesQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubProxyQueryControllerServer).ListRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubProxyQueryController_ListRepositories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubProxyQueryControllerServer).ListRepositories(ctx, req.(*model.ListRepositoriesQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubProxyQueryController_GetRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetRepositoryQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubProxyQueryControllerServer).GetRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubProxyQueryController_GetRepository_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubProxyQueryControllerServer).GetRepository(ctx, req.(*model.GetRepositoryQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubProxyQueryController_GetGithubAppInstallation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GithubAppInstallation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubProxyQueryControllerServer).GetGithubAppInstallation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubProxyQueryController_GetGithubAppInstallation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubProxyQueryControllerServer).GetGithubAppInstallation(ctx, req.(*model.GithubAppInstallation))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubProxyQueryController_GetGithubCodeProjectProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetGithubCodeProjectProfileQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubProxyQueryControllerServer).GetGithubCodeProjectProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubProxyQueryController_GetGithubCodeProjectProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubProxyQueryControllerServer).GetGithubCodeProjectProfile(ctx, req.(*model.GetGithubCodeProjectProfileQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GithubProxyQueryController_ServiceDesc is the grpc.ServiceDesc for GithubProxyQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GithubProxyQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.integration.github.proxy.service.GithubProxyQueryController",
	HandlerType: (*GithubProxyQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listRepositories",
			Handler:    _GithubProxyQueryController_ListRepositories_Handler,
		},
		{
			MethodName: "getRepository",
			Handler:    _GithubProxyQueryController_GetRepository_Handler,
		},
		{
			MethodName: "getGithubAppInstallation",
			Handler:    _GithubProxyQueryController_GetGithubAppInstallation_Handler,
		},
		{
			MethodName: "getGithubCodeProjectProfile",
			Handler:    _GithubProxyQueryController_GetGithubCodeProjectProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/integration/github/proxy/service/query.proto",
}
