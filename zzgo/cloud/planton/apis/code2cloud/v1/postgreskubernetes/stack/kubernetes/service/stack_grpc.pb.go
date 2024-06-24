// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/postgreskubernetes/stack/kubernetes/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/postgreskubernetes/stack/kubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PostgresKubernetesStackController_Execute_FullMethodName = "/cloud.planton.apis.code2cloud.v1.postgreskubernetes.stack.kubernetes.service.PostgresKubernetesStackController/execute"
)

// PostgresKubernetesStackControllerClient is the client API for PostgresKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostgresKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.PostgresKubernetesStackInput, opts ...grpc.CallOption) (PostgresKubernetesStackController_ExecuteClient, error)
}

type postgresKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewPostgresKubernetesStackControllerClient(cc grpc.ClientConnInterface) PostgresKubernetesStackControllerClient {
	return &postgresKubernetesStackControllerClient{cc}
}

func (c *postgresKubernetesStackControllerClient) Execute(ctx context.Context, in *model.PostgresKubernetesStackInput, opts ...grpc.CallOption) (PostgresKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &PostgresKubernetesStackController_ServiceDesc.Streams[0], PostgresKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &postgresKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PostgresKubernetesStackController_ExecuteClient interface {
	Recv() (*model.PostgresKubernetesStackResponse, error)
	grpc.ClientStream
}

type postgresKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *postgresKubernetesStackControllerExecuteClient) Recv() (*model.PostgresKubernetesStackResponse, error) {
	m := new(model.PostgresKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PostgresKubernetesStackControllerServer is the server API for PostgresKubernetesStackController service.
// All implementations should embed UnimplementedPostgresKubernetesStackControllerServer
// for forward compatibility
type PostgresKubernetesStackControllerServer interface {
	Execute(*model.PostgresKubernetesStackInput, PostgresKubernetesStackController_ExecuteServer) error
}

// UnimplementedPostgresKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedPostgresKubernetesStackControllerServer struct {
}

func (UnimplementedPostgresKubernetesStackControllerServer) Execute(*model.PostgresKubernetesStackInput, PostgresKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafePostgresKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostgresKubernetesStackControllerServer will
// result in compilation errors.
type UnsafePostgresKubernetesStackControllerServer interface {
	mustEmbedUnimplementedPostgresKubernetesStackControllerServer()
}

func RegisterPostgresKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv PostgresKubernetesStackControllerServer) {
	s.RegisterService(&PostgresKubernetesStackController_ServiceDesc, srv)
}

func _PostgresKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.PostgresKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostgresKubernetesStackControllerServer).Execute(m, &postgresKubernetesStackControllerExecuteServer{stream})
}

type PostgresKubernetesStackController_ExecuteServer interface {
	Send(*model.PostgresKubernetesStackResponse) error
	grpc.ServerStream
}

type postgresKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *postgresKubernetesStackControllerExecuteServer) Send(m *model.PostgresKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PostgresKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for PostgresKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostgresKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.postgreskubernetes.stack.kubernetes.service.PostgresKubernetesStackController",
	HandlerType: (*PostgresKubernetesStackControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _PostgresKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/postgreskubernetes/stack/kubernetes/service/stack.proto",
}
