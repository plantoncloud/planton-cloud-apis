// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/iam/identity/account/rpc/service.proto

package rpc

import (
	context "context"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MachineAccountCommandController_Create_FullMethodName  = "/cloud.planton.apis.v1.iam.identity.account.rpc.MachineAccountCommandController/create"
	MachineAccountCommandController_Update_FullMethodName  = "/cloud.planton.apis.v1.iam.identity.account.rpc.MachineAccountCommandController/update"
	MachineAccountCommandController_Delete_FullMethodName  = "/cloud.planton.apis.v1.iam.identity.account.rpc.MachineAccountCommandController/delete"
	MachineAccountCommandController_Restore_FullMethodName = "/cloud.planton.apis.v1.iam.identity.account.rpc.MachineAccountCommandController/restore"
)

// MachineAccountCommandControllerClient is the client API for MachineAccountCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MachineAccountCommandControllerClient interface {
	// create new machine account
	Create(ctx context.Context, in *IdentityAccount, opts ...grpc.CallOption) (*IdentityAccount, error)
	// update an existing machine account
	Update(ctx context.Context, in *IdentityAccount, opts ...grpc.CallOption) (*IdentityAccount, error)
	// delete an existing machine account
	Delete(ctx context.Context, in *IdentityAccountEmail, opts ...grpc.CallOption) (*IdentityAccount, error)
	// restore an existing machine account
	Restore(ctx context.Context, in *IdentityAccount, opts ...grpc.CallOption) (*IdentityAccount, error)
}

type machineAccountCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMachineAccountCommandControllerClient(cc grpc.ClientConnInterface) MachineAccountCommandControllerClient {
	return &machineAccountCommandControllerClient{cc}
}

