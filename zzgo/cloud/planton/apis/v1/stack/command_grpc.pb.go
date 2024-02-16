// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/stack/command.proto

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
	StackCommandController_Create_FullMethodName              = "/cloud.planton.apis.v1.stack.StackCommandController/create"
	StackCommandController_Update_FullMethodName              = "/cloud.planton.apis.v1.stack.StackCommandController/update"
	StackCommandController_Delete_FullMethodName              = "/cloud.planton.apis.v1.stack.StackCommandController/delete"
	StackCommandController_DeleteOnPulumiCloud_FullMethodName = "/cloud.planton.apis.v1.stack.StackCommandController/deleteOnPulumiCloud"
)

// StackCommandControllerClient is the client API for StackCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StackCommandControllerClient interface {
	// create stack
	Create(ctx context.Context, in *Stack, opts ...grpc.CallOption) (*Stack, error)
	// update stack
	Update(ctx context.Context, in *Stack, opts ...grpc.CallOption) (*Stack, error)
	// delete stack
	Delete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*Stack, error)
	// delete stack on pulumi cloud
	DeleteOnPulumiCloud(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*Stack, error)
}

type stackCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewStackCommandControllerClient(cc grpc.ClientConnInterface) StackCommandControllerClient {
	return &stackCommandControllerClient{cc}
}

func (c *stackCommandControllerClient) Create(ctx context.Context, in *Stack, opts ...grpc.CallOption) (*Stack, error) {
	out := new(Stack)
	err := c.cc.Invoke(ctx, StackCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackCommandControllerClient) Update(ctx context.Context, in *Stack, opts ...grpc.CallOption) (*Stack, error) {
	out := new(Stack)
	err := c.cc.Invoke(ctx, StackCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackCommandControllerClient) Delete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*Stack, error) {
	out := new(Stack)
	err := c.cc.Invoke(ctx, StackCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackCommandControllerClient) DeleteOnPulumiCloud(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*Stack, error) {
	out := new(Stack)
	err := c.cc.Invoke(ctx, StackCommandController_DeleteOnPulumiCloud_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StackCommandControllerServer is the server API for StackCommandController service.
// All implementations should embed UnimplementedStackCommandControllerServer
// for forward compatibility
type StackCommandControllerServer interface {
	// create stack
	Create(context.Context, *Stack) (*Stack, error)
	// update stack
	Update(context.Context, *Stack) (*Stack, error)
	// delete stack
	Delete(context.Context, *resource.ApiResourceDeleteCommandInput) (*Stack, error)
	// delete stack on pulumi cloud
	DeleteOnPulumiCloud(context.Context, *resource.ApiResourceDeleteCommandInput) (*Stack, error)
}

// UnimplementedStackCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedStackCommandControllerServer struct {
}

func (UnimplementedStackCommandControllerServer) Create(context.Context, *Stack) (*Stack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStackCommandControllerServer) Update(context.Context, *Stack) (*Stack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStackCommandControllerServer) Delete(context.Context, *resource.ApiResourceDeleteCommandInput) (*Stack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStackCommandControllerServer) DeleteOnPulumiCloud(context.Context, *resource.ApiResourceDeleteCommandInput) (*Stack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOnPulumiCloud not implemented")
}

// UnsafeStackCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StackCommandControllerServer will
// result in compilation errors.
type UnsafeStackCommandControllerServer interface {
	mustEmbedUnimplementedStackCommandControllerServer()
}

func RegisterStackCommandControllerServer(s grpc.ServiceRegistrar, srv StackCommandControllerServer) {
	s.RegisterService(&StackCommandController_ServiceDesc, srv)
}

func _StackCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stack)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackCommandControllerServer).Create(ctx, req.(*Stack))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stack)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackCommandControllerServer).Update(ctx, req.(*Stack))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackCommandControllerServer).Delete(ctx, req.(*resource.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackCommandController_DeleteOnPulumiCloud_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackCommandControllerServer).DeleteOnPulumiCloud(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackCommandController_DeleteOnPulumiCloud_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackCommandControllerServer).DeleteOnPulumiCloud(ctx, req.(*resource.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// StackCommandController_ServiceDesc is the grpc.ServiceDesc for StackCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StackCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.stack.StackCommandController",
	HandlerType: (*StackCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _StackCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _StackCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _StackCommandController_Delete_Handler,
		},
		{
			MethodName: "deleteOnPulumiCloud",
			Handler:    _StackCommandController_DeleteOnPulumiCloud_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/stack/command.proto",
}
