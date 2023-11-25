// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/stack/job/enums/executionstatus/execution_status.proto

package executionstatus

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

type StackJobExecutionStatus int32

const (
	StackJobExecutionStatus_STACK_JOB_EXECUTION_STATUS_UNSPECIFIED StackJobExecutionStatus = 0
	StackJobExecutionStatus_pending                                StackJobExecutionStatus = 1
	StackJobExecutionStatus_running                                StackJobExecutionStatus = 2
	StackJobExecutionStatus_succeeded                              StackJobExecutionStatus = 3
	StackJobExecutionStatus_failed                                 StackJobExecutionStatus = 4
	StackJobExecutionStatus_cancelled                              StackJobExecutionStatus = 5
	StackJobExecutionStatus_skipped                                StackJobExecutionStatus = 6
)

// Enum value maps for StackJobExecutionStatus.
var (
	StackJobExecutionStatus_name = map[int32]string{
		0: "STACK_JOB_EXECUTION_STATUS_UNSPECIFIED",
		1: "pending",
		2: "running",
		3: "succeeded",
		4: "failed",
		5: "cancelled",
		6: "skipped",
	}
	StackJobExecutionStatus_value = map[string]int32{
		"STACK_JOB_EXECUTION_STATUS_UNSPECIFIED": 0,
		"pending":                                1,
		"running":                                2,
		"succeeded":                              3,
		"failed":                                 4,
		"cancelled":                              5,
		"skipped":                                6,
	}
)

func (x StackJobExecutionStatus) Enum() *StackJobExecutionStatus {
	p := new(StackJobExecutionStatus)
	*p = x
	return p
}

func (x StackJobExecutionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StackJobExecutionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_enumTypes[0].Descriptor()
}

func (StackJobExecutionStatus) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_enumTypes[0]
}

func (x StackJobExecutionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StackJobExecutionStatus.Descriptor instead.
func (StackJobExecutionStatus) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f,
	0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x96, 0x01, 0x0a, 0x17, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a,
	0x6f, 0x62, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2a, 0x0a, 0x26, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x45,
	0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x72, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x65, 0x64, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10,
	0x05, 0x12, 0x0b, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x10, 0x06, 0x42, 0xc2,
	0x03, 0x0a, 0x43, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x14, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x65,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x53, 0x4a, 0x45, 0x45,
	0xaa, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x4a,
	0x6f, 0x62, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0xca, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4a, 0x6f, 0x62, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x5c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0xe2, 0x02, 0x41, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4a,
	0x6f, 0x62, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x4a, 0x6f, 0x62, 0x3a, 0x3a, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x3a, 0x3a, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_rawDescData = file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_rawDesc
)

func file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_rawDescData
}

var file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_goTypes = []interface{}{
	(StackJobExecutionStatus)(0), // 0: cloud.planton.apis.v1.stack.job.enums.executionstatus.StackJobExecutionStatus
}
var file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_init() }
func file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_init() {
	if File_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto = out.File
	file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_rawDesc = nil
	file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_goTypes = nil
	file_cloud_planton_apis_v1_stack_job_enums_executionstatus_execution_status_proto_depIdxs = nil
}
