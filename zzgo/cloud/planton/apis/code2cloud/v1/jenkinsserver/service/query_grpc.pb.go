// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/jenkinsserver/service/query.proto

package service

import (
	context "context"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/environment/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/jenkinsserver/model"
	model3 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/model"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/rpc"
	model4 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/apiresources/model"
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
	JenkinsServerQueryController_List_FullMethodName                = "/cloud.planton.apis.code2cloud.v1.jenkinsserver.service.JenkinsServerQueryController/list"
	JenkinsServerQueryController_GetById_FullMethodName             = "/cloud.planton.apis.code2cloud.v1.jenkinsserver.service.JenkinsServerQueryController/getById"
	JenkinsServerQueryController_FindByProductId_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.jenkinsserver.service.JenkinsServerQueryController/findByProductId"
	JenkinsServerQueryController_FindByEnvironmentId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.jenkinsserver.service.JenkinsServerQueryController/findByEnvironmentId"
	JenkinsServerQueryController_FindByKubeClusterId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.jenkinsserver.service.JenkinsServerQueryController/findByKubeClusterId"
	JenkinsServerQueryController_FindPods_FullMethodName            = "/cloud.planton.apis.code2cloud.v1.jenkinsserver.service.JenkinsServerQueryController/findPods"
)

// JenkinsServerQueryControllerClient is the client API for JenkinsServerQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JenkinsServerQueryControllerClient interface {
	// list all jenkins-servers on planton cluster for the requested page.
	List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.JenkinsServerList, error)
	// look up jenkins-server using jenkins-server id
	GetById(ctx context.Context, in *model.JenkinsServerId, opts ...grpc.CallOption) (*model.JenkinsServer, error)
	// find jenkins-servers by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.JenkinsServers, error)
	// find jenkins-servers by environment
	FindByEnvironmentId(ctx context.Context, in *model2.EnvironmentId, opts ...grpc.CallOption) (*model.JenkinsServers, error)
	FindByKubeClusterId(ctx context.Context, in *model3.KubeClusterId, opts ...grpc.CallOption) (*model.JenkinsServers, error)
	// lookup pods of a jenkins-server deployed to a environment
	FindPods(ctx context.Context, in *model.JenkinsServerId, opts ...grpc.CallOption) (*model4.Pods, error)
}

type jenkinsServerQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewJenkinsServerQueryControllerClient(cc grpc.ClientConnInterface) JenkinsServerQueryControllerClient {
	return &jenkinsServerQueryControllerClient{cc}
}

