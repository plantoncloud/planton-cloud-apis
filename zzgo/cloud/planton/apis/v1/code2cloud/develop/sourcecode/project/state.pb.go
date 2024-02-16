// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/develop/sourcecode/project/state.proto

package project

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	audit "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/audit"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/field/options"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/metadata/options"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/options"
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

// code-project
type CodeProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// resource metadata
	Metadata *resource.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *CodeProjectSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *CodeProjectStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CodeProject) Reset() {
	*x = CodeProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeProject) ProtoMessage() {}

func (x *CodeProject) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeProject.ProtoReflect.Descriptor instead.
func (*CodeProject) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDescGZIP(), []int{0}
}

func (x *CodeProject) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *CodeProject) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CodeProject) GetMetadata() *resource.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CodeProject) GetSpec() *CodeProjectSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *CodeProject) GetStatus() *CodeProjectStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// code-project spec
type CodeProjectSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// company to which the code project belongs to.
	// value is computed from product.
	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// product to which the code-project belongs to.
	// this is computed from code-server.
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// code server to which the code-project belongs to.
	CodeServerId string `protobuf:"bytes,3,opt,name=code_server_id,json=codeServerId,proto3" json:"code_server_id,omitempty"`
	// path of the group/owner/user on github or gitlab
	// (note) the parent path should start with the code_server_group path
	// when creating a project on gitlab, all the groups in the provided path are created if they do not exist.
	ParentPath string `protobuf:"bytes,4,opt,name=parent_path,json=parentPath,proto3" json:"parent_path,omitempty"`
	// web or clone(http/ssh) for the code project.
	// important: this value is used for reverse lookup to generate the code pipelines and other
	// operations performed using the cli.
	// The cli reads the value using "git remote get-url origin" command.
	// so, the value returned by "git remote get-url origin" should match this attribute.
	// conditionally required while adding(add) a code-project project that already exists.
	// is not required while creating a code project
	CodeProjectUrl string `protobuf:"bytes,5,opt,name=code_project_url,json=codeProjectUrl,proto3" json:"code_project_url,omitempty"`
	// toggle to control code pipeline for the code project.
	IsCodePipelineEnabled bool `protobuf:"varint,6,opt,name=is_code_pipeline_enabled,json=isCodePipelineEnabled,proto3" json:"is_code_pipeline_enabled,omitempty"`
	// toggle to control code pipeline for the code project for merge request commits
	IsReviewCodePipelineEnabled bool `protobuf:"varint,7,opt,name=is_review_code_pipeline_enabled,json=isReviewCodePipelineEnabled,proto3" json:"is_review_code_pipeline_enabled,omitempty"`
	// toggle to control code pipeline for the code project for tags
	IsTagCodePipelineEnabled bool `protobuf:"varint,8,opt,name=is_tag_code_pipeline_enabled,json=isTagCodePipelineEnabled,proto3" json:"is_tag_code_pipeline_enabled,omitempty"`
	// id of the code project on upstream code server.
	// this is the id of repository on github/gitlab etc, if the code project was created as part of code-server sync.
	// value is computed by looking up the repository on the code-server.
	UpstreamCodeProjectId string `protobuf:"bytes,9,opt,name=upstream_code_project_id,json=upstreamCodeProjectId,proto3" json:"upstream_code_project_id,omitempty"`
	// browser url for the code project to be used for opening the code project on github/gitlab/bitbucket etc.
	// value is either populated from the information received about the project on git server or is
	// computed by applying string transformations on the clone url.
	BrowserUrl string `protobuf:"bytes,10,opt,name=browser_url,json=browserUrl,proto3" json:"browser_url,omitempty"`
	// flag to indicate if the code project is a template project.
	// a template project on github is marked as template project and is also a valid cookiecutter project.
	// https://cookiecutter.readthedocs.io
	IsTemplate bool `protobuf:"varint,11,opt,name=is_template,json=isTemplate,proto3" json:"is_template,omitempty"`
	// configuration variables required to render a full project from a template project.
	// https://cookiecutter.readthedocs.io/en/1.7.2/tutorial1.html#cookiecutter-json
	// when a template code project is synchronized from github, if the project contains cookiecutter.json file in the
	// root directory of the project in the default branch, then the json contents are loaded as key value map and are
	// persistent in the planton cloud system.
	// the stored configuration is used to allow the users to provide the values while creating new projects from the
	// template project.
	CookiecutterJson map[string]string `protobuf:"bytes,12,rep,name=cookiecutter_json,json=cookiecutterJson,proto3" json:"cookiecutter_json,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// id of the template code project from which the code project is created.
	// this value is only populated for code projects created using planton cloud.
	TemplateCodeProjectId string `protobuf:"bytes,13,opt,name=template_code_project_id,json=templateCodeProjectId,proto3" json:"template_code_project_id,omitempty"`
	// input values provided when code project is created using a template.
	// this value is only populated for code projects created using planton cloud.
	// conditionally required when creating a code-project from a template.
	CookiecutterInput map[string]string `protobuf:"bytes,14,rep,name=cookiecutter_input,json=cookiecutterInput,proto3" json:"cookiecutter_input,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CodeProjectSpec) Reset() {
	*x = CodeProjectSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeProjectSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeProjectSpec) ProtoMessage() {}

