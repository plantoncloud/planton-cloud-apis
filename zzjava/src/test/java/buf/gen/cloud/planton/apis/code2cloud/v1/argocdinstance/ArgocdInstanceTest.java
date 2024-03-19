package buf.gen.cloud.planton.apis.code2cloud.v1.argocdinstance;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.argocdinstance.model.ArgocdInstance;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class ArgocdInstanceTest {
    @Test
    public void testArgocdInstance_ShouldNotThroughProtoValidationException() {
        var argocdInstance = ArgocdInstance.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(argocdInstance));
    }
}

