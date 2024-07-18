// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/kubernetes/istiohttpendpoint/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/istiohttpendpoint/model"
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
	IstioHttpEndpointCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointCommandController/previewCreate"
	IstioHttpEndpointCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointCommandController/create"
	IstioHttpEndpointCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointCommandController/previewUpdate"
	IstioHttpEndpointCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointCommandController/update"
	IstioHttpEndpointCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointCommandController/previewDelete"
	IstioHttpEndpointCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointCommandController/delete"
	IstioHttpEndpointCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointCommandController/previewRestore"
	IstioHttpEndpointCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointCommandController/restore"
	IstioHttpEndpointCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointCommandController/previewRefresh"
	IstioHttpEndpointCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointCommandController/refresh"
)

// IstioHttpEndpointCommandControllerClient is the client API for IstioHttpEndpointCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IstioHttpEndpointCommandControllerClient interface {
	// preview create istio-http-endpoint
	PreviewCreate(ctx context.Context, in *model.IstioHttpEndpoint, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
	// create istio-http-endpoint
	Create(ctx context.Context, in *model.IstioHttpEndpoint, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
	// preview update an existing istio-http-endpoint
	PreviewUpdate(ctx context.Context, in *model.IstioHttpEndpoint, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
	// update an existing istio-http-endpoint
	Update(ctx context.Context, in *model.IstioHttpEndpoint, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
	// preview delete istio-http-endpoint
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
	// delete istio-http-endpoint
	Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
	// preview restoring a deleted istio-http-endpoint
	PreviewRestore(ctx context.Context, in *model.IstioHttpEndpoint, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
	// restore a deleted istio-http-endpoint
	Restore(ctx context.Context, in *model.IstioHttpEndpoint, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
	// preview refresh a istio-http-endpoint that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
	// refresh a istio-http-endpoint that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
}

type istioHttpEndpointCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIstioHttpEndpointCommandControllerClient(cc grpc.ClientConnInterface) IstioHttpEndpointCommandControllerClient {
	return &istioHttpEndpointCommandControllerClient{cc}
}

func (c *istioHttpEndpointCommandControllerClient) PreviewCreate(ctx context.Context, in *model.IstioHttpEndpoint, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioHttpEndpointCommandControllerClient) Create(ctx context.Context, in *model.IstioHttpEndpoint, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioHttpEndpointCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.IstioHttpEndpoint, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioHttpEndpointCommandControllerClient) Update(ctx context.Context, in *model.IstioHttpEndpoint, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioHttpEndpointCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioHttpEndpointCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioHttpEndpointCommandControllerClient) PreviewRestore(ctx context.Context, in *model.IstioHttpEndpoint, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioHttpEndpointCommandControllerClient) Restore(ctx context.Context, in *model.IstioHttpEndpoint, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioHttpEndpointCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioHttpEndpointCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IstioHttpEndpointCommandControllerServer is the server API for IstioHttpEndpointCommandController service.
// All implementations should embed UnimplementedIstioHttpEndpointCommandControllerServer
// for forward compatibility
type IstioHttpEndpointCommandControllerServer interface {
	// preview create istio-http-endpoint
	PreviewCreate(context.Context, *model.IstioHttpEndpoint) (*model.IstioHttpEndpoint, error)
	// create istio-http-endpoint
	Create(context.Context, *model.IstioHttpEndpoint) (*model.IstioHttpEndpoint, error)
	// preview update an existing istio-http-endpoint
	PreviewUpdate(context.Context, *model.IstioHttpEndpoint) (*model.IstioHttpEndpoint, error)
	// update an existing istio-http-endpoint
	Update(context.Context, *model.IstioHttpEndpoint) (*model.IstioHttpEndpoint, error)
	// preview delete istio-http-endpoint
	PreviewDelete(context.Context, *model1.ApiResourceDeleteInput) (*model.IstioHttpEndpoint, error)
	// delete istio-http-endpoint
	Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.IstioHttpEndpoint, error)
	// preview restoring a deleted istio-http-endpoint
	PreviewRestore(context.Context, *model.IstioHttpEndpoint) (*model.IstioHttpEndpoint, error)
	// restore a deleted istio-http-endpoint
	Restore(context.Context, *model.IstioHttpEndpoint) (*model.IstioHttpEndpoint, error)
	// preview refresh a istio-http-endpoint that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshInput) (*model.IstioHttpEndpoint, error)
	// refresh a istio-http-endpoint that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshInput) (*model.IstioHttpEndpoint, error)
}

// UnimplementedIstioHttpEndpointCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedIstioHttpEndpointCommandControllerServer struct {
}

func (UnimplementedIstioHttpEndpointCommandControllerServer) PreviewCreate(context.Context, *model.IstioHttpEndpoint) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedIstioHttpEndpointCommandControllerServer) Create(context.Context, *model.IstioHttpEndpoint) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedIstioHttpEndpointCommandControllerServer) PreviewUpdate(context.Context, *model.IstioHttpEndpoint) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedIstioHttpEndpointCommandControllerServer) Update(context.Context, *model.IstioHttpEndpoint) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedIstioHttpEndpointCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteInput) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedIstioHttpEndpointCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedIstioHttpEndpointCommandControllerServer) PreviewRestore(context.Context, *model.IstioHttpEndpoint) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedIstioHttpEndpointCommandControllerServer) Restore(context.Context, *model.IstioHttpEndpoint) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedIstioHttpEndpointCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshInput) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedIstioHttpEndpointCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshInput) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeIstioHttpEndpointCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IstioHttpEndpointCommandControllerServer will
// result in compilation errors.
type UnsafeIstioHttpEndpointCommandControllerServer interface {
	mustEmbedUnimplementedIstioHttpEndpointCommandControllerServer()
}

