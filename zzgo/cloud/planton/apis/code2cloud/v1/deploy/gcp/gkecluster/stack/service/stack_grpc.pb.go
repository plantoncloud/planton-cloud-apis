// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/gcp/gkecluster/stack/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/gcp/gkecluster/stack/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GkeClusterGcpStackController_Execute_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.gcp.gkecluster.stack.service.GkeClusterGcpStackController/execute"
)

// GkeClusterGcpStackControllerClient is the client API for GkeClusterGcpStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GkeClusterGcpStackControllerClient interface {
	Execute(ctx context.Context, in *model.GkeClusterStackInput, opts ...grpc.CallOption) (GkeClusterGcpStackController_ExecuteClient, error)
}

type gkeClusterGcpStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGkeClusterGcpStackControllerClient(cc grpc.ClientConnInterface) GkeClusterGcpStackControllerClient {
	return &gkeClusterGcpStackControllerClient{cc}
}

func (c *gkeClusterGcpStackControllerClient) Execute(ctx context.Context, in *model.GkeClusterStackInput, opts ...grpc.CallOption) (GkeClusterGcpStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &GkeClusterGcpStackController_ServiceDesc.Streams[0], GkeClusterGcpStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gkeClusterGcpStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GkeClusterGcpStackController_ExecuteClient interface {
	Recv() (*model.GkeClusterStackResponse, error)
	grpc.ClientStream
}

type gkeClusterGcpStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *gkeClusterGcpStackControllerExecuteClient) Recv() (*model.GkeClusterStackResponse, error) {
	m := new(model.GkeClusterStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GkeClusterGcpStackControllerServer is the server API for GkeClusterGcpStackController service.
// All implementations should embed UnimplementedGkeClusterGcpStackControllerServer
// for forward compatibility
type GkeClusterGcpStackControllerServer interface {
	Execute(*model.GkeClusterStackInput, GkeClusterGcpStackController_ExecuteServer) error
}

// UnimplementedGkeClusterGcpStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGkeClusterGcpStackControllerServer struct {
}

func (UnimplementedGkeClusterGcpStackControllerServer) Execute(*model.GkeClusterStackInput, GkeClusterGcpStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeGkeClusterGcpStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GkeClusterGcpStackControllerServer will
// result in compilation errors.
type UnsafeGkeClusterGcpStackControllerServer interface {
	mustEmbedUnimplementedGkeClusterGcpStackControllerServer()
}

func RegisterGkeClusterGcpStackControllerServer(s grpc.ServiceRegistrar, srv GkeClusterGcpStackControllerServer) {
	s.RegisterService(&GkeClusterGcpStackController_ServiceDesc, srv)
}

func _GkeClusterGcpStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.GkeClusterStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GkeClusterGcpStackControllerServer).Execute(m, &gkeClusterGcpStackControllerExecuteServer{stream})
}

type GkeClusterGcpStackController_ExecuteServer interface {
	Send(*model.GkeClusterStackResponse) error
	grpc.ServerStream
}

type gkeClusterGcpStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *gkeClusterGcpStackControllerExecuteServer) Send(m *model.GkeClusterStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GkeClusterGcpStackController_ServiceDesc is the grpc.ServiceDesc for GkeClusterGcpStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GkeClusterGcpStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.gcp.gkecluster.stack.service.GkeClusterGcpStackController",
	HandlerType: (*GkeClusterGcpStackControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _GkeClusterGcpStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/gcp/gkecluster/stack/service/stack.proto",
}