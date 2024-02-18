package buf.gen.cloud.planton.apis.code2cloud.v1.rediscluster.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.rediscluster.model.RedisCluster;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.rediscluster.model.RedisClusterSpecKubernetesSpecIngressSpec;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertEquals;

public final class RedisClusterTest {

    @Test
    public void testRedisCluster_ShouldNotThroughProtoValidationException() throws ValidationException {
        var input1 = RedisCluster.newBuilder().build();
        Validator validator = new Validator();
        validator.validate(input1);
        assertDoesNotThrow(() -> validator.validate(input1));
    }

    @Test
    public void testRedisClusterSpecKubernetesSpecIngressSpec_ShouldReturnTwoViolations() throws ValidationException {
        var input1 = RedisClusterSpecKubernetesSpecIngressSpec.newBuilder()
                .setIsEnabled(true)
                .build();
        Validator validator = new Validator();
        var validationResult = validator.validate(input1);
        assertEquals(2, validationResult.getViolations().size());
    }

    @Test
    public void testRedisClusterSpecKubernetesSpecIngressSpec_ShouldReturnNoViolations() throws ValidationException {
        var input1 = RedisClusterSpecKubernetesSpecIngressSpec.newBuilder()
                .setIsEnabled(false)
                .build();
        Validator validator = new Validator();
        var validationResult = validator.validate(input1);
        assertEquals(0, validationResult.getViolations().size());
    }

    @Test
    public void testRedisClusterSpecKubernetesSpecIngressSpec_ShouldReturnNoViolationsIfIsEnabledIsNotSet() throws ValidationException {
        var input1 = RedisClusterSpecKubernetesSpecIngressSpec.newBuilder().build();
        Validator validator = new Validator();
        var validationResult = validator.validate(input1);
        assertEquals(0, validationResult.getViolations().size());
    }

}

