// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cloud/planton/apis/v1/code2cloud/deploy/kubecluster/query.proto

package kubecluster

import (
	context "context"
	cloudaccount "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/cloudaccount"
	stream "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/grpc/stream"
	custom "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/protobuf/custom"
	pagination "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/commons/rpc/pagination"
	resource "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/integration/kubernetes/resource"
	company "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/resourcemanager/company"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KubeClusterQueryController_List_FullMethodName                                    = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterQueryController/list"
	KubeClusterQueryController_GetById_FullMethodName                                 = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterQueryController/getById"
	KubeClusterQueryController_FindByCompanyId_FullMethodName                         = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterQueryController/findByCompanyId"
	KubeClusterQueryController_FindByCloudAccountId_FullMethodName                    = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterQueryController/findByCloudAccountId"
	KubeClusterQueryController_FindEnvironmentCreateKubeClusters_FullMethodName       = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterQueryController/findEnvironmentCreateKubeClusters"
	KubeClusterQueryController_GetKubeClusterComponentsByKubeClusterId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterQueryController/getKubeClusterComponentsByKubeClusterId"
	KubeClusterQueryController_FindWorkloadNamespacesByKubeClusterId_FullMethodName   = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterQueryController/findWorkloadNamespacesByKubeClusterId"
	KubeClusterQueryController_FindWorkloadPodsByKubeClusterId_FullMethodName         = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterQueryController/findWorkloadPodsByKubeClusterId"
	KubeClusterQueryController_FindSslCertificatesByKubeClusterId_FullMethodName      = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterQueryController/findSslCertificatesByKubeClusterId"
	KubeClusterQueryController_GetPod_FullMethodName                                  = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterQueryController/getPod"
	KubeClusterQueryController_GetPodLogStream_FullMethodName                         = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterQueryController/getPodLogStream"
)

// KubeClusterQueryControllerClient is the client API for KubeClusterQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubeClusterQueryControllerClient interface {
	// list all kube-clusters on planton cloud for the requested page. This is intended for use on portal.
	List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*KubeClusterList, error)
	// lookup kube-cluster using kube-cluster id
	GetById(ctx context.Context, in *KubeClusterId, opts ...grpc.CallOption) (*KubeCluster, error)
	// find kube-clusters by company id
	FindByCompanyId(ctx context.Context, in *company.CompanyId, opts ...grpc.CallOption) (*KubeClusters, error)
	// find kube-clusters in a cloud account.
	FindByCloudAccountId(ctx context.Context, in *cloudaccount.CloudAccountId, opts ...grpc.CallOption) (*KubeClusters, error)
	// find kube-clusters by company id to create environment.
	// this will be used to populate drop down of kube-clusters in create environment form.
	// the response should only include kube-clusters that a company is authorised to create environment.
	// the authorization is verified by looking up kube-clusters with `company-environment-creator` relation with the company id provided in input.
	FindEnvironmentCreateKubeClusters(ctx context.Context, in *company.CompanyId, opts ...grpc.CallOption) (*KubeClusters, error)
	// lookup components in a kube-cluster of a kube-cluster
	GetKubeClusterComponentsByKubeClusterId(ctx context.Context, in *KubeClusterId, opts ...grpc.CallOption) (*KubeClusterComponents, error)
	// find workload namespaces in a kube-cluster.
	FindWorkloadNamespacesByKubeClusterId(ctx context.Context, in *KubeClusterId, opts ...grpc.CallOption) (*resource.WorkloadNamespaces, error)
	// find workload pods part of all environments hosted in a kube-cluster.
	FindWorkloadPodsByKubeClusterId(ctx context.Context, in *KubeClusterId, opts ...grpc.CallOption) (*resource.WorkloadPods, error)
	// find workload pods part of all environments hosted in a kube-cluster.
	FindSslCertificatesByKubeClusterId(ctx context.Context, in *KubeClusterId, opts ...grpc.CallOption) (*resource.Certificates, error)
	// get a pod details
	GetPod(ctx context.Context, in *ByKubeClusterByNamespaceByPodInput, opts ...grpc.CallOption) (*resource.Pod, error)
	// get a log stream for a pod running in a kube-cluster kube-cluster
	GetPodLogStream(ctx context.Context, in *ByKubeClusterByNamespaceByPodInput, opts ...grpc.CallOption) (KubeClusterQueryController_GetPodLogStreamClient, error)
}

type kubeClusterQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKubeClusterQueryControllerClient(cc grpc.ClientConnInterface) KubeClusterQueryControllerClient {
	return &kubeClusterQueryControllerClient{cc}
}

func (c *kubeClusterQueryControllerClient) List(ctx context.Context, in *pagination.PageInfo, opts ...grpc.CallOption) (*KubeClusterList, error) {
	out := new(KubeClusterList)
	err := c.cc.Invoke(ctx, KubeClusterQueryController_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeClusterQueryControllerClient) GetById(ctx context.Context, in *KubeClusterId, opts ...grpc.CallOption) (*KubeCluster, error) {
	out := new(KubeCluster)
	err := c.cc.Invoke(ctx, KubeClusterQueryController_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeClusterQueryControllerClient) FindByCompanyId(ctx context.Context, in *company.CompanyId, opts ...grpc.CallOption) (*KubeClusters, error) {
	out := new(KubeClusters)
	err := c.cc.Invoke(ctx, KubeClusterQueryController_FindByCompanyId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeClusterQueryControllerClient) FindByCloudAccountId(ctx context.Context, in *cloudaccount.CloudAccountId, opts ...grpc.CallOption) (*KubeClusters, error) {
	out := new(KubeClusters)
	err := c.cc.Invoke(ctx, KubeClusterQueryController_FindByCloudAccountId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeClusterQueryControllerClient) FindEnvironmentCreateKubeClusters(ctx context.Context, in *company.CompanyId, opts ...grpc.CallOption) (*KubeClusters, error) {
	out := new(KubeClusters)
	err := c.cc.Invoke(ctx, KubeClusterQueryController_FindEnvironmentCreateKubeClusters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeClusterQueryControllerClient) GetKubeClusterComponentsByKubeClusterId(ctx context.Context, in *KubeClusterId, opts ...grpc.CallOption) (*KubeClusterComponents, error) {
	out := new(KubeClusterComponents)
	err := c.cc.Invoke(ctx, KubeClusterQueryController_GetKubeClusterComponentsByKubeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeClusterQueryControllerClient) FindWorkloadNamespacesByKubeClusterId(ctx context.Context, in *KubeClusterId, opts ...grpc.CallOption) (*resource.WorkloadNamespaces, error) {
	out := new(resource.WorkloadNamespaces)
	err := c.cc.Invoke(ctx, KubeClusterQueryController_FindWorkloadNamespacesByKubeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeClusterQueryControllerClient) FindWorkloadPodsByKubeClusterId(ctx context.Context, in *KubeClusterId, opts ...grpc.CallOption) (*resource.WorkloadPods, error) {
	out := new(resource.WorkloadPods)
	err := c.cc.Invoke(ctx, KubeClusterQueryController_FindWorkloadPodsByKubeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeClusterQueryControllerClient) FindSslCertificatesByKubeClusterId(ctx context.Context, in *KubeClusterId, opts ...grpc.CallOption) (*resource.Certificates, error) {
	out := new(resource.Certificates)
	err := c.cc.Invoke(ctx, KubeClusterQueryController_FindSslCertificatesByKubeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeClusterQueryControllerClient) GetPod(ctx context.Context, in *ByKubeClusterByNamespaceByPodInput, opts ...grpc.CallOption) (*resource.Pod, error) {
	out := new(resource.Pod)
	err := c.cc.Invoke(ctx, KubeClusterQueryController_GetPod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeClusterQueryControllerClient) GetPodLogStream(ctx context.Context, in *ByKubeClusterByNamespaceByPodInput, opts ...grpc.CallOption) (KubeClusterQueryController_GetPodLogStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &KubeClusterQueryController_ServiceDesc.Streams[0], KubeClusterQueryController_GetPodLogStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &kubeClusterQueryControllerGetPodLogStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KubeClusterQueryController_GetPodLogStreamClient interface {
	Recv() (*stream.OutputLine, error)
	grpc.ClientStream
}

type kubeClusterQueryControllerGetPodLogStreamClient struct {
	grpc.ClientStream
}

func (x *kubeClusterQueryControllerGetPodLogStreamClient) Recv() (*stream.OutputLine, error) {
	m := new(stream.OutputLine)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KubeClusterQueryControllerServer is the server API for KubeClusterQueryController service.
// All implementations should embed UnimplementedKubeClusterQueryControllerServer
// for forward compatibility
type KubeClusterQueryControllerServer interface {
	// list all kube-clusters on planton cloud for the requested page. This is intended for use on portal.
	List(context.Context, *pagination.PageInfo) (*KubeClusterList, error)
	// lookup kube-cluster using kube-cluster id
	GetById(context.Context, *KubeClusterId) (*KubeCluster, error)
	// find kube-clusters by company id
	FindByCompanyId(context.Context, *company.CompanyId) (*KubeClusters, error)
	// find kube-clusters in a cloud account.
	FindByCloudAccountId(context.Context, *cloudaccount.CloudAccountId) (*KubeClusters, error)
	// find kube-clusters by company id to create environment.
	// this will be used to populate drop down of kube-clusters in create environment form.
	// the response should only include kube-clusters that a company is authorised to create environment.
	// the authorization is verified by looking up kube-clusters with `company-environment-creator` relation with the company id provided in input.
	FindEnvironmentCreateKubeClusters(context.Context, *company.CompanyId) (*KubeClusters, error)
	// lookup components in a kube-cluster of a kube-cluster
	GetKubeClusterComponentsByKubeClusterId(context.Context, *KubeClusterId) (*KubeClusterComponents, error)
	// find workload namespaces in a kube-cluster.
	FindWorkloadNamespacesByKubeClusterId(context.Context, *KubeClusterId) (*resource.WorkloadNamespaces, error)
	// find workload pods part of all environments hosted in a kube-cluster.
	FindWorkloadPodsByKubeClusterId(context.Context, *KubeClusterId) (*resource.WorkloadPods, error)
	// find workload pods part of all environments hosted in a kube-cluster.
	FindSslCertificatesByKubeClusterId(context.Context, *KubeClusterId) (*resource.Certificates, error)
	// get a pod details
	GetPod(context.Context, *ByKubeClusterByNamespaceByPodInput) (*resource.Pod, error)
	// get a log stream for a pod running in a kube-cluster kube-cluster
	GetPodLogStream(*ByKubeClusterByNamespaceByPodInput, KubeClusterQueryController_GetPodLogStreamServer) error
}

// UnimplementedKubeClusterQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKubeClusterQueryControllerServer struct {
}

func (UnimplementedKubeClusterQueryControllerServer) List(context.Context, *pagination.PageInfo) (*KubeClusterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedKubeClusterQueryControllerServer) GetById(context.Context, *KubeClusterId) (*KubeCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedKubeClusterQueryControllerServer) FindByCompanyId(context.Context, *company.CompanyId) (*KubeClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCompanyId not implemented")
}
func (UnimplementedKubeClusterQueryControllerServer) FindByCloudAccountId(context.Context, *cloudaccount.CloudAccountId) (*KubeClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCloudAccountId not implemented")
}
func (UnimplementedKubeClusterQueryControllerServer) FindEnvironmentCreateKubeClusters(context.Context, *company.CompanyId) (*KubeClusters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnvironmentCreateKubeClusters not implemented")
}
func (UnimplementedKubeClusterQueryControllerServer) GetKubeClusterComponentsByKubeClusterId(context.Context, *KubeClusterId) (*KubeClusterComponents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKubeClusterComponentsByKubeClusterId not implemented")
}
func (UnimplementedKubeClusterQueryControllerServer) FindWorkloadNamespacesByKubeClusterId(context.Context, *KubeClusterId) (*resource.WorkloadNamespaces, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindWorkloadNamespacesByKubeClusterId not implemented")
}
func (UnimplementedKubeClusterQueryControllerServer) FindWorkloadPodsByKubeClusterId(context.Context, *KubeClusterId) (*resource.WorkloadPods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindWorkloadPodsByKubeClusterId not implemented")
}
func (UnimplementedKubeClusterQueryControllerServer) FindSslCertificatesByKubeClusterId(context.Context, *KubeClusterId) (*resource.Certificates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSslCertificatesByKubeClusterId not implemented")
}
func (UnimplementedKubeClusterQueryControllerServer) GetPod(context.Context, *ByKubeClusterByNamespaceByPodInput) (*resource.Pod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPod not implemented")
}
func (UnimplementedKubeClusterQueryControllerServer) GetPodLogStream(*ByKubeClusterByNamespaceByPodInput, KubeClusterQueryController_GetPodLogStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPodLogStream not implemented")
}

// UnsafeKubeClusterQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubeClusterQueryControllerServer will
// result in compilation errors.
type UnsafeKubeClusterQueryControllerServer interface {
	mustEmbedUnimplementedKubeClusterQueryControllerServer()
}

func RegisterKubeClusterQueryControllerServer(s grpc.ServiceRegistrar, srv KubeClusterQueryControllerServer) {
	s.RegisterService(&KubeClusterQueryController_ServiceDesc, srv)
}

func _KubeClusterQueryController_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pagination.PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterQueryControllerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterQueryController_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterQueryControllerServer).List(ctx, req.(*pagination.PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeClusterQueryController_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterQueryControllerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterQueryController_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterQueryControllerServer).GetById(ctx, req.(*KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeClusterQueryController_FindByCompanyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(company.CompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterQueryControllerServer).FindByCompanyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterQueryController_FindByCompanyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterQueryControllerServer).FindByCompanyId(ctx, req.(*company.CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeClusterQueryController_FindByCloudAccountId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cloudaccount.CloudAccountId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterQueryControllerServer).FindByCloudAccountId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterQueryController_FindByCloudAccountId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterQueryControllerServer).FindByCloudAccountId(ctx, req.(*cloudaccount.CloudAccountId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeClusterQueryController_FindEnvironmentCreateKubeClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(company.CompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterQueryControllerServer).FindEnvironmentCreateKubeClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterQueryController_FindEnvironmentCreateKubeClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterQueryControllerServer).FindEnvironmentCreateKubeClusters(ctx, req.(*company.CompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeClusterQueryController_GetKubeClusterComponentsByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterQueryControllerServer).GetKubeClusterComponentsByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterQueryController_GetKubeClusterComponentsByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterQueryControllerServer).GetKubeClusterComponentsByKubeClusterId(ctx, req.(*KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeClusterQueryController_FindWorkloadNamespacesByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterQueryControllerServer).FindWorkloadNamespacesByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterQueryController_FindWorkloadNamespacesByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterQueryControllerServer).FindWorkloadNamespacesByKubeClusterId(ctx, req.(*KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeClusterQueryController_FindWorkloadPodsByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterQueryControllerServer).FindWorkloadPodsByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterQueryController_FindWorkloadPodsByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterQueryControllerServer).FindWorkloadPodsByKubeClusterId(ctx, req.(*KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeClusterQueryController_FindSslCertificatesByKubeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeClusterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterQueryControllerServer).FindSslCertificatesByKubeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterQueryController_FindSslCertificatesByKubeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterQueryControllerServer).FindSslCertificatesByKubeClusterId(ctx, req.(*KubeClusterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeClusterQueryController_GetPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByKubeClusterByNamespaceByPodInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterQueryControllerServer).GetPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterQueryController_GetPod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterQueryControllerServer).GetPod(ctx, req.(*ByKubeClusterByNamespaceByPodInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeClusterQueryController_GetPodLogStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ByKubeClusterByNamespaceByPodInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KubeClusterQueryControllerServer).GetPodLogStream(m, &kubeClusterQueryControllerGetPodLogStreamServer{stream})
}

type KubeClusterQueryController_GetPodLogStreamServer interface {
	Send(*stream.OutputLine) error
	grpc.ServerStream
}

type kubeClusterQueryControllerGetPodLogStreamServer struct {
	grpc.ServerStream
}

func (x *kubeClusterQueryControllerGetPodLogStreamServer) Send(m *stream.OutputLine) error {
	return x.ServerStream.SendMsg(m)
}

// KubeClusterQueryController_ServiceDesc is the grpc.ServiceDesc for KubeClusterQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubeClusterQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterQueryController",
	HandlerType: (*KubeClusterQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _KubeClusterQueryController_List_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _KubeClusterQueryController_GetById_Handler,
		},
		{
			MethodName: "findByCompanyId",
			Handler:    _KubeClusterQueryController_FindByCompanyId_Handler,
		},
		{
			MethodName: "findByCloudAccountId",
			Handler:    _KubeClusterQueryController_FindByCloudAccountId_Handler,
		},
		{
			MethodName: "findEnvironmentCreateKubeClusters",
			Handler:    _KubeClusterQueryController_FindEnvironmentCreateKubeClusters_Handler,
		},
		{
			MethodName: "getKubeClusterComponentsByKubeClusterId",
			Handler:    _KubeClusterQueryController_GetKubeClusterComponentsByKubeClusterId_Handler,
		},
		{
			MethodName: "findWorkloadNamespacesByKubeClusterId",
			Handler:    _KubeClusterQueryController_FindWorkloadNamespacesByKubeClusterId_Handler,
		},
		{
			MethodName: "findWorkloadPodsByKubeClusterId",
			Handler:    _KubeClusterQueryController_FindWorkloadPodsByKubeClusterId_Handler,
		},
		{
			MethodName: "findSslCertificatesByKubeClusterId",
			Handler:    _KubeClusterQueryController_FindSslCertificatesByKubeClusterId_Handler,
		},
		{
			MethodName: "getPod",
			Handler:    _KubeClusterQueryController_GetPod_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getPodLogStream",
			Handler:       _KubeClusterQueryController_GetPodLogStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/kubecluster/query.proto",
}

const (
	KubeClusterNodePoolGcpQueryController_GetByGcpContainerNodePoolId_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterNodePoolGcpQueryController/getByGcpContainerNodePoolId"
)

// KubeClusterNodePoolGcpQueryControllerClient is the client API for KubeClusterNodePoolGcpQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubeClusterNodePoolGcpQueryControllerClient interface {
	// lookup gcp container node pool env using container-nodepool-id
	GetByGcpContainerNodePoolId(ctx context.Context, in *GetByKubeClusterNodePoolGcpIdInput, opts ...grpc.CallOption) (*KubeClusterNodePoolGcp, error)
}

type kubeClusterNodePoolGcpQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKubeClusterNodePoolGcpQueryControllerClient(cc grpc.ClientConnInterface) KubeClusterNodePoolGcpQueryControllerClient {
	return &kubeClusterNodePoolGcpQueryControllerClient{cc}
}

func (c *kubeClusterNodePoolGcpQueryControllerClient) GetByGcpContainerNodePoolId(ctx context.Context, in *GetByKubeClusterNodePoolGcpIdInput, opts ...grpc.CallOption) (*KubeClusterNodePoolGcp, error) {
	out := new(KubeClusterNodePoolGcp)
	err := c.cc.Invoke(ctx, KubeClusterNodePoolGcpQueryController_GetByGcpContainerNodePoolId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubeClusterNodePoolGcpQueryControllerServer is the server API for KubeClusterNodePoolGcpQueryController service.
// All implementations should embed UnimplementedKubeClusterNodePoolGcpQueryControllerServer
// for forward compatibility
type KubeClusterNodePoolGcpQueryControllerServer interface {
	// lookup gcp container node pool env using container-nodepool-id
	GetByGcpContainerNodePoolId(context.Context, *GetByKubeClusterNodePoolGcpIdInput) (*KubeClusterNodePoolGcp, error)
}

// UnimplementedKubeClusterNodePoolGcpQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedKubeClusterNodePoolGcpQueryControllerServer struct {
}

func (UnimplementedKubeClusterNodePoolGcpQueryControllerServer) GetByGcpContainerNodePoolId(context.Context, *GetByKubeClusterNodePoolGcpIdInput) (*KubeClusterNodePoolGcp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByGcpContainerNodePoolId not implemented")
}

// UnsafeKubeClusterNodePoolGcpQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubeClusterNodePoolGcpQueryControllerServer will
// result in compilation errors.
type UnsafeKubeClusterNodePoolGcpQueryControllerServer interface {
	mustEmbedUnimplementedKubeClusterNodePoolGcpQueryControllerServer()
}

func RegisterKubeClusterNodePoolGcpQueryControllerServer(s grpc.ServiceRegistrar, srv KubeClusterNodePoolGcpQueryControllerServer) {
	s.RegisterService(&KubeClusterNodePoolGcpQueryController_ServiceDesc, srv)
}

func _KubeClusterNodePoolGcpQueryController_GetByGcpContainerNodePoolId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByKubeClusterNodePoolGcpIdInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeClusterNodePoolGcpQueryControllerServer).GetByGcpContainerNodePoolId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeClusterNodePoolGcpQueryController_GetByGcpContainerNodePoolId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeClusterNodePoolGcpQueryControllerServer).GetByGcpContainerNodePoolId(ctx, req.(*GetByKubeClusterNodePoolGcpIdInput))
	}
	return interceptor(ctx, in, info, handler)
}

// KubeClusterNodePoolGcpQueryController_ServiceDesc is the grpc.ServiceDesc for KubeClusterNodePoolGcpQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubeClusterNodePoolGcpQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.KubeClusterNodePoolGcpQueryController",
	HandlerType: (*KubeClusterNodePoolGcpQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getByGcpContainerNodePoolId",
			Handler:    _KubeClusterNodePoolGcpQueryController_GetByGcpContainerNodePoolId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/kubecluster/query.proto",
}

const (
	GcpQueryController_FindRegions_FullMethodName                 = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.GcpQueryController/findRegions"
	GcpQueryController_FindZonesByRegionIdentifier_FullMethodName = "/cloud.planton.apis.v1.code2cloud.deploy.kubecluster.GcpQueryController/findZonesByRegionIdentifier"
)

// GcpQueryControllerClient is the client API for GcpQueryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GcpQueryControllerClient interface {
	// list all gcp regions
	FindRegions(ctx context.Context, in *custom.CustomEmpty, opts ...grpc.CallOption) (*GcpRegions, error)
	// list all zones in a gcp region
	FindZonesByRegionIdentifier(ctx context.Context, in *GcpRegionIdentifier, opts ...grpc.CallOption) (*GcpZones, error)
}

type gcpQueryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGcpQueryControllerClient(cc grpc.ClientConnInterface) GcpQueryControllerClient {
	return &gcpQueryControllerClient{cc}
}

func (c *gcpQueryControllerClient) FindRegions(ctx context.Context, in *custom.CustomEmpty, opts ...grpc.CallOption) (*GcpRegions, error) {
	out := new(GcpRegions)
	err := c.cc.Invoke(ctx, GcpQueryController_FindRegions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcpQueryControllerClient) FindZonesByRegionIdentifier(ctx context.Context, in *GcpRegionIdentifier, opts ...grpc.CallOption) (*GcpZones, error) {
	out := new(GcpZones)
	err := c.cc.Invoke(ctx, GcpQueryController_FindZonesByRegionIdentifier_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GcpQueryControllerServer is the server API for GcpQueryController service.
// All implementations should embed UnimplementedGcpQueryControllerServer
// for forward compatibility
type GcpQueryControllerServer interface {
	// list all gcp regions
	FindRegions(context.Context, *custom.CustomEmpty) (*GcpRegions, error)
	// list all zones in a gcp region
	FindZonesByRegionIdentifier(context.Context, *GcpRegionIdentifier) (*GcpZones, error)
}

// UnimplementedGcpQueryControllerServer should be embedded to have forward compatible implementations.
type UnimplementedGcpQueryControllerServer struct {
}

func (UnimplementedGcpQueryControllerServer) FindRegions(context.Context, *custom.CustomEmpty) (*GcpRegions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRegions not implemented")
}
func (UnimplementedGcpQueryControllerServer) FindZonesByRegionIdentifier(context.Context, *GcpRegionIdentifier) (*GcpZones, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindZonesByRegionIdentifier not implemented")
}

// UnsafeGcpQueryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GcpQueryControllerServer will
// result in compilation errors.
type UnsafeGcpQueryControllerServer interface {
	mustEmbedUnimplementedGcpQueryControllerServer()
}

func RegisterGcpQueryControllerServer(s grpc.ServiceRegistrar, srv GcpQueryControllerServer) {
	s.RegisterService(&GcpQueryController_ServiceDesc, srv)
}

func _GcpQueryController_FindRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(custom.CustomEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpQueryControllerServer).FindRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpQueryController_FindRegions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpQueryControllerServer).FindRegions(ctx, req.(*custom.CustomEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcpQueryController_FindZonesByRegionIdentifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GcpRegionIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcpQueryControllerServer).FindZonesByRegionIdentifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GcpQueryController_FindZonesByRegionIdentifier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcpQueryControllerServer).FindZonesByRegionIdentifier(ctx, req.(*GcpRegionIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

// GcpQueryController_ServiceDesc is the grpc.ServiceDesc for GcpQueryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GcpQueryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.GcpQueryController",
	HandlerType: (*GcpQueryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findRegions",
			Handler:    _GcpQueryController_FindRegions_Handler,
		},
		{
			MethodName: "findZonesByRegionIdentifier",
			Handler:    _GcpQueryController_FindZonesByRegionIdentifier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud/planton/apis/v1/code2cloud/deploy/kubecluster/query.proto",
}