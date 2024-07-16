// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/kubernetes/istiohttpendpoint/stack/service/stack.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/kubernetes/istiohttpendpoint/stack/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IstioHttpEndpointStackController_Execute_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.kubernetes.istiohttpendpoint.stack.service.IstioHttpEndpointStackController/execute"
)

// IstioHttpEndpointStackControllerClient is the client API for IstioHttpEndpointStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IstioHttpEndpointStackControllerClient interface {
	Execute(ctx context.Context, in *model.IstioHttpEndpointStackInput, opts ...grpc.CallOption) (IstioHttpEndpointStackController_ExecuteClient, error)
}

type istioHttpEndpointStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIstioHttpEndpointStackControllerClient(cc grpc.ClientConnInterface) IstioHttpEndpointStackControllerClient {
	return &istioHttpEndpointStackControllerClient{cc}
}

func (c *istioHttpEndpointStackControllerClient) Execute(ctx context.Context, in *model.IstioHttpEndpointStackInput, opts ...grpc.CallOption) (IstioHttpEndpointStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &IstioHttpEndpointStackController_ServiceDesc.Streams[0], IstioHttpEndpointStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &istioHttpEndpointStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IstioHttpEndpointStackController_ExecuteClient interface {
	Recv() (*model.IstioHttpEndpointStackResponse, error)
	grpc.ClientStream
}

type istioHttpEndpointStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *istioHttpEndpointStackControllerExecuteClient) Recv() (*model.IstioHttpEndpointStackResponse, error) {
	m := new(model.IstioHttpEndpointStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IstioHttpEndpointStackControllerServer is the server API for IstioHttpEndpointStackController service.
// All implementations should embed UnimplementedIstioHttpEndpointStackControllerServer
// for forward compatibility
type IstioHttpEndpointStackControllerServer interface {
	Execute(*model.IstioHttpEndpointStackInput, IstioHttpEndpointStackController_ExecuteServer) error
}

// UnimplementedIstioHttpEndpointStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedIstioHttpEndpointStackControllerServer struct {
}

func (UnimplementedIstioHttpEndpointStackControllerServer) Execute(*model.IstioHttpEndpointStackInput, IstioHttpEndpointStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeIstioHttpEndpointStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IstioHttpEndpointStackControllerServer will
// result in compilation errors.
type UnsafeIstioHttpEndpointStackControllerServer interface {
	mustEmbedUnimplementedIstioHttpEndpointStackControllerServer()
}

func RegisterIstioHttpEndpointStackControllerServer(s grpc.ServiceRegistrar, srv IstioHttpEndpointStackControllerServer) {
	s.RegisterService(&IstioHttpEndpointStackController_ServiceDesc, srv)
}

func _IstioHttpEndpointStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.IstioHttpEndpointStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IstioHttpEndpointStackControllerServer).Execute(m, &istioHttpEndpointStackControllerExecuteServer{stream})
}

type IstioHttpEndpointStackController_ExecuteServer interface {
	Send(*model.IstioHttpEndpointStackResponse) error
	grpc.ServerStream
}

type istioHttpEndpointStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *istioHttpEndpointStackControllerExecuteServer) Send(m *model.IstioHttpEndpointStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

// IstioHttpEndpointStackController_ServiceDesc is the grpc.ServiceDesc for IstioHttpEndpointStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IstioHttpEndpointStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.kubernetes.istiohttpendpoint.stack.service.IstioHttpEndpointStackController",
	HandlerType: (*IstioHttpEndpointStackControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _IstioHttpEndpointStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/kubernetes/istiohttpendpoint/stack/service/stack.proto",
}
