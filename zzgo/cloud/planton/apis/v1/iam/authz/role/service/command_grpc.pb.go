// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/iam/authz/role/service/command.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/role/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IamRoleCommandController_Create_FullMethodName = "/cloud.planton.apis.v1.iam.authz.role.service.IamRoleCommandController/create"
	IamRoleCommandController_Update_FullMethodName = "/cloud.planton.apis.v1.iam.authz.role.service.IamRoleCommandController/update"
	IamRoleCommandController_Delete_FullMethodName = "/cloud.planton.apis.v1.iam.authz.role.service.IamRoleCommandController/delete"
)

// IamRoleCommandControllerClient is the client API for IamRoleCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IamRoleCommandControllerClient interface {
	// create iam role
	Create(ctx context.Context, in *model.IamRole, opts ...grpc.CallOption) (*model.IamRole, error)
	// update iam role
	Update(ctx context.Context, in *model.IamRole, opts ...grpc.CallOption) (*model.IamRole, error)
	// delete iam role
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.IamRole, error)
}

type iamRoleCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIamRoleCommandControllerClient(cc grpc.ClientConnInterface) IamRoleCommandControllerClient {
	return &iamRoleCommandControllerClient{cc}
}

func (c *iamRoleCommandControllerClient) Create(ctx context.Context, in *model.IamRole, opts ...grpc.CallOption) (*model.IamRole, error) {
	out := new(model.IamRole)
	err := c.cc.Invoke(ctx, IamRoleCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamRoleCommandControllerClient) Update(ctx context.Context, in *model.IamRole, opts ...grpc.CallOption) (*model.IamRole, error) {
	out := new(model.IamRole)
	err := c.cc.Invoke(ctx, IamRoleCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamRoleCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.IamRole, error) {
	out := new(model.IamRole)
	err := c.cc.Invoke(ctx, IamRoleCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IamRoleCommandControllerServer is the server API for IamRoleCommandController service.
// All implementations should embed UnimplementedIamRoleCommandControllerServer
// for forward compatibility
type IamRoleCommandControllerServer interface {
	// create iam role
	Create(context.Context, *model.IamRole) (*model.IamRole, error)
	// update iam role
	Update(context.Context, *model.IamRole) (*model.IamRole, error)
	// delete iam role
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.IamRole, error)
}

// UnimplementedIamRoleCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedIamRoleCommandControllerServer struct {
}

func (UnimplementedIamRoleCommandControllerServer) Create(context.Context, *model.IamRole) (*model.IamRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedIamRoleCommandControllerServer) Update(context.Context, *model.IamRole) (*model.IamRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedIamRoleCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.IamRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeIamRoleCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IamRoleCommandControllerServer will
// result in compilation errors.
type UnsafeIamRoleCommandControllerServer interface {
	mustEmbedUnimplementedIamRoleCommandControllerServer()
}

func RegisterIamRoleCommandControllerServer(s grpc.ServiceRegistrar, srv IamRoleCommandControllerServer) {
	s.RegisterService(&IamRoleCommandController_ServiceDesc, srv)
}

func _IamRoleCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IamRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamRoleCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamRoleCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamRoleCommandControllerServer).Create(ctx, req.(*model.IamRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamRoleCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IamRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamRoleCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamRoleCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamRoleCommandControllerServer).Update(ctx, req.(*model.IamRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamRoleCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamRoleCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamRoleCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamRoleCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// IamRoleCommandController_ServiceDesc is the grpc.ServiceDesc for IamRoleCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IamRoleCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.iam.authz.role.service.IamRoleCommandController",
	HandlerType: (*IamRoleCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _IamRoleCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _IamRoleCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _IamRoleCommandController_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/iam/authz/role/service/command.proto",
}
