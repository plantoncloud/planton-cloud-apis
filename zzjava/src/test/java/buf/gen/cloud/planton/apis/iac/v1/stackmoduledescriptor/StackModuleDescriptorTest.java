package buf.gen.cloud.planton.apis.iac.v1.stackmoduledescriptor;

import build.buf.gen.cloud.planton.apis.iac.v1.stackmoduledescriptor.model.StackModuleDescriptor;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class StackModuleDescriptorTest {

    @Test
    public void test_ShouldNotThroughProtoValidationException() {
        var input = StackModuleDescriptor.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input));
    }
}
