// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/customendpoint/model/io.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// wrapper for custom-endpoint Domain id
type CustomEndpointId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CustomEndpointId) Reset() {
	*x = CustomEndpointId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomEndpointId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEndpointId) ProtoMessage() {}

func (x *CustomEndpointId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEndpointId.ProtoReflect.Descriptor instead.
func (*CustomEndpointId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescGZIP(), []int{0}
}

func (x *CustomEndpointId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// list of custom-endpoints
type CustomEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*CustomEndpoint `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *CustomEndpoints) Reset() {
	*x = CustomEndpoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEndpoints) ProtoMessage() {}

func (x *CustomEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEndpoints.ProtoReflect.Descriptor instead.
func (*CustomEndpoints) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescGZIP(), []int{1}
}

func (x *CustomEndpoints) GetEntries() []*CustomEndpoint {
	if x != nil {
		return x.Entries
	}
	return nil
}

// response for paginated query to list custom-endpoints
type CustomEndpointList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPages int32             `protobuf:"varint,1,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	Entries    []*CustomEndpoint `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *CustomEndpointList) Reset() {
	*x = CustomEndpointList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomEndpointList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEndpointList) ProtoMessage() {}

func (x *CustomEndpointList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEndpointList.ProtoReflect.Descriptor instead.
func (*CustomEndpointList) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescGZIP(), []int{2}
}

func (x *CustomEndpointList) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *CustomEndpointList) GetEntries() []*CustomEndpoint {
	if x != nil {
		return x.Entries
	}
	return nil
}

