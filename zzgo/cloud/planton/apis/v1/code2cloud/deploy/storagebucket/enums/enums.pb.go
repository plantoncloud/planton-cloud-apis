// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/storagebucket/enums/enums.proto

package enums

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/state/enums/options"
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

type StorageBucketEventType int32

const (
	StorageBucketEventType_STORAGE_BUCKET_EVENT_TYPE_UNSPECIFIED                 StorageBucketEventType = 0
	StorageBucketEventType_STORAGE_BUCKET_EVENT_TYPE_STATE_CREATED               StorageBucketEventType = 1
	StorageBucketEventType_STORAGE_BUCKET_EVENT_TYPE_STATE_UPDATED               StorageBucketEventType = 2
	StorageBucketEventType_STORAGE_BUCKET_EVENT_TYPE_STATE_DELETED               StorageBucketEventType = 3
	StorageBucketEventType_STORAGE_BUCKET_EVENT_TYPE_STATE_RESTORED              StorageBucketEventType = 4
	StorageBucketEventType_STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_PROGRESS_UPDATED  StorageBucketEventType = 5
	StorageBucketEventType_STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_PREVIEW_REQUESTED StorageBucketEventType = 6
	StorageBucketEventType_STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_APPLY_REQUESTED   StorageBucketEventType = 7
	StorageBucketEventType_STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_APPLY_COMPLETED   StorageBucketEventType = 8
	StorageBucketEventType_STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_DESTROY_REQUESTED StorageBucketEventType = 9
	StorageBucketEventType_STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_DESTROY_COMPLETED StorageBucketEventType = 10
)

// Enum value maps for StorageBucketEventType.
var (
	StorageBucketEventType_name = map[int32]string{
		0:  "STORAGE_BUCKET_EVENT_TYPE_UNSPECIFIED",
		1:  "STORAGE_BUCKET_EVENT_TYPE_STATE_CREATED",
		2:  "STORAGE_BUCKET_EVENT_TYPE_STATE_UPDATED",
		3:  "STORAGE_BUCKET_EVENT_TYPE_STATE_DELETED",
		4:  "STORAGE_BUCKET_EVENT_TYPE_STATE_RESTORED",
		5:  "STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_PROGRESS_UPDATED",
		6:  "STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_PREVIEW_REQUESTED",
		7:  "STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_APPLY_REQUESTED",
		8:  "STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_APPLY_COMPLETED",
		9:  "STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_DESTROY_REQUESTED",
		10: "STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_DESTROY_COMPLETED",
	}
	StorageBucketEventType_value = map[string]int32{
		"STORAGE_BUCKET_EVENT_TYPE_UNSPECIFIED":                 0,
		"STORAGE_BUCKET_EVENT_TYPE_STATE_CREATED":               1,
		"STORAGE_BUCKET_EVENT_TYPE_STATE_UPDATED":               2,
		"STORAGE_BUCKET_EVENT_TYPE_STATE_DELETED":               3,
		"STORAGE_BUCKET_EVENT_TYPE_STATE_RESTORED":              4,
		"STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_PROGRESS_UPDATED":  5,
		"STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_PREVIEW_REQUESTED": 6,
		"STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_APPLY_REQUESTED":   7,
		"STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_APPLY_COMPLETED":   8,
		"STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_DESTROY_REQUESTED": 9,
		"STORAGE_BUCKET_EVENT_TYPE_STACK_JOB_DESTROY_COMPLETED": 10,
	}
)

func (x StorageBucketEventType) Enum() *StorageBucketEventType {
	p := new(StorageBucketEventType)
	*p = x
	return p
}

func (x StorageBucketEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageBucketEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_enumTypes[0].Descriptor()
}

func (StorageBucketEventType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_enumTypes[0]
}

func (x StorageBucketEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageBucketEventType.Descriptor instead.
func (StorageBucketEventType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDescGZIP(), []int{0}
}

// provider for storage-bucket
type StorageBucketProvider int32

const (
	StorageBucketProvider_STORAGE_BUCKET_PROVIDER_UNSPECIFIED StorageBucketProvider = 0
	// https://cloud.google.com/storage
	StorageBucketProvider_GCP_GCS StorageBucketProvider = 1
	// https://aws.amazon.com/s3/
	StorageBucketProvider_AWS_S3 StorageBucketProvider = 2
	// https://azure.microsoft.com/en-us/products/storage/blobs
	StorageBucketProvider_AZURE_BLOB_STORAGE StorageBucketProvider = 3
	// https://www.digitalocean.com/products/spaces
	StorageBucketProvider_DIGITAL_OCEAN_OBJECT_STORAGE StorageBucketProvider = 4
)

// Enum value maps for StorageBucketProvider.
var (
	StorageBucketProvider_name = map[int32]string{
		0: "STORAGE_BUCKET_PROVIDER_UNSPECIFIED",
		1: "GCP_GCS",
		2: "AWS_S3",
		3: "AZURE_BLOB_STORAGE",
		4: "DIGITAL_OCEAN_OBJECT_STORAGE",
	}
	StorageBucketProvider_value = map[string]int32{
		"STORAGE_BUCKET_PROVIDER_UNSPECIFIED": 0,
		"GCP_GCS":                             1,
		"AWS_S3":                              2,
		"AZURE_BLOB_STORAGE":                  3,
		"DIGITAL_OCEAN_OBJECT_STORAGE":        4,
	}
)

func (x StorageBucketProvider) Enum() *StorageBucketProvider {
	p := new(StorageBucketProvider)
	*p = x
	return p
}

func (x StorageBucketProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageBucketProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_enumTypes[1].Descriptor()
}

func (StorageBucketProvider) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_enumTypes[1]
}

