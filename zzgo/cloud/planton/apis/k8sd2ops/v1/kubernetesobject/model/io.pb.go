// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/k8sd2ops/v1/kubernetesobject/model/io.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	model1 "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/integration/v1/kubernetes/kubernetesobject/model"
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

// api-resource-kubernetes-object
type ApiResourceKubernetesObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// planton-cloud api-resource
	ApiResource *model.ApiResourceKindApiResourceId `protobuf:"bytes,1,opt,name=api_resource,json=apiResource,proto3" json:"api_resource,omitempty"`
	// kubernetes object
	KubernetesObject *model1.KubernetesObject `protobuf:"bytes,2,opt,name=kubernetes_object,json=kubernetesObject,proto3" json:"kubernetes_object,omitempty"`
}

func (x *ApiResourceKubernetesObject) Reset() {
	*x = ApiResourceKubernetesObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResourceKubernetesObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResourceKubernetesObject) ProtoMessage() {}

func (x *ApiResourceKubernetesObject) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResourceKubernetesObject.ProtoReflect.Descriptor instead.
func (*ApiResourceKubernetesObject) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *ApiResourceKubernetesObject) GetApiResource() *model.ApiResourceKindApiResourceId {
	if x != nil {
		return x.ApiResource
	}
	return nil
}

func (x *ApiResourceKubernetesObject) GetKubernetesObject() *model1.KubernetesObject {
	if x != nil {
		return x.KubernetesObject
	}
	return nil
}

// input for rpc to find all kubernetes objects in a namespace.
type StreamKubernetesObjectsInNamespaceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// planton-cloud api-resource
	ApiResource *model.ApiResourceKindApiResourceId `protobuf:"bytes,1,opt,name=api_resource,json=apiResource,proto3" json:"api_resource,omitempty"`
	// kubernetes namespace
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *StreamKubernetesObjectsInNamespaceInput) Reset() {
	*x = StreamKubernetesObjectsInNamespaceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamKubernetesObjectsInNamespaceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamKubernetesObjectsInNamespaceInput) ProtoMessage() {}

func (x *StreamKubernetesObjectsInNamespaceInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamKubernetesObjectsInNamespaceInput.ProtoReflect.Descriptor instead.
func (*StreamKubernetesObjectsInNamespaceInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *StreamKubernetesObjectsInNamespaceInput) GetApiResource() *model.ApiResourceKindApiResourceId {
	if x != nil {
		return x.ApiResource
	}
	return nil
}

func (x *StreamKubernetesObjectsInNamespaceInput) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// input for rpc to update a kubernetes object in a kube-cluster
type UpdateKubernetesObjectInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// planton-cloud api-resource
	ApiResource *model.ApiResourceKindApiResourceId `protobuf:"bytes,1,opt,name=api_resource,json=apiResource,proto3" json:"api_resource,omitempty"`
	// kubernetes object
	KubernetesObject *model1.KubernetesObject `protobuf:"bytes,2,opt,name=kubernetes_object,json=kubernetesObject,proto3" json:"kubernetes_object,omitempty"`
	// base64 encoded of the updated kubernetes object yaml
	UpdatedObjectYamlBase64 string `protobuf:"bytes,3,opt,name=updated_object_yaml_base64,json=updatedObjectYamlBase64,proto3" json:"updated_object_yaml_base64,omitempty"`
}

func (x *UpdateKubernetesObjectInput) Reset() {
	*x = UpdateKubernetesObjectInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateKubernetesObjectInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateKubernetesObjectInput) ProtoMessage() {}

func (x *UpdateKubernetesObjectInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateKubernetesObjectInput.ProtoReflect.Descriptor instead.
func (*UpdateKubernetesObjectInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateKubernetesObjectInput) GetApiResource() *model.ApiResourceKindApiResourceId {
	if x != nil {
		return x.ApiResource
	}
	return nil
}

func (x *UpdateKubernetesObjectInput) GetKubernetesObject() *model1.KubernetesObject {
	if x != nil {
		return x.KubernetesObject
	}
	return nil
}

func (x *UpdateKubernetesObjectInput) GetUpdatedObjectYamlBase64() string {
	if x != nil {
		return x.UpdatedObjectYamlBase64
	}
	return ""
}

// input for rpc to find list of pods for api-resource
type FindKubernetesPodsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// planton-cloud api-resource
	ApiResource *model.ApiResourceKindApiResourceId `protobuf:"bytes,1,opt,name=api_resource,json=apiResource,proto3" json:"api_resource,omitempty"`
	// kubernetes namespace
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// details of the parent kubernetes object(pod-manager) to be used to filter pods list
	// if a parent is not specified, pods of all pod-managers will be included in the response.
	PodManager *model1.KubernetesObject `protobuf:"bytes,3,opt,name=pod_manager,json=podManager,proto3" json:"pod_manager,omitempty"`
}

