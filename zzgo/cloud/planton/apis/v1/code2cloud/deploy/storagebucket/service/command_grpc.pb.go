// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/storagebucket/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/storagebucket/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StorageBucketCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.service.StorageBucketCommandController/previewCreate"
	StorageBucketCommandController_Create_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.service.StorageBucketCommandController/create"
	StorageBucketCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.service.StorageBucketCommandController/previewUpdate"
	StorageBucketCommandController_Update_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.service.StorageBucketCommandController/update"
	StorageBucketCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.service.StorageBucketCommandController/previewDelete"
	StorageBucketCommandController_Delete_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.service.StorageBucketCommandController/delete"
	StorageBucketCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.service.StorageBucketCommandController/previewRestore"
	StorageBucketCommandController_Restore_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.service.StorageBucketCommandController/restore"
	StorageBucketCommandController_CreateStackJob_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.service.StorageBucketCommandController/createStackJob"
	StorageBucketCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.service.StorageBucketCommandController/previewRefresh"
	StorageBucketCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.v1.code2cloud.deploy.storagebucket.service.StorageBucketCommandController/refresh"
)

// StorageBucketCommandControllerClient is the client API for StorageBucketCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageBucketCommandControllerClient interface {
	// preview create storage-bucket
	PreviewCreate(ctx context.Context, in *model.StorageBucket, opts ...grpc.CallOption) (*model.StorageBucket, error)
	// create storage-bucket
	Create(ctx context.Context, in *model.StorageBucket, opts ...grpc.CallOption) (*model.StorageBucket, error)
	// preview update an existing storage-bucket
	PreviewUpdate(ctx context.Context, in *model.StorageBucket, opts ...grpc.CallOption) (*model.StorageBucket, error)
	// update an existing storage-bucket
	Update(ctx context.Context, in *model.StorageBucket, opts ...grpc.CallOption) (*model.StorageBucket, error)
	// preview delete an existing storage-bucket
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.StorageBucket, error)
	// delete an existing storage-bucket
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.StorageBucket, error)
	// preview restore a deleted storage-bucket
	PreviewRestore(ctx context.Context, in *model.StorageBucket, opts ...grpc.CallOption) (*model.StorageBucket, error)
	// restore a deleted storage-bucket
	Restore(ctx context.Context, in *model.StorageBucket, opts ...grpc.CallOption) (*model.StorageBucket, error)
	// create-stack-job for storage-bucket
	CreateStackJob(ctx context.Context, in *model2.CreateStackJobCommandInput, opts ...grpc.CallOption) (*model.StorageBucket, error)
	// preview refresh a storage-bucket that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.StorageBucket, error)
	// refresh a storage-bucket that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.StorageBucket, error)
}

type storageBucketCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageBucketCommandControllerClient(cc grpc.ClientConnInterface) StorageBucketCommandControllerClient {
	return &storageBucketCommandControllerClient{cc}
}

