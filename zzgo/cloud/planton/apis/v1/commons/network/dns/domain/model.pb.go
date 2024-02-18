// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/commons/network/dns/domain/model.proto

package domain

import (
	enums "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/network/dns/domain/enums"
	record "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/network/dns/record"
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

type DnsDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// domain name example.com
	DomainName string `protobuf:"bytes,1,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	// domain visibility. for private visibility internal reserved ips are added to the zone.
	// for publicly visible domains external reserved domains are added to the zone.
	Visibility enums.DnsDomainVisibility `protobuf:"varint,2,opt,name=visibility,proto3,enum=cloud.planton.apis.commons.network.dns.domain.enums.DnsDomainVisibility" json:"visibility,omitempty"`
	// dns records to be added to the zone
	DnsRecords []*record.DnsRecord `protobuf:"bytes,3,rep,name=dns_records,json=dnsRecords,proto3" json:"dns_records,omitempty"`
}

func (x *DnsDomain) Reset() {
	*x = DnsDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsDomain) ProtoMessage() {}

func (x *DnsDomain) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsDomain.ProtoReflect.Descriptor instead.
func (*DnsDomain) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_rawDescGZIP(), []int{0}
}

func (x *DnsDomain) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *DnsDomain) GetVisibility() enums.DnsDomainVisibility {
	if x != nil {
		return x.Visibility
	}
	return enums.DnsDomainVisibility(0)
}

func (x *DnsDomain) GetDnsRecords() []*record.DnsRecord {
	if x != nil {
		return x.DnsRecords
	}
	return nil
}

var File_cloud_planton_apis_v1_commons_network_dns_domain_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x1a, 0x3c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x42,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x09, 0x44, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x6b, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x6e,
	0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x44,
	0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x5c,
	0x0a, 0x0b, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x6e, 0x73, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x0a, 0x64, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x9a, 0x03, 0x0a,
	0x3e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x42,
	0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x60, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0xa2,
	0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x43, 0x4e, 0x44, 0x44, 0xaa, 0x02, 0x30, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e,
	0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x44, 0x6e, 0x73, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0xca, 0x02, 0x30,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5c, 0x44, 0x6e, 0x73, 0x5c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0xe2, 0x02, 0x3c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x5c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5c, 0x44, 0x6e, 0x73, 0x5c, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x3a, 0x3a, 0x44, 0x6e,
	0x73, 0x3a, 0x3a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_rawDescData = file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_goTypes = []interface{}{
	(*DnsDomain)(nil),              // 0: cloud.planton.apis.commons.network.dns.domain.DnsDomain
	(enums.DnsDomainVisibility)(0), // 1: cloud.planton.apis.commons.network.dns.domain.enums.DnsDomainVisibility
	(*record.DnsRecord)(nil),       // 2: cloud.planton.apis.commons.network.dns.record.DnsRecord
}
var file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_depIdxs = []int32{
	1, // 0: cloud.planton.apis.commons.network.dns.domain.DnsDomain.visibility:type_name -> cloud.planton.apis.commons.network.dns.domain.enums.DnsDomainVisibility
	2, // 1: cloud.planton.apis.commons.network.dns.domain.DnsDomain.dns_records:type_name -> cloud.planton.apis.commons.network.dns.record.DnsRecord
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_init() }
func file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_init() {
	if File_cloud_planton_apis_v1_commons_network_dns_domain_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsDomain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_network_dns_domain_model_proto = out.File
	file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_network_dns_domain_model_proto_depIdxs = nil
}
