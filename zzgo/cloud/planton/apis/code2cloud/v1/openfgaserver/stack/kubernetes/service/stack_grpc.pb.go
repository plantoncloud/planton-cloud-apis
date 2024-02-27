// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/openfgaserver/stack/kubernetes/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/openfgaserver/stack/kubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OpenFGAServerKubernetesStackController_Execute_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.openfgaserver.stack.kubernetes.service.OpenFGAServerKubernetesStackController/execute"
	OpenFGAServerKubernetesStackController_GetStackOutputs_FullMethodName = "/cloud.planton.apis.code2cloud.v1.openfgaserver.stack.kubernetes.service.OpenFGAServerKubernetesStackController/getStackOutputs"
)

// OpenFGAServerKubernetesStackControllerClient is the client API for OpenFGAServerKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenFGAServerKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.OpenFGAServerKubernetesStackInput, opts ...grpc.CallOption) (OpenFGAServerKubernetesStackController_ExecuteClient, error)
	GetStackOutputs(ctx context.Context, in *model.OpenFGAServerKubernetesStackInput, opts ...grpc.CallOption) (*model.OpenFGAServerKubernetesStackOutputs, error)
}

type openFGAServerKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenFGAServerKubernetesStackControllerClient(cc grpc.ClientConnInterface) OpenFGAServerKubernetesStackControllerClient {
	return &openFGAServerKubernetesStackControllerClient{cc}
}

func (c *openFGAServerKubernetesStackControllerClient) Execute(ctx context.Context, in *model.OpenFGAServerKubernetesStackInput, opts ...grpc.CallOption) (OpenFGAServerKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &OpenFGAServerKubernetesStackController_ServiceDesc.Streams[0], OpenFGAServerKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &openFGAServerKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OpenFGAServerKubernetesStackController_ExecuteClient interface {
	Recv() (*model.OpenFGAServerKubernetesStackResponse, error)
	grpc.ClientStream
}

type openFGAServerKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *openFGAServerKubernetesStackControllerExecuteClient) Recv() (*model.OpenFGAServerKubernetesStackResponse, error) {
	m := new(model.OpenFGAServerKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *openFGAServerKubernetesStackControllerClient) GetStackOutputs(ctx context.Context, in *model.OpenFGAServerKubernetesStackInput, opts ...grpc.CallOption) (*model.OpenFGAServerKubernetesStackOutputs, error) {
	out := new(model.OpenFGAServerKubernetesStackOutputs)
	err := c.cc.Invoke(ctx, OpenFGAServerKubernetesStackController_GetStackOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenFGAServerKubernetesStackControllerServer is the server API for OpenFGAServerKubernetesStackController service.
// All implementations should embed UnimplementedOpenFGAServerKubernetesStackControllerServer
// for forward compatibility
type OpenFGAServerKubernetesStackControllerServer interface {
	Execute(*model.OpenFGAServerKubernetesStackInput, OpenFGAServerKubernetesStackController_ExecuteServer) error
	GetStackOutputs(context.Context, *model.OpenFGAServerKubernetesStackInput) (*model.OpenFGAServerKubernetesStackOutputs, error)
}

// UnimplementedOpenFGAServerKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedOpenFGAServerKubernetesStackControllerServer struct {
}

func (UnimplementedOpenFGAServerKubernetesStackControllerServer) Execute(*model.OpenFGAServerKubernetesStackInput, OpenFGAServerKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedOpenFGAServerKubernetesStackControllerServer) GetStackOutputs(context.Context, *model.OpenFGAServerKubernetesStackInput) (*model.OpenFGAServerKubernetesStackOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackOutputs not implemented")
}

// UnsafeOpenFGAServerKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenFGAServerKubernetesStackControllerServer will
// result in compilation errors.
type UnsafeOpenFGAServerKubernetesStackControllerServer interface {
	mustEmbedUnimplementedOpenFGAServerKubernetesStackControllerServer()
}

func RegisterOpenFGAServerKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv OpenFGAServerKubernetesStackControllerServer) {
	s.RegisterService(&OpenFGAServerKubernetesStackController_ServiceDesc, srv)
}

func _OpenFGAServerKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.OpenFGAServerKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OpenFGAServerKubernetesStackControllerServer).Execute(m, &openFGAServerKubernetesStackControllerExecuteServer{stream})
}

type OpenFGAServerKubernetesStackController_ExecuteServer interface {
	Send(*model.OpenFGAServerKubernetesStackResponse) error
	grpc.ServerStream
}

type openFGAServerKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *openFGAServerKubernetesStackControllerExecuteServer) Send(m *model.OpenFGAServerKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OpenFGAServerKubernetesStackController_GetStackOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenFGAServerKubernetesStackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerKubernetesStackControllerServer).GetStackOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerKubernetesStackController_GetStackOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerKubernetesStackControllerServer).GetStackOutputs(ctx, req.(*model.OpenFGAServerKubernetesStackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// OpenFGAServerKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for OpenFGAServerKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenFGAServerKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.openfgaserver.stack.kubernetes.service.OpenFGAServerKubernetesStackController",
	HandlerType: (*OpenFGAServerKubernetesStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStackOutputs",
			Handler:    _OpenFGAServerKubernetesStackController_GetStackOutputs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _OpenFGAServerKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/openfgaserver/stack/kubernetes/service/stack.proto",
}
