// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/mongodbcluster/service/query.proto

package service

import (
	context "context"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/environment/model"
	model3 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/mongodbcluster/model"
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
	MongodbClusterQueryController_List_FullMethodName                = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterQueryController/list"
	MongodbClusterQueryController_GetById_FullMethodName             = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterQueryController/getById"
	MongodbClusterQueryController_FindByProductId_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterQueryController/findByProductId"
	MongodbClusterQueryController_FindByEnvironmentId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterQueryController/findByEnvironmentId"
	MongodbClusterQueryController_FindByKubeClusterId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterQueryController/findByKubeClusterId"
	MongodbClusterQueryController_GetPassword_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterQueryController/getPassword"
	MongodbClusterQueryController_FindPods_FullMethodName            = "/cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterQueryController/findPods"
)

// MongodbClusterQueryControllerClient is the client API for MongodbClusterQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MongodbClusterQueryControllerClient interface {
	// list all mongodb-clusters on planton cluster for the requested page.
	List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.MongodbClusterList, error)
	// look up mongodb-cluster using mongodb-cluster id
	GetById(ctx context.Context, in *model.MongodbClusterId, opts ...grpc.CallOption) (*model.MongodbCluster, error)
	// find mongodb-clusters by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.MongodbClusters, error)
	// find mongodb-clusters by environment
	FindByEnvironmentId(ctx context.Context, in *model2.EnvironmentId, opts ...grpc.CallOption) (*model.MongodbClusters, error)
	FindByKubeClusterId(ctx context.Context, in *model3.KubeClusterId, opts ...grpc.CallOption) (*model.MongodbClusters, error)
	// look up mongodb-cluster sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *model.MongodbClusterId, opts ...grpc.CallOption) (*model.MongodbClusterPassword, error)
	// lookup pods of a mongodb-cluster deployed to a environment
	FindPods(ctx context.Context, in *model.MongodbClusterId, opts ...grpc.CallOption) (*model4.Pods, error)
}

type mongodbClusterQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMongodbClusterQueryControllerClient(cc grpc.ClientConnInterface) MongodbClusterQueryControllerClient {
	return &mongodbClusterQueryControllerClient{cc}
}

