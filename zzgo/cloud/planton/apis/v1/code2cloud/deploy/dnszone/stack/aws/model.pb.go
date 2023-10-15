// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/dnszone/stack/aws/model.proto

package aws

import (
	state "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/dnszone/state"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/network/dns/domain/rpc"
	rpc1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/pulumi/operation/rpc"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/rpc"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/rpc/enums"
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
type DnsZoneAwsStackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack job
	StackJob *rpc.StackJob `protobuf:"bytes,1,opt,name=stack_job,json=stackJob,proto3" json:"stack_job,omitempty"`
	// pulumi stack credentials
	CredentialsInput *DnsZoneAwsStackCredentialsInput `protobuf:"bytes,2,opt,name=credentials_input,json=credentialsInput,proto3" json:"credentials_input,omitempty"`
	// inputs used for creating stack resources
	ResourceInput *DnsZoneAwsStackResourceInput `protobuf:"bytes,3,opt,name=resource_input,json=resourceInput,proto3" json:"resource_input,omitempty"`
}

func (x *DnsZoneAwsStackInput) Reset() {
	*x = DnsZoneAwsStackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsZoneAwsStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsZoneAwsStackInput) ProtoMessage() {}

func (x *DnsZoneAwsStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsZoneAwsStackInput.ProtoReflect.Descriptor instead.
func (*DnsZoneAwsStackInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDescGZIP(), []int{0}
}

func (x *DnsZoneAwsStackInput) GetStackJob() *rpc.StackJob {
	if x != nil {
		return x.StackJob
	}
	return nil
}

func (x *DnsZoneAwsStackInput) GetCredentialsInput() *DnsZoneAwsStackCredentialsInput {
	if x != nil {
		return x.CredentialsInput
	}
	return nil
}

func (x *DnsZoneAwsStackInput) GetResourceInput() *DnsZoneAwsStackResourceInput {
	if x != nil {
		return x.ResourceInput
	}
	return nil
}

// stack credentials input
type DnsZoneAwsStackCredentialsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// google provider credential input
	Aws *rpc1.AwsProviderCredential `protobuf:"bytes,1,opt,name=aws,proto3" json:"aws,omitempty"`
}

func (x *DnsZoneAwsStackCredentialsInput) Reset() {
	*x = DnsZoneAwsStackCredentialsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsZoneAwsStackCredentialsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsZoneAwsStackCredentialsInput) ProtoMessage() {}

func (x *DnsZoneAwsStackCredentialsInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsZoneAwsStackCredentialsInput.ProtoReflect.Descriptor instead.
func (*DnsZoneAwsStackCredentialsInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDescGZIP(), []int{1}
}

func (x *DnsZoneAwsStackCredentialsInput) GetAws() *rpc1.AwsProviderCredential {
	if x != nil {
		return x.Aws
	}
	return nil
}

// stack resource input
type DnsZoneAwsStackResourceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// dns-zone to be added to the company
	DnsZone *state.DnsZoneState `protobuf:"bytes,1,opt,name=dns_zone,json=dnsZone,proto3" json:"dns_zone,omitempty"`
}

func (x *DnsZoneAwsStackResourceInput) Reset() {
	*x = DnsZoneAwsStackResourceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsZoneAwsStackResourceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsZoneAwsStackResourceInput) ProtoMessage() {}

func (x *DnsZoneAwsStackResourceInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsZoneAwsStackResourceInput.ProtoReflect.Descriptor instead.
func (*DnsZoneAwsStackResourceInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDescGZIP(), []int{2}
}

func (x *DnsZoneAwsStackResourceInput) GetDnsZone() *state.DnsZoneState {
	if x != nil {
		return x.DnsZone
	}
	return nil
}

// the stack output only supports domains and their nameservers
// the outputs exclude any dns records for the domains
type DnsZoneAwsStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// zone status populated with the details of the zone created in google cloud dns
	ZoneStatus *state.DnsZoneStatusState `protobuf:"bytes,2,opt,name=zone_status,json=zoneStatus,proto3" json:"zone_status,omitempty"`
}

func (x *DnsZoneAwsStackOutputs) Reset() {
	*x = DnsZoneAwsStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsZoneAwsStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsZoneAwsStackOutputs) ProtoMessage() {}

