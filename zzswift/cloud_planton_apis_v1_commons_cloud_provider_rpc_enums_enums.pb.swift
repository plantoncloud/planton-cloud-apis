// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/commons/cloud/provider/rpc/enums/enums.proto
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

///cloud provider type
///the recommended best practice to prefix the entry with enum name and using uppercase for entry names has been intentionally ignored as the values of entry names are used for naming resources.
public enum Cloud_Planton_Apis_V1_Commons_Cloud_Provider_Rpc_Enums_CloudProvider: SwiftProtobuf.Enum {
  public typealias RawValue = Int

  /// default value when cloud provider name is not set
  case unspecified // = 0

  /// amazon web services cloud provider
  case aws // = 1

  /// google cloud platform cloud provider
  case gcp // = 2

  /// microsoft azure cloud provider
  case azure // = 3

  /// alibaba cloud provider
  case alibaba // = 4

  /// oracle cloud provider
  case oracle // = 5
  case UNRECOGNIZED(Int)

  public init() {
    self = .unspecified
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unspecified
    case 1: self = .aws
    case 2: self = .gcp
    case 3: self = .azure
    case 4: self = .alibaba
    case 5: self = .oracle
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .unspecified: return 0
    case .aws: return 1
    case .gcp: return 2
    case .azure: return 3
    case .alibaba: return 4
    case .oracle: return 5
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Cloud_Planton_Apis_V1_Commons_Cloud_Provider_Rpc_Enums_CloudProvider: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static let allCases: [Cloud_Planton_Apis_V1_Commons_Cloud_Provider_Rpc_Enums_CloudProvider] = [
    .unspecified,
    .aws,
    .gcp,
    .azure,
    .alibaba,
    .oracle,
  ]
}

#endif  // swift(>=4.2)

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Commons_Cloud_Provider_Rpc_Enums_CloudProvider: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

extension Cloud_Planton_Apis_V1_Commons_Cloud_Provider_Rpc_Enums_CloudProvider: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "CLOUD_PROVIDER_UNSPECIFIED"),
    1: .same(proto: "aws"),
    2: .same(proto: "gcp"),
    3: .same(proto: "azure"),
    4: .same(proto: "alibaba"),
    5: .same(proto: "oracle"),
  ]
}
