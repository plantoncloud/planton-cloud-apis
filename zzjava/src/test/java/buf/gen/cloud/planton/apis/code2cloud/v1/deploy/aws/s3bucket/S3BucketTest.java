package buf.gen.cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.aws.s3bucket.model.S3Bucket;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class S3BucketTest {

    @Test
    public void testS3Bucket_ShouldNotThroughProtoValidationException() {
        var input1 = S3Bucket.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

    @Test
    public void testS3Bucket_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var storageBucket = S3Bucket.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(storageBucket);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testS3Bucket_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var storageBucket = S3Bucket.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(storageBucket);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testS3Bucket_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var storageBucket = S3Bucket.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test s3-bucket").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(storageBucket);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testS3Bucket_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var storageBucket = S3Bucket.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test s3-bucket").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(storageBucket);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 3 and 63 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testS3Bucket_ShouldReturnValidationErrorIfNameLengthIsGreaterThan63() throws ValidationException {
        var storageBucket = S3Bucket.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation. adding additional text to make it greater than 63")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test s3-bucket").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(storageBucket);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 3 and 63 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testS3Bucket_ShouldNotReturnValidationErrorIfNameLengthIsLessThan63() throws ValidationException {
        var storageBucket = S3Bucket.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test s3-bucket").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(storageBucket);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 3 and 63 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

}

