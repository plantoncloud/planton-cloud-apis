// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/gcp/gcsbucket/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcefieldoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemessageoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/resourcemanager/v1/environment/model"
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

// gcs-bucket
type GcsBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource-kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata
	// id format "<id-prefix>-<env-id>-<normalized-resource-name>"
	Metadata *model.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *GcsBucketSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *GcsBucketStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GcsBucket) Reset() {
	*x = GcsBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsBucket) ProtoMessage() {}

func (x *GcsBucket) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsBucket.ProtoReflect.Descriptor instead.
func (*GcsBucket) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *GcsBucket) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *GcsBucket) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *GcsBucket) GetMetadata() *model.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GcsBucket) GetSpec() *GcsBucketSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *GcsBucket) GetStatus() *GcsBucketStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// gcs-bucket spec
type GcsBucketSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// environment-info
	EnvironmentInfo *model1.ApiResourceEnvironmentInfo `protobuf:"bytes,99,opt,name=environment_info,json=environmentInfo,proto3" json:"environment_info,omitempty"`
	// stack-job settings
	StackJobSettings *model2.StackJobSettings `protobuf:"bytes,98,opt,name=stack_job_settings,json=stackJobSettings,proto3" json:"stack_job_settings,omitempty"`
	// gcp-credential-id to be used for setting up gcp-provider in stack-job
	GcpCredentialId string `protobuf:"bytes,97,opt,name=gcp_credential_id,json=gcpCredentialId,proto3" json:"gcp_credential_id,omitempty"`
	// gcp project in which the storage-bucket is to be created.
	GcpProjectId string `protobuf:"bytes,1,opt,name=gcp_project_id,json=gcpProjectId,proto3" json:"gcp_project_id,omitempty"`
	// gcp region
	GcpRegion string `protobuf:"bytes,2,opt,name=gcp_region,json=gcpRegion,proto3" json:"gcp_region,omitempty"`
	// flag to indicate if gcs-bucket should have external(public) access.
	// defaults to "false"
	IsPublic bool `protobuf:"varint,3,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
}

func (x *GcsBucketSpec) Reset() {
	*x = GcsBucketSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsBucketSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsBucketSpec) ProtoMessage() {}

func (x *GcsBucketSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsBucketSpec.ProtoReflect.Descriptor instead.
func (*GcsBucketSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *GcsBucketSpec) GetEnvironmentInfo() *model1.ApiResourceEnvironmentInfo {
	if x != nil {
		return x.EnvironmentInfo
	}
	return nil
}

func (x *GcsBucketSpec) GetStackJobSettings() *model2.StackJobSettings {
	if x != nil {
		return x.StackJobSettings
	}
	return nil
}

func (x *GcsBucketSpec) GetGcpCredentialId() string {
	if x != nil {
		return x.GcpCredentialId
	}
	return ""
}

func (x *GcsBucketSpec) GetGcpProjectId() string {
	if x != nil {
		return x.GcpProjectId
	}
	return ""
}

func (x *GcsBucketSpec) GetGcpRegion() string {
	if x != nil {
		return x.GcpRegion
	}
	return ""
}

func (x *GcsBucketSpec) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

// gcs-bucket status
type GcsBucketStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// lifecycle
	Lifecycle *model.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// audit-info
	Audit *model.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// stack-job id
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
	// stack-outputs
	StackOutputs *GcsBucketStackOutputs `protobuf:"bytes,1,opt,name=stack_outputs,json=stackOutputs,proto3" json:"stack_outputs,omitempty"`
}

func (x *GcsBucketStatus) Reset() {
	*x = GcsBucketStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsBucketStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsBucketStatus) ProtoMessage() {}

func (x *GcsBucketStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsBucketStatus.ProtoReflect.Descriptor instead.
func (*GcsBucketStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *GcsBucketStatus) GetLifecycle() *model.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *GcsBucketStatus) GetAudit() *model.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *GcsBucketStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

func (x *GcsBucketStatus) GetStackOutputs() *GcsBucketStackOutputs {
	if x != nil {
		return x.StackOutputs
	}
	return nil
}

// gcs-bucket stack outputs
type GcsBucketStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the storage-bucket created on aws
	BucketId string `protobuf:"bytes,1,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
}

func (x *GcsBucketStackOutputs) Reset() {
	*x = GcsBucketStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsBucketStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsBucketStackOutputs) ProtoMessage() {}

