// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/aws/s3bucket/service/command.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/deploy/aws/s3bucket/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	S3BucketCommandController_PreviewCreate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket.service.S3BucketCommandController/previewCreate"
	S3BucketCommandController_Create_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket.service.S3BucketCommandController/create"
	S3BucketCommandController_PreviewUpdate_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket.service.S3BucketCommandController/previewUpdate"
	S3BucketCommandController_Update_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket.service.S3BucketCommandController/update"
	S3BucketCommandController_PreviewDelete_FullMethodName  = "/cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket.service.S3BucketCommandController/previewDelete"
	S3BucketCommandController_Delete_FullMethodName         = "/cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket.service.S3BucketCommandController/delete"
	S3BucketCommandController_PreviewRestore_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket.service.S3BucketCommandController/previewRestore"
	S3BucketCommandController_Restore_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket.service.S3BucketCommandController/restore"
	S3BucketCommandController_PreviewRefresh_FullMethodName = "/cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket.service.S3BucketCommandController/previewRefresh"
	S3BucketCommandController_Refresh_FullMethodName        = "/cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket.service.S3BucketCommandController/refresh"
)

// S3BucketCommandControllerClient is the client API for S3BucketCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type S3BucketCommandControllerClient interface {
	// preview create s3-bucket
	PreviewCreate(ctx context.Context, in *model.S3Bucket, opts ...grpc.CallOption) (*model.S3Bucket, error)
	// create s3-bucket
	Create(ctx context.Context, in *model.S3Bucket, opts ...grpc.CallOption) (*model.S3Bucket, error)
	// preview update an existing s3-bucket
	PreviewUpdate(ctx context.Context, in *model.S3Bucket, opts ...grpc.CallOption) (*model.S3Bucket, error)
	// update an existing s3-bucket
	Update(ctx context.Context, in *model.S3Bucket, opts ...grpc.CallOption) (*model.S3Bucket, error)
	// preview delete an existing s3-bucket
	PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.S3Bucket, error)
	// delete an existing s3-bucket
	Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.S3Bucket, error)
	// preview restore a deleted s3-bucket
	PreviewRestore(ctx context.Context, in *model.S3Bucket, opts ...grpc.CallOption) (*model.S3Bucket, error)
	// restore a deleted s3-bucket
	Restore(ctx context.Context, in *model.S3Bucket, opts ...grpc.CallOption) (*model.S3Bucket, error)
	// preview refresh a s3-bucket that was previously created
	PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.S3Bucket, error)
	// refresh a s3-bucket that was previously created
	Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.S3Bucket, error)
}

type s3BucketCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewS3BucketCommandControllerClient(cc grpc.ClientConnInterface) S3BucketCommandControllerClient {
	return &s3BucketCommandControllerClient{cc}
}

