// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/integration/cloudprovider/gcp/service.proto

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
	GcpCommandController_AuthorizeCloudAccount_FullMethodName = "/cloud.planton.apis.v1.integration.cloudprovider.gcp.GcpCommandController/authorizeCloudAccount"
)

// GcpCommandControllerClient is the client API for GcpCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GcpCommandControllerClient interface {
	// authorizes the service-account to be used for authenticating with google cloud for the specific cloud-account.
	AuthorizeCloudAccount(ctx context.Context, in *AuthorizeGcpCloudAccountCommandInput, opts ...grpc.CallOption) (*AuthorizeGcpCloudAccountCommandResponse, error)
}

type gcpCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGcpCommandControllerClient(cc grpc.ClientConnInterface) GcpCommandControllerClient {
	return &gcpCommandControllerClient{cc}
}

func (c *gcpCommandControllerClient) AuthorizeCloudAccount(ctx context.Context, in *AuthorizeGcpCloudAccountCommandInput, opts ...grpc.CallOption) (*AuthorizeGcpCloudAccountCommandResponse, error) {
	out := new(AuthorizeGcpCloudAccountCommandResponse)
	err := c.cc.Invoke(ctx, GcpCommandController_AuthorizeCloudAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GcpCommandControllerServer is the server API for GcpCommandController service.
// All implementations should embed UnimplementedGcpCommandControllerServer
// for forward compatibility
type GcpCommandControllerServer interface {
	// authorizes the service-account to be used for authenticating with google cloud for the specific cloud-account.
	AuthorizeCloudAccount(context.Context, *AuthorizeGcpCloudAccountCommandInput) (*AuthorizeGcpCloudAccountCommandResponse, error)
}

// UnimplementedGcpCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGcpCommandControllerServer struct {
}

func (UnimplementedGcpCommandControllerServer) AuthorizeCloudAccount(context.Context, *AuthorizeGcpCloudAccountCommandInput) (*AuthorizeGcpCloudAccountCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeCloudAccount not implemented")
}

// UnsafeGcpCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GcpCommandControllerServer will
// result in compilation errors.
type UnsafeGcpCommandControllerServer interface {
	mustEmbedUnimplementedGcpCommandControllerServer()
}

func RegisterGcpCommandControllerServer(s grpc.ServiceRegistrar, srv GcpCommandControllerServer) {
	s.RegisterService(&GcpCommandController_ServiceDesc, srv)
}

func _GcpCommandController_AuthorizeCloudAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeGcpCloudAccountCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCommandControllerServer).AuthorizeCloudAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCommandController_AuthorizeCloudAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCommandControllerServer).AuthorizeCloudAccount(ctx, req.(*AuthorizeGcpCloudAccountCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GcpCommandController_ServiceDesc is the grpc.ServiceDesc for GcpCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GcpCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.integration.cloudprovider.gcp.GcpCommandController",
	HandlerType: (*GcpCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "authorizeCloudAccount",
			Handler:    _GcpCommandController_AuthorizeCloudAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/integration/cloudprovider/gcp/service.proto",
}
