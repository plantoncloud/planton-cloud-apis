package buf.gen.cloud.planton.apis.connect.v1.pulumibackendcredential;

import build.buf.gen.cloud.planton.apis.connect.v1.pulumibackendcredential.model.PulumiBackendCredential;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class PulumiBackendCredentialTest {

    @Test
    public void test_ShouldNotThroughProtoValidationException() {
        var input = PulumiBackendCredential.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input));
    }
}

