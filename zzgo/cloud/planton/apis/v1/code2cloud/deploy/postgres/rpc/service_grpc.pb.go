// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/postgres/rpc/service.proto

package rpc

import (
	context "context"
	rpc2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/kubecluster/rpc"
	rpc1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/environment/rpc"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/integration/kubernetes/resource"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/product/rpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PostgresClusterCommandController_Create_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterCommandController/create"
	PostgresClusterCommandController_Update_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterCommandController/update"
	PostgresClusterCommandController_Delete_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterCommandController/delete"
	PostgresClusterCommandController_Restore_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterCommandController/restore"
	PostgresClusterCommandController_Restart_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterCommandController/restart"
	PostgresClusterCommandController_Pause_FullMethodName   = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterCommandController/pause"
	PostgresClusterCommandController_Unpause_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterCommandController/unpause"
)

// PostgresClusterCommandControllerClient is the client API for PostgresClusterCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostgresClusterCommandControllerClient interface {
	// create postgres-cluster
	Create(ctx context.Context, in *PostgresCluster, opts ...grpc.CallOption) (*PostgresCluster, error)
	// update an existing postgres-cluster
	Update(ctx context.Context, in *PostgresCluster, opts ...grpc.CallOption) (*PostgresCluster, error)
	// delete an existing postgres-cluster
	Delete(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresCluster, error)
	// restore a deleted postgres-cluster in a environment
	Restore(ctx context.Context, in *PostgresCluster, opts ...grpc.CallOption) (*PostgresCluster, error)
	// restart a postgres-cluster running in a environment.
	// postgres-cluster is restarted by deleting running pods which will be automatically recreated by kubernetes
	Restart(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresCluster, error)
	// pause a postgres-cluster running in a environment.
	// postgres-cluster is paused by scaling down number of replicas of
	// the kubernetes deployment/stateful sets to zero in the environment.
	Pause(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresCluster, error)
	// unpause a previously paused postgres-cluster running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the postgres-cluster.
	Unpause(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresCluster, error)
}

type postgresClusterCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewPostgresClusterCommandControllerClient(cc grpc.ClientConnInterface) PostgresClusterCommandControllerClient {
	return &postgresClusterCommandControllerClient{cc}
}

