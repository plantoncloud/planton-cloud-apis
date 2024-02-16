// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/dnszone/command.proto

package dnszone

import (
	context "context"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
	job "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DnsZoneCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/previewCreate"
	DnsZoneCommandController_Create_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/create"
	DnsZoneCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/previewUpdate"
	DnsZoneCommandController_Update_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/update"
	DnsZoneCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/previewDelete"
	DnsZoneCommandController_Delete_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/delete"
	DnsZoneCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/previewRestore"
	DnsZoneCommandController_Restore_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/restore"
	DnsZoneCommandController_CreateStackJob_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/createStackJob"
	DnsZoneCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/previewRefresh"
	DnsZoneCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController/refresh"
)

// DnsZoneCommandControllerClient is the client API for DnsZoneCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DnsZoneCommandControllerClient interface {
	// preview dns-zone before creating
	PreviewCreate(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error)
	// create a dns-zone
	Create(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error)
	// preview updates to a dns-zone
	PreviewUpdate(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error)
	// update an existing dns-zone
	Update(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error)
	// preview deleting a dns-zone
	PreviewDelete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*DnsZone, error)
	// delete a dns-zone
	Delete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*DnsZone, error)
	// preview restoring a deleted dns-zone
	PreviewRestore(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error)
	// restore a deleted dns-zone
	Restore(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error)
	// create-stack-job for dns-zone
	CreateStackJob(ctx context.Context, in *job.CreateStackJobCommandInput, opts ...grpc.CallOption) (*DnsZone, error)
	// preview refresh a dns-zone that was previously created
	PreviewRefresh(ctx context.Context, in *resource.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*DnsZone, error)
	// refresh a dns-zone that was previously created
	Refresh(ctx context.Context, in *resource.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*DnsZone, error)
}

type dnsZoneCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsZoneCommandControllerClient(cc grpc.ClientConnInterface) DnsZoneCommandControllerClient {
	return &dnsZoneCommandControllerClient{cc}
}

func (c *dnsZoneCommandControllerClient) PreviewCreate(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) Create(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) PreviewUpdate(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) Update(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) PreviewDelete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) Delete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) PreviewRestore(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) Restore(ctx context.Context, in *DnsZone, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) CreateStackJob(ctx context.Context, in *job.CreateStackJobCommandInput, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_CreateStackJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) PreviewRefresh(ctx context.Context, in *resource.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneCommandControllerClient) Refresh(ctx context.Context, in *resource.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsZoneCommandControllerServer is the server API for DnsZoneCommandController service.
// All implementations should embed UnimplementedDnsZoneCommandControllerServer
// for forward compatibility
type DnsZoneCommandControllerServer interface {
	// preview dns-zone before creating
	PreviewCreate(context.Context, *DnsZone) (*DnsZone, error)
	// create a dns-zone
	Create(context.Context, *DnsZone) (*DnsZone, error)
	// preview updates to a dns-zone
	PreviewUpdate(context.Context, *DnsZone) (*DnsZone, error)
	// update an existing dns-zone
	Update(context.Context, *DnsZone) (*DnsZone, error)
	// preview deleting a dns-zone
	PreviewDelete(context.Context, *resource.ApiResourceDeleteCommandInput) (*DnsZone, error)
	// delete a dns-zone
	Delete(context.Context, *resource.ApiResourceDeleteCommandInput) (*DnsZone, error)
	// preview restoring a deleted dns-zone
	PreviewRestore(context.Context, *DnsZone) (*DnsZone, error)
	// restore a deleted dns-zone
	Restore(context.Context, *DnsZone) (*DnsZone, error)
	// create-stack-job for dns-zone
	CreateStackJob(context.Context, *job.CreateStackJobCommandInput) (*DnsZone, error)
	// preview refresh a dns-zone that was previously created
	PreviewRefresh(context.Context, *resource.ApiResourceRefreshCommandInput) (*DnsZone, error)
	// refresh a dns-zone that was previously created
	Refresh(context.Context, *resource.ApiResourceRefreshCommandInput) (*DnsZone, error)
}

// UnimplementedDnsZoneCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedDnsZoneCommandControllerServer struct {
}

func (UnimplementedDnsZoneCommandControllerServer) PreviewCreate(context.Context, *DnsZone) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) Create(context.Context, *DnsZone) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) PreviewUpdate(context.Context, *DnsZone) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) Update(context.Context, *DnsZone) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) PreviewDelete(context.Context, *resource.ApiResourceDeleteCommandInput) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) Delete(context.Context, *resource.ApiResourceDeleteCommandInput) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) PreviewRestore(context.Context, *DnsZone) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) Restore(context.Context, *DnsZone) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) CreateStackJob(context.Context, *job.CreateStackJobCommandInput) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStackJob not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) PreviewRefresh(context.Context, *resource.ApiResourceRefreshCommandInput) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedDnsZoneCommandControllerServer) Refresh(context.Context, *resource.ApiResourceRefreshCommandInput) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeDnsZoneCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DnsZoneCommandControllerServer will
// result in compilation errors.
type UnsafeDnsZoneCommandControllerServer interface {
	mustEmbedUnimplementedDnsZoneCommandControllerServer()
}

func RegisterDnsZoneCommandControllerServer(s grpc.ServiceRegistrar, srv DnsZoneCommandControllerServer) {
	s.RegisterService(&DnsZoneCommandController_ServiceDesc, srv)
}

