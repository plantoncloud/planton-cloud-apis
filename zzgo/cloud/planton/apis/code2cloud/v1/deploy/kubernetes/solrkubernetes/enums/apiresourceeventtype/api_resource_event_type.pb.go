// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/kubernetes/solrkubernetes/enums/apiresourceeventtype/api_resource_event_type.proto

package apiresourceeventtype

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourceeventtypeenumoptions"
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

type SolrKubernetesApiResourceEventType int32

const (
	SolrKubernetesApiResourceEventType_unspecified                         SolrKubernetesApiResourceEventType = 0
	SolrKubernetesApiResourceEventType_created                             SolrKubernetesApiResourceEventType = 1
	SolrKubernetesApiResourceEventType_updated                             SolrKubernetesApiResourceEventType = 2
	SolrKubernetesApiResourceEventType_deleted                             SolrKubernetesApiResourceEventType = 3
	SolrKubernetesApiResourceEventType_restored                            SolrKubernetesApiResourceEventType = 4
	SolrKubernetesApiResourceEventType_refreshed                           SolrKubernetesApiResourceEventType = 5
	SolrKubernetesApiResourceEventType_stack_job_progress_updated          SolrKubernetesApiResourceEventType = 6
	SolrKubernetesApiResourceEventType_stack_job_refresh_requested         SolrKubernetesApiResourceEventType = 7
	SolrKubernetesApiResourceEventType_stack_job_refresh_completed         SolrKubernetesApiResourceEventType = 8
	SolrKubernetesApiResourceEventType_stack_job_apply_preview_requested   SolrKubernetesApiResourceEventType = 9
	SolrKubernetesApiResourceEventType_stack_job_apply_requested           SolrKubernetesApiResourceEventType = 10
	SolrKubernetesApiResourceEventType_stack_job_apply_completed           SolrKubernetesApiResourceEventType = 11
	SolrKubernetesApiResourceEventType_stack_job_destroy_preview_requested SolrKubernetesApiResourceEventType = 12
	SolrKubernetesApiResourceEventType_stack_job_destroy_requested         SolrKubernetesApiResourceEventType = 13
	SolrKubernetesApiResourceEventType_stack_job_destroy_completed         SolrKubernetesApiResourceEventType = 14
	SolrKubernetesApiResourceEventType_paused                              SolrKubernetesApiResourceEventType = 15
	SolrKubernetesApiResourceEventType_unpaused                            SolrKubernetesApiResourceEventType = 16
)

// Enum value maps for SolrKubernetesApiResourceEventType.
var (
	SolrKubernetesApiResourceEventType_name = map[int32]string{
		0:  "unspecified",
		1:  "created",
		2:  "updated",
		3:  "deleted",
		4:  "restored",
		5:  "refreshed",
		6:  "stack_job_progress_updated",
		7:  "stack_job_refresh_requested",
		8:  "stack_job_refresh_completed",
		9:  "stack_job_apply_preview_requested",
		10: "stack_job_apply_requested",
		11: "stack_job_apply_completed",
		12: "stack_job_destroy_preview_requested",
		13: "stack_job_destroy_requested",
		14: "stack_job_destroy_completed",
		15: "paused",
		16: "unpaused",
	}
	SolrKubernetesApiResourceEventType_value = map[string]int32{
		"unspecified":                         0,
		"created":                             1,
		"updated":                             2,
		"deleted":                             3,
		"restored":                            4,
		"refreshed":                           5,
		"stack_job_progress_updated":          6,
		"stack_job_refresh_requested":         7,
		"stack_job_refresh_completed":         8,
		"stack_job_apply_preview_requested":   9,
		"stack_job_apply_requested":           10,
		"stack_job_apply_completed":           11,
		"stack_job_destroy_preview_requested": 12,
		"stack_job_destroy_requested":         13,
		"stack_job_destroy_completed":         14,
		"paused":                              15,
		"unpaused":                            16,
	}
)

func (x SolrKubernetesApiResourceEventType) Enum() *SolrKubernetesApiResourceEventType {
	p := new(SolrKubernetesApiResourceEventType)
	*p = x
	return p
}

func (x SolrKubernetesApiResourceEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SolrKubernetesApiResourceEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_enumTypes[0].Descriptor()
}

func (SolrKubernetesApiResourceEventType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_enumTypes[0]
}

func (x SolrKubernetesApiResourceEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SolrKubernetesApiResourceEventType.Descriptor instead.
func (SolrKubernetesApiResourceEventType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_rawDesc = []byte{
	0x0a, 0x7a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x5c, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x73, 0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x79, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x65,
	0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x8b, 0x05, 0x0a, 0x22, 0x53, 0x6f, 0x6c, 0x72, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b,
	0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x25, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x01, 0x1a, 0x18, 0x80, 0xf9, 0x2b, 0x01,
	0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x03,
	0xa8, 0xf9, 0x2b, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x10,
	0x02, 0x1a, 0x10, 0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0,
	0xf9, 0x2b, 0x03, 0x12, 0x1d, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x03,
	0x1a, 0x10, 0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9,
	0x2b, 0x05, 0x12, 0x1e, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x10, 0x04,
	0x1a, 0x10, 0x88, 0xf9, 0x2b, 0x01, 0x90, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9,
	0x2b, 0x03, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x65, 0x64, 0x10,
	0x05, 0x1a, 0x0c, 0x88, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x03, 0x12,
	0x1e, 0x0a, 0x1a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x10, 0x06, 0x12,
	0x29, 0x0a, 0x1b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x10, 0x07,
	0x1a, 0x08, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x01, 0x12, 0x25, 0x0a, 0x1b, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x08, 0x1a, 0x04, 0xb0, 0xf9, 0x2b,
	0x01, 0x12, 0x2f, 0x0a, 0x21, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x10, 0x09, 0x1a, 0x08, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9,
	0x2b, 0x02, 0x12, 0x2b, 0x0a, 0x19, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x10,
	0x0a, 0x1a, 0x0c, 0x88, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x03, 0x12,
	0x23, 0x0a, 0x19, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x0b, 0x1a, 0x04,
	0xb0, 0xf9, 0x2b, 0x01, 0x12, 0x31, 0x0a, 0x23, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f,
	0x62, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x10, 0x0c, 0x1a, 0x08, 0x98,
	0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x04, 0x12, 0x2d, 0x0a, 0x1b, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x10, 0x0d, 0x1a, 0x0c, 0x88, 0xf9, 0x2b, 0x01, 0x98, 0xf9,
	0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x05, 0x12, 0x25, 0x0a, 0x1b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f,
	0x6a, 0x6f, 0x62, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x0e, 0x1a, 0x04, 0xb0, 0xf9, 0x2b, 0x01, 0x12, 0x18, 0x0a,
	0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x10, 0x0f, 0x1a, 0x0c, 0x88, 0xf9, 0x2b, 0x01, 0x98,
	0xf9, 0x2b, 0x01, 0xa0, 0xf9, 0x2b, 0x03, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x70, 0x61, 0x75,
	0x73, 0x65, 0x64, 0x10, 0x10, 0x1a, 0x0c, 0x88, 0xf9, 0x2b, 0x01, 0x98, 0xf9, 0x2b, 0x01, 0xa0,
	0xf9, 0x2b, 0x03, 0x42, 0xb6, 0x05, 0x0a, 0x6a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x73, 0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x19, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x8c, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0xa2, 0x02,
	0x0a, 0x43, 0x50, 0x41, 0x43, 0x56, 0x44, 0x4b, 0x53, 0x45, 0x41, 0xaa, 0x02, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x53, 0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0xca, 0x02, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c,
	0x53, 0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0xe2, 0x02, 0x68, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x53,
	0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a,
	0x3a, 0x53, 0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_goTypes = []interface{}{
	(SolrKubernetesApiResourceEventType)(0), // 0: cloud.planton.apis.code2cloud.v1.deploy.kubernetes.solrkubernetes.enums.apiresourceeventtype.SolrKubernetesApiResourceEventType
}
var file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_init()
}
func file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_deploy_kubernetes_solrkubernetes_enums_apiresourceeventtype_api_resource_event_type_proto_depIdxs = nil
}
