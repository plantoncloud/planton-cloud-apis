// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/kubernetes/microservicekubernetes/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/kubernetes/microservicekubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MicroserviceKubernetesQueryController_Get_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.microservicekubernetes.service.MicroserviceKubernetesQueryController/get"
	MicroserviceKubernetesQueryController_GetEnvVarMap_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.microservicekubernetes.service.MicroserviceKubernetesQueryController/getEnvVarMap"
)

// MicroserviceKubernetesQueryControllerClient is the client API for MicroserviceKubernetesQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicroserviceKubernetesQueryControllerClient interface {
	// look up microservice-kubernetes using microservice-kubernetes id
	Get(ctx context.Context, in *model.MicroserviceKubernetesId, opts ...grpc.CallOption) (*model.MicroserviceKubernetes, error)
	GetEnvVarMap(ctx context.Context, in *model.GetMicroserviceKubernetesEnvVarMapInput, opts ...grpc.CallOption) (*model.MicroserviceKubernetesEnvVarMap, error)
}

type microserviceKubernetesQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroserviceKubernetesQueryControllerClient(cc grpc.ClientConnInterface) MicroserviceKubernetesQueryControllerClient {
	return &microserviceKubernetesQueryControllerClient{cc}
}

func (c *microserviceKubernetesQueryControllerClient) Get(ctx context.Context, in *model.MicroserviceKubernetesId, opts ...grpc.CallOption) (*model.MicroserviceKubernetes, error) {
	out := new(model.MicroserviceKubernetes)
	err := c.cc.Invoke(ctx, MicroserviceKubernetesQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceKubernetesQueryControllerClient) GetEnvVarMap(ctx context.Context, in *model.GetMicroserviceKubernetesEnvVarMapInput, opts ...grpc.CallOption) (*model.MicroserviceKubernetesEnvVarMap, error) {
	out := new(model.MicroserviceKubernetesEnvVarMap)
	err := c.cc.Invoke(ctx, MicroserviceKubernetesQueryController_GetEnvVarMap_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroserviceKubernetesQueryControllerServer is the server API for MicroserviceKubernetesQueryController service.
// All implementations should embed UnimplementedMicroserviceKubernetesQueryControllerServer
// for forward compatibility
type MicroserviceKubernetesQueryControllerServer interface {
	// look up microservice-kubernetes using microservice-kubernetes id
	Get(context.Context, *model.MicroserviceKubernetesId) (*model.MicroserviceKubernetes, error)
	GetEnvVarMap(context.Context, *model.GetMicroserviceKubernetesEnvVarMapInput) (*model.MicroserviceKubernetesEnvVarMap, error)
}

// UnimplementedMicroserviceKubernetesQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedMicroserviceKubernetesQueryControllerServer struct {
}

func (UnimplementedMicroserviceKubernetesQueryControllerServer) Get(context.Context, *model.MicroserviceKubernetesId) (*model.MicroserviceKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMicroserviceKubernetesQueryControllerServer) GetEnvVarMap(context.Context, *model.GetMicroserviceKubernetesEnvVarMapInput) (*model.MicroserviceKubernetesEnvVarMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnvVarMap not implemented")
}

// UnsafeMicroserviceKubernetesQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroserviceKubernetesQueryControllerServer will
// result in compilation errors.
type UnsafeMicroserviceKubernetesQueryControllerServer interface {
	mustEmbedUnimplementedMicroserviceKubernetesQueryControllerServer()
}

func RegisterMicroserviceKubernetesQueryControllerServer(s grpc.ServiceRegistrar, srv MicroserviceKubernetesQueryControllerServer) {
	s.RegisterService(&MicroserviceKubernetesQueryController_ServiceDesc, srv)
}

func _MicroserviceKubernetesQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MicroserviceKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceKubernetesQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceKubernetesQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceKubernetesQueryControllerServer).Get(ctx, req.(*model.MicroserviceKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceKubernetesQueryController_GetEnvVarMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetMicroserviceKubernetesEnvVarMapInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceKubernetesQueryControllerServer).GetEnvVarMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceKubernetesQueryController_GetEnvVarMap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceKubernetesQueryControllerServer).GetEnvVarMap(ctx, req.(*model.GetMicroserviceKubernetesEnvVarMapInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MicroserviceKubernetesQueryController_ServiceDesc is the grpc.ServiceDesc for MicroserviceKubernetesQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MicroserviceKubernetesQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.kubernetes.microservicekubernetes.service.MicroserviceKubernetesQueryController",
	HandlerType: (*MicroserviceKubernetesQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _MicroserviceKubernetesQueryController_Get_Handler,
		},
		{
			MethodName: "getEnvVarMap",
			Handler:    _MicroserviceKubernetesQueryController_GetEnvVarMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/kubernetes/microservicekubernetes/service/query.proto",
}
