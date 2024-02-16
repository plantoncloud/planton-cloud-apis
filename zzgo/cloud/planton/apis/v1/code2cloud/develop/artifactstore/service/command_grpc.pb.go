// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/develop/artifactstore/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/develop/artifactstore/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ArtifactStoreCommandController_PreviewCreate_FullMethodName                     = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController/previewCreate"
	ArtifactStoreCommandController_Create_FullMethodName                            = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController/create"
	ArtifactStoreCommandController_PreviewUpdate_FullMethodName                     = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController/previewUpdate"
	ArtifactStoreCommandController_Update_FullMethodName                            = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController/update"
	ArtifactStoreCommandController_PreviewDelete_FullMethodName                     = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController/previewDelete"
	ArtifactStoreCommandController_Delete_FullMethodName                            = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController/delete"
	ArtifactStoreCommandController_PreviewRestore_FullMethodName                    = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController/previewRestore"
	ArtifactStoreCommandController_Restore_FullMethodName                           = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController/restore"
	ArtifactStoreCommandController_CreateStackJob_FullMethodName                    = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController/createStackJob"
	ArtifactStoreCommandController_DeleteArtifactStorePackageVersion_FullMethodName = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController/deleteArtifactStorePackageVersion"
	ArtifactStoreCommandController_PreviewRefresh_FullMethodName                    = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController/previewRefresh"
	ArtifactStoreCommandController_Refresh_FullMethodName                           = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController/refresh"
)

// ArtifactStoreCommandControllerClient is the client API for ArtifactStoreCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtifactStoreCommandControllerClient interface {
	// preview create artifact-store
	PreviewCreate(ctx context.Context, in *model.ArtifactStore, opts ...grpc.CallOption) (*model.ArtifactStore, error)
	// create artifact-store
	Create(ctx context.Context, in *model.ArtifactStore, opts ...grpc.CallOption) (*model.ArtifactStore, error)
	// preview update artifact-store
	PreviewUpdate(ctx context.Context, in *model.ArtifactStore, opts ...grpc.CallOption) (*model.ArtifactStore, error)
	// update artifact-store
	Update(ctx context.Context, in *model.ArtifactStore, opts ...grpc.CallOption) (*model.ArtifactStore, error)
	// preview delete an artifact-store.
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.ArtifactStore, error)
	// delete an artifact-store.
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.ArtifactStore, error)
	// preview restoring a deleted artifact-store.
	PreviewRestore(ctx context.Context, in *model.ArtifactStore, opts ...grpc.CallOption) (*model.ArtifactStore, error)
	// restore a deleted artifact-store.
	Restore(ctx context.Context, in *model.ArtifactStore, opts ...grpc.CallOption) (*model.ArtifactStore, error)
	// create-stack-job for artifact-store
	CreateStackJob(ctx context.Context, in *model2.CreateStackJobCommandInput, opts ...grpc.CallOption) (*model.ArtifactStore, error)
	// restore a deleted artifact-store.
	DeleteArtifactStorePackageVersion(ctx context.Context, in *model.DelArtifactStorePackageVersionCommandInput, opts ...grpc.CallOption) (*model.ArtifactStorePackageVersion, error)
	// preview refresh a artifact-store that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.ArtifactStore, error)
	// refresh a artifact-store that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.ArtifactStore, error)
}

type artifactStoreCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewArtifactStoreCommandControllerClient(cc grpc.ClientConnInterface) ArtifactStoreCommandControllerClient {
	return &artifactStoreCommandControllerClient{cc}
}

