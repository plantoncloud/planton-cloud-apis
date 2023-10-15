// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/stack/state/model.proto

package state

import (
	audit "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/audit"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/enums"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/rpc/enums"
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

// stack-job state
type StackJobState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// system-audit information
	SysAudit *audit.SysAudit `protobuf:"bytes,99,opt,name=sys_audit,json=sysAudit,proto3" json:"sys_audit,omitempty"`
	// unique identifier(uuid) for stack-job
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// id of the company to which the stack-job belongs to
	CompanyId string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// id of the product to which the stack-job belongs to
	// note: this attribute is always empty for company level stacks
	ProductId string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// id of the product-env to which the stack-job belongs to
	// note: this attribute is always empty for company level stacks
	EnvironmentId string `protobuf:"bytes,4,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	// resource type for which the stack is executed
	ResourceType string `protobuf:"bytes,5,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// id of the resource for which the stack is executed
	ResourceId string `protobuf:"bytes,6,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// name of the stack
	StackName string `protobuf:"bytes,7,opt,name=stack_name,json=stackName,proto3" json:"stack_name,omitempty"`
	// (required) stack operation type
	OperationType string `protobuf:"bytes,8,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
	// path of the log file in the storage bucket
	LogFile string `protobuf:"bytes,9,opt,name=log_file,json=logFile,proto3" json:"log_file,omitempty"`
	// status of the stack-job
	Status string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	// url of the stack-job on pulumi web console.
	// note: this value is not persisted in the database.
	// the value of this attributes is rendered by replacing the placeholders in
	// "https://app.pulumi.com/${pulumiOrgName}/${pulumiProject}/${pulumiStackName}"
	// value of pulumiOrgName is same for every single stack for each planton-cloud environment.
	// value of pulumiProject is the company-id on planton-cloud.
	// value of pulumiStack is stack_name attribute in this object.
	// ex: https://app.pulumi.com/plantonstack-prod/cookie/afs-cookie-shop-main.ca-planton-hosting-gcp-main.artifact-store
	WebUrl string `protobuf:"bytes,11,opt,name=web_url,json=webUrl,proto3" json:"web_url,omitempty"`
	// content of the log file.
	// note: log content is not persisted in the database.
	// log-content will only be included in response for stack-job details rpc calls.
	LogContent string `protobuf:"bytes,12,opt,name=log_content,json=logContent,proto3" json:"log_content,omitempty"`
	// error message from stack-job. this attribute is only populated when stack-job-status is "failed"
	ErrorMessage string `protobuf:"bytes,13,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// resource-count
	ResourceCount int32 `protobuf:"varint,14,opt,name=resource_count,json=resourceCount,proto3" json:"resource_count,omitempty"`
}

func (x *StackJobState) Reset() {
	*x = StackJobState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_stack_state_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackJobState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackJobState) ProtoMessage() {}

func (x *StackJobState) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_stack_state_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackJobState.ProtoReflect.Descriptor instead.
func (*StackJobState) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_stack_state_model_proto_rawDescGZIP(), []int{0}
}

func (x *StackJobState) GetSysAudit() *audit.SysAudit {
	if x != nil {
		return x.SysAudit
	}
	return nil
}

func (x *StackJobState) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StackJobState) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *StackJobState) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *StackJobState) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *StackJobState) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *StackJobState) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *StackJobState) GetStackName() string {
	if x != nil {
		return x.StackName
	}
	return ""
}

func (x *StackJobState) GetOperationType() string {
	if x != nil {
		return x.OperationType
	}
	return ""
}

func (x *StackJobState) GetLogFile() string {
	if x != nil {
		return x.LogFile
	}
	return ""
}

func (x *StackJobState) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StackJobState) GetWebUrl() string {
	if x != nil {
		return x.WebUrl
	}
	return ""
}

func (x *StackJobState) GetLogContent() string {
	if x != nil {
		return x.LogContent
	}
	return ""
}

func (x *StackJobState) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *StackJobState) GetResourceCount() int32 {
	if x != nil {
		return x.ResourceCount
	}
	return 0
}

var File_cloud_planton_apis_v1_stack_state_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_stack_state_model_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x04, 0x0a, 0x0d, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4a, 0x0a, 0x09,
	0x73, 0x79, 0x73, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x08,
	0x73, 0x79, 0x73, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x77, 0x65, 0x62, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77,
	0x65, 0x62, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0xbc, 0x02, 0x0a, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0xa2, 0x02, 0x06, 0x43, 0x50, 0x41, 0x56, 0x53, 0x53,
	0xaa, 0x02, 0x21, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0xca, 0x02, 0x21, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x65, 0xe2, 0x02, 0x2d, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x26, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_stack_state_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_stack_state_model_proto_rawDescData = file_cloud_planton_apis_v1_stack_state_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_stack_state_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_stack_state_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_stack_state_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_stack_state_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_stack_state_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_stack_state_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloud_planton_apis_v1_stack_state_model_proto_goTypes = []interface{}{
	(*StackJobState)(nil),  // 0: cloud.planton.apis.v1.stack.state.StackJobState
	(*audit.SysAudit)(nil), // 1: cloud.planton.apis.v1.commons.audit.SysAudit
}
var file_cloud_planton_apis_v1_stack_state_model_proto_depIdxs = []int32{
	1, // 0: cloud.planton.apis.v1.stack.state.StackJobState.sys_audit:type_name -> cloud.planton.apis.v1.commons.audit.SysAudit
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_stack_state_model_proto_init() }
func file_cloud_planton_apis_v1_stack_state_model_proto_init() {
	if File_cloud_planton_apis_v1_stack_state_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_stack_state_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackJobState); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_stack_state_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_stack_state_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_stack_state_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_stack_state_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_stack_state_model_proto = out.File
	file_cloud_planton_apis_v1_stack_state_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_stack_state_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_stack_state_model_proto_depIdxs = nil
}
