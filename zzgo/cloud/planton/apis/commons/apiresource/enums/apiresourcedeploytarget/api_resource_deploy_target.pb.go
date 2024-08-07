// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/commons/apiresource/enums/apiresourcedeploytarget/api_resource_deploy_target.proto

package apiresourcedeploytarget

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

type ApiResourceDeploymentTarget int32

const (
	ApiResourceDeploymentTarget_unspecified ApiResourceDeploymentTarget = 0
	ApiResourceDeploymentTarget_aws         ApiResourceDeploymentTarget = 1
	ApiResourceDeploymentTarget_gcp         ApiResourceDeploymentTarget = 2
	ApiResourceDeploymentTarget_azure       ApiResourceDeploymentTarget = 3
	ApiResourceDeploymentTarget_kubernetes  ApiResourceDeploymentTarget = 4
)

// Enum value maps for ApiResourceDeploymentTarget.
var (
	ApiResourceDeploymentTarget_name = map[int32]string{
		0: "unspecified",
		1: "aws",
		2: "gcp",
		3: "azure",
		4: "kubernetes",
	}
	ApiResourceDeploymentTarget_value = map[string]int32{
		"unspecified": 0,
		"aws":         1,
		"gcp":         2,
		"azure":       3,
		"kubernetes":  4,
	}
)

func (x ApiResourceDeploymentTarget) Enum() *ApiResourceDeploymentTarget {
	p := new(ApiResourceDeploymentTarget)
	*p = x
	return p
}

func (x ApiResourceDeploymentTarget) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiResourceDeploymentTarget) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_enumTypes[0].Descriptor()
}

func (ApiResourceDeploymentTarget) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_enumTypes[0]
}

func (x ApiResourceDeploymentTarget) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiResourceDeploymentTarget.Descriptor instead.
func (ApiResourceDeploymentTarget) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_rawDesc = []byte{
	0x0a, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x44, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2a, 0x5b, 0x0a,
	0x1b, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x0f, 0x0a, 0x0b,
	0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x61, 0x77, 0x73, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x67, 0x63, 0x70, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x10, 0x04, 0x42, 0xa2, 0x04, 0x0a, 0x52, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x42, 0x1c, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x74, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a,
	0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x43, 0x41, 0x45,
	0x41, 0xaa, 0x02, 0x44, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0xca, 0x02, 0x44, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0xe2,
	0x02, 0x50, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x41,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x4a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_rawDescData = file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_rawDesc
)

func file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_rawDescData)
	})
	return file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_rawDescData
}

var file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_goTypes = []interface{}{
	(ApiResourceDeploymentTarget)(0), // 0: cloud.planton.apis.commons.apiresource.enums.apiresourcedeploytarget.ApiResourceDeploymentTarget
}
var file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_init()
}
func file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_init() {
	if File_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto = out.File
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_rawDesc = nil
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_goTypes = nil
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcedeploytarget_api_resource_deploy_target_proto_depIdxs = nil
}
