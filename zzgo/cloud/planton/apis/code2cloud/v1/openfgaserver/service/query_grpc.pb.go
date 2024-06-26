// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/openfgaserver/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/openfgaserver/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OpenfgaServerQueryController_GetById_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerQueryController/getById"
	OpenfgaServerQueryController_GetPassword_FullMethodName = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerQueryController/getPassword"
)

// OpenfgaServerQueryControllerClient is the client API for OpenfgaServerQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenfgaServerQueryControllerClient interface {
	// look up openfga-server using openfga-server id
	GetById(ctx context.Context, in *model.OpenfgaServerId, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// look up openfga-server sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *model.OpenfgaServerId, opts ...grpc.CallOption) (*model.OpenfgaServerPassword, error)
}

type openfgaServerQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenfgaServerQueryControllerClient(cc grpc.ClientConnInterface) OpenfgaServerQueryControllerClient {
	return &openfgaServerQueryControllerClient{cc}
}

func (c *openfgaServerQueryControllerClient) GetById(ctx context.Context, in *model.OpenfgaServerId, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerQueryControllerClient) GetPassword(ctx context.Context, in *model.OpenfgaServerId, opts ...grpc.CallOption) (*model.OpenfgaServerPassword, error) {
	out := new(model.OpenfgaServerPassword)
	err := c.cc.Invoke(ctx, OpenfgaServerQueryController_GetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenfgaServerQueryControllerServer is the server API for OpenfgaServerQueryController service.
// All implementations should embed UnimplementedOpenfgaServerQueryControllerServer
// for forward compatibility
type OpenfgaServerQueryControllerServer interface {
	// look up openfga-server using openfga-server id
	GetById(context.Context, *model.OpenfgaServerId) (*model.OpenfgaServer, error)
	// look up openfga-server sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model.OpenfgaServerId) (*model.OpenfgaServerPassword, error)
}

// UnimplementedOpenfgaServerQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedOpenfgaServerQueryControllerServer struct {
}

func (UnimplementedOpenfgaServerQueryControllerServer) GetById(context.Context, *model.OpenfgaServerId) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedOpenfgaServerQueryControllerServer) GetPassword(context.Context, *model.OpenfgaServerId) (*model.OpenfgaServerPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}

// UnsafeOpenfgaServerQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenfgaServerQueryControllerServer will
// result in compilation errors.
type UnsafeOpenfgaServerQueryControllerServer interface {
	mustEmbedUnimplementedOpenfgaServerQueryControllerServer()
}

func RegisterOpenfgaServerQueryControllerServer(s grpc.ServiceRegistrar, srv OpenfgaServerQueryControllerServer) {
	s.RegisterService(&OpenfgaServerQueryController_ServiceDesc, srv)
}

func _OpenfgaServerQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenfgaServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerQueryControllerServer).GetById(ctx, req.(*model.OpenfgaServerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerQueryController_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenfgaServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerQueryControllerServer).GetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerQueryController_GetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerQueryControllerServer).GetPassword(ctx, req.(*model.OpenfgaServerId))
	}
	return interceptor(ctx, in, info, handler)
}

// OpenfgaServerQueryController_ServiceDesc is the grpc.ServiceDesc for OpenfgaServerQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenfgaServerQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerQueryController",
	HandlerType: (*OpenfgaServerQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _OpenfgaServerQueryController_GetById_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _OpenfgaServerQueryController_GetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/openfgaserver/service/query.proto",
}
