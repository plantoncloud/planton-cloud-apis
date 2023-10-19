package buf.gen.cloud.planton.apis.v1.code2cloud.develop.sourcecode.project;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.develop.sourcecode.project.rpc.CodeProject;
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

