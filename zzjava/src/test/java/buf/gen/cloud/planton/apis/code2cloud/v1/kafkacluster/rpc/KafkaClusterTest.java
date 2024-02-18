package buf.gen.cloud.planton.apis.code2cloud.v1.kafkacluster.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.kafkacluster.model.KafkaCluster;
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

