// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/aws/route53zone/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/aws/route53zone/model"
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
	Route53ZoneCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneCommandController/previewCreate"
	Route53ZoneCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneCommandController/create"
	Route53ZoneCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneCommandController/previewUpdate"
	Route53ZoneCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneCommandController/update"
	Route53ZoneCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneCommandController/previewDelete"
	Route53ZoneCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneCommandController/delete"
	Route53ZoneCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneCommandController/previewRestore"
	Route53ZoneCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneCommandController/restore"
	Route53ZoneCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneCommandController/previewRefresh"
	Route53ZoneCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneCommandController/refresh"
)

// Route53ZoneCommandControllerClient is the client API for Route53ZoneCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Route53ZoneCommandControllerClient interface {
	// preview route53-zone before creating
	PreviewCreate(ctx context.Context, in *model.Route53Zone, opts ...grpc.CallOption) (*model.Route53Zone, error)
	// create a route53-zone
	Create(ctx context.Context, in *model.Route53Zone, opts ...grpc.CallOption) (*model.Route53Zone, error)
	// preview updates to a route53-zone
	PreviewUpdate(ctx context.Context, in *model.Route53Zone, opts ...grpc.CallOption) (*model.Route53Zone, error)
	// update an existing route53-zone
	Update(ctx context.Context, in *model.Route53Zone, opts ...grpc.CallOption) (*model.Route53Zone, error)
	// preview deleting a route53-zone
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error)
	// delete a route53-zone
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error)
	// preview restoring a deleted route53-zone
	PreviewRestore(ctx context.Context, in *model.Route53Zone, opts ...grpc.CallOption) (*model.Route53Zone, error)
	// restore a deleted route53-zone
	Restore(ctx context.Context, in *model.Route53Zone, opts ...grpc.CallOption) (*model.Route53Zone, error)
	// preview refresh a route53-zone that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error)
	// refresh a route53-zone that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error)
}

type route53ZoneCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewRoute53ZoneCommandControllerClient(cc grpc.ClientConnInterface) Route53ZoneCommandControllerClient {
	return &route53ZoneCommandControllerClient{cc}
}

func (c *route53ZoneCommandControllerClient) PreviewCreate(ctx context.Context, in *model.Route53Zone, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *route53ZoneCommandControllerClient) Create(ctx context.Context, in *model.Route53Zone, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *route53ZoneCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.Route53Zone, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *route53ZoneCommandControllerClient) Update(ctx context.Context, in *model.Route53Zone, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *route53ZoneCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *route53ZoneCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *route53ZoneCommandControllerClient) PreviewRestore(ctx context.Context, in *model.Route53Zone, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *route53ZoneCommandControllerClient) Restore(ctx context.Context, in *model.Route53Zone, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *route53ZoneCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *route53ZoneCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Route53ZoneCommandControllerServer is the server API for Route53ZoneCommandController service.
// All implementations should embed UnimplementedRoute53ZoneCommandControllerServer
// for forward compatibility
type Route53ZoneCommandControllerServer interface {
	// preview route53-zone before creating
	PreviewCreate(context.Context, *model.Route53Zone) (*model.Route53Zone, error)
	// create a route53-zone
	Create(context.Context, *model.Route53Zone) (*model.Route53Zone, error)
	// preview updates to a route53-zone
	PreviewUpdate(context.Context, *model.Route53Zone) (*model.Route53Zone, error)
	// update an existing route53-zone
	Update(context.Context, *model.Route53Zone) (*model.Route53Zone, error)
	// preview deleting a route53-zone
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.Route53Zone, error)
	// delete a route53-zone
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.Route53Zone, error)
	// preview restoring a deleted route53-zone
	PreviewRestore(context.Context, *model.Route53Zone) (*model.Route53Zone, error)
	// restore a deleted route53-zone
	Restore(context.Context, *model.Route53Zone) (*model.Route53Zone, error)
	// preview refresh a route53-zone that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.Route53Zone, error)
	// refresh a route53-zone that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.Route53Zone, error)
}

// UnimplementedRoute53ZoneCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedRoute53ZoneCommandControllerServer struct {
}

func (UnimplementedRoute53ZoneCommandControllerServer) PreviewCreate(context.Context, *model.Route53Zone) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedRoute53ZoneCommandControllerServer) Create(context.Context, *model.Route53Zone) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRoute53ZoneCommandControllerServer) PreviewUpdate(context.Context, *model.Route53Zone) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedRoute53ZoneCommandControllerServer) Update(context.Context, *model.Route53Zone) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRoute53ZoneCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedRoute53ZoneCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRoute53ZoneCommandControllerServer) PreviewRestore(context.Context, *model.Route53Zone) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedRoute53ZoneCommandControllerServer) Restore(context.Context, *model.Route53Zone) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedRoute53ZoneCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedRoute53ZoneCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeRoute53ZoneCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Route53ZoneCommandControllerServer will
// result in compilation errors.
type UnsafeRoute53ZoneCommandControllerServer interface {
	mustEmbedUnimplementedRoute53ZoneCommandControllerServer()
}

func RegisterRoute53ZoneCommandControllerServer(s grpc.ServiceRegistrar, srv Route53ZoneCommandControllerServer) {
	s.RegisterService(&Route53ZoneCommandController_ServiceDesc, srv)
}

