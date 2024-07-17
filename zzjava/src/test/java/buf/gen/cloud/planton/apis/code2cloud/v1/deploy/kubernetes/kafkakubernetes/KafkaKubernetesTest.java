package buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.model.KafkaKubernetes;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.model.KafkaKubernetesSpec;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.model.KafkaKubernetesSpecBrokerContainerSpec;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.kafkakubernetes.model.KafkaKubernetesSpecZookeeperContainerSpec;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.gen.cloud.planton.apis.commons.kubernetes.model.ContainerResources;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class KafkaKubernetesTest {

    @Test
    public void testKafkaKubernetes_ShouldNotThroughProtoValidationException() {
        var input = KafkaKubernetes.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input));
    }

    @Test
    public void testKafkaKubernetes_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var kafkaKubernetes = KafkaKubernetes.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testKafkaKubernetes_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var kafkaKubernetes = KafkaKubernetes.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testKafkaKubernetes_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var kafkaKubernetes = KafkaKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test kafka-kubernetes").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testKafkaKubernetes_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var kafkaKubernetes = KafkaKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test kafka-kubernetes").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testKafkaKubernetes_ShouldReturnValidationErrorIfNameLengthIsGreaterThan12() throws ValidationException {
        var kafkaKubernetes = KafkaKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation, this is test name to check length validation,this is test name to check length validation")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test kafka-kubernetes").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testKafkaKubernetes_ShouldNotReturnValidationErrorIfNameLengthIsLessThan12() throws ValidationException {
        var kafkaKubernetes = KafkaKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test kafka-kubernetes").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testKafkaKubernetes_ShouldReturnValidationErrorIfContainerReplicasAndDiskSizeIsNotSet() throws ValidationException {
        var kafkaKubernetes = KafkaKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test kafka-kubernetes").build())
                        .build())
                .setSpec(KafkaKubernetesSpec.newBuilder()
                                .setZookeeperContainer(KafkaKubernetesSpecZookeeperContainerSpec.newBuilder().build())
                                .setBrokerContainer(KafkaKubernetesSpecBrokerContainerSpec.newBuilder().build())
                                .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaKubernetes);
        assertTrue(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.broker_container.replicas"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertTrue(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.broker_container.disk_size"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertTrue(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.broker_container.resources"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());

        assertTrue(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.zookeeper_container.replicas"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertTrue(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.zookeeper_container.disk_size"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertTrue(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.zookeeper_container.resources"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
    }


    @Test
    public void testKafkaKubernetes_ShouldNotReturnValidationErrorIfContainerReplicasAndDiskSizeIsSet() throws ValidationException {
        var kafkaKubernetes = KafkaKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test kafka-kubernetes").build())
                        .build())
                .setSpec(KafkaKubernetesSpec.newBuilder()
                                .setZookeeperContainer(KafkaKubernetesSpecZookeeperContainerSpec.newBuilder()
                                        .setDiskSize("8Gi")
                                        .setReplicas(1)
                                        .setResources(ContainerResources.newBuilder().build())
                                        .build())
                                .setBrokerContainer(KafkaKubernetesSpecBrokerContainerSpec.newBuilder()
                                        .setDiskSize("8Gi")
                                        .setReplicas(1)
                                        .setResources(ContainerResources.newBuilder().build())
                                        .build())
                                .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaKubernetes);
        assertFalse(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.broker_container.replicas"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertFalse(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.broker_container.disk_size"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertFalse(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.broker_container.resources"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());

        assertFalse(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.zookeeper_container.replicas"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertFalse(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.zookeeper_container.disk_size"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertFalse(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.zookeeper_container.resources"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
    }



}

