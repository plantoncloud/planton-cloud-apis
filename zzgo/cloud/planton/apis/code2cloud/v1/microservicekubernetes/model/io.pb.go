// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/microservicekubernetes/model/io.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// wrapper for id field of microservice-kubernetes
type MicroserviceKubernetesId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MicroserviceKubernetesId) Reset() {
	*x = MicroserviceKubernetesId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MicroserviceKubernetesId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MicroserviceKubernetesId) ProtoMessage() {}

func (x *MicroserviceKubernetesId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MicroserviceKubernetesId.ProtoReflect.Descriptor instead.
func (*MicroserviceKubernetesId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *MicroserviceKubernetesId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// list of microservice-kubernetes
type MicroserviceKubernetesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*MicroserviceKubernetes `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *MicroserviceKubernetesList) Reset() {
	*x = MicroserviceKubernetesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MicroserviceKubernetesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MicroserviceKubernetesList) ProtoMessage() {}

func (x *MicroserviceKubernetesList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MicroserviceKubernetesList.ProtoReflect.Descriptor instead.
func (*MicroserviceKubernetesList) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *MicroserviceKubernetesList) GetEntries() []*MicroserviceKubernetes {
	if x != nil {
		return x.Entries
	}
	return nil
}

// response for paginated query to list microservice-kubernetes
type MicroserviceKubernetesPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPages int32                     `protobuf:"varint,1,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	Entries    []*MicroserviceKubernetes `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *MicroserviceKubernetesPage) Reset() {
	*x = MicroserviceKubernetesPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MicroserviceKubernetesPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MicroserviceKubernetesPage) ProtoMessage() {}

func (x *MicroserviceKubernetesPage) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MicroserviceKubernetesPage.ProtoReflect.Descriptor instead.
func (*MicroserviceKubernetesPage) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescGZIP(), []int{2}
}

func (x *MicroserviceKubernetesPage) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *MicroserviceKubernetesPage) GetEntries() []*MicroserviceKubernetes {
	if x != nil {
		return x.Entries
	}
	return nil
}

// input for query to get microservice-kubernetes log stream
type GetMicroserviceKubernetesLogStreamQueryInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the microservice-kubernetes
	MicroserviceKubernetesId string `protobuf:"bytes,1,opt,name=microservice_kubernetes_id,json=microserviceKubernetesId,proto3" json:"microservice_kubernetes_id,omitempty"`
	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/
	// the number of lines from the end of the logs to show.
	// if the value is not set or is set to 0, then a default value of 600 seconds is used.
	TailLines int32 `protobuf:"varint,2,opt,name=tail_lines,json=tailLines,proto3" json:"tail_lines,omitempty"`
	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/
	// a relative time in seconds before the current time from which to show logs.
	// If this value precedes the time a pod was started, only logs since the pod start will be returned.
	// If this value is in the future, no logs will be returned.
	// if the value is not set or is set to 0, then a default value of 60 seconds is used.
	SinceSeconds int32 `protobuf:"varint,3,opt,name=since_seconds,json=sinceSeconds,proto3" json:"since_seconds,omitempty"`
}

func (x *GetMicroserviceKubernetesLogStreamQueryInput) Reset() {
	*x = GetMicroserviceKubernetesLogStreamQueryInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMicroserviceKubernetesLogStreamQueryInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMicroserviceKubernetesLogStreamQueryInput) ProtoMessage() {}

