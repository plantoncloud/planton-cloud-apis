// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/iam/identity/connection/state/model.proto
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

public struct Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionState {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// event_type is event that resulted in the new state which often is a result of processing a command
  public var eventType: Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_Enums_IdentityConnectionEventType = .unspecified

  /// resource api version
  public var apiVersion: String = String()

  /// resource kind
  public var kind: String = String()

  /// metadata for the resource
  /// id:
  ///
  /// naming convention "<idc>-<company_id>-<connection_name>".
  /// backend ignores the value provided by the client.
  ///
  /// name:
  ///
  /// a user preferred name of the identity connection.
  /// (important) spaces and special characters are not allowed except for hyphens.
  public var metadata: Cloud_Planton_Apis_V1_Commons_Resource_Metadata {
    get {return _metadata ?? Cloud_Planton_Apis_V1_Commons_Resource_Metadata()}
    set {_metadata = newValue}
  }
  /// Returns true if `metadata` has been explicitly set.
  public var hasMetadata: Bool {return self._metadata != nil}
  /// Clears the value of `metadata`. Subsequent reads from it will return its default value.
  public mutating func clearMetadata() {self._metadata = nil}

  /// resource spec
  public var spec: Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionSpecState {
    get {return _spec ?? Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionSpecState()}
    set {_spec = newValue}
  }
  /// Returns true if `spec` has been explicitly set.
  public var hasSpec: Bool {return self._spec != nil}
  /// Clears the value of `spec`. Subsequent reads from it will return its default value.
  public mutating func clearSpec() {self._spec = nil}

  /// resource status
  public var status: Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionStatusState {
    get {return _status ?? Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionStatusState()}
    set {_status = newValue}
  }
  /// Returns true if `status` has been explicitly set.
  public var hasStatus: Bool {return self._status != nil}
  /// Clears the value of `status`. Subsequent reads from it will return its default value.
  public mutating func clearStatus() {self._status = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _metadata: Cloud_Planton_Apis_V1_Commons_Resource_Metadata? = nil
  fileprivate var _spec: Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionSpecState? = nil
  fileprivate var _status: Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionStatusState? = nil
}

/// identity connection used for linking a company's identity provider with planton cloud.
public struct Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionSpecState {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// id of the company owning the identity connection.
  public var companyID: String {
    get {return _storage._companyID}
    set {_uniqueStorage()._companyID = newValue}
  }

  /// type of the identity connection.
  public var connectionType: String {
    get {return _storage._connectionType}
    set {_uniqueStorage()._connectionType = newValue}
  }

  /// a user preferred name of the identity connection.
  /// (important) spaces and special characters are not allowed except for hyphens.
  public var connectionName: String {
    get {return _storage._connectionName}
    set {_uniqueStorage()._connectionName = newValue}
  }

  /// (read-only) this is generated by planton cloud.
  /// naming convention "<dc>-<company_id>-<connection_name>.
  /// backend ignores the value provided by the client.
  public var identityConnectionID: String {
    get {return _storage._identityConnectionID}
    set {_uniqueStorage()._identityConnectionID = newValue}
  }

  /// (read-only) id assigned by auth0 upon successful creation of the enterprise connection.
  /// this value is required to query auth0 system to get the details of the connection.
  public var identityConnectionIDOnIdp: String {
    get {return _storage._identityConnectionIDOnIdp}
    set {_uniqueStorage()._identityConnectionIDOnIdp = newValue}
  }

  /// (read-only) url of the identity connection on idp.
  /// this is constructed when a client requests for the account object.
  /// the format of the url for machine accounts is https://manage.auth0.com/dashboard/us/<tenant>/connections/enterprise/<connection-type-on-auth0>/<connection-id-on-auth0>/settings
  /// ex: https://manage.auth0.com/dashboard/us/planton-pcs-dev/connections/enterprise/google-apps/con_DBlqRlQ8dsPCvZnj/settings
  public var idpURL: String {
    get {return _storage._idpURL}
    set {_uniqueStorage()._idpURL = newValue}
  }

  /// saml input for sign in url
  public var signInURL: String {
    get {return _storage._signInURL}
    set {_uniqueStorage()._signInURL = newValue}
  }

  /// saml input for x509_signing_cert_base64
  /// this is also used in ping federate
  public var x509SigningCertBase64: String {
    get {return _storage._x509SigningCertBase64}
    set {_uniqueStorage()._x509SigningCertBase64 = newValue}
  }

