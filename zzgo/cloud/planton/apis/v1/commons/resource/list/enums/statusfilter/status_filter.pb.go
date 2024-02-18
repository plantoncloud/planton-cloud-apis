// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/commons/resource/list/enums/statusfilter/status_filter.proto

package statusfilter

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusFilter int32

const (
	StatusFilter_STATUS_FILTER_UNSPECIFIED StatusFilter = 0
	// This value indicates that the search operation should return all resources, regardless of their active status.
	// This option allows clients to see a complete overview of resources, both active and inactive.
	StatusFilter_all StatusFilter = 1
	// When this value is used, the search operation will only return resources that are currently active.
	// This is suitable for scenarios where the client is only interested in resources that are in use or available for use
	StatusFilter_active_only StatusFilter = 2
	// When this value is selected, the search operation will return only resources that are currently inactive.
	// This option is particularly useful when the client wants to review or clean up resources that are no longer in use
	StatusFilter_inactive_only StatusFilter = 3
)

// Enum value maps for StatusFilter.
var (
	StatusFilter_name = map[int32]string{
		0: "STATUS_FILTER_UNSPECIFIED",
		1: "all",
		2: "active_only",
		3: "inactive_only",
	}
	StatusFilter_value = map[string]int32{
		"STATUS_FILTER_UNSPECIFIED": 0,
		"all":                       1,
		"active_only":               2,
		"inactive_only":             3,
	}
)

func (x StatusFilter) Enum() *StatusFilter {
	p := new(StatusFilter)
	*p = x
	return p
}

func (x StatusFilter) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusFilter) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_enumTypes[0].Descriptor()
}

func (StatusFilter) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_enumTypes[0]
}

func (x StatusFilter) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusFilter.Descriptor instead.
func (StatusFilter) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_rawDesc = []byte{
	0x0a, 0x52, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2a, 0x5a, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x10, 0x03,
	0x42, 0xf9, 0x03, 0x0a, 0x4c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67,
	0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x42, 0x11, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x56, 0x43, 0x52, 0x4c,
	0x45, 0x53, 0xaa, 0x02, 0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0xca, 0x02, 0x3f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x4c, 0x69, 0x73,
	0x74, 0x5f, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0xe2, 0x02, 0x4b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x4c,
	0x69, 0x73, 0x74, 0x5f, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x46, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x3a, 0x3a, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_rawDescData = file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_goTypes = []interface{}{
	(StatusFilter)(0), // 0: cloud.planton.apis.commons.resource.list.enums.statusfilter.StatusFilter
}
var file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_init()
}
func file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_init() {
	if File_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto = out.File
	file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_resource_list_enums_statusfilter_status_filter_proto_depIdxs = nil
}
