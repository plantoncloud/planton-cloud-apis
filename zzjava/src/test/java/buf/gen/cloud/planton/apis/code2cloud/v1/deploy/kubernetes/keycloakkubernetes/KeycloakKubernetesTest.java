package buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.keycloakkubernetes;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.keycloakkubernetes.model.KeycloakKubernetes;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.keycloakkubernetes.model.KeycloakKubernetesSpec;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertTrue;

public final class KeycloakKubernetesTest {
    @Test
    public void testKeycloakKubernetes_ShouldNotThroughProtoValidationException() {
        var keycloakkubernetes = KeycloakKubernetes.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(keycloakkubernetes));
    }

    @Test
    public void testKeycloakKubernetes_ShouldThroughErrorIfMetadataIsNotInitialized() throws ValidationException {
        var keycloakkubernetes = KeycloakKubernetes.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(keycloakkubernetes);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testKeycloakKubernetes_ShouldThroughErrorIfKubernetesIsInitializedButContainerIsNotInitialized() throws ValidationException {
        var keycloakkubernetes = KeycloakKubernetes.newBuilder()
                .setSpec(KeycloakKubernetesSpec.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(keycloakkubernetes);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.container"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }
}

