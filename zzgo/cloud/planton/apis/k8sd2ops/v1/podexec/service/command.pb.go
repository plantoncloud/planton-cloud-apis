// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/k8sd2ops/v1/podexec/service/command.proto

package service

import (
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/kubernetesobject/model"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/k8sd2ops/v1/podexec/model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x4c,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x64,
	0x65, 0x78, 0x65, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xd6, 0x04, 0x0a, 0x2d, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x6f, 0x64, 0x45, 0x78, 0x65,
	0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0xc6, 0x01, 0x0a, 0x14, 0x65, 0x78, 0x65, 0x63, 0x49, 0x6e, 0x74, 0x6f, 0x50,
	0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x47, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x64,
	0x65, 0x78, 0x65, 0x63, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x49,
	0x6e, 0x74, 0x6f, 0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x49,
	0x6e, 0x74, 0x6f, 0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0xbb, 0x01, 0x0a, 0x1b,
	0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x49, 0x6e, 0x74, 0x6f, 0x50,
	0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x47, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x64,
	0x65, 0x78, 0x65, 0x63, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x49,
	0x6e, 0x74, 0x6f, 0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x51, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f,
	0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x49,
	0x6e, 0x74, 0x6f, 0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x9d, 0x01, 0x0a, 0x27, 0x62, 0x72,
	0x6f, 0x77, 0x73, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x4e, 0x65, 0x78, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x5a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x64, 0x32,
	0x6f, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x4e, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e,
	0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x8e, 0x03, 0x0a, 0x3c, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b,
	0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x64, 0x65, 0x78,
	0x65, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b,
	0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x64, 0x65, 0x78,
	0x65, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41,
	0x4b, 0x56, 0x50, 0x53, 0xaa, 0x02, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x64, 0x32, 0x6f,
	0x70, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x4b, 0x38, 0x73, 0x64, 0x32,
	0x6f, 0x70, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x4b, 0x38, 0x73, 0x64,
	0x32, 0x6f, 0x70, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x34, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x4b, 0x38, 0x73, 0x64,
	0x32, 0x6f, 0x70, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x50, 0x6f, 0x64, 0x65, 0x78, 0x65,
	0x63, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto_goTypes = []interface{}{
	(*model.ExecIntoPodContainerInput)(nil),                    // 0: cloud.planton.apis.k8sd2ops.v1.podexec.model.ExecIntoPodContainerInput
	(*model.BrowserExecuteNextCommandInPodContainerInput)(nil), // 1: cloud.planton.apis.k8sd2ops.v1.podexec.model.BrowserExecuteNextCommandInPodContainerInput
	(*model1.ExecIntoPodContainerResponse)(nil),                // 2: cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.model.ExecIntoPodContainerResponse
	(*model.BrowserExecIntoPodContainerResponse)(nil),          // 3: cloud.planton.apis.k8sd2ops.v1.podexec.model.BrowserExecIntoPodContainerResponse
	(*emptypb.Empty)(nil),                                      // 4: google.protobuf.Empty
}
var file_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.k8sd2ops.v1.podexec.service.ApiResourceKubernetesPodExecCommandController.execIntoPodContainer:input_type -> cloud.planton.apis.k8sd2ops.v1.podexec.model.ExecIntoPodContainerInput
	0, // 1: cloud.planton.apis.k8sd2ops.v1.podexec.service.ApiResourceKubernetesPodExecCommandController.browserExecIntoPodContainer:input_type -> cloud.planton.apis.k8sd2ops.v1.podexec.model.ExecIntoPodContainerInput
	1, // 2: cloud.planton.apis.k8sd2ops.v1.podexec.service.ApiResourceKubernetesPodExecCommandController.browserExecuteNextCommandInPodContainer:input_type -> cloud.planton.apis.k8sd2ops.v1.podexec.model.BrowserExecuteNextCommandInPodContainerInput
	2, // 3: cloud.planton.apis.k8sd2ops.v1.podexec.service.ApiResourceKubernetesPodExecCommandController.execIntoPodContainer:output_type -> cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.model.ExecIntoPodContainerResponse
	3, // 4: cloud.planton.apis.k8sd2ops.v1.podexec.service.ApiResourceKubernetesPodExecCommandController.browserExecIntoPodContainer:output_type -> cloud.planton.apis.k8sd2ops.v1.podexec.model.BrowserExecIntoPodContainerResponse
	4, // 5: cloud.planton.apis.k8sd2ops.v1.podexec.service.ApiResourceKubernetesPodExecCommandController.browserExecuteNextCommandInPodContainer:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto_init() }
func file_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto_init() {
	if File_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto = out.File
	file_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto_rawDesc = nil
	file_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto_goTypes = nil
	file_cloud_planton_apis_k8sd2ops_v1_podexec_service_command_proto_depIdxs = nil
}
