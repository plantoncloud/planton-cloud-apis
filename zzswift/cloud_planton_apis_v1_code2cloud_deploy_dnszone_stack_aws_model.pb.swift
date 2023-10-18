// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/code2cloud/deploy/dnszone/stack/aws/model.proto
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

/// input for dns-zone stack
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackInput {
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

  /// pulumi stack credentials
  public var credentialsInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackCredentialsInput {
    get {return _credentialsInput ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackCredentialsInput()}
    set {_credentialsInput = newValue}
  }
  /// Returns true if `credentialsInput` has been explicitly set.
  public var hasCredentialsInput: Bool {return self._credentialsInput != nil}
  /// Clears the value of `credentialsInput`. Subsequent reads from it will return its default value.
  public mutating func clearCredentialsInput() {self._credentialsInput = nil}

  /// inputs used for creating stack resources
  public var resourceInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResourceInput {
    get {return _resourceInput ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResourceInput()}
    set {_resourceInput = newValue}
  }
  /// Returns true if `resourceInput` has been explicitly set.
  public var hasResourceInput: Bool {return self._resourceInput != nil}
  /// Clears the value of `resourceInput`. Subsequent reads from it will return its default value.
  public mutating func clearResourceInput() {self._resourceInput = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _stackJob: Cloud_Planton_Apis_V1_Stack_Rpc_StackJob? = nil
  fileprivate var _credentialsInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackCredentialsInput? = nil
  fileprivate var _resourceInput: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResourceInput? = nil
}

/// stack credentials input
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackCredentialsInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///google provider credential input
  public var aws: Cloud_Planton_Apis_V1_Commons_Pulumi_Operation_Rpc_AwsProviderCredential {
    get {return _aws ?? Cloud_Planton_Apis_V1_Commons_Pulumi_Operation_Rpc_AwsProviderCredential()}
    set {_aws = newValue}
  }
  /// Returns true if `aws` has been explicitly set.
  public var hasAws: Bool {return self._aws != nil}
  /// Clears the value of `aws`. Subsequent reads from it will return its default value.
  public mutating func clearAws() {self._aws = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _aws: Cloud_Planton_Apis_V1_Commons_Pulumi_Operation_Rpc_AwsProviderCredential? = nil
}

/// stack resource input
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResourceInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// dns-zone to be added to the company
  public var dnsZone: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_State_DnsZoneState {
    get {return _dnsZone ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_State_DnsZoneState()}
    set {_dnsZone = newValue}
  }
  /// Returns true if `dnsZone` has been explicitly set.
  public var hasDnsZone: Bool {return self._dnsZone != nil}
  /// Clears the value of `dnsZone`. Subsequent reads from it will return its default value.
  public mutating func clearDnsZone() {self._dnsZone = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _dnsZone: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_State_DnsZoneState? = nil
}

///the stack output only supports domains and their nameservers
///the outputs exclude any dns records for the domains
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackOutputs {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// zone status populated with the details of the zone created in google cloud dns
  public var zoneStatus: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_State_DnsZoneStatusState {
    get {return _zoneStatus ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_State_DnsZoneStatusState()}
    set {_zoneStatus = newValue}
  }
  /// Returns true if `zoneStatus` has been explicitly set.
  public var hasZoneStatus: Bool {return self._zoneStatus != nil}
  /// Clears the value of `zoneStatus`. Subsequent reads from it will return its default value.
  public mutating func clearZoneStatus() {self._zoneStatus = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _zoneStatus: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_State_DnsZoneStatusState? = nil
}

///stack response
public struct Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResponse {
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
  public var outputs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackOutputs {
    get {return _outputs ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackOutputs()}
    set {_outputs = newValue}
  }
  /// Returns true if `outputs` has been explicitly set.
  public var hasOutputs: Bool {return self._outputs != nil}
  /// Clears the value of `outputs`. Subsequent reads from it will return its default value.
  public mutating func clearOutputs() {self._outputs = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _progress: Cloud_Planton_Apis_V1_Stack_Rpc_StackProgress? = nil
  fileprivate var _outputs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackOutputs? = nil
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackCredentialsInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResourceInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackOutputs: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResponse: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.code2cloud.deploy.dnszone.stack.aws"

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".DnsZoneAwsStackInput"
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackInput) -> Bool {
    if lhs._stackJob != rhs._stackJob {return false}
    if lhs._credentialsInput != rhs._credentialsInput {return false}
    if lhs._resourceInput != rhs._resourceInput {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackCredentialsInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".DnsZoneAwsStackCredentialsInput"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "aws"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._aws) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._aws {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackCredentialsInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackCredentialsInput) -> Bool {
    if lhs._aws != rhs._aws {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResourceInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".DnsZoneAwsStackResourceInput"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "dns_zone"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._dnsZone) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._dnsZone {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResourceInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResourceInput) -> Bool {
    if lhs._dnsZone != rhs._dnsZone {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackOutputs: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".DnsZoneAwsStackOutputs"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    2: .standard(proto: "zone_status"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 2: try { try decoder.decodeSingularMessageField(value: &self._zoneStatus) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._zoneStatus {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackOutputs, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackOutputs) -> Bool {
    if lhs._zoneStatus != rhs._zoneStatus {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".DnsZoneAwsStackResponse"
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResponse, rhs: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Dnszone_Stack_Aws_DnsZoneAwsStackResponse) -> Bool {
    if lhs._progress != rhs._progress {return false}
    if lhs._outputs != rhs._outputs {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
