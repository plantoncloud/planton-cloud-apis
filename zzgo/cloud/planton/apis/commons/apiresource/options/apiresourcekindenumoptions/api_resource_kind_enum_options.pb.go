// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/commons/apiresource/options/apiresourcekindenumoptions/api_resource_kind_enum_options.proto

package apiresourcekindenumoptions

import (
	apiresourcedeploytarget "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/enums/apiresourcedeploytarget"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         90100,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.id_prefix",
		Tag:           "bytes,90100,opt,name=id_prefix",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcekindenumoptions/api_resource_kind_enum_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         90101,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.kind_name",
		Tag:           "bytes,90101,opt,name=kind_name",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcekindenumoptions/api_resource_kind_enum_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         90102,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.display_name",
		Tag:           "bytes,90102,opt,name=display_name",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcekindenumoptions/api_resource_kind_enum_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         90103,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.api_resource_group",
		Tag:           "bytes,90103,opt,name=api_resource_group",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcekindenumoptions/api_resource_kind_enum_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         90104,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.is_deployable",
		Tag:           "varint,90104,opt,name=is_deployable",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcekindenumoptions/api_resource_kind_enum_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*apiresourcedeploytarget.ApiResourceDeploymentTarget)(nil),
		Field:         90105,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.deployment_target",
		Tag:           "varint,90105,opt,name=deployment_target,enum=cloud.planton.apis.commons.apiresource.enums.apiresourcedeploytarget.ApiResourceDeploymentTarget",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcekindenumoptions/api_resource_kind_enum_options.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string id_prefix = 90100;
	E_IdPrefix = &file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_extTypes[0]
	// optional string kind_name = 90101;
	E_KindName = &file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_extTypes[1]
	// optional string display_name = 90102;
	E_DisplayName = &file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_extTypes[2]
	// optional string api_resource_group = 90103;
	E_ApiResourceGroup = &file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_extTypes[3]
	// optional bool is_deployable = 90104;
	E_IsDeployable = &file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_extTypes[4]
	// optional cloud.planton.apis.commons.apiresource.enums.apiresourcedeploytarget.ApiResourceDeploymentTarget deployment_target = 90105;
	E_DeploymentTarget = &file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_extTypes[5]
)

var File_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_rawDesc = []byte{
	0x0a, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64,
	0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x49, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64,
	0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x40, 0x0a, 0x09, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf4, 0xbf, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x3a, 0x40, 0x0a, 0x09, 0x6b, 0x69, 0x6e, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf5, 0xbf, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6b, 0x69, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x46, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf6, 0xbf, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x3a, 0x51, 0x0a, 0x12, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf7, 0xbf, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x61, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x3a, 0x48, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf8, 0xbf, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x69, 0x73, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x3a, 0xb3, 0x01,
	0x0a, 0x11, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf9, 0xbf, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x61,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x42, 0xc3, 0x04, 0x0a, 0x57, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6b, 0x69, 0x6e, 0x64, 0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x1f, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x79, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a,
	0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69,
	0x6e, 0x64, 0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x07,
	0x43, 0x50, 0x41, 0x43, 0x41, 0x4f, 0x41, 0xaa, 0x02, 0x49, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0xca, 0x02, 0x49, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x5c, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6b, 0x69, 0x6e, 0x64, 0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xe2,
	0x02, 0x55, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x5c, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64,
	0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x4f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x3a, 0x3a, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x41,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x65, 0x6e,
	0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_goTypes = []interface{}{
	(*descriptorpb.EnumValueOptions)(nil),                    // 0: google.protobuf.EnumValueOptions
	(apiresourcedeploytarget.ApiResourceDeploymentTarget)(0), // 1: cloud.planton.apis.commons.apiresource.enums.apiresourcedeploytarget.ApiResourceDeploymentTarget
}
var file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.id_prefix:extendee -> google.protobuf.EnumValueOptions
	0, // 1: cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.kind_name:extendee -> google.protobuf.EnumValueOptions
	0, // 2: cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.display_name:extendee -> google.protobuf.EnumValueOptions
	0, // 3: cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.api_resource_group:extendee -> google.protobuf.EnumValueOptions
	0, // 4: cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.is_deployable:extendee -> google.protobuf.EnumValueOptions
	0, // 5: cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.deployment_target:extendee -> google.protobuf.EnumValueOptions
	1, // 6: cloud.planton.apis.commons.apiresource.options.apiresourcekindenumoptions.deployment_target:type_name -> cloud.planton.apis.commons.apiresource.enums.apiresourcedeploytarget.ApiResourceDeploymentTarget
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	6, // [6:7] is the sub-list for extension type_name
	0, // [0:6] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_init()
}
func file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_init() {
	if File_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_depIdxs,
		ExtensionInfos:    file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_extTypes,
	}.Build()
	File_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto = out.File
	file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_rawDesc = nil
	file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_goTypes = nil
	file_cloud_planton_apis_commons_apiresource_options_apiresourcekindenumoptions_api_resource_kind_enum_options_proto_depIdxs = nil
}
