// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/prometheuskubernetes/stack/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/prometheuskubernetes/stack/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PrometheusKubernetesStackController_Execute_FullMethodName = "/cloud.planton.apis.code2cloud.v1.prometheuskubernetes.stack.service.PrometheusKubernetesStackController/execute"
)

// PrometheusKubernetesStackControllerClient is the client API for PrometheusKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrometheusKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.PrometheusKubernetesStackInput, opts ...grpc.CallOption) (PrometheusKubernetesStackController_ExecuteClient, error)
}

type prometheusKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewPrometheusKubernetesStackControllerClient(cc grpc.ClientConnInterface) PrometheusKubernetesStackControllerClient {
	return &prometheusKubernetesStackControllerClient{cc}
}

func (c *prometheusKubernetesStackControllerClient) Execute(ctx context.Context, in *model.PrometheusKubernetesStackInput, opts ...grpc.CallOption) (PrometheusKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrometheusKubernetesStackController_ServiceDesc.Streams[0], PrometheusKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &prometheusKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrometheusKubernetesStackController_ExecuteClient interface {
	Recv() (*model.PrometheusKubernetesStackResponse, error)
	grpc.ClientStream
}

type prometheusKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *prometheusKubernetesStackControllerExecuteClient) Recv() (*model.PrometheusKubernetesStackResponse, error) {
	m := new(model.PrometheusKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrometheusKubernetesStackControllerServer is the server API for PrometheusKubernetesStackController service.
// All implementations should embed UnimplementedPrometheusKubernetesStackControllerServer
// for forward compatibility
type PrometheusKubernetesStackControllerServer interface {
	Execute(*model.PrometheusKubernetesStackInput, PrometheusKubernetesStackController_ExecuteServer) error
}

// UnimplementedPrometheusKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedPrometheusKubernetesStackControllerServer struct {
}

func (UnimplementedPrometheusKubernetesStackControllerServer) Execute(*model.PrometheusKubernetesStackInput, PrometheusKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafePrometheusKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrometheusKubernetesStackControllerServer will
// result in compilation errors.
type UnsafePrometheusKubernetesStackControllerServer interface {
	mustEmbedUnimplementedPrometheusKubernetesStackControllerServer()
}

func RegisterPrometheusKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv PrometheusKubernetesStackControllerServer) {
	s.RegisterService(&PrometheusKubernetesStackController_ServiceDesc, srv)
}

func _PrometheusKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.PrometheusKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrometheusKubernetesStackControllerServer).Execute(m, &prometheusKubernetesStackControllerExecuteServer{stream})
}

type PrometheusKubernetesStackController_ExecuteServer interface {
	Send(*model.PrometheusKubernetesStackResponse) error
	grpc.ServerStream
}

type prometheusKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *prometheusKubernetesStackControllerExecuteServer) Send(m *model.PrometheusKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PrometheusKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for PrometheusKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrometheusKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.prometheuskubernetes.stack.service.PrometheusKubernetesStackController",
	HandlerType: (*PrometheusKubernetesStackControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _PrometheusKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/prometheuskubernetes/stack/service/stack.proto",
}
