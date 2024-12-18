FROM golang:1.23-alpine AS builder

RUN apk update && apk add postgresql-client

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o cloud_storage ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/cloud_storage .

EXPOSE 8080

CMD ["./cloud_storage"]