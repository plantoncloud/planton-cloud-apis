package buf.gen.cloud.planton.apis.code2cloud.v1.rpc;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.customendpoint.model.CustomEndpoint;
import build.buf.gen.cloud.planton.apis.commons.resource.model.Metadata;
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

