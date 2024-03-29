// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/elasticsearchcluster/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/elasticsearchcluster/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/environment/model"
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
	ElasticsearchClusterQueryController_List_FullMethodName                = "/cloud.planton.apis.code2cloud.v1.elasticsearchcluster.service.ElasticsearchClusterQueryController/list"
	ElasticsearchClusterQueryController_GetById_FullMethodName             = "/cloud.planton.apis.code2cloud.v1.elasticsearchcluster.service.ElasticsearchClusterQueryController/getById"
	ElasticsearchClusterQueryController_FindByProductId_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.elasticsearchcluster.service.ElasticsearchClusterQueryController/findByProductId"
	ElasticsearchClusterQueryController_FindByEnvironmentId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.elasticsearchcluster.service.ElasticsearchClusterQueryController/findByEnvironmentId"
	ElasticsearchClusterQueryController_FindByKubeClusterId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.elasticsearchcluster.service.ElasticsearchClusterQueryController/findByKubeClusterId"
	ElasticsearchClusterQueryController_GetPassword_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.elasticsearchcluster.service.ElasticsearchClusterQueryController/getPassword"
	ElasticsearchClusterQueryController_FindPods_FullMethodName            = "/cloud.planton.apis.code2cloud.v1.elasticsearchcluster.service.ElasticsearchClusterQueryController/findPods"
)

// ElasticsearchClusterQueryControllerClient is the client API for ElasticsearchClusterQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElasticsearchClusterQueryControllerClient interface {
	// list all elasticsearch-clusters on planton cluster for the requested page.
	List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.ElasticsearchClusterList, error)
	// look up elasticsearch-cluster using elasticsearch-cluster id
	GetById(ctx context.Context, in *model.ElasticsearchClusterId, opts ...grpc.CallOption) (*model.ElasticsearchCluster, error)
	// find elasticsearch-clusters by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.ElasticsearchClusters, error)
	// find elasticsearch-clusters by environment
	FindByEnvironmentId(ctx context.Context, in *model2.EnvironmentId, opts ...grpc.CallOption) (*model.ElasticsearchClusters, error)
	FindByKubeClusterId(ctx context.Context, in *model3.KubeClusterId, opts ...grpc.CallOption) (*model.ElasticsearchClusters, error)
	// look up elasticsearch-cluster sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *model.ElasticsearchClusterId, opts ...grpc.CallOption) (*model.ElasticsearchClusterPassword, error)
	// lookup pods of a elasticsearch-cluster deployed to a environment
	FindPods(ctx context.Context, in *model.ElasticsearchClusterId, opts ...grpc.CallOption) (*model4.Pods, error)
}

type elasticsearchClusterQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewElasticsearchClusterQueryControllerClient(cc grpc.ClientConnInterface) ElasticsearchClusterQueryControllerClient {
	return &elasticsearchClusterQueryControllerClient{cc}
}

