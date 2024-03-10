// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/argocdinstance/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/argocdinstance/model"
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
	ArgocdInstanceQueryController_List_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.argocdinstance.service.ArgocdInstanceQueryController/list"
	ArgocdInstanceQueryController_GetById_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.argocdinstance.service.ArgocdInstanceQueryController/getById"
	ArgocdInstanceQueryController_FindPods_FullMethodName = "/cloud.planton.apis.code2cloud.v1.argocdinstance.service.ArgocdInstanceQueryController/findPods"
)

// ArgocdInstanceQueryControllerClient is the client API for ArgocdInstanceQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArgocdInstanceQueryControllerClient interface {
	// list all argocd-instances on planton cluster for the requested page.
	List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.ArgocdInstanceList, error)
	// look up argocd-instance using argocd-instance id
	GetById(ctx context.Context, in *model.ArgocdInstanceId, opts ...grpc.CallOption) (*model.ArgocdInstance, error)
	// lookup pods of a argocd-instance deployed to a environment
	FindPods(ctx context.Context, in *model.ArgocdInstanceId, opts ...grpc.CallOption) (*model1.Pods, error)
}

type argocdInstanceQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewArgocdInstanceQueryControllerClient(cc grpc.ClientConnInterface) ArgocdInstanceQueryControllerClient {
	return &argocdInstanceQueryControllerClient{cc}
}

func (c *argocdInstanceQueryControllerClient) List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.ArgocdInstanceList, error) {
	out := new(model.ArgocdInstanceList)
	err := c.cc.Invoke(ctx, ArgocdInstanceQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *argocdInstanceQueryControllerClient) GetById(ctx context.Context, in *model.ArgocdInstanceId, opts ...grpc.CallOption) (*model.ArgocdInstance, error) {
	out := new(model.ArgocdInstance)
	err := c.cc.Invoke(ctx, ArgocdInstanceQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *argocdInstanceQueryControllerClient) FindPods(ctx context.Context, in *model.ArgocdInstanceId, opts ...grpc.CallOption) (*model1.Pods, error) {
	out := new(model1.Pods)
	err := c.cc.Invoke(ctx, ArgocdInstanceQueryController_FindPods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArgocdInstanceQueryControllerServer is the server API for ArgocdInstanceQueryController service.
// All implementations should embed UnimplementedArgocdInstanceQueryControllerServer
// for forward compatibility
type ArgocdInstanceQueryControllerServer interface {
	// list all argocd-instances on planton cluster for the requested page.
	List(context.Context, *rpc.PageInfo) (*model.ArgocdInstanceList, error)
	// look up argocd-instance using argocd-instance id
	GetById(context.Context, *model.ArgocdInstanceId) (*model.ArgocdInstance, error)
	// lookup pods of a argocd-instance deployed to a environment
	FindPods(context.Context, *model.ArgocdInstanceId) (*model1.Pods, error)
}

// UnimplementedArgocdInstanceQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedArgocdInstanceQueryControllerServer struct {
}

func (UnimplementedArgocdInstanceQueryControllerServer) List(context.Context, *rpc.PageInfo) (*model.ArgocdInstanceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedArgocdInstanceQueryControllerServer) GetById(context.Context, *model.ArgocdInstanceId) (*model.ArgocdInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedArgocdInstanceQueryControllerServer) FindPods(context.Context, *model.ArgocdInstanceId) (*model1.Pods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPods not implemented")
}

// UnsafeArgocdInstanceQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArgocdInstanceQueryControllerServer will
// result in compilation errors.
type UnsafeArgocdInstanceQueryControllerServer interface {
	mustEmbedUnimplementedArgocdInstanceQueryControllerServer()
}

func RegisterArgocdInstanceQueryControllerServer(s grpc.ServiceRegistrar, srv ArgocdInstanceQueryControllerServer) {
	s.RegisterService(&ArgocdInstanceQueryController_ServiceDesc, srv)
}

func _ArgocdInstanceQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArgocdInstanceQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArgocdInstanceQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArgocdInstanceQueryControllerServer).List(ctx, req.(*rpc.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArgocdInstanceQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ArgocdInstanceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArgocdInstanceQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArgocdInstanceQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArgocdInstanceQueryControllerServer).GetById(ctx, req.(*model.ArgocdInstanceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArgocdInstanceQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ArgocdInstanceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArgocdInstanceQueryControllerServer).FindPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArgocdInstanceQueryController_FindPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArgocdInstanceQueryControllerServer).FindPods(ctx, req.(*model.ArgocdInstanceId))
	}
	return interceptor(ctx, in, info, handler)
}

// ArgocdInstanceQueryController_ServiceDesc is the grpc.ServiceDesc for ArgocdInstanceQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArgocdInstanceQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.argocdinstance.service.ArgocdInstanceQueryController",
	HandlerType: (*ArgocdInstanceQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _ArgocdInstanceQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _ArgocdInstanceQueryController_GetById_Handler,
		},
		{
			MethodName: "findPods",
			Handler:    _ArgocdInstanceQueryController_FindPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/argocdinstance/service/query.proto",
}
