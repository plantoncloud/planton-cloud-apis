// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/aws/ekscluster/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	eksworkersmanagementmode "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/ekscluster/enums/eksworkersmanagementmode"
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

// eks-cluster
type EksCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// resource metadata
	Metadata *model.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// resource spec
	Spec *EksClusterSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// resource status
	Status *EksClusterStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *EksCluster) Reset() {
	*x = EksCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EksCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EksCluster) ProtoMessage() {}

func (x *EksCluster) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EksCluster.ProtoReflect.Descriptor instead.
func (*EksCluster) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *EksCluster) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *EksCluster) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *EksCluster) GetMetadata() *model.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *EksCluster) GetSpec() *EksClusterSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *EksCluster) GetStatus() *EksClusterStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// EksClusterSpec is a message type that defines the specifications for a eks-cluster on Planton Cloud.
type EksClusterSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// environment-info
	EnvironmentInfo *model1.ApiResourceEnvironmentInfo `protobuf:"bytes,99,opt,name=environment_info,json=environmentInfo,proto3" json:"environment_info,omitempty"`
	// stack-job settings
	StackJobSettings *model2.StackJobSettings `protobuf:"bytes,98,opt,name=stack_job_settings,json=stackJobSettings,proto3" json:"stack_job_settings,omitempty"`
	// aws-credential-id to be used for setting up aws-provider in stack-job
	AwsCredentialId string `protobuf:"bytes,97,opt,name=aws_credential_id,json=awsCredentialId,proto3" json:"aws_credential_id,omitempty"`
	// valid aws region in which to create the eks-cluster.
	// warning: eks-cluster will recreated if this value is updated.
	// https://aws.amazon.com/about-aws/global-infrastructure/regions_az/
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// (optional) id of the vpc to be used for creating eks-cluster.
	// if an id is not provided, a new vpc will be created.
	// warning: eks-cluster will be recreated if this is updated.
	VpcId string `protobuf:"bytes,2,opt,name=vpc_id,json=vpcId,proto3" json:"vpc_id,omitempty"`
	// aws eks worker node-groups management mode
	WorkersManagementMode eksworkersmanagementmode.EksWorkersManagementMode `protobuf:"varint,3,opt,name=workers_management_mode,json=workersManagementMode,proto3,enum=cloud.planton.apis.code2cloud.v1.aws.ekscluster.enums.eksworkersmanagementmode.EksWorkersManagementMode" json:"workers_management_mode,omitempty"`
}

func (x *EksClusterSpec) Reset() {
	*x = EksClusterSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EksClusterSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EksClusterSpec) ProtoMessage() {}

func (x *EksClusterSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EksClusterSpec.ProtoReflect.Descriptor instead.
func (*EksClusterSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *EksClusterSpec) GetEnvironmentInfo() *model1.ApiResourceEnvironmentInfo {
	if x != nil {
		return x.EnvironmentInfo
	}
	return nil
}

func (x *EksClusterSpec) GetStackJobSettings() *model2.StackJobSettings {
	if x != nil {
		return x.StackJobSettings
	}
	return nil
}

func (x *EksClusterSpec) GetAwsCredentialId() string {
	if x != nil {
		return x.AwsCredentialId
	}
	return ""
}

func (x *EksClusterSpec) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *EksClusterSpec) GetVpcId() string {
	if x != nil {
		return x.VpcId
	}
	return ""
}

func (x *EksClusterSpec) GetWorkersManagementMode() eksworkersmanagementmode.EksWorkersManagementMode {
	if x != nil {
		return x.WorkersManagementMode
	}
	return eksworkersmanagementmode.EksWorkersManagementMode(0)
}

// eks-cluster status
type EksClusterStatus struct {
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
	//
	//	eks-cluster stack-outputs
	StackOutputs *EksClusterStatusStackOutputs `protobuf:"bytes,1,opt,name=stack_outputs,json=stackOutputs,proto3" json:"stack_outputs,omitempty"`
}

func (x *EksClusterStatus) Reset() {
	*x = EksClusterStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EksClusterStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EksClusterStatus) ProtoMessage() {}

