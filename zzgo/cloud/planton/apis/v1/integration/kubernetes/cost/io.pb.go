// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/integration/kubernetes/cost/io.proto

package cost

import (
	enums "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/enums"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// list of cost allocations
type CostAllocations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*CostAllocation `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *CostAllocations) Reset() {
	*x = CostAllocations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostAllocations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostAllocations) ProtoMessage() {}

func (x *CostAllocations) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostAllocations.ProtoReflect.Descriptor instead.
func (*CostAllocations) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDescGZIP(), []int{0}
}

func (x *CostAllocations) GetEntries() []*CostAllocation {
	if x != nil {
		return x.Entries
	}
	return nil
}

// response for paginated query to list cost allocations
type CostAllocationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPages int32             `protobuf:"varint,1,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	Entries    []*CostAllocation `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *CostAllocationList) Reset() {
	*x = CostAllocationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostAllocationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostAllocationList) ProtoMessage() {}

func (x *CostAllocationList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostAllocationList.ProtoReflect.Descriptor instead.
func (*CostAllocationList) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDescGZIP(), []int{1}
}

func (x *CostAllocationList) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *CostAllocationList) GetEntries() []*CostAllocation {
	if x != nil {
		return x.Entries
	}
	return nil
}

// input for paginated rpc to get list of cost allocations for a resources based on provided filters.
type ListByCostAllocationFiltersInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// page info
	PageInfo *pagination.PageInfo `protobuf:"bytes,1,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
	// start timestamp
	StartTs *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	// end timestamp
	EndTs *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_ts,json=endTs,proto3" json:"end_ts,omitempty"`
	// id of the product
	ProductId string `protobuf:"bytes,4,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// id of the environment for which the cost allocations are filtered for the resource requested.
	EnvironmentId string `protobuf:"bytes,5,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	// type of the resource on planton cloud.
	ResourceType enums.ResourceType `protobuf:"varint,6,opt,name=resource_type,json=resourceType,proto3,enum=cloud.planton.apis.v1.commons.resource.enums.ResourceType" json:"resource_type,omitempty"`
	// id of the resource on planton cloud.
	// ex: ms-planton-pcs-product which is the id of the microservice.
	ResourceId string `protobuf:"bytes,7,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ListByCostAllocationFiltersInput) Reset() {
	*x = ListByCostAllocationFiltersInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListByCostAllocationFiltersInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByCostAllocationFiltersInput) ProtoMessage() {}

func (x *ListByCostAllocationFiltersInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListByCostAllocationFiltersInput.ProtoReflect.Descriptor instead.
func (*ListByCostAllocationFiltersInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDescGZIP(), []int{2}
}

func (x *ListByCostAllocationFiltersInput) GetPageInfo() *pagination.PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *ListByCostAllocationFiltersInput) GetStartTs() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTs
	}
	return nil
}

func (x *ListByCostAllocationFiltersInput) GetEndTs() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTs
	}
	return nil
}

func (x *ListByCostAllocationFiltersInput) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ListByCostAllocationFiltersInput) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *ListByCostAllocationFiltersInput) GetResourceType() enums.ResourceType {
	if x != nil {
		return x.ResourceType
	}
	return enums.ResourceType(0)
}

func (x *ListByCostAllocationFiltersInput) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type GetCostAggregateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of the resource on planton cloud.
	ResourceType enums.ResourceType `protobuf:"varint,1,opt,name=resource_type,json=resourceType,proto3,enum=cloud.planton.apis.v1.commons.resource.enums.ResourceType" json:"resource_type,omitempty"`
	// id of the resource on planton cloud.
	// ex: ms-planton-pcs-product which is the id of the microservice.
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// start timestamp
	StartTs *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	// end timestamp
	EndTs *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_ts,json=endTs,proto3" json:"end_ts,omitempty"`
}

func (x *GetCostAggregateInput) Reset() {
	*x = GetCostAggregateInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCostAggregateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCostAggregateInput) ProtoMessage() {}

func (x *GetCostAggregateInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCostAggregateInput.ProtoReflect.Descriptor instead.
func (*GetCostAggregateInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDescGZIP(), []int{3}
}

func (x *GetCostAggregateInput) GetResourceType() enums.ResourceType {
	if x != nil {
		return x.ResourceType
	}
	return enums.ResourceType(0)
}

func (x *GetCostAggregateInput) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *GetCostAggregateInput) GetStartTs() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTs
	}
	return nil
}

func (x *GetCostAggregateInput) GetEndTs() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTs
	}
	return nil
}

