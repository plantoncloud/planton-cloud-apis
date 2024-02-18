// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/commons/resource/field/options/resource_field_options.proto

package options

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

var file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50056,
		Name:          "cloud.planton.apis.commons.resource.field.options.is_required",
		Tag:           "varint,50056,opt,name=is_required",
		Filename:      "cloud/planton/apis/commons/resource/field/options/resource_field_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50057,
		Name:          "cloud.planton.apis.commons.resource.field.options.is_computed",
		Tag:           "varint,50057,opt,name=is_computed",
		Filename:      "cloud/planton/apis/commons/resource/field/options/resource_field_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50058,
		Name:          "cloud.planton.apis.commons.resource.field.options.is_immutable",
		Tag:           "varint,50058,opt,name=is_immutable",
		Filename:      "cloud/planton/apis/commons/resource/field/options/resource_field_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50059,
		Name:          "cloud.planton.apis.commons.resource.field.options.string_regex",
		Tag:           "bytes,50059,opt,name=string_regex",
		Filename:      "cloud/planton/apis/commons/resource/field/options/resource_field_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50060,
		Name:          "cloud.planton.apis.commons.resource.field.options.string_min_length",
		Tag:           "varint,50060,opt,name=string_min_length",
		Filename:      "cloud/planton/apis/commons/resource/field/options/resource_field_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50061,
		Name:          "cloud.planton.apis.commons.resource.field.options.string_max_length",
		Tag:           "varint,50061,opt,name=string_max_length",
		Filename:      "cloud/planton/apis/commons/resource/field/options/resource_field_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50062,
		Name:          "cloud.planton.apis.commons.resource.field.options.field_constraints_description",
		Tag:           "bytes,50062,opt,name=field_constraints_description",
		Filename:      "cloud/planton/apis/commons/resource/field/options/resource_field_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50063,
		Name:          "cloud.planton.apis.commons.resource.field.options.string_default",
		Tag:           "bytes,50063,opt,name=string_default",
		Filename:      "cloud/planton/apis/commons/resource/field/options/resource_field_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50064,
		Name:          "cloud.planton.apis.commons.resource.field.options.int32_default",
		Tag:           "varint,50064,opt,name=int32_default",
		Filename:      "cloud/planton/apis/commons/resource/field/options/resource_field_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int64)(nil),
		Field:         50065,
		Name:          "cloud.planton.apis.commons.resource.field.options.int64_default",
		Tag:           "varint,50065,opt,name=int64_default",
		Filename:      "cloud/planton/apis/commons/resource/field/options/resource_field_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50066,
		Name:          "cloud.planton.apis.commons.resource.field.options.bool_default",
		Tag:           "varint,50066,opt,name=bool_default",
		Filename:      "cloud/planton/apis/commons/resource/field/options/resource_field_options.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool is_required = 50056;
	E_IsRequired = &file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes[0]
	// optional bool is_computed = 50057;
	E_IsComputed = &file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes[1]
	// optional bool is_immutable = 50058;
	E_IsImmutable = &file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes[2]
	// optional string string_regex = 50059;
	E_StringRegex = &file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes[3]
	// optional int32 string_min_length = 50060;
	E_StringMinLength = &file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes[4]
	// optional int32 string_max_length = 50061;
	E_StringMaxLength = &file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes[5]
	// optional string field_constraints_description = 50062;
	E_FieldConstraintsDescription = &file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes[6]
	// optional string string_default = 50063;
	E_StringDefault = &file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes[7]
	// optional int32 int32_default = 50064;
	E_Int32Default = &file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes[8]
	// optional int64 int64_default = 50065;
	E_Int64Default = &file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes[9]
	// optional bool bool_default = 50066;
	E_BoolDefault = &file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes[10]
)

var File_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_rawDesc = []byte{
	0x0a, 0x51, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x40, 0x0a, 0x0b, 0x69,
	0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x88, 0x87, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x3a, 0x40, 0x0a,
	0x0b, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x89, 0x87, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x3a,
	0x42, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8a,
	0x87, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x49, 0x6d, 0x6d, 0x75, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x3a, 0x42, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65,
	0x67, 0x65, 0x78, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x8b, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x67, 0x65, 0x78, 0x3a, 0x4b, 0x0a, 0x11, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8c, 0x87, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x69, 0x6e, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x3a, 0x4b, 0x0a, 0x11, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d,
	0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8d, 0x87, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x3a, 0x63, 0x0a, 0x1d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x8e, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x46, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8f, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0x44,
	0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x90,
	0x87, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x3a, 0x44, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x91, 0x87, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0x42, 0x0a, 0x0c, 0x62, 0x6f,
	0x6f, 0x6c, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x92, 0x87, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0xc1,
	0x03, 0x0a, 0x42, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x19, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x64, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a,
	0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x43,
	0x52, 0x46, 0x4f, 0xaa, 0x02, 0x34, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x34, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x5c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0xe2, 0x02, 0x40, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x5c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x3a, 0x3a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x3a, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_depIdxs = []int32{
	0,  // 0: cloud.planton.apis.commons.resource.field.options.is_required:extendee -> google.protobuf.FieldOptions
	0,  // 1: cloud.planton.apis.commons.resource.field.options.is_computed:extendee -> google.protobuf.FieldOptions
	0,  // 2: cloud.planton.apis.commons.resource.field.options.is_immutable:extendee -> google.protobuf.FieldOptions
	0,  // 3: cloud.planton.apis.commons.resource.field.options.string_regex:extendee -> google.protobuf.FieldOptions
	0,  // 4: cloud.planton.apis.commons.resource.field.options.string_min_length:extendee -> google.protobuf.FieldOptions
	0,  // 5: cloud.planton.apis.commons.resource.field.options.string_max_length:extendee -> google.protobuf.FieldOptions
	0,  // 6: cloud.planton.apis.commons.resource.field.options.field_constraints_description:extendee -> google.protobuf.FieldOptions
	0,  // 7: cloud.planton.apis.commons.resource.field.options.string_default:extendee -> google.protobuf.FieldOptions
	0,  // 8: cloud.planton.apis.commons.resource.field.options.int32_default:extendee -> google.protobuf.FieldOptions
	0,  // 9: cloud.planton.apis.commons.resource.field.options.int64_default:extendee -> google.protobuf.FieldOptions
	0,  // 10: cloud.planton.apis.commons.resource.field.options.bool_default:extendee -> google.protobuf.FieldOptions
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	0,  // [0:11] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_init()
}
func file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_init() {
	if File_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 11,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_depIdxs,
		ExtensionInfos:    file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_extTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto = out.File
	file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_resource_field_options_resource_field_options_proto_depIdxs = nil
}
