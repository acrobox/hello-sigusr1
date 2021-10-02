FROM golang:alpine AS builder
ENV CGO_ENABLED=0
RUN adduser -u 1000 -S user
WORKDIR /build
COPY go.mod /build/go.mod
COPY *.go /build/
RUN go build -a -installsuffix docker -ldflags='-w -s' -o /build/bin/hello-sigusr1 /build

FROM alpine:latest
EXPOSE 8080
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /build/bin/hello-sigusr1 /usr/local/bin/hello-sigusr1
USER user
WORKDIR /home/user
CMD ["/usr/local/bin/hello-sigusr1"]
