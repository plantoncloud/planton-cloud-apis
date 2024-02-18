// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/cloudaccount/enums/kubernetesprovider/kubernetes_provider.proto

package kubernetesprovider

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

// kubernetes provider
type KubernetesProvider int32

const (
	KubernetesProvider_KUBERNETES_PROVIDER_UNSPECIFIED KubernetesProvider = 0
	// https://cloud.google.com/kubernetes-engine
	KubernetesProvider_gcp_gke KubernetesProvider = 1
	// https://aws.amazon.com/eks/
	KubernetesProvider_aws_eks KubernetesProvider = 2
	// https://azure.microsoft.com/en-us/products/kubernetes-service
	KubernetesProvider_azure_aks KubernetesProvider = 3
	// https://docs.digitalocean.com/products/kubernetes/
	KubernetesProvider_digital_ocean_doks KubernetesProvider = 4
)

// Enum value maps for KubernetesProvider.
var (
	KubernetesProvider_name = map[int32]string{
		0: "KUBERNETES_PROVIDER_UNSPECIFIED",
		1: "gcp_gke",
		2: "aws_eks",
		3: "azure_aks",
		4: "digital_ocean_doks",
	}
	KubernetesProvider_value = map[string]int32{
		"KUBERNETES_PROVIDER_UNSPECIFIED": 0,
		"gcp_gke":                         1,
		"aws_eks":                         2,
		"azure_aks":                       3,
		"digital_ocean_doks":              4,
	}
)

func (x KubernetesProvider) Enum() *KubernetesProvider {
	p := new(KubernetesProvider)
	*p = x
	return p
}

func (x KubernetesProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KubernetesProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_enumTypes[0].Descriptor()
}

func (KubernetesProvider) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_enumTypes[0]
}

func (x KubernetesProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KubernetesProvider.Descriptor instead.
func (KubernetesProvider) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_rawDesc = []byte{
	0x0a, 0x60, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x46, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2a, 0x7a, 0x0a, 0x12, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x23, 0x0a, 0x1f, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x5f, 0x50,
	0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x67, 0x63, 0x70, 0x5f, 0x67, 0x6b, 0x65,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x61, 0x77, 0x73, 0x5f, 0x65, 0x6b, 0x73, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x61, 0x6b, 0x73, 0x10, 0x03, 0x12, 0x16,
	0x0a, 0x12, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x5f, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x5f,
	0x64, 0x6f, 0x6b, 0x73, 0x10, 0x04, 0x42, 0xab, 0x04, 0x0a, 0x54, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42,
	0x17, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x76, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x43, 0x56, 0x43, 0x45, 0x4b, 0xaa, 0x02, 0x46,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70,
	0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0xca, 0x02, 0x46, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0xe2,
	0x02, 0x52, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x56, 0x31, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x4d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_goTypes = []interface{}{
	(KubernetesProvider)(0), // 0: cloud.planton.apis.code2cloud.v1.cloudaccount.enums.kubernetesprovider.KubernetesProvider
}
var file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_init()
}
func file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_cloudaccount_enums_kubernetesprovider_kubernetes_provider_proto_depIdxs = nil
}
