// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/microserviceinstance/stack/kubernetes/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/microserviceinstance/stack/kubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MicroserviceInstanceKubernetesStackController_Execute_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.stack.kubernetes.service.MicroserviceInstanceKubernetesStackController/execute"
	MicroserviceInstanceKubernetesStackController_GetStackOutputs_FullMethodName = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.stack.kubernetes.service.MicroserviceInstanceKubernetesStackController/getStackOutputs"
)

// MicroserviceInstanceKubernetesStackControllerClient is the client API for MicroserviceInstanceKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicroserviceInstanceKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.MicroserviceInstanceKubernetesStackInput, opts ...grpc.CallOption) (MicroserviceInstanceKubernetesStackController_ExecuteClient, error)
	GetStackOutputs(ctx context.Context, in *model.MicroserviceInstanceKubernetesStackInput, opts ...grpc.CallOption) (*model.MicroserviceInstanceKubernetesStackOutputs, error)
}

type microserviceInstanceKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroserviceInstanceKubernetesStackControllerClient(cc grpc.ClientConnInterface) MicroserviceInstanceKubernetesStackControllerClient {
	return &microserviceInstanceKubernetesStackControllerClient{cc}
}

func (c *microserviceInstanceKubernetesStackControllerClient) Execute(ctx context.Context, in *model.MicroserviceInstanceKubernetesStackInput, opts ...grpc.CallOption) (MicroserviceInstanceKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &MicroserviceInstanceKubernetesStackController_ServiceDesc.Streams[0], MicroserviceInstanceKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &microserviceInstanceKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MicroserviceInstanceKubernetesStackController_ExecuteClient interface {
	Recv() (*model.MicroserviceInstanceKubernetesStackResponse, error)
	grpc.ClientStream
}

type microserviceInstanceKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *microserviceInstanceKubernetesStackControllerExecuteClient) Recv() (*model.MicroserviceInstanceKubernetesStackResponse, error) {
	m := new(model.MicroserviceInstanceKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *microserviceInstanceKubernetesStackControllerClient) GetStackOutputs(ctx context.Context, in *model.MicroserviceInstanceKubernetesStackInput, opts ...grpc.CallOption) (*model.MicroserviceInstanceKubernetesStackOutputs, error) {
	out := new(model.MicroserviceInstanceKubernetesStackOutputs)
	err := c.cc.Invoke(ctx, MicroserviceInstanceKubernetesStackController_GetStackOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroserviceInstanceKubernetesStackControllerServer is the server API for MicroserviceInstanceKubernetesStackController service.
// All implementations should embed UnimplementedMicroserviceInstanceKubernetesStackControllerServer
// for forward compatibility
type MicroserviceInstanceKubernetesStackControllerServer interface {
	Execute(*model.MicroserviceInstanceKubernetesStackInput, MicroserviceInstanceKubernetesStackController_ExecuteServer) error
	GetStackOutputs(context.Context, *model.MicroserviceInstanceKubernetesStackInput) (*model.MicroserviceInstanceKubernetesStackOutputs, error)
}

// UnimplementedMicroserviceInstanceKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedMicroserviceInstanceKubernetesStackControllerServer struct {
}

func (UnimplementedMicroserviceInstanceKubernetesStackControllerServer) Execute(*model.MicroserviceInstanceKubernetesStackInput, MicroserviceInstanceKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedMicroserviceInstanceKubernetesStackControllerServer) GetStackOutputs(context.Context, *model.MicroserviceInstanceKubernetesStackInput) (*model.MicroserviceInstanceKubernetesStackOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackOutputs not implemented")
}

// UnsafeMicroserviceInstanceKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroserviceInstanceKubernetesStackControllerServer will
// result in compilation errors.
type UnsafeMicroserviceInstanceKubernetesStackControllerServer interface {
	mustEmbedUnimplementedMicroserviceInstanceKubernetesStackControllerServer()
}

func RegisterMicroserviceInstanceKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv MicroserviceInstanceKubernetesStackControllerServer) {
	s.RegisterService(&MicroserviceInstanceKubernetesStackController_ServiceDesc, srv)
}

func _MicroserviceInstanceKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.MicroserviceInstanceKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MicroserviceInstanceKubernetesStackControllerServer).Execute(m, &microserviceInstanceKubernetesStackControllerExecuteServer{stream})
}

type MicroserviceInstanceKubernetesStackController_ExecuteServer interface {
	Send(*model.MicroserviceInstanceKubernetesStackResponse) error
	grpc.ServerStream
}

type microserviceInstanceKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *microserviceInstanceKubernetesStackControllerExecuteServer) Send(m *model.MicroserviceInstanceKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MicroserviceInstanceKubernetesStackController_GetStackOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MicroserviceInstanceKubernetesStackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceKubernetesStackControllerServer).GetStackOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceKubernetesStackController_GetStackOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceKubernetesStackControllerServer).GetStackOutputs(ctx, req.(*model.MicroserviceInstanceKubernetesStackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MicroserviceInstanceKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for MicroserviceInstanceKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MicroserviceInstanceKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.microserviceinstance.stack.kubernetes.service.MicroserviceInstanceKubernetesStackController",
	HandlerType: (*MicroserviceInstanceKubernetesStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStackOutputs",
			Handler:    _MicroserviceInstanceKubernetesStackController_GetStackOutputs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _MicroserviceInstanceKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/microserviceinstance/stack/kubernetes/service/stack.proto",
}