package buf.gen.cloud.planton.apis.resourcemanager.v1.organization;

import build.buf.gen.cloud.planton.apis.resourcemanager.v1.organization.model.Organization;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class OrganizationTest {

    @Test
    public void testOrganization_ShouldNotThroughProtoValidationException() {
        var input1 = Organization.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

