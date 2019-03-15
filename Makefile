all:
	go build -buildmode=c-shared -o out_loki.so .

fast:
	go build out_loki.go

clean:
	rm -rf *.so *.h
