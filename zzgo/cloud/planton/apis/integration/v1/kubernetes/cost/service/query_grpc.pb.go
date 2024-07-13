// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/integration/v1/kubernetes/cost/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/cost/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CostAllocationQueryController_ListByCostAllocationFilters_FullMethodName  = "/cloud.planton.apis.integration.v1.kubernetes.cost.service.CostAllocationQueryController/listByCostAllocationFilters"
	CostAllocationQueryController_GetCostAggregateByResourceId_FullMethodName = "/cloud.planton.apis.integration.v1.kubernetes.cost.service.CostAllocationQueryController/getCostAggregateByResourceId"
)

// CostAllocationQueryControllerClient is the client API for CostAllocationQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CostAllocationQueryControllerClient interface {
	// Get Cost allocation data by search filters
	// This returns a paginated list of cost allocation data
	// Filters include start and end date so company identifier, environment identifier and posting environment identifier
	ListByCostAllocationFilters(ctx context.Context, in *model.ListByCostAllocationFiltersInput, opts ...grpc.CallOption) (*model.CostAllocationList, error)
	// get cost aggregate by timestamp for a given resource.
	// example: Get month to date cost of a postgres cluster
	// example: Get month to date cost of a kafka cluster
	GetCostAggregateByResourceId(ctx context.Context, in *model.GetCostAggregateInput, opts ...grpc.CallOption) (*model.CostAggregate, error)
}

type costAllocationQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCostAllocationQueryControllerClient(cc grpc.ClientConnInterface) CostAllocationQueryControllerClient {
	return &costAllocationQueryControllerClient{cc}
}

func (c *costAllocationQueryControllerClient) ListByCostAllocationFilters(ctx context.Context, in *model.ListByCostAllocationFiltersInput, opts ...grpc.CallOption) (*model.CostAllocationList, error) {
	out := new(model.CostAllocationList)
	err := c.cc.Invoke(ctx, CostAllocationQueryController_ListByCostAllocationFilters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costAllocationQueryControllerClient) GetCostAggregateByResourceId(ctx context.Context, in *model.GetCostAggregateInput, opts ...grpc.CallOption) (*model.CostAggregate, error) {
	out := new(model.CostAggregate)
	err := c.cc.Invoke(ctx, CostAllocationQueryController_GetCostAggregateByResourceId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CostAllocationQueryControllerServer is the server API for CostAllocationQueryController service.
// All implementations should embed UnimplementedCostAllocationQueryControllerServer
// for forward compatibility
type CostAllocationQueryControllerServer interface {
	// Get Cost allocation data by search filters
	// This returns a paginated list of cost allocation data
	// Filters include start and end date so company identifier, environment identifier and posting environment identifier
	ListByCostAllocationFilters(context.Context, *model.ListByCostAllocationFiltersInput) (*model.CostAllocationList, error)
	// get cost aggregate by timestamp for a given resource.
	// example: Get month to date cost of a postgres cluster
	// example: Get month to date cost of a kafka cluster
	GetCostAggregateByResourceId(context.Context, *model.GetCostAggregateInput) (*model.CostAggregate, error)
}

// UnimplementedCostAllocationQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCostAllocationQueryControllerServer struct {
}

func (UnimplementedCostAllocationQueryControllerServer) ListByCostAllocationFilters(context.Context, *model.ListByCostAllocationFiltersInput) (*model.CostAllocationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByCostAllocationFilters not implemented")
}
func (UnimplementedCostAllocationQueryControllerServer) GetCostAggregateByResourceId(context.Context, *model.GetCostAggregateInput) (*model.CostAggregate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCostAggregateByResourceId not implemented")
}

// UnsafeCostAllocationQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CostAllocationQueryControllerServer will
// result in compilation errors.
type UnsafeCostAllocationQueryControllerServer interface {
	mustEmbedUnimplementedCostAllocationQueryControllerServer()
}

func RegisterCostAllocationQueryControllerServer(s grpc.ServiceRegistrar, srv CostAllocationQueryControllerServer) {
	s.RegisterService(&CostAllocationQueryController_ServiceDesc, srv)
}

func _CostAllocationQueryController_ListByCostAllocationFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ListByCostAllocationFiltersInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostAllocationQueryControllerServer).ListByCostAllocationFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostAllocationQueryController_ListByCostAllocationFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostAllocationQueryControllerServer).ListByCostAllocationFilters(ctx, req.(*model.ListByCostAllocationFiltersInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CostAllocationQueryController_GetCostAggregateByResourceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetCostAggregateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostAllocationQueryControllerServer).GetCostAggregateByResourceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostAllocationQueryController_GetCostAggregateByResourceId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostAllocationQueryControllerServer).GetCostAggregateByResourceId(ctx, req.(*model.GetCostAggregateInput))
	}
	return interceptor(ctx, in, info, handler)
}

// CostAllocationQueryController_ServiceDesc is the grpc.ServiceDesc for CostAllocationQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CostAllocationQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.integration.v1.kubernetes.cost.service.CostAllocationQueryController",
	HandlerType: (*CostAllocationQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listByCostAllocationFilters",
			Handler:    _CostAllocationQueryController_ListByCostAllocationFilters_Handler,
		},
		{
			MethodName: "getCostAggregateByResourceId",
			Handler:    _CostAllocationQueryController_GetCostAggregateByResourceId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/integration/v1/kubernetes/cost/service/query.proto",
}
