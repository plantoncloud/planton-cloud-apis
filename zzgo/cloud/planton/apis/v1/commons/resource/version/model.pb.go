// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/commons/resource/version/model.proto

package version

import (
	audit "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/audit"
	enums "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/enums"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/metadata/options"
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

// The ResourceVersion message represents a record of a resource state modification
// in a system's resource. It stores the original and new states of the resource,
// the differences between them in a unified diff format, as well as audit and
// identification information.
type ResourceVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource audit information, which typically includes data about
	// when and by whom the version was created.
	Audit *audit.AuditInfo `protobuf:"bytes,99,opt,name=audit,proto3" json:"audit,omitempty"`
	// A unique identifier (UUID) for the resource-version.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the company to which the resource-version belongs.
	CompanyId string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// The ID of the product to which the resource-version belongs.
	// This attribute is always empty for company level state.
	ProductId string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// The ID of the product environment to which the resource-version belongs.
	// This attribute is always empty for company level state.
	EnvironmentId string `protobuf:"bytes,4,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	// The type of the resource for which the version was created.
	// Resource types are defined in the ResourceType enum.
	ResourceType enums.ResourceType `protobuf:"varint,5,opt,name=resource_type,json=resourceType,proto3,enum=cloud.planton.apis.v1.commons.resource.enums.ResourceType" json:"resource_type,omitempty"`
	// The ID of the resource for which the version was created.
	ResourceId string `protobuf:"bytes,6,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The original state of the resource, represented as a YAML string.
	OriginalStateYaml string `protobuf:"bytes,7,opt,name=original_state_yaml,json=originalStateYaml,proto3" json:"original_state_yaml,omitempty"`
	// The new state of the resource, represented as a YAML string.
	NewStateYaml string `protobuf:"bytes,8,opt,name=new_state_yaml,json=newStateYaml,proto3" json:"new_state_yaml,omitempty"`
	// The differences between the original and new state, represented in a unified diff format.
	DiffUnifiedFormat string `protobuf:"bytes,9,opt,name=diff_unified_format,json=diffUnifiedFormat,proto3" json:"diff_unified_format,omitempty"`
	// id of the stack-job created for the version. this will only be populated for versions created for resources that have stack.
	StackJobId string `protobuf:"bytes,10,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
	// operation done on the state which ended up in creating this new version.
	OperationType string `protobuf:"bytes,11,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
	// version description provided during resource update operation.
	Message string `protobuf:"bytes,12,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResourceVersion) Reset() {
	*x = ResourceVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceVersion) ProtoMessage() {}

func (x *ResourceVersion) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceVersion.ProtoReflect.Descriptor instead.
func (*ResourceVersion) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceVersion) GetAudit() *audit.AuditInfo {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *ResourceVersion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceVersion) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *ResourceVersion) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ResourceVersion) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *ResourceVersion) GetResourceType() enums.ResourceType {
	if x != nil {
		return x.ResourceType
	}
	return enums.ResourceType(0)
}

func (x *ResourceVersion) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *ResourceVersion) GetOriginalStateYaml() string {
	if x != nil {
		return x.OriginalStateYaml
	}
	return ""
}

func (x *ResourceVersion) GetNewStateYaml() string {
	if x != nil {
		return x.NewStateYaml
	}
	return ""
}

func (x *ResourceVersion) GetDiffUnifiedFormat() string {
	if x != nil {
		return x.DiffUnifiedFormat
	}
	return ""
}

func (x *ResourceVersion) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

func (x *ResourceVersion) GetOperationType() string {
	if x != nil {
		return x.OperationType
	}
	return ""
}

func (x *ResourceVersion) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// wrapper for resource-versions id
type ResourceVersionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ResourceVersionId) Reset() {
	*x = ResourceVersionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceVersionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceVersionId) ProtoMessage() {}

func (x *ResourceVersionId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceVersionId.ProtoReflect.Descriptor instead.
func (*ResourceVersionId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceVersionId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// The StateCommitByIdWithContextSize message represents a specific
// resource-version along with the context size associated with it. The context size
// is typically used to control the amount of information or data surrounding a
// particular point of interest in the resource-version.
type ResourceVersionWithContextSizeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier (usually a UUID) for the resource-version.
	// It is used to look up or refer to a specific version.
	VersionId string `protobuf:"bytes,1,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// The context size associated with the resource-version. This could represent
	// the amount of surrounding data or metadata to include when fetching or
	// working with the version. The meaning of this value can vary depending on
	// the specifics of your application.
	ContextSize int32 `protobuf:"varint,2,opt,name=context_size,json=contextSize,proto3" json:"context_size,omitempty"`
}

func (x *ResourceVersionWithContextSizeInput) Reset() {
	*x = ResourceVersionWithContextSizeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceVersionWithContextSizeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceVersionWithContextSizeInput) ProtoMessage() {}

func (x *ResourceVersionWithContextSizeInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceVersionWithContextSizeInput.ProtoReflect.Descriptor instead.
func (*ResourceVersionWithContextSizeInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceVersionWithContextSizeInput) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *ResourceVersionWithContextSizeInput) GetContextSize() int32 {
	if x != nil {
		return x.ContextSize
	}
	return 0
}

