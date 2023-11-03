// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/redis/service.proto

package redis

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

var File_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x1a, 0x38, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70,
	0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfa, 0x0a, 0x0a, 0x1d, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xd8, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a,
	0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e,
	0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x54, 0xc2, 0xb8,
	0x18, 0x50, 0x08, 0x60, 0x10, 0x09, 0x1a, 0x24, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x24, 0x75, 0x6e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x20, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0xbf, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x3b, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x3b, 0xc2, 0xb8, 0x18, 0x37, 0x08, 0x64, 0x10,
	0x15, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x24,
	0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0xbb, 0x01, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x3d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e,
	0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x3b,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52,
	0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x35, 0xc2, 0xb8, 0x18,
	0x31, 0x08, 0x61, 0x10, 0x15, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x24, 0x75, 0x6e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x20, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0xc1, 0x01, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3b,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52,
	0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x3b, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69,
	0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x3c, 0xc2, 0xb8, 0x18, 0x38, 0x08, 0x63,
	0x10, 0x15, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22,
	0x25, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0xbd, 0x01, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x3d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x1a, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x36,
	0xc2, 0xb8, 0x18, 0x32, 0x08, 0x64, 0x10, 0x15, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x25, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x20, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0xb9, 0x01, 0x0a, 0x05, 0x70, 0x61, 0x75, 0x73, 0x65,
	0x12, 0x3d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73,
	0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a,
	0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e,
	0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x34, 0xc2, 0xb8,
	0x18, 0x30, 0x08, 0x64, 0x10, 0x15, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x75,
	0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70,
	0x61, 0x75, 0x73, 0x65, 0x20, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0xbd, 0x01, 0x0a, 0x07, 0x75, 0x6e, 0x70, 0x61, 0x75, 0x73, 0x65, 0x12, 0x3d,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52,
	0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x3b, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x36, 0xc2, 0xb8, 0x18, 0x32,
	0x08, 0x64, 0x10, 0x15, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x25, 0x75, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x6e, 0x70,
	0x61, 0x75, 0x73, 0x65, 0x20, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x32, 0xa5, 0x09, 0x0a, 0x1b, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0x7b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x3f,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52,
	0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0xba, 0x01, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x3d, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69,
	0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x3b, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x33, 0xc2, 0xb8, 0x18, 0x2f, 0x08, 0x62, 0x10,
	0x15, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x72,
	0x65, 0x64, 0x69, 0x73, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x89, 0x01, 0x0a,
	0x0f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x38, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x3c, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x90, 0x01, 0x0a, 0x13, 0x66, 0x69, 0x6e,
	0x64, 0x42, 0x79, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x3c, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x97, 0x01, 0x0a, 0x13,
	0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4b, 0x75, 0x62, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x42, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0xcf, 0x01, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x3d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x1a, 0x43, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72,
	0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3c, 0xc2, 0xb8, 0x18, 0x38, 0x08,
	0x64, 0x10, 0x15, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2b, 0x75, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77,
	0x20, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x20, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0xc0, 0x01, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x64,
	0x50, 0x6f, 0x64, 0x73, 0x12, 0x3d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72,
	0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x1a, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x64, 0x73,
	0x22, 0x38, 0xc2, 0xb8, 0x18, 0x34, 0x08, 0x62, 0x10, 0x15, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x27, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x20, 0x70, 0x6f, 0x64, 0x73, 0x32, 0xa9, 0x03, 0x0a, 0x1b, 0x52,
	0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xc7, 0x01, 0x0a, 0x07, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x1a, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65,
	0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x22, 0x42, 0xc2, 0xb8, 0x18, 0x3e, 0x08, 0x64, 0x10, 0x15, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x2b, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x20, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x20, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x12, 0xbf, 0x01, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x3d,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52,
	0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x3b, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x3a, 0xc2, 0xb8, 0x18, 0x36,
	0x08, 0x64, 0x10, 0x15, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x29, 0x75, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x20, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x20, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x42, 0x88, 0x03, 0x0a, 0x3b, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f,
	0x72, 0x65, 0x64, 0x69, 0x73, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x52, 0xaa,
	0x02, 0x2d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0xca,
	0x02, 0x2d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x52, 0x65, 0x64, 0x69, 0x73, 0xe2,
	0x02, 0x39, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x52, 0x65, 0x64, 0x69, 0x73, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x33, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a, 0x52, 0x65, 0x64, 0x69,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto_goTypes = []interface{}{
	(*RedisCluster)(nil),              // 0: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	(*RedisClusterId)(nil),            // 1: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterId
	(*pagination.PageInfo)(nil),       // 2: cloud.planton.apis.v1.commons.pagination.PageInfo
	(*product.ProductId)(nil),         // 3: cloud.planton.apis.v1.resourcemanager.product.ProductId
	(*environment.EnvironmentId)(nil), // 4: cloud.planton.apis.v1.code2cloud.environment.EnvironmentId
	(*kubecluster.KubeClusterId)(nil), // 5: cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterId
	(*RedisClusterList)(nil),          // 6: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterList
	(*RedisClusters)(nil),             // 7: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusters
	(*RedisClusterPassword)(nil),      // 8: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterPassword
	(*resource.Pods)(nil),             // 9: cloud.planton.apis.v1.integration.kubernetes.resource.Pods
}
var file_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto_depIdxs = []int32{
	0,  // 0: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.create:input_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	0,  // 1: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.update:input_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	1,  // 2: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.delete:input_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterId
	0,  // 3: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.restore:input_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	1,  // 4: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.restart:input_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterId
	1,  // 5: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.pause:input_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterId
	1,  // 6: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.unpause:input_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterId
	2,  // 7: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.list:input_type -> cloud.planton.apis.v1.commons.pagination.PageInfo
	1,  // 8: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.getById:input_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterId
	3,  // 9: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.findByProductId:input_type -> cloud.planton.apis.v1.resourcemanager.product.ProductId
	4,  // 10: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.findByEnvironmentId:input_type -> cloud.planton.apis.v1.code2cloud.environment.EnvironmentId
	5,  // 11: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.findByKubeClusterId:input_type -> cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterId
	1,  // 12: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.getPassword:input_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterId
	1,  // 13: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.findPods:input_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterId
	0,  // 14: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterStackController.preview:input_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	1,  // 15: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterStackController.apply:input_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterId
	0,  // 16: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.create:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	0,  // 17: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.update:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	0,  // 18: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.delete:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	0,  // 19: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.restore:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	0,  // 20: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.restart:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	0,  // 21: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.pause:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	0,  // 22: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterCommandController.unpause:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	6,  // 23: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.list:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterList
	0,  // 24: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.getById:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	7,  // 25: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.findByProductId:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusters
	7,  // 26: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.findByEnvironmentId:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusters
	7,  // 27: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.findByKubeClusterId:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusters
	8,  // 28: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.getPassword:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterPassword
	9,  // 29: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterQueryController.findPods:output_type -> cloud.planton.apis.v1.integration.kubernetes.resource.Pods
	0,  // 30: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterStackController.preview:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	0,  // 31: cloud.planton.apis.v1.code2cloud.deploy.redis.RedisClusterStackController.apply:output_type -> cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster
	16, // [16:32] is the sub-list for method output_type
	0,  // [0:16] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto != nil {
		return
	}
	file_cloud_planton_apis_v1_code2cloud_deploy_redis_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_redis_service_proto_depIdxs = nil
}
