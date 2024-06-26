// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/commons/english/enums/englishacronym/english_acronym.proto

package englishacronym

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

type EnglishAcronym_EnglishAcronymEnum int32

const (
	EnglishAcronym_unspecified EnglishAcronym_EnglishAcronymEnum = 0
	// network
	EnglishAcronym_NW EnglishAcronym_EnglishAcronymEnum = 1
	// rpc
	EnglishAcronym_RPC EnglishAcronym_EnglishAcronymEnum = 2
)

// Enum value maps for EnglishAcronym_EnglishAcronymEnum.
var (
	EnglishAcronym_EnglishAcronymEnum_name = map[int32]string{
		0: "unspecified",
		1: "NW",
		2: "RPC",
	}
	EnglishAcronym_EnglishAcronymEnum_value = map[string]int32{
		"unspecified": 0,
		"NW":          1,
		"RPC":         2,
	}
)

func (x EnglishAcronym_EnglishAcronymEnum) Enum() *EnglishAcronym_EnglishAcronymEnum {
	p := new(EnglishAcronym_EnglishAcronymEnum)
	*p = x
	return p
}

func (x EnglishAcronym_EnglishAcronymEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnglishAcronym_EnglishAcronymEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_enumTypes[0].Descriptor()
}

func (EnglishAcronym_EnglishAcronymEnum) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_enumTypes[0]
}

func (x EnglishAcronym_EnglishAcronymEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnglishAcronym_EnglishAcronymEnum.Descriptor instead.
func (EnglishAcronym_EnglishAcronymEnum) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDescGZIP(), []int{0, 0}
}

// english acronym enums added to avoid typing mistakes for the commonly used acronyms for naming things
// this enum is encapsulated inside a message as a few entries like "name" (a reserved word in few languages) can only be added to the enum if it is inside a message.
// the recommended best practice to prefix the entry with enum name has been intentionally ignored to as the values of the entries are intended to be used to name resources.
type EnglishAcronym struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnglishAcronym) Reset() {
	*x = EnglishAcronym{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnglishAcronym) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnglishAcronym) ProtoMessage() {}

func (x *EnglishAcronym) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnglishAcronym.ProtoReflect.Descriptor instead.
func (*EnglishAcronym) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x67,
	0x6c, 0x69, 0x73, 0x68, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x67, 0x6c, 0x69,
	0x73, 0x68, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2f, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73,
	0x68, 0x5f, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x67, 0x6c,
	0x69, 0x73, 0x68, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73,
	0x68, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x22, 0x48, 0x0a, 0x0e, 0x45, 0x6e, 0x67, 0x6c,
	0x69, 0x73, 0x68, 0x41, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x22, 0x36, 0x0a, 0x12, 0x45, 0x6e,
	0x67, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x0f, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10,
	0x00, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x57, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x50, 0x43,
	0x10, 0x02, 0x42, 0xcb, 0x03, 0x0a, 0x45, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x65, 0x6e,
	0x67, 0x6c, 0x69, 0x73, 0x68, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x42, 0x13, 0x45, 0x6e,
	0x67, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e,
	0x67, 0x6c, 0x69, 0x73, 0x68, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0xa2, 0x02, 0x07, 0x43,
	0x50, 0x41, 0x43, 0x45, 0x45, 0x45, 0xaa, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d,
	0xca, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x45, 0x6e,
	0x67, 0x6c, 0x69, 0x73, 0x68, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x45, 0x6e, 0x67, 0x6c,
	0x69, 0x73, 0x68, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0xe2, 0x02, 0x43, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x5c,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x61, 0x63, 0x72,
	0x6f, 0x6e, 0x79, 0x6d, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x3d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x3a, 0x3a, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x3a, 0x3a, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x61, 0x63, 0x72, 0x6f, 0x6e, 0x79, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDescData = file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDesc
)

func file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDescData)
	})
	return file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDescData
}

var file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_goTypes = []interface{}{
	(EnglishAcronym_EnglishAcronymEnum)(0), // 0: cloud.planton.apis.commons.english.enums.englishacronym.EnglishAcronym.EnglishAcronymEnum
	(*EnglishAcronym)(nil),                 // 1: cloud.planton.apis.commons.english.enums.englishacronym.EnglishAcronym
}
var file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_init()
}
func file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_init() {
	if File_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnglishAcronym); i {
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
			RawDescriptor: file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_enumTypes,
		MessageInfos:      file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto = out.File
	file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_rawDesc = nil
	file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_goTypes = nil
	file_cloud_planton_apis_commons_english_enums_englishacronym_english_acronym_proto_depIdxs = nil
}
