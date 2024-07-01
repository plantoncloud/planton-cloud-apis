// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/keycloakkubernetes/stack/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/keycloakkubernetes/stack/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KeycloakKubernetesStackController_Execute_FullMethodName = "/cloud.planton.apis.code2cloud.v1.keycloakkubernetes.stack.service.KeycloakKubernetesStackController/execute"
)

// KeycloakKubernetesStackControllerClient is the client API for KeycloakKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeycloakKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.KeycloakKubernetesStackInput, opts ...grpc.CallOption) (KeycloakKubernetesStackController_ExecuteClient, error)
}

type keycloakKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKeycloakKubernetesStackControllerClient(cc grpc.ClientConnInterface) KeycloakKubernetesStackControllerClient {
	return &keycloakKubernetesStackControllerClient{cc}
}

func (c *keycloakKubernetesStackControllerClient) Execute(ctx context.Context, in *model.KeycloakKubernetesStackInput, opts ...grpc.CallOption) (KeycloakKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &KeycloakKubernetesStackController_ServiceDesc.Streams[0], KeycloakKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &keycloakKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KeycloakKubernetesStackController_ExecuteClient interface {
	Recv() (*model.KeycloakKubernetesStackResponse, error)
	grpc.ClientStream
}

type keycloakKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *keycloakKubernetesStackControllerExecuteClient) Recv() (*model.KeycloakKubernetesStackResponse, error) {
	m := new(model.KeycloakKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KeycloakKubernetesStackControllerServer is the server API for KeycloakKubernetesStackController service.
// All implementations should embed UnimplementedKeycloakKubernetesStackControllerServer
// for forward compatibility
type KeycloakKubernetesStackControllerServer interface {
	Execute(*model.KeycloakKubernetesStackInput, KeycloakKubernetesStackController_ExecuteServer) error
}

// UnimplementedKeycloakKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKeycloakKubernetesStackControllerServer struct {
}

func (UnimplementedKeycloakKubernetesStackControllerServer) Execute(*model.KeycloakKubernetesStackInput, KeycloakKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeKeycloakKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeycloakKubernetesStackControllerServer will
// result in compilation errors.
type UnsafeKeycloakKubernetesStackControllerServer interface {
	mustEmbedUnimplementedKeycloakKubernetesStackControllerServer()
}

func RegisterKeycloakKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv KeycloakKubernetesStackControllerServer) {
	s.RegisterService(&KeycloakKubernetesStackController_ServiceDesc, srv)
}

func _KeycloakKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.KeycloakKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KeycloakKubernetesStackControllerServer).Execute(m, &keycloakKubernetesStackControllerExecuteServer{stream})
}

type KeycloakKubernetesStackController_ExecuteServer interface {
	Send(*model.KeycloakKubernetesStackResponse) error
	grpc.ServerStream
}

type keycloakKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *keycloakKubernetesStackControllerExecuteServer) Send(m *model.KeycloakKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

// KeycloakKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for KeycloakKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeycloakKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.keycloakkubernetes.stack.service.KeycloakKubernetesStackController",
	HandlerType: (*KeycloakKubernetesStackControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _KeycloakKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/keycloakkubernetes/stack/service/stack.proto",
}
