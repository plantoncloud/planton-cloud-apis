// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/iam/authz/policy/service.proto

package policy

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/extensions"
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

var File_cloud_planton_apis_v1_iam_authz_policy_service_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_iam_authz_policy_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xb7, 0x04, 0x0a, 0x1a, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xa8,
	0x01, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x12, 0x39, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x41, 0x64, 0x64, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x3e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x49, 0x61, 0x6d, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x22, 0x26, 0xc2, 0xb8, 0x18, 0x22, 0x08, 0x40, 0x22, 0x1e, 0x75, 0x6e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x64, 0x64, 0x20, 0x69,
	0x61, 0x6d, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0xbb, 0x01, 0x0a, 0x0e, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x3e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x61, 0x6d, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x3e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x42, 0x79, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x22, 0x29, 0xc2, 0xb8,
	0x18, 0x25, 0x08, 0x40, 0x22, 0x21, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20, 0x69, 0x61, 0x6d,
	0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0xaf, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x42, 0x79, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x22, 0x29,
	0xc2, 0xb8, 0x18, 0x25, 0x08, 0x40, 0x22, 0x21, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x69,
	0x61, 0x6d, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x32, 0xe4, 0x03, 0x0a, 0x18, 0x49, 0x61,
	0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xdf, 0x01, 0x0a, 0x1e, 0x67, 0x65, 0x74, 0x42, 0x79,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42,
	0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x3e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x22,
	0x27, 0xc2, 0xb8, 0x18, 0x23, 0x08, 0x3f, 0x22, 0x1f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x69, 0x61,
	0x6d, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0xe5, 0x01, 0x0a, 0x29, 0x67, 0x65, 0x74,
	0x42, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x79, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x39, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x27, 0xc2, 0xb8, 0x18, 0x23, 0x08, 0x3f, 0x22,
	0x1f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x69, 0x61, 0x6d, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x42, 0xde, 0x02, 0x0a, 0x34, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67,
	0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x56, 0x49, 0x41, 0x50, 0xaa, 0x02, 0x26, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73,
	0x2e, 0x56, 0x31, 0x2e, 0x49, 0x61, 0x6d, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0xca, 0x02, 0x26, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x61, 0x6d,
	0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0xe2, 0x02, 0x32,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x2c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x49,
	0x61, 0x6d, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x3a, 0x3a, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_iam_authz_policy_service_proto_goTypes = []interface{}{
	(*AddIamPolicyInput)(nil),                            // 0: cloud.planton.apis.v1.iam.authz.policy.AddIamPolicyInput
	(*RemoveIamPoliciesInput)(nil),                       // 1: cloud.planton.apis.v1.iam.authz.policy.RemoveIamPoliciesInput
	(*UpdateIamPolicyInput)(nil),                         // 2: cloud.planton.apis.v1.iam.authz.policy.UpdateIamPolicyInput
	(*GetIamPolicyByResourceTypeAndResourceIdInput)(nil), // 3: cloud.planton.apis.v1.iam.authz.policy.GetIamPolicyByResourceTypeAndResourceIdInput
	(*IamPoliciesByPrincipal)(nil),                       // 4: cloud.planton.apis.v1.iam.authz.policy.IamPoliciesByPrincipal
	(*IamPolicyByPrincipal)(nil),                         // 5: cloud.planton.apis.v1.iam.authz.policy.IamPolicyByPrincipal
	(*IamPoliciesByRole)(nil),                            // 6: cloud.planton.apis.v1.iam.authz.policy.IamPoliciesByRole
}
var file_cloud_planton_apis_v1_iam_authz_policy_service_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.iam.authz.policy.IamPolicyCommandController.add:input_type -> cloud.planton.apis.v1.iam.authz.policy.AddIamPolicyInput
	1, // 1: cloud.planton.apis.v1.iam.authz.policy.IamPolicyCommandController.removeMultiple:input_type -> cloud.planton.apis.v1.iam.authz.policy.RemoveIamPoliciesInput
	2, // 2: cloud.planton.apis.v1.iam.authz.policy.IamPolicyCommandController.update:input_type -> cloud.planton.apis.v1.iam.authz.policy.UpdateIamPolicyInput
	3, // 3: cloud.planton.apis.v1.iam.authz.policy.IamPolicyQueryController.getByResourceTypeAndResourceId:input_type -> cloud.planton.apis.v1.iam.authz.policy.GetIamPolicyByResourceTypeAndResourceIdInput
	3, // 4: cloud.planton.apis.v1.iam.authz.policy.IamPolicyQueryController.getByResourceTypeAndResourceIdGroupByRole:input_type -> cloud.planton.apis.v1.iam.authz.policy.GetIamPolicyByResourceTypeAndResourceIdInput
	4, // 5: cloud.planton.apis.v1.iam.authz.policy.IamPolicyCommandController.add:output_type -> cloud.planton.apis.v1.iam.authz.policy.IamPoliciesByPrincipal
	4, // 6: cloud.planton.apis.v1.iam.authz.policy.IamPolicyCommandController.removeMultiple:output_type -> cloud.planton.apis.v1.iam.authz.policy.IamPoliciesByPrincipal
	5, // 7: cloud.planton.apis.v1.iam.authz.policy.IamPolicyCommandController.update:output_type -> cloud.planton.apis.v1.iam.authz.policy.IamPolicyByPrincipal
	4, // 8: cloud.planton.apis.v1.iam.authz.policy.IamPolicyQueryController.getByResourceTypeAndResourceId:output_type -> cloud.planton.apis.v1.iam.authz.policy.IamPoliciesByPrincipal
	6, // 9: cloud.planton.apis.v1.iam.authz.policy.IamPolicyQueryController.getByResourceTypeAndResourceIdGroupByRole:output_type -> cloud.planton.apis.v1.iam.authz.policy.IamPoliciesByRole
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_iam_authz_policy_service_proto_init() }
func file_cloud_planton_apis_v1_iam_authz_policy_service_proto_init() {
	if File_cloud_planton_apis_v1_iam_authz_policy_service_proto != nil {
		return
	}
	file_cloud_planton_apis_v1_iam_authz_policy_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_iam_authz_policy_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cloud_planton_apis_v1_iam_authz_policy_service_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_iam_authz_policy_service_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_iam_authz_policy_service_proto = out.File
	file_cloud_planton_apis_v1_iam_authz_policy_service_proto_rawDesc = nil
	file_cloud_planton_apis_v1_iam_authz_policy_service_proto_goTypes = nil
	file_cloud_planton_apis_v1_iam_authz_policy_service_proto_depIdxs = nil
}
