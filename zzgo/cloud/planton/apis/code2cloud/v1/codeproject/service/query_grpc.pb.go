// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/code2cloud/v1/codeproject/service/query.proto

package service

import (
	context "context"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/codeproject/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/codeserver/model"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/rpc"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/resourcemanager/v1/product/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CodeProjectQueryController_List_FullMethodName                                = "/cloud.planton.apis.code2cloud.v1.codeproject.service.CodeProjectQueryController/list"
	CodeProjectQueryController_GetById_FullMethodName                             = "/cloud.planton.apis.code2cloud.v1.codeproject.service.CodeProjectQueryController/getById"
	CodeProjectQueryController_FindByProductId_FullMethodName                     = "/cloud.planton.apis.code2cloud.v1.codeproject.service.CodeProjectQueryController/findByProductId"
	CodeProjectQueryController_FindByCodeServerId_FullMethodName                  = "/cloud.planton.apis.code2cloud.v1.codeproject.service.CodeProjectQueryController/findByCodeServerId"
	CodeProjectQueryController_FindByCodeProjectUrl_FullMethodName                = "/cloud.planton.apis.code2cloud.v1.codeproject.service.CodeProjectQueryController/findByCodeProjectUrl"
	CodeProjectQueryController_FindTemplateCodeProjectsByProductId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.codeproject.service.CodeProjectQueryController/findTemplateCodeProjectsByProductId"
	CodeProjectQueryController_GetCodeProjectProfileById_FullMethodName           = "/cloud.planton.apis.code2cloud.v1.codeproject.service.CodeProjectQueryController/getCodeProjectProfileById"
)

// CodeProjectQueryControllerClient is the client API for CodeProjectQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeProjectQueryControllerClient interface {
	// list all code projects for the requested page.
	List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.CodeProjectList, error)
	// look up a code project by code project id
	GetById(ctx context.Context, in *model.CodeProjectId, opts ...grpc.CallOption) (*model.CodeProject, error)
	// find code projects by product id
	FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.CodeProjects, error)
	// find code projects by code-server id
	FindByCodeServerId(ctx context.Context, in *model2.CodeServerId, opts ...grpc.CallOption) (*model.CodeProjects, error)
	// look up all code projects by code project url
	FindByCodeProjectUrl(ctx context.Context, in *model.CodeProjectUrl, opts ...grpc.CallOption) (*model.CodeProjects, error)
	// find code project templates by company id to create new code project.
	// this will be used to populate drop down of code project templates in create code project form.
	// the response should only include code project templates that a company is authorised to use for creating new code projects.
	// the authorization is verified by looking up code project template with `code-project-template-user` relation for the company id provided in input.
	FindTemplateCodeProjectsByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.CodeProjects, error)
	// get code project profile by code project id
	GetCodeProjectProfileById(ctx context.Context, in *model.CodeProjectId, opts ...grpc.CallOption) (*model.CodeProjectProfile, error)
}

type codeProjectQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeProjectQueryControllerClient(cc grpc.ClientConnInterface) CodeProjectQueryControllerClient {
	return &codeProjectQueryControllerClient{cc}
}

func (c *codeProjectQueryControllerClient) List(ctx context.Context, in *rpc.PageInfo, opts ...grpc.CallOption) (*model.CodeProjectList, error) {
	out := new(model.CodeProjectList)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectQueryControllerClient) GetById(ctx context.Context, in *model.CodeProjectId, opts ...grpc.CallOption) (*model.CodeProject, error) {
	out := new(model.CodeProject)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectQueryControllerClient) FindByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.CodeProjects, error) {
	out := new(model.CodeProjects)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_FindByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectQueryControllerClient) FindByCodeServerId(ctx context.Context, in *model2.CodeServerId, opts ...grpc.CallOption) (*model.CodeProjects, error) {
	out := new(model.CodeProjects)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_FindByCodeServerId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectQueryControllerClient) FindByCodeProjectUrl(ctx context.Context, in *model.CodeProjectUrl, opts ...grpc.CallOption) (*model.CodeProjects, error) {
	out := new(model.CodeProjects)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_FindByCodeProjectUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectQueryControllerClient) FindTemplateCodeProjectsByProductId(ctx context.Context, in *model1.ProductId, opts ...grpc.CallOption) (*model.CodeProjects, error) {
	out := new(model.CodeProjects)
	err := c.cc.Invoke(ctx, CodeProjectQueryController_FindTemplateCodeProjectsByProductId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectQueryControllerClient) GetCodeProjectProfileById(ctx context.Context, in *model.CodeProjectId, opts ...grpc.CallOption) (*model.CodeProjectProfile, error) {
	out := new(model.CodeProjectProfile)
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
	// list all code projects for the requested page.
	List(context.Context, *rpc.PageInfo) (*model.CodeProjectList, error)
	// look up a code project by code project id
	GetById(context.Context, *model.CodeProjectId) (*model.CodeProject, error)
	// find code projects by product id
	FindByProductId(context.Context, *model1.ProductId) (*model.CodeProjects, error)
	// find code projects by code-server id
	FindByCodeServerId(context.Context, *model2.CodeServerId) (*model.CodeProjects, error)
	// look up all code projects by code project url
	FindByCodeProjectUrl(context.Context, *model.CodeProjectUrl) (*model.CodeProjects, error)
	// find code project templates by company id to create new code project.
	// this will be used to populate drop down of code project templates in create code project form.
	// the response should only include code project templates that a company is authorised to use for creating new code projects.
	// the authorization is verified by looking up code project template with `code-project-template-user` relation for the company id provided in input.
	FindTemplateCodeProjectsByProductId(context.Context, *model1.ProductId) (*model.CodeProjects, error)
	// get code project profile by code project id
	GetCodeProjectProfileById(context.Context, *model.CodeProjectId) (*model.CodeProjectProfile, error)
}

// UnimplementedCodeProjectQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCodeProjectQueryControllerServer struct {
}

