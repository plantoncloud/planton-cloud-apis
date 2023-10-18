// Code generated by protoc-gen-connect-swift. DO NOT EDIT.
//
// Source: cloud/planton/apis/v1/code2cloud/deploy/kafka/rpc/service.proto
//

import Connect
import Foundation
import SwiftProtobuf

///kafka-cluster command controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterCommandControllerClientInterface {

    ///create kafka-cluster
    @available(iOS 13, *)
    func `create`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>

    ///update an existing kafka-cluster
    @available(iOS 13, *)
    func `update`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>

    ///delete an existing kafka-cluster
    @available(iOS 13, *)
    func `delete`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>

    ///restore a deleted kafka-cluster
    @available(iOS 13, *)
    func `restore`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>

    ///restart a kafka-cluster running in a environment.
    ///kafka-cluster is restarted by deleting running "broker" pods which will be automatically recreated by kubernetes
    ///note: zookeeper pods are not deleted.
    @available(iOS 13, *)
    func `restart`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>

    ///pause a kafka-cluster running in a environment.
    ///kafka-cluster is paused by scaling down number of replicas of
    ///the kubernetes deployment/stateful sets to zero in the environment.
    @available(iOS 13, *)
    func `pause`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>

    ///unpause a previously paused kafka-cluster running in a environment.
    ///unpause is done by scaling the number of pods back to the number of
    ///replicas configured for the kafka-cluster.
    @available(iOS 13, *)
    func `unpause`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterCommandControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterCommandControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterCommandControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `create`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController/create", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `update`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController/update", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `delete`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController/delete", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `restore`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController/restore", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `restart`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController/restart", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `pause`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController/pause", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `unpause`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController/unpause", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let create = Connect.MethodSpec(name: "create", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController", type: .unary)
            public static let update = Connect.MethodSpec(name: "update", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController", type: .unary)
            public static let delete = Connect.MethodSpec(name: "delete", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController", type: .unary)
            public static let restore = Connect.MethodSpec(name: "restore", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController", type: .unary)
            public static let restart = Connect.MethodSpec(name: "restart", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController", type: .unary)
            public static let pause = Connect.MethodSpec(name: "pause", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController", type: .unary)
            public static let unpause = Connect.MethodSpec(name: "unpause", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterCommandController", type: .unary)
        }
    }
}

///kafka-cluster query controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterQueryControllerClientInterface {

    ///list all kafka-clusters on planton cloud for the requested page. This is intended for use on portal.
    @available(iOS 13, *)
    func `list`(request: Cloud_Planton_Apis_V1_Commons_Rpc_Pagination_PageInfo, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterList>

    ///look up kafka-cluster using kafka-cluster id
    @available(iOS 13, *)
    func `getByID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>

    ///find kafka-clusters by product id.
    ///response contains only the resources that the authenticated user account has viewer access to.
    @available(iOS 13, *)
    func `findByProductID`(request: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusters>

    ///find kafka-clusters by environment
    @available(iOS 13, *)
    func `findByEnvironmentID`(request: Cloud_Planton_Apis_V1_Code2cloud_Environment_Rpc_EnvironmentId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusters>

    ///find kafka-clusters by kube-cluster
    @available(iOS 13, *)
    func `findByKubeClusterID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusters>

    ///look up kafka-cluster sasl password
    ///password is retrieved from the kubernetes cluster.
    @available(iOS 13, *)
    func `getPassword`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterPassword>

    ///lookup pods of a kafka-cluster deployed to a environment
    @available(iOS 13, *)
    func `findPods`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pods>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterQueryControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterQueryControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterQueryControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `list`(request: Cloud_Planton_Apis_V1_Commons_Rpc_Pagination_PageInfo, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterList> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController/list", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getByID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController/getById", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findByProductID`(request: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusters> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController/findByProductId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findByEnvironmentID`(request: Cloud_Planton_Apis_V1_Code2cloud_Environment_Rpc_EnvironmentId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusters> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController/findByEnvironmentId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findByKubeClusterID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kubecluster_Rpc_KubeClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusters> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController/findByKubeClusterId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getPassword`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterPassword> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController/getPassword", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findPods`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pods> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController/findPods", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let list = Connect.MethodSpec(name: "list", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController", type: .unary)
            public static let getByID = Connect.MethodSpec(name: "getById", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController", type: .unary)
            public static let findByProductID = Connect.MethodSpec(name: "findByProductId", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController", type: .unary)
            public static let findByEnvironmentID = Connect.MethodSpec(name: "findByEnvironmentId", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController", type: .unary)
            public static let findByKubeClusterID = Connect.MethodSpec(name: "findByKubeClusterId", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController", type: .unary)
            public static let getPassword = Connect.MethodSpec(name: "getPassword", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController", type: .unary)
            public static let findPods = Connect.MethodSpec(name: "findPods", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterQueryController", type: .unary)
        }
    }
}

///kafka-topic command controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopicCommandControllerClientInterface {

    ///add a single kafka topic to existing list of kafka topics of a kafka-cluster
    @available(iOS 13, *)
    func `add`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_AddOrUpdateKafkaTopicCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>

    ///add multiple kafka topics to existing list of kafka topics of a kafka-cluster
    @available(iOS 13, *)
    func `addMultiple`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_AddKafkaTopicsCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>

    ///update a kafka topic.
    @available(iOS 13, *)
    func `update`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_AddOrUpdateKafkaTopicCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>

    ///delete a kafka topic.
    @available(iOS 13, *)
    func `delete`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_DeleteOrRestoreKafkaTopicCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopicCommandControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopicCommandControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopicCommandControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `add`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_AddOrUpdateKafkaTopicCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaTopicCommandController/add", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `addMultiple`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_AddKafkaTopicsCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaTopicCommandController/addMultiple", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `update`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_AddOrUpdateKafkaTopicCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaTopicCommandController/update", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `delete`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_DeleteOrRestoreKafkaTopicCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaTopicCommandController/delete", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let add = Connect.MethodSpec(name: "add", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaTopicCommandController", type: .unary)
            public static let addMultiple = Connect.MethodSpec(name: "addMultiple", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaTopicCommandController", type: .unary)
            public static let update = Connect.MethodSpec(name: "update", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaTopicCommandController", type: .unary)
            public static let delete = Connect.MethodSpec(name: "delete", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaTopicCommandController", type: .unary)
        }
    }
}

///kafka-topic query controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopicQueryControllerClientInterface {

    ///find kafka topics by kafka-cluster id
    @available(iOS 13, *)
    func `findByKafkaClusterID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopics>

    ///look up kafka topic using kafka topic id
    @available(iOS 13, *)
    func `getByID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopicId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopic>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopicQueryControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopicQueryControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopicQueryControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `findByKafkaClusterID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopics> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaTopicQueryController/findByKafkaClusterId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getByID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopicId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaTopic> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaTopicQueryController/getById", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let findByKafkaClusterID = Connect.MethodSpec(name: "findByKafkaClusterId", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaTopicQueryController", type: .unary)
            public static let getByID = Connect.MethodSpec(name: "getById", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaTopicQueryController", type: .unary)
        }
    }
}

///kafka-cluster stack controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterStackControllerClientInterface {

    ///preview kafka-cluster stack
    @available(iOS 13, *)
    func `preview`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>

    ///apply kafka-cluster stack
    @available(iOS 13, *)
    func `apply`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterStackControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterStackControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterStackControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `preview`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterStackController/preview", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `apply`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaClusterId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Kafka_Rpc_KafkaCluster> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterStackController/apply", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let preview = Connect.MethodSpec(name: "preview", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterStackController", type: .unary)
            public static let apply = Connect.MethodSpec(name: "apply", service: "cloud.planton.apis.v1.code2cloud.deploy.kafka.rpc.KafkaClusterStackController", type: .unary)
        }
    }
}
