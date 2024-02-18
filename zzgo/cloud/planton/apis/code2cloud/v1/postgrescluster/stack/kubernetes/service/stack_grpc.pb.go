// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/postgrescluster/stack/kubernetes/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/postgrescluster/stack/kubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PostgresClusterKubernetesStackController_Execute_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.postgrescluster.stack.kubernetes.service.PostgresClusterKubernetesStackController/execute"
	PostgresClusterKubernetesStackController_GetStackOutputs_FullMethodName = "/cloud.planton.apis.code2cloud.v1.postgrescluster.stack.kubernetes.service.PostgresClusterKubernetesStackController/getStackOutputs"
)

// PostgresClusterKubernetesStackControllerClient is the client API for PostgresClusterKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostgresClusterKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.PostgresClusterKubernetesStackInput, opts ...grpc.CallOption) (PostgresClusterKubernetesStackController_ExecuteClient, error)
	GetStackOutputs(ctx context.Context, in *model.PostgresClusterKubernetesStackInput, opts ...grpc.CallOption) (*model.PostgresClusterKubernetesStackOutputs, error)
}

type postgresClusterKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewPostgresClusterKubernetesStackControllerClient(cc grpc.ClientConnInterface) PostgresClusterKubernetesStackControllerClient {
	return &postgresClusterKubernetesStackControllerClient{cc}
}

func (c *postgresClusterKubernetesStackControllerClient) Execute(ctx context.Context, in *model.PostgresClusterKubernetesStackInput, opts ...grpc.CallOption) (PostgresClusterKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &PostgresClusterKubernetesStackController_ServiceDesc.Streams[0], PostgresClusterKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &postgresClusterKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PostgresClusterKubernetesStackController_ExecuteClient interface {
	Recv() (*model.PostgresClusterKubernetesStackResponse, error)
	grpc.ClientStream
}

type postgresClusterKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *postgresClusterKubernetesStackControllerExecuteClient) Recv() (*model.PostgresClusterKubernetesStackResponse, error) {
	m := new(model.PostgresClusterKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *postgresClusterKubernetesStackControllerClient) GetStackOutputs(ctx context.Context, in *model.PostgresClusterKubernetesStackInput, opts ...grpc.CallOption) (*model.PostgresClusterKubernetesStackOutputs, error) {
	out := new(model.PostgresClusterKubernetesStackOutputs)
	err := c.cc.Invoke(ctx, PostgresClusterKubernetesStackController_GetStackOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostgresClusterKubernetesStackControllerServer is the server API for PostgresClusterKubernetesStackController service.
// All implementations should embed UnimplementedPostgresClusterKubernetesStackControllerServer
// for forward compatibility
type PostgresClusterKubernetesStackControllerServer interface {
	Execute(*model.PostgresClusterKubernetesStackInput, PostgresClusterKubernetesStackController_ExecuteServer) error
	GetStackOutputs(context.Context, *model.PostgresClusterKubernetesStackInput) (*model.PostgresClusterKubernetesStackOutputs, error)
}

// UnimplementedPostgresClusterKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedPostgresClusterKubernetesStackControllerServer struct {
}

func (UnimplementedPostgresClusterKubernetesStackControllerServer) Execute(*model.PostgresClusterKubernetesStackInput, PostgresClusterKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedPostgresClusterKubernetesStackControllerServer) GetStackOutputs(context.Context, *model.PostgresClusterKubernetesStackInput) (*model.PostgresClusterKubernetesStackOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackOutputs not implemented")
}

// UnsafePostgresClusterKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostgresClusterKubernetesStackControllerServer will
// result in compilation errors.
type UnsafePostgresClusterKubernetesStackControllerServer interface {
	mustEmbedUnimplementedPostgresClusterKubernetesStackControllerServer()
}

func RegisterPostgresClusterKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv PostgresClusterKubernetesStackControllerServer) {
	s.RegisterService(&PostgresClusterKubernetesStackController_ServiceDesc, srv)
}

func _PostgresClusterKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.PostgresClusterKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostgresClusterKubernetesStackControllerServer).Execute(m, &postgresClusterKubernetesStackControllerExecuteServer{stream})
}

type PostgresClusterKubernetesStackController_ExecuteServer interface {
	Send(*model.PostgresClusterKubernetesStackResponse) error
	grpc.ServerStream
}

type postgresClusterKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *postgresClusterKubernetesStackControllerExecuteServer) Send(m *model.PostgresClusterKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PostgresClusterKubernetesStackController_GetStackOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PostgresClusterKubernetesStackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresClusterKubernetesStackControllerServer).GetStackOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgresClusterKubernetesStackController_GetStackOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresClusterKubernetesStackControllerServer).GetStackOutputs(ctx, req.(*model.PostgresClusterKubernetesStackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// PostgresClusterKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for PostgresClusterKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostgresClusterKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.postgrescluster.stack.kubernetes.service.PostgresClusterKubernetesStackController",
	HandlerType: (*PostgresClusterKubernetesStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStackOutputs",
			Handler:    _PostgresClusterKubernetesStackController_GetStackOutputs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _PostgresClusterKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/postgrescluster/stack/kubernetes/service/stack.proto",
}