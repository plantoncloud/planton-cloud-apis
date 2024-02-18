package buf.gen.cloud.planton.apis.code2cloud.v1.codeserver;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.codeserver.model.CodeServer;
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

