// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/azure/akscluster/service/stack.proto

package service

import (
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/azure/akscluster/model"
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

var File_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e,
	0x61, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x6b, 0x73, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xca, 0x01, 0x0a, 0x19, 0x41, 0x6b, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xac, 0x01, 0x0a, 0x07, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x12, 0x4d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x6b, 0x73, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x6b, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x50, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x6b, 0x73, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x6b, 0x73, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x42, 0xd0, 0x03, 0x0a, 0x47, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x6b,
	0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x42, 0x0a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x69,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x43,
	0x56, 0x41, 0x41, 0x53, 0xaa, 0x02, 0x39, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x41, 0x6b,
	0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0xca, 0x02, 0x39, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x56, 0x31, 0x5c, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x5c, 0x41, 0x6b, 0x73, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x45, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69,
	0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c,
	0x41, 0x7a, 0x75, 0x72, 0x65, 0x5c, 0x41, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x40, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x41, 0x7a, 0x75,
	0x72, 0x65, 0x3a, 0x3a, 0x41, 0x6b, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x3a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto_goTypes = []interface{}{
	(*model.AksClusterStackInput)(nil),    // 0: cloud.planton.apis.code2cloud.v1.azure.akscluster.model.AksClusterStackInput
	(*model.AksClusterStackResponse)(nil), // 1: cloud.planton.apis.code2cloud.v1.azure.akscluster.model.AksClusterStackResponse
}
var file_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.code2cloud.v1.azure.akscluster.service.AksClusterStackController.execute:input_type -> cloud.planton.apis.code2cloud.v1.azure.akscluster.model.AksClusterStackInput
	1, // 1: cloud.planton.apis.code2cloud.v1.azure.akscluster.service.AksClusterStackController.execute:output_type -> cloud.planton.apis.code2cloud.v1.azure.akscluster.model.AksClusterStackResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_azure_akscluster_service_stack_proto_depIdxs = nil
}
