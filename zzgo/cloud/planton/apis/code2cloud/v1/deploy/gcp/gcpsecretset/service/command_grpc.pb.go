// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/gcp/gcpsecretset/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/gcp/gcpsecretset/model"
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
	GcpSecretSetCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretset.service.GcpSecretSetCommandController/previewCreate"
	GcpSecretSetCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretset.service.GcpSecretSetCommandController/create"
	GcpSecretSetCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretset.service.GcpSecretSetCommandController/previewUpdate"
	GcpSecretSetCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretset.service.GcpSecretSetCommandController/update"
	GcpSecretSetCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretset.service.GcpSecretSetCommandController/previewDelete"
	GcpSecretSetCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretset.service.GcpSecretSetCommandController/delete"
	GcpSecretSetCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretset.service.GcpSecretSetCommandController/previewRestore"
	GcpSecretSetCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretset.service.GcpSecretSetCommandController/restore"
	GcpSecretSetCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretset.service.GcpSecretSetCommandController/previewRefresh"
	GcpSecretSetCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretset.service.GcpSecretSetCommandController/refresh"
)

// GcpSecretSetCommandControllerClient is the client API for GcpSecretSetCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GcpSecretSetCommandControllerClient interface {
	// preview create gcp-secret-set
	PreviewCreate(ctx context.Context, in *model.GcpSecretSet, opts ...grpc.CallOption) (*model.GcpSecretSet, error)
	// create gcp-secret-set
	Create(ctx context.Context, in *model.GcpSecretSet, opts ...grpc.CallOption) (*model.GcpSecretSet, error)
	// preview update an existing gcp-secret-set
	PreviewUpdate(ctx context.Context, in *model.GcpSecretSet, opts ...grpc.CallOption) (*model.GcpSecretSet, error)
	// update an existing gcp-secret-set
	Update(ctx context.Context, in *model.GcpSecretSet, opts ...grpc.CallOption) (*model.GcpSecretSet, error)
	// preview delete an existing gcp-secret-set
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.GcpSecretSet, error)
	// delete an existing gcp-secret-set
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.GcpSecretSet, error)
	// preview restore a deleted gcp-secret-set
	PreviewRestore(ctx context.Context, in *model.GcpSecretSet, opts ...grpc.CallOption) (*model.GcpSecretSet, error)
	// restore a deleted gcp-secret-set
	Restore(ctx context.Context, in *model.GcpSecretSet, opts ...grpc.CallOption) (*model.GcpSecretSet, error)
	// preview refresh a gcp-secret-set that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.GcpSecretSet, error)
	// refresh a gcp-secret-set that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.GcpSecretSet, error)
}

type gcpSecretSetCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGcpSecretSetCommandControllerClient(cc grpc.ClientConnInterface) GcpSecretSetCommandControllerClient {
	return &gcpSecretSetCommandControllerClient{cc}
}

