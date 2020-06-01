FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o agent-bin cmd/agent/main.go


FROM alpine

WORKDIR /app

COPY --from=builder /build/agent-bin agent
