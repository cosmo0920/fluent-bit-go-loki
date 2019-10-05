ifeq ($(OS),Windows_NT)
    DLLEXT := .dll
    TEST_OPTS := ./... -v
else
    DLLEXT := .so
    TEST_OPTS := -cover -race -coverprofile=coverage.txt -covermode=atomic
endif

VERSION := 0.3.1

# Version info for binaries
GIT_REVISION := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

VPREFIX := github.com/cosmo0920/fluent-bit-go-loki/vendor/github.com/prometheus/common/version
GO_FLAGS := -ldflags "-X $(VPREFIX).Branch=$(GIT_BRANCH) -X $(VPREFIX).Version=$(VERSION) -X $(VPREFIX).Revision=$(GIT_REVISION)" -tags netgo

all: test build
build:
	go build $(GO_FLAGS) -buildmode=c-shared -o out_loki$(DLLEXT) .

fast:
	go build out_loki.go

test:
	go test $(TEST_OPTS)

dep:
	dep ensure

clean:
	rm -rf *.so *.h
