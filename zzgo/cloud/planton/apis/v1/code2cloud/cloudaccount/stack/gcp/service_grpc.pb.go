// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/cloudaccount/stack/gcp/service.proto

package gcp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CloudAccountGcpStackController_Execute_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.cloudaccount.stack.gcp.CloudAccountGcpStackController/execute"
	CloudAccountGcpStackController_GetStackOutputs_FullMethodName = "/cloud.planton.apis.v1.code2cloud.cloudaccount.stack.gcp.CloudAccountGcpStackController/getStackOutputs"
)

// CloudAccountGcpStackControllerClient is the client API for CloudAccountGcpStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudAccountGcpStackControllerClient interface {
	Execute(ctx context.Context, in *CloudAccountGcpStackInput, opts ...grpc.CallOption) (CloudAccountGcpStackController_ExecuteClient, error)
	GetStackOutputs(ctx context.Context, in *CloudAccountGcpStackInput, opts ...grpc.CallOption) (*CloudAccountGcpStackOutputs, error)
}

type cloudAccountGcpStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudAccountGcpStackControllerClient(cc grpc.ClientConnInterface) CloudAccountGcpStackControllerClient {
	return &cloudAccountGcpStackControllerClient{cc}
}

func (c *cloudAccountGcpStackControllerClient) Execute(ctx context.Context, in *CloudAccountGcpStackInput, opts ...grpc.CallOption) (CloudAccountGcpStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &CloudAccountGcpStackController_ServiceDesc.Streams[0], CloudAccountGcpStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &cloudAccountGcpStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CloudAccountGcpStackController_ExecuteClient interface {
	Recv() (*CloudAccountGcpStackResponse, error)
	grpc.ClientStream
}

type cloudAccountGcpStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *cloudAccountGcpStackControllerExecuteClient) Recv() (*CloudAccountGcpStackResponse, error) {
	m := new(CloudAccountGcpStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cloudAccountGcpStackControllerClient) GetStackOutputs(ctx context.Context, in *CloudAccountGcpStackInput, opts ...grpc.CallOption) (*CloudAccountGcpStackOutputs, error) {
	out := new(CloudAccountGcpStackOutputs)
	err := c.cc.Invoke(ctx, CloudAccountGcpStackController_GetStackOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudAccountGcpStackControllerServer is the server API for CloudAccountGcpStackController service.
// All implementations should embed UnimplementedCloudAccountGcpStackControllerServer
// for forward compatibility
type CloudAccountGcpStackControllerServer interface {
	Execute(*CloudAccountGcpStackInput, CloudAccountGcpStackController_ExecuteServer) error
	GetStackOutputs(context.Context, *CloudAccountGcpStackInput) (*CloudAccountGcpStackOutputs, error)
}

// UnimplementedCloudAccountGcpStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCloudAccountGcpStackControllerServer struct {
}

func (UnimplementedCloudAccountGcpStackControllerServer) Execute(*CloudAccountGcpStackInput, CloudAccountGcpStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedCloudAccountGcpStackControllerServer) GetStackOutputs(context.Context, *CloudAccountGcpStackInput) (*CloudAccountGcpStackOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackOutputs not implemented")
}

// UnsafeCloudAccountGcpStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudAccountGcpStackControllerServer will
// result in compilation errors.
type UnsafeCloudAccountGcpStackControllerServer interface {
	mustEmbedUnimplementedCloudAccountGcpStackControllerServer()
}

func RegisterCloudAccountGcpStackControllerServer(s grpc.ServiceRegistrar, srv CloudAccountGcpStackControllerServer) {
	s.RegisterService(&CloudAccountGcpStackController_ServiceDesc, srv)
}

func _CloudAccountGcpStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CloudAccountGcpStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CloudAccountGcpStackControllerServer).Execute(m, &cloudAccountGcpStackControllerExecuteServer{stream})
}

type CloudAccountGcpStackController_ExecuteServer interface {
	Send(*CloudAccountGcpStackResponse) error
	grpc.ServerStream
}

type cloudAccountGcpStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *cloudAccountGcpStackControllerExecuteServer) Send(m *CloudAccountGcpStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CloudAccountGcpStackController_GetStackOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudAccountGcpStackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAccountGcpStackControllerServer).GetStackOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudAccountGcpStackController_GetStackOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAccountGcpStackControllerServer).GetStackOutputs(ctx, req.(*CloudAccountGcpStackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudAccountGcpStackController_ServiceDesc is the grpc.ServiceDesc for CloudAccountGcpStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudAccountGcpStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.cloudaccount.stack.gcp.CloudAccountGcpStackController",
	HandlerType: (*CloudAccountGcpStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStackOutputs",
			Handler:    _CloudAccountGcpStackController_GetStackOutputs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _CloudAccountGcpStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/v1/code2cloud/cloudaccount/stack/gcp/service.proto",
}
