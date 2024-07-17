// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/commons/kubernetes/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	kubernetesworkloadingresstype "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/kubernetes/enums/kubernetesworkloadingresstype"
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

// sidecar object in microservice deployment configuration
// this spec should always match the specification of a kubernetes container spec https://pkg.go.dev/k8s.io/api/core/v1#Container
// warning: sidecar feature does not currently support all features of a kubernetes container spec.
type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the container
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// container image
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// container ports
	Ports []*ContainerPort `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	// container resources
	Resources *ContainerResources `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
	// container environment variables
	Env []*ContainerEnvVar `protobuf:"bytes,5,rep,name=env,proto3" json:"env,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *Container) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Container) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Container) GetPorts() []*ContainerPort {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *Container) GetResources() *ContainerResources {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Container) GetEnv() []*ContainerEnvVar {
	if x != nil {
		return x.Env
	}
	return nil
}

// container cpu and memory resources
type ContainerResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// container resource limits
	// key is either cpu or memory and value is the limits value for the resource
	Limits *CpuMemory `protobuf:"bytes,1,opt,name=limits,proto3" json:"limits,omitempty"`
	// container resource limits
	// key is either cpu or memory and value is the requests value for the resource
	Requests *CpuMemory `protobuf:"bytes,2,opt,name=requests,proto3" json:"requests,omitempty"`
}

func (x *ContainerResources) Reset() {
	*x = ContainerResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerResources) ProtoMessage() {}

func (x *ContainerResources) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerResources.ProtoReflect.Descriptor instead.
func (*ContainerResources) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *ContainerResources) GetLimits() *CpuMemory {
	if x != nil {
		return x.Limits
	}
	return nil
}

func (x *ContainerResources) GetRequests() *CpuMemory {
	if x != nil {
		return x.Requests
	}
	return nil
}

// container env var
type ContainerEnvVar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the environment variable
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// value of the environment variable
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ContainerEnvVar) Reset() {
	*x = ContainerEnvVar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerEnvVar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerEnvVar) ProtoMessage() {}

func (x *ContainerEnvVar) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerEnvVar.ProtoReflect.Descriptor instead.
func (*ContainerEnvVar) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *ContainerEnvVar) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContainerEnvVar) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// container port
type ContainerPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// port name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// container port
	// the attribute names must use camel case to marshal into https://pkg.go.dev/k8s.io/api/core/v1#Container
	ContainerPort int32 `protobuf:"varint,2,opt,name=containerPort,proto3" json:"containerPort,omitempty"`
	// container port protocol
	Protocol string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *ContainerPort) Reset() {
	*x = ContainerPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerPort) ProtoMessage() {}

func (x *ContainerPort) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerPort.ProtoReflect.Descriptor instead.
func (*ContainerPort) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescGZIP(), []int{3}
}

func (x *ContainerPort) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContainerPort) GetContainerPort() int32 {
	if x != nil {
		return x.ContainerPort
	}
	return 0
}

func (x *ContainerPort) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

// container cpu and memory
type CpuMemory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// container cpu
	Cpu string `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	// container memory
	Memory string `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
}

func (x *CpuMemory) Reset() {
	*x = CpuMemory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuMemory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuMemory) ProtoMessage() {}

func (x *CpuMemory) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuMemory.ProtoReflect.Descriptor instead.
func (*CpuMemory) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescGZIP(), []int{4}
}

func (x *CpuMemory) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *CpuMemory) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

// container image
type ContainerImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// image repository
	Repo string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	// image tag
	Tag string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	// image pull secret name
	PullSecretName string `protobuf:"bytes,3,opt,name=pull_secret_name,json=pullSecretName,proto3" json:"pull_secret_name,omitempty"`
}

func (x *ContainerImage) Reset() {
	*x = ContainerImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerImage) ProtoMessage() {}

func (x *ContainerImage) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerImage.ProtoReflect.Descriptor instead.
func (*ContainerImage) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescGZIP(), []int{5}
}

func (x *ContainerImage) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *ContainerImage) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *ContainerImage) GetPullSecretName() string {
	if x != nil {
		return x.PullSecretName
	}
	return ""
}

// ingress spec of planton cloud api resource to be deployed in kubernetes
type IngressSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// toggle to control ingress
	IsEnabled bool `protobuf:"varint,1,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled,omitempty"`
	// ingress type
	IngressType kubernetesworkloadingresstype.KubernetesWorkloadIngressType `protobuf:"varint,2,opt,name=ingress_type,json=ingressType,proto3,enum=cloud.planton.apis.commons.kubernetes.enums.kubernetesworkloadingresstype.KubernetesWorkloadIngressType" json:"ingress_type,omitempty"`
	// endpoint domain to be used for creating internal and external endpoints for planton cloud api resource.
	EndpointDomainName string `protobuf:"bytes,3,opt,name=endpoint_domain_name,json=endpointDomainName,proto3" json:"endpoint_domain_name,omitempty"`
}

func (x *IngressSpec) Reset() {
	*x = IngressSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngressSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngressSpec) ProtoMessage() {}

func (x *IngressSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngressSpec.ProtoReflect.Descriptor instead.
func (*IngressSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescGZIP(), []int{6}
}

func (x *IngressSpec) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

func (x *IngressSpec) GetIngressType() kubernetesworkloadingresstype.KubernetesWorkloadIngressType {
	if x != nil {
		return x.IngressType
	}
	return kubernetesworkloadingresstype.KubernetesWorkloadIngressType(0)
}

func (x *IngressSpec) GetEndpointDomainName() string {
	if x != nil {
		return x.EndpointDomainName
	}
	return ""
}

var File_cloud_planton_apis_commons_kubernetes_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x77, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x50, 0x0a,
	0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12,
	0x5d, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x4e,
	0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x22, 0xb8,
	0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x43, 0x70, 0x75, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x52, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x70, 0x75, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52,
	0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x3b, 0x0a, 0x0f, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x65, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x35, 0x0a,
	0x09, 0x43, 0x70, 0x75, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70,
	0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x22, 0x60, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x28, 0x0a, 0x10,
	0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x75, 0x6c, 0x6c, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa1, 0x04, 0x0a, 0x0b, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x8b, 0x01, 0x0a, 0x0c, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x68, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0xb2, 0x02, 0xba, 0x48, 0xae, 0x02, 0x1a, 0xa3, 0x01, 0x0a,
	0x2d, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x2e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x72,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x20,
	0x26, 0x26, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x29, 0x20, 0x3d, 0x3d, 0x20, 0x30, 0x3f, 0x20, 0x27, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x20, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x69,
	0x73, 0x20, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x20, 0x74, 0x6f, 0x20, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x27, 0x3a, 0x20,
	0x27, 0x27, 0x1a, 0x85, 0x01, 0x0a, 0x25, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x2e, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x5c, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x20, 0x26, 0x26,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x20, 0x3d, 0x3d, 0x20, 0x30, 0x3f, 0x20, 0x27, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x69, 0x73, 0x20, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x79, 0x20, 0x74, 0x6f, 0x20, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x27, 0x3a, 0x20, 0x27, 0x27, 0x42, 0xf8, 0x02, 0x0a, 0x39, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0xa2, 0x02, 0x06, 0x43, 0x50, 0x41, 0x43, 0x4b, 0x4d, 0xaa, 0x02, 0x2b, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x2b, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x30, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescData = file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDesc
)

func file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDescData
}

var file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cloud_planton_apis_commons_kubernetes_model_state_proto_goTypes = []interface{}{
	(*Container)(nil),          // 0: cloud.planton.apis.commons.kubernetes.model.Container
	(*ContainerResources)(nil), // 1: cloud.planton.apis.commons.kubernetes.model.ContainerResources
	(*ContainerEnvVar)(nil),    // 2: cloud.planton.apis.commons.kubernetes.model.ContainerEnvVar
	(*ContainerPort)(nil),      // 3: cloud.planton.apis.commons.kubernetes.model.ContainerPort
	(*CpuMemory)(nil),          // 4: cloud.planton.apis.commons.kubernetes.model.CpuMemory
	(*ContainerImage)(nil),     // 5: cloud.planton.apis.commons.kubernetes.model.ContainerImage
	(*IngressSpec)(nil),        // 6: cloud.planton.apis.commons.kubernetes.model.IngressSpec
	(kubernetesworkloadingresstype.KubernetesWorkloadIngressType)(0), // 7: cloud.planton.apis.commons.kubernetes.enums.kubernetesworkloadingresstype.KubernetesWorkloadIngressType
}
var file_cloud_planton_apis_commons_kubernetes_model_state_proto_depIdxs = []int32{
	3, // 0: cloud.planton.apis.commons.kubernetes.model.Container.ports:type_name -> cloud.planton.apis.commons.kubernetes.model.ContainerPort
	1, // 1: cloud.planton.apis.commons.kubernetes.model.Container.resources:type_name -> cloud.planton.apis.commons.kubernetes.model.ContainerResources
	2, // 2: cloud.planton.apis.commons.kubernetes.model.Container.env:type_name -> cloud.planton.apis.commons.kubernetes.model.ContainerEnvVar
	4, // 3: cloud.planton.apis.commons.kubernetes.model.ContainerResources.limits:type_name -> cloud.planton.apis.commons.kubernetes.model.CpuMemory
	4, // 4: cloud.planton.apis.commons.kubernetes.model.ContainerResources.requests:type_name -> cloud.planton.apis.commons.kubernetes.model.CpuMemory
	7, // 5: cloud.planton.apis.commons.kubernetes.model.IngressSpec.ingress_type:type_name -> cloud.planton.apis.commons.kubernetes.enums.kubernetesworkloadingresstype.KubernetesWorkloadIngressType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_commons_kubernetes_model_state_proto_init() }
func file_cloud_planton_apis_commons_kubernetes_model_state_proto_init() {
	if File_cloud_planton_apis_commons_kubernetes_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container); i {
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
		file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerResources); i {
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
		file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerEnvVar); i {
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
		file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerPort); i {
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
		file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuMemory); i {
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
		file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerImage); i {
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
		file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngressSpec); i {
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
			RawDescriptor: file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_commons_kubernetes_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_commons_kubernetes_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_commons_kubernetes_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_commons_kubernetes_model_state_proto = out.File
	file_cloud_planton_apis_commons_kubernetes_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_commons_kubernetes_model_state_proto_goTypes = nil
	file_cloud_planton_apis_commons_kubernetes_model_state_proto_depIdxs = nil
}