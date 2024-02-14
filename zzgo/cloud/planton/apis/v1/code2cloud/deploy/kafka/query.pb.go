// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/kafka/query.proto

package kafka

import (
	kubecluster "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/kubecluster"
	environment "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/environment"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/extensions"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/integration/kubernetes/resource"
	product "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/product"
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

var File_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto_rawDesc = []byte{
	0x0a, 0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x1a, 0x36, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x69, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa5, 0x09, 0x0a, 0x1b, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x7b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x3f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0xba, 0x01, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x3d,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4b,
	0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x3b, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x33, 0xc2, 0xb8, 0x18, 0x2f,
	0x08, 0x4d, 0x10, 0x0f, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x75, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x69, 0x65,
	0x77, 0x20, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x89, 0x01, 0x0a, 0x0f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x38, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x3c, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x90, 0x01, 0x0a, 0x13,
	0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x1a, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x97,
	0x01, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4b, 0x75, 0x62, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x42, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4b, 0x75, 0x62,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x3c, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0xcf, 0x01, 0x0a, 0x0b, 0x67, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x3d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x43, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3c, 0xc2, 0xb8,
	0x18, 0x38, 0x08, 0x4f, 0x10, 0x0f, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2b, 0x75,
	0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76,
	0x69, 0x65, 0x77, 0x20, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0xc0, 0x01, 0x0a, 0x08, 0x66,
	0x69, 0x6e, 0x64, 0x50, 0x6f, 0x64, 0x73, 0x12, 0x3d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x50,
	0x6f, 0x64, 0x73, 0x22, 0x38, 0xc2, 0xb8, 0x18, 0x34, 0x08, 0x4d, 0x10, 0x0f, 0x1a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x27, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x20, 0x70, 0x6f, 0x64, 0x73, 0x32, 0x9c, 0x03,
	0x0a, 0x19, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xc5, 0x01, 0x0a, 0x14,
	0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x1a, 0x3a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22,
	0x32, 0xc2, 0xb8, 0x18, 0x2e, 0x08, 0x4d, 0x10, 0x0f, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x21, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74,
	0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x20, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x12, 0xb6, 0x01, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x43, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e,
	0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x22,
	0x2b, 0xc2, 0xb8, 0x18, 0x27, 0x08, 0x4d, 0x10, 0x0f, 0x22, 0x21, 0x75, 0x6e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x20, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x42, 0x86, 0x03, 0x0a,
	0x3b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x42, 0x0a, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x56,
	0x43, 0x44, 0x4b, 0xaa, 0x02, 0x2d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0xca, 0x02, 0x2d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0xe2, 0x02, 0x39, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x33, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a,
	0x4b, 0x61, 0x66, 0x6b, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto_goTypes = []interface{}{
	(*pagination.PageInfo)(nil),       // 0: cloud.planton.apis.v1.commons.pagination.PageInfo
	(*KafkaClusterId)(nil),            // 1: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterId
	(*product.ProductId)(nil),         // 2: cloud.planton.apis.v1.resourcemanager.product.ProductId
	(*environment.EnvironmentId)(nil), // 3: cloud.planton.apis.v1.code2cloud.environment.EnvironmentId
	(*kubecluster.KubeClusterId)(nil), // 4: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterId
	(*KafkaTopicQueryInput)(nil),      // 5: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaTopicQueryInput
	(*KafkaClusterList)(nil),          // 6: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterList
	(*KafkaCluster)(nil),              // 7: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaCluster
	(*KafkaClusters)(nil),             // 8: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusters
	(*KafkaClusterPassword)(nil),      // 9: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterPassword
	(*resource.Pods)(nil),             // 10: cloud.planton.apis.v1.integration.kubernetes.resource.Pods
	(*KafkaTopics)(nil),               // 11: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaTopics
	(*KafkaTopic)(nil),                // 12: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaTopic
}
var file_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto_depIdxs = []int32{
	0,  // 0: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.list:input_type -> cloud.planton.apis.v1.commons.pagination.PageInfo
	1,  // 1: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.getById:input_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterId
	2,  // 2: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.findByProductId:input_type -> cloud.planton.apis.v1.resourcemanager.product.ProductId
	3,  // 3: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.findByEnvironmentId:input_type -> cloud.planton.apis.v1.code2cloud.environment.EnvironmentId
	4,  // 4: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.findByKubeClusterId:input_type -> cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterId
	1,  // 5: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.getPassword:input_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterId
	1,  // 6: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.findPods:input_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterId
	1,  // 7: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaTopicQueryController.findByKafkaClusterId:input_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterId
	5,  // 8: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaTopicQueryController.getById:input_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaTopicQueryInput
	6,  // 9: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.list:output_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterList
	7,  // 10: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.getById:output_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaCluster
	8,  // 11: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.findByProductId:output_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusters
	8,  // 12: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.findByEnvironmentId:output_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusters
	8,  // 13: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.findByKubeClusterId:output_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusters
	9,  // 14: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.getPassword:output_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterPassword
	10, // 15: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaClusterQueryController.findPods:output_type -> cloud.planton.apis.v1.integration.kubernetes.resource.Pods
	11, // 16: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaTopicQueryController.findByKafkaClusterId:output_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaTopics
	12, // 17: cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaTopicQueryController.getById:output_type -> cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaTopic
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto != nil {
		return
	}
	file_cloud_planton_apis_v1_code2cloud_deploy_kafka_io_proto_init()
	file_cloud_planton_apis_v1_code2cloud_deploy_kafka_state_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_kafka_query_proto_depIdxs = nil
}