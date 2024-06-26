// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/prometheuskubernetes/model/io.proto

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

// wrapper for id field of prometheus-kubernetes
type PrometheusKubernetesId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PrometheusKubernetesId) Reset() {
	*x = PrometheusKubernetesId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrometheusKubernetesId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrometheusKubernetesId) ProtoMessage() {}

func (x *PrometheusKubernetesId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrometheusKubernetesId.ProtoReflect.Descriptor instead.
func (*PrometheusKubernetesId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *PrometheusKubernetesId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// list of prometheus-kubernetes
type PrometheusKubernetesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*PrometheusKubernetes `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *PrometheusKubernetesList) Reset() {
	*x = PrometheusKubernetesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrometheusKubernetesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrometheusKubernetesList) ProtoMessage() {}

func (x *PrometheusKubernetesList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrometheusKubernetesList.ProtoReflect.Descriptor instead.
func (*PrometheusKubernetesList) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *PrometheusKubernetesList) GetEntries() []*PrometheusKubernetes {
	if x != nil {
		return x.Entries
	}
	return nil
}

// wrapper for prometheus-kubernetes password
type PrometheusKubernetesPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PrometheusKubernetesPassword) Reset() {
	*x = PrometheusKubernetesPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrometheusKubernetesPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrometheusKubernetesPassword) ProtoMessage() {}

func (x *PrometheusKubernetesPassword) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrometheusKubernetesPassword.ProtoReflect.Descriptor instead.
func (*PrometheusKubernetesPassword) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDescGZIP(), []int{2}
}

func (x *PrometheusKubernetesPassword) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// response for paginated query to list prometheus-kubernetes
type PrometheusKubernetesPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPages int32                   `protobuf:"varint,1,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	Entries    []*PrometheusKubernetes `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *PrometheusKubernetesPage) Reset() {
	*x = PrometheusKubernetesPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrometheusKubernetesPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrometheusKubernetesPage) ProtoMessage() {}

func (x *PrometheusKubernetesPage) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrometheusKubernetesPage.ProtoReflect.Descriptor instead.
func (*PrometheusKubernetesPage) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDescGZIP(), []int{3}
}

func (x *PrometheusKubernetesPage) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *PrometheusKubernetesPage) GetEntries() []*PrometheusKubernetes {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDesc = []byte{
	0x0a, 0x44, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68,
	0x65, 0x75, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x1a, 0x47, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x16, 0x50, 0x72, 0x6f,
	0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x87, 0x01, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x6b,
	0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x51, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x1c, 0x50,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0xa8, 0x01, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x6b, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x51, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0xd7, 0x03, 0x0a,
	0x49, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x49, 0x6f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65,
	0x75, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x43, 0x56, 0x50, 0x4d, 0xaa, 0x02, 0x3b, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69,
	0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x3b, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x47, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x72, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x41, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a,
	0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_goTypes = []interface{}{
	(*PrometheusKubernetesId)(nil),       // 0: cloud.planton.apis.code2cloud.v1.prometheuskubernetes.model.PrometheusKubernetesId
	(*PrometheusKubernetesList)(nil),     // 1: cloud.planton.apis.code2cloud.v1.prometheuskubernetes.model.PrometheusKubernetesList
	(*PrometheusKubernetesPassword)(nil), // 2: cloud.planton.apis.code2cloud.v1.prometheuskubernetes.model.PrometheusKubernetesPassword
	(*PrometheusKubernetesPage)(nil),     // 3: cloud.planton.apis.code2cloud.v1.prometheuskubernetes.model.PrometheusKubernetesPage
	(*PrometheusKubernetes)(nil),         // 4: cloud.planton.apis.code2cloud.v1.prometheuskubernetes.model.PrometheusKubernetes
}
var file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.code2cloud.v1.prometheuskubernetes.model.PrometheusKubernetesList.entries:type_name -> cloud.planton.apis.code2cloud.v1.prometheuskubernetes.model.PrometheusKubernetes
	4, // 1: cloud.planton.apis.code2cloud.v1.prometheuskubernetes.model.PrometheusKubernetesPage.entries:type_name -> cloud.planton.apis.code2cloud.v1.prometheuskubernetes.model.PrometheusKubernetes
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto != nil {
		return
	}
	file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrometheusKubernetesId); i {
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
		file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrometheusKubernetesList); i {
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
		file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrometheusKubernetesPassword); i {
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
		file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrometheusKubernetesPage); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_prometheuskubernetes_model_io_proto_depIdxs = nil
}
