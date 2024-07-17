// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/iac/v1/pulumistatebackend/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/enums/apiresourcekind"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcefieldoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemessageoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions"
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

// pulumi-state-backend
type PulumiStateBackend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// resource metadata
	Metadata *model.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// resource spec
	Spec *PulumiStateBackendSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// resource status
	Status *PulumiStateBackendStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PulumiStateBackend) Reset() {
	*x = PulumiStateBackend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PulumiStateBackend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PulumiStateBackend) ProtoMessage() {}

func (x *PulumiStateBackend) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PulumiStateBackend.ProtoReflect.Descriptor instead.
func (*PulumiStateBackend) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *PulumiStateBackend) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *PulumiStateBackend) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *PulumiStateBackend) GetMetadata() *model.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PulumiStateBackend) GetSpec() *PulumiStateBackendSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *PulumiStateBackend) GetStatus() *PulumiStateBackendStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// pulumi-state-backend spec
type PulumiStateBackendSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the organization to which the pulumi-state-backend belongs to.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// grpc endpoint to be used for creating grpc client
	GrpcEndpoint string `protobuf:"bytes,2,opt,name=grpc_endpoint,json=grpcEndpoint,proto3" json:"grpc_endpoint,omitempty"`
	// grpc metadata to be attached to execute rpc call.
	// maintainer of pulumi-state-backend is responsible for
	// matching the grpc-metadata provided by client.
	GrpcMetadata map[string]string `protobuf:"bytes,3,rep,name=grpc_metadata,json=grpcMetadata,proto3" json:"grpc_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PulumiStateBackendSpec) Reset() {
	*x = PulumiStateBackendSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PulumiStateBackendSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PulumiStateBackendSpec) ProtoMessage() {}

func (x *PulumiStateBackendSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PulumiStateBackendSpec.ProtoReflect.Descriptor instead.
func (*PulumiStateBackendSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *PulumiStateBackendSpec) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *PulumiStateBackendSpec) GetGrpcEndpoint() string {
	if x != nil {
		return x.GrpcEndpoint
	}
	return ""
}

func (x *PulumiStateBackendSpec) GetGrpcMetadata() map[string]string {
	if x != nil {
		return x.GrpcMetadata
	}
	return nil
}

// pulumi-state-backend status
type PulumiStateBackendStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource lifecycle
	Lifecycle *model.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// resource audit info
	Audit *model.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
}

func (x *PulumiStateBackendStatus) Reset() {
	*x = PulumiStateBackendStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PulumiStateBackendStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PulumiStateBackendStatus) ProtoMessage() {}

func (x *PulumiStateBackendStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PulumiStateBackendStatus.ProtoReflect.Descriptor instead.
func (*PulumiStateBackendStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *PulumiStateBackendStatus) GetLifecycle() *model.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *PulumiStateBackendStatus) GetAudit() *model.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

var File_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x6c, 0x75,
	0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x75, 0x6c, 0x75,
	0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x54, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6b, 0x69, 0x6e, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6b, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x7c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x03, 0x0a, 0x12, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x12, 0x5d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x5e, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d,
	0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x12, 0x64, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x4c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x75, 0x6c,
	0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x1b, 0x90, 0xb5, 0x18, 0x01, 0x88, 0xa6, 0x1d, 0x23,
	0x92, 0xa6, 0x1d, 0x0f, 0x08, 0x1f, 0x12, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x6f, 0x72, 0x67,
	0x5f, 0x69, 0x64, 0x22, 0xad, 0x02, 0x0a, 0x16, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21,
	0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0xd0, 0xb8, 0x18, 0x01, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x81,
	0x01, 0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x5c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x75, 0x6c, 0x75,
	0x6d, 0x69, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x70,
	0x65, 0x63, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x3f, 0x0a, 0x11, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xd2, 0x01, 0x0a, 0x18, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x60, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x63, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x12, 0x54, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x62, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x42, 0xa4, 0x03, 0x0a, 0x40, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x62, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x69, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2,
	0x02, 0x07, 0x43, 0x50, 0x41, 0x49, 0x56, 0x50, 0x4d, 0xaa, 0x02, 0x32, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x49,
	0x61, 0x63, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02,
	0x32, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x49, 0x61, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x75, 0x6c, 0x75, 0x6d,
	0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5c, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0xe2, 0x02, 0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x49, 0x61, 0x63, 0x5c, 0x56, 0x31, 0x5c,
	0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x38, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x49, 0x61, 0x63,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDescData = file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDesc
)

func file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDescData
}

var file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_goTypes = []interface{}{
	(*PulumiStateBackend)(nil),         // 0: cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackend
	(*PulumiStateBackendSpec)(nil),     // 1: cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackendSpec
	(*PulumiStateBackendStatus)(nil),   // 2: cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackendStatus
	nil,                                // 3: cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackendSpec.GrpcMetadataEntry
	(*model.ApiResourceMetadata)(nil),  // 4: cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	(*model.ApiResourceLifecycle)(nil), // 5: cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	(*model.ApiResourceAudit)(nil),     // 6: cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
}
var file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackend.metadata:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	1, // 1: cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackend.spec:type_name -> cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackendSpec
	2, // 2: cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackend.status:type_name -> cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackendStatus
	3, // 3: cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackendSpec.grpc_metadata:type_name -> cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackendSpec.GrpcMetadataEntry
	5, // 4: cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackendStatus.lifecycle:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	6, // 5: cloud.planton.apis.iac.v1.pulumistatebackend.model.PulumiStateBackendStatus.audit:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_init() }
func file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_init() {
	if File_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PulumiStateBackend); i {
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
		file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PulumiStateBackendSpec); i {
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
		file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PulumiStateBackendStatus); i {
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
			RawDescriptor: file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto = out.File
	file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_goTypes = nil
	file_cloud_planton_apis_iac_v1_pulumistatebackend_model_state_proto_depIdxs = nil
}