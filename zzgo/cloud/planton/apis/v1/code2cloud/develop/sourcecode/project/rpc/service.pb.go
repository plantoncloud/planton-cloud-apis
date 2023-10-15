// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/develop/sourcecode/project/rpc/service.proto

package rpc

import (
	rpc1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/develop/sourcecode/server/rpc"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/extensions"
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

var File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x3f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63,
	0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4b, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xac, 0x0e, 0x0a, 0x1c, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0xd0, 0x01, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x12, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0x2d, 0xc2, 0xb8, 0x18, 0x29, 0x08, 0x14, 0x10, 0x14, 0x22, 0x23,
	0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0xd3, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x4c,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x4c, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2d, 0xc2, 0xb8, 0x18, 0x29,
	0x08, 0x14, 0x10, 0x14, 0x22, 0x23, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x63, 0x6f, 0x64,
	0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xe0, 0x01, 0x0a, 0x06, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x1a, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x3a, 0xc2, 0xb8, 0x18, 0x36, 0x08, 0x18, 0x10, 0x04, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x23, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20,
	0x63, 0x6f, 0x64, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xdc, 0x01, 0x0a,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x34, 0xc2, 0xb8, 0x18, 0x30, 0x08, 0x16, 0x10, 0x04, 0x1a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x63,
	0x6f, 0x64, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xe2, 0x01, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x22, 0x3b, 0xc2, 0xb8, 0x18, 0x37, 0x08, 0x17, 0x10, 0x04, 0x1a, 0x0b, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x24, 0x75, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x92, 0x01, 0x0a, 0x16, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65,
	0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x3c, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x3a, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0xf5, 0x01, 0x0a, 0x1a, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72,
	0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x1a, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x39, 0xc2, 0xb8, 0x18, 0x35, 0x08, 0x18, 0x10, 0x04, 0x1a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x28, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65,
	0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xae, 0x02,
	0x0a, 0x23, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x47, 0xc2, 0xb8, 0x18, 0x43, 0x08, 0x44, 0x10, 0x0c, 0x1a,
	0x15, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x26, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x20, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xfa,
	0x0a, 0x0a, 0x1a, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x90, 0x01,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x50,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0xda, 0x01, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x4e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x4c, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x31, 0xc2, 0xb8, 0x18, 0x2d,
	0x08, 0x15, 0x10, 0x04, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x75, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74,
	0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x9e, 0x01,
	0x0a, 0x0f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a,
	0x4d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0xb1,
	0x01, 0x0a, 0x12, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x1a, 0x4d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x12, 0xe3, 0x01, 0x0a, 0x14, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x4f, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x1a, 0x4d, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x2b, 0xc2, 0xb8, 0x18,
	0x27, 0x08, 0x15, 0x10, 0x04, 0x22, 0x21, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0xb2, 0x01, 0x0a, 0x23, 0x66, 0x69, 0x6e,
	0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x4d,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0xfb, 0x01,
	0x0a, 0x19, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x4e, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x53, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x22, 0x39, 0xc2, 0xb8, 0x18, 0x35, 0x08, 0x15, 0x10, 0x04, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x28, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0xee, 0x03, 0x0a, 0x43,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x72, 0x70, 0x63, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x6f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x72, 0x70, 0x63, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x53, 0x50, 0x52,
	0xaa, 0x02, 0x3f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x52,
	0x70, 0x63, 0xca, 0x02, 0x3f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x5c, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5c, 0x52, 0x70, 0x63, 0xe2, 0x02, 0x4b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x5c,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x47, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x3a, 0x3a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3a,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto_goTypes = []interface{}{
	(*CodeProject)(nil),   // 0: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	(*CodeProjectId)(nil), // 1: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectId
	(*rpc.ProductId)(nil), // 2: cloud.planton.apis.v1.resourcemanager.product.rpc.ProductId
	(*AttachMachineAccountByCodeProjectIdCommandInput)(nil), // 3: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.AttachMachineAccountByCodeProjectIdCommandInput
	(*pagination.PageInfo)(nil),                             // 4: cloud.planton.apis.v1.commons.rpc.pagination.PageInfo
	(*rpc1.CodeServerId)(nil),                               // 5: cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.rpc.CodeServerId
	(*CodeProjectUrl)(nil),                                  // 6: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectUrl
	(*rpc.Product)(nil),                                     // 7: cloud.planton.apis.v1.resourcemanager.product.rpc.Product
	(*CodeProjectList)(nil),                                 // 8: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectList
	(*CodeProjects)(nil),                                    // 9: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjects
	(*CodeProjectProfile)(nil),                              // 10: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectProfile
}
var file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto_depIdxs = []int32{
	0,  // 0: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.add:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	0,  // 1: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.create:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	0,  // 2: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.update:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	1,  // 3: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.delete:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectId
	0,  // 4: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.restore:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	2,  // 5: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.synchronizeByProductId:input_type -> cloud.planton.apis.v1.resourcemanager.product.rpc.ProductId
	1,  // 6: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.synchronizeByCodeProjectId:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectId
	3,  // 7: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.attachMachineAccountByCodeProjectId:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.AttachMachineAccountByCodeProjectIdCommandInput
	4,  // 8: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.list:input_type -> cloud.planton.apis.v1.commons.rpc.pagination.PageInfo
	1,  // 9: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.getById:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectId
	2,  // 10: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.findByProductId:input_type -> cloud.planton.apis.v1.resourcemanager.product.rpc.ProductId
	5,  // 11: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.findByCodeServerId:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.rpc.CodeServerId
	6,  // 12: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.findByCodeProjectUrl:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectUrl
	2,  // 13: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.findTemplateCodeProjectsByProductId:input_type -> cloud.planton.apis.v1.resourcemanager.product.rpc.ProductId
	1,  // 14: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.getCodeProjectProfileById:input_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectId
	0,  // 15: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.add:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	0,  // 16: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.create:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	0,  // 17: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.update:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	0,  // 18: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.delete:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	0,  // 19: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.restore:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	7,  // 20: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.synchronizeByProductId:output_type -> cloud.planton.apis.v1.resourcemanager.product.rpc.Product
	0,  // 21: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.synchronizeByCodeProjectId:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	0,  // 22: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectCommandController.attachMachineAccountByCodeProjectId:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	8,  // 23: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.list:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectList
	0,  // 24: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.getById:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject
	9,  // 25: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.findByProductId:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjects
	9,  // 26: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.findByCodeServerId:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjects
	9,  // 27: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.findByCodeProjectUrl:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjects
	9,  // 28: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.findTemplateCodeProjectsByProductId:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjects
	10, // 29: cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectQueryController.getCodeProjectProfileById:output_type -> cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProjectProfile
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto_init()
}
func file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto != nil {
		return
	}
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_develop_sourcecode_project_rpc_service_proto_depIdxs = nil
}
