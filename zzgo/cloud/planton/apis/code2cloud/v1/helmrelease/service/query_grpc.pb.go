// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/helmrelease/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/helmrelease/model"
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
	HelmReleaseQueryController_List_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.helmrelease.service.HelmReleaseQueryController/list"
	HelmReleaseQueryController_GetById_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.helmrelease.service.HelmReleaseQueryController/getById"
	HelmReleaseQueryController_FindPods_FullMethodName = "/cloud.planton.apis.code2cloud.v1.helmrelease.service.HelmReleaseQueryController/findPods"
)

// HelmReleaseQueryControllerClient is the client API for HelmReleaseQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelmReleaseQueryControllerClient interface {
	// list all helm-releases for the requested page.
	List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.HelmReleaseList, error)
	// look up helm-release using helm-release id
	GetById(ctx context.Context, in *model.HelmReleaseId, opts ...grpc.CallOption) (*model.HelmRelease, error)
	// lookup pods of a helm-release deployed to a environment
	FindPods(ctx context.Context, in *model.HelmReleaseId, opts ...grpc.CallOption) (*model1.Pods, error)
}

type helmReleaseQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewHelmReleaseQueryControllerClient(cc grpc.ClientConnInterface) HelmReleaseQueryControllerClient {
	return &helmReleaseQueryControllerClient{cc}
}

func (c *helmReleaseQueryControllerClient) List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.HelmReleaseList, error) {
	out := new(model.HelmReleaseList)
	err := c.cc.Invoke(ctx, HelmReleaseQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmReleaseQueryControllerClient) GetById(ctx context.Context, in *model.HelmReleaseId, opts ...grpc.CallOption) (*model.HelmRelease, error) {
	out := new(model.HelmRelease)
	err := c.cc.Invoke(ctx, HelmReleaseQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmReleaseQueryControllerClient) FindPods(ctx context.Context, in *model.HelmReleaseId, opts ...grpc.CallOption) (*model1.Pods, error) {
	out := new(model1.Pods)
	err := c.cc.Invoke(ctx, HelmReleaseQueryController_FindPods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelmReleaseQueryControllerServer is the server API for HelmReleaseQueryController service.
// All implementations should embed UnimplementedHelmReleaseQueryControllerServer
// for forward compatibility
type HelmReleaseQueryControllerServer interface {
	// list all helm-releases for the requested page.
	List(context.Context, *rpc.PageInfo) (*model.HelmReleaseList, error)
	// look up helm-release using helm-release id
	GetById(context.Context, *model.HelmReleaseId) (*model.HelmRelease, error)
	// lookup pods of a helm-release deployed to a environment
	FindPods(context.Context, *model.HelmReleaseId) (*model1.Pods, error)
}

// UnimplementedHelmReleaseQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedHelmReleaseQueryControllerServer struct {
}

func (UnimplementedHelmReleaseQueryControllerServer) List(context.Context, *rpc.PageInfo) (*model.HelmReleaseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedHelmReleaseQueryControllerServer) GetById(context.Context, *model.HelmReleaseId) (*model.HelmRelease, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedHelmReleaseQueryControllerServer) FindPods(context.Context, *model.HelmReleaseId) (*model1.Pods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPods not implemented")
}

// UnsafeHelmReleaseQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelmReleaseQueryControllerServer will
// result in compilation errors.
type UnsafeHelmReleaseQueryControllerServer interface {
	mustEmbedUnimplementedHelmReleaseQueryControllerServer()
}

func RegisterHelmReleaseQueryControllerServer(s grpc.ServiceRegistrar, srv HelmReleaseQueryControllerServer) {
	s.RegisterService(&HelmReleaseQueryController_ServiceDesc, srv)
}

func _HelmReleaseQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmReleaseQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelmReleaseQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmReleaseQueryControllerServer).List(ctx, req.(*rpc.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmReleaseQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.HelmReleaseId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmReleaseQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelmReleaseQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmReleaseQueryControllerServer).GetById(ctx, req.(*model.HelmReleaseId))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmReleaseQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.HelmReleaseId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmReleaseQueryControllerServer).FindPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelmReleaseQueryController_FindPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmReleaseQueryControllerServer).FindPods(ctx, req.(*model.HelmReleaseId))
	}
	return interceptor(ctx, in, info, handler)
}

// HelmReleaseQueryController_ServiceDesc is the grpc.ServiceDesc for HelmReleaseQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelmReleaseQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.helmrelease.service.HelmReleaseQueryController",
	HandlerType: (*HelmReleaseQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _HelmReleaseQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _HelmReleaseQueryController_GetById_Handler,
		},
		{
			MethodName: "findPods",
			Handler:    _HelmReleaseQueryController_FindPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/helmrelease/service/query.proto",
}
