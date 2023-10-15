// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/storagebucket/stack/gcp/service.proto

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
	StorageBucketGcpStackController_Execute_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackController/execute"
	StorageBucketGcpStackController_GetStackOutputs_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackController/getStackOutputs"
)

// StorageBucketGcpStackControllerClient is the client API for StorageBucketGcpStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageBucketGcpStackControllerClient interface {
	Execute(ctx context.Context, in *StorageBucketGcpStackInput, opts ...grpc.CallOption) (StorageBucketGcpStackController_ExecuteClient, error)
	GetStackOutputs(ctx context.Context, in *StorageBucketGcpStackInput, opts ...grpc.CallOption) (*StorageBucketGcpStackOutputs, error)
}

type storageBucketGcpStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageBucketGcpStackControllerClient(cc grpc.ClientConnInterface) StorageBucketGcpStackControllerClient {
	return &storageBucketGcpStackControllerClient{cc}
}

func (c *storageBucketGcpStackControllerClient) Execute(ctx context.Context, in *StorageBucketGcpStackInput, opts ...grpc.CallOption) (StorageBucketGcpStackController_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &StorageBucketGcpStackController_ServiceDesc.Streams[0], StorageBucketGcpStackController_Execute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &storageBucketGcpStackControllerExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StorageBucketGcpStackController_ExecuteClient interface {
	Recv() (*StorageBucketGcpStackResponse, error)
	grpc.ClientStream
}

type storageBucketGcpStackControllerExecuteClient struct {
	grpc.ClientStream
}

func (x *storageBucketGcpStackControllerExecuteClient) Recv() (*StorageBucketGcpStackResponse, error) {
	m := new(StorageBucketGcpStackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageBucketGcpStackControllerClient) GetStackOutputs(ctx context.Context, in *StorageBucketGcpStackInput, opts ...grpc.CallOption) (*StorageBucketGcpStackOutputs, error) {
	out := new(StorageBucketGcpStackOutputs)
	err := c.cc.Invoke(ctx, StorageBucketGcpStackController_GetStackOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageBucketGcpStackControllerServer is the server API for StorageBucketGcpStackController service.
// All implementations should embed UnimplementedStorageBucketGcpStackControllerServer
// for forward compatibility
type StorageBucketGcpStackControllerServer interface {
	Execute(*StorageBucketGcpStackInput, StorageBucketGcpStackController_ExecuteServer) error
	GetStackOutputs(context.Context, *StorageBucketGcpStackInput) (*StorageBucketGcpStackOutputs, error)
}

// UnimplementedStorageBucketGcpStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedStorageBucketGcpStackControllerServer struct {
}

func (UnimplementedStorageBucketGcpStackControllerServer) Execute(*StorageBucketGcpStackInput, StorageBucketGcpStackController_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedStorageBucketGcpStackControllerServer) GetStackOutputs(context.Context, *StorageBucketGcpStackInput) (*StorageBucketGcpStackOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStackOutputs not implemented")
}

// UnsafeStorageBucketGcpStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageBucketGcpStackControllerServer will
// result in compilation errors.
type UnsafeStorageBucketGcpStackControllerServer interface {
	mustEmbedUnimplementedStorageBucketGcpStackControllerServer()
}

func RegisterStorageBucketGcpStackControllerServer(s grpc.ServiceRegistrar, srv StorageBucketGcpStackControllerServer) {
	s.RegisterService(&StorageBucketGcpStackController_ServiceDesc, srv)
}

func _StorageBucketGcpStackController_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StorageBucketGcpStackInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageBucketGcpStackControllerServer).Execute(m, &storageBucketGcpStackControllerExecuteServer{stream})
}

type StorageBucketGcpStackController_ExecuteServer interface {
	Send(*StorageBucketGcpStackResponse) error
	grpc.ServerStream
}

type storageBucketGcpStackControllerExecuteServer struct {
	grpc.ServerStream
}

func (x *storageBucketGcpStackControllerExecuteServer) Send(m *StorageBucketGcpStackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StorageBucketGcpStackController_GetStackOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageBucketGcpStackInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBucketGcpStackControllerServer).GetStackOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageBucketGcpStackController_GetStackOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBucketGcpStackControllerServer).GetStackOutputs(ctx, req.(*StorageBucketGcpStackInput))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageBucketGcpStackController_ServiceDesc is the grpc.ServiceDesc for StorageBucketGcpStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageBucketGcpStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackController",
	HandlerType: (*StorageBucketGcpStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStackOutputs",
			Handler:    _StorageBucketGcpStackController_GetStackOutputs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _StorageBucketGcpStackController_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/storagebucket/stack/gcp/service.proto",
}