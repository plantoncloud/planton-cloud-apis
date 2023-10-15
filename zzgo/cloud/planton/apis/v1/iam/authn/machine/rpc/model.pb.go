// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/iam/authn/machine/rpc/model.proto

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

// input for query to get access token for machine identity account
type GetMachineAccessTokenQueryInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// email of machine identity account
	MachineAccountEmail string `protobuf:"bytes,1,opt,name=machine_account_email,json=machineAccountEmail,proto3" json:"machine_account_email,omitempty"`
	// client secret of the machine identity account
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *GetMachineAccessTokenQueryInput) Reset() {
	*x = GetMachineAccessTokenQueryInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMachineAccessTokenQueryInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMachineAccessTokenQueryInput) ProtoMessage() {}

func (x *GetMachineAccessTokenQueryInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMachineAccessTokenQueryInput.ProtoReflect.Descriptor instead.
func (*GetMachineAccessTokenQueryInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDescGZIP(), []int{0}
}

func (x *GetMachineAccessTokenQueryInput) GetMachineAccountEmail() string {
	if x != nil {
		return x.MachineAccountEmail
	}
	return ""
}

func (x *GetMachineAccessTokenQueryInput) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

// wrapper for access token
type AccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AccessToken) Reset() {
	*x = AccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessToken) ProtoMessage() {}

func (x *AccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessToken.ProtoReflect.Descriptor instead.
func (*AccessToken) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDescGZIP(), []int{1}
}

func (x *AccessToken) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x6e, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x22, 0x7a, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x22, 0x23, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xf2, 0x02, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x0a, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x2f, 0x72, 0x70, 0x63, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x49, 0x41, 0x4d,
	0x52, 0xaa, 0x02, 0x2b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x49, 0x61, 0x6d, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6e, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x70, 0x63, 0xca,
	0x02, 0x2b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x41, 0x75, 0x74, 0x68,
	0x6e, 0x5c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5c, 0x52, 0x70, 0x63, 0xe2, 0x02, 0x37,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x5c,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x32, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x3a, 0x3a, 0x49, 0x61, 0x6d, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x3a, 0x3a,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDescData = file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_goTypes = []interface{}{
	(*GetMachineAccessTokenQueryInput)(nil), // 0: cloud.planton.apis.v1.iam.authn.machine.rpc.GetMachineAccessTokenQueryInput
	(*AccessToken)(nil),                     // 1: cloud.planton.apis.v1.iam.authn.machine.rpc.AccessToken
}
var file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_init() }
func file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_init() {
	if File_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMachineAccessTokenQueryInput); i {
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
		file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessToken); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto = out.File
	file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_iam_authn_machine_rpc_model_proto_depIdxs = nil
}
