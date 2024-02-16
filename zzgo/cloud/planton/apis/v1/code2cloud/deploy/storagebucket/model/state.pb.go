// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/storagebucket/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/cloudaccount/provider/enums/storagebucketprovider"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/environment/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/audit/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/field/options"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/metadata/options"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/model"
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

// storage-bucket
type StorageBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// resource metadata
	Metadata *model.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *StorageBucketSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *StorageBucketStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StorageBucket) Reset() {
	*x = StorageBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucket) ProtoMessage() {}

func (x *StorageBucket) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucket.ProtoReflect.Descriptor instead.
func (*StorageBucket) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *StorageBucket) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *StorageBucket) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *StorageBucket) GetMetadata() *model.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *StorageBucket) GetSpec() *StorageBucketSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *StorageBucket) GetStatus() *StorageBucketStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// storage-bucket spec
type StorageBucketSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource parent
	EnvironmentInfo *model1.ResourceEnvironmentInfo `protobuf:"bytes,1,opt,name=environment_info,json=environmentInfo,proto3" json:"environment_info,omitempty"`
	// (optional for create) flag to indicate if storage-bucket should have external(public) access.
	// defaults to "false"
	IsExternal bool `protobuf:"varint,2,opt,name=is_external,json=isExternal,proto3" json:"is_external,omitempty"`
	// storage-bucket spec on gcp
	Gcp *StorageBucketGcpSpec `protobuf:"bytes,4,opt,name=gcp,proto3" json:"gcp,omitempty"`
}

func (x *StorageBucketSpec) Reset() {
	*x = StorageBucketSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageBucketSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucketSpec) ProtoMessage() {}

func (x *StorageBucketSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucketSpec.ProtoReflect.Descriptor instead.
func (*StorageBucketSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *StorageBucketSpec) GetEnvironmentInfo() *model1.ResourceEnvironmentInfo {
	if x != nil {
		return x.EnvironmentInfo
	}
	return nil
}

func (x *StorageBucketSpec) GetIsExternal() bool {
	if x != nil {
		return x.IsExternal
	}
	return false
}

func (x *StorageBucketSpec) GetGcp() *StorageBucketGcpSpec {
	if x != nil {
		return x.Gcp
	}
	return nil
}

// storage-bucket status
type StorageBucketStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource lifecycle
	Lifecycle *model.ResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// resource audit info
	Audit *model2.ResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// id of the stack-job
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
}

func (x *StorageBucketStatus) Reset() {
	*x = StorageBucketStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageBucketStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucketStatus) ProtoMessage() {}

func (x *StorageBucketStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucketStatus.ProtoReflect.Descriptor instead.
func (*StorageBucketStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *StorageBucketStatus) GetLifecycle() *model.ResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *StorageBucketStatus) GetAudit() *model2.ResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *StorageBucketStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

// storage-bucket on gcp
type StorageBucketGcpSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the gcp cloud account in which the storage-bucket is created.
	// this field is computed from the kube-cluster configured for the environment.
	GcpCloudAccountId string `protobuf:"bytes,1,opt,name=gcp_cloud_account_id,json=gcpCloudAccountId,proto3" json:"gcp_cloud_account_id,omitempty"`
	// gcp project in which the storage-bucket is created.
	// this field is computed from the kube-cluster configured for the environment.
	// storage-bucket is created in the same gcp-project in which the container-cluster is created.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// gcp region in which the bucket is to be created.
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *StorageBucketGcpSpec) Reset() {
	*x = StorageBucketGcpSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageBucketGcpSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucketGcpSpec) ProtoMessage() {}

func (x *StorageBucketGcpSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucketGcpSpec.ProtoReflect.Descriptor instead.
func (*StorageBucketGcpSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDescGZIP(), []int{3}
}

func (x *StorageBucketGcpSpec) GetGcpCloudAccountId() string {
	if x != nil {
		return x.GcpCloudAccountId
	}
	return ""
}

func (x *StorageBucketGcpSpec) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *StorageBucketGcpSpec) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

