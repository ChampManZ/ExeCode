# syntax=docker/dockerfile:1

FROM golang:1.19.1 AS builder
WORKDIR /execode/
# Execode server
COPY ./cmd/server ./cmd/server
COPY ./internal ./internal
COPY ./docs ./docs
# Go module files
COPY ./go.* ./
# Dependencies and build
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /execode/app ./
CMD ["./app"]
