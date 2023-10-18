// Code generated by protoc-gen-connect-swift. DO NOT EDIT.
//
// Source: cloud/planton/apis/v1/billing/rpc/service.proto
//

import Connect
import Foundation
import SwiftProtobuf

///billing command controller contains rpc definitions for operations that perform mutating/update operations on the billing domain objects.
public protocol Cloud_Planton_Apis_V1_Billing_Rpc_BillingCommandControllerClientInterface {

    ///update an existing billing account.
    @available(iOS 13, *)
    func `updBillingAccount`(request: Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccount, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccount>

    ///delete an existing billing account.
    ///a billing account can not be deleted without deleting all the subscriptions linked to the billing account.
    @available(iOS 13, *)
    func `delBillingAccount`(request: Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccountId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccount>

    ///restore a previously deleted billing account.
    @available(iOS 13, *)
    func `restoreBillingAccount`(request: Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccountId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccount>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Billing_Rpc_BillingCommandControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Billing_Rpc_BillingCommandControllerClient: Cloud_Planton_Apis_V1_Billing_Rpc_BillingCommandControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `updBillingAccount`(request: Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccount, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccount> {
        return await self.client.unary(path: "cloud.planton.apis.v1.billing.rpc.BillingCommandController/updBillingAccount", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `delBillingAccount`(request: Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccountId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccount> {
        return await self.client.unary(path: "cloud.planton.apis.v1.billing.rpc.BillingCommandController/delBillingAccount", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `restoreBillingAccount`(request: Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccountId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccount> {
        return await self.client.unary(path: "cloud.planton.apis.v1.billing.rpc.BillingCommandController/restoreBillingAccount", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let updBillingAccount = Connect.MethodSpec(name: "updBillingAccount", service: "cloud.planton.apis.v1.billing.rpc.BillingCommandController", type: .unary)
            public static let delBillingAccount = Connect.MethodSpec(name: "delBillingAccount", service: "cloud.planton.apis.v1.billing.rpc.BillingCommandController", type: .unary)
            public static let restoreBillingAccount = Connect.MethodSpec(name: "restoreBillingAccount", service: "cloud.planton.apis.v1.billing.rpc.BillingCommandController", type: .unary)
        }
    }
}

///billing account query controller contains rpc definitions for operations that perform read only operations on the billing domain objects.
public protocol Cloud_Planton_Apis_V1_Billing_Rpc_BillingQueryControllerClientInterface {

    ///retrieve paginated list of all billing accounts on planton cloud. this is intended to be used on portal.
    @available(iOS 13, *)
    func `listBillingAccounts`(request: Cloud_Planton_Apis_V1_Commons_Rpc_Pagination_PageInfo, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccountList>

    ///lookup billing account using billing account id.
    @available(iOS 13, *)
    func `getBillingAccountByBillingAccountID`(request: Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccountId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccount>

    ///lookup billing account using company id.
    @available(iOS 13, *)
    func `getBillingAccountByCompanyID`(request: Cloud_Planton_Apis_V1_Resourcemanager_Company_Rpc_CompanyId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccount>

    ///get customer portal session.
    ///https://stripe.com/docs/api/customer_portal/sessions/create
    @available(iOS 13, *)
    func `getCustomerPortalSessionByCompanyID`(request: Cloud_Planton_Apis_V1_Resourcemanager_Company_Rpc_CompanyId, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_GetCustomerPortalSessionQueryResp>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Billing_Rpc_BillingQueryControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Billing_Rpc_BillingQueryControllerClient: Cloud_Planton_Apis_V1_Billing_Rpc_BillingQueryControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `listBillingAccounts`(request: Cloud_Planton_Apis_V1_Commons_Rpc_Pagination_PageInfo, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccountList> {
        return await self.client.unary(path: "cloud.planton.apis.v1.billing.rpc.BillingQueryController/listBillingAccounts", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getBillingAccountByBillingAccountID`(request: Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccountId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccount> {
        return await self.client.unary(path: "cloud.planton.apis.v1.billing.rpc.BillingQueryController/getBillingAccountByBillingAccountId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getBillingAccountByCompanyID`(request: Cloud_Planton_Apis_V1_Resourcemanager_Company_Rpc_CompanyId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_BillingAccount> {
        return await self.client.unary(path: "cloud.planton.apis.v1.billing.rpc.BillingQueryController/getBillingAccountByCompanyId", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getCustomerPortalSessionByCompanyID`(request: Cloud_Planton_Apis_V1_Resourcemanager_Company_Rpc_CompanyId, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Billing_Rpc_GetCustomerPortalSessionQueryResp> {
        return await self.client.unary(path: "cloud.planton.apis.v1.billing.rpc.BillingQueryController/getCustomerPortalSessionByCompanyId", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let listBillingAccounts = Connect.MethodSpec(name: "listBillingAccounts", service: "cloud.planton.apis.v1.billing.rpc.BillingQueryController", type: .unary)
            public static let getBillingAccountByBillingAccountID = Connect.MethodSpec(name: "getBillingAccountByBillingAccountId", service: "cloud.planton.apis.v1.billing.rpc.BillingQueryController", type: .unary)
            public static let getBillingAccountByCompanyID = Connect.MethodSpec(name: "getBillingAccountByCompanyId", service: "cloud.planton.apis.v1.billing.rpc.BillingQueryController", type: .unary)
            public static let getCustomerPortalSessionByCompanyID = Connect.MethodSpec(name: "getCustomerPortalSessionByCompanyId", service: "cloud.planton.apis.v1.billing.rpc.BillingQueryController", type: .unary)
        }
    }
}
