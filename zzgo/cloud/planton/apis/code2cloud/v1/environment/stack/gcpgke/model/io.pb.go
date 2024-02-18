// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/environment/stack/gcpgke/model/io.proto

package model

import (
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/environment/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model"
	credentials "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/credentials"
	progress "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/progress"
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

type EnvironmentGcpGkeStackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack job
	StackJob *model.StackJob `protobuf:"bytes,1,opt,name=stack_job,json=stackJob,proto3" json:"stack_job,omitempty"`
	// pulumi stack credentials
	CredentialsInput *EnvironmentGcpGkeStackCredentialsInput `protobuf:"bytes,2,opt,name=credentials_input,json=credentialsInput,proto3" json:"credentials_input,omitempty"`
	// inputs used for creating stack resources
	ResourceInput *EnvironmentGcpGkeStackResourceInput `protobuf:"bytes,3,opt,name=resource_input,json=resourceInput,proto3" json:"resource_input,omitempty"`
}

func (x *EnvironmentGcpGkeStackInput) Reset() {
	*x = EnvironmentGcpGkeStackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentGcpGkeStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentGcpGkeStackInput) ProtoMessage() {}

func (x *EnvironmentGcpGkeStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentGcpGkeStackInput.ProtoReflect.Descriptor instead.
func (*EnvironmentGcpGkeStackInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *EnvironmentGcpGkeStackInput) GetStackJob() *model.StackJob {
	if x != nil {
		return x.StackJob
	}
	return nil
}

func (x *EnvironmentGcpGkeStackInput) GetCredentialsInput() *EnvironmentGcpGkeStackCredentialsInput {
	if x != nil {
		return x.CredentialsInput
	}
	return nil
}

func (x *EnvironmentGcpGkeStackInput) GetResourceInput() *EnvironmentGcpGkeStackResourceInput {
	if x != nil {
		return x.ResourceInput
	}
	return nil
}

// stack credentials input
type EnvironmentGcpGkeStackCredentialsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// google provider credential input required to create secret-resources on gcp
	Google *credentials.GoogleProviderCredential `protobuf:"bytes,1,opt,name=google,proto3" json:"google,omitempty"`
	// kubernetes provider credential for creating endpoint-domain resources on kubernetes cluster
	Kubernetes *credentials.KubernetesProviderCredential `protobuf:"bytes,2,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
}

func (x *EnvironmentGcpGkeStackCredentialsInput) Reset() {
	*x = EnvironmentGcpGkeStackCredentialsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentGcpGkeStackCredentialsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentGcpGkeStackCredentialsInput) ProtoMessage() {}

func (x *EnvironmentGcpGkeStackCredentialsInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentGcpGkeStackCredentialsInput.ProtoReflect.Descriptor instead.
func (*EnvironmentGcpGkeStackCredentialsInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *EnvironmentGcpGkeStackCredentialsInput) GetGoogle() *credentials.GoogleProviderCredential {
	if x != nil {
		return x.Google
	}
	return nil
}

func (x *EnvironmentGcpGkeStackCredentialsInput) GetKubernetes() *credentials.KubernetesProviderCredential {
	if x != nil {
		return x.Kubernetes
	}
	return nil
}

// stack resource input
type EnvironmentGcpGkeStackResourceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// environment
	Environment *model1.Environment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *EnvironmentGcpGkeStackResourceInput) Reset() {
	*x = EnvironmentGcpGkeStackResourceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentGcpGkeStackResourceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentGcpGkeStackResourceInput) ProtoMessage() {}

func (x *EnvironmentGcpGkeStackResourceInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentGcpGkeStackResourceInput.ProtoReflect.Descriptor instead.
func (*EnvironmentGcpGkeStackResourceInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDescGZIP(), []int{2}
}

func (x *EnvironmentGcpGkeStackResourceInput) GetEnvironment() *model1.Environment {
	if x != nil {
		return x.Environment
	}
	return nil
}

// environment secrets outputs
type EnvironmentGcpGkeStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// environment secrets
	EnvironmentSecrets []*model1.EnvironmentSecret `protobuf:"bytes,1,rep,name=environment_secrets,json=environmentSecrets,proto3" json:"environment_secrets,omitempty"`
}

func (x *EnvironmentGcpGkeStackOutputs) Reset() {
	*x = EnvironmentGcpGkeStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentGcpGkeStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentGcpGkeStackOutputs) ProtoMessage() {}

func (x *EnvironmentGcpGkeStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentGcpGkeStackOutputs.ProtoReflect.Descriptor instead.
func (*EnvironmentGcpGkeStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDescGZIP(), []int{3}
}

func (x *EnvironmentGcpGkeStackOutputs) GetEnvironmentSecrets() []*model1.EnvironmentSecret {
	if x != nil {
		return x.EnvironmentSecrets
	}
	return nil
}

// stack response
type EnvironmentGcpGkeStackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack-job progress event
	ProgressEvent *progress.StackJobProgressEvent `protobuf:"bytes,1,opt,name=progress_event,json=progressEvent,proto3" json:"progress_event,omitempty"`
	// stack outputs
	Outputs *EnvironmentGcpGkeStackOutputs `protobuf:"bytes,3,opt,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *EnvironmentGcpGkeStackResponse) Reset() {
	*x = EnvironmentGcpGkeStackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentGcpGkeStackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentGcpGkeStackResponse) ProtoMessage() {}

func (x *EnvironmentGcpGkeStackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentGcpGkeStackResponse.ProtoReflect.Descriptor instead.
func (*EnvironmentGcpGkeStackResponse) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDescGZIP(), []int{4}
}

func (x *EnvironmentGcpGkeStackResponse) GetProgressEvent() *progress.StackJobProgressEvent {
	if x != nil {
		return x.ProgressEvent
	}
	return nil
}

func (x *EnvironmentGcpGkeStackResponse) GetOutputs() *EnvironmentGcpGkeStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDesc = []byte{
	0x0a, 0x48, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2f, 0x67, 0x63, 0x70, 0x67, 0x6b, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x67,
	0x63, 0x70, 0x67, 0x6b, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x3e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69,
	0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a,
	0x1b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x63, 0x70, 0x47,
	0x6b, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x4f, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x4a, 0x6f, 0x62, 0x52, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x94, 0x01,
	0x0a, 0x11, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x67, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x67,
	0x63, 0x70, 0x67, 0x6b, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x63, 0x70, 0x47, 0x6b, 0x65, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x8b, 0x01, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x64, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x67, 0x63, 0x70, 0x67, 0x6b, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x63, 0x70, 0x47, 0x6b,
	0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x22, 0x84, 0x02, 0x0a, 0x26, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x47, 0x63, 0x70, 0x47, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x66, 0x0a,
	0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a,
	0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x06, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x72, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69,
	0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x0a, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x23, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x63, 0x70, 0x47, 0x6b, 0x65, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x61, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x1d, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x47, 0x63, 0x70, 0x47, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x76, 0x0a, 0x13, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x12, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x22, 0x8b,
	0x02, 0x0a, 0x1e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x63,
	0x70, 0x47, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6f, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69,
	0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x78, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x5e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x67, 0x63, 0x70, 0x67, 0x6b, 0x65, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x47, 0x63, 0x70, 0x47, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x42, 0xf3, 0x03, 0x0a,
	0x4d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x67, 0x63, 0x70, 0x67, 0x6b, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07,
	0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x67, 0x63,
	0x70, 0x67, 0x6b, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41,
	0x43, 0x56, 0x45, 0x53, 0x47, 0x4d, 0xaa, 0x02, 0x3f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x47, 0x63, 0x70, 0x67,
	0x6b, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x3f, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x47, 0x63,
	0x70, 0x67, 0x6b, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x4b, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c,
	0x47, 0x63, 0x70, 0x67, 0x6b, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x47, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a,
	0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x67, 0x6b, 0x65, 0x3a, 0x3a, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_goTypes = []interface{}{
	(*EnvironmentGcpGkeStackInput)(nil),              // 0: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackInput
	(*EnvironmentGcpGkeStackCredentialsInput)(nil),   // 1: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackCredentialsInput
	(*EnvironmentGcpGkeStackResourceInput)(nil),      // 2: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackResourceInput
	(*EnvironmentGcpGkeStackOutputs)(nil),            // 3: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackOutputs
	(*EnvironmentGcpGkeStackResponse)(nil),           // 4: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackResponse
	(*model.StackJob)(nil),                           // 5: cloud.planton.apis.iac.v1.stackjob.model.StackJob
	(*credentials.GoogleProviderCredential)(nil),     // 6: cloud.planton.apis.iac.v1.stackjob.model.credentials.GoogleProviderCredential
	(*credentials.KubernetesProviderCredential)(nil), // 7: cloud.planton.apis.iac.v1.stackjob.model.credentials.KubernetesProviderCredential
	(*model1.Environment)(nil),                       // 8: cloud.planton.apis.code2cloud.v1.environment.model.Environment
	(*model1.EnvironmentSecret)(nil),                 // 9: cloud.planton.apis.code2cloud.v1.environment.model.EnvironmentSecret
	(*progress.StackJobProgressEvent)(nil),           // 10: cloud.planton.apis.iac.v1.stackjob.model.progress.StackJobProgressEvent
}
var file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_depIdxs = []int32{
	5,  // 0: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackInput.stack_job:type_name -> cloud.planton.apis.iac.v1.stackjob.model.StackJob
	1,  // 1: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackInput.credentials_input:type_name -> cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackCredentialsInput
	2,  // 2: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackInput.resource_input:type_name -> cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackResourceInput
	6,  // 3: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackCredentialsInput.google:type_name -> cloud.planton.apis.iac.v1.stackjob.model.credentials.GoogleProviderCredential
	7,  // 4: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackCredentialsInput.kubernetes:type_name -> cloud.planton.apis.iac.v1.stackjob.model.credentials.KubernetesProviderCredential
	8,  // 5: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackResourceInput.environment:type_name -> cloud.planton.apis.code2cloud.v1.environment.model.Environment
	9,  // 6: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackOutputs.environment_secrets:type_name -> cloud.planton.apis.code2cloud.v1.environment.model.EnvironmentSecret
	10, // 7: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackResponse.progress_event:type_name -> cloud.planton.apis.iac.v1.stackjob.model.progress.StackJobProgressEvent
	3,  // 8: cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackResponse.outputs:type_name -> cloud.planton.apis.code2cloud.v1.environment.stack.gcpgke.model.EnvironmentGcpGkeStackOutputs
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentGcpGkeStackInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentGcpGkeStackCredentialsInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentGcpGkeStackResourceInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentGcpGkeStackOutputs); i {
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
		file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentGcpGkeStackResponse); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_environment_stack_gcpgke_model_io_proto_depIdxs = nil
}
