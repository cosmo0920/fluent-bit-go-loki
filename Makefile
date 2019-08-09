ifeq ($(OS),Windows_NT)
    DLLEXT := .dll
    TEST_OPTS := ./... -v
else
    DLLEXT := .so
    TEST_OPTS := -cover -race -coverprofile=coverage.txt -covermode=atomic
endif

VERSION := 0.2.0
GO_FLAGS := -ldflags "-X main.Version=${VERSION}"

all: test
	go build $(GO_FLAGS) -buildmode=c-shared -o out_loki$(DLLEXT) .

fast:
	go build out_loki.go

test:
	go test $(TEST_OPTS)

dep:
	dep ensure

clean:
	rm -rf *.so *.h
