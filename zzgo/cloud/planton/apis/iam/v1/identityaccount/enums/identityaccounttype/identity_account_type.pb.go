// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/iam/v1/identityaccount/enums/identityaccounttype/identity_account_type.proto

package identityaccounttype

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

type IdentityAccountType int32

const (
	IdentityAccountType_unspecified IdentityAccountType = 0
	IdentityAccountType_user        IdentityAccountType = 1
	IdentityAccountType_machine     IdentityAccountType = 2
)

// Enum value maps for IdentityAccountType.
var (
	IdentityAccountType_name = map[int32]string{
		0: "unspecified",
		1: "user",
		2: "machine",
	}
	IdentityAccountType_value = map[string]int32{
		"unspecified": 0,
		"user":        1,
		"machine":     2,
	}
)

func (x IdentityAccountType) Enum() *IdentityAccountType {
	p := new(IdentityAccountType)
	*p = x
	return p
}

func (x IdentityAccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentityAccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_enumTypes[0].Descriptor()
}

func (IdentityAccountType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_enumTypes[0]
}

func (x IdentityAccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentityAccountType.Descriptor instead.
func (IdentityAccountType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_rawDesc = []byte{
	0x0a, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x3d, 0x0a, 0x13, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x10, 0x02, 0x42, 0x9a, 0x04, 0x0a, 0x51, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x42, 0x18, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x73, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0xa2, 0x02, 0x08, 0x43,
	0x50, 0x41, 0x49, 0x56, 0x49, 0x45, 0x49, 0xaa, 0x02, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x49, 0x61, 0x6d,
	0x2e, 0x56, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0xca, 0x02, 0x43,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74,
	0x79, 0x70, 0x65, 0xe2, 0x02, 0x4f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x56, 0x31, 0x5c,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x4a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x49, 0x61,
	0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x74, 0x79,
	0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_rawDescData = file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_rawDesc
)

func file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_rawDescData)
	})
	return file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_rawDescData
}

var file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_goTypes = []interface{}{
	(IdentityAccountType)(0), // 0: cloud.planton.apis.iam.v1.identityaccount.enums.identityaccounttype.IdentityAccountType
}
var file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_init()
}
func file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_init() {
	if File_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto = out.File
	file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_rawDesc = nil
	file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_goTypes = nil
	file_cloud_planton_apis_iam_v1_identityaccount_enums_identityaccounttype_identity_account_type_proto_depIdxs = nil
}
