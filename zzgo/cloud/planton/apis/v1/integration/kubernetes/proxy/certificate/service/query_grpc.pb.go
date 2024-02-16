// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/integration/kubernetes/proxy/certificate/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/integration/kubernetes/proxy/certificate/model"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/integration/kubernetes/resource"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CertificateProxyQueryController_GetCertificateByNamespaceByName_FullMethodName = "/cloud.planton.apis.v1.integration.kubernetes.proxy.certificate.service.CertificateProxyQueryController/getCertificateByNamespaceByName"
	CertificateProxyQueryController_FindCertificates_FullMethodName                = "/cloud.planton.apis.v1.integration.kubernetes.proxy.certificate.service.CertificateProxyQueryController/findCertificates"
)

// CertificateProxyQueryControllerClient is the client API for CertificateProxyQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificateProxyQueryControllerClient interface {
	GetCertificateByNamespaceByName(ctx context.Context, in *model.GetCertificateByNamespaceByNameQueryInput, opts ...grpc.CallOption) (*resource.Certificate, error)
	FindCertificates(ctx context.Context, in *model.FindCertificatesQueryInput, opts ...grpc.CallOption) (*resource.Certificates, error)
}

type certificateProxyQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateProxyQueryControllerClient(cc grpc.ClientConnInterface) CertificateProxyQueryControllerClient {
	return &certificateProxyQueryControllerClient{cc}
}

func (c *certificateProxyQueryControllerClient) GetCertificateByNamespaceByName(ctx context.Context, in *model.GetCertificateByNamespaceByNameQueryInput, opts ...grpc.CallOption) (*resource.Certificate, error) {
	out := new(resource.Certificate)
	err := c.cc.Invoke(ctx, CertificateProxyQueryController_GetCertificateByNamespaceByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateProxyQueryControllerClient) FindCertificates(ctx context.Context, in *model.FindCertificatesQueryInput, opts ...grpc.CallOption) (*resource.Certificates, error) {
	out := new(resource.Certificates)
	err := c.cc.Invoke(ctx, CertificateProxyQueryController_FindCertificates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateProxyQueryControllerServer is the server API for CertificateProxyQueryController service.
// All implementations should embed UnimplementedCertificateProxyQueryControllerServer
// for forward compatibility
type CertificateProxyQueryControllerServer interface {
	GetCertificateByNamespaceByName(context.Context, *model.GetCertificateByNamespaceByNameQueryInput) (*resource.Certificate, error)
	FindCertificates(context.Context, *model.FindCertificatesQueryInput) (*resource.Certificates, error)
}

// UnimplementedCertificateProxyQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCertificateProxyQueryControllerServer struct {
}

func (UnimplementedCertificateProxyQueryControllerServer) GetCertificateByNamespaceByName(context.Context, *model.GetCertificateByNamespaceByNameQueryInput) (*resource.Certificate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificateByNamespaceByName not implemented")
}
func (UnimplementedCertificateProxyQueryControllerServer) FindCertificates(context.Context, *model.FindCertificatesQueryInput) (*resource.Certificates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCertificates not implemented")
}

// UnsafeCertificateProxyQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateProxyQueryControllerServer will
// result in compilation errors.
type UnsafeCertificateProxyQueryControllerServer interface {
	mustEmbedUnimplementedCertificateProxyQueryControllerServer()
}

func RegisterCertificateProxyQueryControllerServer(s grpc.ServiceRegistrar, srv CertificateProxyQueryControllerServer) {
	s.RegisterService(&CertificateProxyQueryController_ServiceDesc, srv)
}

func _CertificateProxyQueryController_GetCertificateByNamespaceByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetCertificateByNamespaceByNameQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateProxyQueryControllerServer).GetCertificateByNamespaceByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateProxyQueryController_GetCertificateByNamespaceByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateProxyQueryControllerServer).GetCertificateByNamespaceByName(ctx, req.(*model.GetCertificateByNamespaceByNameQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateProxyQueryController_FindCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.FindCertificatesQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateProxyQueryControllerServer).FindCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateProxyQueryController_FindCertificates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateProxyQueryControllerServer).FindCertificates(ctx, req.(*model.FindCertificatesQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateProxyQueryController_ServiceDesc is the grpc.ServiceDesc for CertificateProxyQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateProxyQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.integration.kubernetes.proxy.certificate.service.CertificateProxyQueryController",
	HandlerType: (*CertificateProxyQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCertificateByNamespaceByName",
			Handler:    _CertificateProxyQueryController_GetCertificateByNamespaceByName_Handler,
		},
		{
			MethodName: "findCertificates",
			Handler:    _CertificateProxyQueryController_FindCertificates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/integration/kubernetes/proxy/certificate/service/query.proto",
}