package buf.gen.cloud.planton.apis.code2cloud.v1.environment;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.environment.model.Environment;
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

