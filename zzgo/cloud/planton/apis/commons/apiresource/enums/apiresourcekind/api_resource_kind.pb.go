// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/commons/apiresource/enums/apiresourcekind/api_resource_kind.proto

package apiresourcekind

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcekindenumoptions"
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

type ApiResourceKind int32

const (
	ApiResourceKind_unspecified              ApiResourceKind = 0
	ApiResourceKind_api_resource_version     ApiResourceKind = 1
	ApiResourceKind_artifact_store           ApiResourceKind = 2
	ApiResourceKind_billing_account          ApiResourceKind = 3
	ApiResourceKind_cloud_account            ApiResourceKind = 4
	ApiResourceKind_code_project             ApiResourceKind = 5
	ApiResourceKind_code_server              ApiResourceKind = 6
	ApiResourceKind_organization             ApiResourceKind = 7
	ApiResourceKind_custom_endpoint          ApiResourceKind = 8
	ApiResourceKind_dns_zone                 ApiResourceKind = 9
	ApiResourceKind_environment              ApiResourceKind = 10
	ApiResourceKind_identity_account         ApiResourceKind = 11
	ApiResourceKind_identity_connection      ApiResourceKind = 12
	ApiResourceKind_identity_group           ApiResourceKind = 13
	ApiResourceKind_kube_cluster             ApiResourceKind = 14
	ApiResourceKind_platform                 ApiResourceKind = 15
	ApiResourceKind_stack                    ApiResourceKind = 16
	ApiResourceKind_stack_job                ApiResourceKind = 17
	ApiResourceKind_storage_bucket           ApiResourceKind = 18
	ApiResourceKind_helm_release             ApiResourceKind = 19
	ApiResourceKind_gitlab_kubernetes        ApiResourceKind = 20
	ApiResourceKind_postgres_kubernetes      ApiResourceKind = 21
	ApiResourceKind_jenkins_kubernetes       ApiResourceKind = 22
	ApiResourceKind_kafka_kubernetes         ApiResourceKind = 23
	ApiResourceKind_locust_kubernetes        ApiResourceKind = 24
	ApiResourceKind_microservice_kubernetes  ApiResourceKind = 25
	ApiResourceKind_mongodb_kubernetes       ApiResourceKind = 26
	ApiResourceKind_openfga_kubernetes       ApiResourceKind = 27
	ApiResourceKind_prometheus_kubernetes    ApiResourceKind = 28
	ApiResourceKind_redis_kubernetes         ApiResourceKind = 29
	ApiResourceKind_solr_kubernetes          ApiResourceKind = 30
	ApiResourceKind_argocd_kubernetes        ApiResourceKind = 31
	ApiResourceKind_elasticsearch_kubernetes ApiResourceKind = 32
	ApiResourceKind_grafana_kubernetes       ApiResourceKind = 33
	ApiResourceKind_keycloak_kubernetes      ApiResourceKind = 34
)

// Enum value maps for ApiResourceKind.
var (
	ApiResourceKind_name = map[int32]string{
		0:  "unspecified",
		1:  "api_resource_version",
		2:  "artifact_store",
		3:  "billing_account",
		4:  "cloud_account",
		5:  "code_project",
		6:  "code_server",
		7:  "organization",
		8:  "custom_endpoint",
		9:  "dns_zone",
		10: "environment",
		11: "identity_account",
		12: "identity_connection",
		13: "identity_group",
		14: "kube_cluster",
		15: "platform",
		16: "stack",
		17: "stack_job",
		18: "storage_bucket",
		19: "helm_release",
		20: "gitlab_kubernetes",
		21: "postgres_kubernetes",
		22: "jenkins_kubernetes",
		23: "kafka_kubernetes",
		24: "locust_kubernetes",
		25: "microservice_kubernetes",
		26: "mongodb_kubernetes",
		27: "openfga_kubernetes",
		28: "prometheus_kubernetes",
		29: "redis_kubernetes",
		30: "solr_kubernetes",
		31: "argocd_kubernetes",
		32: "elasticsearch_kubernetes",
		33: "grafana_kubernetes",
		34: "keycloak_kubernetes",
	}
	ApiResourceKind_value = map[string]int32{
		"unspecified":              0,
		"api_resource_version":     1,
		"artifact_store":           2,
		"billing_account":          3,
		"cloud_account":            4,
		"code_project":             5,
		"code_server":              6,
		"organization":             7,
		"custom_endpoint":          8,
		"dns_zone":                 9,
		"environment":              10,
		"identity_account":         11,
		"identity_connection":      12,
		"identity_group":           13,
		"kube_cluster":             14,
		"platform":                 15,
		"stack":                    16,
		"stack_job":                17,
		"storage_bucket":           18,
		"helm_release":             19,
		"gitlab_kubernetes":        20,
		"postgres_kubernetes":      21,
		"jenkins_kubernetes":       22,
		"kafka_kubernetes":         23,
		"locust_kubernetes":        24,
		"microservice_kubernetes":  25,
		"mongodb_kubernetes":       26,
		"openfga_kubernetes":       27,
		"prometheus_kubernetes":    28,
		"redis_kubernetes":         29,
		"solr_kubernetes":          30,
		"argocd_kubernetes":        31,
		"elasticsearch_kubernetes": 32,
		"grafana_kubernetes":       33,
		"keycloak_kubernetes":      34,
	}
)

