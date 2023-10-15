// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/commons/english/rpc/enums/enums.proto

package enums

import (
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

type Word_WordEnum int32

const (
	Word_word_unspecified Word_WordEnum = 0
	Word_app              Word_WordEnum = 1
	Word_application      Word_WordEnum = 2
	Word_aws              Word_WordEnum = 3
	Word_azure            Word_WordEnum = 4
	Word_bill             Word_WordEnum = 5
	Word_company          Word_WordEnum = 6
	Word_config           Word_WordEnum = 7
	Word_cpu              Word_WordEnum = 8
	Word_debug            Word_WordEnum = 9
	Word_docker           Word_WordEnum = 10
	Word_domain           Word_WordEnum = 11
	Word_email            Word_WordEnum = 12
	Word_endpoint         Word_WordEnum = 13
	Word_external         Word_WordEnum = 14
	Word_file             Word_WordEnum = 15
	Word_gcp              Word_WordEnum = 16
	Word_google           Word_WordEnum = 17
	Word_hosting          Word_WordEnum = 18
	Word_hostname         Word_WordEnum = 19
	Word_id               Word_WordEnum = 20
	Word_ingress          Word_WordEnum = 21
	Word_instances        Word_WordEnum = 22
	Word_internal         Word_WordEnum = 23
	Word_ip               Word_WordEnum = 24
	Word_key              Word_WordEnum = 25
	Word_kubeconfig       Word_WordEnum = 26
	Word_kubernetes       Word_WordEnum = 27
	Word_localhost        Word_WordEnum = 28
	Word_machine          Word_WordEnum = 29
	Word_main             Word_WordEnum = 30
	Word_maven            Word_WordEnum = 31
	Word_memory           Word_WordEnum = 32
	Word_microservice     Word_WordEnum = 33
	Word_name             Word_WordEnum = 34
	Word_nameservers      Word_WordEnum = 35
	Word_namespace        Word_WordEnum = 36
	Word_network          Word_WordEnum = 37
	Word_npm              Word_WordEnum = 38
	Word_number           Word_WordEnum = 39
	Word_org              Word_WordEnum = 40
	Word_parent           Word_WordEnum = 41
	Word_postgres         Word_WordEnum = 42
	Word_product          Word_WordEnum = 43
	Word_project          Word_WordEnum = 44
	Word_python           Word_WordEnum = 45
	Word_repo             Word_WordEnum = 46
	Word_review           Word_WordEnum = 47
	Word_server           Word_WordEnum = 48
	Word_servers          Word_WordEnum = 49
	Word_share            Word_WordEnum = 50
	Word_shared_services  Word_WordEnum = 51
	Word_spilo            Word_WordEnum = 52
	Word_spot             Word_WordEnum = 53
	Word_stunnel          Word_WordEnum = 54
	Word_team             Word_WordEnum = 55
	Word_type             Word_WordEnum = 56
	Word_username         Word_WordEnum = 57
	Word_util             Word_WordEnum = 58
	Word_version          Word_WordEnum = 59
	Word_workload         Word_WordEnum = 60
	Word_url              Word_WordEnum = 61
	Word_env              Word_WordEnum = 62
	Word_kafka            Word_WordEnum = 63
	Word_unknown          Word_WordEnum = 64
	Word_resource         Word_WordEnum = 65
	Word_resource_id      Word_WordEnum = 66
	Word_resource_type    Word_WordEnum = 67
	Word_environment      Word_WordEnum = 68
)

// Enum value maps for Word_WordEnum.
var (
	Word_WordEnum_name = map[int32]string{
		0:  "word_unspecified",
		1:  "app",
		2:  "application",
		3:  "aws",
		4:  "azure",
		5:  "bill",
		6:  "company",
		7:  "config",
		8:  "cpu",
		9:  "debug",
		10: "docker",
		11: "domain",
		12: "email",
		13: "endpoint",
		14: "external",
		15: "file",
		16: "gcp",
		17: "google",
		18: "hosting",
		19: "hostname",
		20: "id",
		21: "ingress",
		22: "instances",
		23: "internal",
		24: "ip",
		25: "key",
		26: "kubeconfig",
		27: "kubernetes",
		28: "localhost",
		29: "machine",
		30: "main",
		31: "maven",
		32: "memory",
		33: "microservice",
		34: "name",
		35: "nameservers",
		36: "namespace",
		37: "network",
		38: "npm",
		39: "number",
		40: "org",
		41: "parent",
		42: "postgres",
		43: "product",
		44: "project",
		45: "python",
		46: "repo",
		47: "review",
		48: "server",
		49: "servers",
		50: "share",
		51: "shared_services",
		52: "spilo",
		53: "spot",
		54: "stunnel",
		55: "team",
		56: "type",
		57: "username",
		58: "util",
		59: "version",
		60: "workload",
		61: "url",
		62: "env",
		63: "kafka",
		64: "unknown",
		65: "resource",
		66: "resource_id",
		67: "resource_type",
		68: "environment",
	}
	Word_WordEnum_value = map[string]int32{
		"word_unspecified": 0,
		"app":              1,
		"application":      2,
		"aws":              3,
		"azure":            4,
		"bill":             5,
		"company":          6,
		"config":           7,
		"cpu":              8,
		"debug":            9,
		"docker":           10,
		"domain":           11,
		"email":            12,
		"endpoint":         13,
		"external":         14,
		"file":             15,
		"gcp":              16,
		"google":           17,
		"hosting":          18,
		"hostname":         19,
		"id":               20,
		"ingress":          21,
		"instances":        22,
		"internal":         23,
		"ip":               24,
		"key":              25,
		"kubeconfig":       26,
		"kubernetes":       27,
		"localhost":        28,
		"machine":          29,
		"main":             30,
		"maven":            31,
		"memory":           32,
		"microservice":     33,
		"name":             34,
		"nameservers":      35,
		"namespace":        36,
		"network":          37,
		"npm":              38,
		"number":           39,
		"org":              40,
		"parent":           41,
		"postgres":         42,
		"product":          43,
		"project":          44,
		"python":           45,
		"repo":             46,
		"review":           47,
		"server":           48,
		"servers":          49,
		"share":            50,
		"shared_services":  51,
		"spilo":            52,
		"spot":             53,
		"stunnel":          54,
		"team":             55,
		"type":             56,
		"username":         57,
		"util":             58,
		"version":          59,
		"workload":         60,
		"url":              61,
		"env":              62,
		"kafka":            63,
		"unknown":          64,
		"resource":         65,
		"resource_id":      66,
		"resource_type":    67,
		"environment":      68,
	}
)

func (x Word_WordEnum) Enum() *Word_WordEnum {
	p := new(Word_WordEnum)
	*p = x
	return p
}

func (x Word_WordEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Word_WordEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_enumTypes[0].Descriptor()
}

func (Word_WordEnum) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_enumTypes[0]
}

