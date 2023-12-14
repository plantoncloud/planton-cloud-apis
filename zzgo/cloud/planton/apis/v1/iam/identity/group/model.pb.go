// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/iam/identity/group/model.proto

package group

import (
	audit "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/audit"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/enums"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/field/options"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/options"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
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

// identity-group
type IdentityGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource api version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// resource metadata
	Metadata *resource.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// resource spec
	Spec *IdentityGroupSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// resource status
	Status *IdentityGroupStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *IdentityGroup) Reset() {
	*x = IdentityGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityGroup) ProtoMessage() {}

func (x *IdentityGroup) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityGroup.ProtoReflect.Descriptor instead.
func (*IdentityGroup) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescGZIP(), []int{0}
}

func (x *IdentityGroup) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *IdentityGroup) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *IdentityGroup) GetMetadata() *resource.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *IdentityGroup) GetSpec() *IdentityGroupSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *IdentityGroup) GetStatus() *IdentityGroupStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// specification for identity-group
type IdentityGroupSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// company id of the group
	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *IdentityGroupSpec) Reset() {
	*x = IdentityGroupSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityGroupSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityGroupSpec) ProtoMessage() {}

func (x *IdentityGroupSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityGroupSpec.ProtoReflect.Descriptor instead.
func (*IdentityGroupSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescGZIP(), []int{1}
}

func (x *IdentityGroupSpec) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

// status for identity-group
type IdentityGroupStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// lifecycle of the resource
	Lifecycle *resource.ResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// resource audit info
	Audit *audit.ResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
}

func (x *IdentityGroupStatus) Reset() {
	*x = IdentityGroupStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityGroupStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityGroupStatus) ProtoMessage() {}

func (x *IdentityGroupStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityGroupStatus.ProtoReflect.Descriptor instead.
func (*IdentityGroupStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescGZIP(), []int{2}
}

func (x *IdentityGroupStatus) GetLifecycle() *resource.ResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *IdentityGroupStatus) GetAudit() *audit.ResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

// wrapper for iam group id
type IdentityGroupId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IdentityGroupId) Reset() {
	*x = IdentityGroupId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityGroupId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityGroupId) ProtoMessage() {}

func (x *IdentityGroupId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityGroupId.ProtoReflect.Descriptor instead.
func (*IdentityGroupId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescGZIP(), []int{3}
}

func (x *IdentityGroupId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// wrapper for iam group name
type IdentityGroupName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IdentityGroupName) Reset() {
	*x = IdentityGroupName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityGroupName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityGroupName) ProtoMessage() {}

func (x *IdentityGroupName) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityGroupName.ProtoReflect.Descriptor instead.
func (*IdentityGroupName) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescGZIP(), []int{4}
}

func (x *IdentityGroupName) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// list of iam groups
type IdentityGroups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*IdentityGroup `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *IdentityGroups) Reset() {
	*x = IdentityGroups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityGroups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityGroups) ProtoMessage() {}

func (x *IdentityGroups) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityGroups.ProtoReflect.Descriptor instead.
func (*IdentityGroups) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescGZIP(), []int{5}
}

func (x *IdentityGroups) GetEntries() []*IdentityGroup {
	if x != nil {
		return x.Entries
	}
	return nil
}

// response for paginated rpc query to list iam groups.
type IdentityGroupList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries    []*IdentityGroup `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	TotalPages int32            `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
}

func (x *IdentityGroupList) Reset() {
	*x = IdentityGroupList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityGroupList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityGroupList) ProtoMessage() {}

func (x *IdentityGroupList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityGroupList.ProtoReflect.Descriptor instead.
func (*IdentityGroupList) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescGZIP(), []int{6}
}

func (x *IdentityGroupList) GetEntries() []*IdentityGroup {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *IdentityGroupList) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

type ListWithCompanyIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// company id of the iam group
	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// pagination input
	Page *pagination.PageInfo `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListWithCompanyIdReq) Reset() {
	*x = ListWithCompanyIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWithCompanyIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWithCompanyIdReq) ProtoMessage() {}

func (x *ListWithCompanyIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWithCompanyIdReq.ProtoReflect.Descriptor instead.
func (*ListWithCompanyIdReq) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescGZIP(), []int{7}
}

