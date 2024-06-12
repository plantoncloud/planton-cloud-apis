// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/helmrelease/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/helmrelease/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	HelmReleaseQueryController_GetById_FullMethodName = "/cloud.planton.apis.code2cloud.v1.helmrelease.service.HelmReleaseQueryController/getById"
)

// HelmReleaseQueryControllerClient is the client API for HelmReleaseQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelmReleaseQueryControllerClient interface {
	// look up helm-release using helm-release id
	GetById(ctx context.Context, in *model.HelmReleaseId, opts ...grpc.CallOption) (*model.HelmRelease, error)
}

type helmReleaseQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewHelmReleaseQueryControllerClient(cc grpc.ClientConnInterface) HelmReleaseQueryControllerClient {
	return &helmReleaseQueryControllerClient{cc}
}

func (c *helmReleaseQueryControllerClient) GetById(ctx context.Context, in *model.HelmReleaseId, opts ...grpc.CallOption) (*model.HelmRelease, error) {
	out := new(model.HelmRelease)
	err := c.cc.Invoke(ctx, HelmReleaseQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelmReleaseQueryControllerServer is the server API for HelmReleaseQueryController service.
// All implementations should embed UnimplementedHelmReleaseQueryControllerServer
// for forward compatibility
type HelmReleaseQueryControllerServer interface {
	// look up helm-release using helm-release id
	GetById(context.Context, *model.HelmReleaseId) (*model.HelmRelease, error)
}

// UnimplementedHelmReleaseQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedHelmReleaseQueryControllerServer struct {
}

func (UnimplementedHelmReleaseQueryControllerServer) GetById(context.Context, *model.HelmReleaseId) (*model.HelmRelease, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}

// UnsafeHelmReleaseQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelmReleaseQueryControllerServer will
// result in compilation errors.
type UnsafeHelmReleaseQueryControllerServer interface {
	mustEmbedUnimplementedHelmReleaseQueryControllerServer()
}

func RegisterHelmReleaseQueryControllerServer(s grpc.ServiceRegistrar, srv HelmReleaseQueryControllerServer) {
	s.RegisterService(&HelmReleaseQueryController_ServiceDesc, srv)
}

func _HelmReleaseQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.HelmReleaseId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmReleaseQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelmReleaseQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmReleaseQueryControllerServer).GetById(ctx, req.(*model.HelmReleaseId))
	}
	return interceptor(ctx, in, info, handler)
}

// HelmReleaseQueryController_ServiceDesc is the grpc.ServiceDesc for HelmReleaseQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelmReleaseQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.helmrelease.service.HelmReleaseQueryController",
	HandlerType: (*HelmReleaseQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _HelmReleaseQueryController_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/helmrelease/service/query.proto",
}
