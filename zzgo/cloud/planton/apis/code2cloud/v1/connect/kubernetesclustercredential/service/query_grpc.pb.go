// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/connect/kubernetesclustercredential/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/connect/kubernetesclustercredential/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/apiresources/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KubernetesClusterCredentialQueryController_Get_FullMethodName = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterCredentialQueryController/get"
)

// KubernetesClusterCredentialQueryControllerClient is the client API for KubernetesClusterCredentialQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubernetesClusterCredentialQueryControllerClient interface {
	// look up a kubernetes-cluster-credential by id
	Get(ctx context.Context, in *model.KubernetesClusterCredentialId, opts ...grpc.CallOption) (*model.KubernetesClusterCredential, error)
}

type kubernetesClusterCredentialQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKubernetesClusterCredentialQueryControllerClient(cc grpc.ClientConnInterface) KubernetesClusterCredentialQueryControllerClient {
	return &kubernetesClusterCredentialQueryControllerClient{cc}
}

func (c *kubernetesClusterCredentialQueryControllerClient) Get(ctx context.Context, in *model.KubernetesClusterCredentialId, opts ...grpc.CallOption) (*model.KubernetesClusterCredential, error) {
	out := new(model.KubernetesClusterCredential)
	err := c.cc.Invoke(ctx, KubernetesClusterCredentialQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubernetesClusterCredentialQueryControllerServer is the server API for KubernetesClusterCredentialQueryController service.
// All implementations should embed UnimplementedKubernetesClusterCredentialQueryControllerServer
// for forward compatibility
type KubernetesClusterCredentialQueryControllerServer interface {
	// look up a kubernetes-cluster-credential by id
	Get(context.Context, *model.KubernetesClusterCredentialId) (*model.KubernetesClusterCredential, error)
}

// UnimplementedKubernetesClusterCredentialQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKubernetesClusterCredentialQueryControllerServer struct {
}

func (UnimplementedKubernetesClusterCredentialQueryControllerServer) Get(context.Context, *model.KubernetesClusterCredentialId) (*model.KubernetesClusterCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeKubernetesClusterCredentialQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubernetesClusterCredentialQueryControllerServer will
// result in compilation errors.
type UnsafeKubernetesClusterCredentialQueryControllerServer interface {
	mustEmbedUnimplementedKubernetesClusterCredentialQueryControllerServer()
}

func RegisterKubernetesClusterCredentialQueryControllerServer(s grpc.ServiceRegistrar, srv KubernetesClusterCredentialQueryControllerServer) {
	s.RegisterService(&KubernetesClusterCredentialQueryController_ServiceDesc, srv)
}

func _KubernetesClusterCredentialQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KubernetesClusterCredentialId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterCredentialQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesClusterCredentialQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterCredentialQueryControllerServer).Get(ctx, req.(*model.KubernetesClusterCredentialId))
	}
	return interceptor(ctx, in, info, handler)
}

// KubernetesClusterCredentialQueryController_ServiceDesc is the grpc.ServiceDesc for KubernetesClusterCredentialQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubernetesClusterCredentialQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterCredentialQueryController",
	HandlerType: (*KubernetesClusterCredentialQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _KubernetesClusterCredentialQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/connect/kubernetesclustercredential/service/query.proto",
}

const (
	KubernetesClusterKubernetesObjectQueryController_Get_FullMethodName                     = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectQueryController/get"
	KubernetesClusterKubernetesObjectQueryController_FindNamespaces_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectQueryController/findNamespaces"
	KubernetesClusterKubernetesObjectQueryController_StreamKubernetesObjects_FullMethodName = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectQueryController/streamKubernetesObjects"
	KubernetesClusterKubernetesObjectQueryController_FindPods_FullMethodName                = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectQueryController/findPods"
	KubernetesClusterKubernetesObjectQueryController_GetPod_FullMethodName                  = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectQueryController/getPod"
	KubernetesClusterKubernetesObjectQueryController_StreamPodLogs_FullMethodName           = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectQueryController/streamPodLogs"
)

