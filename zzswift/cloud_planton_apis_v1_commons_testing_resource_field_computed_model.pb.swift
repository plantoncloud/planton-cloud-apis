// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/commons/testing/resource/field/computed/model.proto
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

/// A message with computed fields.
public struct Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_ComputedFieldsTest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var requiredStringField: String = String()

  public var computedStringField: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

/// A message with a nested message containing computed fields.
public struct Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_NestedComputedFieldsTest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var nestedComputedField: Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_ComputedFieldsTest {
    get {return _nestedComputedField ?? Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_ComputedFieldsTest()}
    set {_nestedComputedField = newValue}
  }
  /// Returns true if `nestedComputedField` has been explicitly set.
  public var hasNestedComputedField: Bool {return self._nestedComputedField != nil}
  /// Clears the value of `nestedComputedField`. Subsequent reads from it will return its default value.
  public mutating func clearNestedComputedField() {self._nestedComputedField = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _nestedComputedField: Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_ComputedFieldsTest? = nil
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_ComputedFieldsTest: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_NestedComputedFieldsTest: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.commons.testing.resource.field.computed"

extension Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_ComputedFieldsTest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ComputedFieldsTest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "required_string_field"),
    2: .standard(proto: "computed_string_field"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.requiredStringField) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.computedStringField) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.requiredStringField.isEmpty {
      try visitor.visitSingularStringField(value: self.requiredStringField, fieldNumber: 1)
    }
    if !self.computedStringField.isEmpty {
      try visitor.visitSingularStringField(value: self.computedStringField, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_ComputedFieldsTest, rhs: Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_ComputedFieldsTest) -> Bool {
    if lhs.requiredStringField != rhs.requiredStringField {return false}
    if lhs.computedStringField != rhs.computedStringField {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_NestedComputedFieldsTest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".NestedComputedFieldsTest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "nested_computed_field"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._nestedComputedField) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._nestedComputedField {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_NestedComputedFieldsTest, rhs: Cloud_Planton_Apis_V1_Commons_Testing_Resource_Field_Computed_NestedComputedFieldsTest) -> Bool {
    if lhs._nestedComputedField != rhs._nestedComputedField {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
