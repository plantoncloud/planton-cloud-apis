// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/kubecluster/provider/awsaks/enums/workersmanagementmode/workers_management_mode.proto

package workersmanagementmode

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

type AwsEksWorkersManagementMode int32

const (
	AwsEksWorkersManagementMode_AWS_EKA_WORKERS_MANAGEMENT_MODE_UNSPECIFIED AwsEksWorkersManagementMode = 0
	AwsEksWorkersManagementMode_ec2                                         AwsEksWorkersManagementMode = 1
	AwsEksWorkersManagementMode_fargate                                     AwsEksWorkersManagementMode = 2
	AwsEksWorkersManagementMode_managed_node_groups                         AwsEksWorkersManagementMode = 3
)

// Enum value maps for AwsEksWorkersManagementMode.
var (
	AwsEksWorkersManagementMode_name = map[int32]string{
		0: "AWS_EKA_WORKERS_MANAGEMENT_MODE_UNSPECIFIED",
		1: "ec2",
		2: "fargate",
		3: "managed_node_groups",
	}
	AwsEksWorkersManagementMode_value = map[string]int32{
		"AWS_EKA_WORKERS_MANAGEMENT_MODE_UNSPECIFIED": 0,
		"ec2":                 1,
		"fargate":             2,
		"managed_node_groups": 3,
	}
)

func (x AwsEksWorkersManagementMode) Enum() *AwsEksWorkersManagementMode {
	p := new(AwsEksWorkersManagementMode)
	*p = x
	return p
}

func (x AwsEksWorkersManagementMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AwsEksWorkersManagementMode) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_enumTypes[0].Descriptor()
}

func (AwsEksWorkersManagementMode) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_enumTypes[0]
}

func (x AwsEksWorkersManagementMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AwsEksWorkersManagementMode.Descriptor instead.
func (AwsEksWorkersManagementMode) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_rawDesc = []byte{
	0x0a, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x61, 0x6b, 0x73, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x58, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x61, 0x77, 0x73, 0x61, 0x6b, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x6f,
	0x64, 0x65, 0x2a, 0x7d, 0x0a, 0x1b, 0x41, 0x77, 0x73, 0x45, 0x6b, 0x73, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x57, 0x53, 0x5f, 0x45, 0x4b, 0x41, 0x5f, 0x57, 0x4f, 0x52,
	0x4b, 0x45, 0x52, 0x53, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x65, 0x63, 0x32, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x66,
	0x61, 0x72, 0x67, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x64, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x10,
	0x03, 0x42, 0x9f, 0x05, 0x0a, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x77, 0x73, 0x61, 0x6b, 0x73,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x42, 0x1a, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x88, 0x01, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75,
	0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x61, 0x6b, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x6d, 0x6f, 0x64, 0x65, 0xa2, 0x02, 0x0a, 0x43, 0x50, 0x41, 0x43, 0x56, 0x4b, 0x50, 0x41,
	0x45, 0x57, 0xaa, 0x02, 0x58, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x77, 0x73, 0x61, 0x6b,
	0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0xca, 0x02, 0x58,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31,
	0x5c, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x41, 0x77, 0x73, 0x61, 0x6b, 0x73, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x5c, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0xe2, 0x02, 0x64, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x4b, 0x75, 0x62, 0x65,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x5c, 0x41, 0x77, 0x73, 0x61, 0x6b, 0x73, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d,
	0x6f, 0x64, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x61, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x41,
	0x77, 0x73, 0x61, 0x6b, 0x73, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d,
	0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_goTypes = []interface{}{
	(AwsEksWorkersManagementMode)(0), // 0: cloud.planton.apis.code2cloud.v1.kubecluster.provider.awsaks.enums.workersmanagementmode.AwsEksWorkersManagementMode
}
var file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_init()
}
func file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_kubecluster_provider_awsaks_enums_workersmanagementmode_workers_management_mode_proto_depIdxs = nil
}