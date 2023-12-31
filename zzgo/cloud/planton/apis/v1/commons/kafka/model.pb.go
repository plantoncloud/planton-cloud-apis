// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/commons/kafka/model.proto

package kafka

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// this is used to get host information of the pod
// ip address and port of the pod
type HostInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *HostInfo) Reset() {
	*x = HostInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostInfo) ProtoMessage() {}

func (x *HostInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostInfo.ProtoReflect.Descriptor instead.
func (*HostInfo) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescGZIP(), []int{0}
}

func (x *HostInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *HostInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// this is used as input to pass state store name
type StoreName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreName string `protobuf:"bytes,1,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
}

func (x *StoreName) Reset() {
	*x = StoreName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreName) ProtoMessage() {}

func (x *StoreName) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreName.ProtoReflect.Descriptor instead.
func (*StoreName) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescGZIP(), []int{1}
}

func (x *StoreName) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

// this is used as get topic partition details
type TopicPartition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic     string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Partition int32  `protobuf:"varint,2,opt,name=partition,proto3" json:"partition,omitempty"`
}

func (x *TopicPartition) Reset() {
	*x = TopicPartition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicPartition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicPartition) ProtoMessage() {}

func (x *TopicPartition) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicPartition.ProtoReflect.Descriptor instead.
func (*TopicPartition) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescGZIP(), []int{2}
}

func (x *TopicPartition) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *TopicPartition) GetPartition() int32 {
	if x != nil {
		return x.Partition
	}
	return 0
}

// this is used to get meta information of kafka streams instance
type InstanceMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostInfo               *HostInfo         `protobuf:"bytes,1,opt,name=host_info,json=hostInfo,proto3" json:"host_info,omitempty"`
	StateStoreNames        []string          `protobuf:"bytes,2,rep,name=state_store_names,json=stateStoreNames,proto3" json:"state_store_names,omitempty"`
	TopicPartitions        []*TopicPartition `protobuf:"bytes,3,rep,name=topic_partitions,json=topicPartitions,proto3" json:"topic_partitions,omitempty"`
	StandbyStateStoreNames []string          `protobuf:"bytes,4,rep,name=standby_state_store_names,json=standbyStateStoreNames,proto3" json:"standby_state_store_names,omitempty"`
	StandbyTopicPartitions []*TopicPartition `protobuf:"bytes,5,rep,name=standby_topic_partitions,json=standbyTopicPartitions,proto3" json:"standby_topic_partitions,omitempty"`
}

func (x *InstanceMeta) Reset() {
	*x = InstanceMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceMeta) ProtoMessage() {}

func (x *InstanceMeta) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceMeta.ProtoReflect.Descriptor instead.
func (*InstanceMeta) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescGZIP(), []int{3}
}

func (x *InstanceMeta) GetHostInfo() *HostInfo {
	if x != nil {
		return x.HostInfo
	}
	return nil
}

func (x *InstanceMeta) GetStateStoreNames() []string {
	if x != nil {
		return x.StateStoreNames
	}
	return nil
}

func (x *InstanceMeta) GetTopicPartitions() []*TopicPartition {
	if x != nil {
		return x.TopicPartitions
	}
	return nil
}

func (x *InstanceMeta) GetStandbyStateStoreNames() []string {
	if x != nil {
		return x.StandbyStateStoreNames
	}
	return nil
}

func (x *InstanceMeta) GetStandbyTopicPartitions() []*TopicPartition {
	if x != nil {
		return x.StandbyTopicPartitions
	}
	return nil
}

// this is used to get meta information of kafka streams pod based on key
type KeyQueryMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostInfo     *HostInfo   `protobuf:"bytes,1,opt,name=host_info,json=hostInfo,proto3" json:"host_info,omitempty"`
	StandbyHosts []*HostInfo `protobuf:"bytes,2,rep,name=standby_hosts,json=standbyHosts,proto3" json:"standby_hosts,omitempty"`
	Partition    int32       `protobuf:"varint,3,opt,name=partition,proto3" json:"partition,omitempty"`
}

func (x *KeyQueryMeta) Reset() {
	*x = KeyQueryMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyQueryMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyQueryMeta) ProtoMessage() {}

func (x *KeyQueryMeta) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyQueryMeta.ProtoReflect.Descriptor instead.
func (*KeyQueryMeta) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescGZIP(), []int{4}
}

func (x *KeyQueryMeta) GetHostInfo() *HostInfo {
	if x != nil {
		return x.HostInfo
	}
	return nil
}

func (x *KeyQueryMeta) GetStandbyHosts() []*HostInfo {
	if x != nil {
		return x.StandbyHosts
	}
	return nil
}

func (x *KeyQueryMeta) GetPartition() int32 {
	if x != nil {
		return x.Partition
	}
	return 0
}

// this is used to get meta information of kafka streams pod
type StreamsMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceMeta []*InstanceMeta `protobuf:"bytes,1,rep,name=instance_meta,json=instanceMeta,proto3" json:"instance_meta,omitempty"`
}

func (x *StreamsMeta) Reset() {
	*x = StreamsMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamsMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamsMeta) ProtoMessage() {}

func (x *StreamsMeta) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamsMeta.ProtoReflect.Descriptor instead.
func (*StreamsMeta) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescGZIP(), []int{5}
}

func (x *StreamsMeta) GetInstanceMeta() []*InstanceMeta {
	if x != nil {
		return x.InstanceMeta
	}
	return nil
}

// this is used to pass uuid as key
type Uuid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Uuid) Reset() {
	*x = Uuid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uuid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uuid) ProtoMessage() {}

func (x *Uuid) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uuid.ProtoReflect.Descriptor instead.
func (*Uuid) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescGZIP(), []int{6}
}

func (x *Uuid) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_cloud_planton_apis_v1_commons_kafka_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x23, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x2a, 0x0a, 0x09, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x0e, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x90, 0x03,
	0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x4a,
	0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x5e, 0x0a, 0x10, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x62,
	0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x16, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x62, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x6d, 0x0a, 0x18, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x62, 0x79, 0x5f, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x62,
	0x79, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xcc, 0x01, 0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x4a, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x52, 0x0a,
	0x0d, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x62, 0x79, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x62, 0x79, 0x48, 0x6f, 0x73, 0x74,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x65, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x56,
	0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x1a, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x42, 0xc8, 0x02, 0x0a, 0x31, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0xa2, 0x02, 0x06, 0x43, 0x50,
	0x41, 0x56, 0x43, 0x4b, 0xaa, 0x02, 0x23, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0xca, 0x02, 0x23, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x5c, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0xe2, 0x02, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x5c, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x28, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescData = file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cloud_planton_apis_v1_commons_kafka_model_proto_goTypes = []interface{}{
	(*HostInfo)(nil),       // 0: cloud.planton.apis.v1.commons.kafka.HostInfo
	(*StoreName)(nil),      // 1: cloud.planton.apis.v1.commons.kafka.StoreName
	(*TopicPartition)(nil), // 2: cloud.planton.apis.v1.commons.kafka.TopicPartition
	(*InstanceMeta)(nil),   // 3: cloud.planton.apis.v1.commons.kafka.InstanceMeta
	(*KeyQueryMeta)(nil),   // 4: cloud.planton.apis.v1.commons.kafka.KeyQueryMeta
	(*StreamsMeta)(nil),    // 5: cloud.planton.apis.v1.commons.kafka.StreamsMeta
	(*Uuid)(nil),           // 6: cloud.planton.apis.v1.commons.kafka.Uuid
}
var file_cloud_planton_apis_v1_commons_kafka_model_proto_depIdxs = []int32{
	0, // 0: cloud.planton.apis.v1.commons.kafka.InstanceMeta.host_info:type_name -> cloud.planton.apis.v1.commons.kafka.HostInfo
	2, // 1: cloud.planton.apis.v1.commons.kafka.InstanceMeta.topic_partitions:type_name -> cloud.planton.apis.v1.commons.kafka.TopicPartition
	2, // 2: cloud.planton.apis.v1.commons.kafka.InstanceMeta.standby_topic_partitions:type_name -> cloud.planton.apis.v1.commons.kafka.TopicPartition
	0, // 3: cloud.planton.apis.v1.commons.kafka.KeyQueryMeta.host_info:type_name -> cloud.planton.apis.v1.commons.kafka.HostInfo
	0, // 4: cloud.planton.apis.v1.commons.kafka.KeyQueryMeta.standby_hosts:type_name -> cloud.planton.apis.v1.commons.kafka.HostInfo
	3, // 5: cloud.planton.apis.v1.commons.kafka.StreamsMeta.instance_meta:type_name -> cloud.planton.apis.v1.commons.kafka.InstanceMeta
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_commons_kafka_model_proto_init() }
func file_cloud_planton_apis_v1_commons_kafka_model_proto_init() {
	if File_cloud_planton_apis_v1_commons_kafka_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostInfo); i {
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
		file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreName); i {
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
		file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicPartition); i {
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
		file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceMeta); i {
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
		file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyQueryMeta); i {
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
		file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamsMeta); i {
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
		file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uuid); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_commons_kafka_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_commons_kafka_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_commons_kafka_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_commons_kafka_model_proto = out.File
	file_cloud_planton_apis_v1_commons_kafka_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_commons_kafka_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_commons_kafka_model_proto_depIdxs = nil
}
