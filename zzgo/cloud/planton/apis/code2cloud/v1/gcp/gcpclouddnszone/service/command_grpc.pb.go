// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/gcp/gcpclouddnszone/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpclouddnszone/model"
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
	GcpCloudDnsZoneCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsZoneCommandController/previewCreate"
	GcpCloudDnsZoneCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsZoneCommandController/create"
	GcpCloudDnsZoneCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsZoneCommandController/previewUpdate"
	GcpCloudDnsZoneCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsZoneCommandController/update"
	GcpCloudDnsZoneCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsZoneCommandController/previewDelete"
	GcpCloudDnsZoneCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsZoneCommandController/delete"
	GcpCloudDnsZoneCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsZoneCommandController/previewRestore"
	GcpCloudDnsZoneCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsZoneCommandController/restore"
	GcpCloudDnsZoneCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsZoneCommandController/previewRefresh"
	GcpCloudDnsZoneCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsZoneCommandController/refresh"
)

// GcpCloudDnsZoneCommandControllerClient is the client API for GcpCloudDnsZoneCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GcpCloudDnsZoneCommandControllerClient interface {
	// preview gcp-cloud-dns-zone before creating
	PreviewCreate(ctx context.Context, in *model.GcpCloudDnsZone, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
	// create a gcp-cloud-dns-zone
	Create(ctx context.Context, in *model.GcpCloudDnsZone, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
	// preview updates to a gcp-cloud-dns-zone
	PreviewUpdate(ctx context.Context, in *model.GcpCloudDnsZone, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
	// update an existing gcp-cloud-dns-zone
	Update(ctx context.Context, in *model.GcpCloudDnsZone, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
	// preview deleting a gcp-cloud-dns-zone
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
	// delete a gcp-cloud-dns-zone
	Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
	// preview restoring a deleted gcp-cloud-dns-zone
	PreviewRestore(ctx context.Context, in *model.GcpCloudDnsZone, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
	// restore a deleted gcp-cloud-dns-zone
	Restore(ctx context.Context, in *model.GcpCloudDnsZone, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
	// preview refresh a gcp-cloud-dns-zone that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
	// refresh a gcp-cloud-dns-zone that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
}

type gcpCloudDnsZoneCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGcpCloudDnsZoneCommandControllerClient(cc grpc.ClientConnInterface) GcpCloudDnsZoneCommandControllerClient {
	return &gcpCloudDnsZoneCommandControllerClient{cc}
}

func (c *gcpCloudDnsZoneCommandControllerClient) PreviewCreate(ctx context.Context, in *model.GcpCloudDnsZone, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsZoneCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpCloudDnsZoneCommandControllerClient) Create(ctx context.Context, in *model.GcpCloudDnsZone, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsZoneCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpCloudDnsZoneCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.GcpCloudDnsZone, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsZoneCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpCloudDnsZoneCommandControllerClient) Update(ctx context.Context, in *model.GcpCloudDnsZone, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsZoneCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpCloudDnsZoneCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsZoneCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpCloudDnsZoneCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsZoneCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpCloudDnsZoneCommandControllerClient) PreviewRestore(ctx context.Context, in *model.GcpCloudDnsZone, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsZoneCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpCloudDnsZoneCommandControllerClient) Restore(ctx context.Context, in *model.GcpCloudDnsZone, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsZoneCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpCloudDnsZoneCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsZoneCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpCloudDnsZoneCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsZoneCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GcpCloudDnsZoneCommandControllerServer is the server API for GcpCloudDnsZoneCommandController service.
// All implementations should embed UnimplementedGcpCloudDnsZoneCommandControllerServer
// for forward compatibility
type GcpCloudDnsZoneCommandControllerServer interface {
	// preview gcp-cloud-dns-zone before creating
	PreviewCreate(context.Context, *model.GcpCloudDnsZone) (*model.GcpCloudDnsZone, error)
	// create a gcp-cloud-dns-zone
	Create(context.Context, *model.GcpCloudDnsZone) (*model.GcpCloudDnsZone, error)
	// preview updates to a gcp-cloud-dns-zone
	PreviewUpdate(context.Context, *model.GcpCloudDnsZone) (*model.GcpCloudDnsZone, error)
	// update an existing gcp-cloud-dns-zone
	Update(context.Context, *model.GcpCloudDnsZone) (*model.GcpCloudDnsZone, error)
	// preview deleting a gcp-cloud-dns-zone
	PreviewDelete(context.Context, *model1.ApiResourceDeleteInput) (*model.GcpCloudDnsZone, error)
	// delete a gcp-cloud-dns-zone
	Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.GcpCloudDnsZone, error)
	// preview restoring a deleted gcp-cloud-dns-zone
	PreviewRestore(context.Context, *model.GcpCloudDnsZone) (*model.GcpCloudDnsZone, error)
	// restore a deleted gcp-cloud-dns-zone
	Restore(context.Context, *model.GcpCloudDnsZone) (*model.GcpCloudDnsZone, error)
	// preview refresh a gcp-cloud-dns-zone that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshInput) (*model.GcpCloudDnsZone, error)
	// refresh a gcp-cloud-dns-zone that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshInput) (*model.GcpCloudDnsZone, error)
}

// UnimplementedGcpCloudDnsZoneCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGcpCloudDnsZoneCommandControllerServer struct {
}

func (UnimplementedGcpCloudDnsZoneCommandControllerServer) PreviewCreate(context.Context, *model.GcpCloudDnsZone) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedGcpCloudDnsZoneCommandControllerServer) Create(context.Context, *model.GcpCloudDnsZone) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGcpCloudDnsZoneCommandControllerServer) PreviewUpdate(context.Context, *model.GcpCloudDnsZone) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedGcpCloudDnsZoneCommandControllerServer) Update(context.Context, *model.GcpCloudDnsZone) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGcpCloudDnsZoneCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteInput) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedGcpCloudDnsZoneCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGcpCloudDnsZoneCommandControllerServer) PreviewRestore(context.Context, *model.GcpCloudDnsZone) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedGcpCloudDnsZoneCommandControllerServer) Restore(context.Context, *model.GcpCloudDnsZone) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedGcpCloudDnsZoneCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshInput) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedGcpCloudDnsZoneCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshInput) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeGcpCloudDnsZoneCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GcpCloudDnsZoneCommandControllerServer will
// result in compilation errors.
type UnsafeGcpCloudDnsZoneCommandControllerServer interface {
	mustEmbedUnimplementedGcpCloudDnsZoneCommandControllerServer()
}

