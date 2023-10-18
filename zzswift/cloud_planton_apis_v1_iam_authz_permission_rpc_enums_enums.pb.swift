// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/iam/authz/permission/rpc/enums/enums.proto
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

public enum Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_Enums_IamPermission: SwiftProtobuf.Enum {
  public typealias RawValue = Int
  case unspecified // = 0

  /// artifact_store
  case artifactStoreCreate // = 1
  case artifactStoreCreateCloudAccountGet // = 2
  case artifactStoreDelete // = 3
  case artifactStoreGet // = 4
  case artifactStoreRestore // = 5
  case artifactStoreUpdate // = 6

  /// billing_account
  case billingAccountDelete // = 7
  case billingAccountGet // = 8
  case billingAccountRestore // = 9
  case billingAccountUpdate // = 10

  /// billing_subscription
  case billingSubscriptionDelete // = 11
  case billingSubscriptionGet // = 12
  case billingSubscriptionRestore // = 13
  case billingSubscriptionUpdate // = 14

  /// cloud_account
  case cloudAccountCreate // = 15
  case cloudAccountDelete // = 16
  case cloudAccountGet // = 17
  case cloudAccountRestore // = 18
  case cloudAccountUpdate // = 19

  /// code_project
  case codeProjectCreate // = 20
  case codeProjectGet // = 21
  case codeProjectDelete // = 22
  case codeProjectRestore // = 23
  case codeProjectUpdate // = 24

  /// code_server
  case codeServerCreate // = 25
  case codeServerDelete // = 26
  case codeServerGet // = 27
  case codeServerRestore // = 28
  case codeServerSynchronize // = 29
  case codeServerUpdate // = 30

  /// company
  case companyDelete // = 31
  case companyGet // = 32
  case companyRestore // = 33
  case companyUpdate // = 34

  /// custom_endpoint
  case customEndpointCreate // = 35
  case customEndpointDelete // = 36
  case customEndpointGet // = 37
  case customEndpointRestore // = 38
  case customEndpointUpdate // = 39

  /// dns_zone
  case dnsZoneCreateCloudAccountGet // = 40
  case dnsZoneCreate // = 41
  case dnsZoneDelete // = 42
  case dnsZoneGet // = 43
  case dnsZoneRestore // = 44
  case dnsZoneUpdate // = 45

  /// environment
  case environmentCreateKubeClusterGet // = 46
  case environmentCreate // = 47
  case environmentDelete // = 48
  case environmentGet // = 49
  case environmentUpdate // = 50
  case environmentRestore // = 51

  /// identity_group
  case identityGroupCreate // = 52
  case identityGroupDelete // = 53
  case identityGroupGet // = 54
  case identityGroupUpdate // = 55

  /// kube_cluster
  case kubeClusterCreate // = 56
  case kubeClusterCreateCloudAccountGet // = 57
  case kubeClusterDelete // = 58
  case kubeClusterGet // = 59
  case kubeClusterRestore // = 60
  case kubeClusterUpdate // = 61

  /// iam_policy
  case iamPolicyGet // = 62
  case iamPolicyUpdate // = 63

  /// identity_account
  case identityAccountCreate // = 64
  case identityAccountDelete // = 65
  case identityAccountGet // = 66
  case identityAccountRestore // = 67
  case identityAccountUpdate // = 68

  /// identity_connection
  case identityConnectionCreate // = 69
  case identityConnectionDelete // = 70
  case identityConnectionGet // = 71
  case identityConnectionRestore // = 72
  case identityConnectionUpdate // = 73

  /// kafka_cluster
  case kafkaClusterCreate // = 74
  case kafkaClusterDelete // = 75
  case kafkaClusterGet // = 76
  case kafkaClusterRestore // = 77
  case kafkaClusterUpdate // = 78

  /// list_associates
  case listAssociates // = 79

  /// login_to_back_office
  case loginToBackOffice // = 80

  /// microservice_instance
  case microserviceInstanceCreate // = 81
  case microserviceInstanceDelete // = 82
  case microserviceInstanceGet // = 83
  case microserviceInstanceRestore // = 84
  case microserviceInstanceUpdate // = 85

