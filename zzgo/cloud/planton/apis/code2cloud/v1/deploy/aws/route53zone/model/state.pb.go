// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/aws/route53zone/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcefieldoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemessageoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions"
	dnsrecordtype "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/networking/enums/dnsrecordtype"
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

// route53-zone
type Route53Zone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata for the resource
	// Name is a valid, unique DNS domain name within the Platon Cloud.
	// Id value is automatically computed in the format 'dns-<org_id>-<normalized-domain-name>',
	// and its uniqueness is guaranteed by the backend.
	Metadata *model.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// route53-zone spec
	Spec *Route53ZoneSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// route53-zone status
	Status *Route53ZoneStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Route53Zone) Reset() {
	*x = Route53Zone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route53Zone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route53Zone) ProtoMessage() {}

func (x *Route53Zone) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route53Zone.ProtoReflect.Descriptor instead.
func (*Route53Zone) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *Route53Zone) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Route53Zone) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Route53Zone) GetMetadata() *model.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Route53Zone) GetSpec() *Route53ZoneSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Route53Zone) GetStatus() *Route53ZoneStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Specification for the Route53 Zone
type Route53ZoneSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// environment-info
	EnvironmentInfo *model1.ApiResourceEnvironmentInfo `protobuf:"bytes,1,opt,name=environment_info,json=environmentInfo,proto3" json:"environment_info,omitempty"`
	// The ID of the AWS Cloud Connection where the Managed Zone should be created.
	AwsCredentialId string `protobuf:"bytes,2,opt,name=aws_credential_id,json=awsCredentialId,proto3" json:"aws_credential_id,omitempty"`
	// The DNS records that are added to the Zone.
	Records []*Route53ZoneDnsRecord `protobuf:"bytes,4,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *Route53ZoneSpec) Reset() {
	*x = Route53ZoneSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route53ZoneSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route53ZoneSpec) ProtoMessage() {}

func (x *Route53ZoneSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route53ZoneSpec.ProtoReflect.Descriptor instead.
func (*Route53ZoneSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *Route53ZoneSpec) GetEnvironmentInfo() *model1.ApiResourceEnvironmentInfo {
	if x != nil {
		return x.EnvironmentInfo
	}
	return nil
}

func (x *Route53ZoneSpec) GetAwsCredentialId() string {
	if x != nil {
		return x.AwsCredentialId
	}
	return ""
}

func (x *Route53ZoneSpec) GetRecords() []*Route53ZoneDnsRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

// Status for the Route53 Zone
type Route53ZoneStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource lifecycle
	Lifecycle *model.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// resource audit info
	Audit *model.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// id of the stack-job
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
	// stack-outputs
	//
	//	stack outputs
	StackOutputs *Route53ZoneStatusStackOutputs `protobuf:"bytes,1,opt,name=stack_outputs,json=stackOutputs,proto3" json:"stack_outputs,omitempty"`
}

func (x *Route53ZoneStatus) Reset() {
	*x = Route53ZoneStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route53ZoneStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route53ZoneStatus) ProtoMessage() {}

func (x *Route53ZoneStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route53ZoneStatus.ProtoReflect.Descriptor instead.
func (*Route53ZoneStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *Route53ZoneStatus) GetLifecycle() *model.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *Route53ZoneStatus) GetAudit() *model.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *Route53ZoneStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

func (x *Route53ZoneStatus) GetStackOutputs() *Route53ZoneStatusStackOutputs {
	if x != nil {
		return x.StackOutputs
	}
	return nil
}

// Status for the Route53 Zone
type Route53ZoneStatusStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of nameservers for the Managed Zone created for the DNS Domain.
	Nameservers []string `protobuf:"bytes,1,rep,name=nameservers,proto3" json:"nameservers,omitempty"`
}

func (x *Route53ZoneStatusStackOutputs) Reset() {
	*x = Route53ZoneStatusStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route53ZoneStatusStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route53ZoneStatusStackOutputs) ProtoMessage() {}

func (x *Route53ZoneStatusStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route53ZoneStatusStackOutputs.ProtoReflect.Descriptor instead.
func (*Route53ZoneStatusStackOutputs) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDescGZIP(), []int{3}
}

func (x *Route53ZoneStatusStackOutputs) GetNameservers() []string {
	if x != nil {
		return x.Nameservers
	}
	return nil
}

// route53-zone dns-record
type Route53ZoneDnsRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// dns record type.
	RecordType dnsrecordtype.DnsRecordType `protobuf:"varint,2,opt,name=record_type,json=recordType,proto3,enum=cloud.planton.apis.commons.networking.enums.dnsrecordtype.DnsRecordType" json:"record_type,omitempty"`
	// name of the route53-zone ex: example.com or dev.example.com.
	// this value should always end with a dot.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// values for the route53-zone record.
	// if the route53_zone_record_type is cname then each value in the list should end with a dot.
	Values []string `protobuf:"bytes,4,rep,name=values,proto3" json:"values,omitempty"`
	// ttl for the domain record in seconds.
	TtlSeconds int32 `protobuf:"varint,5,opt,name=ttl_seconds,json=ttlSeconds,proto3" json:"ttl_seconds,omitempty"`
}

func (x *Route53ZoneDnsRecord) Reset() {
	*x = Route53ZoneDnsRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route53ZoneDnsRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route53ZoneDnsRecord) ProtoMessage() {}

func (x *Route53ZoneDnsRecord) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route53ZoneDnsRecord.ProtoReflect.Descriptor instead.
func (*Route53ZoneDnsRecord) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDescGZIP(), []int{4}
}

func (x *Route53ZoneDnsRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Route53ZoneDnsRecord) GetRecordType() dnsrecordtype.DnsRecordType {
	if x != nil {
		return x.RecordType
	}
	return dnsrecordtype.DnsRecordType(0)
}

func (x *Route53ZoneDnsRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Route53ZoneDnsRecord) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Route53ZoneDnsRecord) GetTtlSeconds() int32 {
	if x != nil {
		return x.TtlSeconds
	}
	return 0
}

var File_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDesc = []byte{
	0x0a, 0x49, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x35, 0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33,
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
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x7c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x64, 0x6e, 0x73, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x06, 0x0a, 0x0b,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0xea, 0x03, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x8a, 0x03, 0xba, 0x48, 0x86, 0x03, 0xba, 0x01, 0xa6,
	0x01, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65,
	0x20, 0x61, 0x6e, 0x79, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x44, 0x4e, 0x53, 0x20, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x6b, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e,
	0x28, 0x3f, 0x3a, 0x5b, 0x5f, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x3f, 0x3a, 0x5b,
	0x5f, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x5b, 0x2e, 0x5d, 0x29, 0x2b, 0x28, 0x3f,
	0x3a, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x28, 0x3f, 0x3a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d,
	0x29, 0x3f, 0x29, 0x3f, 0x24, 0x27, 0x29, 0xba, 0x01, 0x6c, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x4e, 0x61, 0x6d, 0x65, 0x20,
	0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x20,
	0x31, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x36, 0x35, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x73, 0x20, 0x6c, 0x6f, 0x6e, 0x67, 0x1a, 0x2c, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x20, 0x26, 0x26,
	0x20, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29,
	0x20, 0x3c, 0x3d, 0x20, 0x36, 0x35, 0xba, 0x01, 0x67, 0x0a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x30, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x20, 0x69, 0x73, 0x20, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x79, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x29,
	0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x62, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x35,
	0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x35, 0x33, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x68, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x50, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x61, 0x77, 0x73, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x2c, 0x90, 0xb5, 0x18,
	0x01, 0x88, 0xa6, 0x1d, 0x26, 0x92, 0xa6, 0x1d, 0x20, 0x08, 0x0a, 0x12, 0x1c, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x69, 0x64, 0x22, 0xb8, 0x02, 0x0a, 0x0f, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x35, 0x33, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x7e, 0x0a,
	0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a,
	0x11, 0x61, 0x77, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0xd0, 0xb8, 0x18, 0x01, 0x52, 0x0f, 0x61, 0x77, 0x73, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x6d, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x61, 0x77, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33, 0x7a, 0x6f, 0x6e, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33, 0x5a, 0x6f,
	0x6e, 0x65, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x22, 0xf1, 0x02, 0x0a, 0x11, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33,
	0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x60, 0x0a, 0x09, 0x6c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x54, 0x0a, 0x05,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x62, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x05, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f,
	0x69, 0x64, 0x18, 0x61, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a,
	0x6f, 0x62, 0x49, 0x64, 0x12, 0x81, 0x01, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5c, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x35, 0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x35, 0x33, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22, 0x41, 0x0a, 0x1d, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x35, 0x33, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0xa4, 0x03, 0x0a, 0x14,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33, 0x5a, 0x6f, 0x6e, 0x65, 0x44, 0x6e, 0x73, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x71, 0x0a, 0x0b, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x48, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x64, 0x6e,
	0x73, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0xc3, 0x01,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0xae, 0x01, 0xba,
	0x48, 0xaa, 0x01, 0xba, 0x01, 0xa3, 0x01, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x4e,
	0x61, 0x6d, 0x65, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x61, 0x6e,
	0x79, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x44, 0x4e, 0x53, 0x20, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x71, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x28, 0x3f, 0x3a, 0x5b, 0x2a, 0x5d, 0x5b, 0x2e,
	0x5d, 0x29, 0x3f, 0x28, 0x3f, 0x3a, 0x5b, 0x5f, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28,
	0x3f, 0x3a, 0x5b, 0x5f, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x36,
	0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x5b, 0x2e, 0x5d, 0x29,
	0x2b, 0x28, 0x3f, 0x3a, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x28, 0x3f, 0x3a, 0x5b, 0x61, 0x2d, 0x7a,
	0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30,
	0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x29, 0x3f, 0x24, 0x27, 0x29, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x74,
	0x74, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x04, 0x80, 0xb9, 0x18, 0x3c, 0x52, 0x0a, 0x74, 0x74, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x42, 0xea, 0x03, 0x0a, 0x4b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x61, 0x77, 0x73,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x6d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x35, 0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2,
	0x02, 0x09, 0x43, 0x50, 0x41, 0x43, 0x56, 0x44, 0x41, 0x52, 0x4d, 0xaa, 0x02, 0x3d, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x41, 0x77, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x35,
	0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x3d, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x35,
	0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x49, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x35,
	0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x45, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a, 0x41, 0x77, 0x73, 0x3a, 0x3a, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x35, 0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_goTypes = []interface{}{
	(*Route53Zone)(nil),                       // 0: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53Zone
	(*Route53ZoneSpec)(nil),                   // 1: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneSpec
	(*Route53ZoneStatus)(nil),                 // 2: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneStatus
	(*Route53ZoneStatusStackOutputs)(nil),     // 3: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneStatusStackOutputs
	(*Route53ZoneDnsRecord)(nil),              // 4: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneDnsRecord
	(*model.ApiResourceMetadata)(nil),         // 5: cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	(*model1.ApiResourceEnvironmentInfo)(nil), // 6: cloud.planton.apis.resourcemanager.v1.environment.model.ApiResourceEnvironmentInfo
	(*model.ApiResourceLifecycle)(nil),        // 7: cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	(*model.ApiResourceAudit)(nil),            // 8: cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	(dnsrecordtype.DnsRecordType)(0),          // 9: cloud.planton.apis.commons.networking.enums.dnsrecordtype.DnsRecordType
}
var file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53Zone.metadata:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	1, // 1: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53Zone.spec:type_name -> cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneSpec
	2, // 2: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53Zone.status:type_name -> cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneStatus
	6, // 3: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneSpec.environment_info:type_name -> cloud.planton.apis.resourcemanager.v1.environment.model.ApiResourceEnvironmentInfo
	4, // 4: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneSpec.records:type_name -> cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneDnsRecord
	7, // 5: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneStatus.lifecycle:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	8, // 6: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneStatus.audit:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	3, // 7: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneStatus.stack_outputs:type_name -> cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneStatusStackOutputs
	9, // 8: cloud.planton.apis.code2cloud.v1.deploy.aws.route53zone.model.Route53ZoneDnsRecord.record_type:type_name -> cloud.planton.apis.commons.networking.enums.dnsrecordtype.DnsRecordType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route53Zone); i {
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
		file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route53ZoneSpec); i {
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
		file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route53ZoneStatus); i {
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
		file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route53ZoneStatusStackOutputs); i {
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
		file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route53ZoneDnsRecord); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_deploy_aws_route53zone_model_state_proto_depIdxs = nil
}
