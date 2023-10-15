// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/iam/authz/role/rpc/model.proto

package rpc

import (
	audit "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/audit"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
	enums "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/enums"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/field/options"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/options"
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
	Metadata *resource.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// resource spec
	Spec *IamRoleSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// resource status
	Status *IamRoleStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *IamRole) Reset() {
	*x = IamRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamRole) ProtoMessage() {}

func (x *IamRole) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[0]
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
	return file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescGZIP(), []int{0}
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

func (x *IamRole) GetMetadata() *resource.Metadata {
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
}

func (x *IamRoleSpec) Reset() {
	*x = IamRoleSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamRoleSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamRoleSpec) ProtoMessage() {}

func (x *IamRoleSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[1]
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
	return file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescGZIP(), []int{1}
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

// status for iam-role
type IamRoleStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// lifecycle of the resource
	Lifecycle *resource.ResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// system audit info
	SysAudit *audit.SysAudit `protobuf:"bytes,98,opt,name=sys_audit,json=sysAudit,proto3" json:"sys_audit,omitempty"`
}

func (x *IamRoleStatus) Reset() {
	*x = IamRoleStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamRoleStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamRoleStatus) ProtoMessage() {}

func (x *IamRoleStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[2]
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
	return file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescGZIP(), []int{2}
}

func (x *IamRoleStatus) GetLifecycle() *resource.ResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *IamRoleStatus) GetSysAudit() *audit.SysAudit {
	if x != nil {
		return x.SysAudit
	}
	return nil
}

// wrapper for iam role id
type IamRoleId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IamRoleId) Reset() {
	*x = IamRoleId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamRoleId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamRoleId) ProtoMessage() {}

func (x *IamRoleId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamRoleId.ProtoReflect.Descriptor instead.
func (*IamRoleId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescGZIP(), []int{3}
}

func (x *IamRoleId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// wrapper for iam resource type input
type IamResourceTypeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IamResourceTypeInput) Reset() {
	*x = IamResourceTypeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamResourceTypeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamResourceTypeInput) ProtoMessage() {}

func (x *IamResourceTypeInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamResourceTypeInput.ProtoReflect.Descriptor instead.
func (*IamResourceTypeInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescGZIP(), []int{4}
}

func (x *IamResourceTypeInput) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// response for rpc query to list all iam roles.
type IamRoles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of iam role entries
	Entries []*IamRole `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *IamRoles) Reset() {
	*x = IamRoles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamRoles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamRoles) ProtoMessage() {}

func (x *IamRoles) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamRoles.ProtoReflect.Descriptor instead.
func (*IamRoles) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescGZIP(), []int{5}
}

func (x *IamRoles) GetEntries() []*IamRole {
	if x != nil {
		return x.Entries
	}
	return nil
}

// response for paginated rpc query to list iam roles.
type IamRoleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of iam role entries
	Entries []*IamRole `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	// total number of pages available to get roles data
	TotalPages int32 `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
}

func (x *IamRoleList) Reset() {
	*x = IamRoleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamRoleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamRoleList) ProtoMessage() {}

func (x *IamRoleList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamRoleList.ProtoReflect.Descriptor instead.
func (*IamRoleList) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescGZIP(), []int{6}
}

func (x *IamRoleList) GetEntries() []*IamRole {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *IamRoleList) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

// input to pass resource type and principal type
type ResourceTypeAndPrincipalTypeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of the resource the role belongs to
	// examples: company/ cloud_account etc
	ResourceType string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// type of the principal that is allowed to link with the resource
	// examples: user/ company etc
	PrincipalType string `protobuf:"bytes,2,opt,name=principal_type,json=principalType,proto3" json:"principal_type,omitempty"`
}

func (x *ResourceTypeAndPrincipalTypeInput) Reset() {
	*x = ResourceTypeAndPrincipalTypeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceTypeAndPrincipalTypeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceTypeAndPrincipalTypeInput) ProtoMessage() {}

func (x *ResourceTypeAndPrincipalTypeInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceTypeAndPrincipalTypeInput.ProtoReflect.Descriptor instead.
func (*ResourceTypeAndPrincipalTypeInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescGZIP(), []int{7}
}

func (x *ResourceTypeAndPrincipalTypeInput) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *ResourceTypeAndPrincipalTypeInput) GetPrincipalType() string {
	if x != nil {
		return x.PrincipalType
	}
	return ""
}

// response to rpc query to get list of all principal types
type PrincipalTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of principal types
	Entries []string `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *PrincipalTypes) Reset() {
	*x = PrincipalTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrincipalTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrincipalTypes) ProtoMessage() {}

func (x *PrincipalTypes) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrincipalTypes.ProtoReflect.Descriptor instead.
func (*PrincipalTypes) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescGZIP(), []int{8}
}

