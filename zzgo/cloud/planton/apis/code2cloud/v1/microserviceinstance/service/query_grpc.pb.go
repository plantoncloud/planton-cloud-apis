// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/microserviceinstance/service/query.proto

package service

import (
	context "context"
	model5 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/codeproject/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/environment/model"
	model3 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/microserviceinstance/model"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/rpc"
	model4 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/apiresources/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/resourcemanager/v1/product/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MicroserviceInstanceQueryController_List_FullMethodName                               = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.service.MicroserviceInstanceQueryController/list"
	MicroserviceInstanceQueryController_GetById_FullMethodName                            = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.service.MicroserviceInstanceQueryController/getById"
	MicroserviceInstanceQueryController_FindByProductId_FullMethodName                    = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.service.MicroserviceInstanceQueryController/findByProductId"
	MicroserviceInstanceQueryController_FindByEnvironmentId_FullMethodName                = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.service.MicroserviceInstanceQueryController/findByEnvironmentId"
	MicroserviceInstanceQueryController_FindByKubeClusterId_FullMethodName                = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.service.MicroserviceInstanceQueryController/findByKubeClusterId"
	MicroserviceInstanceQueryController_FindByEnvironmentIdByCodeProjectId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.service.MicroserviceInstanceQueryController/findByEnvironmentIdByCodeProjectId"
	MicroserviceInstanceQueryController_FindPods_FullMethodName                           = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.service.MicroserviceInstanceQueryController/findPods"
	MicroserviceInstanceQueryController_GetLogStream_FullMethodName                       = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.service.MicroserviceInstanceQueryController/getLogStream"
	MicroserviceInstanceQueryController_GetByCodeProjectId_FullMethodName                 = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.service.MicroserviceInstanceQueryController/getByCodeProjectId"
	MicroserviceInstanceQueryController_FindByCodeProjectUrl_FullMethodName               = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.service.MicroserviceInstanceQueryController/findByCodeProjectUrl"
	MicroserviceInstanceQueryController_GetEnvVarMap_FullMethodName                       = "/cloud.planton.apis.code2cloud.v1.microserviceinstance.service.MicroserviceInstanceQueryController/getEnvVarMap"
)

// MicroserviceInstanceQueryControllerClient is the client API for MicroserviceInstanceQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicroserviceInstanceQueryControllerClient interface {
	// list all microservice-instances on planton instance for the requested page. This is intended for use on portal.
	List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.MicroserviceInstanceList, error)
	// look up microservice-instance using microservice-instance id
	GetById(ctx context.Context, in *model.MicroserviceInstanceId, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// find microservice-instances by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.MicroserviceInstances, error)
	// find microservice-instances by environment
	FindByEnvironmentId(ctx context.Context, in *model2.EnvironmentId, opts ...grpc.CallOption) (*model.MicroserviceInstances, error)
	// find microservice-instances by kube-cluster
	FindByKubeClusterId(ctx context.Context, in *model3.KubeClusterId, opts ...grpc.CallOption) (*model.MicroserviceInstances, error)
	// find microservice-instances in a environment for a code-project
	FindByEnvironmentIdByCodeProjectId(ctx context.Context, in *model.ByEnvironmentIdByCodeProjectIdInput, opts ...grpc.CallOption) (*model.MicroserviceInstances, error)
	// lookup pods of a microservice-instance deployed to a environment
	FindPods(ctx context.Context, in *model.MicroserviceInstanceId, opts ...grpc.CallOption) (*model4.Pods, error)
	// get a log stream for a running instance of a microservice-instance
	GetLogStream(ctx context.Context, in *model.GetMicroserviceInstanceLogStreamQueryInput, opts ...grpc.CallOption) (MicroserviceInstanceQueryController_GetLogStreamClient, error)
	// lookup a microservice-instance by code project id
	GetByCodeProjectId(ctx context.Context, in *model5.CodeProjectId, opts ...grpc.CallOption) (*model.MicroserviceInstance, error)
	// lookup all microservice-instances by code project url
	FindByCodeProjectUrl(ctx context.Context, in *model5.CodeProjectUrl, opts ...grpc.CallOption) (*model.MicroserviceInstances, error)
	GetEnvVarMap(ctx context.Context, in *model.GetMicroserviceInstanceEnvVarMapInput, opts ...grpc.CallOption) (*model.MicroserviceInstanceEnvVarMap, error)
}

type microserviceInstanceQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroserviceInstanceQueryControllerClient(cc grpc.ClientConnInterface) MicroserviceInstanceQueryControllerClient {
	return &microserviceInstanceQueryControllerClient{cc}
}

