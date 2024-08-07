// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/integration/v1/kubernetes/cost/model/io.proto

package model

import (
	apiresourcekind "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/enums/apiresourcekind"
	rpc "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/rpc"
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
		mi := &file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostAllocations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostAllocations) ProtoMessage() {}

func (x *CostAllocations) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[0]
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
	return file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDescGZIP(), []int{0}
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
		mi := &file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostAllocationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostAllocationList) ProtoMessage() {}

func (x *CostAllocationList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[1]
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
	return file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDescGZIP(), []int{1}
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
	PageInfo *rpc.PageInfo `protobuf:"bytes,1,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
	// start timestamp
	StartTs *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	// end timestamp
	EndTs *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_ts,json=endTs,proto3" json:"end_ts,omitempty"`
	// id of the environment for which the cost allocations are filtered for the resource requested.
	EnvId string `protobuf:"bytes,4,opt,name=env_id,json=envId,proto3" json:"env_id,omitempty"`
	// type of the resource on planton-cloud.
	ResourceKind apiresourcekind.ApiResourceKind `protobuf:"varint,5,opt,name=resource_kind,json=resourceKind,proto3,enum=cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind" json:"resource_kind,omitempty"`
	// id of the resource on planton-cloud.
	// ex: msk8s-planton-cloud-prod-iam-main which is the id of the microservice.
	ResourceId string `protobuf:"bytes,6,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ListByCostAllocationFiltersInput) Reset() {
	*x = ListByCostAllocationFiltersInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListByCostAllocationFiltersInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByCostAllocationFiltersInput) ProtoMessage() {}

func (x *ListByCostAllocationFiltersInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[2]
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
	return file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDescGZIP(), []int{2}
}

func (x *ListByCostAllocationFiltersInput) GetPageInfo() *rpc.PageInfo {
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

func (x *ListByCostAllocationFiltersInput) GetEnvId() string {
	if x != nil {
		return x.EnvId
	}
	return ""
}

func (x *ListByCostAllocationFiltersInput) GetResourceKind() apiresourcekind.ApiResourceKind {
	if x != nil {
		return x.ResourceKind
	}
	return apiresourcekind.ApiResourceKind(0)
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

	// type of the resource on planton-cloud.
	ResourceKind apiresourcekind.ApiResourceKind `protobuf:"varint,1,opt,name=resource_kind,json=resourceKind,proto3,enum=cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind" json:"resource_kind,omitempty"`
	// id of the resource on planton-cloud.
	// ex: msk8s-planton-cloud-prod-iam-main which is the id of the microservice.
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// start timestamp
	StartTs *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	// end timestamp
	EndTs *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_ts,json=endTs,proto3" json:"end_ts,omitempty"`
}

func (x *GetCostAggregateInput) Reset() {
	*x = GetCostAggregateInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCostAggregateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCostAggregateInput) ProtoMessage() {}

func (x *GetCostAggregateInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[3]
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
	return file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDescGZIP(), []int{3}
}

func (x *GetCostAggregateInput) GetResourceKind() apiresourcekind.ApiResourceKind {
	if x != nil {
		return x.ResourceKind
	}
	return apiresourcekind.ApiResourceKind(0)
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
		mi := &file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostAggregate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostAggregate) ProtoMessage() {}

func (x *CostAggregate) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[4]
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
	return file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDescGZIP(), []int{4}
}

func (x *CostAggregate) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDesc = []byte{
	0x0a, 0x40, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x54, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x0f, 0x43, 0x6f, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x61, 0x0a, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x73, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x98,
	0x01, 0x0a, 0x12, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x61, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0xff, 0x02, 0x0a, 0x20, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x45,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x73, 0x12, 0x31, 0x0a, 0x06,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x54, 0x73, 0x12,
	0x15, 0x0a, 0x06, 0x65, 0x6e, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6e, 0x76, 0x49, 0x64, 0x12, 0x72, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x96, 0x02, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x72, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x2e, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x73, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x65,
	0x6e, 0x64, 0x54, 0x73, 0x22, 0x25, 0x0a, 0x0d, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xc1, 0x03, 0x0a, 0x45,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x49,
	0x56, 0x4b, 0x43, 0x4d, 0xaa, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02,
	0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c,
	0x56, 0x31, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x43, 0x6f,
	0x73, 0x74, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x43, 0x6f, 0x73, 0x74, 0x5c, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x73, 0x74, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDescData = file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDesc
)

func file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDescData
}

var file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_goTypes = []interface{}{
	(*CostAllocations)(nil),                  // 0: cloud.planton.apis.integration.v1.kubernetes.cost.model.CostAllocations
	(*CostAllocationList)(nil),               // 1: cloud.planton.apis.integration.v1.kubernetes.cost.model.CostAllocationList
	(*ListByCostAllocationFiltersInput)(nil), // 2: cloud.planton.apis.integration.v1.kubernetes.cost.model.ListByCostAllocationFiltersInput
	(*GetCostAggregateInput)(nil),            // 3: cloud.planton.apis.integration.v1.kubernetes.cost.model.GetCostAggregateInput
	(*CostAggregate)(nil),                    // 4: cloud.planton.apis.integration.v1.kubernetes.cost.model.CostAggregate
	(*CostAllocation)(nil),                   // 5: cloud.planton.apis.integration.v1.kubernetes.cost.model.CostAllocation
	(*rpc.PageInfo)(nil),                     // 6: cloud.planton.apis.commons.rpc.PageInfo
	(*timestamppb.Timestamp)(nil),            // 7: google.protobuf.Timestamp
	(apiresourcekind.ApiResourceKind)(0),     // 8: cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind
}
var file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_depIdxs = []int32{
	5, // 0: cloud.planton.apis.integration.v1.kubernetes.cost.model.CostAllocations.entries:type_name -> cloud.planton.apis.integration.v1.kubernetes.cost.model.CostAllocation
	5, // 1: cloud.planton.apis.integration.v1.kubernetes.cost.model.CostAllocationList.entries:type_name -> cloud.planton.apis.integration.v1.kubernetes.cost.model.CostAllocation
	6, // 2: cloud.planton.apis.integration.v1.kubernetes.cost.model.ListByCostAllocationFiltersInput.page_info:type_name -> cloud.planton.apis.commons.rpc.PageInfo
	7, // 3: cloud.planton.apis.integration.v1.kubernetes.cost.model.ListByCostAllocationFiltersInput.start_ts:type_name -> google.protobuf.Timestamp
	7, // 4: cloud.planton.apis.integration.v1.kubernetes.cost.model.ListByCostAllocationFiltersInput.end_ts:type_name -> google.protobuf.Timestamp
	8, // 5: cloud.planton.apis.integration.v1.kubernetes.cost.model.ListByCostAllocationFiltersInput.resource_kind:type_name -> cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind
	8, // 6: cloud.planton.apis.integration.v1.kubernetes.cost.model.GetCostAggregateInput.resource_kind:type_name -> cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind
	7, // 7: cloud.planton.apis.integration.v1.kubernetes.cost.model.GetCostAggregateInput.start_ts:type_name -> google.protobuf.Timestamp
	7, // 8: cloud.planton.apis.integration.v1.kubernetes.cost.model.GetCostAggregateInput.end_ts:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_init() }
func file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_init() {
	if File_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto != nil {
		return
	}
	file_cloud_planton_apis_integration_v1_kubernetes_cost_model_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto = out.File
	file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_goTypes = nil
	file_cloud_planton_apis_integration_v1_kubernetes_cost_model_io_proto_depIdxs = nil
}
