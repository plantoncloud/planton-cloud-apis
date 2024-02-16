// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/integration/kubernetes/cost/model/state.proto

package model

import (
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/audit/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/enums"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination/model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// cost allocation
type CostAllocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource audit info
	Audit                      *model.AuditInfo `protobuf:"bytes,99,opt,name=audit,proto3" json:"audit,omitempty"`
	CompanyId                  string           `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	ProductId                  string           `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	KubeClusterId              string           `protobuf:"bytes,3,opt,name=kube_cluster_id,json=kubeClusterId,proto3" json:"kube_cluster_id,omitempty"`
	EnvironmentId              string           `protobuf:"bytes,4,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	ResourceType               string           `protobuf:"bytes,5,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceId                 string           `protobuf:"bytes,6,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Name                       string           `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	StartTs                    string           `protobuf:"bytes,8,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	EndTs                      string           `protobuf:"bytes,9,opt,name=end_ts,json=endTs,proto3" json:"end_ts,omitempty"`
	Minutes                    float64          `protobuf:"fixed64,10,opt,name=minutes,proto3" json:"minutes,omitempty"`
	CpuCores                   float64          `protobuf:"fixed64,11,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	CpuCoreRequestAverage      float64          `protobuf:"fixed64,12,opt,name=cpu_core_request_average,json=cpuCoreRequestAverage,proto3" json:"cpu_core_request_average,omitempty"`
	CpuCoreUsageAverage        float64          `protobuf:"fixed64,13,opt,name=cpu_core_usage_average,json=cpuCoreUsageAverage,proto3" json:"cpu_core_usage_average,omitempty"`
	CpuCoreHours               float64          `protobuf:"fixed64,14,opt,name=cpu_core_hours,json=cpuCoreHours,proto3" json:"cpu_core_hours,omitempty"`
	CpuCost                    float64          `protobuf:"fixed64,15,opt,name=cpu_cost,json=cpuCost,proto3" json:"cpu_cost,omitempty"`
	CpuCostAdjustment          float64          `protobuf:"fixed64,16,opt,name=cpu_cost_adjustment,json=cpuCostAdjustment,proto3" json:"cpu_cost_adjustment,omitempty"`
	CpuEfficiency              float64          `protobuf:"fixed64,17,opt,name=cpu_efficiency,json=cpuEfficiency,proto3" json:"cpu_efficiency,omitempty"`
	GpuCount                   float64          `protobuf:"fixed64,18,opt,name=gpu_count,json=gpuCount,proto3" json:"gpu_count,omitempty"`
	GpuHours                   float64          `protobuf:"fixed64,19,opt,name=gpu_hours,json=gpuHours,proto3" json:"gpu_hours,omitempty"`
	GpuCost                    float64          `protobuf:"fixed64,20,opt,name=gpu_cost,json=gpuCost,proto3" json:"gpu_cost,omitempty"`
	GpuCostAdjustment          float64          `protobuf:"fixed64,21,opt,name=gpu_cost_adjustment,json=gpuCostAdjustment,proto3" json:"gpu_cost_adjustment,omitempty"`
	NetworkTransferBytes       float64          `protobuf:"fixed64,22,opt,name=network_transfer_bytes,json=networkTransferBytes,proto3" json:"network_transfer_bytes,omitempty"`
	NetworkReceiveBytes        float64          `protobuf:"fixed64,23,opt,name=network_receive_bytes,json=networkReceiveBytes,proto3" json:"network_receive_bytes,omitempty"`
	NetworkCost                float64          `protobuf:"fixed64,24,opt,name=network_cost,json=networkCost,proto3" json:"network_cost,omitempty"`
	NetworkCrossZoneCost       float64          `protobuf:"fixed64,25,opt,name=network_cross_zone_cost,json=networkCrossZoneCost,proto3" json:"network_cross_zone_cost,omitempty"`
	NetworkCrossRegionCost     float64          `protobuf:"fixed64,26,opt,name=network_cross_region_cost,json=networkCrossRegionCost,proto3" json:"network_cross_region_cost,omitempty"`
	NetworkInternetCost        float64          `protobuf:"fixed64,27,opt,name=network_internet_cost,json=networkInternetCost,proto3" json:"network_internet_cost,omitempty"`
	NetworkCostAdjustment      float64          `protobuf:"fixed64,28,opt,name=network_cost_adjustment,json=networkCostAdjustment,proto3" json:"network_cost_adjustment,omitempty"`
	LoadBalancerCost           float64          `protobuf:"fixed64,29,opt,name=load_balancer_cost,json=loadBalancerCost,proto3" json:"load_balancer_cost,omitempty"`
	LoadBalancerCostAdjustment float64          `protobuf:"fixed64,30,opt,name=load_balancer_cost_adjustment,json=loadBalancerCostAdjustment,proto3" json:"load_balancer_cost_adjustment,omitempty"`
	PvBytes                    float64          `protobuf:"fixed64,31,opt,name=pv_bytes,json=pvBytes,proto3" json:"pv_bytes,omitempty"`
	PvByteHours                float64          `protobuf:"fixed64,32,opt,name=pv_byte_hours,json=pvByteHours,proto3" json:"pv_byte_hours,omitempty"`
	PvCost                     float64          `protobuf:"fixed64,33,opt,name=pv_cost,json=pvCost,proto3" json:"pv_cost,omitempty"`
	PvCostAdjustment           float64          `protobuf:"fixed64,34,opt,name=pv_cost_adjustment,json=pvCostAdjustment,proto3" json:"pv_cost_adjustment,omitempty"`
	RamBytes                   float64          `protobuf:"fixed64,35,opt,name=ram_bytes,json=ramBytes,proto3" json:"ram_bytes,omitempty"`
	RamByteRequestAverage      float64          `protobuf:"fixed64,36,opt,name=ram_byte_request_average,json=ramByteRequestAverage,proto3" json:"ram_byte_request_average,omitempty"`
	RamByteUsageAverage        float64          `protobuf:"fixed64,37,opt,name=ram_byte_usage_average,json=ramByteUsageAverage,proto3" json:"ram_byte_usage_average,omitempty"`
	RamByteHours               float64          `protobuf:"fixed64,38,opt,name=ram_byte_hours,json=ramByteHours,proto3" json:"ram_byte_hours,omitempty"`
	RamCost                    float64          `protobuf:"fixed64,39,opt,name=ram_cost,json=ramCost,proto3" json:"ram_cost,omitempty"`
	RamCostAdjustment          float64          `protobuf:"fixed64,40,opt,name=ram_cost_adjustment,json=ramCostAdjustment,proto3" json:"ram_cost_adjustment,omitempty"`
	RamEfficiency              float64          `protobuf:"fixed64,41,opt,name=ram_efficiency,json=ramEfficiency,proto3" json:"ram_efficiency,omitempty"`
	ExternalCost               float64          `protobuf:"fixed64,42,opt,name=external_cost,json=externalCost,proto3" json:"external_cost,omitempty"`
	SharedCost                 float64          `protobuf:"fixed64,43,opt,name=shared_cost,json=sharedCost,proto3" json:"shared_cost,omitempty"`
	TotalCost                  float64          `protobuf:"fixed64,44,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	TotalEfficiency            float64          `protobuf:"fixed64,45,opt,name=total_efficiency,json=totalEfficiency,proto3" json:"total_efficiency,omitempty"`
}

