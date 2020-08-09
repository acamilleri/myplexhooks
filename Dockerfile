FROM golang:1.14-buster AS builder

ARG version

COPY ./ /go/myplexhooks
WORKDIR /go/myplexhooks
RUN go build -ldflags "-X main.version=$version" cmd/myplexhooks.go

FROM busybox:glibc

WORKDIR /app
COPY --from=builder /go/myplexhooks/myplexhooks /app

ENTRYPOINT ["/app/myplexhooks"]