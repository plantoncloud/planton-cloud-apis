// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/integration/kubernetes/resource/certificate.proto

package resource

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

// kubernetes certificate
type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// namespace of the certificate
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// name of the certificate
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// certificate labels
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// primary domain in the certificate
	PrimaryDomain string `protobuf:"bytes,4,opt,name=primary_domain,json=primaryDomain,proto3" json:"primary_domain,omitempty"`
	// additional domain names
	AdditionalDomains []string `protobuf:"bytes,5,rep,name=additional_domains,json=additionalDomains,proto3" json:"additional_domains,omitempty"`
	// certificate issuer reference in the format clusterissuer/issuer-name or issuer/namespace/name
	IssuerRef string `protobuf:"bytes,6,opt,name=issuer_ref,json=issuerRef,proto3" json:"issuer_ref,omitempty"`
	// name of the kubernetes secret in which certificate is stored
	SecretName string `protobuf:"bytes,7,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	// status of the certificate
	IsReady bool `protobuf:"varint,8,opt,name=is_ready,json=isReady,proto3" json:"is_ready,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *Certificate) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Certificate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Certificate) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Certificate) GetPrimaryDomain() string {
	if x != nil {
		return x.PrimaryDomain
	}
	return ""
}

func (x *Certificate) GetAdditionalDomains() []string {
	if x != nil {
		return x.AdditionalDomains
	}
	return nil
}

func (x *Certificate) GetIssuerRef() string {
	if x != nil {
		return x.IssuerRef
	}
	return ""
}

func (x *Certificate) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

func (x *Certificate) GetIsReady() bool {
	if x != nil {
		return x.IsReady
	}
	return false
}

// list of certificates
type Certificates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*Certificate `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *Certificates) Reset() {
	*x = Certificates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificates) ProtoMessage() {}

func (x *Certificates) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificates.ProtoReflect.Descriptor instead.
func (*Certificates) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDescGZIP(), []int{1}
}

func (x *Certificates) GetEntries() []*Certificate {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDesc = []byte{
	0x0a, 0x47, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0x93, 0x03, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x66, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x61,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x52, 0x65, 0x66, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6c, 0x0a, 0x0c, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x5c, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x42, 0xbc, 0x03, 0x0a, 0x43, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x10, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x56, 0x49, 0x4b,
	0x52, 0xaa, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xca, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56,
	0x31, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0xe2, 0x02, 0x41, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDescData = file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDesc
)

func file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDescData
}

var file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_goTypes = []interface{}{
	(*Certificate)(nil),  // 0: cloud.planton.apis.v1.integration.kubernetes.resource.Certificate
	(*Certificates)(nil), // 1: cloud.planton.apis.v1.integration.kubernetes.resource.Certificates
	nil,                  // 2: cloud.planton.apis.v1.integration.kubernetes.resource.Certificate.LabelsEntry
}
var file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_depIdxs = []int32{
	2, // 0: cloud.planton.apis.v1.integration.kubernetes.resource.Certificate.labels:type_name -> cloud.planton.apis.v1.integration.kubernetes.resource.Certificate.LabelsEntry
	0, // 1: cloud.planton.apis.v1.integration.kubernetes.resource.Certificates.entries:type_name -> cloud.planton.apis.v1.integration.kubernetes.resource.Certificate
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_init() }
func file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_init() {
	if File_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
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
		file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificates); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto = out.File
	file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_rawDesc = nil
	file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_goTypes = nil
	file_cloud_planton_apis_v1_integration_kubernetes_resource_certificate_proto_depIdxs = nil
}
