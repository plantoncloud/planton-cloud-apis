// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: cloud/planton/apis/v1/code2cloud/develop/sourcecode/pipeline/rpc/model.proto
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

///input for query rpc to get dynamically generated code pipeline template
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_GenerateCodePipelineTemplateQueryInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///code project profile
  public var codeProjectProfile: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProjectProfile {
    get {return _codeProjectProfile ?? Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProjectProfile()}
    set {_codeProjectProfile = newValue}
  }
  /// Returns true if `codeProjectProfile` has been explicitly set.
  public var hasCodeProjectProfile: Bool {return self._codeProjectProfile != nil}
  /// Clears the value of `codeProjectProfile`. Subsequent reads from it will return its default value.
  public mutating func clearCodeProjectProfile() {self._codeProjectProfile = nil}

  ///git remote url used for looking up code project
  public var gitRemoteURL: String = String()

  ///list of microservice-instances declared in _kustomize directory
  public var microserviceInstances: [Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Rpc_MicroserviceInstance] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _codeProjectProfile: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Project_Rpc_CodeProjectProfile? = nil
}

public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_GenerateCodePipelineTemplateQueryResp {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///input used for generating the template
  public var templateInput: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineBuildTemplateInput {
    get {return _templateInput ?? Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineBuildTemplateInput()}
    set {_templateInput = newValue}
  }
  /// Returns true if `templateInput` has been explicitly set.
  public var hasTemplateInput: Bool {return self._templateInput != nil}
  /// Clears the value of `templateInput`. Subsequent reads from it will return its default value.
  public mutating func clearTemplateInput() {self._templateInput = nil}

  ///template generated based on the provided inputs
  public var templateOutput: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _templateInput: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineBuildTemplateInput? = nil
}

///container(docker) images to be used in code pipeline templates
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineTemplateContainerImages {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var dockerBuild: String = String()

  public var golangBuild: String = String()

  public var javaBuild: String = String()

  public var javascriptBuild: String = String()

  public var plantonCli: String = String()

  public var protobufBuild: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

///gitlab pipeline template is used as input data model for processing gitlab build job templates
///templating frameworks like apache freemarker in java or the built in templating engine in golang can use this data model as input for rendering the template.
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineBuildTemplateInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///name of the cli environment to be used. this can be test or live.
  public var plantonCloudCliEnvironment: String {
    get {return _storage._plantonCloudCliEnvironment}
    set {_uniqueStorage()._plantonCloudCliEnvironment = newValue}
  }

  ///artifact-store to be used by the pipeline
  public var artifactStore: Cloud_Planton_Apis_V1_Code2cloud_Develop_Artifactstore_Rpc_ArtifactStore {
    get {return _storage._artifactStore ?? Cloud_Planton_Apis_V1_Code2cloud_Develop_Artifactstore_Rpc_ArtifactStore()}
    set {_uniqueStorage()._artifactStore = newValue}
  }
  /// Returns true if `artifactStore` has been explicitly set.
  public var hasArtifactStore: Bool {return _storage._artifactStore != nil}
  /// Clears the value of `artifactStore`. Subsequent reads from it will return its default value.
  public mutating func clearArtifactStore() {_uniqueStorage()._artifactStore = nil}

  ///container images(docker) to be used inside code pipeline templates
  public var templateContainerImages: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineTemplateContainerImages {
    get {return _storage._templateContainerImages ?? Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineTemplateContainerImages()}
    set {_uniqueStorage()._templateContainerImages = newValue}
  }
  /// Returns true if `templateContainerImages` has been explicitly set.
  public var hasTemplateContainerImages: Bool {return _storage._templateContainerImages != nil}
  /// Clears the value of `templateContainerImages`. Subsequent reads from it will return its default value.
  public mutating func clearTemplateContainerImages() {_uniqueStorage()._templateContainerImages = nil}

  ///container image to be used in container image build jobs.
  ///container image tag is not included in the value of this attribute.
  ///git commit short sha is used as the tag for the image.
  public var projectContainerImage: String {
    get {return _storage._projectContainerImage}
    set {_uniqueStorage()._projectContainerImage = newValue}
  }

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _storage = _StorageClass.defaultInstance
}

