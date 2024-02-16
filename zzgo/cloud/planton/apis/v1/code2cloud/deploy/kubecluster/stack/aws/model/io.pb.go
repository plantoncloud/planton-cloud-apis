// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/kubecluster/stack/aws/model/io.proto

package model

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/cloudaccount/model"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/kubecluster/model"
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

type KubeClusterAwsStackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack job
	StackJob *model.StackJob `protobuf:"bytes,1,opt,name=stack_job,json=stackJob,proto3" json:"stack_job,omitempty"`
	// pulumi stack credentials
	CredentialsInput *KubeClusterAwsStackCredentialsInput `protobuf:"bytes,2,opt,name=credentials_input,json=credentialsInput,proto3" json:"credentials_input,omitempty"`
	// inputs used for creating stack resources
	ResourceInput *KubeClusterAwsStackResourceInput `protobuf:"bytes,3,opt,name=resource_input,json=resourceInput,proto3" json:"resource_input,omitempty"`
}

func (x *KubeClusterAwsStackInput) Reset() {
	*x = KubeClusterAwsStackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeClusterAwsStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeClusterAwsStackInput) ProtoMessage() {}

func (x *KubeClusterAwsStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeClusterAwsStackInput.ProtoReflect.Descriptor instead.
func (*KubeClusterAwsStackInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *KubeClusterAwsStackInput) GetStackJob() *model.StackJob {
	if x != nil {
		return x.StackJob
	}
	return nil
}

func (x *KubeClusterAwsStackInput) GetCredentialsInput() *KubeClusterAwsStackCredentialsInput {
	if x != nil {
		return x.CredentialsInput
	}
	return nil
}

func (x *KubeClusterAwsStackInput) GetResourceInput() *KubeClusterAwsStackResourceInput {
	if x != nil {
		return x.ResourceInput
	}
	return nil
}

// stack credentials input
type KubeClusterAwsStackCredentialsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// aws provider credential input
	Aws *model1.AwsProviderCredential `protobuf:"bytes,1,opt,name=aws,proto3" json:"aws,omitempty"`
}

func (x *KubeClusterAwsStackCredentialsInput) Reset() {
	*x = KubeClusterAwsStackCredentialsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeClusterAwsStackCredentialsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeClusterAwsStackCredentialsInput) ProtoMessage() {}

func (x *KubeClusterAwsStackCredentialsInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeClusterAwsStackCredentialsInput.ProtoReflect.Descriptor instead.
func (*KubeClusterAwsStackCredentialsInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *KubeClusterAwsStackCredentialsInput) GetAws() *model1.AwsProviderCredential {
	if x != nil {
		return x.Aws
	}
	return nil
}

// stack resource input
type KubeClusterAwsStackResourceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// details of the kube-cluster
	KubeCluster *model2.KubeCluster `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster,proto3" json:"kube_cluster,omitempty"`
}

func (x *KubeClusterAwsStackResourceInput) Reset() {
	*x = KubeClusterAwsStackResourceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeClusterAwsStackResourceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeClusterAwsStackResourceInput) ProtoMessage() {}

func (x *KubeClusterAwsStackResourceInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeClusterAwsStackResourceInput.ProtoReflect.Descriptor instead.
func (*KubeClusterAwsStackResourceInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDescGZIP(), []int{2}
}

func (x *KubeClusterAwsStackResourceInput) GetKubeCluster() *model2.KubeCluster {
	if x != nil {
		return x.KubeCluster
	}
	return nil
}

// stack outputs
type KubeClusterAwsStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the vpc in which the kube-cluster is created
	ClusterVpcId string `protobuf:"bytes,1,opt,name=cluster_vpc_id,json=clusterVpcId,proto3" json:"cluster_vpc_id,omitempty"`
	// kube-cluster endpoint
	ClusterEndpoint string `protobuf:"bytes,2,opt,name=cluster_endpoint,json=clusterEndpoint,proto3" json:"cluster_endpoint,omitempty"`
	// kube-cluster ca data
	ClusterCaData string `protobuf:"bytes,3,opt,name=cluster_ca_data,json=clusterCaData,proto3" json:"cluster_ca_data,omitempty"`
}

func (x *KubeClusterAwsStackOutputs) Reset() {
	*x = KubeClusterAwsStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeClusterAwsStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeClusterAwsStackOutputs) ProtoMessage() {}

