// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/kubernetes/mongodbkubernetes/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/mongodbkubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MongodbKubernetesQueryController_Get_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.kubernetes.mongodbkubernetes.service.MongodbKubernetesQueryController/get"
	MongodbKubernetesQueryController_GetPassword_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kubernetes.mongodbkubernetes.service.MongodbKubernetesQueryController/getPassword"
)

// MongodbKubernetesQueryControllerClient is the client API for MongodbKubernetesQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MongodbKubernetesQueryControllerClient interface {
	// look up mongodb-kubernetes using mongodb-kubernetes id
	Get(ctx context.Context, in *model.MongodbKubernetesId, opts ...grpc.CallOption) (*model.MongodbKubernetes, error)
	// look up mongodb-kubernetes sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *model.MongodbKubernetesId, opts ...grpc.CallOption) (*model.MongodbKubernetesPassword, error)
}

type mongodbKubernetesQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMongodbKubernetesQueryControllerClient(cc grpc.ClientConnInterface) MongodbKubernetesQueryControllerClient {
	return &mongodbKubernetesQueryControllerClient{cc}
}

func (c *mongodbKubernetesQueryControllerClient) Get(ctx context.Context, in *model.MongodbKubernetesId, opts ...grpc.CallOption) (*model.MongodbKubernetes, error) {
	out := new(model.MongodbKubernetes)
	err := c.cc.Invoke(ctx, MongodbKubernetesQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongodbKubernetesQueryControllerClient) GetPassword(ctx context.Context, in *model.MongodbKubernetesId, opts ...grpc.CallOption) (*model.MongodbKubernetesPassword, error) {
	out := new(model.MongodbKubernetesPassword)
	err := c.cc.Invoke(ctx, MongodbKubernetesQueryController_GetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MongodbKubernetesQueryControllerServer is the server API for MongodbKubernetesQueryController service.
// All implementations should embed UnimplementedMongodbKubernetesQueryControllerServer
// for forward compatibility
type MongodbKubernetesQueryControllerServer interface {
	// look up mongodb-kubernetes using mongodb-kubernetes id
	Get(context.Context, *model.MongodbKubernetesId) (*model.MongodbKubernetes, error)
	// look up mongodb-kubernetes sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model.MongodbKubernetesId) (*model.MongodbKubernetesPassword, error)
}

// UnimplementedMongodbKubernetesQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedMongodbKubernetesQueryControllerServer struct {
}

func (UnimplementedMongodbKubernetesQueryControllerServer) Get(context.Context, *model.MongodbKubernetesId) (*model.MongodbKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMongodbKubernetesQueryControllerServer) GetPassword(context.Context, *model.MongodbKubernetesId) (*model.MongodbKubernetesPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}

// UnsafeMongodbKubernetesQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MongodbKubernetesQueryControllerServer will
// result in compilation errors.
type UnsafeMongodbKubernetesQueryControllerServer interface {
	mustEmbedUnimplementedMongodbKubernetesQueryControllerServer()
}

func RegisterMongodbKubernetesQueryControllerServer(s grpc.ServiceRegistrar, srv MongodbKubernetesQueryControllerServer) {
	s.RegisterService(&MongodbKubernetesQueryController_ServiceDesc, srv)
}

func _MongodbKubernetesQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesQueryControllerServer).Get(ctx, req.(*model.MongodbKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongodbKubernetesQueryController_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MongodbKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongodbKubernetesQueryControllerServer).GetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MongodbKubernetesQueryController_GetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongodbKubernetesQueryControllerServer).GetPassword(ctx, req.(*model.MongodbKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

// MongodbKubernetesQueryController_ServiceDesc is the grpc.ServiceDesc for MongodbKubernetesQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MongodbKubernetesQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.kubernetes.mongodbkubernetes.service.MongodbKubernetesQueryController",
	HandlerType: (*MongodbKubernetesQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _MongodbKubernetesQueryController_Get_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _MongodbKubernetesQueryController_GetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/kubernetes/mongodbkubernetes/service/query.proto",
}
