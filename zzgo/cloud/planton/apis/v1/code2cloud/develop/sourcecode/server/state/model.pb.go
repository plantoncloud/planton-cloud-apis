// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/develop/sourcecode/server/state/model.proto

package state

import (
	enums "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/develop/sourcecode/server/state/enums"
	audit "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/audit"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/options"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/state"
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

// code-server state
type CodeServerState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// event-type
	EventType enums.CodeServerEventType `protobuf:"varint,99,opt,name=event_type,json=eventType,proto3,enum=cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.enums.CodeServerEventType" json:"event_type,omitempty"`
	// resource api version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// resource metadata
	Metadata *resource.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *CodeServerSpecState `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *CodeServerStatusState `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CodeServerState) Reset() {
	*x = CodeServerState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeServerState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeServerState) ProtoMessage() {}

func (x *CodeServerState) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeServerState.ProtoReflect.Descriptor instead.
func (*CodeServerState) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDescGZIP(), []int{0}
}

func (x *CodeServerState) GetEventType() enums.CodeServerEventType {
	if x != nil {
		return x.EventType
	}
	return enums.CodeServerEventType(0)
}

func (x *CodeServerState) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *CodeServerState) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CodeServerState) GetMetadata() *resource.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CodeServerState) GetSpec() *CodeServerSpecState {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *CodeServerState) GetStatus() *CodeServerStatusState {
	if x != nil {
		return x.Status
	}
	return nil
}

// code-server spec state
type CodeServerSpecState struct {
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
	CodeServerProvider string `protobuf:"bytes,4,opt,name=code_server_provider,json=codeServerProvider,proto3" json:"code_server_provider,omitempty"`
	// github spec
	Github *CodeServerGithubSpecState `protobuf:"bytes,5,opt,name=github,proto3" json:"github,omitempty"`
	// gitlab spec
	Gitlab *CodeServerGitlabSpecState `protobuf:"bytes,6,opt,name=gitlab,proto3" json:"gitlab,omitempty"`
}

func (x *CodeServerSpecState) Reset() {
	*x = CodeServerSpecState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeServerSpecState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeServerSpecState) ProtoMessage() {}

func (x *CodeServerSpecState) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeServerSpecState.ProtoReflect.Descriptor instead.
func (*CodeServerSpecState) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDescGZIP(), []int{1}
}

func (x *CodeServerSpecState) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CodeServerSpecState) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CodeServerSpecState) GetCodeServerHost() string {
	if x != nil {
		return x.CodeServerHost
	}
	return ""
}

func (x *CodeServerSpecState) GetCodeServerProvider() string {
	if x != nil {
		return x.CodeServerProvider
	}
	return ""
}

func (x *CodeServerSpecState) GetGithub() *CodeServerGithubSpecState {
	if x != nil {
		return x.Github
	}
	return nil
}

func (x *CodeServerSpecState) GetGitlab() *CodeServerGitlabSpecState {
	if x != nil {
		return x.Gitlab
	}
	return nil
}

// github code-server spec
type CodeServerGithubSpecState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the github organization
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// ID of the Github App installation provided by Github after successful app installation via OAuth flow.
	AppInstallId int64 `protobuf:"varint,2,opt,name=app_install_id,json=appInstallId,proto3" json:"app_install_id,omitempty"`
	// Specifies the owner type for Github App installation: User or Org.
	AppInstallOwnerType string `protobuf:"bytes,3,opt,name=app_install_owner_type,json=appInstallOwnerType,proto3" json:"app_install_owner_type,omitempty"`
}

func (x *CodeServerGithubSpecState) Reset() {
	*x = CodeServerGithubSpecState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeServerGithubSpecState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeServerGithubSpecState) ProtoMessage() {}

func (x *CodeServerGithubSpecState) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeServerGithubSpecState.ProtoReflect.Descriptor instead.
func (*CodeServerGithubSpecState) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDescGZIP(), []int{2}
}

func (x *CodeServerGithubSpecState) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *CodeServerGithubSpecState) GetAppInstallId() int64 {
	if x != nil {
		return x.AppInstallId
	}
	return 0
}

func (x *CodeServerGithubSpecState) GetAppInstallOwnerType() string {
	if x != nil {
		return x.AppInstallOwnerType
	}
	return ""
}

// gitlab code-server spec
type CodeServerGitlabSpecState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name or numerical identifier for gitlab group
	// This value is primarily used in the code project synchronization process.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// Access Token to integrate with the Gitlab server.
	// This value is acquired by browser authorization flow when code server is added via Planton Cloud Console App.
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Refresh token used to fetch a new access token when the current one expires.
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *CodeServerGitlabSpecState) Reset() {
	*x = CodeServerGitlabSpecState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeServerGitlabSpecState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeServerGitlabSpecState) ProtoMessage() {}

