// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/connect/v1/azurecredential/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/azurecredential/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AzureCredentialQueryController_Get_FullMethodName = "/cloud.planton.apis.connect.v1.azurecredential.service.AzureCredentialQueryController/get"
)

// AzureCredentialQueryControllerClient is the client API for AzureCredentialQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AzureCredentialQueryControllerClient interface {
	// look up a azure-credential by id
	Get(ctx context.Context, in *model.AzureCredentialId, opts ...grpc.CallOption) (*model.AzureCredential, error)
}

type azureCredentialQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewAzureCredentialQueryControllerClient(cc grpc.ClientConnInterface) AzureCredentialQueryControllerClient {
	return &azureCredentialQueryControllerClient{cc}
}

func (c *azureCredentialQueryControllerClient) Get(ctx context.Context, in *model.AzureCredentialId, opts ...grpc.CallOption) (*model.AzureCredential, error) {
	out := new(model.AzureCredential)
	err := c.cc.Invoke(ctx, AzureCredentialQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AzureCredentialQueryControllerServer is the server API for AzureCredentialQueryController service.
// All implementations should embed UnimplementedAzureCredentialQueryControllerServer
// for forward compatibility
type AzureCredentialQueryControllerServer interface {
	// look up a azure-credential by id
	Get(context.Context, *model.AzureCredentialId) (*model.AzureCredential, error)
}

// UnimplementedAzureCredentialQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedAzureCredentialQueryControllerServer struct {
}

func (UnimplementedAzureCredentialQueryControllerServer) Get(context.Context, *model.AzureCredentialId) (*model.AzureCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeAzureCredentialQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AzureCredentialQueryControllerServer will
// result in compilation errors.
type UnsafeAzureCredentialQueryControllerServer interface {
	mustEmbedUnimplementedAzureCredentialQueryControllerServer()
}

func RegisterAzureCredentialQueryControllerServer(s grpc.ServiceRegistrar, srv AzureCredentialQueryControllerServer) {
	s.RegisterService(&AzureCredentialQueryController_ServiceDesc, srv)
}

func _AzureCredentialQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AzureCredentialId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AzureCredentialQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AzureCredentialQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AzureCredentialQueryControllerServer).Get(ctx, req.(*model.AzureCredentialId))
	}
	return interceptor(ctx, in, info, handler)
}

// AzureCredentialQueryController_ServiceDesc is the grpc.ServiceDesc for AzureCredentialQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AzureCredentialQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.connect.v1.azurecredential.service.AzureCredentialQueryController",
	HandlerType: (*AzureCredentialQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _AzureCredentialQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/connect/v1/azurecredential/service/query.proto",
}