// KubernetesClusterKubernetesObjectQueryControllerClient is the client API for KubernetesClusterKubernetesObjectQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubernetesClusterKubernetesObjectQueryControllerClient interface {
	// get detailed object of a kubernetes object
	Get(ctx context.Context, in *model.KubernetesClusterKubernetesObject, opts ...grpc.CallOption) (*model1.KubernetesObjectDetail, error)
	// find list of namespaces on a kubernetes cluster
	FindNamespaces(ctx context.Context, in *model.KubernetesClusterCredentialId, opts ...grpc.CallOption) (*model1.KubernetesNamespaces, error)
	// stream all kubernetes objects from a kubernetes namespace in kube-cluster.
	// this is a streaming rpc since the lookup involves several kubernetes api-calls to fetch all the kubernetes-api-resources.
	// because of high number of api calls to upstream kubernetes cluster, the response is streamed to the client.
	StreamKubernetesObjects(ctx context.Context, in *model.StreamKubernetesClusterNamespaceKubernetesObjectsInput, opts ...grpc.CallOption) (KubernetesClusterKubernetesObjectQueryController_StreamKubernetesObjectsClient, error)
	// find list of pods in a kube-cluster on the specified filters
	FindPods(ctx context.Context, in *model.FindKubernetesClusterPodsInput, opts ...grpc.CallOption) (*model1.Pods, error)
	// get details of a pod
	GetPod(ctx context.Context, in *model.KubernetesClusterKubernetesObject, opts ...grpc.CallOption) (*model1.Pod, error)
	// stream logs of all kubernetes pods running in a kube-cluster on the specified filters
	StreamPodLogs(ctx context.Context, in *model.StreamKubernetesClusterPodLogsInput, opts ...grpc.CallOption) (KubernetesClusterKubernetesObjectQueryController_StreamPodLogsClient, error)
}

type kubernetesClusterKubernetesObjectQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKubernetesClusterKubernetesObjectQueryControllerClient(cc grpc.ClientConnInterface) KubernetesClusterKubernetesObjectQueryControllerClient {
	return &kubernetesClusterKubernetesObjectQueryControllerClient{cc}
}

