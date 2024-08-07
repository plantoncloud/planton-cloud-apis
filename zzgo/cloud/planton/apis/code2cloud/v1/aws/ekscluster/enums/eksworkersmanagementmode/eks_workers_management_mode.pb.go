// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/aws/ekscluster/enums/eksworkersmanagementmode/eks_workers_management_mode.proto

package eksworkersmanagementmode

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

type EksWorkersManagementMode int32

const (
	EksWorkersManagementMode_unspecified         EksWorkersManagementMode = 0
	EksWorkersManagementMode_ec2                 EksWorkersManagementMode = 1
	EksWorkersManagementMode_fargate             EksWorkersManagementMode = 2
	EksWorkersManagementMode_managed_node_groups EksWorkersManagementMode = 3
)

// Enum value maps for EksWorkersManagementMode.
var (
	EksWorkersManagementMode_name = map[int32]string{
		0: "unspecified",
		1: "ec2",
		2: "fargate",
		3: "managed_node_groups",
	}
	EksWorkersManagementMode_value = map[string]int32{
		"unspecified":         0,
		"ec2":                 1,
		"fargate":             2,
		"managed_node_groups": 3,
	}
)

func (x EksWorkersManagementMode) Enum() *EksWorkersManagementMode {
	p := new(EksWorkersManagementMode)
	*p = x
	return p
}

func (x EksWorkersManagementMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EksWorkersManagementMode) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_enumTypes[0].Descriptor()
}

func (EksWorkersManagementMode) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_enumTypes[0]
}

func (x EksWorkersManagementMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EksWorkersManagementMode.Descriptor instead.
func (EksWorkersManagementMode) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_rawDesc = []byte{
	0x0a, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x65, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6b, 0x73, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x6f, 0x64, 0x65,
	0x2f, 0x65, 0x6b, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x4e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x65, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x65, 0x6b, 0x73, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x6f,
	0x64, 0x65, 0x2a, 0x5a, 0x0a, 0x18, 0x45, 0x6b, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x65, 0x63, 0x32, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x66, 0x61, 0x72, 0x67,
	0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x10, 0x03, 0x42, 0xe3,
	0x04, 0x0a, 0x5c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x65, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x65, 0x6b, 0x73, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x42,
	0x1d, 0x45, 0x6b, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x7e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x65, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6b, 0x73, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x6f, 0x64, 0x65,
	0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x43, 0x56, 0x41, 0x45, 0x45, 0x45, 0xaa, 0x02, 0x4e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69,
	0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e,
	0x41, 0x77, 0x73, 0x2e, 0x45, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x45, 0x6b, 0x73, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0xca, 0x02, 0x4e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31,
	0x5c, 0x41, 0x77, 0x73, 0x5c, 0x45, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x45, 0x6b, 0x73, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0xe2, 0x02,
	0x5a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56,
	0x31, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x45, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x45, 0x6b, 0x73, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x56, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x56, 0x31, 0x3a, 0x3a, 0x41, 0x77, 0x73, 0x3a, 0x3a, 0x45, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x45, 0x6b, 0x73, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x6d, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_goTypes = []interface{}{
	(EksWorkersManagementMode)(0), // 0: cloud.planton.apis.code2cloud.v1.aws.ekscluster.enums.eksworkersmanagementmode.EksWorkersManagementMode
}
var file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_init()
}
func file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_aws_ekscluster_enums_eksworkersmanagementmode_eks_workers_management_mode_proto_depIdxs = nil
}
