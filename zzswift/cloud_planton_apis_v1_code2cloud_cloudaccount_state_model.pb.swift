// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/code2cloud/cloudaccount/state/model.proto
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

/// cloud account state
public struct Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountState {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var eventType: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_Enums_CloudAccountEventType {
    get {return _storage._eventType}
    set {_uniqueStorage()._eventType = newValue}
  }

  /// api version for the resource
  public var apiVersion: String {
    get {return _storage._apiVersion}
    set {_uniqueStorage()._apiVersion = newValue}
  }

  /// kind for the resource
  public var kind: String {
    get {return _storage._kind}
    set {_uniqueStorage()._kind = newValue}
  }

  /// metadata for the resource
  /// cloud-account id value adheres to the following constraints:
  /// 1. The ID must be prefixed with 'ca-<company_id>-', where <company_id> can vary in length.
  /// 2. The ID must not exceed 27 characters in length.
  /// 3. Once a cloud account is added to a company, the ID cannot be changed.
  /// By convention, the cloud_account_name is used as the suffix, but a different suffix may be used if desired.
  /// The constraints are enforced by the regular expression "^ca-.{1,22}$".
  public var metadata: Cloud_Planton_Apis_V1_Commons_Resource_Metadata {
    get {return _storage._metadata ?? Cloud_Planton_Apis_V1_Commons_Resource_Metadata()}
    set {_uniqueStorage()._metadata = newValue}
  }
  /// Returns true if `metadata` has been explicitly set.
  public var hasMetadata: Bool {return _storage._metadata != nil}
  /// Clears the value of `metadata`. Subsequent reads from it will return its default value.
  public mutating func clearMetadata() {_uniqueStorage()._metadata = nil}

  /// cloud account spec state
  public var spec: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountSpecState {
    get {return _storage._spec ?? Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountSpecState()}
    set {_uniqueStorage()._spec = newValue}
  }
  /// Returns true if `spec` has been explicitly set.
  public var hasSpec: Bool {return _storage._spec != nil}
  /// Clears the value of `spec`. Subsequent reads from it will return its default value.
  public mutating func clearSpec() {_uniqueStorage()._spec = nil}

  /// cloud account status state
  public var status: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountStatusState {
    get {return _storage._status ?? Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountStatusState()}
    set {_uniqueStorage()._status = newValue}
  }
  /// Returns true if `status` has been explicitly set.
  public var hasStatus: Bool {return _storage._status != nil}
  /// Clears the value of `status`. Subsequent reads from it will return its default value.
  public mutating func clearStatus() {_uniqueStorage()._status = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _storage = _StorageClass.defaultInstance
}

/// cloud account spec state
public struct Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountSpecState {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// company to which company cloud account belongs to.
  public var companyID: String = String()

  /// cloud provider.
  /// depending on the value of this attribute, cloud_account_<cloud_provider_name> object is required.
  public var cloudProvider: String = String()

  /// (conditionally required) if the value of cloud_provider is gcp then this attribute is required.
  /// this is only set if the cloud provider is gcp and is set to nil in other cases.
  public var gcp: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpSpecState {
    get {return _gcp ?? Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpSpecState()}
    set {_gcp = newValue}
  }
  /// Returns true if `gcp` has been explicitly set.
  public var hasGcp: Bool {return self._gcp != nil}
  /// Clears the value of `gcp`. Subsequent reads from it will return its default value.
  public mutating func clearGcp() {self._gcp = nil}

  /// (conditionally required) if the value of cloud_provider is aws then this attribute is required.
  /// this is only set if the cloud provider is aws and is set to nil in other cases.
  public var aws: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountAwsSpecState {
    get {return _aws ?? Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountAwsSpecState()}
    set {_aws = newValue}
  }
  /// Returns true if `aws` has been explicitly set.
  public var hasAws: Bool {return self._aws != nil}
  /// Clears the value of `aws`. Subsequent reads from it will return its default value.
  public mutating func clearAws() {self._aws = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _gcp: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpSpecState? = nil
  fileprivate var _aws: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountAwsSpecState? = nil
}

/// cloud account status state
public struct Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountStatusState {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// resource lifecycle
  public var lifecycle: Cloud_Planton_Apis_V1_Commons_Resource_ResourceLifecycle {
    get {return _lifecycle ?? Cloud_Planton_Apis_V1_Commons_Resource_ResourceLifecycle()}
    set {_lifecycle = newValue}
  }
  /// Returns true if `lifecycle` has been explicitly set.
  public var hasLifecycle: Bool {return self._lifecycle != nil}
  /// Clears the value of `lifecycle`. Subsequent reads from it will return its default value.
  public mutating func clearLifecycle() {self._lifecycle = nil}

