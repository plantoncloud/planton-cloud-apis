// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/commons/kubernetes/enums/enums.proto

package enums

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

type PodControllerType int32

const (
	PodControllerType_POD_CONTROLLER_TYPE_UNSPECIFIED PodControllerType = 0
	PodControllerType_deployment                      PodControllerType = 1
	PodControllerType_stateful_set                    PodControllerType = 2
	PodControllerType_cron_job                        PodControllerType = 3
)

// Enum value maps for PodControllerType.
var (
	PodControllerType_name = map[int32]string{
		0: "POD_CONTROLLER_TYPE_UNSPECIFIED",
		1: "deployment",
		2: "stateful_set",
		3: "cron_job",
	}
	PodControllerType_value = map[string]int32{
		"POD_CONTROLLER_TYPE_UNSPECIFIED": 0,
		"deployment":                      1,
		"stateful_set":                    2,
		"cron_job":                        3,
	}
)

func (x PodControllerType) Enum() *PodControllerType {
	p := new(PodControllerType)
	*p = x
	return p
}

func (x PodControllerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PodControllerType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_enumTypes[0].Descriptor()
}

func (PodControllerType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_enumTypes[0]
}

func (x PodControllerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PodControllerType.Descriptor instead.
func (PodControllerType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDescGZIP(), []int{0}
}

// kubernetes resource kind
type ResourceKind int32

const (
	ResourceKind_RESOURCE_KIND_UNSPECIFIED ResourceKind = 0
	ResourceKind_Namespace                 ResourceKind = 1
	ResourceKind_Deployment                ResourceKind = 2
	ResourceKind_Service                   ResourceKind = 3
	ResourceKind_Secret                    ResourceKind = 4
	ResourceKind_Ingress                   ResourceKind = 5
	ResourceKind_Issuer                    ResourceKind = 6
	ResourceKind_Certificate               ResourceKind = 7
)

// Enum value maps for ResourceKind.
var (
	ResourceKind_name = map[int32]string{
		0: "RESOURCE_KIND_UNSPECIFIED",
		1: "Namespace",
		2: "Deployment",
		3: "Service",
		4: "Secret",
		5: "Ingress",
		6: "Issuer",
		7: "Certificate",
	}
	ResourceKind_value = map[string]int32{
		"RESOURCE_KIND_UNSPECIFIED": 0,
		"Namespace":                 1,
		"Deployment":                2,
		"Service":                   3,
		"Secret":                    4,
		"Ingress":                   5,
		"Issuer":                    6,
		"Certificate":               7,
	}
)

func (x ResourceKind) Enum() *ResourceKind {
	p := new(ResourceKind)
	*p = x
	return p
}

func (x ResourceKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceKind) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_enumTypes[1].Descriptor()
}

func (ResourceKind) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_enumTypes[1]
}

func (x ResourceKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceKind.Descriptor instead.
func (ResourceKind) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDescGZIP(), []int{1}
}

var File_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0x68, 0x0a, 0x11,
	0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x23, 0x0a, 0x1f, 0x50, 0x4f, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c,
	0x4c, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66,
	0x75, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x72, 0x6f, 0x6e,
	0x5f, 0x6a, 0x6f, 0x62, 0x10, 0x03, 0x2a, 0x8f, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x10, 0x07, 0x42, 0x82, 0x03, 0x0a, 0x32, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42,
	0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x07,
	0x43, 0x50, 0x41, 0x56, 0x43, 0x4b, 0x45, 0xaa, 0x02, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xe2, 0x02, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56,
	0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x34, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDescData = file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_goTypes = []interface{}{
	(PodControllerType)(0), // 0: cloud.planton.apis.v1.commons.kubernetes.enums.PodControllerType
	(ResourceKind)(0),      // 1: cloud.planton.apis.v1.commons.kubernetes.enums.ResourceKind
}
var file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_init() }
func file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_init() {
	if File_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto = out.File
	file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_kubernetes_enums_enums_proto_depIdxs = nil
}