func (c *storageBucketCommandControllerClient) PreviewCreate(ctx context.Context, in *model.StorageBucket, opts ...grpc.CallOption) (*model.StorageBucket, error) {
	out := new(model.StorageBucket)
	err := c.cc.Invoke(ctx, StorageBucketCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBucketCommandControllerClient) Create(ctx context.Context, in *model.StorageBucket, opts ...grpc.CallOption) (*model.StorageBucket, error) {
	out := new(model.StorageBucket)
	err := c.cc.Invoke(ctx, StorageBucketCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBucketCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.StorageBucket, opts ...grpc.CallOption) (*model.StorageBucket, error) {
	out := new(model.StorageBucket)
	err := c.cc.Invoke(ctx, StorageBucketCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBucketCommandControllerClient) Update(ctx context.Context, in *model.StorageBucket, opts ...grpc.CallOption) (*model.StorageBucket, error) {
	out := new(model.StorageBucket)
	err := c.cc.Invoke(ctx, StorageBucketCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBucketCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.StorageBucket, error) {
	out := new(model.StorageBucket)
	err := c.cc.Invoke(ctx, StorageBucketCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBucketCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.StorageBucket, error) {
	out := new(model.StorageBucket)
	err := c.cc.Invoke(ctx, StorageBucketCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBucketCommandControllerClient) PreviewRestore(ctx context.Context, in *model.StorageBucket, opts ...grpc.CallOption) (*model.StorageBucket, error) {
	out := new(model.StorageBucket)
	err := c.cc.Invoke(ctx, StorageBucketCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBucketCommandControllerClient) Restore(ctx context.Context, in *model.StorageBucket, opts ...grpc.CallOption) (*model.StorageBucket, error) {
	out := new(model.StorageBucket)
	err := c.cc.Invoke(ctx, StorageBucketCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBucketCommandControllerClient) CreateStackJob(ctx context.Context, in *model2.CreateStackJobCommandInput, opts ...grpc.CallOption) (*model.StorageBucket, error) {
	out := new(model.StorageBucket)
	err := c.cc.Invoke(ctx, StorageBucketCommandController_CreateStackJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBucketCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.StorageBucket, error) {
	out := new(model.StorageBucket)
	err := c.cc.Invoke(ctx, StorageBucketCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBucketCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.StorageBucket, error) {
	out := new(model.StorageBucket)
	err := c.cc.Invoke(ctx, StorageBucketCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageBucketCommandControllerServer is the server API for StorageBucketCommandController service.
// All implementations should embed UnimplementedStorageBucketCommandControllerServer
// for forward compatibility
type StorageBucketCommandControllerServer interface {
	// preview create storage-bucket
	PreviewCreate(context.Context, *model.StorageBucket) (*model.StorageBucket, error)
	// create storage-bucket
	Create(context.Context, *model.StorageBucket) (*model.StorageBucket, error)
	// preview update an existing storage-bucket
	PreviewUpdate(context.Context, *model.StorageBucket) (*model.StorageBucket, error)
	// update an existing storage-bucket
	Update(context.Context, *model.StorageBucket) (*model.StorageBucket, error)
	// preview delete an existing storage-bucket
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.StorageBucket, error)
	// delete an existing storage-bucket
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.StorageBucket, error)
	// preview restore a deleted storage-bucket
	PreviewRestore(context.Context, *model.StorageBucket) (*model.StorageBucket, error)
	// restore a deleted storage-bucket
	Restore(context.Context, *model.StorageBucket) (*model.StorageBucket, error)
	// create-stack-job for storage-bucket
	CreateStackJob(context.Context, *model2.CreateStackJobCommandInput) (*model.StorageBucket, error)
	// preview refresh a storage-bucket that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.StorageBucket, error)
	// refresh a storage-bucket that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.StorageBucket, error)
}

// UnimplementedStorageBucketCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedStorageBucketCommandControllerServer struct {
}

func (UnimplementedStorageBucketCommandControllerServer) PreviewCreate(context.Context, *model.StorageBucket) (*model.StorageBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedStorageBucketCommandControllerServer) Create(context.Context, *model.StorageBucket) (*model.StorageBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStorageBucketCommandControllerServer) PreviewUpdate(context.Context, *model.StorageBucket) (*model.StorageBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedStorageBucketCommandControllerServer) Update(context.Context, *model.StorageBucket) (*model.StorageBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStorageBucketCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.StorageBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedStorageBucketCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.StorageBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStorageBucketCommandControllerServer) PreviewRestore(context.Context, *model.StorageBucket) (*model.StorageBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedStorageBucketCommandControllerServer) Restore(context.Context, *model.StorageBucket) (*model.StorageBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedStorageBucketCommandControllerServer) CreateStackJob(context.Context, *model2.CreateStackJobCommandInput) (*model.StorageBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStackJob not implemented")
}
func (UnimplementedStorageBucketCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.StorageBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedStorageBucketCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.StorageBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeStorageBucketCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageBucketCommandControllerServer will
// result in compilation errors.
type UnsafeStorageBucketCommandControllerServer interface {
	mustEmbedUnimplementedStorageBucketCommandControllerServer()
}

func RegisterStorageBucketCommandControllerServer(s grpc.ServiceRegistrar, srv StorageBucketCommandControllerServer) {
	s.RegisterService(&StorageBucketCommandController_ServiceDesc, srv)
}

func _StorageBucketCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.StorageBucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBucketCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageBucketCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBucketCommandControllerServer).PreviewCreate(ctx, req.(*model.StorageBucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBucketCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.StorageBucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBucketCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageBucketCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBucketCommandControllerServer).Create(ctx, req.(*model.StorageBucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBucketCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.StorageBucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBucketCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageBucketCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBucketCommandControllerServer).PreviewUpdate(ctx, req.(*model.StorageBucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBucketCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.StorageBucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBucketCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageBucketCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBucketCommandControllerServer).Update(ctx, req.(*model.StorageBucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBucketCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBucketCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageBucketCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBucketCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBucketCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBucketCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageBucketCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBucketCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBucketCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.StorageBucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBucketCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageBucketCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBucketCommandControllerServer).PreviewRestore(ctx, req.(*model.StorageBucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBucketCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.StorageBucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBucketCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageBucketCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBucketCommandControllerServer).Restore(ctx, req.(*model.StorageBucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBucketCommandController_CreateStackJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.CreateStackJobCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBucketCommandControllerServer).CreateStackJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageBucketCommandController_CreateStackJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBucketCommandControllerServer).CreateStackJob(ctx, req.(*model2.CreateStackJobCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBucketCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBucketCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageBucketCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBucketCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBucketCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBucketCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageBucketCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBucketCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageBucketCommandController_ServiceDesc is the grpc.ServiceDesc for StorageBucketCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageBucketCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.storagebucket.service.StorageBucketCommandController",
	HandlerType: (*StorageBucketCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _StorageBucketCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _StorageBucketCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _StorageBucketCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _StorageBucketCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _StorageBucketCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _StorageBucketCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _StorageBucketCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _StorageBucketCommandController_Restore_Handler,
		},
		{
			MethodName: "createStackJob",
			Handler:    _StorageBucketCommandController_CreateStackJob_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _StorageBucketCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _StorageBucketCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/storagebucket/service/command.proto",
}
