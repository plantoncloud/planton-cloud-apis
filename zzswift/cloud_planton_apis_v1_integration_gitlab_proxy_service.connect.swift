// Code generated by protoc-gen-connect-swift. DO NOT EDIT.
//
// Source: cloud/planton/apis/v1/integration/gitlab/proxy/service.proto
//

import Connect
import Foundation
import SwiftProtobuf

///gitlab proxy command controller
public protocol Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GitlabProxyCommandControllerClientInterface {

    ///create new project on gitlab
    ///https://docs.gitlab.com/ee/api/projects.html#create-project
    @available(iOS 13, *)
    func `creProject`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_CreProjectCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProject>

    ///apply a cookiecutter template on a code project created on gitlab
    @available(iOS 13, *)
    func `applyTemplate`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GitlabApplyTemplateCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProjectProfile>

    ///add a list of variables to a gitlab project
    @available(iOS 13, *)
    func `addVariablesToProject`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_AddVariablesToProjectCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProject>

    ///add a list of variables to a gitlab group
    @available(iOS 13, *)
    func `addVariablesToGroup`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_AddVariablesToGroupCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServer>

    ///add a list of files to a gitlab project
    @available(iOS 13, *)
    func `addFilesToProject`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_AddFilesToGitlabProjectCommandInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProject>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GitlabProxyCommandControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GitlabProxyCommandControllerClient: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GitlabProxyCommandControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `creProject`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_CreProjectCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProject> {
        return await self.client.unary(path: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController/creProject", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `applyTemplate`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GitlabApplyTemplateCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProjectProfile> {
        return await self.client.unary(path: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController/applyTemplate", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `addVariablesToProject`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_AddVariablesToProjectCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProject> {
        return await self.client.unary(path: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController/addVariablesToProject", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `addVariablesToGroup`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_AddVariablesToGroupCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServer> {
        return await self.client.unary(path: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController/addVariablesToGroup", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `addFilesToProject`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_AddFilesToGitlabProjectCommandInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProject> {
        return await self.client.unary(path: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController/addFilesToProject", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let creProject = Connect.MethodSpec(name: "creProject", service: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController", type: .unary)
            public static let applyTemplate = Connect.MethodSpec(name: "applyTemplate", service: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController", type: .unary)
            public static let addVariablesToProject = Connect.MethodSpec(name: "addVariablesToProject", service: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController", type: .unary)
            public static let addVariablesToGroup = Connect.MethodSpec(name: "addVariablesToGroup", service: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController", type: .unary)
            public static let addFilesToProject = Connect.MethodSpec(name: "addFilesToProject", service: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyCommandController", type: .unary)
        }
    }
}

///gitlab proxy query controller
public protocol Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GitlabProxyQueryControllerClientInterface {

    ///list projects for the requested group, including projects in sub-groups on gitlab
    ///https://docs.gitlab.com/ee/api/groups.html#list-a-groups-projects
    ///todo: we have to add pagination support for response.
    @available(iOS 13, *)
    func `listProjects`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_ListProjectsQueryInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProjects>

    ///get details of a project on gitlab
    @available(iOS 13, *)
    func `getProject`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GetProjectQueryInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProject>

    ///get details of a group on gitlab
    @available(iOS 13, *)
    func `getGroup`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GetGroupQueryInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServer>

    ///get code project profile of a code-project hosted on gitlab
    @available(iOS 13, *)
    func `getGitlabCodeProjectProfile`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GetGitlabCodeProjectProfileQueryInput, headers: Connect.Headers) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProjectProfile>
}

/// Concrete implementation of `Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GitlabProxyQueryControllerClientInterface`.
public final class Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GitlabProxyQueryControllerClient: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GitlabProxyQueryControllerClientInterface {
    private let client: Connect.ProtocolClientInterface

    public init(client: Connect.ProtocolClientInterface) {
        self.client = client
    }

    @available(iOS 13, *)
    public func `listProjects`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_ListProjectsQueryInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProjects> {
        return await self.client.unary(path: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController/listProjects", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getProject`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GetProjectQueryInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProject> {
        return await self.client.unary(path: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController/getProject", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getGroup`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GetGroupQueryInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Server_Rpc_CodeServer> {
        return await self.client.unary(path: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController/getGroup", request: request, headers: headers)
    }

    @available(iOS 13, *)
    public func `getGitlabCodeProjectProfile`(request: Cloud_Planton_Apis_V1_Integration_Gitlab_Proxy_GetGitlabCodeProjectProfileQueryInput, headers: Connect.Headers = [:]) async -> ResponseMessage<Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProjectProfile> {
        return await self.client.unary(path: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController/getGitlabCodeProjectProfile", request: request, headers: headers)
    }

    public enum Metadata {
        public enum Methods {
            public static let listProjects = Connect.MethodSpec(name: "listProjects", service: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController", type: .unary)
            public static let getProject = Connect.MethodSpec(name: "getProject", service: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController", type: .unary)
            public static let getGroup = Connect.MethodSpec(name: "getGroup", service: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController", type: .unary)
            public static let getGitlabCodeProjectProfile = Connect.MethodSpec(name: "getGitlabCodeProjectProfile", service: "cloud.planton.apis.v1.integration.gitlab.proxy.GitlabProxyQueryController", type: .unary)
        }
    }
}