func (x *EksClusterStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EksClusterStatus.ProtoReflect.Descriptor instead.
func (*EksClusterStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *EksClusterStatus) GetLifecycle() *model.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *EksClusterStatus) GetAudit() *model.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *EksClusterStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

func (x *EksClusterStatus) GetStackOutputs() *EksClusterStatusStackOutputs {
	if x != nil {
		return x.StackOutputs
	}
	return nil
}

// eks-cluster stack-outputs
type EksClusterStatusStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the vpc in which the eks-cluster is created
	ClusterVpcId string `protobuf:"bytes,1,opt,name=cluster_vpc_id,json=clusterVpcId,proto3" json:"cluster_vpc_id,omitempty"`
	// eks-cluster endpoint.
	ClusterEndpoint string `protobuf:"bytes,2,opt,name=cluster_endpoint,json=clusterEndpoint,proto3" json:"cluster_endpoint,omitempty"`
	// eks-cluster certificate-authority-data.
	// this value is updated upon successful eks-cluster creation stack-job.
	ClusterCaData string `protobuf:"bytes,3,opt,name=cluster_ca_data,json=clusterCaData,proto3" json:"cluster_ca_data,omitempty"`
}

func (x *EksClusterStatusStackOutputs) Reset() {
	*x = EksClusterStatusStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EksClusterStatusStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EksClusterStatusStackOutputs) ProtoMessage() {}

