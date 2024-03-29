// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/audit/v1/apiresourceversion/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/audit/v1/apiresourceversion/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ApiResourceVersionCommandController_Upsert_FullMethodName = "/cloud.planton.apis.audit.v1.apiresourceversion.service.ApiResourceVersionCommandController/upsert"
)

// ApiResourceVersionCommandControllerClient is the client API for ApiResourceVersionCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiResourceVersionCommandControllerClient interface {
	// create or update an api-resource-version resource
	Upsert(ctx context.Context, in *model.ApiResourceVersionUpsertInput, opts ...grpc.CallOption) (*model.ApiResourceVersion, error)
}

type apiResourceVersionCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewApiResourceVersionCommandControllerClient(cc grpc.ClientConnInterface) ApiResourceVersionCommandControllerClient {
	return &apiResourceVersionCommandControllerClient{cc}
}

func (c *apiResourceVersionCommandControllerClient) Upsert(ctx context.Context, in *model.ApiResourceVersionUpsertInput, opts ...grpc.CallOption) (*model.ApiResourceVersion, error) {
	out := new(model.ApiResourceVersion)
	err := c.cc.Invoke(ctx, ApiResourceVersionCommandController_Upsert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiResourceVersionCommandControllerServer is the server API for ApiResourceVersionCommandController service.
// All implementations should embed UnimplementedApiResourceVersionCommandControllerServer
// for forward compatibility
type ApiResourceVersionCommandControllerServer interface {
	// create or update an api-resource-version resource
	Upsert(context.Context, *model.ApiResourceVersionUpsertInput) (*model.ApiResourceVersion, error)
}

// UnimplementedApiResourceVersionCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedApiResourceVersionCommandControllerServer struct {
}

func (UnimplementedApiResourceVersionCommandControllerServer) Upsert(context.Context, *model.ApiResourceVersionUpsertInput) (*model.ApiResourceVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}

// UnsafeApiResourceVersionCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiResourceVersionCommandControllerServer will
// result in compilation errors.
type UnsafeApiResourceVersionCommandControllerServer interface {
	mustEmbedUnimplementedApiResourceVersionCommandControllerServer()
}

func RegisterApiResourceVersionCommandControllerServer(s grpc.ServiceRegistrar, srv ApiResourceVersionCommandControllerServer) {
	s.RegisterService(&ApiResourceVersionCommandController_ServiceDesc, srv)
}

func _ApiResourceVersionCommandController_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ApiResourceVersionUpsertInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResourceVersionCommandControllerServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiResourceVersionCommandController_Upsert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResourceVersionCommandControllerServer).Upsert(ctx, req.(*model.ApiResourceVersionUpsertInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiResourceVersionCommandController_ServiceDesc is the grpc.ServiceDesc for ApiResourceVersionCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiResourceVersionCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.audit.v1.apiresourceversion.service.ApiResourceVersionCommandController",
	HandlerType: (*ApiResourceVersionCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "upsert",
			Handler:    _ApiResourceVersionCommandController_Upsert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/audit/v1/apiresourceversion/service/command.proto",
}
