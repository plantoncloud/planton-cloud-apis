// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/code2cloud/deploy/redis/stack/kubernetes/model.proto
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

///input for redis-cluster stack
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///stack job
  public var stackJob: Cloud_Planton_Apis_V1_Stack_Rpc_StackJob {
    get {return _stackJob ?? Cloud_Planton_Apis_V1_Stack_Rpc_StackJob()}
    set {_stackJob = newValue}
  }
  /// Returns true if `stackJob` has been explicitly set.
  public var hasStackJob: Bool {return self._stackJob != nil}
  /// Clears the value of `stackJob`. Subsequent reads from it will return its default value.
  public mutating func clearStackJob() {self._stackJob = nil}

  ///pulumi stack credentials
  public var credentialsInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackCredentialsInput {
    get {return _credentialsInput ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackCredentialsInput()}
    set {_credentialsInput = newValue}
  }
  /// Returns true if `credentialsInput` has been explicitly set.
  public var hasCredentialsInput: Bool {return self._credentialsInput != nil}
  /// Clears the value of `credentialsInput`. Subsequent reads from it will return its default value.
  public mutating func clearCredentialsInput() {self._credentialsInput = nil}

  ///inputs used for creating stack resources
  public var resourceInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResourceInput {
    get {return _resourceInput ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResourceInput()}
    set {_resourceInput = newValue}
  }
  /// Returns true if `resourceInput` has been explicitly set.
  public var hasResourceInput: Bool {return self._resourceInput != nil}
  /// Clears the value of `resourceInput`. Subsequent reads from it will return its default value.
  public mutating func clearResourceInput() {self._resourceInput = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _stackJob: Cloud_Planton_Apis_V1_Stack_Rpc_StackJob? = nil
  fileprivate var _credentialsInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackCredentialsInput? = nil
  fileprivate var _resourceInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResourceInput? = nil
}

///stack credentials input
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackCredentialsInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///kubernetes provider credential for creating redis-cluster resources on container cluster
  public var kubernetes: Cloud_Planton_Apis_V1_Commons_Pulumi_Operation_Rpc_KubernetesProviderCredential {
    get {return _kubernetes ?? Cloud_Planton_Apis_V1_Commons_Pulumi_Operation_Rpc_KubernetesProviderCredential()}
    set {_kubernetes = newValue}
  }
  /// Returns true if `kubernetes` has been explicitly set.
  public var hasKubernetes: Bool {return self._kubernetes != nil}
  /// Clears the value of `kubernetes`. Subsequent reads from it will return its default value.
  public mutating func clearKubernetes() {self._kubernetes = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _kubernetes: Cloud_Planton_Apis_V1_Commons_Pulumi_Operation_Rpc_KubernetesProviderCredential? = nil
}

///stack resource input
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResourceInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///redis-cluster
  public var redisCluster: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_State_RedisClusterState {
    get {return _redisCluster ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_State_RedisClusterState()}
    set {_redisCluster = newValue}
  }
  /// Returns true if `redisCluster` has been explicitly set.
  public var hasRedisCluster: Bool {return self._redisCluster != nil}
  /// Clears the value of `redisCluster`. Subsequent reads from it will return its default value.
  public mutating func clearRedisCluster() {self._redisCluster = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _redisCluster: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_State_RedisClusterState? = nil
}

