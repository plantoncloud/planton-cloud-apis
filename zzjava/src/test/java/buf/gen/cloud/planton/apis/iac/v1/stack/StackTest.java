package buf.gen.cloud.planton.apis.iac.v1.stack;

import build.buf.gen.cloud.planton.apis.iac.v1.stack.model.Stack;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class StackTest {

    @Test
    public void test_ShouldNotThroughProtoValidationException() {
        var input1 = Stack.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

