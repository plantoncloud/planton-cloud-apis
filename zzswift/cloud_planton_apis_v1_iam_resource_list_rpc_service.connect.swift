// Code generated by protoc-gen-connect-swift. DO NOT EDIT.
//
// Source: cloud/planton/apis/v1/iam/resource/list/rpc/service.proto
//

import Connect
import Foundation
import SwiftProtobuf

/// The `listByResourceType` RPC service method enables clients to execute a search operation
/// based on a specific resource type within a specified company's product.
public protocol Cloud_Planton_Apis_V1_Iam_Resource_List_Rpc_ResourceListViewQueryControllerClientInterface {

    /// This method returns a `ResourceList` message, which encapsulates a list of resources that match
    /// the input search parameters. Each resource in the list should match the specified resource type,
    /// and be associated with the specified company and product.
    @available(iOS 13, *)
    func `listByResourceType`(request: Cloud_Planton_Apis_V1_Commons_Resource_List_Rpc_GetByResourceTypeInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Commons_Resource_List_Rpc_RecordList>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Iam_Resource_List_Rpc_ResourceListViewQueryControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Iam_Resource_List_Rpc_ResourceListViewQueryControllerClient: Cloud_Planton_Apis_V1_Iam_Resource_List_Rpc_ResourceListViewQueryControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `listByResourceType`(request: Cloud_Planton_Apis_V1_Commons_Resource_List_Rpc_GetByResourceTypeInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Commons_Resource_List_Rpc_RecordList> {
        return await self.client.unary(path: "cloud.planton.apis.v1.iam.resource.list.rpc.ResourceListViewQueryController/listByResourceType", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let listByResourceType = Connect.MethodSpec(name: "listByResourceType", service: "cloud.planton.apis.v1.iam.resource.list.rpc.ResourceListViewQueryController", type: .unary)
        }
    }
}
