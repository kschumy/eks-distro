BASE_DIRECTORY=$(shell git rev-parse --show-toplevel)
RELEASE_BRANCH?=$(shell cat $(BASE_DIRECTORY)/release/DEFAULT_RELEASE_BRANCH)
RELEASE_ENVIRONMENT?=development
RELEASE?=$(shell cat $(BASE_DIRECTORY)/release/$(RELEASE_BRANCH)/$(RELEASE_ENVIRONMENT)/RELEASE)


.PHONY: generate-release-number
generate-release-number:
	go vet releasecmd
	go run releasecmd/main.go \
		--release-branch=$(RELEASE_BRANCH) \
		--release-environment=$(RELEASE_ENVIRONMENT)