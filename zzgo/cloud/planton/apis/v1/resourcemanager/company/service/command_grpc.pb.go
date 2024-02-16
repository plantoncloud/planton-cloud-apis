// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/resourcemanager/company/service/command.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/company/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CompanyCommandController_Create_FullMethodName  = "/cloud.planton.apis.v1.resourcemanager.company.service.CompanyCommandController/create"
	CompanyCommandController_Update_FullMethodName  = "/cloud.planton.apis.v1.resourcemanager.company.service.CompanyCommandController/update"
	CompanyCommandController_Delete_FullMethodName  = "/cloud.planton.apis.v1.resourcemanager.company.service.CompanyCommandController/delete"
	CompanyCommandController_Restore_FullMethodName = "/cloud.planton.apis.v1.resourcemanager.company.service.CompanyCommandController/restore"
)

// CompanyCommandControllerClient is the client API for CompanyCommandController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyCommandControllerClient interface {
	// create a new company on planton cloud
	Create(ctx context.Context, in *model.Company, opts ...grpc.CallOption) (*model.Company, error)
	// update an existing company on planton cloud
	Update(ctx context.Context, in *model.Company, opts ...grpc.CallOption) (*model.Company, error)
	// delete an existing company on planton cloud using company id
	Delete(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.Company, error)
	// restore a previously deleted company.
	Restore(ctx context.Context, in *model.Company, opts ...grpc.CallOption) (*model.Company, error)
}

type companyCommandControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyCommandControllerClient(cc grpc.ClientConnInterface) CompanyCommandControllerClient {
	return &companyCommandControllerClient{cc}
}

func (c *companyCommandControllerClient) Create(ctx context.Context, in *model.Company, opts ...grpc.CallOption) (*model.Company, error) {
	out := new(model.Company)
	err := c.cc.Invoke(ctx, CompanyCommandController_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyCommandControllerClient) Update(ctx context.Context, in *model.Company, opts ...grpc.CallOption) (*model.Company, error) {
	out := new(model.Company)
	err := c.cc.Invoke(ctx, CompanyCommandController_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyCommandControllerClient) Delete(ctx context.Context, in *model1.ApiResourceRefreshCommandInput, opts ...grpc.CallOption) (*model.Company, error) {
	out := new(model.Company)
	err := c.cc.Invoke(ctx, CompanyCommandController_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyCommandControllerClient) Restore(ctx context.Context, in *model.Company, opts ...grpc.CallOption) (*model.Company, error) {
	out := new(model.Company)
	err := c.cc.Invoke(ctx, CompanyCommandController_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyCommandControllerServer is the server API for CompanyCommandController service.
// All implementations should embed UnimplementedCompanyCommandControllerServer
// for forward compatibility
type CompanyCommandControllerServer interface {
	// create a new company on planton cloud
	Create(context.Context, *model.Company) (*model.Company, error)
	// update an existing company on planton cloud
	Update(context.Context, *model.Company) (*model.Company, error)
	// delete an existing company on planton cloud using company id
	Delete(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.Company, error)
	// restore a previously deleted company.
	Restore(context.Context, *model.Company) (*model.Company, error)
}

// UnimplementedCompanyCommandControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCompanyCommandControllerServer struct {
}

func (UnimplementedCompanyCommandControllerServer) Create(context.Context, *model.Company) (*model.Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCompanyCommandControllerServer) Update(context.Context, *model.Company) (*model.Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCompanyCommandControllerServer) Delete(context.Context, *model1.ApiResourceRefreshCommandInput) (*model.Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCompanyCommandControllerServer) Restore(context.Context, *model.Company) (*model.Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeCompanyCommandControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyCommandControllerServer will
// result in compilation errors.
type UnsafeCompanyCommandControllerServer interface {
	mustEmbedUnimplementedCompanyCommandControllerServer()
}

func RegisterCompanyCommandControllerServer(s grpc.ServiceRegistrar, srv CompanyCommandControllerServer) {
	s.RegisterService(&CompanyCommandController_ServiceDesc, srv)
}

func _CompanyCommandController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Company)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyCommandControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyCommandController_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyCommandControllerServer).Create(ctx, req.(*model.Company))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyCommandController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Company)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyCommandControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyCommandController_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyCommandControllerServer).Update(ctx, req.(*model.Company))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyCommandController_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ApiResourceRefreshCommandInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyCommandControllerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyCommandController_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyCommandControllerServer).Delete(ctx, req.(*model1.ApiResourceRefreshCommandInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyCommandController_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Company)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyCommandControllerServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyCommandController_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyCommandControllerServer).Restore(ctx, req.(*model.Company))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyCommandController_ServiceDesc is the grpc.ServiceDesc for CompanyCommandController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyCommandController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.resourcemanager.company.service.CompanyCommandController",
	HandlerType: (*CompanyCommandControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _CompanyCommandController_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _CompanyCommandController_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _CompanyCommandController_Delete_Handler,
		},
		{
			MethodName: "restore",
			Handler:    _CompanyCommandController_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/resourcemanager/company/service/command.proto",
}
