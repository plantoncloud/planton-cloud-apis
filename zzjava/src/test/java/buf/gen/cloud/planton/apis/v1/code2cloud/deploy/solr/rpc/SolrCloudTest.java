package buf.gen.cloud.planton.apis.v1.code2cloud.deploy.solr.rpc;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.deploy.solr.rpc.SolrCloud;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class SolrCloudTest {

    @Test
    public void testSolrCloud_ShouldNotThroughProtoValidationException() {
        var input1 = SolrCloud.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

