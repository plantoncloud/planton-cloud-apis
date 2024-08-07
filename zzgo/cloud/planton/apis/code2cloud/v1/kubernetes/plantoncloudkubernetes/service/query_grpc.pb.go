// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/kubernetes/plantoncloudkubernetes/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/plantoncloudkubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PlantonCloudKubernetesQueryController_Get_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kubernetes.plantoncloudkubernetes.service.PlantonCloudKubernetesQueryController/get"
)

// PlantonCloudKubernetesQueryControllerClient is the client API for PlantonCloudKubernetesQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlantonCloudKubernetesQueryControllerClient interface {
	// look up planton-cloud-kubernetes using planton-cloud-kubernetes id
	Get(ctx context.Context, in *model.PlantonCloudKubernetesId, opts ...grpc.CallOption) (*model.PlantonCloudKubernetes, error)
}

type plantonCloudKubernetesQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewPlantonCloudKubernetesQueryControllerClient(cc grpc.ClientConnInterface) PlantonCloudKubernetesQueryControllerClient {
	return &plantonCloudKubernetesQueryControllerClient{cc}
}

func (c *plantonCloudKubernetesQueryControllerClient) Get(ctx context.Context, in *model.PlantonCloudKubernetesId, opts ...grpc.CallOption) (*model.PlantonCloudKubernetes, error) {
	out := new(model.PlantonCloudKubernetes)
	err := c.cc.Invoke(ctx, PlantonCloudKubernetesQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlantonCloudKubernetesQueryControllerServer is the server API for PlantonCloudKubernetesQueryController service.
// All implementations should embed UnimplementedPlantonCloudKubernetesQueryControllerServer
// for forward compatibility
type PlantonCloudKubernetesQueryControllerServer interface {
	// look up planton-cloud-kubernetes using planton-cloud-kubernetes id
	Get(context.Context, *model.PlantonCloudKubernetesId) (*model.PlantonCloudKubernetes, error)
}

// UnimplementedPlantonCloudKubernetesQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedPlantonCloudKubernetesQueryControllerServer struct {
}

func (UnimplementedPlantonCloudKubernetesQueryControllerServer) Get(context.Context, *model.PlantonCloudKubernetesId) (*model.PlantonCloudKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafePlantonCloudKubernetesQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlantonCloudKubernetesQueryControllerServer will
// result in compilation errors.
type UnsafePlantonCloudKubernetesQueryControllerServer interface {
	mustEmbedUnimplementedPlantonCloudKubernetesQueryControllerServer()
}

func RegisterPlantonCloudKubernetesQueryControllerServer(s grpc.ServiceRegistrar, srv PlantonCloudKubernetesQueryControllerServer) {
	s.RegisterService(&PlantonCloudKubernetesQueryController_ServiceDesc, srv)
}

func _PlantonCloudKubernetesQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PlantonCloudKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlantonCloudKubernetesQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlantonCloudKubernetesQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlantonCloudKubernetesQueryControllerServer).Get(ctx, req.(*model.PlantonCloudKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

// PlantonCloudKubernetesQueryController_ServiceDesc is the grpc.ServiceDesc for PlantonCloudKubernetesQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlantonCloudKubernetesQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.kubernetes.plantoncloudkubernetes.service.PlantonCloudKubernetesQueryController",
	HandlerType: (*PlantonCloudKubernetesQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _PlantonCloudKubernetesQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/kubernetes/plantoncloudkubernetes/service/query.proto",
}