func (x *KubeClusterAwsStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeClusterAwsStackOutputs.ProtoReflect.Descriptor instead.
func (*KubeClusterAwsStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDescGZIP(), []int{3}
}

func (x *KubeClusterAwsStackOutputs) GetClusterVpcId() string {
	if x != nil {
		return x.ClusterVpcId
	}
	return ""
}

func (x *KubeClusterAwsStackOutputs) GetClusterEndpoint() string {
	if x != nil {
		return x.ClusterEndpoint
	}
	return ""
}

func (x *KubeClusterAwsStackOutputs) GetClusterCaData() string {
	if x != nil {
		return x.ClusterCaData
	}
	return ""
}

// stack response
type KubeClusterAwsStackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack-job progress event
	ProgressEvent *model3.StackJobProgressEvent `protobuf:"bytes,1,opt,name=progress_event,json=progressEvent,proto3" json:"progress_event,omitempty"`
	// stack outputs
	Outputs *KubeClusterAwsStackOutputs `protobuf:"bytes,3,opt,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *KubeClusterAwsStackResponse) Reset() {
	*x = KubeClusterAwsStackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeClusterAwsStackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeClusterAwsStackResponse) ProtoMessage() {}

func (x *KubeClusterAwsStackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeClusterAwsStackResponse.ProtoReflect.Descriptor instead.
func (*KubeClusterAwsStackResponse) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDescGZIP(), []int{4}
}

func (x *KubeClusterAwsStackResponse) GetProgressEvent() *model3.StackJobProgressEvent {
	if x != nil {
		return x.ProgressEvent
	}
	return nil
}

func (x *KubeClusterAwsStackResponse) GetOutputs() *KubeClusterAwsStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x77, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x43,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x1a, 0x3f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x75,
	0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d,
	0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x03, 0x0a, 0x18, 0x4b, 0x75, 0x62, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x77, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x4c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f,
	0x62, 0x12, 0x95, 0x01, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x68, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x77, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x8c, 0x01, 0x0a, 0x0e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61,
	0x77, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x77, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x23, 0x4b, 0x75, 0x62,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x77, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x5d, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x75,
	0x6c, 0x75, 0x6d, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x03, 0x61, 0x77, 0x73, 0x22,
	0x8d, 0x01, 0x0a, 0x20, 0x4b, 0x75, 0x62, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x77, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x69, 0x0a, 0x0c, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x0b, 0x6b, 0x75, 0x62, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22,
	0x95, 0x01, 0x0a, 0x1a, 0x4b, 0x75, 0x62, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x77, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x24,
	0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x76, 0x70, 0x63, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x56,
	0x70, 0x63, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x61, 0x44, 0x61, 0x74, 0x61, 0x22, 0x86, 0x02, 0x0a, 0x1b, 0x4b, 0x75, 0x62, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x77, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x45, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x79, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4b, 0x75, 0x62,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x77, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x42, 0x8d, 0x04, 0x0a, 0x51, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67,
	0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x77, 0x73,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x73, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a,
	0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x77, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x0a, 0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x4b,
	0x53, 0x41, 0x4d, 0xaa, 0x02, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x4b, 0x75,
	0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x41, 0x77, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x43, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56,
	0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2,
	0x02, 0x4f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x77, 0x73, 0x5c,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x4c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x41, 0x77, 0x73, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_goTypes = []interface{}{
	(*KubeClusterAwsStackInput)(nil),            // 0: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackInput
	(*KubeClusterAwsStackCredentialsInput)(nil), // 1: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackCredentialsInput
	(*KubeClusterAwsStackResourceInput)(nil),    // 2: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackResourceInput
	(*KubeClusterAwsStackOutputs)(nil),          // 3: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackOutputs
	(*KubeClusterAwsStackResponse)(nil),         // 4: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackResponse
	(*model.StackJob)(nil),                      // 5: cloud.planton.apis.v1.stack.job.model.StackJob
	(*model1.AwsProviderCredential)(nil),        // 6: cloud.planton.apis.v1.commons.pulumi.operation.model.AwsProviderCredential
	(*model2.KubeCluster)(nil),                  // 7: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.model.KubeCluster
	(*model3.StackJobProgressEvent)(nil),        // 8: cloud.planton.apis.v1.stack.job.progress.model.StackJobProgressEvent
}
var file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackInput.stack_job:type_name -> cloud.planton.apis.v1.stack.job.model.StackJob
	1, // 1: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackInput.credentials_input:type_name -> cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackCredentialsInput
	2, // 2: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackInput.resource_input:type_name -> cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackResourceInput
	6, // 3: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackCredentialsInput.aws:type_name -> cloud.planton.apis.v1.commons.pulumi.operation.model.AwsProviderCredential
	7, // 4: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackResourceInput.kube_cluster:type_name -> cloud.planton.apis.v1.code2cloud.deploy.kubecluster.model.KubeCluster
	8, // 5: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackResponse.progress_event:type_name -> cloud.planton.apis.v1.stack.job.progress.model.StackJobProgressEvent
	3, // 6: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackResponse.outputs:type_name -> cloud.planton.apis.v1.code2cloud.deploy.kubecluster.stack.aws.model.KubeClusterAwsStackOutputs
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeClusterAwsStackInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeClusterAwsStackCredentialsInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeClusterAwsStackResourceInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeClusterAwsStackOutputs); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeClusterAwsStackResponse); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_kubecluster_stack_aws_model_io_proto_depIdxs = nil
}
