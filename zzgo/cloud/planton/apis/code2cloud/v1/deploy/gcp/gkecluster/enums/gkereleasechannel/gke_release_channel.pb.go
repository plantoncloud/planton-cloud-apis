// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/gcp/gkecluster/enums/gkereleasechannel/gke_release_channel.proto

package gkereleasechannel

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

// note: using uppercase for values since this enum is used as input for google cloud apis
type GkeReleaseChannel int32

const (
	GkeReleaseChannel_unspecified GkeReleaseChannel = 0
	GkeReleaseChannel_REGULAR     GkeReleaseChannel = 1
	GkeReleaseChannel_STABLE      GkeReleaseChannel = 2
	GkeReleaseChannel_RAPID       GkeReleaseChannel = 3
)

// Enum value maps for GkeReleaseChannel.
var (
	GkeReleaseChannel_name = map[int32]string{
		0: "unspecified",
		1: "REGULAR",
		2: "STABLE",
		3: "RAPID",
	}
	GkeReleaseChannel_value = map[string]int32{
		"unspecified": 0,
		"REGULAR":     1,
		"STABLE":      2,
		"RAPID":       3,
	}
)

func (x GkeReleaseChannel) Enum() *GkeReleaseChannel {
	p := new(GkeReleaseChannel)
	*p = x
	return p
}

func (x GkeReleaseChannel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GkeReleaseChannel) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_enumTypes[0].Descriptor()
}

func (GkeReleaseChannel) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_enumTypes[0]
}

func (x GkeReleaseChannel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GkeReleaseChannel.Descriptor instead.
func (GkeReleaseChannel) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_rawDesc = []byte{
	0x0a, 0x68, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x6b,
	0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x67,
	0x6b, 0x65, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2f, 0x67, 0x6b, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x4e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x67, 0x6b, 0x65, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2a, 0x48, 0x0a, 0x11, 0x47, 0x6b,
	0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x0f, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x53, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x41, 0x50,
	0x49, 0x44, 0x10, 0x03, 0x42, 0xde, 0x04, 0x0a, 0x5c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x67,
	0x63, 0x70, 0x2e, 0x67, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x67, 0x6b, 0x65, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x16, 0x47, 0x6b, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x7e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x6b, 0x65,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x67, 0x6b,
	0x65, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0xa2,
	0x02, 0x0a, 0x43, 0x50, 0x41, 0x43, 0x56, 0x44, 0x47, 0x47, 0x45, 0x47, 0xaa, 0x02, 0x4e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69,
	0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x47, 0x63, 0x70, 0x2e, 0x47, 0x6b, 0x65, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x47, 0x6b, 0x65, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0xca, 0x02, 0x4e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31,
	0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x6b, 0x65, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x47, 0x6b, 0x65,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0xe2, 0x02,
	0x5a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56,
	0x31, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x6b, 0x65,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x47, 0x6b,
	0x65, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x57, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x56, 0x31, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x3a,
	0x3a, 0x47, 0x6b, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x3a, 0x3a, 0x47, 0x6b, 0x65, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_goTypes = []interface{}{
	(GkeReleaseChannel)(0), // 0: cloud.planton.apis.code2cloud.v1.deploy.gcp.gkecluster.enums.gkereleasechannel.GkeReleaseChannel
}
var file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_init()
}
func file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gkecluster_enums_gkereleasechannel_gke_release_channel_proto_depIdxs = nil
}
