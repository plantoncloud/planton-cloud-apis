// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/cloudaccount/enums/cloudprovider/cloud_provider.proto

package cloudprovider

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

// cloud-provider
type CloudProvider int32

const (
	// default value when cloud provider name is not set
	CloudProvider_unspecified CloudProvider = 0
	// amazon web services cloud provider
	CloudProvider_aws CloudProvider = 1
	// google cloud platform cloud provider
	CloudProvider_gcp CloudProvider = 2
	// microsoft azure cloud provider
	CloudProvider_azure CloudProvider = 3
	// alibaba cloud provider
	CloudProvider_alibaba CloudProvider = 4
	// oracle cloud provider
	CloudProvider_oracle CloudProvider = 5
)

// Enum value maps for CloudProvider.
var (
	CloudProvider_name = map[int32]string{
		0: "unspecified",
		1: "aws",
		2: "gcp",
		3: "azure",
		4: "alibaba",
		5: "oracle",
	}
	CloudProvider_value = map[string]int32{
		"unspecified": 0,
		"aws":         1,
		"gcp":         2,
		"azure":       3,
		"alibaba":     4,
		"oracle":      5,
	}
)

func (x CloudProvider) Enum() *CloudProvider {
	p := new(CloudProvider)
	*p = x
	return p
}

func (x CloudProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CloudProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_enumTypes[0].Descriptor()
}

func (CloudProvider) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_enumTypes[0]
}

func (x CloudProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CloudProvider.Descriptor instead.
func (CloudProvider) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_rawDesc = []byte{
	0x0a, 0x56, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2a, 0x56, 0x0a, 0x0d, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x0a, 0x0b,
	0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x61, 0x77, 0x73, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x67, 0x63, 0x70, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x61, 0x6c,
	0x69, 0x62, 0x61, 0x62, 0x61, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x10, 0x05, 0x42, 0x88, 0x04, 0x0a, 0x4f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x12, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x71, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x43, 0x56, 0x43, 0x45, 0x43, 0xaa, 0x02, 0x41, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0xca,
	0x02, 0x41, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x56, 0x31, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0xe2, 0x02, 0x4d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x48, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_goTypes = []interface{}{
	(CloudProvider)(0), // 0: cloud.planton.apis.code2cloud.v1.cloudaccount.enums.cloudprovider.CloudProvider
}
var file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_init()
}
func file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_cloudprovider_cloud_provider_proto_depIdxs = nil
}
