// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/billing/resource/version/rpc/service.proto

package rpc

import (
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/version/rpc"
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

var File_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x91, 0x03, 0x0a, 0x1e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xb5, 0x01, 0x0a, 0x0d, 0x6c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x5b, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x47, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0xb6, 0x01, 0x0a, 0x16, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x57, 0x69, 0x74,
	0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x57, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x69, 0x7a, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x43, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0xa8, 0x03, 0x0a, 0x40, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x42,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x62, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x72, 0x70, 0x63, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x42, 0x52, 0x56, 0x52, 0xaa, 0x02,
	0x32, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41,
	0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x70, 0x63, 0xca, 0x02, 0x32, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x52, 0x70, 0x63, 0xe2, 0x02, 0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x39, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto_goTypes = []interface{}{
	(*rpc.ListResourceVersionsByFiltersQueryInput)(nil), // 0: cloud.planton.apis.v1.commons.resource.version.rpc.ListResourceVersionsByFiltersQueryInput
	(*rpc.ResourceVersionWithContextSizeInput)(nil),     // 1: cloud.planton.apis.v1.commons.resource.version.rpc.ResourceVersionWithContextSizeInput
	(*rpc.ResourceVersionList)(nil),                     // 2: cloud.planton.apis.v1.commons.resource.version.rpc.ResourceVersionList
	(*rpc.ResourceVersion)(nil),                         // 3: cloud.planton.apis.v1.commons.resource.version.rpc.ResourceVersion
}
var file_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.billing.resource.version.rpc.ResourceVersionQueryController.listByFilters:input_type -> cloud.planton.apis.v1.commons.resource.version.rpc.ListResourceVersionsByFiltersQueryInput
	1, // 1: cloud.planton.apis.v1.billing.resource.version.rpc.ResourceVersionQueryController.getByIdWithContextSize:input_type -> cloud.planton.apis.v1.commons.resource.version.rpc.ResourceVersionWithContextSizeInput
	2, // 2: cloud.planton.apis.v1.billing.resource.version.rpc.ResourceVersionQueryController.listByFilters:output_type -> cloud.planton.apis.v1.commons.resource.version.rpc.ResourceVersionList
	3, // 3: cloud.planton.apis.v1.billing.resource.version.rpc.ResourceVersionQueryController.getByIdWithContextSize:output_type -> cloud.planton.apis.v1.commons.resource.version.rpc.ResourceVersion
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto_init() }
func file_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto_init() {
	if File_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto_depIdxs,
	}.Build()
	File_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto = out.File
	file_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto_rawDesc = nil
	file_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto_goTypes = nil
	file_cloud_planton_apis_v1_billing_resource_version_rpc_service_proto_depIdxs = nil
}