func (x *EksClusterStatusStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EksClusterStatusStackOutputs.ProtoReflect.Descriptor instead.
func (*EksClusterStatusStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDescGZIP(), []int{3}
}

func (x *EksClusterStatusStackOutputs) GetClusterVpcId() string {
	if x != nil {
		return x.ClusterVpcId
	}
	return ""
}

func (x *EksClusterStatusStackOutputs) GetClusterEndpoint() string {
	if x != nil {
		return x.ClusterEndpoint
	}
	return ""
}

func (x *EksClusterStatusStackOutputs) GetClusterCaData() string {
	if x != nil {
		return x.ClusterCaData
	}
	return ""
}

var File_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDesc = []byte{
	0x0a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x65, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x65, 0x6b, 0x73, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x65, 0x6b,
	0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65,
	0x6b, 0x73, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x2f, 0x65, 0x6b, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6b, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x7c, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfd, 0x06, 0x0a, 0x0a, 0x45, 0x6b, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0xcf, 0x04, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xef, 0x03, 0xba,
	0x48, 0xeb, 0x03, 0xba, 0x01, 0x6c, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x4f, 0x6e, 0x6c, 0x79, 0x20, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x63, 0x61, 0x73, 0x65, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x68, 0x79, 0x70, 0x68,
	0x65, 0x6e, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x1a,
	0x21, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x2b, 0x24,
	0x27, 0x29, 0xba, 0x01, 0x53, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x4d, 0x75, 0x73, 0x74, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x63, 0x61, 0x73,
	0x65, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x1e, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61,
	0x2d, 0x7a, 0x5d, 0x2e, 0x2a, 0x24, 0x27, 0x29, 0xba, 0x01, 0x47, 0x0a, 0x0d, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x4d, 0x75, 0x73, 0x74,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20,
	0x68, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x1a, 0x1a, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5b, 0x5e, 0x2d, 0x5d, 0x24,
	0x27, 0x29, 0xba, 0x01, 0x6c, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20,
	0x62, 0x65, 0x20, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x20, 0x31, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x33, 0x30, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x20, 0x6c,
	0x6f, 0x6e, 0x67, 0x1a, 0x2c, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x20, 0x26, 0x26, 0x20, 0x73, 0x69, 0x7a, 0x65,
	0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20, 0x3c, 0x3d, 0x20, 0x33,
	0x30, 0xba, 0x01, 0x67, 0x0a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20,
	0x69, 0x73, 0x20, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x63, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x19, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x29, 0xc8, 0x01, 0x01, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x59, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x65, 0x6b,
	0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45,
	0x6b, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x12, 0x5f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x65, 0x6b, 0x73, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6b, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x3a, 0x2c, 0x90, 0xb5, 0x18, 0x01, 0x88, 0xa6, 0x1d, 0x08, 0x92, 0xa6,
	0x1d, 0x20, 0x08, 0x0a, 0x12, 0x1c, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x5f,
	0x69, 0x64, 0x22, 0x88, 0x04, 0x0a, 0x0e, 0x45, 0x6b, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x7e, 0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x53, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x68, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a,
	0x6f, 0x62, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x62, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x10, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x2a, 0x0a, 0x11, 0x61, 0x77, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x61, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x77, 0x73, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x76,
	0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x70, 0x63,
	0x49, 0x64, 0x12, 0xa8, 0x01, 0x0a, 0x17, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x68, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x65, 0x6b, 0x73, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x65, 0x6b, 0x73,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x45, 0x6b, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x15, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0xe6, 0x02,
	0x0a, 0x10, 0x45, 0x6b, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
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
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x78, 0x0a, 0x0d,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x65, 0x6b, 0x73, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6b, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x1c, 0x45, 0x6b, 0x73, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x76, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x56, 0x70, 0x63, 0x49, 0x64, 0x12, 0x29, 0x0a,
	0x10, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x42, 0xb8, 0x03, 0x0a, 0x43, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67,
	0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x65, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x65, 0x6b, 0x73,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x08,
	0x43, 0x50, 0x41, 0x43, 0x56, 0x41, 0x45, 0x4d, 0xaa, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x77, 0x73, 0x2e,
	0x45, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0xca, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x56, 0x31, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x45, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x41, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x41, 0x77, 0x73, 0x5c,
	0x45, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41,
	0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x41, 0x77, 0x73, 0x3a, 0x3a, 0x45, 0x6b, 0x73, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_goTypes = []interface{}{
	(*EksCluster)(nil),                                     // 0: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksCluster
	(*EksClusterSpec)(nil),                                 // 1: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksClusterSpec
	(*EksClusterStatus)(nil),                               // 2: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksClusterStatus
	(*EksClusterStatusStackOutputs)(nil),                   // 3: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksClusterStatusStackOutputs
	(*model.ApiResourceMetadata)(nil),                      // 4: cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	(*model1.ApiResourceEnvironmentInfo)(nil),              // 5: cloud.planton.apis.resourcemanager.v1.environment.model.ApiResourceEnvironmentInfo
	(*model2.StackJobSettings)(nil),                        // 6: cloud.planton.apis.iac.v1.stackjob.model.StackJobSettings
	(eksworkersmanagementmode.EksWorkersManagementMode)(0), // 7: cloud.planton.apis.code2cloud.v1.aws.ekscluster.enums.eksworkersmanagementmode.EksWorkersManagementMode
	(*model.ApiResourceLifecycle)(nil),                     // 8: cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	(*model.ApiResourceAudit)(nil),                         // 9: cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
}
var file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksCluster.metadata:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	1, // 1: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksCluster.spec:type_name -> cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksClusterSpec
	2, // 2: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksCluster.status:type_name -> cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksClusterStatus
	5, // 3: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksClusterSpec.environment_info:type_name -> cloud.planton.apis.resourcemanager.v1.environment.model.ApiResourceEnvironmentInfo
	6, // 4: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksClusterSpec.stack_job_settings:type_name -> cloud.planton.apis.iac.v1.stackjob.model.StackJobSettings
	7, // 5: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksClusterSpec.workers_management_mode:type_name -> cloud.planton.apis.code2cloud.v1.aws.ekscluster.enums.eksworkersmanagementmode.EksWorkersManagementMode
	8, // 6: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksClusterStatus.lifecycle:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	9, // 7: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksClusterStatus.audit:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	3, // 8: cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksClusterStatus.stack_outputs:type_name -> cloud.planton.apis.code2cloud.v1.aws.ekscluster.model.EksClusterStatusStackOutputs
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EksCluster); i {
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
		file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EksClusterSpec); i {
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
		file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EksClusterStatus); i {
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
		file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EksClusterStatusStackOutputs); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_model_state_proto_depIdxs = nil
}
