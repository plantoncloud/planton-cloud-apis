// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/kubernetes/gitlabkubernetes/service/command.proto

package service

import (
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/gitlabkubernetes/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcerpcmethodoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcerpcserviceoptions"
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

var File_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto_rawDesc = []byte{
	0x0a, 0x52, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x44, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x4b, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x70,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x70, 0x63, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x70, 0x63, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x61, 0x6d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x70, 0x63, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfd, 0x16, 0x0a, 0x21, 0x47, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xa1, 0x02, 0x0a,
	0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x54,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x1a, 0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0x64, 0xc2, 0xb8, 0x18, 0x54,
	0x08, 0x01, 0x10, 0x0a, 0x1a, 0x1c, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x5f,
	0x69, 0x64, 0x22, 0x30, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x20, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0xa0, 0xff, 0x2b, 0x0a, 0xb0, 0xff, 0x2b, 0x01, 0xb8, 0xff, 0x2b, 0x01,
	0x12, 0x92, 0x02, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x54, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x1a, 0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0x5c, 0xc2, 0xb8, 0x18, 0x4c, 0x08, 0x01, 0x10,
	0x0a, 0x1a, 0x1c, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x69, 0x64, 0x22,
	0x28, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0xa0, 0xff, 0x2b, 0x01, 0xb0, 0xff, 0x2b,
	0x01, 0xb8, 0xff, 0x2b, 0x01, 0x12, 0x90, 0x02, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x1a, 0x54, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x22, 0x53, 0xc2, 0xb8, 0x18, 0x43, 0x08, 0x05, 0x10, 0x12, 0x1a, 0x0b, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x30, 0x75, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0xa0, 0xff, 0x2b, 0x0a,
	0xb0, 0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0x12, 0x81, 0x02, 0x0a, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x1a, 0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22,
	0x4b, 0xc2, 0xb8, 0x18, 0x3b, 0x08, 0x05, 0x10, 0x12, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22, 0x28, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0xa0, 0xff, 0x2b, 0x02, 0xb0, 0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0x12, 0x84, 0x02, 0x0a,
	0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x44,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0x57, 0xc2, 0xb8, 0x18, 0x43,
	0x08, 0x02, 0x10, 0x12, 0x1a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x22, 0x30, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x20, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0xa0, 0xff, 0x2b, 0x0d, 0xb0, 0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0xc0,
	0xff, 0x2b, 0x01, 0x12, 0xf5, 0x01, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x44,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0x4f, 0xc2, 0xb8, 0x18, 0x3b,
	0x08, 0x02, 0x10, 0x12, 0x1a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x22, 0x28, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0xa0, 0xff, 0x2b, 0x03, 0xb0,
	0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0xc0, 0xff, 0x2b, 0x01, 0x12, 0x92, 0x02, 0x0a, 0x0e,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x54,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x1a, 0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0x54, 0xc2, 0xb8, 0x18, 0x44,
	0x08, 0x04, 0x10, 0x12, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69,
	0x64, 0x22, 0x31, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x20, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x20, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0xa0, 0xff, 0x2b, 0x0a, 0xb0, 0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01,
	0x12, 0x83, 0x02, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x54, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x1a, 0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0x4c, 0xc2, 0xb8, 0x18, 0x3c, 0x08, 0x04,
	0x10, 0x12, 0x1a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x22,
	0x29, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0xa0, 0xff, 0x2b, 0x04, 0xb0, 0xff,
	0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0x12, 0xf7, 0x01, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x56, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x49, 0x64, 0x1a, 0x54, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x22, 0x3e, 0xc2, 0xb8, 0x18, 0x36, 0x08, 0x05, 0x10, 0x12, 0x1a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x29, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x20, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0xb0, 0xff, 0x2b, 0x02,
	0x12, 0xf2, 0x01, 0x0a, 0x05, 0x70, 0x61, 0x75, 0x73, 0x65, 0x12, 0x43, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0x4e, 0xc2, 0xb8, 0x18, 0x3a, 0x08, 0x05, 0x10, 0x12, 0x1a,
	0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x27, 0x75, 0x6e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x61,
	0x75, 0x73, 0x65, 0x20, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0xa0, 0xff, 0x2b, 0x05, 0xb0, 0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b,
	0x01, 0xc0, 0xff, 0x2b, 0x01, 0x12, 0xf8, 0x01, 0x0a, 0x07, 0x75, 0x6e, 0x70, 0x61, 0x75, 0x73,
	0x65, 0x12, 0x45, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x6e, 0x50, 0x61,
	0x75, 0x73, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x54, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0x50,
	0xc2, 0xb8, 0x18, 0x3c, 0x08, 0x05, 0x10, 0x12, 0x1a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x29, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x6e, 0x70, 0x61, 0x75, 0x73, 0x65, 0x20, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0xa0, 0xff, 0x2b, 0x06, 0xb0, 0xff, 0x2b, 0x02, 0xb8, 0xff, 0x2b, 0x01, 0xc0, 0xff, 0x2b, 0x01,
	0x1a, 0x04, 0xa0, 0xff, 0x2b, 0x12, 0x42, 0x94, 0x04, 0x0a, 0x52, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0c, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x74, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x43, 0x56, 0x4b, 0x47, 0x53, 0xaa, 0x02,
	0x44, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41,
	0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56,
	0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x44, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5c, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x50, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69,
	0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x47, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x4b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x3a, 0x3a, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto_goTypes = []interface{}{
	(*model.GitlabKubernetes)(nil),         // 0: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	(*model1.ApiResourceDeleteInput)(nil),  // 1: cloud.planton.apis.commons.apiresource.model.ApiResourceDeleteInput
	(*model.GitlabKubernetesId)(nil),       // 2: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetesId
	(*model1.ApiResourcePauseInput)(nil),   // 3: cloud.planton.apis.commons.apiresource.model.ApiResourcePauseInput
	(*model1.ApiResourceUnPauseInput)(nil), // 4: cloud.planton.apis.commons.apiresource.model.ApiResourceUnPauseInput
}
var file_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto_depIdxs = []int32{
	0,  // 0: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.previewCreate:input_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 1: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.create:input_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 2: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.previewUpdate:input_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 3: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.update:input_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	1,  // 4: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.previewDelete:input_type -> cloud.planton.apis.commons.apiresource.model.ApiResourceDeleteInput
	1,  // 5: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.delete:input_type -> cloud.planton.apis.commons.apiresource.model.ApiResourceDeleteInput
	0,  // 6: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.previewRestore:input_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 7: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.restore:input_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	2,  // 8: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.restart:input_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetesId
	3,  // 9: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.pause:input_type -> cloud.planton.apis.commons.apiresource.model.ApiResourcePauseInput
	4,  // 10: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.unpause:input_type -> cloud.planton.apis.commons.apiresource.model.ApiResourceUnPauseInput
	0,  // 11: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.previewCreate:output_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 12: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.create:output_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 13: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.previewUpdate:output_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 14: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.update:output_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 15: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.previewDelete:output_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 16: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.delete:output_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 17: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.previewRestore:output_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 18: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.restore:output_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 19: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.restart:output_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 20: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.pause:output_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	0,  // 21: cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.service.GitlabKubernetesCommandController.unpause:output_type -> cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto_init()
}
func file_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_kubernetes_gitlabkubernetes_service_command_proto_depIdxs = nil
}