var File_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDesc = []byte{
	0x0a, 0x47, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x51, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x57, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x04, 0x0a, 0x0d, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0xe5, 0x01, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x90, 0x01, 0xba, 0x48, 0x8c, 0x01, 0xba, 0x01,
	0x82, 0x01, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x38, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20,
	0x61, 0x20, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x4e,
	0x61, 0x6d, 0x65, 0x20, 0x6f, 0x6e, 0x20, 0x41, 0x57, 0x53, 0x2f, 0x47, 0x43, 0x50, 0x2f, 0x41,
	0x5a, 0x55, 0x52, 0x45, 0x2f, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x1a, 0x37, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27,
	0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x3f, 0x3a, 0x5b, 0x61, 0x2d, 0x7a,
	0x30, 0x2d, 0x39, 0x5f, 0x2d, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29,
	0x3f, 0x24, 0x27, 0x29, 0x72, 0x04, 0x10, 0x03, 0x18, 0x3f, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x62, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x68, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x3a, 0x43, 0x8a, 0xb5, 0x18, 0x03, 0x62, 0x6b, 0x74, 0x90, 0xb5, 0x18, 0x00, 0x98,
	0xb5, 0x18, 0x01, 0x88, 0xa6, 0x1d, 0x18, 0x90, 0xa6, 0x1d, 0x00, 0x9a, 0xa6, 0x1d, 0x28, 0x08,
	0x09, 0x12, 0x24, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x91, 0x02, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x76, 0x0a,
	0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x63, 0x0a, 0x03, 0x67, 0x63, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x47,
	0x63, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52, 0x03, 0x67, 0x63, 0x70, 0x22, 0xe6, 0x01, 0x0a, 0x13,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x5d, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x62, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x05, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f,
	0x69, 0x64, 0x18, 0x61, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a,
	0x6f, 0x62, 0x49, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x47, 0x63, 0x70, 0x53, 0x70, 0x65, 0x63, 0x12, 0x35, 0x0a,
	0x14, 0x67, 0x63, 0x70, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18,
	0x01, 0x52, 0x11, 0x67, 0x63, 0x70, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0xdc, 0x03, 0x0a, 0x49, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x6b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x53, 0x4d, 0xaa, 0x02, 0x3b, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73,
	0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56,
	0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x47, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x42, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x3a,
	0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_goTypes = []interface{}{
	(*StorageBucket)(nil),                  // 0: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucket
	(*StorageBucketSpec)(nil),              // 1: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucketSpec
	(*StorageBucketStatus)(nil),            // 2: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucketStatus
	(*StorageBucketGcpSpec)(nil),           // 3: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucketGcpSpec
	(*model.Metadata)(nil),                 // 4: cloud.planton.apis.v1.commons.resource.model.Metadata
	(*model1.ResourceEnvironmentInfo)(nil), // 5: cloud.planton.apis.v1.code2cloud.environment.model.ResourceEnvironmentInfo
	(*model.ResourceLifecycle)(nil),        // 6: cloud.planton.apis.v1.commons.resource.model.ResourceLifecycle
	(*model2.ResourceAudit)(nil),           // 7: cloud.planton.apis.v1.commons.audit.model.ResourceAudit
}
var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucket.metadata:type_name -> cloud.planton.apis.v1.commons.resource.model.Metadata
	1, // 1: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucket.spec:type_name -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucketSpec
	2, // 2: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucket.status:type_name -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucketStatus
	5, // 3: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucketSpec.environment_info:type_name -> cloud.planton.apis.v1.code2cloud.environment.model.ResourceEnvironmentInfo
	3, // 4: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucketSpec.gcp:type_name -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucketGcpSpec
	6, // 5: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucketStatus.lifecycle:type_name -> cloud.planton.apis.v1.commons.resource.model.ResourceLifecycle
	7, // 6: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.model.StorageBucketStatus.audit:type_name -> cloud.planton.apis.v1.commons.audit.model.ResourceAudit
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageBucket); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageBucketSpec); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageBucketStatus); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageBucketGcpSpec); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_state_proto_depIdxs = nil
}
