// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/sourcecode/project/service/query.proto

package service

import (
	context "context"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/sourcecode/project/model"
	model3 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/sourcecode/server/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/product/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CodeProjectQueryController_List_FullMethodName                                = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectQueryController/list"
	CodeProjectQueryController_GetById_FullMethodName                             = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectQueryController/getById"
	CodeProjectQueryController_FindByProductId_FullMethodName                     = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectQueryController/findByProductId"
	CodeProjectQueryController_FindByCodeServerId_FullMethodName                  = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectQueryController/findByCodeServerId"
	CodeProjectQueryController_FindByCodeProjectUrl_FullMethodName                = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectQueryController/findByCodeProjectUrl"
	CodeProjectQueryController_FindTemplateCodeProjectsByProductId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectQueryController/findTemplateCodeProjectsByProductId"
	CodeProjectQueryController_GetCodeProjectProfileById_FullMethodName           = "/cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectQueryController/getCodeProjectProfileById"
)

// CodeProjectQueryControllerClient is the client API for CodeProjectQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeProjectQueryControllerClient interface {
	// list all code projects on planton cloud for the requested page. This is intended for use on portal.
	List(ctx context.Context, in *model.PageInfo, opts ...grpc.CallOption) (*model1.CodeProjectList, error)
	// look up a code project by code project id
	GetById(ctx context.Context, in *model1.CodeProjectId, opts ...grpc.CallOption) (*model1.CodeProject, error)
	// find code projects by product id
	FindByProductId(ctx context.Context, in *model2.ProductId, opts ...grpc.CallOption) (*model1.CodeProjects, error)
	// find code projects by code-server id
	FindByCodeServerId(ctx context.Context, in *model3.CodeServerId, opts ...grpc.CallOption) (*model1.CodeProjects, error)
	// look up all code projects by code project url
	FindByCodeProjectUrl(ctx context.Context, in *model1.CodeProjectUrl, opts ...grpc.CallOption) (*model1.CodeProjects, error)
	// find code project templates by company id to create new code project.
	// this will be used to populate drop down of code project templates in create code project form.
	// the response should only include code project templates that a company is authorised to use for creating new code projects.
	// the authorization is verified by looking up code project template with `code-project-template-user` relation for the company id provided in input.
	FindTemplateCodeProjectsByProductId(ctx context.Context, in *model2.ProductId, opts ...grpc.CallOption) (*model1.CodeProjects, error)
	// get code project profile by code project id
	GetCodeProjectProfileById(ctx context.Context, in *model1.CodeProjectId, opts ...grpc.CallOption) (*model1.CodeProjectProfile, error)
}

type codeProjectQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeProjectQueryControllerClient(cc grpc.ClientConnInterface) CodeProjectQueryControllerClient {
	return &codeProjectQueryControllerClient{cc}
}

func (c *codeProjectQueryControllerClient) List(ctx context.Context, in *model.PageInfo, opts ...grpc.CallOption) (*model1.CodeProjectList, error) {
	out := new(model1.CodeProjectList)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectQueryControllerClient) GetById(ctx context.Context, in *model1.CodeProjectId, opts ...grpc.CallOption) (*model1.CodeProject, error) {
	out := new(model1.CodeProject)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectQueryControllerClient) FindByProductId(ctx context.Context, in *model2.ProductId, opts ...grpc.CallOption) (*model1.CodeProjects, error) {
	out := new(model1.CodeProjects)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectQueryControllerClient) FindByCodeServerId(ctx context.Context, in *model3.CodeServerId, opts ...grpc.CallOption) (*model1.CodeProjects, error) {
	out := new(model1.CodeProjects)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_FindByCodeServerId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectQueryControllerClient) FindByCodeProjectUrl(ctx context.Context, in *model1.CodeProjectUrl, opts ...grpc.CallOption) (*model1.CodeProjects, error) {
	out := new(model1.CodeProjects)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_FindByCodeProjectUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectQueryControllerClient) FindTemplateCodeProjectsByProductId(ctx context.Context, in *model2.ProductId, opts ...grpc.CallOption) (*model1.CodeProjects, error) {
	out := new(model1.CodeProjects)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_FindTemplateCodeProjectsByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectQueryControllerClient) GetCodeProjectProfileById(ctx context.Context, in *model1.CodeProjectId, opts ...grpc.CallOption) (*model1.CodeProjectProfile, error) {
	out := new(model1.CodeProjectProfile)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_GetCodeProjectProfileById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeProjectQueryControllerServer is the server API for CodeProjectQueryController service.
// All implementations should embed UnimplementedCodeProjectQueryControllerServer
// for forward compatibility
type CodeProjectQueryControllerServer interface {
	// list all code projects on planton cloud for the requested page. This is intended for use on portal.
	List(context.Context, *model.PageInfo) (*model1.CodeProjectList, error)
	// look up a code project by code project id
	GetById(context.Context, *model1.CodeProjectId) (*model1.CodeProject, error)
	// find code projects by product id
	FindByProductId(context.Context, *model2.ProductId) (*model1.CodeProjects, error)
	// find code projects by code-server id
	FindByCodeServerId(context.Context, *model3.CodeServerId) (*model1.CodeProjects, error)
	// look up all code projects by code project url
	FindByCodeProjectUrl(context.Context, *model1.CodeProjectUrl) (*model1.CodeProjects, error)
	// find code project templates by company id to create new code project.
	// this will be used to populate drop down of code project templates in create code project form.
	// the response should only include code project templates that a company is authorised to use for creating new code projects.
	// the authorization is verified by looking up code project template with `code-project-template-user` relation for the company id provided in input.
	FindTemplateCodeProjectsByProductId(context.Context, *model2.ProductId) (*model1.CodeProjects, error)
	// get code project profile by code project id
	GetCodeProjectProfileById(context.Context, *model1.CodeProjectId) (*model1.CodeProjectProfile, error)
}

// UnimplementedCodeProjectQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCodeProjectQueryControllerServer struct {
}