  /// saml input for enable_sign_out
  public var enableSignOut: Bool {
    get {return _storage._enableSignOut}
    set {_uniqueStorage()._enableSignOut = newValue}
  }

  /// saml input for sign_out_url
  public var signOutURL: String {
    get {return _storage._signOutURL}
    set {_uniqueStorage()._signOutURL = newValue}
  }

  /// saml input for user_id_attribute
  public var userIDAttribute: String {
    get {return _storage._userIDAttribute}
    set {_uniqueStorage()._userIDAttribute = newValue}
  }

  /// saml input for sign_request
  /// this is also used in ping federate
  public var signRequest: Bool {
    get {return _storage._signRequest}
    set {_uniqueStorage()._signRequest = newValue}
  }

  /// saml input for sign_request_algorithm
  /// this is also used in ping federate
  public var signRequestAlgorithm: String {
    get {return _storage._signRequestAlgorithm}
    set {_uniqueStorage()._signRequestAlgorithm = newValue}
  }

  /// saml input for sign_request_algorithm_digest
  /// this is also used in ping federate
  public var signRequestAlgorithmDigest: String {
    get {return _storage._signRequestAlgorithmDigest}
    set {_uniqueStorage()._signRequestAlgorithmDigest = newValue}
  }

  /// saml input for protocol_binding
  public var protocolBinding: String {
    get {return _storage._protocolBinding}
    set {_uniqueStorage()._protocolBinding = newValue}
  }

  /// ping federate input for server_url
  public var pingFederateServerURL: String {
    get {return _storage._pingFederateServerURL}
    set {_uniqueStorage()._pingFederateServerURL = newValue}
  }

  /// ldap connection input for idp_domains
  public var idpDomains: [String] {
    get {return _storage._idpDomains}
    set {_uniqueStorage()._idpDomains = newValue}
  }

  /// ldap connection input for is_cache_disabled
  public var isCacheDisabled: Bool {
    get {return _storage._isCacheDisabled}
    set {_uniqueStorage()._isCacheDisabled = newValue}
  }

  /// ldap connection input for is_client_ssl_certificate_authentication_required
  public var isClientSslCertificateAuthenticationRequired: Bool {
    get {return _storage._isClientSslCertificateAuthenticationRequired}
    set {_uniqueStorage()._isClientSslCertificateAuthenticationRequired = newValue}
  }

  /// ldap connection input for is_kerberos_enabled
  public var isKerberosEnabled: Bool {
    get {return _storage._isKerberosEnabled}
    set {_uniqueStorage()._isKerberosEnabled = newValue}
  }

  /// open id connection input for issue_url
  public var issueURL: String {
    get {return _storage._issueURL}
    set {_uniqueStorage()._issueURL = newValue}
  }

  /// open id connection input for client_id
  /// this is also used in okta workforce / AzureActiveDirectoryNativeConnection
  public var clientID: String {
    get {return _storage._clientID}
    set {_uniqueStorage()._clientID = newValue}
  }

  /// open id connection input for callback_url
  /// this is also used in okta workforce
  public var callbackURL: String {
    get {return _storage._callbackURL}
    set {_uniqueStorage()._callbackURL = newValue}
  }

  /// adfs connection input for adfs_metadata_url
  public var adfsMetadataURL: String {
    get {return _storage._adfsMetadataURL}
    set {_uniqueStorage()._adfsMetadataURL = newValue}
  }

  /// adfs connection input for adfs_metadata_federation_file
  public var adfsMetadataFederationFileBase64: String {
    get {return _storage._adfsMetadataFederationFileBase64}
    set {_uniqueStorage()._adfsMetadataFederationFileBase64 = newValue}
  }

  /// okta workforce input for okta domain
  public var oktaDomain: String {
    get {return _storage._oktaDomain}
    set {_uniqueStorage()._oktaDomain = newValue}
  }

  /// okta workforce input for client_secret
  /// this is also used in AzureActiveDirectoryNativeConnection
  public var clientSecret: String {
    get {return _storage._clientSecret}
    set {_uniqueStorage()._clientSecret = newValue}
  }

  /// azure active directory native connection input for microsoft_azure_ad_domain
  public var microsoftAzureAdDomain: String {
    get {return _storage._microsoftAzureAdDomain}
    set {_uniqueStorage()._microsoftAzureAdDomain = newValue}
  }

  /// azure active directory native connection input for use_common_endpoint
  public var useCommonEndpoint: Bool {
    get {return _storage._useCommonEndpoint}
    set {_uniqueStorage()._useCommonEndpoint = newValue}
  }

