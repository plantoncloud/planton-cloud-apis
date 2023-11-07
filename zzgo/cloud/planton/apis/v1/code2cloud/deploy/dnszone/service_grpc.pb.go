// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/dnszone/service.proto

package dnszone

import (
	context "context"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	company "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/company"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DnsZoneCommandController_Create_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/create"
	DnsZoneCommandController_Update_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/update"
	DnsZoneCommandController_Delete_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/delete"
	DnsZoneCommandController_Restore_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/restore"
)

// DnsZoneCommandControllerClient is the client API for DnsZoneCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DnsZoneCommandControllerClient interface {
	// add a new dns-zone to a company
	Create(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error)
	// update an existing dns-zone of a company
	// only dns-zone records can be updated as part of this operations
	Update(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error)
	// delete an dns-zone of a company
	Delete(ctx context.Context, in *DnsZoneId, opts ...grpc.CallOption) (*DnsZone, error)
	// restore a delete dns-zone
	Restore(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error)
}

type dnsZoneCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsZoneCommandControllerClient(cc grpc.ClientConnInterface) DnsZoneCommandControllerClient {
	return &dnsZoneCommandControllerClient{cc}
}

func (c *dnsZoneCommandControllerClient) Create(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) Update(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) Delete(ctx context.Context, in *DnsZoneId, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) Restore(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsZoneCommandControllerServer is the server API for DnsZoneCommandController service.
// All implementations should embed UnimplementedDnsZoneCommandControllerServer
// for forward compatibility
type DnsZoneCommandControllerServer interface {
	// add a new dns-zone to a company
	Create(context.Context, *DnsZone) (*DnsZone, error)
	// update an existing dns-zone of a company
	// only dns-zone records can be updated as part of this operations
	Update(context.Context, *DnsZone) (*DnsZone, error)
	// delete an dns-zone of a company
	Delete(context.Context, *DnsZoneId) (*DnsZone, error)
	// restore a delete dns-zone
	Restore(context.Context, *DnsZone) (*DnsZone, error)
}

// UnimplementedDnsZoneCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedDnsZoneCommandControllerServer struct {
}

func (UnimplementedDnsZoneCommandControllerServer) Create(context.Context, *DnsZone) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) Update(context.Context, *DnsZone) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) Delete(context.Context, *DnsZoneId) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) Restore(context.Context, *DnsZone) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeDnsZoneCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DnsZoneCommandControllerServer will
// result in compilation errors.
type UnsafeDnsZoneCommandControllerServer interface {
	mustEmbedUnimplementedDnsZoneCommandControllerServer()
}

func RegisterDnsZoneCommandControllerServer(s grpc.ServiceRegistrar, srv DnsZoneCommandControllerServer) {
	s.RegisterService(&DnsZoneCommandController_ServiceDesc, srv)
}

