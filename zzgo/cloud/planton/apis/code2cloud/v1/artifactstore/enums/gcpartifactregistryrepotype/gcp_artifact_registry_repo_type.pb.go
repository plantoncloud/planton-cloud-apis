// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/artifactstore/enums/gcpartifactregistryrepotype/gcp_artifact_registry_repo_type.proto

package gcpartifactregistryrepotype

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

// note: the values are intentionally left capitalized as required by the gcp artifact-registry apis
type GcpArtifactRegistryRepoType int32

const (
	GcpArtifactRegistryRepoType_GCP_ARTIFACT_REGISTRY_REPO_TYPE_UNSPECIFIED GcpArtifactRegistryRepoType = 0
	GcpArtifactRegistryRepoType_DOCKER                                      GcpArtifactRegistryRepoType = 1
	GcpArtifactRegistryRepoType_NPM                                         GcpArtifactRegistryRepoType = 2
	GcpArtifactRegistryRepoType_MAVEN                                       GcpArtifactRegistryRepoType = 3
	GcpArtifactRegistryRepoType_PYTHON                                      GcpArtifactRegistryRepoType = 4
	GcpArtifactRegistryRepoType_GO                                          GcpArtifactRegistryRepoType = 5
)

// Enum value maps for GcpArtifactRegistryRepoType.
var (
	GcpArtifactRegistryRepoType_name = map[int32]string{
		0: "GCP_ARTIFACT_REGISTRY_REPO_TYPE_UNSPECIFIED",
		1: "DOCKER",
		2: "NPM",
		3: "MAVEN",
		4: "PYTHON",
		5: "GO",
	}
	GcpArtifactRegistryRepoType_value = map[string]int32{
		"GCP_ARTIFACT_REGISTRY_REPO_TYPE_UNSPECIFIED": 0,
		"DOCKER": 1,
		"NPM":    2,
		"MAVEN":  3,
		"PYTHON": 4,
		"GO":     5,
	}
)

func (x GcpArtifactRegistryRepoType) Enum() *GcpArtifactRegistryRepoType {
	p := new(GcpArtifactRegistryRepoType)
	*p = x
	return p
}

func (x GcpArtifactRegistryRepoType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GcpArtifactRegistryRepoType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_enumTypes[0].Descriptor()
}

func (GcpArtifactRegistryRepoType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_enumTypes[0]
}

func (x GcpArtifactRegistryRepoType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GcpArtifactRegistryRepoType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GcpArtifactRegistryRepoType(num)
	return nil
}

// Deprecated: Use GcpArtifactRegistryRepoType.Descriptor instead.
func (GcpArtifactRegistryRepoType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_rawDesc = []byte{
	0x0a, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x72, 0x65, 0x70, 0x6f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x67, 0x63, 0x70, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x50, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x67,
	0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x72, 0x65, 0x70, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x82, 0x01, 0x0a, 0x1b, 0x47,
	0x63, 0x70, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x2b, 0x47, 0x43,
	0x50, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x46, 0x41, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53,
	0x54, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x4f, 0x43, 0x4b, 0x45, 0x52, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x50, 0x4d, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x56, 0x45, 0x4e, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x50,
	0x59, 0x54, 0x48, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x4f, 0x10, 0x05, 0x42,
	0xf1, 0x04, 0x0a, 0x5e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65,
	0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x72, 0x65, 0x70, 0x6f, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x20, 0x47, 0x63, 0x70, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x80, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x67, 0x63,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x72, 0x65, 0x70, 0x6f, 0x74, 0x79, 0x70, 0x65, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x43,
	0x56, 0x41, 0x45, 0x47, 0xaa, 0x02, 0x50, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x47, 0x63, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x72,
	0x65, 0x70, 0x6f, 0x74, 0x79, 0x70, 0x65, 0xca, 0x02, 0x50, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x47,
	0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x72, 0x65, 0x70, 0x6f, 0x74, 0x79, 0x70, 0x65, 0xe2, 0x02, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0x5c, 0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x72, 0x65, 0x70, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x57, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73,
	0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31,
	0x3a, 0x3a, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x72, 0x65, 0x70, 0x6f, 0x74,
	0x79, 0x70, 0x65,
}

var (
	file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_goTypes = []interface{}{
	(GcpArtifactRegistryRepoType)(0), // 0: cloud.planton.apis.code2cloud.v1.artifactstore.enums.gcpartifactregistryrepotype.GcpArtifactRegistryRepoType
}
var file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_init()
}
func file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_artifactstore_enums_gcpartifactregistryrepotype_gcp_artifact_registry_repo_type_proto_depIdxs = nil
}
