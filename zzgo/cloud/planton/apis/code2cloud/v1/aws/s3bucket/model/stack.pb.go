// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/aws/s3bucket/model/stack.proto

package model

import (
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/connect/v1/awscredential/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model"
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

// input for s3-bucket stack
type S3BucketStackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack-job
	StackJob *model.StackJob `protobuf:"bytes,1,opt,name=stack_job,json=stackJob,proto3" json:"stack_job,omitempty"`
	// api-resource
	ApiResource *S3Bucket `protobuf:"bytes,2,opt,name=api_resource,json=apiResource,proto3" json:"api_resource,omitempty"`
	// aws-credential
	AwsCredential *model1.AwsCredential `protobuf:"bytes,3,opt,name=aws_credential,json=awsCredential,proto3" json:"aws_credential,omitempty"`
}

func (x *S3BucketStackInput) Reset() {
	*x = S3BucketStackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3BucketStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3BucketStackInput) ProtoMessage() {}

func (x *S3BucketStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3BucketStackInput.ProtoReflect.Descriptor instead.
func (*S3BucketStackInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDescGZIP(), []int{0}
}

func (x *S3BucketStackInput) GetStackJob() *model.StackJob {
	if x != nil {
		return x.StackJob
	}
	return nil
}

func (x *S3BucketStackInput) GetApiResource() *S3Bucket {
	if x != nil {
		return x.ApiResource
	}
	return nil
}

func (x *S3BucketStackInput) GetAwsCredential() *model1.AwsCredential {
	if x != nil {
		return x.AwsCredential
	}
	return nil
}

// stack response
type S3BucketStackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack-job progress-event
	ProgressEvent *progress.StackJobProgressEvent `protobuf:"bytes,1,opt,name=progress_event,json=progressEvent,proto3" json:"progress_event,omitempty"`
	// stack outputs
	StackOutputs *S3BucketStackOutputs `protobuf:"bytes,2,opt,name=stack_outputs,json=stackOutputs,proto3" json:"stack_outputs,omitempty"`
}

func (x *S3BucketStackResponse) Reset() {
	*x = S3BucketStackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3BucketStackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3BucketStackResponse) ProtoMessage() {}

func (x *S3BucketStackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3BucketStackResponse.ProtoReflect.Descriptor instead.
func (*S3BucketStackResponse) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDescGZIP(), []int{1}
}

func (x *S3BucketStackResponse) GetProgressEvent() *progress.StackJobProgressEvent {
	if x != nil {
		return x.ProgressEvent
	}
	return nil
}

func (x *S3BucketStackResponse) GetStackOutputs() *S3BucketStackOutputs {
	if x != nil {
		return x.StackOutputs
	}
	return nil
}

var File_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x73, 0x33, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x33, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x73, 0x33, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x3d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x73, 0x33, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x02, 0x0a, 0x12, 0x53, 0x33, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x4f, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62,
	0x12, 0x60, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x73, 0x33,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x33, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0b, 0x61, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x67, 0x0a, 0x0e, 0x61, 0x77, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x77, 0x73, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x0d, 0x61, 0x77,
	0x73, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0xf8, 0x01, 0x0a, 0x15,
	0x53, 0x33, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a,
	0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x6e, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x61, 0x77, 0x73, 0x2e, 0x73, 0x33, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x53, 0x33, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x42, 0xac, 0x03, 0x0a, 0x41, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x73, 0x33,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73,
	0x2f, 0x73, 0x33, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2,
	0x02, 0x08, 0x43, 0x50, 0x41, 0x43, 0x56, 0x41, 0x53, 0x4d, 0xaa, 0x02, 0x33, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x77,
	0x73, 0x2e, 0x53, 0x33, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0xca, 0x02, 0x33, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x56, 0x31, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x53, 0x33, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x3f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x53, 0x33,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a,
	0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x41, 0x77, 0x73, 0x3a, 0x3a, 0x53, 0x33, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x3a,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_goTypes = []interface{}{
	(*S3BucketStackInput)(nil),             // 0: cloud.planton.apis.code2cloud.v1.aws.s3bucket.model.S3BucketStackInput
	(*S3BucketStackResponse)(nil),          // 1: cloud.planton.apis.code2cloud.v1.aws.s3bucket.model.S3BucketStackResponse
	(*model.StackJob)(nil),                 // 2: cloud.planton.apis.iac.v1.stackjob.model.StackJob
	(*S3Bucket)(nil),                       // 3: cloud.planton.apis.code2cloud.v1.aws.s3bucket.model.S3Bucket
	(*model1.AwsCredential)(nil),           // 4: cloud.planton.apis.connect.v1.awscredential.model.AwsCredential
	(*progress.StackJobProgressEvent)(nil), // 5: cloud.planton.apis.iac.v1.stackjob.model.progress.StackJobProgressEvent
	(*S3BucketStackOutputs)(nil),           // 6: cloud.planton.apis.code2cloud.v1.aws.s3bucket.model.S3BucketStackOutputs
}
var file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_depIdxs = []int32{
	2, // 0: cloud.planton.apis.code2cloud.v1.aws.s3bucket.model.S3BucketStackInput.stack_job:type_name -> cloud.planton.apis.iac.v1.stackjob.model.StackJob
	3, // 1: cloud.planton.apis.code2cloud.v1.aws.s3bucket.model.S3BucketStackInput.api_resource:type_name -> cloud.planton.apis.code2cloud.v1.aws.s3bucket.model.S3Bucket
	4, // 2: cloud.planton.apis.code2cloud.v1.aws.s3bucket.model.S3BucketStackInput.aws_credential:type_name -> cloud.planton.apis.connect.v1.awscredential.model.AwsCredential
	5, // 3: cloud.planton.apis.code2cloud.v1.aws.s3bucket.model.S3BucketStackResponse.progress_event:type_name -> cloud.planton.apis.iac.v1.stackjob.model.progress.StackJobProgressEvent
	6, // 4: cloud.planton.apis.code2cloud.v1.aws.s3bucket.model.S3BucketStackResponse.stack_outputs:type_name -> cloud.planton.apis.code2cloud.v1.aws.s3bucket.model.S3BucketStackOutputs
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto != nil {
		return
	}
	file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3BucketStackInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3BucketStackResponse); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_aws_s3bucket_model_stack_proto_depIdxs = nil
}