func (x *FindKubernetesPodsInput) Reset() {
	*x = FindKubernetesPodsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindKubernetesPodsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindKubernetesPodsInput) ProtoMessage() {}

func (x *FindKubernetesPodsInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindKubernetesPodsInput.ProtoReflect.Descriptor instead.
func (*FindKubernetesPodsInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDescGZIP(), []int{3}
}

func (x *FindKubernetesPodsInput) GetApiResource() *model.ApiResourceKindApiResourceId {
	if x != nil {
		return x.ApiResource
	}
	return nil
}

func (x *FindKubernetesPodsInput) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *FindKubernetesPodsInput) GetPodManager() *model1.KubernetesObject {
	if x != nil {
		return x.PodManager
	}
	return nil
}

var File_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4c, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x02, 0x0a, 0x1b, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x6d, 0x0a, 0x0c, 0x61, 0x70, 0x69,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x4a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x0b, 0x61, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x11, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x10, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0xb6, 0x01,
	0x0a, 0x27, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x6d, 0x0a, 0x0c, 0x61, 0x70, 0x69,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x4a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x0b, 0x61, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xce, 0x02, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x6d, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x0b, 0x61, 0x70, 0x69, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x11, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x55, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x10, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3b, 0x0a, 0x1a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x79, 0x61, 0x6d,
	0x6c, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x59, 0x61, 0x6d,
	0x6c, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x22, 0x9e, 0x02, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x6f, 0x64, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x6d, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x0b, 0x61, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x76, 0x0a, 0x0b, 0x70, 0x6f, 0x64, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0a, 0x70, 0x6f,
	0x64, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x42, 0xb3, 0x03, 0x0a, 0x43, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x38,
	0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x42, 0x07, 0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x65, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x4b, 0x56, 0x4b, 0x4d, 0xaa, 0x02, 0x35, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69,
	0x73, 0x2e, 0x4b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x4b, 0x38, 0x73, 0x64, 0x32, 0x6f,
	0x70, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x41, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69,
	0x73, 0x5c, 0x4b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x4b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDescData = file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDesc
)

func file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDescData
}

var file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_goTypes = []interface{}{
	(*ApiResourceKubernetesObject)(nil),             // 0: cloud.planton.apis.k8sd2ops.v1.kubernetesobject.model.ApiResourceKubernetesObject
	(*StreamKubernetesObjectsInNamespaceInput)(nil), // 1: cloud.planton.apis.k8sd2ops.v1.kubernetesobject.model.StreamKubernetesObjectsInNamespaceInput
	(*UpdateKubernetesObjectInput)(nil),             // 2: cloud.planton.apis.k8sd2ops.v1.kubernetesobject.model.UpdateKubernetesObjectInput
	(*FindKubernetesPodsInput)(nil),                 // 3: cloud.planton.apis.k8sd2ops.v1.kubernetesobject.model.FindKubernetesPodsInput
	(*model.ApiResourceKindApiResourceId)(nil),      // 4: cloud.planton.apis.commons.apiresource.model.ApiResourceKindApiResourceId
	(*model1.KubernetesObject)(nil),                 // 5: cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.model.KubernetesObject
}
var file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.k8sd2ops.v1.kubernetesobject.model.ApiResourceKubernetesObject.api_resource:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceKindApiResourceId
	5, // 1: cloud.planton.apis.k8sd2ops.v1.kubernetesobject.model.ApiResourceKubernetesObject.kubernetes_object:type_name -> cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.model.KubernetesObject
	4, // 2: cloud.planton.apis.k8sd2ops.v1.kubernetesobject.model.StreamKubernetesObjectsInNamespaceInput.api_resource:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceKindApiResourceId
	4, // 3: cloud.planton.apis.k8sd2ops.v1.kubernetesobject.model.UpdateKubernetesObjectInput.api_resource:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceKindApiResourceId
	5, // 4: cloud.planton.apis.k8sd2ops.v1.kubernetesobject.model.UpdateKubernetesObjectInput.kubernetes_object:type_name -> cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.model.KubernetesObject
	4, // 5: cloud.planton.apis.k8sd2ops.v1.kubernetesobject.model.FindKubernetesPodsInput.api_resource:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceKindApiResourceId
	5, // 6: cloud.planton.apis.k8sd2ops.v1.kubernetesobject.model.FindKubernetesPodsInput.pod_manager:type_name -> cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.model.KubernetesObject
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_init() }
func file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_init() {
	if File_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResourceKubernetesObject); i {
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
		file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamKubernetesObjectsInNamespaceInput); i {
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
		file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateKubernetesObjectInput); i {
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
		file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindKubernetesPodsInput); i {
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
			RawDescriptor: file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto = out.File
	file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_goTypes = nil
	file_cloud_planton_apis_k8sd2ops_v1_kubernetesobject_model_io_proto_depIdxs = nil
}