  /// system audit info
  public var sysAudit: Cloud_Planton_Apis_V1_Commons_Audit_SysAudit {
    get {return _sysAudit ?? Cloud_Planton_Apis_V1_Commons_Audit_SysAudit()}
    set {_sysAudit = newValue}
  }
  /// Returns true if `sysAudit` has been explicitly set.
  public var hasSysAudit: Bool {return self._sysAudit != nil}
  /// Clears the value of `sysAudit`. Subsequent reads from it will return its default value.
  public mutating func clearSysAudit() {self._sysAudit = nil}

  /// id of the stack-job
  public var stackJobID: String = String()

  /// status of a gcp cloud-account.
  /// this is only populated when the parent cloud-account is a gcp cloud account and is set to nil otherwise.
  public var gcp: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpStatusState {
    get {return _gcp ?? Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpStatusState()}
    set {_gcp = newValue}
  }
  /// Returns true if `gcp` has been explicitly set.
  public var hasGcp: Bool {return self._gcp != nil}
  /// Clears the value of `gcp`. Subsequent reads from it will return its default value.
  public mutating func clearGcp() {self._gcp = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _lifecycle: Cloud_Planton_Apis_V1_Commons_Resource_ResourceLifecycle? = nil
  fileprivate var _sysAudit: Cloud_Planton_Apis_V1_Commons_Audit_SysAudit? = nil
  fileprivate var _gcp: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpStatusState? = nil
}

/// gcp cloud account spec state
public struct Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpSpecState {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// The ID of the Google Cloud Organization, required and immutable. For more information,
  /// visit: https://cloud.google.com/resource-manager/docs/creating-managing-organization
  public var orgID: String = String()

  /// The DNS domain of the Google Cloud Organization, optional during creation.
  public var orgDomain: String = String()

  /// The GCP Billing Account ID that will be linked to the projects created for shared services. Required during creation.
  public var billingAccountID: String = String()

  /// Name of the secret containing the value of the service-account key in the secrets manager.
  /// name of the secret is in the format `projects/<project-id>/secrets/<secret-name>`
  /// This is computed automatically by the Platon Cloud system.
  /// The value is derived from the company-status object
  public var serviceAccountKeySecretName: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

/// gcp cloud account status state
public struct Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpStatusState {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// The GCP Project created to host shared services resources (e.g., artifact-store, DNS managed zones).
  /// This is updated after the successful execution of a shared services stack.
  public var sharedServicesProject: Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Project_State_GcpProjectState {
    get {return _sharedServicesProject ?? Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Project_State_GcpProjectState()}
    set {_sharedServicesProject = newValue}
  }
  /// Returns true if `sharedServicesProject` has been explicitly set.
  public var hasSharedServicesProject: Bool {return self._sharedServicesProject != nil}
  /// Clears the value of `sharedServicesProject`. Subsequent reads from it will return its default value.
  public mutating func clearSharedServicesProject() {self._sharedServicesProject = nil}

  /// A GCP Folder is created within the GCP Organization for each Cloud Account added to Platon Cloud.
  /// Multiple Cloud Accounts can share the same Organization ID. Resources created for each Cloud Account are linked to this Folder.
  /// Each Cloud Account's children include the shared services project and one Folder per kube-cluster.
  public var cloudAccountFolder: Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_State_GcpFolderState {
    get {return _cloudAccountFolder ?? Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_State_GcpFolderState()}
    set {_cloudAccountFolder = newValue}
  }
  /// Returns true if `cloudAccountFolder` has been explicitly set.
  public var hasCloudAccountFolder: Bool {return self._cloudAccountFolder != nil}
  /// Clears the value of `cloudAccountFolder`. Subsequent reads from it will return its default value.
  public mutating func clearCloudAccountFolder() {self._cloudAccountFolder = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _sharedServicesProject: Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Project_State_GcpProjectState? = nil
  fileprivate var _cloudAccountFolder: Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_State_GcpFolderState? = nil
}

/// Specification for the Cloud Account in Amazon Web Services (AWS)
public struct Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountAwsSpecState {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// The ID of the AWS Account, required and immutable. For more information,
  public var accountID: String = String()

