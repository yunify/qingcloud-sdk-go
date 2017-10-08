SHELL := /bin/bash

.PHONY: all check vet lint update generate build unit test release clean

PREFIX=qingcloud-sdk-go
VERSION=$(shell cat version.go | grep "Version\ =" | sed -e s/^.*\ //g | sed -e s/\"//g)
DIRS_TO_CHECK=$(shell ls -d */ | grep -vE "vendor|test")
PKGS_TO_CHECK=$(shell go list ./... | grep -vE "vendor|test")
PKGS_TO_RELEASE=$(shell go list ./... | grep -vE "/vendor/|/test")
FILES_TO_RELEASE=$(shell find . -name "*.go" | grep -vE "/vendor/|/test|.*_test.go")
FILES_TO_RELEASE_WITH_VENDOR=$(shell find . -name "*.go" | grep -vE "/test|.*_test.go")
LINT_IGNORE_DOC="service\/.*\.go:.+(comment on exported|should have comment or be unexported)"
LINT_IGNORE_CONFLICT="service\/.*\.go:.+(type name will be used as)"

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  all               to check, build, test and release this SDK"
	@echo "  check             to vet and lint the SDK"
	@echo "  generate          to generate service code"
	@echo "  build             to build the SDK"
	@echo "  unit              to run all sort of unit tests except runtime"
	@echo "  unit-test         to run unit test"
	@echo "  unit-benchmark    to run unit test with benchmark"
	@echo "  unit-coverage     to run unit test with coverage"
	@echo "  unit-race         to run unit test with race"
	@echo "  test              to run service test"
	@echo "  release           to build and release current version"
	@echo "  clean             to clean the coverage files"

all: check build unit release

check: vet lint

vet:
	@echo "go tool vet, skipping vendor packages"
	@go tool vet -all ${DIRS_TO_CHECK}
	@echo "ok"

lint:
	@echo "golint, skipping vendor packages"
	@lint=$$(for pkg in ${PKGS_TO_CHECK}; do golint $${pkg}; done); \
	 lint=$$(echo "$${lint}" | grep -vE -e ${LINT_IGNORE_DOC} -e ${LINT_IGNORE_CONFLICT}); \
	 if [[ -n $${lint} ]]; then echo "$${lint}"; exit 1; fi
	@echo "ok"

generate:
	@if [[ ! -f "$$(which snips)" ]]; then \
		echo "ERROR: Command \"snips\" not found."; \
	fi
	snips -f="./specs/qingcloud/2013-08-30/swagger/api_v2.0.json" -t="./template" -o="./service"
	@find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;
	@echo "ok"

build:
	@echo "build the SDK"
	go build ${PKGS_TO_RELEASE}
	@echo "ok"

unit: unit-test unit-benchmark unit-coverage unit-race

unit-test:
	@echo "run unit test"
	go test -v ${PKGS_TO_CHECK}
	@echo "ok"

unit-benchmark:
	@echo "run unit test with benchmark"
	go test -v -bench=. ${PKGS_TO_CHECK}
	@echo "ok"

unit-coverage:
	@echo "run unit test with coverage"
	for pkg in ${PKGS_TO_CHECK}; do \
		output="coverage$${pkg#github.com/yunify/qingcloud-sdk-go}"; \
		mkdir -p $${output}; \
		go test -v -cover -coverprofile="$${output}/profile.out" $${pkg}; \
		if [[ -e "$${output}/profile.out" ]]; then \
			go tool cover -html="$${output}/profile.out" -o "$${output}/profile.html"; \
		fi; \
	done
	@echo "ok"

unit-race:
	@echo "run unit test with race"
	go test -v -race -cpu=1,2,4 ${PKGS_TO_CHECK}
	@echo "ok"

test:
	pushd "./test"; go test ; popd
	@echo "ok"

release:
	@echo "pack the source code"
	mkdir -p "release"
	@zip -FS "release/${PREFIX}-source-v${VERSION}.zip" ${FILES_TO_RELEASE}
	@echo "ok"

clean:
	rm -rf $${PWD}/coverage
	@echo "ok"
