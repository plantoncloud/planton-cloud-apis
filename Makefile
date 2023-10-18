version ?= local
.PHONY: build
build:
	rm -rf generated zzjava/src/main
	mkdir -p generated zzjava/src/main/java
	buf generate --include-imports
	cp -R generated/java/. zzjava/src/main/java/
	cp -R generated/go/github.com/plantoncloud/planton-cloud-apis/zzgo/. zzgo/
.PHONY: release
release:
	buf push --tag ${version}
	pushd zzjava;rm -rf build;./gradlew publish -Prevision=${version};popd
.PHONY: release-local
release-local:
	pushd zzjava;rm -rf build;./gradlew publishToMavenLocal -Prevision=${version};popd

.PHONY: update
update:
	buf mod update
