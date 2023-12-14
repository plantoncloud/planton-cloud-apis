// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/solr/service.proto

package solr

import (
	context "context"
	kubecluster "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/kubecluster"
	environment "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/environment"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	resource1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/integration/kubernetes/resource"
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
	SolrCloudCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/previewCreate"
	SolrCloudCommandController_Create_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/create"
	SolrCloudCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/previewUpdate"
	SolrCloudCommandController_Update_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/update"
	SolrCloudCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/previewDelete"
	SolrCloudCommandController_Delete_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/delete"
	SolrCloudCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/previewRestore"
	SolrCloudCommandController_Restore_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/restore"
	SolrCloudCommandController_CreateStackJob_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/createStackJob"
	SolrCloudCommandController_Restart_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/restart"
	SolrCloudCommandController_Pause_FullMethodName          = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/pause"
	SolrCloudCommandController_Unpause_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/unpause"
)

// SolrCloudCommandControllerClient is the client API for SolrCloudCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SolrCloudCommandControllerClient interface {
	// preview create solr-cloud
	PreviewCreate(ctx context.Context, in *SolrCloud, opts ...grpc.CallOption) (*SolrCloud, error)
	// create solr-cloud
	Create(ctx context.Context, in *SolrCloud, opts ...grpc.CallOption) (*SolrCloud, error)
	// preview update an existing solr-cloud
	PreviewUpdate(ctx context.Context, in *SolrCloud, opts ...grpc.CallOption) (*SolrCloud, error)
	// update an existing solr-cloud
	Update(ctx context.Context, in *SolrCloud, opts ...grpc.CallOption) (*SolrCloud, error)
	// preview deleting an existing solr-cloud
	PreviewDelete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*SolrCloud, error)
	// delete an existing solr-cloud
	Delete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*SolrCloud, error)
	// preview restoring a deleted solr-cloud
	PreviewRestore(ctx context.Context, in *SolrCloud, opts ...grpc.CallOption) (*SolrCloud, error)
	// restore a deleted solr-cloud
	Restore(ctx context.Context, in *SolrCloud, opts ...grpc.CallOption) (*SolrCloud, error)
	// create-stack-job for solr-cloud
	CreateStackJob(ctx context.Context, in *job.CreateStackJobCommandInput, opts ...grpc.CallOption) (*SolrCloud, error)
	// restart a solr-cloud running in a environment.
	// solr-cloud is restarted by deleting running "solr" pods which will be automatically recreated by kubernetes
	Restart(ctx context.Context, in *SolrCloudId, opts ...grpc.CallOption) (*SolrCloud, error)
	// pause a solr-cloud running in a environment.
	// solr-cloud is paused by scaling down number of replicas of
	// the kubernetes deployment/stateful sets to zero in the environment.
	Pause(ctx context.Context, in *resource.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*SolrCloud, error)
	// unpause a previously paused solr-cloud running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the solr-cloud.
	Unpause(ctx context.Context, in *resource.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*SolrCloud, error)
}

type solrCloudCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewSolrCloudCommandControllerClient(cc grpc.ClientConnInterface) SolrCloudCommandControllerClient {
	return &solrCloudCommandControllerClient{cc}
}

