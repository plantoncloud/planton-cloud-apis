// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/iam/authz/extensions/extensions.proto
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

/// authorization config message to check before executing rpc
public struct Cloud_Planton_Apis_V1_Iam_Authz_Extensions_AuthorizationConfig {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// permission required to run the service
  public var permission: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_Enums_IamPermission = .unspecified

  /// type of the resource that the permission should be granted
  public var resourceType: Cloud_Planton_Apis_V1_Commons_Resource_Enums_ResourceType = .unspecified

  /// path of the field inside the input object to be used for performing authorization.
  public var fieldPath: String = String()

  /// error message to be returned when the permission is not granted to user
  public var errorMsg: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Iam_Authz_Extensions_AuthorizationConfig: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Extension support defined in extensions.proto.

// MARK: - Extension Properties

// Swift Extensions on the exteneded Messages to add easy access to the declared
// extension fields. The names are based on the extension field name from the proto
// declaration. To avoid naming collisions, the names are prefixed with the name of
// the scope where the extend directive occurs.

extension SwiftProtobuf.Google_Protobuf_MethodOptions {

  public var Cloud_Planton_Apis_V1_Iam_Authz_Extensions_authorizationConfig: Cloud_Planton_Apis_V1_Iam_Authz_Extensions_AuthorizationConfig {
    get {return getExtensionValue(ext: Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_authorization_config) ?? Cloud_Planton_Apis_V1_Iam_Authz_Extensions_AuthorizationConfig()}
    set {setExtensionValue(ext: Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_authorization_config, value: newValue)}
  }
  /// Returns true if extension `Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_authorization_config`
  /// has been explicitly set.
  public var hasCloud_Planton_Apis_V1_Iam_Authz_Extensions_authorizationConfig: Bool {
    return hasExtensionValue(ext: Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_authorization_config)
  }
  /// Clears the value of extension `Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_authorization_config`.
  /// Subsequent reads from it will return its default value.
  public mutating func clearCloud_Planton_Apis_V1_Iam_Authz_Extensions_authorizationConfig() {
    clearExtensionValue(ext: Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_authorization_config)
  }

  public var Cloud_Planton_Apis_V1_Iam_Authz_Extensions_isPublic: Bool {
    get {return getExtensionValue(ext: Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_is_public) ?? false}
    set {setExtensionValue(ext: Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_is_public, value: newValue)}
  }
  /// Returns true if extension `Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_is_public`
  /// has been explicitly set.
  public var hasCloud_Planton_Apis_V1_Iam_Authz_Extensions_isPublic: Bool {
    return hasExtensionValue(ext: Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_is_public)
  }
  /// Clears the value of extension `Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_is_public`.
  /// Subsequent reads from it will return its default value.
  public mutating func clearCloud_Planton_Apis_V1_Iam_Authz_Extensions_isPublic() {
    clearExtensionValue(ext: Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_is_public)
  }

}

// MARK: - File's ExtensionMap: Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_Extensions

/// A `SwiftProtobuf.SimpleExtensionMap` that includes all of the extensions defined by
/// this .proto file. It can be used any place an `SwiftProtobuf.ExtensionMap` is needed
/// in parsing, or it can be combined with other `SwiftProtobuf.SimpleExtensionMap`s to create
/// a larger `SwiftProtobuf.SimpleExtensionMap`.
public let Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_Extensions: SwiftProtobuf.SimpleExtensionMap = [
  Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_authorization_config,
  Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_is_public
]

// Extension Objects - The only reason these might be needed is when manually
// constructing a `SimpleExtensionMap`, otherwise, use the above _Extension Properties_
// accessors for the extension fields on the messages directly.

public let Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_authorization_config = SwiftProtobuf.MessageExtension<SwiftProtobuf.OptionalMessageExtensionField<Cloud_Planton_Apis_V1_Iam_Authz_Extensions_AuthorizationConfig>, SwiftProtobuf.Google_Protobuf_MethodOptions>(
  _protobuf_fieldNumber: 50056,
  fieldName: "cloud.planton.apis.v1.iam.authz.extensions.authorization_config"
)

public let Cloud_Planton_Apis_V1_Iam_Authz_Extensions_Extensions_is_public = SwiftProtobuf.MessageExtension<SwiftProtobuf.OptionalExtensionField<SwiftProtobuf.ProtobufBool>, SwiftProtobuf.Google_Protobuf_MethodOptions>(
  _protobuf_fieldNumber: 50057,
  fieldName: "cloud.planton.apis.v1.iam.authz.extensions.is_public"
)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.iam.authz.extensions"

extension Cloud_Planton_Apis_V1_Iam_Authz_Extensions_AuthorizationConfig: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".AuthorizationConfig"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "permission"),
    2: .standard(proto: "resource_type"),
    3: .standard(proto: "field_path"),
    4: .standard(proto: "error_msg"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularEnumField(value: &self.permission) }()
      case 2: try { try decoder.decodeSingularEnumField(value: &self.resourceType) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.fieldPath) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.errorMsg) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.permission != .unspecified {
      try visitor.visitSingularEnumField(value: self.permission, fieldNumber: 1)
    }
    if self.resourceType != .unspecified {
      try visitor.visitSingularEnumField(value: self.resourceType, fieldNumber: 2)
    }
    if !self.fieldPath.isEmpty {
      try visitor.visitSingularStringField(value: self.fieldPath, fieldNumber: 3)
    }
    if !self.errorMsg.isEmpty {
      try visitor.visitSingularStringField(value: self.errorMsg, fieldNumber: 4)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Iam_Authz_Extensions_AuthorizationConfig, rhs: Cloud_Planton_Apis_V1_Iam_Authz_Extensions_AuthorizationConfig) -> Bool {
    if lhs.permission != rhs.permission {return false}
    if lhs.resourceType != rhs.resourceType {return false}
    if lhs.fieldPath != rhs.fieldPath {return false}
    if lhs.errorMsg != rhs.errorMsg {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}