// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/resourcemanager/product/rpc/service.proto

package rpc

import (
	context "context"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/company/rpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProductCommandController_Create_FullMethodName  = "/cloud.planton.apis.v1.resourcemanager.product.rpc.ProductCommandController/create"
	ProductCommandController_Update_FullMethodName  = "/cloud.planton.apis.v1.resourcemanager.product.rpc.ProductCommandController/update"
	ProductCommandController_Delete_FullMethodName  = "/cloud.planton.apis.v1.resourcemanager.product.rpc.ProductCommandController/delete"
	ProductCommandController_Restore_FullMethodName = "/cloud.planton.apis.v1.resourcemanager.product.rpc.ProductCommandController/restore"
)

// ProductCommandControllerClient is the client API for ProductCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCommandControllerClient interface {
	// add a new product to a company
	Create(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error)
	// update an existing product
	Update(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error)
	// delete an existing product
	Delete(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*Product, error)
	// restore a previously deleted product
	Restore(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error)
}

type productCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCommandControllerClient(cc grpc.ClientConnInterface) ProductCommandControllerClient {
	return &productCommandControllerClient{cc}
}

func (c *productCommandControllerClient) Create(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, ProductCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCommandControllerClient) Update(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, ProductCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCommandControllerClient) Delete(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, ProductCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCommandControllerClient) Restore(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, ProductCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCommandControllerServer is the server API for ProductCommandController service.
// All implementations should embed UnimplementedProductCommandControllerServer
// for forward compatibility
type ProductCommandControllerServer interface {
	// add a new product to a company
	Create(context.Context, *Product) (*Product, error)
	// update an existing product
	Update(context.Context, *Product) (*Product, error)
	// delete an existing product
	Delete(context.Context, *ProductId) (*Product, error)
	// restore a previously deleted product
	Restore(context.Context, *Product) (*Product, error)
}

// UnimplementedProductCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedProductCommandControllerServer struct {
}

func (UnimplementedProductCommandControllerServer) Create(context.Context, *Product) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProductCommandControllerServer) Update(context.Context, *Product) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProductCommandControllerServer) Delete(context.Context, *ProductId) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProductCommandControllerServer) Restore(context.Context, *Product) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeProductCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCommandControllerServer will
// result in compilation errors.
type UnsafeProductCommandControllerServer interface {
	mustEmbedUnimplementedProductCommandControllerServer()
}

func RegisterProductCommandControllerServer(s grpc.ServiceRegistrar, srv ProductCommandControllerServer) {
	s.RegisterService(&ProductCommandController_ServiceDesc, srv)
}

func _ProductCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCommandControllerServer).Create(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCommandControllerServer).Update(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCommandControllerServer).Delete(ctx, req.(*ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCommandControllerServer).Restore(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCommandController_ServiceDesc is the grpc.ServiceDesc for ProductCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.resourcemanager.product.rpc.ProductCommandController",
	HandlerType: (*ProductCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _ProductCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _ProductCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _ProductCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _ProductCommandController_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/resourcemanager/product/rpc/service.proto",
}

const (
	ProductQueryController_List_FullMethodName            = "/cloud.planton.apis.v1.resourcemanager.product.rpc.ProductQueryController/list"
	ProductQueryController_GetById_FullMethodName         = "/cloud.planton.apis.v1.resourcemanager.product.rpc.ProductQueryController/getById"
	ProductQueryController_FindByCompanyId_FullMethodName = "/cloud.planton.apis.v1.resourcemanager.product.rpc.ProductQueryController/findByCompanyId"
)

// ProductQueryControllerClient is the client API for ProductQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductQueryControllerClient interface {
	// list all products on planton cloud for the requested page. This is intended to be used on back-office portal.
	List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*ProductList, error)
	// get details of a product by product id
	GetById(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*Product, error)
	// find products by company id.
	// the response should only include the products that the authenticated user has access to.
	FindByCompanyId(ctx context.Context, in *rpc.CompanyId, opts ...grpc.CallOption) (*Products, error)
}

type productQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewProductQueryControllerClient(cc grpc.ClientConnInterface) ProductQueryControllerClient {
	return &productQueryControllerClient{cc}
}

func (c *productQueryControllerClient) List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := c.cc.Invoke(ctx, ProductQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productQueryControllerClient) GetById(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, ProductQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productQueryControllerClient) FindByCompanyId(ctx context.Context, in *rpc.CompanyId, opts ...grpc.CallOption) (*Products, error) {
	out := new(Products)
	err := c.cc.Invoke(ctx, ProductQueryController_FindByCompanyId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductQueryControllerServer is the server API for ProductQueryController service.
// All implementations should embed UnimplementedProductQueryControllerServer
// for forward compatibility
type ProductQueryControllerServer interface {
	// list all products on planton cloud for the requested page. This is intended to be used on back-office portal.
	List(context.Context, *pagination.PageInfo) (*ProductList, error)
	// get details of a product by product id
	GetById(context.Context, *ProductId) (*Product, error)
	// find products by company id.
	// the response should only include the products that the authenticated user has access to.
	FindByCompanyId(context.Context, *rpc.CompanyId) (*Products, error)
}

// UnimplementedProductQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedProductQueryControllerServer struct {
}

func (UnimplementedProductQueryControllerServer) List(context.Context, *pagination.PageInfo) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProductQueryControllerServer) GetById(context.Context, *ProductId) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedProductQueryControllerServer) FindByCompanyId(context.Context, *rpc.CompanyId) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCompanyId not implemented")
}

// UnsafeProductQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductQueryControllerServer will
// result in compilation errors.
type UnsafeProductQueryControllerServer interface {
	mustEmbedUnimplementedProductQueryControllerServer()
}

func RegisterProductQueryControllerServer(s grpc.ServiceRegistrar, srv ProductQueryControllerServer) {
	s.RegisterService(&ProductQueryController_ServiceDesc, srv)
}

func _ProductQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductQueryControllerServer).List(ctx, req.(*pagination.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductQueryControllerServer).GetById(ctx, req.(*ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductQueryController_FindByCompanyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.CompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductQueryControllerServer).FindByCompanyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductQueryController_FindByCompanyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductQueryControllerServer).FindByCompanyId(ctx, req.(*rpc.CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductQueryController_ServiceDesc is the grpc.ServiceDesc for ProductQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.resourcemanager.product.rpc.ProductQueryController",
	HandlerType: (*ProductQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _ProductQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _ProductQueryController_GetById_Handler,
		},
		{
			MethodName: "findByCompanyId",
			Handler:    _ProductQueryController_FindByCompanyId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/resourcemanager/product/rpc/service.proto",
}