  ///AWS Access Key Id
  public var accessKeyID: String = String()

  ///AWS Secret Access Key
  public var secretAccessKey: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountState: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountSpecState: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountStatusState: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpSpecState: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpStatusState: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountAwsSpecState: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.code2cloud.cloudaccount.state"

extension Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountState: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CloudAccountState"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    99: .standard(proto: "event_type"),
    1: .standard(proto: "api_version"),
    2: .same(proto: "kind"),
    3: .same(proto: "metadata"),
    4: .same(proto: "spec"),
    5: .same(proto: "status"),
  ]

  fileprivate class _StorageClass {
    var _eventType: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_Enums_CloudAccountEventType = .unspecified
    var _apiVersion: String = String()
    var _kind: String = String()
    var _metadata: Cloud_Planton_Apis_V1_Commons_Resource_Metadata? = nil
    var _spec: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountSpecState? = nil
    var _status: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountStatusState? = nil

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
      _eventType = source._eventType
      _apiVersion = source._apiVersion
      _kind = source._kind
      _metadata = source._metadata
      _spec = source._spec
      _status = source._status
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
        case 1: try { try decoder.decodeSingularStringField(value: &_storage._apiVersion) }()
        case 2: try { try decoder.decodeSingularStringField(value: &_storage._kind) }()
        case 3: try { try decoder.decodeSingularMessageField(value: &_storage._metadata) }()
        case 4: try { try decoder.decodeSingularMessageField(value: &_storage._spec) }()
        case 5: try { try decoder.decodeSingularMessageField(value: &_storage._status) }()
        case 99: try { try decoder.decodeSingularEnumField(value: &_storage._eventType) }()
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
      if !_storage._apiVersion.isEmpty {
        try visitor.visitSingularStringField(value: _storage._apiVersion, fieldNumber: 1)
      }
      if !_storage._kind.isEmpty {
        try visitor.visitSingularStringField(value: _storage._kind, fieldNumber: 2)
      }
      try { if let v = _storage._metadata {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 3)
      } }()
      try { if let v = _storage._spec {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 4)
      } }()
      try { if let v = _storage._status {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 5)
      } }()
      if _storage._eventType != .unspecified {
        try visitor.visitSingularEnumField(value: _storage._eventType, fieldNumber: 99)
      }
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountState, rhs: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountState) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
        if _storage._eventType != rhs_storage._eventType {return false}
        if _storage._apiVersion != rhs_storage._apiVersion {return false}
        if _storage._kind != rhs_storage._kind {return false}
        if _storage._metadata != rhs_storage._metadata {return false}
        if _storage._spec != rhs_storage._spec {return false}
        if _storage._status != rhs_storage._status {return false}
        return true
      }
      if !storagesAreEqual {return false}
    }
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountSpecState: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CloudAccountSpecState"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "company_id"),
    2: .standard(proto: "cloud_provider"),
    3: .same(proto: "gcp"),
    4: .same(proto: "aws"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.companyID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.cloudProvider) }()
      case 3: try { try decoder.decodeSingularMessageField(value: &self._gcp) }()
      case 4: try { try decoder.decodeSingularMessageField(value: &self._aws) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.companyID.isEmpty {
      try visitor.visitSingularStringField(value: self.companyID, fieldNumber: 1)
    }
    if !self.cloudProvider.isEmpty {
      try visitor.visitSingularStringField(value: self.cloudProvider, fieldNumber: 2)
    }
    try { if let v = self._gcp {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 3)
    } }()
    try { if let v = self._aws {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 4)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountSpecState, rhs: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountSpecState) -> Bool {
    if lhs.companyID != rhs.companyID {return false}
    if lhs.cloudProvider != rhs.cloudProvider {return false}
    if lhs._gcp != rhs._gcp {return false}
    if lhs._aws != rhs._aws {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountStatusState: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CloudAccountStatusState"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    99: .same(proto: "lifecycle"),
    98: .standard(proto: "sys_audit"),
    97: .standard(proto: "stack_job_id"),
    1: .same(proto: "gcp"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._gcp) }()
      case 97: try { try decoder.decodeSingularStringField(value: &self.stackJobID) }()
      case 98: try { try decoder.decodeSingularMessageField(value: &self._sysAudit) }()
      case 99: try { try decoder.decodeSingularMessageField(value: &self._lifecycle) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._gcp {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    if !self.stackJobID.isEmpty {
      try visitor.visitSingularStringField(value: self.stackJobID, fieldNumber: 97)
    }
    try { if let v = self._sysAudit {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 98)
    } }()
    try { if let v = self._lifecycle {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 99)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountStatusState, rhs: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountStatusState) -> Bool {
    if lhs._lifecycle != rhs._lifecycle {return false}
    if lhs._sysAudit != rhs._sysAudit {return false}
    if lhs.stackJobID != rhs.stackJobID {return false}
    if lhs._gcp != rhs._gcp {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpSpecState: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CloudAccountGcpSpecState"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "org_id"),
    2: .standard(proto: "org_domain"),
    3: .standard(proto: "billing_account_id"),
    4: .standard(proto: "service_account_key_secret_name"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.orgID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.orgDomain) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.billingAccountID) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.serviceAccountKeySecretName) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.orgID.isEmpty {
      try visitor.visitSingularStringField(value: self.orgID, fieldNumber: 1)
    }
    if !self.orgDomain.isEmpty {
      try visitor.visitSingularStringField(value: self.orgDomain, fieldNumber: 2)
    }
    if !self.billingAccountID.isEmpty {
      try visitor.visitSingularStringField(value: self.billingAccountID, fieldNumber: 3)
    }
    if !self.serviceAccountKeySecretName.isEmpty {
      try visitor.visitSingularStringField(value: self.serviceAccountKeySecretName, fieldNumber: 4)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpSpecState, rhs: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpSpecState) -> Bool {
    if lhs.orgID != rhs.orgID {return false}
    if lhs.orgDomain != rhs.orgDomain {return false}
    if lhs.billingAccountID != rhs.billingAccountID {return false}
    if lhs.serviceAccountKeySecretName != rhs.serviceAccountKeySecretName {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpStatusState: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CloudAccountGcpStatusState"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "shared_services_project"),
    2: .standard(proto: "cloud_account_folder"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._sharedServicesProject) }()
      case 2: try { try decoder.decodeSingularMessageField(value: &self._cloudAccountFolder) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._sharedServicesProject {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try { if let v = self._cloudAccountFolder {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpStatusState, rhs: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountGcpStatusState) -> Bool {
    if lhs._sharedServicesProject != rhs._sharedServicesProject {return false}
    if lhs._cloudAccountFolder != rhs._cloudAccountFolder {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountAwsSpecState: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CloudAccountAwsSpecState"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "account_id"),
    2: .standard(proto: "access_key_id"),
    3: .standard(proto: "secret_access_key"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.accountID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.accessKeyID) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.secretAccessKey) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.accountID.isEmpty {
      try visitor.visitSingularStringField(value: self.accountID, fieldNumber: 1)
    }
    if !self.accessKeyID.isEmpty {
      try visitor.visitSingularStringField(value: self.accessKeyID, fieldNumber: 2)
    }
    if !self.secretAccessKey.isEmpty {
      try visitor.visitSingularStringField(value: self.secretAccessKey, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountAwsSpecState, rhs: Cloud_Planton_Apis_V1_Code2cloud_Cloudaccount_State_CloudAccountAwsSpecState) -> Bool {
    if lhs.accountID != rhs.accountID {return false}
    if lhs.accessKeyID != rhs.accessKeyID {return false}
    if lhs.secretAccessKey != rhs.secretAccessKey {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}