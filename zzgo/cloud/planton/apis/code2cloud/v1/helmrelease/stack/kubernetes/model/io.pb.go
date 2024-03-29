// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/helmrelease/stack/kubernetes/model/io.proto

package model

import (
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/helmrelease/model"
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

// input for helm-release stack
type HelmReleaseKubernetesStackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack job
	StackJob *model.StackJob `protobuf:"bytes,1,opt,name=stack_job,json=stackJob,proto3" json:"stack_job,omitempty"`
	// pulumi stack credentials
	CredentialsInput *HelmReleaseKubernetesStackCredentialsInput `protobuf:"bytes,2,opt,name=credentials_input,json=credentialsInput,proto3" json:"credentials_input,omitempty"`
	// helm-release
	HelmRelease *model1.HelmRelease `protobuf:"bytes,3,opt,name=helm_release,json=helmRelease,proto3" json:"helm_release,omitempty"`
}

func (x *HelmReleaseKubernetesStackInput) Reset() {
	*x = HelmReleaseKubernetesStackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmReleaseKubernetesStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmReleaseKubernetesStackInput) ProtoMessage() {}

func (x *HelmReleaseKubernetesStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmReleaseKubernetesStackInput.ProtoReflect.Descriptor instead.
func (*HelmReleaseKubernetesStackInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *HelmReleaseKubernetesStackInput) GetStackJob() *model.StackJob {
	if x != nil {
		return x.StackJob
	}
	return nil
}

func (x *HelmReleaseKubernetesStackInput) GetCredentialsInput() *HelmReleaseKubernetesStackCredentialsInput {
	if x != nil {
		return x.CredentialsInput
	}
	return nil
}

func (x *HelmReleaseKubernetesStackInput) GetHelmRelease() *model1.HelmRelease {
	if x != nil {
		return x.HelmRelease
	}
	return nil
}

// stack credentials input
type HelmReleaseKubernetesStackCredentialsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// kubernetes provider credential for creating helm-release resources on kubernetes cluster
	Kubernetes *credentials.KubernetesProviderCredential `protobuf:"bytes,1,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
}

func (x *HelmReleaseKubernetesStackCredentialsInput) Reset() {
	*x = HelmReleaseKubernetesStackCredentialsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmReleaseKubernetesStackCredentialsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmReleaseKubernetesStackCredentialsInput) ProtoMessage() {}

func (x *HelmReleaseKubernetesStackCredentialsInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmReleaseKubernetesStackCredentialsInput.ProtoReflect.Descriptor instead.
func (*HelmReleaseKubernetesStackCredentialsInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *HelmReleaseKubernetesStackCredentialsInput) GetKubernetes() *credentials.KubernetesProviderCredential {
	if x != nil {
		return x.Kubernetes
	}
	return nil
}

// helm-release stack outputs
type HelmReleaseKubernetesStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HelmReleaseKubernetesStackOutputs) Reset() {
	*x = HelmReleaseKubernetesStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmReleaseKubernetesStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmReleaseKubernetesStackOutputs) ProtoMessage() {}

func (x *HelmReleaseKubernetesStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmReleaseKubernetesStackOutputs.ProtoReflect.Descriptor instead.
func (*HelmReleaseKubernetesStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDescGZIP(), []int{2}
}

// stack response
type HelmReleaseKubernetesStackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack-job progress event
	ProgressEvent *progress.StackJobProgressEvent `protobuf:"bytes,1,opt,name=progress_event,json=progressEvent,proto3" json:"progress_event,omitempty"`
	// stack outputs
	Outputs *HelmReleaseKubernetesStackOutputs `protobuf:"bytes,3,opt,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *HelmReleaseKubernetesStackResponse) Reset() {
	*x = HelmReleaseKubernetesStackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmReleaseKubernetesStackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmReleaseKubernetesStackResponse) ProtoMessage() {}

func (x *HelmReleaseKubernetesStackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmReleaseKubernetesStackResponse.ProtoReflect.Descriptor instead.
func (*HelmReleaseKubernetesStackResponse) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDescGZIP(), []int{3}
}

