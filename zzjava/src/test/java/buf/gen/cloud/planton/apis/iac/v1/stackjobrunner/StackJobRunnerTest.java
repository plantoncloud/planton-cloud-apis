package buf.gen.cloud.planton.apis.iac.v1.stackjobrunner;

import build.buf.gen.cloud.planton.apis.iac.v1.stackjobrunner.model.StackJobRunner;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class StackJobRunnerTest {

    @Test
    public void test_ShouldNotThroughProtoValidationException() {
        var input = StackJobRunner.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input));
    }
}

