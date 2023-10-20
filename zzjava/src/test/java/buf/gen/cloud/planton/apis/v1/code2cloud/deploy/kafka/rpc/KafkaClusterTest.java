package buf.gen.cloud.planton.apis.v1.code2cloud.deploy.kafka;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.deploy.kafka.KafkaCluster;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class KafkaClusterTest {

    @Test
    public void testKafkaCluster_ShouldNotThroughProtoValidationException() {
        var input1 = KafkaCluster.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