func (x *CodeProjectSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeProjectSpec.ProtoReflect.Descriptor instead.
func (*CodeProjectSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDescGZIP(), []int{1}
}

func (x *CodeProjectSpec) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CodeProjectSpec) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CodeProjectSpec) GetCodeServerId() string {
	if x != nil {
		return x.CodeServerId
	}
	return ""
}

func (x *CodeProjectSpec) GetParentPath() string {
	if x != nil {
		return x.ParentPath
	}
	return ""
}

func (x *CodeProjectSpec) GetCodeProjectUrl() string {
	if x != nil {
		return x.CodeProjectUrl
	}
	return ""
}

func (x *CodeProjectSpec) GetIsCodePipelineEnabled() bool {
	if x != nil {
		return x.IsCodePipelineEnabled
	}
	return false
}

func (x *CodeProjectSpec) GetIsReviewCodePipelineEnabled() bool {
	if x != nil {
		return x.IsReviewCodePipelineEnabled
	}
	return false
}

func (x *CodeProjectSpec) GetIsTagCodePipelineEnabled() bool {
	if x != nil {
		return x.IsTagCodePipelineEnabled
	}
	return false
}

func (x *CodeProjectSpec) GetUpstreamCodeProjectId() string {
	if x != nil {
		return x.UpstreamCodeProjectId
	}
	return ""
}

func (x *CodeProjectSpec) GetBrowserUrl() string {
	if x != nil {
		return x.BrowserUrl
	}
	return ""
}

func (x *CodeProjectSpec) GetIsTemplate() bool {
	if x != nil {
		return x.IsTemplate
	}
	return false
}

func (x *CodeProjectSpec) GetCookiecutterJson() map[string]string {
	if x != nil {
		return x.CookiecutterJson
	}
	return nil
}

func (x *CodeProjectSpec) GetTemplateCodeProjectId() string {
	if x != nil {
		return x.TemplateCodeProjectId
	}
	return ""
}

func (x *CodeProjectSpec) GetCookiecutterInput() map[string]string {
	if x != nil {
		return x.CookiecutterInput
	}
	return nil
}

// code-project status
type CodeProjectStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource lifecycle
	Lifecycle *resource.ResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// resource audit info
	Audit *audit.ResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
}

func (x *CodeProjectStatus) Reset() {
	*x = CodeProjectStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeProjectStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeProjectStatus) ProtoMessage() {}

func (x *CodeProjectStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeProjectStatus.ProtoReflect.Descriptor instead.
func (*CodeProjectStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDescGZIP(), []int{2}
}

func (x *CodeProjectStatus) GetLifecycle() *resource.ResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *CodeProjectStatus) GetAudit() *audit.ResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

