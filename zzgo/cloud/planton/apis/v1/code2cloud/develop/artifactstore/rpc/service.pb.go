// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/develop/artifactstore/rpc/service.proto

package rpc

import (
	rpc1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/develop/sourcecode/project/rpc"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/extensions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/company/rpc"
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

var File_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x48, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3a, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x4b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xde, 0x09, 0x0a, 0x1e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x12, 0xe0, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x49,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x49, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x22, 0x40, 0xc2, 0xb8, 0x18, 0x3c, 0x08, 0x01, 0x10, 0x14, 0x1a, 0x0f,
	0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x22,
	0x25, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0xdc, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x49, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x49, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x3c, 0xc2, 0xb8, 0x18, 0x38, 0x08, 0x06, 0x10,
	0x01, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x25,
	0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2d,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0xd8, 0x01, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x1a, 0x49, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x36, 0xc2, 0xb8, 0x18, 0x32, 0x08, 0x03,
	0x10, 0x01, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x25, 0x75, 0x6e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x20, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0xde, 0x01, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x49, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x49, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x22, 0x3d, 0xc2, 0xb8, 0x18, 0x39, 0x08, 0x05, 0x10, 0x01, 0x1a, 0x0b, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x26, 0x75, 0x6e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x20, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2d, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0xbd, 0x02, 0x0a, 0x21, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x57, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x57, 0xc2, 0xb8, 0x18, 0x53, 0x08, 0x06,
	0x10, 0x01, 0x1a, 0x11, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x3a, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x72,
	0x6f, 0x6d, 0x20, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2d, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x32, 0xf1, 0x0a, 0x0a, 0x1c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0xd7, 0x01, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x4b,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x1a, 0x49, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x34, 0xc2, 0xb8, 0x18, 0x30, 0x08, 0x04, 0x10, 0x01,
	0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x8d, 0x01, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x4d, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x9b, 0x01, 0x0a,
	0x0f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x4a,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0xb3, 0x01, 0x0a, 0x14, 0x66,
	0x69, 0x6e, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x6c, 0x12, 0x4f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x55, 0x72, 0x6c, 0x1a, 0x4a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x12, 0xd9, 0x01, 0x0a, 0x1d, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x5e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x58, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0xd1, 0x01, 0x0a,
	0x19, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x5e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65, 0x70,
	0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x54, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0xe2, 0x01, 0x0a, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x5b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xe4, 0x03, 0x0a, 0x1c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xe4, 0x01, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x49, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x49, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x43, 0xc2, 0xb8, 0x18, 0x3f, 0x08, 0x06,
	0x10, 0x01, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22,
	0x2c, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x20, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x12, 0xdc, 0x01,
	0x0a, 0x05, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x1a, 0x49, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22,
	0x3b, 0xc2, 0xb8, 0x18, 0x37, 0x08, 0x06, 0x10, 0x01, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x2a, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74,
	0x6f, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x20, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x42, 0xce, 0x03, 0x0a,
	0x3e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x42,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x6a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x70, 0x63, 0xa2, 0x02, 0x08, 0x43, 0x50,
	0x41, 0x56, 0x43, 0x44, 0x41, 0x52, 0xaa, 0x02, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x52, 0x70, 0x63, 0xca, 0x02, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x5c, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x52, 0x70, 0x63,
	0xe2, 0x02, 0x46, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x5c, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x41, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x3a, 0x3a, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto_goTypes = []interface{}{
	(*ArtifactStore)(nil),                              // 0: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore
	(*ArtifactStoreId)(nil),                            // 1: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreId
	(*DelArtifactStorePackageVersionCommandInput)(nil), // 2: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.DelArtifactStorePackageVersionCommandInput
	(*pagination.PageInfo)(nil),                        // 3: cloud.planton.apis.v1.commons.rpc.pagination.PageInfo
	(*rpc.ProductId)(nil),                              // 4: cloud.planton.apis.v1.resourcemanager.product.rpc.ProductId
	(*rpc1.CodeProjectUrl)(nil),                        // 5: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectUrl
	(*ListByArtifactStoreIdRepoNameInput)(nil),         // 6: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ListByArtifactStoreIdRepoNameInput
	(*ListByArtifactStoreIdPackageNameInput)(nil),      // 7: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ListByArtifactStoreIdPackageNameInput
	(*ArtifactStorePackageVersion)(nil),                // 8: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStorePackageVersion
	(*ArtifactStoreList)(nil),                          // 9: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreList
	(*ArtifactStores)(nil),                             // 10: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStores
	(*ArtifactStoreDockerImageList)(nil),               // 11: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreDockerImageList
	(*ArtifactStorePackageList)(nil),                   // 12: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStorePackageList
	(*ArtifactStorePackageVersionList)(nil),            // 13: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStorePackageVersionList
}
var file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto_depIdxs = []int32{
	0,  // 0: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController.create:input_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore
	0,  // 1: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController.update:input_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore
	1,  // 2: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController.delete:input_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreId
	0,  // 3: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController.restore:input_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore
	2,  // 4: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController.deleteArtifactStorePackageVersion:input_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.DelArtifactStorePackageVersionCommandInput
	1,  // 5: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.getById:input_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreId
	3,  // 6: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.list:input_type -> cloud.planton.apis.v1.commons.rpc.pagination.PageInfo
	4,  // 7: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.findByProductId:input_type -> cloud.planton.apis.v1.resourcemanager.product.rpc.ProductId
	5,  // 8: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.findByCodeProjectUrl:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectUrl
	6,  // 9: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.listArtifactStoreDockerImages:input_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ListByArtifactStoreIdRepoNameInput
	6,  // 10: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.listArtifactStorePackages:input_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ListByArtifactStoreIdRepoNameInput
	7,  // 11: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.listArtifactStorePackageVersions:input_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ListByArtifactStoreIdPackageNameInput
	0,  // 12: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreStackController.preview:input_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore
	1,  // 13: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreStackController.apply:input_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreId
	0,  // 14: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController.create:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore
	0,  // 15: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController.update:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore
	0,  // 16: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController.delete:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore
	0,  // 17: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController.restore:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore
	8,  // 18: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreCommandController.deleteArtifactStorePackageVersion:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStorePackageVersion
	0,  // 19: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.getById:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore
	9,  // 20: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.list:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreList
	10, // 21: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.findByProductId:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStores
	10, // 22: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.findByCodeProjectUrl:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStores
	11, // 23: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.listArtifactStoreDockerImages:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreDockerImageList
	12, // 24: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.listArtifactStorePackages:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStorePackageList
	13, // 25: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreQueryController.listArtifactStorePackageVersions:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStorePackageVersionList
	0,  // 26: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreStackController.preview:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore
	0,  // 27: cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStoreStackController.apply:output_type -> cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto != nil {
		return
	}
	file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_develop_artifactstore_rpc_service_proto_depIdxs = nil
}
