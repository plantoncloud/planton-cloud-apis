// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/iam/v1/identityconnection/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iam/v1/identityconnection/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IdentityConnectionQueryController_GetById_FullMethodName     = "/cloud.planton.apis.iam.v1.identityconnection.service.IdentityConnectionQueryController/getById"
	IdentityConnectionQueryController_FindByOrgId_FullMethodName = "/cloud.planton.apis.iam.v1.identityconnection.service.IdentityConnectionQueryController/findByOrgId"
)

// IdentityConnectionQueryControllerClient is the client API for IdentityConnectionQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityConnectionQueryControllerClient interface {
	// lookup identity account by id.
	GetById(ctx context.Context, in *model.IdentityConnectionId, opts ...grpc.CallOption) (*model.IdentityConnection, error)
	// retrieve paginated list of all identity connections on planton-cloud.
	FindByOrgId(ctx context.Context, in *model.IdentityConnectionOrgId, opts ...grpc.CallOption) (*model.IdentityConnections, error)
}

type identityConnectionQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityConnectionQueryControllerClient(cc grpc.ClientConnInterface) IdentityConnectionQueryControllerClient {
	return &identityConnectionQueryControllerClient{cc}
}

func (c *identityConnectionQueryControllerClient) GetById(ctx context.Context, in *model.IdentityConnectionId, opts ...grpc.CallOption) (*model.IdentityConnection, error) {
	out := new(model.IdentityConnection)
	err := c.cc.Invoke(ctx, IdentityConnectionQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityConnectionQueryControllerClient) FindByOrgId(ctx context.Context, in *model.IdentityConnectionOrgId, opts ...grpc.CallOption) (*model.IdentityConnections, error) {
	out := new(model.IdentityConnections)
	err := c.cc.Invoke(ctx, IdentityConnectionQueryController_FindByOrgId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityConnectionQueryControllerServer is the server API for IdentityConnectionQueryController service.
// All implementations should embed UnimplementedIdentityConnectionQueryControllerServer
// for forward compatibility
type IdentityConnectionQueryControllerServer interface {
	// lookup identity account by id.
	GetById(context.Context, *model.IdentityConnectionId) (*model.IdentityConnection, error)
	// retrieve paginated list of all identity connections on planton-cloud.
	FindByOrgId(context.Context, *model.IdentityConnectionOrgId) (*model.IdentityConnections, error)
}

// UnimplementedIdentityConnectionQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedIdentityConnectionQueryControllerServer struct {
}

func (UnimplementedIdentityConnectionQueryControllerServer) GetById(context.Context, *model.IdentityConnectionId) (*model.IdentityConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedIdentityConnectionQueryControllerServer) FindByOrgId(context.Context, *model.IdentityConnectionOrgId) (*model.IdentityConnections, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByOrgId not implemented")
}

// UnsafeIdentityConnectionQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityConnectionQueryControllerServer will
// result in compilation errors.
type UnsafeIdentityConnectionQueryControllerServer interface {
	mustEmbedUnimplementedIdentityConnectionQueryControllerServer()
}

func RegisterIdentityConnectionQueryControllerServer(s grpc.ServiceRegistrar, srv IdentityConnectionQueryControllerServer) {
	s.RegisterService(&IdentityConnectionQueryController_ServiceDesc, srv)
}

func _IdentityConnectionQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IdentityConnectionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityConnectionQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityConnectionQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityConnectionQueryControllerServer).GetById(ctx, req.(*model.IdentityConnectionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityConnectionQueryController_FindByOrgId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IdentityConnectionOrgId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityConnectionQueryControllerServer).FindByOrgId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityConnectionQueryController_FindByOrgId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityConnectionQueryControllerServer).FindByOrgId(ctx, req.(*model.IdentityConnectionOrgId))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityConnectionQueryController_ServiceDesc is the grpc.ServiceDesc for IdentityConnectionQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityConnectionQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.iam.v1.identityconnection.service.IdentityConnectionQueryController",
	HandlerType: (*IdentityConnectionQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _IdentityConnectionQueryController_GetById_Handler,
		},
		{
			MethodName: "findByOrgId",
			Handler:    _IdentityConnectionQueryController_FindByOrgId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/iam/v1/identityconnection/service/query.proto",
}
