// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/code2cloud/deploy/endpoint/custom/state/enums/enums.proto
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

public enum Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_State_Enums_CustomEndpointEventType: SwiftProtobuf.Enum {
  public typealias RawValue = Int
  case unspecified // = 0
  case stateCreated // = 1
  case stateUpdated // = 2
  case stateDeleted // = 3
  case stateRestored // = 4
  case stackJobProgressUpdated // = 5
  case stackJobPreviewRequested // = 6
  case stackJobApplyRequested // = 7
  case stackJobApplyCompleted // = 8
  case stackJobDestroyRequested // = 9
  case stackJobDestroyCompleted // = 10
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
    case 4: self = .stateRestored
    case 5: self = .stackJobProgressUpdated
    case 6: self = .stackJobPreviewRequested
    case 7: self = .stackJobApplyRequested
    case 8: self = .stackJobApplyCompleted
    case 9: self = .stackJobDestroyRequested
    case 10: self = .stackJobDestroyCompleted
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .unspecified: return 0
    case .stateCreated: return 1
    case .stateUpdated: return 2
    case .stateDeleted: return 3
    case .stateRestored: return 4
    case .stackJobProgressUpdated: return 5
    case .stackJobPreviewRequested: return 6
    case .stackJobApplyRequested: return 7
    case .stackJobApplyCompleted: return 8
    case .stackJobDestroyRequested: return 9
    case .stackJobDestroyCompleted: return 10
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_State_Enums_CustomEndpointEventType: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static let allCases: [Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_State_Enums_CustomEndpointEventType] = [
    .unspecified,
    .stateCreated,
    .stateUpdated,
    .stateDeleted,
    .stateRestored,
    .stackJobProgressUpdated,
    .stackJobPreviewRequested,
    .stackJobApplyRequested,
    .stackJobApplyCompleted,
    .stackJobDestroyRequested,
    .stackJobDestroyCompleted,
  ]
}

#endif  // swift(>=4.2)

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_State_Enums_CustomEndpointEventType: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

extension Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_State_Enums_CustomEndpointEventType: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "CUSTOM_ENDPOINT_EVENT_TYPE_UNSPECIFIED"),
    1: .same(proto: "CUSTOM_ENDPOINT_EVENT_TYPE_STATE_CREATED"),
    2: .same(proto: "CUSTOM_ENDPOINT_EVENT_TYPE_STATE_UPDATED"),
    3: .same(proto: "CUSTOM_ENDPOINT_EVENT_TYPE_STATE_DELETED"),
    4: .same(proto: "CUSTOM_ENDPOINT_EVENT_TYPE_STATE_RESTORED"),
    5: .same(proto: "CUSTOM_ENDPOINT_EVENT_TYPE_STACK_JOB_PROGRESS_UPDATED"),
    6: .same(proto: "CUSTOM_ENDPOINT_EVENT_TYPE_STACK_JOB_PREVIEW_REQUESTED"),
    7: .same(proto: "CUSTOM_ENDPOINT_EVENT_TYPE_STACK_JOB_APPLY_REQUESTED"),
    8: .same(proto: "CUSTOM_ENDPOINT_EVENT_TYPE_STACK_JOB_APPLY_COMPLETED"),
    9: .same(proto: "CUSTOM_ENDPOINT_EVENT_TYPE_STACK_JOB_DESTROY_REQUESTED"),
    10: .same(proto: "CUSTOM_ENDPOINT_EVENT_TYPE_STACK_JOB_DESTROY_COMPLETED"),
  ]
}
