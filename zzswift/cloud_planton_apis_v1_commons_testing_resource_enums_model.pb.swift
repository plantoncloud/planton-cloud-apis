// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/commons/testing/resource/enums/model.proto
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

public enum Cloud_Planton_Apis_V1_Commons_Testing_Resource_Enums_TestEventType: SwiftProtobuf.Enum {
  public typealias RawValue = Int
  case unspecified // = 0
  case stateCreated // = 1
  case stateUpdated // = 2
  case UNRECOGNIZED(Int)

  public init() {
    self = .unspecified
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unspecified
    case 1: self = .stateCreated
    case 2: self = .stateUpdated
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .unspecified: return 0
    case .stateCreated: return 1
    case .stateUpdated: return 2
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Cloud_Planton_Apis_V1_Commons_Testing_Resource_Enums_TestEventType: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static let allCases: [Cloud_Planton_Apis_V1_Commons_Testing_Resource_Enums_TestEventType] = [
    .unspecified,
    .stateCreated,
    .stateUpdated,
  ]
}

#endif  // swift(>=4.2)

/// A message with computed fields.
public struct Cloud_Planton_Apis_V1_Commons_Testing_Resource_Enums_EnumFieldsTest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var eventType: Cloud_Planton_Apis_V1_Commons_Testing_Resource_Enums_TestEventType = .unspecified

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Commons_Testing_Resource_Enums_TestEventType: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Commons_Testing_Resource_Enums_EnumFieldsTest: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.commons.testing.resource.enums"

extension Cloud_Planton_Apis_V1_Commons_Testing_Resource_Enums_TestEventType: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "TEST_EVENT_TYPE_UNSPECIFIED"),
    1: .same(proto: "TEST_EVENT_TYPE_STATE_CREATED"),
    2: .same(proto: "TEST_EVENT_TYPE_STATE_UPDATED"),
  ]
}

extension Cloud_Planton_Apis_V1_Commons_Testing_Resource_Enums_EnumFieldsTest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".EnumFieldsTest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "event_type"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularEnumField(value: &self.eventType) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.eventType != .unspecified {
      try visitor.visitSingularEnumField(value: self.eventType, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Testing_Resource_Enums_EnumFieldsTest, rhs: Cloud_Planton_Apis_V1_Commons_Testing_Resource_Enums_EnumFieldsTest) -> Bool {
    if lhs.eventType != rhs.eventType {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
