// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/iam/identity/connection/rpc/enums/enums.proto

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

type IdentityConnectionType int32

const (
	IdentityConnectionType_IDENTITY_CONNECTION_TYPE_UNSPECIFIED                   IdentityConnectionType = 0
	IdentityConnectionType_IDENTITY_CONNECTION_TYPE_SAML                          IdentityConnectionType = 1
	IdentityConnectionType_IDENTITY_CONNECTION_TYPE_LDAP                          IdentityConnectionType = 2
	IdentityConnectionType_IDENTITY_CONNECTION_TYPE_OPENID                        IdentityConnectionType = 3
	IdentityConnectionType_IDENTITY_CONNECTION_TYPE_ADFS                          IdentityConnectionType = 4
	IdentityConnectionType_IDENTITY_CONNECTION_TYPE_OKTA_WORKFORCE                IdentityConnectionType = 5
	IdentityConnectionType_IDENTITY_CONNECTION_TYPE_AZURE_ACTIVE_DIRECTORY_NATIVE IdentityConnectionType = 6
	IdentityConnectionType_IDENTITY_CONNECTION_TYPE_PING_FEDERATE                 IdentityConnectionType = 7
)

// Enum value maps for IdentityConnectionType.
var (
	IdentityConnectionType_name = map[int32]string{
		0: "IDENTITY_CONNECTION_TYPE_UNSPECIFIED",
		1: "IDENTITY_CONNECTION_TYPE_SAML",
		2: "IDENTITY_CONNECTION_TYPE_LDAP",
		3: "IDENTITY_CONNECTION_TYPE_OPENID",
		4: "IDENTITY_CONNECTION_TYPE_ADFS",
		5: "IDENTITY_CONNECTION_TYPE_OKTA_WORKFORCE",
		6: "IDENTITY_CONNECTION_TYPE_AZURE_ACTIVE_DIRECTORY_NATIVE",
		7: "IDENTITY_CONNECTION_TYPE_PING_FEDERATE",
	}
	IdentityConnectionType_value = map[string]int32{
		"IDENTITY_CONNECTION_TYPE_UNSPECIFIED":                   0,
		"IDENTITY_CONNECTION_TYPE_SAML":                          1,
		"IDENTITY_CONNECTION_TYPE_LDAP":                          2,
		"IDENTITY_CONNECTION_TYPE_OPENID":                        3,
		"IDENTITY_CONNECTION_TYPE_ADFS":                          4,
		"IDENTITY_CONNECTION_TYPE_OKTA_WORKFORCE":                5,
		"IDENTITY_CONNECTION_TYPE_AZURE_ACTIVE_DIRECTORY_NATIVE": 6,
		"IDENTITY_CONNECTION_TYPE_PING_FEDERATE":                 7,
	}
)

func (x IdentityConnectionType) Enum() *IdentityConnectionType {
	p := new(IdentityConnectionType)
	*p = x
	return p
}

func (x IdentityConnectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentityConnectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_enumTypes[0].Descriptor()
}

func (IdentityConnectionType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_enumTypes[0]
}

