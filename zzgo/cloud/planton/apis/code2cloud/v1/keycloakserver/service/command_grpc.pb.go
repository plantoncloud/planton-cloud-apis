// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/keycloakserver/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/keycloakserver/model"
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
	KeycloakServerCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/previewCreate"
	KeycloakServerCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/create"
	KeycloakServerCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/previewUpdate"
	KeycloakServerCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/update"
	KeycloakServerCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/previewDelete"
	KeycloakServerCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/delete"
	KeycloakServerCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/previewRestore"
	KeycloakServerCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/restore"
	KeycloakServerCommandController_Restart_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/restart"
	KeycloakServerCommandController_Pause_FullMethodName          = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/pause"
	KeycloakServerCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/unpause"
	KeycloakServerCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/previewRefresh"
	KeycloakServerCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController/refresh"
)

// KeycloakServerCommandControllerClient is the client API for KeycloakServerCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeycloakServerCommandControllerClient interface {
	// preview creating keycloak-server
	PreviewCreate(ctx context.Context, in *model.KeycloakServer, opts ...grpc.CallOption) (*model.KeycloakServer, error)
	// create keycloak-server
	Create(ctx context.Context, in *model.KeycloakServer, opts ...grpc.CallOption) (*model.KeycloakServer, error)
	// preview updating an existing keycloak-server
	PreviewUpdate(ctx context.Context, in *model.KeycloakServer, opts ...grpc.CallOption) (*model.KeycloakServer, error)
	// update an existing keycloak-server
	Update(ctx context.Context, in *model.KeycloakServer, opts ...grpc.CallOption) (*model.KeycloakServer, error)
	// preview deleting an existing keycloak-server
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.KeycloakServer, error)
	// delete an existing keycloak-server
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.KeycloakServer, error)
	// preview restoring a previously deleted keycloak-server
	PreviewRestore(ctx context.Context, in *model.KeycloakServer, opts ...grpc.CallOption) (*model.KeycloakServer, error)
	// restore a previously deleted keycloak-server
	Restore(ctx context.Context, in *model.KeycloakServer, opts ...grpc.CallOption) (*model.KeycloakServer, error)
	// restart a keycloak-server running in a environment.
	// keycloak-server is restarted by deleting running "keycloak" pods which will be automatically recreated by kubernetes
	Restart(ctx context.Context, in *model.KeycloakServerId, opts ...grpc.CallOption) (*model.KeycloakServer, error)
	// pause a keycloak-server running in a environment.
	// keycloak-server is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.KeycloakServer, error)
	// unpause a previously paused keycloak-server running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the keycloak-server.
	Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.KeycloakServer, error)
	// preview refresh a keycloak-server that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.KeycloakServer, error)
	// refresh a keycloak-server that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.KeycloakServer, error)
}

type keycloakServerCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKeycloakServerCommandControllerClient(cc grpc.ClientConnInterface) KeycloakServerCommandControllerClient {
	return &keycloakServerCommandControllerClient{cc}
}

func (c *keycloakServerCommandControllerClient) PreviewCreate(ctx context.Context, in *model.KeycloakServer, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycloakServerCommandControllerClient) Create(ctx context.Context, in *model.KeycloakServer, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycloakServerCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.KeycloakServer, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycloakServerCommandControllerClient) Update(ctx context.Context, in *model.KeycloakServer, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycloakServerCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycloakServerCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycloakServerCommandControllerClient) PreviewRestore(ctx context.Context, in *model.KeycloakServer, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycloakServerCommandControllerClient) Restore(ctx context.Context, in *model.KeycloakServer, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycloakServerCommandControllerClient) Restart(ctx context.Context, in *model.KeycloakServerId, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycloakServerCommandControllerClient) Pause(ctx context.Context, in *model1.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycloakServerCommandControllerClient) Unpause(ctx context.Context, in *model1.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycloakServerCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycloakServerCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.KeycloakServer, error) {
	out := new(model.KeycloakServer)
	err := c.cc.Invoke(ctx, KeycloakServerCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeycloakServerCommandControllerServer is the server API for KeycloakServerCommandController service.
// All implementations should embed UnimplementedKeycloakServerCommandControllerServer
// for forward compatibility
type KeycloakServerCommandControllerServer interface {
	// preview creating keycloak-server
	PreviewCreate(context.Context, *model.KeycloakServer) (*model.KeycloakServer, error)
	// create keycloak-server
	Create(context.Context, *model.KeycloakServer) (*model.KeycloakServer, error)
	// preview updating an existing keycloak-server
	PreviewUpdate(context.Context, *model.KeycloakServer) (*model.KeycloakServer, error)
	// update an existing keycloak-server
	Update(context.Context, *model.KeycloakServer) (*model.KeycloakServer, error)
	// preview deleting an existing keycloak-server
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.KeycloakServer, error)
	// delete an existing keycloak-server
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.KeycloakServer, error)
	// preview restoring a previously deleted keycloak-server
	PreviewRestore(context.Context, *model.KeycloakServer) (*model.KeycloakServer, error)
	// restore a previously deleted keycloak-server
	Restore(context.Context, *model.KeycloakServer) (*model.KeycloakServer, error)
	// restart a keycloak-server running in a environment.
	// keycloak-server is restarted by deleting running "keycloak" pods which will be automatically recreated by kubernetes
	Restart(context.Context, *model.KeycloakServerId) (*model.KeycloakServer, error)
	// pause a keycloak-server running in a environment.
	// keycloak-server is paused by scaling down number of replicas of
	// the kubernetes stateful sets to zero in the environment.
	Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.KeycloakServer, error)
	// unpause a previously paused keycloak-server running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the keycloak-server.
	Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.KeycloakServer, error)
	// preview refresh a keycloak-server that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.KeycloakServer, error)
	// refresh a keycloak-server that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.KeycloakServer, error)
}

// UnimplementedKeycloakServerCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKeycloakServerCommandControllerServer struct {
}

func (UnimplementedKeycloakServerCommandControllerServer) PreviewCreate(context.Context, *model.KeycloakServer) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedKeycloakServerCommandControllerServer) Create(context.Context, *model.KeycloakServer) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedKeycloakServerCommandControllerServer) PreviewUpdate(context.Context, *model.KeycloakServer) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedKeycloakServerCommandControllerServer) Update(context.Context, *model.KeycloakServer) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedKeycloakServerCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedKeycloakServerCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKeycloakServerCommandControllerServer) PreviewRestore(context.Context, *model.KeycloakServer) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedKeycloakServerCommandControllerServer) Restore(context.Context, *model.KeycloakServer) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedKeycloakServerCommandControllerServer) Restart(context.Context, *model.KeycloakServerId) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedKeycloakServerCommandControllerServer) Pause(context.Context, *model1.ApiResourcePauseCommandInput) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedKeycloakServerCommandControllerServer) Unpause(context.Context, *model1.ApiResourceUnPauseCommandInput) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}
func (UnimplementedKeycloakServerCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedKeycloakServerCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.KeycloakServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeKeycloakServerCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeycloakServerCommandControllerServer will
// result in compilation errors.
type UnsafeKeycloakServerCommandControllerServer interface {
	mustEmbedUnimplementedKeycloakServerCommandControllerServer()
}

func RegisterKeycloakServerCommandControllerServer(s grpc.ServiceRegistrar, srv KeycloakServerCommandControllerServer) {
	s.RegisterService(&KeycloakServerCommandController_ServiceDesc, srv)
}

func _KeycloakServerCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KeycloakServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).PreviewCreate(ctx, req.(*model.KeycloakServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycloakServerCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KeycloakServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).Create(ctx, req.(*model.KeycloakServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycloakServerCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KeycloakServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).PreviewUpdate(ctx, req.(*model.KeycloakServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycloakServerCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KeycloakServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).Update(ctx, req.(*model.KeycloakServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycloakServerCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycloakServerCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycloakServerCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KeycloakServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).PreviewRestore(ctx, req.(*model.KeycloakServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycloakServerCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KeycloakServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).Restore(ctx, req.(*model.KeycloakServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycloakServerCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.KeycloakServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).Restart(ctx, req.(*model.KeycloakServerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycloakServerCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourcePauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).Pause(ctx, req.(*model1.ApiResourcePauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycloakServerCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceUnPauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).Unpause(ctx, req.(*model1.ApiResourceUnPauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycloakServerCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycloakServerCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycloakServerCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeycloakServerCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycloakServerCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// KeycloakServerCommandController_ServiceDesc is the grpc.ServiceDesc for KeycloakServerCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeycloakServerCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.keycloakserver.service.KeycloakServerCommandController",
	HandlerType: (*KeycloakServerCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _KeycloakServerCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _KeycloakServerCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _KeycloakServerCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _KeycloakServerCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _KeycloakServerCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _KeycloakServerCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _KeycloakServerCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _KeycloakServerCommandController_Restore_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _KeycloakServerCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _KeycloakServerCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _KeycloakServerCommandController_Unpause_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _KeycloakServerCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _KeycloakServerCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/keycloakserver/service/command.proto",
}
