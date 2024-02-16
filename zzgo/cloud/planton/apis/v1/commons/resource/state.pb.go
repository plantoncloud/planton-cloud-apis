// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/commons/resource/state.proto

package resource

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

// metadata for all resource objects on planton-cloud
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the resource
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// id of the resource
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// labels for the resource
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// metadata for resource version
	Version *MetadataResourceVersion `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_resource_state_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Metadata) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Metadata) GetVersion() *MetadataResourceVersion {
	if x != nil {
		return x.Version
	}
	return nil
}

// resource lifecycle
type ResourceLifecycle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// flag to indicate if the resource is not deleted and is active
	IsActive bool `protobuf:"varint,1,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *ResourceLifecycle) Reset() {
	*x = ResourceLifecycle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceLifecycle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceLifecycle) ProtoMessage() {}

func (x *ResourceLifecycle) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceLifecycle.ProtoReflect.Descriptor instead.
func (*ResourceLifecycle) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_resource_state_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceLifecycle) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

// runnable resource lifecycle
type RunnableResourceLifecycle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// flag to indicate if the resource is not deleted and is active
	IsActive bool `protobuf:"varint,1,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// flag to indicate if the resource is paused
	IsPaused bool `protobuf:"varint,2,opt,name=is_paused,json=isPaused,proto3" json:"is_paused,omitempty"`
}

func (x *RunnableResourceLifecycle) Reset() {
	*x = RunnableResourceLifecycle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunnableResourceLifecycle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunnableResourceLifecycle) ProtoMessage() {}

func (x *RunnableResourceLifecycle) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunnableResourceLifecycle.ProtoReflect.Descriptor instead.
func (*RunnableResourceLifecycle) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_resource_state_proto_rawDescGZIP(), []int{2}
}

func (x *RunnableResourceLifecycle) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *RunnableResourceLifecycle) GetIsPaused() bool {
	if x != nil {
		return x.IsPaused
	}
	return false
}

// version information for api-resource
type MetadataResourceVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MetadataResourceVersion) Reset() {
	*x = MetadataResourceVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataResourceVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataResourceVersion) ProtoMessage() {}

func (x *MetadataResourceVersion) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataResourceVersion.ProtoReflect.Descriptor instead.
func (*MetadataResourceVersion) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_resource_state_proto_rawDescGZIP(), []int{3}
}

func (x *MetadataResourceVersion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MetadataResourceVersion) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_cloud_planton_apis_v1_commons_resource_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_resource_state_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x51, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa0, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x54, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x59, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x30, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x22, 0x55, 0x0a, 0x19, 0x52, 0x75, 0x6e, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x61, 0x75, 0x73, 0x65, 0x64, 0x22, 0x53, 0x0a, 0x17, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xc0, 0xb8, 0x18, 0x00, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0xda, 0x02, 0x0a, 0x34, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67,
	0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xa2,
	0x02, 0x06, 0x43, 0x50, 0x41, 0x56, 0x43, 0x52, 0xaa, 0x02, 0x26, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0xca, 0x02, 0x26, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xe2, 0x02, 0x32, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x2b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_resource_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_resource_state_proto_rawDescData = file_cloud_planton_apis_v1_commons_resource_state_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_resource_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_resource_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_resource_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_resource_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_resource_state_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloud_planton_apis_v1_commons_resource_state_proto_goTypes = []interface{}{
	(*Metadata)(nil),                  // 0: cloud.planton.apis.v1.commons.resource.Metadata
	(*ResourceLifecycle)(nil),         // 1: cloud.planton.apis.v1.commons.resource.ResourceLifecycle
	(*RunnableResourceLifecycle)(nil), // 2: cloud.planton.apis.v1.commons.resource.RunnableResourceLifecycle
	(*MetadataResourceVersion)(nil),   // 3: cloud.planton.apis.v1.commons.resource.MetadataResourceVersion
	nil,                               // 4: cloud.planton.apis.v1.commons.resource.Metadata.LabelsEntry
}
var file_cloud_planton_apis_v1_commons_resource_state_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.v1.commons.resource.Metadata.labels:type_name -> cloud.planton.apis.v1.commons.resource.Metadata.LabelsEntry
	3, // 1: cloud.planton.apis.v1.commons.resource.Metadata.version:type_name -> cloud.planton.apis.v1.commons.resource.MetadataResourceVersion
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_resource_state_proto_init() }
func file_cloud_planton_apis_v1_commons_resource_state_proto_init() {
	if File_cloud_planton_apis_v1_commons_resource_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceLifecycle); i {
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
		file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunnableResourceLifecycle); i {
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
		file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataResourceVersion); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_commons_resource_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_resource_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_resource_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_commons_resource_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_resource_state_proto = out.File
	file_cloud_planton_apis_v1_commons_resource_state_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_resource_state_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_resource_state_proto_depIdxs = nil
}
