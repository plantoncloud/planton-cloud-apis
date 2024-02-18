// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/dnszone/service/query.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/dnszone/model"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/rpc"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/resourcemanager/v1/company/model"
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
	DnsZoneQueryController_FindByCompanyId_FullMethodName                     = "/cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneQueryController/findByCompanyId"
	DnsZoneQueryController_GetById_FullMethodName                             = "/cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneQueryController/getById"
	DnsZoneQueryController_GetExactOrParentDnsZoneByDomainName_FullMethodName = "/cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneQueryController/getExactOrParentDnsZoneByDomainName"
	DnsZoneQueryController_List_FullMethodName                                = "/cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneQueryController/list"
	DnsZoneQueryController_IsNameserversDelegated_FullMethodName              = "/cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneQueryController/isNameserversDelegated"
)

// DnsZoneQueryControllerClient is the client API for DnsZoneQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DnsZoneQueryControllerClient interface {
	// todo: add authorization
	// find dns-zones by company id
	// the response should only include dns-zones in a company that the authenticated user account has viewer access to.
	FindByCompanyId(ctx context.Context, in *model.CompanyId, opts ...grpc.CallOption) (*model1.DnsZones, error)
	// todo: add authorization
	// get details of a dns-zone id
	GetById(ctx context.Context, in *model1.DnsZoneId, opts ...grpc.CallOption) (*model1.DnsZone, error)
	// todo: add authorization
	// get details of the exact or a parent of the provided dns-zone name
	GetExactOrParentDnsZoneByDomainName(ctx context.Context, in *model1.DnsDomainName, opts ...grpc.CallOption) (*model1.DnsZone, error)
	// todo: add authorization
	// list all dns-zones for the requested page. This is intended to be used on back-office portal.
	List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model1.DnsZoneList, error)
	// todo: add authorization
	// checks if the nameservers for the dns-zone are resolving to the nameservers of the managed zone.
	IsNameserversDelegated(ctx context.Context, in *model1.DnsZoneId, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
}

type dnsZoneQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsZoneQueryControllerClient(cc grpc.ClientConnInterface) DnsZoneQueryControllerClient {
	return &dnsZoneQueryControllerClient{cc}
}

func (c *dnsZoneQueryControllerClient) FindByCompanyId(ctx context.Context, in *model.CompanyId, opts ...grpc.CallOption) (*model1.DnsZones, error) {
	out := new(model1.DnsZones)
	err := c.cc.Invoke(ctx, DnsZoneQueryController_FindByCompanyId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneQueryControllerClient) GetById(ctx context.Context, in *model1.DnsZoneId, opts ...grpc.CallOption) (*model1.DnsZone, error) {
	out := new(model1.DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneQueryControllerClient) GetExactOrParentDnsZoneByDomainName(ctx context.Context, in *model1.DnsDomainName, opts ...grpc.CallOption) (*model1.DnsZone, error) {
	out := new(model1.DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneQueryController_GetExactOrParentDnsZoneByDomainName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneQueryControllerClient) List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model1.DnsZoneList, error) {
	out := new(model1.DnsZoneList)
	err := c.cc.Invoke(ctx, DnsZoneQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneQueryControllerClient) IsNameserversDelegated(ctx context.Context, in *model1.DnsZoneId, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
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
	FindByCompanyId(context.Context, *model.CompanyId) (*model1.DnsZones, error)
	// todo: add authorization
	// get details of a dns-zone id
	GetById(context.Context, *model1.DnsZoneId) (*model1.DnsZone, error)
	// todo: add authorization
	// get details of the exact or a parent of the provided dns-zone name
	GetExactOrParentDnsZoneByDomainName(context.Context, *model1.DnsDomainName) (*model1.DnsZone, error)
	// todo: add authorization
	// list all dns-zones for the requested page. This is intended to be used on back-office portal.
	List(context.Context, *rpc.PageInfo) (*model1.DnsZoneList, error)
	// todo: add authorization
	// checks if the nameservers for the dns-zone are resolving to the nameservers of the managed zone.
	IsNameserversDelegated(context.Context, *model1.DnsZoneId) (*wrapperspb.BoolValue, error)
}

// UnimplementedDnsZoneQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedDnsZoneQueryControllerServer struct {
}

func (UnimplementedDnsZoneQueryControllerServer) FindByCompanyId(context.Context, *model.CompanyId) (*model1.DnsZones, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCompanyId not implemented")
}
func (UnimplementedDnsZoneQueryControllerServer) GetById(context.Context, *model1.DnsZoneId) (*model1.DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedDnsZoneQueryControllerServer) GetExactOrParentDnsZoneByDomainName(context.Context, *model1.DnsDomainName) (*model1.DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExactOrParentDnsZoneByDomainName not implemented")
}
func (UnimplementedDnsZoneQueryControllerServer) List(context.Context, *rpc.PageInfo) (*model1.DnsZoneList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDnsZoneQueryControllerServer) IsNameserversDelegated(context.Context, *model1.DnsZoneId) (*wrapperspb.BoolValue, error) {
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
	in := new(model.CompanyId)
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
		return srv.(DnsZoneQueryControllerServer).FindByCompanyId(ctx, req.(*model.CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.DnsZoneId)
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
		return srv.(DnsZoneQueryControllerServer).GetById(ctx, req.(*model1.DnsZoneId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneQueryController_GetExactOrParentDnsZoneByDomainName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.DnsDomainName)
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
		return srv.(DnsZoneQueryControllerServer).GetExactOrParentDnsZoneByDomainName(ctx, req.(*model1.DnsDomainName))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.PageInfo)
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
		return srv.(DnsZoneQueryControllerServer).List(ctx, req.(*rpc.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneQueryController_IsNameserversDelegated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.DnsZoneId)
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
		return srv.(DnsZoneQueryControllerServer).IsNameserversDelegated(ctx, req.(*model1.DnsZoneId))
	}
	return interceptor(ctx, in, info, handler)
}

// DnsZoneQueryController_ServiceDesc is the grpc.ServiceDesc for DnsZoneQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DnsZoneQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneQueryController",
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
	Metadata: "cloud/planton/apis/code2cloud/v1/dnszone/service/query.proto",
}