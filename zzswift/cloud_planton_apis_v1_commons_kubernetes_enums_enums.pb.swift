// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/commons/kubernetes/enums/enums.proto
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

public enum Cloud_Planton_Apis_V1_Commons_Kubernetes_Enums_PodControllerType: SwiftProtobuf.Enum {
  public typealias RawValue = Int
  case unspecified // = 0
  case deployment // = 1
  case statefulSet // = 2
  case cronJob // = 3
  case UNRECOGNIZED(Int)

  public init() {
    self = .unspecified
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unspecified
    case 1: self = .deployment
    case 2: self = .statefulSet
    case 3: self = .cronJob
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .unspecified: return 0
    case .deployment: return 1
    case .statefulSet: return 2
    case .cronJob: return 3
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Cloud_Planton_Apis_V1_Commons_Kubernetes_Enums_PodControllerType: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static let allCases: [Cloud_Planton_Apis_V1_Commons_Kubernetes_Enums_PodControllerType] = [
    .unspecified,
    .deployment,
    .statefulSet,
    .cronJob,
  ]
}

#endif  // swift(>=4.2)

/// kubernetes resource kind
public enum Cloud_Planton_Apis_V1_Commons_Kubernetes_Enums_ResourceKind: SwiftProtobuf.Enum {
  public typealias RawValue = Int
  case unspecified // = 0
  case namespace // = 1
  case deployment // = 2
  case service // = 3
  case secret // = 4
  case ingress // = 5
  case issuer // = 6
  case certificate // = 7
  case UNRECOGNIZED(Int)

  public init() {
    self = .unspecified
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unspecified
    case 1: self = .namespace
    case 2: self = .deployment
    case 3: self = .service
    case 4: self = .secret
    case 5: self = .ingress
    case 6: self = .issuer
    case 7: self = .certificate
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .unspecified: return 0
    case .namespace: return 1
    case .deployment: return 2
    case .service: return 3
    case .secret: return 4
    case .ingress: return 5
    case .issuer: return 6
    case .certificate: return 7
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Cloud_Planton_Apis_V1_Commons_Kubernetes_Enums_ResourceKind: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static let allCases: [Cloud_Planton_Apis_V1_Commons_Kubernetes_Enums_ResourceKind] = [
    .unspecified,
    .namespace,
    .deployment,
    .service,
    .secret,
    .ingress,
    .issuer,
    .certificate,
  ]
}

#endif  // swift(>=4.2)

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Commons_Kubernetes_Enums_PodControllerType: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Commons_Kubernetes_Enums_ResourceKind: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

extension Cloud_Planton_Apis_V1_Commons_Kubernetes_Enums_PodControllerType: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "POD_CONTROLLER_TYPE_UNSPECIFIED"),
    1: .same(proto: "deployment"),
    2: .same(proto: "stateful_set"),
    3: .same(proto: "cron_job"),
  ]
}

extension Cloud_Planton_Apis_V1_Commons_Kubernetes_Enums_ResourceKind: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "RESOURCE_KIND_UNSPECIFIED"),
    1: .same(proto: "Namespace"),
    2: .same(proto: "Deployment"),
    3: .same(proto: "Service"),
    4: .same(proto: "Secret"),
    5: .same(proto: "Ingress"),
    6: .same(proto: "Issuer"),
    7: .same(proto: "Certificate"),
  ]
}