// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/stack/service.proto

package stack

import (
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
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

var File_cloud_planton_apis_v1_stack_service_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_stack_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x1a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb4, 0x03, 0x0a, 0x16, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x50, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x1a, 0x22,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x12, 0x50, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x1a, 0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x12, 0x73, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x45,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x80, 0x01, 0x0a, 0x13, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x6e, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x12, 0x45, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x32, 0xd6, 0x01, 0x0a,
	0x14, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x24, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x1a, 0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x69, 0x0a, 0x0f, 0x67, 0x65,
	0x74, 0x42, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x32, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x1a, 0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x42, 0x98, 0x02, 0x0a, 0x29, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0xa2, 0x02, 0x05, 0x43, 0x50, 0x41, 0x56, 0x53, 0xaa, 0x02, 0x1b, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0xca, 0x02, 0x1b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0xe2, 0x02, 0x27, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_stack_service_proto_goTypes = []interface{}{
	(*Stack)(nil), // 0: cloud.planton.apis.v1.stack.Stack
	(*resource.ApiResourceDeleteCommandInput)(nil), // 1: cloud.planton.apis.v1.commons.resource.ApiResourceDeleteCommandInput
	(*StackId)(nil),             // 2: cloud.planton.apis.v1.stack.StackId
	(*resource.ResourceId)(nil), // 3: cloud.planton.apis.v1.commons.resource.ResourceId
}
var file_cloud_planton_apis_v1_stack_service_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.stack.StackCommandController.create:input_type -> cloud.planton.apis.v1.stack.Stack
	0, // 1: cloud.planton.apis.v1.stack.StackCommandController.update:input_type -> cloud.planton.apis.v1.stack.Stack
	1, // 2: cloud.planton.apis.v1.stack.StackCommandController.delete:input_type -> cloud.planton.apis.v1.commons.resource.ApiResourceDeleteCommandInput
	1, // 3: cloud.planton.apis.v1.stack.StackCommandController.deleteOnPulumiCloud:input_type -> cloud.planton.apis.v1.commons.resource.ApiResourceDeleteCommandInput
	2, // 4: cloud.planton.apis.v1.stack.StackQueryController.getById:input_type -> cloud.planton.apis.v1.stack.StackId
	3, // 5: cloud.planton.apis.v1.stack.StackQueryController.getByResourceId:input_type -> cloud.planton.apis.v1.commons.resource.ResourceId
	0, // 6: cloud.planton.apis.v1.stack.StackCommandController.create:output_type -> cloud.planton.apis.v1.stack.Stack
	0, // 7: cloud.planton.apis.v1.stack.StackCommandController.update:output_type -> cloud.planton.apis.v1.stack.Stack
	0, // 8: cloud.planton.apis.v1.stack.StackCommandController.delete:output_type -> cloud.planton.apis.v1.stack.Stack
	0, // 9: cloud.planton.apis.v1.stack.StackCommandController.deleteOnPulumiCloud:output_type -> cloud.planton.apis.v1.stack.Stack
	0, // 10: cloud.planton.apis.v1.stack.StackQueryController.getById:output_type -> cloud.planton.apis.v1.stack.Stack
	0, // 11: cloud.planton.apis.v1.stack.StackQueryController.getByResourceId:output_type -> cloud.planton.apis.v1.stack.Stack
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_stack_service_proto_init() }
func file_cloud_planton_apis_v1_stack_service_proto_init() {
	if File_cloud_planton_apis_v1_stack_service_proto != nil {
		return
	}
	file_cloud_planton_apis_v1_stack_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_stack_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cloud_planton_apis_v1_stack_service_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_stack_service_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_stack_service_proto = out.File
	file_cloud_planton_apis_v1_stack_service_proto_rawDesc = nil
	file_cloud_planton_apis_v1_stack_service_proto_goTypes = nil
	file_cloud_planton_apis_v1_stack_service_proto_depIdxs = nil
}
