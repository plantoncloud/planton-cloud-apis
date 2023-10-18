// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/code2cloud/deploy/endpoint/custom/stack/kubernetes/model.proto
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

///custom-endpoint container stack input
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackInput {
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
  public var credentialsInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackCredentialsInput {
    get {return _credentialsInput ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackCredentialsInput()}
    set {_credentialsInput = newValue}
  }
  /// Returns true if `credentialsInput` has been explicitly set.
  public var hasCredentialsInput: Bool {return self._credentialsInput != nil}
  /// Clears the value of `credentialsInput`. Subsequent reads from it will return its default value.
  public mutating func clearCredentialsInput() {self._credentialsInput = nil}

  ///inputs used for creating stack resources
  public var resourceInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResourceInput {
    get {return _resourceInput ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResourceInput()}
    set {_resourceInput = newValue}
  }
  /// Returns true if `resourceInput` has been explicitly set.
  public var hasResourceInput: Bool {return self._resourceInput != nil}
  /// Clears the value of `resourceInput`. Subsequent reads from it will return its default value.
  public mutating func clearResourceInput() {self._resourceInput = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _stackJob: Cloud_Planton_Apis_V1_Stack_Rpc_StackJob? = nil
  fileprivate var _credentialsInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackCredentialsInput? = nil
  fileprivate var _resourceInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResourceInput? = nil
}

///custom-endpoint container stack credentials input
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackCredentialsInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///kubernetes provider credential for creating endpoint domain resources on container cluster
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

///custom-endpoint container stack resource input
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResourceInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///custom-endpoint
  public var customEndpoint: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_State_CustomEndpointState {
    get {return _customEndpoint ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_State_CustomEndpointState()}
    set {_customEndpoint = newValue}
  }
  /// Returns true if `customEndpoint` has been explicitly set.
  public var hasCustomEndpoint: Bool {return self._customEndpoint != nil}
  /// Clears the value of `customEndpoint`. Subsequent reads from it will return its default value.
  public mutating func clearCustomEndpoint() {self._customEndpoint = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _customEndpoint: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_State_CustomEndpointState? = nil
}

///custom-endpoint container stack outputs
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackOutputs {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///custom-endpoint
  public var customEndpointStatus: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_State_CustomEndpointStatusState {
    get {return _customEndpointStatus ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_State_CustomEndpointStatusState()}
    set {_customEndpointStatus = newValue}
  }
  /// Returns true if `customEndpointStatus` has been explicitly set.
  public var hasCustomEndpointStatus: Bool {return self._customEndpointStatus != nil}
  /// Clears the value of `customEndpointStatus`. Subsequent reads from it will return its default value.
  public mutating func clearCustomEndpointStatus() {self._customEndpointStatus = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _customEndpointStatus: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_State_CustomEndpointStatusState? = nil
}

///stack response
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///stack progress
  public var progress: Cloud_Planton_Apis_V1_Stack_Rpc_StackProgress {
    get {return _progress ?? Cloud_Planton_Apis_V1_Stack_Rpc_StackProgress()}
    set {_progress = newValue}
  }
  /// Returns true if `progress` has been explicitly set.
  public var hasProgress: Bool {return self._progress != nil}
  /// Clears the value of `progress`. Subsequent reads from it will return its default value.
  public mutating func clearProgress() {self._progress = nil}

  ///stack outputs
  public var outputs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackOutputs {
    get {return _outputs ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackOutputs()}
    set {_outputs = newValue}
  }
  /// Returns true if `outputs` has been explicitly set.
  public var hasOutputs: Bool {return self._outputs != nil}
  /// Clears the value of `outputs`. Subsequent reads from it will return its default value.
  public mutating func clearOutputs() {self._outputs = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _progress: Cloud_Planton_Apis_V1_Stack_Rpc_StackProgress? = nil
  fileprivate var _outputs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackOutputs? = nil
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackCredentialsInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResourceInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackOutputs: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResponse: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.stack.kubernetes"

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CustomEndpointKubernetesStackInput"
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackInput) -> Bool {
    if lhs._stackJob != rhs._stackJob {return false}
    if lhs._credentialsInput != rhs._credentialsInput {return false}
    if lhs._resourceInput != rhs._resourceInput {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackCredentialsInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CustomEndpointKubernetesStackCredentialsInput"
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackCredentialsInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackCredentialsInput) -> Bool {
    if lhs._kubernetes != rhs._kubernetes {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResourceInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CustomEndpointKubernetesStackResourceInput"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "custom_endpoint"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._customEndpoint) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._customEndpoint {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResourceInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResourceInput) -> Bool {
    if lhs._customEndpoint != rhs._customEndpoint {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackOutputs: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CustomEndpointKubernetesStackOutputs"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "custom_endpoint_status"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._customEndpointStatus) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._customEndpointStatus {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackOutputs, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackOutputs) -> Bool {
    if lhs._customEndpointStatus != rhs._customEndpointStatus {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CustomEndpointKubernetesStackResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "progress"),
    3: .same(proto: "outputs"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._progress) }()
      case 3: try { try decoder.decodeSingularMessageField(value: &self._outputs) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._progress {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try { if let v = self._outputs {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 3)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResponse, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Stack_Kubernetes_CustomEndpointKubernetesStackResponse) -> Bool {
    if lhs._progress != rhs._progress {return false}
    if lhs._outputs != rhs._outputs {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}