func (c *solrCloudCommandControllerClient) PreviewCreate(ctx context.Context, in *SolrCloud, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudCommandControllerClient) Create(ctx context.Context, in *SolrCloud, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudCommandControllerClient) PreviewUpdate(ctx context.Context, in *SolrCloud, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudCommandControllerClient) Update(ctx context.Context, in *SolrCloud, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudCommandControllerClient) PreviewDelete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudCommandControllerClient) Delete(ctx context.Context, in *resource.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudCommandControllerClient) PreviewRestore(ctx context.Context, in *SolrCloud, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudCommandControllerClient) Restore(ctx context.Context, in *SolrCloud, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudCommandControllerClient) CreateStackJob(ctx context.Context, in *job.CreateStackJobCommandInput, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_CreateStackJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudCommandControllerClient) Restart(ctx context.Context, in *SolrCloudId, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudCommandControllerClient) Pause(ctx context.Context, in *resource.ApiResourcePauseCommandInput, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_Pause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudCommandControllerClient) Unpause(ctx context.Context, in *resource.ApiResourceUnPauseCommandInput, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_Unpause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolrCloudCommandControllerServer is the server API for SolrCloudCommandController service.
// All implementations should embed UnimplementedSolrCloudCommandControllerServer
// for forward compatibility
type SolrCloudCommandControllerServer interface {
	// preview create solr-cloud
	PreviewCreate(context.Context, *SolrCloud) (*SolrCloud, error)
	// create solr-cloud
	Create(context.Context, *SolrCloud) (*SolrCloud, error)
	// preview update an existing solr-cloud
	PreviewUpdate(context.Context, *SolrCloud) (*SolrCloud, error)
	// update an existing solr-cloud
	Update(context.Context, *SolrCloud) (*SolrCloud, error)
	// preview deleting an existing solr-cloud
	PreviewDelete(context.Context, *resource.ApiResourceDeleteCommandInput) (*SolrCloud, error)
	// delete an existing solr-cloud
	Delete(context.Context, *resource.ApiResourceDeleteCommandInput) (*SolrCloud, error)
	// preview restoring a deleted solr-cloud
	PreviewRestore(context.Context, *SolrCloud) (*SolrCloud, error)
	// restore a deleted solr-cloud
	Restore(context.Context, *SolrCloud) (*SolrCloud, error)
	// create-stack-job for solr-cloud
	CreateStackJob(context.Context, *job.CreateStackJobCommandInput) (*SolrCloud, error)
	// restart a solr-cloud running in a environment.
	// solr-cloud is restarted by deleting running "solr" pods which will be automatically recreated by kubernetes
	Restart(context.Context, *SolrCloudId) (*SolrCloud, error)
	// pause a solr-cloud running in a environment.
	// solr-cloud is paused by scaling down number of replicas of
	// the kubernetes deployment/stateful sets to zero in the environment.
	Pause(context.Context, *resource.ApiResourcePauseCommandInput) (*SolrCloud, error)
	// unpause a previously paused solr-cloud running in a environment.
	// unpause is done by scaling the number of pods back to the number of
	// replicas configured for the solr-cloud.
	Unpause(context.Context, *resource.ApiResourceUnPauseCommandInput) (*SolrCloud, error)
}

// UnimplementedSolrCloudCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedSolrCloudCommandControllerServer struct {
}

func (UnimplementedSolrCloudCommandControllerServer) PreviewCreate(context.Context, *SolrCloud) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedSolrCloudCommandControllerServer) Create(context.Context, *SolrCloud) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSolrCloudCommandControllerServer) PreviewUpdate(context.Context, *SolrCloud) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedSolrCloudCommandControllerServer) Update(context.Context, *SolrCloud) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSolrCloudCommandControllerServer) PreviewDelete(context.Context, *resource.ApiResourceDeleteCommandInput) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedSolrCloudCommandControllerServer) Delete(context.Context, *resource.ApiResourceDeleteCommandInput) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSolrCloudCommandControllerServer) PreviewRestore(context.Context, *SolrCloud) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedSolrCloudCommandControllerServer) Restore(context.Context, *SolrCloud) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedSolrCloudCommandControllerServer) CreateStackJob(context.Context, *job.CreateStackJobCommandInput) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStackJob not implemented")
}
func (UnimplementedSolrCloudCommandControllerServer) Restart(context.Context, *SolrCloudId) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedSolrCloudCommandControllerServer) Pause(context.Context, *resource.ApiResourcePauseCommandInput) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedSolrCloudCommandControllerServer) Unpause(context.Context, *resource.ApiResourceUnPauseCommandInput) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpause not implemented")
}

// UnsafeSolrCloudCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SolrCloudCommandControllerServer will
// result in compilation errors.
type UnsafeSolrCloudCommandControllerServer interface {
	mustEmbedUnimplementedSolrCloudCommandControllerServer()
}

func RegisterSolrCloudCommandControllerServer(s grpc.ServiceRegistrar, srv SolrCloudCommandControllerServer) {
	s.RegisterService(&SolrCloudCommandController_ServiceDesc, srv)
}

