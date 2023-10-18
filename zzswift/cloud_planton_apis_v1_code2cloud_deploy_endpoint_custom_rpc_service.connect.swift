// Code generated by protoc-gen-connect-swift. DO NOT EDIT.
//
// Source: cloud/planton/apis/v1/code2cloud/deploy/endpoint/custom/rpc/service.proto
//

import Connect
import Foundation
import SwiftProtobuf

///custom-endpoint command controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointCommandControllerClientInterface {

    ///create custom-endpoint
    @available(iOS 13, *)
    func `create`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint>

    ///update an existing custom-endpoint
    @available(iOS 13, *)
    func `update`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint>

    ///delete custom-endpoint
    @available(iOS 13, *)
    func `delete`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint>

    ///restore a deleted custom-endpoint
    @available(iOS 13, *)
    func `restore`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointCommandControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointCommandControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointCommandControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `create`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController/create", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `update`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController/update", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `delete`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController/delete", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `restore`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController/restore", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let create = Connect.MethodSpec(name: "create", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController", type: .unary)
            public static let update = Connect.MethodSpec(name: "update", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController", type: .unary)
            public static let delete = Connect.MethodSpec(name: "delete", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController", type: .unary)
            public static let restore = Connect.MethodSpec(name: "restore", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointCommandController", type: .unary)
        }
    }
}

///custom-endpoint query controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointQueryControllerClientInterface {

    ///list all custom-endpoints on planton cloud for the requested page. This is intended for use on portal.
    @available(iOS 13, *)
    func `list`(request: Cloud_Planton_Apis_V1_Commons_Rpc_Pagination_PageInfo, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointList>

    ///look up custom-endpoint using custom-endpoint id
    @available(iOS 13, *)
    func `getByID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint>

    ///find custom-endpoints by product id.
    ///response contains only the endpoint domains that the authenticated user account id as viewer access to.
    @available(iOS 13, *)
    func `findByProductID`(request: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoints>

    ///check cert status for custom-endpoint
    @available(iOS 13, *)
    func `getCustomEndpointCertStatus`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointCertStatus>

    ///check status of dns resolution for custom-endpoint.
    ///confirms if the dns of the custom-endpoint domain is resolving to the correct address.
    @available(iOS 13, *)
    func `getCustomEndpointDsnResolutionStatus`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointDnsResolutionStatus>

    ///find certificates for custom-endpoint
    @available(iOS 13, *)
    func `findCustomEndpointCertificates`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Certificates>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointQueryControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointQueryControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointQueryControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `list`(request: Cloud_Planton_Apis_V1_Commons_Rpc_Pagination_PageInfo, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointList> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController/list", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getByID`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController/getById", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findByProductID`(request: Cloud_Planton_Apis_V1_Resourcemanager_Product_Rpc_ProductId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoints> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController/findByProductId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getCustomEndpointCertStatus`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointCertStatus> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController/getCustomEndpointCertStatus", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getCustomEndpointDsnResolutionStatus`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointDnsResolutionStatus> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController/getCustomEndpointDsnResolutionStatus", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `findCustomEndpointCertificates`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Integration_Kubernetes_Resource_Certificates> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController/findCustomEndpointCertificates", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let list = Connect.MethodSpec(name: "list", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController", type: .unary)
            public static let getByID = Connect.MethodSpec(name: "getById", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController", type: .unary)
            public static let findByProductID = Connect.MethodSpec(name: "findByProductId", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController", type: .unary)
            public static let getCustomEndpointCertStatus = Connect.MethodSpec(name: "getCustomEndpointCertStatus", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController", type: .unary)
            public static let getCustomEndpointDsnResolutionStatus = Connect.MethodSpec(name: "getCustomEndpointDsnResolutionStatus", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController", type: .unary)
            public static let findCustomEndpointCertificates = Connect.MethodSpec(name: "findCustomEndpointCertificates", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointQueryController", type: .unary)
        }
    }
}

///custom-endpoint routes command controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointRouteCommandControllerClientInterface {

    ///add a route to a custom-endpoint
    @available(iOS 13, *)
    func `add`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_AddOrUpdateCustomEndpointRouteCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint>

    ///update an existing route of a custom-endpoint
    @available(iOS 13, *)
    func `update`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_AddOrUpdateCustomEndpointRouteCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint>

    ///delete a route for a custom-endpoint.
    @available(iOS 13, *)
    func `delete`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_DeleteOrRestoreCustomEndpointRouteCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointRouteCommandControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointRouteCommandControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointRouteCommandControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `add`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_AddOrUpdateCustomEndpointRouteCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointRouteCommandController/add", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `update`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_AddOrUpdateCustomEndpointRouteCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointRouteCommandController/update", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `delete`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_DeleteOrRestoreCustomEndpointRouteCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointRouteCommandController/delete", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let add = Connect.MethodSpec(name: "add", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointRouteCommandController", type: .unary)
            public static let update = Connect.MethodSpec(name: "update", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointRouteCommandController", type: .unary)
            public static let delete = Connect.MethodSpec(name: "delete", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointRouteCommandController", type: .unary)
        }
    }
}

///custom-endpoint stack controller
public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointStackControllerClientInterface {

    ///preview custom-endpoint stack
    @available(iOS 13, *)
    func `preview`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint>

    ///apply custom-endpoint stack
    @available(iOS 13, *)
    func `apply`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointStackControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointStackControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointStackControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `preview`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointStackController/preview", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `apply`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpointId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Endpoint_Custom_Rpc_CustomEndpoint> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointStackController/apply", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let preview = Connect.MethodSpec(name: "preview", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointStackController", type: .unary)
            public static let apply = Connect.MethodSpec(name: "apply", service: "cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpointStackController", type: .unary)
        }
    }
}
