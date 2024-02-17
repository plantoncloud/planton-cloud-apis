// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/artifactstore/stack/gcp/model/io.proto

package model

import (
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/artifactstore/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/pulumi/operation/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job/model"
	model3 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job/progress/model"
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

// gcp artifact repo stack input
type ArtifactStoreGcpStackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack job
	StackJob *model.StackJob `protobuf:"bytes,1,opt,name=stack_job,json=stackJob,proto3" json:"stack_job,omitempty"`
	// pulumi stack credentials
	CredentialsInput *ArtifactStoreGcpStackCredentialsInput `protobuf:"bytes,2,opt,name=credentials_input,json=credentialsInput,proto3" json:"credentials_input,omitempty"`
	// inputs used for creating stack resources
	ResourceInput *ArtifactStoreGcpStackResourceInput `protobuf:"bytes,3,opt,name=resource_input,json=resourceInput,proto3" json:"resource_input,omitempty"`
}

func (x *ArtifactStoreGcpStackInput) Reset() {
	*x = ArtifactStoreGcpStackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactStoreGcpStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactStoreGcpStackInput) ProtoMessage() {}

func (x *ArtifactStoreGcpStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactStoreGcpStackInput.ProtoReflect.Descriptor instead.
func (*ArtifactStoreGcpStackInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *ArtifactStoreGcpStackInput) GetStackJob() *model.StackJob {
	if x != nil {
		return x.StackJob
	}
	return nil
}

func (x *ArtifactStoreGcpStackInput) GetCredentialsInput() *ArtifactStoreGcpStackCredentialsInput {
	if x != nil {
		return x.CredentialsInput
	}
	return nil
}

func (x *ArtifactStoreGcpStackInput) GetResourceInput() *ArtifactStoreGcpStackResourceInput {
	if x != nil {
		return x.ResourceInput
	}
	return nil
}

// stack credentials input
type ArtifactStoreGcpStackCredentialsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// google provider credential input
	// this is only populated when the artifact-store repo provider is gcp.
	Google *model1.GoogleProviderCredential `protobuf:"bytes,1,opt,name=google,proto3" json:"google,omitempty"`
}

func (x *ArtifactStoreGcpStackCredentialsInput) Reset() {
	*x = ArtifactStoreGcpStackCredentialsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactStoreGcpStackCredentialsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactStoreGcpStackCredentialsInput) ProtoMessage() {}

func (x *ArtifactStoreGcpStackCredentialsInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactStoreGcpStackCredentialsInput.ProtoReflect.Descriptor instead.
func (*ArtifactStoreGcpStackCredentialsInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *ArtifactStoreGcpStackCredentialsInput) GetGoogle() *model1.GoogleProviderCredential {
	if x != nil {
		return x.Google
	}
	return nil
}

// stack resource input
type ArtifactStoreGcpStackResourceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArtifactStore *model2.ArtifactStore `protobuf:"bytes,1,opt,name=artifact_store,json=artifactStore,proto3" json:"artifact_store,omitempty"`
}

func (x *ArtifactStoreGcpStackResourceInput) Reset() {
	*x = ArtifactStoreGcpStackResourceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactStoreGcpStackResourceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactStoreGcpStackResourceInput) ProtoMessage() {}

func (x *ArtifactStoreGcpStackResourceInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactStoreGcpStackResourceInput.ProtoReflect.Descriptor instead.
func (*ArtifactStoreGcpStackResourceInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDescGZIP(), []int{2}
}

func (x *ArtifactStoreGcpStackResourceInput) GetArtifactStore() *model2.ArtifactStore {
	if x != nil {
		return x.ArtifactStore
	}
	return nil
}

// gcp artifact repo stack outputs
type ArtifactStoreGcpStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GcpArtifactRegistryStatus *model2.ArtifactStoreGcpArtifactRegistryStatus `protobuf:"bytes,1,opt,name=gcp_artifact_registry_status,json=gcpArtifactRegistryStatus,proto3" json:"gcp_artifact_registry_status,omitempty"`
}

func (x *ArtifactStoreGcpStackOutputs) Reset() {
	*x = ArtifactStoreGcpStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactStoreGcpStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactStoreGcpStackOutputs) ProtoMessage() {}

