// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/audit/v1/apiresourceversion/model/io.proto

package model

import (
	apiresourcekind "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/enums/apiresourcekind"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/rpc"
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

// wrapper for api-resource-version id
type ApiResourceVersionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ApiResourceVersionId) Reset() {
	*x = ApiResourceVersionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResourceVersionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResourceVersionId) ProtoMessage() {}

func (x *ApiResourceVersionId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResourceVersionId.ProtoReflect.Descriptor instead.
func (*ApiResourceVersionId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *ApiResourceVersionId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// The ApiResourceVersionWithContextSizeInput message represents a specific
// api-resource-version along with the context size associated with it. The context size
// is typically used to control the amount of information or data surrounding a
// particular point of interest in the api-resource-version.
type ApiResourceVersionWithContextSizeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier (usually a UUID) for the api-resource-version.
	// It is used to look up or refer to a specific version.
	VersionId string `protobuf:"bytes,1,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// The context size associated with the api-resource-version. This could represent
	// the amount of surrounding data or metadata to include when fetching or
	// working with the version. The meaning of this value can vary depending on
	// the specifics of your application.
	ContextSize int32 `protobuf:"varint,2,opt,name=context_size,json=contextSize,proto3" json:"context_size,omitempty"`
}

func (x *ApiResourceVersionWithContextSizeInput) Reset() {
	*x = ApiResourceVersionWithContextSizeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResourceVersionWithContextSizeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResourceVersionWithContextSizeInput) ProtoMessage() {}

func (x *ApiResourceVersionWithContextSizeInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResourceVersionWithContextSizeInput.ProtoReflect.Descriptor instead.
func (*ApiResourceVersionWithContextSizeInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *ApiResourceVersionWithContextSizeInput) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *ApiResourceVersionWithContextSizeInput) GetContextSize() int32 {
	if x != nil {
		return x.ContextSize
	}
	return 0
}

// list api-resource-versions by resource type and id query input
type ListApiResourceVersionsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// page info
	PageInfo *rpc.PageInfo `protobuf:"bytes,1,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
	// type of resource
	Kind apiresourcekind.ApiResourceKind `protobuf:"varint,2,opt,name=kind,proto3,enum=cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind" json:"kind,omitempty"`
	// id of the resource
	ResourceId string `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ListApiResourceVersionsInput) Reset() {
	*x = ListApiResourceVersionsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApiResourceVersionsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiResourceVersionsInput) ProtoMessage() {}

func (x *ListApiResourceVersionsInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiResourceVersionsInput.ProtoReflect.Descriptor instead.
func (*ListApiResourceVersionsInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescGZIP(), []int{2}
}

func (x *ListApiResourceVersionsInput) GetPageInfo() *rpc.PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *ListApiResourceVersionsInput) GetKind() apiresourcekind.ApiResourceKind {
	if x != nil {
		return x.Kind
	}
	return apiresourcekind.ApiResourceKind(0)
}

func (x *ListApiResourceVersionsInput) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

// paginated list of api-resource-versions
type ApiResourceVersionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// total number of pages
	TotalPages int32 `protobuf:"varint,1,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	// list of api-resource-versions
	Entries []*ApiResourceVersion `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ApiResourceVersionList) Reset() {
	*x = ApiResourceVersionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResourceVersionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResourceVersionList) ProtoMessage() {}

func (x *ApiResourceVersionList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResourceVersionList.ProtoReflect.Descriptor instead.
func (*ApiResourceVersionList) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescGZIP(), []int{3}
}

func (x *ApiResourceVersionList) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *ApiResourceVersionList) GetEntries() []*ApiResourceVersion {
	if x != nil {
		return x.Entries
	}
	return nil
}

// count api-resource-versions by resource type and id query input
type GetApiResourceVersionCountInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of resource
	Kind apiresourcekind.ApiResourceKind `protobuf:"varint,5,opt,name=kind,proto3,enum=cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind" json:"kind,omitempty"`
	// id of the resource
	Id string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetApiResourceVersionCountInput) Reset() {
	*x = GetApiResourceVersionCountInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApiResourceVersionCountInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiResourceVersionCountInput) ProtoMessage() {}

func (x *GetApiResourceVersionCountInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiResourceVersionCountInput.ProtoReflect.Descriptor instead.
func (*GetApiResourceVersionCountInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescGZIP(), []int{4}
}

func (x *GetApiResourceVersionCountInput) GetKind() apiresourcekind.ApiResourceKind {
	if x != nil {
		return x.Kind
	}
	return apiresourcekind.ApiResourceKind(0)
}

func (x *GetApiResourceVersionCountInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// wrapper for api-resource-version count
type ApiResourceVersionCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// value for the count of resource versions
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ApiResourceVersionCount) Reset() {
	*x = ApiResourceVersionCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResourceVersionCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResourceVersionCount) ProtoMessage() {}

func (x *ApiResourceVersionCount) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResourceVersionCount.ProtoReflect.Descriptor instead.
func (*ApiResourceVersionCount) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescGZIP(), []int{5}
}

func (x *ApiResourceVersionCount) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type UpdateStackJobIdInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StackJobId string `protobuf:"bytes,1,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
}

func (x *UpdateStackJobIdInput) Reset() {
	*x = UpdateStackJobIdInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStackJobIdInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStackJobIdInput) ProtoMessage() {}

func (x *UpdateStackJobIdInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStackJobIdInput.ProtoReflect.Descriptor instead.
func (*UpdateStackJobIdInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateStackJobIdInput) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

var File_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x40, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x54, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6b, 0x69, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c,
	0x0a, 0x14, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6a, 0x0a, 0x26,
	0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x69, 0x7a,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xe9, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x45, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x61, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x16, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x62, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x48, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x61, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x17, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a,
	0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x42, 0xad, 0x03, 0x0a, 0x42, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07,
	0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x64, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2,
	0x02, 0x07, 0x43, 0x50, 0x41, 0x41, 0x56, 0x41, 0x4d, 0xaa, 0x02, 0x34, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0xca, 0x02, 0x34, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x41,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x40, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69,
	0x73, 0x3a, 0x3a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescData = file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDesc
)

func file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDescData
}