func (c *postgresClusterCommandControllerClient) Create(ctx context.Context, in *PostgresCluster, opts ...grpc.CallOption) (*PostgresCluster, error) {
	out := new(PostgresCluster)
	err := c.cc.Invoke(ctx, PostgresClusterCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterCommandControllerClient) Update(ctx context.Context, in *PostgresCluster, opts ...grpc.CallOption) (*PostgresCluster, error) {
	out := new(PostgresCluster)
	err := c.cc.Invoke(ctx, PostgresClusterCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterCommandControllerClient) Delete(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresCluster, error) {
	out := new(PostgresCluster)
	err := c.cc.Invoke(ctx, PostgresClusterCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterCommandControllerClient) Restore(ctx context.Context, in *PostgresCluster, opts ...grpc.CallOption) (*PostgresCluster, error) {
	out := new(PostgresCluster)
	err := c.cc.Invoke(ctx, PostgresClusterCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterCommandControllerClient) Restart(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresCluster, error) {
	out := new(PostgresCluster)
	err := c.cc.Invoke(ctx, PostgresClusterCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterCommandControllerClient) Pause(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresCluster, error) {
	out := new(PostgresCluster)
	err := c.cc.Invoke(ctx, PostgresClusterCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterCommandControllerClient) Unpause(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresCluster, error) {
	out := new(PostgresCluster)
	err := c.cc.Invoke(ctx, PostgresClusterCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostgresClusterCommandControllerServer is the server API for PostgresClusterCommandController service.
// All implementations should embed UnimplementedPostgresClusterCommandControllerServer
// for forward compatibility
type PostgresClusterCommandControllerServer interface {
	// create postgres-cluster
	Create(context.Context, *PostgresCluster) (*PostgresCluster, error)
	// update an existing postgres-cluster
	Update(context.Context, *PostgresCluster) (*PostgresCluster, error)
	// delete an existing postgres-cluster
	Delete(context.Context, *PostgresClusterId) (*PostgresCluster, error)
	// restore a deleted postgres-cluster in a environment
	Restore(context.Context, *PostgresCluster) (*PostgresCluster, error)
	// restart a postgres-cluster running in a environment.
	// postgres-cluster is restarted by deleting running pods which will be automatically recreated by kubernetes
	Restart(context.Context, *PostgresClusterId) (*PostgresCluster, error)
	// pause a postgres-cluster running in a environment.
	// postgres-cluster is paused by scaling down number of replicas of
	// the kubernetes deployment/stateful sets to zero in the environment.
	Pause(context.Context, *PostgresClusterId) (*PostgresCluster, error)
	// unpause a previously paused postgres-cluster running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the postgres-cluster.
	Unpause(context.Context, *PostgresClusterId) (*PostgresCluster, error)
}

// UnimplementedPostgresClusterCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedPostgresClusterCommandControllerServer struct {
}

func (UnimplementedPostgresClusterCommandControllerServer) Create(context.Context, *PostgresCluster) (*PostgresCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPostgresClusterCommandControllerServer) Update(context.Context, *PostgresCluster) (*PostgresCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPostgresClusterCommandControllerServer) Delete(context.Context, *PostgresClusterId) (*PostgresCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPostgresClusterCommandControllerServer) Restore(context.Context, *PostgresCluster) (*PostgresCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedPostgresClusterCommandControllerServer) Restart(context.Context, *PostgresClusterId) (*PostgresCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedPostgresClusterCommandControllerServer) Pause(context.Context, *PostgresClusterId) (*PostgresCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedPostgresClusterCommandControllerServer) Unpause(context.Context, *PostgresClusterId) (*PostgresCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}

// UnsafePostgresClusterCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostgresClusterCommandControllerServer will
// result in compilation errors.
type UnsafePostgresClusterCommandControllerServer interface {
	mustEmbedUnimplementedPostgresClusterCommandControllerServer()
}

func RegisterPostgresClusterCommandControllerServer(s grpc.ServiceRegistrar, srv PostgresClusterCommandControllerServer) {
	s.RegisterService(&PostgresClusterCommandController_ServiceDesc, srv)
}

func _PostgresClusterCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresCluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterCommandControllerServer).Create(ctx, req.(*PostgresCluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresCluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterCommandControllerServer).Update(ctx, req.(*PostgresCluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterCommandControllerServer).Delete(ctx, req.(*PostgresClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresCluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterCommandControllerServer).Restore(ctx, req.(*PostgresCluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterCommandControllerServer).Restart(ctx, req.(*PostgresClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterCommandControllerServer).Pause(ctx, req.(*PostgresClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterCommandControllerServer).Unpause(ctx, req.(*PostgresClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

// PostgresClusterCommandController_ServiceDesc is the grpc.ServiceDesc for PostgresClusterCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostgresClusterCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterCommandController",
	HandlerType: (*PostgresClusterCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _PostgresClusterCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _PostgresClusterCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _PostgresClusterCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _PostgresClusterCommandController_Restore_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _PostgresClusterCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _PostgresClusterCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _PostgresClusterCommandController_Unpause_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/postgres/rpc/service.proto",
}

const (
	PostgresClusterQueryController_List_FullMethodName                = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterQueryController/list"
	PostgresClusterQueryController_GetById_FullMethodName             = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterQueryController/getById"
	PostgresClusterQueryController_FindByProductId_FullMethodName     = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterQueryController/findByProductId"
	PostgresClusterQueryController_FindByEnvironmentId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterQueryController/findByEnvironmentId"
	PostgresClusterQueryController_FindByKubeClusterId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterQueryController/findByKubeClusterId"
	PostgresClusterQueryController_GetPassword_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterQueryController/getPassword"
	PostgresClusterQueryController_FindPods_FullMethodName            = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterQueryController/findPods"
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
	FindByProductId(ctx context.Context, in *rpc.ProductId, opts ...grpc.CallOption) (*PostgresClusters, error)
	// find postgres-clusters by environment id.
	// response contains only objects that the authenticated user account id has viewer access to.
	FindByEnvironmentId(ctx context.Context, in *rpc1.EnvironmentId, opts ...grpc.CallOption) (*PostgresClusters, error)
	// find postgres-clusters by kube-cluster
	FindByKubeClusterId(ctx context.Context, in *rpc2.KubeClusterId, opts ...grpc.CallOption) (*PostgresClusters, error)
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

func (c *postgresClusterQueryControllerClient) FindByProductId(ctx context.Context, in *rpc.ProductId, opts ...grpc.CallOption) (*PostgresClusters, error) {
	out := new(PostgresClusters)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) FindByEnvironmentId(ctx context.Context, in *rpc1.EnvironmentId, opts ...grpc.CallOption) (*PostgresClusters, error) {
	out := new(PostgresClusters)
	err := c.cc.Invoke(ctx, PostgresClusterQueryController_FindByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterQueryControllerClient) FindByKubeClusterId(ctx context.Context, in *rpc2.KubeClusterId, opts ...grpc.CallOption) (*PostgresClusters, error) {
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
	FindByProductId(context.Context, *rpc.ProductId) (*PostgresClusters, error)
	// find postgres-clusters by environment id.
	// response contains only objects that the authenticated user account id has viewer access to.
	FindByEnvironmentId(context.Context, *rpc1.EnvironmentId) (*PostgresClusters, error)
	// find postgres-clusters by kube-cluster
	FindByKubeClusterId(context.Context, *rpc2.KubeClusterId) (*PostgresClusters, error)
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
func (UnimplementedPostgresClusterQueryControllerServer) FindByProductId(context.Context, *rpc.ProductId) (*PostgresClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) FindByEnvironmentId(context.Context, *rpc1.EnvironmentId) (*PostgresClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEnvironmentId not implemented")
}
func (UnimplementedPostgresClusterQueryControllerServer) FindByKubeClusterId(context.Context, *rpc2.KubeClusterId) (*PostgresClusters, error) {
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
	in := new(rpc.ProductId)
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
		return srv.(PostgresClusterQueryControllerServer).FindByProductId(ctx, req.(*rpc.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_FindByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc1.EnvironmentId)
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
		return srv.(PostgresClusterQueryControllerServer).FindByEnvironmentId(ctx, req.(*rpc1.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterQueryController_FindByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc2.KubeClusterId)
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
		return srv.(PostgresClusterQueryControllerServer).FindByKubeClusterId(ctx, req.(*rpc2.KubeClusterId))
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
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterQueryController",
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
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/postgres/rpc/service.proto",
}

const (
	PostgresClusterStackController_Preview_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterStackController/preview"
	PostgresClusterStackController_Apply_FullMethodName   = "/cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterStackController/apply"
)

// PostgresClusterStackControllerClient is the client API for PostgresClusterStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostgresClusterStackControllerClient interface {
	// preview postgres-cluster stack
	Preview(ctx context.Context, in *PostgresCluster, opts ...grpc.CallOption) (*PostgresCluster, error)
	// apply postgres-cluster stack
	Apply(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresCluster, error)
}

type postgresClusterStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewPostgresClusterStackControllerClient(cc grpc.ClientConnInterface) PostgresClusterStackControllerClient {
	return &postgresClusterStackControllerClient{cc}
}

func (c *postgresClusterStackControllerClient) Preview(ctx context.Context, in *PostgresCluster, opts ...grpc.CallOption) (*PostgresCluster, error) {
	out := new(PostgresCluster)
	err := c.cc.Invoke(ctx, PostgresClusterStackController_Preview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresClusterStackControllerClient) Apply(ctx context.Context, in *PostgresClusterId, opts ...grpc.CallOption) (*PostgresCluster, error) {
	out := new(PostgresCluster)
	err := c.cc.Invoke(ctx, PostgresClusterStackController_Apply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostgresClusterStackControllerServer is the server API for PostgresClusterStackController service.
// All implementations should embed UnimplementedPostgresClusterStackControllerServer
// for forward compatibility
type PostgresClusterStackControllerServer interface {
	// preview postgres-cluster stack
	Preview(context.Context, *PostgresCluster) (*PostgresCluster, error)
	// apply postgres-cluster stack
	Apply(context.Context, *PostgresClusterId) (*PostgresCluster, error)
}

// UnimplementedPostgresClusterStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedPostgresClusterStackControllerServer struct {
}

func (UnimplementedPostgresClusterStackControllerServer) Preview(context.Context, *PostgresCluster) (*PostgresCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Preview not implemented")
}
func (UnimplementedPostgresClusterStackControllerServer) Apply(context.Context, *PostgresClusterId) (*PostgresCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}

// UnsafePostgresClusterStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostgresClusterStackControllerServer will
// result in compilation errors.
type UnsafePostgresClusterStackControllerServer interface {
	mustEmbedUnimplementedPostgresClusterStackControllerServer()
}

func RegisterPostgresClusterStackControllerServer(s grpc.ServiceRegistrar, srv PostgresClusterStackControllerServer) {
	s.RegisterService(&PostgresClusterStackController_ServiceDesc, srv)
}

func _PostgresClusterStackController_Preview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresCluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterStackControllerServer).Preview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterStackController_Preview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterStackControllerServer).Preview(ctx, req.(*PostgresCluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresClusterStackController_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterStackControllerServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterStackController_Apply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterStackControllerServer).Apply(ctx, req.(*PostgresClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

// PostgresClusterStackController_ServiceDesc is the grpc.ServiceDesc for PostgresClusterStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostgresClusterStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.postgres.rpc.PostgresClusterStackController",
	HandlerType: (*PostgresClusterStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "preview",
			Handler:    _PostgresClusterStackController_Preview_Handler,
		},
		{
			MethodName: "apply",
			Handler:    _PostgresClusterStackController_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/postgres/rpc/service.proto",
}
