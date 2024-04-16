package buf.gen.cloud.planton.apis.code2cloud.v1.argocdinstance;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.argocdinstance.model.ArgocdInstance;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class ArgocdInstanceTest {
    @Test
    public void testArgocdInstance_ShouldNotThroughProtoValidationException() {
        var argocdInstance = ArgocdInstance.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(argocdInstance));
    }

    @Test
    public void testArgocdInstance_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var argocdInstance  = ArgocdInstance.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(argocdInstance );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("Name is mandatory")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testArgocdInstance_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var argocdInstance  = ArgocdInstance.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(argocdInstance );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testArgocdInstance_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var argocdInstance  = ArgocdInstance.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test argocd Instance ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(argocdInstance );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testArgocdInstance_ShouldReturnValidationErrorIfNameIsInvalid() throws ValidationException {
        var argocdInstance  = ArgocdInstance.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("$5asdasd")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test argocd Instance ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(argocdInstance);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Invalid Name - `$5asdasd`.Only lowercase letters, numbers, and hyphens are allowed")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }
}