func (x IdentityConnectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentityConnectionType.Descriptor instead.
func (IdentityConnectionType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDescGZIP(), []int{0}
}

type SamlSignRequestAlgorithm int32

const (
	SamlSignRequestAlgorithm_SAML_SIGN_REQUEST_ALGORITHM_UNSPECIFIED SamlSignRequestAlgorithm = 0
	SamlSignRequestAlgorithm_SAML_SIGN_REQUEST_ALGORITHM_RSA_SHA256  SamlSignRequestAlgorithm = 1
	SamlSignRequestAlgorithm_SAML_SIGN_REQUEST_ALGORITHM_RSA_SHA1    SamlSignRequestAlgorithm = 2
)

// Enum value maps for SamlSignRequestAlgorithm.
var (
	SamlSignRequestAlgorithm_name = map[int32]string{
		0: "SAML_SIGN_REQUEST_ALGORITHM_UNSPECIFIED",
		1: "SAML_SIGN_REQUEST_ALGORITHM_RSA_SHA256",
		2: "SAML_SIGN_REQUEST_ALGORITHM_RSA_SHA1",
	}
	SamlSignRequestAlgorithm_value = map[string]int32{
		"SAML_SIGN_REQUEST_ALGORITHM_UNSPECIFIED": 0,
		"SAML_SIGN_REQUEST_ALGORITHM_RSA_SHA256":  1,
		"SAML_SIGN_REQUEST_ALGORITHM_RSA_SHA1":    2,
	}
)

func (x SamlSignRequestAlgorithm) Enum() *SamlSignRequestAlgorithm {
	p := new(SamlSignRequestAlgorithm)
	*p = x
	return p
}

func (x SamlSignRequestAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SamlSignRequestAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_enumTypes[1].Descriptor()
}

func (SamlSignRequestAlgorithm) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_enumTypes[1]
}

func (x SamlSignRequestAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SamlSignRequestAlgorithm.Descriptor instead.
func (SamlSignRequestAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDescGZIP(), []int{1}
}

type SamlSignRequestAlgorithmDigest int32

const (
	SamlSignRequestAlgorithmDigest_SAML_SIGN_REQUEST_ALGORITHM_DIGEST_UNSPECIFIED SamlSignRequestAlgorithmDigest = 0
	SamlSignRequestAlgorithmDigest_SAML_SIGN_REQUEST_ALGORITHM_DIGEST_SHA256      SamlSignRequestAlgorithmDigest = 1
	SamlSignRequestAlgorithmDigest_SAML_SIGN_REQUEST_ALGORITHM_DIGEST_SHA1        SamlSignRequestAlgorithmDigest = 2
)

// Enum value maps for SamlSignRequestAlgorithmDigest.
var (
	SamlSignRequestAlgorithmDigest_name = map[int32]string{
		0: "SAML_SIGN_REQUEST_ALGORITHM_DIGEST_UNSPECIFIED",
		1: "SAML_SIGN_REQUEST_ALGORITHM_DIGEST_SHA256",
		2: "SAML_SIGN_REQUEST_ALGORITHM_DIGEST_SHA1",
	}
	SamlSignRequestAlgorithmDigest_value = map[string]int32{
		"SAML_SIGN_REQUEST_ALGORITHM_DIGEST_UNSPECIFIED": 0,
		"SAML_SIGN_REQUEST_ALGORITHM_DIGEST_SHA256":      1,
		"SAML_SIGN_REQUEST_ALGORITHM_DIGEST_SHA1":        2,
	}
)

func (x SamlSignRequestAlgorithmDigest) Enum() *SamlSignRequestAlgorithmDigest {
	p := new(SamlSignRequestAlgorithmDigest)
	*p = x
	return p
}

func (x SamlSignRequestAlgorithmDigest) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SamlSignRequestAlgorithmDigest) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_enumTypes[2].Descriptor()
}

func (SamlSignRequestAlgorithmDigest) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_enumTypes[2]
}

func (x SamlSignRequestAlgorithmDigest) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SamlSignRequestAlgorithmDigest.Descriptor instead.
func (SamlSignRequestAlgorithmDigest) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDescGZIP(), []int{2}
}

type SamlProtocolBinding int32

const (
	SamlProtocolBinding_SAML_PROTOCOL_BINDING_UNSPECIFIED   SamlProtocolBinding = 0
	SamlProtocolBinding_SAML_PROTOCOL_BINDING_HTTP_REDIRECT SamlProtocolBinding = 1
	SamlProtocolBinding_SAML_PROTOCOL_BINDING_HTTP_POST     SamlProtocolBinding = 2
)

