// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/cloudaccount/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/cloudaccount/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/resourcemanager/v1/company/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CloudAccountQueryController_GetById_FullMethodName                              = "/cloud.planton.apis.code2cloud.v1.cloudaccount.service.CloudAccountQueryController/getById"
	CloudAccountQueryController_FindByCompanyId_FullMethodName                      = "/cloud.planton.apis.code2cloud.v1.cloudaccount.service.CloudAccountQueryController/findByCompanyId"
	CloudAccountQueryController_FindArtifactStoreCreateCloudAccounts_FullMethodName = "/cloud.planton.apis.code2cloud.v1.cloudaccount.service.CloudAccountQueryController/findArtifactStoreCreateCloudAccounts"
	CloudAccountQueryController_FindKubeClusterCreateCloudAccounts_FullMethodName   = "/cloud.planton.apis.code2cloud.v1.cloudaccount.service.CloudAccountQueryController/findKubeClusterCreateCloudAccounts"
	CloudAccountQueryController_FindDnsZoneCreateCloudAccounts_FullMethodName       = "/cloud.planton.apis.code2cloud.v1.cloudaccount.service.CloudAccountQueryController/findDnsZoneCreateCloudAccounts"
)

// CloudAccountQueryControllerClient is the client API for CloudAccountQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudAccountQueryControllerClient interface {
	// look up a cloud-account by id
	GetById(ctx context.Context, in *model.CloudAccountId, opts ...grpc.CallOption) (*model.CloudAccount, error)
	// find cloud-accounts by company id.
	// the response should only include cloud-accounts in a company that the authenticated user account has viewer access to.
	// authorization is handled internally by running get authorized cloud account ids
	FindByCompanyId(ctx context.Context, in *model1.CompanyId, opts ...grpc.CallOption) (*model.CloudAccountList, error)
	// find cloud-accounts by company id to create artifact store.
	// this will be used to populate drop down of cloud-accounts in create artifact store form.
	// the response should only include cloud-accounts that a company is authorised to create artifact stores.
	// the authorization is verified by looking up cloud-accounts with `company-artifact-creator` relation for the company id provided in input.
	// the response should only include public attributes of a cloud-account. all non-public attributes should be excluded from the response.
	FindArtifactStoreCreateCloudAccounts(ctx context.Context, in *model1.CompanyId, opts ...grpc.CallOption) (*model.CloudAccountList, error)
	// find cloud-accounts by company id to create kube-cluster.
	// this will be used to populate drop down of cloud-accounts in create kube-cluster form.
	// the response should only include cloud-accounts that a company is authorised to create kube-cluster.
	// the authorization is verified by looking up cloud-accounts with `company-kube-cluster-creator` relation for the company id provided in input.
	// the response should only include public attributes of a cloud-account. all non-public attributes should be excluded from the response.
	FindKubeClusterCreateCloudAccounts(ctx context.Context, in *model1.CompanyId, opts ...grpc.CallOption) (*model.CloudAccountList, error)
	// find cloud-accounts by company id to create dns managed zone.
	// this will be used to populate drop down of cloud-accounts in create dns managed zone form.
	// the response should only include cloud-accounts that a company is authorised to create dns managed zone.
	// the authorization is verified by looking up cloud-accounts with `company-dns-managed-zone-creator` relation for the company id provided in input.
	// the response should only include public attributes of a cloud-account. all non-public attributes should be excluded from the response.
	FindDnsZoneCreateCloudAccounts(ctx context.Context, in *model1.CompanyId, opts ...grpc.CallOption) (*model.CloudAccountList, error)
}

type cloudAccountQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudAccountQueryControllerClient(cc grpc.ClientConnInterface) CloudAccountQueryControllerClient {
	return &cloudAccountQueryControllerClient{cc}
}

