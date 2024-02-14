// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/develop/sourcecode/server/command.proto

package server

import (
	context "context"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CodeServerCommandController_Create_FullMethodName                             = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.CodeServerCommandController/create"
	CodeServerCommandController_Update_FullMethodName                             = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.CodeServerCommandController/update"
	CodeServerCommandController_Delete_FullMethodName                             = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.CodeServerCommandController/delete"
	CodeServerCommandController_Restore_FullMethodName                            = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.CodeServerCommandController/restore"
	CodeServerCommandController_SynchronizeCodeProjects_FullMethodName            = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.CodeServerCommandController/synchronizeCodeProjects"
	CodeServerCommandController_AttachMachineAccountByCodeServerId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.CodeServerCommandController/attachMachineAccountByCodeServerId"
)

// CodeServerCommandControllerClient is the client API for CodeServerCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeServerCommandControllerClient interface {
	// create code-server
	Create(ctx context.Context, in *CodeServer, opts ...grpc.CallOption) (*CodeServer, error)
	// update an existing code-server
	Update(ctx context.Context, in *CodeServer, opts ...grpc.CallOption) (*CodeServer, error)
	// delete an existing code-server
	Delete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*CodeServer, error)
	// restore a deleted code server
	Restore(ctx context.Context, in *CodeServer, opts ...grpc.CallOption) (*CodeServer, error)
	// synchronize code projects from the server.
	SynchronizeCodeProjects(ctx context.Context, in *CodeServerId, opts ...grpc.CallOption) (*CodeServer, error)
	// attach a machine account to a code server by adding client-id and client-secret as
	// variables on organization on github or group on gitlab etc
	AttachMachineAccountByCodeServerId(ctx context.Context, in *AttachMachineAccountByCodeServerIdCommandInput, opts ...grpc.CallOption) (*CodeServer, error)
}

type codeServerCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeServerCommandControllerClient(cc grpc.ClientConnInterface) CodeServerCommandControllerClient {
	return &codeServerCommandControllerClient{cc}
}

func (c *codeServerCommandControllerClient) Create(ctx context.Context, in *CodeServer, opts ...grpc.CallOption) (*CodeServer, error) {
	out := new(CodeServer)
	err := c.cc.Invoke(ctx, CodeServerCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServerCommandControllerClient) Update(ctx context.Context, in *CodeServer, opts ...grpc.CallOption) (*CodeServer, error) {
	out := new(CodeServer)
	err := c.cc.Invoke(ctx, CodeServerCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServerCommandControllerClient) Delete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*CodeServer, error) {
	out := new(CodeServer)
	err := c.cc.Invoke(ctx, CodeServerCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServerCommandControllerClient) Restore(ctx context.Context, in *CodeServer, opts ...grpc.CallOption) (*CodeServer, error) {
	out := new(CodeServer)
	err := c.cc.Invoke(ctx, CodeServerCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServerCommandControllerClient) SynchronizeCodeProjects(ctx context.Context, in *CodeServerId, opts ...grpc.CallOption) (*CodeServer, error) {
	out := new(CodeServer)
	err := c.cc.Invoke(ctx, CodeServerCommandController_SynchronizeCodeProjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServerCommandControllerClient) AttachMachineAccountByCodeServerId(ctx context.Context, in *AttachMachineAccountByCodeServerIdCommandInput, opts ...grpc.CallOption) (*CodeServer, error) {
	out := new(CodeServer)
	err := c.cc.Invoke(ctx, CodeServerCommandController_AttachMachineAccountByCodeServerId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeServerCommandControllerServer is the server API for CodeServerCommandController service.
// All implementations should embed UnimplementedCodeServerCommandControllerServer
// for forward compatibility
type CodeServerCommandControllerServer interface {
	// create code-server
	Create(context.Context, *CodeServer) (*CodeServer, error)
	// update an existing code-server
	Update(context.Context, *CodeServer) (*CodeServer, error)
	// delete an existing code-server
	Delete(context.Context, *resource.ApiResourceDeleteCommandInput) (*CodeServer, error)
	// restore a deleted code server
	Restore(context.Context, *CodeServer) (*CodeServer, error)
	// synchronize code projects from the server.
	SynchronizeCodeProjects(context.Context, *CodeServerId) (*CodeServer, error)
	// attach a machine account to a code server by adding client-id and client-secret as
	// variables on organization on github or group on gitlab etc
	AttachMachineAccountByCodeServerId(context.Context, *AttachMachineAccountByCodeServerIdCommandInput) (*CodeServer, error)
}

// UnimplementedCodeServerCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCodeServerCommandControllerServer struct {
}

func (UnimplementedCodeServerCommandControllerServer) Create(context.Context, *CodeServer) (*CodeServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCodeServerCommandControllerServer) Update(context.Context, *CodeServer) (*CodeServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCodeServerCommandControllerServer) Delete(context.Context, *resource.ApiResourceDeleteCommandInput) (*CodeServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCodeServerCommandControllerServer) Restore(context.Context, *CodeServer) (*CodeServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedCodeServerCommandControllerServer) SynchronizeCodeProjects(context.Context, *CodeServerId) (*CodeServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SynchronizeCodeProjects not implemented")
}
func (UnimplementedCodeServerCommandControllerServer) AttachMachineAccountByCodeServerId(context.Context, *AttachMachineAccountByCodeServerIdCommandInput) (*CodeServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachMachineAccountByCodeServerId not implemented")
}

// UnsafeCodeServerCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeServerCommandControllerServer will
// result in compilation errors.
type UnsafeCodeServerCommandControllerServer interface {
	mustEmbedUnimplementedCodeServerCommandControllerServer()
}

func RegisterCodeServerCommandControllerServer(s grpc.ServiceRegistrar, srv CodeServerCommandControllerServer) {
	s.RegisterService(&CodeServerCommandController_ServiceDesc, srv)
}

func _CodeServerCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServerCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeServerCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServerCommandControllerServer).Create(ctx, req.(*CodeServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeServerCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServerCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeServerCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServerCommandControllerServer).Update(ctx, req.(*CodeServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeServerCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServerCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeServerCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServerCommandControllerServer).Delete(ctx, req.(*resource.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeServerCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServerCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeServerCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServerCommandControllerServer).Restore(ctx, req.(*CodeServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeServerCommandController_SynchronizeCodeProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServerCommandControllerServer).SynchronizeCodeProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeServerCommandController_SynchronizeCodeProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServerCommandControllerServer).SynchronizeCodeProjects(ctx, req.(*CodeServerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeServerCommandController_AttachMachineAccountByCodeServerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachMachineAccountByCodeServerIdCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServerCommandControllerServer).AttachMachineAccountByCodeServerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeServerCommandController_AttachMachineAccountByCodeServerId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServerCommandControllerServer).AttachMachineAccountByCodeServerId(ctx, req.(*AttachMachineAccountByCodeServerIdCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeServerCommandController_ServiceDesc is the grpc.ServiceDesc for CodeServerCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeServerCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.CodeServerCommandController",
	HandlerType: (*CodeServerCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _CodeServerCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _CodeServerCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _CodeServerCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _CodeServerCommandController_Restore_Handler,
		},
		{
			MethodName: "synchronizeCodeProjects",
			Handler:    _CodeServerCommandController_SynchronizeCodeProjects_Handler,
		},
		{
			MethodName: "attachMachineAccountByCodeServerId",
			Handler:    _CodeServerCommandController_AttachMachineAccountByCodeServerId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/develop/sourcecode/server/command.proto",
}