func (c *microserviceInstanceQueryControllerClient) List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.MicroserviceInstanceList, error) {
	out := new(model.MicroserviceInstanceList)
	err := c.cc.Invoke(ctx, MicroserviceInstanceQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceQueryControllerClient) GetById(ctx context.Context, in *model.MicroserviceInstanceId, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceQueryControllerClient) FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.MicroserviceInstances, error) {
	out := new(model.MicroserviceInstances)
	err := c.cc.Invoke(ctx, MicroserviceInstanceQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceQueryControllerClient) FindByEnvironmentId(ctx context.Context, in *model2.EnvironmentId, opts ...grpc.CallOption) (*model.MicroserviceInstances, error) {
	out := new(model.MicroserviceInstances)
	err := c.cc.Invoke(ctx, MicroserviceInstanceQueryController_FindByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceQueryControllerClient) FindByKubeClusterId(ctx context.Context, in *model3.KubeClusterId, opts ...grpc.CallOption) (*model.MicroserviceInstances, error) {
	out := new(model.MicroserviceInstances)
	err := c.cc.Invoke(ctx, MicroserviceInstanceQueryController_FindByKubeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceQueryControllerClient) FindByEnvironmentIdByCodeProjectId(ctx context.Context, in *model.ByEnvironmentIdByCodeProjectIdInput, opts ...grpc.CallOption) (*model.MicroserviceInstances, error) {
	out := new(model.MicroserviceInstances)
	err := c.cc.Invoke(ctx, MicroserviceInstanceQueryController_FindByEnvironmentIdByCodeProjectId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceQueryControllerClient) FindPods(ctx context.Context, in *model.MicroserviceInstanceId, opts ...grpc.CallOption) (*model4.Pods, error) {
	out := new(model4.Pods)
	err := c.cc.Invoke(ctx, MicroserviceInstanceQueryController_FindPods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceQueryControllerClient) GetLogStream(ctx context.Context, in *model.GetMicroserviceInstanceLogStreamQueryInput, opts ...grpc.CallOption) (MicroserviceInstanceQueryController_GetLogStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &MicroserviceInstanceQueryController_ServiceDesc.Streams[0], MicroserviceInstanceQueryController_GetLogStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &microserviceInstanceQueryControllerGetLogStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MicroserviceInstanceQueryController_GetLogStreamClient interface {
	Recv() (*wrapperspb.StringValue, error)
	grpc.ClientStream
}

type microserviceInstanceQueryControllerGetLogStreamClient struct {
	grpc.ClientStream
}

func (x *microserviceInstanceQueryControllerGetLogStreamClient) Recv() (*wrapperspb.StringValue, error) {
	m := new(wrapperspb.StringValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *microserviceInstanceQueryControllerClient) GetByCodeProjectId(ctx context.Context, in *model5.CodeProjectId, opts ...grpc.CallOption) (*model.MicroserviceInstance, error) {
	out := new(model.MicroserviceInstance)
	err := c.cc.Invoke(ctx, MicroserviceInstanceQueryController_GetByCodeProjectId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceQueryControllerClient) FindByCodeProjectUrl(ctx context.Context, in *model5.CodeProjectUrl, opts ...grpc.CallOption) (*model.MicroserviceInstances, error) {
	out := new(model.MicroserviceInstances)
	err := c.cc.Invoke(ctx, MicroserviceInstanceQueryController_FindByCodeProjectUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microserviceInstanceQueryControllerClient) GetEnvVarMap(ctx context.Context, in *model.GetMicroserviceInstanceEnvVarMapInput, opts ...grpc.CallOption) (*model.MicroserviceInstanceEnvVarMap, error) {
	out := new(model.MicroserviceInstanceEnvVarMap)
	err := c.cc.Invoke(ctx, MicroserviceInstanceQueryController_GetEnvVarMap_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroserviceInstanceQueryControllerServer is the server API for MicroserviceInstanceQueryController service.
// All implementations should embed UnimplementedMicroserviceInstanceQueryControllerServer
// for forward compatibility
type MicroserviceInstanceQueryControllerServer interface {
	// list all microservice-instances on planton instance for the requested page. This is intended for use on portal.
	List(context.Context, *rpc.PageInfo) (*model.MicroserviceInstanceList, error)
	// look up microservice-instance using microservice-instance id
	GetById(context.Context, *model.MicroserviceInstanceId) (*model.MicroserviceInstance, error)
	// find microservice-instances by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(context.Context, *model1.ProductId) (*model.MicroserviceInstances, error)
	// find microservice-instances by environment
	FindByEnvironmentId(context.Context, *model2.EnvironmentId) (*model.MicroserviceInstances, error)
	// find microservice-instances by kube-cluster
	FindByKubeClusterId(context.Context, *model3.KubeClusterId) (*model.MicroserviceInstances, error)
	// find microservice-instances in a environment for a code-project
	FindByEnvironmentIdByCodeProjectId(context.Context, *model.ByEnvironmentIdByCodeProjectIdInput) (*model.MicroserviceInstances, error)
	// lookup pods of a microservice-instance deployed to a environment
	FindPods(context.Context, *model.MicroserviceInstanceId) (*model4.Pods, error)
	// get a log stream for a running instance of a microservice-instance
	GetLogStream(*model.GetMicroserviceInstanceLogStreamQueryInput, MicroserviceInstanceQueryController_GetLogStreamServer) error
	// lookup a microservice-instance by code project id
	GetByCodeProjectId(context.Context, *model5.CodeProjectId) (*model.MicroserviceInstance, error)
	// lookup all microservice-instances by code project url
	FindByCodeProjectUrl(context.Context, *model5.CodeProjectUrl) (*model.MicroserviceInstances, error)
	GetEnvVarMap(context.Context, *model.GetMicroserviceInstanceEnvVarMapInput) (*model.MicroserviceInstanceEnvVarMap, error)
}

// UnimplementedMicroserviceInstanceQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedMicroserviceInstanceQueryControllerServer struct {
}

func (UnimplementedMicroserviceInstanceQueryControllerServer) List(context.Context, *rpc.PageInfo) (*model.MicroserviceInstanceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMicroserviceInstanceQueryControllerServer) GetById(context.Context, *model.MicroserviceInstanceId) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedMicroserviceInstanceQueryControllerServer) FindByProductId(context.Context, *model1.ProductId) (*model.MicroserviceInstances, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedMicroserviceInstanceQueryControllerServer) FindByEnvironmentId(context.Context, *model2.EnvironmentId) (*model.MicroserviceInstances, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEnvironmentId not implemented")
}
func (UnimplementedMicroserviceInstanceQueryControllerServer) FindByKubeClusterId(context.Context, *model3.KubeClusterId) (*model.MicroserviceInstances, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKubeClusterId not implemented")
}
func (UnimplementedMicroserviceInstanceQueryControllerServer) FindByEnvironmentIdByCodeProjectId(context.Context, *model.ByEnvironmentIdByCodeProjectIdInput) (*model.MicroserviceInstances, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEnvironmentIdByCodeProjectId not implemented")
}
func (UnimplementedMicroserviceInstanceQueryControllerServer) FindPods(context.Context, *model.MicroserviceInstanceId) (*model4.Pods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPods not implemented")
}
func (UnimplementedMicroserviceInstanceQueryControllerServer) GetLogStream(*model.GetMicroserviceInstanceLogStreamQueryInput, MicroserviceInstanceQueryController_GetLogStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLogStream not implemented")
}
func (UnimplementedMicroserviceInstanceQueryControllerServer) GetByCodeProjectId(context.Context, *model5.CodeProjectId) (*model.MicroserviceInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByCodeProjectId not implemented")
}
func (UnimplementedMicroserviceInstanceQueryControllerServer) FindByCodeProjectUrl(context.Context, *model5.CodeProjectUrl) (*model.MicroserviceInstances, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCodeProjectUrl not implemented")
}
func (UnimplementedMicroserviceInstanceQueryControllerServer) GetEnvVarMap(context.Context, *model.GetMicroserviceInstanceEnvVarMapInput) (*model.MicroserviceInstanceEnvVarMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnvVarMap not implemented")
}

// UnsafeMicroserviceInstanceQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroserviceInstanceQueryControllerServer will
// result in compilation errors.
type UnsafeMicroserviceInstanceQueryControllerServer interface {
	mustEmbedUnimplementedMicroserviceInstanceQueryControllerServer()
}

func RegisterMicroserviceInstanceQueryControllerServer(s grpc.ServiceRegistrar, srv MicroserviceInstanceQueryControllerServer) {
	s.RegisterService(&MicroserviceInstanceQueryController_ServiceDesc, srv)
}

func _MicroserviceInstanceQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceQueryControllerServer).List(ctx, req.(*rpc.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MicroserviceInstanceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceQueryControllerServer).GetById(ctx, req.(*model.MicroserviceInstanceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceQueryControllerServer).FindByProductId(ctx, req.(*model1.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceQueryController_FindByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.EnvironmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceQueryControllerServer).FindByEnvironmentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceQueryController_FindByEnvironmentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceQueryControllerServer).FindByEnvironmentId(ctx, req.(*model2.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceQueryController_FindByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model3.KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceQueryControllerServer).FindByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceQueryController_FindByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceQueryControllerServer).FindByKubeClusterId(ctx, req.(*model3.KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceQueryController_FindByEnvironmentIdByCodeProjectId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ByEnvironmentIdByCodeProjectIdInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceQueryControllerServer).FindByEnvironmentIdByCodeProjectId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceQueryController_FindByEnvironmentIdByCodeProjectId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceQueryControllerServer).FindByEnvironmentIdByCodeProjectId(ctx, req.(*model.ByEnvironmentIdByCodeProjectIdInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.MicroserviceInstanceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceQueryControllerServer).FindPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceQueryController_FindPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceQueryControllerServer).FindPods(ctx, req.(*model.MicroserviceInstanceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceQueryController_GetLogStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(model.GetMicroserviceInstanceLogStreamQueryInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MicroserviceInstanceQueryControllerServer).GetLogStream(m, &microserviceInstanceQueryControllerGetLogStreamServer{stream})
}

type MicroserviceInstanceQueryController_GetLogStreamServer interface {
	Send(*wrapperspb.StringValue) error
	grpc.ServerStream
}

type microserviceInstanceQueryControllerGetLogStreamServer struct {
	grpc.ServerStream
}

func (x *microserviceInstanceQueryControllerGetLogStreamServer) Send(m *wrapperspb.StringValue) error {
	return x.ServerStream.SendMsg(m)
}

func _MicroserviceInstanceQueryController_GetByCodeProjectId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model5.CodeProjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceQueryControllerServer).GetByCodeProjectId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceQueryController_GetByCodeProjectId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceQueryControllerServer).GetByCodeProjectId(ctx, req.(*model5.CodeProjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceQueryController_FindByCodeProjectUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model5.CodeProjectUrl)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceQueryControllerServer).FindByCodeProjectUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceQueryController_FindByCodeProjectUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceQueryControllerServer).FindByCodeProjectUrl(ctx, req.(*model5.CodeProjectUrl))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroserviceInstanceQueryController_GetEnvVarMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetMicroserviceInstanceEnvVarMapInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceInstanceQueryControllerServer).GetEnvVarMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroserviceInstanceQueryController_GetEnvVarMap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceInstanceQueryControllerServer).GetEnvVarMap(ctx, req.(*model.GetMicroserviceInstanceEnvVarMapInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MicroserviceInstanceQueryController_ServiceDesc is the grpc.ServiceDesc for MicroserviceInstanceQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MicroserviceInstanceQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.microserviceinstance.service.MicroserviceInstanceQueryController",
	HandlerType: (*MicroserviceInstanceQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _MicroserviceInstanceQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _MicroserviceInstanceQueryController_GetById_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _MicroserviceInstanceQueryController_FindByProductId_Handler,
		},
		{
			MethodName: "findByEnvironmentId",
			Handler:    _MicroserviceInstanceQueryController_FindByEnvironmentId_Handler,
		},
		{
			MethodName: "findByKubeClusterId",
			Handler:    _MicroserviceInstanceQueryController_FindByKubeClusterId_Handler,
		},
		{
			MethodName: "findByEnvironmentIdByCodeProjectId",
			Handler:    _MicroserviceInstanceQueryController_FindByEnvironmentIdByCodeProjectId_Handler,
		},
		{
			MethodName: "findPods",
			Handler:    _MicroserviceInstanceQueryController_FindPods_Handler,
		},
		{
			MethodName: "getByCodeProjectId",
			Handler:    _MicroserviceInstanceQueryController_GetByCodeProjectId_Handler,
		},
		{
			MethodName: "findByCodeProjectUrl",
			Handler:    _MicroserviceInstanceQueryController_FindByCodeProjectUrl_Handler,
		},
		{
			MethodName: "getEnvVarMap",
			Handler:    _MicroserviceInstanceQueryController_GetEnvVarMap_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getLogStream",
			Handler:       _MicroserviceInstanceQueryController_GetLogStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/code2cloud/v1/microserviceinstance/service/query.proto",
}
