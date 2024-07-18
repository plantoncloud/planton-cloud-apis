// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/gcp/gcpclouddnszone/model/io.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	dnsrecordtype "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/networking/enums/dnsrecordtype"
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

// wrapper for gcp-cloud-dns-zone id
type GcpCloudDnsZoneId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GcpCloudDnsZoneId) Reset() {
	*x = GcpCloudDnsZoneId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpCloudDnsZoneId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpCloudDnsZoneId) ProtoMessage() {}

func (x *GcpCloudDnsZoneId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpCloudDnsZoneId.ProtoReflect.Descriptor instead.
func (*GcpCloudDnsZoneId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *GcpCloudDnsZoneId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// list of gcp-cloud-dns-zones
type GcpCloudDnsZoneList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*GcpCloudDnsZone `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *GcpCloudDnsZoneList) Reset() {
	*x = GcpCloudDnsZoneList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpCloudDnsZoneList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpCloudDnsZoneList) ProtoMessage() {}

func (x *GcpCloudDnsZoneList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpCloudDnsZoneList.ProtoReflect.Descriptor instead.
func (*GcpCloudDnsZoneList) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *GcpCloudDnsZoneList) GetEntries() []*GcpCloudDnsZone {
	if x != nil {
		return x.Entries
	}
	return nil
}

// AddOrUpdateDnsRecordInput is used to encapsulate the details needed to either add
// a new DNS record or update an existing one within a specific DNS zone.
// This message is typically used to transmit data from the client to the server
// during an add or update or restore operation on a DNS record within a specified DNS zone.
type AddOrUpdateGcpCloudDnsRecordInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field contains the unique identifier of the DNS zone within which the DNS record
	// is to be added or updated. The ID should be a valid DNS zone ID, typically obtained
	// from the DNS zone entity itself. This field is used by the server to determine
	// the correct DNS zone where the DNS record is to be added or updated.
	GcpCloudDnsZoneId string `protobuf:"bytes,1,opt,name=gcp_cloud_dns_zone_id,json=gcpCloudDnsZoneId,proto3" json:"gcp_cloud_dns_zone_id,omitempty"`
	// This field contains the actual DNS record that is to be added or updated. The DNS record
	// object should be populated with all necessary fields (like name, type, TTL, data etc.),
	// which the server uses to create a new DNS record or update an existing one.
	Record *GcpCloudDnsDnsRecord `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
	// A descriptive message explaining the reason for the change or operation.
	// This is used for history logging purposes.
	VersionMessage string `protobuf:"bytes,3,opt,name=version_message,json=versionMessage,proto3" json:"version_message,omitempty"`
}

func (x *AddOrUpdateGcpCloudDnsRecordInput) Reset() {
	*x = AddOrUpdateGcpCloudDnsRecordInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrUpdateGcpCloudDnsRecordInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrUpdateGcpCloudDnsRecordInput) ProtoMessage() {}

func (x *AddOrUpdateGcpCloudDnsRecordInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrUpdateGcpCloudDnsRecordInput.ProtoReflect.Descriptor instead.
func (*AddOrUpdateGcpCloudDnsRecordInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDescGZIP(), []int{2}
}

func (x *AddOrUpdateGcpCloudDnsRecordInput) GetGcpCloudDnsZoneId() string {
	if x != nil {
		return x.GcpCloudDnsZoneId
	}
	return ""
}

func (x *AddOrUpdateGcpCloudDnsRecordInput) GetRecord() *GcpCloudDnsDnsRecord {
	if x != nil {
		return x.Record
	}
	return nil
}

func (x *AddOrUpdateGcpCloudDnsRecordInput) GetVersionMessage() string {
	if x != nil {
		return x.VersionMessage
	}
	return ""
}

// DeleteOrRestoreDnsRecordInput is a protobuf message that encapsulates the data
// necessary for either soft deleting or restoring a specific DNS record within a DNS zone.
// This message is typically used to send the relevant details from the client to the server.
type DeleteGcpCloudDnsRecordInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field carries the unique identifier of the DNS zone from which the DNS record
	// is to be either soft deleted or restored. The server uses this ID to find the appropriate DNS zone.
	// This is a required field.
	GcpCloudDnsZoneId string `protobuf:"bytes,1,opt,name=gcp_cloud_dns_zone_id,json=gcpCloudDnsZoneId,proto3" json:"gcp_cloud_dns_zone_id,omitempty"`
	// This field holds the type of the DNS record that is to be either soft deleted or restored.
	// DNS record types include but are not limited to A, AAAA, CNAME, MX, etc.
	// The server uses this information to find the DNS record of the correct type within the
	// specified DNS zone. This is a required field.
	RecordType dnsrecordtype.DnsRecordType `protobuf:"varint,2,opt,name=record_type,json=recordType,proto3,enum=cloud.planton.apis.commons.networking.enums.dnsrecordtype.DnsRecordType" json:"record_type,omitempty"`
	// This field contains the name of the DNS record that is to be either soft deleted or restored.
	// The DNS record name is a significant identifier used by the server to find the specific record
	// within the specified DNS zone and of the provided type to either soft delete or restore.
	// This is a required field.
	RecordName string `protobuf:"bytes,3,opt,name=record_name,json=recordName,proto3" json:"record_name,omitempty"`
	// A descriptive message explaining the reason for the change or operation.
	// This is used for history logging purposes.
	VersionMessage string `protobuf:"bytes,4,opt,name=version_message,json=versionMessage,proto3" json:"version_message,omitempty"`
}

func (x *DeleteGcpCloudDnsRecordInput) Reset() {
	*x = DeleteGcpCloudDnsRecordInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGcpCloudDnsRecordInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGcpCloudDnsRecordInput) ProtoMessage() {}

func (x *DeleteGcpCloudDnsRecordInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGcpCloudDnsRecordInput.ProtoReflect.Descriptor instead.
func (*DeleteGcpCloudDnsRecordInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteGcpCloudDnsRecordInput) GetGcpCloudDnsZoneId() string {
	if x != nil {
		return x.GcpCloudDnsZoneId
	}
	return ""
}

func (x *DeleteGcpCloudDnsRecordInput) GetRecordType() dnsrecordtype.DnsRecordType {
	if x != nil {
		return x.RecordType
	}
	return dnsrecordtype.DnsRecordType(0)
}

func (x *DeleteGcpCloudDnsRecordInput) GetRecordName() string {
	if x != nil {
		return x.RecordName
	}
	return ""
}

func (x *DeleteGcpCloudDnsRecordInput) GetVersionMessage() string {
	if x != nil {
		return x.VersionMessage
	}
	return ""
}

var File_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDesc = []byte{
	0x0a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x64,
	0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x64, 0x6e, 0x73,
	0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x64, 0x6e, 0x73, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x11, 0x47, 0x63, 0x70, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7c, 0x0a, 0x13, 0x47, 0x63,
	0x70, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x65, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x47, 0x63, 0x70, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x52,
	0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x80, 0x02, 0x0a, 0x21, 0x41, 0x64, 0x64,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x63, 0x70, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x38,
	0x0a, 0x15, 0x67, 0x63, 0x70, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x64, 0x6e, 0x73, 0x5f,
	0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x11, 0x67, 0x63, 0x70, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44,
	0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x70, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e,
	0x67, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x63, 0x70, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x6e,
	0x73, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2f, 0x0a, 0x0f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xaa, 0x02, 0x0a, 0x1c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x63, 0x70, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x6e,
	0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x15,
	0x67, 0x63, 0x70, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x7a, 0x6f,
	0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x11, 0x67, 0x63, 0x70, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x6e, 0x73,
	0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x76, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x48, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x64, 0x6e, 0x73, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x82, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x0f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0xd3, 0x03, 0x0a, 0x48, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e,
	0x67, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x6a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x64,
	0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x08, 0x43,
	0x50, 0x41, 0x43, 0x56, 0x47, 0x47, 0x4d, 0xaa, 0x02, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x47, 0x63, 0x70, 0x2e, 0x47,
	0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x70, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0xe2, 0x02, 0x46, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x41, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69,
	0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56,
	0x31, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_goTypes = []interface{}{
	(*GcpCloudDnsZoneId)(nil),                 // 0: cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.model.GcpCloudDnsZoneId
	(*GcpCloudDnsZoneList)(nil),               // 1: cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.model.GcpCloudDnsZoneList
	(*AddOrUpdateGcpCloudDnsRecordInput)(nil), // 2: cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.model.AddOrUpdateGcpCloudDnsRecordInput
	(*DeleteGcpCloudDnsRecordInput)(nil),      // 3: cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.model.DeleteGcpCloudDnsRecordInput
	(*GcpCloudDnsZone)(nil),                   // 4: cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.model.GcpCloudDnsZone
	(*GcpCloudDnsDnsRecord)(nil),              // 5: cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.model.GcpCloudDnsDnsRecord
	(dnsrecordtype.DnsRecordType)(0),          // 6: cloud.planton.apis.commons.networking.enums.dnsrecordtype.DnsRecordType
}
var file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.model.GcpCloudDnsZoneList.entries:type_name -> cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.model.GcpCloudDnsZone
	5, // 1: cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.model.AddOrUpdateGcpCloudDnsRecordInput.record:type_name -> cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.model.GcpCloudDnsDnsRecord
	6, // 2: cloud.planton.apis.code2cloud.v1.gcp.gcpclouddnszone.model.DeleteGcpCloudDnsRecordInput.record_type:type_name -> cloud.planton.apis.commons.networking.enums.dnsrecordtype.DnsRecordType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto != nil {
		return
	}
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpCloudDnsZoneId); i {
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
		file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpCloudDnsZoneList); i {
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
		file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrUpdateGcpCloudDnsRecordInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGcpCloudDnsRecordInput); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpclouddnszone_model_io_proto_depIdxs = nil
}
