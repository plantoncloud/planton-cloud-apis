// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/v1/resourcemanager/product/model.proto

package product

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	audit "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/audit"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/field/options"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/metadata/options"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/resource/options"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack"
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

// product
type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// resource metadata
	Metadata *resource.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *ProductSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *ProductStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Product) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Product) GetMetadata() *resource.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Product) GetSpec() *ProductSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Product) GetStatus() *ProductStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// product spec
type ProductSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// company to which the product belongs to.
	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// product key is a three character key for the product.
	// this key is unique with in the company.
	// product key length should be minimum 2 chars and maximum of 4 chars.
	// product-key is suffixed to the company-id to create product-id
	ProductKey string `protobuf:"bytes,2,opt,name=product_key,json=productKey,proto3" json:"product_key,omitempty"`
	// product description.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ProductSpec) Reset() {
	*x = ProductSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSpec) ProtoMessage() {}

func (x *ProductSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSpec.ProtoReflect.Descriptor instead.
func (*ProductSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescGZIP(), []int{1}
}

func (x *ProductSpec) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *ProductSpec) GetProductKey() string {
	if x != nil {
		return x.ProductKey
	}
	return ""
}

func (x *ProductSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// product status
type ProductStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource lifecycle
	Lifecycle *resource.ResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// resource audit info
	Audit *audit.ResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// id of the stack-job
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
}

func (x *ProductStatus) Reset() {
	*x = ProductStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductStatus) ProtoMessage() {}

func (x *ProductStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductStatus.ProtoReflect.Descriptor instead.
func (*ProductStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescGZIP(), []int{2}
}

func (x *ProductStatus) GetLifecycle() *resource.ResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *ProductStatus) GetAudit() *audit.ResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *ProductStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

// list of products
type Products struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*Product `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *Products) Reset() {
	*x = Products{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Products) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Products) ProtoMessage() {}

func (x *Products) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Products.ProtoReflect.Descriptor instead.
func (*Products) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescGZIP(), []int{3}
}

func (x *Products) GetEntries() []*Product {
	if x != nil {
		return x.Entries
	}
	return nil
}

// wrapper for product id
type ProductId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ProductId) Reset() {
	*x = ProductId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductId) ProtoMessage() {}

func (x *ProductId) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductId.ProtoReflect.Descriptor instead.
func (*ProductId) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescGZIP(), []int{4}
}

func (x *ProductId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// response for paginated query to get list of products
type ProductList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPages int32      `protobuf:"varint,1,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	Entries    []*Product `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ProductList) Reset() {
	*x = ProductList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductList) ProtoMessage() {}

func (x *ProductList) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductList.ProtoReflect.Descriptor instead.
func (*ProductList) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescGZIP(), []int{5}
}

func (x *ProductList) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *ProductList) GetEntries() []*Product {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_cloud_planton_apis_v1_resourcemanager_product_model_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDesc = []byte{
	0x0a, 0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x51, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x57, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x25, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52,
	0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc0, 0xb8, 0x18, 0x01, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x4e, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x12, 0x54, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x27, 0x90, 0xb5, 0x18, 0x00, 0x98,
	0xb5, 0x18, 0x01, 0x88, 0xa6, 0x1d, 0x14, 0x90, 0xa6, 0x1d, 0x00, 0x9a, 0xa6, 0x1d, 0x13, 0x08,
	0x06, 0x12, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x69, 0x64, 0x22, 0xe7, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x27, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xc0, 0xb8, 0x18, 0x01, 0xd0, 0xb8, 0x18, 0x01,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x8c, 0x01, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x6b, 0xba, 0x48, 0x68, 0xba, 0x01, 0x5f, 0x0a, 0x10, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x4f, 0x6e, 0x6c,
	0x79, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x20, 0x6c, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x20,
	0x61, 0x72, 0x65, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x1a, 0x1b, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a,
	0x30, 0x2d, 0x39, 0x5d, 0x2b, 0x24, 0x27, 0x29, 0x72, 0x04, 0x10, 0x02, 0x18, 0x05, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd4, 0x01, 0x0a,
	0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x57,
	0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x63, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x18, 0x62, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x61, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4a, 0x6f,
	0x62, 0x49, 0x64, 0x22, 0x5c, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12,
	0x50, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x22, 0x21, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0x84, 0x03, 0x0a, 0x3b, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0xa2, 0x02, 0x06, 0x43, 0x50, 0x41, 0x56, 0x52, 0x50, 0xaa, 0x02, 0x2d,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70,
	0x69, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0xca, 0x02, 0x2d,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0xe2, 0x02, 0x39,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70,
	0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x32, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescData = file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDesc
)

func file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescData)
	})
	return file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDescData
}

