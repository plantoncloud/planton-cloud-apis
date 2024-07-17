package buf.gen.cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpclouddnszone;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.gcp.gcpclouddnszone.model.GcpCloudDnsZone;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class GcpCloudDnsZoneTest {

    @Test
    public void testGcpCloudDnsZone_ShouldNotThroughProtoValidationException() {
        var input1 = GcpCloudDnsZone.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

    @Test
    public void testGcpCloudDnsZone_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var dnsZone = GcpCloudDnsZone.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(dnsZone);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testGcpCloudDnsZone_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var dnsZone = GcpCloudDnsZone.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(dnsZone);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testGcpCloudDnsZone_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var dnsZone = GcpCloudDnsZone.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gcp-cloud-dns-zone ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(dnsZone);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testGcpCloudDnsZone_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var dnsZone = GcpCloudDnsZone.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gcp-cloud-dns-zone ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(dnsZone);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 65 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testGcpCloudDnsZone_ShouldReturnValidationErrorIfNameLengthIsGreaterThan65() throws ValidationException {
        var dnsZone = GcpCloudDnsZone.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation. adding additional text to make it greater than 65")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gcp-cloud-dns-zone ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(dnsZone);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 65 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testGcpCloudDnsZone_ShouldNotReturnValidationErrorIfNameLengthIsLessThan65() throws ValidationException {
        var dnsZone = GcpCloudDnsZone.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gcp-cloud-dns-zone ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(dnsZone);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 65 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

}

