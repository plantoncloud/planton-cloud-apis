// Code generated by protoc-gen-connect-swift. DO NOT EDIT.
//
// Source: cloud/planton/apis/v1/iam/identity/connection/rpc/service.proto
//

import Connect
import Foundation
import SwiftProtobuf

/// identity connection command controller
public protocol Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionCommandControllerClientInterface {

    /// create new identity connection
    @available(iOS 13, *)
    func `create`(request: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection>

    /// update an existing identity connection
    @available(iOS 13, *)
    func `update`(request: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection>

    /// delete an existing identity connection
    @available(iOS 13, *)
    func `delete`(request: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection>

    /// restore an existing identity connection
    @available(iOS 13, *)
    func `restore`(request: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionCommandControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionCommandControllerClient: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionCommandControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `create`(request: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.identity.connection.rpc.IdentityConnectionCommandController/create", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `update`(request: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.identity.connection.rpc.IdentityConnectionCommandController/update", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `delete`(request: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.identity.connection.rpc.IdentityConnectionCommandController/delete", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `restore`(request: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.identity.connection.rpc.IdentityConnectionCommandController/restore", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let create = Connect.MethodSpec(name: "create", service: "cloud.planton.apis.v1.iam.identity.connection.rpc.IdentityConnectionCommandController", type: .unary)
            public static let update = Connect.MethodSpec(name: "update", service: "cloud.planton.apis.v1.iam.identity.connection.rpc.IdentityConnectionCommandController", type: .unary)
            public static let delete = Connect.MethodSpec(name: "delete", service: "cloud.planton.apis.v1.iam.identity.connection.rpc.IdentityConnectionCommandController", type: .unary)
            public static let restore = Connect.MethodSpec(name: "restore", service: "cloud.planton.apis.v1.iam.identity.connection.rpc.IdentityConnectionCommandController", type: .unary)
        }
    }
}

/// identity connection query controller
public protocol Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionQueryControllerClientInterface {

    /// lookup identity account by id.
    @available(iOS 13, *)
    func `getByID`(request: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection>

    /// retrieve paginated list of all identity connections on planton cloud. this is intended for use on portal.
    @available(iOS 13, *)
    func `findByCompanyID`(request: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionCompanyId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnections>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionQueryControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionQueryControllerClient: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionQueryControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `getByID`(request: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnection> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.identity.connection.rpc.IdentityConnectionQueryController/getById", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findByCompanyID`(request: Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnectionCompanyId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Iam_Identity_Connection_Rpc_IdentityConnections> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.identity.connection.rpc.IdentityConnectionQueryController/findByCompanyId", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let getByID = Connect.MethodSpec(name: "getById", service: "cloud.planton.apis.v1.iam.identity.connection.rpc.IdentityConnectionQueryController", type: .unary)
            public static let findByCompanyID = Connect.MethodSpec(name: "findByCompanyId", service: "cloud.planton.apis.v1.iam.identity.connection.rpc.IdentityConnectionQueryController", type: .unary)
        }
    }
}
