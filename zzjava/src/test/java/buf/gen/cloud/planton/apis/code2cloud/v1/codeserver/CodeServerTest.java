package buf.gen.cloud.planton.apis.code2cloud.v1.codeserver;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.codeserver.model.CodeServer;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class CodeServerTest {

    @Test
    public void testCodeServer_ShouldNotThroughProtoValidationException() {
        var input1 = CodeServer.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

    @Test
    public void testCodeServer_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var codeServer = CodeServer.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(codeServer);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("Name is mandatory")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testCodeServer_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var codeServer = CodeServer.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(codeServer);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testCodeServer_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var codeServer = CodeServer.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test code server ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(codeServer);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testCodeServer_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var codeServer = CodeServer.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test code server ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(codeServer);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 12 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testCodeServer_ShouldReturnValidationErrorIfNameLengthIsGreaterThan12() throws ValidationException {
        var codeServer = CodeServer.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test code server ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(codeServer);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 12 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testCodeServer_ShouldNotReturnValidationErrorIfNameLengthIsLessThan12() throws ValidationException {
        var codeServer = CodeServer.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test code server ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(codeServer);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 12 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

}

