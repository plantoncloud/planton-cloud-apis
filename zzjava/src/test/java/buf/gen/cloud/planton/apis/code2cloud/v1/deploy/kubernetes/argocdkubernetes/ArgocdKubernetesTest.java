package buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.argocdkubernetes;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.argocdkubernetes.model.ArgocdKubernetes;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.argocdkubernetes.model.ArgocdKubernetesSpec;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class ArgocdKubernetesTest {
    @Test
    public void testArgocdKubernetes_ShouldNotThroughProtoValidationException() {
        var argocdkubernetes = ArgocdKubernetes.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(argocdkubernetes));
    }

    @Test
    public void testArgocdKubernetes_ShouldThroughErrorIfMetadataIsNotInitialized() throws ValidationException {
        var argocdkubernetes = ArgocdKubernetes.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(argocdkubernetes);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testArgocdKubernetes_ShouldThroughErrorIfKubernetesIsInitializedButContainerIsNotInitialized() throws ValidationException {
        var argocdkubernetes = ArgocdKubernetes.newBuilder()
                .setSpec(ArgocdKubernetesSpec.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(argocdkubernetes);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.container"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }
}

