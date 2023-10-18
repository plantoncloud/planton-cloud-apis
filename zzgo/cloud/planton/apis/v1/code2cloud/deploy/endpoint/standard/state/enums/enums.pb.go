// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/endpoint/standard/state/enums/enums.proto

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

type StandardEndpointEventType int32

const (
	StandardEndpointEventType_STANDARD_ENDPOINT_EVENT_TYPE_UNSPECIFIED                 StandardEndpointEventType = 0
	StandardEndpointEventType_STANDARD_ENDPOINT_EVENT_TYPE_STATE_CREATED               StandardEndpointEventType = 1
	StandardEndpointEventType_STANDARD_ENDPOINT_EVENT_TYPE_STATE_UPDATED               StandardEndpointEventType = 2
	StandardEndpointEventType_STANDARD_ENDPOINT_EVENT_TYPE_STATE_DELETED               StandardEndpointEventType = 3
	StandardEndpointEventType_STANDARD_ENDPOINT_EVENT_TYPE_STATE_RESTORED              StandardEndpointEventType = 4
	StandardEndpointEventType_STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_PROGRESS_UPDATED  StandardEndpointEventType = 5
	StandardEndpointEventType_STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_PREVIEW_REQUESTED StandardEndpointEventType = 6
	StandardEndpointEventType_STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_APPLY_REQUESTED   StandardEndpointEventType = 7
	StandardEndpointEventType_STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_APPLY_COMPLETED   StandardEndpointEventType = 8
	StandardEndpointEventType_STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_DESTROY_REQUESTED StandardEndpointEventType = 9
	StandardEndpointEventType_STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_DESTROY_COMPLETED StandardEndpointEventType = 10
)

// Enum value maps for StandardEndpointEventType.
var (
	StandardEndpointEventType_name = map[int32]string{
		0:  "STANDARD_ENDPOINT_EVENT_TYPE_UNSPECIFIED",
		1:  "STANDARD_ENDPOINT_EVENT_TYPE_STATE_CREATED",
		2:  "STANDARD_ENDPOINT_EVENT_TYPE_STATE_UPDATED",
		3:  "STANDARD_ENDPOINT_EVENT_TYPE_STATE_DELETED",
		4:  "STANDARD_ENDPOINT_EVENT_TYPE_STATE_RESTORED",
		5:  "STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_PROGRESS_UPDATED",
		6:  "STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_PREVIEW_REQUESTED",
		7:  "STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_APPLY_REQUESTED",
		8:  "STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_APPLY_COMPLETED",
		9:  "STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_DESTROY_REQUESTED",
		10: "STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_DESTROY_COMPLETED",
	}
	StandardEndpointEventType_value = map[string]int32{
		"STANDARD_ENDPOINT_EVENT_TYPE_UNSPECIFIED":                 0,
		"STANDARD_ENDPOINT_EVENT_TYPE_STATE_CREATED":               1,
		"STANDARD_ENDPOINT_EVENT_TYPE_STATE_UPDATED":               2,
		"STANDARD_ENDPOINT_EVENT_TYPE_STATE_DELETED":               3,
		"STANDARD_ENDPOINT_EVENT_TYPE_STATE_RESTORED":              4,
		"STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_PROGRESS_UPDATED":  5,
		"STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_PREVIEW_REQUESTED": 6,
		"STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_APPLY_REQUESTED":   7,
		"STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_APPLY_COMPLETED":   8,
		"STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_DESTROY_REQUESTED": 9,
		"STANDARD_ENDPOINT_EVENT_TYPE_STACK_JOB_DESTROY_COMPLETED": 10,
	}
)

func (x StandardEndpointEventType) Enum() *StandardEndpointEventType {
	p := new(StandardEndpointEventType)
	*p = x
	return p
}

func (x StandardEndpointEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StandardEndpointEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_enumTypes[0].Descriptor()
}

func (StandardEndpointEventType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_enumTypes[0]
}

func (x StandardEndpointEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StandardEndpointEventType.Descriptor instead.
func (StandardEndpointEventType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_rawDesc = []byte{
	0x0a, 0x51, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x5d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x85, 0x06, 0x0a, 0x19, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x28, 0x53, 0x54, 0x41, 0x4e, 0x44,
	0x41, 0x52, 0x44, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x48, 0x0a, 0x2a, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52,
	0x44, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x01, 0x1a, 0x18, 0x80, 0xf9, 0x2b, 0x01, 0x88, 0xf9, 0x2b, 0x01, 0x90,
	0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x03, 0xa8, 0xf9, 0x2b, 0x01, 0x12,
	0x44, 0x0a, 0x2a, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x5f, 0x45, 0x4e, 0x44, 0x50,
	0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x1a,
	0x14, 0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b,
	0x03, 0xa8, 0xf9, 0x2b, 0x01, 0x12, 0x40, 0x0a, 0x2a, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52,
	0x44, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x44, 0x10, 0x03, 0x1a, 0x10, 0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9, 0x2b, 0x01, 0x98,
	0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x04, 0x12, 0x41, 0x0a, 0x2b, 0x53, 0x54, 0x41, 0x4e, 0x44,
	0x41, 0x52, 0x44, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45,
	0x53, 0x54, 0x4f, 0x52, 0x45, 0x44, 0x10, 0x04, 0x1a, 0x10, 0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9,
	0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x03, 0x12, 0x41, 0x0a, 0x37, 0x53, 0x54,
	0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x43, 0x4b,
	0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x05, 0x1a, 0x04, 0x88, 0xf9, 0x2b, 0x01, 0x12, 0x46, 0x0a,
	0x38, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49,
	0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x06, 0x1a, 0x08, 0x98, 0xf9, 0x2b,
	0x01, 0xa0, 0xf9, 0x2b, 0x02, 0x12, 0x48, 0x0a, 0x36, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52,
	0x44, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f,
	0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10,
	0x07, 0x1a, 0x0c, 0x88, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x03, 0x12,
	0x40, 0x0a, 0x36, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x5f, 0x45, 0x4e, 0x44, 0x50,
	0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x08, 0x1a, 0x04, 0x88, 0xf9, 0x2b,
	0x01, 0x12, 0x4a, 0x0a, 0x38, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x5f, 0x45, 0x4e,
	0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x44, 0x45, 0x53, 0x54,
	0x52, 0x4f, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x09, 0x1a,
	0x0c, 0x88, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x04, 0x12, 0x42, 0x0a,
	0x38, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49,
	0x4e, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x52, 0x4f, 0x59, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x1a, 0x04, 0x88, 0xf9, 0x2b,
	0x01, 0x42, 0x9c, 0x04, 0x0a, 0x53, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x75, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02,
	0x0a, 0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x45, 0x53, 0x53, 0x45, 0xaa, 0x02, 0x45, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73,
	0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xca, 0x02, 0x45, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5c, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5c,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xe2, 0x02, 0x51, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5c, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x5c, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_goTypes = []interface{}{
	(StandardEndpointEventType)(0), // 0: cloud.planton.apis.v1.code2cloud.deploy.endpoint.standard.state.enums.StandardEndpointEventType
}
var file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_init()
}
func file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_standard_state_enums_enums_proto_depIdxs = nil
}
