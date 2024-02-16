// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/resourcemanager/product/query.proto

package product

import (
	context "context"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	company "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/company"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProductQueryController_List_FullMethodName            = "/cloud.planton.apis.v1.resourcemanager.product.ProductQueryController/list"
	ProductQueryController_GetById_FullMethodName         = "/cloud.planton.apis.v1.resourcemanager.product.ProductQueryController/getById"
	ProductQueryController_FindByCompanyId_FullMethodName = "/cloud.planton.apis.v1.resourcemanager.product.ProductQueryController/findByCompanyId"
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
	FindByCompanyId(ctx context.Context, in *company.CompanyId, opts ...grpc.CallOption) (*Products, error)
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

func (c *productQueryControllerClient) FindByCompanyId(ctx context.Context, in *company.CompanyId, opts ...grpc.CallOption) (*Products, error) {
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
	FindByCompanyId(context.Context, *company.CompanyId) (*Products, error)
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
func (UnimplementedProductQueryControllerServer) FindByCompanyId(context.Context, *company.CompanyId) (*Products, error) {
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
	in := new(company.CompanyId)
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
		return srv.(ProductQueryControllerServer).FindByCompanyId(ctx, req.(*company.CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductQueryController_ServiceDesc is the grpc.ServiceDesc for ProductQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.resourcemanager.product.ProductQueryController",
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
	Metadata: "cloud/planton/apis/v1/resourcemanager/product/query.proto",
}