func _DnsZoneCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).Create(ctx, req.(*DnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).Update(ctx, req.(*DnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZoneId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).Delete(ctx, req.(*DnsZoneId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).Restore(ctx, req.(*DnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

// DnsZoneCommandController_ServiceDesc is the grpc.ServiceDesc for DnsZoneCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DnsZoneCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController",
	HandlerType: (*DnsZoneCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _DnsZoneCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _DnsZoneCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _DnsZoneCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _DnsZoneCommandController_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/dnszone/service.proto",
}

const (
	DnsZoneQueryController_FindByCompanyId_FullMethodName                     = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneQueryController/findByCompanyId"
	DnsZoneQueryController_GetById_FullMethodName                             = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneQueryController/getById"
	DnsZoneQueryController_GetExactOrParentDnsZoneByDomainName_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneQueryController/getExactOrParentDnsZoneByDomainName"
	DnsZoneQueryController_List_FullMethodName                                = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneQueryController/list"
	DnsZoneQueryController_IsNameserversDelegated_FullMethodName              = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneQueryController/isNameserversDelegated"
)

// DnsZoneQueryControllerClient is the client API for DnsZoneQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DnsZoneQueryControllerClient interface {
	// todo: add authorization
	// find dns-zones by company id
	// the response should only include dns-zones in a company that the authenticated user account has viewer access to.
	FindByCompanyId(ctx context.Context, in *company.CompanyId, opts ...grpc.CallOption) (*DnsZones, error)
	// todo: add authorization
	// get details of a dns-zone id
	GetById(ctx context.Context, in *DnsZoneId, opts ...grpc.CallOption) (*DnsZone, error)
	// todo: add authorization
	// get details of the exact or a parent of the provided dns-zone name
	GetExactOrParentDnsZoneByDomainName(ctx context.Context, in *DnsDomainName, opts ...grpc.CallOption) (*DnsZone, error)
	// todo: add authorization
	// list all dns-zones for the requested page. This is intended to be used on back-office portal.
	List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*DnsZoneList, error)
	// todo: add authorization
	// checks if the nameservers for the dns-zone are resolving to the nameservers of the managed zone.
	IsNameserversDelegated(ctx context.Context, in *DnsZoneId, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
}

type dnsZoneQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsZoneQueryControllerClient(cc grpc.ClientConnInterface) DnsZoneQueryControllerClient {
	return &dnsZoneQueryControllerClient{cc}
}

func (c *dnsZoneQueryControllerClient) FindByCompanyId(ctx context.Context, in *company.CompanyId, opts ...grpc.CallOption) (*DnsZones, error) {
	out := new(DnsZones)
	err := c.cc.Invoke(ctx, DnsZoneQueryController_FindByCompanyId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneQueryControllerClient) GetById(ctx context.Context, in *DnsZoneId, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneQueryControllerClient) GetExactOrParentDnsZoneByDomainName(ctx context.Context, in *DnsDomainName, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneQueryController_GetExactOrParentDnsZoneByDomainName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneQueryControllerClient) List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*DnsZoneList, error) {
	out := new(DnsZoneList)
	err := c.cc.Invoke(ctx, DnsZoneQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneQueryControllerClient) IsNameserversDelegated(ctx context.Context, in *DnsZoneId, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, DnsZoneQueryController_IsNameserversDelegated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsZoneQueryControllerServer is the server API for DnsZoneQueryController service.
// All implementations should embed UnimplementedDnsZoneQueryControllerServer
// for forward compatibility
type DnsZoneQueryControllerServer interface {
	// todo: add authorization
	// find dns-zones by company id
	// the response should only include dns-zones in a company that the authenticated user account has viewer access to.
	FindByCompanyId(context.Context, *company.CompanyId) (*DnsZones, error)
	// todo: add authorization
	// get details of a dns-zone id
	GetById(context.Context, *DnsZoneId) (*DnsZone, error)
	// todo: add authorization
	// get details of the exact or a parent of the provided dns-zone name
	GetExactOrParentDnsZoneByDomainName(context.Context, *DnsDomainName) (*DnsZone, error)
	// todo: add authorization
	// list all dns-zones for the requested page. This is intended to be used on back-office portal.
	List(context.Context, *pagination.PageInfo) (*DnsZoneList, error)
	// todo: add authorization
	// checks if the nameservers for the dns-zone are resolving to the nameservers of the managed zone.
	IsNameserversDelegated(context.Context, *DnsZoneId) (*wrapperspb.BoolValue, error)
}

// UnimplementedDnsZoneQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedDnsZoneQueryControllerServer struct {
}

func (UnimplementedDnsZoneQueryControllerServer) FindByCompanyId(context.Context, *company.CompanyId) (*DnsZones, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCompanyId not implemented")
}
func (UnimplementedDnsZoneQueryControllerServer) GetById(context.Context, *DnsZoneId) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedDnsZoneQueryControllerServer) GetExactOrParentDnsZoneByDomainName(context.Context, *DnsDomainName) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExactOrParentDnsZoneByDomainName not implemented")
}
func (UnimplementedDnsZoneQueryControllerServer) List(context.Context, *pagination.PageInfo) (*DnsZoneList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDnsZoneQueryControllerServer) IsNameserversDelegated(context.Context, *DnsZoneId) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsNameserversDelegated not implemented")
}

// UnsafeDnsZoneQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DnsZoneQueryControllerServer will
// result in compilation errors.
type UnsafeDnsZoneQueryControllerServer interface {
	mustEmbedUnimplementedDnsZoneQueryControllerServer()
}

func RegisterDnsZoneQueryControllerServer(s grpc.ServiceRegistrar, srv DnsZoneQueryControllerServer) {
	s.RegisterService(&DnsZoneQueryController_ServiceDesc, srv)
}

func _DnsZoneQueryController_FindByCompanyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(company.CompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneQueryControllerServer).FindByCompanyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneQueryController_FindByCompanyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneQueryControllerServer).FindByCompanyId(ctx, req.(*company.CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZoneId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneQueryControllerServer).GetById(ctx, req.(*DnsZoneId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneQueryController_GetExactOrParentDnsZoneByDomainName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsDomainName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneQueryControllerServer).GetExactOrParentDnsZoneByDomainName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneQueryController_GetExactOrParentDnsZoneByDomainName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneQueryControllerServer).GetExactOrParentDnsZoneByDomainName(ctx, req.(*DnsDomainName))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneQueryControllerServer).List(ctx, req.(*pagination.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneQueryController_IsNameserversDelegated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZoneId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneQueryControllerServer).IsNameserversDelegated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneQueryController_IsNameserversDelegated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneQueryControllerServer).IsNameserversDelegated(ctx, req.(*DnsZoneId))
	}
	return interceptor(ctx, in, info, handler)
}

