// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/stack/job/service.proto

package job

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/grpc/stream"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/pulumi/engine"
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

var File_cloud_planton_apis_v1_stack_job_service_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_stack_job_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f,
	0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62,
	0x1a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xe6, 0x02, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x5e, 0x0a,
	0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a,
	0x6f, 0x62, 0x1a, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x5e, 0x0a,
	0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a,
	0x6f, 0x62, 0x1a, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x88, 0x01,
	0x0a, 0x12, 0x73, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x47, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a,
	0x6f, 0x62, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x29, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x32, 0x87, 0x04, 0x0a, 0x17, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x81, 0x01, 0x0a, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x73, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x61, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64,
	0x1a, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x7f, 0x0a, 0x16, 0x67,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62,
	0x49, 0x64, 0x1a, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x12, 0x83, 0x01, 0x0a,
	0x1c, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2b, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x1a, 0x34, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x30, 0x01, 0x42, 0xb2, 0x02, 0x0a, 0x2d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x6a, 0x6f, 0x62, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0xa2, 0x02, 0x06, 0x43, 0x50, 0x41, 0x56, 0x53, 0x4a, 0xaa, 0x02,
	0x1f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41,
	0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x4a, 0x6f, 0x62,
	0xca, 0x02, 0x1f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4a,
	0x6f, 0x62, 0xe2, 0x02, 0x2b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x5c, 0x4a, 0x6f, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x24, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x3a, 0x3a, 0x4a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_stack_job_service_proto_goTypes = []interface{}{
	(*StackJob)(nil), // 0: cloud.planton.apis.v1.stack.job.StackJob
	(*SetStackJobExecutionStatusCommandInput)(nil), // 1: cloud.planton.apis.v1.stack.job.SetStackJobExecutionStatusCommandInput
	(*ListStackJobsByFiltersQueryInput)(nil),       // 2: cloud.planton.apis.v1.stack.job.ListStackJobsByFiltersQueryInput
	(*StackJobId)(nil),                             // 3: cloud.planton.apis.v1.stack.job.StackJobId
	(*StackJobList)(nil),                           // 4: cloud.planton.apis.v1.stack.job.StackJobList
	(*StackJobProgressEvent)(nil),                  // 5: cloud.planton.apis.v1.stack.job.StackJobProgressEvent
	(*StackJobLogSnapshot)(nil),                    // 6: cloud.planton.apis.v1.stack.job.StackJobLogSnapshot
}
var file_cloud_planton_apis_v1_stack_job_service_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.stack.job.StackJobCommandController.create:input_type -> cloud.planton.apis.v1.stack.job.StackJob
	0, // 1: cloud.planton.apis.v1.stack.job.StackJobCommandController.update:input_type -> cloud.planton.apis.v1.stack.job.StackJob
	1, // 2: cloud.planton.apis.v1.stack.job.StackJobCommandController.setExecutionStatus:input_type -> cloud.planton.apis.v1.stack.job.SetStackJobExecutionStatusCommandInput
	2, // 3: cloud.planton.apis.v1.stack.job.StackJobQueryController.listByFilters:input_type -> cloud.planton.apis.v1.stack.job.ListStackJobsByFiltersQueryInput
	3, // 4: cloud.planton.apis.v1.stack.job.StackJobQueryController.getById:input_type -> cloud.planton.apis.v1.stack.job.StackJobId
	3, // 5: cloud.planton.apis.v1.stack.job.StackJobQueryController.getProgressEventStream:input_type -> cloud.planton.apis.v1.stack.job.StackJobId
	3, // 6: cloud.planton.apis.v1.stack.job.StackJobQueryController.getStackJobLogSnapshotStream:input_type -> cloud.planton.apis.v1.stack.job.StackJobId
	0, // 7: cloud.planton.apis.v1.stack.job.StackJobCommandController.create:output_type -> cloud.planton.apis.v1.stack.job.StackJob
	0, // 8: cloud.planton.apis.v1.stack.job.StackJobCommandController.update:output_type -> cloud.planton.apis.v1.stack.job.StackJob
	0, // 9: cloud.planton.apis.v1.stack.job.StackJobCommandController.setExecutionStatus:output_type -> cloud.planton.apis.v1.stack.job.StackJob
	4, // 10: cloud.planton.apis.v1.stack.job.StackJobQueryController.listByFilters:output_type -> cloud.planton.apis.v1.stack.job.StackJobList
	0, // 11: cloud.planton.apis.v1.stack.job.StackJobQueryController.getById:output_type -> cloud.planton.apis.v1.stack.job.StackJob
	5, // 12: cloud.planton.apis.v1.stack.job.StackJobQueryController.getProgressEventStream:output_type -> cloud.planton.apis.v1.stack.job.StackJobProgressEvent
	6, // 13: cloud.planton.apis.v1.stack.job.StackJobQueryController.getStackJobLogSnapshotStream:output_type -> cloud.planton.apis.v1.stack.job.StackJobLogSnapshot
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_stack_job_service_proto_init() }
func file_cloud_planton_apis_v1_stack_job_service_proto_init() {
	if File_cloud_planton_apis_v1_stack_job_service_proto != nil {
		return
	}
	file_cloud_planton_apis_v1_stack_job_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_stack_job_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cloud_planton_apis_v1_stack_job_service_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_stack_job_service_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_stack_job_service_proto = out.File
	file_cloud_planton_apis_v1_stack_job_service_proto_rawDesc = nil
	file_cloud_planton_apis_v1_stack_job_service_proto_goTypes = nil
	file_cloud_planton_apis_v1_stack_job_service_proto_depIdxs = nil
}
