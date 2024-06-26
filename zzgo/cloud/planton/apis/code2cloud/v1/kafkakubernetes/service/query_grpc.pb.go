// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/kafkakubernetes/service/query.proto

package service

import (
	context "context"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/environment/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kafkakubernetes/model"
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
	KafkaKubernetesQueryController_GetById_FullMethodName             = "/cloud.planton.apis.code2cloud.v1.kafkakubernetes.service.KafkaKubernetesQueryController/getById"
	KafkaKubernetesQueryController_FindByProductId_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.kafkakubernetes.service.KafkaKubernetesQueryController/findByProductId"
	KafkaKubernetesQueryController_FindByEnvironmentId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kafkakubernetes.service.KafkaKubernetesQueryController/findByEnvironmentId"
	KafkaKubernetesQueryController_FindByKubeClusterId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kafkakubernetes.service.KafkaKubernetesQueryController/findByKubeClusterId"
	KafkaKubernetesQueryController_GetPassword_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.kafkakubernetes.service.KafkaKubernetesQueryController/getPassword"
)

// KafkaKubernetesQueryControllerClient is the client API for KafkaKubernetesQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KafkaKubernetesQueryControllerClient interface {
	// look up kafka-kubernetes using kafka-kubernetes id
	GetById(ctx context.Context, in *model.KafkaKubernetesId, opts ...grpc.CallOption) (*model.KafkaKubernetes, error)
	// find kafka-kubernetess by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.KafkaKubernetesPage, error)
	// find kafka-kubernetess by environment
	FindByEnvironmentId(ctx context.Context, in *model2.EnvironmentId, opts ...grpc.CallOption) (*model.KafkaKubernetesPage, error)
	// find kafka-kubernetess by kube-cluster
	FindByKubeClusterId(ctx context.Context, in *model3.KubeClusterId, opts ...grpc.CallOption) (*model.KafkaKubernetesPage, error)
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

