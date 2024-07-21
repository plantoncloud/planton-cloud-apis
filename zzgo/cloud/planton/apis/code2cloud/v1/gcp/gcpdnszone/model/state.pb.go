// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/gcp/gcpdnszone/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcefieldoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemessageoptions"
	dnsrecordtype "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/networking/enums/dnsrecordtype"
	model2 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/resourcemanager/v1/environment/model"
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

// gcp-dns-zone
type GcpDnsZone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource-kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata
	// id format "<id-prefix>-<env-id>-<normalized-resource-name>"
	Metadata *model.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *GcpDnsZoneSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *GcpDnsZoneStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GcpDnsZone) Reset() {
	*x = GcpDnsZone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpDnsZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpDnsZone) ProtoMessage() {}

func (x *GcpDnsZone) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpDnsZone.ProtoReflect.Descriptor instead.
func (*GcpDnsZone) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *GcpDnsZone) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *GcpDnsZone) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *GcpDnsZone) GetMetadata() *model.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GcpDnsZone) GetSpec() *GcpDnsZoneSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *GcpDnsZone) GetStatus() *GcpDnsZoneStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// gcp-dns-zone spec
type GcpDnsZoneSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// environment-info
	EnvironmentInfo *model1.ApiResourceEnvironmentInfo `protobuf:"bytes,99,opt,name=environment_info,json=environmentInfo,proto3" json:"environment_info,omitempty"`
	// stack-job settings
	StackJobSettings *model2.StackJobSettings `protobuf:"bytes,98,opt,name=stack_job_settings,json=stackJobSettings,proto3" json:"stack_job_settings,omitempty"`
	// gcp-credential-id to be used for setting up gcp-provider in stack-job
	GcpCredentialId string `protobuf:"bytes,97,opt,name=gcp_credential_id,json=gcpCredentialId,proto3" json:"gcp_credential_id,omitempty"`
	// The ID of the GCP Project where the Managed Zone is created.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// An optional list of GCP service accounts to be granted permissions to manage DNS records in the Managed Zone.
	// These accounts are primarily created as workload identities like cert-manager,
	// and are added when new environments are created or updated.
	IamServiceAccounts []string `protobuf:"bytes,2,rep,name=iam_service_accounts,json=iamServiceAccounts,proto3" json:"iam_service_accounts,omitempty"`
	// The DNS records that are added to the Zone.
	Records []*GcpDnsRecord `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *GcpDnsZoneSpec) Reset() {
	*x = GcpDnsZoneSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpDnsZoneSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpDnsZoneSpec) ProtoMessage() {}

func (x *GcpDnsZoneSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpDnsZoneSpec.ProtoReflect.Descriptor instead.
func (*GcpDnsZoneSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *GcpDnsZoneSpec) GetEnvironmentInfo() *model1.ApiResourceEnvironmentInfo {
	if x != nil {
		return x.EnvironmentInfo
	}
	return nil
}

func (x *GcpDnsZoneSpec) GetStackJobSettings() *model2.StackJobSettings {
	if x != nil {
		return x.StackJobSettings
	}
	return nil
}

func (x *GcpDnsZoneSpec) GetGcpCredentialId() string {
	if x != nil {
		return x.GcpCredentialId
	}
	return ""
}

func (x *GcpDnsZoneSpec) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GcpDnsZoneSpec) GetIamServiceAccounts() []string {
	if x != nil {
		return x.IamServiceAccounts
	}
	return nil
}

func (x *GcpDnsZoneSpec) GetRecords() []*GcpDnsRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

// gcp-dns-zone status
type GcpDnsZoneStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// lifecycle
	Lifecycle *model.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// audit-info
	Audit *model.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// stack-job id
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
	// stack-outputs
	//
	//	stack outputs
	StackOutputs *GcpDnsZoneStackOutputs `protobuf:"bytes,1,opt,name=stack_outputs,json=stackOutputs,proto3" json:"stack_outputs,omitempty"`
}

func (x *GcpDnsZoneStatus) Reset() {
	*x = GcpDnsZoneStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpDnsZoneStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpDnsZoneStatus) ProtoMessage() {}

func (x *GcpDnsZoneStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpDnsZoneStatus.ProtoReflect.Descriptor instead.
func (*GcpDnsZoneStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *GcpDnsZoneStatus) GetLifecycle() *model.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *GcpDnsZoneStatus) GetAudit() *model.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *GcpDnsZoneStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

func (x *GcpDnsZoneStatus) GetStackOutputs() *GcpDnsZoneStackOutputs {
	if x != nil {
		return x.StackOutputs
	}
	return nil
}

// gcp-dns-zone stack-outputs
type GcpDnsZoneStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of nameservers for the Managed Zone created for the DNS Domain.
	Nameservers []string `protobuf:"bytes,1,rep,name=nameservers,proto3" json:"nameservers,omitempty"`
}

func (x *GcpDnsZoneStackOutputs) Reset() {
	*x = GcpDnsZoneStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpDnsZoneStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpDnsZoneStackOutputs) ProtoMessage() {}

func (x *GcpDnsZoneStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpDnsZoneStackOutputs.ProtoReflect.Descriptor instead.
func (*GcpDnsZoneStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDescGZIP(), []int{3}
}

func (x *GcpDnsZoneStackOutputs) GetNameservers() []string {
	if x != nil {
		return x.Nameservers
	}
	return nil
}

// gcp-dns-zone dns-record
type GcpDnsRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// dns record type.
	RecordType dnsrecordtype.DnsRecordType `protobuf:"varint,2,opt,name=record_type,json=recordType,proto3,enum=cloud.planton.apis.commons.networking.enums.dnsrecordtype.DnsRecordType" json:"record_type,omitempty"`
	// name of the gcp-dns-zone ex: example.com or dev.example.com.
	// this value should always end with a dot.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// values for the gcp-dns-zone record.
	// if the gcp_dns_zone_record_type is cname then each value in the list should end with a dot.
	Values []string `protobuf:"bytes,4,rep,name=values,proto3" json:"values,omitempty"`
	// ttl for the domain record in seconds.
	TtlSeconds int32 `protobuf:"varint,5,opt,name=ttl_seconds,json=ttlSeconds,proto3" json:"ttl_seconds,omitempty"`
}

func (x *GcpDnsRecord) Reset() {
	*x = GcpDnsRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpDnsRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpDnsRecord) ProtoMessage() {}

func (x *GcpDnsRecord) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpDnsRecord.ProtoReflect.Descriptor instead.
func (*GcpDnsRecord) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDescGZIP(), []int{4}
}

func (x *GcpDnsRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GcpDnsRecord) GetRecordType() dnsrecordtype.DnsRecordType {
	if x != nil {
		return x.RecordType
	}
	return dnsrecordtype.DnsRecordType(0)
}

func (x *GcpDnsRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GcpDnsRecord) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *GcpDnsRecord) GetTtlSeconds() int32 {
	if x != nil {
		return x.TtlSeconds
	}
	return 0
}

var File_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDesc = []byte{
	0x0a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e,
	0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x64, 0x6e, 0x73,
	0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6b, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x64, 0x6e, 0x73, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x94, 0x06, 0x0a, 0x0a, 0x47, 0x63, 0x70, 0x44, 0x6e, 0x73, 0x5a, 0x6f,
	0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0xea, 0x03, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x8a, 0x03,
	0xba, 0x48, 0x86, 0x03, 0xba, 0x01, 0xa6, 0x01, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x73, 0x68,
	0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x61, 0x6e, 0x79, 0x20, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x20, 0x44, 0x4e, 0x53, 0x20, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x20, 0x4e, 0x61, 0x6d,
	0x65, 0x1a, 0x6b, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x28, 0x3f, 0x3a, 0x5b, 0x5f, 0x61, 0x2d, 0x7a, 0x30,
	0x2d, 0x39, 0x5d, 0x28, 0x3f, 0x3a, 0x5b, 0x5f, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d,
	0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f,
	0x5b, 0x2e, 0x5d, 0x29, 0x2b, 0x28, 0x3f, 0x3a, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x28, 0x3f, 0x3a,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x29, 0x3f, 0x24, 0x27, 0x29, 0xba, 0x01,
	0x6c, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2d, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x62,
	0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x20, 0x31, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x36, 0x35, 0x20,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x20, 0x6c, 0x6f, 0x6e, 0x67, 0x1a,
	0x2c, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29,
	0x20, 0x3e, 0x20, 0x30, 0x20, 0x26, 0x26, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20, 0x3c, 0x3d, 0x20, 0x36, 0x35, 0xba, 0x01, 0x67,
	0x0a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x69, 0x73, 0x20, 0x6d,
	0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x68, 0x61,
	0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x29, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x59, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x64, 0x6e, 0x73,
	0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x63, 0x70, 0x44, 0x6e,
	0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12,
	0x5f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x47, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e,
	0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x63, 0x70, 0x44, 0x6e, 0x73, 0x5a, 0x6f,
	0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x3a, 0x28, 0x88, 0xa6, 0x1d, 0x0d, 0x92, 0xa6, 0x1d, 0x20, 0x08, 0x0a, 0x12, 0x1c, 0x73, 0x70,
	0x65, 0x63, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x69, 0x64, 0x22, 0xdc, 0x03, 0x0a, 0x0e, 0x47,
	0x63, 0x70, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x7e, 0x0a,
	0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x68, 0x0a,
	0x12, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x62, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69,
	0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6a, 0x6f, 0x62, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x10, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x63, 0x70, 0x5f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x61, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x67, 0x63, 0x70, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x61, 0x6d, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x69, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x5d, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67,
	0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x47, 0x63, 0x70, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xe0, 0x02, 0x0a, 0x10, 0x47, 0x63,
	0x70, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x60,
	0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x63, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x12, 0x54, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x62, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52,
	0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f,
	0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x61, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x72, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x4d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e,
	0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x63, 0x70, 0x44, 0x6e, 0x73, 0x5a, 0x6f,
	0x6e, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x0c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22, 0x3a, 0x0a, 0x16,
	0x47, 0x63, 0x70, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x9c, 0x03, 0x0a, 0x0c, 0x47, 0x63, 0x70,
	0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x71, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x48, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x64, 0x6e, 0x73, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0xc3, 0x01, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0xae, 0x01, 0xba, 0x48, 0xaa, 0x01, 0xba, 0x01, 0xa3, 0x01, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x28, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20,
	0x62, 0x65, 0x20, 0x61, 0x6e, 0x79, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x44, 0x4e, 0x53,
	0x20, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x71, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x28, 0x3f, 0x3a,
	0x5b, 0x2a, 0x5d, 0x5b, 0x2e, 0x5d, 0x29, 0x3f, 0x28, 0x3f, 0x3a, 0x5b, 0x5f, 0x61, 0x2d, 0x7a,
	0x30, 0x2d, 0x39, 0x5d, 0x28, 0x3f, 0x3a, 0x5b, 0x5f, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d,
	0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29,
	0x3f, 0x5b, 0x2e, 0x5d, 0x29, 0x2b, 0x28, 0x3f, 0x3a, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x28, 0x3f,
	0x3a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x29, 0x3f, 0x24, 0x27, 0x29, 0xc8,
	0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x25, 0x0a, 0x0b, 0x74, 0x74, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x04, 0x80, 0xb9, 0x18, 0x3c, 0x52, 0x0a, 0x74, 0x74, 0x6c,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0xb8, 0x03, 0x0a, 0x43, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67,
	0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x65, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x43, 0x56, 0x47, 0x47, 0x4d, 0xaa,
	0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x56, 0x31, 0x2e, 0x47, 0x63, 0x70, 0x2e, 0x47, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e,
	0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47,
	0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2,
	0x02, 0x41, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e,
	0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x3c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x3a,
	0x3a, 0x47, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x3a, 0x3a, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_goTypes = []interface{}{
	(*GcpDnsZone)(nil),                        // 0: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZone
	(*GcpDnsZoneSpec)(nil),                    // 1: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZoneSpec
	(*GcpDnsZoneStatus)(nil),                  // 2: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZoneStatus
	(*GcpDnsZoneStackOutputs)(nil),            // 3: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZoneStackOutputs
	(*GcpDnsRecord)(nil),                      // 4: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsRecord
	(*model.ApiResourceMetadata)(nil),         // 5: cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	(*model1.ApiResourceEnvironmentInfo)(nil), // 6: cloud.planton.apis.resourcemanager.v1.environment.model.ApiResourceEnvironmentInfo
	(*model2.StackJobSettings)(nil),           // 7: cloud.planton.apis.iac.v1.stackjob.model.StackJobSettings
	(*model.ApiResourceLifecycle)(nil),        // 8: cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	(*model.ApiResourceAudit)(nil),            // 9: cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	(dnsrecordtype.DnsRecordType)(0),          // 10: cloud.planton.apis.commons.networking.enums.dnsrecordtype.DnsRecordType
}
var file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_depIdxs = []int32{
	5,  // 0: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZone.metadata:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	1,  // 1: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZone.spec:type_name -> cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZoneSpec
	2,  // 2: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZone.status:type_name -> cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZoneStatus
	6,  // 3: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZoneSpec.environment_info:type_name -> cloud.planton.apis.resourcemanager.v1.environment.model.ApiResourceEnvironmentInfo
	7,  // 4: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZoneSpec.stack_job_settings:type_name -> cloud.planton.apis.iac.v1.stackjob.model.StackJobSettings
	4,  // 5: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZoneSpec.records:type_name -> cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsRecord
	8,  // 6: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZoneStatus.lifecycle:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	9,  // 7: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZoneStatus.audit:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	3,  // 8: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZoneStatus.stack_outputs:type_name -> cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsZoneStackOutputs
	10, // 9: cloud.planton.apis.code2cloud.v1.gcp.gcpdnszone.model.GcpDnsRecord.record_type:type_name -> cloud.planton.apis.commons.networking.enums.dnsrecordtype.DnsRecordType
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpDnsZone); i {
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
		file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpDnsZoneSpec); i {
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
		file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpDnsZoneStatus); i {
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
		file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpDnsZoneStackOutputs); i {
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
		file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpDnsRecord); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_gcp_gcpdnszone_model_state_proto_depIdxs = nil
}
