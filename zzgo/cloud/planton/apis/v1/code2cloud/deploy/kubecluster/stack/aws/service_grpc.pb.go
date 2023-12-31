// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/kubecluster/stack/aws/service.proto

package aws

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KubeClusterAwsStackController_Execute_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.KubeClusterAwsStackController/execute"
	KubeClusterAwsStackController_GetStackOutputs_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.KubeClusterAwsStackController/getStackOutputs"
)

// KubeClusterAwsStackControllerClient is the client API for KubeClusterAwsStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubeClusterAwsStackControllerClient interface {
	Execute(ctx context.Context, in *KubeClusterAwsStackInput, opts ...grpc.CallOption) (KubeClusterAwsStackController_ExecuteClient, error)
	GetStackOutputs(ctx context.Context, in *KubeClusterAwsStackInput, opts ...grpc.CallOption) (*KubeClusterAwsStackOutputs, error)
}

type kubeClusterAwsStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKubeClusterAwsStackControllerClient(cc grpc.ClientConnInterface) KubeClusterAwsStackControllerClient {
	return &kubeClusterAwsStackControllerClient{cc}
}

func (c *kubeClusterAwsStackControllerClient) Execute(ctx context.Context, in *KubeClusterAwsStackInput, opts ...grpc.CallOption) (KubeClusterAwsStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &KubeClusterAwsStackController_ServiceDesc.Streams[0], KubeClusterAwsStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &kubeClusterAwsStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KubeClusterAwsStackController_ExecuteClient interface {
	Recv() (*KubeClusterAwsStackResponse, error)
	grpc.ClientStream
}

type kubeClusterAwsStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *kubeClusterAwsStackControllerExecuteClient) Recv() (*KubeClusterAwsStackResponse, error) {
	m := new(KubeClusterAwsStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kubeClusterAwsStackControllerClient) GetStackOutputs(ctx context.Context, in *KubeClusterAwsStackInput, opts ...grpc.CallOption) (*KubeClusterAwsStackOutputs, error) {
	out := new(KubeClusterAwsStackOutputs)
	err := c.cc.Invoke(ctx, KubeClusterAwsStackController_GetStackOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubeClusterAwsStackControllerServer is the server API for KubeClusterAwsStackController service.
// All implementations should embed UnimplementedKubeClusterAwsStackControllerServer
// for forward compatibility
type KubeClusterAwsStackControllerServer interface {
	Execute(*KubeClusterAwsStackInput, KubeClusterAwsStackController_ExecuteServer) error
	GetStackOutputs(context.Context, *KubeClusterAwsStackInput) (*KubeClusterAwsStackOutputs, error)
}

// UnimplementedKubeClusterAwsStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKubeClusterAwsStackControllerServer struct {
}

func (UnimplementedKubeClusterAwsStackControllerServer) Execute(*KubeClusterAwsStackInput, KubeClusterAwsStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedKubeClusterAwsStackControllerServer) GetStackOutputs(context.Context, *KubeClusterAwsStackInput) (*KubeClusterAwsStackOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackOutputs not implemented")
}

// UnsafeKubeClusterAwsStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubeClusterAwsStackControllerServer will
// result in compilation errors.
type UnsafeKubeClusterAwsStackControllerServer interface {
	mustEmbedUnimplementedKubeClusterAwsStackControllerServer()
}

func RegisterKubeClusterAwsStackControllerServer(s grpc.ServiceRegistrar, srv KubeClusterAwsStackControllerServer) {
	s.RegisterService(&KubeClusterAwsStackController_ServiceDesc, srv)
}

func _KubeClusterAwsStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KubeClusterAwsStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KubeClusterAwsStackControllerServer).Execute(m, &kubeClusterAwsStackControllerExecuteServer{stream})
}

type KubeClusterAwsStackController_ExecuteServer interface {
	Send(*KubeClusterAwsStackResponse) error
	grpc.ServerStream
}

type kubeClusterAwsStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *kubeClusterAwsStackControllerExecuteServer) Send(m *KubeClusterAwsStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KubeClusterAwsStackController_GetStackOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeClusterAwsStackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterAwsStackControllerServer).GetStackOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterAwsStackController_GetStackOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterAwsStackControllerServer).GetStackOutputs(ctx, req.(*KubeClusterAwsStackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// KubeClusterAwsStackController_ServiceDesc is the grpc.ServiceDesc for KubeClusterAwsStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubeClusterAwsStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.KubeClusterAwsStackController",
	HandlerType: (*KubeClusterAwsStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStackOutputs",
			Handler:    _KubeClusterAwsStackController_GetStackOutputs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _KubeClusterAwsStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/kubecluster/stack/aws/service.proto",
}
