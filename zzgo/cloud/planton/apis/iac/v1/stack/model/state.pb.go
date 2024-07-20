// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/iac/v1/stack/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	apiresourcekind "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/enums/apiresourcekind"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcefieldoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemessageoptions"
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

// stack
type Stack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource-kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata
	// id format "<id-prefix>-<normalized-resource-name>"
	Metadata *model.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *StackSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *model.ApiResourceLifecycleAndAuditStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Stack) Reset() {
	*x = Stack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_iac_v1_stack_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stack) ProtoMessage() {}

func (x *Stack) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_iac_v1_stack_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stack.ProtoReflect.Descriptor instead.
func (*Stack) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *Stack) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Stack) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Stack) GetMetadata() *model.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Stack) GetSpec() *StackSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Stack) GetStatus() *model.ApiResourceLifecycleAndAuditStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// stack spec
type StackSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of the resource for which the stack belongs to.
	ResourceKind apiresourcekind.ApiResourceKind `protobuf:"varint,1,opt,name=resource_kind,json=resourceKind,proto3,enum=cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind" json:"resource_kind,omitempty"`
	// id of the resource for which the stack belongs to.
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// url of the stack-job on pulumi web console.
	// note: this value is not persisted in the database.
	// the value of this attributes is rendered by replacing the placeholders in
	// "https://app.pulumi.com/${pulumiOrgName}/${pulumiProject}/${pulumiStackName}"
	// value of pulumiOrgName is same for every single stack for each planton-cloud environment.
	// value of pulumiProject is the organization-id on planton-cloud.
	// value of pulumiStack is stack_name attribute in this object.
	// ex: https://app.pulumi.com/plantonstack-prod/cookie/afs-cookie-shop-main.ca-planton-hosting-gcp-main.artifact-store
	PulumiCloudUrl string `protobuf:"bytes,3,opt,name=pulumi_cloud_url,json=pulumiCloudUrl,proto3" json:"pulumi_cloud_url,omitempty"`
	// id of the organization to which the resource belongs to.
	// this is used as project name when creating stacks on pulumi cloud.
	OrgId string `protobuf:"bytes,4,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// id of the identity-account that created the resource.
	// this is used to configure the audit-actor before persisting the state of stack resource.
	ApiResourceCreatedBy string `protobuf:"bytes,5,opt,name=api_resource_created_by,json=apiResourceCreatedBy,proto3" json:"api_resource_created_by,omitempty"`
}

func (x *StackSpec) Reset() {
	*x = StackSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_iac_v1_stack_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackSpec) ProtoMessage() {}

func (x *StackSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_iac_v1_stack_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackSpec.ProtoReflect.Descriptor instead.
func (*StackSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *StackSpec) GetResourceKind() apiresourcekind.ApiResourceKind {
	if x != nil {
		return x.ResourceKind
	}
	return apiresourcekind.ApiResourceKind(0)
}

func (x *StackSpec) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *StackSpec) GetPulumiCloudUrl() string {
	if x != nil {
		return x.PulumiCloudUrl
	}
	return ""
}

func (x *StackSpec) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *StackSpec) GetApiResourceCreatedBy() string {
	if x != nil {
		return x.ApiResourceCreatedBy
	}
	return ""
}

var File_cloud_planton_apis_iac_v1_stack_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x54, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6b, 0x69, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x6b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70,
	0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x5d, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x44, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69,
	0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x68, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x50, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x41, 0x6e, 0x64, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x04, 0x88, 0xa6, 0x1d,
	0x28, 0x22, 0xad, 0x02, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x72, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b,
	0x69, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x10, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x5f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xba, 0x48, 0x03, 0xd0, 0x01, 0x01, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x0e, 0x70, 0x75, 0x6c, 0x75,
	0x6d, 0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49,
	0x64, 0x12, 0x3e, 0x0a, 0x17, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x00, 0xd0, 0xb8, 0x18, 0x01, 0x52, 0x14, 0x61, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x42, 0xd6, 0x02, 0x0a, 0x33, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02,
	0x07, 0x43, 0x50, 0x41, 0x49, 0x56, 0x53, 0x4d, 0xaa, 0x02, 0x25, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x49, 0x61,
	0x63, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0xca, 0x02, 0x25, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x49, 0x61, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x31, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x49, 0x61,
	0x63, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2b, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41,
	0x70, 0x69, 0x73, 0x3a, 0x3a, 0x49, 0x61, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDescData = file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDesc
)

func file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDescData
}

var file_cloud_planton_apis_iac_v1_stack_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_planton_apis_iac_v1_stack_model_state_proto_goTypes = []interface{}{
	(*Stack)(nil),                                    // 0: cloud.planton.apis.iac.v1.stack.model.Stack
	(*StackSpec)(nil),                                // 1: cloud.planton.apis.iac.v1.stack.model.StackSpec
	(*model.ApiResourceMetadata)(nil),                // 2: cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	(*model.ApiResourceLifecycleAndAuditStatus)(nil), // 3: cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycleAndAuditStatus
	(apiresourcekind.ApiResourceKind)(0),             // 4: cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind
}
var file_cloud_planton_apis_iac_v1_stack_model_state_proto_depIdxs = []int32{
	2, // 0: cloud.planton.apis.iac.v1.stack.model.Stack.metadata:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	1, // 1: cloud.planton.apis.iac.v1.stack.model.Stack.spec:type_name -> cloud.planton.apis.iac.v1.stack.model.StackSpec
	3, // 2: cloud.planton.apis.iac.v1.stack.model.Stack.status:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycleAndAuditStatus
	4, // 3: cloud.planton.apis.iac.v1.stack.model.StackSpec.resource_kind:type_name -> cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_iac_v1_stack_model_state_proto_init() }
func file_cloud_planton_apis_iac_v1_stack_model_state_proto_init() {
	if File_cloud_planton_apis_iac_v1_stack_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_iac_v1_stack_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stack); i {
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
		file_cloud_planton_apis_iac_v1_stack_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackSpec); i {
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
			RawDescriptor: file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_iac_v1_stack_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_iac_v1_stack_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_iac_v1_stack_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_iac_v1_stack_model_state_proto = out.File
	file_cloud_planton_apis_iac_v1_stack_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_iac_v1_stack_model_state_proto_goTypes = nil
	file_cloud_planton_apis_iac_v1_stack_model_state_proto_depIdxs = nil
}
