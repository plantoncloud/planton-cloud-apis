package buf.gen.cloud.planton.apis.v1.code2cloud.deploy.storagebucket;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.deploy.storagebucket.StorageBucket;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class StorageBucketTest {

    @Test
    public void testStorageBucket_ShouldNotThroughProtoValidationException() {
        var input1 = StorageBucket.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

