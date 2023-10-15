// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/commons/testing/resource/field/computed/model.proto

package computed

import (
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/field/options"
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

// A message with computed fields.
type ComputedFieldsTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequiredStringField string `protobuf:"bytes,1,opt,name=required_string_field,json=requiredStringField,proto3" json:"required_string_field,omitempty"`
	ComputedStringField string `protobuf:"bytes,2,opt,name=computed_string_field,json=computedStringField,proto3" json:"computed_string_field,omitempty"`
}

func (x *ComputedFieldsTest) Reset() {
	*x = ComputedFieldsTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputedFieldsTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputedFieldsTest) ProtoMessage() {}

func (x *ComputedFieldsTest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputedFieldsTest.ProtoReflect.Descriptor instead.
func (*ComputedFieldsTest) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDescGZIP(), []int{0}
}

func (x *ComputedFieldsTest) GetRequiredStringField() string {
	if x != nil {
		return x.RequiredStringField
	}
	return ""
}

func (x *ComputedFieldsTest) GetComputedStringField() string {
	if x != nil {
		return x.ComputedStringField
	}
	return ""
}

// A message with a nested message containing computed fields.
type NestedComputedFieldsTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NestedComputedField *ComputedFieldsTest `protobuf:"bytes,1,opt,name=nested_computed_field,json=nestedComputedField,proto3" json:"nested_computed_field,omitempty"`
}

func (x *NestedComputedFieldsTest) Reset() {
	*x = NestedComputedFieldsTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedComputedFieldsTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedComputedFieldsTest) ProtoMessage() {}

func (x *NestedComputedFieldsTest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedComputedFieldsTest.ProtoReflect.Descriptor instead.
func (*NestedComputedFieldsTest) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDescGZIP(), []int{1}
}

func (x *NestedComputedFieldsTest) GetNestedComputedField() *ComputedFieldsTest {
	if x != nil {
		return x.NestedComputedField
	}
	return nil
}

var File_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDesc = []byte{
	0x0a, 0x49, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x1a, 0x51, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01,
	0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x54, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x15, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x38,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8,
	0xb8, 0x18, 0x01, 0x52, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x18, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x54, 0x65, 0x73, 0x74, 0x12, 0x85, 0x01, 0x0a, 0x15, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x54, 0x65, 0x73, 0x74, 0x52, 0x13, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0xea, 0x03,
	0x0a, 0x4b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x42, 0x0a, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41,
	0x56, 0x43, 0x54, 0x52, 0x46, 0x43, 0xaa, 0x02, 0x3d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0xca, 0x02, 0x3d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x5c, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0xe2, 0x02, 0x49, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x5c, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x45, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x3a,
	0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x3a, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDescData = file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_goTypes = []interface{}{
	(*ComputedFieldsTest)(nil),       // 0: cloud.planton.apis.v1.commons.testing.resource.field.computed.ComputedFieldsTest
	(*NestedComputedFieldsTest)(nil), // 1: cloud.planton.apis.v1.commons.testing.resource.field.computed.NestedComputedFieldsTest
}
var file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.commons.testing.resource.field.computed.NestedComputedFieldsTest.nested_computed_field:type_name -> cloud.planton.apis.v1.commons.testing.resource.field.computed.ComputedFieldsTest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_init() }
func file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_init() {
	if File_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputedFieldsTest); i {
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
		file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedComputedFieldsTest); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto = out.File
	file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_testing_resource_field_computed_model_proto_depIdxs = nil
}