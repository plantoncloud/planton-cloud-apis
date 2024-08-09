// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/aws/awssecretset/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/awssecretset/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AwsSecretSetQueryController_Get_FullMethodName = "/cloud.planton.apis.code2cloud.v1.aws.awssecretset.service.AwsSecretSetQueryController/get"
)

// AwsSecretSetQueryControllerClient is the client API for AwsSecretSetQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AwsSecretSetQueryControllerClient interface {
	// look up a aws-secret-set using aws-secret-set id
	Get(ctx context.Context, in *model.AwsSecretSetId, opts ...grpc.CallOption) (*model.AwsSecretSet, error)
}

type awsSecretSetQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewAwsSecretSetQueryControllerClient(cc grpc.ClientConnInterface) AwsSecretSetQueryControllerClient {
	return &awsSecretSetQueryControllerClient{cc}
}

func (c *awsSecretSetQueryControllerClient) Get(ctx context.Context, in *model.AwsSecretSetId, opts ...grpc.CallOption) (*model.AwsSecretSet, error) {
	out := new(model.AwsSecretSet)
	err := c.cc.Invoke(ctx, AwsSecretSetQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AwsSecretSetQueryControllerServer is the server API for AwsSecretSetQueryController service.
// All implementations should embed UnimplementedAwsSecretSetQueryControllerServer
// for forward compatibility
type AwsSecretSetQueryControllerServer interface {
	// look up a aws-secret-set using aws-secret-set id
	Get(context.Context, *model.AwsSecretSetId) (*model.AwsSecretSet, error)
}

// UnimplementedAwsSecretSetQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedAwsSecretSetQueryControllerServer struct {
}

func (UnimplementedAwsSecretSetQueryControllerServer) Get(context.Context, *model.AwsSecretSetId) (*model.AwsSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeAwsSecretSetQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AwsSecretSetQueryControllerServer will
// result in compilation errors.
type UnsafeAwsSecretSetQueryControllerServer interface {
	mustEmbedUnimplementedAwsSecretSetQueryControllerServer()
}

func RegisterAwsSecretSetQueryControllerServer(s grpc.ServiceRegistrar, srv AwsSecretSetQueryControllerServer) {
	s.RegisterService(&AwsSecretSetQueryController_ServiceDesc, srv)
}

func _AwsSecretSetQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AwsSecretSetId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsSecretSetQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AwsSecretSetQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsSecretSetQueryControllerServer).Get(ctx, req.(*model.AwsSecretSetId))
	}
	return interceptor(ctx, in, info, handler)
}

// AwsSecretSetQueryController_ServiceDesc is the grpc.ServiceDesc for AwsSecretSetQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AwsSecretSetQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.aws.awssecretset.service.AwsSecretSetQueryController",
	HandlerType: (*AwsSecretSetQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _AwsSecretSetQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/aws/awssecretset/service/query.proto",
}
