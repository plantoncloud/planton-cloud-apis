// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/dnszone/stack/gcp/model.proto

package gcp

import (
	dnszone "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/dnszone"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/network/dns/domain"
	operation "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/pulumi/operation"
	job "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job"
	progress "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job/progress"
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

// input for dns-zone stack
type DnsZoneGcpStackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack job
	StackJob *job.StackJob `protobuf:"bytes,1,opt,name=stack_job,json=stackJob,proto3" json:"stack_job,omitempty"`
	// pulumi stack credentials
	CredentialsInput *DnsZoneGcpStackCredentialsInput `protobuf:"bytes,2,opt,name=credentials_input,json=credentialsInput,proto3" json:"credentials_input,omitempty"`
	// inputs used for creating stack resources
	ResourceInput *DnsZoneGcpStackResourceInput `protobuf:"bytes,3,opt,name=resource_input,json=resourceInput,proto3" json:"resource_input,omitempty"`
}

func (x *DnsZoneGcpStackInput) Reset() {
	*x = DnsZoneGcpStackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsZoneGcpStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsZoneGcpStackInput) ProtoMessage() {}

func (x *DnsZoneGcpStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsZoneGcpStackInput.ProtoReflect.Descriptor instead.
func (*DnsZoneGcpStackInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDescGZIP(), []int{0}
}

func (x *DnsZoneGcpStackInput) GetStackJob() *job.StackJob {
	if x != nil {
		return x.StackJob
	}
	return nil
}

func (x *DnsZoneGcpStackInput) GetCredentialsInput() *DnsZoneGcpStackCredentialsInput {
	if x != nil {
		return x.CredentialsInput
	}
	return nil
}

func (x *DnsZoneGcpStackInput) GetResourceInput() *DnsZoneGcpStackResourceInput {
	if x != nil {
		return x.ResourceInput
	}
	return nil
}

// stack credentials input
type DnsZoneGcpStackCredentialsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// google provider credential input
	Google *operation.GoogleProviderCredential `protobuf:"bytes,1,opt,name=google,proto3" json:"google,omitempty"`
}

func (x *DnsZoneGcpStackCredentialsInput) Reset() {
	*x = DnsZoneGcpStackCredentialsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsZoneGcpStackCredentialsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsZoneGcpStackCredentialsInput) ProtoMessage() {}

func (x *DnsZoneGcpStackCredentialsInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsZoneGcpStackCredentialsInput.ProtoReflect.Descriptor instead.
func (*DnsZoneGcpStackCredentialsInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDescGZIP(), []int{1}
}

func (x *DnsZoneGcpStackCredentialsInput) GetGoogle() *operation.GoogleProviderCredential {
	if x != nil {
		return x.Google
	}
	return nil
}

// stack resource input
type DnsZoneGcpStackResourceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// dns-zone to be added to the company
	DnsZone *dnszone.DnsZone `protobuf:"bytes,1,opt,name=dns_zone,json=dnsZone,proto3" json:"dns_zone,omitempty"`
}

func (x *DnsZoneGcpStackResourceInput) Reset() {
	*x = DnsZoneGcpStackResourceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsZoneGcpStackResourceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsZoneGcpStackResourceInput) ProtoMessage() {}

func (x *DnsZoneGcpStackResourceInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsZoneGcpStackResourceInput.ProtoReflect.Descriptor instead.
func (*DnsZoneGcpStackResourceInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDescGZIP(), []int{2}
}

func (x *DnsZoneGcpStackResourceInput) GetDnsZone() *dnszone.DnsZone {
	if x != nil {
		return x.DnsZone
	}
	return nil
}

// the stack output only supports domains and their nameservers
// the outputs exclude any dns records for the domains
type DnsZoneGcpStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// zone status populated with the details of the zone created in google cloud dns
	ZoneStatus *dnszone.DnsZoneStatus `protobuf:"bytes,2,opt,name=zone_status,json=zoneStatus,proto3" json:"zone_status,omitempty"`
}

func (x *DnsZoneGcpStackOutputs) Reset() {
	*x = DnsZoneGcpStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsZoneGcpStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsZoneGcpStackOutputs) ProtoMessage() {}

