// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/integration/kubernetes/resource/pod.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

/// kubernetes pod
public struct Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// namespace of the pod
  public var namespace: String = String()

  /// id of the pod
  public var podID: String = String()

  /// pod labels
  public var labels: Dictionary<String,String> = [:]

  /// status of the pod
  public var status: String = String()

  ///reason for the current status. this information is useful when pod in not in running state.
  public var statusReason: String = String()

  ///description of the reason for the current status. this information is useful when pod in not in running state.
  public var statusMessage: String = String()

  /// containers in the pod
  public var containers: [Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_PodContainer] = []

  ///value derived from containers for easy consumption by clients.
  ///displaying the the list of containers in a pod in <ready>/<total> format is useful. this attribute is added to
  ///make it easy to display the contents in that format.
  public var containersInReadyState: String = String()

  ///sum of the restart counts of each individual container in the pod.
  ///this attribute is being added to make it easy for user-facing clients to display this information.
  public var containersRestartCount: Int32 = 0

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///container in a kubernetes pod
public struct Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_PodContainer {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///name of the container
  public var name: String = String()

  ///image used for the container
  public var image: String = String()

  ///a container can be in running, waiting or terminated status.
  public var status: String = String()

  ///reason for the current status of the container.
  ///this is only relevant for waiting and terminated statuses.
  public var statusReason: String = String()

  ///description of the reason for the current status. this information is useful when pod in not in running state.
  public var statusMessage: String = String()