func (c *machineAccountCommandControllerClient) Create(ctx context.Context, in *IdentityAccount, opts ...grpc.CallOption) (*IdentityAccount, error) {
	out := new(IdentityAccount)
	err := c.cc.Invoke(ctx, MachineAccountCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineAccountCommandControllerClient) Update(ctx context.Context, in *IdentityAccount, opts ...grpc.CallOption) (*IdentityAccount, error) {
	out := new(IdentityAccount)
	err := c.cc.Invoke(ctx, MachineAccountCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineAccountCommandControllerClient) Delete(ctx context.Context, in *IdentityAccountEmail, opts ...grpc.CallOption) (*IdentityAccount, error) {
	out := new(IdentityAccount)
	err := c.cc.Invoke(ctx, MachineAccountCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineAccountCommandControllerClient) Restore(ctx context.Context, in *IdentityAccount, opts ...grpc.CallOption) (*IdentityAccount, error) {
	out := new(IdentityAccount)
	err := c.cc.Invoke(ctx, MachineAccountCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MachineAccountCommandControllerServer is the server API for MachineAccountCommandController service.
// All implementations should embed UnimplementedMachineAccountCommandControllerServer
// for forward compatibility
type MachineAccountCommandControllerServer interface {
	// create new machine account
	Create(context.Context, *IdentityAccount) (*IdentityAccount, error)
	// update an existing machine account
	Update(context.Context, *IdentityAccount) (*IdentityAccount, error)
	// delete an existing machine account
	Delete(context.Context, *IdentityAccountEmail) (*IdentityAccount, error)
	// restore an existing machine account
	Restore(context.Context, *IdentityAccount) (*IdentityAccount, error)
}

// UnimplementedMachineAccountCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedMachineAccountCommandControllerServer struct {
}

func (UnimplementedMachineAccountCommandControllerServer) Create(context.Context, *IdentityAccount) (*IdentityAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMachineAccountCommandControllerServer) Update(context.Context, *IdentityAccount) (*IdentityAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMachineAccountCommandControllerServer) Delete(context.Context, *IdentityAccountEmail) (*IdentityAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMachineAccountCommandControllerServer) Restore(context.Context, *IdentityAccount) (*IdentityAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeMachineAccountCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MachineAccountCommandControllerServer will
// result in compilation errors.
type UnsafeMachineAccountCommandControllerServer interface {
	mustEmbedUnimplementedMachineAccountCommandControllerServer()
}

func RegisterMachineAccountCommandControllerServer(s grpc.ServiceRegistrar, srv MachineAccountCommandControllerServer) {
	s.RegisterService(&MachineAccountCommandController_ServiceDesc, srv)
}

func _MachineAccountCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineAccountCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineAccountCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineAccountCommandControllerServer).Create(ctx, req.(*IdentityAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineAccountCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineAccountCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineAccountCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineAccountCommandControllerServer).Update(ctx, req.(*IdentityAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineAccountCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityAccountEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineAccountCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineAccountCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineAccountCommandControllerServer).Delete(ctx, req.(*IdentityAccountEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineAccountCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineAccountCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineAccountCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineAccountCommandControllerServer).Restore(ctx, req.(*IdentityAccount))
	}
	return interceptor(ctx, in, info, handler)
}

// MachineAccountCommandController_ServiceDesc is the grpc.ServiceDesc for MachineAccountCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MachineAccountCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.iam.identity.account.rpc.MachineAccountCommandController",
	HandlerType: (*MachineAccountCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _MachineAccountCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _MachineAccountCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _MachineAccountCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _MachineAccountCommandController_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/iam/identity/account/rpc/service.proto",
}

const (
	MachineAccountQueryController_List_FullMethodName                                 = "/cloud.planton.apis.v1.iam.identity.account.rpc.MachineAccountQueryController/list"
	MachineAccountQueryController_GetById_FullMethodName                              = "/cloud.planton.apis.v1.iam.identity.account.rpc.MachineAccountQueryController/getById"
	MachineAccountQueryController_FindByCompanyId_FullMethodName                      = "/cloud.planton.apis.v1.iam.identity.account.rpc.MachineAccountQueryController/findByCompanyId"
	MachineAccountQueryController_GetByEmail_FullMethodName                           = "/cloud.planton.apis.v1.iam.identity.account.rpc.MachineAccountQueryController/getByEmail"
	MachineAccountQueryController_GetByCompanyByName_FullMethodName                   = "/cloud.planton.apis.v1.iam.identity.account.rpc.MachineAccountQueryController/getByCompanyByName"
	MachineAccountQueryController_GetClientSecretByMachineAccountEmail_FullMethodName = "/cloud.planton.apis.v1.iam.identity.account.rpc.MachineAccountQueryController/getClientSecretByMachineAccountEmail"
)

// MachineAccountQueryControllerClient is the client API for MachineAccountQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MachineAccountQueryControllerClient interface {
	// retrieve paginated list of all machine accounts on planton cloud. this is intended for use on portal.
	List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*IdentityAccountsList, error)
	// lookup machine account by identity account id.
	GetById(ctx context.Context, in *IdentityAccountId, opts ...grpc.CallOption) (*IdentityAccount, error)
	// retrieve paginated list of all machine accounts on planton cloud. this is intended for use on portal.
	FindByCompanyId(ctx context.Context, in *MachineAccountCompanyId, opts ...grpc.CallOption) (*IdentityAccounts, error)
	// lookup machine account by identity account email.
	GetByEmail(ctx context.Context, in *IdentityAccountEmail, opts ...grpc.CallOption) (*IdentityAccount, error)
	// lookup machine-account by company and name.
	GetByCompanyByName(ctx context.Context, in *GetByCompanyByNameQueryInput, opts ...grpc.CallOption) (*IdentityAccount, error)
	// lookup the client secret for the machine account.
	// since client_secret for machine account is not stored in planton database, the client secret is
	// retrieved from the idp account and is passed on to the client as response.
	GetClientSecretByMachineAccountEmail(ctx context.Context, in *IdentityAccountEmail, opts ...grpc.CallOption) (*MachineAccountClientSecret, error)
}

type machineAccountQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMachineAccountQueryControllerClient(cc grpc.ClientConnInterface) MachineAccountQueryControllerClient {
	return &machineAccountQueryControllerClient{cc}
}

func (c *machineAccountQueryControllerClient) List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*IdentityAccountsList, error) {
	out := new(IdentityAccountsList)
	err := c.cc.Invoke(ctx, MachineAccountQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineAccountQueryControllerClient) GetById(ctx context.Context, in *IdentityAccountId, opts ...grpc.CallOption) (*IdentityAccount, error) {
	out := new(IdentityAccount)
	err := c.cc.Invoke(ctx, MachineAccountQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineAccountQueryControllerClient) FindByCompanyId(ctx context.Context, in *MachineAccountCompanyId, opts ...grpc.CallOption) (*IdentityAccounts, error) {
	out := new(IdentityAccounts)
	err := c.cc.Invoke(ctx, MachineAccountQueryController_FindByCompanyId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineAccountQueryControllerClient) GetByEmail(ctx context.Context, in *IdentityAccountEmail, opts ...grpc.CallOption) (*IdentityAccount, error) {
	out := new(IdentityAccount)
	err := c.cc.Invoke(ctx, MachineAccountQueryController_GetByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineAccountQueryControllerClient) GetByCompanyByName(ctx context.Context, in *GetByCompanyByNameQueryInput, opts ...grpc.CallOption) (*IdentityAccount, error) {
	out := new(IdentityAccount)
	err := c.cc.Invoke(ctx, MachineAccountQueryController_GetByCompanyByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineAccountQueryControllerClient) GetClientSecretByMachineAccountEmail(ctx context.Context, in *IdentityAccountEmail, opts ...grpc.CallOption) (*MachineAccountClientSecret, error) {
	out := new(MachineAccountClientSecret)
	err := c.cc.Invoke(ctx, MachineAccountQueryController_GetClientSecretByMachineAccountEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MachineAccountQueryControllerServer is the server API for MachineAccountQueryController service.
// All implementations should embed UnimplementedMachineAccountQueryControllerServer
// for forward compatibility
type MachineAccountQueryControllerServer interface {
	// retrieve paginated list of all machine accounts on planton cloud. this is intended for use on portal.
	List(context.Context, *pagination.PageInfo) (*IdentityAccountsList, error)
	// lookup machine account by identity account id.
	GetById(context.Context, *IdentityAccountId) (*IdentityAccount, error)
	// retrieve paginated list of all machine accounts on planton cloud. this is intended for use on portal.
	FindByCompanyId(context.Context, *MachineAccountCompanyId) (*IdentityAccounts, error)
	// lookup machine account by identity account email.
	GetByEmail(context.Context, *IdentityAccountEmail) (*IdentityAccount, error)
	// lookup machine-account by company and name.
	GetByCompanyByName(context.Context, *GetByCompanyByNameQueryInput) (*IdentityAccount, error)
	// lookup the client secret for the machine account.
	// since client_secret for machine account is not stored in planton database, the client secret is
	// retrieved from the idp account and is passed on to the client as response.
	GetClientSecretByMachineAccountEmail(context.Context, *IdentityAccountEmail) (*MachineAccountClientSecret, error)
}

// UnimplementedMachineAccountQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedMachineAccountQueryControllerServer struct {
}

func (UnimplementedMachineAccountQueryControllerServer) List(context.Context, *pagination.PageInfo) (*IdentityAccountsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMachineAccountQueryControllerServer) GetById(context.Context, *IdentityAccountId) (*IdentityAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedMachineAccountQueryControllerServer) FindByCompanyId(context.Context, *MachineAccountCompanyId) (*IdentityAccounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCompanyId not implemented")
}
func (UnimplementedMachineAccountQueryControllerServer) GetByEmail(context.Context, *IdentityAccountEmail) (*IdentityAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByEmail not implemented")
}
func (UnimplementedMachineAccountQueryControllerServer) GetByCompanyByName(context.Context, *GetByCompanyByNameQueryInput) (*IdentityAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByCompanyByName not implemented")
}
func (UnimplementedMachineAccountQueryControllerServer) GetClientSecretByMachineAccountEmail(context.Context, *IdentityAccountEmail) (*MachineAccountClientSecret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientSecretByMachineAccountEmail not implemented")
}

// UnsafeMachineAccountQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MachineAccountQueryControllerServer will
// result in compilation errors.
type UnsafeMachineAccountQueryControllerServer interface {
	mustEmbedUnimplementedMachineAccountQueryControllerServer()
}

func RegisterMachineAccountQueryControllerServer(s grpc.ServiceRegistrar, srv MachineAccountQueryControllerServer) {
	s.RegisterService(&MachineAccountQueryController_ServiceDesc, srv)
}

func _MachineAccountQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineAccountQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineAccountQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineAccountQueryControllerServer).List(ctx, req.(*pagination.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineAccountQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityAccountId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineAccountQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineAccountQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineAccountQueryControllerServer).GetById(ctx, req.(*IdentityAccountId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineAccountQueryController_FindByCompanyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineAccountCompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineAccountQueryControllerServer).FindByCompanyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineAccountQueryController_FindByCompanyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineAccountQueryControllerServer).FindByCompanyId(ctx, req.(*MachineAccountCompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineAccountQueryController_GetByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityAccountEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineAccountQueryControllerServer).GetByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineAccountQueryController_GetByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineAccountQueryControllerServer).GetByEmail(ctx, req.(*IdentityAccountEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineAccountQueryController_GetByCompanyByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByCompanyByNameQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineAccountQueryControllerServer).GetByCompanyByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineAccountQueryController_GetByCompanyByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineAccountQueryControllerServer).GetByCompanyByName(ctx, req.(*GetByCompanyByNameQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineAccountQueryController_GetClientSecretByMachineAccountEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityAccountEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineAccountQueryControllerServer).GetClientSecretByMachineAccountEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineAccountQueryController_GetClientSecretByMachineAccountEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineAccountQueryControllerServer).GetClientSecretByMachineAccountEmail(ctx, req.(*IdentityAccountEmail))
	}
	return interceptor(ctx, in, info, handler)
}

// MachineAccountQueryController_ServiceDesc is the grpc.ServiceDesc for MachineAccountQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MachineAccountQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.iam.identity.account.rpc.MachineAccountQueryController",
	HandlerType: (*MachineAccountQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _MachineAccountQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _MachineAccountQueryController_GetById_Handler,
		},
		{
			MethodName: "findByCompanyId",
			Handler:    _MachineAccountQueryController_FindByCompanyId_Handler,
		},
		{
			MethodName: "getByEmail",
			Handler:    _MachineAccountQueryController_GetByEmail_Handler,
		},
		{
			MethodName: "getByCompanyByName",
			Handler:    _MachineAccountQueryController_GetByCompanyByName_Handler,
		},
		{
			MethodName: "getClientSecretByMachineAccountEmail",
			Handler:    _MachineAccountQueryController_GetClientSecretByMachineAccountEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/iam/identity/account/rpc/service.proto",
}

const (
	UserAccountQueryController_List_FullMethodName                      = "/cloud.planton.apis.v1.iam.identity.account.rpc.UserAccountQueryController/list"
	UserAccountQueryController_GetById_FullMethodName                   = "/cloud.planton.apis.v1.iam.identity.account.rpc.UserAccountQueryController/getById"
	UserAccountQueryController_GetByEmail_FullMethodName                = "/cloud.planton.apis.v1.iam.identity.account.rpc.UserAccountQueryController/getByEmail"
	UserAccountQueryController_IsBackofficeUser_FullMethodName          = "/cloud.planton.apis.v1.iam.identity.account.rpc.UserAccountQueryController/isBackofficeUser"
	UserAccountQueryController_ListAssociatesByCompanyId_FullMethodName = "/cloud.planton.apis.v1.iam.identity.account.rpc.UserAccountQueryController/listAssociatesByCompanyId"
)

// UserAccountQueryControllerClient is the client API for UserAccountQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAccountQueryControllerClient interface {
	// retrieve paginated list of all user accounts on planton cloud. this is intended for use on portal.
	List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*IdentityAccountsList, error)
	// todo: add authorization
	// lookup user-account by identity account id.
	GetById(ctx context.Context, in *IdentityAccountId, opts ...grpc.CallOption) (*IdentityAccount, error)
	// todo: add authorization
	// lookup user-account by identity account email.
	GetByEmail(ctx context.Context, in *IdentityAccountEmail, opts ...grpc.CallOption) (*IdentityAccount, error)
	// this is to check if a user is authorized to login to back office or not
	// like other queries this rpc does not return any data
	// instead it tries to check the authorization based on config given below
	// if rpc returns authorization error then the user is not supposed to login to backoffice
	// otherwise if rpc returns boolean response then the user is allowed to login to backoffice.
	IsBackofficeUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	// retrieve paginated list of all associate accounts of a company. this is intended for use on portal.
	ListAssociatesByCompanyId(ctx context.Context, in *ListWithIdentityCompanyId, opts ...grpc.CallOption) (*IdentityAccountsList, error)
}

type userAccountQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAccountQueryControllerClient(cc grpc.ClientConnInterface) UserAccountQueryControllerClient {
	return &userAccountQueryControllerClient{cc}
}

func (c *userAccountQueryControllerClient) List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*IdentityAccountsList, error) {
	out := new(IdentityAccountsList)
	err := c.cc.Invoke(ctx, UserAccountQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountQueryControllerClient) GetById(ctx context.Context, in *IdentityAccountId, opts ...grpc.CallOption) (*IdentityAccount, error) {
	out := new(IdentityAccount)
	err := c.cc.Invoke(ctx, UserAccountQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountQueryControllerClient) GetByEmail(ctx context.Context, in *IdentityAccountEmail, opts ...grpc.CallOption) (*IdentityAccount, error) {
	out := new(IdentityAccount)
	err := c.cc.Invoke(ctx, UserAccountQueryController_GetByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountQueryControllerClient) IsBackofficeUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, UserAccountQueryController_IsBackofficeUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountQueryControllerClient) ListAssociatesByCompanyId(ctx context.Context, in *ListWithIdentityCompanyId, opts ...grpc.CallOption) (*IdentityAccountsList, error) {
	out := new(IdentityAccountsList)
	err := c.cc.Invoke(ctx, UserAccountQueryController_ListAssociatesByCompanyId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAccountQueryControllerServer is the server API for UserAccountQueryController service.
// All implementations should embed UnimplementedUserAccountQueryControllerServer
// for forward compatibility
type UserAccountQueryControllerServer interface {
	// retrieve paginated list of all user accounts on planton cloud. this is intended for use on portal.
	List(context.Context, *pagination.PageInfo) (*IdentityAccountsList, error)
	// todo: add authorization
	// lookup user-account by identity account id.
	GetById(context.Context, *IdentityAccountId) (*IdentityAccount, error)
	// todo: add authorization
	// lookup user-account by identity account email.
	GetByEmail(context.Context, *IdentityAccountEmail) (*IdentityAccount, error)
	// this is to check if a user is authorized to login to back office or not
	// like other queries this rpc does not return any data
	// instead it tries to check the authorization based on config given below
	// if rpc returns authorization error then the user is not supposed to login to backoffice
	// otherwise if rpc returns boolean response then the user is allowed to login to backoffice.
	IsBackofficeUser(context.Context, *emptypb.Empty) (*wrapperspb.BoolValue, error)
	// retrieve paginated list of all associate accounts of a company. this is intended for use on portal.
	ListAssociatesByCompanyId(context.Context, *ListWithIdentityCompanyId) (*IdentityAccountsList, error)
}

// UnimplementedUserAccountQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedUserAccountQueryControllerServer struct {
}

func (UnimplementedUserAccountQueryControllerServer) List(context.Context, *pagination.PageInfo) (*IdentityAccountsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserAccountQueryControllerServer) GetById(context.Context, *IdentityAccountId) (*IdentityAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedUserAccountQueryControllerServer) GetByEmail(context.Context, *IdentityAccountEmail) (*IdentityAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByEmail not implemented")
}
func (UnimplementedUserAccountQueryControllerServer) IsBackofficeUser(context.Context, *emptypb.Empty) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsBackofficeUser not implemented")
}
func (UnimplementedUserAccountQueryControllerServer) ListAssociatesByCompanyId(context.Context, *ListWithIdentityCompanyId) (*IdentityAccountsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAssociatesByCompanyId not implemented")
}

// UnsafeUserAccountQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAccountQueryControllerServer will
// result in compilation errors.
type UnsafeUserAccountQueryControllerServer interface {
	mustEmbedUnimplementedUserAccountQueryControllerServer()
}

func RegisterUserAccountQueryControllerServer(s grpc.ServiceRegistrar, srv UserAccountQueryControllerServer) {
	s.RegisterService(&UserAccountQueryController_ServiceDesc, srv)
}

func _UserAccountQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAccountQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountQueryControllerServer).List(ctx, req.(*pagination.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityAccountId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAccountQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountQueryControllerServer).GetById(ctx, req.(*IdentityAccountId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountQueryController_GetByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityAccountEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountQueryControllerServer).GetByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAccountQueryController_GetByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountQueryControllerServer).GetByEmail(ctx, req.(*IdentityAccountEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountQueryController_IsBackofficeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountQueryControllerServer).IsBackofficeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAccountQueryController_IsBackofficeUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountQueryControllerServer).IsBackofficeUser(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountQueryController_ListAssociatesByCompanyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWithIdentityCompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountQueryControllerServer).ListAssociatesByCompanyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAccountQueryController_ListAssociatesByCompanyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountQueryControllerServer).ListAssociatesByCompanyId(ctx, req.(*ListWithIdentityCompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAccountQueryController_ServiceDesc is the grpc.ServiceDesc for UserAccountQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAccountQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.iam.identity.account.rpc.UserAccountQueryController",
	HandlerType: (*UserAccountQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _UserAccountQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _UserAccountQueryController_GetById_Handler,
		},
		{
			MethodName: "getByEmail",
			Handler:    _UserAccountQueryController_GetByEmail_Handler,
		},
		{
			MethodName: "isBackofficeUser",
			Handler:    _UserAccountQueryController_IsBackofficeUser_Handler,
		},
		{
			MethodName: "listAssociatesByCompanyId",
			Handler:    _UserAccountQueryController_ListAssociatesByCompanyId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/iam/identity/account/rpc/service.proto",
}
