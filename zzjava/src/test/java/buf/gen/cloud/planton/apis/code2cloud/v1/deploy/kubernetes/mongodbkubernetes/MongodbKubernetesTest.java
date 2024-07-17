package buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes;

import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.model.MongodbKubernetes;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.model.MongodbKubernetesSpec;
import build.buf.gen.cloud.planton.apis.code2cloud.v1.deploy.kubernetes.mongodbkubernetes.model.MongodbKubernetesSpecContainerSpec;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadata;
import build.buf.gen.cloud.planton.apis.commons.apiresource.model.ApiResourceMetadataVersion;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public final class MongodbKubernetesTest {

    @Test
    public void testMongodbKubernetes_ShouldNotThroughProtoValidationException() {
        var input1 = MongodbKubernetes.newBuilder().build();
        Validator validator = new Validator();
        assertDoesNotThrow(() -> validator.validate(input1));
    }

    @Test
    public void testMongodbKubernetes_ShouldReturnValidationErrorIfMetadataIsNotPassed() throws ValidationException {
        var mongodbKubernetes  = MongodbKubernetes.newBuilder().build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getFieldPath().equals("metadata"))
                .filter(violation -> violation.getMessage().equals("value is required")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbKubernetes_ShouldReturnValidationErrorIfVersionMessageIsNotPassed() throws ValidationException {
        var mongodbKubernetes  = MongodbKubernetes.newBuilder().setMetadata(ApiResourceMetadata.newBuilder().build()).build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbKubernetes_ShouldNotReturnValidationErrorIfVersionMessageIsPassed() throws ValidationException {
        var mongodbKubernetes  = MongodbKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb-kubernetes ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.version.message"))
                .filter(violation -> violation.getMessage().equals("Version message is mandatory and cannot be empty")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbKubernetes_ShouldReturnValidationErrorIfNameIsEmpty() throws ValidationException {
        var mongodbKubernetes  = MongodbKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb-kubernetes ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbKubernetes_ShouldNotReturnValidationErrorIfNameLengthIsLessThan100() throws ValidationException {
        var mongodbKubernetes  = MongodbKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb-kubernetes ").build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("metadata.name"))
                .filter(violation -> violation.getMessage().equals("Name must be between 1 and 30 characters long")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbKubernetes_ShouldNotReturnValidationErrorIfDiskSizeIsEmpty() throws ValidationException {
        var mongodbKubernetes  = MongodbKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb-kubernetes ").build())
                        .build())
                .setSpec(MongodbKubernetesSpec.newBuilder()
                                .setContainer(MongodbKubernetesSpecContainerSpec.newBuilder()
                                        .setIsPersistenceEnabled(false)
                                        .build())
                                .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("spec.kubernetes.container.disk_size.mandatory"))
                .filter(violation -> violation.getMessage().equals("disk size is mandatory to enable persistence")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbKubernetes_ShouldReturnValidationErrorIfDiskSizeIsEmpty() throws ValidationException {
        var mongodbKubernetes  = MongodbKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb-kubernetes ").build())
                        .build())
                .setSpec(MongodbKubernetesSpec.newBuilder()
                        .setContainer(MongodbKubernetesSpecContainerSpec.newBuilder()
                                .setIsPersistenceEnabled(true)
                                .build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("spec.kubernetes.container.disk_size.mandatory"))
                .filter(violation -> violation.getMessage().equals("disk size is mandatory to enable persistence")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }


    @Test
    public void testMongodbKubernetes_ShouldNotReturnValidationErrorIfDiskSizeIsNotEmpty() throws ValidationException {
        var mongodbKubernetes  = MongodbKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb-kubernetes ").build())
                        .build())
                .setSpec(MongodbKubernetesSpec.newBuilder()
                        .setContainer(MongodbKubernetesSpecContainerSpec.newBuilder()
                                .setIsPersistenceEnabled(true)
                                .setDiskSize("8Gi")
                                .build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("spec.kubernetes.container.disk_size.is_valid"))
                .filter(violation -> violation.getMessage().equals("disk size value is invalid")).findFirst();
        assertFalse(versionMessageViolation.isPresent());
    }

    @Test
    public void testMongodbKubernetes_ShouldReturnValidationErrorIfDiskSizeIsNotEmptyAndInvalid() throws ValidationException {
        var mongodbKubernetes  = MongodbKubernetes.newBuilder()
                .setMetadata(ApiResourceMetadata.newBuilder()
                        .setName("test")
                        .setVersion(ApiResourceMetadataVersion.newBuilder().setMessage(" test mongodb-kubernetes ").build())
                        .build())
                .setSpec(MongodbKubernetesSpec.newBuilder()
                        .setContainer(MongodbKubernetesSpecContainerSpec.newBuilder()
                                .setIsPersistenceEnabled(true)
                                .setDiskSize("Gi8")
                                .build())
                        .build())
                .build();
        Validator validator = new Validator();
        var result = validator.validate(mongodbKubernetes );
        var versionMessageViolation = result.getViolations().stream()
                .filter(violation -> violation.getConstraintId().equals("spec.kubernetes.container.disk_size.is_valid"))
                .filter(violation -> violation.getMessage().equals("disk size value is invalid")).findFirst();
        assertTrue(versionMessageViolation.isPresent());
    }

}

