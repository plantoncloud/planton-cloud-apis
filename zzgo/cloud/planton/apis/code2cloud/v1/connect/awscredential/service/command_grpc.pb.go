// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/connect/awscredential/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/connect/awscredential/model"
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
	AwsCredentialCommandController_Create_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.connect.awscredential.service.AwsCredentialCommandController/create"
	AwsCredentialCommandController_Update_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.connect.awscredential.service.AwsCredentialCommandController/update"
	AwsCredentialCommandController_Delete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.connect.awscredential.service.AwsCredentialCommandController/delete"
	AwsCredentialCommandController_Restore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.connect.awscredential.service.AwsCredentialCommandController/restore"
)

// AwsCredentialCommandControllerClient is the client API for AwsCredentialCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AwsCredentialCommandControllerClient interface {
	// create a aws-credential resource
	Create(ctx context.Context, in *model.AwsCredential, opts ...grpc.CallOption) (*model.AwsCredential, error)
	// update an existing aws-credential
	Update(ctx context.Context, in *model.AwsCredential, opts ...grpc.CallOption) (*model.AwsCredential, error)
	// delete a aws-credential that was previously created
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.AwsCredential, error)
	// restore a deleted aws-credential.
	Restore(ctx context.Context, in *model.AwsCredential, opts ...grpc.CallOption) (*model.AwsCredential, error)
}

type awsCredentialCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewAwsCredentialCommandControllerClient(cc grpc.ClientConnInterface) AwsCredentialCommandControllerClient {
	return &awsCredentialCommandControllerClient{cc}
}

func (c *awsCredentialCommandControllerClient) Create(ctx context.Context, in *model.AwsCredential, opts ...grpc.CallOption) (*model.AwsCredential, error) {
	out := new(model.AwsCredential)
	err := c.cc.Invoke(ctx, AwsCredentialCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awsCredentialCommandControllerClient) Update(ctx context.Context, in *model.AwsCredential, opts ...grpc.CallOption) (*model.AwsCredential, error) {
	out := new(model.AwsCredential)
	err := c.cc.Invoke(ctx, AwsCredentialCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awsCredentialCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.AwsCredential, error) {
	out := new(model.AwsCredential)
	err := c.cc.Invoke(ctx, AwsCredentialCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awsCredentialCommandControllerClient) Restore(ctx context.Context, in *model.AwsCredential, opts ...grpc.CallOption) (*model.AwsCredential, error) {
	out := new(model.AwsCredential)
	err := c.cc.Invoke(ctx, AwsCredentialCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AwsCredentialCommandControllerServer is the server API for AwsCredentialCommandController service.
// All implementations should embed UnimplementedAwsCredentialCommandControllerServer
// for forward compatibility
type AwsCredentialCommandControllerServer interface {
	// create a aws-credential resource
	Create(context.Context, *model.AwsCredential) (*model.AwsCredential, error)
	// update an existing aws-credential
	Update(context.Context, *model.AwsCredential) (*model.AwsCredential, error)
	// delete a aws-credential that was previously created
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.AwsCredential, error)
	// restore a deleted aws-credential.
	Restore(context.Context, *model.AwsCredential) (*model.AwsCredential, error)
}

// UnimplementedAwsCredentialCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedAwsCredentialCommandControllerServer struct {
}

func (UnimplementedAwsCredentialCommandControllerServer) Create(context.Context, *model.AwsCredential) (*model.AwsCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAwsCredentialCommandControllerServer) Update(context.Context, *model.AwsCredential) (*model.AwsCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAwsCredentialCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.AwsCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAwsCredentialCommandControllerServer) Restore(context.Context, *model.AwsCredential) (*model.AwsCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeAwsCredentialCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AwsCredentialCommandControllerServer will
// result in compilation errors.
type UnsafeAwsCredentialCommandControllerServer interface {
	mustEmbedUnimplementedAwsCredentialCommandControllerServer()
}

func RegisterAwsCredentialCommandControllerServer(s grpc.ServiceRegistrar, srv AwsCredentialCommandControllerServer) {
	s.RegisterService(&AwsCredentialCommandController_ServiceDesc, srv)
}

func _AwsCredentialCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AwsCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsCredentialCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AwsCredentialCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsCredentialCommandControllerServer).Create(ctx, req.(*model.AwsCredential))
	}
	return interceptor(ctx, in, info, handler)
}

func _AwsCredentialCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AwsCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsCredentialCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AwsCredentialCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsCredentialCommandControllerServer).Update(ctx, req.(*model.AwsCredential))
	}
	return interceptor(ctx, in, info, handler)
}

func _AwsCredentialCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsCredentialCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AwsCredentialCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsCredentialCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AwsCredentialCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AwsCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsCredentialCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AwsCredentialCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsCredentialCommandControllerServer).Restore(ctx, req.(*model.AwsCredential))
	}
	return interceptor(ctx, in, info, handler)
}

// AwsCredentialCommandController_ServiceDesc is the grpc.ServiceDesc for AwsCredentialCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AwsCredentialCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.connect.awscredential.service.AwsCredentialCommandController",
	HandlerType: (*AwsCredentialCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _AwsCredentialCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _AwsCredentialCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _AwsCredentialCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _AwsCredentialCommandController_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/connect/awscredential/service/command.proto",
}
