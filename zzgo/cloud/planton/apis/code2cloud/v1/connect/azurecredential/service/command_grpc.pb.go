// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/connect/azurecredential/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/connect/azurecredential/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AzureCredentialCommandController_Create_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.connect.azurecredential.service.AzureCredentialCommandController/create"
	AzureCredentialCommandController_Update_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.connect.azurecredential.service.AzureCredentialCommandController/update"
	AzureCredentialCommandController_Delete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.connect.azurecredential.service.AzureCredentialCommandController/delete"
	AzureCredentialCommandController_Restore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.connect.azurecredential.service.AzureCredentialCommandController/restore"
)

// AzureCredentialCommandControllerClient is the client API for AzureCredentialCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AzureCredentialCommandControllerClient interface {
	// create a azure-credential resource
	Create(ctx context.Context, in *model.AzureCredential, opts ...grpc.CallOption) (*model.AzureCredential, error)
	// update an existing azure-credential
	Update(ctx context.Context, in *model.AzureCredential, opts ...grpc.CallOption) (*model.AzureCredential, error)
	// delete a azure-credential that was previously created
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.AzureCredential, error)
	// restore a deleted azure-credential.
	Restore(ctx context.Context, in *model.AzureCredential, opts ...grpc.CallOption) (*model.AzureCredential, error)
}

type azureCredentialCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewAzureCredentialCommandControllerClient(cc grpc.ClientConnInterface) AzureCredentialCommandControllerClient {
	return &azureCredentialCommandControllerClient{cc}
}

func (c *azureCredentialCommandControllerClient) Create(ctx context.Context, in *model.AzureCredential, opts ...grpc.CallOption) (*model.AzureCredential, error) {
	out := new(model.AzureCredential)
	err := c.cc.Invoke(ctx, AzureCredentialCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *azureCredentialCommandControllerClient) Update(ctx context.Context, in *model.AzureCredential, opts ...grpc.CallOption) (*model.AzureCredential, error) {
	out := new(model.AzureCredential)
	err := c.cc.Invoke(ctx, AzureCredentialCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *azureCredentialCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.AzureCredential, error) {
	out := new(model.AzureCredential)
	err := c.cc.Invoke(ctx, AzureCredentialCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *azureCredentialCommandControllerClient) Restore(ctx context.Context, in *model.AzureCredential, opts ...grpc.CallOption) (*model.AzureCredential, error) {
	out := new(model.AzureCredential)
	err := c.cc.Invoke(ctx, AzureCredentialCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AzureCredentialCommandControllerServer is the server API for AzureCredentialCommandController service.
// All implementations should embed UnimplementedAzureCredentialCommandControllerServer
// for forward compatibility
type AzureCredentialCommandControllerServer interface {
	// create a azure-credential resource
	Create(context.Context, *model.AzureCredential) (*model.AzureCredential, error)
	// update an existing azure-credential
	Update(context.Context, *model.AzureCredential) (*model.AzureCredential, error)
	// delete a azure-credential that was previously created
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.AzureCredential, error)
	// restore a deleted azure-credential.
	Restore(context.Context, *model.AzureCredential) (*model.AzureCredential, error)
}

// UnimplementedAzureCredentialCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedAzureCredentialCommandControllerServer struct {
}

func (UnimplementedAzureCredentialCommandControllerServer) Create(context.Context, *model.AzureCredential) (*model.AzureCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAzureCredentialCommandControllerServer) Update(context.Context, *model.AzureCredential) (*model.AzureCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAzureCredentialCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.AzureCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAzureCredentialCommandControllerServer) Restore(context.Context, *model.AzureCredential) (*model.AzureCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeAzureCredentialCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AzureCredentialCommandControllerServer will
// result in compilation errors.
type UnsafeAzureCredentialCommandControllerServer interface {
	mustEmbedUnimplementedAzureCredentialCommandControllerServer()
}

func RegisterAzureCredentialCommandControllerServer(s grpc.ServiceRegistrar, srv AzureCredentialCommandControllerServer) {
	s.RegisterService(&AzureCredentialCommandController_ServiceDesc, srv)
}

func _AzureCredentialCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AzureCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AzureCredentialCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AzureCredentialCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AzureCredentialCommandControllerServer).Create(ctx, req.(*model.AzureCredential))
	}
	return interceptor(ctx, in, info, handler)
}

func _AzureCredentialCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AzureCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AzureCredentialCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AzureCredentialCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AzureCredentialCommandControllerServer).Update(ctx, req.(*model.AzureCredential))
	}
	return interceptor(ctx, in, info, handler)
}

func _AzureCredentialCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AzureCredentialCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AzureCredentialCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AzureCredentialCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AzureCredentialCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AzureCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AzureCredentialCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AzureCredentialCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AzureCredentialCommandControllerServer).Restore(ctx, req.(*model.AzureCredential))
	}
	return interceptor(ctx, in, info, handler)
}

// AzureCredentialCommandController_ServiceDesc is the grpc.ServiceDesc for AzureCredentialCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AzureCredentialCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.connect.azurecredential.service.AzureCredentialCommandController",
	HandlerType: (*AzureCredentialCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _AzureCredentialCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _AzureCredentialCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _AzureCredentialCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _AzureCredentialCommandController_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/connect/azurecredential/service/command.proto",
}