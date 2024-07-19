// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/gcp/gcpartifactregistry/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpartifactregistry/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GcpArtifactRegistryQueryController_Get_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryQueryController/get"
)

// GcpArtifactRegistryQueryControllerClient is the client API for GcpArtifactRegistryQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GcpArtifactRegistryQueryControllerClient interface {
	// get gcp-artifact-registry by id
	Get(ctx context.Context, in *model.GcpArtifactRegistryId, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error)
}

type gcpArtifactRegistryQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGcpArtifactRegistryQueryControllerClient(cc grpc.ClientConnInterface) GcpArtifactRegistryQueryControllerClient {
	return &gcpArtifactRegistryQueryControllerClient{cc}
}

func (c *gcpArtifactRegistryQueryControllerClient) Get(ctx context.Context, in *model.GcpArtifactRegistryId, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error) {
	out := new(model.GcpArtifactRegistry)
	err := c.cc.Invoke(ctx, GcpArtifactRegistryQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GcpArtifactRegistryQueryControllerServer is the server API for GcpArtifactRegistryQueryController service.
// All implementations should embed UnimplementedGcpArtifactRegistryQueryControllerServer
// for forward compatibility
type GcpArtifactRegistryQueryControllerServer interface {
	// get gcp-artifact-registry by id
	Get(context.Context, *model.GcpArtifactRegistryId) (*model.GcpArtifactRegistry, error)
}

// UnimplementedGcpArtifactRegistryQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGcpArtifactRegistryQueryControllerServer struct {
}

func (UnimplementedGcpArtifactRegistryQueryControllerServer) Get(context.Context, *model.GcpArtifactRegistryId) (*model.GcpArtifactRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeGcpArtifactRegistryQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GcpArtifactRegistryQueryControllerServer will
// result in compilation errors.
type UnsafeGcpArtifactRegistryQueryControllerServer interface {
	mustEmbedUnimplementedGcpArtifactRegistryQueryControllerServer()
}

func RegisterGcpArtifactRegistryQueryControllerServer(s grpc.ServiceRegistrar, srv GcpArtifactRegistryQueryControllerServer) {
	s.RegisterService(&GcpArtifactRegistryQueryController_ServiceDesc, srv)
}

func _GcpArtifactRegistryQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpArtifactRegistryId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpArtifactRegistryQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpArtifactRegistryQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpArtifactRegistryQueryControllerServer).Get(ctx, req.(*model.GcpArtifactRegistryId))
	}
	return interceptor(ctx, in, info, handler)
}

// GcpArtifactRegistryQueryController_ServiceDesc is the grpc.ServiceDesc for GcpArtifactRegistryQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GcpArtifactRegistryQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryQueryController",
	HandlerType: (*GcpArtifactRegistryQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _GcpArtifactRegistryQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/gcp/gcpartifactregistry/service/query.proto",
}