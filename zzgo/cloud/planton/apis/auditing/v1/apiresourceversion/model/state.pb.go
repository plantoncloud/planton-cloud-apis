// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/auditing/v1/apiresourceversion/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	apiresourcekind "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/enums/apiresourcekind"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcefieldoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemessageoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemetadataoptions"
	stackjobexecutionstatustype "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/enums/stackjobexecutionstatustype"
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

// The ApiResourceVersion message represents a record of a resource state modification
// in a system's resource. It stores the original and new states of the resource,
// the differences between them in a unified diff format, as well as audit and
// identification information.
type ApiResourceVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind     string                     `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata *model.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *ApiResourceVersionSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *ApiResourceVersionSpec `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ApiResourceVersion) Reset() {
	*x = ApiResourceVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResourceVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResourceVersion) ProtoMessage() {}

func (x *ApiResourceVersion) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResourceVersion.ProtoReflect.Descriptor instead.
func (*ApiResourceVersion) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *ApiResourceVersion) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *ApiResourceVersion) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ApiResourceVersion) GetMetadata() *model.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ApiResourceVersion) GetSpec() *ApiResourceVersionSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *ApiResourceVersion) GetStatus() *ApiResourceVersionSpec {
	if x != nil {
		return x.Status
	}
	return nil
}

type ApiResourceVersionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the company to which the api-resource-version belongs.
	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// The ID of the product to which the api-resource-version belongs.
	// This attribute is always empty for company level state.
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// The ID of the product environment to which the api-resource-version belongs.
	// This attribute is always empty for company level state.
	EnvironmentId string `protobuf:"bytes,3,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	// The type of the resource for which the version was created.
	// Resource types are defined in the ApiResourceKind enum.
	ResourceKind apiresourcekind.ApiResourceKind `protobuf:"varint,4,opt,name=resource_kind,json=resourceKind,proto3,enum=cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind" json:"resource_kind,omitempty"`
	// The ID of the resource for which the version was created.
	ResourceId string `protobuf:"bytes,5,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The original state of the resource, represented as a YAML string.
	OriginalStateYaml string `protobuf:"bytes,6,opt,name=original_state_yaml,json=originalStateYaml,proto3" json:"original_state_yaml,omitempty"`
	// The new state of the resource, represented as a YAML string.
	NewStateYaml string `protobuf:"bytes,7,opt,name=new_state_yaml,json=newStateYaml,proto3" json:"new_state_yaml,omitempty"`
	// The differences between the original and new state, represented in a unified diff format.
	DiffUnifiedFormat string `protobuf:"bytes,8,opt,name=diff_unified_format,json=diffUnifiedFormat,proto3" json:"diff_unified_format,omitempty"`
	// operation done on the state which ended up in creating this new version.
	OperationType string `protobuf:"bytes,9,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
	// unique identifier for the stack job.
	StackJobId string `protobuf:"bytes,10,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
}

func (x *ApiResourceVersionSpec) Reset() {
	*x = ApiResourceVersionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResourceVersionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResourceVersionSpec) ProtoMessage() {}

func (x *ApiResourceVersionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResourceVersionSpec.ProtoReflect.Descriptor instead.
func (*ApiResourceVersionSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *ApiResourceVersionSpec) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *ApiResourceVersionSpec) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ApiResourceVersionSpec) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *ApiResourceVersionSpec) GetResourceKind() apiresourcekind.ApiResourceKind {
	if x != nil {
		return x.ResourceKind
	}
	return apiresourcekind.ApiResourceKind(0)
}

func (x *ApiResourceVersionSpec) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *ApiResourceVersionSpec) GetOriginalStateYaml() string {
	if x != nil {
		return x.OriginalStateYaml
	}
	return ""
}

func (x *ApiResourceVersionSpec) GetNewStateYaml() string {
	if x != nil {
		return x.NewStateYaml
	}
	return ""
}

func (x *ApiResourceVersionSpec) GetDiffUnifiedFormat() string {
	if x != nil {
		return x.DiffUnifiedFormat
	}
	return ""
}

func (x *ApiResourceVersionSpec) GetOperationType() string {
	if x != nil {
		return x.OperationType
	}
	return ""
}

func (x *ApiResourceVersionSpec) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

// api-resource-version status
type ApiResourceVersionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource lifecycle
	Lifecycle *model.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// resource audit info
	Audit *model.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// represents the current status of the stack job.
	Status stackjobexecutionstatustype.StackJobExecutionStatusType `protobuf:"varint,1,opt,name=status,proto3,enum=cloud.planton.apis.iac.v1.stackjob.enums.stackjobexecutionstatustype.StackJobExecutionStatusType" json:"status,omitempty"`
}

