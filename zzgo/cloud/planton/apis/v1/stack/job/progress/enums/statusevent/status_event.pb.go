// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/stack/job/progress/enums/statusevent/status_event.proto

package statusevent

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

type StackJobProgressStatusEventType int32

const (
	StackJobProgressStatusEventType_STACK_JOB_PROGRESS_STATUS_EVENT_TYPE_UNSPECIFIED StackJobProgressStatusEventType = 0
	StackJobProgressStatusEventType_job_started                                      StackJobProgressStatusEventType = 1
	StackJobProgressStatusEventType_job_failed                                       StackJobProgressStatusEventType = 2
	StackJobProgressStatusEventType_job_succeeded                                    StackJobProgressStatusEventType = 3
	StackJobProgressStatusEventType_refresh_started                                  StackJobProgressStatusEventType = 4
	StackJobProgressStatusEventType_refresh_errors_reported                          StackJobProgressStatusEventType = 5
	StackJobProgressStatusEventType_refresh_succeeded                                StackJobProgressStatusEventType = 6
	StackJobProgressStatusEventType_refresh_failed                                   StackJobProgressStatusEventType = 7
	StackJobProgressStatusEventType_apply_preview_started                            StackJobProgressStatusEventType = 8
	StackJobProgressStatusEventType_apply_preview_errors_reported                    StackJobProgressStatusEventType = 9
	StackJobProgressStatusEventType_apply_preview_succeeded                          StackJobProgressStatusEventType = 10
	StackJobProgressStatusEventType_apply_preview_failed                             StackJobProgressStatusEventType = 11
	StackJobProgressStatusEventType_apply_started                                    StackJobProgressStatusEventType = 12
	StackJobProgressStatusEventType_apply_errors_reported                            StackJobProgressStatusEventType = 13
	StackJobProgressStatusEventType_apply_succeeded                                  StackJobProgressStatusEventType = 14
	StackJobProgressStatusEventType_apply_failed                                     StackJobProgressStatusEventType = 15
	StackJobProgressStatusEventType_destroy_preview_started                          StackJobProgressStatusEventType = 16
	StackJobProgressStatusEventType_destroy_preview_errors_reported                  StackJobProgressStatusEventType = 17
	StackJobProgressStatusEventType_destroy_preview_succeeded                        StackJobProgressStatusEventType = 18
	StackJobProgressStatusEventType_destroy_preview_failed                           StackJobProgressStatusEventType = 19
	StackJobProgressStatusEventType_destroy_started                                  StackJobProgressStatusEventType = 20
	StackJobProgressStatusEventType_destroy_errors_reported                          StackJobProgressStatusEventType = 21
	StackJobProgressStatusEventType_destroy_succeeded                                StackJobProgressStatusEventType = 22
	StackJobProgressStatusEventType_destroy_failed                                   StackJobProgressStatusEventType = 23
)

// Enum value maps for StackJobProgressStatusEventType.
var (
	StackJobProgressStatusEventType_name = map[int32]string{
		0:  "STACK_JOB_PROGRESS_STATUS_EVENT_TYPE_UNSPECIFIED",
		1:  "job_started",
		2:  "job_failed",
		3:  "job_succeeded",
		4:  "refresh_started",
		5:  "refresh_errors_reported",
		6:  "refresh_succeeded",
		7:  "refresh_failed",
		8:  "apply_preview_started",
		9:  "apply_preview_errors_reported",
		10: "apply_preview_succeeded",
		11: "apply_preview_failed",
		12: "apply_started",
		13: "apply_errors_reported",
		14: "apply_succeeded",
		15: "apply_failed",
		16: "destroy_preview_started",
		17: "destroy_preview_errors_reported",
		18: "destroy_preview_succeeded",
		19: "destroy_preview_failed",
		20: "destroy_started",
		21: "destroy_errors_reported",
		22: "destroy_succeeded",
		23: "destroy_failed",
	}
	StackJobProgressStatusEventType_value = map[string]int32{
		"STACK_JOB_PROGRESS_STATUS_EVENT_TYPE_UNSPECIFIED": 0,
		"job_started":                     1,
		"job_failed":                      2,
		"job_succeeded":                   3,
		"refresh_started":                 4,
		"refresh_errors_reported":         5,
		"refresh_succeeded":               6,
		"refresh_failed":                  7,
		"apply_preview_started":           8,
		"apply_preview_errors_reported":   9,
		"apply_preview_succeeded":         10,
		"apply_preview_failed":            11,
		"apply_started":                   12,
		"apply_errors_reported":           13,
		"apply_succeeded":                 14,
		"apply_failed":                    15,
		"destroy_preview_started":         16,
		"destroy_preview_errors_reported": 17,
		"destroy_preview_succeeded":       18,
		"destroy_preview_failed":          19,
		"destroy_started":                 20,
		"destroy_errors_reported":         21,
		"destroy_succeeded":               22,
		"destroy_failed":                  23,
	}
)

