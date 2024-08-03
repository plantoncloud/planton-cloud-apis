// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/resourcemanager/v1/organization/service/query.proto

package service

import (
	context "context"
	protobuf "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/protobuf"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/resourcemanager/v1/organization/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OrganizationQueryController_Get_FullMethodName               = "/cloud.planton.apis.resourcemanager.v1.organization.service.OrganizationQueryController/get"
	OrganizationQueryController_FindOrganizations_FullMethodName = "/cloud.planton.apis.resourcemanager.v1.organization.service.OrganizationQueryController/findOrganizations"
)

// OrganizationQueryControllerClient is the client API for OrganizationQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationQueryControllerClient interface {
	// get a organization using organization id
	Get(ctx context.Context, in *model.OrgId, opts ...grpc.CallOption) (*model.Organization, error)
	FindOrganizations(ctx context.Context, in *protobuf.CustomEmpty, opts ...grpc.CallOption) (*model.Organizations, error)
}

type organizationQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationQueryControllerClient(cc grpc.ClientConnInterface) OrganizationQueryControllerClient {
	return &organizationQueryControllerClient{cc}
}

func (c *organizationQueryControllerClient) Get(ctx context.Context, in *model.OrgId, opts ...grpc.CallOption) (*model.Organization, error) {
	out := new(model.Organization)
	err := c.cc.Invoke(ctx, OrganizationQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationQueryControllerClient) FindOrganizations(ctx context.Context, in *protobuf.CustomEmpty, opts ...grpc.CallOption) (*model.Organizations, error) {
	out := new(model.Organizations)
	err := c.cc.Invoke(ctx, OrganizationQueryController_FindOrganizations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationQueryControllerServer is the server API for OrganizationQueryController service.
// All implementations should embed UnimplementedOrganizationQueryControllerServer
// for forward compatibility
type OrganizationQueryControllerServer interface {
	// get a organization using organization id
	Get(context.Context, *model.OrgId) (*model.Organization, error)
	FindOrganizations(context.Context, *protobuf.CustomEmpty) (*model.Organizations, error)
}

// UnimplementedOrganizationQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedOrganizationQueryControllerServer struct {
}

func (UnimplementedOrganizationQueryControllerServer) Get(context.Context, *model.OrgId) (*model.Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedOrganizationQueryControllerServer) FindOrganizations(context.Context, *protobuf.CustomEmpty) (*model.Organizations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrganizations not implemented")
}

// UnsafeOrganizationQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationQueryControllerServer will
// result in compilation errors.
type UnsafeOrganizationQueryControllerServer interface {
	mustEmbedUnimplementedOrganizationQueryControllerServer()
}

func RegisterOrganizationQueryControllerServer(s grpc.ServiceRegistrar, srv OrganizationQueryControllerServer) {
	s.RegisterService(&OrganizationQueryController_ServiceDesc, srv)
}

func _OrganizationQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OrgId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrganizationQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationQueryControllerServer).Get(ctx, req.(*model.OrgId))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationQueryController_FindOrganizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.CustomEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationQueryControllerServer).FindOrganizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrganizationQueryController_FindOrganizations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationQueryControllerServer).FindOrganizations(ctx, req.(*protobuf.CustomEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationQueryController_ServiceDesc is the grpc.ServiceDesc for OrganizationQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.resourcemanager.v1.organization.service.OrganizationQueryController",
	HandlerType: (*OrganizationQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _OrganizationQueryController_Get_Handler,
		},
		{
			MethodName: "findOrganizations",
			Handler:    _OrganizationQueryController_FindOrganizations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/resourcemanager/v1/organization/service/query.proto",
}
