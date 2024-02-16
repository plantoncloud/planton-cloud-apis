// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/develop/artifactstore/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/develop/artifactstore/model"
	model3 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/develop/sourcecode/project/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/product/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ArtifactStoreQueryController_GetById_FullMethodName                          = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreQueryController/getById"
	ArtifactStoreQueryController_List_FullMethodName                             = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreQueryController/list"
	ArtifactStoreQueryController_FindByProductId_FullMethodName                  = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreQueryController/findByProductId"
	ArtifactStoreQueryController_FindByCodeProjectUrl_FullMethodName             = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreQueryController/findByCodeProjectUrl"
	ArtifactStoreQueryController_ListArtifactStoreDockerImages_FullMethodName    = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreQueryController/listArtifactStoreDockerImages"
	ArtifactStoreQueryController_ListArtifactStorePackages_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreQueryController/listArtifactStorePackages"
	ArtifactStoreQueryController_ListArtifactStorePackageVersions_FullMethodName = "/cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreQueryController/listArtifactStorePackageVersions"
)

// ArtifactStoreQueryControllerClient is the client API for ArtifactStoreQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtifactStoreQueryControllerClient interface {
	// get artifact-store by id
	GetById(ctx context.Context, in *model.ArtifactStoreId, opts ...grpc.CallOption) (*model.ArtifactStore, error)
	// list all artifact-stores on planton cloud for the requested page. This is intended for use on portal.
	List(ctx context.Context, in *model1.PageInfo, opts ...grpc.CallOption) (*model.ArtifactStoreList, error)
	// look up artifact-stores by product id.
	FindByProductId(ctx context.Context, in *model2.ProductId, opts ...grpc.CallOption) (*model.ArtifactStores, error)
	// look up artifact-stores by code project url
	FindByCodeProjectUrl(ctx context.Context, in *model3.CodeProjectUrl, opts ...grpc.CallOption) (*model.ArtifactStores, error)
	// list docker images from the artifact-store provided in the input
	// (proxy google artifact-registry server)
	ListArtifactStoreDockerImages(ctx context.Context, in *model.ListByArtifactStoreIdRepoNameInput, opts ...grpc.CallOption) (*model.ArtifactStoreDockerImageList, error)
	// list maven/npm/python packages from the corresponding repositories of artifact-store provided in the input
	// (proxy google artifact-registry server)
	ListArtifactStorePackages(ctx context.Context, in *model.ListByArtifactStoreIdRepoNameInput, opts ...grpc.CallOption) (*model.ArtifactStorePackageList, error)
	// list maven/npm/python package versions from the corresponding repositories of artifact-store provided in the input
	// (proxy google artifact-registry server)
	ListArtifactStorePackageVersions(ctx context.Context, in *model.ListByArtifactStoreIdPackageNameInput, opts ...grpc.CallOption) (*model.ArtifactStorePackageVersionList, error)
}

type artifactStoreQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewArtifactStoreQueryControllerClient(cc grpc.ClientConnInterface) ArtifactStoreQueryControllerClient {
	return &artifactStoreQueryControllerClient{cc}
}

