package buf.gen.cloud.planton.apis.v1.resourcemanager.product;

import build.buf.gen.cloud.planton.apis.v1.resourcemanager.product.rpc.Product;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class ProductTest {

    @Test
    public void testProduct_ShouldNotThroughProtoValidationException() {
        var input1 = Product.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

