ifeq ($(OS),Windows_NT)
    DLLEXT := .dll
    TEST_OPTS := ./... -v
else
    DLLEXT := .so
    TEST_OPTS := -cover -race -coverprofile=coverage.txt -covermode=atomic
endif

all: test
	go build -buildmode=c-shared -o out_loki$(DLLEXT) .

fast:
	go build out_loki.go

test:
	go test $(TEST_OPTS)

dep:
	dep ensure

clean:
	rm -rf *.so *.h
