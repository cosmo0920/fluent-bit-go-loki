ifeq ($(OS),Windows_NT)
    DLLEXT := .dll
else
    DLLEXT := .so
endif

all: test
	go build -buildmode=c-shared -o out_loki$(DLLEXT) .

fast:
	go build out_loki.go

test:
	go test -cover -race -coverprofile=coverage.txt -covermode=atomic

dep:
	dep ensure

clean:
	rm -rf *.so *.h
