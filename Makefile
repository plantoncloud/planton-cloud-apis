version ?= local
.PHONY: build
build:	
	buf generate --include-imports
	cp -R generated/go/github.com/plantoncloud/planton-cloud-apis/zzgo/. zzgo/
.PHONY: deploy
deploy:
	buf push --tag ${version}
	pushd zzjava;rm -rf build;./gradlew publish -Prevision=${version};popd
.PHONY: deploy-local
deploy-local:
	pushd zzjava;rm -rf build;./gradlew publishToMavenLocal -Prevision=${version};popd

.PHONY: update
update:
	buf mod update

.PHONY: release
release:
	buf push --tag ${version}
