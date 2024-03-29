// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/rediscluster/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/rediscluster/model"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/rpc"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/apiresources/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RedisClusterQueryController_List_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.rediscluster.service.RedisClusterQueryController/list"
	RedisClusterQueryController_GetById_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.rediscluster.service.RedisClusterQueryController/getById"
	RedisClusterQueryController_GetPassword_FullMethodName = "/cloud.planton.apis.code2cloud.v1.rediscluster.service.RedisClusterQueryController/getPassword"
	RedisClusterQueryController_FindPods_FullMethodName    = "/cloud.planton.apis.code2cloud.v1.rediscluster.service.RedisClusterQueryController/findPods"
)

// RedisClusterQueryControllerClient is the client API for RedisClusterQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RedisClusterQueryControllerClient interface {
	// list all redis-clusters on planton cluster for the requested page.
	List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.RedisClusterList, error)
	// look up redis-cluster using redis-cluster id
	GetById(ctx context.Context, in *model.RedisClusterId, opts ...grpc.CallOption) (*model.RedisCluster, error)
	// look up redis-cluster sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *model.RedisClusterId, opts ...grpc.CallOption) (*model.RedisClusterPassword, error)
	// lookup pods of a redis-cluster deployed to a environment
	FindPods(ctx context.Context, in *model.RedisClusterId, opts ...grpc.CallOption) (*model1.Pods, error)
}

type redisClusterQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewRedisClusterQueryControllerClient(cc grpc.ClientConnInterface) RedisClusterQueryControllerClient {
	return &redisClusterQueryControllerClient{cc}
}

func (c *redisClusterQueryControllerClient) List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.RedisClusterList, error) {
	out := new(model.RedisClusterList)
	err := c.cc.Invoke(ctx, RedisClusterQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisClusterQueryControllerClient) GetById(ctx context.Context, in *model.RedisClusterId, opts ...grpc.CallOption) (*model.RedisCluster, error) {
	out := new(model.RedisCluster)
	err := c.cc.Invoke(ctx, RedisClusterQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisClusterQueryControllerClient) GetPassword(ctx context.Context, in *model.RedisClusterId, opts ...grpc.CallOption) (*model.RedisClusterPassword, error) {
	out := new(model.RedisClusterPassword)
	err := c.cc.Invoke(ctx, RedisClusterQueryController_GetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisClusterQueryControllerClient) FindPods(ctx context.Context, in *model.RedisClusterId, opts ...grpc.CallOption) (*model1.Pods, error) {
	out := new(model1.Pods)
	err := c.cc.Invoke(ctx, RedisClusterQueryController_FindPods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedisClusterQueryControllerServer is the server API for RedisClusterQueryController service.
// All implementations should embed UnimplementedRedisClusterQueryControllerServer
// for forward compatibility
type RedisClusterQueryControllerServer interface {
	// list all redis-clusters on planton cluster for the requested page.
	List(context.Context, *rpc.PageInfo) (*model.RedisClusterList, error)
	// look up redis-cluster using redis-cluster id
	GetById(context.Context, *model.RedisClusterId) (*model.RedisCluster, error)
	// look up redis-cluster sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model.RedisClusterId) (*model.RedisClusterPassword, error)
	// lookup pods of a redis-cluster deployed to a environment
	FindPods(context.Context, *model.RedisClusterId) (*model1.Pods, error)
}

// UnimplementedRedisClusterQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedRedisClusterQueryControllerServer struct {
}

func (UnimplementedRedisClusterQueryControllerServer) List(context.Context, *rpc.PageInfo) (*model.RedisClusterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRedisClusterQueryControllerServer) GetById(context.Context, *model.RedisClusterId) (*model.RedisCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedRedisClusterQueryControllerServer) GetPassword(context.Context, *model.RedisClusterId) (*model.RedisClusterPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}
func (UnimplementedRedisClusterQueryControllerServer) FindPods(context.Context, *model.RedisClusterId) (*model1.Pods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPods not implemented")
}

// UnsafeRedisClusterQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedisClusterQueryControllerServer will
// result in compilation errors.
type UnsafeRedisClusterQueryControllerServer interface {
	mustEmbedUnimplementedRedisClusterQueryControllerServer()
}

func RegisterRedisClusterQueryControllerServer(s grpc.ServiceRegistrar, srv RedisClusterQueryControllerServer) {
	s.RegisterService(&RedisClusterQueryController_ServiceDesc, srv)
}

func _RedisClusterQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisClusterQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisClusterQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisClusterQueryControllerServer).List(ctx, req.(*rpc.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisClusterQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.RedisClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisClusterQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisClusterQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisClusterQueryControllerServer).GetById(ctx, req.(*model.RedisClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisClusterQueryController_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.RedisClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisClusterQueryControllerServer).GetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisClusterQueryController_GetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisClusterQueryControllerServer).GetPassword(ctx, req.(*model.RedisClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisClusterQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.RedisClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisClusterQueryControllerServer).FindPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisClusterQueryController_FindPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisClusterQueryControllerServer).FindPods(ctx, req.(*model.RedisClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

// RedisClusterQueryController_ServiceDesc is the grpc.ServiceDesc for RedisClusterQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedisClusterQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.rediscluster.service.RedisClusterQueryController",
	HandlerType: (*RedisClusterQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _RedisClusterQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _RedisClusterQueryController_GetById_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _RedisClusterQueryController_GetPassword_Handler,
		},
		{
			MethodName: "findPods",
			Handler:    _RedisClusterQueryController_FindPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/rediscluster/service/query.proto",
}
