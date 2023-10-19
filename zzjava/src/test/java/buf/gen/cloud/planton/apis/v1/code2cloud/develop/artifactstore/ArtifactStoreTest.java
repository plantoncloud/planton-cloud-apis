package buf.gen.cloud.planton.apis.v1.code2cloud.develop.artifactstore;

import build.buf.gen.cloud.planton.apis.v1.code2cloud.develop.artifactstore.rpc.ArtifactStore;
import build.buf.protovalidate.Validator;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public final class ArtifactStoreTest {

    @Test
    public void testArtifactStore_ShouldNotThroughProtoValidationException() {
        var input1 = ArtifactStore.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

}

