package buf.gen.cloud.planton.apis.v1.code2cloud.deploy.redis;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.deploy.redis.RedisCluster;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class RedisClusterTest {

    @Test
    public void testRedisCluster_ShouldNotThroughProtoValidationException() {
        var input1 = RedisCluster.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

