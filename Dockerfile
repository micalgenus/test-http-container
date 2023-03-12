FROM golang:1.20 as builder

COPY . /home
WORKDIR /home
RUN go build -v ./...

FROM ubuntu:22.10
COPY --from=builder /home/test-http-container /usr/local/bin/test-http-container
WORKDIR /

ENTRYPOINT /usr/local/bin/test-http-container
