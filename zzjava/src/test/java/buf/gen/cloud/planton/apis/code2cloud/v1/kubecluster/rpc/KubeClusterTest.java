package buf.gen.cloud.planton.apis.code2cloud.v1.kubecluster.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.kubecluster.model.KubeCluster;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class KubeClusterTest {

    @Test
    public void testKubeCluster_ShouldNotThroughProtoValidationException() {
        var input1 = KubeCluster.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}
