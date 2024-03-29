// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/environment/stack/gcpgke/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/environment/stack/gcpgke/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EnvironmentGcpGkeStackController_Execute_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.service.EnvironmentGcpGkeStackController/execute"
	EnvironmentGcpGkeStackController_GetStackOutputs_FullMethodName = "/cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.service.EnvironmentGcpGkeStackController/getStackOutputs"
)

// EnvironmentGcpGkeStackControllerClient is the client API for EnvironmentGcpGkeStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentGcpGkeStackControllerClient interface {
	Execute(ctx context.Context, in *model.EnvironmentGcpGkeStackInput, opts ...grpc.CallOption) (EnvironmentGcpGkeStackController_ExecuteClient, error)
	GetStackOutputs(ctx context.Context, in *model.EnvironmentGcpGkeStackInput, opts ...grpc.CallOption) (*model.EnvironmentGcpGkeStackOutputs, error)
}

type environmentGcpGkeStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentGcpGkeStackControllerClient(cc grpc.ClientConnInterface) EnvironmentGcpGkeStackControllerClient {
	return &environmentGcpGkeStackControllerClient{cc}
}

func (c *environmentGcpGkeStackControllerClient) Execute(ctx context.Context, in *model.EnvironmentGcpGkeStackInput, opts ...grpc.CallOption) (EnvironmentGcpGkeStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &EnvironmentGcpGkeStackController_ServiceDesc.Streams[0], EnvironmentGcpGkeStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &environmentGcpGkeStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EnvironmentGcpGkeStackController_ExecuteClient interface {
	Recv() (*model.EnvironmentGcpGkeStackResponse, error)
	grpc.ClientStream
}

type environmentGcpGkeStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *environmentGcpGkeStackControllerExecuteClient) Recv() (*model.EnvironmentGcpGkeStackResponse, error) {
	m := new(model.EnvironmentGcpGkeStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *environmentGcpGkeStackControllerClient) GetStackOutputs(ctx context.Context, in *model.EnvironmentGcpGkeStackInput, opts ...grpc.CallOption) (*model.EnvironmentGcpGkeStackOutputs, error) {
	out := new(model.EnvironmentGcpGkeStackOutputs)
	err := c.cc.Invoke(ctx, EnvironmentGcpGkeStackController_GetStackOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentGcpGkeStackControllerServer is the server API for EnvironmentGcpGkeStackController service.
// All implementations should embed UnimplementedEnvironmentGcpGkeStackControllerServer
// for forward compatibility
type EnvironmentGcpGkeStackControllerServer interface {
	Execute(*model.EnvironmentGcpGkeStackInput, EnvironmentGcpGkeStackController_ExecuteServer) error
	GetStackOutputs(context.Context, *model.EnvironmentGcpGkeStackInput) (*model.EnvironmentGcpGkeStackOutputs, error)
}

// UnimplementedEnvironmentGcpGkeStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedEnvironmentGcpGkeStackControllerServer struct {
}

func (UnimplementedEnvironmentGcpGkeStackControllerServer) Execute(*model.EnvironmentGcpGkeStackInput, EnvironmentGcpGkeStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedEnvironmentGcpGkeStackControllerServer) GetStackOutputs(context.Context, *model.EnvironmentGcpGkeStackInput) (*model.EnvironmentGcpGkeStackOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackOutputs not implemented")
}

// UnsafeEnvironmentGcpGkeStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvironmentGcpGkeStackControllerServer will
// result in compilation errors.
type UnsafeEnvironmentGcpGkeStackControllerServer interface {
	mustEmbedUnimplementedEnvironmentGcpGkeStackControllerServer()
}

func RegisterEnvironmentGcpGkeStackControllerServer(s grpc.ServiceRegistrar, srv EnvironmentGcpGkeStackControllerServer) {
	s.RegisterService(&EnvironmentGcpGkeStackController_ServiceDesc, srv)
}

func _EnvironmentGcpGkeStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.EnvironmentGcpGkeStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EnvironmentGcpGkeStackControllerServer).Execute(m, &environmentGcpGkeStackControllerExecuteServer{stream})
}

type EnvironmentGcpGkeStackController_ExecuteServer interface {
	Send(*model.EnvironmentGcpGkeStackResponse) error
	grpc.ServerStream
}

type environmentGcpGkeStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *environmentGcpGkeStackControllerExecuteServer) Send(m *model.EnvironmentGcpGkeStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EnvironmentGcpGkeStackController_GetStackOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EnvironmentGcpGkeStackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentGcpGkeStackControllerServer).GetStackOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentGcpGkeStackController_GetStackOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentGcpGkeStackControllerServer).GetStackOutputs(ctx, req.(*model.EnvironmentGcpGkeStackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// EnvironmentGcpGkeStackController_ServiceDesc is the grpc.ServiceDesc for EnvironmentGcpGkeStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnvironmentGcpGkeStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.service.EnvironmentGcpGkeStackController",
	HandlerType: (*EnvironmentGcpGkeStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStackOutputs",
			Handler:    _EnvironmentGcpGkeStackController_GetStackOutputs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _EnvironmentGcpGkeStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/environment/stack/gcpgke/service/stack.proto",
}
