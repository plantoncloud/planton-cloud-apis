// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/commons/cloud/gcp/resource/project/rpc/model.proto

package rpc

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
type GcpProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the gcp project
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// id of the gcp project
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// number of the gcp project
	Number string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	// id of the billing account
	BillingAccountId string `protobuf:"bytes,4,opt,name=billing_account_id,json=billingAccountId,proto3" json:"billing_account_id,omitempty"`
}

func (x *GcpProject) Reset() {
	*x = GcpProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpProject) ProtoMessage() {}

func (x *GcpProject) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpProject.ProtoReflect.Descriptor instead.
func (*GcpProject) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDescGZIP(), []int{0}
}

func (x *GcpProject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GcpProject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GcpProject) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *GcpProject) GetBillingAccountId() string {
	if x != nil {
		return x.BillingAccountId
	}
	return ""
}

// list of gcp projects
type GcpProjects struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*GcpProject `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GcpProjects) Reset() {
	*x = GcpProjects{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpProjects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpProjects) ProtoMessage() {}

func (x *GcpProjects) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpProjects.ProtoReflect.Descriptor instead.
func (*GcpProjects) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDescGZIP(), []int{1}
}

func (x *GcpProjects) GetList() []*GcpProject {
	if x != nil {
		return x.List
	}
	return nil
}

var File_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDesc = []byte{
	0x0a, 0x48, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3c, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x63, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x22, 0x76, 0x0a, 0x0a, 0x47, 0x63, 0x70, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x6b, 0x0a, 0x0b, 0x47, 0x63, 0x70, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12,
	0x5c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x63, 0x70,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x42, 0xdc, 0x03,
	0x0a, 0x40, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x72,
	0x70, 0x63, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x72, 0x70, 0x63, 0xa2, 0x02,
	0x0a, 0x43, 0x50, 0x41, 0x56, 0x43, 0x43, 0x47, 0x52, 0x50, 0x52, 0xaa, 0x02, 0x3c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73,
	0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x47, 0x63, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x52, 0x70, 0x63, 0xca, 0x02, 0x3c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x47, 0x63, 0x70, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x52, 0x70, 0x63, 0xe2, 0x02, 0x48, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56,
	0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x47, 0x63, 0x70, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x45, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x47, 0x63, 0x70, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x3a,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDescData = file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_goTypes = []interface{}{
	(*GcpProject)(nil),  // 0: cloud.planton.apis.v1.commons.cloud.gcp.resource.project.rpc.GcpProject
	(*GcpProjects)(nil), // 1: cloud.planton.apis.v1.commons.cloud.gcp.resource.project.rpc.GcpProjects
}
var file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.commons.cloud.gcp.resource.project.rpc.GcpProjects.list:type_name -> cloud.planton.apis.v1.commons.cloud.gcp.resource.project.rpc.GcpProject
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_init() }
func file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_init() {
	if File_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpProject); i {
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
		file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpProjects); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto = out.File
	file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_cloud_gcp_resource_project_rpc_model_proto_depIdxs = nil
}
