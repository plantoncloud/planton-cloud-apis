// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/aws/route53zone/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/aws/route53zone/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Route53ZoneQueryController_Get_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneQueryController/get"
)

// Route53ZoneQueryControllerClient is the client API for Route53ZoneQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Route53ZoneQueryControllerClient interface {
	// get details of a route53-zone id
	Get(ctx context.Context, in *model.Route53ZoneId, opts ...grpc.CallOption) (*model.Route53Zone, error)
}

type route53ZoneQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewRoute53ZoneQueryControllerClient(cc grpc.ClientConnInterface) Route53ZoneQueryControllerClient {
	return &route53ZoneQueryControllerClient{cc}
}

func (c *route53ZoneQueryControllerClient) Get(ctx context.Context, in *model.Route53ZoneId, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Route53ZoneQueryControllerServer is the server API for Route53ZoneQueryController service.
// All implementations should embed UnimplementedRoute53ZoneQueryControllerServer
// for forward compatibility
type Route53ZoneQueryControllerServer interface {
	// get details of a route53-zone id
	Get(context.Context, *model.Route53ZoneId) (*model.Route53Zone, error)
}

// UnimplementedRoute53ZoneQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedRoute53ZoneQueryControllerServer struct {
}

func (UnimplementedRoute53ZoneQueryControllerServer) Get(context.Context, *model.Route53ZoneId) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeRoute53ZoneQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Route53ZoneQueryControllerServer will
// result in compilation errors.
type UnsafeRoute53ZoneQueryControllerServer interface {
	mustEmbedUnimplementedRoute53ZoneQueryControllerServer()
}

func RegisterRoute53ZoneQueryControllerServer(s grpc.ServiceRegistrar, srv Route53ZoneQueryControllerServer) {
	s.RegisterService(&Route53ZoneQueryController_ServiceDesc, srv)
}

func _Route53ZoneQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Route53ZoneId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneQueryControllerServer).Get(ctx, req.(*model.Route53ZoneId))
	}
	return interceptor(ctx, in, info, handler)
}

// Route53ZoneQueryController_ServiceDesc is the grpc.ServiceDesc for Route53ZoneQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Route53ZoneQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneQueryController",
	HandlerType: (*Route53ZoneQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _Route53ZoneQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/aws/route53zone/service/query.proto",
}
