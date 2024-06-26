// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/elasticsearchcluster/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/elasticsearchcluster/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ElasticsearchClusterQueryController_GetById_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.elasticsearchcluster.service.ElasticsearchClusterQueryController/getById"
	ElasticsearchClusterQueryController_GetPassword_FullMethodName = "/cloud.planton.apis.code2cloud.v1.elasticsearchcluster.service.ElasticsearchClusterQueryController/getPassword"
)

// ElasticsearchClusterQueryControllerClient is the client API for ElasticsearchClusterQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElasticsearchClusterQueryControllerClient interface {
	// look up elasticsearch-cluster using elasticsearch-cluster id
	GetById(ctx context.Context, in *model.ElasticsearchClusterId, opts ...grpc.CallOption) (*model.ElasticsearchCluster, error)
	// look up elasticsearch-cluster sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *model.ElasticsearchClusterId, opts ...grpc.CallOption) (*model.ElasticsearchClusterPassword, error)
}

type elasticsearchClusterQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewElasticsearchClusterQueryControllerClient(cc grpc.ClientConnInterface) ElasticsearchClusterQueryControllerClient {
	return &elasticsearchClusterQueryControllerClient{cc}
}

func (c *elasticsearchClusterQueryControllerClient) GetById(ctx context.Context, in *model.ElasticsearchClusterId, opts ...grpc.CallOption) (*model.ElasticsearchCluster, error) {
	out := new(model.ElasticsearchCluster)
	err := c.cc.Invoke(ctx, ElasticsearchClusterQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchClusterQueryControllerClient) GetPassword(ctx context.Context, in *model.ElasticsearchClusterId, opts ...grpc.CallOption) (*model.ElasticsearchClusterPassword, error) {
	out := new(model.ElasticsearchClusterPassword)
	err := c.cc.Invoke(ctx, ElasticsearchClusterQueryController_GetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElasticsearchClusterQueryControllerServer is the server API for ElasticsearchClusterQueryController service.
// All implementations should embed UnimplementedElasticsearchClusterQueryControllerServer
// for forward compatibility
type ElasticsearchClusterQueryControllerServer interface {
	// look up elasticsearch-cluster using elasticsearch-cluster id
	GetById(context.Context, *model.ElasticsearchClusterId) (*model.ElasticsearchCluster, error)
	// look up elasticsearch-cluster sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model.ElasticsearchClusterId) (*model.ElasticsearchClusterPassword, error)
}

// UnimplementedElasticsearchClusterQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedElasticsearchClusterQueryControllerServer struct {
}

func (UnimplementedElasticsearchClusterQueryControllerServer) GetById(context.Context, *model.ElasticsearchClusterId) (*model.ElasticsearchCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedElasticsearchClusterQueryControllerServer) GetPassword(context.Context, *model.ElasticsearchClusterId) (*model.ElasticsearchClusterPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}

// UnsafeElasticsearchClusterQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElasticsearchClusterQueryControllerServer will
// result in compilation errors.
type UnsafeElasticsearchClusterQueryControllerServer interface {
	mustEmbedUnimplementedElasticsearchClusterQueryControllerServer()
}

func RegisterElasticsearchClusterQueryControllerServer(s grpc.ServiceRegistrar, srv ElasticsearchClusterQueryControllerServer) {
	s.RegisterService(&ElasticsearchClusterQueryController_ServiceDesc, srv)
}

func _ElasticsearchClusterQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ElasticsearchClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchClusterQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchClusterQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchClusterQueryControllerServer).GetById(ctx, req.(*model.ElasticsearchClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchClusterQueryController_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ElasticsearchClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchClusterQueryControllerServer).GetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchClusterQueryController_GetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchClusterQueryControllerServer).GetPassword(ctx, req.(*model.ElasticsearchClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

// ElasticsearchClusterQueryController_ServiceDesc is the grpc.ServiceDesc for ElasticsearchClusterQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ElasticsearchClusterQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.elasticsearchcluster.service.ElasticsearchClusterQueryController",
	HandlerType: (*ElasticsearchClusterQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _ElasticsearchClusterQueryController_GetById_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _ElasticsearchClusterQueryController_GetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/elasticsearchcluster/service/query.proto",
}