// DnsZoneQueryController_ServiceDesc is the grpc.ServiceDesc for DnsZoneQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DnsZoneQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneQueryController",
	HandlerType: (*DnsZoneQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findByCompanyId",
			Handler:    _DnsZoneQueryController_FindByCompanyId_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _DnsZoneQueryController_GetById_Handler,
		},
		{
			MethodName: "getExactOrParentDnsZoneByDomainName",
			Handler:    _DnsZoneQueryController_GetExactOrParentDnsZoneByDomainName_Handler,
		},
		{
			MethodName: "list",
			Handler:    _DnsZoneQueryController_List_Handler,
		},
		{
			MethodName: "isNameserversDelegated",
			Handler:    _DnsZoneQueryController_IsNameserversDelegated_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/dnszone/service.proto",
}

const (
	DnsRecordCommandController_Add_FullMethodName    = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsRecordCommandController/add"
	DnsRecordCommandController_Update_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsRecordCommandController/update"
	DnsRecordCommandController_Delete_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsRecordCommandController/delete"
)

// DnsRecordCommandControllerClient is the client API for DnsRecordCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DnsRecordCommandControllerClient interface {
	// add a new dns-zone to a company
	Add(ctx context.Context, in *AddOrUpdateDnsRecordCommandInput, opts ...grpc.CallOption) (*DnsZone, error)
	// update an existing dns-zone of a company
	// only dns-zone records can be updated as part of this operations
	Update(ctx context.Context, in *AddOrUpdateDnsRecordCommandInput, opts ...grpc.CallOption) (*DnsZone, error)
	// delete an dns-zone of a company
	Delete(ctx context.Context, in *DeleteDnsRecordCommandInput, opts ...grpc.CallOption) (*DnsZone, error)
}

type dnsRecordCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsRecordCommandControllerClient(cc grpc.ClientConnInterface) DnsRecordCommandControllerClient {
	return &dnsRecordCommandControllerClient{cc}
}

func (c *dnsRecordCommandControllerClient) Add(ctx context.Context, in *AddOrUpdateDnsRecordCommandInput, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsRecordCommandController_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsRecordCommandControllerClient) Update(ctx context.Context, in *AddOrUpdateDnsRecordCommandInput, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsRecordCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsRecordCommandControllerClient) Delete(ctx context.Context, in *DeleteDnsRecordCommandInput, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsRecordCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsRecordCommandControllerServer is the server API for DnsRecordCommandController service.
// All implementations should embed UnimplementedDnsRecordCommandControllerServer
// for forward compatibility
type DnsRecordCommandControllerServer interface {
	// add a new dns-zone to a company
	Add(context.Context, *AddOrUpdateDnsRecordCommandInput) (*DnsZone, error)
	// update an existing dns-zone of a company
	// only dns-zone records can be updated as part of this operations
	Update(context.Context, *AddOrUpdateDnsRecordCommandInput) (*DnsZone, error)
	// delete an dns-zone of a company
	Delete(context.Context, *DeleteDnsRecordCommandInput) (*DnsZone, error)
}

// UnimplementedDnsRecordCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedDnsRecordCommandControllerServer struct {
}