func (c *kubernetesClusterKubernetesObjectQueryControllerClient) Get(ctx context.Context, in *model.KubernetesClusterKubernetesObject, opts ...grpc.CallOption) (*model1.KubernetesObjectDetail, error) {
	out := new(model1.KubernetesObjectDetail)
	err := c.cc.Invoke(ctx, KubernetesClusterKubernetesObjectQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterKubernetesObjectQueryControllerClient) FindNamespaces(ctx context.Context, in *model.KubernetesClusterCredentialId, opts ...grpc.CallOption) (*model1.KubernetesNamespaces, error) {
	out := new(model1.KubernetesNamespaces)
	err := c.cc.Invoke(ctx, KubernetesClusterKubernetesObjectQueryController_FindNamespaces_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterKubernetesObjectQueryControllerClient) StreamKubernetesObjects(ctx context.Context, in *model.StreamKubernetesClusterNamespaceKubernetesObjectsInput, opts ...grpc.CallOption) (KubernetesClusterKubernetesObjectQueryController_StreamKubernetesObjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &KubernetesClusterKubernetesObjectQueryController_ServiceDesc.Streams[0], KubernetesClusterKubernetesObjectQueryController_StreamKubernetesObjects_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &kubernetesClusterKubernetesObjectQueryControllerStreamKubernetesObjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KubernetesClusterKubernetesObjectQueryController_StreamKubernetesObjectsClient interface {
	Recv() (*model1.KubernetesObjects, error)
	grpc.ClientStream
}

type kubernetesClusterKubernetesObjectQueryControllerStreamKubernetesObjectsClient struct {
	grpc.ClientStream
}

func (x *kubernetesClusterKubernetesObjectQueryControllerStreamKubernetesObjectsClient) Recv() (*model1.KubernetesObjects, error) {
	m := new(model1.KubernetesObjects)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kubernetesClusterKubernetesObjectQueryControllerClient) FindPods(ctx context.Context, in *model.FindKubernetesClusterPodsInput, opts ...grpc.CallOption) (*model1.Pods, error) {
	out := new(model1.Pods)
	err := c.cc.Invoke(ctx, KubernetesClusterKubernetesObjectQueryController_FindPods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterKubernetesObjectQueryControllerClient) GetPod(ctx context.Context, in *model.KubernetesClusterKubernetesObject, opts ...grpc.CallOption) (*model1.Pod, error) {
	out := new(model1.Pod)
	err := c.cc.Invoke(ctx, KubernetesClusterKubernetesObjectQueryController_GetPod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterKubernetesObjectQueryControllerClient) StreamPodLogs(ctx context.Context, in *model.StreamKubernetesClusterPodLogsInput, opts ...grpc.CallOption) (KubernetesClusterKubernetesObjectQueryController_StreamPodLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &KubernetesClusterKubernetesObjectQueryController_ServiceDesc.Streams[1], KubernetesClusterKubernetesObjectQueryController_StreamPodLogs_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &kubernetesClusterKubernetesObjectQueryControllerStreamPodLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KubernetesClusterKubernetesObjectQueryController_StreamPodLogsClient interface {
	Recv() (*model1.PodLogLine, error)
	grpc.ClientStream
}

type kubernetesClusterKubernetesObjectQueryControllerStreamPodLogsClient struct {
	grpc.ClientStream
}

func (x *kubernetesClusterKubernetesObjectQueryControllerStreamPodLogsClient) Recv() (*model1.PodLogLine, error) {
	m := new(model1.PodLogLine)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KubernetesClusterKubernetesObjectQueryControllerServer is the server API for KubernetesClusterKubernetesObjectQueryController service.
// All implementations should embed UnimplementedKubernetesClusterKubernetesObjectQueryControllerServer
// for forward compatibility
type KubernetesClusterKubernetesObjectQueryControllerServer interface {
	// get detailed object of a kubernetes object
	Get(context.Context, *model.KubernetesClusterKubernetesObject) (*model1.KubernetesObjectDetail, error)
	// find list of namespaces on a kubernetes cluster
	FindNamespaces(context.Context, *model.KubernetesClusterCredentialId) (*model1.KubernetesNamespaces, error)
	// stream all kubernetes objects from a kubernetes namespace in kube-cluster.
	// this is a streaming rpc since the lookup involves several kubernetes api-calls to fetch all the kubernetes-api-resources.
	// because of high number of api calls to upstream kubernetes cluster, the response is streamed to the client.
	StreamKubernetesObjects(*model.StreamKubernetesClusterNamespaceKubernetesObjectsInput, KubernetesClusterKubernetesObjectQueryController_StreamKubernetesObjectsServer) error
	// find list of pods in a kube-cluster on the specified filters
	FindPods(context.Context, *model.FindKubernetesClusterPodsInput) (*model1.Pods, error)
	// get details of a pod
	GetPod(context.Context, *model.KubernetesClusterKubernetesObject) (*model1.Pod, error)
	// stream logs of all kubernetes pods running in a kube-cluster on the specified filters
	StreamPodLogs(*model.StreamKubernetesClusterPodLogsInput, KubernetesClusterKubernetesObjectQueryController_StreamPodLogsServer) error
}

// UnimplementedKubernetesClusterKubernetesObjectQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKubernetesClusterKubernetesObjectQueryControllerServer struct {
}

func (UnimplementedKubernetesClusterKubernetesObjectQueryControllerServer) Get(context.Context, *model.KubernetesClusterKubernetesObject) (*model1.KubernetesObjectDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKubernetesClusterKubernetesObjectQueryControllerServer) FindNamespaces(context.Context, *model.KubernetesClusterCredentialId) (*model1.KubernetesNamespaces, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNamespaces not implemented")
}
func (UnimplementedKubernetesClusterKubernetesObjectQueryControllerServer) StreamKubernetesObjects(*model.StreamKubernetesClusterNamespaceKubernetesObjectsInput, KubernetesClusterKubernetesObjectQueryController_StreamKubernetesObjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamKubernetesObjects not implemented")
}
func (UnimplementedKubernetesClusterKubernetesObjectQueryControllerServer) FindPods(context.Context, *model.FindKubernetesClusterPodsInput) (*model1.Pods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPods not implemented")
}
func (UnimplementedKubernetesClusterKubernetesObjectQueryControllerServer) GetPod(context.Context, *model.KubernetesClusterKubernetesObject) (*model1.Pod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPod not implemented")
}
func (UnimplementedKubernetesClusterKubernetesObjectQueryControllerServer) StreamPodLogs(*model.StreamKubernetesClusterPodLogsInput, KubernetesClusterKubernetesObjectQueryController_StreamPodLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPodLogs not implemented")
}

// UnsafeKubernetesClusterKubernetesObjectQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubernetesClusterKubernetesObjectQueryControllerServer will
// result in compilation errors.
type UnsafeKubernetesClusterKubernetesObjectQueryControllerServer interface {
	mustEmbedUnimplementedKubernetesClusterKubernetesObjectQueryControllerServer()
}

func RegisterKubernetesClusterKubernetesObjectQueryControllerServer(s grpc.ServiceRegistrar, srv KubernetesClusterKubernetesObjectQueryControllerServer) {
	s.RegisterService(&KubernetesClusterKubernetesObjectQueryController_ServiceDesc, srv)
}

func _KubernetesClusterKubernetesObjectQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KubernetesClusterKubernetesObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterKubernetesObjectQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesClusterKubernetesObjectQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterKubernetesObjectQueryControllerServer).Get(ctx, req.(*model.KubernetesClusterKubernetesObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterKubernetesObjectQueryController_FindNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KubernetesClusterCredentialId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterKubernetesObjectQueryControllerServer).FindNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesClusterKubernetesObjectQueryController_FindNamespaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterKubernetesObjectQueryControllerServer).FindNamespaces(ctx, req.(*model.KubernetesClusterCredentialId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterKubernetesObjectQueryController_StreamKubernetesObjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.StreamKubernetesClusterNamespaceKubernetesObjectsInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KubernetesClusterKubernetesObjectQueryControllerServer).StreamKubernetesObjects(m, &kubernetesClusterKubernetesObjectQueryControllerStreamKubernetesObjectsServer{stream})
}

type KubernetesClusterKubernetesObjectQueryController_StreamKubernetesObjectsServer interface {
	Send(*model1.KubernetesObjects) error
	grpc.ServerStream
}

type kubernetesClusterKubernetesObjectQueryControllerStreamKubernetesObjectsServer struct {
	grpc.ServerStream
}

func (x *kubernetesClusterKubernetesObjectQueryControllerStreamKubernetesObjectsServer) Send(m *model1.KubernetesObjects) error {
	return x.ServerStream.SendMsg(m)
}

func _KubernetesClusterKubernetesObjectQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.FindKubernetesClusterPodsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterKubernetesObjectQueryControllerServer).FindPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesClusterKubernetesObjectQueryController_FindPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterKubernetesObjectQueryControllerServer).FindPods(ctx, req.(*model.FindKubernetesClusterPodsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterKubernetesObjectQueryController_GetPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KubernetesClusterKubernetesObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterKubernetesObjectQueryControllerServer).GetPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesClusterKubernetesObjectQueryController_GetPod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterKubernetesObjectQueryControllerServer).GetPod(ctx, req.(*model.KubernetesClusterKubernetesObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterKubernetesObjectQueryController_StreamPodLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.StreamKubernetesClusterPodLogsInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KubernetesClusterKubernetesObjectQueryControllerServer).StreamPodLogs(m, &kubernetesClusterKubernetesObjectQueryControllerStreamPodLogsServer{stream})
}

type KubernetesClusterKubernetesObjectQueryController_StreamPodLogsServer interface {
	Send(*model1.PodLogLine) error
	grpc.ServerStream
}

type kubernetesClusterKubernetesObjectQueryControllerStreamPodLogsServer struct {
	grpc.ServerStream
}

func (x *kubernetesClusterKubernetesObjectQueryControllerStreamPodLogsServer) Send(m *model1.PodLogLine) error {
	return x.ServerStream.SendMsg(m)
}

// KubernetesClusterKubernetesObjectQueryController_ServiceDesc is the grpc.ServiceDesc for KubernetesClusterKubernetesObjectQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubernetesClusterKubernetesObjectQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectQueryController",
	HandlerType: (*KubernetesClusterKubernetesObjectQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _KubernetesClusterKubernetesObjectQueryController_Get_Handler,
		},
		{
			MethodName: "findNamespaces",
			Handler:    _KubernetesClusterKubernetesObjectQueryController_FindNamespaces_Handler,
		},
		{
			MethodName: "findPods",
			Handler:    _KubernetesClusterKubernetesObjectQueryController_FindPods_Handler,
		},
		{
			MethodName: "getPod",
			Handler:    _KubernetesClusterKubernetesObjectQueryController_GetPod_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "streamKubernetesObjects",
			Handler:       _KubernetesClusterKubernetesObjectQueryController_StreamKubernetesObjects_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "streamPodLogs",
			Handler:       _KubernetesClusterKubernetesObjectQueryController_StreamPodLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/connect/kubernetesclustercredential/service/query.proto",
}