  ///number of times pod has restarted.
  public var restartCount: Int32 = 0

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

/// list of kubernetes pods
public struct Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pods {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var entries: [Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

/// workload pods
public struct Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_WorkloadPods {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///all pods with workload=microservice in a kube-cluster are included.
  public var microservicePods: [Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod] = []

  ///postgres cluster pods include all ancillary pods associated with the postgres cluster including entity-operator etc
  public var postgresClusterPods: [Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod] = []

  ///kafka cluster pods include all ancillary pods associated with the kafka cluster including schema-registry, kowl etc
  public var kafkaClusterPods: [Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod] = []

  ///all pods associated with solr-cloud deployments
  public var solrCloudPods: [Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod] = []

  ///all pods associated with redis-cluster deployments
  public var redisClusterPods: [Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_PodContainer: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pods: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_WorkloadPods: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.integration.kubernetes.resource"

extension Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".Pod"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "namespace"),
    2: .standard(proto: "pod_id"),
    3: .same(proto: "labels"),
    4: .same(proto: "status"),
    5: .standard(proto: "status_reason"),
    6: .standard(proto: "status_message"),
    7: .same(proto: "containers"),
    8: .standard(proto: "containers_in_ready_state"),
    9: .standard(proto: "containers_restart_count"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.namespace) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.podID) }()
      case 3: try { try decoder.decodeMapField(fieldType: SwiftProtobuf._ProtobufMap<SwiftProtobuf.ProtobufString,SwiftProtobuf.ProtobufString>.self, value: &self.labels) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.status) }()
      case 5: try { try decoder.decodeSingularStringField(value: &self.statusReason) }()
      case 6: try { try decoder.decodeSingularStringField(value: &self.statusMessage) }()
      case 7: try { try decoder.decodeRepeatedMessageField(value: &self.containers) }()
      case 8: try { try decoder.decodeSingularStringField(value: &self.containersInReadyState) }()
      case 9: try { try decoder.decodeSingularInt32Field(value: &self.containersRestartCount) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.namespace.isEmpty {
      try visitor.visitSingularStringField(value: self.namespace, fieldNumber: 1)
    }
    if !self.podID.isEmpty {
      try visitor.visitSingularStringField(value: self.podID, fieldNumber: 2)
    }
    if !self.labels.isEmpty {
      try visitor.visitMapField(fieldType: SwiftProtobuf._ProtobufMap<SwiftProtobuf.ProtobufString,SwiftProtobuf.ProtobufString>.self, value: self.labels, fieldNumber: 3)
    }
    if !self.status.isEmpty {
      try visitor.visitSingularStringField(value: self.status, fieldNumber: 4)
    }
    if !self.statusReason.isEmpty {
      try visitor.visitSingularStringField(value: self.statusReason, fieldNumber: 5)
    }
    if !self.statusMessage.isEmpty {
      try visitor.visitSingularStringField(value: self.statusMessage, fieldNumber: 6)
    }
    if !self.containers.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.containers, fieldNumber: 7)
    }
    if !self.containersInReadyState.isEmpty {
      try visitor.visitSingularStringField(value: self.containersInReadyState, fieldNumber: 8)
    }
    if self.containersRestartCount != 0 {
      try visitor.visitSingularInt32Field(value: self.containersRestartCount, fieldNumber: 9)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod, rhs: Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pod) -> Bool {
    if lhs.namespace != rhs.namespace {return false}
    if lhs.podID != rhs.podID {return false}
    if lhs.labels != rhs.labels {return false}
    if lhs.status != rhs.status {return false}
    if lhs.statusReason != rhs.statusReason {return false}
    if lhs.statusMessage != rhs.statusMessage {return false}
    if lhs.containers != rhs.containers {return false}
    if lhs.containersInReadyState != rhs.containersInReadyState {return false}
    if lhs.containersRestartCount != rhs.containersRestartCount {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_PodContainer: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".PodContainer"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "name"),
    2: .same(proto: "image"),
    3: .same(proto: "status"),
    4: .standard(proto: "status_reason"),
    5: .standard(proto: "status_message"),
    6: .standard(proto: "restart_count"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.name) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.image) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.status) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.statusReason) }()
      case 5: try { try decoder.decodeSingularStringField(value: &self.statusMessage) }()
      case 6: try { try decoder.decodeSingularInt32Field(value: &self.restartCount) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 1)
    }
    if !self.image.isEmpty {
      try visitor.visitSingularStringField(value: self.image, fieldNumber: 2)
    }
    if !self.status.isEmpty {
      try visitor.visitSingularStringField(value: self.status, fieldNumber: 3)
    }
    if !self.statusReason.isEmpty {
      try visitor.visitSingularStringField(value: self.statusReason, fieldNumber: 4)
    }
    if !self.statusMessage.isEmpty {
      try visitor.visitSingularStringField(value: self.statusMessage, fieldNumber: 5)
    }
    if self.restartCount != 0 {
      try visitor.visitSingularInt32Field(value: self.restartCount, fieldNumber: 6)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_PodContainer, rhs: Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_PodContainer) -> Bool {
    if lhs.name != rhs.name {return false}
    if lhs.image != rhs.image {return false}
    if lhs.status != rhs.status {return false}
    if lhs.statusReason != rhs.statusReason {return false}
    if lhs.statusMessage != rhs.statusMessage {return false}
    if lhs.restartCount != rhs.restartCount {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pods: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".Pods"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "entries"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeRepeatedMessageField(value: &self.entries) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.entries.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.entries, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pods, rhs: Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Pods) -> Bool {
    if lhs.entries != rhs.entries {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_WorkloadPods: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".WorkloadPods"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "microservice_pods"),
    2: .standard(proto: "postgres_cluster_pods"),
    3: .standard(proto: "kafka_cluster_pods"),
    4: .standard(proto: "solr_cloud_pods"),
    5: .standard(proto: "redis_cluster_pods"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeRepeatedMessageField(value: &self.microservicePods) }()
      case 2: try { try decoder.decodeRepeatedMessageField(value: &self.postgresClusterPods) }()
      case 3: try { try decoder.decodeRepeatedMessageField(value: &self.kafkaClusterPods) }()
      case 4: try { try decoder.decodeRepeatedMessageField(value: &self.solrCloudPods) }()
      case 5: try { try decoder.decodeRepeatedMessageField(value: &self.redisClusterPods) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.microservicePods.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.microservicePods, fieldNumber: 1)
    }
    if !self.postgresClusterPods.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.postgresClusterPods, fieldNumber: 2)
    }
    if !self.kafkaClusterPods.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.kafkaClusterPods, fieldNumber: 3)
    }
    if !self.solrCloudPods.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.solrCloudPods, fieldNumber: 4)
    }
    if !self.redisClusterPods.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.redisClusterPods, fieldNumber: 5)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_WorkloadPods, rhs: Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_WorkloadPods) -> Bool {
    if lhs.microservicePods != rhs.microservicePods {return false}
    if lhs.postgresClusterPods != rhs.postgresClusterPods {return false}
    if lhs.kafkaClusterPods != rhs.kafkaClusterPods {return false}
    if lhs.solrCloudPods != rhs.solrCloudPods {return false}
    if lhs.redisClusterPods != rhs.redisClusterPods {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
