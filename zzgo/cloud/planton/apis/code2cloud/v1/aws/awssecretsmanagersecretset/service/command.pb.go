// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/aws/awssecretsmanagersecretset/service/command.proto

package service

import (
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/awssecretsmanagersecretset/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcerpcmethodoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcerpcserviceoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iam/v1/iampolicy/options/rpcauthorizationmethodoptions"
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

var File_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto_rawDesc = []byte{
	0x0a, 0x55, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65,
	0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x47, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77,
	0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x4e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65,
	0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x51, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65,
	0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x72, 0x70, 0x63, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x70,
	0x63, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x61, 0x6d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x72, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xf4, 0x17, 0x0a, 0x2b, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12,
	0xc8, 0x02, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x53, 0x65, 0x74, 0x1a, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77, 0x73,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x22, 0x71, 0xc2, 0xb8, 0x18, 0x61, 0x08, 0x01, 0x10,
	0x0a, 0x1a, 0x1c, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x69, 0x64, 0x22,
	0x3d, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20,
	0x61, 0x77, 0x73, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x74, 0xa0, 0xff,
	0x2b, 0x0b, 0xb0, 0xff, 0x2b, 0x01, 0xb8, 0xff, 0x2b, 0x01, 0x12, 0xb9, 0x02, 0x0a, 0x06, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77,
	0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x1a, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e,
	0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x22, 0x69, 0xc2, 0xb8, 0x18,
	0x59, 0x08, 0x01, 0x10, 0x0a, 0x1a, 0x1c, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x5f, 0x69, 0x64, 0x22, 0x35, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x77, 0x73, 0x2d,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x74, 0xa0, 0xff, 0x2b, 0x01, 0xb0, 0xff,
	0x2b, 0x01, 0xb8, 0xff, 0x2b, 0x01, 0x12, 0xb7, 0x02, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e,
	0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x1a, 0x61, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61,
	0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x22, 0x60,
	0xc2, 0xb8, 0x18, 0x50, 0x08, 0x05, 0x10, 0x05, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x3d, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x20, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x77, 0x73, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x2d, 0x73, 0x65, 0x74, 0xa0, 0xff, 0x2b, 0x0b, 0xb0, 0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01,
	0x12, 0xa8, 0x02, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x61, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61,
	0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x1a, 0x61,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65,
	0x74, 0x22, 0x58, 0xc2, 0xb8, 0x18, 0x48, 0x08, 0x05, 0x10, 0x05, 0x1a, 0x0b, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x35, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x20, 0x61, 0x77, 0x73, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2d, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x74, 0xa0,
	0xff, 0x2b, 0x02, 0xb0, 0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0x12, 0x9e, 0x02, 0x0a, 0x0d,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x44, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77, 0x73, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x22, 0x64, 0xc2, 0xb8, 0x18, 0x50, 0x08, 0x02, 0x10, 0x05,
	0x1a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x3d, 0x75,
	0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x77,
	0x73, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x74, 0xa0, 0xff, 0x2b, 0x0e,
	0xb0, 0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0xc0, 0xff, 0x2b, 0x01, 0x12, 0x8f, 0x02, 0x0a,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x61, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74,
	0x22, 0x5c, 0xc2, 0xb8, 0x18, 0x48, 0x08, 0x02, 0x10, 0x05, 0x1a, 0x0b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x35, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x61, 0x77, 0x73, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x74, 0xa0, 0xff,
	0x2b, 0x03, 0xb0, 0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0xc0, 0xff, 0x2b, 0x01, 0x12, 0xb9,
	0x02, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x53, 0x65, 0x74, 0x1a, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77, 0x73,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x22, 0x61, 0xc2, 0xb8, 0x18, 0x51, 0x08, 0x04, 0x10,
	0x05, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x3e,
	0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x20, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20,
	0x61, 0x77, 0x73, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x74, 0xa0, 0xff,
	0x2b, 0x0b, 0xb0, 0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0x12, 0xaa, 0x02, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77,
	0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x1a, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73,
	0x2e, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x22, 0x59, 0xc2, 0xb8,
	0x18, 0x49, 0x08, 0x04, 0x10, 0x05, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x69, 0x64, 0x22, 0x36, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x77, 0x73,
	0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x74, 0xa0, 0xff, 0x2b, 0x04, 0xb0,
	0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0x12, 0xa1, 0x02, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x45, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x53, 0x65, 0x74, 0x22, 0x65, 0xc2, 0xb8, 0x18, 0x51, 0x08, 0x05, 0x10, 0x05, 0x1a, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x3e, 0x75, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x61, 0x77, 0x73,
	0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x74, 0xa0, 0xff, 0x2b, 0x0b, 0xb0,
	0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0xc0, 0xff, 0x2b, 0x01, 0x12, 0x92, 0x02, 0x0a, 0x07,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x45, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x61,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65,
	0x74, 0x22, 0x5d, 0xc2, 0xb8, 0x18, 0x49, 0x08, 0x05, 0x10, 0x05, 0x1a, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x36, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x20, 0x61, 0x77, 0x73, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2d, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x74,
	0xa0, 0xff, 0x2b, 0x05, 0xb0, 0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0xc0, 0xff, 0x2b, 0x01,
	0x1a, 0x04, 0xa0, 0xff, 0x2b, 0x05, 0x42, 0xa6, 0x04, 0x0a, 0x55, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77,
	0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x77, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65,
	0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x43,
	0x56, 0x41, 0x41, 0x53, 0xaa, 0x02, 0x47, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x77, 0x73, 0x2e, 0x41, 0x77, 0x73, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xca, 0x02,
	0x47, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56,
	0x31, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x41, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x53, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x41, 0x77, 0x73, 0x5c,
	0x41, 0x77, 0x73, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x41, 0x77, 0x73, 0x3a, 0x3a, 0x41, 0x77, 0x73, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x65, 0x74, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto_goTypes = []interface{}{
	(*model.AwsSecretsManagerSecretSet)(nil), // 0: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	(*model1.ApiResourceDeleteInput)(nil),    // 1: cloud.planton.apis.commons.apiresource.model.ApiResourceDeleteInput
	(*model1.ApiResourceRefreshInput)(nil),   // 2: cloud.planton.apis.commons.apiresource.model.ApiResourceRefreshInput
}
var file_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto_depIdxs = []int32{
	0,  // 0: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.previewCreate:input_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 1: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.create:input_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 2: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.previewUpdate:input_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 3: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.update:input_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	1,  // 4: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.previewDelete:input_type -> cloud.planton.apis.commons.apiresource.model.ApiResourceDeleteInput
	1,  // 5: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.delete:input_type -> cloud.planton.apis.commons.apiresource.model.ApiResourceDeleteInput
	0,  // 6: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.previewRestore:input_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 7: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.restore:input_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	2,  // 8: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.previewRefresh:input_type -> cloud.planton.apis.commons.apiresource.model.ApiResourceRefreshInput
	2,  // 9: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.refresh:input_type -> cloud.planton.apis.commons.apiresource.model.ApiResourceRefreshInput
	0,  // 10: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.previewCreate:output_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 11: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.create:output_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 12: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.previewUpdate:output_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 13: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.update:output_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 14: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.previewDelete:output_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 15: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.delete:output_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 16: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.previewRestore:output_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 17: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.restore:output_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 18: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.previewRefresh:output_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	0,  // 19: cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.service.AwsSecretsManagerSecretSetCommandController.refresh:output_type -> cloud.planton.apis.code2cloud.v1.aws.awssecretsmanagersecretset.model.AwsSecretsManagerSecretSet
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto_init()
}
func file_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_aws_awssecretsmanagersecretset_service_command_proto_depIdxs = nil
}
