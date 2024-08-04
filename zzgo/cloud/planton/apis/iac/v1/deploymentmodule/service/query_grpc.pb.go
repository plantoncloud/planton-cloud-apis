// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/iac/v1/deploymentmodule/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/deploymentmodule/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DeploymentModuleQueryController_Get_FullMethodName = "/cloud.planton.apis.iac.v1.deploymentmodule.service.DeploymentModuleQueryController/get"
)

// DeploymentModuleQueryControllerClient is the client API for DeploymentModuleQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeploymentModuleQueryControllerClient interface {
	// lookup deployment-module using deployment-module id
	Get(ctx context.Context, in *model.DeploymentModuleId, opts ...grpc.CallOption) (*model.DeploymentModule, error)
}

type deploymentModuleQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentModuleQueryControllerClient(cc grpc.ClientConnInterface) DeploymentModuleQueryControllerClient {
	return &deploymentModuleQueryControllerClient{cc}
}

func (c *deploymentModuleQueryControllerClient) Get(ctx context.Context, in *model.DeploymentModuleId, opts ...grpc.CallOption) (*model.DeploymentModule, error) {
	out := new(model.DeploymentModule)
	err := c.cc.Invoke(ctx, DeploymentModuleQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentModuleQueryControllerServer is the server API for DeploymentModuleQueryController service.
// All implementations should embed UnimplementedDeploymentModuleQueryControllerServer
// for forward compatibility
type DeploymentModuleQueryControllerServer interface {
	// lookup deployment-module using deployment-module id
	Get(context.Context, *model.DeploymentModuleId) (*model.DeploymentModule, error)
}

// UnimplementedDeploymentModuleQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedDeploymentModuleQueryControllerServer struct {
}

func (UnimplementedDeploymentModuleQueryControllerServer) Get(context.Context, *model.DeploymentModuleId) (*model.DeploymentModule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeDeploymentModuleQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeploymentModuleQueryControllerServer will
// result in compilation errors.
type UnsafeDeploymentModuleQueryControllerServer interface {
	mustEmbedUnimplementedDeploymentModuleQueryControllerServer()
}

func RegisterDeploymentModuleQueryControllerServer(s grpc.ServiceRegistrar, srv DeploymentModuleQueryControllerServer) {
	s.RegisterService(&DeploymentModuleQueryController_ServiceDesc, srv)
}

func _DeploymentModuleQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.DeploymentModuleId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentModuleQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeploymentModuleQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentModuleQueryControllerServer).Get(ctx, req.(*model.DeploymentModuleId))
	}
	return interceptor(ctx, in, info, handler)
}

// DeploymentModuleQueryController_ServiceDesc is the grpc.ServiceDesc for DeploymentModuleQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeploymentModuleQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.iac.v1.deploymentmodule.service.DeploymentModuleQueryController",
	HandlerType: (*DeploymentModuleQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _DeploymentModuleQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/iac/v1/deploymentmodule/service/query.proto",
}