// Enum value maps for SamlProtocolBinding.
var (
	SamlProtocolBinding_name = map[int32]string{
		0: "SAML_PROTOCOL_BINDING_UNSPECIFIED",
		1: "SAML_PROTOCOL_BINDING_HTTP_REDIRECT",
		2: "SAML_PROTOCOL_BINDING_HTTP_POST",
	}
	SamlProtocolBinding_value = map[string]int32{
		"SAML_PROTOCOL_BINDING_UNSPECIFIED":   0,
		"SAML_PROTOCOL_BINDING_HTTP_REDIRECT": 1,
		"SAML_PROTOCOL_BINDING_HTTP_POST":     2,
	}
)

func (x SamlProtocolBinding) Enum() *SamlProtocolBinding {
	p := new(SamlProtocolBinding)
	*p = x
	return p
}

func (x SamlProtocolBinding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SamlProtocolBinding) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_enumTypes[3].Descriptor()
}

func (SamlProtocolBinding) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_enumTypes[3]
}

func (x SamlProtocolBinding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SamlProtocolBinding.Descriptor instead.
func (SamlProtocolBinding) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDescGZIP(), []int{3}
}

type MicrosoftIdentityApi int32

const (
	MicrosoftIdentityApi_MICROSOFT_IDENTITY_API_UNSPECIFIED MicrosoftIdentityApi = 0
	MicrosoftIdentityApi_MICROSOFT_IDENTITY_API_V1          MicrosoftIdentityApi = 1
	MicrosoftIdentityApi_MICROSOFT_IDENTITY_API_V2          MicrosoftIdentityApi = 2
)

// Enum value maps for MicrosoftIdentityApi.
var (
	MicrosoftIdentityApi_name = map[int32]string{
		0: "MICROSOFT_IDENTITY_API_UNSPECIFIED",
		1: "MICROSOFT_IDENTITY_API_V1",
		2: "MICROSOFT_IDENTITY_API_V2",
	}
	MicrosoftIdentityApi_value = map[string]int32{
		"MICROSOFT_IDENTITY_API_UNSPECIFIED": 0,
		"MICROSOFT_IDENTITY_API_V1":          1,
		"MICROSOFT_IDENTITY_API_V2":          2,
	}
)

func (x MicrosoftIdentityApi) Enum() *MicrosoftIdentityApi {
	p := new(MicrosoftIdentityApi)
	*p = x
	return p
}

func (x MicrosoftIdentityApi) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MicrosoftIdentityApi) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_enumTypes[4].Descriptor()
}

func (MicrosoftIdentityApi) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_enumTypes[4]
}

func (x MicrosoftIdentityApi) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MicrosoftIdentityApi.Descriptor instead.
func (MicrosoftIdentityApi) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDescGZIP(), []int{4}
}

