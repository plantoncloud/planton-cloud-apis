// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/argocdinstance/stack/kubernetes/model/io.proto

package model

import (
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/argocdinstance/model"
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

// input for argocd-instance stack
type ArgoCDInstanceKubernetesStackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack job
	StackJob *model.StackJob `protobuf:"bytes,1,opt,name=stack_job,json=stackJob,proto3" json:"stack_job,omitempty"`
	// pulumi stack credentials
	CredentialsInput *ArgoCDInstanceKubernetesStackCredentialsInput `protobuf:"bytes,2,opt,name=credentials_input,json=credentialsInput,proto3" json:"credentials_input,omitempty"`
	// inputs used for creating stack resources
	ResourceInput *ArgoCDInstanceKubernetesStackResourceInput `protobuf:"bytes,3,opt,name=resource_input,json=resourceInput,proto3" json:"resource_input,omitempty"`
}

func (x *ArgoCDInstanceKubernetesStackInput) Reset() {
	*x = ArgoCDInstanceKubernetesStackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArgoCDInstanceKubernetesStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArgoCDInstanceKubernetesStackInput) ProtoMessage() {}

func (x *ArgoCDInstanceKubernetesStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArgoCDInstanceKubernetesStackInput.ProtoReflect.Descriptor instead.
func (*ArgoCDInstanceKubernetesStackInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *ArgoCDInstanceKubernetesStackInput) GetStackJob() *model.StackJob {
	if x != nil {
		return x.StackJob
	}
	return nil
}

func (x *ArgoCDInstanceKubernetesStackInput) GetCredentialsInput() *ArgoCDInstanceKubernetesStackCredentialsInput {
	if x != nil {
		return x.CredentialsInput
	}
	return nil
}

func (x *ArgoCDInstanceKubernetesStackInput) GetResourceInput() *ArgoCDInstanceKubernetesStackResourceInput {
	if x != nil {
		return x.ResourceInput
	}
	return nil
}

// stack credentials input
type ArgoCDInstanceKubernetesStackCredentialsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// kubernetes provider credential for creating argocd-instance resources on kubernetes cluster
	Kubernetes *credentials.KubernetesProviderCredential `protobuf:"bytes,1,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
}

func (x *ArgoCDInstanceKubernetesStackCredentialsInput) Reset() {
	*x = ArgoCDInstanceKubernetesStackCredentialsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArgoCDInstanceKubernetesStackCredentialsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArgoCDInstanceKubernetesStackCredentialsInput) ProtoMessage() {}

func (x *ArgoCDInstanceKubernetesStackCredentialsInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArgoCDInstanceKubernetesStackCredentialsInput.ProtoReflect.Descriptor instead.
func (*ArgoCDInstanceKubernetesStackCredentialsInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *ArgoCDInstanceKubernetesStackCredentialsInput) GetKubernetes() *credentials.KubernetesProviderCredential {
	if x != nil {
		return x.Kubernetes
	}
	return nil
}

// stack resource input
type ArgoCDInstanceKubernetesStackResourceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// argocd-instance
	ArgocdInstance *model1.ArgoCDInstance `protobuf:"bytes,1,opt,name=argocd_instance,json=argocdInstance,proto3" json:"argocd_instance,omitempty"`
}

func (x *ArgoCDInstanceKubernetesStackResourceInput) Reset() {
	*x = ArgoCDInstanceKubernetesStackResourceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArgoCDInstanceKubernetesStackResourceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArgoCDInstanceKubernetesStackResourceInput) ProtoMessage() {}

func (x *ArgoCDInstanceKubernetesStackResourceInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArgoCDInstanceKubernetesStackResourceInput.ProtoReflect.Descriptor instead.
func (*ArgoCDInstanceKubernetesStackResourceInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDescGZIP(), []int{2}
}

func (x *ArgoCDInstanceKubernetesStackResourceInput) GetArgocdInstance() *model1.ArgoCDInstance {
	if x != nil {
		return x.ArgocdInstance
	}
	return nil
}

// argocd-instance stack outputs
type ArgoCDInstanceKubernetesStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// status of the argocd-instance upon stack-job
	ArgocdInstanceStatus *model1.ArgoCDInstanceStatus `protobuf:"bytes,1,opt,name=argocd_instance_status,json=argocdInstanceStatus,proto3" json:"argocd_instance_status,omitempty"`
}

func (x *ArgoCDInstanceKubernetesStackOutputs) Reset() {
	*x = ArgoCDInstanceKubernetesStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArgoCDInstanceKubernetesStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArgoCDInstanceKubernetesStackOutputs) ProtoMessage() {}

func (x *ArgoCDInstanceKubernetesStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArgoCDInstanceKubernetesStackOutputs.ProtoReflect.Descriptor instead.
func (*ArgoCDInstanceKubernetesStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDescGZIP(), []int{3}
}

func (x *ArgoCDInstanceKubernetesStackOutputs) GetArgocdInstanceStatus() *model1.ArgoCDInstanceStatus {
	if x != nil {
		return x.ArgocdInstanceStatus
	}
	return nil
}

// stack response
type ArgoCDInstanceKubernetesStackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack-job progress event
	ProgressEvent *progress.StackJobProgressEvent `protobuf:"bytes,1,opt,name=progress_event,json=progressEvent,proto3" json:"progress_event,omitempty"`
	// stack outputs
	Outputs *ArgoCDInstanceKubernetesStackOutputs `protobuf:"bytes,3,opt,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *ArgoCDInstanceKubernetesStackResponse) Reset() {
	*x = ArgoCDInstanceKubernetesStackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArgoCDInstanceKubernetesStackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArgoCDInstanceKubernetesStackResponse) ProtoMessage() {}

func (x *ArgoCDInstanceKubernetesStackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArgoCDInstanceKubernetesStackResponse.ProtoReflect.Descriptor instead.
func (*ArgoCDInstanceKubernetesStackResponse) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDescGZIP(), []int{4}
}

func (x *ArgoCDInstanceKubernetesStackResponse) GetProgressEvent() *progress.StackJobProgressEvent {
	if x != nil {
		return x.ProgressEvent
	}
	return nil
}

func (x *ArgoCDInstanceKubernetesStackResponse) GetOutputs() *ArgoCDInstanceKubernetesStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x46, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x67, 0x6f,
	0x63, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x03,
	0x0a, 0x22, 0x41, 0x72, 0x67, 0x6f, 0x43, 0x44, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x4f, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f,
	0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x08, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0xa2, 0x01, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x75, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x72, 0x67, 0x6f, 0x43,
	0x44, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x99, 0x01, 0x0a, 0x0e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x72, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x72, 0x67,
	0x6f, 0x43, 0x44, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0xa3, 0x01, 0x0a, 0x2d, 0x41, 0x72, 0x67, 0x6f, 0x43,
	0x44, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x72, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f,
	0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0x9c, 0x01, 0x0a,
	0x2a, 0x41, 0x72, 0x67, 0x6f, 0x43, 0x44, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x6e, 0x0a, 0x0f, 0x61,
	0x72, 0x67, 0x6f, 0x63, 0x64, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x72, 0x67,
	0x6f, 0x43, 0x44, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x61, 0x72, 0x67,
	0x6f, 0x63, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x24,
	0x41, 0x72, 0x67, 0x6f, 0x43, 0x44, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x12, 0x81, 0x01, 0x0a, 0x16, 0x61, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x72,
	0x67, 0x6f, 0x43, 0x44, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x14, 0x61, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa1, 0x02, 0x0a, 0x25, 0x41, 0x72, 0x67,
	0x6f, 0x43, 0x44, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6f, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x86, 0x01, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x6c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x72, 0x67, 0x6f, 0x43, 0x44, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x42, 0x9d, 0x04, 0x0a,
	0x54, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x76, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x43, 0x56,
	0x41, 0x53, 0x4b, 0x4d, 0xaa, 0x02, 0x46, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x46,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31,
	0x5c, 0x41, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5c,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x52, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x41, 0x72, 0x67, 0x6f, 0x63, 0x64,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x4e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x56, 0x31, 0x3a, 0x3a, 0x41, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_goTypes = []interface{}{
	(*ArgoCDInstanceKubernetesStackInput)(nil),            // 0: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackInput
	(*ArgoCDInstanceKubernetesStackCredentialsInput)(nil), // 1: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackCredentialsInput
	(*ArgoCDInstanceKubernetesStackResourceInput)(nil),    // 2: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackResourceInput
	(*ArgoCDInstanceKubernetesStackOutputs)(nil),          // 3: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackOutputs
	(*ArgoCDInstanceKubernetesStackResponse)(nil),         // 4: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackResponse
	(*model.StackJob)(nil),                                // 5: cloud.planton.apis.iac.v1.stackjob.model.StackJob
	(*credentials.KubernetesProviderCredential)(nil),      // 6: cloud.planton.apis.iac.v1.stackjob.model.credentials.KubernetesProviderCredential
	(*model1.ArgoCDInstance)(nil),                         // 7: cloud.planton.apis.code2cloud.v1.argocdinstance.model.ArgoCDInstance
	(*model1.ArgoCDInstanceStatus)(nil),                   // 8: cloud.planton.apis.code2cloud.v1.argocdinstance.model.ArgoCDInstanceStatus
	(*progress.StackJobProgressEvent)(nil),                // 9: cloud.planton.apis.iac.v1.stackjob.model.progress.StackJobProgressEvent
}
var file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackInput.stack_job:type_name -> cloud.planton.apis.iac.v1.stackjob.model.StackJob
	1, // 1: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackInput.credentials_input:type_name -> cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackCredentialsInput
	2, // 2: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackInput.resource_input:type_name -> cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackResourceInput
	6, // 3: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackCredentialsInput.kubernetes:type_name -> cloud.planton.apis.iac.v1.stackjob.model.credentials.KubernetesProviderCredential
	7, // 4: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackResourceInput.argocd_instance:type_name -> cloud.planton.apis.code2cloud.v1.argocdinstance.model.ArgoCDInstance
	8, // 5: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackOutputs.argocd_instance_status:type_name -> cloud.planton.apis.code2cloud.v1.argocdinstance.model.ArgoCDInstanceStatus
	9, // 6: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackResponse.progress_event:type_name -> cloud.planton.apis.iac.v1.stackjob.model.progress.StackJobProgressEvent
	3, // 7: cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackResponse.outputs:type_name -> cloud.planton.apis.code2cloud.v1.argocdinstance.stack.kubernetes.model.ArgoCDInstanceKubernetesStackOutputs
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_init()
}
func file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArgoCDInstanceKubernetesStackInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArgoCDInstanceKubernetesStackCredentialsInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArgoCDInstanceKubernetesStackResourceInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArgoCDInstanceKubernetesStackOutputs); i {
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
		file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArgoCDInstanceKubernetesStackResponse); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_argocdinstance_stack_kubernetes_model_io_proto_depIdxs = nil
}