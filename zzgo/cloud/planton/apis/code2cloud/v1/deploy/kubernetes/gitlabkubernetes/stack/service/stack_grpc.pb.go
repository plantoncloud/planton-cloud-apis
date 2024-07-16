// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/kubernetes/gitlabkubernetes/stack/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/kubernetes/gitlabkubernetes/stack/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GitlabKubernetesStackController_Execute_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.gitlabkubernetes.stack.service.GitlabKubernetesStackController/execute"
)

// GitlabKubernetesStackControllerClient is the client API for GitlabKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitlabKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.GitlabKubernetesStackInput, opts ...grpc.CallOption) (GitlabKubernetesStackController_ExecuteClient, error)
}

type gitlabKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGitlabKubernetesStackControllerClient(cc grpc.ClientConnInterface) GitlabKubernetesStackControllerClient {
	return &gitlabKubernetesStackControllerClient{cc}
}

func (c *gitlabKubernetesStackControllerClient) Execute(ctx context.Context, in *model.GitlabKubernetesStackInput, opts ...grpc.CallOption) (GitlabKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &GitlabKubernetesStackController_ServiceDesc.Streams[0], GitlabKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gitlabKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GitlabKubernetesStackController_ExecuteClient interface {
	Recv() (*model.GitlabKubernetesStackResponse, error)
	grpc.ClientStream
}

type gitlabKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *gitlabKubernetesStackControllerExecuteClient) Recv() (*model.GitlabKubernetesStackResponse, error) {
	m := new(model.GitlabKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GitlabKubernetesStackControllerServer is the server API for GitlabKubernetesStackController service.
// All implementations should embed UnimplementedGitlabKubernetesStackControllerServer
// for forward compatibility
type GitlabKubernetesStackControllerServer interface {
	Execute(*model.GitlabKubernetesStackInput, GitlabKubernetesStackController_ExecuteServer) error
}

// UnimplementedGitlabKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGitlabKubernetesStackControllerServer struct {
}

func (UnimplementedGitlabKubernetesStackControllerServer) Execute(*model.GitlabKubernetesStackInput, GitlabKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeGitlabKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitlabKubernetesStackControllerServer will
// result in compilation errors.
type UnsafeGitlabKubernetesStackControllerServer interface {
	mustEmbedUnimplementedGitlabKubernetesStackControllerServer()
}

func RegisterGitlabKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv GitlabKubernetesStackControllerServer) {
	s.RegisterService(&GitlabKubernetesStackController_ServiceDesc, srv)
}

func _GitlabKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.GitlabKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GitlabKubernetesStackControllerServer).Execute(m, &gitlabKubernetesStackControllerExecuteServer{stream})
}

type GitlabKubernetesStackController_ExecuteServer interface {
	Send(*model.GitlabKubernetesStackResponse) error
	grpc.ServerStream
}

type gitlabKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *gitlabKubernetesStackControllerExecuteServer) Send(m *model.GitlabKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GitlabKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for GitlabKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitlabKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.kubernetes.gitlabkubernetes.stack.service.GitlabKubernetesStackController",
	HandlerType: (*GitlabKubernetesStackControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _GitlabKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/kubernetes/gitlabkubernetes/stack/service/stack.proto",
}