func (c *cloudAccountQueryControllerClient) GetById(ctx context.Context, in *model.CloudAccountId, opts ...grpc.CallOption) (*model.CloudAccount, error) {
	out := new(model.CloudAccount)
	err := c.cc.Invoke(ctx, CloudAccountQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAccountQueryControllerClient) FindByCompanyId(ctx context.Context, in *model1.CompanyId, opts ...grpc.CallOption) (*model.CloudAccountList, error) {
	out := new(model.CloudAccountList)
	err := c.cc.Invoke(ctx, CloudAccountQueryController_FindByCompanyId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAccountQueryControllerClient) FindArtifactStoreCreateCloudAccounts(ctx context.Context, in *model1.CompanyId, opts ...grpc.CallOption) (*model.CloudAccountList, error) {
	out := new(model.CloudAccountList)
	err := c.cc.Invoke(ctx, CloudAccountQueryController_FindArtifactStoreCreateCloudAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAccountQueryControllerClient) FindKubeClusterCreateCloudAccounts(ctx context.Context, in *model1.CompanyId, opts ...grpc.CallOption) (*model.CloudAccountList, error) {
	out := new(model.CloudAccountList)
	err := c.cc.Invoke(ctx, CloudAccountQueryController_FindKubeClusterCreateCloudAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAccountQueryControllerClient) FindDnsZoneCreateCloudAccounts(ctx context.Context, in *model1.CompanyId, opts ...grpc.CallOption) (*model.CloudAccountList, error) {
	out := new(model.CloudAccountList)
	err := c.cc.Invoke(ctx, CloudAccountQueryController_FindDnsZoneCreateCloudAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudAccountQueryControllerServer is the server API for CloudAccountQueryController service.
// All implementations should embed UnimplementedCloudAccountQueryControllerServer
// for forward compatibility
type CloudAccountQueryControllerServer interface {
	// look up a cloud-account by id
	GetById(context.Context, *model.CloudAccountId) (*model.CloudAccount, error)
	// find cloud-accounts by company id.
	// the response should only include cloud-accounts in a company that the authenticated user account has viewer access to.
	// authorization is handled internally by running get authorized cloud account ids
	FindByCompanyId(context.Context, *model1.CompanyId) (*model.CloudAccountList, error)
	// find cloud-accounts by company id to create artifact store.
	// this will be used to populate drop down of cloud-accounts in create artifact store form.
	// the response should only include cloud-accounts that a company is authorised to create artifact stores.
	// the authorization is verified by looking up cloud-accounts with `company-artifact-creator` relation for the company id provided in input.
	// the response should only include public attributes of a cloud-account. all non-public attributes should be excluded from the response.
	FindArtifactStoreCreateCloudAccounts(context.Context, *model1.CompanyId) (*model.CloudAccountList, error)
	// find cloud-accounts by company id to create kube-cluster.
	// this will be used to populate drop down of cloud-accounts in create kube-cluster form.
	// the response should only include cloud-accounts that a company is authorised to create kube-cluster.
	// the authorization is verified by looking up cloud-accounts with `company-kube-cluster-creator` relation for the company id provided in input.
	// the response should only include public attributes of a cloud-account. all non-public attributes should be excluded from the response.
	FindKubeClusterCreateCloudAccounts(context.Context, *model1.CompanyId) (*model.CloudAccountList, error)
	// find cloud-accounts by company id to create dns managed zone.
	// this will be used to populate drop down of cloud-accounts in create dns managed zone form.
	// the response should only include cloud-accounts that a company is authorised to create dns managed zone.
	// the authorization is verified by looking up cloud-accounts with `company-dns-managed-zone-creator` relation for the company id provided in input.
	// the response should only include public attributes of a cloud-account. all non-public attributes should be excluded from the response.
	FindDnsZoneCreateCloudAccounts(context.Context, *model1.CompanyId) (*model.CloudAccountList, error)
}

// UnimplementedCloudAccountQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCloudAccountQueryControllerServer struct {
}

func (UnimplementedCloudAccountQueryControllerServer) GetById(context.Context, *model.CloudAccountId) (*model.CloudAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedCloudAccountQueryControllerServer) FindByCompanyId(context.Context, *model1.CompanyId) (*model.CloudAccountList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCompanyId not implemented")
}
func (UnimplementedCloudAccountQueryControllerServer) FindArtifactStoreCreateCloudAccounts(context.Context, *model1.CompanyId) (*model.CloudAccountList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindArtifactStoreCreateCloudAccounts not implemented")
}
func (UnimplementedCloudAccountQueryControllerServer) FindKubeClusterCreateCloudAccounts(context.Context, *model1.CompanyId) (*model.CloudAccountList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindKubeClusterCreateCloudAccounts not implemented")
}
func (UnimplementedCloudAccountQueryControllerServer) FindDnsZoneCreateCloudAccounts(context.Context, *model1.CompanyId) (*model.CloudAccountList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDnsZoneCreateCloudAccounts not implemented")
}

// UnsafeCloudAccountQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudAccountQueryControllerServer will
// result in compilation errors.
type UnsafeCloudAccountQueryControllerServer interface {
	mustEmbedUnimplementedCloudAccountQueryControllerServer()
}

func RegisterCloudAccountQueryControllerServer(s grpc.ServiceRegistrar, srv CloudAccountQueryControllerServer) {
	s.RegisterService(&CloudAccountQueryController_ServiceDesc, srv)
}

func _CloudAccountQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CloudAccountId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAccountQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudAccountQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAccountQueryControllerServer).GetById(ctx, req.(*model.CloudAccountId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAccountQueryController_FindByCompanyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.CompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAccountQueryControllerServer).FindByCompanyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudAccountQueryController_FindByCompanyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAccountQueryControllerServer).FindByCompanyId(ctx, req.(*model1.CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAccountQueryController_FindArtifactStoreCreateCloudAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.CompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAccountQueryControllerServer).FindArtifactStoreCreateCloudAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudAccountQueryController_FindArtifactStoreCreateCloudAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAccountQueryControllerServer).FindArtifactStoreCreateCloudAccounts(ctx, req.(*model1.CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAccountQueryController_FindKubeClusterCreateCloudAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.CompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAccountQueryControllerServer).FindKubeClusterCreateCloudAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudAccountQueryController_FindKubeClusterCreateCloudAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAccountQueryControllerServer).FindKubeClusterCreateCloudAccounts(ctx, req.(*model1.CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAccountQueryController_FindDnsZoneCreateCloudAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.CompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAccountQueryControllerServer).FindDnsZoneCreateCloudAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudAccountQueryController_FindDnsZoneCreateCloudAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAccountQueryControllerServer).FindDnsZoneCreateCloudAccounts(ctx, req.(*model1.CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudAccountQueryController_ServiceDesc is the grpc.ServiceDesc for CloudAccountQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudAccountQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.cloudaccount.service.CloudAccountQueryController",
	HandlerType: (*CloudAccountQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getById",
			Handler:    _CloudAccountQueryController_GetById_Handler,
		},
		{
			MethodName: "findByCompanyId",
			Handler:    _CloudAccountQueryController_FindByCompanyId_Handler,
		},
		{
			MethodName: "findArtifactStoreCreateCloudAccounts",
			Handler:    _CloudAccountQueryController_FindArtifactStoreCreateCloudAccounts_Handler,
		},
		{
			MethodName: "findKubeClusterCreateCloudAccounts",
			Handler:    _CloudAccountQueryController_FindKubeClusterCreateCloudAccounts_Handler,
		},
		{
			MethodName: "findDnsZoneCreateCloudAccounts",
			Handler:    _CloudAccountQueryController_FindDnsZoneCreateCloudAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/cloudaccount/service/query.proto",
}
