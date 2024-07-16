// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/kubernetes/argocdkubernetes/stack/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/kubernetes/argocdkubernetes/stack/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ArgocdKubernetesStackController_Execute_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.argocdkubernetes.stack.service.ArgocdKubernetesStackController/execute"
)

// ArgocdKubernetesStackControllerClient is the client API for ArgocdKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArgocdKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.ArgocdKubernetesStackInput, opts ...grpc.CallOption) (ArgocdKubernetesStackController_ExecuteClient, error)
}

type argocdKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewArgocdKubernetesStackControllerClient(cc grpc.ClientConnInterface) ArgocdKubernetesStackControllerClient {
	return &argocdKubernetesStackControllerClient{cc}
}

func (c *argocdKubernetesStackControllerClient) Execute(ctx context.Context, in *model.ArgocdKubernetesStackInput, opts ...grpc.CallOption) (ArgocdKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &ArgocdKubernetesStackController_ServiceDesc.Streams[0], ArgocdKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &argocdKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArgocdKubernetesStackController_ExecuteClient interface {
	Recv() (*model.ArgocdKubernetesStackResponse, error)
	grpc.ClientStream
}

type argocdKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *argocdKubernetesStackControllerExecuteClient) Recv() (*model.ArgocdKubernetesStackResponse, error) {
	m := new(model.ArgocdKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ArgocdKubernetesStackControllerServer is the server API for ArgocdKubernetesStackController service.
// All implementations should embed UnimplementedArgocdKubernetesStackControllerServer
// for forward compatibility
type ArgocdKubernetesStackControllerServer interface {
	Execute(*model.ArgocdKubernetesStackInput, ArgocdKubernetesStackController_ExecuteServer) error
}

// UnimplementedArgocdKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedArgocdKubernetesStackControllerServer struct {
}

func (UnimplementedArgocdKubernetesStackControllerServer) Execute(*model.ArgocdKubernetesStackInput, ArgocdKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeArgocdKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArgocdKubernetesStackControllerServer will
// result in compilation errors.
type UnsafeArgocdKubernetesStackControllerServer interface {
	mustEmbedUnimplementedArgocdKubernetesStackControllerServer()
}

func RegisterArgocdKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv ArgocdKubernetesStackControllerServer) {
	s.RegisterService(&ArgocdKubernetesStackController_ServiceDesc, srv)
}

func _ArgocdKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.ArgocdKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArgocdKubernetesStackControllerServer).Execute(m, &argocdKubernetesStackControllerExecuteServer{stream})
}

type ArgocdKubernetesStackController_ExecuteServer interface {
	Send(*model.ArgocdKubernetesStackResponse) error
	grpc.ServerStream
}

type argocdKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *argocdKubernetesStackControllerExecuteServer) Send(m *model.ArgocdKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ArgocdKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for ArgocdKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArgocdKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.kubernetes.argocdkubernetes.stack.service.ArgocdKubernetesStackController",
	HandlerType: (*ArgocdKubernetesStackControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _ArgocdKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/kubernetes/argocdkubernetes/stack/service/stack.proto",
}
