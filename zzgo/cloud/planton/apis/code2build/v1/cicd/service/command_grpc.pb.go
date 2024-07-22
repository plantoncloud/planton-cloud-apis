// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2build/v1/cicd/service/command.proto

package service

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const ()

// CiCdCommandControllerClient is the client API for CiCdCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CiCdCommandControllerClient interface {
}

type ciCdCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCiCdCommandControllerClient(cc grpc.ClientConnInterface) CiCdCommandControllerClient {
	return &ciCdCommandControllerClient{cc}
}

// CiCdCommandControllerServer is the server API for CiCdCommandController service.
// All implementations should embed UnimplementedCiCdCommandControllerServer
// for forward compatibility
type CiCdCommandControllerServer interface {
}

// UnimplementedCiCdCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCiCdCommandControllerServer struct {
}

// UnsafeCiCdCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CiCdCommandControllerServer will
// result in compilation errors.
type UnsafeCiCdCommandControllerServer interface {
	mustEmbedUnimplementedCiCdCommandControllerServer()
}

func RegisterCiCdCommandControllerServer(s grpc.ServiceRegistrar, srv CiCdCommandControllerServer) {
	s.RegisterService(&CiCdCommandController_ServiceDesc, srv)
}

// CiCdCommandController_ServiceDesc is the grpc.ServiceDesc for CiCdCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CiCdCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2build.v1.cicd.service.CiCdCommandController",
	HandlerType: (*CiCdCommandControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cloud/planton/apis/code2build/v1/cicd/service/command.proto",
}