func (x *ApiResourceVersionStatus) Reset() {
	*x = ApiResourceVersionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResourceVersionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResourceVersionStatus) ProtoMessage() {}

func (x *ApiResourceVersionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResourceVersionStatus.ProtoReflect.Descriptor instead.
func (*ApiResourceVersionStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *ApiResourceVersionStatus) GetLifecycle() *model.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *ApiResourceVersionStatus) GetAudit() *model.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *ApiResourceVersionStatus) GetStatus() stackjobexecutionstatustype.StackJobExecutionStatusType {
	if x != nil {
		return x.Status
	}
	return stackjobexecutionstatustype.StackJobExecutionStatusType(0)
}

var File_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDesc = []byte{
	0x0a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x54, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
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
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x03, 0x0a, 0x12, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x20, 0xba, 0x48, 0x1d, 0x72, 0x1b, 0x0a, 0x19, 0x61, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2d,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xba, 0x48,
	0x16, 0x72, 0x14, 0x0a, 0x12, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x5d, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x63, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x67, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x4f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x0f, 0x8a, 0xb5, 0x18, 0x03,
	0x76, 0x65, 0x72, 0x88, 0xa6, 0x1d, 0x01, 0x90, 0xa6, 0x1d, 0x00, 0x22, 0xe1, 0x03, 0x0a, 0x16,
	0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x72, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e,
	0x64, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e,
	0x64, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x79, 0x61, 0x6d, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x59, 0x61, 0x6d, 0x6c,
	0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x79, 0x61,
	0x6d, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x59, 0x61, 0x6d, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x66, 0x66, 0x5f, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x69, 0x66, 0x66, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x22,
	0xcd, 0x02, 0x0a, 0x18, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x60, 0x0a, 0x09,
	0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x42, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x54,
	0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x62, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x05, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x12, 0x79, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x4a, 0x6f, 0x62, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0xc2, 0x03, 0x0a, 0x45, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65,
	0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x41, 0x56, 0x41, 0x4d, 0xaa, 0x02, 0x37, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x69,
	0x6e, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02,
	0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x5c,
	0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDescData = file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDesc
)

func file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDescData
}

var file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_goTypes = []interface{}{
	(*ApiResourceVersion)(nil),                                   // 0: cloud.planton.apis.auditing.v1.apiresourceversion.model.ApiResourceVersion
	(*ApiResourceVersionSpec)(nil),                               // 1: cloud.planton.apis.auditing.v1.apiresourceversion.model.ApiResourceVersionSpec
	(*ApiResourceVersionStatus)(nil),                             // 2: cloud.planton.apis.auditing.v1.apiresourceversion.model.ApiResourceVersionStatus
	(*model.ApiResourceMetadata)(nil),                            // 3: cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	(apiresourcekind.ApiResourceKind)(0),                         // 4: cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind
	(*model.ApiResourceLifecycle)(nil),                           // 5: cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	(*model.ApiResourceAudit)(nil),                               // 6: cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	(stackjobexecutionstatustype.StackJobExecutionStatusType)(0), // 7: cloud.planton.apis.iac.v1.stackjob.enums.stackjobexecutionstatustype.StackJobExecutionStatusType
}
var file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_depIdxs = []int32{
	3, // 0: cloud.planton.apis.auditing.v1.apiresourceversion.model.ApiResourceVersion.metadata:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	1, // 1: cloud.planton.apis.auditing.v1.apiresourceversion.model.ApiResourceVersion.spec:type_name -> cloud.planton.apis.auditing.v1.apiresourceversion.model.ApiResourceVersionSpec
	1, // 2: cloud.planton.apis.auditing.v1.apiresourceversion.model.ApiResourceVersion.status:type_name -> cloud.planton.apis.auditing.v1.apiresourceversion.model.ApiResourceVersionSpec
	4, // 3: cloud.planton.apis.auditing.v1.apiresourceversion.model.ApiResourceVersionSpec.resource_kind:type_name -> cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind
	5, // 4: cloud.planton.apis.auditing.v1.apiresourceversion.model.ApiResourceVersionStatus.lifecycle:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	6, // 5: cloud.planton.apis.auditing.v1.apiresourceversion.model.ApiResourceVersionStatus.audit:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	7, // 6: cloud.planton.apis.auditing.v1.apiresourceversion.model.ApiResourceVersionStatus.status:type_name -> cloud.planton.apis.iac.v1.stackjob.enums.stackjobexecutionstatustype.StackJobExecutionStatusType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_init() }
func file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_init() {
	if File_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResourceVersion); i {
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
		file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResourceVersionSpec); i {
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
		file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResourceVersionStatus); i {
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
			RawDescriptor: file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto = out.File
	file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_goTypes = nil
	file_cloud_planton_apis_auditing_v1_apiresourceversion_model_state_proto_depIdxs = nil
}
