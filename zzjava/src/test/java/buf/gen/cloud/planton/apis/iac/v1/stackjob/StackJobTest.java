package buf.gen.cloud.planton.apis.iac.v1.stackjob;

import build.buf.gen.cloud.planton.apis.iac.v1.stackjob.model.StackJob;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class StackJobTest {

    @Test
    public void test_ShouldNotThroughProtoValidationException() {
        var input = StackJob.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input));
    }
}

