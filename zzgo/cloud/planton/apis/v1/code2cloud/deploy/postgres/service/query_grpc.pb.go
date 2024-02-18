// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/postgres/service/query.proto

package service

import (
	context "context"
	model4 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/kubecluster/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/postgres/model"
	model3 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/environment/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/rpc/pagination/model"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/integration/kubernetes/resource"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/product/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PostgresClusterQueryController_List_FullMethodName                = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.service.PostgresClusterQueryController/list"
	PostgresClusterQueryController_GetById_FullMethodName             = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.service.PostgresClusterQueryController/getById"
	PostgresClusterQueryController_FindByProductId_FullMethodName     = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.service.PostgresClusterQueryController/findByProductId"
	PostgresClusterQueryController_FindByEnvironmentId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.service.PostgresClusterQueryController/findByEnvironmentId"
	PostgresClusterQueryController_FindByKubeClusterId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.service.PostgresClusterQueryController/findByKubeClusterId"
	PostgresClusterQueryController_GetPassword_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.service.PostgresClusterQueryController/getPassword"
	PostgresClusterQueryController_FindPods_FullMethodName            = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.service.PostgresClusterQueryController/findPods"
)

// PostgresClusterQueryControllerClient is the client API for PostgresClusterQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostgresClusterQueryControllerClient interface {
	// list all postgres-clusters on planton cloud for the requested page. This is intended for use on portal.
	List(ctx context.Context, in *model.PageInfo, opts ...grpc.CallOption) (*model1.PostgresClusterList, error)
	// look up a postgres-cluster using postgres-cluster id
	GetById(ctx context.Context, in *model1.PostgresClusterId, opts ...grpc.CallOption) (*model1.PostgresCluster, error)
	// find postgres-clusters by product id.
	// response contains only objects that the authenticated user account id has viewer access to.
	FindByProductId(ctx context.Context, in *model2.ProductId, opts ...grpc.CallOption) (*model1.PostgresClusters, error)
	// find postgres-clusters by environment id.
	// response contains only objects that the authenticated user account id has viewer access to.
	FindByEnvironmentId(ctx context.Context, in *model3.EnvironmentId, opts ...grpc.CallOption) (*model1.PostgresClusters, error)
	// find postgres-clusters by kube-cluster
	FindByKubeClusterId(ctx context.Context, in *model4.KubeClusterId, opts ...grpc.CallOption) (*model1.PostgresClusters, error)
	// look up postgres-cluster password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *model1.PostgresClusterId, opts ...grpc.CallOption) (*model1.PostgresClusterPassword, error)
	// lookup pods of a postgres-cluster deployment
	FindPods(ctx context.Context, in *model1.PostgresClusterId, opts ...grpc.CallOption) (*resource.Pods, error)
}

type postgresClusterQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewPostgresClusterQueryControllerClient(cc grpc.ClientConnInterface) PostgresClusterQueryControllerClient {
	return &postgresClusterQueryControllerClient{cc}
}

func (c *postgresClusterQueryControllerClient) List(ctx context.Context, in *model.PageInfo, opts ...grpc.CallOption) (*model1.PostgresClusterList, error) {
	out := new(model1.PostgresClusterList)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) GetById(ctx context.Context, in *model1.PostgresClusterId, opts ...grpc.CallOption) (*model1.PostgresCluster, error) {
	out := new(model1.PostgresCluster)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) FindByProductId(ctx context.Context, in *model2.ProductId, opts ...grpc.CallOption) (*model1.PostgresClusters, error) {
	out := new(model1.PostgresClusters)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) FindByEnvironmentId(ctx context.Context, in *model3.EnvironmentId, opts ...grpc.CallOption) (*model1.PostgresClusters, error) {
	out := new(model1.PostgresClusters)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_FindByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) FindByKubeClusterId(ctx context.Context, in *model4.KubeClusterId, opts ...grpc.CallOption) (*model1.PostgresClusters, error) {
	out := new(model1.PostgresClusters)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_FindByKubeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) GetPassword(ctx context.Context, in *model1.PostgresClusterId, opts ...grpc.CallOption) (*model1.PostgresClusterPassword, error) {
	out := new(model1.PostgresClusterPassword)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_GetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) FindPods(ctx context.Context, in *model1.PostgresClusterId, opts ...grpc.CallOption) (*resource.Pods, error) {
	out := new(resource.Pods)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_FindPods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostgresClusterQueryControllerServer is the server API for PostgresClusterQueryController service.
// All implementations should embed UnimplementedPostgresClusterQueryControllerServer
// for forward compatibility
type PostgresClusterQueryControllerServer interface {
	// list all postgres-clusters on planton cloud for the requested page. This is intended for use on portal.
	List(context.Context, *model.PageInfo) (*model1.PostgresClusterList, error)
	// look up a postgres-cluster using postgres-cluster id
	GetById(context.Context, *model1.PostgresClusterId) (*model1.PostgresCluster, error)
	// find postgres-clusters by product id.
	// response contains only objects that the authenticated user account id has viewer access to.
	FindByProductId(context.Context, *model2.ProductId) (*model1.PostgresClusters, error)
	// find postgres-clusters by environment id.
	// response contains only objects that the authenticated user account id has viewer access to.
	FindByEnvironmentId(context.Context, *model3.EnvironmentId) (*model1.PostgresClusters, error)
	// find postgres-clusters by kube-cluster
	FindByKubeClusterId(context.Context, *model4.KubeClusterId) (*model1.PostgresClusters, error)
	// look up postgres-cluster password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model1.PostgresClusterId) (*model1.PostgresClusterPassword, error)
	// lookup pods of a postgres-cluster deployment
	FindPods(context.Context, *model1.PostgresClusterId) (*resource.Pods, error)
}

// UnimplementedPostgresClusterQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedPostgresClusterQueryControllerServer struct {
}

