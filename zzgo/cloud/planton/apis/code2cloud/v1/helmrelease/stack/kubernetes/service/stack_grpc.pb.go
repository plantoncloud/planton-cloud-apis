// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/helmrelease/stack/kubernetes/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/helmrelease/stack/kubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	HelmReleaseKubernetesStackController_Execute_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.service.HelmReleaseKubernetesStackController/execute"
	HelmReleaseKubernetesStackController_GetStackOutputs_FullMethodName = "/cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.service.HelmReleaseKubernetesStackController/getStackOutputs"
)

// HelmReleaseKubernetesStackControllerClient is the client API for HelmReleaseKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelmReleaseKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.HelmReleaseKubernetesStackInput, opts ...grpc.CallOption) (HelmReleaseKubernetesStackController_ExecuteClient, error)
	GetStackOutputs(ctx context.Context, in *model.HelmReleaseKubernetesStackInput, opts ...grpc.CallOption) (*model.HelmReleaseKubernetesStackOutputs, error)
}

type helmReleaseKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewHelmReleaseKubernetesStackControllerClient(cc grpc.ClientConnInterface) HelmReleaseKubernetesStackControllerClient {
	return &helmReleaseKubernetesStackControllerClient{cc}
}

func (c *helmReleaseKubernetesStackControllerClient) Execute(ctx context.Context, in *model.HelmReleaseKubernetesStackInput, opts ...grpc.CallOption) (HelmReleaseKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelmReleaseKubernetesStackController_ServiceDesc.Streams[0], HelmReleaseKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helmReleaseKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelmReleaseKubernetesStackController_ExecuteClient interface {
	Recv() (*model.HelmReleaseKubernetesStackResponse, error)
	grpc.ClientStream
}

type helmReleaseKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *helmReleaseKubernetesStackControllerExecuteClient) Recv() (*model.HelmReleaseKubernetesStackResponse, error) {
	m := new(model.HelmReleaseKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helmReleaseKubernetesStackControllerClient) GetStackOutputs(ctx context.Context, in *model.HelmReleaseKubernetesStackInput, opts ...grpc.CallOption) (*model.HelmReleaseKubernetesStackOutputs, error) {
	out := new(model.HelmReleaseKubernetesStackOutputs)
	err := c.cc.Invoke(ctx, HelmReleaseKubernetesStackController_GetStackOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelmReleaseKubernetesStackControllerServer is the server API for HelmReleaseKubernetesStackController service.
// All implementations should embed UnimplementedHelmReleaseKubernetesStackControllerServer
// for forward compatibility
type HelmReleaseKubernetesStackControllerServer interface {
	Execute(*model.HelmReleaseKubernetesStackInput, HelmReleaseKubernetesStackController_ExecuteServer) error
	GetStackOutputs(context.Context, *model.HelmReleaseKubernetesStackInput) (*model.HelmReleaseKubernetesStackOutputs, error)
}

// UnimplementedHelmReleaseKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedHelmReleaseKubernetesStackControllerServer struct {
}

func (UnimplementedHelmReleaseKubernetesStackControllerServer) Execute(*model.HelmReleaseKubernetesStackInput, HelmReleaseKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedHelmReleaseKubernetesStackControllerServer) GetStackOutputs(context.Context, *model.HelmReleaseKubernetesStackInput) (*model.HelmReleaseKubernetesStackOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackOutputs not implemented")
}

// UnsafeHelmReleaseKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelmReleaseKubernetesStackControllerServer will
// result in compilation errors.
type UnsafeHelmReleaseKubernetesStackControllerServer interface {
	mustEmbedUnimplementedHelmReleaseKubernetesStackControllerServer()
}

func RegisterHelmReleaseKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv HelmReleaseKubernetesStackControllerServer) {
	s.RegisterService(&HelmReleaseKubernetesStackController_ServiceDesc, srv)
}

func _HelmReleaseKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.HelmReleaseKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelmReleaseKubernetesStackControllerServer).Execute(m, &helmReleaseKubernetesStackControllerExecuteServer{stream})
}

type HelmReleaseKubernetesStackController_ExecuteServer interface {
	Send(*model.HelmReleaseKubernetesStackResponse) error
	grpc.ServerStream
}

type helmReleaseKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *helmReleaseKubernetesStackControllerExecuteServer) Send(m *model.HelmReleaseKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelmReleaseKubernetesStackController_GetStackOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.HelmReleaseKubernetesStackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmReleaseKubernetesStackControllerServer).GetStackOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelmReleaseKubernetesStackController_GetStackOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmReleaseKubernetesStackControllerServer).GetStackOutputs(ctx, req.(*model.HelmReleaseKubernetesStackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// HelmReleaseKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for HelmReleaseKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelmReleaseKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.service.HelmReleaseKubernetesStackController",
	HandlerType: (*HelmReleaseKubernetesStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStackOutputs",
			Handler:    _HelmReleaseKubernetesStackController_GetStackOutputs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _HelmReleaseKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/helmrelease/stack/kubernetes/service/stack.proto",
}