func (x Word_WordEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Word_WordEnum.Descriptor instead.
func (Word_WordEnum) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDescGZIP(), []int{0, 0}
}

type Acronym_AcronymEnum int32

const (
	Acronym_ACRONYM_UNSPECIFIED Acronym_AcronymEnum = 0
	// network
	Acronym_NW Acronym_AcronymEnum = 1
	// rpc
	Acronym_RPC Acronym_AcronymEnum = 2
)

// Enum value maps for Acronym_AcronymEnum.
var (
	Acronym_AcronymEnum_name = map[int32]string{
		0: "ACRONYM_UNSPECIFIED",
		1: "NW",
		2: "RPC",
	}
	Acronym_AcronymEnum_value = map[string]int32{
		"ACRONYM_UNSPECIFIED": 0,
		"NW":                  1,
		"RPC":                 2,
	}
)

func (x Acronym_AcronymEnum) Enum() *Acronym_AcronymEnum {
	p := new(Acronym_AcronymEnum)
	*p = x
	return p
}

func (x Acronym_AcronymEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Acronym_AcronymEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_enumTypes[1].Descriptor()
}

func (Acronym_AcronymEnum) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_enumTypes[1]
}

func (x Acronym_AcronymEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Acronym_AcronymEnum.Descriptor instead.
func (Acronym_AcronymEnum) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDescGZIP(), []int{1, 0}
}

// english word enums added to avoid typing mistakes for the commonly used words for naming things
// this enum is encapsulated inside a message as a few entries like "name" (a reserved word in few languages) can only be added to the enum if it is inside a message.
// the recommended best practice to prefix the entry with enum name has been intentionally ignored to as the values of the entries are intended to be used to name resources.
type Word struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Word) Reset() {
	*x = Word{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Word) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Word) ProtoMessage() {}

func (x *Word) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Word.ProtoReflect.Descriptor instead.
func (*Word) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDescGZIP(), []int{0}
}

// english acronym enums added to avoid typing mistakes for the commonly used acronyms for naming things
// this enum is encapsulated inside a message as a few entries like "name" (a reserved word in few languages) can only be added to the enum if it is inside a message.
// the recommended best practice to prefix the entry with enum name has been intentionally ignored to as the values of the entries are intended to be used to name resources.
type Acronym struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Acronym) Reset() {
	*x = Acronym{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acronym) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acronym) ProtoMessage() {}

func (x *Acronym) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acronym.ProtoReflect.Descriptor instead.
func (*Acronym) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDescGZIP(), []int{1}
}