///input used for rendering job template for microservice deployment to each environment
public struct Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineDeployTemplateJobInput {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///container images(docker) to be used inside code pipeline templates
  public var templateContainerImages: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineTemplateContainerImages {
    get {return _templateContainerImages ?? Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineTemplateContainerImages()}
    set {_templateContainerImages = newValue}
  }
  /// Returns true if `templateContainerImages` has been explicitly set.
  public var hasTemplateContainerImages: Bool {return self._templateContainerImages != nil}
  /// Clears the value of `templateContainerImages`. Subsequent reads from it will return its default value.
  public mutating func clearTemplateContainerImages() {self._templateContainerImages = nil}

  ///environment
  public var environment: Cloud_Planton_Apis_V1_Code2cloud_Environment_Rpc_Environment {
    get {return _environment ?? Cloud_Planton_Apis_V1_Code2cloud_Environment_Rpc_Environment()}
    set {_environment = newValue}
  }
  /// Returns true if `environment` has been explicitly set.
  public var hasEnvironment: Bool {return self._environment != nil}
  /// Clears the value of `environment`. Subsequent reads from it will return its default value.
  public mutating func clearEnvironment() {self._environment = nil}

  ///microservice-instance
  public var microserviceInstance: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Rpc_MicroserviceInstance {
    get {return _microserviceInstance ?? Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Rpc_MicroserviceInstance()}
    set {_microserviceInstance = newValue}
  }
  /// Returns true if `microserviceInstance` has been explicitly set.
  public var hasMicroserviceInstance: Bool {return self._microserviceInstance != nil}
  /// Clears the value of `microserviceInstance`. Subsequent reads from it will return its default value.
  public mutating func clearMicroserviceInstance() {self._microserviceInstance = nil}

  ///container image tag is not included in the value of this attribute.
  ///git commit short sha is used as the tag for the image.
  public var projectContainerImage: String = String()

  ///flag to indicate if manual gate is enabled for the deployment environment.
  ///this flag is used for setting `when` attribute of the gitlab pipeline https://docs.gitlab.com/ee/ci/yaml/#when
  public var isManualGateRequired: Bool = false

