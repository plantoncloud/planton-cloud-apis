// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/connect/gcpcredential/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/connect/gcpcredential/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GcpCredentialQueryController_Get_FullMethodName = "/cloud.planton.apis.code2cloud.v1.connect.gcpcredential.service.GcpCredentialQueryController/get"
)

// GcpCredentialQueryControllerClient is the client API for GcpCredentialQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GcpCredentialQueryControllerClient interface {
	// look up a gcp-credential by id
	Get(ctx context.Context, in *model.GcpCredentialId, opts ...grpc.CallOption) (*model.GcpCredential, error)
}

type gcpCredentialQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGcpCredentialQueryControllerClient(cc grpc.ClientConnInterface) GcpCredentialQueryControllerClient {
	return &gcpCredentialQueryControllerClient{cc}
}

func (c *gcpCredentialQueryControllerClient) Get(ctx context.Context, in *model.GcpCredentialId, opts ...grpc.CallOption) (*model.GcpCredential, error) {
	out := new(model.GcpCredential)
	err := c.cc.Invoke(ctx, GcpCredentialQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GcpCredentialQueryControllerServer is the server API for GcpCredentialQueryController service.
// All implementations should embed UnimplementedGcpCredentialQueryControllerServer
// for forward compatibility
type GcpCredentialQueryControllerServer interface {
	// look up a gcp-credential by id
	Get(context.Context, *model.GcpCredentialId) (*model.GcpCredential, error)
}

// UnimplementedGcpCredentialQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGcpCredentialQueryControllerServer struct {
}

func (UnimplementedGcpCredentialQueryControllerServer) Get(context.Context, *model.GcpCredentialId) (*model.GcpCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeGcpCredentialQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GcpCredentialQueryControllerServer will
// result in compilation errors.
type UnsafeGcpCredentialQueryControllerServer interface {
	mustEmbedUnimplementedGcpCredentialQueryControllerServer()
}

func RegisterGcpCredentialQueryControllerServer(s grpc.ServiceRegistrar, srv GcpCredentialQueryControllerServer) {
	s.RegisterService(&GcpCredentialQueryController_ServiceDesc, srv)
}

func _GcpCredentialQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpCredentialId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCredentialQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCredentialQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCredentialQueryControllerServer).Get(ctx, req.(*model.GcpCredentialId))
	}
	return interceptor(ctx, in, info, handler)
}

// GcpCredentialQueryController_ServiceDesc is the grpc.ServiceDesc for GcpCredentialQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GcpCredentialQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.connect.gcpcredential.service.GcpCredentialQueryController",
	HandlerType: (*GcpCredentialQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _GcpCredentialQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/connect/gcpcredential/service/query.proto",
}
