// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/search/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/search/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SearchQueryController_SearchByText_FullMethodName                 = "/cloud.planton.apis.v1.search.service.SearchQueryController/searchByText"
	SearchQueryController_SearchByResourceType_FullMethodName         = "/cloud.planton.apis.v1.search.service.SearchQueryController/searchByResourceType"
	SearchQueryController_SearchIdentityAccountByEmail_FullMethodName = "/cloud.planton.apis.v1.search.service.SearchQueryController/searchIdentityAccountByEmail"
)

// SearchQueryControllerClient is the client API for SearchQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchQueryControllerClient interface {
	// SearchByText is an RPC method that takes a SearchByTextInput message
	// containing the company identifier, product identifier, and search text.
	// This method is responsible for performing a text-based search query
	// related to the specified company and product, and it returns a response
	// containing the search results.
	SearchByText(ctx context.Context, in *model.SearchByTextInput, opts ...grpc.CallOption) (*model.ResourceList, error)
	// This method returns a `ResourceList` message, which encapsulates a list of resources that match
	// the input search parameters. Each resource in the list should match the specified resource type,
	// and be associated with the specified company and product.
	SearchByResourceType(ctx context.Context, in *model.SearchByResourceTypeInput, opts ...grpc.CallOption) (*model.ResourceList, error)
	// This method returns a `ResourceList` message, which encapsulates a list of identities that match
	// the input search parameters. Each identity in the list should be associated with the specified company
	// and match the specified email or part thereof.
	SearchIdentityAccountByEmail(ctx context.Context, in *model.SearchIdentityAccountByEmailInput, opts ...grpc.CallOption) (*model.ResourceList, error)
}

type searchQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchQueryControllerClient(cc grpc.ClientConnInterface) SearchQueryControllerClient {
	return &searchQueryControllerClient{cc}
}

func (c *searchQueryControllerClient) SearchByText(ctx context.Context, in *model.SearchByTextInput, opts ...grpc.CallOption) (*model.ResourceList, error) {
	out := new(model.ResourceList)
	err := c.cc.Invoke(ctx, SearchQueryController_SearchByText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchQueryControllerClient) SearchByResourceType(ctx context.Context, in *model.SearchByResourceTypeInput, opts ...grpc.CallOption) (*model.ResourceList, error) {
	out := new(model.ResourceList)
	err := c.cc.Invoke(ctx, SearchQueryController_SearchByResourceType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchQueryControllerClient) SearchIdentityAccountByEmail(ctx context.Context, in *model.SearchIdentityAccountByEmailInput, opts ...grpc.CallOption) (*model.ResourceList, error) {
	out := new(model.ResourceList)
	err := c.cc.Invoke(ctx, SearchQueryController_SearchIdentityAccountByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchQueryControllerServer is the server API for SearchQueryController service.
// All implementations should embed UnimplementedSearchQueryControllerServer
// for forward compatibility
type SearchQueryControllerServer interface {
	// SearchByText is an RPC method that takes a SearchByTextInput message
	// containing the company identifier, product identifier, and search text.
	// This method is responsible for performing a text-based search query
	// related to the specified company and product, and it returns a response
	// containing the search results.
	SearchByText(context.Context, *model.SearchByTextInput) (*model.ResourceList, error)
	// This method returns a `ResourceList` message, which encapsulates a list of resources that match
	// the input search parameters. Each resource in the list should match the specified resource type,
	// and be associated with the specified company and product.
	SearchByResourceType(context.Context, *model.SearchByResourceTypeInput) (*model.ResourceList, error)
	// This method returns a `ResourceList` message, which encapsulates a list of identities that match
	// the input search parameters. Each identity in the list should be associated with the specified company
	// and match the specified email or part thereof.
	SearchIdentityAccountByEmail(context.Context, *model.SearchIdentityAccountByEmailInput) (*model.ResourceList, error)
}

// UnimplementedSearchQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedSearchQueryControllerServer struct {
}

func (UnimplementedSearchQueryControllerServer) SearchByText(context.Context, *model.SearchByTextInput) (*model.ResourceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchByText not implemented")
}
func (UnimplementedSearchQueryControllerServer) SearchByResourceType(context.Context, *model.SearchByResourceTypeInput) (*model.ResourceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchByResourceType not implemented")
}
func (UnimplementedSearchQueryControllerServer) SearchIdentityAccountByEmail(context.Context, *model.SearchIdentityAccountByEmailInput) (*model.ResourceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchIdentityAccountByEmail not implemented")
}

// UnsafeSearchQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchQueryControllerServer will
// result in compilation errors.
type UnsafeSearchQueryControllerServer interface {
	mustEmbedUnimplementedSearchQueryControllerServer()
}

func RegisterSearchQueryControllerServer(s grpc.ServiceRegistrar, srv SearchQueryControllerServer) {
	s.RegisterService(&SearchQueryController_ServiceDesc, srv)
}

func _SearchQueryController_SearchByText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.SearchByTextInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchQueryControllerServer).SearchByText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchQueryController_SearchByText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchQueryControllerServer).SearchByText(ctx, req.(*model.SearchByTextInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchQueryController_SearchByResourceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.SearchByResourceTypeInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchQueryControllerServer).SearchByResourceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchQueryController_SearchByResourceType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchQueryControllerServer).SearchByResourceType(ctx, req.(*model.SearchByResourceTypeInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchQueryController_SearchIdentityAccountByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.SearchIdentityAccountByEmailInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchQueryControllerServer).SearchIdentityAccountByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchQueryController_SearchIdentityAccountByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchQueryControllerServer).SearchIdentityAccountByEmail(ctx, req.(*model.SearchIdentityAccountByEmailInput))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchQueryController_ServiceDesc is the grpc.ServiceDesc for SearchQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.search.service.SearchQueryController",
	HandlerType: (*SearchQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "searchByText",
			Handler:    _SearchQueryController_SearchByText_Handler,
		},
		{
			MethodName: "searchByResourceType",
			Handler:    _SearchQueryController_SearchByResourceType_Handler,
		},
		{
			MethodName: "searchIdentityAccountByEmail",
			Handler:    _SearchQueryController_SearchIdentityAccountByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/search/service/query.proto",
}