// list resource-versions by resource type and id query input
type ListResourceVersionsByFiltersQueryInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// page info
	PageInfo *pagination.PageInfo `protobuf:"bytes,1,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
	// id of the company to filter the resource-versions
	CompanyId string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// id of the product to filter the resource-versions
	ProductId string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// id of the product env to filter the resource-versions
	EnvironmentId string `protobuf:"bytes,4,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	// type of resource
	ResourceType enums.ResourceType `protobuf:"varint,5,opt,name=resource_type,json=resourceType,proto3,enum=cloud.planton.apis.v1.commons.resource.enums.ResourceType" json:"resource_type,omitempty"`
	// id of the resource
	ResourceId string `protobuf:"bytes,6,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ListResourceVersionsByFiltersQueryInput) Reset() {
	*x = ListResourceVersionsByFiltersQueryInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourceVersionsByFiltersQueryInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourceVersionsByFiltersQueryInput) ProtoMessage() {}

func (x *ListResourceVersionsByFiltersQueryInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourceVersionsByFiltersQueryInput.ProtoReflect.Descriptor instead.
func (*ListResourceVersionsByFiltersQueryInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDescGZIP(), []int{3}
}

func (x *ListResourceVersionsByFiltersQueryInput) GetPageInfo() *pagination.PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *ListResourceVersionsByFiltersQueryInput) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *ListResourceVersionsByFiltersQueryInput) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ListResourceVersionsByFiltersQueryInput) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *ListResourceVersionsByFiltersQueryInput) GetResourceType() enums.ResourceType {
	if x != nil {
		return x.ResourceType
	}
	return enums.ResourceType(0)
}

func (x *ListResourceVersionsByFiltersQueryInput) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

// paginated list of resource-versions
type ResourceVersionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// total number of pages
	TotalPages int32 `protobuf:"varint,1,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	// list of resource-versions
	Entries []*ResourceVersion `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ResourceVersionList) Reset() {
	*x = ResourceVersionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceVersionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceVersionList) ProtoMessage() {}

func (x *ResourceVersionList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceVersionList.ProtoReflect.Descriptor instead.
func (*ResourceVersionList) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDescGZIP(), []int{4}
}

func (x *ResourceVersionList) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *ResourceVersionList) GetEntries() []*ResourceVersion {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_cloud_planton_apis_v1_commons_resource_version_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x57, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x04, 0x0a, 0x0f, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44,
	0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x3a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x79, 0x61,
	0x6d, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x59, 0x61, 0x6d, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x6e,
	0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x79, 0x61, 0x6d, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x59, 0x61, 0x6d,
	0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x66, 0x66, 0x5f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x64, 0x69, 0x66, 0x66, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f,
	0x62, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x3a, 0x07, 0x8a, 0xb5, 0x18, 0x03, 0x76, 0x65, 0x72, 0x22, 0x29, 0x0a,
	0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x67, 0x0a, 0x23, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0xe1, 0x02, 0x0a, 0x27, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x4f, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x59,
	0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0x8c, 0x03, 0x0a, 0x3c, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x56, 0x43,
	0x52, 0x56, 0xaa, 0x02, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0xca, 0x02, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x34, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a,
	0x3a, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDescData = file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_commons_resource_version_model_proto_goTypes = []interface{}{
	(*ResourceVersion)(nil),                         // 0: cloud.planton.apis.v1.commons.resource.version.ResourceVersion
	(*ResourceVersionId)(nil),                       // 1: cloud.planton.apis.v1.commons.resource.version.ResourceVersionId
	(*ResourceVersionWithContextSizeInput)(nil),     // 2: cloud.planton.apis.v1.commons.resource.version.ResourceVersionWithContextSizeInput
	(*ListResourceVersionsByFiltersQueryInput)(nil), // 3: cloud.planton.apis.v1.commons.resource.version.ListResourceVersionsByFiltersQueryInput
	(*ResourceVersionList)(nil),                     // 4: cloud.planton.apis.v1.commons.resource.version.ResourceVersionList
	(*audit.AuditInfo)(nil),                         // 5: cloud.planton.apis.v1.commons.audit.AuditInfo
	(enums.ResourceType)(0),                         // 6: cloud.planton.apis.v1.commons.resource.enums.ResourceType
	(*pagination.PageInfo)(nil),                     // 7: cloud.planton.apis.v1.commons.pagination.PageInfo
}
var file_cloud_planton_apis_v1_commons_resource_version_model_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.v1.commons.resource.version.ResourceVersion.audit:type_name -> cloud.planton.apis.v1.commons.audit.AuditInfo
	6, // 1: cloud.planton.apis.v1.commons.resource.version.ResourceVersion.resource_type:type_name -> cloud.planton.apis.v1.commons.resource.enums.ResourceType
	7, // 2: cloud.planton.apis.v1.commons.resource.version.ListResourceVersionsByFiltersQueryInput.page_info:type_name -> cloud.planton.apis.v1.commons.pagination.PageInfo
	6, // 3: cloud.planton.apis.v1.commons.resource.version.ListResourceVersionsByFiltersQueryInput.resource_type:type_name -> cloud.planton.apis.v1.commons.resource.enums.ResourceType
	0, // 4: cloud.planton.apis.v1.commons.resource.version.ResourceVersionList.entries:type_name -> cloud.planton.apis.v1.commons.resource.version.ResourceVersion
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_resource_version_model_proto_init() }
func file_cloud_planton_apis_v1_commons_resource_version_model_proto_init() {
	if File_cloud_planton_apis_v1_commons_resource_version_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceVersion); i {
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
		file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceVersionId); i {
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
		file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceVersionWithContextSizeInput); i {
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
		file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourceVersionsByFiltersQueryInput); i {
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
		file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceVersionList); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_resource_version_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_resource_version_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_commons_resource_version_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_resource_version_model_proto = out.File
	file_cloud_planton_apis_v1_commons_resource_version_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_resource_version_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_resource_version_model_proto_depIdxs = nil
}
