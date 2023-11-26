// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/storagebucket/service.proto

package storagebucket

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/kubecluster"
	environment "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/environment"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/method/options"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/extensions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/integration/kubernetes/resource"
	product "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/product"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job"
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

var File_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto_rawDesc = []byte{
	0x0a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x3f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x49, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
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
	0x72, 0x63, 0x65, 0x2f, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8d, 0x07, 0x0a, 0x1e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xed, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22,
	0x57, 0xc2, 0xb8, 0x18, 0x53, 0x08, 0x6a, 0x10, 0x09, 0x1a, 0x24, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22,
	0x27, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73,
	0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0xd2, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22,
	0x3c, 0xc2, 0xb8, 0x18, 0x38, 0x08, 0x6e, 0x10, 0x18, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x25, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0xce, 0x01,
	0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x46, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64,
	0x1a, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x36, 0xc2, 0xb8, 0x18, 0x32, 0x08, 0x6b, 0x10, 0x18,
	0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x25, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0xd4,
	0x01, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x44, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x1a, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x3d, 0xc2, 0xb8, 0x18, 0x39, 0x08, 0x6d, 0x10, 0x18,
	0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x26, 0x75,
	0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x72,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2d, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x32, 0xa6, 0x05, 0x0a, 0x1c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x84, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x48, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0xcd, 0x01,
	0x0a, 0x07, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x46, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49,
	0x64, 0x1a, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x34, 0xc2, 0xb8, 0x18, 0x30, 0x08, 0x6c, 0x10,
	0x18, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x92, 0x01,
	0x0a, 0x0f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x38, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x45, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x12, 0x99, 0x01, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3b, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x45, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x32, 0xe9,
	0x03, 0x0a, 0x1c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12,
	0xf3, 0x01, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x44, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x1a, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x5c, 0xc2, 0xb8, 0x18, 0x58, 0x08, 0x6a, 0x10,
	0x09, 0x1a, 0x24, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x2c, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x20, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x20,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x12, 0xd2, 0x01, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x12,
	0x46, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x1a, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x3b, 0xc2,
	0xb8, 0x18, 0x37, 0x08, 0x6e, 0x10, 0x18, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2a,
	0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2d, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x20, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x42, 0xb8, 0x03, 0x0a, 0x43, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a,
	0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x56,
	0x43, 0x44, 0x53, 0xaa, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0xca, 0x02, 0x35, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0xe2, 0x02, 0x41, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto_goTypes = []interface{}{
	(*StorageBucket)(nil),             // 0: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket
	(*StorageBucketId)(nil),           // 1: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketId
	(*pagination.PageInfo)(nil),       // 2: cloud.planton.apis.v1.commons.pagination.PageInfo
	(*product.ProductId)(nil),         // 3: cloud.planton.apis.v1.resourcemanager.product.ProductId
	(*environment.EnvironmentId)(nil), // 4: cloud.planton.apis.v1.code2cloud.environment.EnvironmentId
	(*StorageBucketList)(nil),         // 5: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketList
	(*StorageBuckets)(nil),            // 6: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBuckets
}
var file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto_depIdxs = []int32{
	0,  // 0: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketCommandController.create:input_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket
	0,  // 1: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketCommandController.update:input_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket
	1,  // 2: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketCommandController.delete:input_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketId
	0,  // 3: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketCommandController.restore:input_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket
	2,  // 4: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketQueryController.list:input_type -> cloud.planton.apis.v1.commons.pagination.PageInfo
	1,  // 5: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketQueryController.getById:input_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketId
	3,  // 6: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketQueryController.findByProductId:input_type -> cloud.planton.apis.v1.resourcemanager.product.ProductId
	4,  // 7: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketQueryController.findByEnvironmentId:input_type -> cloud.planton.apis.v1.code2cloud.environment.EnvironmentId
	0,  // 8: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketStackController.preview:input_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket
	1,  // 9: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketStackController.apply:input_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketId
	0,  // 10: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketCommandController.create:output_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket
	0,  // 11: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketCommandController.update:output_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket
	0,  // 12: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketCommandController.delete:output_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket
	0,  // 13: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketCommandController.restore:output_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket
	5,  // 14: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketQueryController.list:output_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketList
	0,  // 15: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketQueryController.getById:output_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket
	6,  // 16: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketQueryController.findByProductId:output_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBuckets
	6,  // 17: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketQueryController.findByEnvironmentId:output_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBuckets
	0,  // 18: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketStackController.preview:output_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket
	0,  // 19: cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucketStackController.apply:output_type -> cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto != nil {
		return
	}
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_storagebucket_service_proto_depIdxs = nil
}
