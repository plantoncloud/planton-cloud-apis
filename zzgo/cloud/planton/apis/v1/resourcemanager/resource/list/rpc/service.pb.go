// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/resourcemanager/resource/list/rpc/service.proto

package rpc

import (
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/list/rpc"
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

var File_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x72, 0x70, 0x63,
	0x1a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbe, 0x01,
	0x0a, 0x1f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69,
	0x65, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x12, 0x9a, 0x01, 0x0a, 0x12, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x47, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0xbe,
	0x03, 0x0a, 0x3b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x0c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x67,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x2f, 0x72, 0x70, 0x63, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x52, 0x52,
	0x4c, 0x52, 0xaa, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x52, 0x70, 0x63, 0xca, 0x02, 0x38, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69,
	0x73, 0x5c, 0x56, 0x31, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x4c, 0x69,
	0x73, 0x74, 0x5f, 0x5c, 0x52, 0x70, 0x63, 0xe2, 0x02, 0x44, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x4c, 0x69, 0x73, 0x74, 0x5f, 0x5c, 0x52,
	0x70, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x3a, 0x3a, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto_goTypes = []interface{}{
	(*rpc.GetByResourceTypeInput)(nil), // 0: cloud.planton.apis.v1.commons.resource.list.rpc.GetByResourceTypeInput
	(*rpc.RecordList)(nil),             // 1: cloud.planton.apis.v1.commons.resource.list.rpc.RecordList
}
var file_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.resourcemanager.resource.list.rpc.ResourceListViewQueryController.listByResourceType:input_type -> cloud.planton.apis.v1.commons.resource.list.rpc.GetByResourceTypeInput
	1, // 1: cloud.planton.apis.v1.resourcemanager.resource.list.rpc.ResourceListViewQueryController.listByResourceType:output_type -> cloud.planton.apis.v1.commons.resource.list.rpc.RecordList
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto_init() }
func file_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto_init() {
	if File_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto = out.File
	file_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto_rawDesc = nil
	file_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto_goTypes = nil
	file_cloud_planton_apis_v1_resourcemanager_resource_list_rpc_service_proto_depIdxs = nil
}
