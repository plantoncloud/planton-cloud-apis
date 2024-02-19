package buf.gen.cloud.planton.apis.code2cloud.v1.rediscluster.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.rediscluster.model.RedisCluster;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.rediscluster.model.RedisClusterSpecKubernetesSpecIngressSpec;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class RedisClusterTest {

    @Test
    public void testRedisCluster_ShouldNotThroughProtoValidationException() throws ValidationException {
        var input1 = RedisCluster.newBuilder().build();
        Validator validator = new Validator();
        validator.validate(input1);
        assertDoesNotThrow(() -> validator.validate(input1));
    }

    @Test
    public void testRedisClusterSpecKubernetesSpecIngressSpec_ShouldReturnTwoViolations() throws ValidationException {
        var input1 = RedisClusterSpecKubernetesSpecIngressSpec.newBuilder()
                .setIsEnabled(true)
                .build();
        Validator validator = new Validator();
        var validationResult = validator.validate(input1);
        assertEquals(2, validationResult.getViolations().size());
    }

    @Test
    public void testRedisClusterSpecKubernetesSpecIngressSpec_ShouldReturnNoViolations() throws ValidationException {
        var input1 = RedisClusterSpecKubernetesSpecIngressSpec.newBuilder()
                .setIsEnabled(false)
                .build();
        Validator validator = new Validator();
        var validationResult = validator.validate(input1);
        assertEquals(0, validationResult.getViolations().size());
    }

    @Test
    public void testRedisClusterSpecKubernetesSpecIngressSpec_ShouldReturnNoViolationsIfIsEnabledIsNotSet() throws ValidationException {
        var input1 = RedisClusterSpecKubernetesSpecIngressSpec.newBuilder().build();
        Validator validator = new Validator();
        var validationResult = validator.validate(input1);
        assertEquals(0, validationResult.getViolations().size());
    }

    @Test
    public void testRedisCluster_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var redisCluster = RedisCluster.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(redisCluster);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("Name is mandatory")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testRedisCluster_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var redisCluster = RedisCluster.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(redisCluster);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testRedisCluster_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var redisCluster = RedisCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test redis cluster").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(redisCluster);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testRedisCluster_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var redisCluster = RedisCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test redis cluster").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(redisCluster);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 12 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testRedisCluster_ShouldReturnValidationErrorIfNameLengthIsGreaterThan12() throws ValidationException {
        var redisCluster = RedisCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test redis cluster").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(redisCluster);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 12 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testRedisCluster_ShouldNotReturnValidationErrorIfNameLengthIsLessThan12() throws ValidationException {
        var redisCluster = RedisCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test redis cluster").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(redisCluster);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 12 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

}

