// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/gcp/gcpsecretsmanagersecretset/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/gcp/gcpsecretsmanagersecretset/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GcpSecretsManagerSecretSetQueryController_Get_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetQueryController/get"
)

// GcpSecretsManagerSecretSetQueryControllerClient is the client API for GcpSecretsManagerSecretSetQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GcpSecretsManagerSecretSetQueryControllerClient interface {
	// look up a gcp-secrets-manager-secret-set using gcp-secrets-manager-secret-set id
	Get(ctx context.Context, in *model.GcpSecretsManagerSecretSetId, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error)
}

type gcpSecretsManagerSecretSetQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGcpSecretsManagerSecretSetQueryControllerClient(cc grpc.ClientConnInterface) GcpSecretsManagerSecretSetQueryControllerClient {
	return &gcpSecretsManagerSecretSetQueryControllerClient{cc}
}

func (c *gcpSecretsManagerSecretSetQueryControllerClient) Get(ctx context.Context, in *model.GcpSecretsManagerSecretSetId, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error) {
	out := new(model.GcpSecretsManagerSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretsManagerSecretSetQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GcpSecretsManagerSecretSetQueryControllerServer is the server API for GcpSecretsManagerSecretSetQueryController service.
// All implementations should embed UnimplementedGcpSecretsManagerSecretSetQueryControllerServer
// for forward compatibility
type GcpSecretsManagerSecretSetQueryControllerServer interface {
	// look up a gcp-secrets-manager-secret-set using gcp-secrets-manager-secret-set id
	Get(context.Context, *model.GcpSecretsManagerSecretSetId) (*model.GcpSecretsManagerSecretSet, error)
}

// UnimplementedGcpSecretsManagerSecretSetQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGcpSecretsManagerSecretSetQueryControllerServer struct {
}

func (UnimplementedGcpSecretsManagerSecretSetQueryControllerServer) Get(context.Context, *model.GcpSecretsManagerSecretSetId) (*model.GcpSecretsManagerSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeGcpSecretsManagerSecretSetQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GcpSecretsManagerSecretSetQueryControllerServer will
// result in compilation errors.
type UnsafeGcpSecretsManagerSecretSetQueryControllerServer interface {
	mustEmbedUnimplementedGcpSecretsManagerSecretSetQueryControllerServer()
}

func RegisterGcpSecretsManagerSecretSetQueryControllerServer(s grpc.ServiceRegistrar, srv GcpSecretsManagerSecretSetQueryControllerServer) {
	s.RegisterService(&GcpSecretsManagerSecretSetQueryController_ServiceDesc, srv)
}

func _GcpSecretsManagerSecretSetQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretsManagerSecretSetId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretsManagerSecretSetQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretsManagerSecretSetQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretsManagerSecretSetQueryControllerServer).Get(ctx, req.(*model.GcpSecretsManagerSecretSetId))
	}
	return interceptor(ctx, in, info, handler)
}

// GcpSecretsManagerSecretSetQueryController_ServiceDesc is the grpc.ServiceDesc for GcpSecretsManagerSecretSetQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GcpSecretsManagerSecretSetQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetQueryController",
	HandlerType: (*GcpSecretsManagerSecretSetQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _GcpSecretsManagerSecretSetQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/gcp/gcpsecretsmanagersecretset/service/query.proto",
}