func RegisterIstioHttpEndpointCommandControllerServer(s grpc.ServiceRegistrar, srv IstioHttpEndpointCommandControllerServer) {
	s.RegisterService(&IstioHttpEndpointCommandController_ServiceDesc, srv)
}

func _IstioHttpEndpointCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IstioHttpEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointCommandControllerServer).PreviewCreate(ctx, req.(*model.IstioHttpEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioHttpEndpointCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IstioHttpEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointCommandControllerServer).Create(ctx, req.(*model.IstioHttpEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioHttpEndpointCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IstioHttpEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointCommandControllerServer).PreviewUpdate(ctx, req.(*model.IstioHttpEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioHttpEndpointCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IstioHttpEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointCommandControllerServer).Update(ctx, req.(*model.IstioHttpEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioHttpEndpointCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioHttpEndpointCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioHttpEndpointCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IstioHttpEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointCommandControllerServer).PreviewRestore(ctx, req.(*model.IstioHttpEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioHttpEndpointCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.IstioHttpEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointCommandControllerServer).Restore(ctx, req.(*model.IstioHttpEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioHttpEndpointCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioHttpEndpointCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshInput))
	}
	return interceptor(ctx, in, info, handler)
}

// IstioHttpEndpointCommandController_ServiceDesc is the grpc.ServiceDesc for IstioHttpEndpointCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IstioHttpEndpointCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointCommandController",
	HandlerType: (*IstioHttpEndpointCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _IstioHttpEndpointCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _IstioHttpEndpointCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _IstioHttpEndpointCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _IstioHttpEndpointCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _IstioHttpEndpointCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _IstioHttpEndpointCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _IstioHttpEndpointCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _IstioHttpEndpointCommandController_Restore_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _IstioHttpEndpointCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _IstioHttpEndpointCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/kubernetes/istiohttpendpoint/service/command.proto",
}

const (
	IstioHttpEndpointRouteCommandController_Add_FullMethodName    = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointRouteCommandController/add"
	IstioHttpEndpointRouteCommandController_Update_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointRouteCommandController/update"
	IstioHttpEndpointRouteCommandController_Delete_FullMethodName = "/cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointRouteCommandController/delete"
)

// IstioHttpEndpointRouteCommandControllerClient is the client API for IstioHttpEndpointRouteCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IstioHttpEndpointRouteCommandControllerClient interface {
	// add a route to a istio-http-endpoint
	Add(ctx context.Context, in *model.AddOrUpdateIstioHttpEndpointRouteInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
	// update an existing route of a istio-http-endpoint
	Update(ctx context.Context, in *model.AddOrUpdateIstioHttpEndpointRouteInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
	// delete a route for a istio-http-endpoint.
	Delete(ctx context.Context, in *model.DeleteOrRestoreIstioHttpEndpointRouteInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error)
}

type istioHttpEndpointRouteCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIstioHttpEndpointRouteCommandControllerClient(cc grpc.ClientConnInterface) IstioHttpEndpointRouteCommandControllerClient {
	return &istioHttpEndpointRouteCommandControllerClient{cc}
}

func (c *istioHttpEndpointRouteCommandControllerClient) Add(ctx context.Context, in *model.AddOrUpdateIstioHttpEndpointRouteInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointRouteCommandController_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioHttpEndpointRouteCommandControllerClient) Update(ctx context.Context, in *model.AddOrUpdateIstioHttpEndpointRouteInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointRouteCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioHttpEndpointRouteCommandControllerClient) Delete(ctx context.Context, in *model.DeleteOrRestoreIstioHttpEndpointRouteInput, opts ...grpc.CallOption) (*model.IstioHttpEndpoint, error) {
	out := new(model.IstioHttpEndpoint)
	err := c.cc.Invoke(ctx, IstioHttpEndpointRouteCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IstioHttpEndpointRouteCommandControllerServer is the server API for IstioHttpEndpointRouteCommandController service.
// All implementations should embed UnimplementedIstioHttpEndpointRouteCommandControllerServer
// for forward compatibility
type IstioHttpEndpointRouteCommandControllerServer interface {
	// add a route to a istio-http-endpoint
	Add(context.Context, *model.AddOrUpdateIstioHttpEndpointRouteInput) (*model.IstioHttpEndpoint, error)
	// update an existing route of a istio-http-endpoint
	Update(context.Context, *model.AddOrUpdateIstioHttpEndpointRouteInput) (*model.IstioHttpEndpoint, error)
	// delete a route for a istio-http-endpoint.
	Delete(context.Context, *model.DeleteOrRestoreIstioHttpEndpointRouteInput) (*model.IstioHttpEndpoint, error)
}

// UnimplementedIstioHttpEndpointRouteCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedIstioHttpEndpointRouteCommandControllerServer struct {
}

func (UnimplementedIstioHttpEndpointRouteCommandControllerServer) Add(context.Context, *model.AddOrUpdateIstioHttpEndpointRouteInput) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedIstioHttpEndpointRouteCommandControllerServer) Update(context.Context, *model.AddOrUpdateIstioHttpEndpointRouteInput) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedIstioHttpEndpointRouteCommandControllerServer) Delete(context.Context, *model.DeleteOrRestoreIstioHttpEndpointRouteInput) (*model.IstioHttpEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeIstioHttpEndpointRouteCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IstioHttpEndpointRouteCommandControllerServer will
// result in compilation errors.
type UnsafeIstioHttpEndpointRouteCommandControllerServer interface {
	mustEmbedUnimplementedIstioHttpEndpointRouteCommandControllerServer()
}

func RegisterIstioHttpEndpointRouteCommandControllerServer(s grpc.ServiceRegistrar, srv IstioHttpEndpointRouteCommandControllerServer) {
	s.RegisterService(&IstioHttpEndpointRouteCommandController_ServiceDesc, srv)
}

func _IstioHttpEndpointRouteCommandController_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AddOrUpdateIstioHttpEndpointRouteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointRouteCommandControllerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointRouteCommandController_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointRouteCommandControllerServer).Add(ctx, req.(*model.AddOrUpdateIstioHttpEndpointRouteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioHttpEndpointRouteCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AddOrUpdateIstioHttpEndpointRouteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointRouteCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointRouteCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointRouteCommandControllerServer).Update(ctx, req.(*model.AddOrUpdateIstioHttpEndpointRouteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioHttpEndpointRouteCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.DeleteOrRestoreIstioHttpEndpointRouteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioHttpEndpointRouteCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IstioHttpEndpointRouteCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioHttpEndpointRouteCommandControllerServer).Delete(ctx, req.(*model.DeleteOrRestoreIstioHttpEndpointRouteInput))
	}
	return interceptor(ctx, in, info, handler)
}

// IstioHttpEndpointRouteCommandController_ServiceDesc is the grpc.ServiceDesc for IstioHttpEndpointRouteCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IstioHttpEndpointRouteCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.kubernetes.istiohttpendpoint.service.IstioHttpEndpointRouteCommandController",
	HandlerType: (*IstioHttpEndpointRouteCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _IstioHttpEndpointRouteCommandController_Add_Handler,
		},
		{
			MethodName: "update",
			Handler:    _IstioHttpEndpointRouteCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _IstioHttpEndpointRouteCommandController_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/kubernetes/istiohttpendpoint/service/command.proto",
}
