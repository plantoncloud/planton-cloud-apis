// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/solr/state/keys/solr_cloud_state_key.proto

package keys

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

// wrapper for solr-cloud id used as key in solr messages
type SolrCloudStateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SolrCloudStateKey) Reset() {
	*x = SolrCloudStateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolrCloudStateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolrCloudStateKey) ProtoMessage() {}

func (x *SolrCloudStateKey) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolrCloudStateKey.ProtoReflect.Descriptor instead.
func (*SolrCloudStateKey) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_rawDescGZIP(), []int{0}
}

func (x *SolrCloudStateKey) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_rawDesc = []byte{
	0x0a, 0x52, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x73, 0x6f, 0x6c, 0x72, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x72, 0x5f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x6f,
	0x6c, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x29, 0x0a,
	0x11, 0x53, 0x6f, 0x6c, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xc8, 0x03, 0x0a, 0x3b, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x72, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x16, 0x53, 0x6f, 0x6c, 0x72, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a,
	0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x73, 0x6f, 0x6c, 0x72,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0xa2, 0x02, 0x09, 0x43, 0x50,
	0x41, 0x56, 0x43, 0x44, 0x53, 0x53, 0x4b, 0xaa, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x53, 0x6f, 0x6c, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x4b, 0x65, 0x79,
	0x73, 0xca, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x53, 0x6f, 0x6c, 0x72,
	0x5c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x5c, 0x4b, 0x65, 0x79, 0x73, 0xe2, 0x02, 0x43, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5c, 0x53, 0x6f, 0x6c, 0x72, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x5c, 0x4b, 0x65, 0x79, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x3f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x3a, 0x3a, 0x53, 0x6f, 0x6c, 0x72, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x4b,
	0x65, 0x79, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_goTypes = []interface{}{
	(*SolrCloudStateKey)(nil), // 0: cloud.planton.apis.v1.code2cloud.deploy.solr.state.keys.SolrCloudStateKey
}
var file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_init()
}
func file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolrCloudStateKey); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_deploy_solr_state_keys_solr_cloud_state_key_proto_depIdxs = nil
}