func (x *ListWithCompanyIdReq) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *ListWithCompanyIdReq) GetPage() *pagination.PageInfo {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_cloud_planton_apis_v1_iam_identity_group_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
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
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x02, 0x0a, 0x0d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x25, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01,
	0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x4f, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x55, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x1b, 0x88, 0xa6,
	0x1d, 0x0e, 0x9a, 0xa6, 0x1d, 0x13, 0x08, 0x06, 0x12, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x11, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x12, 0x27,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xc0, 0xb8, 0x18, 0x01, 0xd0, 0xb8, 0x18, 0x01, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x13, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x57, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x63, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x18, 0x62, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x05, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x22, 0x27, 0x0a, 0x0f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x63, 0x0a, 0x0e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x51, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x11,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x51, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x61, 0x67, 0x65, 0x73, 0x22, 0x7d, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x42, 0xe8, 0x02, 0x0a, 0x36, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42,
	0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x56, 0x49, 0x49,
	0x47, 0xaa, 0x02, 0x28, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x49, 0x61, 0x6d, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0xca, 0x02, 0x28, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69,
	0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0xe2, 0x02, 0x34, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x49, 0x61, 0x6d, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x49, 0x61, 0x6d, 0x3a, 0x3a,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescData = file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cloud_planton_apis_v1_iam_identity_group_model_proto_goTypes = []interface{}{
	(*IdentityGroup)(nil),              // 0: cloud.planton.apis.v1.iam.identity.group.IdentityGroup
	(*IdentityGroupSpec)(nil),          // 1: cloud.planton.apis.v1.iam.identity.group.IdentityGroupSpec
	(*IdentityGroupStatus)(nil),        // 2: cloud.planton.apis.v1.iam.identity.group.IdentityGroupStatus
	(*IdentityGroupId)(nil),            // 3: cloud.planton.apis.v1.iam.identity.group.IdentityGroupId
	(*IdentityGroupName)(nil),          // 4: cloud.planton.apis.v1.iam.identity.group.IdentityGroupName
	(*IdentityGroups)(nil),             // 5: cloud.planton.apis.v1.iam.identity.group.IdentityGroups
	(*IdentityGroupList)(nil),          // 6: cloud.planton.apis.v1.iam.identity.group.IdentityGroupList
	(*ListWithCompanyIdReq)(nil),       // 7: cloud.planton.apis.v1.iam.identity.group.ListWithCompanyIdReq
	(*resource.Metadata)(nil),          // 8: cloud.planton.apis.v1.commons.resource.Metadata
	(*resource.ResourceLifecycle)(nil), // 9: cloud.planton.apis.v1.commons.resource.ResourceLifecycle
	(*audit.ResourceAudit)(nil),        // 10: cloud.planton.apis.v1.commons.audit.ResourceAudit
	(*pagination.PageInfo)(nil),        // 11: cloud.planton.apis.v1.commons.pagination.PageInfo
}
var file_cloud_planton_apis_v1_iam_identity_group_model_proto_depIdxs = []int32{
	8,  // 0: cloud.planton.apis.v1.iam.identity.group.IdentityGroup.metadata:type_name -> cloud.planton.apis.v1.commons.resource.Metadata
	1,  // 1: cloud.planton.apis.v1.iam.identity.group.IdentityGroup.spec:type_name -> cloud.planton.apis.v1.iam.identity.group.IdentityGroupSpec
	2,  // 2: cloud.planton.apis.v1.iam.identity.group.IdentityGroup.status:type_name -> cloud.planton.apis.v1.iam.identity.group.IdentityGroupStatus
	9,  // 3: cloud.planton.apis.v1.iam.identity.group.IdentityGroupStatus.lifecycle:type_name -> cloud.planton.apis.v1.commons.resource.ResourceLifecycle
	10, // 4: cloud.planton.apis.v1.iam.identity.group.IdentityGroupStatus.audit:type_name -> cloud.planton.apis.v1.commons.audit.ResourceAudit
	0,  // 5: cloud.planton.apis.v1.iam.identity.group.IdentityGroups.entries:type_name -> cloud.planton.apis.v1.iam.identity.group.IdentityGroup
	0,  // 6: cloud.planton.apis.v1.iam.identity.group.IdentityGroupList.entries:type_name -> cloud.planton.apis.v1.iam.identity.group.IdentityGroup
	11, // 7: cloud.planton.apis.v1.iam.identity.group.ListWithCompanyIdReq.page:type_name -> cloud.planton.apis.v1.commons.pagination.PageInfo
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_iam_identity_group_model_proto_init() }
func file_cloud_planton_apis_v1_iam_identity_group_model_proto_init() {
	if File_cloud_planton_apis_v1_iam_identity_group_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityGroup); i {
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
		file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityGroupSpec); i {
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
		file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityGroupStatus); i {
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
		file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityGroupId); i {
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
		file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityGroupName); i {
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
		file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityGroups); i {
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
		file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityGroupList); i {
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
		file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWithCompanyIdReq); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_iam_identity_group_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_iam_identity_group_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_iam_identity_group_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_iam_identity_group_model_proto = out.File
	file_cloud_planton_apis_v1_iam_identity_group_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_iam_identity_group_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_iam_identity_group_model_proto_depIdxs = nil
}
