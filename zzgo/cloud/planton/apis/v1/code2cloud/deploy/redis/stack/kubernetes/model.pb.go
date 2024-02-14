// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/redis/stack/kubernetes/model.proto

package kubernetes

import (
	redis "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/redis"
	operation "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/pulumi/operation"
	job "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job"
	progress "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job/progress"
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

// input for redis-cluster stack
type RedisClusterKubernetesStackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack job
	StackJob *job.StackJob `protobuf:"bytes,1,opt,name=stack_job,json=stackJob,proto3" json:"stack_job,omitempty"`
	// pulumi stack credentials
	CredentialsInput *RedisClusterKubernetesStackCredentialsInput `protobuf:"bytes,2,opt,name=credentials_input,json=credentialsInput,proto3" json:"credentials_input,omitempty"`
	// inputs used for creating stack resources
	ResourceInput *RedisClusterKubernetesStackResourceInput `protobuf:"bytes,3,opt,name=resource_input,json=resourceInput,proto3" json:"resource_input,omitempty"`
}

func (x *RedisClusterKubernetesStackInput) Reset() {
	*x = RedisClusterKubernetesStackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisClusterKubernetesStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisClusterKubernetesStackInput) ProtoMessage() {}

func (x *RedisClusterKubernetesStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisClusterKubernetesStackInput.ProtoReflect.Descriptor instead.
func (*RedisClusterKubernetesStackInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDescGZIP(), []int{0}
}

func (x *RedisClusterKubernetesStackInput) GetStackJob() *job.StackJob {
	if x != nil {
		return x.StackJob
	}
	return nil
}

func (x *RedisClusterKubernetesStackInput) GetCredentialsInput() *RedisClusterKubernetesStackCredentialsInput {
	if x != nil {
		return x.CredentialsInput
	}
	return nil
}

func (x *RedisClusterKubernetesStackInput) GetResourceInput() *RedisClusterKubernetesStackResourceInput {
	if x != nil {
		return x.ResourceInput
	}
	return nil
}

// stack credentials input
type RedisClusterKubernetesStackCredentialsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// kubernetes provider credential for creating redis-cluster resources on kubernetes cluster
	Kubernetes *operation.KubernetesProviderCredential `protobuf:"bytes,1,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
}

func (x *RedisClusterKubernetesStackCredentialsInput) Reset() {
	*x = RedisClusterKubernetesStackCredentialsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisClusterKubernetesStackCredentialsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisClusterKubernetesStackCredentialsInput) ProtoMessage() {}

func (x *RedisClusterKubernetesStackCredentialsInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisClusterKubernetesStackCredentialsInput.ProtoReflect.Descriptor instead.
func (*RedisClusterKubernetesStackCredentialsInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDescGZIP(), []int{1}
}

func (x *RedisClusterKubernetesStackCredentialsInput) GetKubernetes() *operation.KubernetesProviderCredential {
	if x != nil {
		return x.Kubernetes
	}
	return nil
}

// stack resource input
type RedisClusterKubernetesStackResourceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// redis-cluster
	RedisCluster *redis.RedisCluster `protobuf:"bytes,1,opt,name=redis_cluster,json=redisCluster,proto3" json:"redis_cluster,omitempty"`
}

func (x *RedisClusterKubernetesStackResourceInput) Reset() {
	*x = RedisClusterKubernetesStackResourceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisClusterKubernetesStackResourceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisClusterKubernetesStackResourceInput) ProtoMessage() {}

func (x *RedisClusterKubernetesStackResourceInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisClusterKubernetesStackResourceInput.ProtoReflect.Descriptor instead.
func (*RedisClusterKubernetesStackResourceInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDescGZIP(), []int{2}
}

func (x *RedisClusterKubernetesStackResourceInput) GetRedisCluster() *redis.RedisCluster {
	if x != nil {
		return x.RedisCluster
	}
	return nil
}

// redis-cluster stack outputs
type RedisClusterKubernetesStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// status of the redis-cluster upon stack-job
	RedisClusterStatus *redis.RedisClusterStatus `protobuf:"bytes,1,opt,name=redis_cluster_status,json=redisClusterStatus,proto3" json:"redis_cluster_status,omitempty"`
}

func (x *RedisClusterKubernetesStackOutputs) Reset() {
	*x = RedisClusterKubernetesStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisClusterKubernetesStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisClusterKubernetesStackOutputs) ProtoMessage() {}

func (x *RedisClusterKubernetesStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisClusterKubernetesStackOutputs.ProtoReflect.Descriptor instead.
func (*RedisClusterKubernetesStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDescGZIP(), []int{3}
}

func (x *RedisClusterKubernetesStackOutputs) GetRedisClusterStatus() *redis.RedisClusterStatus {
	if x != nil {
		return x.RedisClusterStatus
	}
	return nil
}

// stack response
type RedisClusterKubernetesStackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack-job progress event
	ProgressEvent *progress.StackJobProgressEvent `protobuf:"bytes,1,opt,name=progress_event,json=progressEvent,proto3" json:"progress_event,omitempty"`
	// stack outputs
	Outputs *RedisClusterKubernetesStackOutputs `protobuf:"bytes,3,opt,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *RedisClusterKubernetesStackResponse) Reset() {
	*x = RedisClusterKubernetesStackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisClusterKubernetesStackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisClusterKubernetesStackResponse) ProtoMessage() {}

func (x *RedisClusterKubernetesStackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisClusterKubernetesStackResponse.ProtoReflect.Descriptor instead.
func (*RedisClusterKubernetesStackResponse) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDescGZIP(), []int{4}
}

func (x *RedisClusterKubernetesStackResponse) GetProgressEvent() *progress.StackJobProgressEvent {
	if x != nil {
		return x.ProgressEvent
	}
	return nil
}

func (x *RedisClusterKubernetesStackResponse) GetOutputs() *RedisClusterKubernetesStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x1a, 0x39, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x03, 0x0a, 0x20, 0x52, 0x65, 0x64, 0x69, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x46, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x4a, 0x6f, 0x62, 0x12, 0x98, 0x01, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x6b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x10, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x8f,
	0x01, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x68, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x22, 0x9b, 0x01, 0x0a, 0x2b, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x6c, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0x8c,
	0x01, 0x0a, 0x28, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x60, 0x0a, 0x0d, 0x72,
	0x65, 0x64, 0x69, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x0c, 0x72, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x99, 0x01,
	0x0a, 0x22, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x12, 0x73, 0x0a, 0x14, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x12, 0x72, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8b, 0x02, 0x0a, 0x23, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x66, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x7c, 0x0a, 0x07, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x62, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69,
	0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x07,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x42, 0xf0, 0x03, 0x0a, 0x4c, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x52,
	0x53, 0x4b, 0xaa, 0x02, 0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x52, 0x65, 0x64,
	0x69, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0xca, 0x02, 0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0xe2, 0x02, 0x4a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x52,
	0x65, 0x64, 0x69, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x46, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x3a, 0x3a, 0x52, 0x65, 0x64, 0x69, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_goTypes = []interface{}{
	(*RedisClusterKubernetesStackInput)(nil),            // 0: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackInput
	(*RedisClusterKubernetesStackCredentialsInput)(nil), // 1: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackCredentialsInput
	(*RedisClusterKubernetesStackResourceInput)(nil),    // 2: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackResourceInput
	(*RedisClusterKubernetesStackOutputs)(nil),          // 3: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackOutputs
	(*RedisClusterKubernetesStackResponse)(nil),         // 4: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackResponse
	(*job.StackJob)(nil),                                // 5: cloud.planton.apis.v1.stack.job.StackJob
	(*operation.KubernetesProviderCredential)(nil),      // 6: cloud.planton.apis.v1.commons.pulumi.operation.KubernetesProviderCredential
	(*redis.RedisCluster)(nil),                          // 7: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	(*redis.RedisClusterStatus)(nil),                    // 8: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterStatus
	(*progress.StackJobProgressEvent)(nil),              // 9: cloud.planton.apis.v1.stack.job.progress.StackJobProgressEvent
}
var file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackInput.stack_job:type_name -> cloud.planton.apis.v1.stack.job.StackJob
	1, // 1: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackInput.credentials_input:type_name -> cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackCredentialsInput
	2, // 2: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackInput.resource_input:type_name -> cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackResourceInput
	6, // 3: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackCredentialsInput.kubernetes:type_name -> cloud.planton.apis.v1.commons.pulumi.operation.KubernetesProviderCredential
	7, // 4: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackResourceInput.redis_cluster:type_name -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	8, // 5: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackOutputs.redis_cluster_status:type_name -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterStatus
	9, // 6: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackResponse.progress_event:type_name -> cloud.planton.apis.v1.stack.job.progress.StackJobProgressEvent
	3, // 7: cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackResponse.outputs:type_name -> cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes.RedisClusterKubernetesStackOutputs
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisClusterKubernetesStackInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisClusterKubernetesStackCredentialsInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisClusterKubernetesStackResourceInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisClusterKubernetesStackOutputs); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisClusterKubernetesStackResponse); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_redis_stack_kubernetes_model_proto_depIdxs = nil
}
