// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/commons/cloud/gcp/resource/folder/rpc/model.proto
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

/// gcp project
/// https://cloud.google.com/resource-manager/reference/rest/v2/folders
public struct Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_Rpc_GcpFolder {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// id of the gcp project.
  /// ex: 498520640386.
  public var id: String = String()

  /// display name for the folder.
  /// (important) a two character suffix starting with a hyphen is added at the end of the
  /// display name to allow recreating folders with same resource identifier on planton cloud.
  public var displayName: String = String()

  /// parent of the folder.
  /// parent can be either a gcp organization or another gcp folder.
  public var parent: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

/// list of gcp folder
public struct Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_Rpc_GcpFolders {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var entries: [Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_Rpc_GcpFolder] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_Rpc_GcpFolder: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_Rpc_GcpFolders: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.commons.cloud.gcp.resource.folder.rpc"

extension Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_Rpc_GcpFolder: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".GcpFolder"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "id"),
    2: .standard(proto: "display_name"),
    3: .same(proto: "parent"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.id) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.displayName) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.parent) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.id.isEmpty {
      try visitor.visitSingularStringField(value: self.id, fieldNumber: 1)
    }
    if !self.displayName.isEmpty {
      try visitor.visitSingularStringField(value: self.displayName, fieldNumber: 2)
    }
    if !self.parent.isEmpty {
      try visitor.visitSingularStringField(value: self.parent, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_Rpc_GcpFolder, rhs: Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_Rpc_GcpFolder) -> Bool {
    if lhs.id != rhs.id {return false}
    if lhs.displayName != rhs.displayName {return false}
    if lhs.parent != rhs.parent {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_Rpc_GcpFolders: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".GcpFolders"
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_Rpc_GcpFolders, rhs: Cloud_Planton_Apis_V1_Commons_Cloud_Gcp_Resource_Folder_Rpc_GcpFolders) -> Bool {
    if lhs.entries != rhs.entries {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
