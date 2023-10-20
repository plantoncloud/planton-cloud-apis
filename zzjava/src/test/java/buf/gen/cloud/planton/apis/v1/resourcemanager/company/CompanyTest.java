package buf.gen.cloud.planton.apis.v1.resourcemanager.company;

import build.buf.gen.cloud.planton.apis.v1.resourcemanager.company.Company;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class CompanyTest {

    @Test
    public void testCompany_ShouldNotThroughProtoValidationException() {
        var input1 = Company.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