func (x *CostAllocation) Reset() {
	*x = CostAllocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostAllocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostAllocation) ProtoMessage() {}

func (x *CostAllocation) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostAllocation.ProtoReflect.Descriptor instead.
func (*CostAllocation) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *CostAllocation) GetAudit() *model.AuditInfo {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *CostAllocation) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CostAllocation) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CostAllocation) GetKubeClusterId() string {
	if x != nil {
		return x.KubeClusterId
	}
	return ""
}

func (x *CostAllocation) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *CostAllocation) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *CostAllocation) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *CostAllocation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CostAllocation) GetStartTs() string {
	if x != nil {
		return x.StartTs
	}
	return ""
}

func (x *CostAllocation) GetEndTs() string {
	if x != nil {
		return x.EndTs
	}
	return ""
}

func (x *CostAllocation) GetMinutes() float64 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

func (x *CostAllocation) GetCpuCores() float64 {
	if x != nil {
		return x.CpuCores
	}
	return 0
}

func (x *CostAllocation) GetCpuCoreRequestAverage() float64 {
	if x != nil {
		return x.CpuCoreRequestAverage
	}
	return 0
}

func (x *CostAllocation) GetCpuCoreUsageAverage() float64 {
	if x != nil {
		return x.CpuCoreUsageAverage
	}
	return 0
}