func (x *CodeServerGitlabSpecState) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeServerGitlabSpecState.ProtoReflect.Descriptor instead.
func (*CodeServerGitlabSpecState) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDescGZIP(), []int{3}
}

func (x *CodeServerGitlabSpecState) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *CodeServerGitlabSpecState) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CodeServerGitlabSpecState) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

// code-server status state
type CodeServerStatusState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource lifecycle
	Lifecycle *resource.ResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// system audit info
	SysAudit *audit.SysAudit `protobuf:"bytes,98,opt,name=sys_audit,json=sysAudit,proto3" json:"sys_audit,omitempty"`
	// id of the stack-job
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
}

func (x *CodeServerStatusState) Reset() {
	*x = CodeServerStatusState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeServerStatusState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeServerStatusState) ProtoMessage() {}

func (x *CodeServerStatusState) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeServerStatusState.ProtoReflect.Descriptor instead.
func (*CodeServerStatusState) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDescGZIP(), []int{4}
}

func (x *CodeServerStatusState) GetLifecycle() *resource.ResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *CodeServerStatusState) GetSysAudit() *audit.SysAudit {
	if x != nil {
		return x.SysAudit
	}
	return nil
}

func (x *CodeServerStatusState) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

var File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x40,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x1a, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8d, 0x04, 0x0a, 0x0f, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x7a, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x69, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x55, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x70,
	0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x6f, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x57, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x1f,
	0x88, 0xa6, 0x1d, 0x05, 0x90, 0xa6, 0x1d, 0x00, 0x9a, 0xa6, 0x1d, 0x13, 0x08, 0x14, 0x12, 0x0f,
	0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x22,
	0x99, 0x03, 0x0a, 0x13, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x70,
	0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x14, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63,
	0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x73, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x5b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x70, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x12, 0x73, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x53, 0x70, 0x65, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x06, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x22, 0x8d, 0x01, 0x0a, 0x19,
	0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x53, 0x70, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x7e, 0x0a, 0x19, 0x43,
	0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x53,
	0x70, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xde, 0x01, 0x0a, 0x15,
	0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x57, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x4a,
	0x0a, 0x09, 0x73, 0x79, 0x73, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x62, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x52, 0x08, 0x73, 0x79, 0x73, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x61, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x42, 0xf2, 0x03, 0x0a,
	0x44, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x70, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x53, 0x53,
	0x53, 0xaa, 0x02, 0x40, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0xca, 0x02, 0x40, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x5c,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x65, 0xe2, 0x02, 0x4c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x5c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x48, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x3a, 0x3a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_goTypes = []interface{}{
	(*CodeServerState)(nil),            // 0: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerState
	(*CodeServerSpecState)(nil),        // 1: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerSpecState
	(*CodeServerGithubSpecState)(nil),  // 2: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerGithubSpecState
	(*CodeServerGitlabSpecState)(nil),  // 3: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerGitlabSpecState
	(*CodeServerStatusState)(nil),      // 4: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerStatusState
	(enums.CodeServerEventType)(0),     // 5: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.enums.CodeServerEventType
	(*resource.Metadata)(nil),          // 6: cloud.planton.apis.v1.commons.resource.Metadata
	(*resource.ResourceLifecycle)(nil), // 7: cloud.planton.apis.v1.commons.resource.ResourceLifecycle
	(*audit.SysAudit)(nil),             // 8: cloud.planton.apis.v1.commons.audit.SysAudit
}
var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerState.event_type:type_name -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.enums.CodeServerEventType
	6, // 1: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerState.metadata:type_name -> cloud.planton.apis.v1.commons.resource.Metadata
	1, // 2: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerState.spec:type_name -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerSpecState
	4, // 3: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerState.status:type_name -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerStatusState
	2, // 4: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerSpecState.github:type_name -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerGithubSpecState
	3, // 5: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerSpecState.gitlab:type_name -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerGitlabSpecState
	7, // 6: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerStatusState.lifecycle:type_name -> cloud.planton.apis.v1.commons.resource.ResourceLifecycle
	8, // 7: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.state.CodeServerStatusState.sys_audit:type_name -> cloud.planton.apis.v1.commons.audit.SysAudit
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeServerState); i {
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
		file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeServerSpecState); i {
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
		file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeServerGithubSpecState); i {
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
		file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeServerGitlabSpecState); i {
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
		file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeServerStatusState); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_state_model_proto_depIdxs = nil
}
