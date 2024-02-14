// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/stack/query.proto

package stack

import (
	context "context"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StackQueryController_GetById_FullMethodName         = "/cloud.planton.apis.v1.stack.StackQueryController/getById"
	StackQueryController_GetByResourceId_FullMethodName = "/cloud.planton.apis.v1.stack.StackQueryController/getByResourceId"
)

// StackQueryControllerClient is the client API for StackQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StackQueryControllerClient interface {
	// look up stack by stack-id
	GetById(ctx context.Context, in *StackId, opts ...grpc.CallOption) (*Stack, error)
	// lookup a stack by resource-id
	GetByResourceId(ctx context.Context, in *resource.ResourceId, opts ...grpc.CallOption) (*Stack, error)
}

type stackQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewStackQueryControllerClient(cc grpc.ClientConnInterface) StackQueryControllerClient {
	return &stackQueryControllerClient{cc}
}

func (c *stackQueryControllerClient) GetById(ctx context.Context, in *StackId, opts ...grpc.CallOption) (*Stack, error) {
	out := new(Stack)
	err := c.cc.Invoke(ctx, StackQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackQueryControllerClient) GetByResourceId(ctx context.Context, in *resource.ResourceId, opts ...grpc.CallOption) (*Stack, error) {
	out := new(Stack)
	err := c.cc.Invoke(ctx, StackQueryController_GetByResourceId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StackQueryControllerServer is the server API for StackQueryController service.
// All implementations should embed UnimplementedStackQueryControllerServer
// for forward compatibility
type StackQueryControllerServer interface {
	// look up stack by stack-id
	GetById(context.Context, *StackId) (*Stack, error)
	// lookup a stack by resource-id
	GetByResourceId(context.Context, *resource.ResourceId) (*Stack, error)
}

// UnimplementedStackQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedStackQueryControllerServer struct {
}

func (UnimplementedStackQueryControllerServer) GetById(context.Context, *StackId) (*Stack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedStackQueryControllerServer) GetByResourceId(context.Context, *resource.ResourceId) (*Stack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByResourceId not implemented")
}

// UnsafeStackQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StackQueryControllerServer will
// result in compilation errors.
type UnsafeStackQueryControllerServer interface {
	mustEmbedUnimplementedStackQueryControllerServer()
}

func RegisterStackQueryControllerServer(s grpc.ServiceRegistrar, srv StackQueryControllerServer) {
	s.RegisterService(&StackQueryController_ServiceDesc, srv)
}

func _StackQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackQueryControllerServer).GetById(ctx, req.(*StackId))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackQueryController_GetByResourceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackQueryControllerServer).GetByResourceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackQueryController_GetByResourceId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackQueryControllerServer).GetByResourceId(ctx, req.(*resource.ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

// StackQueryController_ServiceDesc is the grpc.ServiceDesc for StackQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StackQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.stack.StackQueryController",
	HandlerType: (*StackQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _StackQueryController_GetById_Handler,
		},
		{
			MethodName: "getByResourceId",
			Handler:    _StackQueryController_GetByResourceId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/stack/query.proto",
}