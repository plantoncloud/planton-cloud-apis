package buf.gen.cloud.planton.apis.code2cloud.v1.microserviceinstance.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.microserviceinstance.model.MicroserviceInstance;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class MicroserviceInstanceTest {

    @Test
    public void testMicroserviceInstance_ShouldNotThroughProtoValidationException() {
        var input1 = MicroserviceInstance.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

    @Test
    public void testMicroserviceInstance_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var microserviceInstance  = MicroserviceInstance.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(microserviceInstance );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("Name is mandatory")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testMicroserviceInstance_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var microserviceInstance  = MicroserviceInstance.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(microserviceInstance );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testMicroserviceInstance_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var microserviceInstance  = MicroserviceInstance.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test microservice instance ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(microserviceInstance );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testMicroserviceInstance_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var microserviceInstance  = MicroserviceInstance.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test microservice instance ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(microserviceInstance );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 12 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testMicroserviceInstance_ShouldReturnValidationErrorIfNameLengthIsGreaterThan12() throws ValidationException {
        var microserviceInstance  = MicroserviceInstance.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test microservice instance ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(microserviceInstance );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 12 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testMicroserviceInstance_ShouldNotReturnValidationErrorIfNameLengthIsLessThan12() throws ValidationException {
        var microserviceInstance  = MicroserviceInstance.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test microservice instance ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(microserviceInstance );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 12 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

}

