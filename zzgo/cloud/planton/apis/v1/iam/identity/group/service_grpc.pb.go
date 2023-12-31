// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/iam/identity/group/service.proto

package group

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
	IdentityGroupCommandController_Create_FullMethodName = "/cloud.planton.apis.v1.iam.identity.group.IdentityGroupCommandController/create"
	IdentityGroupCommandController_Update_FullMethodName = "/cloud.planton.apis.v1.iam.identity.group.IdentityGroupCommandController/update"
	IdentityGroupCommandController_Delete_FullMethodName = "/cloud.planton.apis.v1.iam.identity.group.IdentityGroupCommandController/delete"
)

// IdentityGroupCommandControllerClient is the client API for IdentityGroupCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityGroupCommandControllerClient interface {
	// create iam group
	Create(ctx context.Context, in *IdentityGroup, opts ...grpc.CallOption) (*IdentityGroup, error)
	// update iam group
	Update(ctx context.Context, in *IdentityGroup, opts ...grpc.CallOption) (*IdentityGroup, error)
	// delete iam group
	Delete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*IdentityGroup, error)
}

type identityGroupCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityGroupCommandControllerClient(cc grpc.ClientConnInterface) IdentityGroupCommandControllerClient {
	return &identityGroupCommandControllerClient{cc}
}

