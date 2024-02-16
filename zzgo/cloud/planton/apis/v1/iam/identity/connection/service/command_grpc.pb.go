// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/iam/identity/connection/service/command.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/identity/connection/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IdentityConnectionCommandController_Create_FullMethodName  = "/cloud.planton.apis.v1.iam.identity.connection.service.IdentityConnectionCommandController/create"
	IdentityConnectionCommandController_Update_FullMethodName  = "/cloud.planton.apis.v1.iam.identity.connection.service.IdentityConnectionCommandController/update"
	IdentityConnectionCommandController_Delete_FullMethodName  = "/cloud.planton.apis.v1.iam.identity.connection.service.IdentityConnectionCommandController/delete"
	IdentityConnectionCommandController_Restore_FullMethodName = "/cloud.planton.apis.v1.iam.identity.connection.service.IdentityConnectionCommandController/restore"
)

// IdentityConnectionCommandControllerClient is the client API for IdentityConnectionCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityConnectionCommandControllerClient interface {
	// create new identity connection
	Create(ctx context.Context, in *model.IdentityConnection, opts ...grpc.CallOption) (*model.IdentityConnection, error)
	// update an existing identity connection
	Update(ctx context.Context, in *model.IdentityConnection, opts ...grpc.CallOption) (*model.IdentityConnection, error)
	// delete an existing identity connection
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.IdentityConnection, error)
	// restore an existing identity connection
	Restore(ctx context.Context, in *model.IdentityConnection, opts ...grpc.CallOption) (*model.IdentityConnection, error)
}

type identityConnectionCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityConnectionCommandControllerClient(cc grpc.ClientConnInterface) IdentityConnectionCommandControllerClient {
	return &identityConnectionCommandControllerClient{cc}
}

func (c *identityConnectionCommandControllerClient) Create(ctx context.Context, in *model.IdentityConnection, opts ...grpc.CallOption) (*model.IdentityConnection, error) {
	out := new(model.IdentityConnection)
	err := c.cc.Invoke(ctx, IdentityConnectionCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityConnectionCommandControllerClient) Update(ctx context.Context, in *model.IdentityConnection, opts ...grpc.CallOption) (*model.IdentityConnection, error) {
	out := new(model.IdentityConnection)
	err := c.cc.Invoke(ctx, IdentityConnectionCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityConnectionCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.IdentityConnection, error) {
	out := new(model.IdentityConnection)
	err := c.cc.Invoke(ctx, IdentityConnectionCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityConnectionCommandControllerClient) Restore(ctx context.Context, in *model.IdentityConnection, opts ...grpc.CallOption) (*model.IdentityConnection, error) {
	out := new(model.IdentityConnection)
	err := c.cc.Invoke(ctx, IdentityConnectionCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityConnectionCommandControllerServer is the server API for IdentityConnectionCommandController service.
// All implementations should embed UnimplementedIdentityConnectionCommandControllerServer
// for forward compatibility
type IdentityConnectionCommandControllerServer interface {
	// create new identity connection
	Create(context.Context, *model.IdentityConnection) (*model.IdentityConnection, error)
	// update an existing identity connection
	Update(context.Context, *model.IdentityConnection) (*model.IdentityConnection, error)
	// delete an existing identity connection
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.IdentityConnection, error)
	// restore an existing identity connection
	Restore(context.Context, *model.IdentityConnection) (*model.IdentityConnection, error)
}

// UnimplementedIdentityConnectionCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedIdentityConnectionCommandControllerServer struct {
}

func (UnimplementedIdentityConnectionCommandControllerServer) Create(context.Context, *model.IdentityConnection) (*model.IdentityConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedIdentityConnectionCommandControllerServer) Update(context.Context, *model.IdentityConnection) (*model.IdentityConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedIdentityConnectionCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.IdentityConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedIdentityConnectionCommandControllerServer) Restore(context.Context, *model.IdentityConnection) (*model.IdentityConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeIdentityConnectionCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityConnectionCommandControllerServer will
// result in compilation errors.
type UnsafeIdentityConnectionCommandControllerServer interface {
	mustEmbedUnimplementedIdentityConnectionCommandControllerServer()
}

func RegisterIdentityConnectionCommandControllerServer(s grpc.ServiceRegistrar, srv IdentityConnectionCommandControllerServer) {
	s.RegisterService(&IdentityConnectionCommandController_ServiceDesc, srv)
}

func _IdentityConnectionCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IdentityConnection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityConnectionCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityConnectionCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityConnectionCommandControllerServer).Create(ctx, req.(*model.IdentityConnection))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityConnectionCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IdentityConnection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityConnectionCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityConnectionCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityConnectionCommandControllerServer).Update(ctx, req.(*model.IdentityConnection))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityConnectionCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityConnectionCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityConnectionCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityConnectionCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityConnectionCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IdentityConnection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityConnectionCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityConnectionCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityConnectionCommandControllerServer).Restore(ctx, req.(*model.IdentityConnection))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityConnectionCommandController_ServiceDesc is the grpc.ServiceDesc for IdentityConnectionCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityConnectionCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.iam.identity.connection.service.IdentityConnectionCommandController",
	HandlerType: (*IdentityConnectionCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _IdentityConnectionCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _IdentityConnectionCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _IdentityConnectionCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _IdentityConnectionCommandController_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/iam/identity/connection/service/command.proto",
}
