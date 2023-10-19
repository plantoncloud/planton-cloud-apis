// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/commons/testing/resource/metadata/model.proto

package metadata

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/metadata/options"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/testing/resource/enums"
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

type ResourceIdExtractionTestWithOutIdPrefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *ResourceIdExtractionTestWithOutIdPrefix) Reset() {
	*x = ResourceIdExtractionTestWithOutIdPrefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceIdExtractionTestWithOutIdPrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceIdExtractionTestWithOutIdPrefix) ProtoMessage() {}

func (x *ResourceIdExtractionTestWithOutIdPrefix) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceIdExtractionTestWithOutIdPrefix.ProtoReflect.Descriptor instead.
func (*ResourceIdExtractionTestWithOutIdPrefix) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceIdExtractionTestWithOutIdPrefix) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ResourceIdExtractionTestWithOutIdPrefix) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type ResourceIdExtractionTestWithIdPrefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *ResourceIdExtractionTestWithIdPrefix) Reset() {
	*x = ResourceIdExtractionTestWithIdPrefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceIdExtractionTestWithIdPrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceIdExtractionTestWithIdPrefix) ProtoMessage() {}

func (x *ResourceIdExtractionTestWithIdPrefix) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceIdExtractionTestWithIdPrefix.ProtoReflect.Descriptor instead.
func (*ResourceIdExtractionTestWithIdPrefix) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceIdExtractionTestWithIdPrefix) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ResourceIdExtractionTestWithIdPrefix) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type RegexNameFieldProtoValidateTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *MetadataTest `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *RegexNameFieldProtoValidateTest) Reset() {
	*x = RegexNameFieldProtoValidateTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegexNameFieldProtoValidateTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegexNameFieldProtoValidateTest) ProtoMessage() {}

func (x *RegexNameFieldProtoValidateTest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegexNameFieldProtoValidateTest.ProtoReflect.Descriptor instead.
func (*RegexNameFieldProtoValidateTest) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDescGZIP(), []int{2}
}

func (x *RegexNameFieldProtoValidateTest) GetMetadata() *MetadataTest {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type SplitRegexNameFieldProtoValidateTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *MetadataTest `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *SplitRegexNameFieldProtoValidateTest) Reset() {
	*x = SplitRegexNameFieldProtoValidateTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitRegexNameFieldProtoValidateTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitRegexNameFieldProtoValidateTest) ProtoMessage() {}

func (x *SplitRegexNameFieldProtoValidateTest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitRegexNameFieldProtoValidateTest.ProtoReflect.Descriptor instead.
func (*SplitRegexNameFieldProtoValidateTest) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDescGZIP(), []int{3}
}

func (x *SplitRegexNameFieldProtoValidateTest) GetMetadata() *MetadataTest {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type MetadataTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MetadataTest) Reset() {
	*x = MetadataTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataTest) ProtoMessage() {}

func (x *MetadataTest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataTest.ProtoReflect.Descriptor instead.
func (*MetadataTest) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDescGZIP(), []int{4}
}