func (x *PrincipalTypes) GetEntries() []string {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63,
	0x1a, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x51, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x07, 0x49, 0x61,
	0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01,
	0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x61,
	0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12,
	0x4f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x61, 0x6d, 0x52, 0x6f,
	0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xb2, 0x01, 0x0a, 0x0b, 0x49, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x0d, 0x49, 0x61, 0x6d, 0x52, 0x6f, 0x6c,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x57, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x12, 0x4a, 0x0a, 0x09, 0x73, 0x79, 0x73, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x62, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x52, 0x08, 0x73, 0x79, 0x73, 0x41, 0x75, 0x64, 0x69, 0x74, 0x22, 0x21, 0x0a, 0x09,
	0x49, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x2c, 0x0a, 0x14, 0x49, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x57, 0x0a,
	0x08, 0x49, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x7b, 0x0a, 0x0b, 0x49, 0x61, 0x6d, 0x52, 0x6f, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x49, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61,
	0x67, 0x65, 0x73, 0x22, 0x6f, 0x0a, 0x21, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x42, 0xea, 0x02, 0x0a, 0x36, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67,
	0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x0a, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x49, 0x41, 0x52, 0x52, 0xaa, 0x02,
	0x28, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41,
	0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x49, 0x61, 0x6d, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x70, 0x63, 0xca, 0x02, 0x28, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56,
	0x31, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x52, 0x6f, 0x6c, 0x65,
	0x5c, 0x52, 0x70, 0x63, 0xe2, 0x02, 0x34, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x61, 0x6d,
	0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x52, 0x6f, 0x6c, 0x65, 0x5c, 0x52, 0x70, 0x63, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2f, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x49, 0x61, 0x6d, 0x3a, 0x3a, 0x41, 0x75, 0x74,
	0x68, 0x7a, 0x3a, 0x3a, 0x52, 0x6f, 0x6c, 0x65, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescData = file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_goTypes = []interface{}{
	(*IamRole)(nil),                           // 0: cloud.planton.apis.v1.iam.authz.role.rpc.IamRole
	(*IamRoleSpec)(nil),                       // 1: cloud.planton.apis.v1.iam.authz.role.rpc.IamRoleSpec
	(*IamRoleStatus)(nil),                     // 2: cloud.planton.apis.v1.iam.authz.role.rpc.IamRoleStatus
	(*IamRoleId)(nil),                         // 3: cloud.planton.apis.v1.iam.authz.role.rpc.IamRoleId
	(*IamResourceTypeInput)(nil),              // 4: cloud.planton.apis.v1.iam.authz.role.rpc.IamResourceTypeInput
	(*IamRoles)(nil),                          // 5: cloud.planton.apis.v1.iam.authz.role.rpc.IamRoles
	(*IamRoleList)(nil),                       // 6: cloud.planton.apis.v1.iam.authz.role.rpc.IamRoleList
	(*ResourceTypeAndPrincipalTypeInput)(nil), // 7: cloud.planton.apis.v1.iam.authz.role.rpc.ResourceTypeAndPrincipalTypeInput
	(*PrincipalTypes)(nil),                    // 8: cloud.planton.apis.v1.iam.authz.role.rpc.PrincipalTypes
	(*resource.Metadata)(nil),                 // 9: cloud.planton.apis.v1.commons.resource.Metadata
	(enums.ResourceType)(0),                   // 10: cloud.planton.apis.v1.commons.resource.enums.ResourceType
	(*resource.ResourceLifecycle)(nil),        // 11: cloud.planton.apis.v1.commons.resource.ResourceLifecycle
	(*audit.SysAudit)(nil),                    // 12: cloud.planton.apis.v1.commons.audit.SysAudit
}
var file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_depIdxs = []int32{
	9,  // 0: cloud.planton.apis.v1.iam.authz.role.rpc.IamRole.metadata:type_name -> cloud.planton.apis.v1.commons.resource.Metadata
	1,  // 1: cloud.planton.apis.v1.iam.authz.role.rpc.IamRole.spec:type_name -> cloud.planton.apis.v1.iam.authz.role.rpc.IamRoleSpec
	2,  // 2: cloud.planton.apis.v1.iam.authz.role.rpc.IamRole.status:type_name -> cloud.planton.apis.v1.iam.authz.role.rpc.IamRoleStatus
	10, // 3: cloud.planton.apis.v1.iam.authz.role.rpc.IamRoleSpec.resource_type:type_name -> cloud.planton.apis.v1.commons.resource.enums.ResourceType
	11, // 4: cloud.planton.apis.v1.iam.authz.role.rpc.IamRoleStatus.lifecycle:type_name -> cloud.planton.apis.v1.commons.resource.ResourceLifecycle
	12, // 5: cloud.planton.apis.v1.iam.authz.role.rpc.IamRoleStatus.sys_audit:type_name -> cloud.planton.apis.v1.commons.audit.SysAudit
	0,  // 6: cloud.planton.apis.v1.iam.authz.role.rpc.IamRoles.entries:type_name -> cloud.planton.apis.v1.iam.authz.role.rpc.IamRole
	0,  // 7: cloud.planton.apis.v1.iam.authz.role.rpc.IamRoleList.entries:type_name -> cloud.planton.apis.v1.iam.authz.role.rpc.IamRole
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_init() }
func file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_init() {
	if File_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamRoleId); i {
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
		file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamResourceTypeInput); i {
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
		file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamRoles); i {
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
		file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamRoleList); i {
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
		file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceTypeAndPrincipalTypeInput); i {
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
		file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrincipalTypes); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto = out.File
	file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_iam_authz_role_rpc_model_proto_depIdxs = nil
}