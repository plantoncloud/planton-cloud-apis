package buf.gen.cloud.planton.apis.iac.v1.deploymentmodule;

import build.buf.gen.cloud.planton.apis.iac.v1.deploymentmodule.model.DeploymentModule;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class DeploymentModuleTest {

    @Test
    public void test_ShouldNotThroughProtoValidationException() {
        var input = DeploymentModule.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input));
    }
}
