// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/develop/sourcecode/project/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/develop/sourcecode/project/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/product/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CodeProjectCommandController_Add_FullMethodName                                 = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController/add"
	CodeProjectCommandController_Create_FullMethodName                              = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController/create"
	CodeProjectCommandController_Update_FullMethodName                              = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController/update"
	CodeProjectCommandController_Delete_FullMethodName                              = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController/delete"
	CodeProjectCommandController_Restore_FullMethodName                             = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController/restore"
	CodeProjectCommandController_SynchronizeByProductId_FullMethodName              = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController/synchronizeByProductId"
	CodeProjectCommandController_SynchronizeByCodeProjectId_FullMethodName          = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController/synchronizeByCodeProjectId"
	CodeProjectCommandController_AttachMachineAccountByCodeProjectId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController/attachMachineAccountByCodeProjectId"
)

// CodeProjectCommandControllerClient is the client API for CodeProjectCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeProjectCommandControllerClient interface {
	// add code-project that already exists on the code-server.
	Add(ctx context.Context, in *model.CodeProject, opts ...grpc.CallOption) (*model.CodeProject, error)
	// create a code-project on selected code-server.
	// new code projects created from planton cloud, can also choose an available code project template.
	Create(ctx context.Context, in *model.CodeProject, opts ...grpc.CallOption) (*model.CodeProject, error)
	// update an existing code-project
	Update(ctx context.Context, in *model.CodeProject, opts ...grpc.CallOption) (*model.CodeProject, error)
	// delete an existing code project.
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.CodeProject, error)
	// restore a deleted code project of a product.
	Restore(ctx context.Context, in *model.CodeProject, opts ...grpc.CallOption) (*model.CodeProject, error)
	// synchronize code projects of a product.
	// this operation will run synchronization process on all code-servers of the product.
	SynchronizeByProductId(ctx context.Context, in *model2.ProductId, opts ...grpc.CallOption) (*model2.Product, error)
	// synchronize code project with its counterpart on the code-server
	SynchronizeByCodeProjectId(ctx context.Context, in *model.CodeProjectId, opts ...grpc.CallOption) (*model.CodeProject, error)
	// attach a machine account to a code project on the code-server by adding client-id and client-secret as
	// variables on the repository or project on github, gitlab etc
	AttachMachineAccountByCodeProjectId(ctx context.Context, in *model.AttachMachineAccountByCodeProjectIdCommandInput, opts ...grpc.CallOption) (*model.CodeProject, error)
}

type codeProjectCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeProjectCommandControllerClient(cc grpc.ClientConnInterface) CodeProjectCommandControllerClient {
	return &codeProjectCommandControllerClient{cc}
}

