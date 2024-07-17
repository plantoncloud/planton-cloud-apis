// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/connect/kubernetesclustercredential/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/connect/kubernetesclustercredential/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/apiresources/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KubernetesClusterCredentialCommandController_Create_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterCredentialCommandController/create"
	KubernetesClusterCredentialCommandController_Update_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterCredentialCommandController/update"
	KubernetesClusterCredentialCommandController_Delete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterCredentialCommandController/delete"
	KubernetesClusterCredentialCommandController_Restore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterCredentialCommandController/restore"
)

// KubernetesClusterCredentialCommandControllerClient is the client API for KubernetesClusterCredentialCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubernetesClusterCredentialCommandControllerClient interface {
	// create a kubernetes-cluster-credential resource
	Create(ctx context.Context, in *model.KubernetesClusterCredential, opts ...grpc.CallOption) (*model.KubernetesClusterCredential, error)
	// update an existing kubernetes-cluster-credential
	Update(ctx context.Context, in *model.KubernetesClusterCredential, opts ...grpc.CallOption) (*model.KubernetesClusterCredential, error)
	// delete a kubernetes-cluster-credential that was previously created
	// warning: deleting a kubernetes-cluster-credential from planton-cloud destroys the resources created by planton-cloud in the account
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.KubernetesClusterCredential, error)
	// restore a deleted kubernetes-cluster-credential.
	Restore(ctx context.Context, in *model.KubernetesClusterCredential, opts ...grpc.CallOption) (*model.KubernetesClusterCredential, error)
}

type kubernetesClusterCredentialCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKubernetesClusterCredentialCommandControllerClient(cc grpc.ClientConnInterface) KubernetesClusterCredentialCommandControllerClient {
	return &kubernetesClusterCredentialCommandControllerClient{cc}
}

func (c *kubernetesClusterCredentialCommandControllerClient) Create(ctx context.Context, in *model.KubernetesClusterCredential, opts ...grpc.CallOption) (*model.KubernetesClusterCredential, error) {
	out := new(model.KubernetesClusterCredential)
	err := c.cc.Invoke(ctx, KubernetesClusterCredentialCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterCredentialCommandControllerClient) Update(ctx context.Context, in *model.KubernetesClusterCredential, opts ...grpc.CallOption) (*model.KubernetesClusterCredential, error) {
	out := new(model.KubernetesClusterCredential)
	err := c.cc.Invoke(ctx, KubernetesClusterCredentialCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterCredentialCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.KubernetesClusterCredential, error) {
	out := new(model.KubernetesClusterCredential)
	err := c.cc.Invoke(ctx, KubernetesClusterCredentialCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterCredentialCommandControllerClient) Restore(ctx context.Context, in *model.KubernetesClusterCredential, opts ...grpc.CallOption) (*model.KubernetesClusterCredential, error) {
	out := new(model.KubernetesClusterCredential)
	err := c.cc.Invoke(ctx, KubernetesClusterCredentialCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubernetesClusterCredentialCommandControllerServer is the server API for KubernetesClusterCredentialCommandController service.
// All implementations should embed UnimplementedKubernetesClusterCredentialCommandControllerServer
// for forward compatibility
type KubernetesClusterCredentialCommandControllerServer interface {
	// create a kubernetes-cluster-credential resource
	Create(context.Context, *model.KubernetesClusterCredential) (*model.KubernetesClusterCredential, error)
	// update an existing kubernetes-cluster-credential
	Update(context.Context, *model.KubernetesClusterCredential) (*model.KubernetesClusterCredential, error)
	// delete a kubernetes-cluster-credential that was previously created
	// warning: deleting a kubernetes-cluster-credential from planton-cloud destroys the resources created by planton-cloud in the account
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.KubernetesClusterCredential, error)
	// restore a deleted kubernetes-cluster-credential.
	Restore(context.Context, *model.KubernetesClusterCredential) (*model.KubernetesClusterCredential, error)
}

// UnimplementedKubernetesClusterCredentialCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKubernetesClusterCredentialCommandControllerServer struct {
}

func (UnimplementedKubernetesClusterCredentialCommandControllerServer) Create(context.Context, *model.KubernetesClusterCredential) (*model.KubernetesClusterCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedKubernetesClusterCredentialCommandControllerServer) Update(context.Context, *model.KubernetesClusterCredential) (*model.KubernetesClusterCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedKubernetesClusterCredentialCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.KubernetesClusterCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKubernetesClusterCredentialCommandControllerServer) Restore(context.Context, *model.KubernetesClusterCredential) (*model.KubernetesClusterCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeKubernetesClusterCredentialCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubernetesClusterCredentialCommandControllerServer will
// result in compilation errors.
type UnsafeKubernetesClusterCredentialCommandControllerServer interface {
	mustEmbedUnimplementedKubernetesClusterCredentialCommandControllerServer()
}

func RegisterKubernetesClusterCredentialCommandControllerServer(s grpc.ServiceRegistrar, srv KubernetesClusterCredentialCommandControllerServer) {
	s.RegisterService(&KubernetesClusterCredentialCommandController_ServiceDesc, srv)
}

func _KubernetesClusterCredentialCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KubernetesClusterCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterCredentialCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesClusterCredentialCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterCredentialCommandControllerServer).Create(ctx, req.(*model.KubernetesClusterCredential))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterCredentialCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KubernetesClusterCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterCredentialCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesClusterCredentialCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterCredentialCommandControllerServer).Update(ctx, req.(*model.KubernetesClusterCredential))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterCredentialCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterCredentialCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesClusterCredentialCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterCredentialCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterCredentialCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KubernetesClusterCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterCredentialCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesClusterCredentialCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterCredentialCommandControllerServer).Restore(ctx, req.(*model.KubernetesClusterCredential))
	}
	return interceptor(ctx, in, info, handler)
}