func (x *DnsZoneGcpStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsZoneGcpStackOutputs.ProtoReflect.Descriptor instead.
func (*DnsZoneGcpStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDescGZIP(), []int{3}
}

func (x *DnsZoneGcpStackOutputs) GetZoneStatus() *dnszone.DnsZoneStatus {
	if x != nil {
		return x.ZoneStatus
	}
	return nil
}

// stack response
type DnsZoneGcpStackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack-job progress event
	ProgressEvent *progress.StackJobProgressEvent `protobuf:"bytes,1,opt,name=progress_event,json=progressEvent,proto3" json:"progress_event,omitempty"`
	// stack outputs
	Outputs *DnsZoneGcpStackOutputs `protobuf:"bytes,3,opt,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *DnsZoneGcpStackResponse) Reset() {
	*x = DnsZoneGcpStackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsZoneGcpStackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsZoneGcpStackResponse) ProtoMessage() {}

func (x *DnsZoneGcpStackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsZoneGcpStackResponse.ProtoReflect.Descriptor instead.
func (*DnsZoneGcpStackResponse) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDescGZIP(), []int{4}
}

func (x *DnsZoneGcpStackResponse) GetProgressEvent() *progress.StackJobProgressEvent {
	if x != nil {
		return x.ProgressEvent
	}
	return nil
}

func (x *DnsZoneGcpStackResponse) GetOutputs() *DnsZoneGcpStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x67,
	0x63, 0x70, 0x1a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x64, 0x6e, 0x73, 0x7a,
	0x6f, 0x6e, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x75, 0x6c,
	0x75, 0x6d, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x02, 0x0a,
	0x14, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x46, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a,
	0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x4a, 0x6f, 0x62, 0x52, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x87, 0x01,
	0x0a, 0x11, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x63, 0x70,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x7e, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x57, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e,
	0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x44, 0x6e, 0x73, 0x5a,
	0x6f, 0x6e, 0x65, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x1f, 0x44, 0x6e, 0x73, 0x5a,
	0x6f, 0x6e, 0x65, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x60, 0x0a, 0x06, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x75, 0x6c, 0x75,
	0x6d, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x22, 0x73, 0x0a,
	0x1c, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x53, 0x0a,
	0x08, 0x64, 0x6e, 0x73, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e,
	0x65, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x07, 0x64, 0x6e, 0x73, 0x5a, 0x6f,
	0x6e, 0x65, 0x22, 0x79, 0x0a, 0x16, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x63, 0x70,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x5f, 0x0a, 0x0b,
	0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x64, 0x6e, 0x73, 0x7a,
	0x6f, 0x6e, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x0a, 0x7a, 0x6f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xee, 0x01,
	0x0a, 0x17, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x0e, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x6b, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x51, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x64, 0x6e, 0x73,
	0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x44,
	0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x63, 0x70, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x42, 0xd2,
	0x03, 0x0a, 0x47, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x67, 0x63, 0x70, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2f, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f,
	0x67, 0x63, 0x70, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x56, 0x43, 0x44, 0x44, 0x53, 0x47, 0xaa,
	0x02, 0x39, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x44, 0x6e, 0x73, 0x7a, 0x6f, 0x6e,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x47, 0x63, 0x70, 0xca, 0x02, 0x39, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x44, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x5c, 0x47, 0x63, 0x70, 0xe2, 0x02, 0x45, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x5c, 0x44, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c,
	0x47, 0x63, 0x70, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x41, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a,
	0x44, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a,
	0x47, 0x63, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_goTypes = []interface{}{
	(*DnsZoneGcpStackInput)(nil),               // 0: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackInput
	(*DnsZoneGcpStackCredentialsInput)(nil),    // 1: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackCredentialsInput
	(*DnsZoneGcpStackResourceInput)(nil),       // 2: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackResourceInput
	(*DnsZoneGcpStackOutputs)(nil),             // 3: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackOutputs
	(*DnsZoneGcpStackResponse)(nil),            // 4: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackResponse
	(*job.StackJob)(nil),                       // 5: cloud.planton.apis.v1.stack.job.StackJob
	(*operation.GoogleProviderCredential)(nil), // 6: cloud.planton.apis.v1.commons.pulumi.operation.GoogleProviderCredential
	(*dnszone.DnsZone)(nil),                    // 7: cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZone
	(*dnszone.DnsZoneStatus)(nil),              // 8: cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneStatus
	(*progress.StackJobProgressEvent)(nil),     // 9: cloud.planton.apis.v1.stack.job.progress.StackJobProgressEvent
}
var file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackInput.stack_job:type_name -> cloud.planton.apis.v1.stack.job.StackJob
	1, // 1: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackInput.credentials_input:type_name -> cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackCredentialsInput
	2, // 2: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackInput.resource_input:type_name -> cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackResourceInput
	6, // 3: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackCredentialsInput.google:type_name -> cloud.planton.apis.v1.commons.pulumi.operation.GoogleProviderCredential
	7, // 4: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackResourceInput.dns_zone:type_name -> cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZone
	8, // 5: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackOutputs.zone_status:type_name -> cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZoneStatus
	9, // 6: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackResponse.progress_event:type_name -> cloud.planton.apis.v1.stack.job.progress.StackJobProgressEvent
	3, // 7: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackResponse.outputs:type_name -> cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.gcp.DnsZoneGcpStackOutputs
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsZoneGcpStackInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsZoneGcpStackCredentialsInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsZoneGcpStackResourceInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsZoneGcpStackOutputs); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsZoneGcpStackResponse); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_gcp_model_proto_depIdxs = nil
}
