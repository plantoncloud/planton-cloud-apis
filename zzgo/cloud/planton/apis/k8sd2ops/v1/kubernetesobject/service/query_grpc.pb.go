// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/k8sd2ops/v1/kubernetesobject/service/query.proto

package service

import (
	context "context"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/kubernetesclustercredential/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/kubernetesobject/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/k8sd2ops/v1/kubernetesobject/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ApiResourceKubernetesObjectQueryController_Get_FullMethodName                     = "/cloud.planton.apis.k8sd2ops.v1.service.ApiResourceKubernetesObjectQueryController/get"
	ApiResourceKubernetesObjectQueryController_FindNamespaces_FullMethodName          = "/cloud.planton.apis.k8sd2ops.v1.service.ApiResourceKubernetesObjectQueryController/findNamespaces"
	ApiResourceKubernetesObjectQueryController_StreamKubernetesObjects_FullMethodName = "/cloud.planton.apis.k8sd2ops.v1.service.ApiResourceKubernetesObjectQueryController/streamKubernetesObjects"
	ApiResourceKubernetesObjectQueryController_FindPods_FullMethodName                = "/cloud.planton.apis.k8sd2ops.v1.service.ApiResourceKubernetesObjectQueryController/findPods"
	ApiResourceKubernetesObjectQueryController_GetPod_FullMethodName                  = "/cloud.planton.apis.k8sd2ops.v1.service.ApiResourceKubernetesObjectQueryController/getPod"
)

// ApiResourceKubernetesObjectQueryControllerClient is the client API for ApiResourceKubernetesObjectQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiResourceKubernetesObjectQueryControllerClient interface {
	// get detailed object of a kubernetes object
	Get(ctx context.Context, in *model.ApiResourceKubernetesObject, opts ...grpc.CallOption) (*model1.KubernetesObjectDetail, error)
	// find list of namespaces on a kubernetes cluster
	FindNamespaces(ctx context.Context, in *model2.KubernetesClusterCredentialId, opts ...grpc.CallOption) (*model1.KubernetesNamespaceList, error)
	// stream all kubernetes objects from a kubernetes namespace in kube-cluster.
	// this is a streaming rpc since the lookup involves several kubernetes api-calls to fetch all the kubernetes-api-resources.
	// because of high number of api calls to upstream kubernetes cluster, the response is streamed to the client.
	StreamKubernetesObjects(ctx context.Context, in *model.StreamKubernetesObjectsInNamespaceInput, opts ...grpc.CallOption) (ApiResourceKubernetesObjectQueryController_StreamKubernetesObjectsClient, error)
	// find list of pods in a kube-cluster on the specified filters
	FindPods(ctx context.Context, in *model.FindKubernetesPodsInput, opts ...grpc.CallOption) (*model1.KubernetesPodList, error)
	// get details of a pod
	GetPod(ctx context.Context, in *model.ApiResourceKubernetesObject, opts ...grpc.CallOption) (*model1.KubernetesPod, error)
}

type apiResourceKubernetesObjectQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewApiResourceKubernetesObjectQueryControllerClient(cc grpc.ClientConnInterface) ApiResourceKubernetesObjectQueryControllerClient {
	return &apiResourceKubernetesObjectQueryControllerClient{cc}
}

