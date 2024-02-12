// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/stack/job/progress/model.proto

package progress

import (
	operationtype "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job/enums/operationtype"
	event "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job/progress/enums/event"
	statusevent "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job/progress/enums/statusevent"
	engine "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/pulumi/engine"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// stack-job progress event
type StackJobProgressEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of the event
	EventType event.StackJobProgressEventType `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=cloud.planton.apis.v1.stack.job.progress.enums.event.StackJobProgressEventType" json:"event_type,omitempty"`
	// stack-job operation for which generated progress event
	OperationType operationtype.StackJobOperationType `protobuf:"varint,2,opt,name=operation_type,json=operationType,proto3,enum=cloud.planton.apis.v1.stack.job.enums.operationtype.StackJobOperationType" json:"operation_type,omitempty"`
	// stack-job progress status is populated when event_type is "STACK_JOB_PROGRESS_EVENT_TYPE_STATUS_EVENT"
	ProgressStatusPayload *StackJobProgressStatusPayload `protobuf:"bytes,3,opt,name=progress_status_payload,json=progressStatusPayload,proto3" json:"progress_status_payload,omitempty"`
	// pulumi engine event is only populated when event_type is "STACK_JOB_PROGRESS_EVENT_TYPE_PULUMI_ENGINE_EVENT"
	EngineEventPayload *engine.EngineEvent `protobuf:"bytes,4,opt,name=engine_event_payload,json=engineEventPayload,proto3" json:"engine_event_payload,omitempty"`
	// The time when the progress event was first received and written to persistent storage.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *StackJobProgressEvent) Reset() {
	*x = StackJobProgressEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_stack_job_progress_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackJobProgressEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackJobProgressEvent) ProtoMessage() {}

func (x *StackJobProgressEvent) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_stack_job_progress_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackJobProgressEvent.ProtoReflect.Descriptor instead.
func (*StackJobProgressEvent) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDescGZIP(), []int{0}
}

func (x *StackJobProgressEvent) GetEventType() event.StackJobProgressEventType {
	if x != nil {
		return x.EventType
	}
	return event.StackJobProgressEventType(0)
}

func (x *StackJobProgressEvent) GetOperationType() operationtype.StackJobOperationType {
	if x != nil {
		return x.OperationType
	}
	return operationtype.StackJobOperationType(0)
}

func (x *StackJobProgressEvent) GetProgressStatusPayload() *StackJobProgressStatusPayload {
	if x != nil {
		return x.ProgressStatusPayload
	}
	return nil
}

func (x *StackJobProgressEvent) GetEngineEventPayload() *engine.EngineEvent {
	if x != nil {
		return x.EngineEventPayload
	}
	return nil
}

func (x *StackJobProgressEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// stack-job progress status payload
type StackJobProgressStatusPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of the stack-job progress status event
	EventType statusevent.StackJobProgressStatusEventType `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=cloud.planton.apis.v1.stack.job.progress.enums.statusevent.StackJobProgressStatusEventType" json:"event_type,omitempty"`
	// errors is populated for errors_reported or failed event types
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *StackJobProgressStatusPayload) Reset() {
	*x = StackJobProgressStatusPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_stack_job_progress_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackJobProgressStatusPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackJobProgressStatusPayload) ProtoMessage() {}

func (x *StackJobProgressStatusPayload) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_stack_job_progress_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackJobProgressStatusPayload.ProtoReflect.Descriptor instead.
func (*StackJobProgressStatusPayload) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDescGZIP(), []int{1}
}

func (x *StackJobProgressStatusPayload) GetEventType() statusevent.StackJobProgressStatusEventType {
	if x != nil {
		return x.EventType
	}
	return statusevent.StackJobProgressStatusEventType(0)
}

