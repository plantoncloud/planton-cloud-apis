package buf.gen.cloud.planton.apis.code2cloud.v1.dnszone.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.dnszone.model.DnsZone;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class DnsZoneTest {

    @Test
    public void testDnsZone_ShouldNotThroughProtoValidationException() {
        var input1 = DnsZone.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

    @Test
    public void testDnsZone_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var dnsZone = DnsZone.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(dnsZone);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("Name is mandatory")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testDnsZone_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var dnsZone = DnsZone.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(dnsZone);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testDnsZone_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var dnsZone = DnsZone.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test dns zone ").build())
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
    public void testDnsZone_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var dnsZone = DnsZone.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test dns zone ").build())
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
    public void testDnsZone_ShouldReturnValidationErrorIfNameLengthIsGreaterThan65() throws ValidationException {
        var dnsZone = DnsZone.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation. adding additional text to make it greater than 65")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test dns zone ").build())
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
    public void testDnsZone_ShouldNotReturnValidationErrorIfNameLengthIsLessThan65() throws ValidationException {
        var dnsZone = DnsZone.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test dns zone ").build())
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

