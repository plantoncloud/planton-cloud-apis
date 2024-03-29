// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/gitlabserver/stack/kubernetes/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gitlabserver/stack/kubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GitlabServerKubernetesStackController_Execute_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gitlabserver.stack.kubernetes.service.GitlabServerKubernetesStackController/execute"
	GitlabServerKubernetesStackController_GetStackOutputs_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gitlabserver.stack.kubernetes.service.GitlabServerKubernetesStackController/getStackOutputs"
)

// GitlabServerKubernetesStackControllerClient is the client API for GitlabServerKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitlabServerKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.GitlabServerKubernetesStackInput, opts ...grpc.CallOption) (GitlabServerKubernetesStackController_ExecuteClient, error)
	GetStackOutputs(ctx context.Context, in *model.GitlabServerKubernetesStackInput, opts ...grpc.CallOption) (*model.GitlabServerKubernetesStackOutputs, error)
}

type gitlabServerKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGitlabServerKubernetesStackControllerClient(cc grpc.ClientConnInterface) GitlabServerKubernetesStackControllerClient {
	return &gitlabServerKubernetesStackControllerClient{cc}
}

func (c *gitlabServerKubernetesStackControllerClient) Execute(ctx context.Context, in *model.GitlabServerKubernetesStackInput, opts ...grpc.CallOption) (GitlabServerKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &GitlabServerKubernetesStackController_ServiceDesc.Streams[0], GitlabServerKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gitlabServerKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GitlabServerKubernetesStackController_ExecuteClient interface {
	Recv() (*model.GitlabServerKubernetesStackResponse, error)
	grpc.ClientStream
}

type gitlabServerKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *gitlabServerKubernetesStackControllerExecuteClient) Recv() (*model.GitlabServerKubernetesStackResponse, error) {
	m := new(model.GitlabServerKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gitlabServerKubernetesStackControllerClient) GetStackOutputs(ctx context.Context, in *model.GitlabServerKubernetesStackInput, opts ...grpc.CallOption) (*model.GitlabServerKubernetesStackOutputs, error) {
	out := new(model.GitlabServerKubernetesStackOutputs)
	err := c.cc.Invoke(ctx, GitlabServerKubernetesStackController_GetStackOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitlabServerKubernetesStackControllerServer is the server API for GitlabServerKubernetesStackController service.
// All implementations should embed UnimplementedGitlabServerKubernetesStackControllerServer
// for forward compatibility
type GitlabServerKubernetesStackControllerServer interface {
	Execute(*model.GitlabServerKubernetesStackInput, GitlabServerKubernetesStackController_ExecuteServer) error
	GetStackOutputs(context.Context, *model.GitlabServerKubernetesStackInput) (*model.GitlabServerKubernetesStackOutputs, error)
}

// UnimplementedGitlabServerKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGitlabServerKubernetesStackControllerServer struct {
}

func (UnimplementedGitlabServerKubernetesStackControllerServer) Execute(*model.GitlabServerKubernetesStackInput, GitlabServerKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedGitlabServerKubernetesStackControllerServer) GetStackOutputs(context.Context, *model.GitlabServerKubernetesStackInput) (*model.GitlabServerKubernetesStackOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackOutputs not implemented")
}

// UnsafeGitlabServerKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitlabServerKubernetesStackControllerServer will
// result in compilation errors.
type UnsafeGitlabServerKubernetesStackControllerServer interface {
	mustEmbedUnimplementedGitlabServerKubernetesStackControllerServer()
}

func RegisterGitlabServerKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv GitlabServerKubernetesStackControllerServer) {
	s.RegisterService(&GitlabServerKubernetesStackController_ServiceDesc, srv)
}

func _GitlabServerKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.GitlabServerKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GitlabServerKubernetesStackControllerServer).Execute(m, &gitlabServerKubernetesStackControllerExecuteServer{stream})
}

type GitlabServerKubernetesStackController_ExecuteServer interface {
	Send(*model.GitlabServerKubernetesStackResponse) error
	grpc.ServerStream
}

type gitlabServerKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *gitlabServerKubernetesStackControllerExecuteServer) Send(m *model.GitlabServerKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GitlabServerKubernetesStackController_GetStackOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabServerKubernetesStackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServerKubernetesStackControllerServer).GetStackOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabServerKubernetesStackController_GetStackOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServerKubernetesStackControllerServer).GetStackOutputs(ctx, req.(*model.GitlabServerKubernetesStackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GitlabServerKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for GitlabServerKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitlabServerKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.gitlabserver.stack.kubernetes.service.GitlabServerKubernetesStackController",
	HandlerType: (*GitlabServerKubernetesStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStackOutputs",
			Handler:    _GitlabServerKubernetesStackController_GetStackOutputs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _GitlabServerKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/gitlabserver/stack/kubernetes/service/stack.proto",
}