var file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_goTypes = []interface{}{
	(*ApiResourceVersionId)(nil),                   // 0: cloud.planton.apis.audit.v1.apiresourceversion.model.ApiResourceVersionId
	(*ApiResourceVersionWithContextSizeInput)(nil), // 1: cloud.planton.apis.audit.v1.apiresourceversion.model.ApiResourceVersionWithContextSizeInput
	(*ListApiResourceVersionsInput)(nil),           // 2: cloud.planton.apis.audit.v1.apiresourceversion.model.ListApiResourceVersionsInput
	(*ApiResourceVersionList)(nil),                 // 3: cloud.planton.apis.audit.v1.apiresourceversion.model.ApiResourceVersionList
	(*GetApiResourceVersionCountInput)(nil),        // 4: cloud.planton.apis.audit.v1.apiresourceversion.model.GetApiResourceVersionCountInput
	(*ApiResourceVersionCount)(nil),                // 5: cloud.planton.apis.audit.v1.apiresourceversion.model.ApiResourceVersionCount
	(*UpdateStackJobIdInput)(nil),                  // 6: cloud.planton.apis.audit.v1.apiresourceversion.model.UpdateStackJobIdInput
	(*rpc.PageInfo)(nil),                           // 7: cloud.planton.apis.commons.rpc.PageInfo
	(apiresourcekind.ApiResourceKind)(0),           // 8: cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind
	(*ApiResourceVersion)(nil),                     // 9: cloud.planton.apis.audit.v1.apiresourceversion.model.ApiResourceVersion
}
var file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_depIdxs = []int32{
	7, // 0: cloud.planton.apis.audit.v1.apiresourceversion.model.ListApiResourceVersionsInput.page_info:type_name -> cloud.planton.apis.commons.rpc.PageInfo
	8, // 1: cloud.planton.apis.audit.v1.apiresourceversion.model.ListApiResourceVersionsInput.kind:type_name -> cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind
	9, // 2: cloud.planton.apis.audit.v1.apiresourceversion.model.ApiResourceVersionList.entries:type_name -> cloud.planton.apis.audit.v1.apiresourceversion.model.ApiResourceVersion
	8, // 3: cloud.planton.apis.audit.v1.apiresourceversion.model.GetApiResourceVersionCountInput.kind:type_name -> cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_init() }
func file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_init() {
	if File_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto != nil {
		return
	}
	file_cloud_planton_apis_audit_v1_apiresourceversion_model_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResourceVersionId); i {
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
		file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResourceVersionWithContextSizeInput); i {
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
		file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApiResourceVersionsInput); i {
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
		file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResourceVersionList); i {
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
		file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApiResourceVersionCountInput); i {
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
		file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResourceVersionCount); i {
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
		file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStackJobIdInput); i {
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
			RawDescriptor: file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto = out.File
	file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_goTypes = nil
	file_cloud_planton_apis_audit_v1_apiresourceversion_model_io_proto_depIdxs = nil
}
