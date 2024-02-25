// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/iac/v1/stackjob/service/query.proto

package service

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model"
	progress "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/progress"
	pulumiengine "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/progress/pulumiengine"
	snapshot "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model/progress/snapshot"
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

var File_cloud_planton_apis_iac_v1_stackjob_service_query_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_iac_v1_stackjob_service_query_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a,
	0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x47, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f,
	0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2f, 0x69,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8c, 0x09,
	0x0a, 0x17, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x93, 0x01, 0x0a, 0x0d, 0x6c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4a, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x4a, 0x6f, 0x62, 0x73, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x73, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x34, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64,
	0x1a, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x9a, 0x01, 0x0a, 0x16, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x34, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x4a, 0x6f, 0x62, 0x49, 0x64, 0x1a, 0x48, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a,
	0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x30,
	0x01, 0x12, 0xb1, 0x01, 0x0a, 0x21, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f,
	0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x34, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x1a, 0x54, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a,
	0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x30, 0x01, 0x12, 0xae, 0x01, 0x0a, 0x20, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x4d, 0x54, 0x44, 0x42,
	0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x4c, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f,
	0x62, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x4d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x4d, 0x54, 0x44, 0x12, 0xb9, 0x01, 0x0a, 0x21, 0x67, 0x65, 0x74, 0x50, 0x75,
	0x6c, 0x75, 0x6d, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x50, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f,
	0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6c, 0x75, 0x6d,
	0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x42,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x75, 0x6c, 0x75, 0x6d, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0xa7, 0x01, 0x0a, 0x1e, 0x67, 0x65, 0x74, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x79, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x4a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x34, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x1a, 0x4f, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x75, 0x6c,
	0x75, 0x6d, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0xf4, 0x02, 0x0a,
	0x38, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f,
	0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x49, 0x56, 0x53, 0x53, 0xaa, 0x02, 0x2a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70,
	0x69, 0x73, 0x2e, 0x49, 0x61, 0x63, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x6a,
	0x6f, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x2a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x49, 0x61, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x36, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x49, 0x61, 0x63,
	0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x30, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x49, 0x61, 0x63, 0x3a, 0x3a, 0x56, 0x31,
	0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_iac_v1_stackjob_service_query_proto_goTypes = []interface{}{
	(*model.ListStackJobsByFiltersQueryInput)(nil),       // 0: cloud.planton.apis.iac.v1.stackjob.model.ListStackJobsByFiltersQueryInput
	(*model.StackJobId)(nil),                             // 1: cloud.planton.apis.iac.v1.stackjob.model.StackJobId
	(*model.GetStackJobMinutesByCompanyIdInput)(nil),     // 2: cloud.planton.apis.iac.v1.stackjob.model.GetStackJobMinutesByCompanyIdInput
	(*model.GetPulumiResourceCountByCompanyIdInput)(nil), // 3: cloud.planton.apis.iac.v1.stackjob.model.GetPulumiResourceCountByCompanyIdInput
	(*model.StackJobList)(nil),                           // 4: cloud.planton.apis.iac.v1.stackjob.model.StackJobList
	(*model.StackJob)(nil),                               // 5: cloud.planton.apis.iac.v1.stackjob.model.StackJob
	(*progress.StackJobProgressEvent)(nil),               // 6: cloud.planton.apis.iac.v1.stackjob.model.progress.StackJobProgressEvent
	(*snapshot.StackJobProgressSnapshot)(nil),            // 7: cloud.planton.apis.iac.v1.stackjob.model.progress.snapshot.StackJobProgressSnapshot
	(*model.StackJobMinutesMTD)(nil),                     // 8: cloud.planton.apis.iac.v1.stackjob.model.StackJobMinutesMTD
	(*model.TotalPulumiResourceCount)(nil),               // 9: cloud.planton.apis.iac.v1.stackjob.model.TotalPulumiResourceCount
	(*pulumiengine.PulumiResources)(nil),                 // 10: cloud.planton.apis.iac.v1.stackjob.model.progress.pulumiengine.PulumiResources
}
var file_cloud_planton_apis_iac_v1_stackjob_service_query_proto_depIdxs = []int32{
	0,  // 0: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.listByFilters:input_type -> cloud.planton.apis.iac.v1.stackjob.model.ListStackJobsByFiltersQueryInput
	1,  // 1: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.getById:input_type -> cloud.planton.apis.iac.v1.stackjob.model.StackJobId
	1,  // 2: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.getProgressEventStream:input_type -> cloud.planton.apis.iac.v1.stackjob.model.StackJobId
	1,  // 3: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.getStackJobProgressSnapshotStream:input_type -> cloud.planton.apis.iac.v1.stackjob.model.StackJobId
	2,  // 4: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.getStackJobMinutesMTDByCompanyId:input_type -> cloud.planton.apis.iac.v1.stackjob.model.GetStackJobMinutesByCompanyIdInput
	3,  // 5: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.getPulumiResourceCountByCompanyId:input_type -> cloud.planton.apis.iac.v1.stackjob.model.GetPulumiResourceCountByCompanyIdInput
	1,  // 6: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.getPulumiResourcesByStackJobId:input_type -> cloud.planton.apis.iac.v1.stackjob.model.StackJobId
	4,  // 7: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.listByFilters:output_type -> cloud.planton.apis.iac.v1.stackjob.model.StackJobList
	5,  // 8: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.getById:output_type -> cloud.planton.apis.iac.v1.stackjob.model.StackJob
	6,  // 9: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.getProgressEventStream:output_type -> cloud.planton.apis.iac.v1.stackjob.model.progress.StackJobProgressEvent
	7,  // 10: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.getStackJobProgressSnapshotStream:output_type -> cloud.planton.apis.iac.v1.stackjob.model.progress.snapshot.StackJobProgressSnapshot
	8,  // 11: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.getStackJobMinutesMTDByCompanyId:output_type -> cloud.planton.apis.iac.v1.stackjob.model.StackJobMinutesMTD
	9,  // 12: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.getPulumiResourceCountByCompanyId:output_type -> cloud.planton.apis.iac.v1.stackjob.model.TotalPulumiResourceCount
	10, // 13: cloud.planton.apis.iac.v1.stackjob.service.StackJobQueryController.getPulumiResourcesByStackJobId:output_type -> cloud.planton.apis.iac.v1.stackjob.model.progress.pulumiengine.PulumiResources
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_iac_v1_stackjob_service_query_proto_init() }
func file_cloud_planton_apis_iac_v1_stackjob_service_query_proto_init() {
	if File_cloud_planton_apis_iac_v1_stackjob_service_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_iac_v1_stackjob_service_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_planton_apis_iac_v1_stackjob_service_query_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_iac_v1_stackjob_service_query_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_iac_v1_stackjob_service_query_proto = out.File
	file_cloud_planton_apis_iac_v1_stackjob_service_query_proto_rawDesc = nil
	file_cloud_planton_apis_iac_v1_stackjob_service_query_proto_goTypes = nil
	file_cloud_planton_apis_iac_v1_stackjob_service_query_proto_depIdxs = nil
}
