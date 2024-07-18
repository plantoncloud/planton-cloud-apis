package buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.prometheuskubernetes;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.prometheuskubernetes.model.PrometheusKubernetes;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.prometheuskubernetes.model.PrometheusKubernetesSpec;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertTrue;

public final class PrometheusKubernetesTest {
    @Test
    public void testPrometheusKubernetes_ShouldNotThroughProtoValidationException() {
        var prometheuskubernetes = PrometheusKubernetes.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(prometheuskubernetes));
    }

    @Test
    public void testPrometheusKubernetes_ShouldThroughErrorIfMetadataIsNotInitialized() throws ValidationException {
        var prometheuskubernetes = PrometheusKubernetes.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(prometheuskubernetes);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testPrometheusKubernetes_ShouldThroughErrorIfKubernetesIsInitializedButContainerIsNotInitialized() throws ValidationException {
        var prometheuskubernetes = PrometheusKubernetes.newBuilder()
                .setSpec(PrometheusKubernetesSpec.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(prometheuskubernetes);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.container"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }
}

