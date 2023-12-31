// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/stack/job/service.proto

package job

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StackJobCommandController_Create_FullMethodName             = "/cloud.planton.apis.v1.stack.job.StackJobCommandController/create"
	StackJobCommandController_Update_FullMethodName             = "/cloud.planton.apis.v1.stack.job.StackJobCommandController/update"
	StackJobCommandController_SetExecutionStatus_FullMethodName = "/cloud.planton.apis.v1.stack.job.StackJobCommandController/setExecutionStatus"
)

// StackJobCommandControllerClient is the client API for StackJobCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StackJobCommandControllerClient interface {
	// create stack-job
	Create(ctx context.Context, in *StackJob, opts ...grpc.CallOption) (*StackJob, error)
	// update stack-job
	Update(ctx context.Context, in *StackJob, opts ...grpc.CallOption) (*StackJob, error)
	// rpc to set execution-status for a stack-job.
	// since execution-status is an attribute in status of stack-job, it is not possible to update it using update rpc.
	SetExecutionStatus(ctx context.Context, in *SetStackJobExecutionStatusCommandInput, opts ...grpc.CallOption) (*StackJob, error)
}

type stackJobCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewStackJobCommandControllerClient(cc grpc.ClientConnInterface) StackJobCommandControllerClient {
	return &stackJobCommandControllerClient{cc}
}

