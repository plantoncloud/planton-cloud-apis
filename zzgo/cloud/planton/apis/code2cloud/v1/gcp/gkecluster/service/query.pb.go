// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/gcp/gkecluster/service/query.proto

package service

import (
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gkecluster/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcerpcmethodoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcerpcserviceoptions"
	protobuf "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/protobuf"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iam/v1/iampolicy/options/rpcauthorizationmethodoptions"
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

var File_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto_rawDesc = []byte{
	0x0a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x6b, 0x65, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x3e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x70, 0x63,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72,
	0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xea, 0x01, 0x0a, 0x19,
	0x47, 0x6b, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xc6, 0x01, 0x0a, 0x03, 0x67, 0x65,
	0x74, 0x12, 0x43, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x6b, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x6b,
	0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47,
	0x6b, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x37, 0xc2, 0xb8, 0x18, 0x2f, 0x08,
	0x02, 0x10, 0x11, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x75, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x6c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x20, 0x67, 0x6b, 0x65, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0xb0, 0xff,
	0x2b, 0x03, 0x1a, 0x04, 0xa0, 0xff, 0x2b, 0x11, 0x32, 0xa8, 0x02, 0x0a, 0x21, 0x47, 0x6b, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x82,
	0x02, 0x0a, 0x1b, 0x67, 0x65, 0x74, 0x42, 0x79, 0x47, 0x63, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x55,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x47, 0x6b, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x64,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x49, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x6b, 0x65,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x6b,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c,
	0x22, 0x41, 0xc2, 0xb8, 0x18, 0x3d, 0x08, 0x02, 0x10, 0x11, 0x1a, 0x0e, 0x67, 0x6b, 0x65, 0x5f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x27, 0x75, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x20, 0x6e, 0x6f, 0x64, 0x65, 0x20, 0x70,
	0x6f, 0x6f, 0x6c, 0x32, 0xd8, 0x02, 0x0a, 0x12, 0x47, 0x63, 0x70, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x8b, 0x01, 0x0a, 0x0b, 0x66,
	0x69, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x44, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x67, 0x63, 0x70, 0x2e, 0x67, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x63, 0x70, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x12, 0xb3, 0x01, 0x0a, 0x1b, 0x66, 0x69, 0x6e,
	0x64, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x42, 0x79, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x4a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e,
	0x67, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x47, 0x63, 0x70, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x1a, 0x42, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x6b, 0x65, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x63, 0x70,
	0x5a, 0x6f, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x42, 0xc4,
	0x03, 0x0a, 0x45, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x6b, 0x65,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2,
	0x02, 0x08, 0x43, 0x50, 0x41, 0x43, 0x56, 0x47, 0x47, 0x53, 0xaa, 0x02, 0x37, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x47, 0x63,
	0x70, 0x2e, 0x47, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x6b, 0x65, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02,
	0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x47, 0x63, 0x70,
	0x3a, 0x3a, 0x47, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto_goTypes = []interface{}{
	(*model.GkeClusterId)(nil),                   // 0: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GkeClusterId
	(*model.GetByGkeClusterNodePoolIdInput)(nil), // 1: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GetByGkeClusterNodePoolIdInput
	(*protobuf.CustomEmpty)(nil),                 // 2: cloud.planton.apis.commons.protobuf.CustomEmpty
	(*model.GcpRegionIdentifier)(nil),            // 3: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GcpRegionIdentifier
	(*model.GkeCluster)(nil),                     // 4: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GkeCluster
	(*model.GkeClusterNodePool)(nil),             // 5: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GkeClusterNodePool
	(*model.GcpRegionList)(nil),                  // 6: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GcpRegionList
	(*model.GcpZoneList)(nil),                    // 7: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GcpZoneList
}
var file_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.service.GkeClusterQueryController.get:input_type -> cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GkeClusterId
	1, // 1: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.service.GkeClusterNodePoolQueryController.getByGcpContainerNodePoolId:input_type -> cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GetByGkeClusterNodePoolIdInput
	2, // 2: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.service.GcpQueryController.findRegions:input_type -> cloud.planton.apis.commons.protobuf.CustomEmpty
	3, // 3: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.service.GcpQueryController.findZonesByRegionIdentifier:input_type -> cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GcpRegionIdentifier
	4, // 4: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.service.GkeClusterQueryController.get:output_type -> cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GkeCluster
	5, // 5: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.service.GkeClusterNodePoolQueryController.getByGcpContainerNodePoolId:output_type -> cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GkeClusterNodePool
	6, // 6: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.service.GcpQueryController.findRegions:output_type -> cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GcpRegionList
	7, // 7: cloud.planton.apis.code2cloud.v1.gcp.gkecluster.service.GcpQueryController.findZonesByRegionIdentifier:output_type -> cloud.planton.apis.code2cloud.v1.gcp.gkecluster.model.GcpZoneList
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_gcp_gkecluster_service_query_proto_depIdxs = nil
}
