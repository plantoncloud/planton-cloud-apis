// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/iac/v1/stackjob/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StackJobCommandController_Create_FullMethodName        = "/cloud.planton.apis.iac.v1.stack.job.service.StackJobCommandController/create"
	StackJobCommandController_Update_FullMethodName        = "/cloud.planton.apis.iac.v1.stack.job.service.StackJobCommandController/update"
	StackJobCommandController_NotifyRunning_FullMethodName = "/cloud.planton.apis.iac.v1.stack.job.service.StackJobCommandController/notifyRunning"
)

// StackJobCommandControllerClient is the client API for StackJobCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StackJobCommandControllerClient interface {
	// create stack-job
	Create(ctx context.Context, in *model.StackJob, opts ...grpc.CallOption) (*model.StackJob, error)
	// update stack-job
	Update(ctx context.Context, in *model.StackJob, opts ...grpc.CallOption) (*model.StackJob, error)
	// rpc to set execution-status for a stack-job.
	// since the start of running of a stack-job happens in other services, stack service is notified of that event.
	// upon receiving this notification, stack service starts a stack-job progress follower that
	// takes care of updating the status of stack-job in the database.
	NotifyRunning(ctx context.Context, in *model.StackJobId, opts ...grpc.CallOption) (*model.StackJob, error)
}

type stackJobCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewStackJobCommandControllerClient(cc grpc.ClientConnInterface) StackJobCommandControllerClient {
	return &stackJobCommandControllerClient{cc}
}

func (c *stackJobCommandControllerClient) Create(ctx context.Context, in *model.StackJob, opts ...grpc.CallOption) (*model.StackJob, error) {
	out := new(model.StackJob)
	err := c.cc.Invoke(ctx, StackJobCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackJobCommandControllerClient) Update(ctx context.Context, in *model.StackJob, opts ...grpc.CallOption) (*model.StackJob, error) {
	out := new(model.StackJob)
	err := c.cc.Invoke(ctx, StackJobCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackJobCommandControllerClient) NotifyRunning(ctx context.Context, in *model.StackJobId, opts ...grpc.CallOption) (*model.StackJob, error) {
	out := new(model.StackJob)
	err := c.cc.Invoke(ctx, StackJobCommandController_NotifyRunning_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StackJobCommandControllerServer is the server API for StackJobCommandController service.
// All implementations should embed UnimplementedStackJobCommandControllerServer
// for forward compatibility
type StackJobCommandControllerServer interface {
	// create stack-job
	Create(context.Context, *model.StackJob) (*model.StackJob, error)
	// update stack-job
	Update(context.Context, *model.StackJob) (*model.StackJob, error)
	// rpc to set execution-status for a stack-job.
	// since the start of running of a stack-job happens in other services, stack service is notified of that event.
	// upon receiving this notification, stack service starts a stack-job progress follower that
	// takes care of updating the status of stack-job in the database.
	NotifyRunning(context.Context, *model.StackJobId) (*model.StackJob, error)
}

// UnimplementedStackJobCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedStackJobCommandControllerServer struct {
}

func (UnimplementedStackJobCommandControllerServer) Create(context.Context, *model.StackJob) (*model.StackJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStackJobCommandControllerServer) Update(context.Context, *model.StackJob) (*model.StackJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStackJobCommandControllerServer) NotifyRunning(context.Context, *model.StackJobId) (*model.StackJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyRunning not implemented")
}

// UnsafeStackJobCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StackJobCommandControllerServer will
// result in compilation errors.
type UnsafeStackJobCommandControllerServer interface {
	mustEmbedUnimplementedStackJobCommandControllerServer()
}

func RegisterStackJobCommandControllerServer(s grpc.ServiceRegistrar, srv StackJobCommandControllerServer) {
	s.RegisterService(&StackJobCommandController_ServiceDesc, srv)
}

func _StackJobCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.StackJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackJobCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackJobCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackJobCommandControllerServer).Create(ctx, req.(*model.StackJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackJobCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.StackJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackJobCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackJobCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackJobCommandControllerServer).Update(ctx, req.(*model.StackJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackJobCommandController_NotifyRunning_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.StackJobId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackJobCommandControllerServer).NotifyRunning(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackJobCommandController_NotifyRunning_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackJobCommandControllerServer).NotifyRunning(ctx, req.(*model.StackJobId))
	}
	return interceptor(ctx, in, info, handler)
}

// StackJobCommandController_ServiceDesc is the grpc.ServiceDesc for StackJobCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StackJobCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.iac.v1.stack.job.service.StackJobCommandController",
	HandlerType: (*StackJobCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _StackJobCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _StackJobCommandController_Update_Handler,
		},
		{
			MethodName: "notifyRunning",
			Handler:    _StackJobCommandController_NotifyRunning_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/iac/v1/stackjob/service/command.proto",
}