func RegisterGcpCloudDnsZoneCommandControllerServer(s grpc.ServiceRegistrar, srv GcpCloudDnsZoneCommandControllerServer) {
	s.RegisterService(&GcpCloudDnsZoneCommandController_ServiceDesc, srv)
}

func _GcpCloudDnsZoneCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpCloudDnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsZoneCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsZoneCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsZoneCommandControllerServer).PreviewCreate(ctx, req.(*model.GcpCloudDnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpCloudDnsZoneCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpCloudDnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsZoneCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsZoneCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsZoneCommandControllerServer).Create(ctx, req.(*model.GcpCloudDnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpCloudDnsZoneCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpCloudDnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsZoneCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsZoneCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsZoneCommandControllerServer).PreviewUpdate(ctx, req.(*model.GcpCloudDnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpCloudDnsZoneCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpCloudDnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsZoneCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsZoneCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsZoneCommandControllerServer).Update(ctx, req.(*model.GcpCloudDnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpCloudDnsZoneCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsZoneCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsZoneCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsZoneCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpCloudDnsZoneCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsZoneCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsZoneCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsZoneCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpCloudDnsZoneCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpCloudDnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsZoneCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsZoneCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsZoneCommandControllerServer).PreviewRestore(ctx, req.(*model.GcpCloudDnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpCloudDnsZoneCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpCloudDnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsZoneCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsZoneCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsZoneCommandControllerServer).Restore(ctx, req.(*model.GcpCloudDnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpCloudDnsZoneCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsZoneCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsZoneCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsZoneCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpCloudDnsZoneCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsZoneCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsZoneCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsZoneCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GcpCloudDnsZoneCommandController_ServiceDesc is the grpc.ServiceDesc for GcpCloudDnsZoneCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GcpCloudDnsZoneCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsZoneCommandController",
	HandlerType: (*GcpCloudDnsZoneCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _GcpCloudDnsZoneCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _GcpCloudDnsZoneCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _GcpCloudDnsZoneCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _GcpCloudDnsZoneCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _GcpCloudDnsZoneCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _GcpCloudDnsZoneCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _GcpCloudDnsZoneCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _GcpCloudDnsZoneCommandController_Restore_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _GcpCloudDnsZoneCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _GcpCloudDnsZoneCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/gcp/gcpclouddnszone/service/command.proto",
}

const (
	GcpCloudDnsRecordCommandController_Add_FullMethodName    = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsRecordCommandController/add"
	GcpCloudDnsRecordCommandController_Update_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsRecordCommandController/update"
	GcpCloudDnsRecordCommandController_Delete_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsRecordCommandController/delete"
)

// GcpCloudDnsRecordCommandControllerClient is the client API for GcpCloudDnsRecordCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GcpCloudDnsRecordCommandControllerClient interface {
	// add a new dns-record to gcp-cloud-dns-zone
	Add(ctx context.Context, in *model.AddOrUpdateGcpCloudDnsRecordInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
	// update an existing dns-record in gcp-cloud-dns-zone
	Update(ctx context.Context, in *model.AddOrUpdateGcpCloudDnsRecordInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
	// delete a dns-record from a gcp-cloud-dns-zone
	Delete(ctx context.Context, in *model.DeleteGcpCloudDnsRecordInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error)
}

type gcpCloudDnsRecordCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGcpCloudDnsRecordCommandControllerClient(cc grpc.ClientConnInterface) GcpCloudDnsRecordCommandControllerClient {
	return &gcpCloudDnsRecordCommandControllerClient{cc}
}

func (c *gcpCloudDnsRecordCommandControllerClient) Add(ctx context.Context, in *model.AddOrUpdateGcpCloudDnsRecordInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsRecordCommandController_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpCloudDnsRecordCommandControllerClient) Update(ctx context.Context, in *model.AddOrUpdateGcpCloudDnsRecordInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsRecordCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpCloudDnsRecordCommandControllerClient) Delete(ctx context.Context, in *model.DeleteGcpCloudDnsRecordInput, opts ...grpc.CallOption) (*model.GcpCloudDnsZone, error) {
	out := new(model.GcpCloudDnsZone)
	err := c.cc.Invoke(ctx, GcpCloudDnsRecordCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GcpCloudDnsRecordCommandControllerServer is the server API for GcpCloudDnsRecordCommandController service.
// All implementations should embed UnimplementedGcpCloudDnsRecordCommandControllerServer
// for forward compatibility
type GcpCloudDnsRecordCommandControllerServer interface {
	// add a new dns-record to gcp-cloud-dns-zone
	Add(context.Context, *model.AddOrUpdateGcpCloudDnsRecordInput) (*model.GcpCloudDnsZone, error)
	// update an existing dns-record in gcp-cloud-dns-zone
	Update(context.Context, *model.AddOrUpdateGcpCloudDnsRecordInput) (*model.GcpCloudDnsZone, error)
	// delete a dns-record from a gcp-cloud-dns-zone
	Delete(context.Context, *model.DeleteGcpCloudDnsRecordInput) (*model.GcpCloudDnsZone, error)
}

// UnimplementedGcpCloudDnsRecordCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGcpCloudDnsRecordCommandControllerServer struct {
}

func (UnimplementedGcpCloudDnsRecordCommandControllerServer) Add(context.Context, *model.AddOrUpdateGcpCloudDnsRecordInput) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedGcpCloudDnsRecordCommandControllerServer) Update(context.Context, *model.AddOrUpdateGcpCloudDnsRecordInput) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGcpCloudDnsRecordCommandControllerServer) Delete(context.Context, *model.DeleteGcpCloudDnsRecordInput) (*model.GcpCloudDnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeGcpCloudDnsRecordCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GcpCloudDnsRecordCommandControllerServer will
// result in compilation errors.
type UnsafeGcpCloudDnsRecordCommandControllerServer interface {
	mustEmbedUnimplementedGcpCloudDnsRecordCommandControllerServer()
}

func RegisterGcpCloudDnsRecordCommandControllerServer(s grpc.ServiceRegistrar, srv GcpCloudDnsRecordCommandControllerServer) {
	s.RegisterService(&GcpCloudDnsRecordCommandController_ServiceDesc, srv)
}

func _GcpCloudDnsRecordCommandController_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AddOrUpdateGcpCloudDnsRecordInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsRecordCommandControllerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsRecordCommandController_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsRecordCommandControllerServer).Add(ctx, req.(*model.AddOrUpdateGcpCloudDnsRecordInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpCloudDnsRecordCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AddOrUpdateGcpCloudDnsRecordInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsRecordCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsRecordCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsRecordCommandControllerServer).Update(ctx, req.(*model.AddOrUpdateGcpCloudDnsRecordInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpCloudDnsRecordCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.DeleteGcpCloudDnsRecordInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpCloudDnsRecordCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpCloudDnsRecordCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpCloudDnsRecordCommandControllerServer).Delete(ctx, req.(*model.DeleteGcpCloudDnsRecordInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GcpCloudDnsRecordCommandController_ServiceDesc is the grpc.ServiceDesc for GcpCloudDnsRecordCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GcpCloudDnsRecordCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.service.GcpCloudDnsRecordCommandController",
	HandlerType: (*GcpCloudDnsRecordCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _GcpCloudDnsRecordCommandController_Add_Handler,
		},
		{
			MethodName: "update",
			Handler:    _GcpCloudDnsRecordCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _GcpCloudDnsRecordCommandController_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/gcp/gcpclouddnszone/service/command.proto",
}