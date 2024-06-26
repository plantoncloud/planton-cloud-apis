// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/integration/v1/github/service/query.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/codeproject/model"
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
	GithubQueryController_ListRepositories_FullMethodName            = "/cloud.planton.apis.integration.v1.github.service.GithubQueryController/listRepositories"
	GithubQueryController_GetRepository_FullMethodName               = "/cloud.planton.apis.integration.v1.github.service.GithubQueryController/getRepository"
	GithubQueryController_GetGithubAppInstallation_FullMethodName    = "/cloud.planton.apis.integration.v1.github.service.GithubQueryController/getGithubAppInstallation"
	GithubQueryController_GetGithubCodeProjectProfile_FullMethodName = "/cloud.planton.apis.integration.v1.github.service.GithubQueryController/getGithubCodeProjectProfile"
)

// GithubQueryControllerClient is the client API for GithubQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GithubQueryControllerClient interface {
	// list repositories for the requested organization or user on github
	// https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#list-organization-repositories
	// todo: we have to add pagination support for response.
	ListRepositories(ctx context.Context, in *model.ListRepositoriesQueryInput, opts ...grpc.CallOption) (*model1.CodeProjectList, error)
	// get the details of a repository on github
	GetRepository(ctx context.Context, in *model.GetRepositoryQueryInput, opts ...grpc.CallOption) (*model1.CodeProject, error)
	// get details of a github app installation
	GetGithubAppInstallation(ctx context.Context, in *model.GithubAppInstallation, opts ...grpc.CallOption) (*model.GithubAppInstallation, error)
	// get code project profile of a code-project hosted on github
	GetGithubCodeProjectProfile(ctx context.Context, in *model.GetGithubCodeProjectProfileQueryInput, opts ...grpc.CallOption) (*model1.CodeProjectProfile, error)
}

type githubQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGithubQueryControllerClient(cc grpc.ClientConnInterface) GithubQueryControllerClient {
	return &githubQueryControllerClient{cc}
}

func (c *githubQueryControllerClient) ListRepositories(ctx context.Context, in *model.ListRepositoriesQueryInput, opts ...grpc.CallOption) (*model1.CodeProjectList, error) {
	out := new(model1.CodeProjectList)
	err := c.cc.Invoke(ctx, GithubQueryController_ListRepositories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubQueryControllerClient) GetRepository(ctx context.Context, in *model.GetRepositoryQueryInput, opts ...grpc.CallOption) (*model1.CodeProject, error) {
	out := new(model1.CodeProject)
	err := c.cc.Invoke(ctx, GithubQueryController_GetRepository_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubQueryControllerClient) GetGithubAppInstallation(ctx context.Context, in *model.GithubAppInstallation, opts ...grpc.CallOption) (*model.GithubAppInstallation, error) {
	out := new(model.GithubAppInstallation)
	err := c.cc.Invoke(ctx, GithubQueryController_GetGithubAppInstallation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubQueryControllerClient) GetGithubCodeProjectProfile(ctx context.Context, in *model.GetGithubCodeProjectProfileQueryInput, opts ...grpc.CallOption) (*model1.CodeProjectProfile, error) {
	out := new(model1.CodeProjectProfile)
	err := c.cc.Invoke(ctx, GithubQueryController_GetGithubCodeProjectProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GithubQueryControllerServer is the server API for GithubQueryController service.
// All implementations should embed UnimplementedGithubQueryControllerServer
// for forward compatibility
type GithubQueryControllerServer interface {
	// list repositories for the requested organization or user on github
	// https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#list-organization-repositories
	// todo: we have to add pagination support for response.
	ListRepositories(context.Context, *model.ListRepositoriesQueryInput) (*model1.CodeProjectList, error)
	// get the details of a repository on github
	GetRepository(context.Context, *model.GetRepositoryQueryInput) (*model1.CodeProject, error)
	// get details of a github app installation
	GetGithubAppInstallation(context.Context, *model.GithubAppInstallation) (*model.GithubAppInstallation, error)
	// get code project profile of a code-project hosted on github
	GetGithubCodeProjectProfile(context.Context, *model.GetGithubCodeProjectProfileQueryInput) (*model1.CodeProjectProfile, error)
}

// UnimplementedGithubQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGithubQueryControllerServer struct {
}

func (UnimplementedGithubQueryControllerServer) ListRepositories(context.Context, *model.ListRepositoriesQueryInput) (*model1.CodeProjectList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepositories not implemented")
}
func (UnimplementedGithubQueryControllerServer) GetRepository(context.Context, *model.GetRepositoryQueryInput) (*model1.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepository not implemented")
}
func (UnimplementedGithubQueryControllerServer) GetGithubAppInstallation(context.Context, *model.GithubAppInstallation) (*model.GithubAppInstallation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGithubAppInstallation not implemented")
}
func (UnimplementedGithubQueryControllerServer) GetGithubCodeProjectProfile(context.Context, *model.GetGithubCodeProjectProfileQueryInput) (*model1.CodeProjectProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGithubCodeProjectProfile not implemented")
}

// UnsafeGithubQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GithubQueryControllerServer will
// result in compilation errors.
type UnsafeGithubQueryControllerServer interface {
	mustEmbedUnimplementedGithubQueryControllerServer()
}

func RegisterGithubQueryControllerServer(s grpc.ServiceRegistrar, srv GithubQueryControllerServer) {
	s.RegisterService(&GithubQueryController_ServiceDesc, srv)
}

func _GithubQueryController_ListRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ListRepositoriesQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubQueryControllerServer).ListRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubQueryController_ListRepositories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubQueryControllerServer).ListRepositories(ctx, req.(*model.ListRepositoriesQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubQueryController_GetRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetRepositoryQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubQueryControllerServer).GetRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubQueryController_GetRepository_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubQueryControllerServer).GetRepository(ctx, req.(*model.GetRepositoryQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubQueryController_GetGithubAppInstallation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GithubAppInstallation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubQueryControllerServer).GetGithubAppInstallation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubQueryController_GetGithubAppInstallation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubQueryControllerServer).GetGithubAppInstallation(ctx, req.(*model.GithubAppInstallation))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubQueryController_GetGithubCodeProjectProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetGithubCodeProjectProfileQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubQueryControllerServer).GetGithubCodeProjectProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubQueryController_GetGithubCodeProjectProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubQueryControllerServer).GetGithubCodeProjectProfile(ctx, req.(*model.GetGithubCodeProjectProfileQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GithubQueryController_ServiceDesc is the grpc.ServiceDesc for GithubQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GithubQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.integration.v1.github.service.GithubQueryController",
	HandlerType: (*GithubQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listRepositories",
			Handler:    _GithubQueryController_ListRepositories_Handler,
		},
		{
			MethodName: "getRepository",
			Handler:    _GithubQueryController_GetRepository_Handler,
		},
		{
			MethodName: "getGithubAppInstallation",
			Handler:    _GithubQueryController_GetGithubAppInstallation_Handler,
		},
		{
			MethodName: "getGithubCodeProjectProfile",
			Handler:    _GithubQueryController_GetGithubCodeProjectProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/integration/v1/github/service/query.proto",
}
