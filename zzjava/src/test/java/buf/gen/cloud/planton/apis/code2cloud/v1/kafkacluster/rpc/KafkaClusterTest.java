package buf.gen.cloud.planton.apis.code2cloud.v1.kafkacluster.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.kafkacluster.model.*;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubecluster.model.ContainerResources;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class KafkaClusterTest {

    @Test
    public void testKafkaCluster_ShouldNotThroughProtoValidationException() {
        var input1 = KafkaCluster.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

    @Test
    public void testKafkaCluster_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var kafkaCluster = KafkaCluster.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaCluster);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testKafkaCluster_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var kafkaCluster = KafkaCluster.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaCluster);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testKafkaCluster_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var kafkaCluster = KafkaCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test kafka cluster").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaCluster);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testKafkaCluster_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var kafkaCluster = KafkaCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test kafka cluster").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaCluster);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testKafkaCluster_ShouldReturnValidationErrorIfNameLengthIsGreaterThan12() throws ValidationException {
        var kafkaCluster = KafkaCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation, this is test name to check length validation,this is test name to check length validation")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test kafka cluster").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaCluster);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testKafkaCluster_ShouldNotReturnValidationErrorIfNameLengthIsLessThan12() throws ValidationException {
        var kafkaCluster = KafkaCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test kafka cluster").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaCluster);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testKafkaCluster_ShouldReturnValidationErrorIfContainerReplicasAndDiskSizeIsNotSet() throws ValidationException {
        var kafkaCluster = KafkaCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test kafka cluster").build())
                        .build())
                .setSpec(KafkaClusterSpec.newBuilder()
                        .setKubernetes(KafkaClusterSpecKubernetesSpec.newBuilder()
                                .setZookeeperContainer(KafkaClusterSpecKubernetesSpecZookeeperContainerSpec.newBuilder().build())
                                .setKafkaBrokerContainer(KafkaClusterSpecKubernetesSpecKafkaBrokerContainerSpec.newBuilder().build())
                                .build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaCluster);
        assertTrue(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.kafka_broker_container.replicas"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertTrue(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.kafka_broker_container.disk_size"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertTrue(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.kafka_broker_container.resources"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());

        assertTrue(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.zookeeper_container.replicas"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertTrue(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.zookeeper_container.disk_size"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertTrue(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.zookeeper_container.resources"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
    }


    @Test
    public void testKafkaCluster_ShouldNotReturnValidationErrorIfContainerReplicasAndDiskSizeIsSet() throws ValidationException {
        var kafkaCluster = KafkaCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test kafka cluster").build())
                        .build())
                .setSpec(KafkaClusterSpec.newBuilder()
                        .setKubernetes(KafkaClusterSpecKubernetesSpec.newBuilder()
                                .setZookeeperContainer(KafkaClusterSpecKubernetesSpecZookeeperContainerSpec.newBuilder()
                                        .setDiskSize("8Gi")
                                        .setReplicas(1)
                                        .setResources(ContainerResources.newBuilder().build())
                                        .build())
                                .setKafkaBrokerContainer(KafkaClusterSpecKubernetesSpecKafkaBrokerContainerSpec.newBuilder()
                                        .setDiskSize("8Gi")
                                        .setReplicas(1)
                                        .setResources(ContainerResources.newBuilder().build())
                                        .build())
                                .build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(kafkaCluster);
        assertFalse(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.kafka_broker_container.replicas"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertFalse(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.kafka_broker_container.disk_size"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertFalse(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.kafka_broker_container.resources"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());

        assertFalse(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.zookeeper_container.replicas"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertFalse(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.zookeeper_container.disk_size"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
        assertFalse(result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.zookeeper_container.resources"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst().isPresent());
    }



}