var File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDesc = []byte{
	0x0a, 0x47, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x51, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x57, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x04, 0x0a, 0x0b,
	0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0xdb, 0x01, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x8c, 0x01, 0xba, 0x48, 0x88, 0x01, 0xba, 0x01, 0x7f, 0x0a, 0x0d, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x4d, 0x75, 0x73, 0x74,
	0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x6e, 0x20, 0x47,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x1a, 0x42, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a,
	0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d,
	0x39, 0x5f, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x39, 0x38, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d,
	0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x3f, 0x24, 0x27, 0x29, 0x72, 0x04, 0x10, 0x01, 0x18, 0x63,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x60, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x66, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x3a, 0x2d, 0x8a, 0xb5, 0x18, 0x02, 0x63, 0x70, 0x90, 0xb5, 0x18, 0x00,
	0x98, 0xb5, 0x18, 0x01, 0x88, 0xa6, 0x1d, 0x04, 0x90, 0xa6, 0x1d, 0x00, 0x9a, 0xa6, 0x1d, 0x13,
	0x08, 0x14, 0x12, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x22, 0x93, 0x08, 0x0a, 0x0f, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x23, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18,
	0x01, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xc0, 0xb8, 0x18, 0x01, 0xd0,
	0xb8, 0x18, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55,
	0x72, 0x6c, 0x12, 0x37, 0x0a, 0x18, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x1f, 0x69,
	0x73, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x69, 0x73, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f,
	0x64, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x3e, 0x0a, 0x1c, 0x69, 0x73, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x69, 0x73, 0x54, 0x61, 0x67, 0x43, 0x6f,
	0x64, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x3d, 0x0a, 0x18, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x15, 0x75, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0b, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x0a, 0x62, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x95, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x62, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x4a,
	0x73, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x10,
	0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x4a, 0x73, 0x6f, 0x6e,
	0x12, 0x37, 0x0a, 0x18, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x92, 0x01, 0x0a, 0x12, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x63, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x63, 0x75, 0x74, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x63, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x43,
	0x0a, 0x15, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x4a, 0x73,
	0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x44, 0x0a, 0x16, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x63, 0x75, 0x74,
	0x74, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb6, 0x01, 0x0a, 0x11, 0x43, 0x6f,
	0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x57, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x63, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x18, 0x62, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x05, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x42, 0xdc, 0x03, 0x0a, 0x49, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0xa2, 0x02, 0x08, 0x43, 0x50,
	0x41, 0x56, 0x43, 0x44, 0x53, 0x50, 0xaa, 0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0xca, 0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x5c,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0xe2, 0x02, 0x47, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x5c, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x42, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41,
	0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x3a, 0x3a, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_goTypes = []interface{}{
	(*CodeProject)(nil),                // 0: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProject
	(*CodeProjectSpec)(nil),            // 1: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProjectSpec
	(*CodeProjectStatus)(nil),          // 2: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProjectStatus
	nil,                                // 3: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProjectSpec.CookiecutterJsonEntry
	nil,                                // 4: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProjectSpec.CookiecutterInputEntry
	(*resource.Metadata)(nil),          // 5: cloud.planton.apis.v1.commons.resource.Metadata
	(*resource.ResourceLifecycle)(nil), // 6: cloud.planton.apis.v1.commons.resource.ResourceLifecycle
	(*audit.ResourceAudit)(nil),        // 7: cloud.planton.apis.v1.commons.audit.ResourceAudit
}
var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProject.metadata:type_name -> cloud.planton.apis.v1.commons.resource.Metadata
	1, // 1: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProject.spec:type_name -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProjectSpec
	2, // 2: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProject.status:type_name -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProjectStatus
	3, // 3: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProjectSpec.cookiecutter_json:type_name -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProjectSpec.CookiecutterJsonEntry
	4, // 4: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProjectSpec.cookiecutter_input:type_name -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProjectSpec.CookiecutterInputEntry
	6, // 5: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProjectStatus.lifecycle:type_name -> cloud.planton.apis.v1.commons.resource.ResourceLifecycle
	7, // 6: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.CodeProjectStatus.audit:type_name -> cloud.planton.apis.v1.commons.audit.ResourceAudit
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeProject); i {
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
		file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeProjectSpec); i {
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
		file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeProjectStatus); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_state_proto_depIdxs = nil
}
