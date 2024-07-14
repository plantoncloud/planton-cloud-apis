// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/openfgakubernetes/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/openfgakubernetes/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OpenfgaKubernetesQueryController_GetById_FullMethodName     = "/cloud.planton.apis.code2cloud.v1.openfgakubernetes.service.OpenfgaKubernetesQueryController/getById"
	OpenfgaKubernetesQueryController_GetPassword_FullMethodName = "/cloud.planton.apis.code2cloud.v1.openfgakubernetes.service.OpenfgaKubernetesQueryController/getPassword"
)

// OpenfgaKubernetesQueryControllerClient is the client API for OpenfgaKubernetesQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenfgaKubernetesQueryControllerClient interface {
	// look up openfga-kubernetes using openfga-kubernetes id
	GetById(ctx context.Context, in *model.OpenfgaKubernetesId, opts ...grpc.CallOption) (*model.OpenfgaKubernetes, error)
	// look up openfga-kubernetes sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(ctx context.Context, in *model.OpenfgaKubernetesId, opts ...grpc.CallOption) (*model.OpenfgaKubernetesPassword, error)
}

type openfgaKubernetesQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenfgaKubernetesQueryControllerClient(cc grpc.ClientConnInterface) OpenfgaKubernetesQueryControllerClient {
	return &openfgaKubernetesQueryControllerClient{cc}
}

func (c *openfgaKubernetesQueryControllerClient) GetById(ctx context.Context, in *model.OpenfgaKubernetesId, opts ...grpc.CallOption) (*model.OpenfgaKubernetes, error) {
	out := new(model.OpenfgaKubernetes)
	err := c.cc.Invoke(ctx, OpenfgaKubernetesQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaKubernetesQueryControllerClient) GetPassword(ctx context.Context, in *model.OpenfgaKubernetesId, opts ...grpc.CallOption) (*model.OpenfgaKubernetesPassword, error) {
	out := new(model.OpenfgaKubernetesPassword)
	err := c.cc.Invoke(ctx, OpenfgaKubernetesQueryController_GetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenfgaKubernetesQueryControllerServer is the server API for OpenfgaKubernetesQueryController service.
// All implementations should embed UnimplementedOpenfgaKubernetesQueryControllerServer
// for forward compatibility
type OpenfgaKubernetesQueryControllerServer interface {
	// look up openfga-kubernetes using openfga-kubernetes id
	GetById(context.Context, *model.OpenfgaKubernetesId) (*model.OpenfgaKubernetes, error)
	// look up openfga-kubernetes sasl password
	// password is retrieved from the kubernetes cluster.
	GetPassword(context.Context, *model.OpenfgaKubernetesId) (*model.OpenfgaKubernetesPassword, error)
}

// UnimplementedOpenfgaKubernetesQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedOpenfgaKubernetesQueryControllerServer struct {
}

func (UnimplementedOpenfgaKubernetesQueryControllerServer) GetById(context.Context, *model.OpenfgaKubernetesId) (*model.OpenfgaKubernetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedOpenfgaKubernetesQueryControllerServer) GetPassword(context.Context, *model.OpenfgaKubernetesId) (*model.OpenfgaKubernetesPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}

// UnsafeOpenfgaKubernetesQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenfgaKubernetesQueryControllerServer will
// result in compilation errors.
type UnsafeOpenfgaKubernetesQueryControllerServer interface {
	mustEmbedUnimplementedOpenfgaKubernetesQueryControllerServer()
}

func RegisterOpenfgaKubernetesQueryControllerServer(s grpc.ServiceRegistrar, srv OpenfgaKubernetesQueryControllerServer) {
	s.RegisterService(&OpenfgaKubernetesQueryController_ServiceDesc, srv)
}

func _OpenfgaKubernetesQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenfgaKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaKubernetesQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaKubernetesQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaKubernetesQueryControllerServer).GetById(ctx, req.(*model.OpenfgaKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaKubernetesQueryController_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenfgaKubernetesId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaKubernetesQueryControllerServer).GetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaKubernetesQueryController_GetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaKubernetesQueryControllerServer).GetPassword(ctx, req.(*model.OpenfgaKubernetesId))
	}
	return interceptor(ctx, in, info, handler)
}

// OpenfgaKubernetesQueryController_ServiceDesc is the grpc.ServiceDesc for OpenfgaKubernetesQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenfgaKubernetesQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.openfgakubernetes.service.OpenfgaKubernetesQueryController",
	HandlerType: (*OpenfgaKubernetesQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _OpenfgaKubernetesQueryController_GetById_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _OpenfgaKubernetesQueryController_GetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/openfgakubernetes/service/query.proto",
}