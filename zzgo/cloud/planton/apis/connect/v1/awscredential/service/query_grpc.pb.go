// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/connect/v1/awscredential/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/awscredential/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AwsCredentialQueryController_Get_FullMethodName = "/cloud.planton.apis.connect.v1.awscredential.service.AwsCredentialQueryController/get"
)

// AwsCredentialQueryControllerClient is the client API for AwsCredentialQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AwsCredentialQueryControllerClient interface {
	// look up a aws-credential by id
	Get(ctx context.Context, in *model.AwsCredentialId, opts ...grpc.CallOption) (*model.AwsCredential, error)
}

type awsCredentialQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewAwsCredentialQueryControllerClient(cc grpc.ClientConnInterface) AwsCredentialQueryControllerClient {
	return &awsCredentialQueryControllerClient{cc}
}

func (c *awsCredentialQueryControllerClient) Get(ctx context.Context, in *model.AwsCredentialId, opts ...grpc.CallOption) (*model.AwsCredential, error) {
	out := new(model.AwsCredential)
	err := c.cc.Invoke(ctx, AwsCredentialQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AwsCredentialQueryControllerServer is the server API for AwsCredentialQueryController service.
// All implementations should embed UnimplementedAwsCredentialQueryControllerServer
// for forward compatibility
type AwsCredentialQueryControllerServer interface {
	// look up a aws-credential by id
	Get(context.Context, *model.AwsCredentialId) (*model.AwsCredential, error)
}

// UnimplementedAwsCredentialQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedAwsCredentialQueryControllerServer struct {
}

func (UnimplementedAwsCredentialQueryControllerServer) Get(context.Context, *model.AwsCredentialId) (*model.AwsCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeAwsCredentialQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AwsCredentialQueryControllerServer will
// result in compilation errors.
type UnsafeAwsCredentialQueryControllerServer interface {
	mustEmbedUnimplementedAwsCredentialQueryControllerServer()
}

func RegisterAwsCredentialQueryControllerServer(s grpc.ServiceRegistrar, srv AwsCredentialQueryControllerServer) {
	s.RegisterService(&AwsCredentialQueryController_ServiceDesc, srv)
}

func _AwsCredentialQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AwsCredentialId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsCredentialQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AwsCredentialQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsCredentialQueryControllerServer).Get(ctx, req.(*model.AwsCredentialId))
	}
	return interceptor(ctx, in, info, handler)
}

// AwsCredentialQueryController_ServiceDesc is the grpc.ServiceDesc for AwsCredentialQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AwsCredentialQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.connect.v1.awscredential.service.AwsCredentialQueryController",
	HandlerType: (*AwsCredentialQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _AwsCredentialQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/connect/v1/awscredential/service/query.proto",
}
