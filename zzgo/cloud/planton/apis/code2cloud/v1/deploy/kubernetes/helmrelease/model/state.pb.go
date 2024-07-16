// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/kubernetes/helmrelease/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemessageoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions"
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

// helm-release
type HelmRelease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// resource metadata
	Metadata *model.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *HelmReleaseSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *HelmReleaseStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HelmRelease) Reset() {
	*x = HelmRelease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmRelease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmRelease) ProtoMessage() {}

func (x *HelmRelease) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmRelease.ProtoReflect.Descriptor instead.
func (*HelmRelease) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *HelmRelease) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *HelmRelease) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *HelmRelease) GetMetadata() *model.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *HelmRelease) GetSpec() *HelmReleaseSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *HelmRelease) GetStatus() *HelmReleaseStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// helm-release spec
type HelmReleaseSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// environment-info
	EnvironmentInfo *model1.ApiResourceEnvironmentInfo `protobuf:"bytes,1,opt,name=environment_info,json=environmentInfo,proto3" json:"environment_info,omitempty"`
	// helm-chart spec
	HelmChart *HelmReleaseSpecChartSpec `protobuf:"bytes,2,opt,name=helm_chart,json=helmChart,proto3" json:"helm_chart,omitempty"`
}

func (x *HelmReleaseSpec) Reset() {
	*x = HelmReleaseSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmReleaseSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmReleaseSpec) ProtoMessage() {}

func (x *HelmReleaseSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmReleaseSpec.ProtoReflect.Descriptor instead.
func (*HelmReleaseSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *HelmReleaseSpec) GetEnvironmentInfo() *model1.ApiResourceEnvironmentInfo {
	if x != nil {
		return x.EnvironmentInfo
	}
	return nil
}

func (x *HelmReleaseSpec) GetHelmChart() *HelmReleaseSpecChartSpec {
	if x != nil {
		return x.HelmChart
	}
	return nil
}

// helm-release status.
type HelmReleaseStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource lifecycle
	Lifecycle *model.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// resource audit info
	Audit *model.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// id of the stack-job
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
	// stack-outputs
	// helm-release stack-outputs
	StackOutputs *HelmReleaseStatusStackOutputs `protobuf:"bytes,1,opt,name=stack_outputs,json=stackOutputs,proto3" json:"stack_outputs,omitempty"`
}

func (x *HelmReleaseStatus) Reset() {
	*x = HelmReleaseStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmReleaseStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmReleaseStatus) ProtoMessage() {}

func (x *HelmReleaseStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmReleaseStatus.ProtoReflect.Descriptor instead.
func (*HelmReleaseStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *HelmReleaseStatus) GetLifecycle() *model.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *HelmReleaseStatus) GetAudit() *model.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *HelmReleaseStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

func (x *HelmReleaseStatus) GetStackOutputs() *HelmReleaseStatusStackOutputs {
	if x != nil {
		return x.StackOutputs
	}
	return nil
}

