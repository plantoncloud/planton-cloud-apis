package buf.gen.cloud.planton.apis.v1.code2cloud.environment;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.environment.rpc.Environment;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class EnvironmentTest {

    @Test
    public void testEnvironment_ShouldNotThroughProtoValidationException() {
        var input1 = Environment.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

