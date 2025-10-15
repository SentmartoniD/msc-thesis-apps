FROM golang:latest AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth-service ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/auth-service .

EXPOSE 8000

CMD ["./auth-service"]
