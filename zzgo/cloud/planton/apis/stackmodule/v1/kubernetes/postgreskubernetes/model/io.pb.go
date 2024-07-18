// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/stackmodule/v1/kubernetes/postgreskubernetes/model/io.proto

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

// wrapper for id field of postgres-kubernetes
type PostgresKubernetesId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PostgresKubernetesId) Reset() {
	*x = PostgresKubernetesId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostgresKubernetesId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresKubernetesId) ProtoMessage() {}

func (x *PostgresKubernetesId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresKubernetesId.ProtoReflect.Descriptor instead.
func (*PostgresKubernetesId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *PostgresKubernetesId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// wrapper for postgres-kubernetes password
type PostgresKubernetesPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PostgresKubernetesPassword) Reset() {
	*x = PostgresKubernetesPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostgresKubernetesPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresKubernetesPassword) ProtoMessage() {}

func (x *PostgresKubernetesPassword) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresKubernetesPassword.ProtoReflect.Descriptor instead.
func (*PostgresKubernetesPassword) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *PostgresKubernetesPassword) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x51, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x14, 0x50, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x32, 0x0a,
	0x1a, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x95, 0x04, 0x0a, 0x53, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x49, 0x6f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x75, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x08, 0x43, 0x50,
	0x41, 0x53, 0x56, 0x4b, 0x50, 0x4d, 0xaa, 0x02, 0x45, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02,
	0x45, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x50, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x51, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x4c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69,
	0x73, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a,
	0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDescData = file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDesc
)

func file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDescData
}

var file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_goTypes = []interface{}{
	(*PostgresKubernetesId)(nil),       // 0: cloud.planton.apis.stackmodule.v1.kubernetes.postgreskubernetes.model.PostgresKubernetesId
	(*PostgresKubernetesPassword)(nil), // 1: cloud.planton.apis.stackmodule.v1.kubernetes.postgreskubernetes.model.PostgresKubernetesPassword
}
var file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_init()
}
func file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_init() {
	if File_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto != nil {
		return
	}
	file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostgresKubernetesId); i {
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
		file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostgresKubernetesPassword); i {
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
			RawDescriptor: file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto = out.File
	file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_goTypes = nil
	file_cloud_planton_apis_stackmodule_v1_kubernetes_postgreskubernetes_model_io_proto_depIdxs = nil
}
