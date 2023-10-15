// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/integration/gitlab/proxy/service.proto

package proxy

import (
	context "context"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/develop/sourcecode/project/rpc"
	rpc1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/develop/sourcecode/server/rpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GitlabProxyCommandController_CreProject_FullMethodName            = "/cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController/creProject"
	GitlabProxyCommandController_ApplyTemplate_FullMethodName         = "/cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController/applyTemplate"
	GitlabProxyCommandController_AddVariablesToProject_FullMethodName = "/cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController/addVariablesToProject"
	GitlabProxyCommandController_AddVariablesToGroup_FullMethodName   = "/cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController/addVariablesToGroup"
	GitlabProxyCommandController_AddFilesToProject_FullMethodName     = "/cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController/addFilesToProject"
)

// GitlabProxyCommandControllerClient is the client API for GitlabProxyCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitlabProxyCommandControllerClient interface {
	// create new project on gitlab
	// https://docs.gitlab.com/ee/api/projects.html#create-project
	CreProject(ctx context.Context, in *CreProjectCommandInput, opts ...grpc.CallOption) (*rpc.CodeProject, error)
	// apply a cookiecutter template on a code project created on gitlab
	ApplyTemplate(ctx context.Context, in *GitlabApplyTemplateCommandInput, opts ...grpc.CallOption) (*rpc.CodeProjectProfile, error)
	// add a list of variables to a gitlab project
	AddVariablesToProject(ctx context.Context, in *AddVariablesToProjectCommandInput, opts ...grpc.CallOption) (*rpc.CodeProject, error)
	// add a list of variables to a gitlab group
	AddVariablesToGroup(ctx context.Context, in *AddVariablesToGroupCommandInput, opts ...grpc.CallOption) (*rpc1.CodeServer, error)
	// add a list of files to a gitlab project
	AddFilesToProject(ctx context.Context, in *AddFilesToGitlabProjectCommandInput, opts ...grpc.CallOption) (*rpc.CodeProject, error)
}

type gitlabProxyCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGitlabProxyCommandControllerClient(cc grpc.ClientConnInterface) GitlabProxyCommandControllerClient {
	return &gitlabProxyCommandControllerClient{cc}
}

