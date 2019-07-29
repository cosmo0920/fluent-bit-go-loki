FROM golang:1.12.7-stretch AS build-env
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOPATH=/go
ADD . /go/src/github.com/cosmo0920/fluent-bit-go-loki
WORKDIR /go/src/github.com/cosmo0920/fluent-bit-go-loki
RUN go build -buildmode=c-shared -o out_loki.so .

FROM fluent/fluent-bit:1.2
MAINTAINER Hiroshi Hatake <cosmo0920.wp[at]gmail.com>
COPY --from=build-env /go/src/github.com/cosmo0920/fluent-bit-go-loki/out_loki.so /usr/lib/x86_64-linux-gnu/
COPY docker/fluent-bit-loki.conf \
     /fluent-bit/etc/
EXPOSE 2020

# Entry point
CMD ["/fluent-bit/bin/fluent-bit", "-c", "/fluent-bit/etc/fluent-bit-loki.conf", "-e", "/usr/lib/x86_64-linux-gnu/out_loki.so"]