func (c *mongodbClusterQueryControllerClient) List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.MongodbClusterList, error) {
	out := new(model.MongodbClusterList)
	err := c.cc.Invoke(ctx, MongodbClusterQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterQueryControllerClient) GetById(ctx context.Context, in *model.MongodbClusterId, opts ...grpc.CallOption) (*model.MongodbCluster, error) {
	out := new(model.MongodbCluster)
	err := c.cc.Invoke(ctx, MongodbClusterQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterQueryControllerClient) FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.MongodbClusters, error) {
	out := new(model.MongodbClusters)
	err := c.cc.Invoke(ctx, MongodbClusterQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterQueryControllerClient) FindByEnvironmentId(ctx context.Context, in *model2.EnvironmentId, opts ...grpc.CallOption) (*model.MongodbClusters, error) {
	out := new(model.MongodbClusters)
	err := c.cc.Invoke(ctx, MongodbClusterQueryController_FindByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterQueryControllerClient) FindByKubeClusterId(ctx context.Context, in *model3.KubeClusterId, opts ...grpc.CallOption) (*model.MongodbClusters, error) {
	out := new(model.MongodbClusters)
	err := c.cc.Invoke(ctx, MongodbClusterQueryController_FindByKubeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterQueryControllerClient) GetPassword(ctx context.Context, in *model.MongodbClusterId, opts ...grpc.CallOption) (*model.MongodbClusterPassword, error) {
	out := new(model.MongodbClusterPassword)
	err := c.cc.Invoke(ctx, MongodbClusterQueryController_GetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbClusterQueryControllerClient) FindPods(ctx context.Context, in *model.MongodbClusterId, opts ...grpc.CallOption) (*model4.Pods, error) {
	out := new(model4.Pods)
	err := c.cc.Invoke(ctx, MongodbClusterQueryController_FindPods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MongodbClusterQueryControllerServer is the server API for MongodbClusterQueryController service.
// All implementations should embed UnimplementedMongodbClusterQueryControllerServer
// for forward compatibility
type MongodbClusterQueryControllerServer interface {
	// list all mongodb-clusters on planton cluster for the requested page.
	List(context.Context, *rpc.PageInfo) (*model.MongodbClusterList, error)
	// look up mongodb-cluster using mongodb-cluster id
	GetById(context.Context, *model.MongodbClusterId) (*model.MongodbCluster, error)
	// find mongodb-clusters by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(context.Context, *model1.ProductId) (*model.MongodbClusters, error)
	// find mongodb-clusters by environment
	FindByEnvironmentId(context.Context, *model2.EnvironmentId) (*model.MongodbClusters, error)
	FindByKubeClusterId(context.Context, *model3.KubeClusterId) (*model.MongodbClusters, error)
	// look up mongodb-cluster sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model.MongodbClusterId) (*model.MongodbClusterPassword, error)
	// lookup pods of a mongodb-cluster deployed to a environment
	FindPods(context.Context, *model.MongodbClusterId) (*model4.Pods, error)
}

// UnimplementedMongodbClusterQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedMongodbClusterQueryControllerServer struct {
}

func (UnimplementedMongodbClusterQueryControllerServer) List(context.Context, *rpc.PageInfo) (*model.MongodbClusterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMongodbClusterQueryControllerServer) GetById(context.Context, *model.MongodbClusterId) (*model.MongodbCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedMongodbClusterQueryControllerServer) FindByProductId(context.Context, *model1.ProductId) (*model.MongodbClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedMongodbClusterQueryControllerServer) FindByEnvironmentId(context.Context, *model2.EnvironmentId) (*model.MongodbClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEnvironmentId not implemented")
}
func (UnimplementedMongodbClusterQueryControllerServer) FindByKubeClusterId(context.Context, *model3.KubeClusterId) (*model.MongodbClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKubeClusterId not implemented")
}
func (UnimplementedMongodbClusterQueryControllerServer) GetPassword(context.Context, *model.MongodbClusterId) (*model.MongodbClusterPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}
func (UnimplementedMongodbClusterQueryControllerServer) FindPods(context.Context, *model.MongodbClusterId) (*model4.Pods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPods not implemented")
}

// UnsafeMongodbClusterQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MongodbClusterQueryControllerServer will
// result in compilation errors.
type UnsafeMongodbClusterQueryControllerServer interface {
	mustEmbedUnimplementedMongodbClusterQueryControllerServer()
}

func RegisterMongodbClusterQueryControllerServer(s grpc.ServiceRegistrar, srv MongodbClusterQueryControllerServer) {
	s.RegisterService(&MongodbClusterQueryController_ServiceDesc, srv)
}

func _MongodbClusterQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterQueryControllerServer).List(ctx, req.(*rpc.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterQueryControllerServer).GetById(ctx, req.(*model.MongodbClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterQueryControllerServer).FindByProductId(ctx, req.(*model1.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterQueryController_FindByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.EnvironmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterQueryControllerServer).FindByEnvironmentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterQueryController_FindByEnvironmentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterQueryControllerServer).FindByEnvironmentId(ctx, req.(*model2.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterQueryController_FindByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model3.KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterQueryControllerServer).FindByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterQueryController_FindByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterQueryControllerServer).FindByKubeClusterId(ctx, req.(*model3.KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterQueryController_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterQueryControllerServer).GetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterQueryController_GetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterQueryControllerServer).GetPassword(ctx, req.(*model.MongodbClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbClusterQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbClusterQueryControllerServer).FindPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbClusterQueryController_FindPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbClusterQueryControllerServer).FindPods(ctx, req.(*model.MongodbClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

// MongodbClusterQueryController_ServiceDesc is the grpc.ServiceDesc for MongodbClusterQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MongodbClusterQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.mongodbcluster.service.MongodbClusterQueryController",
	HandlerType: (*MongodbClusterQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _MongodbClusterQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _MongodbClusterQueryController_GetById_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _MongodbClusterQueryController_FindByProductId_Handler,
		},
		{
			MethodName: "findByEnvironmentId",
			Handler:    _MongodbClusterQueryController_FindByEnvironmentId_Handler,
		},
		{
			MethodName: "findByKubeClusterId",
			Handler:    _MongodbClusterQueryController_FindByKubeClusterId_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _MongodbClusterQueryController_GetPassword_Handler,
		},
		{
			MethodName: "findPods",
			Handler:    _MongodbClusterQueryController_FindPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/mongodbcluster/service/query.proto",
}
