// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/dnszone/service/command.proto

package service

import (
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/dnszone/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcerpcmethodoptions"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model"
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

var File_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x30, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6e,
	0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x72, 0x70, 0x63, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x70, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xd9, 0x11, 0x0a, 0x18, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xc9, 0x01, 0x0a,
	0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x37,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x1a, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f,
	0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65,
	0x22, 0x46, 0xc2, 0xb8, 0x18, 0x3e, 0x08, 0x2a, 0x10, 0x07, 0x1a, 0x0f, 0x73, 0x70, 0x65, 0x63,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x22, 0x27, 0x75, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x64, 0x6e, 0x73, 0x2d,
	0x7a, 0x6f, 0x6e, 0x65, 0xa0, 0xff, 0x2b, 0x0b, 0x12, 0xba, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x1a, 0x37, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e,
	0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x3e, 0xc2, 0xb8, 0x18, 0x36, 0x08, 0x2a, 0x10, 0x07, 0x1a,
	0x0f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64,
	0x22, 0x1f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74,
	0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x64, 0x6e, 0x73, 0x2d, 0x7a, 0x6f, 0x6e,
	0x65, 0xa0, 0xff, 0x2b, 0x01, 0x12, 0xc5, 0x01, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f,
	0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65,
	0x1a, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x42, 0xc2, 0xb8, 0x18, 0x3a, 0x08,
	0x2e, 0x10, 0x09, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64,
	0x22, 0x27, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74,
	0x6f, 0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x20, 0x64, 0x6e, 0x73, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0xa0, 0xff, 0x2b, 0x0b, 0x12, 0xb6, 0x01,
	0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a,
	0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e,
	0x65, 0x1a, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x3a, 0xc2, 0xb8, 0x18, 0x32,
	0x08, 0x2e, 0x10, 0x09, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69,
	0x64, 0x22, 0x1f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x64, 0x6e, 0x73, 0x2d, 0x7a, 0x6f,
	0x6e, 0x65, 0xa0, 0xff, 0x2b, 0x02, 0x12, 0xd9, 0x01, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x42,
	0xc2, 0xb8, 0x18, 0x3a, 0x08, 0x2b, 0x10, 0x09, 0x1a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x27, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x20, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x64, 0x6e, 0x73, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0xa0, 0xff,
	0x2b, 0x0e, 0x12, 0xca, 0x01, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x4b, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e,
	0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a,
	0x6f, 0x6e, 0x65, 0x22, 0x3a, 0xc2, 0xb8, 0x18, 0x32, 0x08, 0x2b, 0x10, 0x09, 0x1a, 0x0b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x1f, 0x75, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x20, 0x64, 0x6e, 0x73, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0xa0, 0xff, 0x2b, 0x03, 0x12,
	0xc7, 0x01, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x1a, 0x37, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64,
	0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73,
	0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x43, 0xc2, 0xb8, 0x18, 0x3b, 0x08, 0x2d, 0x10, 0x09, 0x1a, 0x0b,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x28, 0x75, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x20, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x64, 0x6e, 0x73,
	0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0xa0, 0xff, 0x2b, 0x0b, 0x12, 0xb8, 0x01, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x1a, 0x37,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x3b, 0xc2, 0xb8, 0x18, 0x33, 0x08, 0x2d, 0x10,
	0x09, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x20,
	0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20,
	0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x64, 0x6e, 0x73, 0x2d, 0x7a, 0x6f, 0x6e, 0x65,
	0xa0, 0xff, 0x2b, 0x04, 0x12, 0xd5, 0x01, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f,
	0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x37, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44,
	0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x44, 0xc2, 0xb8, 0x18, 0x40, 0x08, 0x2e, 0x10, 0x09,
	0x1a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x2d, 0x75,
	0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2d, 0x6a, 0x6f, 0x62, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x64, 0x6e, 0x73, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0xdc, 0x01, 0x0a,
	0x0e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12,
	0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x37, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44,
	0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x43, 0xc2, 0xb8, 0x18, 0x3b, 0x08, 0x2e, 0x10, 0x09,
	0x1a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x28, 0x75,
	0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x64,
	0x6e, 0x73, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0xa0, 0xff, 0x2b, 0x0b, 0x12, 0xcd, 0x01, 0x0a, 0x07,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x3b,
	0xc2, 0xb8, 0x18, 0x33, 0x08, 0x2e, 0x10, 0x09, 0x1a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x64,
	0x6e, 0x73, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0xa0, 0xff, 0x2b, 0x05, 0x32, 0x81, 0x05, 0x0a, 0x1a,
	0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xc7, 0x01, 0x0a, 0x03, 0x61,
	0x64, 0x64, 0x12, 0x50, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x35, 0xc2,
	0xb8, 0x18, 0x31, 0x08, 0x2e, 0x10, 0x09, 0x1a, 0x0b, 0x64, 0x6e, 0x73, 0x5f, 0x7a, 0x6f, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x22, 0x1e, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x64, 0x64, 0x20, 0x64, 0x6e, 0x73, 0x2d, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0xcd, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x50, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6e, 0x73, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x38, 0xc2, 0xb8, 0x18, 0x34,
	0x08, 0x2e, 0x10, 0x09, 0x1a, 0x0b, 0x64, 0x6e, 0x73, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x22, 0x21, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x64, 0x6e, 0x73, 0x2d, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0xc8, 0x01, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x37, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x6e,
	0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x38, 0xc2, 0xb8, 0x18, 0x34, 0x08, 0x2e, 0x10, 0x09, 0x1a,
	0x0b, 0x64, 0x6e, 0x73, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x21, 0x75, 0x6e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x20, 0x64, 0x6e, 0x73, 0x2d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42,
	0x9a, 0x03, 0x0a, 0x3e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65,
	0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x60, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a,
	0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x43, 0x56, 0x44, 0x53, 0xaa, 0x02,
	0x30, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41,
	0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56,
	0x31, 0x2e, 0x44, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0xca, 0x02, 0x30, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x56, 0x31, 0x5c, 0x44, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x3c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x44, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x36, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x44, 0x6e, 0x73, 0x7a,
	0x6f, 0x6e, 0x65, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto_goTypes = []interface{}{
	(*model.DnsZone)(nil),                          // 0: cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	(*model1.ApiResourceDeleteCommandInput)(nil),   // 1: cloud.planton.apis.commons.apiresource.model.ApiResourceDeleteCommandInput
	(*model2.CreateStackJobCommandInput)(nil),      // 2: cloud.planton.apis.iac.v1.stackjob.model.CreateStackJobCommandInput
	(*model1.ApiResourceRefreshCommandInput)(nil),  // 3: cloud.planton.apis.commons.apiresource.model.ApiResourceRefreshCommandInput
	(*model.AddOrUpdateDnsRecordCommandInput)(nil), // 4: cloud.planton.apis.code2cloud.v1.dnszone.model.AddOrUpdateDnsRecordCommandInput
	(*model.DeleteDnsRecordCommandInput)(nil),      // 5: cloud.planton.apis.code2cloud.v1.dnszone.model.DeleteDnsRecordCommandInput
}
var file_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto_depIdxs = []int32{
	0,  // 0: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.previewCreate:input_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 1: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.create:input_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 2: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.previewUpdate:input_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 3: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.update:input_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	1,  // 4: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.previewDelete:input_type -> cloud.planton.apis.commons.apiresource.model.ApiResourceDeleteCommandInput
	1,  // 5: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.delete:input_type -> cloud.planton.apis.commons.apiresource.model.ApiResourceDeleteCommandInput
	0,  // 6: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.previewRestore:input_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 7: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.restore:input_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	2,  // 8: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.createStackJob:input_type -> cloud.planton.apis.iac.v1.stackjob.model.CreateStackJobCommandInput
	3,  // 9: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.previewRefresh:input_type -> cloud.planton.apis.commons.apiresource.model.ApiResourceRefreshCommandInput
	3,  // 10: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.refresh:input_type -> cloud.planton.apis.commons.apiresource.model.ApiResourceRefreshCommandInput
	4,  // 11: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsRecordCommandController.add:input_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.AddOrUpdateDnsRecordCommandInput
	4,  // 12: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsRecordCommandController.update:input_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.AddOrUpdateDnsRecordCommandInput
	5,  // 13: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsRecordCommandController.delete:input_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DeleteDnsRecordCommandInput
	0,  // 14: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.previewCreate:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 15: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.create:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 16: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.previewUpdate:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 17: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.update:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 18: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.previewDelete:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 19: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.delete:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 20: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.previewRestore:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 21: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.restore:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 22: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.createStackJob:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 23: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.previewRefresh:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 24: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsZoneCommandController.refresh:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 25: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsRecordCommandController.add:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 26: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsRecordCommandController.update:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	0,  // 27: cloud.planton.apis.code2cloud.v1.dnszone.service.DnsRecordCommandController.delete:output_type -> cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_dnszone_service_command_proto_depIdxs = nil
}