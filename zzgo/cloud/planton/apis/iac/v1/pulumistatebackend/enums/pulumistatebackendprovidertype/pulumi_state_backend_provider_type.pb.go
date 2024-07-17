// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/iac/v1/pulumistatebackend/enums/pulumistatebackendprovidertype/pulumi_state_backend_provider_type.proto

package pulumistatebackendprovidertype

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

// https://www.pulumi.com/docs/concepts/state/#pulumi-cloud-backend
type PulumiStateBackendProviderType int32

const (
	PulumiStateBackendProviderType_unspecified          PulumiStateBackendProviderType = 0
	PulumiStateBackendProviderType_pulumi_cloud         PulumiStateBackendProviderType = 1
	PulumiStateBackendProviderType_aws_s3               PulumiStateBackendProviderType = 2
	PulumiStateBackendProviderType_google_cloud_storage PulumiStateBackendProviderType = 3
	PulumiStateBackendProviderType_azure_blob_storage   PulumiStateBackendProviderType = 4
)

// Enum value maps for PulumiStateBackendProviderType.
var (
	PulumiStateBackendProviderType_name = map[int32]string{
		0: "unspecified",
		1: "pulumi_cloud",
		2: "aws_s3",
		3: "google_cloud_storage",
		4: "azure_blob_storage",
	}
	PulumiStateBackendProviderType_value = map[string]int32{
		"unspecified":          0,
		"pulumi_cloud":         1,
		"aws_s3":               2,
		"google_cloud_storage": 3,
		"azure_blob_storage":   4,
	}
)

func (x PulumiStateBackendProviderType) Enum() *PulumiStateBackendProviderType {
	p := new(PulumiStateBackendProviderType)
	*p = x
	return p
}

func (x PulumiStateBackendProviderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PulumiStateBackendProviderType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_enumTypes[0].Descriptor()
}

func (PulumiStateBackendProviderType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_enumTypes[0]
}

func (x PulumiStateBackendProviderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PulumiStateBackendProviderType.Descriptor instead.
func (PulumiStateBackendProviderType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_rawDesc = []byte{
	0x0a, 0x7a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x6c, 0x75,
	0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x51, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x2a,
	0x81, 0x01, 0x0a, 0x1e, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x5f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x61, 0x77, 0x73, 0x5f, 0x73, 0x33, 0x10,
	0x02, 0x12, 0x18, 0x0a, 0x14, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x61,
	0x7a, 0x75, 0x72, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x10, 0x04, 0x42, 0xfa, 0x04, 0x0a, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x42, 0x23, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x81,
	0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d,
	0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x74, 0x79,
	0x70, 0x65, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x49, 0x56, 0x50, 0x45, 0x50, 0xaa, 0x02, 0x51,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70,
	0x69, 0x73, 0x2e, 0x49, 0x61, 0x63, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x74, 0x79, 0x70,
	0x65, 0xca, 0x02, 0x51, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x49, 0x61, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x75,
	0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x74, 0x79, 0x70, 0x65, 0xe2, 0x02, 0x5d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x49, 0x61, 0x63, 0x5c, 0x56,
	0x31, 0x5c, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x50, 0x75, 0x6c, 0x75, 0x6d,
	0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x58, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x49, 0x61,
	0x63, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x3a, 0x3a, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_rawDescData = file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_rawDesc
)

func file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_rawDescData)
	})
	return file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_rawDescData
}

var file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_goTypes = []interface{}{
	(PulumiStateBackendProviderType)(0), // 0: cloud.planton.apis.iac.v1.pulumistatebackend.enums.pulumistatebackendprovidertype.PulumiStateBackendProviderType
}
var file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_init()
}
func file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_init() {
	if File_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto = out.File
	file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_rawDesc = nil
	file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_goTypes = nil
	file_cloud_planton_apis_iac_v1_pulumistatebackend_enums_pulumistatebackendprovidertype_pulumi_state_backend_provider_type_proto_depIdxs = nil
}
