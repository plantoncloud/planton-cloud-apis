package buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.grafanakubernetes;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.grafanakubernetes.model.GrafanaKubernetes;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.grafanakubernetes.model.GrafanaKubernetesSpec;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertTrue;

public final class GrafanaKubernetesTest {
    @Test
    public void testGrafanaKubernetes_ShouldNotThroughProtoValidationException() {
        var grafanakubernetes = GrafanaKubernetes.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(grafanakubernetes));
    }

    @Test
    public void testGrafanaKubernetes_ShouldThroughErrorIfMetadataIsNotInitialized() throws ValidationException {
        var grafanakubernetes = GrafanaKubernetes.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(grafanakubernetes);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testGrafanaKubernetes_ShouldThroughErrorIfKubernetesIsInitializedButContainerIsNotInitialized() throws ValidationException {
        var grafanakubernetes = GrafanaKubernetes.newBuilder()
                .setSpec(GrafanaKubernetesSpec.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(grafanakubernetes);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.container"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }
}

