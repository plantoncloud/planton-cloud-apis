// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/customendpoint/service/query.proto

package service

import (
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/customendpoint/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/extensions"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/integration/kubernetes/resource"
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

var File_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x45, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x48, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x47, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xef,
	0x0a, 0x0a, 0x1d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x96, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x50, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0xde, 0x01, 0x0a, 0x07, 0x67, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x22, 0x35, 0xc2, 0xb8, 0x18, 0x31, 0x08, 0x26, 0x10, 0x07, 0x1a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x24, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0xa0, 0x01, 0x0a, 0x0f, 0x66,
	0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x3e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x4d,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x90, 0x02,
	0x0a, 0x1b, 0x67, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x56, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x49, 0xc2, 0xb8, 0x18, 0x45, 0x08, 0x26, 0x10, 0x07, 0x1a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x9e, 0x02, 0x0a, 0x24, 0x67, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x73, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x5f, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x45, 0xc2, 0xb8, 0x18, 0x41,
	0x08, 0x26, 0x10, 0x07, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x34, 0x75, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x20, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0xfc, 0x01, 0x0a, 0x1e, 0x66, 0x69, 0x6e, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x49, 0x64, 0x1a, 0x43, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x22, 0x45, 0xc2, 0xb8, 0x18, 0x41, 0x08,
	0x26, 0x10, 0x07, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x34, 0x75, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x6c, 0x6f, 0x6f, 0x6b,
	0x20, 0x75, 0x70, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x20, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x42, 0xee, 0x03, 0x0a, 0x4c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67,
	0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x6e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2,
	0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x43, 0x53, 0xaa, 0x02, 0x3e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e,
	0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x3e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x4a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69,
	0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x45, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto_goTypes = []interface{}{
	(*model.PageInfo)(nil),                           // 0: cloud.planton.apis.v1.commons.rpc.pagination.model.PageInfo
	(*model1.CustomEndpointId)(nil),                  // 1: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpointId
	(*model2.ProductId)(nil),                         // 2: cloud.planton.apis.v1.resourcemanager.product.model.ProductId
	(*model1.CustomEndpointList)(nil),                // 3: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpointList
	(*model1.CustomEndpoint)(nil),                    // 4: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpoint
	(*model1.CustomEndpoints)(nil),                   // 5: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpoints
	(*model1.CustomEndpointCertStatus)(nil),          // 6: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpointCertStatus
	(*model1.CustomEndpointDnsResolutionStatus)(nil), // 7: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpointDnsResolutionStatus
	(*resource.Certificates)(nil),                    // 8: cloud.planton.apis.v1.integration.kubernetes.resource.Certificates
}
var file_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.service.CustomEndpointQueryController.list:input_type -> cloud.planton.apis.v1.commons.rpc.pagination.model.PageInfo
	1, // 1: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.service.CustomEndpointQueryController.getById:input_type -> cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpointId
	2, // 2: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.service.CustomEndpointQueryController.findByProductId:input_type -> cloud.planton.apis.v1.resourcemanager.product.model.ProductId
	1, // 3: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.service.CustomEndpointQueryController.getCustomEndpointCertStatus:input_type -> cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpointId
	1, // 4: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.service.CustomEndpointQueryController.getCustomEndpointDsnResolutionStatus:input_type -> cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpointId
	1, // 5: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.service.CustomEndpointQueryController.findCustomEndpointCertificates:input_type -> cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpointId
	3, // 6: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.service.CustomEndpointQueryController.list:output_type -> cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpointList
	4, // 7: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.service.CustomEndpointQueryController.getById:output_type -> cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpoint
	5, // 8: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.service.CustomEndpointQueryController.findByProductId:output_type -> cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpoints
	6, // 9: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.service.CustomEndpointQueryController.getCustomEndpointCertStatus:output_type -> cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpointCertStatus
	7, // 10: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.service.CustomEndpointQueryController.getCustomEndpointDsnResolutionStatus:output_type -> cloud.planton.apis.v1.code2cloud.deploy.customendpoint.model.CustomEndpointDnsResolutionStatus
	8, // 11: cloud.planton.apis.v1.code2cloud.deploy.customendpoint.service.CustomEndpointQueryController.findCustomEndpointCertificates:output_type -> cloud.planton.apis.v1.integration.kubernetes.resource.Certificates
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_customendpoint_service_query_proto_depIdxs = nil
}
