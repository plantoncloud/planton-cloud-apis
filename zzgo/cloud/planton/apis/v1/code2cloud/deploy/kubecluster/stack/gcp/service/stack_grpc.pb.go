// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/kubecluster/stack/gcp/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/kubecluster/stack/gcp/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KubeClusterGcpStackController_Execute_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.gcp.service.KubeClusterGcpStackController/execute"
	KubeClusterGcpStackController_GetStackOutputs_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.gcp.service.KubeClusterGcpStackController/getStackOutputs"
)

// KubeClusterGcpStackControllerClient is the client API for KubeClusterGcpStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubeClusterGcpStackControllerClient interface {
	Execute(ctx context.Context, in *model.KubeClusterGcpStackInput, opts ...grpc.CallOption) (KubeClusterGcpStackController_ExecuteClient, error)
	GetStackOutputs(ctx context.Context, in *model.KubeClusterGcpStackInput, opts ...grpc.CallOption) (*model.KubeClusterGcpStackOutputs, error)
}

type kubeClusterGcpStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKubeClusterGcpStackControllerClient(cc grpc.ClientConnInterface) KubeClusterGcpStackControllerClient {
	return &kubeClusterGcpStackControllerClient{cc}
}

func (c *kubeClusterGcpStackControllerClient) Execute(ctx context.Context, in *model.KubeClusterGcpStackInput, opts ...grpc.CallOption) (KubeClusterGcpStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &KubeClusterGcpStackController_ServiceDesc.Streams[0], KubeClusterGcpStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &kubeClusterGcpStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KubeClusterGcpStackController_ExecuteClient interface {
	Recv() (*model.KubeClusterGcpStackResponse, error)
	grpc.ClientStream
}

type kubeClusterGcpStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *kubeClusterGcpStackControllerExecuteClient) Recv() (*model.KubeClusterGcpStackResponse, error) {
	m := new(model.KubeClusterGcpStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kubeClusterGcpStackControllerClient) GetStackOutputs(ctx context.Context, in *model.KubeClusterGcpStackInput, opts ...grpc.CallOption) (*model.KubeClusterGcpStackOutputs, error) {
	out := new(model.KubeClusterGcpStackOutputs)
	err := c.cc.Invoke(ctx, KubeClusterGcpStackController_GetStackOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubeClusterGcpStackControllerServer is the server API for KubeClusterGcpStackController service.
// All implementations should embed UnimplementedKubeClusterGcpStackControllerServer
// for forward compatibility
type KubeClusterGcpStackControllerServer interface {
	Execute(*model.KubeClusterGcpStackInput, KubeClusterGcpStackController_ExecuteServer) error
	GetStackOutputs(context.Context, *model.KubeClusterGcpStackInput) (*model.KubeClusterGcpStackOutputs, error)
}

// UnimplementedKubeClusterGcpStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKubeClusterGcpStackControllerServer struct {
}

func (UnimplementedKubeClusterGcpStackControllerServer) Execute(*model.KubeClusterGcpStackInput, KubeClusterGcpStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedKubeClusterGcpStackControllerServer) GetStackOutputs(context.Context, *model.KubeClusterGcpStackInput) (*model.KubeClusterGcpStackOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackOutputs not implemented")
}

// UnsafeKubeClusterGcpStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubeClusterGcpStackControllerServer will
// result in compilation errors.
type UnsafeKubeClusterGcpStackControllerServer interface {
	mustEmbedUnimplementedKubeClusterGcpStackControllerServer()
}

func RegisterKubeClusterGcpStackControllerServer(s grpc.ServiceRegistrar, srv KubeClusterGcpStackControllerServer) {
	s.RegisterService(&KubeClusterGcpStackController_ServiceDesc, srv)
}

func _KubeClusterGcpStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.KubeClusterGcpStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KubeClusterGcpStackControllerServer).Execute(m, &kubeClusterGcpStackControllerExecuteServer{stream})
}

type KubeClusterGcpStackController_ExecuteServer interface {
	Send(*model.KubeClusterGcpStackResponse) error
	grpc.ServerStream
}

type kubeClusterGcpStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *kubeClusterGcpStackControllerExecuteServer) Send(m *model.KubeClusterGcpStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KubeClusterGcpStackController_GetStackOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KubeClusterGcpStackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterGcpStackControllerServer).GetStackOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterGcpStackController_GetStackOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterGcpStackControllerServer).GetStackOutputs(ctx, req.(*model.KubeClusterGcpStackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// KubeClusterGcpStackController_ServiceDesc is the grpc.ServiceDesc for KubeClusterGcpStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubeClusterGcpStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.gcp.service.KubeClusterGcpStackController",
	HandlerType: (*KubeClusterGcpStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStackOutputs",
			Handler:    _KubeClusterGcpStackController_GetStackOutputs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _KubeClusterGcpStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/kubecluster/stack/gcp/service/stack.proto",
}
