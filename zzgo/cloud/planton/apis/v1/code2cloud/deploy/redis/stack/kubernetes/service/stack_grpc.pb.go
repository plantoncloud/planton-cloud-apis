// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/redis/stack/kubernetes/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/redis/stack/kubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RedisClusterKubernetesStackController_Execute_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.service.RedisClusterKubernetesStackController/execute"
	RedisClusterKubernetesStackController_GetStackOutputs_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.service.RedisClusterKubernetesStackController/getStackOutputs"
)

// RedisClusterKubernetesStackControllerClient is the client API for RedisClusterKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RedisClusterKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.RedisClusterKubernetesStackInput, opts ...grpc.CallOption) (RedisClusterKubernetesStackController_ExecuteClient, error)
	GetStackOutputs(ctx context.Context, in *model.RedisClusterKubernetesStackInput, opts ...grpc.CallOption) (*model.RedisClusterKubernetesStackOutputs, error)
}

type redisClusterKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewRedisClusterKubernetesStackControllerClient(cc grpc.ClientConnInterface) RedisClusterKubernetesStackControllerClient {
	return &redisClusterKubernetesStackControllerClient{cc}
}

func (c *redisClusterKubernetesStackControllerClient) Execute(ctx context.Context, in *model.RedisClusterKubernetesStackInput, opts ...grpc.CallOption) (RedisClusterKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &RedisClusterKubernetesStackController_ServiceDesc.Streams[0], RedisClusterKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &redisClusterKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RedisClusterKubernetesStackController_ExecuteClient interface {
	Recv() (*model.RedisClusterKubernetesStackResponse, error)
	grpc.ClientStream
}

type redisClusterKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *redisClusterKubernetesStackControllerExecuteClient) Recv() (*model.RedisClusterKubernetesStackResponse, error) {
	m := new(model.RedisClusterKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *redisClusterKubernetesStackControllerClient) GetStackOutputs(ctx context.Context, in *model.RedisClusterKubernetesStackInput, opts ...grpc.CallOption) (*model.RedisClusterKubernetesStackOutputs, error) {
	out := new(model.RedisClusterKubernetesStackOutputs)
	err := c.cc.Invoke(ctx, RedisClusterKubernetesStackController_GetStackOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedisClusterKubernetesStackControllerServer is the server API for RedisClusterKubernetesStackController service.
// All implementations should embed UnimplementedRedisClusterKubernetesStackControllerServer
// for forward compatibility
type RedisClusterKubernetesStackControllerServer interface {
	Execute(*model.RedisClusterKubernetesStackInput, RedisClusterKubernetesStackController_ExecuteServer) error
	GetStackOutputs(context.Context, *model.RedisClusterKubernetesStackInput) (*model.RedisClusterKubernetesStackOutputs, error)
}

// UnimplementedRedisClusterKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedRedisClusterKubernetesStackControllerServer struct {
}

func (UnimplementedRedisClusterKubernetesStackControllerServer) Execute(*model.RedisClusterKubernetesStackInput, RedisClusterKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedRedisClusterKubernetesStackControllerServer) GetStackOutputs(context.Context, *model.RedisClusterKubernetesStackInput) (*model.RedisClusterKubernetesStackOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackOutputs not implemented")
}

// UnsafeRedisClusterKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedisClusterKubernetesStackControllerServer will
// result in compilation errors.
type UnsafeRedisClusterKubernetesStackControllerServer interface {
	mustEmbedUnimplementedRedisClusterKubernetesStackControllerServer()
}

func RegisterRedisClusterKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv RedisClusterKubernetesStackControllerServer) {
	s.RegisterService(&RedisClusterKubernetesStackController_ServiceDesc, srv)
}

func _RedisClusterKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.RedisClusterKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RedisClusterKubernetesStackControllerServer).Execute(m, &redisClusterKubernetesStackControllerExecuteServer{stream})
}

type RedisClusterKubernetesStackController_ExecuteServer interface {
	Send(*model.RedisClusterKubernetesStackResponse) error
	grpc.ServerStream
}

type redisClusterKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *redisClusterKubernetesStackControllerExecuteServer) Send(m *model.RedisClusterKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RedisClusterKubernetesStackController_GetStackOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.RedisClusterKubernetesStackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisClusterKubernetesStackControllerServer).GetStackOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisClusterKubernetesStackController_GetStackOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisClusterKubernetesStackControllerServer).GetStackOutputs(ctx, req.(*model.RedisClusterKubernetesStackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// RedisClusterKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for RedisClusterKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedisClusterKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.service.RedisClusterKubernetesStackController",
	HandlerType: (*RedisClusterKubernetesStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStackOutputs",
			Handler:    _RedisClusterKubernetesStackController_GetStackOutputs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _RedisClusterKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/redis/stack/kubernetes/service/stack.proto",
}