func (c *s3BucketCommandControllerClient) PreviewCreate(ctx context.Context, in *model.S3Bucket, opts ...grpc.CallOption) (*model.S3Bucket, error) {
	out := new(model.S3Bucket)
	err := c.cc.Invoke(ctx, S3BucketCommandController_PreviewCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3BucketCommandControllerClient) Create(ctx context.Context, in *model.S3Bucket, opts ...grpc.CallOption) (*model.S3Bucket, error) {
	out := new(model.S3Bucket)
	err := c.cc.Invoke(ctx, S3BucketCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3BucketCommandControllerClient) PreviewUpdate(ctx context.Context, in *model.S3Bucket, opts ...grpc.CallOption) (*model.S3Bucket, error) {
	out := new(model.S3Bucket)
	err := c.cc.Invoke(ctx, S3BucketCommandController_PreviewUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3BucketCommandControllerClient) Update(ctx context.Context, in *model.S3Bucket, opts ...grpc.CallOption) (*model.S3Bucket, error) {
	out := new(model.S3Bucket)
	err := c.cc.Invoke(ctx, S3BucketCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3BucketCommandControllerClient) PreviewDelete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.S3Bucket, error) {
	out := new(model.S3Bucket)
	err := c.cc.Invoke(ctx, S3BucketCommandController_PreviewDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3BucketCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceDeleteCommandInput, opts ...grpc.CallOption) (*model.S3Bucket, error) {
	out := new(model.S3Bucket)
	err := c.cc.Invoke(ctx, S3BucketCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3BucketCommandControllerClient) PreviewRestore(ctx context.Context, in *model.S3Bucket, opts ...grpc.CallOption) (*model.S3Bucket, error) {
	out := new(model.S3Bucket)
	err := c.cc.Invoke(ctx, S3BucketCommandController_PreviewRestore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3BucketCommandControllerClient) Restore(ctx context.Context, in *model.S3Bucket, opts ...grpc.CallOption) (*model.S3Bucket, error) {
	out := new(model.S3Bucket)
	err := c.cc.Invoke(ctx, S3BucketCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3BucketCommandControllerClient) PreviewRefresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.S3Bucket, error) {
	out := new(model.S3Bucket)
	err := c.cc.Invoke(ctx, S3BucketCommandController_PreviewRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3BucketCommandControllerClient) Refresh(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.S3Bucket, error) {
	out := new(model.S3Bucket)
	err := c.cc.Invoke(ctx, S3BucketCommandController_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S3BucketCommandControllerServer is the server API for S3BucketCommandController service.
// All implementations should embed UnimplementedS3BucketCommandControllerServer
// for forward compatibility
type S3BucketCommandControllerServer interface {
	// preview create s3-bucket
	PreviewCreate(context.Context, *model.S3Bucket) (*model.S3Bucket, error)
	// create s3-bucket
	Create(context.Context, *model.S3Bucket) (*model.S3Bucket, error)
	// preview update an existing s3-bucket
	PreviewUpdate(context.Context, *model.S3Bucket) (*model.S3Bucket, error)
	// update an existing s3-bucket
	Update(context.Context, *model.S3Bucket) (*model.S3Bucket, error)
	// preview delete an existing s3-bucket
	PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.S3Bucket, error)
	// delete an existing s3-bucket
	Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.S3Bucket, error)
	// preview restore a deleted s3-bucket
	PreviewRestore(context.Context, *model.S3Bucket) (*model.S3Bucket, error)
	// restore a deleted s3-bucket
	Restore(context.Context, *model.S3Bucket) (*model.S3Bucket, error)
	// preview refresh a s3-bucket that was previously created
	PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.S3Bucket, error)
	// refresh a s3-bucket that was previously created
	Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.S3Bucket, error)
}

// UnimplementedS3BucketCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedS3BucketCommandControllerServer struct {
}

func (UnimplementedS3BucketCommandControllerServer) PreviewCreate(context.Context, *model.S3Bucket) (*model.S3Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewCreate not implemented")
}
func (UnimplementedS3BucketCommandControllerServer) Create(context.Context, *model.S3Bucket) (*model.S3Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedS3BucketCommandControllerServer) PreviewUpdate(context.Context, *model.S3Bucket) (*model.S3Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewUpdate not implemented")
}
func (UnimplementedS3BucketCommandControllerServer) Update(context.Context, *model.S3Bucket) (*model.S3Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedS3BucketCommandControllerServer) PreviewDelete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.S3Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDelete not implemented")
}
func (UnimplementedS3BucketCommandControllerServer) Delete(context.Context, *model1.ApiResourceDeleteCommandInput) (*model.S3Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedS3BucketCommandControllerServer) PreviewRestore(context.Context, *model.S3Bucket) (*model.S3Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRestore not implemented")
}
func (UnimplementedS3BucketCommandControllerServer) Restore(context.Context, *model.S3Bucket) (*model.S3Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedS3BucketCommandControllerServer) PreviewRefresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.S3Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewRefresh not implemented")
}
func (UnimplementedS3BucketCommandControllerServer) Refresh(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.S3Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

// UnsafeS3BucketCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to S3BucketCommandControllerServer will
// result in compilation errors.
type UnsafeS3BucketCommandControllerServer interface {
	mustEmbedUnimplementedS3BucketCommandControllerServer()
}

func RegisterS3BucketCommandControllerServer(s grpc.ServiceRegistrar, srv S3BucketCommandControllerServer) {
	s.RegisterService(&S3BucketCommandController_ServiceDesc, srv)
}

func _S3BucketCommandController_PreviewCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.S3Bucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3BucketCommandControllerServer).PreviewCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3BucketCommandController_PreviewCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3BucketCommandControllerServer).PreviewCreate(ctx, req.(*model.S3Bucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3BucketCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.S3Bucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3BucketCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3BucketCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3BucketCommandControllerServer).Create(ctx, req.(*model.S3Bucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3BucketCommandController_PreviewUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.S3Bucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3BucketCommandControllerServer).PreviewUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3BucketCommandController_PreviewUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3BucketCommandControllerServer).PreviewUpdate(ctx, req.(*model.S3Bucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3BucketCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.S3Bucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3BucketCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3BucketCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3BucketCommandControllerServer).Update(ctx, req.(*model.S3Bucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3BucketCommandController_PreviewDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3BucketCommandControllerServer).PreviewDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3BucketCommandController_PreviewDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3BucketCommandControllerServer).PreviewDelete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3BucketCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceDeleteCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3BucketCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3BucketCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3BucketCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceDeleteCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3BucketCommandController_PreviewRestore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.S3Bucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3BucketCommandControllerServer).PreviewRestore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3BucketCommandController_PreviewRestore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3BucketCommandControllerServer).PreviewRestore(ctx, req.(*model.S3Bucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3BucketCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.S3Bucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3BucketCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3BucketCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3BucketCommandControllerServer).Restore(ctx, req.(*model.S3Bucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3BucketCommandController_PreviewRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3BucketCommandControllerServer).PreviewRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3BucketCommandController_PreviewRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3BucketCommandControllerServer).PreviewRefresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3BucketCommandController_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3BucketCommandControllerServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3BucketCommandController_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3BucketCommandControllerServer).Refresh(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// S3BucketCommandController_ServiceDesc is the grpc.ServiceDesc for S3BucketCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var S3BucketCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket.service.S3BucketCommandController",
	HandlerType: (*S3BucketCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "previewCreate",
			Handler:    _S3BucketCommandController_PreviewCreate_Handler,
		},
		{
			MethodName: "create",
			Handler:    _S3BucketCommandController_Create_Handler,
		},
		{
			MethodName: "previewUpdate",
			Handler:    _S3BucketCommandController_PreviewUpdate_Handler,
		},
		{
			MethodName: "update",
			Handler:    _S3BucketCommandController_Update_Handler,
		},
		{
			MethodName: "previewDelete",
			Handler:    _S3BucketCommandController_PreviewDelete_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _S3BucketCommandController_Delete_Handler,
		},
		{
			MethodName: "previewRestore",
			Handler:    _S3BucketCommandController_PreviewRestore_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _S3BucketCommandController_Restore_Handler,
		},
		{
			MethodName: "previewRefresh",
			Handler:    _S3BucketCommandController_PreviewRefresh_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _S3BucketCommandController_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/deploy/aws/s3bucket/service/command.proto",
}
