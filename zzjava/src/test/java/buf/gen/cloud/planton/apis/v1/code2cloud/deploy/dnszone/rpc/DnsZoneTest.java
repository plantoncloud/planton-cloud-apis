package buf.gen.cloud.planton.apis.v1.code2cloud.deploy.dnszone;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.deploy.dnszone.DnsZone;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class DnsZoneTest {

    @Test
    public void testDnsZone_ShouldNotThroughProtoValidationException() {
        var input1 = DnsZone.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

