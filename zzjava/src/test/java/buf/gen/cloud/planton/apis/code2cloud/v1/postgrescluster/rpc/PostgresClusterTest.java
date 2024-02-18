package buf.gen.cloud.planton.apis.code2cloud.v1.postgrescluster.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.postgrescluster.model.PostgresCluster;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class PostgresClusterTest {

    @Test
    public void testPostgresCluster_ShouldNotThroughProtoValidationException() {
        var input1 = PostgresCluster.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

