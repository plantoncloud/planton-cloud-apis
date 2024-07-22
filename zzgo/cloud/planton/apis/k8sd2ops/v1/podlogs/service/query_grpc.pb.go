// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/k8sd2ops/v1/podlogs/service/query.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/kubernetesobject/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/k8sd2ops/v1/podlogs/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ApiResourceKubernetesPodLogsQueryController_StreamPodLogs_FullMethodName = "/cloud.planton.apis.k8sd2ops.v1.podlogs.service.ApiResourceKubernetesPodLogsQueryController/streamPodLogs"
)

// ApiResourceKubernetesPodLogsQueryControllerClient is the client API for ApiResourceKubernetesPodLogsQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiResourceKubernetesPodLogsQueryControllerClient interface {
	// stream logs of all kubernetes pods running in a kubernetes-cluster on the specified filters
	StreamPodLogs(ctx context.Context, in *model.StreamApiResourceKubernetesPodLogsInput, opts ...grpc.CallOption) (ApiResourceKubernetesPodLogsQueryController_StreamPodLogsClient, error)
}

type apiResourceKubernetesPodLogsQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewApiResourceKubernetesPodLogsQueryControllerClient(cc grpc.ClientConnInterface) ApiResourceKubernetesPodLogsQueryControllerClient {
	return &apiResourceKubernetesPodLogsQueryControllerClient{cc}
}

func (c *apiResourceKubernetesPodLogsQueryControllerClient) StreamPodLogs(ctx context.Context, in *model.StreamApiResourceKubernetesPodLogsInput, opts ...grpc.CallOption) (ApiResourceKubernetesPodLogsQueryController_StreamPodLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ApiResourceKubernetesPodLogsQueryController_ServiceDesc.Streams[0], ApiResourceKubernetesPodLogsQueryController_StreamPodLogs_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &apiResourceKubernetesPodLogsQueryControllerStreamPodLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApiResourceKubernetesPodLogsQueryController_StreamPodLogsClient interface {
	Recv() (*model1.PodLogLine, error)
	grpc.ClientStream
}

type apiResourceKubernetesPodLogsQueryControllerStreamPodLogsClient struct {
	grpc.ClientStream
}

func (x *apiResourceKubernetesPodLogsQueryControllerStreamPodLogsClient) Recv() (*model1.PodLogLine, error) {
	m := new(model1.PodLogLine)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ApiResourceKubernetesPodLogsQueryControllerServer is the server API for ApiResourceKubernetesPodLogsQueryController service.
// All implementations should embed UnimplementedApiResourceKubernetesPodLogsQueryControllerServer
// for forward compatibility
type ApiResourceKubernetesPodLogsQueryControllerServer interface {
	// stream logs of all kubernetes pods running in a kubernetes-cluster on the specified filters
	StreamPodLogs(*model.StreamApiResourceKubernetesPodLogsInput, ApiResourceKubernetesPodLogsQueryController_StreamPodLogsServer) error
}

// UnimplementedApiResourceKubernetesPodLogsQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedApiResourceKubernetesPodLogsQueryControllerServer struct {
}

func (UnimplementedApiResourceKubernetesPodLogsQueryControllerServer) StreamPodLogs(*model.StreamApiResourceKubernetesPodLogsInput, ApiResourceKubernetesPodLogsQueryController_StreamPodLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPodLogs not implemented")
}

// UnsafeApiResourceKubernetesPodLogsQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiResourceKubernetesPodLogsQueryControllerServer will
// result in compilation errors.
type UnsafeApiResourceKubernetesPodLogsQueryControllerServer interface {
	mustEmbedUnimplementedApiResourceKubernetesPodLogsQueryControllerServer()
}

func RegisterApiResourceKubernetesPodLogsQueryControllerServer(s grpc.ServiceRegistrar, srv ApiResourceKubernetesPodLogsQueryControllerServer) {
	s.RegisterService(&ApiResourceKubernetesPodLogsQueryController_ServiceDesc, srv)
}

func _ApiResourceKubernetesPodLogsQueryController_StreamPodLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.StreamApiResourceKubernetesPodLogsInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApiResourceKubernetesPodLogsQueryControllerServer).StreamPodLogs(m, &apiResourceKubernetesPodLogsQueryControllerStreamPodLogsServer{stream})
}

type ApiResourceKubernetesPodLogsQueryController_StreamPodLogsServer interface {
	Send(*model1.PodLogLine) error
	grpc.ServerStream
}

type apiResourceKubernetesPodLogsQueryControllerStreamPodLogsServer struct {
	grpc.ServerStream
}

func (x *apiResourceKubernetesPodLogsQueryControllerStreamPodLogsServer) Send(m *model1.PodLogLine) error {
	return x.ServerStream.SendMsg(m)
}

// ApiResourceKubernetesPodLogsQueryController_ServiceDesc is the grpc.ServiceDesc for ApiResourceKubernetesPodLogsQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiResourceKubernetesPodLogsQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.k8sd2ops.v1.podlogs.service.ApiResourceKubernetesPodLogsQueryController",
	HandlerType: (*ApiResourceKubernetesPodLogsQueryControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "streamPodLogs",
			Handler:       _ApiResourceKubernetesPodLogsQueryController_StreamPodLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/k8sd2ops/v1/podlogs/service/query.proto",
}