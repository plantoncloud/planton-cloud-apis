// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/develop/sourcecode/server/rpc/enums/enums.proto

package enums

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

type CodeServerProvider int32

const (
	CodeServerProvider_CODE_SERVER_PROVIDER_UNSPECIFIED CodeServerProvider = 0
	CodeServerProvider_CODE_SERVER_PROVIDER_GITHUB      CodeServerProvider = 1
	CodeServerProvider_CODE_SERVER_PROVIDER_GITLAB      CodeServerProvider = 2
)

// Enum value maps for CodeServerProvider.
var (
	CodeServerProvider_name = map[int32]string{
		0: "CODE_SERVER_PROVIDER_UNSPECIFIED",
		1: "CODE_SERVER_PROVIDER_GITHUB",
		2: "CODE_SERVER_PROVIDER_GITLAB",
	}
	CodeServerProvider_value = map[string]int32{
		"CODE_SERVER_PROVIDER_UNSPECIFIED": 0,
		"CODE_SERVER_PROVIDER_GITHUB":      1,
		"CODE_SERVER_PROVIDER_GITLAB":      2,
	}
)

func (x CodeServerProvider) Enum() *CodeServerProvider {
	p := new(CodeServerProvider)
	*p = x
	return p
}

func (x CodeServerProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodeServerProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_enumTypes[0].Descriptor()
}

func (CodeServerProvider) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_enumTypes[0]
}

func (x CodeServerProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodeServerProvider.Descriptor instead.
func (CodeServerProvider) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDescGZIP(), []int{0}
}

type GithubAppInstallOwnerType int32

const (
	GithubAppInstallOwnerType_GITHUB_APP_INSTALLATION_OWNER_TYPE_UNSPECIFIED GithubAppInstallOwnerType = 0
	GithubAppInstallOwnerType_GITHUB_APP_INSTALLATION_OWNER_TYPE_USER        GithubAppInstallOwnerType = 1
	GithubAppInstallOwnerType_GITHUB_APP_INSTALLATION_OWNER_ORG              GithubAppInstallOwnerType = 2
)

// Enum value maps for GithubAppInstallOwnerType.
var (
	GithubAppInstallOwnerType_name = map[int32]string{
		0: "GITHUB_APP_INSTALLATION_OWNER_TYPE_UNSPECIFIED",
		1: "GITHUB_APP_INSTALLATION_OWNER_TYPE_USER",
		2: "GITHUB_APP_INSTALLATION_OWNER_ORG",
	}
	GithubAppInstallOwnerType_value = map[string]int32{
		"GITHUB_APP_INSTALLATION_OWNER_TYPE_UNSPECIFIED": 0,
		"GITHUB_APP_INSTALLATION_OWNER_TYPE_USER":        1,
		"GITHUB_APP_INSTALLATION_OWNER_ORG":              2,
	}
)

func (x GithubAppInstallOwnerType) Enum() *GithubAppInstallOwnerType {
	p := new(GithubAppInstallOwnerType)
	*p = x
	return p
}

func (x GithubAppInstallOwnerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GithubAppInstallOwnerType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_enumTypes[1].Descriptor()
}

func (GithubAppInstallOwnerType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_enumTypes[1]
}

func (x GithubAppInstallOwnerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GithubAppInstallOwnerType.Descriptor instead.
func (GithubAppInstallOwnerType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDescGZIP(), []int{1}
}

var File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDesc = []byte{
	0x0a, 0x50, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x44, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0x7c, 0x0a, 0x12, 0x43, 0x6f, 0x64, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x20, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x50, 0x52,
	0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x47, 0x49, 0x54,
	0x48, 0x55, 0x42, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x47, 0x49,
	0x54, 0x4c, 0x41, 0x42, 0x10, 0x02, 0x2a, 0xa3, 0x01, 0x0a, 0x19, 0x47, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x2e, 0x47, 0x49, 0x54, 0x48, 0x55, 0x42, 0x5f, 0x41,
	0x50, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4f, 0x57, 0x4e, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2b, 0x0a, 0x27, 0x47, 0x49, 0x54, 0x48,
	0x55, 0x42, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x47, 0x49, 0x54, 0x48, 0x55, 0x42, 0x5f,
	0x41, 0x50, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x5f, 0x4f, 0x52, 0x47, 0x10, 0x02, 0x42, 0x8c, 0x04, 0x0a,
	0x48, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x74, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x0a,
	0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x53, 0x53, 0x52, 0x45, 0xaa, 0x02, 0x44, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e,
	0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x70, 0x63, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xca, 0x02, 0x44, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x5c, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x52,
	0x70, 0x63, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xe2, 0x02, 0x50, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x5c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x4d, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x3a, 0x3a, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a,
	0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_goTypes = []interface{}{
	(CodeServerProvider)(0),        // 0: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.rpc.enums.CodeServerProvider
	(GithubAppInstallOwnerType)(0), // 1: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.rpc.enums.GithubAppInstallOwnerType
}
var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_init()
}
func file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_server_rpc_enums_enums_proto_depIdxs = nil
}
