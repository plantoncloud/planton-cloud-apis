// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/iam/authz/permission/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/permission/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IamPermissionQueryController_List_FullMethodName                = "/cloud.planton.apis.v1.iam.authz.permission.service.IamPermissionQueryController/list"
	IamPermissionQueryController_GetById_FullMethodName             = "/cloud.planton.apis.v1.iam.authz.permission.service.IamPermissionQueryController/getById"
	IamPermissionQueryController_GetByPermissionCode_FullMethodName = "/cloud.planton.apis.v1.iam.authz.permission.service.IamPermissionQueryController/getByPermissionCode"
)

// IamPermissionQueryControllerClient is the client API for IamPermissionQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IamPermissionQueryControllerClient interface {
	// retrieve paginated list of all iam permissions on planton cloud. this is intended for use on portal.
	List(ctx context.Context, in *model.PageInfo, opts ...grpc.CallOption) (*model1.IamPermissionList, error)
	// lookup iam permission by permission id
	GetById(ctx context.Context, in *model1.IamPermissionId, opts ...grpc.CallOption) (*model1.IamPermission, error)
	// lookup iam permission by permission code
	GetByPermissionCode(ctx context.Context, in *model1.IamPermissionCode, opts ...grpc.CallOption) (*model1.IamPermission, error)
}

type iamPermissionQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIamPermissionQueryControllerClient(cc grpc.ClientConnInterface) IamPermissionQueryControllerClient {
	return &iamPermissionQueryControllerClient{cc}
}

func (c *iamPermissionQueryControllerClient) List(ctx context.Context, in *model.PageInfo, opts ...grpc.CallOption) (*model1.IamPermissionList, error) {
	out := new(model1.IamPermissionList)
	err := c.cc.Invoke(ctx, IamPermissionQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamPermissionQueryControllerClient) GetById(ctx context.Context, in *model1.IamPermissionId, opts ...grpc.CallOption) (*model1.IamPermission, error) {
	out := new(model1.IamPermission)
	err := c.cc.Invoke(ctx, IamPermissionQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamPermissionQueryControllerClient) GetByPermissionCode(ctx context.Context, in *model1.IamPermissionCode, opts ...grpc.CallOption) (*model1.IamPermission, error) {
	out := new(model1.IamPermission)
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
	List(context.Context, *model.PageInfo) (*model1.IamPermissionList, error)
	// lookup iam permission by permission id
	GetById(context.Context, *model1.IamPermissionId) (*model1.IamPermission, error)
	// lookup iam permission by permission code
	GetByPermissionCode(context.Context, *model1.IamPermissionCode) (*model1.IamPermission, error)
}

// UnimplementedIamPermissionQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedIamPermissionQueryControllerServer struct {
}

func (UnimplementedIamPermissionQueryControllerServer) List(context.Context, *model.PageInfo) (*model1.IamPermissionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedIamPermissionQueryControllerServer) GetById(context.Context, *model1.IamPermissionId) (*model1.IamPermission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedIamPermissionQueryControllerServer) GetByPermissionCode(context.Context, *model1.IamPermissionCode) (*model1.IamPermission, error) {
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
	in := new(model.PageInfo)
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
		return srv.(IamPermissionQueryControllerServer).List(ctx, req.(*model.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamPermissionQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.IamPermissionId)
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
		return srv.(IamPermissionQueryControllerServer).GetById(ctx, req.(*model1.IamPermissionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamPermissionQueryController_GetByPermissionCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.IamPermissionCode)
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
		return srv.(IamPermissionQueryControllerServer).GetByPermissionCode(ctx, req.(*model1.IamPermissionCode))
	}
	return interceptor(ctx, in, info, handler)
}

// IamPermissionQueryController_ServiceDesc is the grpc.ServiceDesc for IamPermissionQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IamPermissionQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.iam.authz.permission.service.IamPermissionQueryController",
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
	Metadata: "cloud/planton/apis/v1/iam/authz/permission/service/query.proto",
}
