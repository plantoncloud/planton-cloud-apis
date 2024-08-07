package buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.rediskubernetes.model.RedisKubernetes;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.gen.cloud.planton.apis.commons.kubernetes.model.IngressSpec;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class RedisKubernetesTest {

    @Test
    public void testRedisKubernetes_ShouldNotThroughProtoValidationException() throws ValidationException {
        var input = RedisKubernetes.newBuilder().build();
        Validator validator = new Validator();
        validator.validate(input);
        assertDoesNotThrow(() -> validator.validate(input));
    }

    @Test
    public void testRedisKubernetesSpecKubernetesSpecIngressSpec_ShouldReturnTwoViolations() throws ValidationException {
        var input = IngressSpec.newBuilder()
                .setIsEnabled(true)
                .build();
        Validator validator = new Validator();
        var validationResult = validator.validate(input);
        assertEquals(1, validationResult.getViolations().size());
    }

    @Test
    public void testRedisKubernetesSpecKubernetesSpecIngressSpec_ShouldReturnNoViolations() throws ValidationException {
        var input = IngressSpec.newBuilder()
                .setIsEnabled(false)
                .build();
        Validator validator = new Validator();
        var validationResult = validator.validate(input);
        assertEquals(0, validationResult.getViolations().size());
    }

    @Test
    public void testRedisKubernetesSpecKubernetesSpecIngressSpec_ShouldReturnNoViolationsIfIsEnabledIsNotSet() throws ValidationException {
        var input = IngressSpec.newBuilder().build();
        Validator validator = new Validator();
        var validationResult = validator.validate(input);
        assertEquals(0, validationResult.getViolations().size());
    }

    @Test
    public void testRedisKubernetes_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var redisKubernetes = RedisKubernetes.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(redisKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testRedisKubernetes_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var redisKubernetes = RedisKubernetes.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(redisKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testRedisKubernetes_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var redisKubernetes = RedisKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test redis-kubernetes").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(redisKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testRedisKubernetes_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var redisKubernetes = RedisKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test redis-kubernetes").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(redisKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testRedisKubernetes_ShouldReturnValidationErrorIfNameLengthIsGreaterThan12() throws ValidationException {
        var redisKubernetes = RedisKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation, this is test name to check length validation, this is test name to check length validation")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test redis-kubernetes").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(redisKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testRedisKubernetes_ShouldNotReturnValidationErrorIfNameLengthIsLessThan12() throws ValidationException {
        var redisKubernetes = RedisKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test redis-kubernetes").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(redisKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

}

