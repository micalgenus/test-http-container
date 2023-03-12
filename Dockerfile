FROM go:1.20 as builder

COPY . /go
WORKDIR /go
RUN go build -v ./...

FROM ubuntu:22.10
COPY --from=builder /go/test-http-container /usr/local/bin/test-http-container
WORKDIR /

ENTRYPOINT /usr/local/bin/test-http-container