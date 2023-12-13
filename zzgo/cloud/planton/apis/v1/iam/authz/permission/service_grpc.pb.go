// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/iam/authz/permission/service.proto

package permission

import (
	context "context"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IamPermissionCommandController_Create_FullMethodName = "/cloud.planton.apis.v1.iam.authz.permission.IamPermissionCommandController/create"
	IamPermissionCommandController_Update_FullMethodName = "/cloud.planton.apis.v1.iam.authz.permission.IamPermissionCommandController/update"
	IamPermissionCommandController_Delete_FullMethodName = "/cloud.planton.apis.v1.iam.authz.permission.IamPermissionCommandController/delete"
)

// IamPermissionCommandControllerClient is the client API for IamPermissionCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IamPermissionCommandControllerClient interface {
	// create iam permission
	Create(ctx context.Context, in *IamPermission, opts ...grpc.CallOption) (*IamPermission, error)
	// update iam permission
	Update(ctx context.Context, in *IamPermission, opts ...grpc.CallOption) (*IamPermission, error)
	// delete iam permission
	Delete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*IamPermission, error)
}

type iamPermissionCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIamPermissionCommandControllerClient(cc grpc.ClientConnInterface) IamPermissionCommandControllerClient {
	return &iamPermissionCommandControllerClient{cc}
}

func (c *iamPermissionCommandControllerClient) Create(ctx context.Context, in *IamPermission, opts ...grpc.CallOption) (*IamPermission, error) {
	out := new(IamPermission)
	err := c.cc.Invoke(ctx, IamPermissionCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamPermissionCommandControllerClient) Update(ctx context.Context, in *IamPermission, opts ...grpc.CallOption) (*IamPermission, error) {
	out := new(IamPermission)
	err := c.cc.Invoke(ctx, IamPermissionCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamPermissionCommandControllerClient) Delete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*IamPermission, error) {
	out := new(IamPermission)
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
	Create(context.Context, *IamPermission) (*IamPermission, error)
	// update iam permission
	Update(context.Context, *IamPermission) (*IamPermission, error)
	// delete iam permission
	Delete(context.Context, *resource.ApiResourceDeleteCommandInput) (*IamPermission, error)
}

// UnimplementedIamPermissionCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedIamPermissionCommandControllerServer struct {
}

