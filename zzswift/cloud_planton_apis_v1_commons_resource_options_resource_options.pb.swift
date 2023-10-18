// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/commons/resource/options/resource_options.proto
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

/// A message defining the ownership of a resource.
public struct Cloud_Planton_Apis_V1_Commons_Resource_Options_Owner {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// The type of the owner, defined in the ResourceType enum.
  public var type: Cloud_Planton_Apis_V1_Commons_Resource_Enums_ResourceType = .unspecified

  /// The field path of the owner id in a resource.
  public var idFieldPath: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Commons_Resource_Options_Owner: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Extension support defined in resource_options.proto.

// MARK: - Extension Properties

// Swift Extensions on the exteneded Messages to add easy access to the declared
// extension fields. The names are based on the extension field name from the proto
// declaration. To avoid naming collisions, the names are prefixed with the name of
// the scope where the extend directive occurs.

extension SwiftProtobuf.Google_Protobuf_MessageOptions {

  /// An identifier for the type of resource.
  public var Cloud_Planton_Apis_V1_Commons_Resource_Options_resourceType: Cloud_Planton_Apis_V1_Commons_Resource_Enums_ResourceType {
    get {return getExtensionValue(ext: Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_resource_type) ?? .unspecified}
    set {setExtensionValue(ext: Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_resource_type, value: newValue)}
  }
  /// Returns true if extension `Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_resource_type`
  /// has been explicitly set.
  public var hasCloud_Planton_Apis_V1_Commons_Resource_Options_resourceType: Bool {
    return hasExtensionValue(ext: Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_resource_type)
  }
  /// Clears the value of extension `Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_resource_type`.
  /// Subsequent reads from it will return its default value.
  public mutating func clearCloud_Planton_Apis_V1_Commons_Resource_Options_resourceType() {
    clearExtensionValue(ext: Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_resource_type)
  }

  /// Whether the resource is runnable.
  /// All resources that require runtime platform to run should set this to true.
  /// For example, a microservice-instance requires a runtime platform to run and so it is considered runnable.
  public var Cloud_Planton_Apis_V1_Commons_Resource_Options_isRunnable: Bool {
    get {return getExtensionValue(ext: Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_is_runnable) ?? false}
    set {setExtensionValue(ext: Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_is_runnable, value: newValue)}
  }
  /// Returns true if extension `Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_is_runnable`
  /// has been explicitly set.
  public var hasCloud_Planton_Apis_V1_Commons_Resource_Options_isRunnable: Bool {
    return hasExtensionValue(ext: Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_is_runnable)
  }
  /// Clears the value of extension `Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_is_runnable`.
  /// Subsequent reads from it will return its default value.
  public mutating func clearCloud_Planton_Apis_V1_Commons_Resource_Options_isRunnable() {
    clearExtensionValue(ext: Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_is_runnable)
  }

  /// The owner information for the resource.
  public var Cloud_Planton_Apis_V1_Commons_Resource_Options_owner: Cloud_Planton_Apis_V1_Commons_Resource_Options_Owner {
    get {return getExtensionValue(ext: Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_owner) ?? Cloud_Planton_Apis_V1_Commons_Resource_Options_Owner()}
    set {setExtensionValue(ext: Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_owner, value: newValue)}
  }
  /// Returns true if extension `Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_owner`
  /// has been explicitly set.
  public var hasCloud_Planton_Apis_V1_Commons_Resource_Options_owner: Bool {
    return hasExtensionValue(ext: Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_owner)
  }
  /// Clears the value of extension `Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_owner`.
  /// Subsequent reads from it will return its default value.
  public mutating func clearCloud_Planton_Apis_V1_Commons_Resource_Options_owner() {
    clearExtensionValue(ext: Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_owner)
  }

}

// MARK: - File's ExtensionMap: Cloud_Planton_Apis_V1_Commons_Resource_Options_ResourceOptions_Extensions

/// A `SwiftProtobuf.SimpleExtensionMap` that includes all of the extensions defined by
/// this .proto file. It can be used any place an `SwiftProtobuf.ExtensionMap` is needed
/// in parsing, or it can be combined with other `SwiftProtobuf.SimpleExtensionMap`s to create
/// a larger `SwiftProtobuf.SimpleExtensionMap`.
public let Cloud_Planton_Apis_V1_Commons_Resource_Options_ResourceOptions_Extensions: SwiftProtobuf.SimpleExtensionMap = [
  Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_resource_type,
  Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_is_runnable,
  Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_owner
]

// Extension Objects - The only reason these might be needed is when manually
// constructing a `SimpleExtensionMap`, otherwise, use the above _Extension Properties_
// accessors for the extension fields on the messages directly.

/// An identifier for the type of resource.
public let Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_resource_type = SwiftProtobuf.MessageExtension<SwiftProtobuf.OptionalEnumExtensionField<Cloud_Planton_Apis_V1_Commons_Resource_Enums_ResourceType>, SwiftProtobuf.Google_Protobuf_MessageOptions>(
  _protobuf_fieldNumber: 60001,
  fieldName: "cloud.planton.apis.v1.commons.resource.options.resource_type"
)

/// Whether the resource is runnable.
/// All resources that require runtime platform to run should set this to true.
/// For example, a microservice-instance requires a runtime platform to run and so it is considered runnable.
public let Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_is_runnable = SwiftProtobuf.MessageExtension<SwiftProtobuf.OptionalExtensionField<SwiftProtobuf.ProtobufBool>, SwiftProtobuf.Google_Protobuf_MessageOptions>(
  _protobuf_fieldNumber: 60002,
  fieldName: "cloud.planton.apis.v1.commons.resource.options.is_runnable"
)

/// The owner information for the resource.
public let Cloud_Planton_Apis_V1_Commons_Resource_Options_Extensions_owner = SwiftProtobuf.MessageExtension<SwiftProtobuf.OptionalMessageExtensionField<Cloud_Planton_Apis_V1_Commons_Resource_Options_Owner>, SwiftProtobuf.Google_Protobuf_MessageOptions>(
  _protobuf_fieldNumber: 60003,
  fieldName: "cloud.planton.apis.v1.commons.resource.options.owner"
)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.commons.resource.options"

extension Cloud_Planton_Apis_V1_Commons_Resource_Options_Owner: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".Owner"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "type"),
    2: .standard(proto: "id_field_path"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularEnumField(value: &self.type) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.idFieldPath) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.type != .unspecified {
      try visitor.visitSingularEnumField(value: self.type, fieldNumber: 1)
    }
    if !self.idFieldPath.isEmpty {
      try visitor.visitSingularStringField(value: self.idFieldPath, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Resource_Options_Owner, rhs: Cloud_Planton_Apis_V1_Commons_Resource_Options_Owner) -> Bool {
    if lhs.type != rhs.type {return false}
    if lhs.idFieldPath != rhs.idFieldPath {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}