func (x StorageBucketProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageBucketProvider.Descriptor instead.
func (StorageBucketProvider) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDescGZIP(), []int{1}
}

var File_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDesc = []byte{
	0x0a, 0x47, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x5d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xe1, 0x05, 0x0a, 0x16, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x29, 0x0a, 0x25, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x55, 0x43, 0x4b,
	0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x45, 0x0a, 0x27, 0x53,
	0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x1a, 0x18, 0x80, 0xf9, 0x2b, 0x01, 0x88, 0xf9,
	0x2b, 0x01, 0x90, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x03, 0xa8, 0xf9,
	0x2b, 0x01, 0x12, 0x41, 0x0a, 0x27, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x55,
	0x43, 0x4b, 0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x1a,
	0x14, 0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b,
	0x03, 0xa8, 0xf9, 0x2b, 0x01, 0x12, 0x3d, 0x0a, 0x27, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45,
	0x5f, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44,
	0x10, 0x03, 0x1a, 0x10, 0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01,
	0xa0, 0xf9, 0x2b, 0x04, 0x12, 0x3e, 0x0a, 0x28, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f,
	0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x44,
	0x10, 0x04, 0x1a, 0x10, 0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01,
	0xa0, 0xf9, 0x2b, 0x03, 0x12, 0x3e, 0x0a, 0x34, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f,
	0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x50, 0x52, 0x4f, 0x47,
	0x52, 0x45, 0x53, 0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x05, 0x1a, 0x04,
	0x88, 0xf9, 0x2b, 0x01, 0x12, 0x43, 0x0a, 0x35, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f,
	0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x50, 0x52, 0x45, 0x56,
	0x49, 0x45, 0x57, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x06, 0x1a,
	0x08, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x02, 0x12, 0x45, 0x0a, 0x33, 0x53, 0x54, 0x4f,
	0x52, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42,
	0x5f, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44,
	0x10, 0x07, 0x1a, 0x0c, 0x88, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x03,
	0x12, 0x3d, 0x0a, 0x33, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x55, 0x43, 0x4b,
	0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x08, 0x1a, 0x04, 0x88, 0xf9, 0x2b, 0x01, 0x12,
	0x47, 0x0a, 0x35, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x55, 0x43, 0x4b, 0x45,
	0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x52, 0x4f, 0x59, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x09, 0x1a, 0x0c, 0x88, 0xf9, 0x2b, 0x01,
	0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x04, 0x12, 0x3f, 0x0a, 0x35, 0x53, 0x54, 0x4f, 0x52,
	0x41, 0x47, 0x45, 0x5f, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f,
	0x44, 0x45, 0x53, 0x54, 0x52, 0x4f, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x0a, 0x1a, 0x04, 0x88, 0xf9, 0x2b, 0x01, 0x2a, 0x93, 0x01, 0x0a, 0x15, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x23, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x42,
	0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x47, 0x43, 0x50, 0x5f, 0x47, 0x43, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x57, 0x53,
	0x5f, 0x53, 0x33, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x5a, 0x55, 0x52, 0x45, 0x5f, 0x42,
	0x4c, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10, 0x03, 0x12, 0x20, 0x0a,
	0x1c, 0x44, 0x49, 0x47, 0x49, 0x54, 0x41, 0x4c, 0x5f, 0x4f, 0x43, 0x45, 0x41, 0x4e, 0x5f, 0x4f,
	0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10, 0x04, 0x42,
	0xdc, 0x03, 0x0a, 0x49, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65,
	0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x0a, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x43,
	0x44, 0x53, 0x45, 0xaa, 0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xca, 0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xe2,
	0x02, 0x47, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x42, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_goTypes = []interface{}{
	(StorageBucketEventType)(0), // 0: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.enums.StorageBucketEventType
	(StorageBucketProvider)(0),  // 1: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.enums.StorageBucketProvider
}
var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_enums_enums_proto_depIdxs = nil
}