func (c *codeProjectCommandControllerClient) Add(ctx context.Context, in *model.CodeProject, opts ...grpc.CallOption) (*model.CodeProject, error) {
	out := new(model.CodeProject)
	err := c.cc.Invoke(ctx, CodeProjectCommandController_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectCommandControllerClient) Create(ctx context.Context, in *model.CodeProject, opts ...grpc.CallOption) (*model.CodeProject, error) {
	out := new(model.CodeProject)
	err := c.cc.Invoke(ctx, CodeProjectCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectCommandControllerClient) Update(ctx context.Context, in *model.CodeProject, opts ...grpc.CallOption) (*model.CodeProject, error) {
	out := new(model.CodeProject)
	err := c.cc.Invoke(ctx, CodeProjectCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.CodeProject, error) {
	out := new(model.CodeProject)
	err := c.cc.Invoke(ctx, CodeProjectCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectCommandControllerClient) Restore(ctx context.Context, in *model.CodeProject, opts ...grpc.CallOption) (*model.CodeProject, error) {
	out := new(model.CodeProject)
	err := c.cc.Invoke(ctx, CodeProjectCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectCommandControllerClient) SynchronizeByProductId(ctx context.Context, in *model2.ProductId, opts ...grpc.CallOption) (*model2.Product, error) {
	out := new(model2.Product)
	err := c.cc.Invoke(ctx, CodeProjectCommandController_SynchronizeByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectCommandControllerClient) SynchronizeByCodeProjectId(ctx context.Context, in *model.CodeProjectId, opts ...grpc.CallOption) (*model.CodeProject, error) {
	out := new(model.CodeProject)
	err := c.cc.Invoke(ctx, CodeProjectCommandController_SynchronizeByCodeProjectId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectCommandControllerClient) AttachMachineAccountByCodeProjectId(ctx context.Context, in *model.AttachMachineAccountByCodeProjectIdCommandInput, opts ...grpc.CallOption) (*model.CodeProject, error) {
	out := new(model.CodeProject)
	err := c.cc.Invoke(ctx, CodeProjectCommandController_AttachMachineAccountByCodeProjectId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeProjectCommandControllerServer is the server API for CodeProjectCommandController service.
// All implementations should embed UnimplementedCodeProjectCommandControllerServer
// for forward compatibility
type CodeProjectCommandControllerServer interface {
	// add code-project that already exists on the code-server.
	Add(context.Context, *model.CodeProject) (*model.CodeProject, error)
	// create a code-project on selected code-server.
	// new code projects created from planton cloud, can also choose an available code project template.
	Create(context.Context, *model.CodeProject) (*model.CodeProject, error)
	// update an existing code-project
	Update(context.Context, *model.CodeProject) (*model.CodeProject, error)
	// delete an existing code project.
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.CodeProject, error)
	// restore a deleted code project of a product.
	Restore(context.Context, *model.CodeProject) (*model.CodeProject, error)
	// synchronize code projects of a product.
	// this operation will run synchronization process on all code-servers of the product.
	SynchronizeByProductId(context.Context, *model2.ProductId) (*model2.Product, error)
	// synchronize code project with its counterpart on the code-server
	SynchronizeByCodeProjectId(context.Context, *model.CodeProjectId) (*model.CodeProject, error)
	// attach a machine account to a code project on the code-server by adding client-id and client-secret as
	// variables on the repository or project on github, gitlab etc
	AttachMachineAccountByCodeProjectId(context.Context, *model.AttachMachineAccountByCodeProjectIdCommandInput) (*model.CodeProject, error)
}

// UnimplementedCodeProjectCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCodeProjectCommandControllerServer struct {
}

func (UnimplementedCodeProjectCommandControllerServer) Add(context.Context, *model.CodeProject) (*model.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCodeProjectCommandControllerServer) Create(context.Context, *model.CodeProject) (*model.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCodeProjectCommandControllerServer) Update(context.Context, *model.CodeProject) (*model.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCodeProjectCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCodeProjectCommandControllerServer) Restore(context.Context, *model.CodeProject) (*model.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedCodeProjectCommandControllerServer) SynchronizeByProductId(context.Context, *model2.ProductId) (*model2.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SynchronizeByProductId not implemented")
}
func (UnimplementedCodeProjectCommandControllerServer) SynchronizeByCodeProjectId(context.Context, *model.CodeProjectId) (*model.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SynchronizeByCodeProjectId not implemented")
}
func (UnimplementedCodeProjectCommandControllerServer) AttachMachineAccountByCodeProjectId(context.Context, *model.AttachMachineAccountByCodeProjectIdCommandInput) (*model.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachMachineAccountByCodeProjectId not implemented")
}

// UnsafeCodeProjectCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeProjectCommandControllerServer will
// result in compilation errors.
type UnsafeCodeProjectCommandControllerServer interface {
	mustEmbedUnimplementedCodeProjectCommandControllerServer()
}

func RegisterCodeProjectCommandControllerServer(s grpc.ServiceRegistrar, srv CodeProjectCommandControllerServer) {
	s.RegisterService(&CodeProjectCommandController_ServiceDesc, srv)
}

func _CodeProjectCommandController_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CodeProject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectCommandControllerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectCommandController_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectCommandControllerServer).Add(ctx, req.(*model.CodeProject))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CodeProject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectCommandControllerServer).Create(ctx, req.(*model.CodeProject))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CodeProject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectCommandControllerServer).Update(ctx, req.(*model.CodeProject))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CodeProject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectCommandControllerServer).Restore(ctx, req.(*model.CodeProject))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectCommandController_SynchronizeByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectCommandControllerServer).SynchronizeByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectCommandController_SynchronizeByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectCommandControllerServer).SynchronizeByProductId(ctx, req.(*model2.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectCommandController_SynchronizeByCodeProjectId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CodeProjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectCommandControllerServer).SynchronizeByCodeProjectId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectCommandController_SynchronizeByCodeProjectId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectCommandControllerServer).SynchronizeByCodeProjectId(ctx, req.(*model.CodeProjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectCommandController_AttachMachineAccountByCodeProjectId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AttachMachineAccountByCodeProjectIdCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectCommandControllerServer).AttachMachineAccountByCodeProjectId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectCommandController_AttachMachineAccountByCodeProjectId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectCommandControllerServer).AttachMachineAccountByCodeProjectId(ctx, req.(*model.AttachMachineAccountByCodeProjectIdCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeProjectCommandController_ServiceDesc is the grpc.ServiceDesc for CodeProjectCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeProjectCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController",
	HandlerType: (*CodeProjectCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _CodeProjectCommandController_Add_Handler,
		},
		{
			MethodName: "create",
			Handler:    _CodeProjectCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _CodeProjectCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _CodeProjectCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _CodeProjectCommandController_Restore_Handler,
		},
		{
			MethodName: "synchronizeByProductId",
			Handler:    _CodeProjectCommandController_SynchronizeByProductId_Handler,
		},
		{
			MethodName: "synchronizeByCodeProjectId",
			Handler:    _CodeProjectCommandController_SynchronizeByCodeProjectId_Handler,
		},
		{
			MethodName: "attachMachineAccountByCodeProjectId",
			Handler:    _CodeProjectCommandController_AttachMachineAccountByCodeProjectId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/develop/sourcecode/project/service/command.proto",
}
