// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/deploy/gcp/gcpartifactregistry/model/io.proto

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

// wrapper for gcp-artifact-registry id
type GcpArtifactRegistryId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GcpArtifactRegistryId) Reset() {
	*x = GcpArtifactRegistryId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpArtifactRegistryId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpArtifactRegistryId) ProtoMessage() {}

func (x *GcpArtifactRegistryId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpArtifactRegistryId.ProtoReflect.Descriptor instead.
func (*GcpArtifactRegistryId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *GcpArtifactRegistryId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// list of gcp-artifact-registry resources
type GcpArtifactRegistryList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*GcpArtifactRegistry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *GcpArtifactRegistryList) Reset() {
	*x = GcpArtifactRegistryList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpArtifactRegistryList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpArtifactRegistryList) ProtoMessage() {}

func (x *GcpArtifactRegistryList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpArtifactRegistryList.ProtoReflect.Descriptor instead.
func (*GcpArtifactRegistryList) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *GcpArtifactRegistryList) GetEntries() []*GcpArtifactRegistry {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x45, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x51, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x15, 0x47, 0x63, 0x70, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8f,
	0x01, 0x0a, 0x17, 0x47, 0x63, 0x70, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x74, 0x0a, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x47, 0x63, 0x70, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x42, 0x97, 0x04, 0x0a, 0x53, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67,
	0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67,
	0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x75, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x67, 0x63, 0x70,
	0x2f, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41,
	0x43, 0x56, 0x44, 0x47, 0x47, 0x4d, 0xaa, 0x02, 0x45, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x47, 0x63, 0x70, 0x2e, 0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02,
	0x45, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56,
	0x31, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x51, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x4d, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69,
	0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56,
	0x31, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x3a, 0x3a,
	0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_goTypes = []interface{}{
	(*GcpArtifactRegistryId)(nil),   // 0: cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpartifactregistry.model.GcpArtifactRegistryId
	(*GcpArtifactRegistryList)(nil), // 1: cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpartifactregistry.model.GcpArtifactRegistryList
	(*GcpArtifactRegistry)(nil),     // 2: cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpartifactregistry.model.GcpArtifactRegistry
}
var file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_depIdxs = []int32{
	2, // 0: cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpartifactregistry.model.GcpArtifactRegistryList.entries:type_name -> cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpartifactregistry.model.GcpArtifactRegistry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_init()
}
func file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto != nil {
		return
	}
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpArtifactRegistryId); i {
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
		file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpArtifactRegistryList); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_deploy_gcp_gcpartifactregistry_model_io_proto_depIdxs = nil
}
