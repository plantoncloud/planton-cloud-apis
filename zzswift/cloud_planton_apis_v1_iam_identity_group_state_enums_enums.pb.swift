// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/iam/identity/group/state/enums/enums.proto
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

public enum Cloud_Planton_Apis_V1_Iam_Identity_Group_State_Enums_IdentityGroupEventType: SwiftProtobuf.Enum {
  public typealias RawValue = Int
  case unspecified // = 0
  case stateCreated // = 1
  case stateUpdated // = 2
  case stateDeleted // = 3
  case UNRECOGNIZED(Int)

  public init() {
    self = .unspecified
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unspecified
    case 1: self = .stateCreated
    case 2: self = .stateUpdated
    case 3: self = .stateDeleted
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .unspecified: return 0
    case .stateCreated: return 1
    case .stateUpdated: return 2
    case .stateDeleted: return 3
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Cloud_Planton_Apis_V1_Iam_Identity_Group_State_Enums_IdentityGroupEventType: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static let allCases: [Cloud_Planton_Apis_V1_Iam_Identity_Group_State_Enums_IdentityGroupEventType] = [
    .unspecified,
    .stateCreated,
    .stateUpdated,
    .stateDeleted,
  ]
}

#endif  // swift(>=4.2)

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Iam_Identity_Group_State_Enums_IdentityGroupEventType: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

extension Cloud_Planton_Apis_V1_Iam_Identity_Group_State_Enums_IdentityGroupEventType: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "IDENTITY_GROUP_EVENT_TYPE_UNSPECIFIED"),
    1: .same(proto: "IDENTITY_GROUP_EVENT_TYPE_STATE_CREATED"),
    2: .same(proto: "IDENTITY_GROUP_EVENT_TYPE_STATE_UPDATED"),
    3: .same(proto: "IDENTITY_GROUP_EVENT_TYPE_STATE_DELETED"),
  ]
}
