package buf.gen.cloud.planton.apis.code2cloud.v1.artifactstore;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.artifactstore.model.ArtifactStore;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class ArtifactStoreTest {
    @Test
    public void testArtifactStore_ShouldNotThroughProtoValidationException() {
        var artifactStore = ArtifactStore.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(artifactStore));
    }

    @Test
    public void testArtifactStore_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var artifactStore = ArtifactStore.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(artifactStore);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("Name is mandatory")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testArtifactStore_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var artifactStore = ArtifactStore.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(artifactStore);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testArtifactStore_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var artifactStore = ArtifactStore.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test artifact store").build())
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
    public void testArtifactStore_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var artifactStore = ArtifactStore.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test artifact store").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(artifactStore);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 12 characters long")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testArtifactStore_ShouldReturnValidationErrorIfNameLengthIsGreaterThan12() throws ValidationException {
        var artifactStore = ArtifactStore.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test artifact store").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(artifactStore);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 12 characters long")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testArtifactStore_ShouldNotReturnValidationErrorIfNameLengthIsLessThan12() throws ValidationException {
        var artifactStore = ArtifactStore.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test artifact store").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(artifactStore);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 12 characters long")).findFirst();
        assertFalse(optionalViolation.isPresent());
    }
    
    

}