func _Route53ZoneCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Route53Zone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneCommandControllerServer).PreviewCreate(ctx, req.(*model.Route53Zone))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route53ZoneCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Route53Zone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneCommandControllerServer).Create(ctx, req.(*model.Route53Zone))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route53ZoneCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Route53Zone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneCommandControllerServer).PreviewUpdate(ctx, req.(*model.Route53Zone))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route53ZoneCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Route53Zone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneCommandControllerServer).Update(ctx, req.(*model.Route53Zone))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route53ZoneCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route53ZoneCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route53ZoneCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Route53Zone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneCommandControllerServer).PreviewRestore(ctx, req.(*model.Route53Zone))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route53ZoneCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Route53Zone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneCommandControllerServer).Restore(ctx, req.(*model.Route53Zone))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route53ZoneCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route53ZoneCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Route53ZoneCommandController_ServiceDesc is the grpc.ServiceDesc for Route53ZoneCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Route53ZoneCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneCommandController",
	HandlerType: (*Route53ZoneCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _Route53ZoneCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _Route53ZoneCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _Route53ZoneCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Route53ZoneCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _Route53ZoneCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Route53ZoneCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _Route53ZoneCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _Route53ZoneCommandController_Restore_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _Route53ZoneCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _Route53ZoneCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/aws/route53zone/service/command.proto",
}

const (
	Route53ZoneRecordCommandController_Add_FullMethodName    = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneRecordCommandController/add"
	Route53ZoneRecordCommandController_Update_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneRecordCommandController/update"
	Route53ZoneRecordCommandController_Delete_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneRecordCommandController/delete"
)

// Route53ZoneRecordCommandControllerClient is the client API for Route53ZoneRecordCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Route53ZoneRecordCommandControllerClient interface {
	// add a new dns-record to route53-zone
	Add(ctx context.Context, in *model.AddOrUpdateDnsRecordCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error)
	// update an existing dns-record in route53-zone
	Update(ctx context.Context, in *model.AddOrUpdateDnsRecordCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error)
	// delete a dns-record from a route53-zone
	Delete(ctx context.Context, in *model.DeleteDnsRecordCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error)
}

type route53ZoneRecordCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewRoute53ZoneRecordCommandControllerClient(cc grpc.ClientConnInterface) Route53ZoneRecordCommandControllerClient {
	return &route53ZoneRecordCommandControllerClient{cc}
}

func (c *route53ZoneRecordCommandControllerClient) Add(ctx context.Context, in *model.AddOrUpdateDnsRecordCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneRecordCommandController_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *route53ZoneRecordCommandControllerClient) Update(ctx context.Context, in *model.AddOrUpdateDnsRecordCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneRecordCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *route53ZoneRecordCommandControllerClient) Delete(ctx context.Context, in *model.DeleteDnsRecordCommandInput, opts ...grpc.CallOption) (*model.Route53Zone, error) {
	out := new(model.Route53Zone)
	err := c.cc.Invoke(ctx, Route53ZoneRecordCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Route53ZoneRecordCommandControllerServer is the server API for Route53ZoneRecordCommandController service.
// All implementations should embed UnimplementedRoute53ZoneRecordCommandControllerServer
// for forward compatibility
type Route53ZoneRecordCommandControllerServer interface {
	// add a new dns-record to route53-zone
	Add(context.Context, *model.AddOrUpdateDnsRecordCommandInput) (*model.Route53Zone, error)
	// update an existing dns-record in route53-zone
	Update(context.Context, *model.AddOrUpdateDnsRecordCommandInput) (*model.Route53Zone, error)
	// delete a dns-record from a route53-zone
	Delete(context.Context, *model.DeleteDnsRecordCommandInput) (*model.Route53Zone, error)
}

// UnimplementedRoute53ZoneRecordCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedRoute53ZoneRecordCommandControllerServer struct {
}

func (UnimplementedRoute53ZoneRecordCommandControllerServer) Add(context.Context, *model.AddOrUpdateDnsRecordCommandInput) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedRoute53ZoneRecordCommandControllerServer) Update(context.Context, *model.AddOrUpdateDnsRecordCommandInput) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRoute53ZoneRecordCommandControllerServer) Delete(context.Context, *model.DeleteDnsRecordCommandInput) (*model.Route53Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeRoute53ZoneRecordCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Route53ZoneRecordCommandControllerServer will
// result in compilation errors.
type UnsafeRoute53ZoneRecordCommandControllerServer interface {
	mustEmbedUnimplementedRoute53ZoneRecordCommandControllerServer()
}

func RegisterRoute53ZoneRecordCommandControllerServer(s grpc.ServiceRegistrar, srv Route53ZoneRecordCommandControllerServer) {
	s.RegisterService(&Route53ZoneRecordCommandController_ServiceDesc, srv)
}

func _Route53ZoneRecordCommandController_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AddOrUpdateDnsRecordCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneRecordCommandControllerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneRecordCommandController_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneRecordCommandControllerServer).Add(ctx, req.(*model.AddOrUpdateDnsRecordCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route53ZoneRecordCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.AddOrUpdateDnsRecordCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneRecordCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneRecordCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneRecordCommandControllerServer).Update(ctx, req.(*model.AddOrUpdateDnsRecordCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route53ZoneRecordCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.DeleteDnsRecordCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Route53ZoneRecordCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route53ZoneRecordCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Route53ZoneRecordCommandControllerServer).Delete(ctx, req.(*model.DeleteDnsRecordCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Route53ZoneRecordCommandController_ServiceDesc is the grpc.ServiceDesc for Route53ZoneRecordCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Route53ZoneRecordCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.service.Route53ZoneRecordCommandController",
	HandlerType: (*Route53ZoneRecordCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _Route53ZoneRecordCommandController_Add_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Route53ZoneRecordCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Route53ZoneRecordCommandController_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/aws/route53zone/service/command.proto",
}