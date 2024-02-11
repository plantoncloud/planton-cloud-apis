// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/stack/job/enums/operationtype/stackjoboperationtypeoptions/stack_job_operation_type_options.proto

package stackjoboperationtypeoptions

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

var file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50050,
		Name:          "cloud.planton.apis.v1.stack.job.enums.operationtype.stackjoboperationtypeoptions.is_refresh_required_required",
		Tag:           "varint,50050,opt,name=is_refresh_required_required",
		Filename:      "cloud/planton/apis/v1/stack/job/enums/operationtype/stackjoboperationtypeoptions/stack_job_operation_type_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50051,
		Name:          "cloud.planton.apis.v1.stack.job.enums.operationtype.stackjoboperationtypeoptions.is_apply_preview_required",
		Tag:           "varint,50051,opt,name=is_apply_preview_required",
		Filename:      "cloud/planton/apis/v1/stack/job/enums/operationtype/stackjoboperationtypeoptions/stack_job_operation_type_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50052,
		Name:          "cloud.planton.apis.v1.stack.job.enums.operationtype.stackjoboperationtypeoptions.is_destroy_preview_required",
		Tag:           "varint,50052,opt,name=is_destroy_preview_required",
		Filename:      "cloud/planton/apis/v1/stack/job/enums/operationtype/stackjoboperationtypeoptions/stack_job_operation_type_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50053,
		Name:          "cloud.planton.apis.v1.stack.job.enums.operationtype.stackjoboperationtypeoptions.is_apply_required",
		Tag:           "varint,50053,opt,name=is_apply_required",
		Filename:      "cloud/planton/apis/v1/stack/job/enums/operationtype/stackjoboperationtypeoptions/stack_job_operation_type_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50054,
		Name:          "cloud.planton.apis.v1.stack.job.enums.operationtype.stackjoboperationtypeoptions.is_destroy_required",
		Tag:           "varint,50054,opt,name=is_destroy_required",
		Filename:      "cloud/planton/apis/v1/stack/job/enums/operationtype/stackjoboperationtypeoptions/stack_job_operation_type_options.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional bool is_refresh_required_required = 50050;
	E_IsRefreshRequiredRequired = &file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_extTypes[0]
	// optional bool is_apply_preview_required = 50051;
	E_IsApplyPreviewRequired = &file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_extTypes[1]
	// optional bool is_destroy_preview_required = 50052;
	E_IsDestroyPreviewRequired = &file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_extTypes[2]
	// optional bool is_apply_required = 50053;
	E_IsApplyRequired = &file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_extTypes[3]
	// optional bool is_destroy_required = 50054;
	E_IsDestroyRequired = &file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_extTypes[4]
)

var File_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_rawDesc = []byte{
	0x0a, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f,
	0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x50, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x74, 0x79, 0x70, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x64, 0x0a,
	0x1c, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x21, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x82, 0x87, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x69, 0x73, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x3a, 0x5e, 0x0a, 0x19, 0x69, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x83, 0x87, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x69, 0x73, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x3a, 0x62, 0x0a, 0x1b, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f,
	0x79, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x84, 0x87, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x69,
	0x73, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x3a, 0x4f, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x61, 0x70,
	0x70, 0x6c, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x21, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x85, 0x87, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x3a, 0x53, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x64,
	0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12,
	0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x86, 0x87, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x44, 0x65,
	0x73, 0x74, 0x72, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x42, 0xf4, 0x04,
	0x0a, 0x5e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x21, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x80, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x6a, 0x6f, 0x62, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70,
	0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x56, 0x53,
	0x4a, 0x45, 0x4f, 0x53, 0xaa, 0x02, 0x50, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x6a, 0x6f, 0x62, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x50, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4a, 0x6f, 0x62, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x5c, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74,
	0x79, 0x70, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4a, 0x6f, 0x62, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x5c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65,
	0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x58, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x4a, 0x6f, 0x62,
	0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_goTypes = []interface{}{
	(*descriptorpb.EnumValueOptions)(nil), // 0: google.protobuf.EnumValueOptions
}
var file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.stack.job.enums.operationtype.stackjoboperationtypeoptions.is_refresh_required_required:extendee -> google.protobuf.EnumValueOptions
	0, // 1: cloud.planton.apis.v1.stack.job.enums.operationtype.stackjoboperationtypeoptions.is_apply_preview_required:extendee -> google.protobuf.EnumValueOptions
	0, // 2: cloud.planton.apis.v1.stack.job.enums.operationtype.stackjoboperationtypeoptions.is_destroy_preview_required:extendee -> google.protobuf.EnumValueOptions
	0, // 3: cloud.planton.apis.v1.stack.job.enums.operationtype.stackjoboperationtypeoptions.is_apply_required:extendee -> google.protobuf.EnumValueOptions
	0, // 4: cloud.planton.apis.v1.stack.job.enums.operationtype.stackjoboperationtypeoptions.is_destroy_required:extendee -> google.protobuf.EnumValueOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	0, // [0:5] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_init()
}
func file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_init() {
	if File_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_depIdxs,
		ExtensionInfos:    file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_extTypes,
	}.Build()
	File_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto = out.File
	file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_rawDesc = nil
	file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_goTypes = nil
	file_cloud_planton_apis_v1_stack_job_enums_operationtype_stackjoboperationtypeoptions_stack_job_operation_type_options_proto_depIdxs = nil
}
