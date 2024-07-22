// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2build/v1/buildengine/service/query.proto

package service

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const ()

// BuildEngineQueryControllerClient is the client API for BuildEngineQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuildEngineQueryControllerClient interface {
}

type buildEngineQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildEngineQueryControllerClient(cc grpc.ClientConnInterface) BuildEngineQueryControllerClient {
	return &buildEngineQueryControllerClient{cc}
}

// BuildEngineQueryControllerServer is the server API for BuildEngineQueryController service.
// All implementations should embed UnimplementedBuildEngineQueryControllerServer
// for forward compatibility
type BuildEngineQueryControllerServer interface {
}

// UnimplementedBuildEngineQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedBuildEngineQueryControllerServer struct {
}

// UnsafeBuildEngineQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuildEngineQueryControllerServer will
// result in compilation errors.
type UnsafeBuildEngineQueryControllerServer interface {
	mustEmbedUnimplementedBuildEngineQueryControllerServer()
}

func RegisterBuildEngineQueryControllerServer(s grpc.ServiceRegistrar, srv BuildEngineQueryControllerServer) {
	s.RegisterService(&BuildEngineQueryController_ServiceDesc, srv)
}

// BuildEngineQueryController_ServiceDesc is the grpc.ServiceDesc for BuildEngineQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BuildEngineQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2build.v1.buildengine.service.BuildEngineQueryController",
	HandlerType: (*BuildEngineQueryControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cloud/planton/apis/code2build/v1/buildengine/service/query.proto",
}
