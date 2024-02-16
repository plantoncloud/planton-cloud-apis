// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/commons/resource/enums/event/event.proto

package event

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/enums/options"
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

// enumeration of the union of all event types used across all the api resources.
// this is used for specifying the event type produced by rpc methods in command controllers.
// using different enum for these event types while specifying in rpc methods and individual api resource would
// work because the events are eventually converted to strings during message passing.
type ApiResourceEventType int32

const (
	ApiResourceEventType_API_RESOURCE_EVENT_TYPE_UNSPECIFIED ApiResourceEventType = 0
	ApiResourceEventType_created                             ApiResourceEventType = 1
	ApiResourceEventType_updated                             ApiResourceEventType = 2
	ApiResourceEventType_deleted                             ApiResourceEventType = 3
	ApiResourceEventType_restored                            ApiResourceEventType = 4
	ApiResourceEventType_refreshed                           ApiResourceEventType = 5
	ApiResourceEventType_paused                              ApiResourceEventType = 6
	ApiResourceEventType_unpaused                            ApiResourceEventType = 7
	ApiResourceEventType_stack_job_progress_updated          ApiResourceEventType = 8
	ApiResourceEventType_stack_job_refresh_requested         ApiResourceEventType = 9
	ApiResourceEventType_stack_job_refresh_completed         ApiResourceEventType = 10
	ApiResourceEventType_stack_job_apply_preview_requested   ApiResourceEventType = 11
	ApiResourceEventType_stack_job_apply_requested           ApiResourceEventType = 12
	ApiResourceEventType_stack_job_apply_completed           ApiResourceEventType = 13
	ApiResourceEventType_stack_job_destroy_preview_requested ApiResourceEventType = 14
	ApiResourceEventType_stack_job_destroy_requested         ApiResourceEventType = 15
	ApiResourceEventType_stack_job_destroy_completed         ApiResourceEventType = 16
)

// Enum value maps for ApiResourceEventType.
var (
	ApiResourceEventType_name = map[int32]string{
		0:  "API_RESOURCE_EVENT_TYPE_UNSPECIFIED",
		1:  "created",
		2:  "updated",
		3:  "deleted",
		4:  "restored",
		5:  "refreshed",
		6:  "paused",
		7:  "unpaused",
		8:  "stack_job_progress_updated",
		9:  "stack_job_refresh_requested",
		10: "stack_job_refresh_completed",
		11: "stack_job_apply_preview_requested",
		12: "stack_job_apply_requested",
		13: "stack_job_apply_completed",
		14: "stack_job_destroy_preview_requested",
		15: "stack_job_destroy_requested",
		16: "stack_job_destroy_completed",
	}
	ApiResourceEventType_value = map[string]int32{
		"API_RESOURCE_EVENT_TYPE_UNSPECIFIED": 0,
		"created":                             1,
		"updated":                             2,
		"deleted":                             3,
		"restored":                            4,
		"refreshed":                           5,
		"paused":                              6,
		"unpaused":                            7,
		"stack_job_progress_updated":          8,
		"stack_job_refresh_requested":         9,
		"stack_job_refresh_completed":         10,
		"stack_job_apply_preview_requested":   11,
		"stack_job_apply_requested":           12,
		"stack_job_apply_completed":           13,
		"stack_job_destroy_preview_requested": 14,
		"stack_job_destroy_requested":         15,
		"stack_job_destroy_completed":         16,
	}
)

func (x ApiResourceEventType) Enum() *ApiResourceEventType {
	p := new(ApiResourceEventType)
	*p = x
	return p
}

func (x ApiResourceEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiResourceEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_enumTypes[0].Descriptor()
}

func (ApiResourceEventType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_enumTypes[0]
}

func (x ApiResourceEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiResourceEventType.Descriptor instead.
func (ApiResourceEventType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_v1_commons_resource_enums_event_event_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x5b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0xcf, 0x03, 0x0a, 0x14, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x23, 0x41, 0x50,
	0x49, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x65, 0x64, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65,
	0x64, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x75, 0x6e, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x10,
	0x07, 0x12, 0x1e, 0x0a, 0x1a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x10,
	0x08, 0x12, 0x1f, 0x0a, 0x1b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x10, 0x09, 0x12, 0x1f, 0x0a, 0x1b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x10, 0x0a, 0x12, 0x25, 0x0a, 0x21, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62,
	0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x10, 0x0b, 0x12, 0x1d, 0x0a, 0x19, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x10, 0x0c, 0x12, 0x1d, 0x0a, 0x19, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x0d, 0x12, 0x27, 0x0a, 0x23, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x10,
	0x0e, 0x12, 0x1f, 0x0a, 0x1b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x64,
	0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x10, 0x0f, 0x12, 0x1f, 0x0a, 0x1b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f,
	0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x10, 0x10, 0x42, 0xa6, 0x03, 0x0a, 0x40, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41,
	0x56, 0x43, 0x52, 0x45, 0x45, 0xaa, 0x02, 0x32, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x32, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0xe2,
	0x02, 0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x39, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x3a,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_rawDescData = file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_goTypes = []interface{}{
	(ApiResourceEventType)(0), // 0: cloud.planton.apis.v1.commons.resource.enums.event.ApiResourceEventType
}
var file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_init() }
func file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_init() {
	if File_cloud_planton_apis_v1_commons_resource_enums_event_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_resource_enums_event_event_proto = out.File
	file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_resource_enums_event_event_proto_depIdxs = nil
}
