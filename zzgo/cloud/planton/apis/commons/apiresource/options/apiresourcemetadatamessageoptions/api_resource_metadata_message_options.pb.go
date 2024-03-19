// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto

package apiresourcemetadatamessageoptions

import (
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

var file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50001,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.id_prefix",
		Tag:           "bytes,50001,opt,name=id_prefix",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50002,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.is_id_required",
		Tag:           "varint,50002,opt,name=is_id_required",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50003,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.is_id_computed",
		Tag:           "varint,50003,opt,name=is_id_computed",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50004,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.id_regex",
		Tag:           "bytes,50004,opt,name=id_regex",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50005,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.id_min_length",
		Tag:           "varint,50005,opt,name=id_min_length",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50006,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.id_max_length",
		Tag:           "varint,50006,opt,name=id_max_length",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50007,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.id_constraints_description",
		Tag:           "bytes,50007,opt,name=id_constraints_description",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50008,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.name_regex",
		Tag:           "bytes,50008,opt,name=name_regex",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50009,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.name_min_length",
		Tag:           "varint,50009,opt,name=name_min_length",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50010,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.name_max_length",
		Tag:           "varint,50010,opt,name=name_max_length",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50011,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.name_constraints_description",
		Tag:           "bytes,50011,opt,name=name_constraints_description",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: ([]string)(nil),
		Field:         50012,
		Name:          "cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.id_field_paths_to_concat",
		Tag:           "bytes,50012,rep,name=id_field_paths_to_concat",
		Filename:      "cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions/api_resource_metadata_message_options.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional string id_prefix = 50001;
	E_IdPrefix = &file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes[0]
	// optional bool is_id_required = 50002;
	E_IsIdRequired = &file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes[1]
	// optional bool is_id_computed = 50003;
	E_IsIdComputed = &file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes[2]
	// optional string id_regex = 50004;
	E_IdRegex = &file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes[3]
	// optional int32 id_min_length = 50005;
	E_IdMinLength = &file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes[4]
	// optional int32 id_max_length = 50006;
	E_IdMaxLength = &file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes[5]
	// optional string id_constraints_description = 50007;
	E_IdConstraintsDescription = &file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes[6]
	// optional string name_regex = 50008;
	E_NameRegex = &file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes[7]
	// optional int32 name_min_length = 50009;
	E_NameMinLength = &file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes[8]
	// optional int32 name_max_length = 50010;
	E_NameMaxLength = &file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes[9]
	// optional string name_constraints_description = 50011;
	E_NameConstraintsDescription = &file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes[10]
	// repeated string id_field_paths_to_concat = 50012;
	E_IdFieldPathsToConcat = &file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes[11]
)

var File_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_rawDesc = []byte{
	0x0a, 0x7c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x50,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3a, 0x3e, 0x0a, 0x09, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x3a, 0x47, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0x86, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x3a, 0x47, 0x0a, 0x0e, 0x69,
	0x73, 0x5f, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3,
	0x86, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x49, 0x64, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x64, 0x3a, 0x3c, 0x0a, 0x08, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xd4, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x52, 0x65, 0x67,
	0x65, 0x78, 0x3a, 0x45, 0x0a, 0x0d, 0x69, 0x64, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd5, 0x86, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x64,
	0x4d, 0x69, 0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3a, 0x45, 0x0a, 0x0d, 0x69, 0x64, 0x5f,
	0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd6, 0x86, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x64, 0x4d, 0x61, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x3a, 0x5f, 0x0a, 0x1a, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd7, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x69, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x3a, 0x40, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xd8, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x67, 0x65, 0x78, 0x3a, 0x49, 0x0a, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x5f,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd9, 0x86, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x4d, 0x69, 0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3a, 0x49,
	0x0a, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xda, 0x86, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x61, 0x6d, 0x65,
	0x4d, 0x61, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3a, 0x63, 0x0a, 0x1c, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdb, 0x86, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x1a, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x59,
	0x0a, 0x18, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73,
	0x5f, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdc, 0x86, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x14, 0x69, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68,
	0x73, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x42, 0xf5, 0x04, 0x0a, 0x5e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x26, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x80, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x43,
	0x41, 0x4f, 0x41, 0xaa, 0x02, 0x50, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x50, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x56, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a,
	0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x3a, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x3a,
	0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_goTypes = []interface{}{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
}
var file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_depIdxs = []int32{
	0,  // 0: cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.id_prefix:extendee -> google.protobuf.MessageOptions
	0,  // 1: cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.is_id_required:extendee -> google.protobuf.MessageOptions
	0,  // 2: cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.is_id_computed:extendee -> google.protobuf.MessageOptions
	0,  // 3: cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.id_regex:extendee -> google.protobuf.MessageOptions
	0,  // 4: cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.id_min_length:extendee -> google.protobuf.MessageOptions
	0,  // 5: cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.id_max_length:extendee -> google.protobuf.MessageOptions
	0,  // 6: cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.id_constraints_description:extendee -> google.protobuf.MessageOptions
	0,  // 7: cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.name_regex:extendee -> google.protobuf.MessageOptions
	0,  // 8: cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.name_min_length:extendee -> google.protobuf.MessageOptions
	0,  // 9: cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.name_max_length:extendee -> google.protobuf.MessageOptions
	0,  // 10: cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.name_constraints_description:extendee -> google.protobuf.MessageOptions
	0,  // 11: cloud.planton.apis.commons.apiresource.options.apiresourcemetadatamessageoptions.id_field_paths_to_concat:extendee -> google.protobuf.MessageOptions
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	0,  // [0:12] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_init()
}
func file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_init() {
	if File_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 12,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_depIdxs,
		ExtensionInfos:    file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_extTypes,
	}.Build()
	File_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto = out.File
	file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_rawDesc = nil
	file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_goTypes = nil
	file_cloud_planton_apis_commons_apiresource_options_apiresourcemetadatamessageoptions_api_resource_metadata_message_options_proto_depIdxs = nil
}