func (c *artifactStoreQueryControllerClient) GetById(ctx context.Context, in *model.ArtifactStoreId, opts ...grpc.CallOption) (*model.ArtifactStore, error) {
	out := new(model.ArtifactStore)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) List(ctx context.Context, in *model1.PageInfo, opts ...grpc.CallOption) (*model.ArtifactStoreList, error) {
	out := new(model.ArtifactStoreList)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) FindByProductId(ctx context.Context, in *model2.ProductId, opts ...grpc.CallOption) (*model.ArtifactStores, error) {
	out := new(model.ArtifactStores)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) FindByCodeProjectUrl(ctx context.Context, in *model3.CodeProjectUrl, opts ...grpc.CallOption) (*model.ArtifactStores, error) {
	out := new(model.ArtifactStores)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_FindByCodeProjectUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) ListArtifactStoreDockerImages(ctx context.Context, in *model.ListByArtifactStoreIdRepoNameInput, opts ...grpc.CallOption) (*model.ArtifactStoreDockerImageList, error) {
	out := new(model.ArtifactStoreDockerImageList)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_ListArtifactStoreDockerImages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) ListArtifactStorePackages(ctx context.Context, in *model.ListByArtifactStoreIdRepoNameInput, opts ...grpc.CallOption) (*model.ArtifactStorePackageList, error) {
	out := new(model.ArtifactStorePackageList)
	err := c.cc.Invoke(ctx, ArtifactStoreQueryController_ListArtifactStorePackages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactStoreQueryControllerClient) ListArtifactStorePackageVersions(ctx context.Context, in *model.ListByArtifactStoreIdPackageNameInput, opts ...grpc.CallOption) (*model.ArtifactStorePackageVersionList, error) {
	out := new(model.ArtifactStorePackageVersionList)
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
	GetById(context.Context, *model.ArtifactStoreId) (*model.ArtifactStore, error)
	// list all artifact-stores on planton cloud for the requested page. This is intended for use on portal.
	List(context.Context, *model1.PageInfo) (*model.ArtifactStoreList, error)
	// look up artifact-stores by product id.
	FindByProductId(context.Context, *model2.ProductId) (*model.ArtifactStores, error)
	// look up artifact-stores by code project url
	FindByCodeProjectUrl(context.Context, *model3.CodeProjectUrl) (*model.ArtifactStores, error)
	// list docker images from the artifact-store provided in the input
	// (proxy google artifact-registry server)
	ListArtifactStoreDockerImages(context.Context, *model.ListByArtifactStoreIdRepoNameInput) (*model.ArtifactStoreDockerImageList, error)
	// list maven/npm/python packages from the corresponding repositories of artifact-store provided in the input
	// (proxy google artifact-registry server)
	ListArtifactStorePackages(context.Context, *model.ListByArtifactStoreIdRepoNameInput) (*model.ArtifactStorePackageList, error)
	// list maven/npm/python package versions from the corresponding repositories of artifact-store provided in the input
	// (proxy google artifact-registry server)
	ListArtifactStorePackageVersions(context.Context, *model.ListByArtifactStoreIdPackageNameInput) (*model.ArtifactStorePackageVersionList, error)
}

// UnimplementedArtifactStoreQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedArtifactStoreQueryControllerServer struct {
}

func (UnimplementedArtifactStoreQueryControllerServer) GetById(context.Context, *model.ArtifactStoreId) (*model.ArtifactStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) List(context.Context, *model1.PageInfo) (*model.ArtifactStoreList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) FindByProductId(context.Context, *model2.ProductId) (*model.ArtifactStores, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) FindByCodeProjectUrl(context.Context, *model3.CodeProjectUrl) (*model.ArtifactStores, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCodeProjectUrl not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) ListArtifactStoreDockerImages(context.Context, *model.ListByArtifactStoreIdRepoNameInput) (*model.ArtifactStoreDockerImageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtifactStoreDockerImages not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) ListArtifactStorePackages(context.Context, *model.ListByArtifactStoreIdRepoNameInput) (*model.ArtifactStorePackageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtifactStorePackages not implemented")
}
func (UnimplementedArtifactStoreQueryControllerServer) ListArtifactStorePackageVersions(context.Context, *model.ListByArtifactStoreIdPackageNameInput) (*model.ArtifactStorePackageVersionList, error) {
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
	in := new(model.ArtifactStoreId)
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
		return srv.(ArtifactStoreQueryControllerServer).GetById(ctx, req.(*model.ArtifactStoreId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.PageInfo)
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
		return srv.(ArtifactStoreQueryControllerServer).List(ctx, req.(*model1.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.ProductId)
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
		return srv.(ArtifactStoreQueryControllerServer).FindByProductId(ctx, req.(*model2.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_FindByCodeProjectUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model3.CodeProjectUrl)
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
		return srv.(ArtifactStoreQueryControllerServer).FindByCodeProjectUrl(ctx, req.(*model3.CodeProjectUrl))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_ListArtifactStoreDockerImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ListByArtifactStoreIdRepoNameInput)
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
		return srv.(ArtifactStoreQueryControllerServer).ListArtifactStoreDockerImages(ctx, req.(*model.ListByArtifactStoreIdRepoNameInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_ListArtifactStorePackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ListByArtifactStoreIdRepoNameInput)
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
		return srv.(ArtifactStoreQueryControllerServer).ListArtifactStorePackages(ctx, req.(*model.ListByArtifactStoreIdRepoNameInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactStoreQueryController_ListArtifactStorePackageVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ListByArtifactStoreIdPackageNameInput)
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
		return srv.(ArtifactStoreQueryControllerServer).ListArtifactStorePackageVersions(ctx, req.(*model.ListByArtifactStoreIdPackageNameInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtifactStoreQueryController_ServiceDesc is the grpc.ServiceDesc for ArtifactStoreQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtifactStoreQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.develop.artifactstore.service.ArtifactStoreQueryController",
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
	Metadata: "cloud/planton/apis/v1/code2cloud/develop/artifactstore/service/query.proto",
}