var file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cloud_planton_apis_v1_resourcemanager_product_model_proto_goTypes = []interface{}{
	(*Product)(nil),                    // 0: cloud.planton.apis.v1.resourcemanager.product.Product
	(*ProductSpec)(nil),                // 1: cloud.planton.apis.v1.resourcemanager.product.ProductSpec
	(*ProductStatus)(nil),              // 2: cloud.planton.apis.v1.resourcemanager.product.ProductStatus
	(*Products)(nil),                   // 3: cloud.planton.apis.v1.resourcemanager.product.Products
	(*ProductId)(nil),                  // 4: cloud.planton.apis.v1.resourcemanager.product.ProductId
	(*ProductList)(nil),                // 5: cloud.planton.apis.v1.resourcemanager.product.ProductList
	(*resource.Metadata)(nil),          // 6: cloud.planton.apis.v1.commons.resource.Metadata
	(*resource.ResourceLifecycle)(nil), // 7: cloud.planton.apis.v1.commons.resource.ResourceLifecycle
	(*audit.ResourceAudit)(nil),        // 8: cloud.planton.apis.v1.commons.audit.ResourceAudit
}
var file_cloud_planton_apis_v1_resourcemanager_product_model_proto_depIdxs = []int32{
	6, // 0: cloud.planton.apis.v1.resourcemanager.product.Product.metadata:type_name -> cloud.planton.apis.v1.commons.resource.Metadata
	1, // 1: cloud.planton.apis.v1.resourcemanager.product.Product.spec:type_name -> cloud.planton.apis.v1.resourcemanager.product.ProductSpec
	2, // 2: cloud.planton.apis.v1.resourcemanager.product.Product.status:type_name -> cloud.planton.apis.v1.resourcemanager.product.ProductStatus
	7, // 3: cloud.planton.apis.v1.resourcemanager.product.ProductStatus.lifecycle:type_name -> cloud.planton.apis.v1.commons.resource.ResourceLifecycle
	8, // 4: cloud.planton.apis.v1.resourcemanager.product.ProductStatus.audit:type_name -> cloud.planton.apis.v1.commons.audit.ResourceAudit
	0, // 5: cloud.planton.apis.v1.resourcemanager.product.Products.entries:type_name -> cloud.planton.apis.v1.resourcemanager.product.Product
	0, // 6: cloud.planton.apis.v1.resourcemanager.product.ProductList.entries:type_name -> cloud.planton.apis.v1.resourcemanager.product.Product
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_v1_resourcemanager_product_model_proto_init() }
func file_cloud_planton_apis_v1_resourcemanager_product_model_proto_init() {
	if File_cloud_planton_apis_v1_resourcemanager_product_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSpec); i {
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
		file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductStatus); i {
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
		file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Products); i {
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
		file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductId); i {
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
		file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductList); i {
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
			RawDescriptor: file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_v1_resourcemanager_product_model_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_v1_resourcemanager_product_model_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_v1_resourcemanager_product_model_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_v1_resourcemanager_product_model_proto = out.File
	file_cloud_planton_apis_v1_resourcemanager_product_model_proto_rawDesc = nil
	file_cloud_planton_apis_v1_resourcemanager_product_model_proto_goTypes = nil
	file_cloud_planton_apis_v1_resourcemanager_product_model_proto_depIdxs = nil
}
