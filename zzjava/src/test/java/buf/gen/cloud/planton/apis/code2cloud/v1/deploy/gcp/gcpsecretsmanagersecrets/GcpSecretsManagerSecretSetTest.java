package buf.gen.cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretsmanagersecrets;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpsecretsmanagersecretset.model.GcpSecretsManagerSecretSet;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class GcpSecretsManagerSecretSetTest {

    @Test
    public void testGcpSecretsManagerSecretSet_ShouldNotThroughProtoValidationException() {
        var input1 = GcpSecretsManagerSecretSet.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

    @Test
    public void testGcpSecretsManagerSecretSet_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var cloudAccount = GcpSecretsManagerSecretSet.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(cloudAccount);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testGcpSecretsManagerSecretSet_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var cloudAccount = GcpSecretsManagerSecretSet.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(cloudAccount);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testGcpSecretsManagerSecretSet_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var cloudAccount = GcpSecretsManagerSecretSet.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gcp-secrets-manager-secret-set").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(cloudAccount);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testGcpSecretsManagerSecretSet_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var cloudAccount = GcpSecretsManagerSecretSet.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gcp-secrets-manager-secret-set").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(cloudAccount);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testGcpSecretsManagerSecretSet_ShouldReturnValidationErrorIfNameLengthIsGreaterThan12() throws ValidationException {
        var cloudAccount = GcpSecretsManagerSecretSet.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation, this is test name to check length validation, this is test name to check length validation")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gcp-secrets-manager-secret-set").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(cloudAccount);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testGcpSecretsManagerSecretSet_ShouldNotReturnValidationErrorIfNameLengthIsLessThan12() throws ValidationException {
        var cloudAccount = GcpSecretsManagerSecretSet.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gcp-secrets-manager-secret-set").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(cloudAccount);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }
}