func _DnsZoneCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).PreviewCreate(ctx, req.(*DnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).Create(ctx, req.(*DnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).PreviewUpdate(ctx, req.(*DnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).Update(ctx, req.(*DnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).PreviewDelete(ctx, req.(*resource.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).Delete(ctx, req.(*resource.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).PreviewRestore(ctx, req.(*DnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).Restore(ctx, req.(*DnsZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_CreateStackJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(job.CreateStackJobCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).CreateStackJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_CreateStackJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).CreateStackJob(ctx, req.(*job.CreateStackJobCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).PreviewRefresh(ctx, req.(*resource.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneCommandControllerServer).Refresh(ctx, req.(*resource.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// DnsZoneCommandController_ServiceDesc is the grpc.ServiceDesc for DnsZoneCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DnsZoneCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneCommandController",
	HandlerType: (*DnsZoneCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _DnsZoneCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _DnsZoneCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _DnsZoneCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _DnsZoneCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _DnsZoneCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _DnsZoneCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _DnsZoneCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _DnsZoneCommandController_Restore_Handler,
		},
		{
			MethodName: "createStackJob",
			Handler:    _DnsZoneCommandController_CreateStackJob_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _DnsZoneCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _DnsZoneCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/dnszone/command.proto",
}

const (
	DnsRecordCommandController_Add_FullMethodName    = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsRecordCommandController/add"
	DnsRecordCommandController_Update_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsRecordCommandController/update"
	DnsRecordCommandController_Delete_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsRecordCommandController/delete"
)

// DnsRecordCommandControllerClient is the client API for DnsRecordCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DnsRecordCommandControllerClient interface {
	// add a new dns-record to dns-zone
	Add(ctx context.Context, in *AddOrUpdateDnsRecordCommandInput, opts ...grpc.CallOption) (*DnsZone, error)
	// update an existing dns-record in dns-zone
	Update(ctx context.Context, in *AddOrUpdateDnsRecordCommandInput, opts ...grpc.CallOption) (*DnsZone, error)
	// delete a dns-record from a dns-zone
	Delete(ctx context.Context, in *DeleteDnsRecordCommandInput, opts ...grpc.CallOption) (*DnsZone, error)
}

type dnsRecordCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsRecordCommandControllerClient(cc grpc.ClientConnInterface) DnsRecordCommandControllerClient {
	return &dnsRecordCommandControllerClient{cc}
}

func (c *dnsRecordCommandControllerClient) Add(ctx context.Context, in *AddOrUpdateDnsRecordCommandInput, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsRecordCommandController_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsRecordCommandControllerClient) Update(ctx context.Context, in *AddOrUpdateDnsRecordCommandInput, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsRecordCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsRecordCommandControllerClient) Delete(ctx context.Context, in *DeleteDnsRecordCommandInput, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsRecordCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsRecordCommandControllerServer is the server API for DnsRecordCommandController service.
// All implementations should embed UnimplementedDnsRecordCommandControllerServer
// for forward compatibility
type DnsRecordCommandControllerServer interface {
	// add a new dns-record to dns-zone
	Add(context.Context, *AddOrUpdateDnsRecordCommandInput) (*DnsZone, error)
	// update an existing dns-record in dns-zone
	Update(context.Context, *AddOrUpdateDnsRecordCommandInput) (*DnsZone, error)
	// delete a dns-record from a dns-zone
	Delete(context.Context, *DeleteDnsRecordCommandInput) (*DnsZone, error)
}

// UnimplementedDnsRecordCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedDnsRecordCommandControllerServer struct {
}

func (UnimplementedDnsRecordCommandControllerServer) Add(context.Context, *AddOrUpdateDnsRecordCommandInput) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedDnsRecordCommandControllerServer) Update(context.Context, *AddOrUpdateDnsRecordCommandInput) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDnsRecordCommandControllerServer) Delete(context.Context, *DeleteDnsRecordCommandInput) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeDnsRecordCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DnsRecordCommandControllerServer will
// result in compilation errors.
type UnsafeDnsRecordCommandControllerServer interface {
	mustEmbedUnimplementedDnsRecordCommandControllerServer()
}

func RegisterDnsRecordCommandControllerServer(s grpc.ServiceRegistrar, srv DnsRecordCommandControllerServer) {
	s.RegisterService(&DnsRecordCommandController_ServiceDesc, srv)
}

func _DnsRecordCommandController_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateDnsRecordCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsRecordCommandControllerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsRecordCommandController_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsRecordCommandControllerServer).Add(ctx, req.(*AddOrUpdateDnsRecordCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsRecordCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateDnsRecordCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsRecordCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsRecordCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsRecordCommandControllerServer).Update(ctx, req.(*AddOrUpdateDnsRecordCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsRecordCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDnsRecordCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsRecordCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsRecordCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsRecordCommandControllerServer).Delete(ctx, req.(*DeleteDnsRecordCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// DnsRecordCommandController_ServiceDesc is the grpc.ServiceDesc for DnsRecordCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DnsRecordCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsRecordCommandController",
	HandlerType: (*DnsRecordCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _DnsRecordCommandController_Add_Handler,
		},
		{
			MethodName: "update",
			Handler:    _DnsRecordCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _DnsRecordCommandController_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/dnszone/command.proto",
}