func (c *stackJobCommandControllerClient) Create(ctx context.Context, in *StackJob, opts ...grpc.CallOption) (*StackJob, error) {
	out := new(StackJob)
	err := c.cc.Invoke(ctx, StackJobCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackJobCommandControllerClient) Update(ctx context.Context, in *StackJob, opts ...grpc.CallOption) (*StackJob, error) {
	out := new(StackJob)
	err := c.cc.Invoke(ctx, StackJobCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackJobCommandControllerClient) SetExecutionStatus(ctx context.Context, in *SetStackJobExecutionStatusCommandInput, opts ...grpc.CallOption) (*StackJob, error) {
	out := new(StackJob)
	err := c.cc.Invoke(ctx, StackJobCommandController_SetExecutionStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StackJobCommandControllerServer is the server API for StackJobCommandController service.
// All implementations should embed UnimplementedStackJobCommandControllerServer
// for forward compatibility
type StackJobCommandControllerServer interface {
	// create stack-job
	Create(context.Context, *StackJob) (*StackJob, error)
	// update stack-job
	Update(context.Context, *StackJob) (*StackJob, error)
	// rpc to set execution-status for a stack-job.
	// since execution-status is an attribute in status of stack-job, it is not possible to update it using update rpc.
	SetExecutionStatus(context.Context, *SetStackJobExecutionStatusCommandInput) (*StackJob, error)
}

// UnimplementedStackJobCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedStackJobCommandControllerServer struct {
}

func (UnimplementedStackJobCommandControllerServer) Create(context.Context, *StackJob) (*StackJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStackJobCommandControllerServer) Update(context.Context, *StackJob) (*StackJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStackJobCommandControllerServer) SetExecutionStatus(context.Context, *SetStackJobExecutionStatusCommandInput) (*StackJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetExecutionStatus not implemented")
}

// UnsafeStackJobCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StackJobCommandControllerServer will
// result in compilation errors.
type UnsafeStackJobCommandControllerServer interface {
	mustEmbedUnimplementedStackJobCommandControllerServer()
}

func RegisterStackJobCommandControllerServer(s grpc.ServiceRegistrar, srv StackJobCommandControllerServer) {
	s.RegisterService(&StackJobCommandController_ServiceDesc, srv)
}

func _StackJobCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackJobCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackJobCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackJobCommandControllerServer).Create(ctx, req.(*StackJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackJobCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackJobCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackJobCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackJobCommandControllerServer).Update(ctx, req.(*StackJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackJobCommandController_SetExecutionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStackJobExecutionStatusCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackJobCommandControllerServer).SetExecutionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackJobCommandController_SetExecutionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackJobCommandControllerServer).SetExecutionStatus(ctx, req.(*SetStackJobExecutionStatusCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

// StackJobCommandController_ServiceDesc is the grpc.ServiceDesc for StackJobCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StackJobCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.stack.job.StackJobCommandController",
	HandlerType: (*StackJobCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _StackJobCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _StackJobCommandController_Update_Handler,
		},
		{
			MethodName: "setExecutionStatus",
			Handler:    _StackJobCommandController_SetExecutionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/stack/job/service.proto",
}

const (
	StackJobQueryController_ListByFilters_FullMethodName                = "/cloud.planton.apis.v1.stack.job.StackJobQueryController/listByFilters"
	StackJobQueryController_GetById_FullMethodName                      = "/cloud.planton.apis.v1.stack.job.StackJobQueryController/getById"
	StackJobQueryController_GetProgressEventStream_FullMethodName       = "/cloud.planton.apis.v1.stack.job.StackJobQueryController/getProgressEventStream"
	StackJobQueryController_GetStackJobLogSnapshotStream_FullMethodName = "/cloud.planton.apis.v1.stack.job.StackJobQueryController/getStackJobLogSnapshotStream"
)

// StackJobQueryControllerClient is the client API for StackJobQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StackJobQueryControllerClient interface {
	// list of stack-jobs
	// todo: need to add authorization
	ListByFilters(ctx context.Context, in *ListStackJobsByFiltersQueryInput, opts ...grpc.CallOption) (*StackJobList, error)
	// look up stack-job by stack-job-id
	// todo: need to add authorization
	GetById(ctx context.Context, in *StackJobId, opts ...grpc.CallOption) (*StackJob, error)
	// rpc to get stack-job progress event stream
	// this rpc is required to support pulumi progress or diff view in the cli client which is an interactive client.
	GetProgressEventStream(ctx context.Context, in *StackJobId, opts ...grpc.CallOption) (StackJobQueryController_GetProgressEventStreamClient, error)
	// rpc to get snapshot of stack-job log
	// response is optimized to display the stack-job log on non-interactive client-apps
	GetStackJobLogSnapshotStream(ctx context.Context, in *StackJobId, opts ...grpc.CallOption) (StackJobQueryController_GetStackJobLogSnapshotStreamClient, error)
}

type stackJobQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewStackJobQueryControllerClient(cc grpc.ClientConnInterface) StackJobQueryControllerClient {
	return &stackJobQueryControllerClient{cc}
}

func (c *stackJobQueryControllerClient) ListByFilters(ctx context.Context, in *ListStackJobsByFiltersQueryInput, opts ...grpc.CallOption) (*StackJobList, error) {
	out := new(StackJobList)
	err := c.cc.Invoke(ctx, StackJobQueryController_ListByFilters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackJobQueryControllerClient) GetById(ctx context.Context, in *StackJobId, opts ...grpc.CallOption) (*StackJob, error) {
	out := new(StackJob)
	err := c.cc.Invoke(ctx, StackJobQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackJobQueryControllerClient) GetProgressEventStream(ctx context.Context, in *StackJobId, opts ...grpc.CallOption) (StackJobQueryController_GetProgressEventStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StackJobQueryController_ServiceDesc.Streams[0], StackJobQueryController_GetProgressEventStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &stackJobQueryControllerGetProgressEventStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StackJobQueryController_GetProgressEventStreamClient interface {
	Recv() (*StackJobProgressEvent, error)
	grpc.ClientStream
}

type stackJobQueryControllerGetProgressEventStreamClient struct {
	grpc.ClientStream
}

func (x *stackJobQueryControllerGetProgressEventStreamClient) Recv() (*StackJobProgressEvent, error) {
	m := new(StackJobProgressEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stackJobQueryControllerClient) GetStackJobLogSnapshotStream(ctx context.Context, in *StackJobId, opts ...grpc.CallOption) (StackJobQueryController_GetStackJobLogSnapshotStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StackJobQueryController_ServiceDesc.Streams[1], StackJobQueryController_GetStackJobLogSnapshotStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &stackJobQueryControllerGetStackJobLogSnapshotStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StackJobQueryController_GetStackJobLogSnapshotStreamClient interface {
	Recv() (*StackJobLogSnapshot, error)
	grpc.ClientStream
}

type stackJobQueryControllerGetStackJobLogSnapshotStreamClient struct {
	grpc.ClientStream
}

func (x *stackJobQueryControllerGetStackJobLogSnapshotStreamClient) Recv() (*StackJobLogSnapshot, error) {
	m := new(StackJobLogSnapshot)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StackJobQueryControllerServer is the server API for StackJobQueryController service.
// All implementations should embed UnimplementedStackJobQueryControllerServer
// for forward compatibility
type StackJobQueryControllerServer interface {
	// list of stack-jobs
	// todo: need to add authorization
	ListByFilters(context.Context, *ListStackJobsByFiltersQueryInput) (*StackJobList, error)
	// look up stack-job by stack-job-id
	// todo: need to add authorization
	GetById(context.Context, *StackJobId) (*StackJob, error)
	// rpc to get stack-job progress event stream
	// this rpc is required to support pulumi progress or diff view in the cli client which is an interactive client.
	GetProgressEventStream(*StackJobId, StackJobQueryController_GetProgressEventStreamServer) error
	// rpc to get snapshot of stack-job log
	// response is optimized to display the stack-job log on non-interactive client-apps
	GetStackJobLogSnapshotStream(*StackJobId, StackJobQueryController_GetStackJobLogSnapshotStreamServer) error
}

// UnimplementedStackJobQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedStackJobQueryControllerServer struct {
}

func (UnimplementedStackJobQueryControllerServer) ListByFilters(context.Context, *ListStackJobsByFiltersQueryInput) (*StackJobList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByFilters not implemented")
}
func (UnimplementedStackJobQueryControllerServer) GetById(context.Context, *StackJobId) (*StackJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedStackJobQueryControllerServer) GetProgressEventStream(*StackJobId, StackJobQueryController_GetProgressEventStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProgressEventStream not implemented")
}
func (UnimplementedStackJobQueryControllerServer) GetStackJobLogSnapshotStream(*StackJobId, StackJobQueryController_GetStackJobLogSnapshotStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStackJobLogSnapshotStream not implemented")
}

// UnsafeStackJobQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StackJobQueryControllerServer will
// result in compilation errors.
type UnsafeStackJobQueryControllerServer interface {
	mustEmbedUnimplementedStackJobQueryControllerServer()
}

func RegisterStackJobQueryControllerServer(s grpc.ServiceRegistrar, srv StackJobQueryControllerServer) {
	s.RegisterService(&StackJobQueryController_ServiceDesc, srv)
}

func _StackJobQueryController_ListByFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStackJobsByFiltersQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackJobQueryControllerServer).ListByFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackJobQueryController_ListByFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackJobQueryControllerServer).ListByFilters(ctx, req.(*ListStackJobsByFiltersQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackJobQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackJobId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackJobQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StackJobQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackJobQueryControllerServer).GetById(ctx, req.(*StackJobId))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackJobQueryController_GetProgressEventStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StackJobId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StackJobQueryControllerServer).GetProgressEventStream(m, &stackJobQueryControllerGetProgressEventStreamServer{stream})
}

type StackJobQueryController_GetProgressEventStreamServer interface {
	Send(*StackJobProgressEvent) error
	grpc.ServerStream
}

type stackJobQueryControllerGetProgressEventStreamServer struct {
	grpc.ServerStream
}

func (x *stackJobQueryControllerGetProgressEventStreamServer) Send(m *StackJobProgressEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _StackJobQueryController_GetStackJobLogSnapshotStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StackJobId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StackJobQueryControllerServer).GetStackJobLogSnapshotStream(m, &stackJobQueryControllerGetStackJobLogSnapshotStreamServer{stream})
}

type StackJobQueryController_GetStackJobLogSnapshotStreamServer interface {
	Send(*StackJobLogSnapshot) error
	grpc.ServerStream
}

type stackJobQueryControllerGetStackJobLogSnapshotStreamServer struct {
	grpc.ServerStream
}

func (x *stackJobQueryControllerGetStackJobLogSnapshotStreamServer) Send(m *StackJobLogSnapshot) error {
	return x.ServerStream.SendMsg(m)
}

// StackJobQueryController_ServiceDesc is the grpc.ServiceDesc for StackJobQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StackJobQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.stack.job.StackJobQueryController",
	HandlerType: (*StackJobQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listByFilters",
			Handler:    _StackJobQueryController_ListByFilters_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _StackJobQueryController_GetById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getProgressEventStream",
			Handler:       _StackJobQueryController_GetProgressEventStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "getStackJobLogSnapshotStream",
			Handler:       _StackJobQueryController_GetStackJobLogSnapshotStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/v1/stack/job/service.proto",
}
