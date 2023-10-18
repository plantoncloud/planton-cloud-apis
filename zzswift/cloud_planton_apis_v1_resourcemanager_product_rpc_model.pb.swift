// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/resourcemanager/product/rpc/model.proto
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

///product
public struct Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_Product {
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
  public var spec: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductSpec {
    get {return _storage._spec ?? Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductSpec()}
    set {_uniqueStorage()._spec = newValue}
  }
  /// Returns true if `spec` has been explicitly set.
  public var hasSpec: Bool {return _storage._spec != nil}
  /// Clears the value of `spec`. Subsequent reads from it will return its default value.
  public mutating func clearSpec() {_uniqueStorage()._spec = nil}

  ///status
  public var status: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductStatus {
    get {return _storage._status ?? Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductStatus()}
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

///product spec
public struct Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductSpec {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///company to which the product belongs to.
  public var companyID: String = String()

  ///product key is a three character key for the product.
  ///this key is unique with in the company.
  ///product key length should be minimum 2 chars and maximum of 4 chars.
  ///product-key is suffixed to the company-id to create product-id
  public var productKey: String = String()

  ///product description.
  public var description_p: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///product status
public struct Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductStatus {
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

///list of products
public struct Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_Products {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var entries: [Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_Product] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///wrapper for product id
public struct Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductId {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var value: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///response for paginated query to get list of products
public struct Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductList {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var totalPages: Int32 = 0

  public var entries: [Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_Product] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_Product: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductSpec: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductStatus: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_Products: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductId: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductList: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.resourcemanager.product.rpc"

extension Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_Product: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".Product"
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
    var _spec: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductSpec? = nil
    var _status: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductStatus? = nil

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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_Product, rhs: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_Product) -> Bool {
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

extension Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductSpec: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ProductSpec"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "company_id"),
    2: .standard(proto: "product_key"),
    3: .same(proto: "description"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.companyID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.productKey) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.description_p) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.companyID.isEmpty {
      try visitor.visitSingularStringField(value: self.companyID, fieldNumber: 1)
    }
    if !self.productKey.isEmpty {
      try visitor.visitSingularStringField(value: self.productKey, fieldNumber: 2)
    }
    if !self.description_p.isEmpty {
      try visitor.visitSingularStringField(value: self.description_p, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductSpec, rhs: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductSpec) -> Bool {
    if lhs.companyID != rhs.companyID {return false}
    if lhs.productKey != rhs.productKey {return false}
    if lhs.description_p != rhs.description_p {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductStatus: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ProductStatus"
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductStatus, rhs: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductStatus) -> Bool {
    if lhs._lifecycle != rhs._lifecycle {return false}
    if lhs._sysAudit != rhs._sysAudit {return false}
    if lhs.stackJobID != rhs.stackJobID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_Products: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".Products"
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_Products, rhs: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_Products) -> Bool {
    if lhs.entries != rhs.entries {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductId: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ProductId"
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductId, rhs: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductId) -> Bool {
    if lhs.value != rhs.value {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductList: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ProductList"
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

  public static func ==(lhs: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductList, rhs: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductList) -> Bool {
    if lhs.totalPages != rhs.totalPages {return false}
    if lhs.entries != rhs.entries {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
