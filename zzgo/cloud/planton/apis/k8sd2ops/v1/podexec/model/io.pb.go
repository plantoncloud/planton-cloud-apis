// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/k8sd2ops/v1/podexec/model/io.proto

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

// input for rpc to exec into a kube-cluster pod container
// input for rpc to exec into a pod container that belongs to a planton-cloud api-resource
type ExecIntoPodContainerInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// planton-cloud api-resource
	ApiResource *model.ApiResourceKindApiResourceId `protobuf:"bytes,1,opt,name=api_resource,json=apiResource,proto3" json:"api_resource,omitempty"`
	// options required for performing exec into a pod container
	Options *model1.PodContainerExecOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *ExecIntoPodContainerInput) Reset() {
	*x = ExecIntoPodContainerInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecIntoPodContainerInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecIntoPodContainerInput) ProtoMessage() {}

func (x *ExecIntoPodContainerInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecIntoPodContainerInput.ProtoReflect.Descriptor instead.
func (*ExecIntoPodContainerInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *ExecIntoPodContainerInput) GetApiResource() *model.ApiResourceKindApiResourceId {
	if x != nil {
		return x.ApiResource
	}
	return nil
}

func (x *ExecIntoPodContainerInput) GetOptions() *model1.PodContainerExecOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

// id of the stream created for the shell session
type BrowserExecIntoPodContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pod container exec shell session id
	ShellSessionId string `protobuf:"bytes,1,opt,name=shell_session_id,json=shellSessionId,proto3" json:"shell_session_id,omitempty"`
	// shell command execution response
	CommandResponse *model1.ExecIntoPodContainerResponse `protobuf:"bytes,2,opt,name=command_response,json=commandResponse,proto3" json:"command_response,omitempty"`
}

func (x *BrowserExecIntoPodContainerResponse) Reset() {
	*x = BrowserExecIntoPodContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrowserExecIntoPodContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowserExecIntoPodContainerResponse) ProtoMessage() {}

func (x *BrowserExecIntoPodContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowserExecIntoPodContainerResponse.ProtoReflect.Descriptor instead.
func (*BrowserExecIntoPodContainerResponse) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *BrowserExecIntoPodContainerResponse) GetShellSessionId() string {
	if x != nil {
		return x.ShellSessionId
	}
	return ""
}

func (x *BrowserExecIntoPodContainerResponse) GetCommandResponse() *model1.ExecIntoPodContainerResponse {
	if x != nil {
		return x.CommandResponse
	}
	return nil
}

// input for rpc to execute next command for a shell session in a pod container from a browser
type BrowserExecuteNextCommandInPodContainerInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// planton-cloud api-resource
	ApiResource *model.ApiResourceKindApiResourceId `protobuf:"bytes,1,opt,name=api_resource,json=apiResource,proto3" json:"api_resource,omitempty"`
	// pod container exec shell session id
	ShellSessionId string `protobuf:"bytes,2,opt,name=shell_session_id,json=shellSessionId,proto3" json:"shell_session_id,omitempty"`
	// command to execute inside an existing shell session of a pod container
	Command string `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *BrowserExecuteNextCommandInPodContainerInput) Reset() {
	*x = BrowserExecuteNextCommandInPodContainerInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrowserExecuteNextCommandInPodContainerInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowserExecuteNextCommandInPodContainerInput) ProtoMessage() {}

func (x *BrowserExecuteNextCommandInPodContainerInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowserExecuteNextCommandInPodContainerInput.ProtoReflect.Descriptor instead.
func (*BrowserExecuteNextCommandInPodContainerInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDescGZIP(), []int{2}
}

func (x *BrowserExecuteNextCommandInPodContainerInput) GetApiResource() *model.ApiResourceKindApiResourceId {
	if x != nil {
		return x.ApiResource
	}
	return nil
}

func (x *BrowserExecuteNextCommandInPodContainerInput) GetShellSessionId() string {
	if x != nil {
		return x.ShellSessionId
	}
	return ""
}

func (x *BrowserExecuteNextCommandInPodContainerInput) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

var File_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x64,
	0x32, 0x6f, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4c, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02, 0x0a, 0x19, 0x45, 0x78, 0x65, 0x63,
	0x49, 0x6e, 0x74, 0x6f, 0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x6d, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x0b, 0x61, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6f, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xde, 0x01, 0x0a,
	0x23, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x49, 0x6e, 0x74, 0x6f,
	0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x73, 0x68, 0x65, 0x6c, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x8c,
	0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x78, 0x65, 0x63, 0x49, 0x6e, 0x74, 0x6f, 0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe1, 0x01,
	0x0a, 0x2c, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x4e, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x50, 0x6f, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x6d,
	0x0a, 0x0c, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b,
	0x69, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x52, 0x0b, 0x61, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x42, 0xfd, 0x02, 0x0a, 0x3a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x42, 0x07, 0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x64, 0x65,
	0x78, 0x65, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x4b,
	0x56, 0x50, 0x4d, 0xaa, 0x02, 0x2c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70,
	0x73, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0xca, 0x02, 0x2c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x4b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x50, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x5c, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0xe2, 0x02, 0x38, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x4b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x50, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x32, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41,
	0x70, 0x69, 0x73, 0x3a, 0x3a, 0x4b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x3a, 0x3a, 0x50, 0x6f, 0x64, 0x65, 0x78, 0x65, 0x63, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDescData = file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDesc
)

func file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDescData
}

var file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_goTypes = []interface{}{
	(*ExecIntoPodContainerInput)(nil),                    // 0: cloud.planton.apis.k8sd2ops.v1.podexec.model.ExecIntoPodContainerInput
	(*BrowserExecIntoPodContainerResponse)(nil),          // 1: cloud.planton.apis.k8sd2ops.v1.podexec.model.BrowserExecIntoPodContainerResponse
	(*BrowserExecuteNextCommandInPodContainerInput)(nil), // 2: cloud.planton.apis.k8sd2ops.v1.podexec.model.BrowserExecuteNextCommandInPodContainerInput
	(*model.ApiResourceKindApiResourceId)(nil),           // 3: cloud.planton.apis.commons.apiresource.model.ApiResourceKindApiResourceId
	(*model1.PodContainerExecOptions)(nil),               // 4: cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.model.PodContainerExecOptions
	(*model1.ExecIntoPodContainerResponse)(nil),          // 5: cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.model.ExecIntoPodContainerResponse
}
var file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_depIdxs = []int32{
	3, // 0: cloud.planton.apis.k8sd2ops.v1.podexec.model.ExecIntoPodContainerInput.api_resource:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceKindApiResourceId
	4, // 1: cloud.planton.apis.k8sd2ops.v1.podexec.model.ExecIntoPodContainerInput.options:type_name -> cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.model.PodContainerExecOptions
	5, // 2: cloud.planton.apis.k8sd2ops.v1.podexec.model.BrowserExecIntoPodContainerResponse.command_response:type_name -> cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.model.ExecIntoPodContainerResponse
	3, // 3: cloud.planton.apis.k8sd2ops.v1.podexec.model.BrowserExecuteNextCommandInPodContainerInput.api_resource:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceKindApiResourceId
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_init() }
func file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_init() {
	if File_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecIntoPodContainerInput); i {
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
		file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrowserExecIntoPodContainerResponse); i {
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
		file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrowserExecuteNextCommandInPodContainerInput); i {
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
			RawDescriptor: file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto = out.File
	file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_goTypes = nil
	file_cloud_planton_apis_k8sd2ops_v1_podexec_model_io_proto_depIdxs = nil
}
