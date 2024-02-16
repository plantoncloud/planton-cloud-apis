// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/solr/command.proto

package solr

import (
	context "context"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
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
	SolrCloudCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/previewRefresh"
	SolrCloudCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.deploy.solr.SolrCloudCommandController/refresh"
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
	// preview refresh a solr-cloud that was previously created
	PreviewRefresh(ctx context.Context, in *resource.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*SolrCloud, error)
	// refresh a solr-cloud that was previously created
	Refresh(ctx context.Context, in *resource.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*SolrCloud, error)
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

func (c *solrCloudCommandControllerClient) PreviewRefresh(ctx context.Context, in *resource.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solrCloudCommandControllerClient) Refresh(ctx context.Context, in *resource.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*SolrCloud, error) {
	out := new(SolrCloud)
	err := c.cc.Invoke(ctx, SolrCloudCommandController_Refresh_FullMethodName, in, out, opts...)
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
	// preview refresh a solr-cloud that was previously created
	PreviewRefresh(context.Context, *resource.ApiResourceRefreshCommandInput) (*SolrCloud, error)
	// refresh a solr-cloud that was previously created
	Refresh(context.Context, *resource.ApiResourceRefreshCommandInput) (*SolrCloud, error)
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
func (UnimplementedSolrCloudCommandControllerServer) PreviewRefresh(context.Context, *resource.ApiResourceRefreshCommandInput) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedSolrCloudCommandControllerServer) Refresh(context.Context, *resource.ApiResourceRefreshCommandInput) (*SolrCloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
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

func _SolrCloudCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).PreviewRefresh(ctx, req.(*resource.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolrCloudCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resource.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolrCloudCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SolrCloudCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolrCloudCommandControllerServer).Refresh(ctx, req.(*resource.ApiResourceRefreshCommandInput))
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
		{
			MethodName: "previewRefresh",
			Handler:    _SolrCloudCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _SolrCloudCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/solr/command.proto",
}
