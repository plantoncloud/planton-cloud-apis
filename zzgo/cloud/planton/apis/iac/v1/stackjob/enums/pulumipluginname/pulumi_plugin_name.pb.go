// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/iac/v1/stackjob/enums/pulumipluginname/pulumi_plugin_name.proto

package pulumipluginname

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

// pulumi plugin name
// the recommended best practice to prefix the entry with enum name has been intentionally ignored to ensure that the enum's value matches the name of the pulumi plugin.
type PulumiPluginName int32

const (
	PulumiPluginName_unspecified   PulumiPluginName = 0
	PulumiPluginName_RANDOM        PulumiPluginName = 1
	PulumiPluginName_KUBERNETES    PulumiPluginName = 2
	PulumiPluginName_GCP           PulumiPluginName = 3
	PulumiPluginName_GOOGLE_NATIVE PulumiPluginName = 4
	PulumiPluginName_AWS           PulumiPluginName = 5
	PulumiPluginName_AWS_NATIVE    PulumiPluginName = 6
	PulumiPluginName_AZURE         PulumiPluginName = 7
	PulumiPluginName_AZURE_NATIVE  PulumiPluginName = 8
)

// Enum value maps for PulumiPluginName.
var (
	PulumiPluginName_name = map[int32]string{
		0: "unspecified",
		1: "RANDOM",
		2: "KUBERNETES",
		3: "GCP",
		4: "GOOGLE_NATIVE",
		5: "AWS",
		6: "AWS_NATIVE",
		7: "AZURE",
		8: "AZURE_NATIVE",
	}
	PulumiPluginName_value = map[string]int32{
		"unspecified":   0,
		"RANDOM":        1,
		"KUBERNETES":    2,
		"GCP":           3,
		"GOOGLE_NATIVE": 4,
		"AWS":           5,
		"AWS_NATIVE":    6,
		"AZURE":         7,
		"AZURE_NATIVE":  8,
	}
)

func (x PulumiPluginName) Enum() *PulumiPluginName {
	p := new(PulumiPluginName)
	*p = x
	return p
}

func (x PulumiPluginName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PulumiPluginName) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_enumTypes[0].Descriptor()
}

func (PulumiPluginName) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_enumTypes[0]
}

func (x PulumiPluginName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PulumiPluginName.Descriptor instead.
func (PulumiPluginName) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_rawDesc = []byte{
	0x0a, 0x52, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d,
	0x69, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x75, 0x6c, 0x75,
	0x6d, 0x69, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x75, 0x6c, 0x75, 0x6d, 0x69, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x2a,
	0x91, 0x01, 0x0a, 0x10, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x10,
	0x02, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x43, 0x50, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x4f,
	0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12, 0x07, 0x0a,
	0x03, 0x41, 0x57, 0x53, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x57, 0x53, 0x5f, 0x4e, 0x41,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x5a, 0x55, 0x52, 0x45, 0x10,
	0x07, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x5a, 0x55, 0x52, 0x45, 0x5f, 0x4e, 0x41, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x08, 0x42, 0xdb, 0x03, 0x0a, 0x47, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x75, 0x6c, 0x75, 0x6d, 0x69, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x15, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x6e,
	0x61, 0x6d, 0x65, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x49, 0x56, 0x53, 0x45, 0x50, 0xaa, 0x02,
	0x39, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41,
	0x70, 0x69, 0x73, 0x2e, 0x49, 0x61, 0x63, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x6a, 0x6f, 0x62, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0xca, 0x02, 0x39, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x49, 0x61, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x5c,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0xe2, 0x02, 0x45, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x49, 0x61, 0x63, 0x5c,
	0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x5c, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0x5c, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x6e, 0x61,
	0x6d, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x40, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x49, 0x61, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a,
	0x3a, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x6e, 0x61, 0x6d,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_rawDescData = file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_rawDesc
)

func file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_rawDescData)
	})
	return file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_rawDescData
}

var file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_goTypes = []interface{}{
	(PulumiPluginName)(0), // 0: cloud.planton.apis.iac.v1.stackjob.enums.pulumipluginname.PulumiPluginName
}
var file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_init()
}
func file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_init() {
	if File_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto = out.File
	file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_rawDesc = nil
	file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_goTypes = nil
	file_cloud_planton_apis_iac_v1_stackjob_enums_pulumipluginname_pulumi_plugin_name_proto_depIdxs = nil
}