func (UnimplementedCodeProjectQueryControllerServer) List(context.Context, *rpc.PageInfo) (*model.CodeProjectList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCodeProjectQueryControllerServer) GetById(context.Context, *model.CodeProjectId) (*model.CodeProject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedCodeProjectQueryControllerServer) FindByProductId(context.Context, *model1.ProductId) (*model.CodeProjects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductId not implemented")
}
func (UnimplementedCodeProjectQueryControllerServer) FindByCodeServerId(context.Context, *model2.CodeServerId) (*model.CodeProjects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCodeServerId not implemented")
}
func (UnimplementedCodeProjectQueryControllerServer) FindByCodeProjectUrl(context.Context, *model.CodeProjectUrl) (*model.CodeProjects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCodeProjectUrl not implemented")
}
func (UnimplementedCodeProjectQueryControllerServer) FindTemplateCodeProjectsByProductId(context.Context, *model1.ProductId) (*model.CodeProjects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTemplateCodeProjectsByProductId not implemented")
}
func (UnimplementedCodeProjectQueryControllerServer) GetCodeProjectProfileById(context.Context, *model.CodeProjectId) (*model.CodeProjectProfile, error) {
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
	in := new(rpc.PageInfo)
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
		return srv.(CodeProjectQueryControllerServer).List(ctx, req.(*rpc.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CodeProjectId)
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
		return srv.(CodeProjectQueryControllerServer).GetById(ctx, req.(*model.CodeProjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectQueryController_FindByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ProductId)
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
		return srv.(CodeProjectQueryControllerServer).FindByProductId(ctx, req.(*model1.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectQueryController_FindByCodeServerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model2.CodeServerId)
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
		return srv.(CodeProjectQueryControllerServer).FindByCodeServerId(ctx, req.(*model2.CodeServerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectQueryController_FindByCodeProjectUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CodeProjectUrl)
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
		return srv.(CodeProjectQueryControllerServer).FindByCodeProjectUrl(ctx, req.(*model.CodeProjectUrl))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectQueryController_FindTemplateCodeProjectsByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model1.ProductId)
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
		return srv.(CodeProjectQueryControllerServer).FindTemplateCodeProjectsByProductId(ctx, req.(*model1.ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectQueryController_GetCodeProjectProfileById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CodeProjectId)
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
		return srv.(CodeProjectQueryControllerServer).GetCodeProjectProfileById(ctx, req.(*model.CodeProjectId))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeProjectQueryController_ServiceDesc is the grpc.ServiceDesc for CodeProjectQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeProjectQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.codeproject.service.CodeProjectQueryController",
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
	Metadata: "cloud/planton/apis/code2cloud/v1/codeproject/service/query.proto",
}

const (
	CodeProjectPipelineQueryController_GenerateCodePipelineTemplate_FullMethodName    = "/cloud.planton.apis.code2cloud.v1.codeproject.service.CodeProjectPipelineQueryController/generateCodePipelineTemplate"
	CodeProjectPipelineQueryController_GetPipelineFilesByCodeProjectId_FullMethodName = "/cloud.planton.apis.code2cloud.v1.codeproject.service.CodeProjectPipelineQueryController/getPipelineFilesByCodeProjectId"
)

// CodeProjectPipelineQueryControllerClient is the client API for CodeProjectPipelineQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeProjectPipelineQueryControllerClient interface {
	// generate code pipeline template for a code project
	GenerateCodePipelineTemplate(ctx context.Context, in *model.GenerateCodePipelineTemplateQueryInput, opts ...grpc.CallOption) (*model.GenerateCodePipelineTemplateQueryResp, error)
	// get code-pipeline template files for a code project
	GetPipelineFilesByCodeProjectId(ctx context.Context, in *model.CodeProjectId, opts ...grpc.CallOption) (*model.CodeProjectFiles, error)
}

type codeProjectPipelineQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeProjectPipelineQueryControllerClient(cc grpc.ClientConnInterface) CodeProjectPipelineQueryControllerClient {
	return &codeProjectPipelineQueryControllerClient{cc}
}

func (c *codeProjectPipelineQueryControllerClient) GenerateCodePipelineTemplate(ctx context.Context, in *model.GenerateCodePipelineTemplateQueryInput, opts ...grpc.CallOption) (*model.GenerateCodePipelineTemplateQueryResp, error) {
	out := new(model.GenerateCodePipelineTemplateQueryResp)
	err := c.cc.Invoke(ctx, CodeProjectPipelineQueryController_GenerateCodePipelineTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeProjectPipelineQueryControllerClient) GetPipelineFilesByCodeProjectId(ctx context.Context, in *model.CodeProjectId, opts ...grpc.CallOption) (*model.CodeProjectFiles, error) {
	out := new(model.CodeProjectFiles)
	err := c.cc.Invoke(ctx, CodeProjectPipelineQueryController_GetPipelineFilesByCodeProjectId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeProjectPipelineQueryControllerServer is the server API for CodeProjectPipelineQueryController service.
// All implementations should embed UnimplementedCodeProjectPipelineQueryControllerServer
// for forward compatibility
type CodeProjectPipelineQueryControllerServer interface {
	// generate code pipeline template for a code project
	GenerateCodePipelineTemplate(context.Context, *model.GenerateCodePipelineTemplateQueryInput) (*model.GenerateCodePipelineTemplateQueryResp, error)
	// get code-pipeline template files for a code project
	GetPipelineFilesByCodeProjectId(context.Context, *model.CodeProjectId) (*model.CodeProjectFiles, error)
}

// UnimplementedCodeProjectPipelineQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedCodeProjectPipelineQueryControllerServer struct {
}

func (UnimplementedCodeProjectPipelineQueryControllerServer) GenerateCodePipelineTemplate(context.Context, *model.GenerateCodePipelineTemplateQueryInput) (*model.GenerateCodePipelineTemplateQueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateCodePipelineTemplate not implemented")
}
func (UnimplementedCodeProjectPipelineQueryControllerServer) GetPipelineFilesByCodeProjectId(context.Context, *model.CodeProjectId) (*model.CodeProjectFiles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipelineFilesByCodeProjectId not implemented")
}

// UnsafeCodeProjectPipelineQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeProjectPipelineQueryControllerServer will
// result in compilation errors.
type UnsafeCodeProjectPipelineQueryControllerServer interface {
	mustEmbedUnimplementedCodeProjectPipelineQueryControllerServer()
}

func RegisterCodeProjectPipelineQueryControllerServer(s grpc.ServiceRegistrar, srv CodeProjectPipelineQueryControllerServer) {
	s.RegisterService(&CodeProjectPipelineQueryController_ServiceDesc, srv)
}

func _CodeProjectPipelineQueryController_GenerateCodePipelineTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GenerateCodePipelineTemplateQueryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectPipelineQueryControllerServer).GenerateCodePipelineTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectPipelineQueryController_GenerateCodePipelineTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectPipelineQueryControllerServer).GenerateCodePipelineTemplate(ctx, req.(*model.GenerateCodePipelineTemplateQueryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeProjectPipelineQueryController_GetPipelineFilesByCodeProjectId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.CodeProjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeProjectPipelineQueryControllerServer).GetPipelineFilesByCodeProjectId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeProjectPipelineQueryController_GetPipelineFilesByCodeProjectId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeProjectPipelineQueryControllerServer).GetPipelineFilesByCodeProjectId(ctx, req.(*model.CodeProjectId))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeProjectPipelineQueryController_ServiceDesc is the grpc.ServiceDesc for CodeProjectPipelineQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeProjectPipelineQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.code2cloud.v1.codeproject.service.CodeProjectPipelineQueryController",
	HandlerType: (*CodeProjectPipelineQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "generateCodePipelineTemplate",
			Handler:    _CodeProjectPipelineQueryController_GenerateCodePipelineTemplate_Handler,
		},
		{
			MethodName: "getPipelineFilesByCodeProjectId",
			Handler:    _CodeProjectPipelineQueryController_GetPipelineFilesByCodeProjectId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/code2cloud/v1/codeproject/service/query.proto",
}