func (x *CostAllocation) GetCpuCoreHours() float64 {
	if x != nil {
		return x.CpuCoreHours
	}
	return 0
}

func (x *CostAllocation) GetCpuCost() float64 {
	if x != nil {
		return x.CpuCost
	}
	return 0
}

func (x *CostAllocation) GetCpuCostAdjustment() float64 {
	if x != nil {
		return x.CpuCostAdjustment
	}
	return 0
}

func (x *CostAllocation) GetCpuEfficiency() float64 {
	if x != nil {
		return x.CpuEfficiency
	}
	return 0
}

func (x *CostAllocation) GetGpuCount() float64 {
	if x != nil {
		return x.GpuCount
	}
	return 0
}

func (x *CostAllocation) GetGpuHours() float64 {
	if x != nil {
		return x.GpuHours
	}
	return 0
}

func (x *CostAllocation) GetGpuCost() float64 {
	if x != nil {
		return x.GpuCost
	}
	return 0
}

func (x *CostAllocation) GetGpuCostAdjustment() float64 {
	if x != nil {
		return x.GpuCostAdjustment
	}
	return 0
}

func (x *CostAllocation) GetNetworkTransferBytes() float64 {
	if x != nil {
		return x.NetworkTransferBytes
	}
	return 0
}

func (x *CostAllocation) GetNetworkReceiveBytes() float64 {
	if x != nil {
		return x.NetworkReceiveBytes
	}
	return 0
}

func (x *CostAllocation) GetNetworkCost() float64 {
	if x != nil {
		return x.NetworkCost
	}
	return 0
}

func (x *CostAllocation) GetNetworkCrossZoneCost() float64 {
	if x != nil {
		return x.NetworkCrossZoneCost
	}
	return 0
}

func (x *CostAllocation) GetNetworkCrossRegionCost() float64 {
	if x != nil {
		return x.NetworkCrossRegionCost
	}
	return 0
}

func (x *CostAllocation) GetNetworkInternetCost() float64 {
	if x != nil {
		return x.NetworkInternetCost
	}
	return 0
}

func (x *CostAllocation) GetNetworkCostAdjustment() float64 {
	if x != nil {
		return x.NetworkCostAdjustment
	}
	return 0
}

func (x *CostAllocation) GetLoadBalancerCost() float64 {
	if x != nil {
		return x.LoadBalancerCost
	}
	return 0
}

func (x *CostAllocation) GetLoadBalancerCostAdjustment() float64 {
	if x != nil {
		return x.LoadBalancerCostAdjustment
	}
	return 0
}

func (x *CostAllocation) GetPvBytes() float64 {
	if x != nil {
		return x.PvBytes
	}
	return 0
}

func (x *CostAllocation) GetPvByteHours() float64 {
	if x != nil {
		return x.PvByteHours
	}
	return 0
}

func (x *CostAllocation) GetPvCost() float64 {
	if x != nil {
		return x.PvCost
	}
	return 0
}

func (x *CostAllocation) GetPvCostAdjustment() float64 {
	if x != nil {
		return x.PvCostAdjustment
	}
	return 0
}

func (x *CostAllocation) GetRamBytes() float64 {
	if x != nil {
		return x.RamBytes
	}
	return 0
}

func (x *CostAllocation) GetRamByteRequestAverage() float64 {
	if x != nil {
		return x.RamByteRequestAverage
	}
	return 0
}

func (x *CostAllocation) GetRamByteUsageAverage() float64 {
	if x != nil {
		return x.RamByteUsageAverage
	}
	return 0
}

func (x *CostAllocation) GetRamByteHours() float64 {
	if x != nil {
		return x.RamByteHours
	}
	return 0
}

func (x *CostAllocation) GetRamCost() float64 {
	if x != nil {
		return x.RamCost
	}
	return 0
}

func (x *CostAllocation) GetRamCostAdjustment() float64 {
	if x != nil {
		return x.RamCostAdjustment
	}
	return 0
}