func (UnimplementedPostgresClusterQueryControllerServer) List(context.Context, *model.PageInfo) (*model1.PostgresClusterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) GetById(context.Context, *model1.PostgresClusterId) (*model1.PostgresCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) FindByProductId(context.Context, *model2.ProductId) (*model1.PostgresClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) FindByEnvironmentId(context.Context, *model3.EnvironmentId) (*model1.PostgresClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEnvironmentId not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) FindByKubeClusterId(context.Context, *model4.KubeClusterId) (*model1.PostgresClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKubeClusterId not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) GetPassword(context.Context, *model1.PostgresClusterId) (*model1.PostgresClusterPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) FindPods(context.Context, *model1.PostgresClusterId) (*resource.Pods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPods not implemented")
}

// UnsafePostgresClusterQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostgresClusterQueryControllerServer will
// result in compilation errors.
type UnsafePostgresClusterQueryControllerServer interface {
	mustEmbedUnimplementedPostgresClusterQueryControllerServer()
}

func RegisterPostgresClusterQueryControllerServer(s grpc.ServiceRegistrar, srv PostgresClusterQueryControllerServer) {
	s.RegisterService(&PostgresClusterQueryController_ServiceDesc, srv)
}

func _PostgresClusterQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterQueryControllerServer).List(ctx, req.(*model.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.PostgresClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterQueryControllerServer).GetById(ctx, req.(*model1.PostgresClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterQueryControllerServer).FindByProductId(ctx, req.(*model2.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_FindByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model3.EnvironmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterQueryControllerServer).FindByEnvironmentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterQueryController_FindByEnvironmentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterQueryControllerServer).FindByEnvironmentId(ctx, req.(*model3.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_FindByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model4.KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterQueryControllerServer).FindByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterQueryController_FindByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterQueryControllerServer).FindByKubeClusterId(ctx, req.(*model4.KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.PostgresClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterQueryControllerServer).GetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterQueryController_GetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterQueryControllerServer).GetPassword(ctx, req.(*model1.PostgresClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.PostgresClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterQueryControllerServer).FindPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterQueryController_FindPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterQueryControllerServer).FindPods(ctx, req.(*model1.PostgresClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

// PostgresClusterQueryController_ServiceDesc is the grpc.ServiceDesc for PostgresClusterQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostgresClusterQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.postgres.service.PostgresClusterQueryController",
	HandlerType: (*PostgresClusterQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _PostgresClusterQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _PostgresClusterQueryController_GetById_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _PostgresClusterQueryController_FindByProductId_Handler,
		},
		{
			MethodName: "findByEnvironmentId",
			Handler:    _PostgresClusterQueryController_FindByEnvironmentId_Handler,
		},
		{
			MethodName: "findByKubeClusterId",
			Handler:    _PostgresClusterQueryController_FindByKubeClusterId_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _PostgresClusterQueryController_GetPassword_Handler,
		},
		{
			MethodName: "findPods",
			Handler:    _PostgresClusterQueryController_FindPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/postgres/service/query.proto",
}
