// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/iam/authz/policy/service/query.proto

package service

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/extensions"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/iam/authz/policy/model"
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

var File_cloud_planton_apis_v1_iam_authz_policy_service_query_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_iam_authz_policy_service_query_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x3b, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8e, 0x09, 0x0a, 0x18, 0x49,
	0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xeb, 0x01, 0x0a, 0x1e, 0x67, 0x65, 0x74, 0x42,
	0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x5a, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x61, 0x6d,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x42, 0x79, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x22, 0x27, 0xc2, 0xb8,
	0x18, 0x23, 0x08, 0x3f, 0x22, 0x1f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x69, 0x61, 0x6d, 0x20, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0xf1, 0x01, 0x0a, 0x29, 0x67, 0x65, 0x74, 0x42, 0x79, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x5a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42,
	0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x3f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49,
	0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65,
	0x22, 0x27, 0xc2, 0xb8, 0x18, 0x23, 0x08, 0x3f, 0x22, 0x1f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x69,
	0x61, 0x6d, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0xf8, 0x01, 0x0a, 0x30, 0x67, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x5a,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x79, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x3f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x27, 0xc2, 0xb8, 0x18,
	0x23, 0x08, 0x3f, 0x22, 0x1f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x69, 0x61, 0x6d, 0x20, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0xbc, 0x01, 0x0a, 0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x3a, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x73, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x22, 0x28, 0xc2, 0xb8, 0x18, 0x24, 0x08,
	0x3f, 0x22, 0x20, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x69, 0x61, 0x6d, 0x20, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0xd5, 0x01, 0x0a, 0x19, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x73, 0x12, 0x40, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x43, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0x31, 0xc2, 0xb8, 0x18, 0x2d, 0x08, 0x3f,
	0x22, 0x29, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74,
	0x6f, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x8e, 0x03, 0x0a, 0x3c,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0a, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41,
	0x56, 0x49, 0x41, 0x50, 0x53, 0xaa, 0x02, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x49, 0x61,
	0x6d, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49,
	0x61, 0x6d, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x49, 0x61, 0x6d, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x49, 0x61, 0x6d, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x3a, 0x3a, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_iam_authz_policy_service_query_proto_goTypes = []interface{}{
	(*model.GetIamPolicyByResourceTypeAndResourceIdInput)(nil), // 0: cloud.planton.apis.v1.iam.authz.policy.model.GetIamPolicyByResourceTypeAndResourceIdInput
	(*model.AuthorizationInput)(nil),                           // 1: cloud.planton.apis.v1.iam.authz.policy.model.AuthorizationInput
	(*model.IamPoliciesByPrincipal)(nil),                       // 2: cloud.planton.apis.v1.iam.authz.policy.model.IamPoliciesByPrincipal
	(*model.IamPoliciesByRole)(nil),                            // 3: cloud.planton.apis.v1.iam.authz.policy.model.IamPoliciesByRole
	(*model.IsAuthorized)(nil),                                 // 4: cloud.planton.apis.v1.iam.authz.policy.model.IsAuthorized
	(*model.AuthorizedResourceIds)(nil),                        // 5: cloud.planton.apis.v1.iam.authz.policy.model.AuthorizedResourceIds
}
var file_cloud_planton_apis_v1_iam_authz_policy_service_query_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.iam.authz.policy.service.IamPolicyQueryController.getByResourceTypeAndResourceId:input_type -> cloud.planton.apis.v1.iam.authz.policy.model.GetIamPolicyByResourceTypeAndResourceIdInput
	0, // 1: cloud.planton.apis.v1.iam.authz.policy.service.IamPolicyQueryController.getByResourceTypeAndResourceIdGroupByRole:input_type -> cloud.planton.apis.v1.iam.authz.policy.model.GetIamPolicyByResourceTypeAndResourceIdInput
	0, // 2: cloud.planton.apis.v1.iam.authz.policy.service.IamPolicyQueryController.getCompanyByResourceTypeAndResourceIdGroupByRole:input_type -> cloud.planton.apis.v1.iam.authz.policy.model.GetIamPolicyByResourceTypeAndResourceIdInput
	1, // 3: cloud.planton.apis.v1.iam.authz.policy.service.IamPolicyQueryController.checkAuthorization:input_type -> cloud.planton.apis.v1.iam.authz.policy.model.AuthorizationInput
	1, // 4: cloud.planton.apis.v1.iam.authz.policy.service.IamPolicyQueryController.listAuthorizedResourceIds:input_type -> cloud.planton.apis.v1.iam.authz.policy.model.AuthorizationInput
	2, // 5: cloud.planton.apis.v1.iam.authz.policy.service.IamPolicyQueryController.getByResourceTypeAndResourceId:output_type -> cloud.planton.apis.v1.iam.authz.policy.model.IamPoliciesByPrincipal
	3, // 6: cloud.planton.apis.v1.iam.authz.policy.service.IamPolicyQueryController.getByResourceTypeAndResourceIdGroupByRole:output_type -> cloud.planton.apis.v1.iam.authz.policy.model.IamPoliciesByRole
	3, // 7: cloud.planton.apis.v1.iam.authz.policy.service.IamPolicyQueryController.getCompanyByResourceTypeAndResourceIdGroupByRole:output_type -> cloud.planton.apis.v1.iam.authz.policy.model.IamPoliciesByRole
	4, // 8: cloud.planton.apis.v1.iam.authz.policy.service.IamPolicyQueryController.checkAuthorization:output_type -> cloud.planton.apis.v1.iam.authz.policy.model.IsAuthorized
	5, // 9: cloud.planton.apis.v1.iam.authz.policy.service.IamPolicyQueryController.listAuthorizedResourceIds:output_type -> cloud.planton.apis.v1.iam.authz.policy.model.AuthorizedResourceIds
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_iam_authz_policy_service_query_proto_init() }
func file_cloud_planton_apis_v1_iam_authz_policy_service_query_proto_init() {
	if File_cloud_planton_apis_v1_iam_authz_policy_service_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_iam_authz_policy_service_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_planton_apis_v1_iam_authz_policy_service_query_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_iam_authz_policy_service_query_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_iam_authz_policy_service_query_proto = out.File
	file_cloud_planton_apis_v1_iam_authz_policy_service_query_proto_rawDesc = nil
	file_cloud_planton_apis_v1_iam_authz_policy_service_query_proto_goTypes = nil
	file_cloud_planton_apis_v1_iam_authz_policy_service_query_proto_depIdxs = nil
}