  /// postgres_cluster
  case postgresClusterCreate // = 86
  case postgresClusterDelete // = 87
  case postgresClusterGet // = 88
  case postgresClusterRestore // = 89
  case postgresClusterUpdate // = 90

  /// product
  case productCreate // = 91
  case productGet // = 92
  case productUpdate // = 93
  case productDelete // = 94
  case productRestore // = 95

  /// redis_cluster
  case redisClusterCreate // = 96
  case redisClusterDelete // = 97
  case redisClusterGet // = 98
  case redisClusterRestore // = 99
  case redisClusterUpdate // = 100

  /// solr_cloud
  case solrCloudCreate // = 101
  case solrCloudDelete // = 102
  case solrCloudGet // = 103
  case solrCloudRestore // = 104
  case solrCloudUpdate // = 105

  /// standard_endpoint
  case standardEndpointCreate // = 106
  case standardEndpointDelete // = 107
  case standardEndpointGet // = 108
  case standardEndpointRestore // = 109
  case standardEndpointUpdate // = 110

  /// storage_bucket
  case storageBucketCreate // = 111
  case storageBucketDelete // = 112
  case storageBucketGet // = 113
  case storageBucketRestore // = 114
  case storageBucketUpdate // = 115
  case UNRECOGNIZED(Int)