// helm-chart spec
type HelmReleaseSpecChartSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// helm chart repo
	Repo string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	// helm chart name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// helm chart version
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// map of key values for helm input
	Values map[string]string `protobuf:"bytes,4,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HelmReleaseSpecChartSpec) Reset() {
	*x = HelmReleaseSpecChartSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmReleaseSpecChartSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmReleaseSpecChartSpec) ProtoMessage() {}

func (x *HelmReleaseSpecChartSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmReleaseSpecChartSpec.ProtoReflect.Descriptor instead.
func (*HelmReleaseSpecChartSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDescGZIP(), []int{3}
}

func (x *HelmReleaseSpecChartSpec) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *HelmReleaseSpecChartSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelmReleaseSpecChartSpec) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *HelmReleaseSpecChartSpec) GetValues() map[string]string {
	if x != nil {
		return x.Values
	}
	return nil
}

// helm-release status on kubernetes cluster.
type HelmReleaseStatusStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// namespace in which the helm release has been created.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *HelmReleaseStatusStackOutputs) Reset() {
	*x = HelmReleaseStatusStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmReleaseStatusStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmReleaseStatusStackOutputs) ProtoMessage() {}

func (x *HelmReleaseStatusStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmReleaseStatusStackOutputs.ProtoReflect.Descriptor instead.
func (*HelmReleaseStatusStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDescGZIP(), []int{4}
}

func (x *HelmReleaseStatusStackOutputs) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDesc = []byte{
	0x0a, 0x50, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x44, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x7c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x07, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0xcf, 0x04, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xef, 0x03, 0xba, 0x48, 0xeb,
	0x03, 0xba, 0x01, 0x6c, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x38, 0x4f, 0x6e, 0x6c, 0x79, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x63,
	0x61, 0x73, 0x65, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x68, 0x79, 0x70, 0x68, 0x65, 0x6e,
	0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x1a, 0x21, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x2b, 0x24, 0x27, 0x29,
	0xba, 0x01, 0x53, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x4d, 0x75, 0x73, 0x74, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x20, 0x77,
	0x69, 0x74, 0x68, 0x20, 0x61, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x20,
	0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x1e, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a,
	0x5d, 0x2e, 0x2a, 0x24, 0x27, 0x29, 0xba, 0x01, 0x47, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x4d, 0x75, 0x73, 0x74, 0x20, 0x6e,
	0x6f, 0x74, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20, 0x68, 0x79,
	0x70, 0x68, 0x65, 0x6e, 0x1a, 0x1a, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5b, 0x5e, 0x2d, 0x5d, 0x24, 0x27, 0x29,
	0xba, 0x01, 0x6c, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2d, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65,
	0x20, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x20, 0x31, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x33,
	0x30, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x20, 0x6c, 0x6f, 0x6e,
	0x67, 0x1a, 0x2c, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x20, 0x26, 0x26, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20, 0x3c, 0x3d, 0x20, 0x33, 0x30, 0xba,
	0x01, 0x67, 0x0a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x69, 0x73,
	0x20, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19,
	0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x29, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x69, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x48, 0x65, 0x6c, 0x6d,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x6f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x57, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x3a, 0x2c, 0x90, 0xb5, 0x18, 0x01, 0x88, 0xa6, 0x1d, 0x13, 0x92, 0xa6, 0x1d, 0x20,
	0x08, 0x0a, 0x12, 0x1c, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x69, 0x64,
	0x22, 0x99, 0x02, 0x0a, 0x0f, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x7e, 0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x53,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x85, 0x01, 0x0a, 0x0a, 0x68, 0x65, 0x6c, 0x6d, 0x5f, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5e, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x68,
	0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x09, 0x68, 0x65, 0x6c, 0x6d, 0x43, 0x68, 0x61, 0x72, 0x74, 0x22, 0xf8, 0x02, 0x0a,
	0x11, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x60, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18,
	0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x12, 0x54, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x62, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x61, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x88, 0x01, 0x0a,
	0x0d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x63, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6d, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x48, 0x65, 0x6c, 0x6d,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22, 0x9c, 0x02, 0x0a, 0x18, 0x48, 0x65, 0x6c, 0x6d,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x82, 0x01, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x6a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c,
	0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x48,
	0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x43, 0x68,
	0x61, 0x72, 0x74, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x1d, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x94, 0x04, 0x0a, 0x52, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6d, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x74, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x68,
	0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x43, 0x56, 0x44, 0x4b, 0x48, 0x4d, 0xaa, 0x02, 0x44, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69,
	0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0xca, 0x02, 0x44, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x48, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x50, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c,
	0x48, 0x65, 0x6c, 0x6d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x4c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a,
	0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x48, 0x65, 0x6c, 0x6d, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_goTypes = []interface{}{
	(*HelmRelease)(nil),                       // 0: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmRelease
	(*HelmReleaseSpec)(nil),                   // 1: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseSpec
	(*HelmReleaseStatus)(nil),                 // 2: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseStatus
	(*HelmReleaseSpecChartSpec)(nil),          // 3: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseSpecChartSpec
	(*HelmReleaseStatusStackOutputs)(nil),     // 4: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseStatusStackOutputs
	nil,                                       // 5: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseSpecChartSpec.ValuesEntry
	(*model.ApiResourceMetadata)(nil),         // 6: cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	(*model1.ApiResourceEnvironmentInfo)(nil), // 7: cloud.planton.apis.resourcemanager.v1.environment.model.ApiResourceEnvironmentInfo
	(*model.ApiResourceLifecycle)(nil),        // 8: cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	(*model.ApiResourceAudit)(nil),            // 9: cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
}
var file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_depIdxs = []int32{
	6, // 0: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmRelease.metadata:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	1, // 1: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmRelease.spec:type_name -> cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseSpec
	2, // 2: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmRelease.status:type_name -> cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseStatus
	7, // 3: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseSpec.environment_info:type_name -> cloud.planton.apis.resourcemanager.v1.environment.model.ApiResourceEnvironmentInfo
	3, // 4: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseSpec.helm_chart:type_name -> cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseSpecChartSpec
	8, // 5: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseStatus.lifecycle:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	9, // 6: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseStatus.audit:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	4, // 7: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseStatus.stack_outputs:type_name -> cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseStatusStackOutputs
	5, // 8: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseSpecChartSpec.values:type_name -> cloud.planton.apis.code2cloud.v1.deploy.kubernetes.helmrelease.model.HelmReleaseSpecChartSpec.ValuesEntry
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_init()
}
func file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmRelease); i {
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
		file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmReleaseSpec); i {
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
		file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmReleaseStatus); i {
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
		file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmReleaseSpecChartSpec); i {
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
		file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmReleaseStatusStackOutputs); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_helmrelease_model_state_proto_depIdxs = nil
}
