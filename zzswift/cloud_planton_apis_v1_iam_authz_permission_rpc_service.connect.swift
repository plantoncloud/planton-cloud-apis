// Code generated by protoc-gen-connect-swift. DO NOT EDIT.
//
// Source: cloud/planton/apis/v1/iam/authz/permission/rpc/service.proto
//

import Connect
import Foundation
import SwiftProtobuf

///iam command controller
public protocol Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionCommandControllerClientInterface {

    /// create iam permission
    @available(iOS 13, *)
    func `create`(request: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission>

    /// update iam permission
    @available(iOS 13, *)
    func `update`(request: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission>

    /// delete iam permission
    @available(iOS 13, *)
    func `delete`(request: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionCommandControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionCommandControllerClient: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionCommandControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `create`(request: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionCommandController/create", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `update`(request: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionCommandController/update", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `delete`(request: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionCommandController/delete", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let create = Connect.MethodSpec(name: "create", service: "cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionCommandController", type: .unary)
            public static let update = Connect.MethodSpec(name: "update", service: "cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionCommandController", type: .unary)
            public static let delete = Connect.MethodSpec(name: "delete", service: "cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionCommandController", type: .unary)
        }
    }
}

/// iam query controller
public protocol Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionQueryControllerClientInterface {

    /// retrieve paginated list of all iam permissions on planton cloud. this is intended for use on portal.
    @available(iOS 13, *)
    func `list`(request: Cloud_Planton_Apis_V1_Commons_Rpc_Pagination_PageInfo, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionList>

    /// lookup iam permission by permission id
    @available(iOS 13, *)
    func `getByID`(request: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission>

    /// lookup iam permission by permission code
    @available(iOS 13, *)
    func `getByPermissionCode`(request: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionCode, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionQueryControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionQueryControllerClient: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionQueryControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `list`(request: Cloud_Planton_Apis_V1_Commons_Rpc_Pagination_PageInfo, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionList> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionQueryController/list", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getByID`(request: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionQueryController/getById", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getByPermissionCode`(request: Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermissionCode, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Authz_Permission_Rpc_IamPermission> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionQueryController/getByPermissionCode", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let list = Connect.MethodSpec(name: "list", service: "cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionQueryController", type: .unary)
            public static let getByID = Connect.MethodSpec(name: "getById", service: "cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionQueryController", type: .unary)
            public static let getByPermissionCode = Connect.MethodSpec(name: "getByPermissionCode", service: "cloud.planton.apis.v1.iam.authz.permission.rpc.IamPermissionQueryController", type: .unary)
        }
    }
}