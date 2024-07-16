// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/kubernetes/kafkakubernetes/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/kubernetes/kafkakubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KafkaKubernetesQueryController_Get_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesQueryController/get"
	KafkaKubernetesQueryController_GetPassword_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesQueryController/getPassword"
)

// KafkaKubernetesQueryControllerClient is the client API for KafkaKubernetesQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KafkaKubernetesQueryControllerClient interface {
	// look up kafka-kubernetes using kafka-kubernetes id
	Get(ctx context.Context, in *model.KafkaKubernetesId, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// look up kafka-kubernetes sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *model.KafkaKubernetesId, opts ...grpc.CallOption) (*model.KafkaKubernetesPassword, error)
}

type kafkaKubernetesQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKafkaKubernetesQueryControllerClient(cc grpc.ClientConnInterface) KafkaKubernetesQueryControllerClient {
	return &kafkaKubernetesQueryControllerClient{cc}
}

func (c *kafkaKubernetesQueryControllerClient) Get(ctx context.Context, in *model.KafkaKubernetesId, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesQueryControllerClient) GetPassword(ctx context.Context, in *model.KafkaKubernetesId, opts ...grpc.CallOption) (*model.KafkaKubernetesPassword, error) {
	out := new(model.KafkaKubernetesPassword)
	err := c.cc.Invoke(ctx, KafkaKubernetesQueryController_GetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafkaKubernetesQueryControllerServer is the server API for KafkaKubernetesQueryController service.
// All implementations should embed UnimplementedKafkaKubernetesQueryControllerServer
// for forward compatibility
type KafkaKubernetesQueryControllerServer interface {
	// look up kafka-kubernetes using kafka-kubernetes id
	Get(context.Context, *model.KafkaKubernetesId) (*model.KafkaKubernetes, error)
	// look up kafka-kubernetes sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model.KafkaKubernetesId) (*model.KafkaKubernetesPassword, error)
}

// UnimplementedKafkaKubernetesQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKafkaKubernetesQueryControllerServer struct {
}

func (UnimplementedKafkaKubernetesQueryControllerServer) Get(context.Context, *model.KafkaKubernetesId) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKafkaKubernetesQueryControllerServer) GetPassword(context.Context, *model.KafkaKubernetesId) (*model.KafkaKubernetesPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}

// UnsafeKafkaKubernetesQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KafkaKubernetesQueryControllerServer will
// result in compilation errors.
type UnsafeKafkaKubernetesQueryControllerServer interface {
	mustEmbedUnimplementedKafkaKubernetesQueryControllerServer()
}

func RegisterKafkaKubernetesQueryControllerServer(s grpc.ServiceRegistrar, srv KafkaKubernetesQueryControllerServer) {
	s.RegisterService(&KafkaKubernetesQueryController_ServiceDesc, srv)
}

func _KafkaKubernetesQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesQueryControllerServer).Get(ctx, req.(*model.KafkaKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesQueryController_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesQueryControllerServer).GetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesQueryController_GetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesQueryControllerServer).GetPassword(ctx, req.(*model.KafkaKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

// KafkaKubernetesQueryController_ServiceDesc is the grpc.ServiceDesc for KafkaKubernetesQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KafkaKubernetesQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaKubernetesQueryController",
	HandlerType: (*KafkaKubernetesQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _KafkaKubernetesQueryController_Get_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _KafkaKubernetesQueryController_GetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/kubernetes/kafkakubernetes/service/query.proto",
}

const (
	KafkaTopicQueryController_Get_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaTopicQueryController/get"
)

// KafkaTopicQueryControllerClient is the client API for KafkaTopicQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KafkaTopicQueryControllerClient interface {
	// look up kafka topic using kafka topic id
	Get(ctx context.Context, in *model.KafkaTopicQueryInput, opts ...grpc.CallOption) (*model.KafkaTopic, error)
}

type kafkaTopicQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKafkaTopicQueryControllerClient(cc grpc.ClientConnInterface) KafkaTopicQueryControllerClient {
	return &kafkaTopicQueryControllerClient{cc}
}

func (c *kafkaTopicQueryControllerClient) Get(ctx context.Context, in *model.KafkaTopicQueryInput, opts ...grpc.CallOption) (*model.KafkaTopic, error) {
	out := new(model.KafkaTopic)
	err := c.cc.Invoke(ctx, KafkaTopicQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafkaTopicQueryControllerServer is the server API for KafkaTopicQueryController service.
// All implementations should embed UnimplementedKafkaTopicQueryControllerServer
// for forward compatibility
type KafkaTopicQueryControllerServer interface {
	// look up kafka topic using kafka topic id
	Get(context.Context, *model.KafkaTopicQueryInput) (*model.KafkaTopic, error)
}

// UnimplementedKafkaTopicQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKafkaTopicQueryControllerServer struct {
}

func (UnimplementedKafkaTopicQueryControllerServer) Get(context.Context, *model.KafkaTopicQueryInput) (*model.KafkaTopic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeKafkaTopicQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KafkaTopicQueryControllerServer will
// result in compilation errors.
type UnsafeKafkaTopicQueryControllerServer interface {
	mustEmbedUnimplementedKafkaTopicQueryControllerServer()
}

func RegisterKafkaTopicQueryControllerServer(s grpc.ServiceRegistrar, srv KafkaTopicQueryControllerServer) {
	s.RegisterService(&KafkaTopicQueryController_ServiceDesc, srv)
}

func _KafkaTopicQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaTopicQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaTopicQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaTopicQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaTopicQueryControllerServer).Get(ctx, req.(*model.KafkaTopicQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

// KafkaTopicQueryController_ServiceDesc is the grpc.ServiceDesc for KafkaTopicQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KafkaTopicQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.service.KafkaTopicQueryController",
	HandlerType: (*KafkaTopicQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _KafkaTopicQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/kubernetes/kafkakubernetes/service/query.proto",
}
