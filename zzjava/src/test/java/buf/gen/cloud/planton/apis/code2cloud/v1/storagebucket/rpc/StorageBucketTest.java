package buf.gen.cloud.planton.apis.code2cloud.v1.storagebucket.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.storagebucket.model.StorageBucket;
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

