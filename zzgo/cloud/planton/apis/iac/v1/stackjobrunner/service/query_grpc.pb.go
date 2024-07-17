// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/iac/v1/stackjobrunner/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjobrunner/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StackJobRunnerQueryController_Get_FullMethodName = "/cloud.planton.apis.iac.v1.stackjobrunner.service.StackJobRunnerQueryController/get"
)

// StackJobRunnerQueryControllerClient is the client API for StackJobRunnerQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StackJobRunnerQueryControllerClient interface {
	// look up a stack-job-runner using id
	Get(ctx context.Context, in *model.StackJobRunnerId, opts ...grpc.CallOption) (*model.StackJobRunner, error)
}

type stackJobRunnerQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewStackJobRunnerQueryControllerClient(cc grpc.ClientConnInterface) StackJobRunnerQueryControllerClient {
	return &stackJobRunnerQueryControllerClient{cc}
}

func (c *stackJobRunnerQueryControllerClient) Get(ctx context.Context, in *model.StackJobRunnerId, opts ...grpc.CallOption) (*model.StackJobRunner, error) {
	out := new(model.StackJobRunner)
	err := c.cc.Invoke(ctx, StackJobRunnerQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StackJobRunnerQueryControllerServer is the server API for StackJobRunnerQueryController service.
// All implementations should embed UnimplementedStackJobRunnerQueryControllerServer
// for forward compatibility
type StackJobRunnerQueryControllerServer interface {
	// look up a stack-job-runner using id
	Get(context.Context, *model.StackJobRunnerId) (*model.StackJobRunner, error)
}

// UnimplementedStackJobRunnerQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedStackJobRunnerQueryControllerServer struct {
}

func (UnimplementedStackJobRunnerQueryControllerServer) Get(context.Context, *model.StackJobRunnerId) (*model.StackJobRunner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeStackJobRunnerQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StackJobRunnerQueryControllerServer will
// result in compilation errors.
type UnsafeStackJobRunnerQueryControllerServer interface {
	mustEmbedUnimplementedStackJobRunnerQueryControllerServer()
}

func RegisterStackJobRunnerQueryControllerServer(s grpc.ServiceRegistrar, srv StackJobRunnerQueryControllerServer) {
	s.RegisterService(&StackJobRunnerQueryController_ServiceDesc, srv)
}

func _StackJobRunnerQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.StackJobRunnerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackJobRunnerQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackJobRunnerQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackJobRunnerQueryControllerServer).Get(ctx, req.(*model.StackJobRunnerId))
	}
	return interceptor(ctx, in, info, handler)
}

// StackJobRunnerQueryController_ServiceDesc is the grpc.ServiceDesc for StackJobRunnerQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StackJobRunnerQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.iac.v1.stackjobrunner.service.StackJobRunnerQueryController",
	HandlerType: (*StackJobRunnerQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _StackJobRunnerQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/iac/v1/stackjobrunner/service/query.proto",
}
