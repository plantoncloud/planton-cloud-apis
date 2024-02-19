// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/codeserver/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	codeserverprovider "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/codeserver/enums/codeserverprovider"
	githubappinstallownertype "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/codeserver/enums/githubappinstallownertype"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcefieldoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemessageoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemetadataoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// code-server
type CodeServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// resource metadata
	Metadata *model.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *CodeServerSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *CodeServerStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CodeServer) Reset() {
	*x = CodeServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeServer) ProtoMessage() {}

func (x *CodeServer) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeServer.ProtoReflect.Descriptor instead.
func (*CodeServer) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *CodeServer) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *CodeServer) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CodeServer) GetMetadata() *model.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CodeServer) GetSpec() *CodeServerSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *CodeServer) GetStatus() *CodeServerStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// code-server spec
type CodeServerSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the company to which the code server belongs.
	// This value is computed from the product.
	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// Specifies the product to which the code server belongs.
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// Specifies the host of the code server, such as "https://github.com".
	CodeServerHost string `protobuf:"bytes,3,opt,name=code_server_host,json=codeServerHost,proto3" json:"code_server_host,omitempty"`
	// Specifies the provider for the code server.
	CodeServerProvider codeserverprovider.CodeServerProvider `protobuf:"varint,4,opt,name=code_server_provider,json=codeServerProvider,proto3,enum=cloud.planton.apis.code2cloud.v1.codeserver.enums.codeserverprovider.CodeServerProvider" json:"code_server_provider,omitempty"`
	// github spec
	Github *CodeServerGithubSpec `protobuf:"bytes,5,opt,name=github,proto3" json:"github,omitempty"`
	// gitlab spec
	Gitlab *CodeServerGitlabSpec `protobuf:"bytes,6,opt,name=gitlab,proto3" json:"gitlab,omitempty"`
}

func (x *CodeServerSpec) Reset() {
	*x = CodeServerSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeServerSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeServerSpec) ProtoMessage() {}

func (x *CodeServerSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeServerSpec.ProtoReflect.Descriptor instead.
func (*CodeServerSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *CodeServerSpec) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CodeServerSpec) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CodeServerSpec) GetCodeServerHost() string {
	if x != nil {
		return x.CodeServerHost
	}
	return ""
}

func (x *CodeServerSpec) GetCodeServerProvider() codeserverprovider.CodeServerProvider {
	if x != nil {
		return x.CodeServerProvider
	}
	return codeserverprovider.CodeServerProvider(0)
}

func (x *CodeServerSpec) GetGithub() *CodeServerGithubSpec {
	if x != nil {
		return x.Github
	}
	return nil
}

func (x *CodeServerSpec) GetGitlab() *CodeServerGitlabSpec {
	if x != nil {
		return x.Gitlab
	}
	return nil
}

// github code-server spec
type CodeServerGithubSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the github organization
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// ID of the Github App installation provided by Github after successful app installation via OAuth flow.
	AppInstallId int64 `protobuf:"varint,2,opt,name=app_install_id,json=appInstallId,proto3" json:"app_install_id,omitempty"`
	// Specifies the owner type for Github App installation: User or Org.
	AppInstallOwnerType githubappinstallownertype.GithubAppInstallOwnerType `protobuf:"varint,3,opt,name=app_install_owner_type,json=appInstallOwnerType,proto3,enum=cloud.planton.apis.code2cloud.v1.codeserver.enums.githubappinstallownertype.GithubAppInstallOwnerType" json:"app_install_owner_type,omitempty"`
}

func (x *CodeServerGithubSpec) Reset() {
	*x = CodeServerGithubSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeServerGithubSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeServerGithubSpec) ProtoMessage() {}

func (x *CodeServerGithubSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeServerGithubSpec.ProtoReflect.Descriptor instead.
func (*CodeServerGithubSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *CodeServerGithubSpec) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *CodeServerGithubSpec) GetAppInstallId() int64 {
	if x != nil {
		return x.AppInstallId
	}
	return 0
}

func (x *CodeServerGithubSpec) GetAppInstallOwnerType() githubappinstallownertype.GithubAppInstallOwnerType {
	if x != nil {
		return x.AppInstallOwnerType
	}
	return githubappinstallownertype.GithubAppInstallOwnerType(0)
}

// gitlab code-server spec
type CodeServerGitlabSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name or numerical identifier for gitlab group
	// This value is primarily used in the code project synchronization process.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// Access token to integrate with Gitlab Server.
	// This value is acquired by browser authorization flow when code server is added via Planton Cloud Console App.
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Refresh token used to fetch a new access token when the current one expires.
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *CodeServerGitlabSpec) Reset() {
	*x = CodeServerGitlabSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeServerGitlabSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeServerGitlabSpec) ProtoMessage() {}

func (x *CodeServerGitlabSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeServerGitlabSpec.ProtoReflect.Descriptor instead.
func (*CodeServerGitlabSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDescGZIP(), []int{3}
}

func (x *CodeServerGitlabSpec) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *CodeServerGitlabSpec) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CodeServerGitlabSpec) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

// code-server status
type CodeServerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource lifecycle
	Lifecycle *model.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// resource audit info
	Audit *model.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// id of the stack-job
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
}

func (x *CodeServerStatus) Reset() {
	*x = CodeServerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeServerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeServerStatus) ProtoMessage() {}