func (x *CostAllocation) GetRamEfficiency() float64 {
	if x != nil {
		return x.RamEfficiency
	}
	return 0
}

func (x *CostAllocation) GetExternalCost() float64 {
	if x != nil {
		return x.ExternalCost
	}
	return 0
}

func (x *CostAllocation) GetSharedCost() float64 {
	if x != nil {
		return x.SharedCost
	}
	return 0
}

func (x *CostAllocation) GetTotalCost() float64 {
	if x != nil {
		return x.TotalCost
	}
	return 0
}

func (x *CostAllocation) GetTotalEfficiency() float64 {
	if x != nil {
		return x.TotalEfficiency
	}
	return 0
}

var File_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_rawDesc = []byte{
	0x0a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x0e,
	0x0a, 0x0e, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x4a, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6b, 0x75,
	0x62, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6b, 0x75, 0x62, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x73, 0x12, 0x15,
	0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6e, 0x64, 0x54, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x18,
	0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x15,
	0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x70,
	0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x63,
	0x70, 0x75, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x73,
	0x74, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x70, 0x75, 0x5f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0d, 0x63, 0x70, 0x75, 0x45, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x67, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x70, 0x75, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x67, 0x70, 0x75, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x67, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x67, 0x70, 0x75, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x67, 0x70, 0x75, 0x5f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x67, 0x70, 0x75, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x64, 0x6a,
	0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x32, 0x0a,
	0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x73,
	0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x43, 0x6f, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x63, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x72,
	0x6f, 0x73, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x16,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x6a, 0x75, 0x73,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x15, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10,
	0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x41, 0x0a, 0x1d, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x1a, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x76, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x70, 0x76, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x22,
	0x0a, 0x0d, 0x70, 0x76, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x70, 0x76, 0x42, 0x79, 0x74, 0x65, 0x48, 0x6f, 0x75,
	0x72, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x76, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x21, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x76, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x70,
	0x76, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x22, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x70, 0x76, 0x43, 0x6f, 0x73, 0x74, 0x41,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x6d,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x23, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x72, 0x61,
	0x6d, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x18, 0x24, 0x20, 0x01, 0x28, 0x01, 0x52, 0x15, 0x72, 0x61, 0x6d, 0x42, 0x79, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x33, 0x0a, 0x16, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x25, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x13, 0x72, 0x61, 0x6d, 0x42, 0x79, 0x74, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x26, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x72, 0x61,
	0x6d, 0x42, 0x79, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61,
	0x6d, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x27, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x72, 0x61,
	0x6d, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x11, 0x72, 0x61, 0x6d, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x64, 0x6a, 0x75, 0x73,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x61, 0x6d, 0x5f, 0x65, 0x66, 0x66,
	0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x29, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x72,
	0x61, 0x6d, 0x45, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x2a, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x73, 0x74,
	0x18, 0x2b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x6f,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74,
	0x18, 0x2c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x66, 0x66, 0x69, 0x63,
	0x69, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x45, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x42, 0xc4, 0x03, 0x0a,
	0x45, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x73, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x08,
	0x43, 0x50, 0x41, 0x56, 0x49, 0x4b, 0x43, 0x4d, 0xaa, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0xca, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x5c, 0x43, 0x6f, 0x73, 0x74, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x43, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69,
	0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x43, 0x6f, 0x73, 0x74,
	0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x73, 0x74, 0x3a, 0x3a, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_rawDescData = file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_rawDesc
)

func file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_rawDescData
}

var file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_goTypes = []interface{}{
	(*CostAllocation)(nil),  // 0: cloud.planton.apis.v1.integration.kubernetes.cost.model.CostAllocation
	(*model.AuditInfo)(nil), // 1: cloud.planton.apis.v1.commons.audit.model.AuditInfo
}
var file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_depIdxs = []int32{
	1, // 0: cloud.planton.apis.v1.integration.kubernetes.cost.model.CostAllocation.audit:type_name -> cloud.planton.apis.v1.commons.audit.model.AuditInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_init() }
func file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_init() {
	if File_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostAllocation); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto = out.File
	file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_goTypes = nil
	file_cloud_planton_apis_v1_integration_kubernetes_cost_model_state_proto_depIdxs = nil
}