func (x *GcsBucketStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsBucketStackOutputs.ProtoReflect.Descriptor instead.
func (*GcsBucketStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDescGZIP(), []int{3}
}

func (x *GcsBucketStackOutputs) GetBucketId() string {
	if x != nil {
		return x.BucketId
	}
	return ""
}

var File_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDesc = []byte{
	0x0a, 0x40, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x73, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x7c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f,
	0x06, 0x0a, 0x09, 0x47, 0x63, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x43, 0x0a, 0x0b,
	0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x22, 0xba, 0x48, 0x1f, 0x72, 0x1d, 0x0a, 0x1b, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x31, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0xc6, 0x03, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xe6, 0x02, 0xba, 0x48,
	0xe2, 0x02, 0xba, 0x01, 0x82, 0x01, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74,
	0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x6e, 0x20, 0x41, 0x57, 0x53, 0x2f, 0x47,
	0x43, 0x50, 0x2f, 0x41, 0x5a, 0x55, 0x52, 0x45, 0x2f, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x1a,
	0x37, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x3f, 0x3a,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x2d, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30,
	0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x24, 0x27, 0x29, 0xba, 0x01, 0x6c, 0x0a, 0x0d, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x4e, 0x61, 0x6d, 0x65,
	0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e,
	0x20, 0x33, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x36, 0x33, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x73, 0x20, 0x6c, 0x6f, 0x6e, 0x67, 0x1a, 0x2c, 0x73, 0x69, 0x7a, 0x65, 0x28,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20, 0x3e, 0x20, 0x32, 0x20, 0x26,
	0x26, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x29, 0x20, 0x3c, 0x3d, 0x20, 0x36, 0x33, 0xba, 0x01, 0x67, 0x0a, 0x18, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x69, 0x73, 0x20, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x79, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65,
	0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x29, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x57,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x47, 0x63, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x5d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67,
	0x63, 0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47,
	0x63, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x28, 0x88, 0xa6, 0x1d, 0x0f, 0x92, 0xa6, 0x1d, 0x20,
	0x08, 0x0a, 0x12, 0x1c, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x69, 0x64,
	0x22, 0x97, 0x03, 0x0a, 0x0d, 0x47, 0x63, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x7e, 0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x68, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x62, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a,
	0x6f, 0x62, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x10, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2a, 0x0a, 0x11,
	0x67, 0x63, 0x70, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x61, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x63, 0x70, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0e, 0x67, 0x63, 0x70, 0x5f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0c, 0x67, 0x63, 0x70, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0a, 0x67, 0x63, 0x70, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x09, 0x67, 0x63, 0x70, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0xdd, 0x02, 0x0a, 0x0f, 0x47,
	0x63, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x60,
	0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x63, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x12, 0x54, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x62, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52,
	0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f,
	0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x61, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x70, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x63, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x0c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22, 0x34, 0x0a, 0x15, 0x47, 0x63,
	0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64,
	0x42, 0xb2, 0x03, 0x0a, 0x42, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67,
	0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x73, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x64, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x73, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x08, 0x43, 0x50,
	0x41, 0x43, 0x56, 0x47, 0x47, 0x4d, 0xaa, 0x02, 0x34, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x47, 0x63, 0x70, 0x2e, 0x47, 0x63,
	0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x34,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x40, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x73, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a,
	0x47, 0x63, 0x70, 0x3a, 0x3a, 0x47, 0x63, 0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x3a,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_goTypes = []interface{}{
	(*GcsBucket)(nil),                         // 0: cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucket
	(*GcsBucketSpec)(nil),                     // 1: cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucketSpec
	(*GcsBucketStatus)(nil),                   // 2: cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucketStatus
	(*GcsBucketStackOutputs)(nil),             // 3: cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucketStackOutputs
	(*model.ApiResourceMetadata)(nil),         // 4: cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	(*model1.ApiResourceEnvironmentInfo)(nil), // 5: cloud.planton.apis.resourcemanager.v1.environment.model.ApiResourceEnvironmentInfo
	(*model2.StackJobSettings)(nil),           // 6: cloud.planton.apis.iac.v1.stackjob.model.StackJobSettings
	(*model.ApiResourceLifecycle)(nil),        // 7: cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	(*model.ApiResourceAudit)(nil),            // 8: cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
}
var file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucket.metadata:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	1, // 1: cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucket.spec:type_name -> cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucketSpec
	2, // 2: cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucket.status:type_name -> cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucketStatus
	5, // 3: cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucketSpec.environment_info:type_name -> cloud.planton.apis.resourcemanager.v1.environment.model.ApiResourceEnvironmentInfo
	6, // 4: cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucketSpec.stack_job_settings:type_name -> cloud.planton.apis.iac.v1.stackjob.model.StackJobSettings
	7, // 5: cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucketStatus.lifecycle:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	8, // 6: cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucketStatus.audit:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	3, // 7: cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucketStatus.stack_outputs:type_name -> cloud.planton.apis.code2cloud.v1.gcp.gcsbucket.model.GcsBucketStackOutputs
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsBucket); i {
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
		file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsBucketSpec); i {
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
		file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsBucketStatus); i {
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
		file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsBucketStackOutputs); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_gcp_gcsbucket_model_state_proto_depIdxs = nil
}
