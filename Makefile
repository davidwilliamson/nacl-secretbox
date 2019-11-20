MAJOR?=0
MINOR?=1 

VERSION=$(MAJOR).$(MINOR)

APP_NAME = $(shell basename $$PWD)

# directories in this repo for artifacts
BIN_DIR = "bin"
TEST_RESULTS_DIR = "test-results"

# for bumping the version number in 'make next-vers'
$(eval MINOR_NEXT=$(shell echo $$(( ${MINOR} + 1)) ) )
VERSION_NEXT = $(MAJOR).$(MINOR_NEXT)

# Find go files and packages in this repo for go test
# GOPACKAGES=$(shell go list ${PROJECT_DIR}/... | grep -v -e /vendor/ )
GOPACKAGES=$(shell go list | grep -v -e /vendor/ )
GOFILES = $(shell find . -type f -name '*.go')
GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "*/vendor/*")

# Current git state of repo
GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

##########################
# Help screen. invoked by any of: 'make', 'make default', 'make help'
##########################
.PHONY: default
default:
	@echo "${APP_NAME} version ${VERSION}"
	@echo "Please specify a target. The choices are:"
	@echo "---------- Testing ---------"
	@echo "test       : Run go unit tests"
	@echo "testv      : Run go unit tests in verbose mode"
	@echo "test-cover : Run go unit tests and open browser with code coverage"
	@echo "test-static: Run go static analysis tools: go vet, golint"
	@echo "check-fmt  : Run go fmt and report. Does not alter files."
	@echo "fmt        : Run go fmt and modify files (ignores vendor/)"
	@echo "---------- Builds ----------"
	@echo "clean      : Remove files in ./${BIN_DIR}"
	@echo "build      : Build the app; binary is ./${BIN_DIR}/${APP_NAME}"
	@echo "run        : Build and run ./${BIN_DIR}/${APP_NAME}"
	@echo "---------- Release ---------"
	@echo "tag        : Git tag master branch with ${VERSION}"

.PHONY: help
help: default
	@echo ""

#################################
# Build targets for local machine
#################################
.PHONY: clean
clean:
	@echo "+ $@"
	@echo "rm -f ./${BIN_DIR}/*"
	@rm -f ./${BIN_DIR}/*

.PHONY: build
build: clean version-check
	@echo "+ $@"
	@go build -o ./${BIN_DIR}/${APP_NAME} -ldflags "-X main.buildVersion=${VERSION} -X main.buildDate=`date +%Y%m%d.%H%M%S`" .
	@ls -l ./${BIN_DIR}/${APP_NAME}
	@echo "to run: ./${BIN_DIR}/${APP_NAME}"

.PHONY: run
run: build
	@./${BIN_DIR}/${APP_NAME}

#################################
# test targets
#################################
.PHONY: test
test:
	@echo "+ $@"
	@go test -coverprofile=${TEST_RESULTS_DIR}/coverage.out .

.PHONY: testv
testv:
	@echo "+ $@"
	@go test -v -coverprofile=${TEST_RESULTS_DIR}/coverage.out .
.PHONY: test-cover
test-cover: test
	@echo "+ $@"
	@go tool cover -html=${TEST_RESULTS_DIR}/coverage.out

.PHONY: check-fmt
check-fmt:
	@echo "+ $@"
	@gofmt -d ${GOFILES_NOVENDOR}
	@gofmt -l ${GOFILES_NOVENDOR} | (! grep . -q) || (echo "Code differs from gofmt's style. Run 'make fmt'" && false)
	@echo "go fmt check OK"

# Runs gofmt -w on the project's source code, modifying any files that do not
# match its style.
.PHONY: fmt
fmt:
	@echo "+ $@"
	@gofmt -l -w ${GOFILES_NOVENDOR}
	@goimports -l -w ${GOFILES_NOVENDOR}

.PHONY: test-static
test-static: check-fmt
	@echo "+ $@"
	@echo "go vet ${GOPACKAGES}"
	@go vet ${GOPACKAGES}
	@echo "golint ${GOPACKAGES}"
	@golint ${GOPACKAGES}
	@echo "ineffassign"
	@ineffassign .

#################################
# release targets
#################################
.PHONY: tag
tag: version-check branch-check
	@echo "+ $@"
	@git fetch --all
	@echo "Tag commit ${GIT_COMMIT} as version ${VERSION}"
	@git tag release/${VERSION} ${GIT_COMMIT}
	@git tag -l -n
	@git push --tags origin

#################################
# Utilities
#################################
.PHONY: version-check
version-check:
	@echo "+ $@"
	if [ -z "${VERSION}" ]; then \
	  echo "VERSION is not set" ; \
	  false ; \
	else \
	  echo "VERSION is ${VERSION}"; \
	fi

.PHONY: branch-check
branch-check:
	@echo "+ $@"
	@echo "git branch is ${GIT_BRANCH}" ; \
	if [ "${GIT_BRANCH}" = 'master' ]; then \
	  echo "Verified on master branch" ; \
	else  \
	  echo "must be on master branch" ; \
	  false ; \
	fi
