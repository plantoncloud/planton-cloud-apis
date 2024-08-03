// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/connect/v1/gcpcredential/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcefieldoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemessageoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions"
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

// gcp-credential
type GcpCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource-kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata
	// id format "<id-prefix>-<org-id>-<normalized-resource-name>"
	Metadata *model.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *GcpCredentialSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *model.ApiResourceLifecycleAndAuditStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GcpCredential) Reset() {
	*x = GcpCredential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpCredential) ProtoMessage() {}

func (x *GcpCredential) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpCredential.ProtoReflect.Descriptor instead.
func (*GcpCredential) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *GcpCredential) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *GcpCredential) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *GcpCredential) GetMetadata() *model.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GcpCredential) GetSpec() *GcpCredentialSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *GcpCredential) GetStatus() *model.ApiResourceLifecycleAndAuditStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// gcp-credential spec
type GcpCredentialSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the organization on planton cloud to which the credential belongs to
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// The ID of the Google Cloud Organization, required and immutable. For more information,
	// visit: https://cloud.google.com/resource-manager/docs/creating-managing-organization
	GcpOrganizationId string `protobuf:"bytes,2,opt,name=gcp_organization_id,json=gcpOrganizationId,proto3" json:"gcp_organization_id,omitempty"`
	// base64 encoded google service account key
	ServiceAccountKeyBase64 string `protobuf:"bytes,3,opt,name=service_account_key_base64,json=serviceAccountKeyBase64,proto3" json:"service_account_key_base64,omitempty"`
}

func (x *GcpCredentialSpec) Reset() {
	*x = GcpCredentialSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpCredentialSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpCredentialSpec) ProtoMessage() {}

func (x *GcpCredentialSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpCredentialSpec.ProtoReflect.Descriptor instead.
func (*GcpCredentialSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *GcpCredentialSpec) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *GcpCredentialSpec) GetGcpOrganizationId() string {
	if x != nil {
		return x.GcpOrganizationId
	}
	return ""
}

func (x *GcpCredentialSpec) GetServiceAccountKeyBase64() string {
	if x != nil {
		return x.ServiceAccountKeyBase64
	}
	return ""
}

var File_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x63, 0x70, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x67,
	0x63, 0x70, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x6b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x7c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x07,
	0x0a, 0x0d, 0x47, 0x63, 0x70, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12,
	0x40, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0xba, 0x48, 0x1c, 0x72, 0x1a, 0x0a, 0x18, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x31, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0xd9, 0x04, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xf9, 0x03, 0xba, 0x48,
	0xf5, 0x03, 0xba, 0x01, 0x6c, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x4f, 0x6e, 0x6c, 0x79, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x63, 0x61, 0x73, 0x65, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x68, 0x79, 0x70, 0x68, 0x65,
	0x6e, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x1a, 0x21,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x2b, 0x24, 0x27,
	0x29, 0xba, 0x01, 0x58, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x27, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x4d, 0x75, 0x73, 0x74, 0x20, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x63, 0x61, 0x73, 0x65, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x1e, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28,
	0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x2e, 0x2a, 0x24, 0x27, 0x29, 0xba, 0x01, 0x4c, 0x0a,
	0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x4e, 0x61, 0x6d, 0x65, 0x20, 0x4d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x6e,
	0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20, 0x68, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x1a,
	0x1a, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x28, 0x27, 0x5b, 0x5e, 0x2d, 0x5d, 0x24, 0x27, 0x29, 0xba, 0x01, 0x6c, 0x0a, 0x0d,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x4e,
	0x61, 0x6d, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x62, 0x65, 0x74, 0x77,
	0x65, 0x65, 0x6e, 0x20, 0x31, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x33, 0x30, 0x20, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x20, 0x6c, 0x6f, 0x6e, 0x67, 0x1a, 0x2c, 0x73, 0x69,
	0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20, 0x3e, 0x20,
	0x30, 0x20, 0x26, 0x26, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x29, 0x20, 0x3c, 0x3d, 0x20, 0x33, 0x30, 0xba, 0x01, 0x67, 0x0a, 0x18, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x69, 0x73, 0x20, 0x6d, 0x61, 0x6e, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x79, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x68, 0x61, 0x73, 0x28, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x29, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x58, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x44, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x67, 0x63, 0x70, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x47, 0x63, 0x70, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x68, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x41,
	0x6e, 0x64, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x17, 0x88, 0xa6, 0x1d, 0x0c, 0x92, 0xa6, 0x1d, 0x0f, 0x08,
	0x20, 0x12, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x22, 0x83,
	0x02, 0x0a, 0x11, 0x47, 0x63, 0x70, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x12, 0x8b, 0x01, 0x0a, 0x13, 0x67, 0x63, 0x70, 0x5f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x5b, 0xba, 0x48, 0x54, 0xba, 0x01, 0x4b, 0x0a, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x67, 0x63, 0x70, 0x2e, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x12, 0x1e, 0x4f, 0x6e, 0x6c, 0x79,
	0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x20, 0x28, 0x30, 0x2d, 0x39, 0x29, 0x20, 0x61,
	0x72, 0x65, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x1a, 0x18, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x5d,
	0x2b, 0x24, 0x27, 0x29, 0xc8, 0x01, 0x01, 0xd0, 0x01, 0x01, 0xd0, 0xb8, 0x18, 0x01, 0x52, 0x11,
	0x67, 0x63, 0x70, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x41, 0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x17, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x61,
	0x73, 0x65, 0x36, 0x34, 0x42, 0x9e, 0x03, 0x0a, 0x3f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x63, 0x70, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x61, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x63, 0x70, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x43,
	0x56, 0x47, 0x4d, 0xaa, 0x02, 0x31, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x56, 0x31, 0x2e, 0x47, 0x63, 0x70, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x31, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x63, 0x70, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x3d, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x63, 0x70, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x37, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x47, 0x63, 0x70, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x3a, 0x3a,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDescData = file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDesc
)

func file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDescData
}

var file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_goTypes = []interface{}{
	(*GcpCredential)(nil),                            // 0: cloud.planton.apis.connect.v1.gcpcredential.model.GcpCredential
	(*GcpCredentialSpec)(nil),                        // 1: cloud.planton.apis.connect.v1.gcpcredential.model.GcpCredentialSpec
	(*model.ApiResourceMetadata)(nil),                // 2: cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	(*model.ApiResourceLifecycleAndAuditStatus)(nil), // 3: cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycleAndAuditStatus
}
var file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_depIdxs = []int32{
	2, // 0: cloud.planton.apis.connect.v1.gcpcredential.model.GcpCredential.metadata:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	1, // 1: cloud.planton.apis.connect.v1.gcpcredential.model.GcpCredential.spec:type_name -> cloud.planton.apis.connect.v1.gcpcredential.model.GcpCredentialSpec
	3, // 2: cloud.planton.apis.connect.v1.gcpcredential.model.GcpCredential.status:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycleAndAuditStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_init() }
func file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_init() {
	if File_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpCredential); i {
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
		file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpCredentialSpec); i {
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
			RawDescriptor: file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto = out.File
	file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_goTypes = nil
	file_cloud_planton_apis_connect_v1_gcpcredential_model_state_proto_depIdxs = nil
}
