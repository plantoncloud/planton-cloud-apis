// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/iam/authz/permission/service/command.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/permission/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IamPermissionCommandController_Create_FullMethodName = "/cloud.planton.apis.v1.iam.authz.permission.service.IamPermissionCommandController/create"
	IamPermissionCommandController_Update_FullMethodName = "/cloud.planton.apis.v1.iam.authz.permission.service.IamPermissionCommandController/update"
	IamPermissionCommandController_Delete_FullMethodName = "/cloud.planton.apis.v1.iam.authz.permission.service.IamPermissionCommandController/delete"
)

// IamPermissionCommandControllerClient is the client API for IamPermissionCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IamPermissionCommandControllerClient interface {
	// create iam permission
	Create(ctx context.Context, in *model.IamPermission, opts ...grpc.CallOption) (*model.IamPermission, error)
	// update iam permission
	Update(ctx context.Context, in *model.IamPermission, opts ...grpc.CallOption) (*model.IamPermission, error)
	// delete iam permission
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.IamPermission, error)
}

type iamPermissionCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIamPermissionCommandControllerClient(cc grpc.ClientConnInterface) IamPermissionCommandControllerClient {
	return &iamPermissionCommandControllerClient{cc}
}

func (c *iamPermissionCommandControllerClient) Create(ctx context.Context, in *model.IamPermission, opts ...grpc.CallOption) (*model.IamPermission, error) {
	out := new(model.IamPermission)
	err := c.cc.Invoke(ctx, IamPermissionCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamPermissionCommandControllerClient) Update(ctx context.Context, in *model.IamPermission, opts ...grpc.CallOption) (*model.IamPermission, error) {
	out := new(model.IamPermission)
	err := c.cc.Invoke(ctx, IamPermissionCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamPermissionCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.IamPermission, error) {
	out := new(model.IamPermission)
	err := c.cc.Invoke(ctx, IamPermissionCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IamPermissionCommandControllerServer is the server API for IamPermissionCommandController service.
// All implementations should embed UnimplementedIamPermissionCommandControllerServer
// for forward compatibility
type IamPermissionCommandControllerServer interface {
	// create iam permission
	Create(context.Context, *model.IamPermission) (*model.IamPermission, error)
	// update iam permission
	Update(context.Context, *model.IamPermission) (*model.IamPermission, error)
	// delete iam permission
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.IamPermission, error)
}

// UnimplementedIamPermissionCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedIamPermissionCommandControllerServer struct {
}

func (UnimplementedIamPermissionCommandControllerServer) Create(context.Context, *model.IamPermission) (*model.IamPermission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedIamPermissionCommandControllerServer) Update(context.Context, *model.IamPermission) (*model.IamPermission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedIamPermissionCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.IamPermission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeIamPermissionCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IamPermissionCommandControllerServer will
// result in compilation errors.
type UnsafeIamPermissionCommandControllerServer interface {
	mustEmbedUnimplementedIamPermissionCommandControllerServer()
}

func RegisterIamPermissionCommandControllerServer(s grpc.ServiceRegistrar, srv IamPermissionCommandControllerServer) {
	s.RegisterService(&IamPermissionCommandController_ServiceDesc, srv)
}

func _IamPermissionCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IamPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamPermissionCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamPermissionCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamPermissionCommandControllerServer).Create(ctx, req.(*model.IamPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamPermissionCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IamPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamPermissionCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamPermissionCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamPermissionCommandControllerServer).Update(ctx, req.(*model.IamPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamPermissionCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamPermissionCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamPermissionCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamPermissionCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// IamPermissionCommandController_ServiceDesc is the grpc.ServiceDesc for IamPermissionCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IamPermissionCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.iam.authz.permission.service.IamPermissionCommandController",
	HandlerType: (*IamPermissionCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _IamPermissionCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _IamPermissionCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _IamPermissionCommandController_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/iam/authz/permission/service/command.proto",
}
