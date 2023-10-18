// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/commons/testing/resource/field/immutable/model.proto

package immutable

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

// New message to test immutable fields.
type ImmutableFieldsTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImmutableStringField string `protobuf:"bytes,1,opt,name=immutable_string_field,json=immutableStringField,proto3" json:"immutable_string_field,omitempty"`
	MutableStringField   string `protobuf:"bytes,2,opt,name=mutable_string_field,json=mutableStringField,proto3" json:"mutable_string_field,omitempty"`
	ImmutableIntField    int32  `protobuf:"varint,3,opt,name=immutable_int_field,json=immutableIntField,proto3" json:"immutable_int_field,omitempty"`
	MutableIntField      int32  `protobuf:"varint,4,opt,name=mutable_int_field,json=mutableIntField,proto3" json:"mutable_int_field,omitempty"`
}

func (x *ImmutableFieldsTest) Reset() {
	*x = ImmutableFieldsTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImmutableFieldsTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImmutableFieldsTest) ProtoMessage() {}

func (x *ImmutableFieldsTest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImmutableFieldsTest.ProtoReflect.Descriptor instead.
func (*ImmutableFieldsTest) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDescGZIP(), []int{0}
}

func (x *ImmutableFieldsTest) GetImmutableStringField() string {
	if x != nil {
		return x.ImmutableStringField
	}
	return ""
}

func (x *ImmutableFieldsTest) GetMutableStringField() string {
	if x != nil {
		return x.MutableStringField
	}
	return ""
}

func (x *ImmutableFieldsTest) GetImmutableIntField() int32 {
	if x != nil {
		return x.ImmutableIntField
	}
	return 0
}

func (x *ImmutableFieldsTest) GetMutableIntField() int32 {
	if x != nil {
		return x.MutableIntField
	}
	return 0
}

// New message to test nested immutable fields.
type NestedImmutableFieldsTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NestedImmutableFields *ImmutableFieldsTest `protobuf:"bytes,1,opt,name=nested_immutable_fields,json=nestedImmutableFields,proto3" json:"nested_immutable_fields,omitempty"`
	MutableStringField    string               `protobuf:"bytes,2,opt,name=mutable_string_field,json=mutableStringField,proto3" json:"mutable_string_field,omitempty"`
}

func (x *NestedImmutableFieldsTest) Reset() {
	*x = NestedImmutableFieldsTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedImmutableFieldsTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedImmutableFieldsTest) ProtoMessage() {}

func (x *NestedImmutableFieldsTest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedImmutableFieldsTest.ProtoReflect.Descriptor instead.
func (*NestedImmutableFieldsTest) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDescGZIP(), []int{1}
}

func (x *NestedImmutableFieldsTest) GetNestedImmutableFields() *ImmutableFieldsTest {
	if x != nil {
		return x.NestedImmutableFields
	}
	return nil
}

func (x *NestedImmutableFieldsTest) GetMutableStringField() string {
	if x != nil {
		return x.MutableStringField
	}
	return ""
}

var File_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x2e, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x51, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe5, 0x01, 0x0a, 0x13, 0x49, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x54, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x16, 0x69, 0x6d, 0x6d, 0x75, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xd0, 0xb8, 0x18, 0x01, 0x52, 0x14, 0x69,
	0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x34, 0x0a, 0x13, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x04, 0xd0, 0xb8, 0x18, 0x01, 0x52, 0x11, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d,
	0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49,
	0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xdb, 0x01, 0x0a, 0x19, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x49, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x54, 0x65, 0x73, 0x74, 0x12, 0x8b, 0x01, 0x0a, 0x17, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x5f, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x69,
	0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x49, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x54, 0x65, 0x73, 0x74, 0x52, 0x15, 0x6e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x49, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0xf0, 0x03, 0x0a, 0x4c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x69, 0x6d, 0x6d,
	0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x69, 0x6d, 0x6d, 0x75, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x56, 0x43, 0x54, 0x52, 0x46, 0x49,
	0xaa, 0x02, 0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x49, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0xca, 0x02, 0x3e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x5c, 0x49, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0xe2, 0x02, 0x4a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x5c, 0x49, 0x6d, 0x6d, 0x75, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x46, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x3a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x3a, 0x49,
	0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDescData = file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_goTypes = []interface{}{
	(*ImmutableFieldsTest)(nil),       // 0: cloud.planton.apis.v1.commons.testing.resource.field.immutable.ImmutableFieldsTest
	(*NestedImmutableFieldsTest)(nil), // 1: cloud.planton.apis.v1.commons.testing.resource.field.immutable.NestedImmutableFieldsTest
}
var file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.commons.testing.resource.field.immutable.NestedImmutableFieldsTest.nested_immutable_fields:type_name -> cloud.planton.apis.v1.commons.testing.resource.field.immutable.ImmutableFieldsTest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_init() }
func file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_init() {
	if File_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImmutableFieldsTest); i {
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
		file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedImmutableFieldsTest); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto = out.File
	file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_testing_resource_field_immutable_model_proto_depIdxs = nil
}
