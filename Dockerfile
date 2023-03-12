FROM go:1.20 as builder

COPY . /go
WORKDIR /go
RUN go build -v ./...

FROM ubuntu:22.10
COPY --from=builder /go /go
WORKDIR /go