func (c *artifactStoreCommandControllerClient) PreviewCreate(ctx context.Context, in *model.ArtifactStore, opts ...grpc.CallOption) (*model.ArtifactStore, error) {
	out := new(model.ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) Create(ctx context.Context, in *model.ArtifactStore, opts ...grpc.CallOption) (*model.ArtifactStore, error) {
	out := new(model.ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.ArtifactStore, opts ...grpc.CallOption) (*model.ArtifactStore, error) {
	out := new(model.ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) Update(ctx context.Context, in *model.ArtifactStore, opts ...grpc.CallOption) (*model.ArtifactStore, error) {
	out := new(model.ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.ArtifactStore, error) {
	out := new(model.ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.ArtifactStore, error) {
	out := new(model.ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) PreviewRestore(ctx context.Context, in *model.ArtifactStore, opts ...grpc.CallOption) (*model.ArtifactStore, error) {
	out := new(model.ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) Restore(ctx context.Context, in *model.ArtifactStore, opts ...grpc.CallOption) (*model.ArtifactStore, error) {
	out := new(model.ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) CreateStackJob(ctx context.Context, in *model2.CreateStackJobCommandInput, opts ...grpc.CallOption) (*model.ArtifactStore, error) {
	out := new(model.ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_CreateStackJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) DeleteArtifactStorePackageVersion(ctx context.Context, in *model.DelArtifactStorePackageVersionCommandInput, opts ...grpc.CallOption) (*model.ArtifactStorePackageVersion, error) {
	out := new(model.ArtifactStorePackageVersion)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_DeleteArtifactStorePackageVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.ArtifactStore, error) {
	out := new(model.ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.ArtifactStore, error) {
	out := new(model.ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtifactStoreCommandControllerServer is the server API for ArtifactStoreCommandController service.
// All implementations should embed UnimplementedArtifactStoreCommandControllerServer
// for forward compatibility
type ArtifactStoreCommandControllerServer interface {
	// preview create artifact-store
	PreviewCreate(context.Context, *model.ArtifactStore) (*model.ArtifactStore, error)
	// create artifact-store
	Create(context.Context, *model.ArtifactStore) (*model.ArtifactStore, error)
	// preview update artifact-store
	PreviewUpdate(context.Context, *model.ArtifactStore) (*model.ArtifactStore, error)
	// update artifact-store
	Update(context.Context, *model.ArtifactStore) (*model.ArtifactStore, error)
	// preview delete an artifact-store.
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.ArtifactStore, error)
	// delete an artifact-store.
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.ArtifactStore, error)
	// preview restoring a deleted artifact-store.
	PreviewRestore(context.Context, *model.ArtifactStore) (*model.ArtifactStore, error)
	// restore a deleted artifact-store.
	Restore(context.Context, *model.ArtifactStore) (*model.ArtifactStore, error)
	// create-stack-job for artifact-store
	CreateStackJob(context.Context, *model2.CreateStackJobCommandInput) (*model.ArtifactStore, error)
	// restore a deleted artifact-store.
	DeleteArtifactStorePackageVersion(context.Context, *model.DelArtifactStorePackageVersionCommandInput) (*model.ArtifactStorePackageVersion, error)
	// preview refresh a artifact-store that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.ArtifactStore, error)
	// refresh a artifact-store that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.ArtifactStore, error)
}

// UnimplementedArtifactStoreCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedArtifactStoreCommandControllerServer struct {
}

func (UnimplementedArtifactStoreCommandControllerServer) PreviewCreate(context.Context, *model.ArtifactStore) (*model.ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) Create(context.Context, *model.ArtifactStore) (*model.ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) PreviewUpdate(context.Context, *model.ArtifactStore) (*model.ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) Update(context.Context, *model.ArtifactStore) (*model.ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) PreviewRestore(context.Context, *model.ArtifactStore) (*model.ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) Restore(context.Context, *model.ArtifactStore) (*model.ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) CreateStackJob(context.Context, *model2.CreateStackJobCommandInput) (*model.ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStackJob not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) DeleteArtifactStorePackageVersion(context.Context, *model.DelArtifactStorePackageVersionCommandInput) (*model.ArtifactStorePackageVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArtifactStorePackageVersion not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeArtifactStoreCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtifactStoreCommandControllerServer will
// result in compilation errors.
type UnsafeArtifactStoreCommandControllerServer interface {
	mustEmbedUnimplementedArtifactStoreCommandControllerServer()
}

func RegisterArtifactStoreCommandControllerServer(s grpc.ServiceRegistrar, srv ArtifactStoreCommandControllerServer) {
	s.RegisterService(&ArtifactStoreCommandController_ServiceDesc, srv)
}

func _ArtifactStoreCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ArtifactStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreCommandControllerServer).PreviewCreate(ctx, req.(*model.ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ArtifactStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreCommandControllerServer).Create(ctx, req.(*model.ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ArtifactStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreCommandControllerServer).PreviewUpdate(ctx, req.(*model.ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ArtifactStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreCommandControllerServer).Update(ctx, req.(*model.ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ArtifactStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreCommandControllerServer).PreviewRestore(ctx, req.(*model.ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ArtifactStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreCommandControllerServer).Restore(ctx, req.(*model.ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_CreateStackJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.CreateStackJobCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreCommandControllerServer).CreateStackJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreCommandController_CreateStackJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreCommandControllerServer).CreateStackJob(ctx, req.(*model2.CreateStackJobCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_DeleteArtifactStorePackageVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.DelArtifactStorePackageVersionCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreCommandControllerServer).DeleteArtifactStorePackageVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreCommandController_DeleteArtifactStorePackageVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreCommandControllerServer).DeleteArtifactStorePackageVersion(ctx, req.(*model.DelArtifactStorePackageVersionCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtifactStoreCommandController_ServiceDesc is the grpc.ServiceDesc for ArtifactStoreCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtifactStoreCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreCommandController",
	HandlerType: (*ArtifactStoreCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _ArtifactStoreCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _ArtifactStoreCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _ArtifactStoreCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _ArtifactStoreCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _ArtifactStoreCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _ArtifactStoreCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _ArtifactStoreCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _ArtifactStoreCommandController_Restore_Handler,
		},
		{
			MethodName: "createStackJob",
			Handler:    _ArtifactStoreCommandController_CreateStackJob_Handler,
		},
		{
			MethodName: "deleteArtifactStorePackageVersion",
			Handler:    _ArtifactStoreCommandController_DeleteArtifactStorePackageVersion_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _ArtifactStoreCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _ArtifactStoreCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/develop/artifactstore/service/command.proto",
}
