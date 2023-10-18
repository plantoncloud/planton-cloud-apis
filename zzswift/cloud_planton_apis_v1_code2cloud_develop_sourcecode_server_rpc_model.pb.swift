// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/code2cloud/develop/sourcecode/server/rpc/model.proto
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

///code-server
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServer {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///resource api-version
  public var apiVersion: String {
    get {return _storage._apiVersion}
    set {_uniqueStorage()._apiVersion = newValue}
  }

  ///resource kind
  public var kind: String {
    get {return _storage._kind}
    set {_uniqueStorage()._kind = newValue}
  }

  ///resource metadata
  public var metadata: Cloud_Planton_Apis_V1_Commons_Resource_Metadata {
    get {return _storage._metadata ?? Cloud_Planton_Apis_V1_Commons_Resource_Metadata()}
    set {_uniqueStorage()._metadata = newValue}
  }
  /// Returns true if `metadata` has been explicitly set.
  public var hasMetadata: Bool {return _storage._metadata != nil}
  /// Clears the value of `metadata`. Subsequent reads from it will return its default value.
  public mutating func clearMetadata() {_uniqueStorage()._metadata = nil}

  ///spec
  public var spec: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerSpec {
    get {return _storage._spec ?? Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerSpec()}
    set {_uniqueStorage()._spec = newValue}
  }
  /// Returns true if `spec` has been explicitly set.
  public var hasSpec: Bool {return _storage._spec != nil}
  /// Clears the value of `spec`. Subsequent reads from it will return its default value.
  public mutating func clearSpec() {_uniqueStorage()._spec = nil}

  ///status
  public var status: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerStatus {
    get {return _storage._status ?? Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerStatus()}
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

///code-server spec
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerSpec {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///Specifies the company to which the code server belongs.
  ///This value is computed from the product.
  public var companyID: String = String()

  ///Specifies the product to which the code server belongs.
  public var productID: String = String()

  ///Specifies the host of the code server, such as "https://github.com".
  public var codeServerHost: String = String()

  ///Specifies the provider for the code server.
  public var codeServerProvider: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_Enums_CodeServerProvider = .unspecified

  ///github spec
  public var github: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGithubSpec {
    get {return _github ?? Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGithubSpec()}
    set {_github = newValue}
  }
  /// Returns true if `github` has been explicitly set.
  public var hasGithub: Bool {return self._github != nil}
  /// Clears the value of `github`. Subsequent reads from it will return its default value.
  public mutating func clearGithub() {self._github = nil}

  ///gitlab spec
  public var gitlab: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGitlabSpec {
    get {return _gitlab ?? Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGitlabSpec()}
    set {_gitlab = newValue}
  }
  /// Returns true if `gitlab` has been explicitly set.
  public var hasGitlab: Bool {return self._gitlab != nil}
  /// Clears the value of `gitlab`. Subsequent reads from it will return its default value.
  public mutating func clearGitlab() {self._gitlab = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _github: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGithubSpec? = nil
  fileprivate var _gitlab: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGitlabSpec? = nil
}

///github code-server spec
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGithubSpec {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///id of the github organization
  public var orgID: String = String()

  ///ID of the Github App installation provided by Github after successful app installation via OAuth flow.
  public var appInstallID: Int64 = 0

  ///Specifies the owner type for Github App installation: User or Org.
  public var appInstallOwnerType: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_Enums_GithubAppInstallOwnerType = .githubAppInstallationOwnerTypeUnspecified

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///gitlab code-server spec
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGitlabSpec {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///name or numerical identifier for gitlab group
  ///This value is primarily used in the code project synchronization process.
  public var groupID: String = String()

  ///Access token to integrate with Gitlab Server.
  ///This value is acquired by browser authorization flow when code server is added via Planton Cloud Console App.
  public var accessToken: String = String()

  ///Refresh token used to fetch a new access token when the current one expires.
  public var refreshToken: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///code-server status
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerStatus {
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

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _lifecycle: Cloud_Planton_Apis_V1_Commons_Resource_ResourceLifecycle? = nil
  fileprivate var _sysAudit: Cloud_Planton_Apis_V1_Commons_Audit_SysAudit? = nil
}

///wrapper for id field of postgres-cluster
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerId {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var value: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///list of postgres-clusters
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServers {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var entries: [Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServer] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///response for paginated query to list postgres-clusters
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerList {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var totalPages: Int32 = 0

  public var entries: [Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServer] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///input for command to attach a machine account on upstream code-server(github/gitlab)
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_AttachMachineAccountByCodeServerIdCommandInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///email of the machine account to be attached to the code-project
  public var machineAccountEmail: String = String()

  ///id of the code server to which the machine account is to be attached
  public var codeServerID: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServer: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerSpec: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGithubSpec: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGitlabSpec: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerStatus: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerId: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServers: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerList: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_AttachMachineAccountByCodeServerIdCommandInput: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.rpc"

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServer: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CodeServer"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "api_version"),
    2: .same(proto: "kind"),
    3: .same(proto: "metadata"),
    4: .same(proto: "spec"),
    5: .same(proto: "status"),
  ]

  fileprivate class _StorageClass {
    var _apiVersion: String = String()
    var _kind: String = String()
    var _metadata: Cloud_Planton_Apis_V1_Commons_Resource_Metadata? = nil
    var _spec: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerSpec? = nil
    var _status: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerStatus? = nil

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
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
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServer, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServer) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
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

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerSpec: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CodeServerSpec"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "company_id"),
    2: .standard(proto: "product_id"),
    3: .standard(proto: "code_server_host"),
    4: .standard(proto: "code_server_provider"),
    5: .same(proto: "github"),
    6: .same(proto: "gitlab"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.companyID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.productID) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.codeServerHost) }()
      case 4: try { try decoder.decodeSingularEnumField(value: &self.codeServerProvider) }()
      case 5: try { try decoder.decodeSingularMessageField(value: &self._github) }()
      case 6: try { try decoder.decodeSingularMessageField(value: &self._gitlab) }()
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
    if !self.productID.isEmpty {
      try visitor.visitSingularStringField(value: self.productID, fieldNumber: 2)
    }
    if !self.codeServerHost.isEmpty {
      try visitor.visitSingularStringField(value: self.codeServerHost, fieldNumber: 3)
    }
    if self.codeServerProvider != .unspecified {
      try visitor.visitSingularEnumField(value: self.codeServerProvider, fieldNumber: 4)
    }
    try { if let v = self._github {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 5)
    } }()
    try { if let v = self._gitlab {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 6)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerSpec, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerSpec) -> Bool {
    if lhs.companyID != rhs.companyID {return false}
    if lhs.productID != rhs.productID {return false}
    if lhs.codeServerHost != rhs.codeServerHost {return false}
    if lhs.codeServerProvider != rhs.codeServerProvider {return false}
    if lhs._github != rhs._github {return false}
    if lhs._gitlab != rhs._gitlab {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGithubSpec: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CodeServerGithubSpec"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "org_id"),
    2: .standard(proto: "app_install_id"),
    3: .standard(proto: "app_install_owner_type"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.orgID) }()
      case 2: try { try decoder.decodeSingularInt64Field(value: &self.appInstallID) }()
      case 3: try { try decoder.decodeSingularEnumField(value: &self.appInstallOwnerType) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.orgID.isEmpty {
      try visitor.visitSingularStringField(value: self.orgID, fieldNumber: 1)
    }
    if self.appInstallID != 0 {
      try visitor.visitSingularInt64Field(value: self.appInstallID, fieldNumber: 2)
    }
    if self.appInstallOwnerType != .githubAppInstallationOwnerTypeUnspecified {
      try visitor.visitSingularEnumField(value: self.appInstallOwnerType, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGithubSpec, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGithubSpec) -> Bool {
    if lhs.orgID != rhs.orgID {return false}
    if lhs.appInstallID != rhs.appInstallID {return false}
    if lhs.appInstallOwnerType != rhs.appInstallOwnerType {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGitlabSpec: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CodeServerGitlabSpec"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "group_id"),
    2: .standard(proto: "access_token"),
    3: .standard(proto: "refresh_token"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.groupID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.accessToken) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.refreshToken) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.groupID.isEmpty {
      try visitor.visitSingularStringField(value: self.groupID, fieldNumber: 1)
    }
    if !self.accessToken.isEmpty {
      try visitor.visitSingularStringField(value: self.accessToken, fieldNumber: 2)
    }
    if !self.refreshToken.isEmpty {
      try visitor.visitSingularStringField(value: self.refreshToken, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGitlabSpec, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerGitlabSpec) -> Bool {
    if lhs.groupID != rhs.groupID {return false}
    if lhs.accessToken != rhs.accessToken {return false}
    if lhs.refreshToken != rhs.refreshToken {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerStatus: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CodeServerStatus"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    99: .same(proto: "lifecycle"),
    98: .standard(proto: "sys_audit"),
    97: .standard(proto: "stack_job_id"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerStatus, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerStatus) -> Bool {
    if lhs._lifecycle != rhs._lifecycle {return false}
    if lhs._sysAudit != rhs._sysAudit {return false}
    if lhs.stackJobID != rhs.stackJobID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerId: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CodeServerId"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "value"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.value) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.value.isEmpty {
      try visitor.visitSingularStringField(value: self.value, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerId, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerId) -> Bool {
    if lhs.value != rhs.value {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServers: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CodeServers"
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServers, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServers) -> Bool {
    if lhs.entries != rhs.entries {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerList: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CodeServerList"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "total_pages"),
    2: .same(proto: "entries"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularInt32Field(value: &self.totalPages) }()
      case 2: try { try decoder.decodeRepeatedMessageField(value: &self.entries) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.totalPages != 0 {
      try visitor.visitSingularInt32Field(value: self.totalPages, fieldNumber: 1)
    }
    if !self.entries.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.entries, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerList, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServerList) -> Bool {
    if lhs.totalPages != rhs.totalPages {return false}
    if lhs.entries != rhs.entries {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_AttachMachineAccountByCodeServerIdCommandInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".AttachMachineAccountByCodeServerIdCommandInput"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "machine_account_email"),
    2: .standard(proto: "code_server_id"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.machineAccountEmail) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.codeServerID) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.machineAccountEmail.isEmpty {
      try visitor.visitSingularStringField(value: self.machineAccountEmail, fieldNumber: 1)
    }
    if !self.codeServerID.isEmpty {
      try visitor.visitSingularStringField(value: self.codeServerID, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_AttachMachineAccountByCodeServerIdCommandInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_AttachMachineAccountByCodeServerIdCommandInput) -> Bool {
    if lhs.machineAccountEmail != rhs.machineAccountEmail {return false}
    if lhs.codeServerID != rhs.codeServerID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