func (x *MetadataTest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetadataTest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDesc = []byte{
	0x0a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x57,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x27, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x75, 0x74, 0x49, 0x64, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x6b, 0x0a, 0x24, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x49, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x07, 0x8a, 0xb5, 0x18, 0x03, 0x68, 0x63, 0x2d, 0x22, 0x81,
	0x03, 0x0a, 0x1f, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x12, 0xdd, 0x02, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x54, 0x65, 0x73, 0x74, 0x42, 0xf9, 0x01, 0xba,
	0x48, 0xf5, 0x01, 0xba, 0x01, 0xf1, 0x01, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xa7, 0x01, 0x41, 0x74, 0x20, 0x6c, 0x65, 0x61, 0x73,
	0x74, 0x20, 0x31, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x20, 0x4c,
	0x6f, 0x77, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73,
	0x2c, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x68,
	0x79, 0x70, 0x68, 0x65, 0x6e, 0x73, 0x2e, 0x20, 0x4d, 0x75, 0x73, 0x74, 0x20, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x63, 0x61, 0x73,
	0x65, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x20, 0x4d, 0x75, 0x73, 0x74, 0x20, 0x6e,
	0x6f, 0x74, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x20, 0x6f, 0x72, 0x20, 0x65, 0x6e, 0x64, 0x20,
	0x77, 0x69, 0x74, 0x68, 0x20, 0x68, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x2e, 0x20, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x75, 0x6e, 0x64, 0x65,
	0x72, 0x20, 0x31, 0x35, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x1a, 0x36, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x31, 0x33, 0x7d, 0x5b, 0x61, 0x2d,
	0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x27, 0x29, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xbc, 0x04, 0x0a, 0x24, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x52, 0x65, 0x67, 0x65,
	0x78, 0x4e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x12, 0x93, 0x04, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x54, 0x65, 0x73, 0x74, 0x42, 0xaf, 0x03, 0xba, 0x48, 0xab, 0x03, 0xba, 0x01, 0x5a, 0x0a,
	0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x62,
	0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x20, 0x31, 0x20, 0x74, 0x6f, 0x20, 0x31, 0x35, 0x20, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x1e, 0x74, 0x68, 0x69, 0x73, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x2e,
	0x7b, 0x31, 0x2c, 0x31, 0x35, 0x7d, 0x24, 0x27, 0x29, 0xba, 0x01, 0x6b, 0x0a, 0x0d, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x4f, 0x6e, 0x6c,
	0x79, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x20, 0x6c, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x73, 0x2c, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x61, 0x6e,
	0x64, 0x20, 0x68, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x20, 0x61, 0x72, 0x65, 0x20, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x1a, 0x21, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d,
	0x39, 0x2d, 0x5d, 0x2b, 0x24, 0x27, 0x29, 0xba, 0x01, 0x51, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x4d, 0x75, 0x73, 0x74, 0x20,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x63, 0x61, 0x73, 0x65, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x1e, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27,
	0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x2e, 0x2a, 0x24, 0x27, 0x29, 0xba, 0x01, 0x45, 0x0a, 0x0d,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x4d,
	0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68,
	0x20, 0x68, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x1a, 0x1a, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5b, 0x5e, 0x2d, 0x5d,
	0x24, 0x27, 0x29, 0xba, 0x01, 0x41, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x4d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x27, 0x2e, 0x2e, 0x27, 0x1a, 0x19, 0x21, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x73, 0x28, 0x27, 0x2e, 0x2e, 0x27, 0x29, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x32, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x54, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0xc4, 0x03, 0x0a, 0x45, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x67, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x43, 0x54, 0x52,
	0x4d, 0xaa, 0x02, 0x37, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xca, 0x02, 0x37, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x54, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xe2, 0x02, 0x43, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a,
	0x3a, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x3a, 0x3a, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDescData = file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_goTypes = []interface{}{
	(*ResourceIdExtractionTestWithOutIdPrefix)(nil), // 0: cloud.planton.apis.v1.commons.testing.resource.metadata.ResourceIdExtractionTestWithOutIdPrefix
	(*ResourceIdExtractionTestWithIdPrefix)(nil),    // 1: cloud.planton.apis.v1.commons.testing.resource.metadata.ResourceIdExtractionTestWithIdPrefix
	(*RegexNameFieldProtoValidateTest)(nil),         // 2: cloud.planton.apis.v1.commons.testing.resource.metadata.RegexNameFieldProtoValidateTest
	(*SplitRegexNameFieldProtoValidateTest)(nil),    // 3: cloud.planton.apis.v1.commons.testing.resource.metadata.SplitRegexNameFieldProtoValidateTest
	(*MetadataTest)(nil),                            // 4: cloud.planton.apis.v1.commons.testing.resource.metadata.MetadataTest
}
var file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.v1.commons.testing.resource.metadata.RegexNameFieldProtoValidateTest.metadata:type_name -> cloud.planton.apis.v1.commons.testing.resource.metadata.MetadataTest
	4, // 1: cloud.planton.apis.v1.commons.testing.resource.metadata.SplitRegexNameFieldProtoValidateTest.metadata:type_name -> cloud.planton.apis.v1.commons.testing.resource.metadata.MetadataTest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_init() }
func file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_init() {
	if File_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceIdExtractionTestWithOutIdPrefix); i {
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
		file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceIdExtractionTestWithIdPrefix); i {
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
		file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegexNameFieldProtoValidateTest); i {
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
		file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitRegexNameFieldProtoValidateTest); i {
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
		file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataTest); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto = out.File
	file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_testing_resource_metadata_model_proto_depIdxs = nil
}
