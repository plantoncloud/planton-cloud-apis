package buf.gen.cloud.planton.apis.code2cloud.v1.microserviceinstance.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.microserviceinstance.model.MicroserviceInstance;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class MicroserviceInstanceTest {

    @Test
    public void testMicroserviceInstance_ShouldNotThroughProtoValidationException() {
        var input1 = MicroserviceInstance.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