  ///name of the cli environment to be used. this can be test or live.
  public var plantonCloudCliEnvironment: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _templateContainerImages: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineTemplateContainerImages? = nil
  fileprivate var _environment: Cloud_Planton_Apis_V1_Code2cloud_Environment_Rpc_Environment? = nil
  fileprivate var _microserviceInstance: Cloud_Planton_Apis_V1_Code2cloud_Deploy_Microservice_Rpc_MicroserviceInstance? = nil
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_GenerateCodePipelineTemplateQueryInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_GenerateCodePipelineTemplateQueryResp: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineTemplateContainerImages: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineBuildTemplateInput: @unchecked Sendable {}
extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineDeployTemplateJobInput: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.planton.apis.v1.code2cloud.develop.sourcecode.pipeline.common.rpc"

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_GenerateCodePipelineTemplateQueryInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".GenerateCodePipelineTemplateQueryInput"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "code_project_profile"),
    2: .standard(proto: "git_remote_url"),
    3: .standard(proto: "microservice_instances"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._codeProjectProfile) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.gitRemoteURL) }()
      case 3: try { try decoder.decodeRepeatedMessageField(value: &self.microserviceInstances) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._codeProjectProfile {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    if !self.gitRemoteURL.isEmpty {
      try visitor.visitSingularStringField(value: self.gitRemoteURL, fieldNumber: 2)
    }
    if !self.microserviceInstances.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.microserviceInstances, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_GenerateCodePipelineTemplateQueryInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_GenerateCodePipelineTemplateQueryInput) -> Bool {
    if lhs._codeProjectProfile != rhs._codeProjectProfile {return false}
    if lhs.gitRemoteURL != rhs.gitRemoteURL {return false}
    if lhs.microserviceInstances != rhs.microserviceInstances {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_GenerateCodePipelineTemplateQueryResp: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".GenerateCodePipelineTemplateQueryResp"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "template_input"),
    2: .standard(proto: "template_output"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._templateInput) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.templateOutput) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._templateInput {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    if !self.templateOutput.isEmpty {
      try visitor.visitSingularStringField(value: self.templateOutput, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_GenerateCodePipelineTemplateQueryResp, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_GenerateCodePipelineTemplateQueryResp) -> Bool {
    if lhs._templateInput != rhs._templateInput {return false}
    if lhs.templateOutput != rhs.templateOutput {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineTemplateContainerImages: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CodePipelineTemplateContainerImages"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "docker_build"),
    2: .standard(proto: "golang_build"),
    3: .standard(proto: "java_build"),
    4: .standard(proto: "javascript_build"),
    5: .standard(proto: "planton_cli"),
    6: .standard(proto: "protobuf_build"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.dockerBuild) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.golangBuild) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.javaBuild) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.javascriptBuild) }()
      case 5: try { try decoder.decodeSingularStringField(value: &self.plantonCli) }()
      case 6: try { try decoder.decodeSingularStringField(value: &self.protobufBuild) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.dockerBuild.isEmpty {
      try visitor.visitSingularStringField(value: self.dockerBuild, fieldNumber: 1)
    }
    if !self.golangBuild.isEmpty {
      try visitor.visitSingularStringField(value: self.golangBuild, fieldNumber: 2)
    }
    if !self.javaBuild.isEmpty {
      try visitor.visitSingularStringField(value: self.javaBuild, fieldNumber: 3)
    }
    if !self.javascriptBuild.isEmpty {
      try visitor.visitSingularStringField(value: self.javascriptBuild, fieldNumber: 4)
    }
    if !self.plantonCli.isEmpty {
      try visitor.visitSingularStringField(value: self.plantonCli, fieldNumber: 5)
    }
    if !self.protobufBuild.isEmpty {
      try visitor.visitSingularStringField(value: self.protobufBuild, fieldNumber: 6)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineTemplateContainerImages, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineTemplateContainerImages) -> Bool {
    if lhs.dockerBuild != rhs.dockerBuild {return false}
    if lhs.golangBuild != rhs.golangBuild {return false}
    if lhs.javaBuild != rhs.javaBuild {return false}
    if lhs.javascriptBuild != rhs.javascriptBuild {return false}
    if lhs.plantonCli != rhs.plantonCli {return false}
    if lhs.protobufBuild != rhs.protobufBuild {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineBuildTemplateInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CodePipelineBuildTemplateInput"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "planton_cloud_cli_environment"),
    2: .standard(proto: "artifact_store"),
    3: .standard(proto: "template_container_images"),
    4: .standard(proto: "project_container_image"),
  ]

  fileprivate class _StorageClass {
    var _plantonCloudCliEnvironment: String = String()
    var _artifactStore: Cloud_Planton_Apis_V1_Code2cloud_Develop_Artifactstore_Rpc_ArtifactStore? = nil
    var _templateContainerImages: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineTemplateContainerImages? = nil
    var _projectContainerImage: String = String()

    static let defaultInstance = _StorageClass()

    private init() {}

    init(copying source: _StorageClass) {
      _plantonCloudCliEnvironment = source._plantonCloudCliEnvironment
      _artifactStore = source._artifactStore
      _templateContainerImages = source._templateContainerImages
      _projectContainerImage = source._projectContainerImage
    }
  }

  fileprivate mutating func _uniqueStorage() -> _StorageClass {
    if !isKnownUniquelyReferenced(&_storage) {
      _storage = _StorageClass(copying: _storage)
    }
    return _storage
  }

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    _ = _uniqueStorage()
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      while let fieldNumber = try decoder.nextFieldNumber() {
        // The use of inline closures is to circumvent an issue where the compiler
        // allocates stack space for every case branch when no optimizations are
        // enabled. https://github.com/apple/swift-protobuf/issues/1034
        switch fieldNumber {
        case 1: try { try decoder.decodeSingularStringField(value: &_storage._plantonCloudCliEnvironment) }()
        case 2: try { try decoder.decodeSingularMessageField(value: &_storage._artifactStore) }()
        case 3: try { try decoder.decodeSingularMessageField(value: &_storage._templateContainerImages) }()
        case 4: try { try decoder.decodeSingularStringField(value: &_storage._projectContainerImage) }()
        default: break
        }
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try withExtendedLifetime(_storage) { (_storage: _StorageClass) in
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every if/case branch local when no optimizations
      // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
      // https://github.com/apple/swift-protobuf/issues/1182
      if !_storage._plantonCloudCliEnvironment.isEmpty {
        try visitor.visitSingularStringField(value: _storage._plantonCloudCliEnvironment, fieldNumber: 1)
      }
      try { if let v = _storage._artifactStore {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
      } }()
      try { if let v = _storage._templateContainerImages {
        try visitor.visitSingularMessageField(value: v, fieldNumber: 3)
      } }()
      if !_storage._projectContainerImage.isEmpty {
        try visitor.visitSingularStringField(value: _storage._projectContainerImage, fieldNumber: 4)
      }
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineBuildTemplateInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineBuildTemplateInput) -> Bool {
    if lhs._storage !== rhs._storage {
      let storagesAreEqual: Bool = withExtendedLifetime((lhs._storage, rhs._storage)) { (_args: (_StorageClass, _StorageClass)) in
        let _storage = _args.0
        let rhs_storage = _args.1
        if _storage._plantonCloudCliEnvironment != rhs_storage._plantonCloudCliEnvironment {return false}
        if _storage._artifactStore != rhs_storage._artifactStore {return false}
        if _storage._templateContainerImages != rhs_storage._templateContainerImages {return false}
        if _storage._projectContainerImage != rhs_storage._projectContainerImage {return false}
        return true
      }
      if !storagesAreEqual {return false}
    }
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineDeployTemplateJobInput: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CodePipelineDeployTemplateJobInput"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "template_container_images"),
    2: .same(proto: "environment"),
    3: .standard(proto: "microservice_instance"),
    4: .standard(proto: "project_container_image"),
    5: .standard(proto: "is_manual_gate_required"),
    6: .standard(proto: "planton_cloud_cli_environment"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._templateContainerImages) }()
      case 2: try { try decoder.decodeSingularMessageField(value: &self._environment) }()
      case 3: try { try decoder.decodeSingularMessageField(value: &self._microserviceInstance) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.projectContainerImage) }()
      case 5: try { try decoder.decodeSingularBoolField(value: &self.isManualGateRequired) }()
      case 6: try { try decoder.decodeSingularStringField(value: &self.plantonCloudCliEnvironment) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._templateContainerImages {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try { if let v = self._environment {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    } }()
    try { if let v = self._microserviceInstance {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 3)
    } }()
    if !self.projectContainerImage.isEmpty {
      try visitor.visitSingularStringField(value: self.projectContainerImage, fieldNumber: 4)
    }
    if self.isManualGateRequired != false {
      try visitor.visitSingularBoolField(value: self.isManualGateRequired, fieldNumber: 5)
    }
    if !self.plantonCloudCliEnvironment.isEmpty {
      try visitor.visitSingularStringField(value: self.plantonCloudCliEnvironment, fieldNumber: 6)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineDeployTemplateJobInput, rhs: Cloud_Planton_Apis_V1_Code2cloud_Develop_Sourcecode_Pipeline_Common_Rpc_CodePipelineDeployTemplateJobInput) -> Bool {
    if lhs._templateContainerImages != rhs._templateContainerImages {return false}
    if lhs._environment != rhs._environment {return false}
    if lhs._microserviceInstance != rhs._microserviceInstance {return false}
    if lhs.projectContainerImage != rhs.projectContainerImage {return false}
    if lhs.isManualGateRequired != rhs.isManualGateRequired {return false}
    if lhs.plantonCloudCliEnvironment != rhs.plantonCloudCliEnvironment {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
