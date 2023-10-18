// Code generated by protoc-gen-connect-swift. DO NOT EDIT.
//
// Source: cloud/planton/apis/v1/code2cloud/deploy/microservice/stack/kubernetes/service.proto
//

import Connect
import Foundation
import SwiftProtobuf

public protocol Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Stack_Kubernetes_MicroserviceInstanceKubernetesStackControllerClientInterface {

    @available(iOS 13, *)
    func `execute`(headers: Connect.Headers) -> any Connect.ServerOnlyAsyncStreamInterface<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Stack_Kubernetes_MicroserviceInstanceKubernetesStackInput, Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Stack_Kubernetes_MicroserviceInstanceKubernetesStackResponse>

    @available(iOS 13, *)
    func `getStackOutputs`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Stack_Kubernetes_MicroserviceInstanceKubernetesStackInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Stack_Kubernetes_MicroserviceInstanceKubernetesStackOutputs>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Stack_Kubernetes_MicroserviceInstanceKubernetesStackControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Stack_Kubernetes_MicroserviceInstanceKubernetesStackControllerClient: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Stack_Kubernetes_MicroserviceInstanceKubernetesStackControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `execute`(headers: Connect.Headers = [:]) -> any Connect.ServerOnlyAsyncStreamInterface<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Stack_Kubernetes_MicroserviceInstanceKubernetesStackInput, Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Stack_Kubernetes_MicroserviceInstanceKubernetesStackResponse> {
        return self.client.serverOnlyStream(path: "cloud.planton.apis.v1.code2cloud.deploy.microservice.stack.kubernetes.MicroserviceInstanceKubernetesStackController/execute", headers: headers)
    }

    @available(iOS 13, *)
    public func `getStackOutputs`(request: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Stack_Kubernetes_MicroserviceInstanceKubernetesStackInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Stack_Kubernetes_MicroserviceInstanceKubernetesStackOutputs> {
        return await self.client.unary(path: "cloud.planton.apis.v1.code2cloud.deploy.microservice.stack.kubernetes.MicroserviceInstanceKubernetesStackController/getStackOutputs", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let execute = Connect.MethodSpec(name: "execute", service: "cloud.planton.apis.v1.code2cloud.deploy.microservice.stack.kubernetes.MicroserviceInstanceKubernetesStackController", type: .serverStream)
            public static let getStackOutputs = Connect.MethodSpec(name: "getStackOutputs", service: "cloud.planton.apis.v1.code2cloud.deploy.microservice.stack.kubernetes.MicroserviceInstanceKubernetesStackController", type: .unary)
        }
    }
}