  public init() {
    self = .unspecified
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unspecified
    case 1: self = .artifactStoreCreate
    case 2: self = .artifactStoreCreateCloudAccountGet
    case 3: self = .artifactStoreDelete
    case 4: self = .artifactStoreGet
    case 5: self = .artifactStoreRestore
    case 6: self = .artifactStoreUpdate
    case 7: self = .billingAccountDelete
    case 8: self = .billingAccountGet
    case 9: self = .billingAccountRestore
    case 10: self = .billingAccountUpdate
    case 11: self = .billingSubscriptionDelete
    case 12: self = .billingSubscriptionGet
    case 13: self = .billingSubscriptionRestore
    case 14: self = .billingSubscriptionUpdate
    case 15: self = .cloudAccountCreate
    case 16: self = .cloudAccountDelete
    case 17: self = .cloudAccountGet
    case 18: self = .cloudAccountRestore
    case 19: self = .cloudAccountUpdate
    case 20: self = .codeProjectCreate
    case 21: self = .codeProjectGet
    case 22: self = .codeProjectDelete
    case 23: self = .codeProjectRestore
    case 24: self = .codeProjectUpdate
    case 25: self = .codeServerCreate
    case 26: self = .codeServerDelete
    case 27: self = .codeServerGet
    case 28: self = .codeServerRestore
    case 29: self = .codeServerSynchronize
    case 30: self = .codeServerUpdate
    case 31: self = .companyDelete
    case 32: self = .companyGet
    case 33: self = .companyRestore
    case 34: self = .companyUpdate
    case 35: self = .customEndpointCreate
    case 36: self = .customEndpointDelete
    case 37: self = .customEndpointGet
    case 38: self = .customEndpointRestore
    case 39: self = .customEndpointUpdate
    case 40: self = .dnsZoneCreateCloudAccountGet
    case 41: self = .dnsZoneCreate
    case 42: self = .dnsZoneDelete
    case 43: self = .dnsZoneGet
    case 44: self = .dnsZoneRestore
    case 45: self = .dnsZoneUpdate
    case 46: self = .environmentCreateKubeClusterGet
    case 47: self = .environmentCreate
    case 48: self = .environmentDelete
    case 49: self = .environmentGet
    case 50: self = .environmentUpdate
    case 51: self = .environmentRestore
    case 52: self = .identityGroupCreate
    case 53: self = .identityGroupDelete
    case 54: self = .identityGroupGet
    case 55: self = .identityGroupUpdate
    case 56: self = .kubeClusterCreate
    case 57: self = .kubeClusterCreateCloudAccountGet
    case 58: self = .kubeClusterDelete
    case 59: self = .kubeClusterGet
    case 60: self = .kubeClusterRestore
    case 61: self = .kubeClusterUpdate
    case 62: self = .iamPolicyGet
    case 63: self = .iamPolicyUpdate
    case 64: self = .identityAccountCreate
    case 65: self = .identityAccountDelete
    case 66: self = .identityAccountGet
    case 67: self = .identityAccountRestore
    case 68: self = .identityAccountUpdate
    case 69: self = .identityConnectionCreate
    case 70: self = .identityConnectionDelete
    case 71: self = .identityConnectionGet
    case 72: self = .identityConnectionRestore
    case 73: self = .identityConnectionUpdate
    case 74: self = .kafkaClusterCreate
    case 75: self = .kafkaClusterDelete
    case 76: self = .kafkaClusterGet
    case 77: self = .kafkaClusterRestore
    case 78: self = .kafkaClusterUpdate
    case 79: self = .listAssociates
    case 80: self = .loginToBackOffice
    case 81: self = .microserviceInstanceCreate
    case 82: self = .microserviceInstanceDelete
    case 83: self = .microserviceInstanceGet
    case 84: self = .microserviceInstanceRestore
    case 85: self = .microserviceInstanceUpdate
    case 86: self = .postgresClusterCreate
    case 87: self = .postgresClusterDelete
    case 88: self = .postgresClusterGet
    case 89: self = .postgresClusterRestore
    case 90: self = .postgresClusterUpdate
    case 91: self = .productCreate
    case 92: self = .productGet
    case 93: self = .productUpdate
    case 94: self = .productDelete
    case 95: self = .productRestore
    case 96: self = .redisClusterCreate
    case 97: self = .redisClusterDelete
    case 98: self = .redisClusterGet
    case 99: self = .redisClusterRestore
    case 100: self = .redisClusterUpdate
    case 101: self = .solrCloudCreate
    case 102: self = .solrCloudDelete
    case 103: self = .solrCloudGet
    case 104: self = .solrCloudRestore
    case 105: self = .solrCloudUpdate
    case 106: self = .standardEndpointCreate
    case 107: self = .standardEndpointDelete
    case 108: self = .standardEndpointGet
    case 109: self = .standardEndpointRestore
    case 110: self = .standardEndpointUpdate
    case 111: self = .storageBucketCreate
    case 112: self = .storageBucketDelete
    case 113: self = .storageBucketGet
    case 114: self = .storageBucketRestore
    case 115: self = .storageBucketUpdate
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .unspecified: return 0
    case .artifactStoreCreate: return 1
    case .artifactStoreCreateCloudAccountGet: return 2
    case .artifactStoreDelete: return 3
    case .artifactStoreGet: return 4
    case .artifactStoreRestore: return 5
    case .artifactStoreUpdate: return 6
    case .billingAccountDelete: return 7
    case .billingAccountGet: return 8
    case .billingAccountRestore: return 9
    case .billingAccountUpdate: return 10
    case .billingSubscriptionDelete: return 11
    case .billingSubscriptionGet: return 12
    case .billingSubscriptionRestore: return 13
    case .billingSubscriptionUpdate: return 14
    case .cloudAccountCreate: return 15
    case .cloudAccountDelete: return 16
    case .cloudAccountGet: return 17
    case .cloudAccountRestore: return 18
    case .cloudAccountUpdate: return 19
    case .codeProjectCreate: return 20
    case .codeProjectGet: return 21
    case .codeProjectDelete: return 22
    case .codeProjectRestore: return 23
    case .codeProjectUpdate: return 24
    case .codeServerCreate: return 25
    case .codeServerDelete: return 26
    case .codeServerGet: return 27
    case .codeServerRestore: return 28
    case .codeServerSynchronize: return 29
    case .codeServerUpdate: return 30
    case .companyDelete: return 31
    case .companyGet: return 32
    case .companyRestore: return 33
    case .companyUpdate: return 34
    case .customEndpointCreate: return 35
    case .customEndpointDelete: return 36
    case .customEndpointGet: return 37
    case .customEndpointRestore: return 38
    case .customEndpointUpdate: return 39
    case .dnsZoneCreateCloudAccountGet: return 40
    case .dnsZoneCreate: return 41
    case .dnsZoneDelete: return 42
    case .dnsZoneGet: return 43
    case .dnsZoneRestore: return 44
    case .dnsZoneUpdate: return 45
    case .environmentCreateKubeClusterGet: return 46
    case .environmentCreate: return 47
    case .environmentDelete: return 48
    case .environmentGet: return 49
    case .environmentUpdate: return 50
    case .environmentRestore: return 51
    case .identityGroupCreate: return 52
    case .identityGroupDelete: return 53
    case .identityGroupGet: return 54
    case .identityGroupUpdate: return 55
    case .kubeClusterCreate: return 56
    case .kubeClusterCreateCloudAccountGet: return 57
    case .kubeClusterDelete: return 58
    case .kubeClusterGet: return 59
    case .kubeClusterRestore: return 60
    case .kubeClusterUpdate: return 61
    case .iamPolicyGet: return 62
    case .iamPolicyUpdate: return 63
    case .identityAccountCreate: return 64
    case .identityAccountDelete: return 65
    case .identityAccountGet: return 66
    case .identityAccountRestore: return 67
    case .identityAccountUpdate: return 68
    case .identityConnectionCreate: return 69
    case .identityConnectionDelete: return 70
    case .identityConnectionGet: return 71
    case .identityConnectionRestore: return 72
    case .identityConnectionUpdate: return 73
    case .kafkaClusterCreate: return 74
    case .kafkaClusterDelete: return 75
    case .kafkaClusterGet: return 76
    case .kafkaClusterRestore: return 77
    case .kafkaClusterUpdate: return 78
    case .listAssociates: return 79
    case .loginToBackOffice: return 80
    case .microserviceInstanceCreate: return 81
    case .microserviceInstanceDelete: return 82
    case .microserviceInstanceGet: return 83
    case .microserviceInstanceRestore: return 84
    case .microserviceInstanceUpdate: return 85
    case .postgresClusterCreate: return 86
    case .postgresClusterDelete: return 87
    case .postgresClusterGet: return 88
    case .postgresClusterRestore: return 89
    case .postgresClusterUpdate: return 90
    case .productCreate: return 91
    case .productGet: return 92
    case .productUpdate: return 93
    case .productDelete: return 94
    case .productRestore: return 95
    case .redisClusterCreate: return 96
    case .redisClusterDelete: return 97
    case .redisClusterGet: return 98
    case .redisClusterRestore: return 99
    case .redisClusterUpdate: return 100
    case .solrCloudCreate: return 101
    case .solrCloudDelete: return 102
    case .solrCloudGet: return 103
    case .solrCloudRestore: return 104
    case .solrCloudUpdate: return 105
    case .standardEndpointCreate: return 106
    case .standardEndpointDelete: return 107
    case .standardEndpointGet: return 108
    case .standardEndpointRestore: return 109
    case .standardEndpointUpdate: return 110
    case .storageBucketCreate: return 111
    case .storageBucketDelete: return 112
    case .storageBucketGet: return 113
    case .storageBucketRestore: return 114
    case .storageBucketUpdate: return 115
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_Enums_IamPermission: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static let allCases: [Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_Enums_IamPermission] = [
    .unspecified,
    .artifactStoreCreate,
    .artifactStoreCreateCloudAccountGet,
    .artifactStoreDelete,
    .artifactStoreGet,
    .artifactStoreRestore,
    .artifactStoreUpdate,
    .billingAccountDelete,
    .billingAccountGet,
    .billingAccountRestore,
    .billingAccountUpdate,
    .billingSubscriptionDelete,
    .billingSubscriptionGet,
    .billingSubscriptionRestore,
    .billingSubscriptionUpdate,
    .cloudAccountCreate,
    .cloudAccountDelete,
    .cloudAccountGet,
    .cloudAccountRestore,
    .cloudAccountUpdate,
    .codeProjectCreate,
    .codeProjectGet,
    .codeProjectDelete,
    .codeProjectRestore,
    .codeProjectUpdate,
    .codeServerCreate,
    .codeServerDelete,
    .codeServerGet,
    .codeServerRestore,
    .codeServerSynchronize,
    .codeServerUpdate,
    .companyDelete,
    .companyGet,
    .companyRestore,
    .companyUpdate,
    .customEndpointCreate,
    .customEndpointDelete,
    .customEndpointGet,
    .customEndpointRestore,
    .customEndpointUpdate,
    .dnsZoneCreateCloudAccountGet,
    .dnsZoneCreate,
    .dnsZoneDelete,
    .dnsZoneGet,
    .dnsZoneRestore,
    .dnsZoneUpdate,
    .environmentCreateKubeClusterGet,
    .environmentCreate,
    .environmentDelete,
    .environmentGet,
    .environmentUpdate,
    .environmentRestore,
    .identityGroupCreate,
    .identityGroupDelete,
    .identityGroupGet,
    .identityGroupUpdate,
    .kubeClusterCreate,
    .kubeClusterCreateCloudAccountGet,
    .kubeClusterDelete,
    .kubeClusterGet,
    .kubeClusterRestore,
    .kubeClusterUpdate,
    .iamPolicyGet,
    .iamPolicyUpdate,
    .identityAccountCreate,
    .identityAccountDelete,
    .identityAccountGet,
    .identityAccountRestore,
    .identityAccountUpdate,
    .identityConnectionCreate,
    .identityConnectionDelete,
    .identityConnectionGet,
    .identityConnectionRestore,
    .identityConnectionUpdate,
    .kafkaClusterCreate,
    .kafkaClusterDelete,
    .kafkaClusterGet,
    .kafkaClusterRestore,
    .kafkaClusterUpdate,
    .listAssociates,
    .loginToBackOffice,
    .microserviceInstanceCreate,
    .microserviceInstanceDelete,
    .microserviceInstanceGet,
    .microserviceInstanceRestore,
    .microserviceInstanceUpdate,
    .postgresClusterCreate,
    .postgresClusterDelete,
    .postgresClusterGet,
    .postgresClusterRestore,
    .postgresClusterUpdate,
    .productCreate,
    .productGet,
    .productUpdate,
    .productDelete,
    .productRestore,
    .redisClusterCreate,
    .redisClusterDelete,
    .redisClusterGet,
    .redisClusterRestore,
    .redisClusterUpdate,
    .solrCloudCreate,
    .solrCloudDelete,
    .solrCloudGet,
    .solrCloudRestore,
    .solrCloudUpdate,
    .standardEndpointCreate,
    .standardEndpointDelete,
    .standardEndpointGet,
    .standardEndpointRestore,
    .standardEndpointUpdate,
    .storageBucketCreate,
    .storageBucketDelete,
    .storageBucketGet,
    .storageBucketRestore,
    .storageBucketUpdate,
  ]
}

#endif  // swift(>=4.2)

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_Enums_IamPermission: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

extension Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_Enums_IamPermission: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "IAM_PERMISSION_UNSPECIFIED"),
    1: .same(proto: "artifact_store_create"),
    2: .same(proto: "artifact_store_create_cloud_account_get"),
    3: .same(proto: "artifact_store_delete"),
    4: .same(proto: "artifact_store_get"),
    5: .same(proto: "artifact_store_restore"),
    6: .same(proto: "artifact_store_update"),
    7: .same(proto: "billing_account_delete"),
    8: .same(proto: "billing_account_get"),
    9: .same(proto: "billing_account_restore"),
    10: .same(proto: "billing_account_update"),
    11: .same(proto: "billing_subscription_delete"),
    12: .same(proto: "billing_subscription_get"),
    13: .same(proto: "billing_subscription_restore"),
    14: .same(proto: "billing_subscription_update"),
    15: .same(proto: "cloud_account_create"),
    16: .same(proto: "cloud_account_delete"),
    17: .same(proto: "cloud_account_get"),
    18: .same(proto: "cloud_account_restore"),
    19: .same(proto: "cloud_account_update"),
    20: .same(proto: "code_project_create"),
    21: .same(proto: "code_project_get"),
    22: .same(proto: "code_project_delete"),
    23: .same(proto: "code_project_restore"),
    24: .same(proto: "code_project_update"),
    25: .same(proto: "code_server_create"),
    26: .same(proto: "code_server_delete"),
    27: .same(proto: "code_server_get"),
    28: .same(proto: "code_server_restore"),
    29: .same(proto: "code_server_synchronize"),
    30: .same(proto: "code_server_update"),
    31: .same(proto: "company_delete"),
    32: .same(proto: "company_get"),
    33: .same(proto: "company_restore"),
    34: .same(proto: "company_update"),
    35: .same(proto: "custom_endpoint_create"),
    36: .same(proto: "custom_endpoint_delete"),
    37: .same(proto: "custom_endpoint_get"),
    38: .same(proto: "custom_endpoint_restore"),
    39: .same(proto: "custom_endpoint_update"),
    40: .same(proto: "dns_zone_create_cloud_account_get"),
    41: .same(proto: "dns_zone_create"),
    42: .same(proto: "dns_zone_delete"),
    43: .same(proto: "dns_zone_get"),
    44: .same(proto: "dns_zone_restore"),
    45: .same(proto: "dns_zone_update"),
    46: .same(proto: "environment_create_kube_cluster_get"),
    47: .same(proto: "environment_create"),
    48: .same(proto: "environment_delete"),
    49: .same(proto: "environment_get"),
    50: .same(proto: "environment_update"),
    51: .same(proto: "environment_restore"),
    52: .same(proto: "identity_group_create"),
    53: .same(proto: "identity_group_delete"),
    54: .same(proto: "identity_group_get"),
    55: .same(proto: "identity_group_update"),
    56: .same(proto: "kube_cluster_create"),
    57: .same(proto: "kube_cluster_create_cloud_account_get"),
    58: .same(proto: "kube_cluster_delete"),
    59: .same(proto: "kube_cluster_get"),
    60: .same(proto: "kube_cluster_restore"),
    61: .same(proto: "kube_cluster_update"),
    62: .same(proto: "iam_policy_get"),
    63: .same(proto: "iam_policy_update"),
    64: .same(proto: "identity_account_create"),
    65: .same(proto: "identity_account_delete"),
    66: .same(proto: "identity_account_get"),
    67: .same(proto: "identity_account_restore"),
    68: .same(proto: "identity_account_update"),
    69: .same(proto: "identity_connection_create"),
    70: .same(proto: "identity_connection_delete"),
    71: .same(proto: "identity_connection_get"),
    72: .same(proto: "identity_connection_restore"),
    73: .same(proto: "identity_connection_update"),
    74: .same(proto: "kafka_cluster_create"),
    75: .same(proto: "kafka_cluster_delete"),
    76: .same(proto: "kafka_cluster_get"),
    77: .same(proto: "kafka_cluster_restore"),
    78: .same(proto: "kafka_cluster_update"),
    79: .same(proto: "list_associates"),
    80: .same(proto: "login_to_back_office"),
    81: .same(proto: "microservice_instance_create"),
    82: .same(proto: "microservice_instance_delete"),
    83: .same(proto: "microservice_instance_get"),
    84: .same(proto: "microservice_instance_restore"),
    85: .same(proto: "microservice_instance_update"),
    86: .same(proto: "postgres_cluster_create"),
    87: .same(proto: "postgres_cluster_delete"),
    88: .same(proto: "postgres_cluster_get"),
    89: .same(proto: "postgres_cluster_restore"),
    90: .same(proto: "postgres_cluster_update"),
    91: .same(proto: "product_create"),
    92: .same(proto: "product_get"),
    93: .same(proto: "product_update"),
    94: .same(proto: "product_delete"),
    95: .same(proto: "product_restore"),
    96: .same(proto: "redis_cluster_create"),
    97: .same(proto: "redis_cluster_delete"),
    98: .same(proto: "redis_cluster_get"),
    99: .same(proto: "redis_cluster_restore"),
    100: .same(proto: "redis_cluster_update"),
    101: .same(proto: "solr_cloud_create"),
    102: .same(proto: "solr_cloud_delete"),
    103: .same(proto: "solr_cloud_get"),
    104: .same(proto: "solr_cloud_restore"),
    105: .same(proto: "solr_cloud_update"),
    106: .same(proto: "standard_endpoint_create"),
    107: .same(proto: "standard_endpoint_delete"),
    108: .same(proto: "standard_endpoint_get"),
    109: .same(proto: "standard_endpoint_restore"),
    110: .same(proto: "standard_endpoint_update"),
    111: .same(proto: "storage_bucket_create"),
    112: .same(proto: "storage_bucket_delete"),
    113: .same(proto: "storage_bucket_get"),
    114: .same(proto: "storage_bucket_restore"),
    115: .same(proto: "storage_bucket_update"),
  ]
}