func (c *gitlabProxyCommandControllerClient) CreProject(ctx context.Context, in *CreProjectCommandInput, opts ...grpc.CallOption) (*rpc.CodeProject, error) {
	out := new(rpc.CodeProject)
	err := c.cc.Invoke(ctx, GitlabProxyCommandController_CreProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabProxyCommandControllerClient) ApplyTemplate(ctx context.Context, in *GitlabApplyTemplateCommandInput, opts ...grpc.CallOption) (*rpc.CodeProjectProfile, error) {
	out := new(rpc.CodeProjectProfile)
	err := c.cc.Invoke(ctx, GitlabProxyCommandController_ApplyTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabProxyCommandControllerClient) AddVariablesToProject(ctx context.Context, in *AddVariablesToProjectCommandInput, opts ...grpc.CallOption) (*rpc.CodeProject, error) {
	out := new(rpc.CodeProject)
	err := c.cc.Invoke(ctx, GitlabProxyCommandController_AddVariablesToProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabProxyCommandControllerClient) AddVariablesToGroup(ctx context.Context, in *AddVariablesToGroupCommandInput, opts ...grpc.CallOption) (*rpc1.CodeServer, error) {
	out := new(rpc1.CodeServer)
	err := c.cc.Invoke(ctx, GitlabProxyCommandController_AddVariablesToGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabProxyCommandControllerClient) AddFilesToProject(ctx context.Context, in *AddFilesToGitlabProjectCommandInput, opts ...grpc.CallOption) (*rpc.CodeProject, error) {
	out := new(rpc.CodeProject)
	err := c.cc.Invoke(ctx, GitlabProxyCommandController_AddFilesToProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitlabProxyCommandControllerServer is the server API for GitlabProxyCommandController service.
// All implementations should embed UnimplementedGitlabProxyCommandControllerServer
// for forward compatibility
type GitlabProxyCommandControllerServer interface {
	// create new project on gitlab
	// https://docs.gitlab.com/ee/api/projects.html#create-project
	CreProject(context.Context, *CreProjectCommandInput) (*rpc.CodeProject, error)
	// apply a cookiecutter template on a code project created on gitlab
	ApplyTemplate(context.Context, *GitlabApplyTemplateCommandInput) (*rpc.CodeProjectProfile, error)
	// add a list of variables to a gitlab project
	AddVariablesToProject(context.Context, *AddVariablesToProjectCommandInput) (*rpc.CodeProject, error)
	// add a list of variables to a gitlab group
	AddVariablesToGroup(context.Context, *AddVariablesToGroupCommandInput) (*rpc1.CodeServer, error)
	// add a list of files to a gitlab project
	AddFilesToProject(context.Context, *AddFilesToGitlabProjectCommandInput) (*rpc.CodeProject, error)
}

// UnimplementedGitlabProxyCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGitlabProxyCommandControllerServer struct {
}

func (UnimplementedGitlabProxyCommandControllerServer) CreProject(context.Context, *CreProjectCommandInput) (*rpc.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreProject not implemented")
}
func (UnimplementedGitlabProxyCommandControllerServer) ApplyTemplate(context.Context, *GitlabApplyTemplateCommandInput) (*rpc.CodeProjectProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyTemplate not implemented")
}
func (UnimplementedGitlabProxyCommandControllerServer) AddVariablesToProject(context.Context, *AddVariablesToProjectCommandInput) (*rpc.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVariablesToProject not implemented")
}
func (UnimplementedGitlabProxyCommandControllerServer) AddVariablesToGroup(context.Context, *AddVariablesToGroupCommandInput) (*rpc1.CodeServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVariablesToGroup not implemented")
}
func (UnimplementedGitlabProxyCommandControllerServer) AddFilesToProject(context.Context, *AddFilesToGitlabProjectCommandInput) (*rpc.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFilesToProject not implemented")
}

// UnsafeGitlabProxyCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitlabProxyCommandControllerServer will
// result in compilation errors.
type UnsafeGitlabProxyCommandControllerServer interface {
	mustEmbedUnimplementedGitlabProxyCommandControllerServer()
}

func RegisterGitlabProxyCommandControllerServer(s grpc.ServiceRegistrar, srv GitlabProxyCommandControllerServer) {
	s.RegisterService(&GitlabProxyCommandController_ServiceDesc, srv)
}

func _GitlabProxyCommandController_CreProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreProjectCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabProxyCommandControllerServer).CreProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabProxyCommandController_CreProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabProxyCommandControllerServer).CreProject(ctx, req.(*CreProjectCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabProxyCommandController_ApplyTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GitlabApplyTemplateCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabProxyCommandControllerServer).ApplyTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabProxyCommandController_ApplyTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabProxyCommandControllerServer).ApplyTemplate(ctx, req.(*GitlabApplyTemplateCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabProxyCommandController_AddVariablesToProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVariablesToProjectCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabProxyCommandControllerServer).AddVariablesToProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabProxyCommandController_AddVariablesToProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabProxyCommandControllerServer).AddVariablesToProject(ctx, req.(*AddVariablesToProjectCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabProxyCommandController_AddVariablesToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVariablesToGroupCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabProxyCommandControllerServer).AddVariablesToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabProxyCommandController_AddVariablesToGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabProxyCommandControllerServer).AddVariablesToGroup(ctx, req.(*AddVariablesToGroupCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabProxyCommandController_AddFilesToProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFilesToGitlabProjectCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabProxyCommandControllerServer).AddFilesToProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabProxyCommandController_AddFilesToProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabProxyCommandControllerServer).AddFilesToProject(ctx, req.(*AddFilesToGitlabProjectCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GitlabProxyCommandController_ServiceDesc is the grpc.ServiceDesc for GitlabProxyCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitlabProxyCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController",
	HandlerType: (*GitlabProxyCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "creProject",
			Handler:    _GitlabProxyCommandController_CreProject_Handler,
		},
		{
			MethodName: "applyTemplate",
			Handler:    _GitlabProxyCommandController_ApplyTemplate_Handler,
		},
		{
			MethodName: "addVariablesToProject",
			Handler:    _GitlabProxyCommandController_AddVariablesToProject_Handler,
		},
		{
			MethodName: "addVariablesToGroup",
			Handler:    _GitlabProxyCommandController_AddVariablesToGroup_Handler,
		},
		{
			MethodName: "addFilesToProject",
			Handler:    _GitlabProxyCommandController_AddFilesToProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/integration/gitlab/proxy/service.proto",
}

const (
	GitlabProxyQueryController_ListProjects_FullMethodName                = "/cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController/listProjects"
	GitlabProxyQueryController_GetProject_FullMethodName                  = "/cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController/getProject"
	GitlabProxyQueryController_GetGroup_FullMethodName                    = "/cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController/getGroup"
	GitlabProxyQueryController_GetGitlabCodeProjectProfile_FullMethodName = "/cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController/getGitlabCodeProjectProfile"
)

// GitlabProxyQueryControllerClient is the client API for GitlabProxyQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitlabProxyQueryControllerClient interface {
	// list projects for the requested group, including projects in sub-groups on gitlab
	// https://docs.gitlab.com/ee/api/groups.html#list-a-groups-projects
	// todo: we have to add pagination support for response.
	ListProjects(ctx context.Context, in *ListProjectsQueryInput, opts ...grpc.CallOption) (*rpc.CodeProjects, error)
	// get details of a project on gitlab
	GetProject(ctx context.Context, in *GetProjectQueryInput, opts ...grpc.CallOption) (*rpc.CodeProject, error)
	// get details of a group on gitlab
	GetGroup(ctx context.Context, in *GetGroupQueryInput, opts ...grpc.CallOption) (*rpc1.CodeServer, error)
	// get code project profile of a code-project hosted on gitlab
	GetGitlabCodeProjectProfile(ctx context.Context, in *GetGitlabCodeProjectProfileQueryInput, opts ...grpc.CallOption) (*rpc.CodeProjectProfile, error)
}

type gitlabProxyQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGitlabProxyQueryControllerClient(cc grpc.ClientConnInterface) GitlabProxyQueryControllerClient {
	return &gitlabProxyQueryControllerClient{cc}
}

func (c *gitlabProxyQueryControllerClient) ListProjects(ctx context.Context, in *ListProjectsQueryInput, opts ...grpc.CallOption) (*rpc.CodeProjects, error) {
	out := new(rpc.CodeProjects)
	err := c.cc.Invoke(ctx, GitlabProxyQueryController_ListProjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabProxyQueryControllerClient) GetProject(ctx context.Context, in *GetProjectQueryInput, opts ...grpc.CallOption) (*rpc.CodeProject, error) {
	out := new(rpc.CodeProject)
	err := c.cc.Invoke(ctx, GitlabProxyQueryController_GetProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabProxyQueryControllerClient) GetGroup(ctx context.Context, in *GetGroupQueryInput, opts ...grpc.CallOption) (*rpc1.CodeServer, error) {
	out := new(rpc1.CodeServer)
	err := c.cc.Invoke(ctx, GitlabProxyQueryController_GetGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabProxyQueryControllerClient) GetGitlabCodeProjectProfile(ctx context.Context, in *GetGitlabCodeProjectProfileQueryInput, opts ...grpc.CallOption) (*rpc.CodeProjectProfile, error) {
	out := new(rpc.CodeProjectProfile)
	err := c.cc.Invoke(ctx, GitlabProxyQueryController_GetGitlabCodeProjectProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitlabProxyQueryControllerServer is the server API for GitlabProxyQueryController service.
// All implementations should embed UnimplementedGitlabProxyQueryControllerServer
// for forward compatibility
type GitlabProxyQueryControllerServer interface {
	// list projects for the requested group, including projects in sub-groups on gitlab
	// https://docs.gitlab.com/ee/api/groups.html#list-a-groups-projects
	// todo: we have to add pagination support for response.
	ListProjects(context.Context, *ListProjectsQueryInput) (*rpc.CodeProjects, error)
	// get details of a project on gitlab
	GetProject(context.Context, *GetProjectQueryInput) (*rpc.CodeProject, error)
	// get details of a group on gitlab
	GetGroup(context.Context, *GetGroupQueryInput) (*rpc1.CodeServer, error)
	// get code project profile of a code-project hosted on gitlab
	GetGitlabCodeProjectProfile(context.Context, *GetGitlabCodeProjectProfileQueryInput) (*rpc.CodeProjectProfile, error)
}

// UnimplementedGitlabProxyQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGitlabProxyQueryControllerServer struct {
}

func (UnimplementedGitlabProxyQueryControllerServer) ListProjects(context.Context, *ListProjectsQueryInput) (*rpc.CodeProjects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (UnimplementedGitlabProxyQueryControllerServer) GetProject(context.Context, *GetProjectQueryInput) (*rpc.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedGitlabProxyQueryControllerServer) GetGroup(context.Context, *GetGroupQueryInput) (*rpc1.CodeServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedGitlabProxyQueryControllerServer) GetGitlabCodeProjectProfile(context.Context, *GetGitlabCodeProjectProfileQueryInput) (*rpc.CodeProjectProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGitlabCodeProjectProfile not implemented")
}

// UnsafeGitlabProxyQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitlabProxyQueryControllerServer will
// result in compilation errors.
type UnsafeGitlabProxyQueryControllerServer interface {
	mustEmbedUnimplementedGitlabProxyQueryControllerServer()
}

func RegisterGitlabProxyQueryControllerServer(s grpc.ServiceRegistrar, srv GitlabProxyQueryControllerServer) {
	s.RegisterService(&GitlabProxyQueryController_ServiceDesc, srv)
}

func _GitlabProxyQueryController_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabProxyQueryControllerServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabProxyQueryController_ListProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabProxyQueryControllerServer).ListProjects(ctx, req.(*ListProjectsQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabProxyQueryController_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabProxyQueryControllerServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabProxyQueryController_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabProxyQueryControllerServer).GetProject(ctx, req.(*GetProjectQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabProxyQueryController_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabProxyQueryControllerServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabProxyQueryController_GetGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabProxyQueryControllerServer).GetGroup(ctx, req.(*GetGroupQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabProxyQueryController_GetGitlabCodeProjectProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGitlabCodeProjectProfileQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabProxyQueryControllerServer).GetGitlabCodeProjectProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabProxyQueryController_GetGitlabCodeProjectProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabProxyQueryControllerServer).GetGitlabCodeProjectProfile(ctx, req.(*GetGitlabCodeProjectProfileQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GitlabProxyQueryController_ServiceDesc is the grpc.ServiceDesc for GitlabProxyQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitlabProxyQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController",
	HandlerType: (*GitlabProxyQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listProjects",
			Handler:    _GitlabProxyQueryController_ListProjects_Handler,
		},
		{
			MethodName: "getProject",
			Handler:    _GitlabProxyQueryController_GetProject_Handler,
		},
		{
			MethodName: "getGroup",
			Handler:    _GitlabProxyQueryController_GetGroup_Handler,
		},
		{
			MethodName: "getGitlabCodeProjectProfile",
			Handler:    _GitlabProxyQueryController_GetGitlabCodeProjectProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/integration/gitlab/proxy/service.proto",
}
