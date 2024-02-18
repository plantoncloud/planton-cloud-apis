package buf.gen.cloud.planton.apis.resourcemanager.product;

import build.buf.gen.cloud.planton.apis.resourcemanager.v1.product.model.Product;
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