func (c *apiResourceKubernetesObjectQueryControllerClient) Get(ctx context.Context, in *model.ApiResourceKubernetesObject, opts ...grpc.CallOption) (*model1.KubernetesObjectDetail, error) {
	out := new(model1.KubernetesObjectDetail)
	err := c.cc.Invoke(ctx, ApiResourceKubernetesObjectQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiResourceKubernetesObjectQueryControllerClient) FindNamespaces(ctx context.Context, in *model2.KubernetesClusterCredentialId, opts ...grpc.CallOption) (*model1.KubernetesNamespaceList, error) {
	out := new(model1.KubernetesNamespaceList)
	err := c.cc.Invoke(ctx, ApiResourceKubernetesObjectQueryController_FindNamespaces_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiResourceKubernetesObjectQueryControllerClient) StreamKubernetesObjects(ctx context.Context, in *model.StreamKubernetesObjectsInNamespaceInput, opts ...grpc.CallOption) (ApiResourceKubernetesObjectQueryController_StreamKubernetesObjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ApiResourceKubernetesObjectQueryController_ServiceDesc.Streams[0], ApiResourceKubernetesObjectQueryController_StreamKubernetesObjects_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &apiResourceKubernetesObjectQueryControllerStreamKubernetesObjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApiResourceKubernetesObjectQueryController_StreamKubernetesObjectsClient interface {
	Recv() (*model1.KubernetesObjectList, error)
	grpc.ClientStream
}

type apiResourceKubernetesObjectQueryControllerStreamKubernetesObjectsClient struct {
	grpc.ClientStream
}

func (x *apiResourceKubernetesObjectQueryControllerStreamKubernetesObjectsClient) Recv() (*model1.KubernetesObjectList, error) {
	m := new(model1.KubernetesObjectList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *apiResourceKubernetesObjectQueryControllerClient) FindPods(ctx context.Context, in *model.FindKubernetesPodsInput, opts ...grpc.CallOption) (*model1.KubernetesPodList, error) {
	out := new(model1.KubernetesPodList)
	err := c.cc.Invoke(ctx, ApiResourceKubernetesObjectQueryController_FindPods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiResourceKubernetesObjectQueryControllerClient) GetPod(ctx context.Context, in *model.ApiResourceKubernetesObject, opts ...grpc.CallOption) (*model1.KubernetesPod, error) {
	out := new(model1.KubernetesPod)
	err := c.cc.Invoke(ctx, ApiResourceKubernetesObjectQueryController_GetPod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiResourceKubernetesObjectQueryControllerServer is the server API for ApiResourceKubernetesObjectQueryController service.
// All implementations should embed UnimplementedApiResourceKubernetesObjectQueryControllerServer
// for forward compatibility
type ApiResourceKubernetesObjectQueryControllerServer interface {
	// get detailed object of a kubernetes object
	Get(context.Context, *model.ApiResourceKubernetesObject) (*model1.KubernetesObjectDetail, error)
	// find list of namespaces on a kubernetes cluster
	FindNamespaces(context.Context, *model2.KubernetesClusterCredentialId) (*model1.KubernetesNamespaceList, error)
	// stream all kubernetes objects from a kubernetes namespace in kube-cluster.
	// this is a streaming rpc since the lookup involves several kubernetes api-calls to fetch all the kubernetes-api-resources.
	// because of high number of api calls to upstream kubernetes cluster, the response is streamed to the client.
	StreamKubernetesObjects(*model.StreamKubernetesObjectsInNamespaceInput, ApiResourceKubernetesObjectQueryController_StreamKubernetesObjectsServer) error
	// find list of pods in a kube-cluster on the specified filters
	FindPods(context.Context, *model.FindKubernetesPodsInput) (*model1.KubernetesPodList, error)
	// get details of a pod
	GetPod(context.Context, *model.ApiResourceKubernetesObject) (*model1.KubernetesPod, error)
}

// UnimplementedApiResourceKubernetesObjectQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedApiResourceKubernetesObjectQueryControllerServer struct {
}

func (UnimplementedApiResourceKubernetesObjectQueryControllerServer) Get(context.Context, *model.ApiResourceKubernetesObject) (*model1.KubernetesObjectDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedApiResourceKubernetesObjectQueryControllerServer) FindNamespaces(context.Context, *model2.KubernetesClusterCredentialId) (*model1.KubernetesNamespaceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNamespaces not implemented")
}
func (UnimplementedApiResourceKubernetesObjectQueryControllerServer) StreamKubernetesObjects(*model.StreamKubernetesObjectsInNamespaceInput, ApiResourceKubernetesObjectQueryController_StreamKubernetesObjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamKubernetesObjects not implemented")
}
func (UnimplementedApiResourceKubernetesObjectQueryControllerServer) FindPods(context.Context, *model.FindKubernetesPodsInput) (*model1.KubernetesPodList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPods not implemented")
}
func (UnimplementedApiResourceKubernetesObjectQueryControllerServer) GetPod(context.Context, *model.ApiResourceKubernetesObject) (*model1.KubernetesPod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPod not implemented")
}

// UnsafeApiResourceKubernetesObjectQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiResourceKubernetesObjectQueryControllerServer will
// result in compilation errors.
type UnsafeApiResourceKubernetesObjectQueryControllerServer interface {
	mustEmbedUnimplementedApiResourceKubernetesObjectQueryControllerServer()
}

func RegisterApiResourceKubernetesObjectQueryControllerServer(s grpc.ServiceRegistrar, srv ApiResourceKubernetesObjectQueryControllerServer) {
	s.RegisterService(&ApiResourceKubernetesObjectQueryController_ServiceDesc, srv)
}

func _ApiResourceKubernetesObjectQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ApiResourceKubernetesObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceKubernetesObjectQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceKubernetesObjectQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceKubernetesObjectQueryControllerServer).Get(ctx, req.(*model.ApiResourceKubernetesObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiResourceKubernetesObjectQueryController_FindNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.KubernetesClusterCredentialId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceKubernetesObjectQueryControllerServer).FindNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceKubernetesObjectQueryController_FindNamespaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceKubernetesObjectQueryControllerServer).FindNamespaces(ctx, req.(*model2.KubernetesClusterCredentialId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiResourceKubernetesObjectQueryController_StreamKubernetesObjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.StreamKubernetesObjectsInNamespaceInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApiResourceKubernetesObjectQueryControllerServer).StreamKubernetesObjects(m, &apiResourceKubernetesObjectQueryControllerStreamKubernetesObjectsServer{stream})
}

type ApiResourceKubernetesObjectQueryController_StreamKubernetesObjectsServer interface {
	Send(*model1.KubernetesObjectList) error
	grpc.ServerStream
}

type apiResourceKubernetesObjectQueryControllerStreamKubernetesObjectsServer struct {
	grpc.ServerStream
}

func (x *apiResourceKubernetesObjectQueryControllerStreamKubernetesObjectsServer) Send(m *model1.KubernetesObjectList) error {
	return x.ServerStream.SendMsg(m)
}

func _ApiResourceKubernetesObjectQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.FindKubernetesPodsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceKubernetesObjectQueryControllerServer).FindPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceKubernetesObjectQueryController_FindPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceKubernetesObjectQueryControllerServer).FindPods(ctx, req.(*model.FindKubernetesPodsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiResourceKubernetesObjectQueryController_GetPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ApiResourceKubernetesObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceKubernetesObjectQueryControllerServer).GetPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceKubernetesObjectQueryController_GetPod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceKubernetesObjectQueryControllerServer).GetPod(ctx, req.(*model.ApiResourceKubernetesObject))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiResourceKubernetesObjectQueryController_ServiceDesc is the grpc.ServiceDesc for ApiResourceKubernetesObjectQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiResourceKubernetesObjectQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.k8sd2ops.v1.service.ApiResourceKubernetesObjectQueryController",
	HandlerType: (*ApiResourceKubernetesObjectQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _ApiResourceKubernetesObjectQueryController_Get_Handler,
		},
		{
			MethodName: "findNamespaces",
			Handler:    _ApiResourceKubernetesObjectQueryController_FindNamespaces_Handler,
		},
		{
			MethodName: "findPods",
			Handler:    _ApiResourceKubernetesObjectQueryController_FindPods_Handler,
		},
		{
			MethodName: "getPod",
			Handler:    _ApiResourceKubernetesObjectQueryController_GetPod_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "streamKubernetesObjects",
			Handler:       _ApiResourceKubernetesObjectQueryController_StreamKubernetesObjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/k8sd2ops/v1/kubernetesobject/service/query.proto",
}