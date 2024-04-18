// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/openfgaserver/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/openfgaserver/model"
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
	OpenfgaServerCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/previewCreate"
	OpenfgaServerCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/create"
	OpenfgaServerCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/previewUpdate"
	OpenfgaServerCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/update"
	OpenfgaServerCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/previewDelete"
	OpenfgaServerCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/delete"
	OpenfgaServerCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/previewRestore"
	OpenfgaServerCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/restore"
	OpenfgaServerCommandController_Restart_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/restart"
	OpenfgaServerCommandController_Pause_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/pause"
	OpenfgaServerCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/unpause"
	OpenfgaServerCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/previewRefresh"
	OpenfgaServerCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController/refresh"
)

// OpenfgaServerCommandControllerClient is the client API for OpenfgaServerCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenfgaServerCommandControllerClient interface {
	// preview creating openfga-server
	PreviewCreate(ctx context.Context, in *model.OpenfgaServer, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// create openfga-server
	Create(ctx context.Context, in *model.OpenfgaServer, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// preview updating an existing openfga-server
	PreviewUpdate(ctx context.Context, in *model.OpenfgaServer, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// update an existing openfga-server
	Update(ctx context.Context, in *model.OpenfgaServer, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// preview deleting an existing openfga-server
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// delete an existing openfga-server
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// preview restoring a previously deleted openfga-server
	PreviewRestore(ctx context.Context, in *model.OpenfgaServer, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// restore a previously deleted openfga-server
	Restore(ctx context.Context, in *model.OpenfgaServer, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// restart a openfga-server running in a environment.
	// openfga-server is restarted by deleting running "openfga" pods which will be automatically recreated by kubernetes
	Restart(ctx context.Context, in *model.OpenfgaServerId, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// pause a openfga-server running in a environment.
	// openfga-server is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// unpause a previously paused openfga-server running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the openfga-server.
	Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// preview refresh a openfga-server that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
	// refresh a openfga-server that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.OpenfgaServer, error)
}

type openfgaServerCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenfgaServerCommandControllerClient(cc grpc.ClientConnInterface) OpenfgaServerCommandControllerClient {
	return &openfgaServerCommandControllerClient{cc}
}

func (c *openfgaServerCommandControllerClient) PreviewCreate(ctx context.Context, in *model.OpenfgaServer, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerCommandControllerClient) Create(ctx context.Context, in *model.OpenfgaServer, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.OpenfgaServer, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerCommandControllerClient) Update(ctx context.Context, in *model.OpenfgaServer, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerCommandControllerClient) PreviewRestore(ctx context.Context, in *model.OpenfgaServer, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerCommandControllerClient) Restore(ctx context.Context, in *model.OpenfgaServer, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerCommandControllerClient) Restart(ctx context.Context, in *model.OpenfgaServerId, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerCommandControllerClient) Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerCommandControllerClient) Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openfgaServerCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.OpenfgaServer, error) {
	out := new(model.OpenfgaServer)
	err := c.cc.Invoke(ctx, OpenfgaServerCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenfgaServerCommandControllerServer is the server API for OpenfgaServerCommandController service.
// All implementations should embed UnimplementedOpenfgaServerCommandControllerServer
// for forward compatibility
type OpenfgaServerCommandControllerServer interface {
	// preview creating openfga-server
	PreviewCreate(context.Context, *model.OpenfgaServer) (*model.OpenfgaServer, error)
	// create openfga-server
	Create(context.Context, *model.OpenfgaServer) (*model.OpenfgaServer, error)
	// preview updating an existing openfga-server
	PreviewUpdate(context.Context, *model.OpenfgaServer) (*model.OpenfgaServer, error)
	// update an existing openfga-server
	Update(context.Context, *model.OpenfgaServer) (*model.OpenfgaServer, error)
	// preview deleting an existing openfga-server
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.OpenfgaServer, error)
	// delete an existing openfga-server
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.OpenfgaServer, error)
	// preview restoring a previously deleted openfga-server
	PreviewRestore(context.Context, *model.OpenfgaServer) (*model.OpenfgaServer, error)
	// restore a previously deleted openfga-server
	Restore(context.Context, *model.OpenfgaServer) (*model.OpenfgaServer, error)
	// restart a openfga-server running in a environment.
	// openfga-server is restarted by deleting running "openfga" pods which will be automatically recreated by kubernetes
	Restart(context.Context, *model.OpenfgaServerId) (*model.OpenfgaServer, error)
	// pause a openfga-server running in a environment.
	// openfga-server is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.OpenfgaServer, error)
	// unpause a previously paused openfga-server running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the openfga-server.
	Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.OpenfgaServer, error)
	// preview refresh a openfga-server that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.OpenfgaServer, error)
	// refresh a openfga-server that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.OpenfgaServer, error)
}

// UnimplementedOpenfgaServerCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedOpenfgaServerCommandControllerServer struct {
}

func (UnimplementedOpenfgaServerCommandControllerServer) PreviewCreate(context.Context, *model.OpenfgaServer) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedOpenfgaServerCommandControllerServer) Create(context.Context, *model.OpenfgaServer) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOpenfgaServerCommandControllerServer) PreviewUpdate(context.Context, *model.OpenfgaServer) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedOpenfgaServerCommandControllerServer) Update(context.Context, *model.OpenfgaServer) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOpenfgaServerCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedOpenfgaServerCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedOpenfgaServerCommandControllerServer) PreviewRestore(context.Context, *model.OpenfgaServer) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedOpenfgaServerCommandControllerServer) Restore(context.Context, *model.OpenfgaServer) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedOpenfgaServerCommandControllerServer) Restart(context.Context, *model.OpenfgaServerId) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedOpenfgaServerCommandControllerServer) Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedOpenfgaServerCommandControllerServer) Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}
func (UnimplementedOpenfgaServerCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedOpenfgaServerCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.OpenfgaServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeOpenfgaServerCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenfgaServerCommandControllerServer will
// result in compilation errors.
type UnsafeOpenfgaServerCommandControllerServer interface {
	mustEmbedUnimplementedOpenfgaServerCommandControllerServer()
}

func RegisterOpenfgaServerCommandControllerServer(s grpc.ServiceRegistrar, srv OpenfgaServerCommandControllerServer) {
	s.RegisterService(&OpenfgaServerCommandController_ServiceDesc, srv)
}

func _OpenfgaServerCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenfgaServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).PreviewCreate(ctx, req.(*model.OpenfgaServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenfgaServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).Create(ctx, req.(*model.OpenfgaServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenfgaServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).PreviewUpdate(ctx, req.(*model.OpenfgaServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenfgaServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).Update(ctx, req.(*model.OpenfgaServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenfgaServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).PreviewRestore(ctx, req.(*model.OpenfgaServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenfgaServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).Restore(ctx, req.(*model.OpenfgaServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenfgaServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).Restart(ctx, req.(*model.OpenfgaServerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourcePauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).Pause(ctx, req.(*model1.ApiResourcePauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceUnPauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).Unpause(ctx, req.(*model1.ApiResourceUnPauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenfgaServerCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenfgaServerCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenfgaServerCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenfgaServerCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// OpenfgaServerCommandController_ServiceDesc is the grpc.ServiceDesc for OpenfgaServerCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenfgaServerCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenfgaServerCommandController",
	HandlerType: (*OpenfgaServerCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _OpenfgaServerCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _OpenfgaServerCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _OpenfgaServerCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _OpenfgaServerCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _OpenfgaServerCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _OpenfgaServerCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _OpenfgaServerCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _OpenfgaServerCommandController_Restore_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _OpenfgaServerCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _OpenfgaServerCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _OpenfgaServerCommandController_Unpause_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _OpenfgaServerCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _OpenfgaServerCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/openfgaserver/service/command.proto",
}
