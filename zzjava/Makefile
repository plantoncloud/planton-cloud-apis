version?=local
.PHONY: build
build:
	./gradlew build

.PHONY: test
test:
	./gradlew check

.PHONY: release
release:
	./gradlew publish -Prevision=${version}