func (x *CodeServerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeServerStatus.ProtoReflect.Descriptor instead.
func (*CodeServerStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDescGZIP(), []int{4}
}

func (x *CodeServerStatus) GetLifecycle() *model.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *CodeServerStatus) GetAudit() *model.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *CodeServerStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

var File_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x67, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x75, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x05, 0x0a, 0x0a,
	0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0b, 0x61, 0x70,
	0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x22, 0xba, 0x48, 0x1f, 0x72, 0x1d, 0x0a, 0x1b, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x76, 0x31, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x25, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xba,
	0x48, 0x0e, 0x72, 0x0c, 0x0a, 0x0a, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0xf9, 0x02, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x99, 0x02, 0xba,
	0x48, 0x95, 0x02, 0xba, 0x01, 0x6c, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x4f, 0x6e, 0x6c, 0x79, 0x20, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x63, 0x61, 0x73, 0x65, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x68, 0x79, 0x70, 0x68,
	0x65, 0x6e, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x1a,
	0x21, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x2b, 0x24,
	0x27, 0x29, 0xba, 0x01, 0x53, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x4d, 0x75, 0x73, 0x74, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x63, 0x61, 0x73,
	0x65, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x1e, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61,
	0x2d, 0x7a, 0x5d, 0x2e, 0x2a, 0x24, 0x27, 0x29, 0xba, 0x01, 0x47, 0x0a, 0x0d, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x4d, 0x75, 0x73, 0x74,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20,
	0x68, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x1a, 0x1a, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5b, 0x5e, 0x2d, 0x5d, 0x24,
	0x27, 0x29, 0x72, 0x04, 0x10, 0x01, 0x18, 0x0a, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x55, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x5b, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x25, 0x8a, 0xb5, 0x18, 0x02, 0x63, 0x73, 0x98, 0xb5,
	0x18, 0x01, 0x88, 0xa6, 0x1d, 0x06, 0x9a, 0xa6, 0x1d, 0x13, 0x08, 0x13, 0x12, 0x0f, 0x73, 0x70,
	0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x22, 0xeb, 0x03,
	0x0a, 0x0e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x23, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xc0, 0xb8, 0x18, 0x01, 0xd0,
	0xb8, 0x18, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x32,
	0x0a, 0x10, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xc0, 0xb8, 0x18, 0x01, 0xd0, 0xb8,
	0x18, 0x01, 0x52, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x6f,
	0x73, 0x74, 0x12, 0x94, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x58, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x08, 0xc0, 0xb8, 0x18,
	0x01, 0xd0, 0xb8, 0x18, 0x01, 0x52, 0x12, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x5f, 0x0a, 0x06, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x12, 0x5f, 0x0a, 0x06, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x06, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x22, 0x83, 0x02, 0x0a, 0x14,
	0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52,
	0x0c, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0xa1, 0x01,
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x66,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x47, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x13, 0x61, 0x70,
	0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x8b, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1f, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8,
	0x18, 0x01, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x29, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18,
	0x01, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xec, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x60, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x54, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18,
	0x62, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x61, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x42, 0x9e,
	0x03, 0x0a, 0x3f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x61, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x43, 0x56, 0x43, 0x4d, 0xaa, 0x02, 0x31,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70,
	0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0xca, 0x02, 0x31, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x3d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_goTypes = []interface{}{
	(*CodeServer)(nil),                                       // 0: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServer
	(*CodeServerSpec)(nil),                                   // 1: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerSpec
	(*CodeServerGithubSpec)(nil),                             // 2: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerGithubSpec
	(*CodeServerGitlabSpec)(nil),                             // 3: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerGitlabSpec
	(*CodeServerStatus)(nil),                                 // 4: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerStatus
	(*model.ApiResourceMetadata)(nil),                        // 5: cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	(codeserverprovider.CodeServerProvider)(0),               // 6: cloud.planton.apis.code2cloud.v1.codeserver.enums.codeserverprovider.CodeServerProvider
	(githubappinstallownertype.GithubAppInstallOwnerType)(0), // 7: cloud.planton.apis.code2cloud.v1.codeserver.enums.githubappinstallownertype.GithubAppInstallOwnerType
	(*model.ApiResourceLifecycle)(nil),                       // 8: cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	(*model.ApiResourceAudit)(nil),                           // 9: cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
}
var file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServer.metadata:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	1, // 1: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServer.spec:type_name -> cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerSpec
	4, // 2: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServer.status:type_name -> cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerStatus
	6, // 3: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerSpec.code_server_provider:type_name -> cloud.planton.apis.code2cloud.v1.codeserver.enums.codeserverprovider.CodeServerProvider
	2, // 4: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerSpec.github:type_name -> cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerGithubSpec
	3, // 5: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerSpec.gitlab:type_name -> cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerGitlabSpec
	7, // 6: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerGithubSpec.app_install_owner_type:type_name -> cloud.planton.apis.code2cloud.v1.codeserver.enums.githubappinstallownertype.GithubAppInstallOwnerType
	8, // 7: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerStatus.lifecycle:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	9, // 8: cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServerStatus.audit:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeServer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeServerSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeServerGithubSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeServerGitlabSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeServerStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_codeserver_model_state_proto_depIdxs = nil
}
