// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/connect/v1/pulumibackendcredential/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/pulumibackendcredential/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PulumiBackendCredentialCommandController_Create_FullMethodName  = "/cloud.planton.apis.connect.v1.pulumibackendcredential.service.PulumiBackendCredentialCommandController/create"
	PulumiBackendCredentialCommandController_Update_FullMethodName  = "/cloud.planton.apis.connect.v1.pulumibackendcredential.service.PulumiBackendCredentialCommandController/update"
	PulumiBackendCredentialCommandController_Delete_FullMethodName  = "/cloud.planton.apis.connect.v1.pulumibackendcredential.service.PulumiBackendCredentialCommandController/delete"
	PulumiBackendCredentialCommandController_Restore_FullMethodName = "/cloud.planton.apis.connect.v1.pulumibackendcredential.service.PulumiBackendCredentialCommandController/restore"
)

// PulumiBackendCredentialCommandControllerClient is the client API for PulumiBackendCredentialCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PulumiBackendCredentialCommandControllerClient interface {
	// create pulumi-backend-credential
	Create(ctx context.Context, in *model.PulumiBackendCredential, opts ...grpc.CallOption) (*model.PulumiBackendCredential, error)
	// update pulumi-backend-credential
	Update(ctx context.Context, in *model.PulumiBackendCredential, opts ...grpc.CallOption) (*model.PulumiBackendCredential, error)
	// delete pulumi-backend-credential
	Delete(ctx context.Context, in *model.PulumiBackendCredentialId, opts ...grpc.CallOption) (*model.PulumiBackendCredential, error)
	// restore pulumi-backend-credential
	Restore(ctx context.Context, in *model.PulumiBackendCredential, opts ...grpc.CallOption) (*model.PulumiBackendCredential, error)
}

type pulumiBackendCredentialCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewPulumiBackendCredentialCommandControllerClient(cc grpc.ClientConnInterface) PulumiBackendCredentialCommandControllerClient {
	return &pulumiBackendCredentialCommandControllerClient{cc}
}

func (c *pulumiBackendCredentialCommandControllerClient) Create(ctx context.Context, in *model.PulumiBackendCredential, opts ...grpc.CallOption) (*model.PulumiBackendCredential, error) {
	out := new(model.PulumiBackendCredential)
	err := c.cc.Invoke(ctx, PulumiBackendCredentialCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pulumiBackendCredentialCommandControllerClient) Update(ctx context.Context, in *model.PulumiBackendCredential, opts ...grpc.CallOption) (*model.PulumiBackendCredential, error) {
	out := new(model.PulumiBackendCredential)
	err := c.cc.Invoke(ctx, PulumiBackendCredentialCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pulumiBackendCredentialCommandControllerClient) Delete(ctx context.Context, in *model.PulumiBackendCredentialId, opts ...grpc.CallOption) (*model.PulumiBackendCredential, error) {
	out := new(model.PulumiBackendCredential)
	err := c.cc.Invoke(ctx, PulumiBackendCredentialCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pulumiBackendCredentialCommandControllerClient) Restore(ctx context.Context, in *model.PulumiBackendCredential, opts ...grpc.CallOption) (*model.PulumiBackendCredential, error) {
	out := new(model.PulumiBackendCredential)
	err := c.cc.Invoke(ctx, PulumiBackendCredentialCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PulumiBackendCredentialCommandControllerServer is the server API for PulumiBackendCredentialCommandController service.
// All implementations should embed UnimplementedPulumiBackendCredentialCommandControllerServer
// for forward compatibility
type PulumiBackendCredentialCommandControllerServer interface {
	// create pulumi-backend-credential
	Create(context.Context, *model.PulumiBackendCredential) (*model.PulumiBackendCredential, error)
	// update pulumi-backend-credential
	Update(context.Context, *model.PulumiBackendCredential) (*model.PulumiBackendCredential, error)
	// delete pulumi-backend-credential
	Delete(context.Context, *model.PulumiBackendCredentialId) (*model.PulumiBackendCredential, error)
	// restore pulumi-backend-credential
	Restore(context.Context, *model.PulumiBackendCredential) (*model.PulumiBackendCredential, error)
}

// UnimplementedPulumiBackendCredentialCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedPulumiBackendCredentialCommandControllerServer struct {
}

func (UnimplementedPulumiBackendCredentialCommandControllerServer) Create(context.Context, *model.PulumiBackendCredential) (*model.PulumiBackendCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPulumiBackendCredentialCommandControllerServer) Update(context.Context, *model.PulumiBackendCredential) (*model.PulumiBackendCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPulumiBackendCredentialCommandControllerServer) Delete(context.Context, *model.PulumiBackendCredentialId) (*model.PulumiBackendCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPulumiBackendCredentialCommandControllerServer) Restore(context.Context, *model.PulumiBackendCredential) (*model.PulumiBackendCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafePulumiBackendCredentialCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PulumiBackendCredentialCommandControllerServer will
// result in compilation errors.
type UnsafePulumiBackendCredentialCommandControllerServer interface {
	mustEmbedUnimplementedPulumiBackendCredentialCommandControllerServer()
}

func RegisterPulumiBackendCredentialCommandControllerServer(s grpc.ServiceRegistrar, srv PulumiBackendCredentialCommandControllerServer) {
	s.RegisterService(&PulumiBackendCredentialCommandController_ServiceDesc, srv)
}

func _PulumiBackendCredentialCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PulumiBackendCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PulumiBackendCredentialCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PulumiBackendCredentialCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PulumiBackendCredentialCommandControllerServer).Create(ctx, req.(*model.PulumiBackendCredential))
	}
	return interceptor(ctx, in, info, handler)
}

func _PulumiBackendCredentialCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PulumiBackendCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PulumiBackendCredentialCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PulumiBackendCredentialCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PulumiBackendCredentialCommandControllerServer).Update(ctx, req.(*model.PulumiBackendCredential))
	}
	return interceptor(ctx, in, info, handler)
}

func _PulumiBackendCredentialCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PulumiBackendCredentialId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PulumiBackendCredentialCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PulumiBackendCredentialCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PulumiBackendCredentialCommandControllerServer).Delete(ctx, req.(*model.PulumiBackendCredentialId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PulumiBackendCredentialCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PulumiBackendCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PulumiBackendCredentialCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PulumiBackendCredentialCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PulumiBackendCredentialCommandControllerServer).Restore(ctx, req.(*model.PulumiBackendCredential))
	}
	return interceptor(ctx, in, info, handler)
}

// PulumiBackendCredentialCommandController_ServiceDesc is the grpc.ServiceDesc for PulumiBackendCredentialCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PulumiBackendCredentialCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.connect.v1.pulumibackendcredential.service.PulumiBackendCredentialCommandController",
	HandlerType: (*PulumiBackendCredentialCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _PulumiBackendCredentialCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _PulumiBackendCredentialCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _PulumiBackendCredentialCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _PulumiBackendCredentialCommandController_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/connect/v1/pulumibackendcredential/service/command.proto",
}