func (c *gcpSecretSetCommandControllerClient) PreviewCreate(ctx context.Context, in *model.GcpSecretSet, opts ...grpc.CallOption) (*model.GcpSecretSet, error) {
	out := new(model.GcpSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretSetCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretSetCommandControllerClient) Create(ctx context.Context, in *model.GcpSecretSet, opts ...grpc.CallOption) (*model.GcpSecretSet, error) {
	out := new(model.GcpSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretSetCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretSetCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.GcpSecretSet, opts ...grpc.CallOption) (*model.GcpSecretSet, error) {
	out := new(model.GcpSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretSetCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretSetCommandControllerClient) Update(ctx context.Context, in *model.GcpSecretSet, opts ...grpc.CallOption) (*model.GcpSecretSet, error) {
	out := new(model.GcpSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretSetCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretSetCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.GcpSecretSet, error) {
	out := new(model.GcpSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretSetCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretSetCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.GcpSecretSet, error) {
	out := new(model.GcpSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretSetCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretSetCommandControllerClient) PreviewRestore(ctx context.Context, in *model.GcpSecretSet, opts ...grpc.CallOption) (*model.GcpSecretSet, error) {
	out := new(model.GcpSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretSetCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretSetCommandControllerClient) Restore(ctx context.Context, in *model.GcpSecretSet, opts ...grpc.CallOption) (*model.GcpSecretSet, error) {
	out := new(model.GcpSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretSetCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretSetCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.GcpSecretSet, error) {
	out := new(model.GcpSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretSetCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpSecretSetCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.GcpSecretSet, error) {
	out := new(model.GcpSecretSet)
	err := c.cc.Invoke(ctx, GcpSecretSetCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GcpSecretSetCommandControllerServer is the server API for GcpSecretSetCommandController service.
// All implementations should embed UnimplementedGcpSecretSetCommandControllerServer
// for forward compatibility
type GcpSecretSetCommandControllerServer interface {
	// preview create gcp-secret-set
	PreviewCreate(context.Context, *model.GcpSecretSet) (*model.GcpSecretSet, error)
	// create gcp-secret-set
	Create(context.Context, *model.GcpSecretSet) (*model.GcpSecretSet, error)
	// preview update an existing gcp-secret-set
	PreviewUpdate(context.Context, *model.GcpSecretSet) (*model.GcpSecretSet, error)
	// update an existing gcp-secret-set
	Update(context.Context, *model.GcpSecretSet) (*model.GcpSecretSet, error)
	// preview delete an existing gcp-secret-set
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.GcpSecretSet, error)
	// delete an existing gcp-secret-set
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.GcpSecretSet, error)
	// preview restore a deleted gcp-secret-set
	PreviewRestore(context.Context, *model.GcpSecretSet) (*model.GcpSecretSet, error)
	// restore a deleted gcp-secret-set
	Restore(context.Context, *model.GcpSecretSet) (*model.GcpSecretSet, error)
	// preview refresh a gcp-secret-set that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.GcpSecretSet, error)
	// refresh a gcp-secret-set that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.GcpSecretSet, error)
}

// UnimplementedGcpSecretSetCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGcpSecretSetCommandControllerServer struct {
}

func (UnimplementedGcpSecretSetCommandControllerServer) PreviewCreate(context.Context, *model.GcpSecretSet) (*model.GcpSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedGcpSecretSetCommandControllerServer) Create(context.Context, *model.GcpSecretSet) (*model.GcpSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGcpSecretSetCommandControllerServer) PreviewUpdate(context.Context, *model.GcpSecretSet) (*model.GcpSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedGcpSecretSetCommandControllerServer) Update(context.Context, *model.GcpSecretSet) (*model.GcpSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGcpSecretSetCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.GcpSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedGcpSecretSetCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.GcpSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGcpSecretSetCommandControllerServer) PreviewRestore(context.Context, *model.GcpSecretSet) (*model.GcpSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedGcpSecretSetCommandControllerServer) Restore(context.Context, *model.GcpSecretSet) (*model.GcpSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedGcpSecretSetCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.GcpSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedGcpSecretSetCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.GcpSecretSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeGcpSecretSetCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GcpSecretSetCommandControllerServer will
// result in compilation errors.
type UnsafeGcpSecretSetCommandControllerServer interface {
	mustEmbedUnimplementedGcpSecretSetCommandControllerServer()
}

func RegisterGcpSecretSetCommandControllerServer(s grpc.ServiceRegistrar, srv GcpSecretSetCommandControllerServer) {
	s.RegisterService(&GcpSecretSetCommandController_ServiceDesc, srv)
}

func _GcpSecretSetCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretSetCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretSetCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretSetCommandControllerServer).PreviewCreate(ctx, req.(*model.GcpSecretSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretSetCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretSetCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretSetCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretSetCommandControllerServer).Create(ctx, req.(*model.GcpSecretSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretSetCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretSetCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretSetCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretSetCommandControllerServer).PreviewUpdate(ctx, req.(*model.GcpSecretSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretSetCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretSetCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretSetCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretSetCommandControllerServer).Update(ctx, req.(*model.GcpSecretSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretSetCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretSetCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretSetCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretSetCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretSetCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretSetCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretSetCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretSetCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretSetCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretSetCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretSetCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretSetCommandControllerServer).PreviewRestore(ctx, req.(*model.GcpSecretSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretSetCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpSecretSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretSetCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretSetCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretSetCommandControllerServer).Restore(ctx, req.(*model.GcpSecretSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretSetCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretSetCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretSetCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretSetCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpSecretSetCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpSecretSetCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpSecretSetCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpSecretSetCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GcpSecretSetCommandController_ServiceDesc is the grpc.ServiceDesc for GcpSecretSetCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GcpSecretSetCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretset.service.GcpSecretSetCommandController",
	HandlerType: (*GcpSecretSetCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _GcpSecretSetCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _GcpSecretSetCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _GcpSecretSetCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _GcpSecretSetCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _GcpSecretSetCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _GcpSecretSetCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _GcpSecretSetCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _GcpSecretSetCommandController_Restore_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _GcpSecretSetCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _GcpSecretSetCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/gcp/gcpsecretset/service/command.proto",
}
