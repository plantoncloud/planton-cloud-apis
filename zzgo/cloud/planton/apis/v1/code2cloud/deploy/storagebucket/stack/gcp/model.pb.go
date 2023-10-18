// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/storagebucket/stack/gcp/model.proto

package gcp

import (
	state "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/storagebucket/state"
	rpc1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/pulumi/operation/rpc"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/rpc"
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

// input for storage-bucket stack
type StorageBucketGcpStackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack job
	StackJob *rpc.StackJob `protobuf:"bytes,1,opt,name=stack_job,json=stackJob,proto3" json:"stack_job,omitempty"`
	// pulumi stack credentials
	CredentialsInput *StorageBucketGcpStackCredentialsInput `protobuf:"bytes,2,opt,name=credentials_input,json=credentialsInput,proto3" json:"credentials_input,omitempty"`
	// inputs used for creating stack resources
	ResourceInput *StorageBucketGcpStackResourceInput `protobuf:"bytes,3,opt,name=resource_input,json=resourceInput,proto3" json:"resource_input,omitempty"`
}

func (x *StorageBucketGcpStackInput) Reset() {
	*x = StorageBucketGcpStackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageBucketGcpStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucketGcpStackInput) ProtoMessage() {}

func (x *StorageBucketGcpStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucketGcpStackInput.ProtoReflect.Descriptor instead.
func (*StorageBucketGcpStackInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDescGZIP(), []int{0}
}

func (x *StorageBucketGcpStackInput) GetStackJob() *rpc.StackJob {
	if x != nil {
		return x.StackJob
	}
	return nil
}

func (x *StorageBucketGcpStackInput) GetCredentialsInput() *StorageBucketGcpStackCredentialsInput {
	if x != nil {
		return x.CredentialsInput
	}
	return nil
}

func (x *StorageBucketGcpStackInput) GetResourceInput() *StorageBucketGcpStackResourceInput {
	if x != nil {
		return x.ResourceInput
	}
	return nil
}

// stack credentials input
type StorageBucketGcpStackCredentialsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// google provider credential input
	Google *rpc1.GoogleProviderCredential `protobuf:"bytes,1,opt,name=google,proto3" json:"google,omitempty"`
}

func (x *StorageBucketGcpStackCredentialsInput) Reset() {
	*x = StorageBucketGcpStackCredentialsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageBucketGcpStackCredentialsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucketGcpStackCredentialsInput) ProtoMessage() {}

func (x *StorageBucketGcpStackCredentialsInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucketGcpStackCredentialsInput.ProtoReflect.Descriptor instead.
func (*StorageBucketGcpStackCredentialsInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDescGZIP(), []int{1}
}

func (x *StorageBucketGcpStackCredentialsInput) GetGoogle() *rpc1.GoogleProviderCredential {
	if x != nil {
		return x.Google
	}
	return nil
}

// stack resource input
type StorageBucketGcpStackResourceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// storage-bucket
	StorageBucket *state.StorageBucketState `protobuf:"bytes,1,opt,name=storage_bucket,json=storageBucket,proto3" json:"storage_bucket,omitempty"`
}

func (x *StorageBucketGcpStackResourceInput) Reset() {
	*x = StorageBucketGcpStackResourceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageBucketGcpStackResourceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucketGcpStackResourceInput) ProtoMessage() {}

func (x *StorageBucketGcpStackResourceInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucketGcpStackResourceInput.ProtoReflect.Descriptor instead.
func (*StorageBucketGcpStackResourceInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDescGZIP(), []int{2}
}

func (x *StorageBucketGcpStackResourceInput) GetStorageBucket() *state.StorageBucketState {
	if x != nil {
		return x.StorageBucket
	}
	return nil
}

// storage-bucket stack outputs
type StorageBucketGcpStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StorageBucketGcpStackOutputs) Reset() {
	*x = StorageBucketGcpStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageBucketGcpStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucketGcpStackOutputs) ProtoMessage() {}

