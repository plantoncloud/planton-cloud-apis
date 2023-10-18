// Code generated by protoc-gen-connect-swift. DO NOT EDIT.
//
// Source: cloud/planton/apis/v1/code2cloud/deploy/kubecluster/rpc/service.proto
//

import Connect
import Foundation
import SwiftProtobuf

/// kube-cluster command controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterCommandControllerClientInterface {

    /// create a new kube-cluster.
    @available(iOS 13, *)
    func `create`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster>

    /// update an existing kube-cluster.
    @available(iOS 13, *)
    func `update`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster>

    /// delete a kube-cluster.
    @available(iOS 13, *)
    func `delete`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster>

    /// restore a deleted kube-cluster.
    @available(iOS 13, *)
    func `restore`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster>

    ///pause a kube-cluster.
    ///a kube-cluster is paused by setting the number of nodes in each node pool of the kube-cluster to zero.
    ///microservice, database and kafka cluster workload pods will be deleted as there wont be any nodes to run on.
    ///when the kube-cluster is resumed, the pods come back up online automatically when nodes become available.
    ///when a kube-cluster is paused, cloud provider will not charge for the compute resources(cpu & memory) but
    /// may continue to charge a modest operational fee for the cluster.
    @available(iOS 13, *)
    func `pause`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster>

    ///unpause a kube-cluster.
    ///a kube-cluster is resumed by setting the number of nodes in each node pool of the kube-cluster to the
    /// values configured for the kube-cluster.
    ///when the kube-cluster is resumed, the pods come back up online automatically when nodes become available.
    @available(iOS 13, *)
    func `unpause`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster>

    /// delete a namespace in kube-cluster kube-cluster
    @available(iOS 13, *)
    func `deleteNamespace`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_ByKubeClusterByNamespaceInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_WorkloadNamespace>

    /// delete a pod in kube-cluster kube-cluster
    @available(iOS 13, *)
    func `deletePod`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_ByKubeClusterByNamespaceByPodInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterCommandControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterCommandControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterCommandControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `create`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController/create", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `update`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController/update", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `delete`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController/delete", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `restore`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController/restore", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `pause`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController/pause", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `unpause`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController/unpause", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `deleteNamespace`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_ByKubeClusterByNamespaceInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_WorkloadNamespace> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController/deleteNamespace", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `deletePod`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_ByKubeClusterByNamespaceByPodInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController/deletePod", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let create = Connect.MethodSpec(name: "create", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController", type: .unary)
            public static let update = Connect.MethodSpec(name: "update", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController", type: .unary)
            public static let delete = Connect.MethodSpec(name: "delete", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController", type: .unary)
            public static let restore = Connect.MethodSpec(name: "restore", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController", type: .unary)
            public static let pause = Connect.MethodSpec(name: "pause", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController", type: .unary)
            public static let unpause = Connect.MethodSpec(name: "unpause", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController", type: .unary)
            public static let deleteNamespace = Connect.MethodSpec(name: "deleteNamespace", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController", type: .unary)
            public static let deletePod = Connect.MethodSpec(name: "deletePod", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterCommandController", type: .unary)
        }
    }
}

/// kube-cluster query controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterQueryControllerClientInterface {

    /// list all kube-clusters on planton cloud for the requested page. This is intended for use on portal.
    @available(iOS 13, *)
    func `list`(request: Cloud_Planton_Apis_V1_Commons_Rpc_Pagination_PageInfo, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterList>

    /// lookup kube-cluster using kube-cluster id
    @available(iOS 13, *)
    func `getByID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster>

    /// find kube-clusters by company id
    @available(iOS 13, *)
    func `findByCompanyID`(request: Cloud_Planton_Apis_V1_Resourcemanager_Company_Rpc_CompanyId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusters>

    /// find kube-clusters in a cloud account.
    @available(iOS 13, *)
    func `findByCloudAccountID`(request: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_Rpc_CloudAccountId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusters>

    /// find kube-clusters by company id to create environment.
    /// this will be used to populate drop down of kube-clusters in create environment form.
    /// the response should only include kube-clusters that a company is authorised to create environment.
    /// the authorization is verified by looking up kube-clusters with `company-environment-creator` relation with the company id provided in input.
    @available(iOS 13, *)
    func `findEnvironmentCreateKubeClusters`(request: Cloud_Planton_Apis_V1_Resourcemanager_Company_Rpc_CompanyId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusters>

    /// lookup components in a kube-cluster of a kube-cluster
    @available(iOS 13, *)
    func `getKubeClusterComponentsByKubeClusterID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterComponents>

    /// find workload namespaces in a kube-cluster.
    @available(iOS 13, *)
    func `findWorkloadNamespacesByKubeClusterID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_WorkloadNamespaces>

    /// find workload pods part of all environments hosted in a kube-cluster.
    @available(iOS 13, *)
    func `findWorkloadPodsByKubeClusterID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_WorkloadPods>

    /// find workload pods part of all environments hosted in a kube-cluster.
    @available(iOS 13, *)
    func `findSslCertificatesByKubeClusterID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Certificates>

    /// get a pod details
    @available(iOS 13, *)
    func `getPod`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_ByKubeClusterByNamespaceByPodInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod>

    /// get a log stream for a pod running in a kube-cluster kube-cluster
    @available(iOS 13, *)
    func `getPodLogStream`(headers: Connect.Headers) -> any Connect.ServerOnlyAsyncStreamInterface<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_ByKubeClusterByNamespaceByPodInput, Cloud_Planton_Apis_V1_Commons_Grpc_Stream_OutputLine>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterQueryControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterQueryControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterQueryControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `list`(request: Cloud_Planton_Apis_V1_Commons_Rpc_Pagination_PageInfo, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterList> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController/list", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getByID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController/getById", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findByCompanyID`(request: Cloud_Planton_Apis_V1_Resourcemanager_Company_Rpc_CompanyId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusters> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController/findByCompanyId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findByCloudAccountID`(request: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_Rpc_CloudAccountId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusters> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController/findByCloudAccountId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findEnvironmentCreateKubeClusters`(request: Cloud_Planton_Apis_V1_Resourcemanager_Company_Rpc_CompanyId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusters> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController/findEnvironmentCreateKubeClusters", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getKubeClusterComponentsByKubeClusterID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterComponents> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController/getKubeClusterComponentsByKubeClusterId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findWorkloadNamespacesByKubeClusterID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_WorkloadNamespaces> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController/findWorkloadNamespacesByKubeClusterId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findWorkloadPodsByKubeClusterID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_WorkloadPods> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController/findWorkloadPodsByKubeClusterId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findSslCertificatesByKubeClusterID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Certificates> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController/findSslCertificatesByKubeClusterId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getPod`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_ByKubeClusterByNamespaceByPodInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController/getPod", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getPodLogStream`(headers: Connect.Headers = [:]) -> any Connect.ServerOnlyAsyncStreamInterface<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_ByKubeClusterByNamespaceByPodInput, Cloud_Planton_Apis_V1_Commons_Grpc_Stream_OutputLine> {
        return self.client.serverOnlyStream(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController/getPodLogStream", headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let list = Connect.MethodSpec(name: "list", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController", type: .unary)
            public static let getByID = Connect.MethodSpec(name: "getById", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController", type: .unary)
            public static let findByCompanyID = Connect.MethodSpec(name: "findByCompanyId", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController", type: .unary)
            public static let findByCloudAccountID = Connect.MethodSpec(name: "findByCloudAccountId", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController", type: .unary)
            public static let findEnvironmentCreateKubeClusters = Connect.MethodSpec(name: "findEnvironmentCreateKubeClusters", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController", type: .unary)
            public static let getKubeClusterComponentsByKubeClusterID = Connect.MethodSpec(name: "getKubeClusterComponentsByKubeClusterId", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController", type: .unary)
            public static let findWorkloadNamespacesByKubeClusterID = Connect.MethodSpec(name: "findWorkloadNamespacesByKubeClusterId", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController", type: .unary)
            public static let findWorkloadPodsByKubeClusterID = Connect.MethodSpec(name: "findWorkloadPodsByKubeClusterId", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController", type: .unary)
            public static let findSslCertificatesByKubeClusterID = Connect.MethodSpec(name: "findSslCertificatesByKubeClusterId", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController", type: .unary)
            public static let getPod = Connect.MethodSpec(name: "getPod", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController", type: .unary)
            public static let getPodLogStream = Connect.MethodSpec(name: "getPodLogStream", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterQueryController", type: .serverStream)
        }
    }
}

/// gcp kube-cluster container node pool command controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterNodePoolGcpCommandControllerClientInterface {

    /// add a node pool to a kube-cluster in a kube-cluster
    @available(iOS 13, *)
    func `add`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_AddOrUpdateKubeClusterNodePoolGcpCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster>

    /// update a node pool of a kube-cluster in a kube-cluster
    @available(iOS 13, *)
    func `update`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_AddOrUpdateKubeClusterNodePoolGcpCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster>

    /// delete a node pool from a kube-cluster in a kube-cluster
    @available(iOS 13, *)
    func `delete`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_DeleteKubeClusterNodePoolGcpCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterNodePoolGcpCommandControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterNodePoolGcpCommandControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterNodePoolGcpCommandControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `add`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_AddOrUpdateKubeClusterNodePoolGcpCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterNodePoolGcpCommandController/add", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `update`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_AddOrUpdateKubeClusterNodePoolGcpCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterNodePoolGcpCommandController/update", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `delete`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_DeleteKubeClusterNodePoolGcpCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterNodePoolGcpCommandController/delete", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let add = Connect.MethodSpec(name: "add", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterNodePoolGcpCommandController", type: .unary)
            public static let update = Connect.MethodSpec(name: "update", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterNodePoolGcpCommandController", type: .unary)
            public static let delete = Connect.MethodSpec(name: "delete", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterNodePoolGcpCommandController", type: .unary)
        }
    }
}

/// gcp kube-cluster container node pool query controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterNodePoolGcpQueryControllerClientInterface {

    /// lookup gcp container node pool env using container-nodepool-id
    @available(iOS 13, *)
    func `getByGcpContainerNodePoolID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_GetByKubeClusterNodePoolGcpIdInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterNodePoolGcp>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterNodePoolGcpQueryControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterNodePoolGcpQueryControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterNodePoolGcpQueryControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `getByGcpContainerNodePoolID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_GetByKubeClusterNodePoolGcpIdInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterNodePoolGcp> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterNodePoolGcpQueryController/getByGcpContainerNodePoolId", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let getByGcpContainerNodePoolID = Connect.MethodSpec(name: "getByGcpContainerNodePoolId", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterNodePoolGcpQueryController", type: .unary)
        }
    }
}

/// gcp query controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_GcpQueryControllerClientInterface {

    /// list all gcp regions
    @available(iOS 13, *)
    func `findRegions`(request: Cloud_Planton_Apis_V1_Commons_Protobuf_Custom_CustomEmpty, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_GcpRegions>

    /// list all zones in a gcp region
    @available(iOS 13, *)
    func `findZonesByRegionIdentifier`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_GcpRegionIdentifier, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_GcpZones>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_GcpQueryControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_GcpQueryControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_GcpQueryControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `findRegions`(request: Cloud_Planton_Apis_V1_Commons_Protobuf_Custom_CustomEmpty, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_GcpRegions> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.GcpQueryController/findRegions", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findZonesByRegionIdentifier`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_GcpRegionIdentifier, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_GcpZones> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.GcpQueryController/findZonesByRegionIdentifier", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let findRegions = Connect.MethodSpec(name: "findRegions", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.GcpQueryController", type: .unary)
            public static let findZonesByRegionIdentifier = Connect.MethodSpec(name: "findZonesByRegionIdentifier", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.GcpQueryController", type: .unary)
        }
    }
}

/// kube-cluster stack controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterStackControllerClientInterface {

    ///preview kube-cluster stack with provided spec
    @available(iOS 13, *)
    func `preview`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster>

    ///apply kube-cluster stack
    @available(iOS 13, *)
    func `apply`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterStackControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterStackControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterStackControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `preview`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterStackController/preview", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `apply`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterStackController/apply", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let preview = Connect.MethodSpec(name: "preview", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterStackController", type: .unary)
            public static let apply = Connect.MethodSpec(name: "apply", service: "cloud.planton.apis.v1.code2cloud.deploy.kubecluster.rpc.KubeClusterStackController", type: .unary)
        }
    }
}