func (c *elasticsearchClusterQueryControllerClient) List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.ElasticsearchClusterList, error) {
	out := new(model.ElasticsearchClusterList)
	err := c.cc.Invoke(ctx, ElasticsearchClusterQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchClusterQueryControllerClient) GetById(ctx context.Context, in *model.ElasticsearchClusterId, opts ...grpc.CallOption) (*model.ElasticsearchCluster, error) {
	out := new(model.ElasticsearchCluster)
	err := c.cc.Invoke(ctx, ElasticsearchClusterQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchClusterQueryControllerClient) FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.ElasticsearchClusters, error) {
	out := new(model.ElasticsearchClusters)
	err := c.cc.Invoke(ctx, ElasticsearchClusterQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchClusterQueryControllerClient) FindByEnvironmentId(ctx context.Context, in *model2.EnvironmentId, opts ...grpc.CallOption) (*model.ElasticsearchClusters, error) {
	out := new(model.ElasticsearchClusters)
	err := c.cc.Invoke(ctx, ElasticsearchClusterQueryController_FindByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticsearchClusterQueryControllerClient) FindByKubeClusterId(ctx context.Context, in *model3.KubeClusterId, opts ...grpc.CallOption) (*model.ElasticsearchClusters, error) {
	out := new(model.ElasticsearchClusters)
	err := c.cc.Invoke(ctx, ElasticsearchClusterQueryController_FindByKubeClusterId_FullMethodName, in, out, opts...)
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

func (c *elasticsearchClusterQueryControllerClient) FindPods(ctx context.Context, in *model.ElasticsearchClusterId, opts ...grpc.CallOption) (*model4.Pods, error) {
	out := new(model4.Pods)
	err := c.cc.Invoke(ctx, ElasticsearchClusterQueryController_FindPods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElasticsearchClusterQueryControllerServer is the server API for ElasticsearchClusterQueryController service.
// All implementations should embed UnimplementedElasticsearchClusterQueryControllerServer
// for forward compatibility
type ElasticsearchClusterQueryControllerServer interface {
	// list all elasticsearch-clusters on planton cluster for the requested page.
	List(context.Context, *rpc.PageInfo) (*model.ElasticsearchClusterList, error)
	// look up elasticsearch-cluster using elasticsearch-cluster id
	GetById(context.Context, *model.ElasticsearchClusterId) (*model.ElasticsearchCluster, error)
	// find elasticsearch-clusters by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(context.Context, *model1.ProductId) (*model.ElasticsearchClusters, error)
	// find elasticsearch-clusters by environment
	FindByEnvironmentId(context.Context, *model2.EnvironmentId) (*model.ElasticsearchClusters, error)
	FindByKubeClusterId(context.Context, *model3.KubeClusterId) (*model.ElasticsearchClusters, error)
	// look up elasticsearch-cluster sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model.ElasticsearchClusterId) (*model.ElasticsearchClusterPassword, error)
	// lookup pods of a elasticsearch-cluster deployed to a environment
	FindPods(context.Context, *model.ElasticsearchClusterId) (*model4.Pods, error)
}

// UnimplementedElasticsearchClusterQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedElasticsearchClusterQueryControllerServer struct {
}

func (UnimplementedElasticsearchClusterQueryControllerServer) List(context.Context, *rpc.PageInfo) (*model.ElasticsearchClusterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedElasticsearchClusterQueryControllerServer) GetById(context.Context, *model.ElasticsearchClusterId) (*model.ElasticsearchCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedElasticsearchClusterQueryControllerServer) FindByProductId(context.Context, *model1.ProductId) (*model.ElasticsearchClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedElasticsearchClusterQueryControllerServer) FindByEnvironmentId(context.Context, *model2.EnvironmentId) (*model.ElasticsearchClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEnvironmentId not implemented")
}
func (UnimplementedElasticsearchClusterQueryControllerServer) FindByKubeClusterId(context.Context, *model3.KubeClusterId) (*model.ElasticsearchClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKubeClusterId not implemented")
}
func (UnimplementedElasticsearchClusterQueryControllerServer) GetPassword(context.Context, *model.ElasticsearchClusterId) (*model.ElasticsearchClusterPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}
func (UnimplementedElasticsearchClusterQueryControllerServer) FindPods(context.Context, *model.ElasticsearchClusterId) (*model4.Pods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPods not implemented")
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

func _ElasticsearchClusterQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchClusterQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchClusterQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchClusterQueryControllerServer).List(ctx, req.(*rpc.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
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

func _ElasticsearchClusterQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchClusterQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchClusterQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchClusterQueryControllerServer).FindByProductId(ctx, req.(*model1.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchClusterQueryController_FindByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.EnvironmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchClusterQueryControllerServer).FindByEnvironmentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchClusterQueryController_FindByEnvironmentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchClusterQueryControllerServer).FindByEnvironmentId(ctx, req.(*model2.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticsearchClusterQueryController_FindByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model3.KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchClusterQueryControllerServer).FindByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchClusterQueryController_FindByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchClusterQueryControllerServer).FindByKubeClusterId(ctx, req.(*model3.KubeClusterId))
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

func _ElasticsearchClusterQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ElasticsearchClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticsearchClusterQueryControllerServer).FindPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElasticsearchClusterQueryController_FindPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticsearchClusterQueryControllerServer).FindPods(ctx, req.(*model.ElasticsearchClusterId))
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
			MethodName: "list",
			Handler:    _ElasticsearchClusterQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _ElasticsearchClusterQueryController_GetById_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _ElasticsearchClusterQueryController_FindByProductId_Handler,
		},
		{
			MethodName: "findByEnvironmentId",
			Handler:    _ElasticsearchClusterQueryController_FindByEnvironmentId_Handler,
		},
		{
			MethodName: "findByKubeClusterId",
			Handler:    _ElasticsearchClusterQueryController_FindByKubeClusterId_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _ElasticsearchClusterQueryController_GetPassword_Handler,
		},
		{
			MethodName: "findPods",
			Handler:    _ElasticsearchClusterQueryController_FindPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/elasticsearchcluster/service/query.proto",
}