func (c *identityGroupCommandControllerClient) Create(ctx context.Context, in *IdentityGroup, opts ...grpc.CallOption) (*IdentityGroup, error) {
	out := new(IdentityGroup)
	err := c.cc.Invoke(ctx, IdentityGroupCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityGroupCommandControllerClient) Update(ctx context.Context, in *IdentityGroup, opts ...grpc.CallOption) (*IdentityGroup, error) {
	out := new(IdentityGroup)
	err := c.cc.Invoke(ctx, IdentityGroupCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityGroupCommandControllerClient) Delete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*IdentityGroup, error) {
	out := new(IdentityGroup)
	err := c.cc.Invoke(ctx, IdentityGroupCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityGroupCommandControllerServer is the server API for IdentityGroupCommandController service.
// All implementations should embed UnimplementedIdentityGroupCommandControllerServer
// for forward compatibility
type IdentityGroupCommandControllerServer interface {
	// create iam group
	Create(context.Context, *IdentityGroup) (*IdentityGroup, error)
	// update iam group
	Update(context.Context, *IdentityGroup) (*IdentityGroup, error)
	// delete iam group
	Delete(context.Context, *resource.ApiResourceDeleteCommandInput) (*IdentityGroup, error)
}

// UnimplementedIdentityGroupCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedIdentityGroupCommandControllerServer struct {
}

func (UnimplementedIdentityGroupCommandControllerServer) Create(context.Context, *IdentityGroup) (*IdentityGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedIdentityGroupCommandControllerServer) Update(context.Context, *IdentityGroup) (*IdentityGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedIdentityGroupCommandControllerServer) Delete(context.Context, *resource.ApiResourceDeleteCommandInput) (*IdentityGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeIdentityGroupCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityGroupCommandControllerServer will
// result in compilation errors.
type UnsafeIdentityGroupCommandControllerServer interface {
	mustEmbedUnimplementedIdentityGroupCommandControllerServer()
}

func RegisterIdentityGroupCommandControllerServer(s grpc.ServiceRegistrar, srv IdentityGroupCommandControllerServer) {
	s.RegisterService(&IdentityGroupCommandController_ServiceDesc, srv)
}

func _IdentityGroupCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityGroupCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityGroupCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityGroupCommandControllerServer).Create(ctx, req.(*IdentityGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityGroupCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityGroupCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityGroupCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityGroupCommandControllerServer).Update(ctx, req.(*IdentityGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityGroupCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityGroupCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityGroupCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityGroupCommandControllerServer).Delete(ctx, req.(*resource.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityGroupCommandController_ServiceDesc is the grpc.ServiceDesc for IdentityGroupCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityGroupCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.iam.identity.group.IdentityGroupCommandController",
	HandlerType: (*IdentityGroupCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _IdentityGroupCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _IdentityGroupCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _IdentityGroupCommandController_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/iam/identity/group/service.proto",
}

const (
	IdentityGroupQueryController_List_FullMethodName            = "/cloud.planton.apis.v1.iam.identity.group.IdentityGroupQueryController/list"
	IdentityGroupQueryController_GetById_FullMethodName         = "/cloud.planton.apis.v1.iam.identity.group.IdentityGroupQueryController/getById"
	IdentityGroupQueryController_ListByCompanyId_FullMethodName = "/cloud.planton.apis.v1.iam.identity.group.IdentityGroupQueryController/listByCompanyId"
)

// IdentityGroupQueryControllerClient is the client API for IdentityGroupQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityGroupQueryControllerClient interface {
	// retrieve paginated list of all iam groups on planton cloud. this is intended for use on portal.
	List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*IdentityGroupList, error)
	// lookup iam group by group id
	GetById(ctx context.Context, in *IdentityGroupId, opts ...grpc.CallOption) (*IdentityGroup, error)
	// list iam groups by company
	ListByCompanyId(ctx context.Context, in *ListWithCompanyIdReq, opts ...grpc.CallOption) (*IdentityGroupList, error)
}

type identityGroupQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityGroupQueryControllerClient(cc grpc.ClientConnInterface) IdentityGroupQueryControllerClient {
	return &identityGroupQueryControllerClient{cc}
}

func (c *identityGroupQueryControllerClient) List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*IdentityGroupList, error) {
	out := new(IdentityGroupList)
	err := c.cc.Invoke(ctx, IdentityGroupQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityGroupQueryControllerClient) GetById(ctx context.Context, in *IdentityGroupId, opts ...grpc.CallOption) (*IdentityGroup, error) {
	out := new(IdentityGroup)
	err := c.cc.Invoke(ctx, IdentityGroupQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityGroupQueryControllerClient) ListByCompanyId(ctx context.Context, in *ListWithCompanyIdReq, opts ...grpc.CallOption) (*IdentityGroupList, error) {
	out := new(IdentityGroupList)
	err := c.cc.Invoke(ctx, IdentityGroupQueryController_ListByCompanyId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityGroupQueryControllerServer is the server API for IdentityGroupQueryController service.
// All implementations should embed UnimplementedIdentityGroupQueryControllerServer
// for forward compatibility
type IdentityGroupQueryControllerServer interface {
	// retrieve paginated list of all iam groups on planton cloud. this is intended for use on portal.
	List(context.Context, *pagination.PageInfo) (*IdentityGroupList, error)
	// lookup iam group by group id
	GetById(context.Context, *IdentityGroupId) (*IdentityGroup, error)
	// list iam groups by company
	ListByCompanyId(context.Context, *ListWithCompanyIdReq) (*IdentityGroupList, error)
}

// UnimplementedIdentityGroupQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedIdentityGroupQueryControllerServer struct {
}

func (UnimplementedIdentityGroupQueryControllerServer) List(context.Context, *pagination.PageInfo) (*IdentityGroupList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedIdentityGroupQueryControllerServer) GetById(context.Context, *IdentityGroupId) (*IdentityGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedIdentityGroupQueryControllerServer) ListByCompanyId(context.Context, *ListWithCompanyIdReq) (*IdentityGroupList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByCompanyId not implemented")
}

// UnsafeIdentityGroupQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityGroupQueryControllerServer will
// result in compilation errors.
type UnsafeIdentityGroupQueryControllerServer interface {
	mustEmbedUnimplementedIdentityGroupQueryControllerServer()
}

func RegisterIdentityGroupQueryControllerServer(s grpc.ServiceRegistrar, srv IdentityGroupQueryControllerServer) {
	s.RegisterService(&IdentityGroupQueryController_ServiceDesc, srv)
}

func _IdentityGroupQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityGroupQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityGroupQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityGroupQueryControllerServer).List(ctx, req.(*pagination.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityGroupQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityGroupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityGroupQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityGroupQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityGroupQueryControllerServer).GetById(ctx, req.(*IdentityGroupId))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityGroupQueryController_ListByCompanyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWithCompanyIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityGroupQueryControllerServer).ListByCompanyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityGroupQueryController_ListByCompanyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityGroupQueryControllerServer).ListByCompanyId(ctx, req.(*ListWithCompanyIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityGroupQueryController_ServiceDesc is the grpc.ServiceDesc for IdentityGroupQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityGroupQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.iam.identity.group.IdentityGroupQueryController",
	HandlerType: (*IdentityGroupQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _IdentityGroupQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _IdentityGroupQueryController_GetById_Handler,
		},
		{
			MethodName: "listByCompanyId",
			Handler:    _IdentityGroupQueryController_ListByCompanyId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/iam/identity/group/service.proto",
}
