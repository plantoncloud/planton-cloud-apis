// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/develop/sourcecode/project/service/command.proto

package service

import (
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/develop/sourcecode/project/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/extensions"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/product/model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto_rawDesc = []byte{
	0x0a, 0x51, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xb7, 0x0e, 0x0a, 0x1c, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x12, 0xd4, 0x01, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x12, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2d, 0xc2, 0xb8, 0x18, 0x29, 0x08,
	0x14, 0x10, 0x14, 0x22, 0x23, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x63, 0x6f, 0x64, 0x65,
	0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xd7, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x1a, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x22, 0x2d, 0xc2, 0xb8, 0x18, 0x29, 0x08, 0x14, 0x10, 0x14, 0x22, 0x23, 0x75,
	0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0xe4, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x4e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3a, 0xc2,
	0xb8, 0x18, 0x36, 0x08, 0x18, 0x10, 0x04, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x69, 0x64, 0x22, 0x23, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x63, 0x6f, 0x64,
	0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xe2, 0x01, 0x0a, 0x06, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x3a, 0xc2, 0xb8, 0x18, 0x36, 0x08, 0x16, 0x10, 0x04, 0x1a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x23, 0x75, 0x6e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xe6,
	0x01, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x4e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x4e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3b, 0xc2, 0xb8, 0x18, 0x37,
	0x08, 0x17, 0x10, 0x04, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69,
	0x64, 0x22, 0x24, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x96, 0x01, 0x0a, 0x16, 0x73, 0x79, 0x6e, 0x63,
	0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x3e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x1a, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0xf9, 0x01, 0x0a, 0x1a, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65,
	0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x50, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x1a, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x39, 0xc2, 0xb8, 0x18, 0x35, 0x08, 0x18, 0x10, 0x04, 0x1a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x28, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x20, 0x74, 0x6f, 0x20, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x20,
	0x63, 0x6f, 0x64, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x9b, 0x02, 0x0a,
	0x23, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x72, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x30, 0xc2, 0xb8, 0x18, 0x2c, 0x08, 0x45,
	0x10, 0x0c, 0x22, 0x26, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x20, 0x74, 0x6f, 0x20, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x20, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x90, 0x04, 0x0a, 0x51, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x73, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x53, 0x50,
	0x53, 0xaa, 0x02, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x5c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x4f,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x5c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x4b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x3a,
	0x3a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto_goTypes = []interface{}{
	(*model.CodeProject)(nil),                                     // 0: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProject
	(*model1.ApiResourceRefreshCommandInput)(nil),                 // 1: cloud.planton.apis.v1.commons.resource.model.ApiResourceRefreshCommandInput
	(*model2.ProductId)(nil),                                      // 2: cloud.planton.apis.v1.resourcemanager.product.model.ProductId
	(*model.CodeProjectId)(nil),                                   // 3: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProjectId
	(*model.AttachMachineAccountByCodeProjectIdCommandInput)(nil), // 4: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.AttachMachineAccountByCodeProjectIdCommandInput
	(*model2.Product)(nil),                                        // 5: cloud.planton.apis.v1.resourcemanager.product.model.Product
}
var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.add:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProject
	0, // 1: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.create:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProject
	0, // 2: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.update:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProject
	1, // 3: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.delete:input_type -> cloud.planton.apis.v1.commons.resource.model.ApiResourceRefreshCommandInput
	0, // 4: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.restore:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProject
	2, // 5: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.synchronizeByProductId:input_type -> cloud.planton.apis.v1.resourcemanager.product.model.ProductId
	3, // 6: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.synchronizeByCodeProjectId:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProjectId
	4, // 7: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.attachMachineAccountByCodeProjectId:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.AttachMachineAccountByCodeProjectIdCommandInput
	0, // 8: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.add:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProject
	0, // 9: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.create:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProject
	0, // 10: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.update:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProject
	0, // 11: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.delete:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProject
	0, // 12: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.restore:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProject
	5, // 13: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.synchronizeByProductId:output_type -> cloud.planton.apis.v1.resourcemanager.product.model.Product
	0, // 14: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.synchronizeByCodeProjectId:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProject
	0, // 15: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.service.CodeProjectCommandController.attachMachineAccountByCodeProjectId:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.model.CodeProject
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto_init()
}
func file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_service_command_proto_depIdxs = nil
}
