// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/connect/v1/gitlabcredential/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/gitlabcredential/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GitlabCredentialQueryController_Get_FullMethodName = "/cloud.planton.apis.connect.v1.gitlabcredential.service.GitlabCredentialQueryController/get"
)

// GitlabCredentialQueryControllerClient is the client API for GitlabCredentialQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitlabCredentialQueryControllerClient interface {
	// look up a gitlab-credential by id
	Get(ctx context.Context, in *model.GitlabCredentialId, opts ...grpc.CallOption) (*model.GitlabCredential, error)
}

type gitlabCredentialQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGitlabCredentialQueryControllerClient(cc grpc.ClientConnInterface) GitlabCredentialQueryControllerClient {
	return &gitlabCredentialQueryControllerClient{cc}
}

func (c *gitlabCredentialQueryControllerClient) Get(ctx context.Context, in *model.GitlabCredentialId, opts ...grpc.CallOption) (*model.GitlabCredential, error) {
	out := new(model.GitlabCredential)
	err := c.cc.Invoke(ctx, GitlabCredentialQueryController_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitlabCredentialQueryControllerServer is the server API for GitlabCredentialQueryController service.
// All implementations should embed UnimplementedGitlabCredentialQueryControllerServer
// for forward compatibility
type GitlabCredentialQueryControllerServer interface {
	// look up a gitlab-credential by id
	Get(context.Context, *model.GitlabCredentialId) (*model.GitlabCredential, error)
}

// UnimplementedGitlabCredentialQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGitlabCredentialQueryControllerServer struct {
}

func (UnimplementedGitlabCredentialQueryControllerServer) Get(context.Context, *model.GitlabCredentialId) (*model.GitlabCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeGitlabCredentialQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitlabCredentialQueryControllerServer will
// result in compilation errors.
type UnsafeGitlabCredentialQueryControllerServer interface {
	mustEmbedUnimplementedGitlabCredentialQueryControllerServer()
}

func RegisterGitlabCredentialQueryControllerServer(s grpc.ServiceRegistrar, srv GitlabCredentialQueryControllerServer) {
	s.RegisterService(&GitlabCredentialQueryController_ServiceDesc, srv)
}

func _GitlabCredentialQueryController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GitlabCredentialId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabCredentialQueryControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitlabCredentialQueryController_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabCredentialQueryControllerServer).Get(ctx, req.(*model.GitlabCredentialId))
	}
	return interceptor(ctx, in, info, handler)
}

// GitlabCredentialQueryController_ServiceDesc is the grpc.ServiceDesc for GitlabCredentialQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitlabCredentialQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.connect.v1.gitlabcredential.service.GitlabCredentialQueryController",
	HandlerType: (*GitlabCredentialQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _GitlabCredentialQueryController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/connect/v1/gitlabcredential/service/query.proto",
}
