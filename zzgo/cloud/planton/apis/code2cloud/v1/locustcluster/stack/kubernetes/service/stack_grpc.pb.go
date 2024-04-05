// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/locustcluster/stack/kubernetes/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/locustcluster/stack/kubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	LocustClusterKubernetesStackController_Execute_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.locustcluster.stack.kubernetes.service.LocustClusterKubernetesStackController/execute"
	LocustClusterKubernetesStackController_GetStackOutputs_FullMethodName = "/cloud.planton.apis.code2cloud.v1.locustcluster.stack.kubernetes.service.LocustClusterKubernetesStackController/getStackOutputs"
)

// LocustClusterKubernetesStackControllerClient is the client API for LocustClusterKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocustClusterKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.LocustClusterKubernetesStackInput, opts ...grpc.CallOption) (LocustClusterKubernetesStackController_ExecuteClient, error)
	GetStackOutputs(ctx context.Context, in *model.LocustClusterKubernetesStackInput, opts ...grpc.CallOption) (*model.LocustClusterKubernetesStackOutputs, error)
}

type locustClusterKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewLocustClusterKubernetesStackControllerClient(cc grpc.ClientConnInterface) LocustClusterKubernetesStackControllerClient {
	return &locustClusterKubernetesStackControllerClient{cc}
}

func (c *locustClusterKubernetesStackControllerClient) Execute(ctx context.Context, in *model.LocustClusterKubernetesStackInput, opts ...grpc.CallOption) (LocustClusterKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &LocustClusterKubernetesStackController_ServiceDesc.Streams[0], LocustClusterKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &locustClusterKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LocustClusterKubernetesStackController_ExecuteClient interface {
	Recv() (*model.LocustClusterKubernetesStackResponse, error)
	grpc.ClientStream
}

type locustClusterKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *locustClusterKubernetesStackControllerExecuteClient) Recv() (*model.LocustClusterKubernetesStackResponse, error) {
	m := new(model.LocustClusterKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *locustClusterKubernetesStackControllerClient) GetStackOutputs(ctx context.Context, in *model.LocustClusterKubernetesStackInput, opts ...grpc.CallOption) (*model.LocustClusterKubernetesStackOutputs, error) {
	out := new(model.LocustClusterKubernetesStackOutputs)
	err := c.cc.Invoke(ctx, LocustClusterKubernetesStackController_GetStackOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocustClusterKubernetesStackControllerServer is the server API for LocustClusterKubernetesStackController service.
// All implementations should embed UnimplementedLocustClusterKubernetesStackControllerServer
// for forward compatibility
type LocustClusterKubernetesStackControllerServer interface {
	Execute(*model.LocustClusterKubernetesStackInput, LocustClusterKubernetesStackController_ExecuteServer) error
	GetStackOutputs(context.Context, *model.LocustClusterKubernetesStackInput) (*model.LocustClusterKubernetesStackOutputs, error)
}

// UnimplementedLocustClusterKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedLocustClusterKubernetesStackControllerServer struct {
}

func (UnimplementedLocustClusterKubernetesStackControllerServer) Execute(*model.LocustClusterKubernetesStackInput, LocustClusterKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedLocustClusterKubernetesStackControllerServer) GetStackOutputs(context.Context, *model.LocustClusterKubernetesStackInput) (*model.LocustClusterKubernetesStackOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackOutputs not implemented")
}

// UnsafeLocustClusterKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocustClusterKubernetesStackControllerServer will
// result in compilation errors.
type UnsafeLocustClusterKubernetesStackControllerServer interface {
	mustEmbedUnimplementedLocustClusterKubernetesStackControllerServer()
}

func RegisterLocustClusterKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv LocustClusterKubernetesStackControllerServer) {
	s.RegisterService(&LocustClusterKubernetesStackController_ServiceDesc, srv)
}

func _LocustClusterKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.LocustClusterKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LocustClusterKubernetesStackControllerServer).Execute(m, &locustClusterKubernetesStackControllerExecuteServer{stream})
}

type LocustClusterKubernetesStackController_ExecuteServer interface {
	Send(*model.LocustClusterKubernetesStackResponse) error
	grpc.ServerStream
}

type locustClusterKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *locustClusterKubernetesStackControllerExecuteServer) Send(m *model.LocustClusterKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LocustClusterKubernetesStackController_GetStackOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.LocustClusterKubernetesStackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocustClusterKubernetesStackControllerServer).GetStackOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocustClusterKubernetesStackController_GetStackOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocustClusterKubernetesStackControllerServer).GetStackOutputs(ctx, req.(*model.LocustClusterKubernetesStackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// LocustClusterKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for LocustClusterKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocustClusterKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.locustcluster.stack.kubernetes.service.LocustClusterKubernetesStackController",
	HandlerType: (*LocustClusterKubernetesStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStackOutputs",
			Handler:    _LocustClusterKubernetesStackController_GetStackOutputs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _LocustClusterKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/locustcluster/stack/kubernetes/service/stack.proto",
}