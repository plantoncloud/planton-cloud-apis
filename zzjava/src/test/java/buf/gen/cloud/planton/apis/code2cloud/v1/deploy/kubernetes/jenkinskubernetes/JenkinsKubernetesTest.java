package buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.jenkinskubernetes;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.jenkinskubernetes.model.JenkinsKubernetes;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class JenkinsKubernetesTest {

    @Test
    public void testJenkinsKubernetes_ShouldNotThroughProtoValidationException() {
        var input = JenkinsKubernetes.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input));
    }

    @Test
    public void testJenkinsKubernetes_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var jenkinsKubernetes  = JenkinsKubernetes.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(jenkinsKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testJenkinsKubernetes_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var jenkinsKubernetes  = JenkinsKubernetes.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(jenkinsKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testJenkinsKubernetes_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var jenkinsKubernetes  = JenkinsKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test jenkins-kubernetes ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(jenkinsKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testJenkinsKubernetes_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var jenkinsKubernetes  = JenkinsKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test jenkins-kubernetes ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(jenkinsKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testJenkinsKubernetes_ShouldNotReturnValidationErrorIfNameLengthIsLessThan100() throws ValidationException {
        var jenkinsKubernetes  = JenkinsKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test jenkins-kubernetes ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(jenkinsKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    

}