///redis-cluster stack outputs
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackOutputs {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///status of the redis-cluster upon stack-job
  public var redisClusterStatus: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_State_RedisClusterStatusState {
    get {return _redisClusterStatus ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_State_RedisClusterStatusState()}
    set {_redisClusterStatus = newValue}
  }
  /// Returns true if `redisClusterStatus` has been explicitly set.
  public var hasRedisClusterStatus: Bool {return self._redisClusterStatus != nil}
  /// Clears the value of `redisClusterStatus`. Subsequent reads from it will return its default value.
  public mutating func clearRedisClusterStatus() {self._redisClusterStatus = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _redisClusterStatus: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_State_RedisClusterStatusState? = nil
}

///stack response
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///stack progress
  public var progress: Cloud_Planton_Apis_V1_Stack_Rpc_StackProgress {
    get {return _storage._progress ?? Cloud_Planton_Apis_V1_Stack_Rpc_StackProgress()}
    set {_uniqueStorage()._progress = newValue}
  }
  /// Returns true if `progress` has been explicitly set.
  public var hasProgress: Bool {return _storage._progress != nil}
  /// Clears the value of `progress`. Subsequent reads from it will return its default value.
  public mutating func clearProgress() {_uniqueStorage()._progress = nil}

  ///stack outputs
  public var outputs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackOutputs {
    get {return _storage._outputs ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackOutputs()}
    set {_uniqueStorage()._outputs = newValue}
  }
  /// Returns true if `outputs` has been explicitly set.
  public var hasOutputs: Bool {return _storage._outputs != nil}
  /// Clears the value of `outputs`. Subsequent reads from it will return its default value.
  public mutating func clearOutputs() {_uniqueStorage()._outputs = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _storage = _StorageClass.defaultInstance
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackCredentialsInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResourceInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackOutputs: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResponse: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.code2cloud.deploy.redis.stack.kubernetes"

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".RedisClusterKubernetesStackInput"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "stack_job"),
    2: .standard(proto: "credentials_input"),
    3: .standard(proto: "resource_input"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._stackJob) }()
      case 2: try { try decoder.decodeSingularMessageField(value: &self._credentialsInput) }()
      case 3: try { try decoder.decodeSingularMessageField(value: &self._resourceInput) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._stackJob {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try { if let v = self._credentialsInput {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    } }()
    try { if let v = self._resourceInput {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 3)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackInput) -> Bool {
    if lhs._stackJob != rhs._stackJob {return false}
    if lhs._credentialsInput != rhs._credentialsInput {return false}
    if lhs._resourceInput != rhs._resourceInput {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackCredentialsInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".RedisClusterKubernetesStackCredentialsInput"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "kubernetes"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._kubernetes) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._kubernetes {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackCredentialsInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackCredentialsInput) -> Bool {
    if lhs._kubernetes != rhs._kubernetes {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResourceInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".RedisClusterKubernetesStackResourceInput"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "redis_cluster"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._redisCluster) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._redisCluster {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResourceInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResourceInput) -> Bool {
    if lhs._redisCluster != rhs._redisCluster {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackOutputs: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".RedisClusterKubernetesStackOutputs"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "redis_cluster_status"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._redisClusterStatus) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._redisClusterStatus {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackOutputs, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackOutputs) -> Bool {
    if lhs._redisClusterStatus != rhs._redisClusterStatus {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".RedisClusterKubernetesStackResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "progress"),
    3: .same(proto: "outputs"),
  ]

  fileprivate class _StorageClass {
    var _progress: Cloud_Planton_Apis_V1_Stack_Rpc_StackProgress? = nil
    var _outputs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackOutputs? = nil

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
      _progress = source._progress
      _outputs = source._outputs
    }
  }

  fileprivate mutating func _uniqueStorage() -> _StorageClass {
    if !isKnownUniquelyReferenced(&_storage) {
      _storage = _StorageClass(copying: _storage)
    }
    return _storage
  }

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    _ = _uniqueStorage()
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      while let fieldNumber = try decoder.nextFieldNumber() {
        // The use of inline closures is to circumvent an issue where the compiler
        // allocates stack space for every case branch when no optimizations are
        // enabled. https://github.com/apple/swift-protobuf/issues/1034
        switch fieldNumber {
        case 1: try { try decoder.decodeSingularMessageField(value: &_storage._progress) }()
        case 3: try { try decoder.decodeSingularMessageField(value: &_storage._outputs) }()
        default: break
        }
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every if/case branch local when no optimizations
      // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
      // https://github.com/apple/swift-protobuf/issues/1182
      try { if let v = _storage._progress {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
      } }()
      try { if let v = _storage._outputs {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 3)
      } }()
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResponse, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Redis_Stack_Kubernetes_RedisClusterKubernetesStackResponse) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
        if _storage._progress != rhs_storage._progress {return false}
        if _storage._outputs != rhs_storage._outputs {return false}
        return true
      }
      if !storagesAreEqual {return false}
    }
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