func (UnimplementedIamPermissionCommandControllerServer) Create(context.Context, *IamPermission) (*IamPermission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedIamPermissionCommandControllerServer) Update(context.Context, *IamPermission) (*IamPermission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedIamPermissionCommandControllerServer) Delete(context.Context, *resource.ApiResourceDeleteCommandInput) (*IamPermission, error) {
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
	in := new(IamPermission)
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
		return srv.(IamPermissionCommandControllerServer).Create(ctx, req.(*IamPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamPermissionCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IamPermission)
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
		return srv.(IamPermissionCommandControllerServer).Update(ctx, req.(*IamPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamPermissionCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceDeleteCommandInput)
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
		return srv.(IamPermissionCommandControllerServer).Delete(ctx, req.(*resource.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// IamPermissionCommandController_ServiceDesc is the grpc.ServiceDesc for IamPermissionCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IamPermissionCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.iam.authz.permission.IamPermissionCommandController",
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
	Metadata: "cloud/planton/apis/v1/iam/authz/permission/service.proto",
}

const (
	IamPermissionQueryController_List_FullMethodName                = "/cloud.planton.apis.v1.iam.authz.permission.IamPermissionQueryController/list"
	IamPermissionQueryController_GetById_FullMethodName             = "/cloud.planton.apis.v1.iam.authz.permission.IamPermissionQueryController/getById"
	IamPermissionQueryController_GetByPermissionCode_FullMethodName = "/cloud.planton.apis.v1.iam.authz.permission.IamPermissionQueryController/getByPermissionCode"
)

// IamPermissionQueryControllerClient is the client API for IamPermissionQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IamPermissionQueryControllerClient interface {
	// retrieve paginated list of all iam permissions on planton cloud. this is intended for use on portal.
	List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*IamPermissionList, error)
	// lookup iam permission by permission id
	GetById(ctx context.Context, in *IamPermissionId, opts ...grpc.CallOption) (*IamPermission, error)
	// lookup iam permission by permission code
	GetByPermissionCode(ctx context.Context, in *IamPermissionCode, opts ...grpc.CallOption) (*IamPermission, error)
}

type iamPermissionQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIamPermissionQueryControllerClient(cc grpc.ClientConnInterface) IamPermissionQueryControllerClient {
	return &iamPermissionQueryControllerClient{cc}
}

func (c *iamPermissionQueryControllerClient) List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*IamPermissionList, error) {
	out := new(IamPermissionList)
	err := c.cc.Invoke(ctx, IamPermissionQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamPermissionQueryControllerClient) GetById(ctx context.Context, in *IamPermissionId, opts ...grpc.CallOption) (*IamPermission, error) {
	out := new(IamPermission)
	err := c.cc.Invoke(ctx, IamPermissionQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamPermissionQueryControllerClient) GetByPermissionCode(ctx context.Context, in *IamPermissionCode, opts ...grpc.CallOption) (*IamPermission, error) {
	out := new(IamPermission)
	err := c.cc.Invoke(ctx, IamPermissionQueryController_GetByPermissionCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IamPermissionQueryControllerServer is the server API for IamPermissionQueryController service.
// All implementations should embed UnimplementedIamPermissionQueryControllerServer
// for forward compatibility
type IamPermissionQueryControllerServer interface {
	// retrieve paginated list of all iam permissions on planton cloud. this is intended for use on portal.
	List(context.Context, *pagination.PageInfo) (*IamPermissionList, error)
	// lookup iam permission by permission id
	GetById(context.Context, *IamPermissionId) (*IamPermission, error)
	// lookup iam permission by permission code
	GetByPermissionCode(context.Context, *IamPermissionCode) (*IamPermission, error)
}

// UnimplementedIamPermissionQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedIamPermissionQueryControllerServer struct {
}

func (UnimplementedIamPermissionQueryControllerServer) List(context.Context, *pagination.PageInfo) (*IamPermissionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedIamPermissionQueryControllerServer) GetById(context.Context, *IamPermissionId) (*IamPermission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedIamPermissionQueryControllerServer) GetByPermissionCode(context.Context, *IamPermissionCode) (*IamPermission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByPermissionCode not implemented")
}

// UnsafeIamPermissionQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IamPermissionQueryControllerServer will
// result in compilation errors.
type UnsafeIamPermissionQueryControllerServer interface {
	mustEmbedUnimplementedIamPermissionQueryControllerServer()
}

func RegisterIamPermissionQueryControllerServer(s grpc.ServiceRegistrar, srv IamPermissionQueryControllerServer) {
	s.RegisterService(&IamPermissionQueryController_ServiceDesc, srv)
}

func _IamPermissionQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamPermissionQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamPermissionQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamPermissionQueryControllerServer).List(ctx, req.(*pagination.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamPermissionQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IamPermissionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamPermissionQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamPermissionQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamPermissionQueryControllerServer).GetById(ctx, req.(*IamPermissionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamPermissionQueryController_GetByPermissionCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IamPermissionCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamPermissionQueryControllerServer).GetByPermissionCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IamPermissionQueryController_GetByPermissionCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamPermissionQueryControllerServer).GetByPermissionCode(ctx, req.(*IamPermissionCode))
	}
	return interceptor(ctx, in, info, handler)
}

// IamPermissionQueryController_ServiceDesc is the grpc.ServiceDesc for IamPermissionQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IamPermissionQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.iam.authz.permission.IamPermissionQueryController",
	HandlerType: (*IamPermissionQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _IamPermissionQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _IamPermissionQueryController_GetById_Handler,
		},
		{
			MethodName: "getByPermissionCode",
			Handler:    _IamPermissionQueryController_GetByPermissionCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/iam/authz/permission/service.proto",
}
