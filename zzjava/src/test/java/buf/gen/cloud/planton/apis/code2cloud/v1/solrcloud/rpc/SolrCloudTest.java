package buf.gen.cloud.planton.apis.code2cloud.v1.solrcloud.rpc;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.solrcloud.model.SolrCloud;
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
