package buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.helmrelease;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubernetes.helmrelease.model.HelmRelease;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertTrue;

public final class HelmReleaseTest {
    @Test
    public void testHelmRelease_ShouldNotThroughProtoValidationException() {
        var helmrelease = HelmRelease.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(helmrelease));
    }

    @Test
    public void testHelmRelease_ShouldThroughErrorIfMetadataIsNotInitialized() throws ValidationException {
        var helmrelease = HelmRelease.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(helmrelease);
        var optionalViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(optionalViolation.isPresent());
    }
}

