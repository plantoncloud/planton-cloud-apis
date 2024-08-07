// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/k8sd2ops/v1/kubernetesobject/service/command.proto

package service

import (
	context "context"
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
	ApiResourceKubernetesObjectCommandController_Update_FullMethodName = "/cloud.planton.apis.k8sd2ops.v1.kubernetesobject.service.ApiResourceKubernetesObjectCommandController/update"
	ApiResourceKubernetesObjectCommandController_Delete_FullMethodName = "/cloud.planton.apis.k8sd2ops.v1.kubernetesobject.service.ApiResourceKubernetesObjectCommandController/delete"
)

// ApiResourceKubernetesObjectCommandControllerClient is the client API for ApiResourceKubernetesObjectCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiResourceKubernetesObjectCommandControllerClient interface {
	// update kubernetes object in a kube-cluster
	Update(ctx context.Context, in *model.UpdateKubernetesObjectInput, opts ...grpc.CallOption) (*model1.KubernetesObject, error)
	// delete kubernetes object in a kube-cluster
	Delete(ctx context.Context, in *model.ApiResourceKubernetesObject, opts ...grpc.CallOption) (*model1.KubernetesObject, error)
}

type apiResourceKubernetesObjectCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewApiResourceKubernetesObjectCommandControllerClient(cc grpc.ClientConnInterface) ApiResourceKubernetesObjectCommandControllerClient {
	return &apiResourceKubernetesObjectCommandControllerClient{cc}
}

func (c *apiResourceKubernetesObjectCommandControllerClient) Update(ctx context.Context, in *model.UpdateKubernetesObjectInput, opts ...grpc.CallOption) (*model1.KubernetesObject, error) {
	out := new(model1.KubernetesObject)
	err := c.cc.Invoke(ctx, ApiResourceKubernetesObjectCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiResourceKubernetesObjectCommandControllerClient) Delete(ctx context.Context, in *model.ApiResourceKubernetesObject, opts ...grpc.CallOption) (*model1.KubernetesObject, error) {
	out := new(model1.KubernetesObject)
	err := c.cc.Invoke(ctx, ApiResourceKubernetesObjectCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiResourceKubernetesObjectCommandControllerServer is the server API for ApiResourceKubernetesObjectCommandController service.
// All implementations should embed UnimplementedApiResourceKubernetesObjectCommandControllerServer
// for forward compatibility
type ApiResourceKubernetesObjectCommandControllerServer interface {
	// update kubernetes object in a kube-cluster
	Update(context.Context, *model.UpdateKubernetesObjectInput) (*model1.KubernetesObject, error)
	// delete kubernetes object in a kube-cluster
	Delete(context.Context, *model.ApiResourceKubernetesObject) (*model1.KubernetesObject, error)
}

// UnimplementedApiResourceKubernetesObjectCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedApiResourceKubernetesObjectCommandControllerServer struct {
}

func (UnimplementedApiResourceKubernetesObjectCommandControllerServer) Update(context.Context, *model.UpdateKubernetesObjectInput) (*model1.KubernetesObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedApiResourceKubernetesObjectCommandControllerServer) Delete(context.Context, *model.ApiResourceKubernetesObject) (*model1.KubernetesObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeApiResourceKubernetesObjectCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiResourceKubernetesObjectCommandControllerServer will
// result in compilation errors.
type UnsafeApiResourceKubernetesObjectCommandControllerServer interface {
	mustEmbedUnimplementedApiResourceKubernetesObjectCommandControllerServer()
}

func RegisterApiResourceKubernetesObjectCommandControllerServer(s grpc.ServiceRegistrar, srv ApiResourceKubernetesObjectCommandControllerServer) {
	s.RegisterService(&ApiResourceKubernetesObjectCommandController_ServiceDesc, srv)
}

func _ApiResourceKubernetesObjectCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.UpdateKubernetesObjectInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceKubernetesObjectCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceKubernetesObjectCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceKubernetesObjectCommandControllerServer).Update(ctx, req.(*model.UpdateKubernetesObjectInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiResourceKubernetesObjectCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ApiResourceKubernetesObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceKubernetesObjectCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceKubernetesObjectCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceKubernetesObjectCommandControllerServer).Delete(ctx, req.(*model.ApiResourceKubernetesObject))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiResourceKubernetesObjectCommandController_ServiceDesc is the grpc.ServiceDesc for ApiResourceKubernetesObjectCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiResourceKubernetesObjectCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.k8sd2ops.v1.kubernetesobject.service.ApiResourceKubernetesObjectCommandController",
	HandlerType: (*ApiResourceKubernetesObjectCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "update",
			Handler:    _ApiResourceKubernetesObjectCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _ApiResourceKubernetesObjectCommandController_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/k8sd2ops/v1/kubernetesobject/service/command.proto",
}
