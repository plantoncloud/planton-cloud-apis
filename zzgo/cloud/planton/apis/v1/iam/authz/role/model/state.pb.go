// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/iam/authz/role/model/state.proto

package model

import (
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/audit/model"
	enums "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/enums"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/field/options"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/model"
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

// iam-role
type IamRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource api version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata for the resource
	// id:
	// identifier for iam role
	//
	// name:
	//
	// name of the iam role.
	Metadata *model.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// resource spec
	Spec *IamRoleSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// resource status
	Status *IamRoleStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *IamRole) Reset() {
	*x = IamRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamRole) ProtoMessage() {}

func (x *IamRole) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamRole.ProtoReflect.Descriptor instead.
func (*IamRole) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *IamRole) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *IamRole) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *IamRole) GetMetadata() *model.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *IamRole) GetSpec() *IamRoleSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *IamRole) GetStatus() *IamRoleStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// specification for iam-role
type IamRoleSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// code of role that is used in FGA model
	RoleCode string `protobuf:"bytes,1,opt,name=role_code,json=roleCode,proto3" json:"role_code,omitempty"`
	// identifies the principal type that would be used to create relation with the resource
	// example values are user/company etc.
	PrincipalType string `protobuf:"bytes,2,opt,name=principal_type,json=principalType,proto3" json:"principal_type,omitempty"`
	// type of the resource the role belongs to
	ResourceType enums.ResourceType `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=cloud.planton.apis.v1.commons.resource.enums.ResourceType" json:"resource_type,omitempty"`
	// description of iam role
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// flag to identify if the role can be used to grant public access
	// by public it is to grant access to all companies in the platform
	IsPublic bool `protobuf:"varint,5,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
}

func (x *IamRoleSpec) Reset() {
	*x = IamRoleSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamRoleSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamRoleSpec) ProtoMessage() {}

func (x *IamRoleSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamRoleSpec.ProtoReflect.Descriptor instead.
func (*IamRoleSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *IamRoleSpec) GetRoleCode() string {
	if x != nil {
		return x.RoleCode
	}
	return ""
}

func (x *IamRoleSpec) GetPrincipalType() string {
	if x != nil {
		return x.PrincipalType
	}
	return ""
}

func (x *IamRoleSpec) GetResourceType() enums.ResourceType {
	if x != nil {
		return x.ResourceType
	}
	return enums.ResourceType(0)
}

func (x *IamRoleSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *IamRoleSpec) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

// status for iam-role
type IamRoleStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// lifecycle of the resource
	Lifecycle *model.ResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// resource audit info
	Audit *model1.ResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
}

func (x *IamRoleStatus) Reset() {
	*x = IamRoleStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamRoleStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamRoleStatus) ProtoMessage() {}

func (x *IamRoleStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamRoleStatus.ProtoReflect.Descriptor instead.
func (*IamRoleStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *IamRoleStatus) GetLifecycle() *model.ResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *IamRoleStatus) GetAudit() *model1.ResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

var File_cloud_planton_apis_v1_iam_authz_role_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x51, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x07, 0x49, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x25,
	0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12,
	0x52, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x4b, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49,
	0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x12, 0x51, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x61,
	0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0xf1, 0x01, 0x0a, 0x0b, 0x49, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69,
	0x70, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0xbe, 0x01, 0x0a, 0x0d, 0x49, 0x61, 0x6d, 0x52,
	0x6f, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5d, 0x0a, 0x09, 0x6c, 0x69, 0x66,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x18, 0x62, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x42, 0xf6, 0x02, 0x0a, 0x38, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2,
	0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x49, 0x41, 0x52, 0x4d, 0xaa, 0x02, 0x2a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e,
	0x56, 0x31, 0x2e, 0x49, 0x61, 0x6d, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x2a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x49, 0x61, 0x6d, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x52, 0x6f, 0x6c, 0x65, 0x5c, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x36, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x61, 0x6d,
	0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x52, 0x6f, 0x6c, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x31,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a,
	0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x49, 0x61, 0x6d, 0x3a, 0x3a, 0x41,
	0x75, 0x74, 0x68, 0x7a, 0x3a, 0x3a, 0x52, 0x6f, 0x6c, 0x65, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDescData = file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDesc
)

func file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDescData
}

var file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_goTypes = []interface{}{
	(*IamRole)(nil),                 // 0: cloud.planton.apis.v1.iam.authz.role.model.IamRole
	(*IamRoleSpec)(nil),             // 1: cloud.planton.apis.v1.iam.authz.role.model.IamRoleSpec
	(*IamRoleStatus)(nil),           // 2: cloud.planton.apis.v1.iam.authz.role.model.IamRoleStatus
	(*model.Metadata)(nil),          // 3: cloud.planton.apis.v1.commons.resource.model.Metadata
	(enums.ResourceType)(0),         // 4: cloud.planton.apis.v1.commons.resource.enums.ResourceType
	(*model.ResourceLifecycle)(nil), // 5: cloud.planton.apis.v1.commons.resource.model.ResourceLifecycle
	(*model1.ResourceAudit)(nil),    // 6: cloud.planton.apis.v1.commons.audit.model.ResourceAudit
}
var file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_depIdxs = []int32{
	3, // 0: cloud.planton.apis.v1.iam.authz.role.model.IamRole.metadata:type_name -> cloud.planton.apis.v1.commons.resource.model.Metadata
	1, // 1: cloud.planton.apis.v1.iam.authz.role.model.IamRole.spec:type_name -> cloud.planton.apis.v1.iam.authz.role.model.IamRoleSpec
	2, // 2: cloud.planton.apis.v1.iam.authz.role.model.IamRole.status:type_name -> cloud.planton.apis.v1.iam.authz.role.model.IamRoleStatus
	4, // 3: cloud.planton.apis.v1.iam.authz.role.model.IamRoleSpec.resource_type:type_name -> cloud.planton.apis.v1.commons.resource.enums.ResourceType
	5, // 4: cloud.planton.apis.v1.iam.authz.role.model.IamRoleStatus.lifecycle:type_name -> cloud.planton.apis.v1.commons.resource.model.ResourceLifecycle
	6, // 5: cloud.planton.apis.v1.iam.authz.role.model.IamRoleStatus.audit:type_name -> cloud.planton.apis.v1.commons.audit.model.ResourceAudit
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_init() }
func file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_init() {
	if File_cloud_planton_apis_v1_iam_authz_role_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamRole); i {
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
		file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamRoleSpec); i {
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
		file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamRoleStatus); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_iam_authz_role_model_state_proto = out.File
	file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_goTypes = nil
	file_cloud_planton_apis_v1_iam_authz_role_model_state_proto_depIdxs = nil
}
