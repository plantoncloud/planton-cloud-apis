package buf.gen.cloud.planton.apis.v1.code2cloud.deploy.endpoint.standard;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.deploy.endpoint.standard.StandardEndpoint;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class StandardEndpointTest {

    @Test
    public void testStandardEndpoint_ShouldNotThroughProtoValidationException() {
        var input1 = StandardEndpoint.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

