// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/integration/v1/kubernetes/kubernetesobject/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/kubernetesobject/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KubernetesObjectCommandController_Update_FullMethodName               = "/cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.service.KubernetesObjectCommandController/update"
	KubernetesObjectCommandController_Delete_FullMethodName               = "/cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.service.KubernetesObjectCommandController/delete"
	KubernetesObjectCommandController_ExecIntoPodContainer_FullMethodName = "/cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.service.KubernetesObjectCommandController/execIntoPodContainer"
)

// KubernetesObjectCommandControllerClient is the client API for KubernetesObjectCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubernetesObjectCommandControllerClient interface {
	Update(ctx context.Context, in *model.UpdateKubernetesObjectWithKubeConfigInput, opts ...grpc.CallOption) (*model.KubernetesObject, error)
	Delete(ctx context.Context, in *model.DeleteKubernetesObjectWithKubeConfigInput, opts ...grpc.CallOption) (*model.KubernetesObject, error)
	// mimic kubectl exec
	ExecIntoPodContainer(ctx context.Context, opts ...grpc.CallOption) (KubernetesObjectCommandController_ExecIntoPodContainerClient, error)
}

type kubernetesObjectCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKubernetesObjectCommandControllerClient(cc grpc.ClientConnInterface) KubernetesObjectCommandControllerClient {
	return &kubernetesObjectCommandControllerClient{cc}
}

func (c *kubernetesObjectCommandControllerClient) Update(ctx context.Context, in *model.UpdateKubernetesObjectWithKubeConfigInput, opts ...grpc.CallOption) (*model.KubernetesObject, error) {
	out := new(model.KubernetesObject)
	err := c.cc.Invoke(ctx, KubernetesObjectCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesObjectCommandControllerClient) Delete(ctx context.Context, in *model.DeleteKubernetesObjectWithKubeConfigInput, opts ...grpc.CallOption) (*model.KubernetesObject, error) {
	out := new(model.KubernetesObject)
	err := c.cc.Invoke(ctx, KubernetesObjectCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesObjectCommandControllerClient) ExecIntoPodContainer(ctx context.Context, opts ...grpc.CallOption) (KubernetesObjectCommandController_ExecIntoPodContainerClient, error) {
	stream, err := c.cc.NewStream(ctx, &KubernetesObjectCommandController_ServiceDesc.Streams[0], KubernetesObjectCommandController_ExecIntoPodContainer_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &kubernetesObjectCommandControllerExecIntoPodContainerClient{stream}
	return x, nil
}

type KubernetesObjectCommandController_ExecIntoPodContainerClient interface {
	Send(*model.ExecIntoPodContainerWithKubeConfigInput) error
	Recv() (*model.ExecIntoPodContainerResponse, error)
	grpc.ClientStream
}

type kubernetesObjectCommandControllerExecIntoPodContainerClient struct {
	grpc.ClientStream
}

func (x *kubernetesObjectCommandControllerExecIntoPodContainerClient) Send(m *model.ExecIntoPodContainerWithKubeConfigInput) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kubernetesObjectCommandControllerExecIntoPodContainerClient) Recv() (*model.ExecIntoPodContainerResponse, error) {
	m := new(model.ExecIntoPodContainerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KubernetesObjectCommandControllerServer is the server API for KubernetesObjectCommandController service.
// All implementations should embed UnimplementedKubernetesObjectCommandControllerServer
// for forward compatibility
type KubernetesObjectCommandControllerServer interface {
	Update(context.Context, *model.UpdateKubernetesObjectWithKubeConfigInput) (*model.KubernetesObject, error)
	Delete(context.Context, *model.DeleteKubernetesObjectWithKubeConfigInput) (*model.KubernetesObject, error)
	// mimic kubectl exec
	ExecIntoPodContainer(KubernetesObjectCommandController_ExecIntoPodContainerServer) error
}

// UnimplementedKubernetesObjectCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKubernetesObjectCommandControllerServer struct {
}

func (UnimplementedKubernetesObjectCommandControllerServer) Update(context.Context, *model.UpdateKubernetesObjectWithKubeConfigInput) (*model.KubernetesObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedKubernetesObjectCommandControllerServer) Delete(context.Context, *model.DeleteKubernetesObjectWithKubeConfigInput) (*model.KubernetesObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKubernetesObjectCommandControllerServer) ExecIntoPodContainer(KubernetesObjectCommandController_ExecIntoPodContainerServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecIntoPodContainer not implemented")
}

// UnsafeKubernetesObjectCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubernetesObjectCommandControllerServer will
// result in compilation errors.
type UnsafeKubernetesObjectCommandControllerServer interface {
	mustEmbedUnimplementedKubernetesObjectCommandControllerServer()
}

func RegisterKubernetesObjectCommandControllerServer(s grpc.ServiceRegistrar, srv KubernetesObjectCommandControllerServer) {
	s.RegisterService(&KubernetesObjectCommandController_ServiceDesc, srv)
}

func _KubernetesObjectCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.UpdateKubernetesObjectWithKubeConfigInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesObjectCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesObjectCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesObjectCommandControllerServer).Update(ctx, req.(*model.UpdateKubernetesObjectWithKubeConfigInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesObjectCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.DeleteKubernetesObjectWithKubeConfigInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesObjectCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesObjectCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesObjectCommandControllerServer).Delete(ctx, req.(*model.DeleteKubernetesObjectWithKubeConfigInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesObjectCommandController_ExecIntoPodContainer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KubernetesObjectCommandControllerServer).ExecIntoPodContainer(&kubernetesObjectCommandControllerExecIntoPodContainerServer{stream})
}

type KubernetesObjectCommandController_ExecIntoPodContainerServer interface {
	Send(*model.ExecIntoPodContainerResponse) error
	Recv() (*model.ExecIntoPodContainerWithKubeConfigInput, error)
	grpc.ServerStream
}

type kubernetesObjectCommandControllerExecIntoPodContainerServer struct {
	grpc.ServerStream
}

func (x *kubernetesObjectCommandControllerExecIntoPodContainerServer) Send(m *model.ExecIntoPodContainerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kubernetesObjectCommandControllerExecIntoPodContainerServer) Recv() (*model.ExecIntoPodContainerWithKubeConfigInput, error) {
	m := new(model.ExecIntoPodContainerWithKubeConfigInput)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KubernetesObjectCommandController_ServiceDesc is the grpc.ServiceDesc for KubernetesObjectCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubernetesObjectCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.service.KubernetesObjectCommandController",
	HandlerType: (*KubernetesObjectCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "update",
			Handler:    _KubernetesObjectCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _KubernetesObjectCommandController_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execIntoPodContainer",
			Handler:       _KubernetesObjectCommandController_ExecIntoPodContainer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/integration/v1/kubernetes/kubernetesobject/service/command.proto",
}