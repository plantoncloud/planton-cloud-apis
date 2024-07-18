package buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.postgreskubernetes;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.postgreskubernetes.model.PostgresKubernetes;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class PostgresKubernetesTest {

    @Test
    public void testPostgresKubernetes_ShouldNotThroughProtoValidationException() {
        var input = PostgresKubernetes.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input));
    }

    @Test
    public void testPostgresKubernetes_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var postgresKubernetes = PostgresKubernetes.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(postgresKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testPostgresKubernetes_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var postgresKubernetes = PostgresKubernetes.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(postgresKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testPostgresKubernetes_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var postgresKubernetes = PostgresKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test postgres kubernetes").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(postgresKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testPostgresKubernetes_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var postgresKubernetes = PostgresKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test postgres kubernetes").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(postgresKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testPostgresKubernetes_ShouldReturnValidationErrorIfNameLengthIsGreaterThan12() throws ValidationException {
        var postgresKubernetes = PostgresKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation, this is test name to check length validation, this is test name to check length validation")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test postgres kubernetes").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(postgresKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testPostgresKubernetes_ShouldNotReturnValidationErrorIfNameLengthIsLessThan12() throws ValidationException {
        var postgresKubernetes = PostgresKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test postgres kubernetes").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(postgresKubernetes);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }
}

