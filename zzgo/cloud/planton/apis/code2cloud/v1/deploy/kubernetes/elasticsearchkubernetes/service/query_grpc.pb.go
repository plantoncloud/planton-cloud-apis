// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/kubernetes/elasticsearchkubernetes/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/kubernetes/elasticsearchkubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ElasticsearchKubernetesQueryController_Get_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesQueryController/get"
	ElasticsearchKubernetesQueryController_GetPassword_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesQueryController/getPassword"
)

// ElasticsearchKubernetesQueryControllerClient is the client API for ElasticsearchKubernetesQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElasticsearchKubernetesQueryControllerClient interface {
	// look up elasticsearch-kubernetes using elasticsearch-kubernetes id
	Get(ctx context.Context, in *model.ElasticsearchKubernetesId, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error)
	// look up elasticsearch-kubernetes sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *model.ElasticsearchKubernetesId, opts ...grpc.CallOption) (*model.ElasticsearchKubernetesPassword, error)
}

type elasticsearchKubernetesQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewElasticsearchKubernetesQueryControllerClient(cc grpc.ClientConnInterface) ElasticsearchKubernetesQueryControllerClient {
	return &elasticsearchKubernetesQueryControllerClient{cc}
}

func (c *elasticsearchKubernetesQueryControllerClient) Get(ctx context.Context, in *model.ElasticsearchKubernetesId, opts ...grpc.CallOption) (*model.ElasticsearchKubernetes, error) {
	out := new(model.ElasticsearchKubernetes)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchKubernetesQueryControllerClient) GetPassword(ctx context.Context, in *model.ElasticsearchKubernetesId, opts ...grpc.CallOption) (*model.ElasticsearchKubernetesPassword, error) {
	out := new(model.ElasticsearchKubernetesPassword)
	err := c.cc.Invoke(ctx, ElasticsearchKubernetesQueryController_GetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElasticsearchKubernetesQueryControllerServer is the server API for ElasticsearchKubernetesQueryController service.
// All implementations should embed UnimplementedElasticsearchKubernetesQueryControllerServer
// for forward compatibility
type ElasticsearchKubernetesQueryControllerServer interface {
	// look up elasticsearch-kubernetes using elasticsearch-kubernetes id
	Get(context.Context, *model.ElasticsearchKubernetesId) (*model.ElasticsearchKubernetes, error)
	// look up elasticsearch-kubernetes sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model.ElasticsearchKubernetesId) (*model.ElasticsearchKubernetesPassword, error)
}

// UnimplementedElasticsearchKubernetesQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedElasticsearchKubernetesQueryControllerServer struct {
}

func (UnimplementedElasticsearchKubernetesQueryControllerServer) Get(context.Context, *model.ElasticsearchKubernetesId) (*model.ElasticsearchKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedElasticsearchKubernetesQueryControllerServer) GetPassword(context.Context, *model.ElasticsearchKubernetesId) (*model.ElasticsearchKubernetesPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}

// UnsafeElasticsearchKubernetesQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElasticsearchKubernetesQueryControllerServer will
// result in compilation errors.
type UnsafeElasticsearchKubernetesQueryControllerServer interface {
	mustEmbedUnimplementedElasticsearchKubernetesQueryControllerServer()
}

func RegisterElasticsearchKubernetesQueryControllerServer(s grpc.ServiceRegistrar, srv ElasticsearchKubernetesQueryControllerServer) {
	s.RegisterService(&ElasticsearchKubernetesQueryController_ServiceDesc, srv)
}

func _ElasticsearchKubernetesQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ElasticsearchKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesQueryControllerServer).Get(ctx, req.(*model.ElasticsearchKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchKubernetesQueryController_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ElasticsearchKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchKubernetesQueryControllerServer).GetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchKubernetesQueryController_GetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchKubernetesQueryControllerServer).GetPassword(ctx, req.(*model.ElasticsearchKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

// ElasticsearchKubernetesQueryController_ServiceDesc is the grpc.ServiceDesc for ElasticsearchKubernetesQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ElasticsearchKubernetesQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.kubernetes.elasticsearchkubernetes.service.ElasticsearchKubernetesQueryController",
	HandlerType: (*ElasticsearchKubernetesQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _ElasticsearchKubernetesQueryController_Get_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _ElasticsearchKubernetesQueryController_GetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/kubernetes/elasticsearchkubernetes/service/query.proto",
}
