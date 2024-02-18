// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/commons/network/dns/domain/enums/enums.proto

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

// dns domain visibility used for configuring dns domains visibility
type DnsDomainVisibility int32

const (
	DnsDomainVisibility_DNS_DOMAIN_VISIBILITY_UNSPECIFIED DnsDomainVisibility = 0
	DnsDomainVisibility_DNS_DOMAIN_VISIBILITY_INTERNAL    DnsDomainVisibility = 1
	DnsDomainVisibility_DNS_DOMAIN_VISIBILITY_EXTERNAL    DnsDomainVisibility = 2
)

// Enum value maps for DnsDomainVisibility.
var (
	DnsDomainVisibility_name = map[int32]string{
		0: "DNS_DOMAIN_VISIBILITY_UNSPECIFIED",
		1: "DNS_DOMAIN_VISIBILITY_INTERNAL",
		2: "DNS_DOMAIN_VISIBILITY_EXTERNAL",
	}
	DnsDomainVisibility_value = map[string]int32{
		"DNS_DOMAIN_VISIBILITY_UNSPECIFIED": 0,
		"DNS_DOMAIN_VISIBILITY_INTERNAL":    1,
		"DNS_DOMAIN_VISIBILITY_EXTERNAL":    2,
	}
)

func (x DnsDomainVisibility) Enum() *DnsDomainVisibility {
	p := new(DnsDomainVisibility)
	*p = x
	return p
}

func (x DnsDomainVisibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DnsDomainVisibility) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_enumTypes[0].Descriptor()
}

func (DnsDomainVisibility) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_enumTypes[0]
}

func (x DnsDomainVisibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DnsDomainVisibility.Descriptor instead.
func (DnsDomainVisibility) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_rawDesc = []byte{
	0x0a, 0x42, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x6e, 0x73, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0x84, 0x01, 0x0a,
	0x13, 0x44, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x21, 0x44, 0x4e, 0x53, 0x5f, 0x44, 0x4f, 0x4d, 0x41,
	0x49, 0x4e, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x44,
	0x4e, 0x53, 0x5f, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49,
	0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12,
	0x22, 0x0a, 0x1e, 0x44, 0x4e, 0x53, 0x5f, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x56, 0x49,
	0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x10, 0x02, 0x42, 0xc0, 0x03, 0x0a, 0x44, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x6e, 0x73, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x0a, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x66, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x56, 0x43, 0x4e, 0x44, 0x44, 0x45, 0xaa, 0x02,
	0x36, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41,
	0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x44, 0x6e, 0x73, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x36, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5c,
	0x44, 0x6e, 0x73, 0x5c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xe2, 0x02, 0x42, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x5c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5c, 0x44, 0x6e, 0x73, 0x5c, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x3a, 0x3a, 0x44, 0x6e, 0x73, 0x3a, 0x3a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_rawDescData = file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_goTypes = []interface{}{
	(DnsDomainVisibility)(0), // 0: cloud.planton.apis.commons.network.dns.domain.enums.DnsDomainVisibility
}
var file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_init() }
func file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_init() {
	if File_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto = out.File
	file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_network_dns_domain_enums_enums_proto_depIdxs = nil
}
