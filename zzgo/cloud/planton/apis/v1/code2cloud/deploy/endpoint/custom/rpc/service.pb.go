// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/endpoint/custom/rpc/service.proto

package rpc

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/environment/rpc"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/extensions"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/integration/kubernetes/resource"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/product/rpc"
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

var File_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x49, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x47, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x47, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb3,
	0x07, 0x0a, 0x1f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0xe5, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x4b, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x41, 0xc2, 0xb8, 0x18, 0x3d, 0x08, 0x23, 0x10,
	0x14, 0x1a, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x22, 0x26, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0xe1, 0x01, 0x0a, 0x06, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x1a, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22,
	0x3d, 0xc2, 0xb8, 0x18, 0x39, 0x08, 0x27, 0x10, 0x07, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x26, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0xdd,
	0x01, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x4d, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x37, 0xc2, 0xb8, 0x18, 0x33, 0x08, 0x24, 0x10, 0x07, 0x1a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0xe3,
	0x01, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x4b, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x22, 0x3e, 0xc2, 0xb8, 0x18, 0x3a, 0x08, 0x26, 0x10, 0x07, 0x1a, 0x0b,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x27, 0x75, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x32, 0xde, 0x0a, 0x0a, 0x1d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x8f, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x4f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0xdc, 0x01, 0x0a, 0x07, 0x67, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x4d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x49, 0x64, 0x1a, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x22, 0x35, 0xc2, 0xb8, 0x18, 0x31, 0x08, 0x25, 0x10, 0x07, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x24, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x9d, 0x01, 0x0a, 0x0f, 0x66, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x3c, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x8e, 0x02, 0x0a, 0x1b, 0x67, 0x65, 0x74, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x65, 0x72,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x55, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x49, 0xc2,
	0xb8, 0x18, 0x45, 0x08, 0x25, 0x10, 0x07, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38,
	0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x9c, 0x02, 0x0a, 0x24, 0x67, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x73,
	0x6e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x4d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64,
	0x1a, 0x5e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x45, 0xc2, 0xb8, 0x18, 0x41, 0x08, 0x25, 0x10, 0x07, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x34, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65,
	0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0xfb, 0x01, 0x0a, 0x1e, 0x66, 0x69, 0x6e, 0x64,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x4d, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x43, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x22, 0x45,
	0xc2, 0xb8, 0x18, 0x41, 0x08, 0x25, 0x10, 0x07, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x34, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x6c, 0x6f, 0x6f, 0x6b, 0x20, 0x75, 0x70, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x73, 0x32, 0xce, 0x06, 0x0a, 0x24, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x87,
	0x02, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x12, 0x67, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x4a, 0xc2, 0xb8,
	0x18, 0x46, 0x08, 0x27, 0x10, 0x07, 0x1a, 0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x2c, 0x75, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x20, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x8a, 0x02, 0x0a, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x67, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x4b, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x4a, 0xc2, 0xb8, 0x18, 0x46, 0x08,
	0x27, 0x10, 0x07, 0x1a, 0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x2c, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x8e, 0x02, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x6b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x4b, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x4a, 0xc2, 0xb8, 0x18, 0x46,
	0x08, 0x27, 0x10, 0x07, 0x1a, 0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x2c, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x20, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0xef, 0x03, 0x0a, 0x1d, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xe9, 0x01, 0x0a, 0x07, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x1a, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x44,
	0xc2, 0xb8, 0x18, 0x40, 0x08, 0x27, 0x10, 0x07, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x2d, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x20, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x12, 0xe1, 0x01, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x4d,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x4b, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x3c, 0xc2, 0xb8, 0x18, 0x38,
	0x08, 0x27, 0x10, 0x07, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2b, 0x75, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x20, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x42, 0xe0, 0x03, 0x0a, 0x49, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2f,
	0x72, 0x70, 0x63, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x45, 0x43, 0x52, 0xaa,
	0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x52, 0x70, 0x63, 0xca, 0x02, 0x3b,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5c, 0x52, 0x70, 0x63, 0xe2, 0x02, 0x47, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5c, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3a, 0x3a,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto_goTypes = []interface{}{
	(*CustomEndpoint)(nil),                                 // 0: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	(*CustomEndpointId)(nil),                               // 1: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointId
	(*pagination.PageInfo)(nil),                            // 2: cloud.planton.apis.v1.commons.rpc.pagination.PageInfo
	(*rpc.ProductId)(nil),                                  // 3: cloud.planton.apis.v1.resourcemanager.product.rpc.ProductId
	(*AddOrUpdateCustomEndpointRouteCommandInput)(nil),     // 4: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.AddOrUpdateCustomEndpointRouteCommandInput
	(*DeleteOrRestoreCustomEndpointRouteCommandInput)(nil), // 5: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.DeleteOrRestoreCustomEndpointRouteCommandInput
	(*CustomEndpointList)(nil),                             // 6: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointList
	(*CustomEndpoints)(nil),                                // 7: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoints
	(*CustomEndpointCertStatus)(nil),                       // 8: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCertStatus
	(*CustomEndpointDnsResolutionStatus)(nil),              // 9: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointDnsResolutionStatus
	(*resource.Certificates)(nil),                          // 10: cloud.planton.apis.v1.integration.kubernetes.resource.Certificates
}
var file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto_depIdxs = []int32{
	0,  // 0: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController.create:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	0,  // 1: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController.update:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	1,  // 2: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController.delete:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointId
	0,  // 3: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController.restore:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	2,  // 4: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController.list:input_type -> cloud.planton.apis.v1.commons.rpc.pagination.PageInfo
	1,  // 5: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController.getById:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointId
	3,  // 6: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController.findByProductId:input_type -> cloud.planton.apis.v1.resourcemanager.product.rpc.ProductId
	1,  // 7: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController.getCustomEndpointCertStatus:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointId
	1,  // 8: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController.getCustomEndpointDsnResolutionStatus:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointId
	1,  // 9: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController.findCustomEndpointCertificates:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointId
	4,  // 10: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointRouteCommandController.add:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.AddOrUpdateCustomEndpointRouteCommandInput
	4,  // 11: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointRouteCommandController.update:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.AddOrUpdateCustomEndpointRouteCommandInput
	5,  // 12: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointRouteCommandController.delete:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.DeleteOrRestoreCustomEndpointRouteCommandInput
	0,  // 13: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointStackController.preview:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	1,  // 14: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointStackController.apply:input_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointId
	0,  // 15: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController.create:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	0,  // 16: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController.update:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	0,  // 17: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController.delete:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	0,  // 18: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController.restore:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	6,  // 19: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController.list:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointList
	0,  // 20: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController.getById:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	7,  // 21: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController.findByProductId:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoints
	8,  // 22: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController.getCustomEndpointCertStatus:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCertStatus
	9,  // 23: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController.getCustomEndpointDsnResolutionStatus:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointDnsResolutionStatus
	10, // 24: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController.findCustomEndpointCertificates:output_type -> cloud.planton.apis.v1.integration.kubernetes.resource.Certificates
	0,  // 25: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointRouteCommandController.add:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	0,  // 26: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointRouteCommandController.update:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	0,  // 27: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointRouteCommandController.delete:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	0,  // 28: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointStackController.preview:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	0,  // 29: cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointStackController.apply:output_type -> cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto != nil {
		return
	}
	file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_endpoint_custom_rpc_service_proto_depIdxs = nil
}