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
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OpenFGAServerCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/previewCreate"
	OpenFGAServerCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/create"
	OpenFGAServerCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/previewUpdate"
	OpenFGAServerCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/update"
	OpenFGAServerCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/previewDelete"
	OpenFGAServerCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/delete"
	OpenFGAServerCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/previewRestore"
	OpenFGAServerCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/restore"
	OpenFGAServerCommandController_CreateStackJob_FullMethodName = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/createStackJob"
	OpenFGAServerCommandController_Restart_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/restart"
	OpenFGAServerCommandController_Pause_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/pause"
	OpenFGAServerCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/unpause"
	OpenFGAServerCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/previewRefresh"
	OpenFGAServerCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController/refresh"
)

// OpenFGAServerCommandControllerClient is the client API for OpenFGAServerCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenFGAServerCommandControllerClient interface {
	// preview creating openfga-server
	PreviewCreate(ctx context.Context, in *model.OpenFGAServer, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// create openfga-server
	Create(ctx context.Context, in *model.OpenFGAServer, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// preview updating an existing openfga-server
	PreviewUpdate(ctx context.Context, in *model.OpenFGAServer, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// update an existing openfga-server
	Update(ctx context.Context, in *model.OpenFGAServer, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// preview deleting an existing openfga-server
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// delete an existing openfga-server
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// preview restoring a previously deleted openfga-server
	PreviewRestore(ctx context.Context, in *model.OpenFGAServer, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// restore a previously deleted openfga-server
	Restore(ctx context.Context, in *model.OpenFGAServer, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// create-stack-job for openfga-server
	CreateStackJob(ctx context.Context, in *model2.CreateStackJobCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// restart a openfga-server running in a environment.
	// openfga-server is restarted by deleting running "openfga" pods which will be automatically recreated by kubernetes
	Restart(ctx context.Context, in *model.OpenFGAServerId, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// pause a openfga-server running in a environment.
	// openfga-server is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// unpause a previously paused openfga-server running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the openfga-server.
	Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// preview refresh a openfga-server that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
	// refresh a openfga-server that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error)
}

type openFGAServerCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenFGAServerCommandControllerClient(cc grpc.ClientConnInterface) OpenFGAServerCommandControllerClient {
	return &openFGAServerCommandControllerClient{cc}
}

func (c *openFGAServerCommandControllerClient) PreviewCreate(ctx context.Context, in *model.OpenFGAServer, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) Create(ctx context.Context, in *model.OpenFGAServer, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.OpenFGAServer, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) Update(ctx context.Context, in *model.OpenFGAServer, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) PreviewRestore(ctx context.Context, in *model.OpenFGAServer, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) Restore(ctx context.Context, in *model.OpenFGAServer, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) CreateStackJob(ctx context.Context, in *model2.CreateStackJobCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_CreateStackJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) Restart(ctx context.Context, in *model.OpenFGAServerId, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openFGAServerCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.OpenFGAServer, error) {
	out := new(model.OpenFGAServer)
	err := c.cc.Invoke(ctx, OpenFGAServerCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenFGAServerCommandControllerServer is the server API for OpenFGAServerCommandController service.
// All implementations should embed UnimplementedOpenFGAServerCommandControllerServer
// for forward compatibility
type OpenFGAServerCommandControllerServer interface {
	// preview creating openfga-server
	PreviewCreate(context.Context, *model.OpenFGAServer) (*model.OpenFGAServer, error)
	// create openfga-server
	Create(context.Context, *model.OpenFGAServer) (*model.OpenFGAServer, error)
	// preview updating an existing openfga-server
	PreviewUpdate(context.Context, *model.OpenFGAServer) (*model.OpenFGAServer, error)
	// update an existing openfga-server
	Update(context.Context, *model.OpenFGAServer) (*model.OpenFGAServer, error)
	// preview deleting an existing openfga-server
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.OpenFGAServer, error)
	// delete an existing openfga-server
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.OpenFGAServer, error)
	// preview restoring a previously deleted openfga-server
	PreviewRestore(context.Context, *model.OpenFGAServer) (*model.OpenFGAServer, error)
	// restore a previously deleted openfga-server
	Restore(context.Context, *model.OpenFGAServer) (*model.OpenFGAServer, error)
	// create-stack-job for openfga-server
	CreateStackJob(context.Context, *model2.CreateStackJobCommandInput) (*model.OpenFGAServer, error)
	// restart a openfga-server running in a environment.
	// openfga-server is restarted by deleting running "openfga" pods which will be automatically recreated by kubernetes
	Restart(context.Context, *model.OpenFGAServerId) (*model.OpenFGAServer, error)
	// pause a openfga-server running in a environment.
	// openfga-server is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.OpenFGAServer, error)
	// unpause a previously paused openfga-server running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the openfga-server.
	Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.OpenFGAServer, error)
	// preview refresh a openfga-server that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.OpenFGAServer, error)
	// refresh a openfga-server that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.OpenFGAServer, error)
}

// UnimplementedOpenFGAServerCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedOpenFGAServerCommandControllerServer struct {
}

func (UnimplementedOpenFGAServerCommandControllerServer) PreviewCreate(context.Context, *model.OpenFGAServer) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) Create(context.Context, *model.OpenFGAServer) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) PreviewUpdate(context.Context, *model.OpenFGAServer) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) Update(context.Context, *model.OpenFGAServer) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) PreviewRestore(context.Context, *model.OpenFGAServer) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) Restore(context.Context, *model.OpenFGAServer) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) CreateStackJob(context.Context, *model2.CreateStackJobCommandInput) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStackJob not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) Restart(context.Context, *model.OpenFGAServerId) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedOpenFGAServerCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.OpenFGAServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeOpenFGAServerCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenFGAServerCommandControllerServer will
// result in compilation errors.
type UnsafeOpenFGAServerCommandControllerServer interface {
	mustEmbedUnimplementedOpenFGAServerCommandControllerServer()
}

