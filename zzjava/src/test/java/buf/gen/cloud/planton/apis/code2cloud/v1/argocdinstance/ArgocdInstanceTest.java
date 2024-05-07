package buf.gen.cloud.planton.apis.code2cloud.v1.argocdinstance;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.argocdinstance.model.ArgocdInstance;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.argocdinstance.model.ArgocdInstanceSpec;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.argocdinstance.model.ArgocdInstanceSpecKubernetesSpec;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class ArgocdInstanceTest {
    @Test
    public void testArgocdInstance_ShouldNotThroughProtoValidationException() {
        var argocdInstance = ArgocdInstance.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(argocdInstance));
    }

    @Test
    public void testArgocdInstance_ShouldThroughErrorIfMetadataIsNotInitialized() throws ValidationException {
        var argocdInstance = ArgocdInstance.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(argocdInstance);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testArgocdInstance_ShouldThroughErrorIfKubernetesIsInitializedButContainerIsNotInitialized() throws ValidationException {
        var argocdInstance = ArgocdInstance.newBuilder()
                .setSpec(ArgocdInstanceSpec.newBuilder()
                        .setKubernetes(ArgocdInstanceSpecKubernetesSpec.newBuilder().build()).build()).build();
        Validator validator = new Validator();
        var result = validator.validate(argocdInstance);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.kubernetes.argocd_container"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }
}