func (UnimplementedCodeProjectQueryControllerServer) List(context.Context, *model.PageInfo) (*model1.CodeProjectList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCodeProjectQueryControllerServer) GetById(context.Context, *model1.CodeProjectId) (*model1.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedCodeProjectQueryControllerServer) FindByProductId(context.Context, *model2.ProductId) (*model1.CodeProjects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedCodeProjectQueryControllerServer) FindByCodeServerId(context.Context, *model3.CodeServerId) (*model1.CodeProjects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCodeServerId not implemented")
}
func (UnimplementedCodeProjectQueryControllerServer) FindByCodeProjectUrl(context.Context, *model1.CodeProjectUrl) (*model1.CodeProjects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCodeProjectUrl not implemented")
}
func (UnimplementedCodeProjectQueryControllerServer) FindTemplateCodeProjectsByProductId(context.Context, *model2.ProductId) (*model1.CodeProjects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTemplateCodeProjectsByProductId not implemented")
}
func (UnimplementedCodeProjectQueryControllerServer) GetCodeProjectProfileById(context.Context, *model1.CodeProjectId) (*model1.CodeProjectProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCodeProjectProfileById not implemented")
}

// UnsafeCodeProjectQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeProjectQueryControllerServer will
// result in compilation errors.
type UnsafeCodeProjectQueryControllerServer interface {
	mustEmbedUnimplementedCodeProjectQueryControllerServer()
}

func RegisterCodeProjectQueryControllerServer(s grpc.ServiceRegistrar, srv CodeProjectQueryControllerServer) {
	s.RegisterService(&CodeProjectQueryController_ServiceDesc, srv)
}

func _CodeProjectQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectQueryControllerServer).List(ctx, req.(*model.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.CodeProjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectQueryControllerServer).GetById(ctx, req.(*model1.CodeProjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectQueryControllerServer).FindByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectQueryController_FindByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectQueryControllerServer).FindByProductId(ctx, req.(*model2.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectQueryController_FindByCodeServerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model3.CodeServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectQueryControllerServer).FindByCodeServerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectQueryController_FindByCodeServerId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectQueryControllerServer).FindByCodeServerId(ctx, req.(*model3.CodeServerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectQueryController_FindByCodeProjectUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.CodeProjectUrl)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectQueryControllerServer).FindByCodeProjectUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectQueryController_FindByCodeProjectUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectQueryControllerServer).FindByCodeProjectUrl(ctx, req.(*model1.CodeProjectUrl))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectQueryController_FindTemplateCodeProjectsByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectQueryControllerServer).FindTemplateCodeProjectsByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectQueryController_FindTemplateCodeProjectsByProductId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectQueryControllerServer).FindTemplateCodeProjectsByProductId(ctx, req.(*model2.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectQueryController_GetCodeProjectProfileById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.CodeProjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectQueryControllerServer).GetCodeProjectProfileById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectQueryController_GetCodeProjectProfileById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectQueryControllerServer).GetCodeProjectProfileById(ctx, req.(*model1.CodeProjectId))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeProjectQueryController_ServiceDesc is the grpc.ServiceDesc for CodeProjectQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeProjectQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectQueryController",
	HandlerType: (*CodeProjectQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _CodeProjectQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _CodeProjectQueryController_GetById_Handler,
		},
		{
			MethodName: "findByProductId",
			Handler:    _CodeProjectQueryController_FindByProductId_Handler,
		},
		{
			MethodName: "findByCodeServerId",
			Handler:    _CodeProjectQueryController_FindByCodeServerId_Handler,
		},
		{
			MethodName: "findByCodeProjectUrl",
			Handler:    _CodeProjectQueryController_FindByCodeProjectUrl_Handler,
		},
		{
			MethodName: "findTemplateCodeProjectsByProductId",
			Handler:    _CodeProjectQueryController_FindTemplateCodeProjectsByProductId_Handler,
		},
		{
			MethodName: "getCodeProjectProfileById",
			Handler:    _CodeProjectQueryController_GetCodeProjectProfileById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/sourcecode/project/service/query.proto",
}
