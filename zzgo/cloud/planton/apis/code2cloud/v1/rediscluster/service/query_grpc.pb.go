// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/rediscluster/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/rediscluster/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RedisClusterQueryController_GetById_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.rediscluster.service.RedisClusterQueryController/getById"
	RedisClusterQueryController_GetPassword_FullMethodName = "/cloud.planton.apis.code2cloud.v1.rediscluster.service.RedisClusterQueryController/getPassword"
)

// RedisClusterQueryControllerClient is the client API for RedisClusterQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RedisClusterQueryControllerClient interface {
	// look up redis-cluster using redis-cluster id
	GetById(ctx context.Context, in *model.RedisClusterId, opts ...grpc.CallOption) (*model.RedisCluster, error)
	// look up redis-cluster sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *model.RedisClusterId, opts ...grpc.CallOption) (*model.RedisClusterPassword, error)
}

type redisClusterQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewRedisClusterQueryControllerClient(cc grpc.ClientConnInterface) RedisClusterQueryControllerClient {
	return &redisClusterQueryControllerClient{cc}
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

// RedisClusterQueryControllerServer is the server API for RedisClusterQueryController service.
// All implementations should embed UnimplementedRedisClusterQueryControllerServer
// for forward compatibility
type RedisClusterQueryControllerServer interface {
	// look up redis-cluster using redis-cluster id
	GetById(context.Context, *model.RedisClusterId) (*model.RedisCluster, error)
	// look up redis-cluster sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model.RedisClusterId) (*model.RedisClusterPassword, error)
}

// UnimplementedRedisClusterQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedRedisClusterQueryControllerServer struct {
}

func (UnimplementedRedisClusterQueryControllerServer) GetById(context.Context, *model.RedisClusterId) (*model.RedisCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedRedisClusterQueryControllerServer) GetPassword(context.Context, *model.RedisClusterId) (*model.RedisClusterPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
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

// RedisClusterQueryController_ServiceDesc is the grpc.ServiceDesc for RedisClusterQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedisClusterQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.rediscluster.service.RedisClusterQueryController",
	HandlerType: (*RedisClusterQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _RedisClusterQueryController_GetById_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _RedisClusterQueryController_GetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/rediscluster/service/query.proto",
}
