// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/develop/artifactstore/rpc/service.proto

package rpc

import (
	context "context"
	rpc1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/develop/sourcecode/project/rpc"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/product/rpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ArtifactStoreCommandController_Create_FullMethodName                            = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController/create"
	ArtifactStoreCommandController_Update_FullMethodName                            = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController/update"
	ArtifactStoreCommandController_Delete_FullMethodName                            = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController/delete"
	ArtifactStoreCommandController_Restore_FullMethodName                           = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController/restore"
	ArtifactStoreCommandController_DeleteArtifactStorePackageVersion_FullMethodName = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController/deleteArtifactStorePackageVersion"
)

// ArtifactStoreCommandControllerClient is the client API for ArtifactStoreCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtifactStoreCommandControllerClient interface {
	// create artifact-store
	Create(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error)
	// update artifact-store
	Update(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error)
	// delete an artifact-store.
	Delete(ctx context.Context, in *ArtifactStoreId, opts ...grpc.CallOption) (*ArtifactStore, error)
	// restore a deleted artifact-store.
	Restore(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error)
	// restore a deleted artifact-store.
	DeleteArtifactStorePackageVersion(ctx context.Context, in *DelArtifactStorePackageVersionCommandInput, opts ...grpc.CallOption) (*ArtifactStorePackageVersion, error)
}

type artifactStoreCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewArtifactStoreCommandControllerClient(cc grpc.ClientConnInterface) ArtifactStoreCommandControllerClient {
	return &artifactStoreCommandControllerClient{cc}
}

func (c *artifactStoreCommandControllerClient) Create(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_Create_FullMethodName, in, out, opts...)
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

func (c *artifactStoreCommandControllerClient) Delete(ctx context.Context, in *ArtifactStoreId, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreCommandController_Delete_FullMethodName, in, out, opts...)
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
	// create artifact-store
	Create(context.Context, *ArtifactStore) (*ArtifactStore, error)
	// update artifact-store
	Update(context.Context, *ArtifactStore) (*ArtifactStore, error)
	// delete an artifact-store.
	Delete(context.Context, *ArtifactStoreId) (*ArtifactStore, error)
	// restore a deleted artifact-store.
	Restore(context.Context, *ArtifactStore) (*ArtifactStore, error)
	// restore a deleted artifact-store.
	DeleteArtifactStorePackageVersion(context.Context, *DelArtifactStorePackageVersionCommandInput) (*ArtifactStorePackageVersion, error)
}

// UnimplementedArtifactStoreCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedArtifactStoreCommandControllerServer struct {
}

func (UnimplementedArtifactStoreCommandControllerServer) Create(context.Context, *ArtifactStore) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) Update(context.Context, *ArtifactStore) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) Delete(context.Context, *ArtifactStoreId) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedArtifactStoreCommandControllerServer) Restore(context.Context, *ArtifactStore) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
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
	ServiceName: "cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController",
	HandlerType: (*ArtifactStoreCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _ArtifactStoreCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _ArtifactStoreCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _ArtifactStoreCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _ArtifactStoreCommandController_Restore_Handler,
		},
		{
			MethodName: "deleteArtifactStorePackageVersion",
			Handler:    _ArtifactStoreCommandController_DeleteArtifactStorePackageVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/develop/artifactstore/rpc/service.proto",
}