func (c *kafkaKubernetesQueryControllerClient) GetById(ctx context.Context, in *model.KafkaKubernetesId, opts ...grpc.CallOption) (*model.KafkaKubernetes, error) {
	out := new(model.KafkaKubernetes)
	err := c.cc.Invoke(ctx, KafkaKubernetesQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesQueryControllerClient) FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.KafkaKubernetesPage, error) {
	out := new(model.KafkaKubernetesPage)
	err := c.cc.Invoke(ctx, KafkaKubernetesQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesQueryControllerClient) FindByEnvironmentId(ctx context.Context, in *model2.EnvironmentId, opts ...grpc.CallOption) (*model.KafkaKubernetesPage, error) {
	out := new(model.KafkaKubernetesPage)
	err := c.cc.Invoke(ctx, KafkaKubernetesQueryController_FindByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaKubernetesQueryControllerClient) FindByKubeClusterId(ctx context.Context, in *model3.KubeClusterId, opts ...grpc.CallOption) (*model.KafkaKubernetesPage, error) {
	out := new(model.KafkaKubernetesPage)
	err := c.cc.Invoke(ctx, KafkaKubernetesQueryController_FindByKubeClusterId_FullMethodName, in, out, opts...)
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
	GetById(context.Context, *model.KafkaKubernetesId) (*model.KafkaKubernetes, error)
	// find kafka-kubernetess by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(context.Context, *model1.ProductId) (*model.KafkaKubernetesPage, error)
	// find kafka-kubernetess by environment
	FindByEnvironmentId(context.Context, *model2.EnvironmentId) (*model.KafkaKubernetesPage, error)
	// find kafka-kubernetess by kube-cluster
	FindByKubeClusterId(context.Context, *model3.KubeClusterId) (*model.KafkaKubernetesPage, error)
	// look up kafka-kubernetes sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model.KafkaKubernetesId) (*model.KafkaKubernetesPassword, error)
}

// UnimplementedKafkaKubernetesQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKafkaKubernetesQueryControllerServer struct {
}

func (UnimplementedKafkaKubernetesQueryControllerServer) GetById(context.Context, *model.KafkaKubernetesId) (*model.KafkaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedKafkaKubernetesQueryControllerServer) FindByProductId(context.Context, *model1.ProductId) (*model.KafkaKubernetesPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedKafkaKubernetesQueryControllerServer) FindByEnvironmentId(context.Context, *model2.EnvironmentId) (*model.KafkaKubernetesPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEnvironmentId not implemented")
}
func (UnimplementedKafkaKubernetesQueryControllerServer) FindByKubeClusterId(context.Context, *model3.KubeClusterId) (*model.KafkaKubernetesPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKubeClusterId not implemented")
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

func _KafkaKubernetesQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesQueryControllerServer).GetById(ctx, req.(*model.KafkaKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesQueryControllerServer).FindByProductId(ctx, req.(*model1.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesQueryController_FindByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.EnvironmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesQueryControllerServer).FindByEnvironmentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesQueryController_FindByEnvironmentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesQueryControllerServer).FindByEnvironmentId(ctx, req.(*model2.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaKubernetesQueryController_FindByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model3.KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaKubernetesQueryControllerServer).FindByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaKubernetesQueryController_FindByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaKubernetesQueryControllerServer).FindByKubeClusterId(ctx, req.(*model3.KubeClusterId))
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
	ServiceName: "cloud.planton.apis.code2cloud.v1.kafkakubernetes.service.KafkaKubernetesQueryController",
	HandlerType: (*KafkaKubernetesQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _KafkaKubernetesQueryController_GetById_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _KafkaKubernetesQueryController_FindByProductId_Handler,
		},
		{
			MethodName: "findByEnvironmentId",
			Handler:    _KafkaKubernetesQueryController_FindByEnvironmentId_Handler,
		},
		{
			MethodName: "findByKubeClusterId",
			Handler:    _KafkaKubernetesQueryController_FindByKubeClusterId_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _KafkaKubernetesQueryController_GetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/kafkakubernetes/service/query.proto",
}

const (
	KafkaTopicQueryController_FindByKafkaKubernetesId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kafkakubernetes.service.KafkaTopicQueryController/findByKafkaKubernetesId"
	KafkaTopicQueryController_GetById_FullMethodName                 = "/cloud.planton.apis.code2cloud.v1.kafkakubernetes.service.KafkaTopicQueryController/getById"
)

// KafkaTopicQueryControllerClient is the client API for KafkaTopicQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KafkaTopicQueryControllerClient interface {
	// find kafka topics by kafka-kubernetes id
	FindByKafkaKubernetesId(ctx context.Context, in *model.KafkaKubernetesId, opts ...grpc.CallOption) (*model.KafkaTopicList, error)
	// look up kafka topic using kafka topic id
	GetById(ctx context.Context, in *model.KafkaTopicQueryInput, opts ...grpc.CallOption) (*model.KafkaTopic, error)
}

type kafkaTopicQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKafkaTopicQueryControllerClient(cc grpc.ClientConnInterface) KafkaTopicQueryControllerClient {
	return &kafkaTopicQueryControllerClient{cc}
}

func (c *kafkaTopicQueryControllerClient) FindByKafkaKubernetesId(ctx context.Context, in *model.KafkaKubernetesId, opts ...grpc.CallOption) (*model.KafkaTopicList, error) {
	out := new(model.KafkaTopicList)
	err := c.cc.Invoke(ctx, KafkaTopicQueryController_FindByKafkaKubernetesId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaTopicQueryControllerClient) GetById(ctx context.Context, in *model.KafkaTopicQueryInput, opts ...grpc.CallOption) (*model.KafkaTopic, error) {
	out := new(model.KafkaTopic)
	err := c.cc.Invoke(ctx, KafkaTopicQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafkaTopicQueryControllerServer is the server API for KafkaTopicQueryController service.
// All implementations should embed UnimplementedKafkaTopicQueryControllerServer
// for forward compatibility
type KafkaTopicQueryControllerServer interface {
	// find kafka topics by kafka-kubernetes id
	FindByKafkaKubernetesId(context.Context, *model.KafkaKubernetesId) (*model.KafkaTopicList, error)
	// look up kafka topic using kafka topic id
	GetById(context.Context, *model.KafkaTopicQueryInput) (*model.KafkaTopic, error)
}

// UnimplementedKafkaTopicQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKafkaTopicQueryControllerServer struct {
}

func (UnimplementedKafkaTopicQueryControllerServer) FindByKafkaKubernetesId(context.Context, *model.KafkaKubernetesId) (*model.KafkaTopicList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKafkaKubernetesId not implemented")
}
func (UnimplementedKafkaTopicQueryControllerServer) GetById(context.Context, *model.KafkaTopicQueryInput) (*model.KafkaTopic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
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

func _KafkaTopicQueryController_FindByKafkaKubernetesId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaTopicQueryControllerServer).FindByKafkaKubernetesId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaTopicQueryController_FindByKafkaKubernetesId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaTopicQueryControllerServer).FindByKafkaKubernetesId(ctx, req.(*model.KafkaKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaTopicQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KafkaTopicQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaTopicQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaTopicQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaTopicQueryControllerServer).GetById(ctx, req.(*model.KafkaTopicQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

// KafkaTopicQueryController_ServiceDesc is the grpc.ServiceDesc for KafkaTopicQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KafkaTopicQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.kafkakubernetes.service.KafkaTopicQueryController",
	HandlerType: (*KafkaTopicQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findByKafkaKubernetesId",
			Handler:    _KafkaTopicQueryController_FindByKafkaKubernetesId_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _KafkaTopicQueryController_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/kafkakubernetes/service/query.proto",
}
