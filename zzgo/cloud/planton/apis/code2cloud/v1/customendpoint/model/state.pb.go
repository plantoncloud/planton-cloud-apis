// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cloud/planton/apis/code2cloud/v1/customendpoint/model/state.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	model "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/model"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcefieldoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemessageoptions"
	_ "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/apiresource/options/apiresourcemetadatamessageoptions"
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

// custom-endpoint
type CustomEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// resource metadata
	Metadata *model.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *CustomEndpointSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status *CustomEndpointStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CustomEndpoint) Reset() {
	*x = CustomEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEndpoint) ProtoMessage() {}

func (x *CustomEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEndpoint.ProtoReflect.Descriptor instead.
func (*CustomEndpoint) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDescGZIP(), []int{0}
}

func (x *CustomEndpoint) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *CustomEndpoint) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CustomEndpoint) GetMetadata() *model.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CustomEndpoint) GetSpec() *CustomEndpointSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *CustomEndpoint) GetStatus() *CustomEndpointStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// custom-endpoint spec
type CustomEndpointSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the company to which the custom-endpoint belongs to.
	// value is computed from the product.
	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// the product to which the custom-endpoint to be added.
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// the kube-cluster in which the custom-endpoint resources are created in.
	// value is computed from the backend-environment.
	KubeClusterId string `protobuf:"bytes,3,opt,name=kube_cluster_id,json=kubeClusterId,proto3" json:"kube_cluster_id,omitempty"`
	// custom-endpoint to which the requests to the endpoint domain are to be routed to.
	BackendEnvironmentId string `protobuf:"bytes,4,opt,name=backend_environment_id,json=backendEnvironmentId,proto3" json:"backend_environment_id,omitempty"`
	// (optional for create) flag to toggle tls for custom-endpoint.
	// defaults to "false".
	// (important note) certificates are not created for endpoints that do not need tls.
	// (important note) endpoint domains with out tls enabled are not eligible to be used for
	// creating endpoints for postgres-clusters or kafka-clusters.
	IsTlsEnabled bool `protobuf:"varint,5,opt,name=is_tls_enabled,json=isTlsEnabled,proto3" json:"is_tls_enabled,omitempty"`
	// flag to control virtual service configuration compatible for grpc-web clients.
	// grpc-web clients would rely on extra headers added by envoy proxy.
	IsGrpcWebCompatible bool `protobuf:"varint,6,opt,name=is_grpc_web_compatible,json=isGrpcWebCompatible,proto3" json:"is_grpc_web_compatible,omitempty"`
	// id of the project on google cloud containing the dns zone for the endpoint-domain.
	// this value is looked up from the dns-domains in the company dns data.
	// this value is required for configuring the cert-issuer to perform dns validations.
	DnsZoneGcpProjectId string `protobuf:"bytes,7,opt,name=dns_zone_gcp_project_id,json=dnsZoneGcpProjectId,proto3" json:"dns_zone_gcp_project_id,omitempty"`
	// external ingress ip.
	// ingress ip address configured for the ingress-endpoint-domain.
	// this value is computed from the kube-cluster.
	ExternalIngressIp string `protobuf:"bytes,8,opt,name=external_ingress_ip,json=externalIngressIp,proto3" json:"external_ingress_ip,omitempty"`
	// email of the service account created for cert-manager component on the kube-cluster in which the
	// custom-endpoint that the custom-endpoint belongs to.
	// this value is used for ensuring that the service account has required permissions to insert dns records in
	// the dns zone to be able to complete dns01 challenges.
	// this attribute is only populated if the custom-endpoint is hosted in a gcp kube-cluster.
	// this value is computed from the kube-cluster.
	CertManagerGsaEmail string `protobuf:"bytes,9,opt,name=cert_manager_gsa_email,json=certManagerGsaEmail,proto3" json:"cert_manager_gsa_email,omitempty"`
	// routes to configure backends for url paths of the domain.
	// routes allow configuring requests to be routed to different microservices based on the url path.
	// ex: if the endpoint domain name is console.example.com, then console.example.com/public/api/* can be
	// to be sent to public-api microservice and console.example.com/private/api/* can be configured to route to
	// route to private-api microservice.
	Routes []*CustomEndpointRoute `protobuf:"bytes,10,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *CustomEndpointSpec) Reset() {
	*x = CustomEndpointSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomEndpointSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEndpointSpec) ProtoMessage() {}

func (x *CustomEndpointSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEndpointSpec.ProtoReflect.Descriptor instead.
func (*CustomEndpointSpec) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDescGZIP(), []int{1}
}

func (x *CustomEndpointSpec) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CustomEndpointSpec) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CustomEndpointSpec) GetKubeClusterId() string {
	if x != nil {
		return x.KubeClusterId
	}
	return ""
}

func (x *CustomEndpointSpec) GetBackendEnvironmentId() string {
	if x != nil {
		return x.BackendEnvironmentId
	}
	return ""
}

func (x *CustomEndpointSpec) GetIsTlsEnabled() bool {
	if x != nil {
		return x.IsTlsEnabled
	}
	return false
}

func (x *CustomEndpointSpec) GetIsGrpcWebCompatible() bool {
	if x != nil {
		return x.IsGrpcWebCompatible
	}
	return false
}

func (x *CustomEndpointSpec) GetDnsZoneGcpProjectId() string {
	if x != nil {
		return x.DnsZoneGcpProjectId
	}
	return ""
}

func (x *CustomEndpointSpec) GetExternalIngressIp() string {
	if x != nil {
		return x.ExternalIngressIp
	}
	return ""
}

func (x *CustomEndpointSpec) GetCertManagerGsaEmail() string {
	if x != nil {
		return x.CertManagerGsaEmail
	}
	return ""
}

func (x *CustomEndpointSpec) GetRoutes() []*CustomEndpointRoute {
	if x != nil {
		return x.Routes
	}
	return nil
}

// custom-endpoint status
type CustomEndpointStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource lifecycle
	Lifecycle *model.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// resource audit info
	Audit *model.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// id of the stack-job
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
}

func (x *CustomEndpointStatus) Reset() {
	*x = CustomEndpointStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomEndpointStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEndpointStatus) ProtoMessage() {}

func (x *CustomEndpointStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEndpointStatus.ProtoReflect.Descriptor instead.
func (*CustomEndpointStatus) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDescGZIP(), []int{2}
}

func (x *CustomEndpointStatus) GetLifecycle() *model.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *CustomEndpointStatus) GetAudit() *model.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *CustomEndpointStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

// custom-endpoint route
type CustomEndpointRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// url path for the custom-endpoint route.
	// (important) url path prefix combined with custom_endpoint_id combined form the composite key.
	// (note) if endpoint domain name is console.example.com, and if the url_path_prefix is /api, then
	// all requests that match console.example.com/api/* are forwarded to the backend microservice configured in
	// this route.
	UrlPathPrefix string `protobuf:"bytes,2,opt,name=url_path_prefix,json=urlPathPrefix,proto3" json:"url_path_prefix,omitempty"`
	// id of the backend microservice-instance to configure the forwarding rule.
	// ex: msi-planton-pcs-prod-company-main
	BackendMicroserviceInstanceId string `protobuf:"bytes,3,opt,name=backend_microservice_instance_id,json=backendMicroserviceInstanceId,proto3" json:"backend_microservice_instance_id,omitempty"`
	// service port to which the requests matching the url
	// defaults to 80 if the input contains 0.
	BackendMicroserviceServicePort int32 `protobuf:"varint,4,opt,name=backend_microservice_service_port,json=backendMicroserviceServicePort,proto3" json:"backend_microservice_service_port,omitempty"`
	// backend kubernetes endpoint to configure the forwarding rule.
	// this value is computed from the status of the configured microservice-instance
	BackendKubernetesEndpoint string `protobuf:"bytes,5,opt,name=backend_kubernetes_endpoint,json=backendKubernetesEndpoint,proto3" json:"backend_kubernetes_endpoint,omitempty"`
}

func (x *CustomEndpointRoute) Reset() {
	*x = CustomEndpointRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomEndpointRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomEndpointRoute) ProtoMessage() {}

func (x *CustomEndpointRoute) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomEndpointRoute.ProtoReflect.Descriptor instead.
func (*CustomEndpointRoute) Descriptor() ([]byte, []int) {
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDescGZIP(), []int{3}
}

func (x *CustomEndpointRoute) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CustomEndpointRoute) GetUrlPathPrefix() string {
	if x != nil {
		return x.UrlPathPrefix
	}
	return ""
}

func (x *CustomEndpointRoute) GetBackendMicroserviceInstanceId() string {
	if x != nil {
		return x.BackendMicroserviceInstanceId
	}
	return ""
}

func (x *CustomEndpointRoute) GetBackendMicroserviceServicePort() int32 {
	if x != nil {
		return x.BackendMicroserviceServicePort
	}
	return 0
}

func (x *CustomEndpointRoute) GetBackendKubernetesEndpoint() string {
	if x != nil {
		return x.BackendKubernetesEndpoint
	}
	return ""
}

var File_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto protoreflect.FileDescriptor

var file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDesc = []byte{
	0x0a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6b, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x7c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x06, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xba,
	0x48, 0x1f, 0x72, 0x1d, 0x0a, 0x1b, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x31, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xba, 0x48, 0x12,
	0x72, 0x10, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0xea, 0x03, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x8a,
	0x03, 0xba, 0x48, 0x86, 0x03, 0xba, 0x01, 0xa6, 0x01, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x73,
	0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x61, 0x6e, 0x79, 0x20, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x20, 0x44, 0x4e, 0x53, 0x20, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x20, 0x4e, 0x61,
	0x6d, 0x65, 0x1a, 0x6b, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x28, 0x3f, 0x3a, 0x5b, 0x5f, 0x61, 0x2d, 0x7a,
	0x30, 0x2d, 0x39, 0x5d, 0x28, 0x3f, 0x3a, 0x5b, 0x5f, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d,
	0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29,
	0x3f, 0x5b, 0x2e, 0x5d, 0x29, 0x2b, 0x28, 0x3f, 0x3a, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x28, 0x3f,
	0x3a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x29, 0x3f, 0x24, 0x27, 0x29, 0xba,
	0x01, 0x6c, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20,
	0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x20, 0x31, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x36, 0x35,
	0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x20, 0x6c, 0x6f, 0x6e, 0x67,
	0x1a, 0x2c, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x29, 0x20, 0x3e, 0x20, 0x30, 0x20, 0x26, 0x26, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20, 0x3c, 0x3d, 0x20, 0x36, 0x35, 0xba, 0x01,
	0x67, 0x0a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x69, 0x73, 0x20,
	0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x68,
	0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x29, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5d, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x12, 0x63, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x26, 0x8a, 0xb5, 0x18, 0x03, 0x63,
	0x65, 0x70, 0x98, 0xb5, 0x18, 0x01, 0x88, 0xa6, 0x1d, 0x08, 0x9a, 0xa6, 0x1d, 0x13, 0x08, 0x13,
	0x12, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x22, 0xc0, 0x04, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x23, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8,
	0x18, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x29, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0xd0, 0xb8, 0x18, 0x01, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0f, 0x6b, 0x75, 0x62, 0x65,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x0d, 0x6b, 0x75, 0x62, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x16, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0xd0, 0xb8,
	0x18, 0x01, 0x52, 0x14, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x74,
	0x6c, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x69, 0x73, 0x54, 0x6c, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x33,
	0x0a, 0x16, 0x69, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x77, 0x65, 0x62, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13,
	0x69, 0x73, 0x47, 0x72, 0x70, 0x63, 0x57, 0x65, 0x62, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69,
	0x62, 0x6c, 0x65, 0x12, 0x3a, 0x0a, 0x17, 0x64, 0x6e, 0x73, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f,
	0x67, 0x63, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x13, 0x64, 0x6e, 0x73, 0x5a,
	0x6f, 0x6e, 0x65, 0x47, 0x63, 0x70, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x34, 0x0a, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8,
	0x18, 0x01, 0x52, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x70, 0x12, 0x39, 0x0a, 0x16, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x67, 0x73, 0x61, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18, 0x01, 0x52, 0x13, 0x63, 0x65, 0x72,
	0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x47, 0x73, 0x61, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x62, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x4a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x22, 0xf0, 0x01, 0x0a, 0x14, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x60, 0x0a,
	0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x42, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12,
	0x54, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x62, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x05,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6a,
	0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x61, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0xc3, 0x02, 0x0a, 0x13, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8, 0x18,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x0f, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0d, 0x75, 0x72, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x4f, 0x0a, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x1d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x21, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x04, 0x80, 0xb9, 0x18, 0x50, 0x52, 0x1e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x44, 0x0a, 0x1b, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xc8, 0xb8,
	0x18, 0x01, 0x52, 0x19, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0xb6, 0x03,
	0x0a, 0x43, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x7a, 0x7a, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xa2, 0x02, 0x07, 0x43, 0x50, 0x41,
	0x43, 0x56, 0x43, 0x4d, 0xaa, 0x02, 0x35, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x35, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69,
	0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5c, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x41, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x41, 0x70, 0x69, 0x73, 0x5c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3b, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x73, 0x3a,
	0x3a, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3a,
	0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDescOnce sync.Once
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDescData = file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDesc
)

func file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDescGZIP() []byte {
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDescOnce.Do(func() {
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDescData)
	})
	return file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDescData
}

var file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_goTypes = []interface{}{
	(*CustomEndpoint)(nil),             // 0: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpoint
	(*CustomEndpointSpec)(nil),         // 1: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointSpec
	(*CustomEndpointStatus)(nil),       // 2: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointStatus
	(*CustomEndpointRoute)(nil),        // 3: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointRoute
	(*model.ApiResourceMetadata)(nil),  // 4: cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	(*model.ApiResourceLifecycle)(nil), // 5: cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	(*model.ApiResourceAudit)(nil),     // 6: cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
}
var file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_depIdxs = []int32{
	4, // 0: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpoint.metadata:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata
	1, // 1: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpoint.spec:type_name -> cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointSpec
	2, // 2: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpoint.status:type_name -> cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointStatus
	3, // 3: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointSpec.routes:type_name -> cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointRoute
	5, // 4: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointStatus.lifecycle:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceLifecycle
	6, // 5: cloud.planton.apis.code2cloud.v1.customendpoint.model.CustomEndpointStatus.audit:type_name -> cloud.planton.apis.commons.apiresource.model.ApiResourceAudit
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_init() }
func file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_init() {
	if File_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomEndpoint); i {
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
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomEndpointSpec); i {
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
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomEndpointStatus); i {
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
		file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomEndpointRoute); i {
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
			RawDescriptor: file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_goTypes,
		DependencyIndexes: file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_depIdxs,
		MessageInfos:      file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_msgTypes,
	}.Build()
	File_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto = out.File
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_rawDesc = nil
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_goTypes = nil
	file_cloud_planton_apis_code2cloud_v1_customendpoint_model_state_proto_depIdxs = nil
}
