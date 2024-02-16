package buf.gen.cloud.planton.apis.v1.code2cloud.develop.sourcecode.server;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.develop.sourcecode.server.model.CodeServer;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class CodeServerTest {

    @Test
    public void testCodeServer_ShouldNotThroughProtoValidationException() {
        var input1 = CodeServer.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

