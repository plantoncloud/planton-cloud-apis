// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/integration/v1/kubernetes/kubernetesobject/enums/podmanagertype/pod_manager_type.proto

package podmanagertype

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

type PodManagerType int32

const (
	PodManagerType_unspecified  PodManagerType = 0
	PodManagerType_deployment   PodManagerType = 1
	PodManagerType_stateful_set PodManagerType = 2
	PodManagerType_cron_job     PodManagerType = 3
)

// Enum value maps for PodManagerType.
var (
	PodManagerType_name = map[int32]string{
		0: "unspecified",
		1: "deployment",
		2: "stateful_set",
		3: "cron_job",
	}
	PodManagerType_value = map[string]int32{
		"unspecified":  0,
		"deployment":   1,
		"stateful_set": 2,
		"cron_job":     3,
	}
)

func (x PodManagerType) Enum() *PodManagerType {
	p := new(PodManagerType)
	*p = x
	return p
}

func (x PodManagerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PodManagerType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_enumTypes[0].Descriptor()
}

func (PodManagerType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_enumTypes[0]
}

func (x PodManagerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PodManagerType.Descriptor instead.
func (PodManagerType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_rawDesc = []byte{
	0x0a, 0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x6f, 0x64, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x6f, 0x64, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x52, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x70, 0x6f, 0x64, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x2a,
	0x51, 0x0a, 0x0e, 0x50, 0x6f, 0x64, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73,
	0x65, 0x74, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x6a, 0x6f, 0x62,
	0x10, 0x03, 0x42, 0xf2, 0x04, 0x0a, 0x60, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x6f, 0x64, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x42, 0x13, 0x50, 0x6f, 0x64, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x82,
	0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x6f, 0x64, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x74,
	0x79, 0x70, 0x65, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x49, 0x56, 0x4b, 0x4b, 0x45, 0x50, 0xaa,
	0x02, 0x52, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x41, 0x70, 0x69, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x56, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x6f, 0x64, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x74, 0x79, 0x70, 0x65, 0xca, 0x02, 0x52, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x50, 0x6f, 0x64, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0xe2, 0x02, 0x5e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c,
	0x50, 0x6f, 0x64, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x5a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69,
	0x73, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a,
	0x56, 0x31, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x50, 0x6f, 0x64, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_rawDescData = file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_rawDesc
)

func file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_rawDescData)
	})
	return file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_rawDescData
}

var file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_goTypes = []interface{}{
	(PodManagerType)(0), // 0: cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.enums.podmanagertype.PodManagerType
}
var file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_init()
}
func file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_init() {
	if File_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto = out.File
	file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_rawDesc = nil
	file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_goTypes = nil
	file_cloud_planton_apis_integration_v1_kubernetes_kubernetesobject_enums_podmanagertype_pod_manager_type_proto_depIdxs = nil
}
