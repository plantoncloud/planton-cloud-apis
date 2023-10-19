package buf.gen.cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.deploy.endpoint.custom.rpc.CustomEndpoint;
import build.buf.gen.cloud.planton.apis.v1.commons.resource.Metadata;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class CustomEndpointTest {

    @Test
    public void testCustomEndpoint_ShouldNotThroughProtoValidationException() {
        var input = CustomEndpoint.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input));
    }

}