const (
	ArtifactStoreQueryController_GetById_FullMethodName                          = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController/getById"
	ArtifactStoreQueryController_List_FullMethodName                             = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController/list"
	ArtifactStoreQueryController_FindByProductId_FullMethodName                  = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController/findByProductId"
	ArtifactStoreQueryController_FindByCodeProjectUrl_FullMethodName             = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController/findByCodeProjectUrl"
	ArtifactStoreQueryController_ListArtifactStoreDockerImages_FullMethodName    = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController/listArtifactStoreDockerImages"
	ArtifactStoreQueryController_ListArtifactStorePackages_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController/listArtifactStorePackages"
	ArtifactStoreQueryController_ListArtifactStorePackageVersions_FullMethodName = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController/listArtifactStorePackageVersions"
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
	FindByProductId(ctx context.Context, in *rpc.ProductId, opts ...grpc.CallOption) (*ArtifactStores, error)
	// look up artifact-stores by code project url
	FindByCodeProjectUrl(ctx context.Context, in *rpc1.CodeProjectUrl, opts ...grpc.CallOption) (*ArtifactStores, error)
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

func (c *artifactStoreQueryControllerClient) FindByProductId(ctx context.Context, in *rpc.ProductId, opts ...grpc.CallOption) (*ArtifactStores, error) {
	out := new(ArtifactStores)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) FindByCodeProjectUrl(ctx context.Context, in *rpc1.CodeProjectUrl, opts ...grpc.CallOption) (*ArtifactStores, error) {
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
	FindByProductId(context.Context, *rpc.ProductId) (*ArtifactStores, error)
	// look up artifact-stores by code project url
	FindByCodeProjectUrl(context.Context, *rpc1.CodeProjectUrl) (*ArtifactStores, error)
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
func (UnimplementedArtifactStoreQueryControllerServer) FindByProductId(context.Context, *rpc.ProductId) (*ArtifactStores, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) FindByCodeProjectUrl(context.Context, *rpc1.CodeProjectUrl) (*ArtifactStores, error) {
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
	in := new(rpc.ProductId)
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
		return srv.(ArtifactStoreQueryControllerServer).FindByProductId(ctx, req.(*rpc.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_FindByCodeProjectUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc1.CodeProjectUrl)
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
		return srv.(ArtifactStoreQueryControllerServer).FindByCodeProjectUrl(ctx, req.(*rpc1.CodeProjectUrl))
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
	ServiceName: "cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController",
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
	Metadata: "cloud/planton/apis/v1/code2cloud/develop/artifactstore/rpc/service.proto",
}

const (
	ArtifactStoreStackController_Preview_FullMethodName = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreStackController/preview"
	ArtifactStoreStackController_Apply_FullMethodName   = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreStackController/apply"
)

// ArtifactStoreStackControllerClient is the client API for ArtifactStoreStackController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtifactStoreStackControllerClient interface {
	// preview artifact-store stack
	Preview(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error)
	// apply artifact-store stack
	Apply(ctx context.Context, in *ArtifactStoreId, opts ...grpc.CallOption) (*ArtifactStore, error)
}

type artifactStoreStackControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewArtifactStoreStackControllerClient(cc grpc.ClientConnInterface) ArtifactStoreStackControllerClient {
	return &artifactStoreStackControllerClient{cc}
}

func (c *artifactStoreStackControllerClient) Preview(ctx context.Context, in *ArtifactStore, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreStackController_Preview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreStackControllerClient) Apply(ctx context.Context, in *ArtifactStoreId, opts ...grpc.CallOption) (*ArtifactStore, error) {
	out := new(ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreStackController_Apply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtifactStoreStackControllerServer is the server API for ArtifactStoreStackController service.
// All implementations should embed UnimplementedArtifactStoreStackControllerServer
// for forward compatibility
type ArtifactStoreStackControllerServer interface {
	// preview artifact-store stack
	Preview(context.Context, *ArtifactStore) (*ArtifactStore, error)
	// apply artifact-store stack
	Apply(context.Context, *ArtifactStoreId) (*ArtifactStore, error)
}

// UnimplementedArtifactStoreStackControllerServer should be embedded to have forward compatible implementations.
type UnimplementedArtifactStoreStackControllerServer struct {
}

func (UnimplementedArtifactStoreStackControllerServer) Preview(context.Context, *ArtifactStore) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Preview not implemented")
}
func (UnimplementedArtifactStoreStackControllerServer) Apply(context.Context, *ArtifactStoreId) (*ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}

// UnsafeArtifactStoreStackControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtifactStoreStackControllerServer will
// result in compilation errors.
type UnsafeArtifactStoreStackControllerServer interface {
	mustEmbedUnimplementedArtifactStoreStackControllerServer()
}

func RegisterArtifactStoreStackControllerServer(s grpc.ServiceRegistrar, srv ArtifactStoreStackControllerServer) {
	s.RegisterService(&ArtifactStoreStackController_ServiceDesc, srv)
}

func _ArtifactStoreStackController_Preview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreStackControllerServer).Preview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreStackController_Preview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreStackControllerServer).Preview(ctx, req.(*ArtifactStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreStackController_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactStoreId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactStoreStackControllerServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactStoreStackController_Apply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactStoreStackControllerServer).Apply(ctx, req.(*ArtifactStoreId))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtifactStoreStackController_ServiceDesc is the grpc.ServiceDesc for ArtifactStoreStackController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtifactStoreStackController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreStackController",
	HandlerType: (*ArtifactStoreStackControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "preview",
			Handler:    _ArtifactStoreStackController_Preview_Handler,
		},
		{
			MethodName: "apply",
			Handler:    _ArtifactStoreStackController_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/develop/artifactstore/rpc/service.proto",
}
