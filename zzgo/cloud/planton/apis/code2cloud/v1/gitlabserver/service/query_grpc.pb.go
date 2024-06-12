// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/gitlabserver/service/query.proto

package service

import (
	context "context"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/environment/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gitlabserver/model"
	model3 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/resourcemanager/v1/product/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GitlabServerQueryController_GetById_FullMethodName             = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerQueryController/getById"
	GitlabServerQueryController_FindByProductId_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerQueryController/findByProductId"
	GitlabServerQueryController_FindByEnvironmentId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerQueryController/findByEnvironmentId"
	GitlabServerQueryController_FindByKubeClusterId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerQueryController/findByKubeClusterId"
	GitlabServerQueryController_GetPassword_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerQueryController/getPassword"
)

// GitlabServerQueryControllerClient is the client API for GitlabServerQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitlabServerQueryControllerClient interface {
	// look up gitlab-server using gitlab-server id
	GetById(ctx context.Context, in *model.GitlabServerId, opts ...grpc.CallOption) (*model.GitlabServer, error)
	// find gitlab-servers by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.GitlabServers, error)
	// find gitlab-servers by environment
	FindByEnvironmentId(ctx context.Context, in *model2.EnvironmentId, opts ...grpc.CallOption) (*model.GitlabServers, error)
	FindByKubeClusterId(ctx context.Context, in *model3.KubeClusterId, opts ...grpc.CallOption) (*model.GitlabServers, error)
	// look up gitlab-server sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *model.GitlabServerId, opts ...grpc.CallOption) (*model.GitlabServerPassword, error)
}

type gitlabServerQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGitlabServerQueryControllerClient(cc grpc.ClientConnInterface) GitlabServerQueryControllerClient {
	return &gitlabServerQueryControllerClient{cc}
}

func (c *gitlabServerQueryControllerClient) GetById(ctx context.Context, in *model.GitlabServerId, opts ...grpc.CallOption) (*model.GitlabServer, error) {
	out := new(model.GitlabServer)
	err := c.cc.Invoke(ctx, GitlabServerQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerQueryControllerClient) FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.GitlabServers, error) {
	out := new(model.GitlabServers)
	err := c.cc.Invoke(ctx, GitlabServerQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerQueryControllerClient) FindByEnvironmentId(ctx context.Context, in *model2.EnvironmentId, opts ...grpc.CallOption) (*model.GitlabServers, error) {
	out := new(model.GitlabServers)
	err := c.cc.Invoke(ctx, GitlabServerQueryController_FindByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerQueryControllerClient) FindByKubeClusterId(ctx context.Context, in *model3.KubeClusterId, opts ...grpc.CallOption) (*model.GitlabServers, error) {
	out := new(model.GitlabServers)
	err := c.cc.Invoke(ctx, GitlabServerQueryController_FindByKubeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabServerQueryControllerClient) GetPassword(ctx context.Context, in *model.GitlabServerId, opts ...grpc.CallOption) (*model.GitlabServerPassword, error) {
	out := new(model.GitlabServerPassword)
	err := c.cc.Invoke(ctx, GitlabServerQueryController_GetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitlabServerQueryControllerServer is the server API for GitlabServerQueryController service.
// All implementations should embed UnimplementedGitlabServerQueryControllerServer
// for forward compatibility
type GitlabServerQueryControllerServer interface {
	// look up gitlab-server using gitlab-server id
	GetById(context.Context, *model.GitlabServerId) (*model.GitlabServer, error)
	// find gitlab-servers by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(context.Context, *model1.ProductId) (*model.GitlabServers, error)
	// find gitlab-servers by environment
	FindByEnvironmentId(context.Context, *model2.EnvironmentId) (*model.GitlabServers, error)
	FindByKubeClusterId(context.Context, *model3.KubeClusterId) (*model.GitlabServers, error)
	// look up gitlab-server sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model.GitlabServerId) (*model.GitlabServerPassword, error)
}

// UnimplementedGitlabServerQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGitlabServerQueryControllerServer struct {
}

func (UnimplementedGitlabServerQueryControllerServer) GetById(context.Context, *model.GitlabServerId) (*model.GitlabServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedGitlabServerQueryControllerServer) FindByProductId(context.Context, *model1.ProductId) (*model.GitlabServers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedGitlabServerQueryControllerServer) FindByEnvironmentId(context.Context, *model2.EnvironmentId) (*model.GitlabServers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEnvironmentId not implemented")
}
func (UnimplementedGitlabServerQueryControllerServer) FindByKubeClusterId(context.Context, *model3.KubeClusterId) (*model.GitlabServers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKubeClusterId not implemented")
}
func (UnimplementedGitlabServerQueryControllerServer) GetPassword(context.Context, *model.GitlabServerId) (*model.GitlabServerPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}

// UnsafeGitlabServerQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitlabServerQueryControllerServer will
// result in compilation errors.
type UnsafeGitlabServerQueryControllerServer interface {
	mustEmbedUnimplementedGitlabServerQueryControllerServer()
}

func RegisterGitlabServerQueryControllerServer(s grpc.ServiceRegistrar, srv GitlabServerQueryControllerServer) {
	s.RegisterService(&GitlabServerQueryController_ServiceDesc, srv)
}

func _GitlabServerQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerQueryControllerServer).GetById(ctx, req.(*model.GitlabServerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerQueryControllerServer).FindByProductId(ctx, req.(*model1.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerQueryController_FindByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.EnvironmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerQueryControllerServer).FindByEnvironmentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerQueryController_FindByEnvironmentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerQueryControllerServer).FindByEnvironmentId(ctx, req.(*model2.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerQueryController_FindByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model3.KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerQueryControllerServer).FindByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerQueryController_FindByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerQueryControllerServer).FindByKubeClusterId(ctx, req.(*model3.KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitlabServerQueryController_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerQueryControllerServer).GetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerQueryController_GetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerQueryControllerServer).GetPassword(ctx, req.(*model.GitlabServerId))
	}
	return interceptor(ctx, in, info, handler)
}

// GitlabServerQueryController_ServiceDesc is the grpc.ServiceDesc for GitlabServerQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitlabServerQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.gitlabserver.service.GitlabServerQueryController",
	HandlerType: (*GitlabServerQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _GitlabServerQueryController_GetById_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _GitlabServerQueryController_FindByProductId_Handler,
		},
		{
			MethodName: "findByEnvironmentId",
			Handler:    _GitlabServerQueryController_FindByEnvironmentId_Handler,
		},
		{
			MethodName: "findByKubeClusterId",
			Handler:    _GitlabServerQueryController_FindByKubeClusterId_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _GitlabServerQueryController_GetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/gitlabserver/service/query.proto",
}