var File_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x67,
	0x6c, 0x69, 0x73, 0x68, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xf3,
	0x06, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x64, 0x22, 0xea, 0x06, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x64,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x75, 0x6e, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x61, 0x70,
	0x70, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x10, 0x03, 0x12, 0x09, 0x0a,
	0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x62, 0x69, 0x6c, 0x6c,
	0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x10, 0x06, 0x12,
	0x0a, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x10, 0x07, 0x12, 0x07, 0x0a, 0x03, 0x63,
	0x70, 0x75, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x10, 0x09, 0x12,
	0x0a, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x10, 0x0d,
	0x12, 0x0c, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x10, 0x0e, 0x12, 0x08,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x10, 0x0f, 0x12, 0x07, 0x0a, 0x03, 0x67, 0x63, 0x70, 0x10,
	0x10, 0x12, 0x0a, 0x0a, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x10, 0x11, 0x12, 0x0b, 0x0a,
	0x07, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x12, 0x12, 0x0c, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x10, 0x13, 0x12, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x14,
	0x12, 0x0b, 0x0a, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x15, 0x12, 0x0d, 0x0a,
	0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x10, 0x16, 0x12, 0x0c, 0x0a, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x10, 0x17, 0x12, 0x06, 0x0a, 0x02, 0x69, 0x70,
	0x10, 0x18, 0x12, 0x07, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x10, 0x19, 0x12, 0x0e, 0x0a, 0x0a, 0x6b,
	0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x10, 0x1a, 0x12, 0x0e, 0x0a, 0x0a, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x10, 0x1b, 0x12, 0x0d, 0x0a, 0x09, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x10, 0x1c, 0x12, 0x0b, 0x0a, 0x07, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x10, 0x1d, 0x12, 0x08, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x10,
	0x1e, 0x12, 0x09, 0x0a, 0x05, 0x6d, 0x61, 0x76, 0x65, 0x6e, 0x10, 0x1f, 0x12, 0x0a, 0x0a, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x10, 0x20, 0x12, 0x10, 0x0a, 0x0c, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0x21, 0x12, 0x08, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x10, 0x22, 0x12, 0x0f, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x10, 0x23, 0x12, 0x0d, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x10, 0x24, 0x12, 0x0b, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x10,
	0x25, 0x12, 0x07, 0x0a, 0x03, 0x6e, 0x70, 0x6d, 0x10, 0x26, 0x12, 0x0a, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x10, 0x27, 0x12, 0x07, 0x0a, 0x03, 0x6f, 0x72, 0x67, 0x10, 0x28, 0x12,
	0x0a, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x10, 0x29, 0x12, 0x0c, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x10, 0x2a, 0x12, 0x0b, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x10, 0x2b, 0x12, 0x0b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x10, 0x2c, 0x12, 0x0a, 0x0a, 0x06, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x10, 0x2d, 0x12,
	0x08, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x10, 0x2e, 0x12, 0x0a, 0x0a, 0x06, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x10, 0x2f, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10,
	0x30, 0x12, 0x0b, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x10, 0x31, 0x12, 0x09,
	0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x10, 0x32, 0x12, 0x13, 0x0a, 0x0f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x10, 0x33, 0x12, 0x09,
	0x0a, 0x05, 0x73, 0x70, 0x69, 0x6c, 0x6f, 0x10, 0x34, 0x12, 0x08, 0x0a, 0x04, 0x73, 0x70, 0x6f,
	0x74, 0x10, 0x35, 0x12, 0x0b, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x10, 0x36,
	0x12, 0x08, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x10, 0x37, 0x12, 0x08, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x10, 0x38, 0x12, 0x0c, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x10, 0x39, 0x12, 0x08, 0x0a, 0x04, 0x75, 0x74, 0x69, 0x6c, 0x10, 0x3a, 0x12, 0x0b, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x3b, 0x12, 0x0c, 0x0a, 0x08, 0x77, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x10, 0x3c, 0x12, 0x07, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x10, 0x3d,
	0x12, 0x07, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x10, 0x3e, 0x12, 0x09, 0x0a, 0x05, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x10, 0x3f, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x40, 0x12, 0x0c, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x10, 0x41, 0x12,
	0x0f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x10, 0x42,
	0x12, 0x11, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x10, 0x43, 0x12, 0x0f, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x10, 0x44, 0x22, 0x42, 0x0a, 0x07, 0x41, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x22,
	0x37, 0x0a, 0x0b, 0x41, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x17,
	0x0a, 0x13, 0x41, 0x43, 0x52, 0x4f, 0x4e, 0x59, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x57, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x10, 0x02, 0x42, 0x94, 0x03, 0x0a, 0x3d, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x0a, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x56,
	0x43, 0x45, 0x52, 0x45, 0xaa, 0x02, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x52, 0x70, 0x63,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x5c, 0x52,
	0x70, 0x63, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xe2, 0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68,
	0x5c, 0x52, 0x70, 0x63, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x36, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x45, 0x6e, 0x67, 0x6c,
	0x69, 0x73, 0x68, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDescData = file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_goTypes = []interface{}{
	(Word_WordEnum)(0),       // 0: cloud.planton.apis.v1.commons.english.rpc.enums.Word.WordEnum
	(Acronym_AcronymEnum)(0), // 1: cloud.planton.apis.v1.commons.english.rpc.enums.Acronym.AcronymEnum
	(*Word)(nil),             // 2: cloud.planton.apis.v1.commons.english.rpc.enums.Word
	(*Acronym)(nil),          // 3: cloud.planton.apis.v1.commons.english.rpc.enums.Acronym
}
var file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_init() }
func file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_init() {
	if File_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Word); i {
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
		file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Acronym); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_enumTypes,
		MessageInfos:      file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto = out.File
	file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_english_rpc_enums_enums_proto_depIdxs = nil
}