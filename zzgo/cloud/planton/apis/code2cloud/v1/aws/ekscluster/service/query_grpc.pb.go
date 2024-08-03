// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/aws/ekscluster/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/ekscluster/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EksClusterQueryController_Get_FullMethodName = "/cloud.planton.apis.code2cloud.v1.aws.ekscluster.service.EksClusterQueryController/get"
)

// EksClusterQueryControllerClient is the client API for EksClusterQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EksClusterQueryControllerClient interface {
	// lookup eks-cluster using eks-cluster id
	Get(ctx context.Context, in *model.EksClusterId, opts ...grpc.CallOption) (*model.EksCluster, error)
}

type eksClusterQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewEksClusterQueryControllerClient(cc grpc.ClientConnInterface) EksClusterQueryControllerClient {
	return &eksClusterQueryControllerClient{cc}
}

func (c *eksClusterQueryControllerClient) Get(ctx context.Context, in *model.EksClusterId, opts ...grpc.CallOption) (*model.EksCluster, error) {
	out := new(model.EksCluster)
	err := c.cc.Invoke(ctx, EksClusterQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EksClusterQueryControllerServer is the server API for EksClusterQueryController service.
// All implementations should embed UnimplementedEksClusterQueryControllerServer
// for forward compatibility
type EksClusterQueryControllerServer interface {
	// lookup eks-cluster using eks-cluster id
	Get(context.Context, *model.EksClusterId) (*model.EksCluster, error)
}

// UnimplementedEksClusterQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedEksClusterQueryControllerServer struct {
}

func (UnimplementedEksClusterQueryControllerServer) Get(context.Context, *model.EksClusterId) (*model.EksCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeEksClusterQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EksClusterQueryControllerServer will
// result in compilation errors.
type UnsafeEksClusterQueryControllerServer interface {
	mustEmbedUnimplementedEksClusterQueryControllerServer()
}

func RegisterEksClusterQueryControllerServer(s grpc.ServiceRegistrar, srv EksClusterQueryControllerServer) {
	s.RegisterService(&EksClusterQueryController_ServiceDesc, srv)
}

func _EksClusterQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EksClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EksClusterQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EksClusterQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EksClusterQueryControllerServer).Get(ctx, req.(*model.EksClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

// EksClusterQueryController_ServiceDesc is the grpc.ServiceDesc for EksClusterQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EksClusterQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.aws.ekscluster.service.EksClusterQueryController",
	HandlerType: (*EksClusterQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _EksClusterQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/aws/ekscluster/service/query.proto",
}
