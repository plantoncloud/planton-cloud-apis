// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/stack/job/progress/enums/event/event.proto

package event

import (
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

// The type of the stack job progress event.
// These events are emitted while a stack job is running.
type StackJobProgressEventType int32

const (
	StackJobProgressEventType_STACK_JOB_PROGRESS_EVENT_TYPE_UNSPECIFIED StackJobProgressEventType = 0
	StackJobProgressEventType_status_event                              StackJobProgressEventType = 1
	StackJobProgressEventType_pulumi_engine_event                       StackJobProgressEventType = 12
)

// Enum value maps for StackJobProgressEventType.
var (
	StackJobProgressEventType_name = map[int32]string{
		0:  "STACK_JOB_PROGRESS_EVENT_TYPE_UNSPECIFIED",
		1:  "status_event",
		12: "pulumi_engine_event",
	}
	StackJobProgressEventType_value = map[string]int32{
		"STACK_JOB_PROGRESS_EVENT_TYPE_UNSPECIFIED": 0,
		"status_event":        1,
		"pulumi_engine_event": 12,
	}
)

func (x StackJobProgressEventType) Enum() *StackJobProgressEventType {
	p := new(StackJobProgressEventType)
	*p = x
	return p
}

func (x StackJobProgressEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StackJobProgressEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_enumTypes[0].Descriptor()
}

func (StackJobProgressEventType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_enumTypes[0]
}

func (x StackJobProgressEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StackJobProgressEventType.Descriptor instead.
func (StackJobProgressEventType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_rawDesc = []byte{
	0x0a, 0x40, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2a, 0x75, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x29, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x4a,
	0x4f, 0x42, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x0c, 0x42,
	0xb4, 0x03, 0x0a, 0x42, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65,
	0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x64, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41,
	0x56, 0x53, 0x4a, 0x50, 0x45, 0x45, 0xaa, 0x02, 0x34, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x34,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4a, 0x6f, 0x62, 0x5c,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0xe2, 0x02, 0x40, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x5c, 0x4a, 0x6f, 0x62, 0x5c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5c,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x4a, 0x6f, 0x62, 0x3a, 0x3a,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a,
	0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_rawDescData = file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_rawDesc
)

func file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_rawDescData
}

var file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_goTypes = []interface{}{
	(StackJobProgressEventType)(0), // 0: cloud.planton.apis.v1.stack.job.progress.enums.event.StackJobProgressEventType
}
var file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_init() }
func file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_init() {
	if File_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto = out.File
	file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_rawDesc = nil
	file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_goTypes = nil
	file_cloud_planton_apis_v1_stack_job_progress_enums_event_event_proto_depIdxs = nil
}