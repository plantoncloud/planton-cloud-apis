// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/kubernetes/rediskubernetes/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/rediskubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RedisKubernetesStackController_Execute_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesStackController/execute"
)

// RedisKubernetesStackControllerClient is the client API for RedisKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RedisKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.RedisKubernetesStackInput, opts ...grpc.CallOption) (RedisKubernetesStackController_ExecuteClient, error)
}

type redisKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewRedisKubernetesStackControllerClient(cc grpc.ClientConnInterface) RedisKubernetesStackControllerClient {
	return &redisKubernetesStackControllerClient{cc}
}

func (c *redisKubernetesStackControllerClient) Execute(ctx context.Context, in *model.RedisKubernetesStackInput, opts ...grpc.CallOption) (RedisKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &RedisKubernetesStackController_ServiceDesc.Streams[0], RedisKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &redisKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RedisKubernetesStackController_ExecuteClient interface {
	Recv() (*model.RedisKubernetesStackResponse, error)
	grpc.ClientStream
}

type redisKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *redisKubernetesStackControllerExecuteClient) Recv() (*model.RedisKubernetesStackResponse, error) {
	m := new(model.RedisKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RedisKubernetesStackControllerServer is the server API for RedisKubernetesStackController service.
// All implementations should embed UnimplementedRedisKubernetesStackControllerServer
// for forward compatibility
type RedisKubernetesStackControllerServer interface {
	Execute(*model.RedisKubernetesStackInput, RedisKubernetesStackController_ExecuteServer) error
}

// UnimplementedRedisKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedRedisKubernetesStackControllerServer struct {
}

func (UnimplementedRedisKubernetesStackControllerServer) Execute(*model.RedisKubernetesStackInput, RedisKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeRedisKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedisKubernetesStackControllerServer will
// result in compilation errors.
type UnsafeRedisKubernetesStackControllerServer interface {
	mustEmbedUnimplementedRedisKubernetesStackControllerServer()
}

func RegisterRedisKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv RedisKubernetesStackControllerServer) {
	s.RegisterService(&RedisKubernetesStackController_ServiceDesc, srv)
}

func _RedisKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.RedisKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RedisKubernetesStackControllerServer).Execute(m, &redisKubernetesStackControllerExecuteServer{stream})
}

type RedisKubernetesStackController_ExecuteServer interface {
	Send(*model.RedisKubernetesStackResponse) error
	grpc.ServerStream
}

type redisKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *redisKubernetesStackControllerExecuteServer) Send(m *model.RedisKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

// RedisKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for RedisKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedisKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.service.RedisKubernetesStackController",
	HandlerType: (*RedisKubernetesStackControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _RedisKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/kubernetes/rediskubernetes/service/stack.proto",
}
