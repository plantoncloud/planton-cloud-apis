// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/gcp/gcpartifactregistry/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpartifactregistry/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/kubernetesdockercredential/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GcpArtifactRegistryCommandController_PreviewCreate_FullMethodName                    = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryCommandController/previewCreate"
	GcpArtifactRegistryCommandController_Create_FullMethodName                           = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryCommandController/create"
	GcpArtifactRegistryCommandController_PreviewUpdate_FullMethodName                    = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryCommandController/previewUpdate"
	GcpArtifactRegistryCommandController_Update_FullMethodName                           = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryCommandController/update"
	GcpArtifactRegistryCommandController_PreviewDelete_FullMethodName                    = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryCommandController/previewDelete"
	GcpArtifactRegistryCommandController_Delete_FullMethodName                           = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryCommandController/delete"
	GcpArtifactRegistryCommandController_PreviewRestore_FullMethodName                   = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryCommandController/previewRestore"
	GcpArtifactRegistryCommandController_Restore_FullMethodName                          = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryCommandController/restore"
	GcpArtifactRegistryCommandController_PreviewRefresh_FullMethodName                   = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryCommandController/previewRefresh"
	GcpArtifactRegistryCommandController_Refresh_FullMethodName                          = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryCommandController/refresh"
	GcpArtifactRegistryCommandController_CreateKubernetesDockerCredential_FullMethodName = "/cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryCommandController/createKubernetesDockerCredential"
)

// GcpArtifactRegistryCommandControllerClient is the client API for GcpArtifactRegistryCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GcpArtifactRegistryCommandControllerClient interface {
	// preview create gcp-artifact-registry
	PreviewCreate(ctx context.Context, in *model.GcpArtifactRegistry, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error)
	// create gcp-artifact-registry
	Create(ctx context.Context, in *model.GcpArtifactRegistry, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error)
	// preview update gcp-artifact-registry
	PreviewUpdate(ctx context.Context, in *model.GcpArtifactRegistry, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error)
	// update gcp-artifact-registry
	Update(ctx context.Context, in *model.GcpArtifactRegistry, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error)
	// preview delete an gcp-artifact-registry.
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error)
	// delete an gcp-artifact-registry.
	Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error)
	// preview restoring a deleted gcp-artifact-registry.
	PreviewRestore(ctx context.Context, in *model.GcpArtifactRegistry, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error)
	// restore a deleted gcp-artifact-registry.
	Restore(ctx context.Context, in *model.GcpArtifactRegistry, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error)
	// preview refresh a gcp-artifact-registry that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error)
	// refresh a gcp-artifact-registry that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error)
	CreateKubernetesDockerCredential(ctx context.Context, in *model.GcpArtifactRegistryId, opts ...grpc.CallOption) (*model2.KubernetesDockerCredential, error)
}

type gcpArtifactRegistryCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGcpArtifactRegistryCommandControllerClient(cc grpc.ClientConnInterface) GcpArtifactRegistryCommandControllerClient {
	return &gcpArtifactRegistryCommandControllerClient{cc}
}

