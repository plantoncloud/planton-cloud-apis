package buf.gen.cloud.planton.apis.code2cloud.v1.mongodbcluster.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.mongodbcluster.model.MongodbCluster;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.mongodbcluster.model.MongodbClusterSpec;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.mongodbcluster.model.MongodbClusterSpecKubernetesSpec;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.mongodbcluster.model.MongodbClusterSpecKubernetesSpecMongodbContainerSpec;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class MongodbClusterTest {

    @Test
    public void testMongodbCluster_ShouldNotThroughProtoValidationException() {
        var input1 = MongodbCluster.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

    @Test
    public void testMongodbCluster_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var mongodbCluster  = MongodbCluster.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbCluster_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var mongodbCluster  = MongodbCluster.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbCluster_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var mongodbCluster  = MongodbCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb cluster ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbCluster_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var mongodbCluster  = MongodbCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb cluster ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbCluster_ShouldNotReturnValidationErrorIfNameLengthIsLessThan100() throws ValidationException {
        var mongodbCluster  = MongodbCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb cluster ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbCluster_ShouldNotReturnValidationErrorIfDiskSizeIsEmpty() throws ValidationException {
        var mongodbCluster  = MongodbCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb cluster ").build())
                        .build())
                .setSpec(MongodbClusterSpec.newBuilder()
                        .setKubernetes(MongodbClusterSpecKubernetesSpec.newBuilder()
                                .setMongodbContainer(MongodbClusterSpecKubernetesSpecMongodbContainerSpec.newBuilder()
                                        .setIsPersistenceEnabled(false)
                                        .build())
                                .build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("spec.kubernetes.container.disk_size.mandatory"))
                .filter(violation -> violation.getMessage().equals("disk size is mandatory to enable persistence")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbCluster_ShouldReturnValidationErrorIfDiskSizeIsEmpty() throws ValidationException {
        var mongodbCluster  = MongodbCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb cluster ").build())
                        .build())
                .setSpec(MongodbClusterSpec.newBuilder()
                        .setKubernetes(MongodbClusterSpecKubernetesSpec.newBuilder()
                                .setMongodbContainer(MongodbClusterSpecKubernetesSpecMongodbContainerSpec.newBuilder()
                                        .setIsPersistenceEnabled(true)
                                        .build())
                                .build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("spec.kubernetes.container.disk_size.mandatory"))
                .filter(violation -> violation.getMessage().equals("disk size is mandatory to enable persistence")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }


    @Test
    public void testMongodbCluster_ShouldNotReturnValidationErrorIfDiskSizeIsNotEmpty() throws ValidationException {
        var mongodbCluster  = MongodbCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb cluster ").build())
                        .build())
                .setSpec(MongodbClusterSpec.newBuilder()
                        .setKubernetes(MongodbClusterSpecKubernetesSpec.newBuilder()
                                .setMongodbContainer(MongodbClusterSpecKubernetesSpecMongodbContainerSpec.newBuilder()
                                        .setIsPersistenceEnabled(true)
                                        .setDiskSize("8Gi")
                                        .build())
                                .build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("spec.kubernetes.container.disk_size.is_valid"))
                .filter(violation -> violation.getMessage().equals("disk size value is invalid")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbCluster_ShouldReturnValidationErrorIfDiskSizeIsNotEmptyAndInvalid() throws ValidationException {
        var mongodbCluster  = MongodbCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb cluster ").build())
                        .build())
                .setSpec(MongodbClusterSpec.newBuilder()
                        .setKubernetes(MongodbClusterSpecKubernetesSpec.newBuilder()
                                .setMongodbContainer(MongodbClusterSpecKubernetesSpecMongodbContainerSpec.newBuilder()
                                        .setIsPersistenceEnabled(true)
                                        .setDiskSize("Gi8")
                                        .build())
                                .build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("spec.kubernetes.container.disk_size.is_valid"))
                .filter(violation -> violation.getMessage().equals("disk size value is invalid")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

}