func (x StackJobProgressStatusEventType) Enum() *StackJobProgressStatusEventType {
	p := new(StackJobProgressStatusEventType)
	*p = x
	return p
}

func (x StackJobProgressStatusEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StackJobProgressStatusEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_enumTypes[0].Descriptor()
}

func (StackJobProgressStatusEventType) Type() protoreflect.EnumType {
	return &file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_enumTypes[0]
}

func (x StackJobProgressStatusEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StackJobProgressStatusEventType.Descriptor instead.
func (StackJobProgressStatusEventType) EnumDescriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_rawDescGZIP(), []int{0}
}

var File_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2a, 0x8c, 0x05, 0x0a, 0x1f,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x34, 0x0a, 0x30, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x50, 0x52, 0x4f,
	0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x10, 0x04, 0x12, 0x1b,
	0x0a, 0x17, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64,
	0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x10, 0x07, 0x12, 0x19, 0x0a, 0x15, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x10,
	0x08, 0x12, 0x21, 0x0a, 0x1d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10,
	0x0a, 0x12, 0x18, 0x0a, 0x14, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x0b, 0x12, 0x11, 0x0a, 0x0d, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x10, 0x0c, 0x12, 0x19,
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x10, 0x0d, 0x12, 0x13, 0x0a, 0x0f, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10, 0x0e, 0x12, 0x10,
	0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x0f,
	0x12, 0x1b, 0x0a, 0x17, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x10, 0x10, 0x12, 0x23, 0x0a,
	0x1f, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x10, 0x11, 0x12, 0x1d, 0x0a, 0x19, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10,
	0x12, 0x12, 0x1a, 0x0a, 0x16, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x13, 0x12, 0x13, 0x0a,
	0x0f, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x10, 0x14, 0x12, 0x1b, 0x0a, 0x17, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x10, 0x15, 0x12,
	0x15, 0x0a, 0x11, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x65, 0x64, 0x65, 0x64, 0x10, 0x16, 0x12, 0x12, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f,
	0x79, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x17, 0x42, 0xde, 0x03, 0x0a, 0x48, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x10, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x74, 0xa2, 0x02, 0x09, 0x43, 0x50, 0x41, 0x56, 0x53, 0x4a,
	0x50, 0x45, 0x53, 0xaa, 0x02, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0xca, 0x02, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4a,
	0x6f, 0x62, 0x5c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5c, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x74, 0xe2, 0x02, 0x46,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x4a, 0x6f, 0x62, 0x5c,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x42, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x4a, 0x6f, 0x62, 0x3a, 0x3a, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_rawDescData = file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_rawDesc
)

func file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_rawDescData
}

var file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_goTypes = []interface{}{
	(StackJobProgressStatusEventType)(0), // 0: cloud.planton.apis.v1.stack.job.progress.enums.statusevent.StackJobProgressStatusEventType
}
var file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_init()
}
func file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_init() {
	if File_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_depIdxs,
		EnumInfos:         file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_enumTypes,
	}.Build()
	File_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto = out.File
	file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_rawDesc = nil
	file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_goTypes = nil
	file_cloud_planton_apis_v1_stack_job_progress_enums_statusevent_status_event_proto_depIdxs = nil
}
