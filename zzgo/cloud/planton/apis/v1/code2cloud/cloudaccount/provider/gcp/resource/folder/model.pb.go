// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/code2cloud/cloudaccount/provider/gcp/resource/folder/model.proto

package folder

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

// gcp project
// https://cloud.google.com/resource-manager/reference/rest/v2/folders
type GcpFolder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the gcp project.
	// ex: 498520640386.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// display name for the folder.
	// (important) a two character suffix starting with a hyphen is added at the end of the
	// display name to allow recreating folders with same resource identifier on planton cloud.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// parent of the folder.
	// parent can be either a gcp organization or another gcp folder.
	Parent string `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *GcpFolder) Reset() {
	*x = GcpFolder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpFolder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpFolder) ProtoMessage() {}

func (x *GcpFolder) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpFolder.ProtoReflect.Descriptor instead.
func (*GcpFolder) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDescGZIP(), []int{0}
}

func (x *GcpFolder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GcpFolder) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *GcpFolder) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

// list of gcp folder
type GcpFolders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*GcpFolder `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *GcpFolders) Reset() {
	*x = GcpFolders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpFolders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpFolders) ProtoMessage() {}

func (x *GcpFolders) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpFolders.ProtoReflect.Descriptor instead.
func (*GcpFolders) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDescGZIP(), []int{1}
}

func (x *GcpFolders) GetEntries() []*GcpFolder {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDesc = []byte{
	0x0a, 0x56, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x67, 0x63, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x22, 0x56, 0x0a, 0x09, 0x47, 0x63, 0x70, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x7d, 0x0a, 0x0a,
	0x47, 0x63, 0x70, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x6f, 0x0a, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x63, 0x70, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0xba, 0x04, 0x0a, 0x58,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x7a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67,
	0x63, 0x70, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0xa2, 0x02, 0x0a, 0x43, 0x50, 0x41, 0x56, 0x43, 0x43, 0x50, 0x47, 0x52, 0x46, 0xaa,
	0x02, 0x4a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x63, 0x70, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0xca, 0x02, 0x4a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69,
	0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5c, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0xe2, 0x02, 0x56, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56,
	0x31, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x53, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x3a, 0x3a, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDescData = file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_goTypes = []interface{}{
	(*GcpFolder)(nil),  // 0: cloud.planton.apis.v1.code2cloud.cloudaccount.provider.gcp.resource.folder.GcpFolder
	(*GcpFolders)(nil), // 1: cloud.planton.apis.v1.code2cloud.cloudaccount.provider.gcp.resource.folder.GcpFolders
}
var file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.code2cloud.cloudaccount.provider.gcp.resource.folder.GcpFolders.entries:type_name -> cloud.planton.apis.v1.code2cloud.cloudaccount.provider.gcp.resource.folder.GcpFolder
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_init()
}
func file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_init() {
	if File_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpFolder); i {
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
		file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpFolders); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto = out.File
	file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_code2cloud_cloudaccount_provider_gcp_resource_folder_model_proto_depIdxs = nil
}