// KubernetesClusterCredentialCommandController_ServiceDesc is the grpc.ServiceDesc for KubernetesClusterCredentialCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubernetesClusterCredentialCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterCredentialCommandController",
	HandlerType: (*KubernetesClusterCredentialCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _KubernetesClusterCredentialCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _KubernetesClusterCredentialCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _KubernetesClusterCredentialCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _KubernetesClusterCredentialCommandController_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/connect/kubernetesclustercredential/service/command.proto",
}

const (
	KubernetesClusterKubernetesObjectCommandController_Update_FullMethodName                                  = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectCommandController/update"
	KubernetesClusterKubernetesObjectCommandController_Delete_FullMethodName                                  = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectCommandController/delete"
	KubernetesClusterKubernetesObjectCommandController_ExecIntoPodContainer_FullMethodName                    = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectCommandController/execIntoPodContainer"
	KubernetesClusterKubernetesObjectCommandController_BrowserExecIntoPodContainer_FullMethodName             = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectCommandController/browserExecIntoPodContainer"
	KubernetesClusterKubernetesObjectCommandController_BrowserExecuteNextCommandInPodContainer_FullMethodName = "/cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectCommandController/browserExecuteNextCommandInPodContainer"
)

// KubernetesClusterKubernetesObjectCommandControllerClient is the client API for KubernetesClusterKubernetesObjectCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubernetesClusterKubernetesObjectCommandControllerClient interface {
	// update kubernetes object in a kube-cluster
	Update(ctx context.Context, in *model.UpdateKubernetesClusterKubernetesObjectInput, opts ...grpc.CallOption) (*model2.KubernetesObject, error)
	// delete kubernetes object in a kube-cluster
	Delete(ctx context.Context, in *model.KubernetesClusterKubernetesObject, opts ...grpc.CallOption) (*model2.KubernetesObject, error)
	// mimic kubectl exec for a container on kube-cluster
	ExecIntoPodContainer(ctx context.Context, opts ...grpc.CallOption) (KubernetesClusterKubernetesObjectCommandController_ExecIntoPodContainerClient, error)
	// *
	// Mimic kubectl exec for a container on a kube-cluster from browsers.
	// This is a workaround to handle the limitation of browsers not supporting bi-directional gRPC streams.
	BrowserExecIntoPodContainer(ctx context.Context, in *model.ExecIntoKubernetesClusterPodContainerInput, opts ...grpc.CallOption) (KubernetesClusterKubernetesObjectCommandController_BrowserExecIntoPodContainerClient, error)
	// *
	// Send the next command to execute for kube-ctl exec.
	// This RPC is used to send input from the client (browser) which originally would have been sent in a bi-directional stream.
	// NOTE: Authorization will be handled based on the api-resource kind and id since the request input is same for
	// all other api-resources and kube-cluster resources.
	BrowserExecuteNextCommandInPodContainer(ctx context.Context, in *model.BrowserExecuteNextCommandInPodContainerInput, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type kubernetesClusterKubernetesObjectCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKubernetesClusterKubernetesObjectCommandControllerClient(cc grpc.ClientConnInterface) KubernetesClusterKubernetesObjectCommandControllerClient {
	return &kubernetesClusterKubernetesObjectCommandControllerClient{cc}
}

func (c *kubernetesClusterKubernetesObjectCommandControllerClient) Update(ctx context.Context, in *model.UpdateKubernetesClusterKubernetesObjectInput, opts ...grpc.CallOption) (*model2.KubernetesObject, error) {
	out := new(model2.KubernetesObject)
	err := c.cc.Invoke(ctx, KubernetesClusterKubernetesObjectCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterKubernetesObjectCommandControllerClient) Delete(ctx context.Context, in *model.KubernetesClusterKubernetesObject, opts ...grpc.CallOption) (*model2.KubernetesObject, error) {
	out := new(model2.KubernetesObject)
	err := c.cc.Invoke(ctx, KubernetesClusterKubernetesObjectCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterKubernetesObjectCommandControllerClient) ExecIntoPodContainer(ctx context.Context, opts ...grpc.CallOption) (KubernetesClusterKubernetesObjectCommandController_ExecIntoPodContainerClient, error) {
	stream, err := c.cc.NewStream(ctx, &KubernetesClusterKubernetesObjectCommandController_ServiceDesc.Streams[0], KubernetesClusterKubernetesObjectCommandController_ExecIntoPodContainer_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &kubernetesClusterKubernetesObjectCommandControllerExecIntoPodContainerClient{stream}
	return x, nil
}

type KubernetesClusterKubernetesObjectCommandController_ExecIntoPodContainerClient interface {
	Send(*model.ExecIntoKubernetesClusterPodContainerInput) error
	Recv() (*model2.ExecIntoPodContainerResponse, error)
	grpc.ClientStream
}

type kubernetesClusterKubernetesObjectCommandControllerExecIntoPodContainerClient struct {
	grpc.ClientStream
}

func (x *kubernetesClusterKubernetesObjectCommandControllerExecIntoPodContainerClient) Send(m *model.ExecIntoKubernetesClusterPodContainerInput) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kubernetesClusterKubernetesObjectCommandControllerExecIntoPodContainerClient) Recv() (*model2.ExecIntoPodContainerResponse, error) {
	m := new(model2.ExecIntoPodContainerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kubernetesClusterKubernetesObjectCommandControllerClient) BrowserExecIntoPodContainer(ctx context.Context, in *model.ExecIntoKubernetesClusterPodContainerInput, opts ...grpc.CallOption) (KubernetesClusterKubernetesObjectCommandController_BrowserExecIntoPodContainerClient, error) {
	stream, err := c.cc.NewStream(ctx, &KubernetesClusterKubernetesObjectCommandController_ServiceDesc.Streams[1], KubernetesClusterKubernetesObjectCommandController_BrowserExecIntoPodContainer_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &kubernetesClusterKubernetesObjectCommandControllerBrowserExecIntoPodContainerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KubernetesClusterKubernetesObjectCommandController_BrowserExecIntoPodContainerClient interface {
	Recv() (*model.BrowserExecIntoPodContainerResponse, error)
	grpc.ClientStream
}

type kubernetesClusterKubernetesObjectCommandControllerBrowserExecIntoPodContainerClient struct {
	grpc.ClientStream
}

func (x *kubernetesClusterKubernetesObjectCommandControllerBrowserExecIntoPodContainerClient) Recv() (*model.BrowserExecIntoPodContainerResponse, error) {
	m := new(model.BrowserExecIntoPodContainerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kubernetesClusterKubernetesObjectCommandControllerClient) BrowserExecuteNextCommandInPodContainer(ctx context.Context, in *model.BrowserExecuteNextCommandInPodContainerInput, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KubernetesClusterKubernetesObjectCommandController_BrowserExecuteNextCommandInPodContainer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubernetesClusterKubernetesObjectCommandControllerServer is the server API for KubernetesClusterKubernetesObjectCommandController service.
// All implementations should embed UnimplementedKubernetesClusterKubernetesObjectCommandControllerServer
// for forward compatibility
type KubernetesClusterKubernetesObjectCommandControllerServer interface {
	// update kubernetes object in a kube-cluster
	Update(context.Context, *model.UpdateKubernetesClusterKubernetesObjectInput) (*model2.KubernetesObject, error)
	// delete kubernetes object in a kube-cluster
	Delete(context.Context, *model.KubernetesClusterKubernetesObject) (*model2.KubernetesObject, error)
	// mimic kubectl exec for a container on kube-cluster
	ExecIntoPodContainer(KubernetesClusterKubernetesObjectCommandController_ExecIntoPodContainerServer) error
	// *
	// Mimic kubectl exec for a container on a kube-cluster from browsers.
	// This is a workaround to handle the limitation of browsers not supporting bi-directional gRPC streams.
	BrowserExecIntoPodContainer(*model.ExecIntoKubernetesClusterPodContainerInput, KubernetesClusterKubernetesObjectCommandController_BrowserExecIntoPodContainerServer) error
	// *
	// Send the next command to execute for kube-ctl exec.
	// This RPC is used to send input from the client (browser) which originally would have been sent in a bi-directional stream.
	// NOTE: Authorization will be handled based on the api-resource kind and id since the request input is same for
	// all other api-resources and kube-cluster resources.
	BrowserExecuteNextCommandInPodContainer(context.Context, *model.BrowserExecuteNextCommandInPodContainerInput) (*emptypb.Empty, error)
}

// UnimplementedKubernetesClusterKubernetesObjectCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKubernetesClusterKubernetesObjectCommandControllerServer struct {
}

func (UnimplementedKubernetesClusterKubernetesObjectCommandControllerServer) Update(context.Context, *model.UpdateKubernetesClusterKubernetesObjectInput) (*model2.KubernetesObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedKubernetesClusterKubernetesObjectCommandControllerServer) Delete(context.Context, *model.KubernetesClusterKubernetesObject) (*model2.KubernetesObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKubernetesClusterKubernetesObjectCommandControllerServer) ExecIntoPodContainer(KubernetesClusterKubernetesObjectCommandController_ExecIntoPodContainerServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecIntoPodContainer not implemented")
}
func (UnimplementedKubernetesClusterKubernetesObjectCommandControllerServer) BrowserExecIntoPodContainer(*model.ExecIntoKubernetesClusterPodContainerInput, KubernetesClusterKubernetesObjectCommandController_BrowserExecIntoPodContainerServer) error {
	return status.Errorf(codes.Unimplemented, "method BrowserExecIntoPodContainer not implemented")
}
func (UnimplementedKubernetesClusterKubernetesObjectCommandControllerServer) BrowserExecuteNextCommandInPodContainer(context.Context, *model.BrowserExecuteNextCommandInPodContainerInput) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BrowserExecuteNextCommandInPodContainer not implemented")
}

// UnsafeKubernetesClusterKubernetesObjectCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubernetesClusterKubernetesObjectCommandControllerServer will
// result in compilation errors.
type UnsafeKubernetesClusterKubernetesObjectCommandControllerServer interface {
	mustEmbedUnimplementedKubernetesClusterKubernetesObjectCommandControllerServer()
}

func RegisterKubernetesClusterKubernetesObjectCommandControllerServer(s grpc.ServiceRegistrar, srv KubernetesClusterKubernetesObjectCommandControllerServer) {
	s.RegisterService(&KubernetesClusterKubernetesObjectCommandController_ServiceDesc, srv)
}

func _KubernetesClusterKubernetesObjectCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.UpdateKubernetesClusterKubernetesObjectInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterKubernetesObjectCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesClusterKubernetesObjectCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterKubernetesObjectCommandControllerServer).Update(ctx, req.(*model.UpdateKubernetesClusterKubernetesObjectInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterKubernetesObjectCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KubernetesClusterKubernetesObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterKubernetesObjectCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesClusterKubernetesObjectCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterKubernetesObjectCommandControllerServer).Delete(ctx, req.(*model.KubernetesClusterKubernetesObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterKubernetesObjectCommandController_ExecIntoPodContainer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KubernetesClusterKubernetesObjectCommandControllerServer).ExecIntoPodContainer(&kubernetesClusterKubernetesObjectCommandControllerExecIntoPodContainerServer{stream})
}

type KubernetesClusterKubernetesObjectCommandController_ExecIntoPodContainerServer interface {
	Send(*model2.ExecIntoPodContainerResponse) error
	Recv() (*model.ExecIntoKubernetesClusterPodContainerInput, error)
	grpc.ServerStream
}

type kubernetesClusterKubernetesObjectCommandControllerExecIntoPodContainerServer struct {
	grpc.ServerStream
}

func (x *kubernetesClusterKubernetesObjectCommandControllerExecIntoPodContainerServer) Send(m *model2.ExecIntoPodContainerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kubernetesClusterKubernetesObjectCommandControllerExecIntoPodContainerServer) Recv() (*model.ExecIntoKubernetesClusterPodContainerInput, error) {
	m := new(model.ExecIntoKubernetesClusterPodContainerInput)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KubernetesClusterKubernetesObjectCommandController_BrowserExecIntoPodContainer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.ExecIntoKubernetesClusterPodContainerInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KubernetesClusterKubernetesObjectCommandControllerServer).BrowserExecIntoPodContainer(m, &kubernetesClusterKubernetesObjectCommandControllerBrowserExecIntoPodContainerServer{stream})
}

type KubernetesClusterKubernetesObjectCommandController_BrowserExecIntoPodContainerServer interface {
	Send(*model.BrowserExecIntoPodContainerResponse) error
	grpc.ServerStream
}

type kubernetesClusterKubernetesObjectCommandControllerBrowserExecIntoPodContainerServer struct {
	grpc.ServerStream
}

func (x *kubernetesClusterKubernetesObjectCommandControllerBrowserExecIntoPodContainerServer) Send(m *model.BrowserExecIntoPodContainerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KubernetesClusterKubernetesObjectCommandController_BrowserExecuteNextCommandInPodContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.BrowserExecuteNextCommandInPodContainerInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterKubernetesObjectCommandControllerServer).BrowserExecuteNextCommandInPodContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubernetesClusterKubernetesObjectCommandController_BrowserExecuteNextCommandInPodContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterKubernetesObjectCommandControllerServer).BrowserExecuteNextCommandInPodContainer(ctx, req.(*model.BrowserExecuteNextCommandInPodContainerInput))
	}
	return interceptor(ctx, in, info, handler)
}

// KubernetesClusterKubernetesObjectCommandController_ServiceDesc is the grpc.ServiceDesc for KubernetesClusterKubernetesObjectCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubernetesClusterKubernetesObjectCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.connect.kubernetesclustercredential.service.KubernetesClusterKubernetesObjectCommandController",
	HandlerType: (*KubernetesClusterKubernetesObjectCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "update",
			Handler:    _KubernetesClusterKubernetesObjectCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _KubernetesClusterKubernetesObjectCommandController_Delete_Handler,
		},
		{
			MethodName: "browserExecuteNextCommandInPodContainer",
			Handler:    _KubernetesClusterKubernetesObjectCommandController_BrowserExecuteNextCommandInPodContainer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execIntoPodContainer",
			Handler:       _KubernetesClusterKubernetesObjectCommandController_ExecIntoPodContainer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "browserExecIntoPodContainer",
			Handler:       _KubernetesClusterKubernetesObjectCommandController_BrowserExecIntoPodContainer_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/connect/kubernetesclustercredential/service/command.proto",
}