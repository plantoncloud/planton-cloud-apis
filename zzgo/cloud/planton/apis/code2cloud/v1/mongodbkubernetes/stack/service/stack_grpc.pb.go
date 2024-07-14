// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/mongodbkubernetes/stack/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/mongodbkubernetes/stack/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MongodbKubernetesStackController_Execute_FullMethodName = "/cloud.planton.apis.code2cloud.v1.mongodbkubernetes.stack.service.MongodbKubernetesStackController/execute"
)

// MongodbKubernetesStackControllerClient is the client API for MongodbKubernetesStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MongodbKubernetesStackControllerClient interface {
	Execute(ctx context.Context, in *model.MongodbKubernetesStackInput, opts ...grpc.CallOption) (MongodbKubernetesStackController_ExecuteClient, error)
}

type mongodbKubernetesStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMongodbKubernetesStackControllerClient(cc grpc.ClientConnInterface) MongodbKubernetesStackControllerClient {
	return &mongodbKubernetesStackControllerClient{cc}
}

func (c *mongodbKubernetesStackControllerClient) Execute(ctx context.Context, in *model.MongodbKubernetesStackInput, opts ...grpc.CallOption) (MongodbKubernetesStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &MongodbKubernetesStackController_ServiceDesc.Streams[0], MongodbKubernetesStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &mongodbKubernetesStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MongodbKubernetesStackController_ExecuteClient interface {
	Recv() (*model.MongodbKubernetesStackResponse, error)
	grpc.ClientStream
}

type mongodbKubernetesStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *mongodbKubernetesStackControllerExecuteClient) Recv() (*model.MongodbKubernetesStackResponse, error) {
	m := new(model.MongodbKubernetesStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MongodbKubernetesStackControllerServer is the server API for MongodbKubernetesStackController service.
// All implementations should embed UnimplementedMongodbKubernetesStackControllerServer
// for forward compatibility
type MongodbKubernetesStackControllerServer interface {
	Execute(*model.MongodbKubernetesStackInput, MongodbKubernetesStackController_ExecuteServer) error
}

// UnimplementedMongodbKubernetesStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedMongodbKubernetesStackControllerServer struct {
}

func (UnimplementedMongodbKubernetesStackControllerServer) Execute(*model.MongodbKubernetesStackInput, MongodbKubernetesStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeMongodbKubernetesStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MongodbKubernetesStackControllerServer will
// result in compilation errors.
type UnsafeMongodbKubernetesStackControllerServer interface {
	mustEmbedUnimplementedMongodbKubernetesStackControllerServer()
}

func RegisterMongodbKubernetesStackControllerServer(s grpc.ServiceRegistrar, srv MongodbKubernetesStackControllerServer) {
	s.RegisterService(&MongodbKubernetesStackController_ServiceDesc, srv)
}

func _MongodbKubernetesStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.MongodbKubernetesStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MongodbKubernetesStackControllerServer).Execute(m, &mongodbKubernetesStackControllerExecuteServer{stream})
}

type MongodbKubernetesStackController_ExecuteServer interface {
	Send(*model.MongodbKubernetesStackResponse) error
	grpc.ServerStream
}

type mongodbKubernetesStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *mongodbKubernetesStackControllerExecuteServer) Send(m *model.MongodbKubernetesStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MongodbKubernetesStackController_ServiceDesc is the grpc.ServiceDesc for MongodbKubernetesStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MongodbKubernetesStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.mongodbkubernetes.stack.service.MongodbKubernetesStackController",
	HandlerType: (*MongodbKubernetesStackControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _MongodbKubernetesStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/mongodbkubernetes/stack/service/stack.proto",
}