func (c *jenkinsServerQueryControllerClient) List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.JenkinsServerList, error) {
	out := new(model.JenkinsServerList)
	err := c.cc.Invoke(ctx, JenkinsServerQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsServerQueryControllerClient) GetById(ctx context.Context, in *model.JenkinsServerId, opts ...grpc.CallOption) (*model.JenkinsServer, error) {
	out := new(model.JenkinsServer)
	err := c.cc.Invoke(ctx, JenkinsServerQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsServerQueryControllerClient) FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.JenkinsServers, error) {
	out := new(model.JenkinsServers)
	err := c.cc.Invoke(ctx, JenkinsServerQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsServerQueryControllerClient) FindByEnvironmentId(ctx context.Context, in *model2.EnvironmentId, opts ...grpc.CallOption) (*model.JenkinsServers, error) {
	out := new(model.JenkinsServers)
	err := c.cc.Invoke(ctx, JenkinsServerQueryController_FindByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsServerQueryControllerClient) FindByKubeClusterId(ctx context.Context, in *model3.KubeClusterId, opts ...grpc.CallOption) (*model.JenkinsServers, error) {
	out := new(model.JenkinsServers)
	err := c.cc.Invoke(ctx, JenkinsServerQueryController_FindByKubeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jenkinsServerQueryControllerClient) FindPods(ctx context.Context, in *model.JenkinsServerId, opts ...grpc.CallOption) (*model4.Pods, error) {
	out := new(model4.Pods)
	err := c.cc.Invoke(ctx, JenkinsServerQueryController_FindPods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JenkinsServerQueryControllerServer is the server API for JenkinsServerQueryController service.
// All implementations should embed UnimplementedJenkinsServerQueryControllerServer
// for forward compatibility
type JenkinsServerQueryControllerServer interface {
	// list all jenkins-servers on planton cluster for the requested page.
	List(context.Context, *rpc.PageInfo) (*model.JenkinsServerList, error)
	// look up jenkins-server using jenkins-server id
	GetById(context.Context, *model.JenkinsServerId) (*model.JenkinsServer, error)
	// find jenkins-servers by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(context.Context, *model1.ProductId) (*model.JenkinsServers, error)
	// find jenkins-servers by environment
	FindByEnvironmentId(context.Context, *model2.EnvironmentId) (*model.JenkinsServers, error)
	FindByKubeClusterId(context.Context, *model3.KubeClusterId) (*model.JenkinsServers, error)
	// lookup pods of a jenkins-server deployed to a environment
	FindPods(context.Context, *model.JenkinsServerId) (*model4.Pods, error)
}

// UnimplementedJenkinsServerQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedJenkinsServerQueryControllerServer struct {
}

func (UnimplementedJenkinsServerQueryControllerServer) List(context.Context, *rpc.PageInfo) (*model.JenkinsServerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedJenkinsServerQueryControllerServer) GetById(context.Context, *model.JenkinsServerId) (*model.JenkinsServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedJenkinsServerQueryControllerServer) FindByProductId(context.Context, *model1.ProductId) (*model.JenkinsServers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedJenkinsServerQueryControllerServer) FindByEnvironmentId(context.Context, *model2.EnvironmentId) (*model.JenkinsServers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEnvironmentId not implemented")
}
func (UnimplementedJenkinsServerQueryControllerServer) FindByKubeClusterId(context.Context, *model3.KubeClusterId) (*model.JenkinsServers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKubeClusterId not implemented")
}
func (UnimplementedJenkinsServerQueryControllerServer) FindPods(context.Context, *model.JenkinsServerId) (*model4.Pods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPods not implemented")
}

// UnsafeJenkinsServerQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JenkinsServerQueryControllerServer will
// result in compilation errors.
type UnsafeJenkinsServerQueryControllerServer interface {
	mustEmbedUnimplementedJenkinsServerQueryControllerServer()
}

func RegisterJenkinsServerQueryControllerServer(s grpc.ServiceRegistrar, srv JenkinsServerQueryControllerServer) {
	s.RegisterService(&JenkinsServerQueryController_ServiceDesc, srv)
}

func _JenkinsServerQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsServerQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsServerQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsServerQueryControllerServer).List(ctx, req.(*rpc.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsServerQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.JenkinsServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsServerQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsServerQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsServerQueryControllerServer).GetById(ctx, req.(*model.JenkinsServerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsServerQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsServerQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsServerQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsServerQueryControllerServer).FindByProductId(ctx, req.(*model1.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsServerQueryController_FindByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.EnvironmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsServerQueryControllerServer).FindByEnvironmentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsServerQueryController_FindByEnvironmentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsServerQueryControllerServer).FindByEnvironmentId(ctx, req.(*model2.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsServerQueryController_FindByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model3.KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsServerQueryControllerServer).FindByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsServerQueryController_FindByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsServerQueryControllerServer).FindByKubeClusterId(ctx, req.(*model3.KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _JenkinsServerQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.JenkinsServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JenkinsServerQueryControllerServer).FindPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JenkinsServerQueryController_FindPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JenkinsServerQueryControllerServer).FindPods(ctx, req.(*model.JenkinsServerId))
	}
	return interceptor(ctx, in, info, handler)
}

// JenkinsServerQueryController_ServiceDesc is the grpc.ServiceDesc for JenkinsServerQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JenkinsServerQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.jenkinsserver.service.JenkinsServerQueryController",
	HandlerType: (*JenkinsServerQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _JenkinsServerQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _JenkinsServerQueryController_GetById_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _JenkinsServerQueryController_FindByProductId_Handler,
		},
		{
			MethodName: "findByEnvironmentId",
			Handler:    _JenkinsServerQueryController_FindByEnvironmentId_Handler,
		},
		{
			MethodName: "findByKubeClusterId",
			Handler:    _JenkinsServerQueryController_FindByKubeClusterId_Handler,
		},
		{
			MethodName: "findPods",
			Handler:    _JenkinsServerQueryController_FindPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/jenkinsserver/service/query.proto",
}
