// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/commons/testing/protobuf/model.proto

package protobuf

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

type TestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RootLevelString string    `protobuf:"bytes,1,opt,name=rootLevelString,proto3" json:"rootLevelString,omitempty"`
	LevelOne        *LevelOne `protobuf:"bytes,2,opt,name=levelOne,proto3" json:"levelOne,omitempty"`
}

func (x *TestMessage) Reset() {
	*x = TestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage) ProtoMessage() {}

func (x *TestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage.ProtoReflect.Descriptor instead.
func (*TestMessage) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDescGZIP(), []int{0}
}

func (x *TestMessage) GetRootLevelString() string {
	if x != nil {
		return x.RootLevelString
	}
	return ""
}

func (x *TestMessage) GetLevelOne() *LevelOne {
	if x != nil {
		return x.LevelOne
	}
	return nil
}

type LevelOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LevelOneString string    `protobuf:"bytes,1,opt,name=levelOneString,proto3" json:"levelOneString,omitempty"`
	LevelTwo       *LevelTwo `protobuf:"bytes,2,opt,name=levelTwo,proto3" json:"levelTwo,omitempty"`
}

func (x *LevelOne) Reset() {
	*x = LevelOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelOne) ProtoMessage() {}

func (x *LevelOne) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelOne.ProtoReflect.Descriptor instead.
func (*LevelOne) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDescGZIP(), []int{1}
}

func (x *LevelOne) GetLevelOneString() string {
	if x != nil {
		return x.LevelOneString
	}
	return ""
}

func (x *LevelOne) GetLevelTwo() *LevelTwo {
	if x != nil {
		return x.LevelTwo
	}
	return nil
}

type LevelTwo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LevelTwoString string      `protobuf:"bytes,1,opt,name=levelTwoString,proto3" json:"levelTwoString,omitempty"`
	LevelThree     *LevelThree `protobuf:"bytes,2,opt,name=levelThree,proto3" json:"levelThree,omitempty"`
}

func (x *LevelTwo) Reset() {
	*x = LevelTwo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelTwo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelTwo) ProtoMessage() {}

func (x *LevelTwo) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelTwo.ProtoReflect.Descriptor instead.
func (*LevelTwo) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDescGZIP(), []int{2}
}

func (x *LevelTwo) GetLevelTwoString() string {
	if x != nil {
		return x.LevelTwoString
	}
	return ""
}

func (x *LevelTwo) GetLevelThree() *LevelThree {
	if x != nil {
		return x.LevelThree
	}
	return nil
}

type LevelThree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LevelThreeString string `protobuf:"bytes,1,opt,name=levelThreeString,proto3" json:"levelThreeString,omitempty"`
	NotAString       int32  `protobuf:"varint,2,opt,name=notAString,proto3" json:"notAString,omitempty"`
}

func (x *LevelThree) Reset() {
	*x = LevelThree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelThree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelThree) ProtoMessage() {}

func (x *LevelThree) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelThree.ProtoReflect.Descriptor instead.
func (*LevelThree) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDescGZIP(), []int{3}
}

func (x *LevelThree) GetLevelThreeString() string {
	if x != nil {
		return x.LevelThreeString
	}
	return ""
}

func (x *LevelThree) GetNotAString() int32 {
	if x != nil {
		return x.NotAString
	}
	return 0
}

var File_cloud_planton_apis_v1_commons_testing_protobuf_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x8d, 0x01, 0x0a,
	0x0b, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x72, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x54, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x4f,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f,
	0x6e, 0x65, 0x52, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x6e, 0x65, 0x22, 0x88, 0x01, 0x0a,
	0x08, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x4f, 0x6e, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x6e, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x54, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x77, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x77, 0x6f, 0x52, 0x08, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x54, 0x77, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x08, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x54, 0x77, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x77, 0x6f,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x54, 0x77, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x5a, 0x0a, 0x0a,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x68, 0x72, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x68, 0x72, 0x65, 0x65, 0x52, 0x0a, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x54, 0x68, 0x72, 0x65, 0x65, 0x22, 0x58, 0x0a, 0x0a, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x54, 0x68, 0x72, 0x65, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x54,
	0x68, 0x72, 0x65, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x68, 0x72, 0x65, 0x65, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x41, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x41, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x42, 0x8c, 0x03, 0x0a, 0x3c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a,
	0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41, 0x56, 0x43, 0x54, 0x50, 0xaa, 0x02, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73,
	0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xca, 0x02, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69,
	0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x54, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xe2, 0x02, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x54, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x34, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a,
	0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDescData = file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_goTypes = []interface{}{
	(*TestMessage)(nil), // 0: cloud.planton.apis.v1.commons.testing.protobuf.TestMessage
	(*LevelOne)(nil),    // 1: cloud.planton.apis.v1.commons.testing.protobuf.LevelOne
	(*LevelTwo)(nil),    // 2: cloud.planton.apis.v1.commons.testing.protobuf.LevelTwo
	(*LevelThree)(nil),  // 3: cloud.planton.apis.v1.commons.testing.protobuf.LevelThree
}
var file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_depIdxs = []int32{
	1, // 0: cloud.planton.apis.v1.commons.testing.protobuf.TestMessage.levelOne:type_name -> cloud.planton.apis.v1.commons.testing.protobuf.LevelOne
	2, // 1: cloud.planton.apis.v1.commons.testing.protobuf.LevelOne.levelTwo:type_name -> cloud.planton.apis.v1.commons.testing.protobuf.LevelTwo
	3, // 2: cloud.planton.apis.v1.commons.testing.protobuf.LevelTwo.levelThree:type_name -> cloud.planton.apis.v1.commons.testing.protobuf.LevelThree
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_init() }
func file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_init() {
	if File_cloud_planton_apis_v1_commons_testing_protobuf_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage); i {
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
		file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelOne); i {
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
		file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelTwo); i {
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
		file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelThree); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_testing_protobuf_model_proto = out.File
	file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_testing_protobuf_model_proto_depIdxs = nil
}
