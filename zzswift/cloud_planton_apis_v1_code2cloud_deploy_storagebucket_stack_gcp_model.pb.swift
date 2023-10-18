// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/code2cloud/deploy/storagebucket/stack/gcp/model.proto
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

///input for storage-bucket stack
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackInput {
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
  public var credentialsInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackCredentialsInput {
    get {return _credentialsInput ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackCredentialsInput()}
    set {_credentialsInput = newValue}
  }
  /// Returns true if `credentialsInput` has been explicitly set.
  public var hasCredentialsInput: Bool {return self._credentialsInput != nil}
  /// Clears the value of `credentialsInput`. Subsequent reads from it will return its default value.
  public mutating func clearCredentialsInput() {self._credentialsInput = nil}

  ///inputs used for creating stack resources
  public var resourceInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResourceInput {
    get {return _resourceInput ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResourceInput()}
    set {_resourceInput = newValue}
  }
  /// Returns true if `resourceInput` has been explicitly set.
  public var hasResourceInput: Bool {return self._resourceInput != nil}
  /// Clears the value of `resourceInput`. Subsequent reads from it will return its default value.
  public mutating func clearResourceInput() {self._resourceInput = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _stackJob: Cloud_Planton_Apis_V1_Stack_Rpc_StackJob? = nil
  fileprivate var _credentialsInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackCredentialsInput? = nil
  fileprivate var _resourceInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResourceInput? = nil
}

///stack credentials input
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackCredentialsInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///google provider credential input
  public var google: Cloud_Planton_Apis_V1_Commons_Pulumi_Operation_Rpc_GoogleProviderCredential {
    get {return _google ?? Cloud_Planton_Apis_V1_Commons_Pulumi_Operation_Rpc_GoogleProviderCredential()}
    set {_google = newValue}
  }
  /// Returns true if `google` has been explicitly set.
  public var hasGoogle: Bool {return self._google != nil}
  /// Clears the value of `google`. Subsequent reads from it will return its default value.
  public mutating func clearGoogle() {self._google = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _google: Cloud_Planton_Apis_V1_Commons_Pulumi_Operation_Rpc_GoogleProviderCredential? = nil
}

///stack resource input
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResourceInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///storage-bucket
  public var storageBucket: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_State_StorageBucketState {
    get {return _storageBucket ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_State_StorageBucketState()}
    set {_storageBucket = newValue}
  }
  /// Returns true if `storageBucket` has been explicitly set.
  public var hasStorageBucket: Bool {return self._storageBucket != nil}
  /// Clears the value of `storageBucket`. Subsequent reads from it will return its default value.
  public mutating func clearStorageBucket() {self._storageBucket = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _storageBucket: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_State_StorageBucketState? = nil
}

///storage-bucket stack outputs
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackOutputs {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///stack response
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResponse {
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
  public var outputs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackOutputs {
    get {return _outputs ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackOutputs()}
    set {_outputs = newValue}
  }
  /// Returns true if `outputs` has been explicitly set.
  public var hasOutputs: Bool {return self._outputs != nil}
  /// Clears the value of `outputs`. Subsequent reads from it will return its default value.
  public mutating func clearOutputs() {self._outputs = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _progress: Cloud_Planton_Apis_V1_Stack_Rpc_StackProgress? = nil
  fileprivate var _outputs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackOutputs? = nil
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackCredentialsInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResourceInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackOutputs: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResponse: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.code2cloud.deploy.storagebucket.stack.gcp"

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".StorageBucketGcpStackInput"
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackInput) -> Bool {
    if lhs._stackJob != rhs._stackJob {return false}
    if lhs._credentialsInput != rhs._credentialsInput {return false}
    if lhs._resourceInput != rhs._resourceInput {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackCredentialsInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".StorageBucketGcpStackCredentialsInput"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "google"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._google) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._google {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackCredentialsInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackCredentialsInput) -> Bool {
    if lhs._google != rhs._google {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResourceInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".StorageBucketGcpStackResourceInput"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "storage_bucket"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._storageBucket) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._storageBucket {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResourceInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResourceInput) -> Bool {
    if lhs._storageBucket != rhs._storageBucket {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackOutputs: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".StorageBucketGcpStackOutputs"
  public static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let _ = try decoder.nextFieldNumber() {
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackOutputs, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackOutputs) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".StorageBucketGcpStackResponse"
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResponse, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Storagebucket_Stack_Gcp_StorageBucketGcpStackResponse) -> Bool {
    if lhs._progress != rhs._progress {return false}
    if lhs._outputs != rhs._outputs {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}