func (x *HelmReleaseKubernetesStackResponse) GetProgressEvent() *progress.StackJobProgressEvent {
	if x != nil {
		return x.ProgressEvent
	}
	return nil
}

func (x *HelmReleaseKubernetesStackResponse) GetOutputs() *HelmReleaseKubernetesStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x43,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x1a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x69, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x02, 0x0a, 0x1f, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x4f, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x52,
	0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x9c, 0x01, 0x0a, 0x11, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x48, 0x65, 0x6c, 0x6d,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x62, 0x0a, 0x0c, 0x68, 0x65, 0x6c, 0x6d,
	0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52,
	0x0b, 0x68, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x22, 0xa0, 0x01, 0x0a,
	0x2a, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x72, 0x0a, 0x0a, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x52, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22,
	0x23, 0x0a, 0x21, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x22, 0x98, 0x02, 0x0a, 0x22, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x0e, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x80, 0x01, 0x0a,
	0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x66,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x42,
	0x8b, 0x04, 0x0a, 0x51, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65,
	0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x73, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x43, 0x56, 0x48, 0x53, 0x4b,
	0x4d, 0xaa, 0x02, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x56, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x48, 0x65, 0x6c, 0x6d, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x4f,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31,
	0x5c, 0x48, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5c, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x4b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x48, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_goTypes = []interface{}{
	(*HelmReleaseKubernetesStackInput)(nil),            // 0: cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.model.HelmReleaseKubernetesStackInput
	(*HelmReleaseKubernetesStackCredentialsInput)(nil), // 1: cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.model.HelmReleaseKubernetesStackCredentialsInput
	(*HelmReleaseKubernetesStackOutputs)(nil),          // 2: cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.model.HelmReleaseKubernetesStackOutputs
	(*HelmReleaseKubernetesStackResponse)(nil),         // 3: cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.model.HelmReleaseKubernetesStackResponse
	(*model.StackJob)(nil),                             // 4: cloud.planton.apis.iac.v1.stackjob.model.StackJob
	(*model1.HelmRelease)(nil),                         // 5: cloud.planton.apis.code2cloud.v1.helmrelease.model.HelmRelease
	(*credentials.KubernetesProviderCredential)(nil),   // 6: cloud.planton.apis.iac.v1.stackjob.model.credentials.KubernetesProviderCredential
	(*progress.StackJobProgressEvent)(nil),             // 7: cloud.planton.apis.iac.v1.stackjob.model.progress.StackJobProgressEvent
}
var file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.model.HelmReleaseKubernetesStackInput.stack_job:type_name -> cloud.planton.apis.iac.v1.stackjob.model.StackJob
	1, // 1: cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.model.HelmReleaseKubernetesStackInput.credentials_input:type_name -> cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.model.HelmReleaseKubernetesStackCredentialsInput
	5, // 2: cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.model.HelmReleaseKubernetesStackInput.helm_release:type_name -> cloud.planton.apis.code2cloud.v1.helmrelease.model.HelmRelease
	6, // 3: cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.model.HelmReleaseKubernetesStackCredentialsInput.kubernetes:type_name -> cloud.planton.apis.iac.v1.stackjob.model.credentials.KubernetesProviderCredential
	7, // 4: cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.model.HelmReleaseKubernetesStackResponse.progress_event:type_name -> cloud.planton.apis.iac.v1.stackjob.model.progress.StackJobProgressEvent
	2, // 5: cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.model.HelmReleaseKubernetesStackResponse.outputs:type_name -> cloud.planton.apis.code2cloud.v1.helmrelease.stack.kubernetes.model.HelmReleaseKubernetesStackOutputs
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmReleaseKubernetesStackInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmReleaseKubernetesStackCredentialsInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmReleaseKubernetesStackOutputs); i {
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
		file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmReleaseKubernetesStackResponse); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_helmrelease_stack_kubernetes_model_io_proto_depIdxs = nil
}
