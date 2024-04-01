// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/locustcluster/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/locustcluster/model"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/rpc"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/apiresources/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	LocustClusterQueryController_List_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.locustcluster.service.LocustClusterQueryController/list"
	LocustClusterQueryController_GetById_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.locustcluster.service.LocustClusterQueryController/getById"
	LocustClusterQueryController_FindPods_FullMethodName = "/cloud.planton.apis.code2cloud.v1.locustcluster.service.LocustClusterQueryController/findPods"
)

// LocustClusterQueryControllerClient is the client API for LocustClusterQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocustClusterQueryControllerClient interface {
	// list all locust-clusters on planton cluster for the requested page.
	List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.LocustClusterList, error)
	// look up locust-cluster using locust-cluster id
	GetById(ctx context.Context, in *model.LocustClusterId, opts ...grpc.CallOption) (*model.LocustCluster, error)
	// lookup pods of a locust-cluster deployed to a environment
	FindPods(ctx context.Context, in *model.LocustClusterId, opts ...grpc.CallOption) (*model1.Pods, error)
}

type locustClusterQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewLocustClusterQueryControllerClient(cc grpc.ClientConnInterface) LocustClusterQueryControllerClient {
	return &locustClusterQueryControllerClient{cc}
}

func (c *locustClusterQueryControllerClient) List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.LocustClusterList, error) {
	out := new(model.LocustClusterList)
	err := c.cc.Invoke(ctx, LocustClusterQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locustClusterQueryControllerClient) GetById(ctx context.Context, in *model.LocustClusterId, opts ...grpc.CallOption) (*model.LocustCluster, error) {
	out := new(model.LocustCluster)
	err := c.cc.Invoke(ctx, LocustClusterQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locustClusterQueryControllerClient) FindPods(ctx context.Context, in *model.LocustClusterId, opts ...grpc.CallOption) (*model1.Pods, error) {
	out := new(model1.Pods)
	err := c.cc.Invoke(ctx, LocustClusterQueryController_FindPods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocustClusterQueryControllerServer is the server API for LocustClusterQueryController service.
// All implementations should embed UnimplementedLocustClusterQueryControllerServer
// for forward compatibility
type LocustClusterQueryControllerServer interface {
	// list all locust-clusters on planton cluster for the requested page.
	List(context.Context, *rpc.PageInfo) (*model.LocustClusterList, error)
	// look up locust-cluster using locust-cluster id
	GetById(context.Context, *model.LocustClusterId) (*model.LocustCluster, error)
	// lookup pods of a locust-cluster deployed to a environment
	FindPods(context.Context, *model.LocustClusterId) (*model1.Pods, error)
}

// UnimplementedLocustClusterQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedLocustClusterQueryControllerServer struct {
}

func (UnimplementedLocustClusterQueryControllerServer) List(context.Context, *rpc.PageInfo) (*model.LocustClusterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLocustClusterQueryControllerServer) GetById(context.Context, *model.LocustClusterId) (*model.LocustCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedLocustClusterQueryControllerServer) FindPods(context.Context, *model.LocustClusterId) (*model1.Pods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPods not implemented")
}

// UnsafeLocustClusterQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocustClusterQueryControllerServer will
// result in compilation errors.
type UnsafeLocustClusterQueryControllerServer interface {
	mustEmbedUnimplementedLocustClusterQueryControllerServer()
}

func RegisterLocustClusterQueryControllerServer(s grpc.ServiceRegistrar, srv LocustClusterQueryControllerServer) {
	s.RegisterService(&LocustClusterQueryController_ServiceDesc, srv)
}

func _LocustClusterQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocustClusterQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocustClusterQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocustClusterQueryControllerServer).List(ctx, req.(*rpc.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocustClusterQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.LocustClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocustClusterQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocustClusterQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocustClusterQueryControllerServer).GetById(ctx, req.(*model.LocustClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocustClusterQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.LocustClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocustClusterQueryControllerServer).FindPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocustClusterQueryController_FindPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocustClusterQueryControllerServer).FindPods(ctx, req.(*model.LocustClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

// LocustClusterQueryController_ServiceDesc is the grpc.ServiceDesc for LocustClusterQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocustClusterQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.locustcluster.service.LocustClusterQueryController",
	HandlerType: (*LocustClusterQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _LocustClusterQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _LocustClusterQueryController_GetById_Handler,
		},
		{
			MethodName: "findPods",
			Handler:    _LocustClusterQueryController_FindPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/locustcluster/service/query.proto",
}