// wrapper for cost aggregate
type CostAggregate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CostAggregate) Reset() {
	*x = CostAggregate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostAggregate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostAggregate) ProtoMessage() {}

func (x *CostAggregate) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostAggregate.ProtoReflect.Descriptor instead.
func (*CostAggregate) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDescGZIP(), []int{4}
}

func (x *CostAggregate) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x1a,
	0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6e, 0x0a, 0x0f, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x5b, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x92, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0xa5, 0x03, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x43, 0x6f, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x4f, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x08, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05,
	0x65, 0x6e, 0x64, 0x54, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x5f, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x83, 0x02,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x73,
	0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x65, 0x6e,
	0x64, 0x54, 0x73, 0x22, 0x25, 0x0a, 0x0d, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x9b, 0x03, 0x0a, 0x3f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x42, 0x07,
	0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x61, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0xa2, 0x02, 0x07, 0x43,
	0x50, 0x41, 0x56, 0x49, 0x4b, 0x43, 0xaa, 0x02, 0x31, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0xca, 0x02, 0x31, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x43, 0x6f, 0x73, 0x74, 0xe2, 0x02,
	0x3d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x43, 0x6f,
	0x73, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDescData = file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDesc
)

func file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDescData
}

var file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_goTypes = []interface{}{
	(*CostAllocations)(nil),                  // 0: cloud.planton.apis.v1.integration.kubernetes.cost.CostAllocations
	(*CostAllocationList)(nil),               // 1: cloud.planton.apis.v1.integration.kubernetes.cost.CostAllocationList
	(*ListByCostAllocationFiltersInput)(nil), // 2: cloud.planton.apis.v1.integration.kubernetes.cost.ListByCostAllocationFiltersInput
	(*GetCostAggregateInput)(nil),            // 3: cloud.planton.apis.v1.integration.kubernetes.cost.GetCostAggregateInput
	(*CostAggregate)(nil),                    // 4: cloud.planton.apis.v1.integration.kubernetes.cost.CostAggregate
	(*CostAllocation)(nil),                   // 5: cloud.planton.apis.v1.integration.kubernetes.cost.CostAllocation
	(*pagination.PageInfo)(nil),              // 6: cloud.planton.apis.v1.commons.pagination.PageInfo
	(*timestamppb.Timestamp)(nil),            // 7: google.protobuf.Timestamp
	(enums.ResourceType)(0),                  // 8: cloud.planton.apis.v1.commons.resource.enums.ResourceType
}
var file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.v1.integration.kubernetes.cost.CostAllocations.entries:type_name -> cloud.planton.apis.v1.integration.kubernetes.cost.CostAllocation
	5, // 1: cloud.planton.apis.v1.integration.kubernetes.cost.CostAllocationList.entries:type_name -> cloud.planton.apis.v1.integration.kubernetes.cost.CostAllocation
	6, // 2: cloud.planton.apis.v1.integration.kubernetes.cost.ListByCostAllocationFiltersInput.page_info:type_name -> cloud.planton.apis.v1.commons.pagination.PageInfo
	7, // 3: cloud.planton.apis.v1.integration.kubernetes.cost.ListByCostAllocationFiltersInput.start_ts:type_name -> google.protobuf.Timestamp
	7, // 4: cloud.planton.apis.v1.integration.kubernetes.cost.ListByCostAllocationFiltersInput.end_ts:type_name -> google.protobuf.Timestamp
	8, // 5: cloud.planton.apis.v1.integration.kubernetes.cost.ListByCostAllocationFiltersInput.resource_type:type_name -> cloud.planton.apis.v1.commons.resource.enums.ResourceType
	8, // 6: cloud.planton.apis.v1.integration.kubernetes.cost.GetCostAggregateInput.resource_type:type_name -> cloud.planton.apis.v1.commons.resource.enums.ResourceType
	7, // 7: cloud.planton.apis.v1.integration.kubernetes.cost.GetCostAggregateInput.start_ts:type_name -> google.protobuf.Timestamp
	7, // 8: cloud.planton.apis.v1.integration.kubernetes.cost.GetCostAggregateInput.end_ts:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_init() }
func file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_init() {
	if File_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto != nil {
		return
	}
	file_cloud_planton_apis_v1_integration_kubernetes_cost_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostAllocations); i {
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
		file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostAllocationList); i {
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
		file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListByCostAllocationFiltersInput); i {
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
		file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCostAggregateInput); i {
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
		file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostAggregate); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto = out.File
	file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_rawDesc = nil
	file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_goTypes = nil
	file_cloud_planton_apis_v1_integration_kubernetes_cost_io_proto_depIdxs = nil
}
