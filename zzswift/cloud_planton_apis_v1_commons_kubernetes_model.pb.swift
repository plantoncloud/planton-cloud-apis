// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/commons/kubernetes/model.proto
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

///sidecar object in microservice deployment configuration
///this spec should always match the specification of a kubernetes container spec https://pkg.go.dev/k8s.io/api/core/v1#Container
///warning: sidecar feature does not currently support all features of a kubernetes container spec.
public struct Cloud_Planton_Apis_V1_Commons_Kubernetes_Container {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var name: String = String()

  public var image: String = String()

  public var ports: [Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerPort] = []

  public var resources: Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerResources {
    get {return _resources ?? Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerResources()}
    set {_resources = newValue}
  }
  /// Returns true if `resources` has been explicitly set.
  public var hasResources: Bool {return self._resources != nil}
  /// Clears the value of `resources`. Subsequent reads from it will return its default value.
  public mutating func clearResources() {self._resources = nil}

  public var env: [Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerEnvVar] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _resources: Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerResources? = nil
}

/// kubernetes container cpu and memory resources
public struct Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerResources {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// container resource limits
  /// key is either cpu or memory and value is the limits value for the resource
  public var limits: Cloud_Planton_Apis_V1_Commons_Kubernetes_CpuMemory {
    get {return _limits ?? Cloud_Planton_Apis_V1_Commons_Kubernetes_CpuMemory()}
    set {_limits = newValue}
  }
  /// Returns true if `limits` has been explicitly set.
  public var hasLimits: Bool {return self._limits != nil}
  /// Clears the value of `limits`. Subsequent reads from it will return its default value.
  public mutating func clearLimits() {self._limits = nil}

  /// container resource limits
  /// key is either cpu or memory and value is the requests value for the resource
  public var requests: Cloud_Planton_Apis_V1_Commons_Kubernetes_CpuMemory {
    get {return _requests ?? Cloud_Planton_Apis_V1_Commons_Kubernetes_CpuMemory()}
    set {_requests = newValue}
  }
  /// Returns true if `requests` has been explicitly set.
  public var hasRequests: Bool {return self._requests != nil}
  /// Clears the value of `requests`. Subsequent reads from it will return its default value.
  public mutating func clearRequests() {self._requests = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _limits: Cloud_Planton_Apis_V1_Commons_Kubernetes_CpuMemory? = nil
  fileprivate var _requests: Cloud_Planton_Apis_V1_Commons_Kubernetes_CpuMemory? = nil
}

public struct Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerEnvVar {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var name: String = String()

  public var value: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerPort {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var name: String = String()

  ///the attribute names must use camel case to marshal into https://pkg.go.dev/k8s.io/api/core/v1#Container
  public var containerPort: Int32 = 0

  public var `protocol`: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Planton_Apis_V1_Commons_Kubernetes_CpuMemory {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var cpu: String = String()

  public var memory: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerImage {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var repo: String = String()

  public var tag: String = String()

  public var pullSecretName: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Commons_Kubernetes_Container: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerResources: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerEnvVar: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerPort: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Commons_Kubernetes_CpuMemory: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerImage: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.commons.kubernetes"

extension Cloud_Planton_Apis_V1_Commons_Kubernetes_Container: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".Container"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "name"),
    2: .same(proto: "image"),
    3: .same(proto: "ports"),
    4: .same(proto: "resources"),
    5: .same(proto: "env"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.name) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.image) }()
      case 3: try { try decoder.decodeRepeatedMessageField(value: &self.ports) }()
      case 4: try { try decoder.decodeSingularMessageField(value: &self._resources) }()
      case 5: try { try decoder.decodeRepeatedMessageField(value: &self.env) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 1)
    }
    if !self.image.isEmpty {
      try visitor.visitSingularStringField(value: self.image, fieldNumber: 2)
    }
    if !self.ports.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.ports, fieldNumber: 3)
    }
    try { if let v = self._resources {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 4)
    } }()
    if !self.env.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.env, fieldNumber: 5)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Kubernetes_Container, rhs: Cloud_Planton_Apis_V1_Commons_Kubernetes_Container) -> Bool {
    if lhs.name != rhs.name {return false}
    if lhs.image != rhs.image {return false}
    if lhs.ports != rhs.ports {return false}
    if lhs._resources != rhs._resources {return false}
    if lhs.env != rhs.env {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerResources: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ContainerResources"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "limits"),
    2: .same(proto: "requests"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._limits) }()
      case 2: try { try decoder.decodeSingularMessageField(value: &self._requests) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._limits {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try { if let v = self._requests {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerResources, rhs: Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerResources) -> Bool {
    if lhs._limits != rhs._limits {return false}
    if lhs._requests != rhs._requests {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerEnvVar: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ContainerEnvVar"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "name"),
    2: .same(proto: "value"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.name) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.value) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 1)
    }
    if !self.value.isEmpty {
      try visitor.visitSingularStringField(value: self.value, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerEnvVar, rhs: Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerEnvVar) -> Bool {
    if lhs.name != rhs.name {return false}
    if lhs.value != rhs.value {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerPort: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ContainerPort"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "name"),
    2: .same(proto: "containerPort"),
    3: .same(proto: "protocol"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.name) }()
      case 2: try { try decoder.decodeSingularInt32Field(value: &self.containerPort) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.`protocol`) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 1)
    }
    if self.containerPort != 0 {
      try visitor.visitSingularInt32Field(value: self.containerPort, fieldNumber: 2)
    }
    if !self.`protocol`.isEmpty {
      try visitor.visitSingularStringField(value: self.`protocol`, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerPort, rhs: Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerPort) -> Bool {
    if lhs.name != rhs.name {return false}
    if lhs.containerPort != rhs.containerPort {return false}
    if lhs.`protocol` != rhs.`protocol` {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Commons_Kubernetes_CpuMemory: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CpuMemory"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "cpu"),
    2: .same(proto: "memory"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.cpu) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.memory) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.cpu.isEmpty {
      try visitor.visitSingularStringField(value: self.cpu, fieldNumber: 1)
    }
    if !self.memory.isEmpty {
      try visitor.visitSingularStringField(value: self.memory, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Kubernetes_CpuMemory, rhs: Cloud_Planton_Apis_V1_Commons_Kubernetes_CpuMemory) -> Bool {
    if lhs.cpu != rhs.cpu {return false}
    if lhs.memory != rhs.memory {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerImage: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ContainerImage"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "repo"),
    2: .same(proto: "tag"),
    3: .standard(proto: "pull_secret_name"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.repo) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.tag) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.pullSecretName) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.repo.isEmpty {
      try visitor.visitSingularStringField(value: self.repo, fieldNumber: 1)
    }
    if !self.tag.isEmpty {
      try visitor.visitSingularStringField(value: self.tag, fieldNumber: 2)
    }
    if !self.pullSecretName.isEmpty {
      try visitor.visitSingularStringField(value: self.pullSecretName, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerImage, rhs: Cloud_Planton_Apis_V1_Commons_Kubernetes_ContainerImage) -> Bool {
    if lhs.repo != rhs.repo {return false}
    if lhs.tag != rhs.tag {return false}
    if lhs.pullSecretName != rhs.pullSecretName {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}