package buf.gen.cloud.planton.apis.connect.v1.gitlabcredential;

import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.gen.cloud.planton.apis.connect.v1.gitlabcredential.model.GitlabCredential;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class GitlabCredentialTest {
    @Test
    public void testGitlabCredential_ShouldNotThroughProtoValidationException() {
        var gitlabCredential = GitlabCredential.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(gitlabCredential));
    }

    @Test
    public void testGitlabCredential_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var gitlabCredential = GitlabCredential.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(gitlabCredential);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testGitlabCredential_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var gitlabCredential = GitlabCredential.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(gitlabCredential);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testGitlabCredential_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var gitlabCredential = GitlabCredential.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gitlab-credential").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(gitlabCredential);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testGitlabCredential_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var gitlabCredential = GitlabCredential.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gitlab-credential").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(gitlabCredential);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testGitlabCredential_ShouldReturnValidationErrorIfNameLengthIsGreaterThan100() throws ValidationException {
        var gitlabCredential = GitlabCredential.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("this is test name to check length validation, this is test name to check length validation, this is test name to check length validation")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gitlab-credential").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(gitlabCredential);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testGitlabCredential_ShouldNotReturnValidationErrorIfNameLengthIsLessThan100() throws ValidationException {
        var gitlabCredential = GitlabCredential.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test gitlab-credential").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(gitlabCredential);
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }
}