func (x *GetMicroserviceKubernetesLogStreamQueryInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMicroserviceKubernetesLogStreamQueryInput.ProtoReflect.Descriptor instead.
func (*GetMicroserviceKubernetesLogStreamQueryInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescGZIP(), []int{3}
}

func (x *GetMicroserviceKubernetesLogStreamQueryInput) GetMicroserviceKubernetesId() string {
	if x != nil {
		return x.MicroserviceKubernetesId
	}
	return ""
}

func (x *GetMicroserviceKubernetesLogStreamQueryInput) GetTailLines() int32 {
	if x != nil {
		return x.TailLines
	}
	return 0
}

func (x *GetMicroserviceKubernetesLogStreamQueryInput) GetSinceSeconds() int32 {
	if x != nil {
		return x.SinceSeconds
	}
	return 0
}

// input for rpc queries that take environment-id and code-project-id as input
type ByEnvironmentIdByCodeProjectIdInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the environment
	EnvironmentId string `protobuf:"bytes,1,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	// id of the code project
	CodeProjectId string `protobuf:"bytes,2,opt,name=code_project_id,json=codeProjectId,proto3" json:"code_project_id,omitempty"`
}

func (x *ByEnvironmentIdByCodeProjectIdInput) Reset() {
	*x = ByEnvironmentIdByCodeProjectIdInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByEnvironmentIdByCodeProjectIdInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByEnvironmentIdByCodeProjectIdInput) ProtoMessage() {}

func (x *ByEnvironmentIdByCodeProjectIdInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByEnvironmentIdByCodeProjectIdInput.ProtoReflect.Descriptor instead.
func (*ByEnvironmentIdByCodeProjectIdInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescGZIP(), []int{4}
}

func (x *ByEnvironmentIdByCodeProjectIdInput) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *ByEnvironmentIdByCodeProjectIdInput) GetCodeProjectId() string {
	if x != nil {
		return x.CodeProjectId
	}
	return ""
}

// wrapper for microservice-kubernetes env var map
type MicroserviceKubernetesEnvVarMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MicroserviceKubernetesEnvVarMap) Reset() {
	*x = MicroserviceKubernetesEnvVarMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MicroserviceKubernetesEnvVarMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MicroserviceKubernetesEnvVarMap) ProtoMessage() {}

func (x *MicroserviceKubernetesEnvVarMap) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MicroserviceKubernetesEnvVarMap.ProtoReflect.Descriptor instead.
func (*MicroserviceKubernetesEnvVarMap) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescGZIP(), []int{5}
}

func (x *MicroserviceKubernetesEnvVarMap) GetValue() map[string]string {
	if x != nil {
		return x.Value
	}
	return nil
}

// input for rpc to get env-var map which is typically required to build .env files
type GetMicroserviceKubernetesEnvVarMapInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the microservice kubernetes
	MicroserviceKubernetesId string `protobuf:"bytes,1,opt,name=microservice_kubernetes_id,json=microserviceKubernetesId,proto3" json:"microservice_kubernetes_id,omitempty"`
	// map of environment variables
	Variables map[string]string `protobuf:"bytes,2,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// map of environment secrets
	Secrets map[string]string `protobuf:"bytes,3,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetMicroserviceKubernetesEnvVarMapInput) Reset() {
	*x = GetMicroserviceKubernetesEnvVarMapInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMicroserviceKubernetesEnvVarMapInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMicroserviceKubernetesEnvVarMapInput) ProtoMessage() {}

func (x *GetMicroserviceKubernetesEnvVarMapInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMicroserviceKubernetesEnvVarMapInput.ProtoReflect.Descriptor instead.
func (*GetMicroserviceKubernetesEnvVarMapInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescGZIP(), []int{6}
}

func (x *GetMicroserviceKubernetesEnvVarMapInput) GetMicroserviceKubernetesId() string {
	if x != nil {
		return x.MicroserviceKubernetesId
	}
	return ""
}

func (x *GetMicroserviceKubernetesEnvVarMapInput) GetVariables() map[string]string {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *GetMicroserviceKubernetesEnvVarMapInput) GetSecrets() map[string]string {
	if x != nil {
		return x.Secrets
	}
	return nil
}

var File_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDesc = []byte{
	0x0a, 0x46, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x49, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x38, 0x0a, 0x18, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x1a, 0x4d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x6f, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x1a, 0x4d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x6f, 0x0a, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x2c, 0x47,
	0x65, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x44, 0x0a, 0x1a, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x18, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0a, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x74,
	0x61, 0x69, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0d, 0x73, 0x69, 0x6e, 0x63,
	0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0c, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x23, 0x42, 0x79, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2d, 0x0a,
	0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0d, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0f,
	0x63, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0d, 0x63,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0xdc, 0x01, 0x0a,
	0x1f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x4d, 0x61, 0x70,
	0x12, 0x7f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x4d, 0x61, 0x70, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8f, 0x04, 0x0a, 0x27,
	0x47, 0x65, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x4d,
	0x61, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x44, 0x0a, 0x1a, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x18, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x49, 0x64, 0x12, 0x93, 0x01,
	0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x75, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x76, 0x56, 0x61,
	0x72, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x76, 0x56, 0x61, 0x72, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xe3, 0x03,
	0x0a, 0x4b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x49,
	0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x43, 0x56, 0x4d,
	0x4d, 0xaa, 0x02, 0x3d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x56, 0x31, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0xca, 0x02, 0x3d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x56, 0x31, 0x5c, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0xe2, 0x02, 0x49, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x56, 0x31, 0x5c, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x43,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a,
	0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_goTypes = []interface{}{
	(*MicroserviceKubernetesId)(nil),                     // 0: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.MicroserviceKubernetesId
	(*MicroserviceKubernetesList)(nil),                   // 1: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.MicroserviceKubernetesList
	(*MicroserviceKubernetesPage)(nil),                   // 2: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.MicroserviceKubernetesPage
	(*GetMicroserviceKubernetesLogStreamQueryInput)(nil), // 3: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.GetMicroserviceKubernetesLogStreamQueryInput
	(*ByEnvironmentIdByCodeProjectIdInput)(nil),          // 4: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.ByEnvironmentIdByCodeProjectIdInput
	(*MicroserviceKubernetesEnvVarMap)(nil),              // 5: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.MicroserviceKubernetesEnvVarMap
	(*GetMicroserviceKubernetesEnvVarMapInput)(nil),      // 6: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.GetMicroserviceKubernetesEnvVarMapInput
	nil,                            // 7: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.MicroserviceKubernetesEnvVarMap.ValueEntry
	nil,                            // 8: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.GetMicroserviceKubernetesEnvVarMapInput.VariablesEntry
	nil,                            // 9: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.GetMicroserviceKubernetesEnvVarMapInput.SecretsEntry
	(*MicroserviceKubernetes)(nil), // 10: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.MicroserviceKubernetes
}
var file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_depIdxs = []int32{
	10, // 0: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.MicroserviceKubernetesList.entries:type_name -> cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.MicroserviceKubernetes
	10, // 1: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.MicroserviceKubernetesPage.entries:type_name -> cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.MicroserviceKubernetes
	7,  // 2: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.MicroserviceKubernetesEnvVarMap.value:type_name -> cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.MicroserviceKubernetesEnvVarMap.ValueEntry
	8,  // 3: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.GetMicroserviceKubernetesEnvVarMapInput.variables:type_name -> cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.GetMicroserviceKubernetesEnvVarMapInput.VariablesEntry
	9,  // 4: cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.GetMicroserviceKubernetesEnvVarMapInput.secrets:type_name -> cloud.planton.apis.code2cloud.v1.microservicekubernetes.model.GetMicroserviceKubernetesEnvVarMapInput.SecretsEntry
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto != nil {
		return
	}
	file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MicroserviceKubernetesId); i {
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
		file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MicroserviceKubernetesList); i {
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
		file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MicroserviceKubernetesPage); i {
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
		file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMicroserviceKubernetesLogStreamQueryInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByEnvironmentIdByCodeProjectIdInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MicroserviceKubernetesEnvVarMap); i {
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
		file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMicroserviceKubernetesEnvVarMapInput); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_microservicekubernetes_model_io_proto_depIdxs = nil
}
