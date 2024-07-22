// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/k8sd2ops/v1/podlogs/model/io.proto

package model

import (
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

// input for rpc to stream logs of a pod on kube-cluster based on specified filters
type StreamApiResourceKubernetesPodLogsInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// planton-cloud api-resource
	ApiResource *model.ApiResourceKindApiResourceId `protobuf:"bytes,1,opt,name=api_resource,json=apiResource,proto3" json:"api_resource,omitempty"`
	// options to configure pod log stream
	Options *model1.PodLogStreamOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *StreamApiResourceKubernetesPodLogsInput) Reset() {
	*x = StreamApiResourceKubernetesPodLogsInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamApiResourceKubernetesPodLogsInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamApiResourceKubernetesPodLogsInput) ProtoMessage() {}

func (x *StreamApiResourceKubernetesPodLogsInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamApiResourceKubernetesPodLogsInput.ProtoReflect.Descriptor instead.
func (*StreamApiResourceKubernetesPodLogsInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *StreamApiResourceKubernetesPodLogsInput) GetApiResource() *model.ApiResourceKindApiResourceId {
	if x != nil {
		return x.ApiResource
	}
	return nil
}

func (x *StreamApiResourceKubernetesPodLogsInput) GetOptions() *model1.PodLogStreamOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

var File_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x64, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x64,
	0x32, 0x6f, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x64, 0x6c, 0x6f, 0x67, 0x73, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4c, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x02, 0x0a, 0x27, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x6f, 0x64, 0x4c, 0x6f, 0x67,
	0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x6d, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x0b, 0x61, 0x70, 0x69, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x58, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6f, 0x64,
	0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0xfd, 0x02, 0x0a, 0x3a, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6b,
	0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x64, 0x6c, 0x6f,
	0x67, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x64, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x4b, 0x56, 0x50, 0x4d, 0xaa, 0x02, 0x2c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73,
	0x2e, 0x4b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x6f, 0x64,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x2c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x4b, 0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x6f, 0x64, 0x6c,
	0x6f, 0x67, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x38, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x4b,
	0x38, 0x73, 0x64, 0x32, 0x6f, 0x70, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x6f, 0x64, 0x6c, 0x6f,
	0x67, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x32, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x4b, 0x38, 0x73,
	0x64, 0x32, 0x6f, 0x70, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x50, 0x6f, 0x64, 0x6c, 0x6f,
	0x67, 0x73, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_rawDescData = file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_rawDesc
)

func file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_rawDescData
}

var file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_goTypes = []interface{}{
	(*StreamApiResourceKubernetesPodLogsInput)(nil), // 0: cloud.planton.apis.k8sd2ops.v1.podlogs.model.StreamApiResourceKubernetesPodLogsInput
	(*model.ApiResourceKindApiResourceId)(nil),      // 1: cloud.planton.apis.commons.apiresource.model.ApiResourceKindApiResourceId
	(*model1.PodLogStreamOptions)(nil),              // 2: cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.model.PodLogStreamOptions
}
var file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_depIdxs = []int32{
	1, // 0: cloud.planton.apis.k8sd2ops.v1.podlogs.model.StreamApiResourceKubernetesPodLogsInput.api_resource:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceKindApiResourceId
	2, // 1: cloud.planton.apis.k8sd2ops.v1.podlogs.model.StreamApiResourceKubernetesPodLogsInput.options:type_name -> cloud.planton.apis.integration.v1.kubernetes.kubernetesobject.model.PodLogStreamOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_init() }
func file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_init() {
	if File_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamApiResourceKubernetesPodLogsInput); i {
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
			RawDescriptor: file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto = out.File
	file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_goTypes = nil
	file_cloud_planton_apis_k8sd2ops_v1_podlogs_model_io_proto_depIdxs = nil
}
