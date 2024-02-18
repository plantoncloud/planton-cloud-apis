package buf.gen.cloud.planton.apis.code2cloud.v1.codeproject;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.codeproject.model.CodeProject;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class CodeProjectTest {

    @Test
    public void testCodeProject_ShouldNotThroughProtoValidationException() {
        var input1 = CodeProject.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

