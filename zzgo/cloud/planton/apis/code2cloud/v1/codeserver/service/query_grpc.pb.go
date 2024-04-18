// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/codeserver/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/codeserver/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/resourcemanager/v1/product/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CodeServerQueryController_GetById_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.codeserver.service.CodeServerQueryController/getById"
	CodeServerQueryController_FindByProductId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.codeserver.service.CodeServerQueryController/findByProductId"
)

// CodeServerQueryControllerClient is the client API for CodeServerQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeServerQueryControllerClient interface {
	// look up a code server using code server id
	GetById(ctx context.Context, in *model.CodeServerId, opts ...grpc.CallOption) (*model.CodeServer, error)
	// find code servers by product id.
	// response contains only objects that the authenticated user account id has viewer access to.
	FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.CodeServers, error)
}

type codeServerQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeServerQueryControllerClient(cc grpc.ClientConnInterface) CodeServerQueryControllerClient {
	return &codeServerQueryControllerClient{cc}
}

func (c *codeServerQueryControllerClient) GetById(ctx context.Context, in *model.CodeServerId, opts ...grpc.CallOption) (*model.CodeServer, error) {
	out := new(model.CodeServer)
	err := c.cc.Invoke(ctx, CodeServerQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServerQueryControllerClient) FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.CodeServers, error) {
	out := new(model.CodeServers)
	err := c.cc.Invoke(ctx, CodeServerQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeServerQueryControllerServer is the server API for CodeServerQueryController service.
// All implementations should embed UnimplementedCodeServerQueryControllerServer
// for forward compatibility
type CodeServerQueryControllerServer interface {
	// look up a code server using code server id
	GetById(context.Context, *model.CodeServerId) (*model.CodeServer, error)
	// find code servers by product id.
	// response contains only objects that the authenticated user account id has viewer access to.
	FindByProductId(context.Context, *model1.ProductId) (*model.CodeServers, error)
}

// UnimplementedCodeServerQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCodeServerQueryControllerServer struct {
}

func (UnimplementedCodeServerQueryControllerServer) GetById(context.Context, *model.CodeServerId) (*model.CodeServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedCodeServerQueryControllerServer) FindByProductId(context.Context, *model1.ProductId) (*model.CodeServers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}

// UnsafeCodeServerQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeServerQueryControllerServer will
// result in compilation errors.
type UnsafeCodeServerQueryControllerServer interface {
	mustEmbedUnimplementedCodeServerQueryControllerServer()
}

func RegisterCodeServerQueryControllerServer(s grpc.ServiceRegistrar, srv CodeServerQueryControllerServer) {
	s.RegisterService(&CodeServerQueryController_ServiceDesc, srv)
}

func _CodeServerQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CodeServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServerQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeServerQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServerQueryControllerServer).GetById(ctx, req.(*model.CodeServerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeServerQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServerQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeServerQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServerQueryControllerServer).FindByProductId(ctx, req.(*model1.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeServerQueryController_ServiceDesc is the grpc.ServiceDesc for CodeServerQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeServerQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.codeserver.service.CodeServerQueryController",
	HandlerType: (*CodeServerQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _CodeServerQueryController_GetById_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _CodeServerQueryController_FindByProductId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/codeserver/service/query.proto",
}