func (x *StorageBucketGcpStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucketGcpStackOutputs.ProtoReflect.Descriptor instead.
func (*StorageBucketGcpStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDescGZIP(), []int{3}
}

// stack response
type StorageBucketGcpStackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack progress
	Progress *rpc.StackJobProgressEvent `protobuf:"bytes,1,opt,name=progress,proto3" json:"progress,omitempty"`
	// stack outputs
	Outputs *StorageBucketGcpStackOutputs `protobuf:"bytes,3,opt,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *StorageBucketGcpStackResponse) Reset() {
	*x = StorageBucketGcpStackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageBucketGcpStackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucketGcpStackResponse) ProtoMessage() {}

func (x *StorageBucketGcpStackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucketGcpStackResponse.ProtoReflect.Descriptor instead.
func (*StorageBucketGcpStackResponse) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDescGZIP(), []int{4}
}

func (x *StorageBucketGcpStackResponse) GetProgress() *rpc.StackJobProgressEvent {
	if x != nil {
		return x.Progress
	}
	return nil
}

func (x *StorageBucketGcpStackResponse) GetOutputs() *StorageBucketGcpStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x67, 0x63,
	0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x67, 0x63, 0x70, 0x1a, 0x3e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x75,
	0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x47, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x03, 0x0a, 0x1a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x46, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f,
	0x62, 0x52, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x93, 0x01, 0x0a, 0x11,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52,
	0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x8a, 0x01, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x63, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x8d,
	0x01, 0x0a, 0x25, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x64, 0x0a, 0x06, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x22, 0x9c,
	0x01, 0x0a, 0x22, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x76, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0d,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x1e, 0x0a,
	0x1c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x47, 0x63,
	0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22, 0xec, 0x01,
	0x0a, 0x1d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x47,
	0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x52, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x77, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x5d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x42, 0xf6, 0x03, 0x0a,
	0x4d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x67, 0x63, 0x70, 0x42, 0x0a,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x67, 0x63, 0x70, 0xa2, 0x02, 0x09,
	0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x53, 0x53, 0x47, 0xaa, 0x02, 0x3f, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56,
	0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x47, 0x63, 0x70, 0xca, 0x02, 0x3f, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x47, 0x63, 0x70, 0xe2, 0x02, 0x4b,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x47, 0x63, 0x70, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x47, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x3a, 0x3a, 0x47, 0x63, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_goTypes = []interface{}{
	(*StorageBucketGcpStackInput)(nil),            // 0: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackInput
	(*StorageBucketGcpStackCredentialsInput)(nil), // 1: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackCredentialsInput
	(*StorageBucketGcpStackResourceInput)(nil),    // 2: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackResourceInput
	(*StorageBucketGcpStackOutputs)(nil),          // 3: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackOutputs
	(*StorageBucketGcpStackResponse)(nil),         // 4: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackResponse
	(*rpc.StackJob)(nil),                          // 5: cloud.planton.apis.v1.stack.rpc.StackJob
	(*rpc1.GoogleProviderCredential)(nil),         // 6: cloud.planton.apis.v1.commons.pulumi.operation.rpc.GoogleProviderCredential
	(*state.StorageBucketState)(nil),              // 7: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.state.StorageBucketState
	(*rpc.StackJobProgressEvent)(nil),             // 8: cloud.planton.apis.v1.stack.rpc.StackJobProgressEvent
}
var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackInput.stack_job:type_name -> cloud.planton.apis.v1.stack.rpc.StackJob
	1, // 1: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackInput.credentials_input:type_name -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackCredentialsInput
	2, // 2: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackInput.resource_input:type_name -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackResourceInput
	6, // 3: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackCredentialsInput.google:type_name -> cloud.planton.apis.v1.commons.pulumi.operation.rpc.GoogleProviderCredential
	7, // 4: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackResourceInput.storage_bucket:type_name -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.state.StorageBucketState
	8, // 5: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackResponse.progress:type_name -> cloud.planton.apis.v1.stack.rpc.StackJobProgressEvent
	3, // 6: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackResponse.outputs:type_name -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp.StorageBucketGcpStackOutputs
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageBucketGcpStackInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageBucketGcpStackCredentialsInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageBucketGcpStackResourceInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageBucketGcpStackOutputs); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageBucketGcpStackResponse); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_stack_gcp_model_proto_depIdxs = nil
}
