package buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.elasticsearchkubernetes;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.elasticsearchkubernetes.model.ElasticsearchKubernetes;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.elasticsearchkubernetes.model.ElasticsearchKubernetesSpec;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertTrue;

public final class ElastisearchKubernetesTest {
    @Test
    public void testElasticsearchKubernetes_ShouldNotThroughProtoValidationException() {
        var elasticsearchkubernetes = ElasticsearchKubernetes.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(elasticsearchkubernetes));
    }

    @Test
    public void testElasticsearchKubernetes_ShouldThroughErrorIfMetadataIsNotInitialized() throws ValidationException {
        var elasticsearchKubernetes = ElasticsearchKubernetes.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(elasticsearchKubernetes);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }

    @Test
    public void testElasticsearchKubernetes_ShouldThroughErrorIfKubernetesIsInitializedButContainerIsNotInitialized() throws ValidationException {
        var elasticsearchKubernetes = ElasticsearchKubernetes.newBuilder()
                .setSpec(ElasticsearchKubernetesSpec.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(elasticsearchKubernetes);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("spec.container"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }
}

