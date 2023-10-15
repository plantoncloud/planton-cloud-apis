// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/iam/authz/permission/rpc/model.proto

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

// iam permission
type IamPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the iam permission
	PermissionId string `protobuf:"bytes,1,opt,name=permission_id,json=permissionId,proto3" json:"permission_id,omitempty"`
	// name of the iam permission
	PermissionName string `protobuf:"bytes,2,opt,name=permission_name,json=permissionName,proto3" json:"permission_name,omitempty"`
	// code of the permission used in FGA model
	PermissionCode string `protobuf:"bytes,3,opt,name=permission_code,json=permissionCode,proto3" json:"permission_code,omitempty"`
	// flag to indicate if the permission is deleted.
	IsActive bool `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *IamPermission) Reset() {
	*x = IamPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamPermission) ProtoMessage() {}

func (x *IamPermission) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamPermission.ProtoReflect.Descriptor instead.
func (*IamPermission) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDescGZIP(), []int{0}
}

func (x *IamPermission) GetPermissionId() string {
	if x != nil {
		return x.PermissionId
	}
	return ""
}

func (x *IamPermission) GetPermissionName() string {
	if x != nil {
		return x.PermissionName
	}
	return ""
}

func (x *IamPermission) GetPermissionCode() string {
	if x != nil {
		return x.PermissionCode
	}
	return ""
}

func (x *IamPermission) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

// wrapper for iam permission id
type IamPermissionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IamPermissionId) Reset() {
	*x = IamPermissionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamPermissionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamPermissionId) ProtoMessage() {}

func (x *IamPermissionId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamPermissionId.ProtoReflect.Descriptor instead.
func (*IamPermissionId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDescGZIP(), []int{1}
}

func (x *IamPermissionId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// wrapper for iam permission code
type IamPermissionCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IamPermissionCode) Reset() {
	*x = IamPermissionCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamPermissionCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamPermissionCode) ProtoMessage() {}

func (x *IamPermissionCode) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamPermissionCode.ProtoReflect.Descriptor instead.
func (*IamPermissionCode) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDescGZIP(), []int{2}
}

func (x *IamPermissionCode) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// list of iam permissions
type IamPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*IamPermission `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *IamPermissions) Reset() {
	*x = IamPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamPermissions) ProtoMessage() {}

func (x *IamPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamPermissions.ProtoReflect.Descriptor instead.
func (*IamPermissions) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDescGZIP(), []int{3}
}

func (x *IamPermissions) GetEntries() []*IamPermission {
	if x != nil {
		return x.Entries
	}
	return nil
}

// response for paginated rpc query to list iam permissions.
type IamPermissionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries    []*IamPermission `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	TotalPages int32            `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
}

func (x *IamPermissionList) Reset() {
	*x = IamPermissionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamPermissionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamPermissionList) ProtoMessage() {}

func (x *IamPermissionList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamPermissionList.ProtoReflect.Descriptor instead.
func (*IamPermissionList) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDescGZIP(), []int{4}
}

func (x *IamPermissionList) GetEntries() []*IamPermission {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *IamPermissionList) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

var File_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x22, 0xa3, 0x01, 0x0a,
	0x0d, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x49,
	0x61, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x69, 0x0a, 0x0e, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x57, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x61, 0x6d, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65,
	0x73, 0x42, 0x84, 0x03, 0x0a, 0x32, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0xa2, 0x02, 0x08, 0x43, 0x50, 0x41, 0x56, 0x49, 0x41, 0x50,
	0x52, 0xaa, 0x02, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x49, 0x61, 0x6d, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x70, 0x63, 0xca, 0x02, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x41,
	0x75, 0x74, 0x68, 0x7a, 0x5c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5c,
	0x52, 0x70, 0x63, 0xe2, 0x02, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x49, 0x61, 0x6d, 0x5c,
	0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5c, 0x52, 0x70, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x49, 0x61, 0x6d,
	0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x3a, 0x3a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDescData = file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_goTypes = []interface{}{
	(*IamPermission)(nil),     // 0: cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermission
	(*IamPermissionId)(nil),   // 1: cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionId
	(*IamPermissionCode)(nil), // 2: cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionCode
	(*IamPermissions)(nil),    // 3: cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissions
	(*IamPermissionList)(nil), // 4: cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionList
}
var file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissions.entries:type_name -> cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermission
	0, // 1: cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionList.entries:type_name -> cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermission
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_init() }
func file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_init() {
	if File_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamPermission); i {
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
		file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamPermissionId); i {
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
		file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamPermissionCode); i {
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
		file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamPermissions); i {
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
		file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamPermissionList); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto = out.File
	file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_iam_authz_permission_rpc_model_proto_depIdxs = nil
}