func _SolrCloudCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolrCloud)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).PreviewCreate(ctx, req.(*SolrCloud))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolrCloud)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).Create(ctx, req.(*SolrCloud))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolrCloud)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).PreviewUpdate(ctx, req.(*SolrCloud))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolrCloud)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).Update(ctx, req.(*SolrCloud))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).PreviewDelete(ctx, req.(*resource.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).Delete(ctx, req.(*resource.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolrCloud)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).PreviewRestore(ctx, req.(*SolrCloud))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolrCloud)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).Restore(ctx, req.(*SolrCloud))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudCommandController_CreateStackJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(job.CreateStackJobCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).CreateStackJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_CreateStackJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).CreateStackJob(ctx, req.(*job.CreateStackJobCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudCommandController_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolrCloudId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).Restart(ctx, req.(*SolrCloudId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudCommandController_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourcePauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).Pause(ctx, req.(*resource.ApiResourcePauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudCommandController_Unpause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceUnPauseCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).Unpause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_Unpause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).Unpause(ctx, req.(*resource.ApiResourceUnPauseCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// SolrCloudCommandController_ServiceDesc is the grpc.ServiceDesc for SolrCloudCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SolrCloudCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController",
	HandlerType: (*SolrCloudCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _SolrCloudCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _SolrCloudCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _SolrCloudCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _SolrCloudCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _SolrCloudCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _SolrCloudCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _SolrCloudCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _SolrCloudCommandController_Restore_Handler,
		},
		{
			MethodName: "createStackJob",
			Handler:    _SolrCloudCommandController_CreateStackJob_Handler,
		},
		{
			MethodName: "restart",
			Handler:    _SolrCloudCommandController_Restart_Handler,
		},
		{
			MethodName: "pause",
			Handler:    _SolrCloudCommandController_Pause_Handler,
		},
		{
			MethodName: "unpause",
			Handler:    _SolrCloudCommandController_Unpause_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/solr/service.proto",
}

const (
	SolrCloudQueryController_List_FullMethodName                = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudQueryController/list"
	SolrCloudQueryController_GetById_FullMethodName             = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudQueryController/getById"
	SolrCloudQueryController_FindByProductId_FullMethodName     = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudQueryController/findByProductId"
	SolrCloudQueryController_FindByEnvironmentId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudQueryController/findByEnvironmentId"
	SolrCloudQueryController_FindByKubeClusterId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudQueryController/findByKubeClusterId"
	SolrCloudQueryController_GetPassword_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudQueryController/getPassword"
	SolrCloudQueryController_FindPods_FullMethodName            = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudQueryController/findPods"
)

// SolrCloudQueryControllerClient is the client API for SolrCloudQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SolrCloudQueryControllerClient interface {
	// list all solr-clouds on planton cloud for the requested page. This is intended for use on portal.
	List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*SolrCloudList, error)
	// look up solr-cloud using solr-cloud id
	GetById(ctx context.Context, in *SolrCloudId, opts ...grpc.CallOption) (*SolrCloud, error)
	// find solr-clouds by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(ctx context.Context, in *product.ProductId, opts ...grpc.CallOption) (*SolrClouds, error)
	// find solr-clouds by environment
	FindByEnvironmentId(ctx context.Context, in *environment.EnvironmentId, opts ...grpc.CallOption) (*SolrClouds, error)
	// find solr-clouds by kubernetes cloud
	FindByKubeClusterId(ctx context.Context, in *kubecluster.KubeClusterId, opts ...grpc.CallOption) (*SolrClouds, error)
	// look up solr-cloud sasl password
	// password is retrieved from the kubernetes cloud.
	GetPassword(ctx context.Context, in *SolrCloudId, opts ...grpc.CallOption) (*SolrCloudPassword, error)
	// lookup pods of a solr-cloud deployed to a environment
	FindPods(ctx context.Context, in *SolrCloudId, opts ...grpc.CallOption) (*resource1.Pods, error)
}

type solrCloudQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewSolrCloudQueryControllerClient(cc grpc.ClientConnInterface) SolrCloudQueryControllerClient {
	return &solrCloudQueryControllerClient{cc}
}

func (c *solrCloudQueryControllerClient) List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*SolrCloudList, error) {
	out := new(SolrCloudList)
	err := c.cc.Invoke(ctx, SolrCloudQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudQueryControllerClient) GetById(ctx context.Context, in *SolrCloudId, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudQueryControllerClient) FindByProductId(ctx context.Context, in *product.ProductId, opts ...grpc.CallOption) (*SolrClouds, error) {
	out := new(SolrClouds)
	err := c.cc.Invoke(ctx, SolrCloudQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudQueryControllerClient) FindByEnvironmentId(ctx context.Context, in *environment.EnvironmentId, opts ...grpc.CallOption) (*SolrClouds, error) {
	out := new(SolrClouds)
	err := c.cc.Invoke(ctx, SolrCloudQueryController_FindByEnvironmentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudQueryControllerClient) FindByKubeClusterId(ctx context.Context, in *kubecluster.KubeClusterId, opts ...grpc.CallOption) (*SolrClouds, error) {
	out := new(SolrClouds)
	err := c.cc.Invoke(ctx, SolrCloudQueryController_FindByKubeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudQueryControllerClient) GetPassword(ctx context.Context, in *SolrCloudId, opts ...grpc.CallOption) (*SolrCloudPassword, error) {
	out := new(SolrCloudPassword)
	err := c.cc.Invoke(ctx, SolrCloudQueryController_GetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudQueryControllerClient) FindPods(ctx context.Context, in *SolrCloudId, opts ...grpc.CallOption) (*resource1.Pods, error) {
	out := new(resource1.Pods)
	err := c.cc.Invoke(ctx, SolrCloudQueryController_FindPods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolrCloudQueryControllerServer is the server API for SolrCloudQueryController service.
// All implementations should embed UnimplementedSolrCloudQueryControllerServer
// for forward compatibility
type SolrCloudQueryControllerServer interface {
	// list all solr-clouds on planton cloud for the requested page. This is intended for use on portal.
	List(context.Context, *pagination.PageInfo) (*SolrCloudList, error)
	// look up solr-cloud using solr-cloud id
	GetById(context.Context, *SolrCloudId) (*SolrCloud, error)
	// find solr-clouds by product id.
	// response contains only the resources that the authenticated user account has viewer access to.
	FindByProductId(context.Context, *product.ProductId) (*SolrClouds, error)
	// find solr-clouds by environment
	FindByEnvironmentId(context.Context, *environment.EnvironmentId) (*SolrClouds, error)
	// find solr-clouds by kubernetes cloud
	FindByKubeClusterId(context.Context, *kubecluster.KubeClusterId) (*SolrClouds, error)
	// look up solr-cloud sasl password
	// password is retrieved from the kubernetes cloud.
	GetPassword(context.Context, *SolrCloudId) (*SolrCloudPassword, error)
	// lookup pods of a solr-cloud deployed to a environment
	FindPods(context.Context, *SolrCloudId) (*resource1.Pods, error)
}

// UnimplementedSolrCloudQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedSolrCloudQueryControllerServer struct {
}

func (UnimplementedSolrCloudQueryControllerServer) List(context.Context, *pagination.PageInfo) (*SolrCloudList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSolrCloudQueryControllerServer) GetById(context.Context, *SolrCloudId) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedSolrCloudQueryControllerServer) FindByProductId(context.Context, *product.ProductId) (*SolrClouds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedSolrCloudQueryControllerServer) FindByEnvironmentId(context.Context, *environment.EnvironmentId) (*SolrClouds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEnvironmentId not implemented")
}
func (UnimplementedSolrCloudQueryControllerServer) FindByKubeClusterId(context.Context, *kubecluster.KubeClusterId) (*SolrClouds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKubeClusterId not implemented")
}
func (UnimplementedSolrCloudQueryControllerServer) GetPassword(context.Context, *SolrCloudId) (*SolrCloudPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassword not implemented")
}
func (UnimplementedSolrCloudQueryControllerServer) FindPods(context.Context, *SolrCloudId) (*resource1.Pods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPods not implemented")
}

// UnsafeSolrCloudQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SolrCloudQueryControllerServer will
// result in compilation errors.
type UnsafeSolrCloudQueryControllerServer interface {
	mustEmbedUnimplementedSolrCloudQueryControllerServer()
}

func RegisterSolrCloudQueryControllerServer(s grpc.ServiceRegistrar, srv SolrCloudQueryControllerServer) {
	s.RegisterService(&SolrCloudQueryController_ServiceDesc, srv)
}

func _SolrCloudQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudQueryControllerServer).List(ctx, req.(*pagination.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolrCloudId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudQueryControllerServer).GetById(ctx, req.(*SolrCloudId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(product.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudQueryControllerServer).FindByProductId(ctx, req.(*product.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudQueryController_FindByEnvironmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(environment.EnvironmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudQueryControllerServer).FindByEnvironmentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudQueryController_FindByEnvironmentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudQueryControllerServer).FindByEnvironmentId(ctx, req.(*environment.EnvironmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudQueryController_FindByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kubecluster.KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudQueryControllerServer).FindByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudQueryController_FindByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudQueryControllerServer).FindByKubeClusterId(ctx, req.(*kubecluster.KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudQueryController_GetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolrCloudId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudQueryControllerServer).GetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudQueryController_GetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudQueryControllerServer).GetPassword(ctx, req.(*SolrCloudId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudQueryController_FindPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolrCloudId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudQueryControllerServer).FindPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudQueryController_FindPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudQueryControllerServer).FindPods(ctx, req.(*SolrCloudId))
	}
	return interceptor(ctx, in, info, handler)
}

// SolrCloudQueryController_ServiceDesc is the grpc.ServiceDesc for SolrCloudQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SolrCloudQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudQueryController",
	HandlerType: (*SolrCloudQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _SolrCloudQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _SolrCloudQueryController_GetById_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _SolrCloudQueryController_FindByProductId_Handler,
		},
		{
			MethodName: "findByEnvironmentId",
			Handler:    _SolrCloudQueryController_FindByEnvironmentId_Handler,
		},
		{
			MethodName: "findByKubeClusterId",
			Handler:    _SolrCloudQueryController_FindByKubeClusterId_Handler,
		},
		{
			MethodName: "getPassword",
			Handler:    _SolrCloudQueryController_GetPassword_Handler,
		},
		{
			MethodName: "findPods",
			Handler:    _SolrCloudQueryController_FindPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/solr/service.proto",
}
