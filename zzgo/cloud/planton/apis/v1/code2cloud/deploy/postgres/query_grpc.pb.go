// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/postgres/query.proto

package postgres

import (
	context "context"
	kubecluster "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/kubecluster"
	environment "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/environment"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/integration/kubernetes/resource"
	product "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/product"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PostgresClusterQueryController_List_FullMethodName                = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.PostgresClusterQueryController/list"
	PostgresClusterQueryController_GetById_FullMethodName             = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.PostgresClusterQueryController/getById"
	PostgresClusterQueryController_FindByProductId_FullMethodName     = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.PostgresClusterQueryController/findByProductId"
	PostgresClusterQueryController_FindByEnvironmentId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.PostgresClusterQueryController/findByEnvironmentId"
	PostgresClusterQueryController_FindByKubeClusterId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.PostgresClusterQueryController/findByKubeClusterId"
	PostgresClusterQueryController_GetPassword_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.PostgresClusterQueryController/getPassword"
	PostgresClusterQueryController_FindPods_FullMethodName            = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.PostgresClusterQueryController/findPods"
)

// PostgresClusterQueryControllerClient is the client API for PostgresClusterQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostgresClusterQueryControllerClient interface {
	// list all postgres-clusters on planton cloud for the requested page. This is intended for use on portal.
	List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*PostgresClusterList, error)
	// look up a postgres-cluster using postgres-cluster id
	GetById(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresCluster, error)
	// find postgres-clusters by product id.
	// response contains only objects that the authenticated user account id has viewer access to.
	FindByProductId(ctx context.Context, in *product.ProductId, opts ...grpc.CallOption) (*PostgresClusters, error)
	// find postgres-clusters by environment id.
	// response contains only objects that the authenticated user account id has viewer access to.
	FindByEnvironmentId(ctx context.Context, in *environment.EnvironmentId, opts ...grpc.CallOption) (*PostgresClusters, error)
	// find postgres-clusters by kube-cluster
	FindByKubeClusterId(ctx context.Context, in *kubecluster.KubeClusterId, opts ...grpc.CallOption) (*PostgresClusters, error)
	// look up postgres-cluster password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresClusterPassword, error)
	// lookup pods of a postgres-cluster deployment
	FindPods(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*resource.Pods, error)
}

type postgresClusterQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewPostgresClusterQueryControllerClient(cc grpc.ClientConnInterface) PostgresClusterQueryControllerClient {
	return &postgresClusterQueryControllerClient{cc}
}

func (c *postgresClusterQueryControllerClient) List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*PostgresClusterList, error) {
	out := new(PostgresClusterList)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) GetById(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresCluster, error) {
	out := new(PostgresCluster)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) FindByProductId(ctx context.Context, in *product.ProductId, opts ...grpc.CallOption) (*PostgresClusters, error) {
	out := new(PostgresClusters)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) FindByEnvironmentId(ctx context.Context, in *environment.EnvironmentId, opts ...grpc.CallOption) (*PostgresClusters, error) {
	out := new(PostgresClusters)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_FindByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) FindByKubeClusterId(ctx context.Context, in *kubecluster.KubeClusterId, opts ...grpc.CallOption) (*PostgresClusters, error) {
	out := new(PostgresClusters)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_FindByKubeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) GetPassword(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresClusterPassword, error) {
	out := new(PostgresClusterPassword)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_GetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) FindPods(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*resource.Pods, error) {
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
	List(context.Context, *pagination.PageInfo) (*PostgresClusterList, error)
	// look up a postgres-cluster using postgres-cluster id
	GetById(context.Context, *PostgresClusterId) (*PostgresCluster, error)
	// find postgres-clusters by product id.
	// response contains only objects that the authenticated user account id has viewer access to.
	FindByProductId(context.Context, *product.ProductId) (*PostgresClusters, error)
	// find postgres-clusters by environment id.
	// response contains only objects that the authenticated user account id has viewer access to.
	FindByEnvironmentId(context.Context, *environment.EnvironmentId) (*PostgresClusters, error)
	// find postgres-clusters by kube-cluster
	FindByKubeClusterId(context.Context, *kubecluster.KubeClusterId) (*PostgresClusters, error)
	// look up postgres-cluster password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *PostgresClusterId) (*PostgresClusterPassword, error)
	// lookup pods of a postgres-cluster deployment
	FindPods(context.Context, *PostgresClusterId) (*resource.Pods, error)
}

// UnimplementedPostgresClusterQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedPostgresClusterQueryControllerServer struct {
}

func (UnimplementedPostgresClusterQueryControllerServer) List(context.Context, *pagination.PageInfo) (*PostgresClusterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) GetById(context.Context, *PostgresClusterId) (*PostgresCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) FindByProductId(context.Context, *product.ProductId) (*PostgresClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) FindByEnvironmentId(context.Context, *environment.EnvironmentId) (*PostgresClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEnvironmentId not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) FindByKubeClusterId(context.Context, *kubecluster.KubeClusterId) (*PostgresClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKubeClusterId not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) GetPassword(context.Context, *PostgresClusterId) (*PostgresClusterPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) FindPods(context.Context, *PostgresClusterId) (*resource.Pods, error) {
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
	in := new(pagination.PageInfo)
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
		return srv.(PostgresClusterQueryControllerServer).List(ctx, req.(*pagination.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresClusterId)
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
		return srv.(PostgresClusterQueryControllerServer).GetById(ctx, req.(*PostgresClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(product.ProductId)
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
		return srv.(PostgresClusterQueryControllerServer).FindByProductId(ctx, req.(*product.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_FindByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(environment.EnvironmentId)
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
		return srv.(PostgresClusterQueryControllerServer).FindByEnvironmentId(ctx, req.(*environment.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_FindByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kubecluster.KubeClusterId)
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
		return srv.(PostgresClusterQueryControllerServer).FindByKubeClusterId(ctx, req.(*kubecluster.KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresClusterId)
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
		return srv.(PostgresClusterQueryControllerServer).GetPassword(ctx, req.(*PostgresClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresClusterId)
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
		return srv.(PostgresClusterQueryControllerServer).FindPods(ctx, req.(*PostgresClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

// PostgresClusterQueryController_ServiceDesc is the grpc.ServiceDesc for PostgresClusterQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostgresClusterQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.postgres.PostgresClusterQueryController",
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
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/postgres/query.proto",
}