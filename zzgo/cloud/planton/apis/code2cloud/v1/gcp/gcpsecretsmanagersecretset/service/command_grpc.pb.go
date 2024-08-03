// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/gcp/gcpsecretsmanagersecretset/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpsecretsmanagersecretset/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GcpSecretsManagerSecretSetCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetCommandController/previewCreate"
	GcpSecretsManagerSecretSetCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetCommandController/create"
	GcpSecretsManagerSecretSetCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetCommandController/previewUpdate"
	GcpSecretsManagerSecretSetCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetCommandController/update"
	GcpSecretsManagerSecretSetCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetCommandController/previewDelete"
	GcpSecretsManagerSecretSetCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetCommandController/delete"
	GcpSecretsManagerSecretSetCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetCommandController/previewRestore"
	GcpSecretsManagerSecretSetCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetCommandController/restore"
	GcpSecretsManagerSecretSetCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetCommandController/previewRefresh"
	GcpSecretsManagerSecretSetCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetCommandController/refresh"
)

// GcpSecretsManagerSecretSetCommandControllerClient is the client API for GcpSecretsManagerSecretSetCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GcpSecretsManagerSecretSetCommandControllerClient interface {
	// preview create gcp-secrets-manager-secret-set
	PreviewCreate(ctx context.Context, in *model.GcpSecretsManagerSecretSet, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error)
	// create gcp-secrets-manager-secret-set
	Create(ctx context.Context, in *model.GcpSecretsManagerSecretSet, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error)
	// preview update an existing gcp-secrets-manager-secret-set
	PreviewUpdate(ctx context.Context, in *model.GcpSecretsManagerSecretSet, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error)
	// update an existing gcp-secrets-manager-secret-set
	Update(ctx context.Context, in *model.GcpSecretsManagerSecretSet, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error)
	// preview delete an existing gcp-secrets-manager-secret-set
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error)
	// delete an existing gcp-secrets-manager-secret-set
	Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error)
	// preview restore a deleted gcp-secrets-manager-secret-set
	PreviewRestore(ctx context.Context, in *model.GcpSecretsManagerSecretSet, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error)
	// restore a deleted gcp-secrets-manager-secret-set
	Restore(ctx context.Context, in *model.GcpSecretsManagerSecretSet, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error)
	// preview refresh a gcp-secrets-manager-secret-set that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error)
	// refresh a gcp-secrets-manager-secret-set that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error)
}

type gcpSecretsManagerSecretSetCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGcpSecretsManagerSecretSetCommandControllerClient(cc grpc.ClientConnInterface) GcpSecretsManagerSecretSetCommandControllerClient {
	return &gcpSecretsManagerSecretSetCommandControllerClient{cc}
}