  /// azure active directory native connection input for identity_api
  public var identityApi: String {
    get {return _storage._identityApi}
    set {_uniqueStorage()._identityApi = newValue}
  }

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _storage = _StorageClass.defaultInstance
}

///status for identity-connection
public struct Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionStatusState {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// lifecycle of the resource
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

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _lifecycle: Cloud_Planton_Apis_V1_Commons_Resource_ResourceLifecycle? = nil
  fileprivate var _sysAudit: Cloud_Planton_Apis_V1_Commons_Audit_SysAudit? = nil
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionState: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionSpecState: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionStatusState: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.iam.identity.connection.state"

extension Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionState: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".IdentityConnectionState"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    99: .standard(proto: "event_type"),
    1: .standard(proto: "api_version"),
    2: .same(proto: "kind"),
    3: .same(proto: "metadata"),
    4: .same(proto: "spec"),
    5: .same(proto: "status"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.apiVersion) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.kind) }()
      case 3: try { try decoder.decodeSingularMessageField(value: &self._metadata) }()
      case 4: try { try decoder.decodeSingularMessageField(value: &self._spec) }()
      case 5: try { try decoder.decodeSingularMessageField(value: &self._status) }()
      case 99: try { try decoder.decodeSingularEnumField(value: &self.eventType) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.apiVersion.isEmpty {
      try visitor.visitSingularStringField(value: self.apiVersion, fieldNumber: 1)
    }
    if !self.kind.isEmpty {
      try visitor.visitSingularStringField(value: self.kind, fieldNumber: 2)
    }
    try { if let v = self._metadata {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 3)
    } }()
    try { if let v = self._spec {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 4)
    } }()
    try { if let v = self._status {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 5)
    } }()
    if self.eventType != .unspecified {
      try visitor.visitSingularEnumField(value: self.eventType, fieldNumber: 99)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionState, rhs: Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionState) -> Bool {
    if lhs.eventType != rhs.eventType {return false}
    if lhs.apiVersion != rhs.apiVersion {return false}
    if lhs.kind != rhs.kind {return false}
    if lhs._metadata != rhs._metadata {return false}
    if lhs._spec != rhs._spec {return false}
    if lhs._status != rhs._status {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionSpecState: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".IdentityConnectionSpecState"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "company_id"),
    2: .standard(proto: "connection_type"),
    3: .standard(proto: "connection_name"),
    4: .standard(proto: "identity_connection_id"),
    5: .standard(proto: "identity_connection_id_on_idp"),
    6: .standard(proto: "idp_url"),
    7: .standard(proto: "sign_in_url"),
    8: .standard(proto: "x509_signing_cert_base64"),
    9: .standard(proto: "enable_sign_out"),
    10: .standard(proto: "sign_out_url"),
    11: .standard(proto: "user_id_attribute"),
    12: .standard(proto: "sign_request"),
    13: .standard(proto: "sign_request_algorithm"),
    14: .standard(proto: "sign_request_algorithm_digest"),
    15: .standard(proto: "protocol_binding"),
    16: .standard(proto: "ping_federate_server_url"),
    17: .standard(proto: "idp_domains"),
    18: .standard(proto: "is_cache_disabled"),
    19: .standard(proto: "is_client_ssl_certificate_authentication_required"),
    20: .standard(proto: "is_kerberos_enabled"),
    21: .standard(proto: "issue_url"),
    22: .standard(proto: "client_id"),
    23: .standard(proto: "callback_url"),
    24: .standard(proto: "adfs_metadata_url"),
    25: .standard(proto: "adfs_metadata_federation_file_base64"),
    26: .standard(proto: "okta_domain"),
    27: .standard(proto: "client_secret"),
    28: .standard(proto: "microsoft_azure_ad_domain"),
    29: .standard(proto: "use_common_endpoint"),
    30: .standard(proto: "identity_api"),
  ]

  fileprivate class _StorageClass {
    var _companyID: String = String()
    var _connectionType: String = String()
    var _connectionName: String = String()
    var _identityConnectionID: String = String()
    var _identityConnectionIDOnIdp: String = String()
    var _idpURL: String = String()
    var _signInURL: String = String()
    var _x509SigningCertBase64: String = String()
    var _enableSignOut: Bool = false
    var _signOutURL: String = String()
    var _userIDAttribute: String = String()
    var _signRequest: Bool = false
    var _signRequestAlgorithm: String = String()
    var _signRequestAlgorithmDigest: String = String()
    var _protocolBinding: String = String()
    var _pingFederateServerURL: String = String()
    var _idpDomains: [String] = []
    var _isCacheDisabled: Bool = false
    var _isClientSslCertificateAuthenticationRequired: Bool = false
    var _isKerberosEnabled: Bool = false
    var _issueURL: String = String()
    var _clientID: String = String()
    var _callbackURL: String = String()
    var _adfsMetadataURL: String = String()
    var _adfsMetadataFederationFileBase64: String = String()
    var _oktaDomain: String = String()
    var _clientSecret: String = String()
    var _microsoftAzureAdDomain: String = String()
    var _useCommonEndpoint: Bool = false
    var _identityApi: String = String()

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
      _companyID = source._companyID
      _connectionType = source._connectionType
      _connectionName = source._connectionName
      _identityConnectionID = source._identityConnectionID
      _identityConnectionIDOnIdp = source._identityConnectionIDOnIdp
      _idpURL = source._idpURL
      _signInURL = source._signInURL
      _x509SigningCertBase64 = source._x509SigningCertBase64
      _enableSignOut = source._enableSignOut
      _signOutURL = source._signOutURL
      _userIDAttribute = source._userIDAttribute
      _signRequest = source._signRequest
      _signRequestAlgorithm = source._signRequestAlgorithm
      _signRequestAlgorithmDigest = source._signRequestAlgorithmDigest
      _protocolBinding = source._protocolBinding
      _pingFederateServerURL = source._pingFederateServerURL
      _idpDomains = source._idpDomains
      _isCacheDisabled = source._isCacheDisabled
      _isClientSslCertificateAuthenticationRequired = source._isClientSslCertificateAuthenticationRequired
      _isKerberosEnabled = source._isKerberosEnabled
      _issueURL = source._issueURL
      _clientID = source._clientID
      _callbackURL = source._callbackURL
      _adfsMetadataURL = source._adfsMetadataURL
      _adfsMetadataFederationFileBase64 = source._adfsMetadataFederationFileBase64
      _oktaDomain = source._oktaDomain
      _clientSecret = source._clientSecret
      _microsoftAzureAdDomain = source._microsoftAzureAdDomain
      _useCommonEndpoint = source._useCommonEndpoint
      _identityApi = source._identityApi
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
        case 1: try { try decoder.decodeSingularStringField(value: &_storage._companyID) }()
        case 2: try { try decoder.decodeSingularStringField(value: &_storage._connectionType) }()
        case 3: try { try decoder.decodeSingularStringField(value: &_storage._connectionName) }()
        case 4: try { try decoder.decodeSingularStringField(value: &_storage._identityConnectionID) }()
        case 5: try { try decoder.decodeSingularStringField(value: &_storage._identityConnectionIDOnIdp) }()
        case 6: try { try decoder.decodeSingularStringField(value: &_storage._idpURL) }()
        case 7: try { try decoder.decodeSingularStringField(value: &_storage._signInURL) }()
        case 8: try { try decoder.decodeSingularStringField(value: &_storage._x509SigningCertBase64) }()
        case 9: try { try decoder.decodeSingularBoolField(value: &_storage._enableSignOut) }()
        case 10: try { try decoder.decodeSingularStringField(value: &_storage._signOutURL) }()
        case 11: try { try decoder.decodeSingularStringField(value: &_storage._userIDAttribute) }()
        case 12: try { try decoder.decodeSingularBoolField(value: &_storage._signRequest) }()
        case 13: try { try decoder.decodeSingularStringField(value: &_storage._signRequestAlgorithm) }()
        case 14: try { try decoder.decodeSingularStringField(value: &_storage._signRequestAlgorithmDigest) }()
        case 15: try { try decoder.decodeSingularStringField(value: &_storage._protocolBinding) }()
        case 16: try { try decoder.decodeSingularStringField(value: &_storage._pingFederateServerURL) }()
        case 17: try { try decoder.decodeRepeatedStringField(value: &_storage._idpDomains) }()
        case 18: try { try decoder.decodeSingularBoolField(value: &_storage._isCacheDisabled) }()
        case 19: try { try decoder.decodeSingularBoolField(value: &_storage._isClientSslCertificateAuthenticationRequired) }()
        case 20: try { try decoder.decodeSingularBoolField(value: &_storage._isKerberosEnabled) }()
        case 21: try { try decoder.decodeSingularStringField(value: &_storage._issueURL) }()
        case 22: try { try decoder.decodeSingularStringField(value: &_storage._clientID) }()
        case 23: try { try decoder.decodeSingularStringField(value: &_storage._callbackURL) }()
        case 24: try { try decoder.decodeSingularStringField(value: &_storage._adfsMetadataURL) }()
        case 25: try { try decoder.decodeSingularStringField(value: &_storage._adfsMetadataFederationFileBase64) }()
        case 26: try { try decoder.decodeSingularStringField(value: &_storage._oktaDomain) }()
        case 27: try { try decoder.decodeSingularStringField(value: &_storage._clientSecret) }()
        case 28: try { try decoder.decodeSingularStringField(value: &_storage._microsoftAzureAdDomain) }()
        case 29: try { try decoder.decodeSingularBoolField(value: &_storage._useCommonEndpoint) }()
        case 30: try { try decoder.decodeSingularStringField(value: &_storage._identityApi) }()
        default: break
        }
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      if !_storage._companyID.isEmpty {
        try visitor.visitSingularStringField(value: _storage._companyID, fieldNumber: 1)
      }
      if !_storage._connectionType.isEmpty {
        try visitor.visitSingularStringField(value: _storage._connectionType, fieldNumber: 2)
      }
      if !_storage._connectionName.isEmpty {
        try visitor.visitSingularStringField(value: _storage._connectionName, fieldNumber: 3)
      }
      if !_storage._identityConnectionID.isEmpty {
        try visitor.visitSingularStringField(value: _storage._identityConnectionID, fieldNumber: 4)
      }
      if !_storage._identityConnectionIDOnIdp.isEmpty {
        try visitor.visitSingularStringField(value: _storage._identityConnectionIDOnIdp, fieldNumber: 5)
      }
      if !_storage._idpURL.isEmpty {
        try visitor.visitSingularStringField(value: _storage._idpURL, fieldNumber: 6)
      }
      if !_storage._signInURL.isEmpty {
        try visitor.visitSingularStringField(value: _storage._signInURL, fieldNumber: 7)
      }
      if !_storage._x509SigningCertBase64.isEmpty {
        try visitor.visitSingularStringField(value: _storage._x509SigningCertBase64, fieldNumber: 8)
      }
      if _storage._enableSignOut != false {
        try visitor.visitSingularBoolField(value: _storage._enableSignOut, fieldNumber: 9)
      }
      if !_storage._signOutURL.isEmpty {
        try visitor.visitSingularStringField(value: _storage._signOutURL, fieldNumber: 10)
      }
      if !_storage._userIDAttribute.isEmpty {
        try visitor.visitSingularStringField(value: _storage._userIDAttribute, fieldNumber: 11)
      }
      if _storage._signRequest != false {
        try visitor.visitSingularBoolField(value: _storage._signRequest, fieldNumber: 12)
      }
      if !_storage._signRequestAlgorithm.isEmpty {
        try visitor.visitSingularStringField(value: _storage._signRequestAlgorithm, fieldNumber: 13)
      }
      if !_storage._signRequestAlgorithmDigest.isEmpty {
        try visitor.visitSingularStringField(value: _storage._signRequestAlgorithmDigest, fieldNumber: 14)
      }
      if !_storage._protocolBinding.isEmpty {
        try visitor.visitSingularStringField(value: _storage._protocolBinding, fieldNumber: 15)
      }
      if !_storage._pingFederateServerURL.isEmpty {
        try visitor.visitSingularStringField(value: _storage._pingFederateServerURL, fieldNumber: 16)
      }
      if !_storage._idpDomains.isEmpty {
        try visitor.visitRepeatedStringField(value: _storage._idpDomains, fieldNumber: 17)
      }
      if _storage._isCacheDisabled != false {
        try visitor.visitSingularBoolField(value: _storage._isCacheDisabled, fieldNumber: 18)
      }
      if _storage._isClientSslCertificateAuthenticationRequired != false {
        try visitor.visitSingularBoolField(value: _storage._isClientSslCertificateAuthenticationRequired, fieldNumber: 19)
      }
      if _storage._isKerberosEnabled != false {
        try visitor.visitSingularBoolField(value: _storage._isKerberosEnabled, fieldNumber: 20)
      }
      if !_storage._issueURL.isEmpty {
        try visitor.visitSingularStringField(value: _storage._issueURL, fieldNumber: 21)
      }
      if !_storage._clientID.isEmpty {
        try visitor.visitSingularStringField(value: _storage._clientID, fieldNumber: 22)
      }
      if !_storage._callbackURL.isEmpty {
        try visitor.visitSingularStringField(value: _storage._callbackURL, fieldNumber: 23)
      }
      if !_storage._adfsMetadataURL.isEmpty {
        try visitor.visitSingularStringField(value: _storage._adfsMetadataURL, fieldNumber: 24)
      }
      if !_storage._adfsMetadataFederationFileBase64.isEmpty {
        try visitor.visitSingularStringField(value: _storage._adfsMetadataFederationFileBase64, fieldNumber: 25)
      }
      if !_storage._oktaDomain.isEmpty {
        try visitor.visitSingularStringField(value: _storage._oktaDomain, fieldNumber: 26)
      }
      if !_storage._clientSecret.isEmpty {
        try visitor.visitSingularStringField(value: _storage._clientSecret, fieldNumber: 27)
      }
      if !_storage._microsoftAzureAdDomain.isEmpty {
        try visitor.visitSingularStringField(value: _storage._microsoftAzureAdDomain, fieldNumber: 28)
      }
      if _storage._useCommonEndpoint != false {
        try visitor.visitSingularBoolField(value: _storage._useCommonEndpoint, fieldNumber: 29)
      }
      if !_storage._identityApi.isEmpty {
        try visitor.visitSingularStringField(value: _storage._identityApi, fieldNumber: 30)
      }
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionSpecState, rhs: Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionSpecState) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
        if _storage._companyID != rhs_storage._companyID {return false}
        if _storage._connectionType != rhs_storage._connectionType {return false}
        if _storage._connectionName != rhs_storage._connectionName {return false}
        if _storage._identityConnectionID != rhs_storage._identityConnectionID {return false}
        if _storage._identityConnectionIDOnIdp != rhs_storage._identityConnectionIDOnIdp {return false}
        if _storage._idpURL != rhs_storage._idpURL {return false}
        if _storage._signInURL != rhs_storage._signInURL {return false}
        if _storage._x509SigningCertBase64 != rhs_storage._x509SigningCertBase64 {return false}
        if _storage._enableSignOut != rhs_storage._enableSignOut {return false}
        if _storage._signOutURL != rhs_storage._signOutURL {return false}
        if _storage._userIDAttribute != rhs_storage._userIDAttribute {return false}
        if _storage._signRequest != rhs_storage._signRequest {return false}
        if _storage._signRequestAlgorithm != rhs_storage._signRequestAlgorithm {return false}
        if _storage._signRequestAlgorithmDigest != rhs_storage._signRequestAlgorithmDigest {return false}
        if _storage._protocolBinding != rhs_storage._protocolBinding {return false}
        if _storage._pingFederateServerURL != rhs_storage._pingFederateServerURL {return false}
        if _storage._idpDomains != rhs_storage._idpDomains {return false}
        if _storage._isCacheDisabled != rhs_storage._isCacheDisabled {return false}
        if _storage._isClientSslCertificateAuthenticationRequired != rhs_storage._isClientSslCertificateAuthenticationRequired {return false}
        if _storage._isKerberosEnabled != rhs_storage._isKerberosEnabled {return false}
        if _storage._issueURL != rhs_storage._issueURL {return false}
        if _storage._clientID != rhs_storage._clientID {return false}
        if _storage._callbackURL != rhs_storage._callbackURL {return false}
        if _storage._adfsMetadataURL != rhs_storage._adfsMetadataURL {return false}
        if _storage._adfsMetadataFederationFileBase64 != rhs_storage._adfsMetadataFederationFileBase64 {return false}
        if _storage._oktaDomain != rhs_storage._oktaDomain {return false}
        if _storage._clientSecret != rhs_storage._clientSecret {return false}
        if _storage._microsoftAzureAdDomain != rhs_storage._microsoftAzureAdDomain {return false}
        if _storage._useCommonEndpoint != rhs_storage._useCommonEndpoint {return false}
        if _storage._identityApi != rhs_storage._identityApi {return false}
        return true
      }
      if !storagesAreEqual {return false}
    }
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionStatusState: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".IdentityConnectionStatusState"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    99: .same(proto: "lifecycle"),
    98: .standard(proto: "sys_audit"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
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
    try { if let v = self._sysAudit {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 98)
    } }()
    try { if let v = self._lifecycle {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 99)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionStatusState, rhs: Cloud_Planton_Apis_V1_Iam_Identity_Connection_State_IdentityConnectionStatusState) -> Bool {
    if lhs._lifecycle != rhs._lifecycle {return false}
    if lhs._sysAudit != rhs._sysAudit {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}