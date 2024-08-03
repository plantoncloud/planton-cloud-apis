package buf.gen.cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.gcp.gcpartifactregistry.model.GcpArtifactRegistry;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class GcpArtifactRegistryTest {
    @Test
    public void testGcpArtifactRegistry_ShouldNotThroughProtoValidationException() {
        var artifactStore = GcpArtifactRegistry.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(artifactStore));
    }

    @Test
    public void testGcpArtifactRegistry_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var artifactStore = GcpArtifactRegistry.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(artifactStore);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testGcpArtifactRegistry_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var artifactStore = GcpArtifactRegistry.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(artifactStore);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testGcpArtifactRegistry_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var artifactStore = GcpArtifactRegistry.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gcp-artifact-registry").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(artifactStore);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(optionalViolation.isPresent());
    }

    @Test
    public void testGcpArtifactRegistry_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var artifactStore = GcpArtifactRegistry.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gcp-artifact-registry").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(artifactStore);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testGcpArtifactRegistry_ShouldReturnValidationErrorIfNameLengthIsGreaterThan100() throws ValidationException {
        var artifactStore = GcpArtifactRegistry.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation, this is test name to check length validation, this is test name to check length validation")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gcp-artifact-registry").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(artifactStore);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testGcpArtifactRegistry_ShouldNotReturnValidationErrorIfNameLengthIsLessThan100() throws ValidationException {
        var artifactStore = GcpArtifactRegistry.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gcp-artifact-registry").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(artifactStore);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertFalse(optionalViolation.isPresent());
    }
}
