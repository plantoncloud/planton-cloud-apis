// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/iam/identity/connection/io.proto

package connection

import (
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
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

// wrapper to get list of identity connections
type IdentityConnections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*IdentityConnection `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *IdentityConnections) Reset() {
	*x = IdentityConnections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityConnections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityConnections) ProtoMessage() {}

func (x *IdentityConnections) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityConnections.ProtoReflect.Descriptor instead.
func (*IdentityConnections) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDescGZIP(), []int{0}
}

func (x *IdentityConnections) GetEntries() []*IdentityConnection {
	if x != nil {
		return x.Entries
	}
	return nil
}

// wrapper for user identity connection id.
type IdentityConnectionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IdentityConnectionId) Reset() {
	*x = IdentityConnectionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityConnectionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityConnectionId) ProtoMessage() {}

func (x *IdentityConnectionId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityConnectionId.ProtoReflect.Descriptor instead.
func (*IdentityConnectionId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDescGZIP(), []int{1}
}

func (x *IdentityConnectionId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// input for paginated queries that require identity account id as input.
type ListWithIdentityConnectionIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// identity connection id
	IdentityConnectionId string               `protobuf:"bytes,1,opt,name=identity_connection_id,json=identityConnectionId,proto3" json:"identity_connection_id,omitempty"`
	Page                 *pagination.PageInfo `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListWithIdentityConnectionIdReq) Reset() {
	*x = ListWithIdentityConnectionIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWithIdentityConnectionIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWithIdentityConnectionIdReq) ProtoMessage() {}

func (x *ListWithIdentityConnectionIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWithIdentityConnectionIdReq.ProtoReflect.Descriptor instead.
func (*ListWithIdentityConnectionIdReq) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDescGZIP(), []int{2}
}

func (x *ListWithIdentityConnectionIdReq) GetIdentityConnectionId() string {
	if x != nil {
		return x.IdentityConnectionId
	}
	return ""
}

func (x *ListWithIdentityConnectionIdReq) GetPage() *pagination.PageInfo {
	if x != nil {
		return x.Page
	}
	return nil
}

type IdentityConnectionCompanyId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IdentityConnectionCompanyId) Reset() {
	*x = IdentityConnectionCompanyId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityConnectionCompanyId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityConnectionCompanyId) ProtoMessage() {}

func (x *IdentityConnectionCompanyId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityConnectionCompanyId.ProtoReflect.Descriptor instead.
func (*IdentityConnectionCompanyId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDescGZIP(), []int{3}
}

func (x *IdentityConnectionCompanyId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_cloud_planton_apis_v1_iam_identity_connection_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x13, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x5b, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x2c, 0x0a,
	0x14, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x1f,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12,
	0x34, 0x0a, 0x16, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x33, 0x0a,
	0x1b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x83, 0x03, 0x0a, 0x3b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x07, 0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x07, 0x43,
	0x50, 0x41, 0x56, 0x49, 0x49, 0x43, 0xaa, 0x02, 0x2d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x49,
	0x61, 0x6d, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xca, 0x02, 0x2d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49,
	0x61, 0x6d, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x39, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49,
	0x61, 0x6d, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x33, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x49,
	0x61, 0x6d, 0x3a, 0x3a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDescData = file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDesc
)

func file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDescData
}

var file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_planton_apis_v1_iam_identity_connection_io_proto_goTypes = []interface{}{
	(*IdentityConnections)(nil),             // 0: cloud.planton.apis.v1.iam.identity.connection.IdentityConnections
	(*IdentityConnectionId)(nil),            // 1: cloud.planton.apis.v1.iam.identity.connection.IdentityConnectionId
	(*ListWithIdentityConnectionIdReq)(nil), // 2: cloud.planton.apis.v1.iam.identity.connection.ListWithIdentityConnectionIdReq
	(*IdentityConnectionCompanyId)(nil),     // 3: cloud.planton.apis.v1.iam.identity.connection.IdentityConnectionCompanyId
	(*IdentityConnection)(nil),              // 4: cloud.planton.apis.v1.iam.identity.connection.IdentityConnection
	(*pagination.PageInfo)(nil),             // 5: cloud.planton.apis.v1.commons.pagination.PageInfo
}
var file_cloud_planton_apis_v1_iam_identity_connection_io_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.v1.iam.identity.connection.IdentityConnections.entries:type_name -> cloud.planton.apis.v1.iam.identity.connection.IdentityConnection
	5, // 1: cloud.planton.apis.v1.iam.identity.connection.ListWithIdentityConnectionIdReq.page:type_name -> cloud.planton.apis.v1.commons.pagination.PageInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_iam_identity_connection_io_proto_init() }
func file_cloud_planton_apis_v1_iam_identity_connection_io_proto_init() {
	if File_cloud_planton_apis_v1_iam_identity_connection_io_proto != nil {
		return
	}
	file_cloud_planton_apis_v1_iam_identity_connection_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityConnections); i {
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
		file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityConnectionId); i {
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
		file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWithIdentityConnectionIdReq); i {
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
		file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityConnectionCompanyId); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_iam_identity_connection_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_iam_identity_connection_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_iam_identity_connection_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_iam_identity_connection_io_proto = out.File
	file_cloud_planton_apis_v1_iam_identity_connection_io_proto_rawDesc = nil
	file_cloud_planton_apis_v1_iam_identity_connection_io_proto_goTypes = nil
	file_cloud_planton_apis_v1_iam_identity_connection_io_proto_depIdxs = nil
}