func (x *ArtifactStoreGcpStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactStoreGcpStackOutputs.ProtoReflect.Descriptor instead.
func (*ArtifactStoreGcpStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDescGZIP(), []int{3}
}

func (x *ArtifactStoreGcpStackOutputs) GetGcpArtifactRegistryStatus() *model2.ArtifactStoreGcpArtifactRegistryStatus {
	if x != nil {
		return x.GcpArtifactRegistryStatus
	}
	return nil
}

// stack response
type ArtifactStoreGcpStackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack-job progress event
	ProgressEvent *model3.StackJobProgressEvent `protobuf:"bytes,1,opt,name=progress_event,json=progressEvent,proto3" json:"progress_event,omitempty"`
	// stack outputs
	Outputs *ArtifactStoreGcpStackOutputs `protobuf:"bytes,3,opt,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *ArtifactStoreGcpStackResponse) Reset() {
	*x = ArtifactStoreGcpStackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactStoreGcpStackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactStoreGcpStackResponse) ProtoMessage() {}

func (x *ArtifactStoreGcpStackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactStoreGcpStackResponse.ProtoReflect.Descriptor instead.
func (*ArtifactStoreGcpStackResponse) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDescGZIP(), []int{4}
}

func (x *ArtifactStoreGcpStackResponse) GetProgressEvent() *model3.StackJobProgressEvent {
	if x != nil {
		return x.ProgressEvent
	}
	return nil
}

func (x *ArtifactStoreGcpStackResponse) GetOutputs() *ArtifactStoreGcpStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x67,
	0x63, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x46, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x48, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f,
	0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b,
	0x03, 0x0a, 0x1a, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x4c, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f,
	0x62, 0x52, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x9a, 0x01, 0x0a, 0x11,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x47, 0x63,
	0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x91, 0x01, 0x0a, 0x0e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x6a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x8f, 0x01, 0x0a,
	0x25, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x47, 0x63,
	0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x66, 0x0a, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x22, 0x98,
	0x01, 0x0a, 0x22, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x72, 0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x0d, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xc6, 0x01, 0x0a, 0x1c, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x47, 0x63, 0x70, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0xa5, 0x01, 0x0a, 0x1c, 0x67,
	0x63, 0x70, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x47, 0x63,
	0x70, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x19, 0x67, 0x63, 0x70, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x8d, 0x02, 0x0a, 0x1d, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x7e, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x42, 0x9f, 0x04, 0x0a, 0x54, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x49, 0x6f, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x76, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02,
	0x0a, 0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x41, 0x53, 0x47, 0x4d, 0xaa, 0x02, 0x46, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73,
	0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x47, 0x63, 0x70, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x46, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x5c,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x52,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x5c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x47, 0x63, 0x70,
	0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x4f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x3a, 0x3a, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x3a, 0x3a, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_goTypes = []interface{}{
	(*ArtifactStoreGcpStackInput)(nil),                    // 0: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackInput
	(*ArtifactStoreGcpStackCredentialsInput)(nil),         // 1: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackCredentialsInput
	(*ArtifactStoreGcpStackResourceInput)(nil),            // 2: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackResourceInput
	(*ArtifactStoreGcpStackOutputs)(nil),                  // 3: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackOutputs
	(*ArtifactStoreGcpStackResponse)(nil),                 // 4: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackResponse
	(*model.StackJob)(nil),                                // 5: cloud.planton.apis.v1.stack.job.model.StackJob
	(*model1.GoogleProviderCredential)(nil),               // 6: cloud.planton.apis.v1.commons.pulumi.operation.model.GoogleProviderCredential
	(*model2.ArtifactStore)(nil),                          // 7: cloud.planton.apis.v1.code2cloud.artifactstore.model.ArtifactStore
	(*model2.ArtifactStoreGcpArtifactRegistryStatus)(nil), // 8: cloud.planton.apis.v1.code2cloud.artifactstore.model.ArtifactStoreGcpArtifactRegistryStatus
	(*model3.StackJobProgressEvent)(nil),                  // 9: cloud.planton.apis.v1.stack.job.progress.model.StackJobProgressEvent
}
var file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackInput.stack_job:type_name -> cloud.planton.apis.v1.stack.job.model.StackJob
	1, // 1: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackInput.credentials_input:type_name -> cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackCredentialsInput
	2, // 2: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackInput.resource_input:type_name -> cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackResourceInput
	6, // 3: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackCredentialsInput.google:type_name -> cloud.planton.apis.v1.commons.pulumi.operation.model.GoogleProviderCredential
	7, // 4: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackResourceInput.artifact_store:type_name -> cloud.planton.apis.v1.code2cloud.artifactstore.model.ArtifactStore
	8, // 5: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackOutputs.gcp_artifact_registry_status:type_name -> cloud.planton.apis.v1.code2cloud.artifactstore.model.ArtifactStoreGcpArtifactRegistryStatus
	9, // 6: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackResponse.progress_event:type_name -> cloud.planton.apis.v1.stack.job.progress.model.StackJobProgressEvent
	3, // 7: cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackResponse.outputs:type_name -> cloud.planton.apis.v1.code2cloud.artifactstore.stack.gcp.model.ArtifactStoreGcpStackOutputs
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_init()
}
func file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactStoreGcpStackInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactStoreGcpStackCredentialsInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactStoreGcpStackResourceInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactStoreGcpStackOutputs); i {
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
		file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactStoreGcpStackResponse); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_stack_gcp_model_io_proto_depIdxs = nil
}
