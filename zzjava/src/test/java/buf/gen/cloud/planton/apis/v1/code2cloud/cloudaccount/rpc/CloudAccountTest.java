package buf.gen.cloud.planton.apis.v1.code2cloud.cloudaccount;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.cloudaccount.CloudAccount;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class CloudAccountTest {

    @Test
    public void testCloudAccount_ShouldNotThroughProtoValidationException() {
        var input1 = CloudAccount.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