func (x *DnsZoneAwsStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsZoneAwsStackOutputs.ProtoReflect.Descriptor instead.
func (*DnsZoneAwsStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDescGZIP(), []int{3}
}

func (x *DnsZoneAwsStackOutputs) GetZoneStatus() *state.DnsZoneStatusState {
	if x != nil {
		return x.ZoneStatus
	}
	return nil
}

// stack response
type DnsZoneAwsStackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stack progress
	Progress *rpc.StackProgress `protobuf:"bytes,1,opt,name=progress,proto3" json:"progress,omitempty"`
	// stack outputs
	Outputs *DnsZoneAwsStackOutputs `protobuf:"bytes,3,opt,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *DnsZoneAwsStackResponse) Reset() {
	*x = DnsZoneAwsStackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsZoneAwsStackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsZoneAwsStackResponse) ProtoMessage() {}

func (x *DnsZoneAwsStackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsZoneAwsStackResponse.ProtoReflect.Descriptor instead.
func (*DnsZoneAwsStackResponse) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDescGZIP(), []int{4}
}

func (x *DnsZoneAwsStackResponse) GetProgress() *rpc.StackProgress {
	if x != nil {
		return x.Progress
	}
	return nil
}

func (x *DnsZoneAwsStackResponse) GetOutputs() *DnsZoneAwsStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61,
	0x77, 0x73, 0x1a, 0x40, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x64, 0x6e,
	0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x02, 0x0a, 0x14, 0x44, 0x6e, 0x73, 0x5a,
	0x6f, 0x6e, 0x65, 0x41, 0x77, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x46, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x08,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x87, 0x01, 0x0a, 0x11, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x64,
	0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x77, 0x73,
	0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x41, 0x77, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x52, 0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x7e, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x57, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x41, 0x77,
	0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x22, 0x7e, 0x0a, 0x1f, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x41, 0x77, 0x73,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x5b, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x49, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x77, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x03, 0x61,
	0x77, 0x73, 0x22, 0x7e, 0x0a, 0x1c, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x41, 0x77, 0x73,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x5e, 0x0a, 0x08, 0x64, 0x6e, 0x73, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x64,
	0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x6e, 0x73,
	0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x07, 0x64, 0x6e, 0x73, 0x5a, 0x6f,
	0x6e, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x16, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x41, 0x77,
	0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x6a, 0x0a,
	0x0b, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x49, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x64, 0x6e, 0x73,
	0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x5a, 0x6f,
	0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x7a,
	0x6f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xd2, 0x01, 0x0a, 0x17, 0x44, 0x6e,
	0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x41, 0x77, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x6b, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x51, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x64, 0x6e, 0x73,
	0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x44,
	0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x41, 0x77, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x42, 0xc8,
	0x03, 0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x64,
	0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x77, 0x73,
	0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x69,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x77, 0x73, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x56,
	0x43, 0x44, 0x44, 0x53, 0x41, 0xaa, 0x02, 0x39, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x44, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x41, 0x77,
	0x73, 0xca, 0x02, 0x39, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x44, 0x6e, 0x73, 0x7a,
	0x6f, 0x6e, 0x65, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x77, 0x73, 0xe2, 0x02, 0x45,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x44, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x5c,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x41, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a, 0x44, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x3a, 0x3a, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x41, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_goTypes = []interface{}{
	(*DnsZoneAwsStackInput)(nil),            // 0: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackInput
	(*DnsZoneAwsStackCredentialsInput)(nil), // 1: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackCredentialsInput
	(*DnsZoneAwsStackResourceInput)(nil),    // 2: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackResourceInput
	(*DnsZoneAwsStackOutputs)(nil),          // 3: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackOutputs
	(*DnsZoneAwsStackResponse)(nil),         // 4: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackResponse
	(*rpc.StackJob)(nil),                    // 5: cloud.planton.apis.v1.stack.rpc.StackJob
	(*rpc1.AwsProviderCredential)(nil),      // 6: cloud.planton.apis.v1.commons.pulumi.operation.rpc.AwsProviderCredential
	(*state.DnsZoneState)(nil),              // 7: cloud.planton.apis.v1.code2cloud.deploy.dnszone.state.DnsZoneState
	(*state.DnsZoneStatusState)(nil),        // 8: cloud.planton.apis.v1.code2cloud.deploy.dnszone.state.DnsZoneStatusState
	(*rpc.StackProgress)(nil),               // 9: cloud.planton.apis.v1.stack.rpc.StackProgress
}
var file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackInput.stack_job:type_name -> cloud.planton.apis.v1.stack.rpc.StackJob
	1, // 1: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackInput.credentials_input:type_name -> cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackCredentialsInput
	2, // 2: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackInput.resource_input:type_name -> cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackResourceInput
	6, // 3: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackCredentialsInput.aws:type_name -> cloud.planton.apis.v1.commons.pulumi.operation.rpc.AwsProviderCredential
	7, // 4: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackResourceInput.dns_zone:type_name -> cloud.planton.apis.v1.code2cloud.deploy.dnszone.state.DnsZoneState
	8, // 5: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackOutputs.zone_status:type_name -> cloud.planton.apis.v1.code2cloud.deploy.dnszone.state.DnsZoneStatusState
	9, // 6: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackResponse.progress:type_name -> cloud.planton.apis.v1.stack.rpc.StackProgress
	3, // 7: cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackResponse.outputs:type_name -> cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws.DnsZoneAwsStackOutputs
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_init() }
func file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsZoneAwsStackInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsZoneAwsStackCredentialsInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsZoneAwsStackResourceInput); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsZoneAwsStackOutputs); i {
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
		file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsZoneAwsStackResponse); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_dnszone_stack_aws_model_proto_depIdxs = nil
}
