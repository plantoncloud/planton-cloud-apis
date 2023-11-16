// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/resourcemanager/company/enums/event/enums.proto

package event

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/enums/options"
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

type CompanyEventType int32

const (
	CompanyEventType_COMPANY_EVENT_TYPE_UNSPECIFIED CompanyEventType = 0
	CompanyEventType_created                        CompanyEventType = 1
	CompanyEventType_updated                        CompanyEventType = 2
	CompanyEventType_deleted                        CompanyEventType = 3
	CompanyEventType_restored                       CompanyEventType = 4
)

// Enum value maps for CompanyEventType.
var (
	CompanyEventType_name = map[int32]string{
		0: "COMPANY_EVENT_TYPE_UNSPECIFIED",
		1: "created",
		2: "updated",
		3: "deleted",
		4: "restored",
	}
	CompanyEventType_value = map[string]int32{
		"COMPANY_EVENT_TYPE_UNSPECIFIED": 0,
		"created":                        1,
		"updated":                        2,
		"deleted":                        3,
		"restored":                       4,
	}
)

func (x CompanyEventType) Enum() *CompanyEventType {
	p := new(CompanyEventType)
	*p = x
	return p
}

func (x CompanyEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompanyEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_enumTypes[0].Descriptor()
}

func (CompanyEventType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_enumTypes[0]
}

func (x CompanyEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompanyEventType.Descriptor instead.
func (CompanyEventType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x5b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0x97, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e, 0x59, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x10, 0x01, 0x1a, 0x0c, 0x80, 0xf9, 0x2b, 0x01, 0x88, 0xf9, 0x2b, 0x01, 0x90,
	0xf9, 0x2b, 0x01, 0x12, 0x15, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x10, 0x02,
	0x1a, 0x08, 0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9, 0x2b, 0x01, 0x12, 0x15, 0x0a, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x03, 0x1a, 0x08, 0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9, 0x2b,
	0x01, 0x12, 0x16, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x10, 0x04, 0x1a,
	0x08, 0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9, 0x2b, 0x01, 0x42, 0xd0, 0x03, 0x0a, 0x47, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0xa2, 0x02,
	0x08, 0x43, 0x50, 0x41, 0x56, 0x52, 0x43, 0x45, 0x45, 0xaa, 0x02, 0x39, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x39, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0xe2, 0x02, 0x45, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x40, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x3a, 0x3a,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_rawDescData = file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_rawDesc
)

func file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_rawDescData
}

var file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_goTypes = []interface{}{
	(CompanyEventType)(0), // 0: cloud.planton.apis.v1.resourcemanager.company.enums.event.CompanyEventType
}
var file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_init() }
func file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_init() {
	if File_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto = out.File
	file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_rawDesc = nil
	file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_goTypes = nil
	file_cloud_planton_apis_v1_resourcemanager_company_enums_event_enums_proto_depIdxs = nil
}