func (x ApiResourceKind) Enum() *ApiResourceKind {
	p := new(ApiResourceKind)
	*p = x
	return p
}

func (x ApiResourceKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiResourceKind) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_enumTypes[0].Descriptor()
}

func (ApiResourceKind) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_enumTypes[0]
}

func (x ApiResourceKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiResourceKind.Descriptor instead.
func (ApiResourceKind) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_rawDesc = []byte{
	0x0a, 0x54, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6b, 0x69, 0x6e, 0x64, 0x1a, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6b, 0x69, 0x6e, 0x64, 0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e,
	0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x99, 0x0d, 0x0a, 0x0f, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x1a, 0x0f, 0xaa, 0xff, 0x2b, 0x0b, 0x75,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x14, 0x61, 0x70,
	0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x10, 0x01, 0x1a, 0x1d, 0xa2, 0xff, 0x2b, 0x03, 0x76, 0x65, 0x72, 0xaa, 0xff, 0x2b,
	0x12, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x10, 0x02, 0x1a, 0x18, 0xa2, 0xff, 0x2b, 0x03, 0x61, 0x66, 0x73,
	0xaa, 0xff, 0x2b, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x10, 0x03, 0x1a, 0x12, 0xaa, 0xff, 0x2b, 0x0e, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x0d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x04, 0x1a, 0x16, 0xa2,
	0xff, 0x2b, 0x02, 0x63, 0x61, 0xaa, 0xff, 0x2b, 0x0c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x10, 0x05, 0x1a, 0x15, 0xa2, 0xff, 0x2b, 0x02, 0x63, 0x70, 0xaa,
	0xff, 0x2b, 0x0b, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25,
	0x0a, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10, 0x06, 0x1a,
	0x14, 0xa2, 0xff, 0x2b, 0x02, 0x63, 0x73, 0xaa, 0xff, 0x2b, 0x0a, 0x43, 0x6f, 0x64, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x07, 0x1a, 0x10, 0xaa, 0xff, 0x2b, 0x0c, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x0f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x10, 0x08, 0x1a, 0x19,
	0xa2, 0xff, 0x2b, 0x03, 0x63, 0x65, 0x70, 0xaa, 0xff, 0x2b, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x08, 0x64, 0x6e, 0x73,
	0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x10, 0x09, 0x1a, 0x12, 0xa2, 0xff, 0x2b, 0x03, 0x64, 0x6e, 0x73,
	0xaa, 0xff, 0x2b, 0x07, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x0a, 0x1a, 0x0f, 0xaa, 0xff,
	0x2b, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x10, 0x0b, 0x1a, 0x13, 0xaa, 0xff, 0x2b, 0x0f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x0c, 0x1a, 0x16, 0xaa, 0xff, 0x2b, 0x12, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x10, 0x0d, 0x1a, 0x11, 0xaa,
	0xff, 0x2b, 0x0d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x28, 0x0a, 0x0c, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x10, 0x0e, 0x1a, 0x16, 0xa2, 0xff, 0x2b, 0x03, 0x6b, 0x38, 0x63, 0xaa, 0xff, 0x2b, 0x0b, 0x4b,
	0x75, 0x62, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x10, 0x0f, 0x1a, 0x0c, 0xaa, 0xff, 0x2b, 0x08, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x10,
	0x10, 0x1a, 0x0f, 0xa2, 0xff, 0x2b, 0x02, 0x73, 0x74, 0xaa, 0xff, 0x2b, 0x05, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x12, 0x21, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x10,
	0x11, 0x1a, 0x12, 0xa2, 0xff, 0x2b, 0x02, 0x73, 0x6a, 0xaa, 0xff, 0x2b, 0x08, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x2c, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x10, 0x12, 0x1a, 0x18, 0xa2, 0xff, 0x2b, 0x03, 0x62,
	0x6b, 0x74, 0xaa, 0xff, 0x2b, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x0c, 0x68, 0x65, 0x6c, 0x6d, 0x5f, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x10, 0x13, 0x1a, 0x19, 0xa2, 0xff, 0x2b, 0x06, 0x68, 0x6c, 0x6d, 0x6b, 0x38,
	0x73, 0xaa, 0xff, 0x2b, 0x0b, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x11, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x10, 0x14, 0x1a, 0x1d, 0xa2, 0xff, 0x2b, 0x05, 0x67, 0x6c, 0x6b,
	0x38, 0x73, 0xaa, 0xff, 0x2b, 0x10, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x13, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x10, 0x15, 0x1a,
	0x1f, 0xa2, 0xff, 0x2b, 0x05, 0x70, 0x67, 0x6b, 0x38, 0x73, 0xaa, 0xff, 0x2b, 0x12, 0x50, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x12, 0x37, 0x0a, 0x12, 0x6a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x5f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x10, 0x16, 0x1a, 0x1f, 0xa2, 0xff, 0x2b, 0x06, 0x6a, 0x65,
	0x6e, 0x6b, 0x38, 0x73, 0xaa, 0xff, 0x2b, 0x11, 0x4a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x10, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x10, 0x17, 0x1a,
	0x1d, 0xa2, 0xff, 0x2b, 0x06, 0x6b, 0x61, 0x66, 0x6b, 0x38, 0x73, 0xaa, 0xff, 0x2b, 0x0f, 0x4b,
	0x61, 0x66, 0x6b, 0x61, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x35,
	0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x10, 0x18, 0x1a, 0x1e, 0xa2, 0xff, 0x2b, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x38,
	0x73, 0xaa, 0xff, 0x2b, 0x10, 0x4c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x17, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x10, 0x19, 0x1a, 0x23, 0xa2, 0xff, 0x2b, 0x05, 0x6d, 0x73, 0x6b, 0x38, 0x73, 0xaa, 0xff, 0x2b,
	0x16, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x12, 0x6d, 0x6f, 0x6e, 0x67, 0x6f,
	0x64, 0x62, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x10, 0x1a, 0x1a,
	0x1f, 0xa2, 0xff, 0x2b, 0x06, 0x6d, 0x64, 0x62, 0x6b, 0x38, 0x73, 0xaa, 0xff, 0x2b, 0x11, 0x4d,
	0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x12, 0x37, 0x0a, 0x12, 0x6f, 0x70, 0x65, 0x6e, 0x66, 0x67, 0x61, 0x5f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x10, 0x1b, 0x1a, 0x1f, 0xa2, 0xff, 0x2b, 0x06, 0x66, 0x67,
	0x61, 0x6b, 0x38, 0x73, 0xaa, 0xff, 0x2b, 0x11, 0x4f, 0x70, 0x65, 0x6e, 0x66, 0x67, 0x61, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x15, 0x70, 0x72, 0x6f,
	0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x10, 0x1c, 0x1a, 0x22, 0xa2, 0xff, 0x2b, 0x06, 0x70, 0x72, 0x6f, 0x6b, 0x38, 0x73,
	0xaa, 0xff, 0x2b, 0x14, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x10, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x10, 0x1d, 0x1a, 0x1d,
	0xa2, 0xff, 0x2b, 0x06, 0x72, 0x65, 0x64, 0x6b, 0x38, 0x73, 0xaa, 0xff, 0x2b, 0x0f, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x31, 0x0a,
	0x0f, 0x73, 0x6f, 0x6c, 0x72, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x10, 0x1e, 0x1a, 0x1c, 0xa2, 0xff, 0x2b, 0x06, 0x73, 0x6f, 0x6c, 0x6b, 0x38, 0x73, 0xaa, 0xff,
	0x2b, 0x0e, 0x53, 0x6f, 0x6c, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x12, 0x35, 0x0a, 0x11, 0x61, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x10, 0x1f, 0x1a, 0x1e, 0xa2, 0xff, 0x2b, 0x06, 0x61, 0x72, 0x67,
	0x6b, 0x38, 0x73, 0xaa, 0xff, 0x2b, 0x10, 0x41, 0x72, 0x67, 0x6f, 0x63, 0x64, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x18, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x10, 0x20, 0x1a, 0x25, 0xa2, 0xff, 0x2b, 0x06, 0x65, 0x6c, 0x61, 0x6b, 0x38,
	0x73, 0xaa, 0xff, 0x2b, 0x17, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x12,
	0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x10, 0x21, 0x1a, 0x1f, 0xa2, 0xff, 0x2b, 0x06, 0x67, 0x72, 0x61, 0x6b, 0x38, 0x73,
	0xaa, 0xff, 0x2b, 0x11, 0x47, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x13, 0x6b, 0x65, 0x79, 0x63, 0x6c, 0x6f, 0x61,
	0x6b, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x10, 0x22, 0x1a, 0x20,
	0xa2, 0xff, 0x2b, 0x06, 0x6b, 0x65, 0x79, 0x6b, 0x38, 0x73, 0xaa, 0xff, 0x2b, 0x12, 0x4b, 0x65,
	0x79, 0x63, 0x6c, 0x6f, 0x61, 0x6b, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x42, 0xea, 0x03, 0x0a, 0x4a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67,
	0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x42,
	0x14, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6b, 0x69, 0x6e, 0x64, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x43, 0x41, 0x45, 0x41, 0xaa,
	0x02, 0x3c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0xca, 0x02,
	0x3c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x41, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0xe2, 0x02, 0x48,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x41, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x42, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a,
	0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_rawDescData = file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_rawDesc
)

func file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_rawDescData)
	})
	return file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_rawDescData
}

var file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_goTypes = []interface{}{
	(ApiResourceKind)(0), // 0: cloud.planton.apis.commons.apiresource.enums.apiresourcekind.ApiResourceKind
}
var file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_init()
}
func file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_init() {
	if File_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto = out.File
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_rawDesc = nil
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_goTypes = nil
	file_cloud_planton_apis_commons_apiresource_enums_apiresourcekind_api_resource_kind_proto_depIdxs = nil
}