func RegisterOpenFGAServerCommandControllerServer(s grpc.ServiceRegistrar, srv OpenFGAServerCommandControllerServer) {
	s.RegisterService(&OpenFGAServerCommandController_ServiceDesc, srv)
}

func _OpenFGAServerCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenFGAServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).PreviewCreate(ctx, req.(*model.OpenFGAServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenFGAServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).Create(ctx, req.(*model.OpenFGAServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenFGAServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).PreviewUpdate(ctx, req.(*model.OpenFGAServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenFGAServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).Update(ctx, req.(*model.OpenFGAServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenFGAServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).PreviewRestore(ctx, req.(*model.OpenFGAServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenFGAServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).Restore(ctx, req.(*model.OpenFGAServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_CreateStackJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.CreateStackJobCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).CreateStackJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_CreateStackJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).CreateStackJob(ctx, req.(*model2.CreateStackJobCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.OpenFGAServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).Restart(ctx, req.(*model.OpenFGAServerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourcePauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).Pause(ctx, req.(*model1.ApiResourcePauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceUnPauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).Unpause(ctx, req.(*model1.ApiResourceUnPauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenFGAServerCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenFGAServerCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenFGAServerCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenFGAServerCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// OpenFGAServerCommandController_ServiceDesc is the grpc.ServiceDesc for OpenFGAServerCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenFGAServerCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.openfgaserver.service.OpenFGAServerCommandController",
	HandlerType: (*OpenFGAServerCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _OpenFGAServerCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _OpenFGAServerCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _OpenFGAServerCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _OpenFGAServerCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _OpenFGAServerCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _OpenFGAServerCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _OpenFGAServerCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _OpenFGAServerCommandController_Restore_Handler,
		},
		{
			MethodName: "createStackJob",
			Handler:    _OpenFGAServerCommandController_CreateStackJob_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _OpenFGAServerCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _OpenFGAServerCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _OpenFGAServerCommandController_Unpause_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _OpenFGAServerCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _OpenFGAServerCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/openfgaserver/service/command.proto",
}
