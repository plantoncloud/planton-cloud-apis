// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/commons/testing/protobuf/model.proto
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

public struct Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_TestMessage {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var rootLevelString: String = String()

  public var levelOne: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelOne {
    get {return _levelOne ?? Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelOne()}
    set {_levelOne = newValue}
  }
  /// Returns true if `levelOne` has been explicitly set.
  public var hasLevelOne: Bool {return self._levelOne != nil}
  /// Clears the value of `levelOne`. Subsequent reads from it will return its default value.
  public mutating func clearLevelOne() {self._levelOne = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _levelOne: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelOne? = nil
}

public struct Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelOne {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var levelOneString: String = String()

  public var levelTwo: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelTwo {
    get {return _levelTwo ?? Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelTwo()}
    set {_levelTwo = newValue}
  }
  /// Returns true if `levelTwo` has been explicitly set.
  public var hasLevelTwo: Bool {return self._levelTwo != nil}
  /// Clears the value of `levelTwo`. Subsequent reads from it will return its default value.
  public mutating func clearLevelTwo() {self._levelTwo = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _levelTwo: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelTwo? = nil
}

public struct Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelTwo {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var levelTwoString: String = String()

  public var levelThree: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelThree {
    get {return _levelThree ?? Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelThree()}
    set {_levelThree = newValue}
  }
  /// Returns true if `levelThree` has been explicitly set.
  public var hasLevelThree: Bool {return self._levelThree != nil}
  /// Clears the value of `levelThree`. Subsequent reads from it will return its default value.
  public mutating func clearLevelThree() {self._levelThree = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _levelThree: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelThree? = nil
}

public struct Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelThree {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var levelThreeString: String = String()

  public var notAstring: Int32 = 0

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_TestMessage: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelOne: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelTwo: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelThree: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.commons.testing.protobuf"

extension Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_TestMessage: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".TestMessage"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "rootLevelString"),
    2: .same(proto: "levelOne"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.rootLevelString) }()
      case 2: try { try decoder.decodeSingularMessageField(value: &self._levelOne) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.rootLevelString.isEmpty {
      try visitor.visitSingularStringField(value: self.rootLevelString, fieldNumber: 1)
    }
    try { if let v = self._levelOne {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_TestMessage, rhs: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_TestMessage) -> Bool {
    if lhs.rootLevelString != rhs.rootLevelString {return false}
    if lhs._levelOne != rhs._levelOne {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelOne: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".LevelOne"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "levelOneString"),
    2: .same(proto: "levelTwo"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.levelOneString) }()
      case 2: try { try decoder.decodeSingularMessageField(value: &self._levelTwo) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.levelOneString.isEmpty {
      try visitor.visitSingularStringField(value: self.levelOneString, fieldNumber: 1)
    }
    try { if let v = self._levelTwo {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelOne, rhs: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelOne) -> Bool {
    if lhs.levelOneString != rhs.levelOneString {return false}
    if lhs._levelTwo != rhs._levelTwo {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelTwo: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".LevelTwo"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "levelTwoString"),
    2: .same(proto: "levelThree"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.levelTwoString) }()
      case 2: try { try decoder.decodeSingularMessageField(value: &self._levelThree) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.levelTwoString.isEmpty {
      try visitor.visitSingularStringField(value: self.levelTwoString, fieldNumber: 1)
    }
    try { if let v = self._levelThree {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelTwo, rhs: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelTwo) -> Bool {
    if lhs.levelTwoString != rhs.levelTwoString {return false}
    if lhs._levelThree != rhs._levelThree {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelThree: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".LevelThree"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "levelThreeString"),
    2: .same(proto: "notAString"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.levelThreeString) }()
      case 2: try { try decoder.decodeSingularInt32Field(value: &self.notAstring) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.levelThreeString.isEmpty {
      try visitor.visitSingularStringField(value: self.levelThreeString, fieldNumber: 1)
    }
    if self.notAstring != 0 {
      try visitor.visitSingularInt32Field(value: self.notAstring, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelThree, rhs: Cloud_Planton_Apis_V1_Commons_Testing_Protobuf_LevelThree) -> Bool {
    if lhs.levelThreeString != rhs.levelThreeString {return false}
    if lhs.notAstring != rhs.notAstring {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
