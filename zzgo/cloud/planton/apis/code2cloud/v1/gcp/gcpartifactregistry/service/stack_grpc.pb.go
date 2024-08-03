// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/gcp/gcpartifactregistry/service/stack.proto

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
	GcpArtifactRegistryStackController_Execute_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryStackController/execute"
)

// GcpArtifactRegistryStackControllerClient is the client API for GcpArtifactRegistryStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GcpArtifactRegistryStackControllerClient interface {
	Execute(ctx context.Context, in *model.GcpArtifactRegistryStackInput, opts ...grpc.CallOption) (GcpArtifactRegistryStackController_ExecuteClient, error)
}

type gcpArtifactRegistryStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGcpArtifactRegistryStackControllerClient(cc grpc.ClientConnInterface) GcpArtifactRegistryStackControllerClient {
	return &gcpArtifactRegistryStackControllerClient{cc}
}

func (c *gcpArtifactRegistryStackControllerClient) Execute(ctx context.Context, in *model.GcpArtifactRegistryStackInput, opts ...grpc.CallOption) (GcpArtifactRegistryStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &GcpArtifactRegistryStackController_ServiceDesc.Streams[0], GcpArtifactRegistryStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gcpArtifactRegistryStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GcpArtifactRegistryStackController_ExecuteClient interface {
	Recv() (*model.GcpArtifactRegistryStackResponse, error)
	grpc.ClientStream
}

type gcpArtifactRegistryStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *gcpArtifactRegistryStackControllerExecuteClient) Recv() (*model.GcpArtifactRegistryStackResponse, error) {
	m := new(model.GcpArtifactRegistryStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GcpArtifactRegistryStackControllerServer is the server API for GcpArtifactRegistryStackController service.
// All implementations should embed UnimplementedGcpArtifactRegistryStackControllerServer
// for forward compatibility
type GcpArtifactRegistryStackControllerServer interface {
	Execute(*model.GcpArtifactRegistryStackInput, GcpArtifactRegistryStackController_ExecuteServer) error
}

// UnimplementedGcpArtifactRegistryStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGcpArtifactRegistryStackControllerServer struct {
}

func (UnimplementedGcpArtifactRegistryStackControllerServer) Execute(*model.GcpArtifactRegistryStackInput, GcpArtifactRegistryStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeGcpArtifactRegistryStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GcpArtifactRegistryStackControllerServer will
// result in compilation errors.
type UnsafeGcpArtifactRegistryStackControllerServer interface {
	mustEmbedUnimplementedGcpArtifactRegistryStackControllerServer()
}

func RegisterGcpArtifactRegistryStackControllerServer(s grpc.ServiceRegistrar, srv GcpArtifactRegistryStackControllerServer) {
	s.RegisterService(&GcpArtifactRegistryStackController_ServiceDesc, srv)
}

func _GcpArtifactRegistryStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.GcpArtifactRegistryStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GcpArtifactRegistryStackControllerServer).Execute(m, &gcpArtifactRegistryStackControllerExecuteServer{stream})
}

type GcpArtifactRegistryStackController_ExecuteServer interface {
	Send(*model.GcpArtifactRegistryStackResponse) error
	grpc.ServerStream
}

type gcpArtifactRegistryStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *gcpArtifactRegistryStackControllerExecuteServer) Send(m *model.GcpArtifactRegistryStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GcpArtifactRegistryStackController_ServiceDesc is the grpc.ServiceDesc for GcpArtifactRegistryStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GcpArtifactRegistryStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryStackController",
	HandlerType: (*GcpArtifactRegistryStackControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _GcpArtifactRegistryStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/gcp/gcpartifactregistry/service/stack.proto",
}