var File_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDesc = []byte{
	0x0a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0xe5,
	0x02, 0x0a, 0x16, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x24, 0x49, 0x44, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x41, 0x4d, 0x4c, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49,
	0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4c, 0x44, 0x41, 0x50, 0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x49, 0x44, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x49, 0x44, 0x10, 0x03, 0x12, 0x21,
	0x0a, 0x1d, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x46, 0x53, 0x10,
	0x04, 0x12, 0x2b, 0x0a, 0x27, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x4b,
	0x54, 0x41, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x10, 0x05, 0x12, 0x3a,
	0x0a, 0x36, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x5a, 0x55, 0x52, 0x45,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x4f, 0x52,
	0x59, 0x5f, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x06, 0x12, 0x2a, 0x0a, 0x26, 0x49, 0x44,
	0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x45, 0x44, 0x45,
	0x52, 0x41, 0x54, 0x45, 0x10, 0x07, 0x2a, 0x9d, 0x01, 0x0a, 0x18, 0x53, 0x61, 0x6d, 0x6c, 0x53,
	0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x12, 0x2b, 0x0a, 0x27, 0x53, 0x41, 0x4d, 0x4c, 0x5f, 0x53, 0x49, 0x47, 0x4e,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54,
	0x48, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x2a, 0x0a, 0x26, 0x53, 0x41, 0x4d, 0x4c, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f,
	0x52, 0x53, 0x41, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x01, 0x12, 0x28, 0x0a, 0x24,
	0x53, 0x41, 0x4d, 0x4c, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x52, 0x53, 0x41, 0x5f,
	0x53, 0x48, 0x41, 0x31, 0x10, 0x02, 0x2a, 0xb0, 0x01, 0x0a, 0x1e, 0x53, 0x61, 0x6d, 0x6c, 0x53,
	0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x2e, 0x53, 0x41, 0x4d,
	0x4c, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x41,
	0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x44, 0x49, 0x47, 0x45, 0x53, 0x54, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2d, 0x0a,
	0x29, 0x53, 0x41, 0x4d, 0x4c, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x44, 0x49, 0x47,
	0x45, 0x53, 0x54, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x01, 0x12, 0x2b, 0x0a, 0x27,
	0x53, 0x41, 0x4d, 0x4c, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x44, 0x49, 0x47, 0x45,
	0x53, 0x54, 0x5f, 0x53, 0x48, 0x41, 0x31, 0x10, 0x02, 0x2a, 0x8a, 0x01, 0x0a, 0x13, 0x53, 0x61,
	0x6d, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x41, 0x4d, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43,
	0x4f, 0x4c, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x27, 0x0a, 0x23, 0x53, 0x41, 0x4d, 0x4c,
	0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10,
	0x01, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x41, 0x4d, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43,
	0x4f, 0x4c, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x5f,
	0x50, 0x4f, 0x53, 0x54, 0x10, 0x02, 0x2a, 0x7c, 0x0a, 0x14, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x6f, 0x66, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x70, 0x69, 0x12, 0x26,
	0x0a, 0x22, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x53, 0x4f, 0x46, 0x54, 0x5f, 0x49, 0x44, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x5f, 0x41, 0x50, 0x49, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x53,
	0x4f, 0x46, 0x54, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x41, 0x50, 0x49,
	0x5f, 0x56, 0x31, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x53, 0x4f,
	0x46, 0x54, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x41, 0x50, 0x49, 0x5f,
	0x56, 0x32, 0x10, 0x02, 0x42, 0xbc, 0x03, 0x0a, 0x3b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x42, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a,
	0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x09, 0x43, 0x50,
	0x41, 0x56, 0x49, 0x49, 0x43, 0x52, 0x45, 0xaa, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e,
	0x49, 0x61, 0x6d, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x70, 0x63, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xca, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xe2, 0x02, 0x43, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x52, 0x70, 0x63, 0x5c,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x3f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x49, 0x61,
	0x6d, 0x3a, 0x3a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDescData = file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDesc
)

func file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDescData
}

var file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_goTypes = []interface{}{
	(IdentityConnectionType)(0),         // 0: cloud.planton.apis.v1.iam.identity.connection.rpc.enums.IdentityConnectionType
	(SamlSignRequestAlgorithm)(0),       // 1: cloud.planton.apis.v1.iam.identity.connection.rpc.enums.SamlSignRequestAlgorithm
	(SamlSignRequestAlgorithmDigest)(0), // 2: cloud.planton.apis.v1.iam.identity.connection.rpc.enums.SamlSignRequestAlgorithmDigest
	(SamlProtocolBinding)(0),            // 3: cloud.planton.apis.v1.iam.identity.connection.rpc.enums.SamlProtocolBinding
	(MicrosoftIdentityApi)(0),           // 4: cloud.planton.apis.v1.iam.identity.connection.rpc.enums.MicrosoftIdentityApi
}
var file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_init() }
func file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_init() {
	if File_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto = out.File
	file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_rawDesc = nil
	file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_goTypes = nil
	file_cloud_planton_apis_v1_iam_identity_connection_rpc_enums_enums_proto_depIdxs = nil
}
