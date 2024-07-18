package buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetes;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.gitlabkubernetes.model.GitlabKubernetesSpec;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertTrue;

public final class GitlabKubernetesTest {
    @Test
    public void testGitlabKubernetes_ShouldNotThroughProtoValidationException() {
        var gitlabkubernetes = GitlabKubernetes.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(gitlabkubernetes));
    }

    @Test
    public void testGitlabKubernetes_ShouldThroughErrorIfMetadataIsNotInitialized() throws ValidationException {
        var gitlabkubernetes = GitlabKubernetes.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(gitlabkubernetes);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testGitlabKubernetes_ShouldThroughErrorIfKubernetesIsInitializedButContainerIsNotInitialized() throws ValidationException {
        var gitlabkubernetes = GitlabKubernetes.newBuilder()
                .setSpec(GitlabKubernetesSpec.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(gitlabkubernetes);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.container"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }
}