func (x *StackJobProgressStatusPayload) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_cloud_planton_apis_v1_stack_job_progress_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x1a, 0x48, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f,
	0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x04, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x6e, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x4f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f,
	0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x71, 0x0a,
	0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x4a, 0x6f, 0x62, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x7f, 0x0a, 0x17, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x47, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x15, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x68, 0x0a, 0x14, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x75,
	0x6c, 0x75, 0x6d, 0x69, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x12, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xb3, 0x01, 0x0a, 0x1d, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a,
	0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x7a, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5b, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f,
	0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0xe8, 0x02, 0x0a, 0x36,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0xa2, 0x02,
	0x07, 0x43, 0x50, 0x41, 0x56, 0x53, 0x4a, 0x50, 0xaa, 0x02, 0x28, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0xca, 0x02, 0x28, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x5c, 0x4a, 0x6f, 0x62, 0x5c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0xe2, 0x02,
	0x34, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41,
	0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4a, 0x6f, 0x62,
	0x5c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x4a, 0x6f, 0x62, 0x3a, 0x3a, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDescData = file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_stack_job_progress_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_planton_apis_v1_stack_job_progress_model_proto_goTypes = []interface{}{
	(*StackJobProgressEvent)(nil),                    // 0: cloud.planton.apis.v1.stack.job.progress.StackJobProgressEvent
	(*StackJobProgressStatusPayload)(nil),            // 1: cloud.planton.apis.v1.stack.job.progress.StackJobProgressStatusPayload
	(event.StackJobProgressEventType)(0),             // 2: cloud.planton.apis.v1.stack.job.progress.enums.event.StackJobProgressEventType
	(operationtype.StackJobOperationType)(0),         // 3: cloud.planton.apis.v1.stack.job.enums.operationtype.StackJobOperationType
	(*engine.EngineEvent)(nil),                       // 4: cloud.planton.apis.v1.stack.pulumi.engine.EngineEvent
	(*timestamppb.Timestamp)(nil),                    // 5: google.protobuf.Timestamp
	(statusevent.StackJobProgressStatusEventType)(0), // 6: cloud.planton.apis.v1.stack.job.progress.enums.statusevent.StackJobProgressStatusEventType
}
var file_cloud_planton_apis_v1_stack_job_progress_model_proto_depIdxs = []int32{
	2, // 0: cloud.planton.apis.v1.stack.job.progress.StackJobProgressEvent.event_type:type_name -> cloud.planton.apis.v1.stack.job.progress.enums.event.StackJobProgressEventType
	3, // 1: cloud.planton.apis.v1.stack.job.progress.StackJobProgressEvent.operation_type:type_name -> cloud.planton.apis.v1.stack.job.enums.operationtype.StackJobOperationType
	1, // 2: cloud.planton.apis.v1.stack.job.progress.StackJobProgressEvent.progress_status_payload:type_name -> cloud.planton.apis.v1.stack.job.progress.StackJobProgressStatusPayload
	4, // 3: cloud.planton.apis.v1.stack.job.progress.StackJobProgressEvent.engine_event_payload:type_name -> cloud.planton.apis.v1.stack.pulumi.engine.EngineEvent
	5, // 4: cloud.planton.apis.v1.stack.job.progress.StackJobProgressEvent.timestamp:type_name -> google.protobuf.Timestamp
	6, // 5: cloud.planton.apis.v1.stack.job.progress.StackJobProgressStatusPayload.event_type:type_name -> cloud.planton.apis.v1.stack.job.progress.enums.statusevent.StackJobProgressStatusEventType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_stack_job_progress_model_proto_init() }
func file_cloud_planton_apis_v1_stack_job_progress_model_proto_init() {
	if File_cloud_planton_apis_v1_stack_job_progress_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_stack_job_progress_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackJobProgressEvent); i {
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
		file_cloud_planton_apis_v1_stack_job_progress_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackJobProgressStatusPayload); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_stack_job_progress_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_stack_job_progress_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_stack_job_progress_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_stack_job_progress_model_proto = out.File
	file_cloud_planton_apis_v1_stack_job_progress_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_stack_job_progress_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_stack_job_progress_model_proto_depIdxs = nil
}
