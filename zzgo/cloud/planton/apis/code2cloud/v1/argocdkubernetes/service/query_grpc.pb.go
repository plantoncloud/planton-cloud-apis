// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/argocdkubernetes/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/argocdkubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ArgocdKubernetesQueryController_GetById_FullMethodName = "/cloud.planton.apis.code2cloud.v1.argocdkubernetes.service.ArgocdKubernetesQueryController/getById"
)

// ArgocdKubernetesQueryControllerClient is the client API for ArgocdKubernetesQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArgocdKubernetesQueryControllerClient interface {
	// look up argocd-kubernetes using argocd-kubernetes id
	GetById(ctx context.Context, in *model.ArgocdKubernetesId, opts ...grpc.CallOption) (*model.ArgocdKubernetes, error)
}

type argocdKubernetesQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewArgocdKubernetesQueryControllerClient(cc grpc.ClientConnInterface) ArgocdKubernetesQueryControllerClient {
	return &argocdKubernetesQueryControllerClient{cc}
}

func (c *argocdKubernetesQueryControllerClient) GetById(ctx context.Context, in *model.ArgocdKubernetesId, opts ...grpc.CallOption) (*model.ArgocdKubernetes, error) {
	out := new(model.ArgocdKubernetes)
	err := c.cc.Invoke(ctx, ArgocdKubernetesQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArgocdKubernetesQueryControllerServer is the server API for ArgocdKubernetesQueryController service.
// All implementations should embed UnimplementedArgocdKubernetesQueryControllerServer
// for forward compatibility
type ArgocdKubernetesQueryControllerServer interface {
	// look up argocd-kubernetes using argocd-kubernetes id
	GetById(context.Context, *model.ArgocdKubernetesId) (*model.ArgocdKubernetes, error)
}

// UnimplementedArgocdKubernetesQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedArgocdKubernetesQueryControllerServer struct {
}

func (UnimplementedArgocdKubernetesQueryControllerServer) GetById(context.Context, *model.ArgocdKubernetesId) (*model.ArgocdKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}

// UnsafeArgocdKubernetesQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArgocdKubernetesQueryControllerServer will
// result in compilation errors.
type UnsafeArgocdKubernetesQueryControllerServer interface {
	mustEmbedUnimplementedArgocdKubernetesQueryControllerServer()
}

func RegisterArgocdKubernetesQueryControllerServer(s grpc.ServiceRegistrar, srv ArgocdKubernetesQueryControllerServer) {
	s.RegisterService(&ArgocdKubernetesQueryController_ServiceDesc, srv)
}

func _ArgocdKubernetesQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ArgocdKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArgocdKubernetesQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArgocdKubernetesQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArgocdKubernetesQueryControllerServer).GetById(ctx, req.(*model.ArgocdKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

// ArgocdKubernetesQueryController_ServiceDesc is the grpc.ServiceDesc for ArgocdKubernetesQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArgocdKubernetesQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.argocdkubernetes.service.ArgocdKubernetesQueryController",
	HandlerType: (*ArgocdKubernetesQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _ArgocdKubernetesQueryController_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/argocdkubernetes/service/query.proto",
}