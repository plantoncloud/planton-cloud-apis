package buf.gen.cloud.planton.apis.code2cloud.v1.configmap;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.configmap.model.ConfigMap;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class ConfigMapTest {
    @Test
    public void testConfigMap_ShouldNotThroughProtoValidationException() {
        var awsCredential = ConfigMap.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(awsCredential));
    }

    @Test
    public void testConfigMap_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var awsCredential = ConfigMap.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(awsCredential);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testConfigMap_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var awsCredential = ConfigMap.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(awsCredential);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testConfigMap_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var awsCredential = ConfigMap.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test config-map").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(awsCredential);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testConfigMap_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var awsCredential = ConfigMap.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test config-map").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(awsCredential);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testConfigMap_ShouldReturnValidationErrorIfNameLengthIsGreaterThan100() throws ValidationException {
        var awsCredential = ConfigMap.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation, this is test name to check length validation, this is test name to check length validation")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test config-map").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(awsCredential);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testConfigMap_ShouldNotReturnValidationErrorIfNameLengthIsLessThan100() throws ValidationException {
        var awsCredential = ConfigMap.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test config-map").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(awsCredential);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }
}
