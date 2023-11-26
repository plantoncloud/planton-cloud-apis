// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/develop/artifactstore/service.proto

package artifactstore

import (
	context "context"
	project "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/develop/sourcecode/project"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	product "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/product"
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
	ArtifactStoreCommandController_PreviewCreate_FullMethodName                     = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreCommandController/previewCreate"
	ArtifactStoreCommandController_Create_FullMethodName                            = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreCommandController/create"
	ArtifactStoreCommandController_PreviewUpdate_FullMethodName                     = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreCommandController/previewUpdate"
	ArtifactStoreCommandController_Update_FullMethodName                            = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreCommandController/update"
	ArtifactStoreCommandController_PreviewDelete_FullMethodName                     = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreCommandController/previewDelete"
	ArtifactStoreCommandController_Delete_FullMethodName                            = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreCommandController/delete"
	ArtifactStoreCommandController_PreviewRestore_FullMethodName                    = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreCommandController/previewRestore"
	ArtifactStoreCommandController_Restore_FullMethodName                           = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreCommandController/restore"
	ArtifactStoreCommandController_CreateStackJob_FullMethodName                    = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreCommandController/createStackJob"
	ArtifactStoreCommandController_DeleteArtifactStorePackageVersion_FullMethodName = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreCommandController/deleteArtifactStorePackageVersion"
)

// ArtifactStoreCommandControllerClient is the client API for ArtifactStoreCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtifactStoreCommandControllerClient interface {
	// preview create artifact-store
	PreviewCreate(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error)
	// create artifact-store
	Create(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error)
	// preview update artifact-store
	PreviewUpdate(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error)
	// update artifact-store
	Update(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error)
	// preview delete an artifact-store.
	PreviewDelete(ctx context.Context, in *ArtifactStoreId, opts ...grpc.CallOption) (*ArtifactStore, error)
	// delete an artifact-store.
	Delete(ctx context.Context, in *ArtifactStoreId, opts ...grpc.CallOption) (*ArtifactStore, error)
	// preview restoring a deleted artifact-store.
	PreviewRestore(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error)
	// restore a deleted artifact-store.
	Restore(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error)
	// create-stack-job for artifact-store
	CreateStackJob(ctx context.Context, in *job.CreateStackJobCommandInput, opts ...grpc.CallOption) (*ArtifactStore, error)
	// restore a deleted artifact-store.
	DeleteArtifactStorePackageVersion(ctx context.Context, in *DelArtifactStorePackageVersionCommandInput, opts ...grpc.CallOption) (*ArtifactStorePackageVersion, error)
}

type artifactStoreCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewArtifactStoreCommandControllerClient(cc grpc.ClientConnInterface) ArtifactStoreCommandControllerClient {
	return &artifactStoreCommandControllerClient{cc}
}

func (c *artifactStoreCommandControllerClient) PreviewCreate(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) Create(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) PreviewUpdate(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) Update(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) PreviewDelete(ctx context.Context, in *ArtifactStoreId, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) Delete(ctx context.Context, in *ArtifactStoreId, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) PreviewRestore(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) Restore(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) CreateStackJob(ctx context.Context, in *job.CreateStackJobCommandInput, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_CreateStackJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreCommandControllerClient) DeleteArtifactStorePackageVersion(ctx context.Context, in *DelArtifactStorePackageVersionCommandInput, opts ...grpc.CallOption) (*ArtifactStorePackageVersion, error) {
	out := new(ArtifactStorePackageVersion)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_DeleteArtifactStorePackageVersion_FullMethodName, in, out, opts...)
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
	PreviewCreate(context.Context, *ArtifactStore) (*ArtifactStore, error)
	// create artifact-store
	Create(context.Context, *ArtifactStore) (*ArtifactStore, error)
	// preview update artifact-store
	PreviewUpdate(context.Context, *ArtifactStore) (*ArtifactStore, error)
	// update artifact-store
	Update(context.Context, *ArtifactStore) (*ArtifactStore, error)
	// preview delete an artifact-store.
	PreviewDelete(context.Context, *ArtifactStoreId) (*ArtifactStore, error)
	// delete an artifact-store.
	Delete(context.Context, *ArtifactStoreId) (*ArtifactStore, error)
	// preview restoring a deleted artifact-store.
	PreviewRestore(context.Context, *ArtifactStore) (*ArtifactStore, error)
	// restore a deleted artifact-store.
	Restore(context.Context, *ArtifactStore) (*ArtifactStore, error)
	// create-stack-job for artifact-store
	CreateStackJob(context.Context, *job.CreateStackJobCommandInput) (*ArtifactStore, error)
	// restore a deleted artifact-store.
	DeleteArtifactStorePackageVersion(context.Context, *DelArtifactStorePackageVersionCommandInput) (*ArtifactStorePackageVersion, error)
}

// UnimplementedArtifactStoreCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedArtifactStoreCommandControllerServer struct {
}

func (UnimplementedArtifactStoreCommandControllerServer) PreviewCreate(context.Context, *ArtifactStore) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) Create(context.Context, *ArtifactStore) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) PreviewUpdate(context.Context, *ArtifactStore) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) Update(context.Context, *ArtifactStore) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) PreviewDelete(context.Context, *ArtifactStoreId) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) Delete(context.Context, *ArtifactStoreId) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) PreviewRestore(context.Context, *ArtifactStore) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) Restore(context.Context, *ArtifactStore) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) CreateStackJob(context.Context, *job.CreateStackJobCommandInput) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStackJob not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) DeleteArtifactStorePackageVersion(context.Context, *DelArtifactStorePackageVersionCommandInput) (*ArtifactStorePackageVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArtifactStorePackageVersion not implemented")
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
	in := new(ArtifactStore)
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
		return srv.(ArtifactStoreCommandControllerServer).PreviewCreate(ctx, req.(*ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactStore)
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
		return srv.(ArtifactStoreCommandControllerServer).Create(ctx, req.(*ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactStore)
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
		return srv.(ArtifactStoreCommandControllerServer).PreviewUpdate(ctx, req.(*ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactStore)
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
		return srv.(ArtifactStoreCommandControllerServer).Update(ctx, req.(*ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactStoreId)
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
		return srv.(ArtifactStoreCommandControllerServer).PreviewDelete(ctx, req.(*ArtifactStoreId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactStoreId)
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
		return srv.(ArtifactStoreCommandControllerServer).Delete(ctx, req.(*ArtifactStoreId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactStore)
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
		return srv.(ArtifactStoreCommandControllerServer).PreviewRestore(ctx, req.(*ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactStore)
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
		return srv.(ArtifactStoreCommandControllerServer).Restore(ctx, req.(*ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_CreateStackJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(job.CreateStackJobCommandInput)
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
		return srv.(ArtifactStoreCommandControllerServer).CreateStackJob(ctx, req.(*job.CreateStackJobCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreCommandController_DeleteArtifactStorePackageVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelArtifactStorePackageVersionCommandInput)
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
		return srv.(ArtifactStoreCommandControllerServer).DeleteArtifactStorePackageVersion(ctx, req.(*DelArtifactStorePackageVersionCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtifactStoreCommandController_ServiceDesc is the grpc.ServiceDesc for ArtifactStoreCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtifactStoreCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreCommandController",
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/develop/artifactstore/service.proto",
}

const (
	ArtifactStoreQueryController_GetById_FullMethodName                          = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreQueryController/getById"
	ArtifactStoreQueryController_List_FullMethodName                             = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreQueryController/list"
	ArtifactStoreQueryController_FindByProductId_FullMethodName                  = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreQueryController/findByProductId"
	ArtifactStoreQueryController_FindByCodeProjectUrl_FullMethodName             = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreQueryController/findByCodeProjectUrl"
	ArtifactStoreQueryController_ListArtifactStoreDockerImages_FullMethodName    = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreQueryController/listArtifactStoreDockerImages"
	ArtifactStoreQueryController_ListArtifactStorePackages_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreQueryController/listArtifactStorePackages"
	ArtifactStoreQueryController_ListArtifactStorePackageVersions_FullMethodName = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreQueryController/listArtifactStorePackageVersions"
)

// ArtifactStoreQueryControllerClient is the client API for ArtifactStoreQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtifactStoreQueryControllerClient interface {
	// get artifact-store by id
	GetById(ctx context.Context, in *ArtifactStoreId, opts ...grpc.CallOption) (*ArtifactStore, error)
	// list all artifact-stores on planton cloud for the requested page. This is intended for use on portal.
	List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*ArtifactStoreList, error)
	// look up artifact-stores by product id.
	FindByProductId(ctx context.Context, in *product.ProductId, opts ...grpc.CallOption) (*ArtifactStores, error)
	// look up artifact-stores by code project url
	FindByCodeProjectUrl(ctx context.Context, in *project.CodeProjectUrl, opts ...grpc.CallOption) (*ArtifactStores, error)
	// list docker images from the artifact-store provided in the input
	// (proxy google artifact-registry server)
	ListArtifactStoreDockerImages(ctx context.Context, in *ListByArtifactStoreIdRepoNameInput, opts ...grpc.CallOption) (*ArtifactStoreDockerImageList, error)
	// list maven/npm/python packages from the corresponding repositories of artifact-store provided in the input
	// (proxy google artifact-registry server)
	ListArtifactStorePackages(ctx context.Context, in *ListByArtifactStoreIdRepoNameInput, opts ...grpc.CallOption) (*ArtifactStorePackageList, error)
	// list maven/npm/python package versions from the corresponding repositories of artifact-store provided in the input
	// (proxy google artifact-registry server)
	ListArtifactStorePackageVersions(ctx context.Context, in *ListByArtifactStoreIdPackageNameInput, opts ...grpc.CallOption) (*ArtifactStorePackageVersionList, error)
}

type artifactStoreQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewArtifactStoreQueryControllerClient(cc grpc.ClientConnInterface) ArtifactStoreQueryControllerClient {
	return &artifactStoreQueryControllerClient{cc}
}

func (c *artifactStoreQueryControllerClient) GetById(ctx context.Context, in *ArtifactStoreId, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*ArtifactStoreList, error) {
	out := new(ArtifactStoreList)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) FindByProductId(ctx context.Context, in *product.ProductId, opts ...grpc.CallOption) (*ArtifactStores, error) {
	out := new(ArtifactStores)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) FindByCodeProjectUrl(ctx context.Context, in *project.CodeProjectUrl, opts ...grpc.CallOption) (*ArtifactStores, error) {
	out := new(ArtifactStores)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_FindByCodeProjectUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) ListArtifactStoreDockerImages(ctx context.Context, in *ListByArtifactStoreIdRepoNameInput, opts ...grpc.CallOption) (*ArtifactStoreDockerImageList, error) {
	out := new(ArtifactStoreDockerImageList)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_ListArtifactStoreDockerImages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) ListArtifactStorePackages(ctx context.Context, in *ListByArtifactStoreIdRepoNameInput, opts ...grpc.CallOption) (*ArtifactStorePackageList, error) {
	out := new(ArtifactStorePackageList)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_ListArtifactStorePackages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) ListArtifactStorePackageVersions(ctx context.Context, in *ListByArtifactStoreIdPackageNameInput, opts ...grpc.CallOption) (*ArtifactStorePackageVersionList, error) {
	out := new(ArtifactStorePackageVersionList)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_ListArtifactStorePackageVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtifactStoreQueryControllerServer is the server API for ArtifactStoreQueryController service.
// All implementations should embed UnimplementedArtifactStoreQueryControllerServer
// for forward compatibility
type ArtifactStoreQueryControllerServer interface {
	// get artifact-store by id
	GetById(context.Context, *ArtifactStoreId) (*ArtifactStore, error)
	// list all artifact-stores on planton cloud for the requested page. This is intended for use on portal.
	List(context.Context, *pagination.PageInfo) (*ArtifactStoreList, error)
	// look up artifact-stores by product id.
	FindByProductId(context.Context, *product.ProductId) (*ArtifactStores, error)
	// look up artifact-stores by code project url
	FindByCodeProjectUrl(context.Context, *project.CodeProjectUrl) (*ArtifactStores, error)
	// list docker images from the artifact-store provided in the input
	// (proxy google artifact-registry server)
	ListArtifactStoreDockerImages(context.Context, *ListByArtifactStoreIdRepoNameInput) (*ArtifactStoreDockerImageList, error)
	// list maven/npm/python packages from the corresponding repositories of artifact-store provided in the input
	// (proxy google artifact-registry server)
	ListArtifactStorePackages(context.Context, *ListByArtifactStoreIdRepoNameInput) (*ArtifactStorePackageList, error)
	// list maven/npm/python package versions from the corresponding repositories of artifact-store provided in the input
	// (proxy google artifact-registry server)
	ListArtifactStorePackageVersions(context.Context, *ListByArtifactStoreIdPackageNameInput) (*ArtifactStorePackageVersionList, error)
}

// UnimplementedArtifactStoreQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedArtifactStoreQueryControllerServer struct {
}

func (UnimplementedArtifactStoreQueryControllerServer) GetById(context.Context, *ArtifactStoreId) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) List(context.Context, *pagination.PageInfo) (*ArtifactStoreList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) FindByProductId(context.Context, *product.ProductId) (*ArtifactStores, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) FindByCodeProjectUrl(context.Context, *project.CodeProjectUrl) (*ArtifactStores, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCodeProjectUrl not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) ListArtifactStoreDockerImages(context.Context, *ListByArtifactStoreIdRepoNameInput) (*ArtifactStoreDockerImageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtifactStoreDockerImages not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) ListArtifactStorePackages(context.Context, *ListByArtifactStoreIdRepoNameInput) (*ArtifactStorePackageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtifactStorePackages not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) ListArtifactStorePackageVersions(context.Context, *ListByArtifactStoreIdPackageNameInput) (*ArtifactStorePackageVersionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtifactStorePackageVersions not implemented")
}

// UnsafeArtifactStoreQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtifactStoreQueryControllerServer will
// result in compilation errors.
type UnsafeArtifactStoreQueryControllerServer interface {
	mustEmbedUnimplementedArtifactStoreQueryControllerServer()
}

func RegisterArtifactStoreQueryControllerServer(s grpc.ServiceRegistrar, srv ArtifactStoreQueryControllerServer) {
	s.RegisterService(&ArtifactStoreQueryController_ServiceDesc, srv)
}

func _ArtifactStoreQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactStoreId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreQueryControllerServer).GetById(ctx, req.(*ArtifactStoreId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreQueryControllerServer).List(ctx, req.(*pagination.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(product.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreQueryControllerServer).FindByProductId(ctx, req.(*product.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_FindByCodeProjectUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(project.CodeProjectUrl)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreQueryControllerServer).FindByCodeProjectUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreQueryController_FindByCodeProjectUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreQueryControllerServer).FindByCodeProjectUrl(ctx, req.(*project.CodeProjectUrl))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_ListArtifactStoreDockerImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListByArtifactStoreIdRepoNameInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreQueryControllerServer).ListArtifactStoreDockerImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreQueryController_ListArtifactStoreDockerImages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreQueryControllerServer).ListArtifactStoreDockerImages(ctx, req.(*ListByArtifactStoreIdRepoNameInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_ListArtifactStorePackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListByArtifactStoreIdRepoNameInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreQueryControllerServer).ListArtifactStorePackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreQueryController_ListArtifactStorePackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreQueryControllerServer).ListArtifactStorePackages(ctx, req.(*ListByArtifactStoreIdRepoNameInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_ListArtifactStorePackageVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListByArtifactStoreIdPackageNameInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreQueryControllerServer).ListArtifactStorePackageVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreQueryController_ListArtifactStorePackageVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreQueryControllerServer).ListArtifactStorePackageVersions(ctx, req.(*ListByArtifactStoreIdPackageNameInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtifactStoreQueryController_ServiceDesc is the grpc.ServiceDesc for ArtifactStoreQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtifactStoreQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.develop.artifactstore.ArtifactStoreQueryController",
	HandlerType: (*ArtifactStoreQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _ArtifactStoreQueryController_GetById_Handler,
		},
		{
			MethodName: "list",
			Handler:    _ArtifactStoreQueryController_List_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _ArtifactStoreQueryController_FindByProductId_Handler,
		},
		{
			MethodName: "findByCodeProjectUrl",
			Handler:    _ArtifactStoreQueryController_FindByCodeProjectUrl_Handler,
		},
		{
			MethodName: "listArtifactStoreDockerImages",
			Handler:    _ArtifactStoreQueryController_ListArtifactStoreDockerImages_Handler,
		},
		{
			MethodName: "listArtifactStorePackages",
			Handler:    _ArtifactStoreQueryController_ListArtifactStorePackages_Handler,
		},
		{
			MethodName: "listArtifactStorePackageVersions",
			Handler:    _ArtifactStoreQueryController_ListArtifactStorePackageVersions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/develop/artifactstore/service.proto",
}