func (c *gcpArtifactRegistryCommandControllerClient) PreviewCreate(ctx context.Context, in *model.GcpArtifactRegistry, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error) {
	out := new(model.GcpArtifactRegistry)
	err := c.cc.Invoke(ctx, GcpArtifactRegistryCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpArtifactRegistryCommandControllerClient) Create(ctx context.Context, in *model.GcpArtifactRegistry, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error) {
	out := new(model.GcpArtifactRegistry)
	err := c.cc.Invoke(ctx, GcpArtifactRegistryCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpArtifactRegistryCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.GcpArtifactRegistry, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error) {
	out := new(model.GcpArtifactRegistry)
	err := c.cc.Invoke(ctx, GcpArtifactRegistryCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpArtifactRegistryCommandControllerClient) Update(ctx context.Context, in *model.GcpArtifactRegistry, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error) {
	out := new(model.GcpArtifactRegistry)
	err := c.cc.Invoke(ctx, GcpArtifactRegistryCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpArtifactRegistryCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error) {
	out := new(model.GcpArtifactRegistry)
	err := c.cc.Invoke(ctx, GcpArtifactRegistryCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpArtifactRegistryCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteInput, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error) {
	out := new(model.GcpArtifactRegistry)
	err := c.cc.Invoke(ctx, GcpArtifactRegistryCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpArtifactRegistryCommandControllerClient) PreviewRestore(ctx context.Context, in *model.GcpArtifactRegistry, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error) {
	out := new(model.GcpArtifactRegistry)
	err := c.cc.Invoke(ctx, GcpArtifactRegistryCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpArtifactRegistryCommandControllerClient) Restore(ctx context.Context, in *model.GcpArtifactRegistry, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error) {
	out := new(model.GcpArtifactRegistry)
	err := c.cc.Invoke(ctx, GcpArtifactRegistryCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpArtifactRegistryCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error) {
	out := new(model.GcpArtifactRegistry)
	err := c.cc.Invoke(ctx, GcpArtifactRegistryCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpArtifactRegistryCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshInput, opts ...grpc.CallOption) (*model.GcpArtifactRegistry, error) {
	out := new(model.GcpArtifactRegistry)
	err := c.cc.Invoke(ctx, GcpArtifactRegistryCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpArtifactRegistryCommandControllerClient) CreateKubernetesDockerCredential(ctx context.Context, in *model.GcpArtifactRegistryId, opts ...grpc.CallOption) (*model2.KubernetesDockerCredential, error) {
	out := new(model2.KubernetesDockerCredential)
	err := c.cc.Invoke(ctx, GcpArtifactRegistryCommandController_CreateKubernetesDockerCredential_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GcpArtifactRegistryCommandControllerServer is the server API for GcpArtifactRegistryCommandController service.
// All implementations should embed UnimplementedGcpArtifactRegistryCommandControllerServer
// for forward compatibility
type GcpArtifactRegistryCommandControllerServer interface {
	// preview create gcp-artifact-registry
	PreviewCreate(context.Context, *model.GcpArtifactRegistry) (*model.GcpArtifactRegistry, error)
	// create gcp-artifact-registry
	Create(context.Context, *model.GcpArtifactRegistry) (*model.GcpArtifactRegistry, error)
	// preview update gcp-artifact-registry
	PreviewUpdate(context.Context, *model.GcpArtifactRegistry) (*model.GcpArtifactRegistry, error)
	// update gcp-artifact-registry
	Update(context.Context, *model.GcpArtifactRegistry) (*model.GcpArtifactRegistry, error)
	// preview delete an gcp-artifact-registry.
	PreviewDelete(context.Context, *model1.ApiResourceDeleteInput) (*model.GcpArtifactRegistry, error)
	// delete an gcp-artifact-registry.
	Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.GcpArtifactRegistry, error)
	// preview restoring a deleted gcp-artifact-registry.
	PreviewRestore(context.Context, *model.GcpArtifactRegistry) (*model.GcpArtifactRegistry, error)
	// restore a deleted gcp-artifact-registry.
	Restore(context.Context, *model.GcpArtifactRegistry) (*model.GcpArtifactRegistry, error)
	// preview refresh a gcp-artifact-registry that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshInput) (*model.GcpArtifactRegistry, error)
	// refresh a gcp-artifact-registry that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshInput) (*model.GcpArtifactRegistry, error)
	CreateKubernetesDockerCredential(context.Context, *model.GcpArtifactRegistryId) (*model2.KubernetesDockerCredential, error)
}

// UnimplementedGcpArtifactRegistryCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGcpArtifactRegistryCommandControllerServer struct {
}

func (UnimplementedGcpArtifactRegistryCommandControllerServer) PreviewCreate(context.Context, *model.GcpArtifactRegistry) (*model.GcpArtifactRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedGcpArtifactRegistryCommandControllerServer) Create(context.Context, *model.GcpArtifactRegistry) (*model.GcpArtifactRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGcpArtifactRegistryCommandControllerServer) PreviewUpdate(context.Context, *model.GcpArtifactRegistry) (*model.GcpArtifactRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedGcpArtifactRegistryCommandControllerServer) Update(context.Context, *model.GcpArtifactRegistry) (*model.GcpArtifactRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGcpArtifactRegistryCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteInput) (*model.GcpArtifactRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedGcpArtifactRegistryCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteInput) (*model.GcpArtifactRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGcpArtifactRegistryCommandControllerServer) PreviewRestore(context.Context, *model.GcpArtifactRegistry) (*model.GcpArtifactRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedGcpArtifactRegistryCommandControllerServer) Restore(context.Context, *model.GcpArtifactRegistry) (*model.GcpArtifactRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedGcpArtifactRegistryCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshInput) (*model.GcpArtifactRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedGcpArtifactRegistryCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshInput) (*model.GcpArtifactRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedGcpArtifactRegistryCommandControllerServer) CreateKubernetesDockerCredential(context.Context, *model.GcpArtifactRegistryId) (*model2.KubernetesDockerCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKubernetesDockerCredential not implemented")
}

// UnsafeGcpArtifactRegistryCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GcpArtifactRegistryCommandControllerServer will
// result in compilation errors.
type UnsafeGcpArtifactRegistryCommandControllerServer interface {
	mustEmbedUnimplementedGcpArtifactRegistryCommandControllerServer()
}

func RegisterGcpArtifactRegistryCommandControllerServer(s grpc.ServiceRegistrar, srv GcpArtifactRegistryCommandControllerServer) {
	s.RegisterService(&GcpArtifactRegistryCommandController_ServiceDesc, srv)
}

func _GcpArtifactRegistryCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpArtifactRegistry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpArtifactRegistryCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpArtifactRegistryCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpArtifactRegistryCommandControllerServer).PreviewCreate(ctx, req.(*model.GcpArtifactRegistry))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpArtifactRegistryCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpArtifactRegistry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpArtifactRegistryCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpArtifactRegistryCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpArtifactRegistryCommandControllerServer).Create(ctx, req.(*model.GcpArtifactRegistry))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpArtifactRegistryCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpArtifactRegistry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpArtifactRegistryCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpArtifactRegistryCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpArtifactRegistryCommandControllerServer).PreviewUpdate(ctx, req.(*model.GcpArtifactRegistry))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpArtifactRegistryCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpArtifactRegistry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpArtifactRegistryCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpArtifactRegistryCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpArtifactRegistryCommandControllerServer).Update(ctx, req.(*model.GcpArtifactRegistry))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpArtifactRegistryCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpArtifactRegistryCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpArtifactRegistryCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpArtifactRegistryCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpArtifactRegistryCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpArtifactRegistryCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpArtifactRegistryCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpArtifactRegistryCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpArtifactRegistryCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpArtifactRegistry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpArtifactRegistryCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpArtifactRegistryCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpArtifactRegistryCommandControllerServer).PreviewRestore(ctx, req.(*model.GcpArtifactRegistry))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpArtifactRegistryCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpArtifactRegistry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpArtifactRegistryCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpArtifactRegistryCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpArtifactRegistryCommandControllerServer).Restore(ctx, req.(*model.GcpArtifactRegistry))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpArtifactRegistryCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpArtifactRegistryCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpArtifactRegistryCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpArtifactRegistryCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpArtifactRegistryCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpArtifactRegistryCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpArtifactRegistryCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpArtifactRegistryCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpArtifactRegistryCommandController_CreateKubernetesDockerCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GcpArtifactRegistryId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpArtifactRegistryCommandControllerServer).CreateKubernetesDockerCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpArtifactRegistryCommandController_CreateKubernetesDockerCredential_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpArtifactRegistryCommandControllerServer).CreateKubernetesDockerCredential(ctx, req.(*model.GcpArtifactRegistryId))
	}
	return interceptor(ctx, in, info, handler)
}

// GcpArtifactRegistryCommandController_ServiceDesc is the grpc.ServiceDesc for GcpArtifactRegistryCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GcpArtifactRegistryCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.service.GcpArtifactRegistryCommandController",
	HandlerType: (*GcpArtifactRegistryCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _GcpArtifactRegistryCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _GcpArtifactRegistryCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _GcpArtifactRegistryCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _GcpArtifactRegistryCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _GcpArtifactRegistryCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _GcpArtifactRegistryCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _GcpArtifactRegistryCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _GcpArtifactRegistryCommandController_Restore_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _GcpArtifactRegistryCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _GcpArtifactRegistryCommandController_Refresh_Handler,
		},
		{
			MethodName: "createKubernetesDockerCredential",
			Handler:    _GcpArtifactRegistryCommandController_CreateKubernetesDockerCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/gcp/gcpartifactregistry/service/command.proto",
}
