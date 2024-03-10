version ?= local
.PHONY: build
build:
	rm -rf internal/generated zzjava/src/main zzgo
	mkdir -p internal/generated zzjava/src/main/java zzgo
	buf generate --include-imports
	cp -R internal/generated/java/. zzjava/src/main/java/
	cp -R internal/generated/go/github.com/plantoncloud/planton-cloud-apis/zzgo/. zzgo/
	pushd zzjava;rm -rf build;./gradlew build;popd
.PHONY: release
release: build
	git tag ${version}
	git push origin ${version}
	buf push --tag ${version}
	pushd zzjava;rm -rf build;./gradlew publish -Prevision=${version};popd

.PHONY: release-local
release-local:
	pushd zzjava;rm -rf build;./gradlew publishToMavenLocal -Prevision=${version};popd

.PHONY: update
update:
	buf mod update