// AddOrUpdateCustomEndpointRouteCommandInput is used to encapsulate the details required
// for adding a new route to a specific custom-endpoint, or updating
// an existing one. This message is typically used to transmit data between the client and
// server during an add or update operation concerning a specific route
// associated with a particular custom-endpoint.
type AddOrUpdateCustomEndpointRouteCommandInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier for the custom-endpoint to which the route
	// needs to be added or updated. This field must be populated with a valid
	// custom-endpoint ID, which can be obtained from the custom-endpoint entity itself.
	// The server uses this ID to identify the correct custom-endpoint where the
	// route needs to be added or updated.
	CustomEndpointId string `protobuf:"bytes,1,opt,name=custom_endpoint_id,json=customEndpointId,proto3" json:"custom_endpoint_id,omitempty"`
	// The route that needs to be added or updated within the product
	// environment. This field should be populated with a valid CustomEndpointRoute object,
	// which encapsulates the details of the route. If an route
	// with the same url path prefix already exists in the custom-endpoint, the value will be updated.
	// Otherwise, a new route will be created with the provided details.
	Route *CustomEndpointRoute `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	// A descriptive message explaining the reason for the change or operation.
	// This is used for history logging purposes.
	VersionMessage string `protobuf:"bytes,3,opt,name=version_message,json=versionMessage,proto3" json:"version_message,omitempty"`
}

func (x *AddOrUpdateCustomEndpointRouteCommandInput) Reset() {
	*x = AddOrUpdateCustomEndpointRouteCommandInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrUpdateCustomEndpointRouteCommandInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrUpdateCustomEndpointRouteCommandInput) ProtoMessage() {}

func (x *AddOrUpdateCustomEndpointRouteCommandInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrUpdateCustomEndpointRouteCommandInput.ProtoReflect.Descriptor instead.
func (*AddOrUpdateCustomEndpointRouteCommandInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescGZIP(), []int{3}
}

func (x *AddOrUpdateCustomEndpointRouteCommandInput) GetCustomEndpointId() string {
	if x != nil {
		return x.CustomEndpointId
	}
	return ""
}

func (x *AddOrUpdateCustomEndpointRouteCommandInput) GetRoute() *CustomEndpointRoute {
	if x != nil {
		return x.Route
	}
	return nil
}

func (x *AddOrUpdateCustomEndpointRouteCommandInput) GetVersionMessage() string {
	if x != nil {
		return x.VersionMessage
	}
	return ""
}

// DeleteOrRestoreCustomEndpointRouteCommandInput is used to encapsulate the details required for
// deleting or restore a route of a specific custom-endpoint.
// This message is typically used to transmit data between the client and the server
// during a delete or restore operation concerning a specific route associated
// with a particular custom-endpoint.
type DeleteOrRestoreCustomEndpointRouteCommandInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier for the custom-endpoint from which the route
	// needs to be deleted. This field must be populated with a valid
	// custom-endpoint ID, which can be obtained from the custom-endpoint entity itself.
	// The server uses this ID to identify the correct custom-endpoint from which
	// the route needs to be deleted.
	CustomEndpointId string `protobuf:"bytes,1,opt,name=custom_endpoint_id,json=customEndpointId,proto3" json:"custom_endpoint_id,omitempty"`
	// The url path prefix of the route that needs to be deleted
	// from the custom-endpoint. This field should be populated with a valid
	// url path prefix, which can be obtained from the route entity itself.
	// The server uses this url path prefix to identify the correct route that
	// needs to be deleted.
	UrlPathPrefix string `protobuf:"bytes,2,opt,name=url_path_prefix,json=urlPathPrefix,proto3" json:"url_path_prefix,omitempty"`
	// A descriptive message explaining the reason for the change or operation.
	// This is used for history logging purposes.
	VersionMessage string `protobuf:"bytes,3,opt,name=version_message,json=versionMessage,proto3" json:"version_message,omitempty"`
}

func (x *DeleteOrRestoreCustomEndpointRouteCommandInput) Reset() {
	*x = DeleteOrRestoreCustomEndpointRouteCommandInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrRestoreCustomEndpointRouteCommandInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrRestoreCustomEndpointRouteCommandInput) ProtoMessage() {}

func (x *DeleteOrRestoreCustomEndpointRouteCommandInput) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrRestoreCustomEndpointRouteCommandInput.ProtoReflect.Descriptor instead.
func (*DeleteOrRestoreCustomEndpointRouteCommandInput) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteOrRestoreCustomEndpointRouteCommandInput) GetCustomEndpointId() string {
	if x != nil {
		return x.CustomEndpointId
	}
	return ""
}

func (x *DeleteOrRestoreCustomEndpointRouteCommandInput) GetUrlPathPrefix() string {
	if x != nil {
		return x.UrlPathPrefix
	}
	return ""
}

func (x *DeleteOrRestoreCustomEndpointRouteCommandInput) GetVersionMessage() string {
	if x != nil {
		return x.VersionMessage
	}
	return ""
}

// ssl certificate status of custom-endpoint
type CustomEndpointCertStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// flag to indicate if the certificate has been issued.
	IsCertIssued bool `protobuf:"varint,1,opt,name=is_cert_issued,json=isCertIssued,proto3" json:"is_cert_issued,omitempty"`
}

func (x *CustomEndpointCertStatus) Reset() {
	*x = CustomEndpointCertStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomEndpointCertStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEndpointCertStatus) ProtoMessage() {}

func (x *CustomEndpointCertStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEndpointCertStatus.ProtoReflect.Descriptor instead.
func (*CustomEndpointCertStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescGZIP(), []int{5}
}

func (x *CustomEndpointCertStatus) GetIsCertIssued() bool {
	if x != nil {
		return x.IsCertIssued
	}
	return false
}

// dns resolution status of endpoint
type CustomEndpointDnsResolutionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// flag to indicate if the dns is resolving to the correct ingress address(ip or cname).
	IsResolvingToIngressAddress bool `protobuf:"varint,1,opt,name=is_resolving_to_ingress_address,json=isResolvingToIngressAddress,proto3" json:"is_resolving_to_ingress_address,omitempty"`
}

func (x *CustomEndpointDnsResolutionStatus) Reset() {
	*x = CustomEndpointDnsResolutionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomEndpointDnsResolutionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEndpointDnsResolutionStatus) ProtoMessage() {}

func (x *CustomEndpointDnsResolutionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEndpointDnsResolutionStatus.ProtoReflect.Descriptor instead.
func (*CustomEndpointDnsResolutionStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescGZIP(), []int{6}
}

func (x *CustomEndpointDnsResolutionStatus) GetIsResolvingToIngressAddress() bool {
	if x != nil {
		return x.IsResolvingToIngressAddress
	}
	return false
}

var File_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x72, 0x0a, 0x0f, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x5f, 0x0a, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x96, 0x01,
	0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x5f, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0xf5, 0x01, 0x0a, 0x2a, 0x41, 0x64, 0x64, 0x4f, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x34, 0x0a, 0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x60, 0x0a, 0x05, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x2f, 0x0a,
	0x0f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc7,
	0x01, 0x0a, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x52, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x34, 0x0a, 0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0f, 0x75, 0x72, 0x6c, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0d, 0x75, 0x72, 0x6c, 0x50, 0x61, 0x74,
	0x68, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x2f, 0x0a, 0x0f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x18, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73,
	0x43, 0x65, 0x72, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x22, 0x69, 0x0a, 0x21, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x44, 0x0a, 0x1f, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x6f, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x69, 0x73, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0xb3, 0x03, 0x0a, 0x43, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x49,
	0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2,
	0x02, 0x07, 0x43, 0x50, 0x41, 0x43, 0x56, 0x43, 0x4d, 0xaa, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0xca, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x41, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43,
	0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5c, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3b,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a,
	0x41, 0x70, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_goTypes = []interface{}{
	(*CustomEndpointId)(nil),                               // 0: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointId
	(*CustomEndpoints)(nil),                                // 1: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpoints
	(*CustomEndpointList)(nil),                             // 2: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointList
	(*AddOrUpdateCustomEndpointRouteCommandInput)(nil),     // 3: cloud.planton.apis.code2cloud.v1.customendpoint.model.AddOrUpdateCustomEndpointRouteCommandInput
	(*DeleteOrRestoreCustomEndpointRouteCommandInput)(nil), // 4: cloud.planton.apis.code2cloud.v1.customendpoint.model.DeleteOrRestoreCustomEndpointRouteCommandInput
	(*CustomEndpointCertStatus)(nil),                       // 5: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointCertStatus
	(*CustomEndpointDnsResolutionStatus)(nil),              // 6: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointDnsResolutionStatus
	(*CustomEndpoint)(nil),                                 // 7: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpoint
	(*CustomEndpointRoute)(nil),                            // 8: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointRoute
}
var file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_depIdxs = []int32{
	7, // 0: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpoints.entries:type_name -> cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpoint
	7, // 1: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointList.entries:type_name -> cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpoint
	8, // 2: cloud.planton.apis.code2cloud.v1.customendpoint.model.AddOrUpdateCustomEndpointRouteCommandInput.route:type_name -> cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointRoute
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto != nil {
		return
	}
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomEndpointId); i {
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
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomEndpoints); i {
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
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomEndpointList); i {
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
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrUpdateCustomEndpointRouteCommandInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOrRestoreCustomEndpointRouteCommandInput); i {
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
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomEndpointCertStatus); i {
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
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomEndpointDnsResolutionStatus); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_io_proto_depIdxs = nil
}