func (UnimplementedDnsRecordCommandControllerServer) Add(context.Context, *AddOrUpdateDnsRecordCommandInput) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedDnsRecordCommandControllerServer) Update(context.Context, *AddOrUpdateDnsRecordCommandInput) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDnsRecordCommandControllerServer) Delete(context.Context, *DeleteDnsRecordCommandInput) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeDnsRecordCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DnsRecordCommandControllerServer will
// result in compilation errors.
type UnsafeDnsRecordCommandControllerServer interface {
	mustEmbedUnimplementedDnsRecordCommandControllerServer()
}

func RegisterDnsRecordCommandControllerServer(s grpc.ServiceRegistrar, srv DnsRecordCommandControllerServer) {
	s.RegisterService(&DnsRecordCommandController_ServiceDesc, srv)
}

func _DnsRecordCommandController_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateDnsRecordCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsRecordCommandControllerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsRecordCommandController_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsRecordCommandControllerServer).Add(ctx, req.(*AddOrUpdateDnsRecordCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsRecordCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateDnsRecordCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsRecordCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsRecordCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsRecordCommandControllerServer).Update(ctx, req.(*AddOrUpdateDnsRecordCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsRecordCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDnsRecordCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsRecordCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsRecordCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsRecordCommandControllerServer).Delete(ctx, req.(*DeleteDnsRecordCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// DnsRecordCommandController_ServiceDesc is the grpc.ServiceDesc for DnsRecordCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DnsRecordCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsRecordCommandController",
	HandlerType: (*DnsRecordCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _DnsRecordCommandController_Add_Handler,
		},
		{
			MethodName: "update",
			Handler:    _DnsRecordCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _DnsRecordCommandController_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/dnszone/service.proto",
}

const (
	DnsZoneStackController_Preview_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneStackController/preview"
	DnsZoneStackController_Apply_FullMethodName   = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneStackController/apply"
)

// DnsZoneStackControllerClient is the client API for DnsZoneStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DnsZoneStackControllerClient interface {
	// preview dns-zone stack
	Preview(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error)
	// apply dns-zone stack
	Apply(ctx context.Context, in *DnsZoneId, opts ...grpc.CallOption) (*DnsZone, error)
}

type dnsZoneStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsZoneStackControllerClient(cc grpc.ClientConnInterface) DnsZoneStackControllerClient {
	return &dnsZoneStackControllerClient{cc}
}

func (c *dnsZoneStackControllerClient) Preview(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneStackController_Preview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneStackControllerClient) Apply(ctx context.Context, in *DnsZoneId, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneStackController_Apply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsZoneStackControllerServer is the server API for DnsZoneStackController service.
// All implementations should embed UnimplementedDnsZoneStackControllerServer
// for forward compatibility
type DnsZoneStackControllerServer interface {
	// preview dns-zone stack
	Preview(context.Context, *DnsZone) (*DnsZone, error)
	// apply dns-zone stack
	Apply(context.Context, *DnsZoneId) (*DnsZone, error)
}

// UnimplementedDnsZoneStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedDnsZoneStackControllerServer struct {
}

func (UnimplementedDnsZoneStackControllerServer) Preview(context.Context, *DnsZone) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Preview not implemented")
}
func (UnimplementedDnsZoneStackControllerServer) Apply(context.Context, *DnsZoneId) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}

// UnsafeDnsZoneStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DnsZoneStackControllerServer will
// result in compilation errors.
type UnsafeDnsZoneStackControllerServer interface {
	mustEmbedUnimplementedDnsZoneStackControllerServer()
}

func RegisterDnsZoneStackControllerServer(s grpc.ServiceRegistrar, srv DnsZoneStackControllerServer) {
	s.RegisterService(&DnsZoneStackController_ServiceDesc, srv)
}

func _DnsZoneStackController_Preview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneStackControllerServer).Preview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneStackController_Preview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneStackControllerServer).Preview(ctx, req.(*DnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneStackController_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZoneId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneStackControllerServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneStackController_Apply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneStackControllerServer).Apply(ctx, req.(*DnsZoneId))
	}
	return interceptor(ctx, in, info, handler)
}

// DnsZoneStackController_ServiceDesc is the grpc.ServiceDesc for DnsZoneStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DnsZoneStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneStackController",
	HandlerType: (*DnsZoneStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "preview",
			Handler:    _DnsZoneStackController_Preview_Handler,
		},
		{
			MethodName: "apply",
			Handler:    _DnsZoneStackController_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/dnszone/service.proto",
}