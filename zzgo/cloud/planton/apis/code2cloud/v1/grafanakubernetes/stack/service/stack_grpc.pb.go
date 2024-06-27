// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/grafanakubernetes/stack/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/grafanakubernetes/stack/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GrafanaKubernetesStackController_Execute_FullMethodName = "/cloud.planton.apis.code2cloud.v1.grafanakubernetes.stack.service.GrafanaKubernetesStackController/execute"
)

// GrafanaKubernetesStackControllerClient is the client API for GrafanaKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrafanaKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.GrafanaKubernetesStackInput, opts ...grpc.CallOption) (GrafanaKubernetesStackController_ExecuteClient, error)
}

type grafanaKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGrafanaKubernetesStackControllerClient(cc grpc.ClientConnInterface) GrafanaKubernetesStackControllerClient {
	return &grafanaKubernetesStackControllerClient{cc}
}

func (c *grafanaKubernetesStackControllerClient) Execute(ctx context.Context, in *model.GrafanaKubernetesStackInput, opts ...grpc.CallOption) (GrafanaKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &GrafanaKubernetesStackController_ServiceDesc.Streams[0], GrafanaKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &grafanaKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GrafanaKubernetesStackController_ExecuteClient interface {
	Recv() (*model.GrafanaKubernetesStackResponse, error)
	grpc.ClientStream
}

type grafanaKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *grafanaKubernetesStackControllerExecuteClient) Recv() (*model.GrafanaKubernetesStackResponse, error) {
	m := new(model.GrafanaKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GrafanaKubernetesStackControllerServer is the server API for GrafanaKubernetesStackController service.
// All implementations should embed UnimplementedGrafanaKubernetesStackControllerServer
// for forward compatibility
type GrafanaKubernetesStackControllerServer interface {
	Execute(*model.GrafanaKubernetesStackInput, GrafanaKubernetesStackController_ExecuteServer) error
}

// UnimplementedGrafanaKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGrafanaKubernetesStackControllerServer struct {
}

func (UnimplementedGrafanaKubernetesStackControllerServer) Execute(*model.GrafanaKubernetesStackInput, GrafanaKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeGrafanaKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrafanaKubernetesStackControllerServer will
// result in compilation errors.
type UnsafeGrafanaKubernetesStackControllerServer interface {
	mustEmbedUnimplementedGrafanaKubernetesStackControllerServer()
}

func RegisterGrafanaKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv GrafanaKubernetesStackControllerServer) {
	s.RegisterService(&GrafanaKubernetesStackController_ServiceDesc, srv)
}

func _GrafanaKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.GrafanaKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrafanaKubernetesStackControllerServer).Execute(m, &grafanaKubernetesStackControllerExecuteServer{stream})
}

type GrafanaKubernetesStackController_ExecuteServer interface {
	Send(*model.GrafanaKubernetesStackResponse) error
	grpc.ServerStream
}

type grafanaKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *grafanaKubernetesStackControllerExecuteServer) Send(m *model.GrafanaKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GrafanaKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for GrafanaKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GrafanaKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.grafanakubernetes.stack.service.GrafanaKubernetesStackController",
	HandlerType: (*GrafanaKubernetesStackControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _GrafanaKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/grafanakubernetes/stack/service/stack.proto",
}