func (c *gcpSecretsManagerSecretSetCommandControllerClient) PreviewCreate(ctx context.Context, in *model.GcpSecretsManagerSecretSet, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error) {
	out := new(model.GcpSecretsManagerSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretsManagerSecretSetCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretsManagerSecretSetCommandControllerClient) Create(ctx context.Context, in *model.GcpSecretsManagerSecretSet, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error) {
	out := new(model.GcpSecretsManagerSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretsManagerSecretSetCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretsManagerSecretSetCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.GcpSecretsManagerSecretSet, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error) {
	out := new(model.GcpSecretsManagerSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretsManagerSecretSetCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretsManagerSecretSetCommandControllerClient) Update(ctx context.Context, in *model.GcpSecretsManagerSecretSet, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error) {
	out := new(model.GcpSecretsManagerSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretsManagerSecretSetCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretsManagerSecretSetCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error) {
	out := new(model.GcpSecretsManagerSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretsManagerSecretSetCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretsManagerSecretSetCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error) {
	out := new(model.GcpSecretsManagerSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretsManagerSecretSetCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretsManagerSecretSetCommandControllerClient) PreviewRestore(ctx context.Context, in *model.GcpSecretsManagerSecretSet, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error) {
	out := new(model.GcpSecretsManagerSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretsManagerSecretSetCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretsManagerSecretSetCommandControllerClient) Restore(ctx context.Context, in *model.GcpSecretsManagerSecretSet, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error) {
	out := new(model.GcpSecretsManagerSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretsManagerSecretSetCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretsManagerSecretSetCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error) {
	out := new(model.GcpSecretsManagerSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretsManagerSecretSetCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretsManagerSecretSetCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.GcpSecretsManagerSecretSet, error) {
	out := new(model.GcpSecretsManagerSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretsManagerSecretSetCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GcpSecretsManagerSecretSetCommandControllerServer is the server API for GcpSecretsManagerSecretSetCommandController service.
// All implementations should embed UnimplementedGcpSecretsManagerSecretSetCommandControllerServer
// for forward compatibility
type GcpSecretsManagerSecretSetCommandControllerServer interface {
	// preview create gcp-secrets-manager-secret-set
	PreviewCreate(context.Context, *model.GcpSecretsManagerSecretSet) (*model.GcpSecretsManagerSecretSet, error)
	// create gcp-secrets-manager-secret-set
	Create(context.Context, *model.GcpSecretsManagerSecretSet) (*model.GcpSecretsManagerSecretSet, error)
	// preview update an existing gcp-secrets-manager-secret-set
	PreviewUpdate(context.Context, *model.GcpSecretsManagerSecretSet) (*model.GcpSecretsManagerSecretSet, error)
	// update an existing gcp-secrets-manager-secret-set
	Update(context.Context, *model.GcpSecretsManagerSecretSet) (*model.GcpSecretsManagerSecretSet, error)
	// preview delete an existing gcp-secrets-manager-secret-set
	PreviewDelete(context.Context, *model1.ApiResourceDeleteInput) (*model.GcpSecretsManagerSecretSet, error)
	// delete an existing gcp-secrets-manager-secret-set
	Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.GcpSecretsManagerSecretSet, error)
	// preview restore a deleted gcp-secrets-manager-secret-set
	PreviewRestore(context.Context, *model.GcpSecretsManagerSecretSet) (*model.GcpSecretsManagerSecretSet, error)
	// restore a deleted gcp-secrets-manager-secret-set
	Restore(context.Context, *model.GcpSecretsManagerSecretSet) (*model.GcpSecretsManagerSecretSet, error)
	// preview refresh a gcp-secrets-manager-secret-set that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshInput) (*model.GcpSecretsManagerSecretSet, error)
	// refresh a gcp-secrets-manager-secret-set that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshInput) (*model.GcpSecretsManagerSecretSet, error)
}

// UnimplementedGcpSecretsManagerSecretSetCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGcpSecretsManagerSecretSetCommandControllerServer struct {
}

func (UnimplementedGcpSecretsManagerSecretSetCommandControllerServer) PreviewCreate(context.Context, *model.GcpSecretsManagerSecretSet) (*model.GcpSecretsManagerSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedGcpSecretsManagerSecretSetCommandControllerServer) Create(context.Context, *model.GcpSecretsManagerSecretSet) (*model.GcpSecretsManagerSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGcpSecretsManagerSecretSetCommandControllerServer) PreviewUpdate(context.Context, *model.GcpSecretsManagerSecretSet) (*model.GcpSecretsManagerSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedGcpSecretsManagerSecretSetCommandControllerServer) Update(context.Context, *model.GcpSecretsManagerSecretSet) (*model.GcpSecretsManagerSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGcpSecretsManagerSecretSetCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteInput) (*model.GcpSecretsManagerSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedGcpSecretsManagerSecretSetCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.GcpSecretsManagerSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGcpSecretsManagerSecretSetCommandControllerServer) PreviewRestore(context.Context, *model.GcpSecretsManagerSecretSet) (*model.GcpSecretsManagerSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedGcpSecretsManagerSecretSetCommandControllerServer) Restore(context.Context, *model.GcpSecretsManagerSecretSet) (*model.GcpSecretsManagerSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedGcpSecretsManagerSecretSetCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshInput) (*model.GcpSecretsManagerSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedGcpSecretsManagerSecretSetCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshInput) (*model.GcpSecretsManagerSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeGcpSecretsManagerSecretSetCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GcpSecretsManagerSecretSetCommandControllerServer will
// result in compilation errors.
type UnsafeGcpSecretsManagerSecretSetCommandControllerServer interface {
	mustEmbedUnimplementedGcpSecretsManagerSecretSetCommandControllerServer()
}

func RegisterGcpSecretsManagerSecretSetCommandControllerServer(s grpc.ServiceRegistrar, srv GcpSecretsManagerSecretSetCommandControllerServer) {
	s.RegisterService(&GcpSecretsManagerSecretSetCommandController_ServiceDesc, srv)
}

func _GcpSecretsManagerSecretSetCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretsManagerSecretSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretsManagerSecretSetCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).PreviewCreate(ctx, req.(*model.GcpSecretsManagerSecretSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretsManagerSecretSetCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretsManagerSecretSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretsManagerSecretSetCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).Create(ctx, req.(*model.GcpSecretsManagerSecretSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretsManagerSecretSetCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretsManagerSecretSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretsManagerSecretSetCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).PreviewUpdate(ctx, req.(*model.GcpSecretsManagerSecretSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretsManagerSecretSetCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretsManagerSecretSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretsManagerSecretSetCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).Update(ctx, req.(*model.GcpSecretsManagerSecretSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretsManagerSecretSetCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretsManagerSecretSetCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretsManagerSecretSetCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretsManagerSecretSetCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretsManagerSecretSetCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretsManagerSecretSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretsManagerSecretSetCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).PreviewRestore(ctx, req.(*model.GcpSecretsManagerSecretSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretsManagerSecretSetCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretsManagerSecretSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretsManagerSecretSetCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).Restore(ctx, req.(*model.GcpSecretsManagerSecretSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretsManagerSecretSetCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretsManagerSecretSetCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretsManagerSecretSetCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretsManagerSecretSetCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretsManagerSecretSetCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GcpSecretsManagerSecretSetCommandController_ServiceDesc is the grpc.ServiceDesc for GcpSecretsManagerSecretSetCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GcpSecretsManagerSecretSetCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanagersecretset.service.GcpSecretsManagerSecretSetCommandController",
	HandlerType: (*GcpSecretsManagerSecretSetCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _GcpSecretsManagerSecretSetCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _GcpSecretsManagerSecretSetCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _GcpSecretsManagerSecretSetCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _GcpSecretsManagerSecretSetCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _GcpSecretsManagerSecretSetCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _GcpSecretsManagerSecretSetCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _GcpSecretsManagerSecretSetCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _GcpSecretsManagerSecretSetCommandController_Restore_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _GcpSecretsManagerSecretSetCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _GcpSecretsManagerSecretSetCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/gcp/gcpsecretsmanagersecretset/service/command.proto",
}
