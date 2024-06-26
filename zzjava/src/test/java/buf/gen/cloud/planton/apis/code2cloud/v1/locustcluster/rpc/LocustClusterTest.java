package buf.gen.cloud.planton.apis.code2cloud.v1.locustcluster.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.locustcluster.model.LocustCluster;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class LocustClusterTest {

    @Test
    public void testLocustCluster_ShouldNotThroughProtoValidationException() {
        var input1 = LocustCluster.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

    @Test
    public void testLocustCluster_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var locustCluster  = LocustCluster.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(locustCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testLocustCluster_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var locustCluster  = LocustCluster.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(locustCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testLocustCluster_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var locustCluster  = LocustCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test locust cluster ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(locustCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testLocustCluster_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var locustCluster  = LocustCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test locust cluster ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(locustCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testLocustCluster_ShouldNotReturnValidationErrorIfNameLengthIsLessThan100() throws ValidationException {
        var locustCluster  = LocustCluster.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test locust cluster ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(locustCluster );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    

}

