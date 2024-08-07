// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/kubernetes/locustkubernetes/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/locustkubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	LocustKubernetesQueryController_Get_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kubernetes.locustkubernetes.service.LocustKubernetesQueryController/get"
)

// LocustKubernetesQueryControllerClient is the client API for LocustKubernetesQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocustKubernetesQueryControllerClient interface {
	// look up locust-kubernetes using locust-kubernetes id
	Get(ctx context.Context, in *model.LocustKubernetesId, opts ...grpc.CallOption) (*model.LocustKubernetes, error)
}

type locustKubernetesQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewLocustKubernetesQueryControllerClient(cc grpc.ClientConnInterface) LocustKubernetesQueryControllerClient {
	return &locustKubernetesQueryControllerClient{cc}
}

func (c *locustKubernetesQueryControllerClient) Get(ctx context.Context, in *model.LocustKubernetesId, opts ...grpc.CallOption) (*model.LocustKubernetes, error) {
	out := new(model.LocustKubernetes)
	err := c.cc.Invoke(ctx, LocustKubernetesQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocustKubernetesQueryControllerServer is the server API for LocustKubernetesQueryController service.
// All implementations should embed UnimplementedLocustKubernetesQueryControllerServer
// for forward compatibility
type LocustKubernetesQueryControllerServer interface {
	// look up locust-kubernetes using locust-kubernetes id
	Get(context.Context, *model.LocustKubernetesId) (*model.LocustKubernetes, error)
}

// UnimplementedLocustKubernetesQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedLocustKubernetesQueryControllerServer struct {
}

func (UnimplementedLocustKubernetesQueryControllerServer) Get(context.Context, *model.LocustKubernetesId) (*model.LocustKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeLocustKubernetesQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocustKubernetesQueryControllerServer will
// result in compilation errors.
type UnsafeLocustKubernetesQueryControllerServer interface {
	mustEmbedUnimplementedLocustKubernetesQueryControllerServer()
}

func RegisterLocustKubernetesQueryControllerServer(s grpc.ServiceRegistrar, srv LocustKubernetesQueryControllerServer) {
	s.RegisterService(&LocustKubernetesQueryController_ServiceDesc, srv)
}

func _LocustKubernetesQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.LocustKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocustKubernetesQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocustKubernetesQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocustKubernetesQueryControllerServer).Get(ctx, req.(*model.LocustKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

// LocustKubernetesQueryController_ServiceDesc is the grpc.ServiceDesc for LocustKubernetesQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocustKubernetesQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.kubernetes.locustkubernetes.service.LocustKubernetesQueryController",
	HandlerType: (*LocustKubernetesQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _LocustKubernetesQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/kubernetes/locustkubernetes/service/query.proto",
}
